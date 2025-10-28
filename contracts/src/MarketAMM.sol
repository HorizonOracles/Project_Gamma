// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "./OutcomeToken.sol";
import "./FeeSplitter.sol";
import "./HorizonPerks.sol";

/**
 * @title MarketAMM
 * @notice Binary Constant Product Market Maker for prediction markets
 * @dev Implements x * y = k AMM for Yes/No outcome tokens
 *
 *      Trading formula: (x + Δx) * (y - Δy) = k (before fees)
 *      Fee: Applied to input amount before swap
 *      LP Tokens: Minted proportionally to liquidity added
 */
contract MarketAMM is ERC20, ReentrancyGuard, Pausable, IERC1155Receiver {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error MarketClosed();
    error MarketResolved();
    error InsufficientLiquidity();
    error SlippageExceeded();
    error InvalidAmount();
    error InvalidReserves();
    error InsufficientLPTokens();
    error MinimumLiquidityRequired();
    error InvalidState();

    // ============ Events ============

    event LiquidityAdded(address indexed provider, uint256 collateralAmount, uint256 lpTokens);
    event LiquidityRemoved(address indexed provider, uint256 lpTokens, uint256 collateralAmount);
    event Trade(
        address indexed trader,
        bool indexed buyYes,
        uint256 collateralIn,
        uint256 tokensOut,
        uint256 fee,
        uint256 price
    );

    // ============ Constants ============

    /// @notice Minimum liquidity locked forever (similar to Uniswap V2)
    uint256 public constant MINIMUM_LIQUIDITY = 1000;

    /// @notice Outcome ID for Yes tokens
    uint256 public constant OUTCOME_YES = 0;

    /// @notice Outcome ID for No tokens
    uint256 public constant OUTCOME_NO = 1;

    // ============ Immutable State ============

    /// @notice Market identifier
    uint256 public immutable marketId;

    /// @notice Collateral token (e.g., USDC)
    IERC20 public immutable collateralToken;

    /// @notice Outcome token contract
    OutcomeToken public immutable outcomeToken;

    /// @notice Fee splitter contract
    FeeSplitter public immutable feeSplitter;

    /// @notice Horizon perks contract (for fee discounts)
    HorizonPerks public immutable horizonPerks;

    /// @notice Market close timestamp
    uint256 public immutable closeTime;

    // ============ Mutable State ============

    /// @notice Reserve of Yes outcome tokens
    uint256 public reserveYes;

    /// @notice Reserve of No outcome tokens
    uint256 public reserveNo;

    /// @notice Total collateral backing the reserves
    uint256 public totalCollateral;

    // ============ Constructor ============

    /**
     * @notice Initializes a new market AMM
     * @param _marketId Market identifier
     * @param _collateralToken Collateral token address
     * @param _outcomeToken Outcome token contract
     * @param _feeSplitter Fee splitter contract
     * @param _horizonPerks Horizon perks contract
     * @param _closeTime Market close timestamp
     */
    constructor(
        uint256 _marketId,
        address _collateralToken,
        address _outcomeToken,
        address _feeSplitter,
        address _horizonPerks,
        uint256 _closeTime
    ) ERC20("Market LP Token", "MKT-LP") {
        marketId = _marketId;
        collateralToken = IERC20(_collateralToken);
        outcomeToken = OutcomeToken(_outcomeToken);
        feeSplitter = FeeSplitter(_feeSplitter);
        horizonPerks = HorizonPerks(_horizonPerks);
        closeTime = _closeTime;
    }

    // ============ Modifiers ============

    modifier beforeClose() {
        if (block.timestamp >= closeTime) revert MarketClosed();
        _;
    }

    modifier notResolved() {
        if (outcomeToken.isResolved(marketId)) revert MarketResolved();
        _;
    }

    // ============ Liquidity Functions ============

    /**
     * @notice Adds liquidity to the market
     * @param amount Amount of collateral to add
     * @return lpTokens Amount of LP tokens minted
     */
    function addLiquidity(uint256 amount) external nonReentrant whenNotPaused notResolved returns (uint256 lpTokens) {
        if (amount == 0) revert InvalidAmount();

        // Transfer collateral from user
        collateralToken.safeTransferFrom(msg.sender, address(this), amount);

        uint256 totalSupply = totalSupply();

        if (totalSupply == 0) {
            // Initial liquidity provision
            // Mint equal amounts of Yes and No tokens
            uint256 initialTokens = amount;

            // Mint outcome tokens
            outcomeToken.mintOutcome(marketId, OUTCOME_YES, address(this), initialTokens);
            outcomeToken.mintOutcome(marketId, OUTCOME_NO, address(this), initialTokens);

            // Set initial reserves
            reserveYes = initialTokens;
            reserveNo = initialTokens;
            totalCollateral = amount;

            // Mint LP tokens (lock minimum liquidity)
            lpTokens = amount - MINIMUM_LIQUIDITY;
            _mint(address(0xdead), MINIMUM_LIQUIDITY); // Lock minimum liquidity
            _mint(msg.sender, lpTokens);
        } else {
            // Subsequent liquidity provision
            // Mint proportional to existing liquidity
            lpTokens = (amount * totalSupply) / totalCollateral;

            // Calculate tokens to mint based on current reserves
            uint256 yesTokens = (amount * reserveYes) / totalCollateral;
            uint256 noTokens = (amount * reserveNo) / totalCollateral;

            // Mint outcome tokens
            outcomeToken.mintOutcome(marketId, OUTCOME_YES, address(this), yesTokens);
            outcomeToken.mintOutcome(marketId, OUTCOME_NO, address(this), noTokens);

            // Update reserves
            reserveYes += yesTokens;
            reserveNo += noTokens;
            totalCollateral += amount;

            // Mint LP tokens
            _mint(msg.sender, lpTokens);
        }

        emit LiquidityAdded(msg.sender, amount, lpTokens);
    }

    /**
     * @notice Removes liquidity from the market
     * @param lpTokens Amount of LP tokens to burn
     * @return collateralOut Amount of collateral returned
     */
    function removeLiquidity(uint256 lpTokens) external nonReentrant returns (uint256 collateralOut) {
        if (lpTokens == 0) revert InvalidAmount();

        uint256 totalSupply = totalSupply();
        if (lpTokens > balanceOf(msg.sender)) revert InsufficientLPTokens();

        // Calculate proportional collateral
        collateralOut = (lpTokens * totalCollateral) / totalSupply;

        // Calculate proportional tokens to burn
        uint256 yesBurn = (lpTokens * reserveYes) / totalSupply;
        uint256 noBurn = (lpTokens * reserveNo) / totalSupply;

        // Update reserves
        reserveYes -= yesBurn;
        reserveNo -= noBurn;
        totalCollateral -= collateralOut;

        // Burn LP tokens
        _burn(msg.sender, lpTokens);

        // Burn outcome tokens
        outcomeToken.burnOutcome(marketId, OUTCOME_YES, address(this), yesBurn);
        outcomeToken.burnOutcome(marketId, OUTCOME_NO, address(this), noBurn);

        // Return collateral
        collateralToken.safeTransfer(msg.sender, collateralOut);

        emit LiquidityRemoved(msg.sender, lpTokens, collateralOut);
    }

    // ============ Trading Functions ============

    /**
     * @notice Buys Yes outcome tokens
     * @param collateralIn Amount of collateral to spend
     * @param minTokensOut Minimum tokens to receive (slippage protection)
     * @return tokensOut Amount of Yes tokens received
     */
    function buyYes(uint256 collateralIn, uint256 minTokensOut)
        external
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        returns (uint256 tokensOut)
    {
        tokensOut = _buy(collateralIn, minTokensOut, true);
    }

    /**
     * @notice Buys No outcome tokens
     * @param collateralIn Amount of collateral to spend
     * @param minTokensOut Minimum tokens to receive (slippage protection)
     * @return tokensOut Amount of No tokens received
     */
    function buyNo(uint256 collateralIn, uint256 minTokensOut)
        external
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        returns (uint256 tokensOut)
    {
        tokensOut = _buy(collateralIn, minTokensOut, false);
    }

    /**
     * @notice Sells Yes outcome tokens
     * @param tokensIn Amount of Yes tokens to sell
     * @param minCollateralOut Minimum collateral to receive (slippage protection)
     * @return collateralOut Amount of collateral received
     */
    function sellYes(uint256 tokensIn, uint256 minCollateralOut)
        external
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        returns (uint256 collateralOut)
    {
        collateralOut = _sell(tokensIn, minCollateralOut, true);
    }

    /**
     * @notice Sells No outcome tokens
     * @param tokensIn Amount of No tokens to sell
     * @param minCollateralOut Minimum collateral to receive (slippage protection)
     * @return collateralOut Amount of collateral received
     */
    function sellNo(uint256 tokensIn, uint256 minCollateralOut)
        external
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        returns (uint256 collateralOut)
    {
        collateralOut = _sell(tokensIn, minCollateralOut, false);
    }

    // ============ Internal Trading Logic ============

    /**
     * @notice Internal buy logic
     * @param collateralIn Amount of collateral to spend
     * @param minTokensOut Minimum tokens to receive
     * @param isBuyYes True to buy Yes, false to buy No
     * @return tokensOut Amount of tokens received
     */
    function _buy(uint256 collateralIn, uint256 minTokensOut, bool isBuyYes) internal returns (uint256 tokensOut) {
        if (collateralIn == 0) revert InvalidAmount();
        if (reserveYes == 0 || reserveNo == 0) revert InsufficientLiquidity();

        // Get user's fee tier
        uint16 feeBps = horizonPerks.feeBpsFor(msg.sender);
        uint16 protocolBps = horizonPerks.protocolBpsFor(msg.sender);

        // Calculate fee
        uint256 fee = (collateralIn * feeBps) / 10000;
        uint256 collateralAfterFee = collateralIn - fee;

        // Transfer collateral from user
        collateralToken.safeTransferFrom(msg.sender, address(this), collateralIn);

        // Distribute fee
        collateralToken.forceApprove(address(feeSplitter), fee);
        feeSplitter.distribute(marketId, address(collateralToken), fee, protocolBps);

        // Calculate tokens out using CPMM formula
        uint256 reserveIn = isBuyYes ? reserveNo : reserveYes;
        uint256 reserveOut = isBuyYes ? reserveYes : reserveNo;

        // Mint new outcome token pairs with collateral
        outcomeToken.mintOutcome(marketId, OUTCOME_YES, address(this), collateralAfterFee);
        outcomeToken.mintOutcome(marketId, OUTCOME_NO, address(this), collateralAfterFee);

        // Update reserves to include new minted tokens
        uint256 newReserveIn = reserveIn + collateralAfterFee;
        uint256 k = reserveIn * reserveOut;

        // Calculate tokens out: k = newReserveIn * newReserveOut
        uint256 newReserveOut = k / newReserveIn;
        tokensOut = reserveOut - newReserveOut;

        // Slippage check
        if (tokensOut < minTokensOut) revert SlippageExceeded();

        // Update reserves
        if (isBuyYes) {
            reserveYes = newReserveOut;
            reserveNo = newReserveIn;
        } else {
            reserveNo = newReserveOut;
            reserveYes = newReserveIn;
        }

        totalCollateral += collateralAfterFee;

        // Transfer tokens to user
        uint256 outcomeId = isBuyYes ? OUTCOME_YES : OUTCOME_NO;
        outcomeToken.burnOutcome(marketId, outcomeId, address(this), tokensOut);
        outcomeToken.mintOutcome(marketId, outcomeId, msg.sender, tokensOut);

        // Calculate price
        uint256 price = (collateralIn * 1e18) / tokensOut;

        emit Trade(msg.sender, isBuyYes, collateralIn, tokensOut, fee, price);
    }

    /**
     * @notice Internal sell logic
     * @param tokensIn Amount of tokens to sell
     * @param minCollateralOut Minimum collateral to receive
     * @param isSellYes True to sell Yes, false to sell No
     * @return collateralOut Amount of collateral received
     */
    function _sell(uint256 tokensIn, uint256 minCollateralOut, bool isSellYes) internal returns (uint256 collateralOut) {
        if (tokensIn == 0) revert InvalidAmount();
        if (reserveYes == 0 || reserveNo == 0) revert InsufficientLiquidity();

        // Get user's fee tier
        uint16 feeBps = horizonPerks.feeBpsFor(msg.sender);
        uint16 protocolBps = horizonPerks.protocolBpsFor(msg.sender);

        // Transfer tokens from user
        uint256 outcomeId = isSellYes ? OUTCOME_YES : OUTCOME_NO;
        outcomeToken.burnOutcome(marketId, outcomeId, msg.sender, tokensIn);
        outcomeToken.mintOutcome(marketId, outcomeId, address(this), tokensIn);

        // Calculate collateral out using CPMM formula
        uint256 reserveIn = isSellYes ? reserveYes : reserveNo;
        uint256 reserveOut = isSellYes ? reserveNo : reserveYes;

        uint256 newReserveIn = reserveIn + tokensIn;
        uint256 k = reserveIn * reserveOut;
        uint256 newReserveOut = k / newReserveIn;
        uint256 tokensOutOpposite = reserveOut - newReserveOut;

        // Burn the token pair and get collateral
        outcomeToken.burnOutcome(marketId, OUTCOME_YES, address(this), tokensOutOpposite);
        outcomeToken.burnOutcome(marketId, OUTCOME_NO, address(this), tokensOutOpposite);

        collateralOut = tokensOutOpposite;

        // Apply fee
        uint256 fee = (collateralOut * feeBps) / 10000;
        collateralOut -= fee;

        // Slippage check
        if (collateralOut < minCollateralOut) revert SlippageExceeded();

        // Update reserves
        if (isSellYes) {
            reserveYes = newReserveIn - tokensOutOpposite;
            reserveNo = newReserveOut;
        } else {
            reserveNo = newReserveIn - tokensOutOpposite;
            reserveYes = newReserveOut;
        }

        totalCollateral -= tokensOutOpposite;

        // Distribute fee
        collateralToken.forceApprove(address(feeSplitter), fee);
        feeSplitter.distribute(marketId, address(collateralToken), fee, protocolBps);

        // Transfer collateral to user
        collateralToken.safeTransfer(msg.sender, collateralOut);

        // Calculate price
        uint256 price = (collateralOut * 1e18) / tokensIn;

        emit Trade(msg.sender, isSellYes, 0, tokensIn, fee, price);
    }

    // ============ View Functions ============

    /**
     * @notice Gets a quote for buying Yes tokens
     * @param collateralIn Amount of collateral to spend
     * @param user Address to check for fee tier (use trader's address for accurate quotes)
     * @return tokensOut Estimated tokens received
     * @return fee Fee amount
     */
    function getQuoteBuyYes(uint256 collateralIn, address user) external view returns (uint256 tokensOut, uint256 fee) {
        return _getQuoteBuy(collateralIn, true, user);
    }

    /**
     * @notice Gets a quote for buying No tokens
     * @param collateralIn Amount of collateral to spend
     * @param user Address to check for fee tier (use trader's address for accurate quotes)
     * @return tokensOut Estimated tokens received
     * @return fee Fee amount
     */
    function getQuoteBuyNo(uint256 collateralIn, address user) external view returns (uint256 tokensOut, uint256 fee) {
        return _getQuoteBuy(collateralIn, false, user);
    }

    /**
     * @notice Gets a quote for selling Yes tokens
     * @param tokensIn Amount of tokens to sell
     * @param user Address to check for fee tier (use trader's address for accurate quotes)
     * @return collateralOut Estimated collateral received
     * @return fee Fee amount
     */
    function getQuoteSellYes(uint256 tokensIn, address user) external view returns (uint256 collateralOut, uint256 fee) {
        return _getQuoteSell(tokensIn, true, user);
    }

    /**
     * @notice Gets a quote for selling No tokens
     * @param tokensIn Amount of tokens to sell
     * @param user Address to check for fee tier (use trader's address for accurate quotes)
     * @return collateralOut Estimated collateral received
     * @return fee Fee amount
     */
    function getQuoteSellNo(uint256 tokensIn, address user) external view returns (uint256 collateralOut, uint256 fee) {
        return _getQuoteSell(tokensIn, false, user);
    }

    /**
     * @notice Internal buy quote calculation
     */
    function _getQuoteBuy(uint256 collateralIn, bool isBuyYes, address user)
        internal
        view
        returns (uint256 tokensOut, uint256 fee)
    {
        if (collateralIn == 0 || reserveYes == 0 || reserveNo == 0) return (0, 0);

        uint16 feeBps = horizonPerks.feeBpsFor(user);
        fee = (collateralIn * feeBps) / 10000;
        uint256 collateralAfterFee = collateralIn - fee;

        uint256 reserveIn = isBuyYes ? reserveNo : reserveYes;
        uint256 reserveOut = isBuyYes ? reserveYes : reserveNo;

        uint256 newReserveIn = reserveIn + collateralAfterFee;
        uint256 k = reserveIn * reserveOut;
        uint256 newReserveOut = k / newReserveIn;
        tokensOut = reserveOut - newReserveOut;
    }

    /**
     * @notice Internal sell quote calculation
     */
    function _getQuoteSell(uint256 tokensIn, bool isSellYes, address user)
        internal
        view
        returns (uint256 collateralOut, uint256 fee)
    {
        if (tokensIn == 0 || reserveYes == 0 || reserveNo == 0) return (0, 0);

        uint256 reserveIn = isSellYes ? reserveYes : reserveNo;
        uint256 reserveOut = isSellYes ? reserveNo : reserveYes;

        uint256 newReserveIn = reserveIn + tokensIn;
        uint256 k = reserveIn * reserveOut;
        uint256 newReserveOut = k / newReserveIn;
        collateralOut = reserveOut - newReserveOut;

        uint16 feeBps = horizonPerks.feeBpsFor(user);
        fee = (collateralOut * feeBps) / 10000;
        collateralOut -= fee;
    }

    /**
     * @notice Gets the current Yes token price
     * @return price Price in 1e18 precision (1e18 = 1.0 = 100%)
     */
    function getYesPrice() external view returns (uint256 price) {
        if (reserveYes == 0 || reserveNo == 0) return 0.5e18; // 50% if no liquidity
        price = (reserveNo * 1e18) / (reserveYes + reserveNo);
    }

    /**
     * @notice Gets the current No token price
     * @return price Price in 1e18 precision
     */
    function getNoPrice() external view returns (uint256 price) {
        if (reserveYes == 0 || reserveNo == 0) return 0.5e18; // 50% if no liquidity
        price = (reserveYes * 1e18) / (reserveYes + reserveNo);
    }

    // ============ Resolution Functions ============

    /**
     * @notice Transfers collateral to OutcomeToken for redemptions after market resolution
     * @dev Should be called after market is resolved to enable winner redemptions
     */
    function fundRedemptions() external nonReentrant {
        if (!outcomeToken.isResolved(marketId)) revert InvalidState();

        uint256 collateralBalance = collateralToken.balanceOf(address(this));
        if (collateralBalance > 0) {
            collateralToken.safeTransfer(address(outcomeToken), collateralBalance);
        }
    }

    // ============ Admin Functions ============

    /**
     * @notice Pauses trading (emergency only)
     */
    function pause() external {
        // In production, this should be onlyOwner or governance
        _pause();
    }

    /**
     * @notice Unpauses trading
     */
    function unpause() external {
        // In production, this should be onlyOwner or governance
        _unpause();
    }

    // ============ ERC1155 Receiver Implementation ============

    /**
     * @notice Handles the receipt of a single ERC1155 token type
     */
    function onERC1155Received(address, address, uint256, uint256, bytes calldata)
        external
        pure
        override
        returns (bytes4)
    {
        return this.onERC1155Received.selector;
    }

    /**
     * @notice Handles the receipt of multiple ERC1155 token types
     */
    function onERC1155BatchReceived(address, address, uint256[] calldata, uint256[] calldata, bytes calldata)
        external
        pure
        override
        returns (bytes4)
    {
        return this.onERC1155BatchReceived.selector;
    }

    /**
     * @notice Indicates whether the contract implements a given interface
     */
    function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
        return interfaceId == type(IERC1155Receiver).interfaceId;
    }
}
