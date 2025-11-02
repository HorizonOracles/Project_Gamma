// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IMarket
 * @notice Common interface for all prediction market types
 * @dev All market implementations (BinaryMarket, MultiChoiceMarket, etc.) should implement this interface
 *      This provides a standardized way to interact with any market type
 */
interface IMarket {
    // ============ Enums ============

    enum MarketType {
        Binary,             // Traditional Yes/No market
        MultiChoice,        // 3-8 discrete outcomes
        LimitOrder,         // Order book based market
        PooledLiquidity,    // Concentrated liquidity AMM
        Dependent,          // Conditional market with parent dependency
        Bracket,            // Value range prediction
        Trend               // Time-weighted outcome tracking
    }

    // ============ Structs ============

    struct MarketInfo {
        uint256 marketId;
        MarketType marketType;
        address collateralToken;
        uint256 closeTime;
        uint256 outcomeCount;
        bool isResolved;
        bool isPaused;
    }

    // ============ Events ============

    event Trade(
        address indexed trader,
        uint256 indexed outcomeId,
        uint256 amountIn,
        uint256 amountOut,
        uint256 fee,
        bool isBuy
    );

    event LiquidityChanged(
        address indexed provider,
        uint256 amount,
        bool isAddition
    );

    // ============ Core Market Functions ============

    /**
     * @notice Gets the market type
     * @return Market type enum value
     */
    function getMarketType() external view returns (MarketType);

    /**
     * @notice Gets comprehensive market information
     * @return MarketInfo struct with market details
     */
    function getMarketInfo() external view returns (MarketInfo memory);

    /**
     * @notice Gets the unique market identifier
     * @return Market ID
     */
    function marketId() external view returns (uint256);

    /**
     * @notice Gets the collateral token address
     * @return Collateral token ERC20 address
     */
    function collateralToken() external view returns (IERC20);

    /**
     * @notice Gets the market close timestamp
     * @return Unix timestamp when market closes
     */
    function closeTime() external view returns (uint256);

    /**
     * @notice Gets the number of possible outcomes
     * @return Total outcome count
     */
    function getOutcomeCount() external view returns (uint256);

    // ============ Trading Functions ============

    /**
     * @notice Buys outcome tokens
     * @param outcomeId ID of outcome to buy
     * @param collateralIn Amount of collateral to spend
     * @param minTokensOut Minimum tokens expected (slippage protection)
     * @return tokensOut Amount of outcome tokens received
     */
    function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) 
        external 
        returns (uint256 tokensOut);

    /**
     * @notice Sells outcome tokens
     * @param outcomeId ID of outcome to sell
     * @param tokensIn Amount of tokens to sell
     * @param minCollateralOut Minimum collateral expected (slippage protection)
     * @return collateralOut Amount of collateral received
     */
    function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut)
        external
        returns (uint256 collateralOut);

    // ============ Liquidity Functions ============

    /**
     * @notice Adds liquidity to the market
     * @param amount Amount of collateral to add
     * @return lpTokens Amount of LP tokens minted
     */
    function addLiquidity(uint256 amount) external returns (uint256 lpTokens);

    /**
     * @notice Removes liquidity from the market
     * @param lpTokens Amount of LP tokens to burn
     * @return collateralOut Amount of collateral returned
     */
    function removeLiquidity(uint256 lpTokens) external returns (uint256 collateralOut);

    // ============ Price & Quote Functions ============

    /**
     * @notice Gets current price for an outcome
     * @param outcomeId Outcome identifier
     * @return price Price in 1e18 precision (1e18 = 100%)
     */
    function getPrice(uint256 outcomeId) external view returns (uint256 price);

    /**
     * @notice Gets quote for buying outcome tokens
     * @param outcomeId Outcome to buy
     * @param collateralIn Amount of collateral to spend
     * @param user User address (for fee tier calculation)
     * @return tokensOut Expected tokens received
     * @return fee Fee amount
     */
    function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user)
        external
        view
        returns (uint256 tokensOut, uint256 fee);

    /**
     * @notice Gets quote for selling outcome tokens
     * @param outcomeId Outcome to sell
     * @param tokensIn Amount of tokens to sell
     * @param user User address (for fee tier calculation)
     * @return collateralOut Expected collateral received
     * @return fee Fee amount
     */
    function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user)
        external
        view
        returns (uint256 collateralOut, uint256 fee);

    // ============ State Management Functions ============

    /**
     * @notice Pauses trading (emergency only)
     */
    function pause() external;

    /**
     * @notice Unpauses trading
     */
    function unpause() external;

    /**
     * @notice Transfers collateral to OutcomeToken for redemptions after resolution
     */
    function fundRedemptions() external;
}
