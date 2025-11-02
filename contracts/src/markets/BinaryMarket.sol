// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../base/BaseMarket.sol";

/**
 * @title BinaryMarket
 * @notice Binary prediction market with static 1:1 pricing
 * @dev Features:
 *      - Static pricing: 1 collateral = 1 share (no AMM curve)
 *      - 2% fee on both buys and sells
 *      - Selling allowed before deadline, locked after
 *      - Standard 1:1 redemption (winners get 1 collateral per share)
 */
contract BinaryMarket is BaseMarket {
    using SafeERC20 for IERC20;

    // ============ Constants ============

    /// @notice Yes outcome ID
    uint256 public constant OUTCOME_YES = 0;

    /// @notice No outcome ID
    uint256 public constant OUTCOME_NO = 1;

    /// @notice Fixed fee: 2% (200 basis points)
    uint256 public constant FIXED_FEE_BPS = 200;

    // ============ State Variables ============

    /// @notice Total collateral in YES pool
    uint256 public yesPool;

    /// @notice Total collateral in NO pool
    uint256 public noPool;

    /// @notice Total shares minted for YES outcome
    uint256 public totalYesShares;

    /// @notice Total shares minted for NO outcome
    uint256 public totalNoShares;

    /// @notice Whether fundRedemptions has been called
    bool public redemptionsFunded;

    // ============ Events ============

    event SharesPurchased(
        address indexed buyer,
        uint256 indexed outcomeId,
        uint256 shares,
        uint256 collateralPaid,
        uint256 fee
    );

    event SharesSold(
        address indexed seller,
        uint256 indexed outcomeId,
        uint256 shares,
        uint256 collateralReceived,
        uint256 fee
    );

    // ============ Errors ============

    error SellingDisabled();
    error InvalidOutcomeCount();

    // ============ Constructor ============

    /**
     * @notice Creates a new binary market
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
    )
        BaseMarket(
            _marketId,
            MarketType.Binary,
            _collateralToken,
            _outcomeToken,
            _feeSplitter,
            _horizonPerks,
            _closeTime,
            2, // Binary market always has 2 outcomes
            "Binary Market LP",
            "BM-LP"
        )
    {
        if (_closeTime <= block.timestamp) revert MarketClosed();
    }

    // ============ Trading Functions ============

    /**
     * @notice Buy shares at static 1:1 price with 2% fee
     * @param outcomeId Outcome to buy (0=YES, 1=NO)
     * @param collateralIn Amount of collateral to spend (including fee)
     * @param minTokensOut Minimum shares expected (slippage protection)
     * @return tokensOut Number of shares received
     */
    function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut)
        public
        override
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        validOutcome(outcomeId)
        returns (uint256 tokensOut)
    {
        _validateAmount(collateralIn);

        // Calculate fee (2%)
        uint256 fee = (collateralIn * FIXED_FEE_BPS) / 10000;
        uint256 collateralAfterFee = collateralIn - fee;

        // Static 1:1: shares = collateral after fee
        tokensOut = collateralAfterFee;

        // Slippage check
        _validateSlippage(tokensOut, minTokensOut);

        // Transfer collateral from buyer
        collateralToken.safeTransferFrom(msg.sender, address(this), collateralIn);

        // Distribute fee
        uint16 protocolBps = horizonPerks.protocolBpsFor(msg.sender);
        _distributeFee(fee, protocolBps);

        // Update pools and mint shares
        if (outcomeId == OUTCOME_YES) {
            yesPool += collateralAfterFee;
            totalYesShares += tokensOut;
        } else {
            noPool += collateralAfterFee;
            totalNoShares += tokensOut;
        }

        totalCollateral += collateralAfterFee;

        // Mint outcome tokens to buyer
        _mintOutcome(outcomeId, msg.sender, tokensOut);

        emit SharesPurchased(msg.sender, outcomeId, tokensOut, collateralIn, fee);
    }

    /**
     * @notice Sell shares at static 1:1 price with 2% fee (only before deadline)
     * @param outcomeId Outcome to sell (0=YES, 1=NO)
     * @param tokensIn Number of shares to sell
     * @param minCollateralOut Minimum collateral expected (slippage protection)
     * @return collateralOut Amount of collateral received
     */
    function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut)
        public
        override
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        validOutcome(outcomeId)
        returns (uint256 collateralOut)
    {
        _validateAmount(tokensIn);

        // Static 1:1: collateral before fee = shares
        uint256 collateralBeforeFee = tokensIn;

        // Calculate fee (2%)
        uint256 fee = (collateralBeforeFee * FIXED_FEE_BPS) / 10000;
        collateralOut = collateralBeforeFee - fee;

        // Slippage check
        _validateSlippage(collateralOut, minCollateralOut);

        // Check pool has enough collateral
        uint256 pool = (outcomeId == OUTCOME_YES) ? yesPool : noPool;
        if (pool < collateralBeforeFee) revert InsufficientLiquidity();

        // Burn shares from seller
        _burnOutcome(outcomeId, msg.sender, tokensIn);

        // Update pools
        if (outcomeId == OUTCOME_YES) {
            yesPool -= collateralBeforeFee;
            totalYesShares -= tokensIn;
        } else {
            noPool -= collateralBeforeFee;
            totalNoShares -= tokensIn;
        }

        totalCollateral -= collateralBeforeFee;

        // Distribute fee
        uint16 protocolBps = horizonPerks.protocolBpsFor(msg.sender);
        _distributeFee(fee, protocolBps);

        // Transfer collateral to seller (after fee)
        collateralToken.safeTransfer(msg.sender, collateralOut);

        emit SharesSold(msg.sender, outcomeId, tokensIn, collateralOut, fee);
    }

    // ============ Convenience Functions ============

    /**
     * @notice Buy YES shares (convenience function)
     * @param collateralIn Amount of collateral to spend (including fee)
     * @param minTokensOut Minimum shares expected
     * @return tokensOut Number of shares received
     */
    function buyYes(uint256 collateralIn, uint256 minTokensOut)
        external
        returns (uint256 tokensOut)
    {
        return buy(OUTCOME_YES, collateralIn, minTokensOut);
    }

    /**
     * @notice Buy NO shares (convenience function)
     * @param collateralIn Amount of collateral to spend (including fee)
     * @param minTokensOut Minimum shares expected
     * @return tokensOut Number of shares received
     */
    function buyNo(uint256 collateralIn, uint256 minTokensOut)
        external
        returns (uint256 tokensOut)
    {
        return buy(OUTCOME_NO, collateralIn, minTokensOut);
    }

    /**
     * @notice Sell YES shares (convenience function)
     * @param tokensIn Number of shares to sell
     * @param minCollateralOut Minimum collateral expected
     * @return collateralOut Amount of collateral received
     */
    function sellYes(uint256 tokensIn, uint256 minCollateralOut)
        external
        returns (uint256 collateralOut)
    {
        return sell(OUTCOME_YES, tokensIn, minCollateralOut);
    }

    /**
     * @notice Sell NO shares (convenience function)
     * @param tokensIn Number of shares to sell
     * @param minCollateralOut Minimum collateral expected
     * @return collateralOut Amount of collateral received
     */
    function sellNo(uint256 tokensIn, uint256 minCollateralOut)
        external
        returns (uint256 collateralOut)
    {
        return sell(OUTCOME_NO, tokensIn, minCollateralOut);
    }

    // ============ Liquidity Functions ============

    /**
     * @notice Add initial liquidity to the market
     * @param amount Amount of collateral to add
     * @return lpTokens LP tokens minted
     */
    function addLiquidity(uint256 amount)
        external
        override
        nonReentrant
        whenNotPaused
        notResolved
        returns (uint256 lpTokens)
    {
        _validateAmount(amount);

        // Transfer collateral
        collateralToken.safeTransferFrom(msg.sender, address(this), amount);

        // Calculate LP tokens to mint
        uint256 _totalSupply = totalSupply();
        if (_totalSupply == 0) {
            // First liquidity provider
            lpTokens = amount;
            if (lpTokens <= MINIMUM_LIQUIDITY) revert MinimumLiquidityRequired();
            
            // Lock minimum liquidity (burn to dead address, not address(0) due to ERC20 restrictions)
            _mint(address(0xdead), MINIMUM_LIQUIDITY);
            lpTokens -= MINIMUM_LIQUIDITY;

            // Split liquidity equally between both outcomes initially
            uint256 halfAmount = amount / 2;
            yesPool = halfAmount;
            noPool = amount - halfAmount;
            
            // Set initial share counts (1:1 with collateral)
            totalYesShares = halfAmount;
            totalNoShares = amount - halfAmount;
        } else {
            // Subsequent liquidity: proportional to existing
            lpTokens = (amount * _totalSupply) / totalCollateral;
            
            // Add proportionally to pools
            uint256 yesAdd = (amount * yesPool) / totalCollateral;
            uint256 noAdd = amount - yesAdd;
            
            yesPool += yesAdd;
            noPool += noAdd;
            
            // Update share counts proportionally
            totalYesShares += yesAdd;
            totalNoShares += noAdd;
        }

        totalCollateral += amount;

        // Mint LP tokens
        _mint(msg.sender, lpTokens);

        emit LiquidityAdded(msg.sender, amount, lpTokens);
    }

    /**
     * @notice Remove liquidity from the market
     * @param lpTokens Amount of LP tokens to burn
     * @return collateralOut Amount of collateral received
     */
    function removeLiquidity(uint256 lpTokens)
        external
        override
        nonReentrant
        returns (uint256 collateralOut)
    {
        _validateAmount(lpTokens);

        uint256 _totalSupply = totalSupply();
        if (lpTokens > balanceOf(msg.sender)) revert InsufficientLPTokens();

        // Calculate share of pool
        collateralOut = (lpTokens * totalCollateral) / _totalSupply;

        if (collateralOut == 0) revert InvalidAmount();

        // Update pools proportionally
        uint256 yesRemove = (lpTokens * yesPool) / _totalSupply;
        uint256 noRemove = (lpTokens * noPool) / _totalSupply;

        yesPool -= yesRemove;
        noPool -= noRemove;
        totalCollateral -= collateralOut;

        // Update share counts proportionally
        uint256 yesSharesRemove = (lpTokens * totalYesShares) / _totalSupply;
        uint256 noSharesRemove = (lpTokens * totalNoShares) / _totalSupply;
        
        totalYesShares -= yesSharesRemove;
        totalNoShares -= noSharesRemove;

        // Burn LP tokens
        _burn(msg.sender, lpTokens);

        // Transfer collateral
        collateralToken.safeTransfer(msg.sender, collateralOut);

        emit LiquidityRemoved(msg.sender, lpTokens, collateralOut);
    }

    // ============ View Functions ============

    /**
     * @notice Get price for an outcome (always 50/50 static pricing)
     * @param outcomeId Outcome ID
     * @return price Price in PRICE_PRECISION (always 0.5 = 5e17)
     */
    function getPrice(uint256 outcomeId)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 price)
    {
        // Static 50/50 pricing
        return PRICE_PRECISION / 2; // 0.5
    }

    /**
     * @notice Get quote for buying shares
     * @param outcomeId Outcome to buy
     * @param collateralIn Amount of collateral (including fee)
     * @param user User address (for fee tier)
     * @return tokensOut Shares that will be received
     * @return fee Fee amount
     */
    function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 tokensOut, uint256 fee)
    {
        fee = (collateralIn * FIXED_FEE_BPS) / 10000;
        tokensOut = collateralIn - fee;
    }

    /**
     * @notice Get quote for selling shares
     * @param outcomeId Outcome to sell
     * @param tokensIn Number of shares to sell
     * @param user User address (for fee tier)
     * @return collateralOut Collateral that will be received
     * @return fee Fee amount
     */
    function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 collateralOut, uint256 fee)
    {
        fee = (tokensIn * FIXED_FEE_BPS) / 10000;
        collateralOut = tokensIn - fee;
    }

    /**
     * @notice Get current reserves for both outcomes
     * @return yesReserve YES pool collateral
     * @return noReserve NO pool collateral
     */
    function getReserves() external view returns (uint256 yesReserve, uint256 noReserve) {
        return (yesPool, noPool);
    }

    // ============ Redemption Functions ============

    /**
     * @notice Transfers collateral to OutcomeToken for standard 1:1 redemptions
     * @dev Must be called after resolution to enable redemptions
     *      Winners receive 1:1 redemption (1 share = 1 collateral)
     *      Can be called multiple times safely (idempotent)
     */
    function fundRedemptions() external override nonReentrant {
        if (!outcomeToken.isResolved(marketId)) revert InvalidState();

        // Transfer all remaining collateral to OutcomeToken for standard 1:1 redemption
        uint256 collateralBalance = collateralToken.balanceOf(address(this));
        if (collateralBalance > 0) {
            collateralToken.safeTransfer(address(outcomeToken), collateralBalance);
            redemptionsFunded = true;
        }
    }

    // ============ Events ============

    event LiquidityAdded(address indexed provider, uint256 collateral, uint256 lpTokens);
    event LiquidityRemoved(address indexed provider, uint256 lpTokens, uint256 collateral);
}
