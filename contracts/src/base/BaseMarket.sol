// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "../interfaces/IMarket.sol";
import "../OutcomeToken.sol";
import "../FeeSplitter.sol";
import "../HorizonPerks.sol";

/**
 * @title BaseMarket
 * @notice Abstract base contract for all prediction market types
 * @dev Provides common functionality for liquidity, fees, access control, and state management
 *      Derived contracts implement specific trading logic (AMM, order book, etc.)
 */
abstract contract BaseMarket is IMarket, ERC20, ReentrancyGuard, Pausable, IERC1155Receiver {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error MarketClosed();
    error MarketResolved();
    error InsufficientLiquidity();
    error SlippageExceeded();
    error InvalidAmount();
    error InvalidOutcomeId();
    error InvalidState();
    error InsufficientLPTokens();
    error MinimumLiquidityRequired();
    error Unauthorized();

    // ============ Constants ============

    /// @notice Minimum liquidity locked forever (similar to Uniswap V2)
    uint256 public constant MINIMUM_LIQUIDITY = 1000;

    /// @notice Precision for price calculations (1e18 = 100%)
    uint256 public constant PRICE_PRECISION = 1e18;

    // ============ Immutable State ============

    /// @notice Market identifier
    uint256 public immutable override marketId;

    /// @notice Collateral token (e.g., USDC)
    IERC20 public immutable override collateralToken;

    /// @notice Outcome token contract
    OutcomeToken public immutable outcomeToken;

    /// @notice Fee splitter contract
    FeeSplitter public immutable feeSplitter;

    /// @notice Horizon perks contract (for fee discounts)
    HorizonPerks public immutable horizonPerks;

    /// @notice Market close timestamp
    uint256 public immutable override closeTime;

    /// @notice Market type
    MarketType public immutable marketType;

    /// @notice Number of possible outcomes
    uint256 public immutable outcomeCount;

    // ============ Mutable State ============

    /// @notice Total collateral backing the reserves
    uint256 public totalCollateral;

    /// @notice Admin address (for pause/unpause)
    address public admin;

    // ============ Constructor ============

    /**
     * @notice Initializes the base market
     * @param _marketId Market identifier
     * @param _marketType Type of market
     * @param _collateralToken Collateral token address
     * @param _outcomeToken Outcome token contract
     * @param _feeSplitter Fee splitter contract
     * @param _horizonPerks Horizon perks contract
     * @param _closeTime Market close timestamp
     * @param _outcomeCount Number of outcomes
     * @param _lpTokenName LP token name
     * @param _lpTokenSymbol LP token symbol
     */
    constructor(
        uint256 _marketId,
        MarketType _marketType,
        address _collateralToken,
        address _outcomeToken,
        address _feeSplitter,
        address _horizonPerks,
        uint256 _closeTime,
        uint256 _outcomeCount,
        string memory _lpTokenName,
        string memory _lpTokenSymbol
    ) ERC20(_lpTokenName, _lpTokenSymbol) {
        marketId = _marketId;
        marketType = _marketType;
        collateralToken = IERC20(_collateralToken);
        outcomeToken = OutcomeToken(_outcomeToken);
        feeSplitter = FeeSplitter(_feeSplitter);
        horizonPerks = HorizonPerks(_horizonPerks);
        closeTime = _closeTime;
        outcomeCount = _outcomeCount;
        admin = msg.sender;
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

    modifier onlyAdmin() {
        if (msg.sender != admin) revert Unauthorized();
        _;
    }

    modifier validOutcome(uint256 outcomeId) {
        if (outcomeId >= outcomeCount) revert InvalidOutcomeId();
        _;
    }

    // ============ IMarket Implementation ============

    /**
     * @inheritdoc IMarket
     */
    function getMarketType() external view override returns (MarketType) {
        return marketType;
    }

    /**
     * @inheritdoc IMarket
     */
    function getMarketInfo() external view override returns (MarketInfo memory) {
        return MarketInfo({
            marketId: marketId,
            marketType: marketType,
            collateralToken: address(collateralToken),
            closeTime: closeTime,
            outcomeCount: outcomeCount,
            isResolved: outcomeToken.isResolved(marketId),
            isPaused: paused()
        });
    }

    /**
     * @inheritdoc IMarket
     */
    function getOutcomeCount() external view override returns (uint256) {
        return outcomeCount;
    }

    // ============ Admin Functions ============

    /**
     * @notice Sets a new admin address
     * @param newAdmin New admin address
     */
    function setAdmin(address newAdmin) external onlyAdmin {
        if (newAdmin == address(0)) revert InvalidAmount();
        admin = newAdmin;
    }

    /**
     * @inheritdoc IMarket
     */
    function pause() external override onlyAdmin {
        _pause();
    }

    /**
     * @inheritdoc IMarket
     */
    function unpause() external override onlyAdmin {
        _unpause();
    }

    /**
     * @inheritdoc IMarket
     */
    function fundRedemptions() external override nonReentrant {
        if (!outcomeToken.isResolved(marketId)) revert InvalidState();

        uint256 collateralBalance = collateralToken.balanceOf(address(this));
        if (collateralBalance > 0) {
            collateralToken.safeTransfer(address(outcomeToken), collateralBalance);
        }
    }

    // ============ Fee Calculation Helpers ============

    /**
     * @notice Calculates fee for a transaction
     * @param amount Amount to calculate fee on
     * @param user User address for fee tier
     * @return fee Fee amount
     * @return protocolBps Protocol fee basis points
     */
    function _calculateFee(uint256 amount, address user) 
        internal 
        view 
        returns (uint256 fee, uint16 protocolBps) 
    {
        uint16 feeBps = horizonPerks.feeBpsFor(user);
        protocolBps = horizonPerks.protocolBpsFor(user);
        fee = (amount * feeBps) / 10000;
    }

    /**
     * @notice Distributes fees through the FeeSplitter
     * @param amount Fee amount to distribute
     * @param protocolBps Protocol fee basis points
     */
    function _distributeFee(uint256 amount, uint16 protocolBps) internal {
        if (amount > 0) {
            collateralToken.forceApprove(address(feeSplitter), amount);
            feeSplitter.distribute(marketId, address(collateralToken), amount, protocolBps);
        }
    }

    // ============ Outcome Token Helpers ============

    /**
     * @notice Mints outcome tokens for the market
     * @param outcomeId Outcome identifier
     * @param to Address to mint to
     * @param amount Amount to mint
     */
    function _mintOutcome(uint256 outcomeId, address to, uint256 amount) internal {
        outcomeToken.mintOutcome(marketId, outcomeId, to, amount);
    }

    /**
     * @notice Burns outcome tokens from the market
     * @param outcomeId Outcome identifier
     * @param from Address to burn from
     * @param amount Amount to burn
     */
    function _burnOutcome(uint256 outcomeId, address from, uint256 amount) internal {
        outcomeToken.burnOutcome(marketId, outcomeId, from, amount);
    }

    /**
     * @notice Transfers outcome tokens between addresses
     * @param outcomeId Outcome identifier
     * @param from Source address
     * @param to Destination address
     * @param amount Amount to transfer
     */
    function _transferOutcome(uint256 outcomeId, address from, address to, uint256 amount) internal {
        _burnOutcome(outcomeId, from, amount);
        _mintOutcome(outcomeId, to, amount);
    }

    // ============ Validation Helpers ============

    /**
     * @notice Validates slippage protection
     * @param amountOut Actual amount out
     * @param minAmountOut Minimum amount expected
     */
    function _validateSlippage(uint256 amountOut, uint256 minAmountOut) internal pure {
        if (amountOut < minAmountOut) revert SlippageExceeded();
    }

    /**
     * @notice Validates amount is non-zero
     * @param amount Amount to validate
     */
    function _validateAmount(uint256 amount) internal pure {
        if (amount == 0) revert InvalidAmount();
    }

    // ============ Abstract Functions (Must be implemented by derived contracts) ============

    /**
     * @notice Buys outcome tokens - implementation specific to market type
     */
    function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) 
        external 
        virtual 
        override 
        returns (uint256 tokensOut);

    /**
     * @notice Sells outcome tokens - implementation specific to market type
     */
    function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut)
        external
        virtual
        override
        returns (uint256 collateralOut);

    /**
     * @notice Adds liquidity - implementation specific to market type
     */
    function addLiquidity(uint256 amount) 
        external 
        virtual 
        override 
        returns (uint256 lpTokens);

    /**
     * @notice Removes liquidity - implementation specific to market type
     */
    function removeLiquidity(uint256 lpTokens) 
        external 
        virtual 
        override 
        returns (uint256 collateralOut);

    /**
     * @notice Gets price for outcome - implementation specific to market type
     */
    function getPrice(uint256 outcomeId) 
        external 
        view 
        virtual 
        override 
        returns (uint256 price);

    /**
     * @notice Gets buy quote - implementation specific to market type
     */
    function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user)
        external
        view
        virtual
        override
        returns (uint256 tokensOut, uint256 fee);

    /**
     * @notice Gets sell quote - implementation specific to market type
     */
    function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user)
        external
        view
        virtual
        override
        returns (uint256 collateralOut, uint256 fee);

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
