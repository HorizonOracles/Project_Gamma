// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../base/BaseMarket.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title PooledLiquidityMarket
 * @notice Concentrated liquidity AMM for prediction markets (Uniswap V3-style)
 * @dev Allows LPs to provide liquidity in specific price ranges for capital efficiency
 * 
 * Key Features:
 * - Concentrated liquidity: LPs choose price ranges
 * - Tick-based pricing: Discrete price levels for efficient storage
 * - Range orders: Limit orders that earn fees as LP positions
 * - Fee accumulation: Per-position fee tracking
 * - Capital efficiency: Higher returns for active liquidity
 * 
 * Architecture:
 * - Ticks: Discrete price points (e.g., every 0.01 = 1% increments)
 * - Positions: LP deposits in [tickLower, tickUpper] range
 * - Virtual reserves: Only active liquidity participates in swaps
 * - Fee growth: Global and per-tick fee accumulation tracking
 */
contract PooledLiquidityMarket is BaseMarket {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error InvalidTickRange();
    error InvalidTick();
    error InsufficientPosition();
    error PositionNotFound();
    error TickSpacingError();

    // ============ Constants ============

    /// @notice Minimum tick value (price = 0.01 or 1%)
    int24 public constant MIN_TICK = -69000; // ~0.01
    
    /// @notice Maximum tick value (price = 0.99 or 99%)
    int24 public constant MAX_TICK = 69000; // ~0.99
    
    /// @notice Tick spacing (10 = 0.1% increments, balances precision vs gas)
    int24 public constant TICK_SPACING = 10;
    
    /// @notice Fee tier for swaps (30 bps = 0.3%)
    uint24 public constant FEE_TIER = 3000;

    /// @notice Fixed point precision for Q64.96 math (2^96)
    uint256 internal constant Q96 = 0x1000000000000000000000000;

    // ============ Structs ============

    /**
     * @notice Liquidity position owned by an LP
     * @param liquidity Amount of liquidity (sqrt(x*y))
     * @param feeGrowthInside0LastX128 Fee growth per unit liquidity when position was last updated (outcome 0)
     * @param feeGrowthInside1LastX128 Fee growth per unit liquidity when position was last updated (outcome 1)
     * @param tokensOwed0 Uncollected fees for outcome 0
     * @param tokensOwed1 Uncollected fees for outcome 1
     */
    struct Position {
        uint128 liquidity;
        uint256 feeGrowthInside0LastX128;
        uint256 feeGrowthInside1LastX128;
        uint128 tokensOwed0;
        uint128 tokensOwed1;
    }

    /**
     * @notice Information about a tick
     * @param liquidityGross Total liquidity that references this tick
     * @param liquidityNet Amount of liquidity added (sub) when tick is crossed left to right (right to left)
     * @param feeGrowthOutside0X128 Fee growth per unit liquidity on the other side of this tick (outcome 0)
     * @param feeGrowthOutside1X128 Fee growth per unit liquidity on the other side of this tick (outcome 1)
     * @param initialized Whether the tick is initialized
     */
    struct TickInfo {
        uint128 liquidityGross;
        int128 liquidityNet;
        uint256 feeGrowthOutside0X128;
        uint256 feeGrowthOutside1X128;
        bool initialized;
    }

    /**
     * @notice Global pool state
     * @param sqrtPriceX96 Current sqrt(price) as Q64.96
     * @param tick Current tick
     * @param liquidity Current active liquidity
     * @param feeGrowthGlobal0X128 Global fee growth per unit liquidity (outcome 0)
     * @param feeGrowthGlobal1X128 Global fee growth per unit liquidity (outcome 1)
     */
    struct PoolState {
        uint160 sqrtPriceX96; // 20 bytes
        int24 tick;           // 3 bytes
        uint128 liquidity;    // 16 bytes (total 39 bytes - won't pack)
        uint256 feeGrowthGlobal0X128; // 32 bytes
        uint256 feeGrowthGlobal1X128; // 32 bytes
    }

    // ============ State Variables ============

    /// @notice Pool state
    PoolState public poolState;

    /// @notice Mapping from tick to tick info
    mapping(int24 => TickInfo) public ticks;

    /// @notice Mapping from position key to position info
    /// @dev positionKey = keccak256(abi.encodePacked(owner, tickLower, tickUpper))
    mapping(bytes32 => Position) public positions;

    /// @notice Reserve balances for each outcome
    mapping(uint256 => uint256) public reserves;

    // ============ Events ============

    event PositionMinted(
        address indexed owner,
        int24 indexed tickLower,
        int24 indexed tickUpper,
        uint128 liquidity,
        uint256 amount0,
        uint256 amount1
    );

    event PositionBurned(
        address indexed owner,
        int24 indexed tickLower,
        int24 indexed tickUpper,
        uint128 liquidity,
        uint256 amount0,
        uint256 amount1
    );

    event FeesCollected(
        address indexed owner,
        int24 indexed tickLower,
        int24 indexed tickUpper,
        uint128 amount0,
        uint128 amount1
    );

    event Swap(
        address indexed trader,
        uint256 indexed outcomeId,
        uint256 amountIn,
        uint256 amountOut,
        uint160 sqrtPriceX96,
        int24 tick
    );

    // ============ Constructor ============

    constructor(
        uint256 _marketId,
        address _collateralToken,
        address _outcomeToken,
        address _feeSplitter,
        address _horizonPerks,
        uint256 _closeTime,
        string memory _lpTokenName,
        string memory _lpTokenSymbol
    )
        BaseMarket(
            _marketId,
            MarketType.PooledLiquidity,
            _collateralToken,
            _outcomeToken,
            _feeSplitter,
            _horizonPerks,
            _closeTime,
            2, // Binary market (2 outcomes)
            _lpTokenName,
            _lpTokenSymbol
        )
    {
        // Initialize pool at 50/50 (tick 0, price = 0.5)
        // Use _getSqrtRatioAtTick(0) but inline to avoid calling external function in constructor
        // For tick 0: price = 0.5, sqrtPrice = sqrt(0.5) ≈ 0.707
        // sqrtPriceX96 = sqrt(0.5e18) * 2^96 / 1e9
        poolState.sqrtPriceX96 = 56022770974670905984299832681; // sqrt(0.5) * 2^96 / 1e9
        poolState.tick = 0;
    }

    // ============ Core Liquidity Functions ============

    /**
     * @notice Mint a new liquidity position
     * @param tickLower Lower tick of the position range
     * @param tickUpper Upper tick of the position range
     * @param liquidityDesired Amount of liquidity to mint
     * @param amount0Max Maximum amount of outcome 0 tokens willing to spend
     * @param amount1Max Maximum amount of outcome 1 tokens willing to spend
     * @return liquidity Amount of liquidity minted
     * @return amount0 Amount of outcome 0 tokens used
     * @return amount1 Amount of outcome 1 tokens used
     */
    function mintPosition(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidityDesired,
        uint256 amount0Max,
        uint256 amount1Max
    )
        external
        nonReentrant
        beforeClose
        notResolved
        whenNotPaused
        returns (
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1
        )
    {
        return _mintPosition(msg.sender, tickLower, tickUpper, liquidityDesired, amount0Max, amount1Max);
    }

    /**
     * @notice Internal function to mint a liquidity position
     */
    function _mintPosition(
        address owner,
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidityDesired,
        uint256 amount0Max,
        uint256 amount1Max
    )
        internal
        returns (
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1
        )
    {
        _validateTickRange(tickLower, tickUpper);
        _validateAmount(liquidityDesired);

        bytes32 positionKey = _getPositionKey(owner, tickLower, tickUpper);
        Position storage position = positions[positionKey];

        // Calculate amounts needed
        (amount0, amount1) = _getAmountsForLiquidity(
            poolState.sqrtPriceX96,
            tickLower,
            tickUpper,
            liquidityDesired
        );

        if (amount0 > amount0Max || amount1 > amount1Max) {
            revert SlippageExceeded();
        }

        // Update ticks
        _updateTick(tickLower, liquidityDesired, false);
        _updateTick(tickUpper, liquidityDesired, true);

        // Update position
        (
            uint256 feeGrowthInside0X128,
            uint256 feeGrowthInside1X128
        ) = _getFeeGrowthInside(tickLower, tickUpper);

        // Cache position liquidity to save gas on repeated reads
        uint128 positionLiquidity = position.liquidity;
        
        position.tokensOwed0 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside0X128,
                position.feeGrowthInside0LastX128
            )
        );
        position.tokensOwed1 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside1X128,
                position.feeGrowthInside1LastX128
            )
        );

        position.liquidity = positionLiquidity + liquidityDesired;
        position.feeGrowthInside0LastX128 = feeGrowthInside0X128;
        position.feeGrowthInside1LastX128 = feeGrowthInside1X128;

        // Update active liquidity if position is in range
        if (poolState.tick >= tickLower && poolState.tick < tickUpper) {
            poolState.liquidity += liquidityDesired;
        }

        // Transfer tokens from owner
        unchecked {
            if (amount0 > 0) {
                collateralToken.safeTransferFrom(owner, address(this), amount0);
                _mintOutcome(0, address(this), amount0);
                reserves[0] += amount0;
                totalCollateral += amount0;
            }
            if (amount1 > 0) {
                collateralToken.safeTransferFrom(owner, address(this), amount1);
                _mintOutcome(1, address(this), amount1);
                reserves[1] += amount1;
                totalCollateral += amount1;
            }
        }

        liquidity = liquidityDesired;

        emit PositionMinted(owner, tickLower, tickUpper, liquidity, amount0, amount1);
    }

    /**
     * @notice Burn liquidity from a position
     * @param tickLower Lower tick of the position
     * @param tickUpper Upper tick of the position
     * @param liquidityToBurn Amount of liquidity to burn
     * @return amount0 Amount of outcome 0 tokens returned
     * @return amount1 Amount of outcome 1 tokens returned
     */
    function burnPosition(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidityToBurn
    )
        external
        nonReentrant
        returns (uint256 amount0, uint256 amount1)
    {
        _validateAmount(liquidityToBurn);

        bytes32 positionKey = _getPositionKey(msg.sender, tickLower, tickUpper);
        Position storage position = positions[positionKey];

        if (position.liquidity < liquidityToBurn) {
            revert InsufficientPosition();
        }

        // Calculate amounts to return
        (amount0, amount1) = _getAmountsForLiquidity(
            poolState.sqrtPriceX96,
            tickLower,
            tickUpper,
            liquidityToBurn
        );

        // Update ticks
        _updateTick(tickLower, liquidityToBurn, false);
        _updateTick(tickUpper, liquidityToBurn, true);

        // Update position
        (
            uint256 feeGrowthInside0X128,
            uint256 feeGrowthInside1X128
        ) = _getFeeGrowthInside(tickLower, tickUpper);

        // Cache position liquidity to save gas on repeated reads
        uint128 positionLiquidity = position.liquidity;
        
        position.tokensOwed0 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside0X128,
                position.feeGrowthInside0LastX128
            )
        );
        position.tokensOwed1 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside1X128,
                position.feeGrowthInside1LastX128
            )
        );

        unchecked {
            position.tokensOwed0 += uint128(amount0);
            position.tokensOwed1 += uint128(amount1);
        }

        position.liquidity = positionLiquidity - liquidityToBurn;
        position.feeGrowthInside0LastX128 = feeGrowthInside0X128;
        position.feeGrowthInside1LastX128 = feeGrowthInside1X128;

        // Update active liquidity if position is in range
        if (poolState.tick >= tickLower && poolState.tick < tickUpper) {
            poolState.liquidity -= liquidityToBurn;
        }

        unchecked {
            reserves[0] -= amount0;
            reserves[1] -= amount1;
        }

        emit PositionBurned(msg.sender, tickLower, tickUpper, liquidityToBurn, amount0, amount1);
    }

    /**
     * @notice Collect fees from a position
     * @param tickLower Lower tick of the position
     * @param tickUpper Upper tick of the position
     * @return amount0 Amount of outcome 0 fees collected
     * @return amount1 Amount of outcome 1 fees collected
     */
    function collectFees(
        int24 tickLower,
        int24 tickUpper
    )
        external
        nonReentrant
        returns (uint128 amount0, uint128 amount1)
    {
        bytes32 positionKey = _getPositionKey(msg.sender, tickLower, tickUpper);
        Position storage position = positions[positionKey];

        // Calculate accrued fees
        (
            uint256 feeGrowthInside0X128,
            uint256 feeGrowthInside1X128
        ) = _getFeeGrowthInside(tickLower, tickUpper);

        // Cache position liquidity to save gas on repeated reads
        uint128 positionLiquidity = position.liquidity;
        
        position.tokensOwed0 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside0X128,
                position.feeGrowthInside0LastX128
            )
        );
        position.tokensOwed1 += uint128(
            _calculateFeesSinceLastUpdate(
                positionLiquidity,
                feeGrowthInside1X128,
                position.feeGrowthInside1LastX128
            )
        );

        amount0 = position.tokensOwed0;
        amount1 = position.tokensOwed1;

        unchecked {
            if (amount0 > 0) {
                position.tokensOwed0 = 0;
                reserves[0] -= amount0;
                _transferOutcome(0, address(this), msg.sender, amount0);
            }

            if (amount1 > 0) {
                position.tokensOwed1 = 0;
                reserves[1] -= amount1;
                _transferOutcome(1, address(this), msg.sender, amount1);
            }
        }

        position.feeGrowthInside0LastX128 = feeGrowthInside0X128;
        position.feeGrowthInside1LastX128 = feeGrowthInside1X128;

        emit FeesCollected(msg.sender, tickLower, tickUpper, amount0, amount1);
    }

    // ============ BaseMarket Implementation ============

    /**
     * @notice Buy outcome tokens by swapping collateral
     */
    function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut)
        external
        override
        nonReentrant
        beforeClose
        notResolved
        whenNotPaused
        validOutcome(outcomeId)
        returns (uint256 tokensOut)
    {
        _validateAmount(collateralIn);

        // Calculate fee
        (uint256 fee, uint16 protocolBps) = _calculateFee(collateralIn, msg.sender);
        uint256 collateralAfterFee = collateralIn - fee;

        // Transfer collateral from user
        collateralToken.safeTransferFrom(msg.sender, address(this), collateralIn);
        
        // Distribute fee
        _distributeFee(fee, protocolBps);

        // Perform swap (updates reserves and calculates tokensOut)
        tokensOut = _swap(outcomeId, collateralAfterFee, true);

        _validateSlippage(tokensOut, minTokensOut);

        // Mint both outcomes from collateral
        _mintOutcome(0, address(this), collateralAfterFee);
        _mintOutcome(1, address(this), collateralAfterFee);
        
        // Transfer desired outcome tokens to user from pool's balance
        outcomeToken.safeTransferFrom(address(this), msg.sender, _getTokenId(outcomeId), tokensOut, "");

        emit Swap(msg.sender, outcomeId, collateralAfterFee, tokensOut, poolState.sqrtPriceX96, poolState.tick);
    }

    /**
     * @notice Sell outcome tokens for collateral
     */
    function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut)
        external
        override
        nonReentrant
        beforeClose
        notResolved
        whenNotPaused
        validOutcome(outcomeId)
        returns (uint256 collateralOut)
    {
        _validateAmount(tokensIn);

        // Transfer outcome tokens from user to pool
        outcomeToken.safeTransferFrom(msg.sender, address(this), _getTokenId(outcomeId), tokensIn, "");

        // Perform swap (updates reserves and calculates collateralOut)
        collateralOut = _swap(outcomeId, tokensIn, false);

        // Calculate fee
        (uint256 fee, uint16 protocolBps) = _calculateFee(collateralOut, msg.sender);
        uint256 collateralAfterFee = collateralOut - fee;

        _validateSlippage(collateralAfterFee, minCollateralOut);

        // Burn both outcome tokens to get collateral back
        _burnOutcome(0, address(this), collateralOut);
        _burnOutcome(1, address(this), collateralOut);

        // Transfer collateral to user
        collateralToken.safeTransfer(msg.sender, collateralAfterFee);

        // Distribute fee
        _distributeFee(fee, protocolBps);

        emit Swap(msg.sender, outcomeId, tokensIn, collateralAfterFee, poolState.sqrtPriceX96, poolState.tick);
    }

    /**
     * @notice Add liquidity (simplified - adds to full range)
     */
    function addLiquidity(uint256 amount)
        external
        override
        nonReentrant
        beforeClose
        notResolved
        whenNotPaused
        returns (uint256 lpTokens)
    {
        // For simplified addLiquidity, use full range position
        int24 tickLower = MIN_TICK / TICK_SPACING * TICK_SPACING;
        int24 tickUpper = MAX_TICK / TICK_SPACING * TICK_SPACING;

        uint128 liquidityDesired = uint128(amount);
        
        // Call internal mint function to preserve msg.sender
        (uint128 liquidity, uint256 amount0, uint256 amount1) = _mintPosition(
            msg.sender,
            tickLower,
            tickUpper,
            liquidityDesired,
            amount,
            amount
        );

        lpTokens = uint256(liquidity);
        
        // Mint LP tokens proportional to liquidity
        _mint(msg.sender, lpTokens);
    }

    /**
     * @notice Remove liquidity (simplified)
     */
    function removeLiquidity(uint256 lpTokens)
        external
        override
        returns (uint256 collateralOut)
    {
        if (balanceOf(msg.sender) < lpTokens) revert InsufficientLPTokens();

        // Calculate pro-rata share BEFORE burning
        uint256 totalSupplyBefore = totalSupply();
        collateralOut = (totalCollateral * lpTokens) / totalSupplyBefore;
        
        if (collateralOut > totalCollateral) revert InsufficientLiquidity();

        // Burn LP tokens
        _burn(msg.sender, lpTokens);
        
        totalCollateral -= collateralOut;
        collateralToken.safeTransfer(msg.sender, collateralOut);
    }

    /**
     * @notice Get current price for an outcome
     */
    function getPrice(uint256 outcomeId)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 price)
    {
        uint160 sqrtPriceX96 = poolState.sqrtPriceX96;
        uint256 priceX96 = (uint256(sqrtPriceX96) * uint256(sqrtPriceX96)) >> 96;
        
        if (outcomeId == 0) {
            // Price of outcome 0
            price = (priceX96 * PRICE_PRECISION) >> 96;
        } else {
            // Price of outcome 1 = 1 - price of outcome 0
            uint256 price0 = (priceX96 * PRICE_PRECISION) >> 96;
            price = PRICE_PRECISION - price0;
        }
    }

    /**
     * @notice Get quote for buying outcome tokens
     */
    function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 tokensOut, uint256 fee)
    {
        (fee,) = _calculateFee(collateralIn, user);
        uint256 collateralAfterFee = collateralIn - fee;
        
        // Simulate swap
        tokensOut = _simulateSwap(outcomeId, collateralAfterFee, true);
    }

    /**
     * @notice Get quote for selling outcome tokens
     */
    function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 collateralOut, uint256 fee)
    {
        // Simulate swap
        collateralOut = _simulateSwap(outcomeId, tokensIn, false);
        
        (fee,) = _calculateFee(collateralOut, user);
        collateralOut -= fee;
    }

    // ============ Internal Functions ============

    /**
     * @notice Perform a swap
     * @param outcomeId Outcome being traded
     * @param amount Amount in (collateral for buy, outcome tokens for sell)
     * @param isBuy True if buying outcome, false if selling
     * @return amountOut Amount out
     */
    function _swap(
        uint256 outcomeId,
        uint256 amount,
        bool isBuy
    ) internal returns (uint256 amountOut) {
        // Prediction market AMM
        // For buy: amount is collateral, amountOut is outcome tokens
        // For sell: amount is outcome tokens, amountOut is collateral
        
        uint256 reserve0 = reserves[0];
        uint256 reserve1 = reserves[1];

        if (reserve0 == 0 || reserve1 == 0) revert InsufficientLiquidity();

        if (isBuy) {
            // Buying outcome tokens with collateral
            // We will mint `amount` of both outcomes from collateral
            // User gets `amountOut` of desired outcome, pool keeps the rest
            // Price calculated using constant product on current reserves
            if (outcomeId == 0) {
                // Buying outcome 0
                // Use constant product: amountOut = (reserve0 * amount) / (reserve1 + amount)
                amountOut = (reserve0 * amount) / (reserve1 + amount);
                
                if (amountOut > reserve0) revert InsufficientLiquidity();
                
                // After minting `amount` of both outcomes:
                // - Pool gives `amountOut` of outcome0 to user: reserve0 decreases
                // - Pool keeps all `amount` of outcome1: reserve1 increases
                reserves[0] = reserve0 - amountOut + amount;  // Add minted amount, subtract given amount
                reserves[1] = reserve1 + amount;              // Add minted amount
            } else {
                // Buying outcome 1
                amountOut = (reserve1 * amount) / (reserve0 + amount);
                
                if (amountOut > reserve1) revert InsufficientLiquidity();
                
                reserves[0] = reserve0 + amount;
                reserves[1] = reserve1 - amountOut + amount;
            }
        } else {
            // Selling outcome tokens for collateral
            // User gives `amount` of outcome tokens
            // Pool will burn `collateralOut` of both outcomes to get collateral
            // User gets `collateralOut` of collateral
            if (outcomeId == 0) {
                // Selling outcome 0
                // Use constant product: collateralOut = (reserve1 * amount) / (reserve0 + amount)
                amountOut = (reserve1 * amount) / (reserve0 + amount);
                
                if (amountOut > reserve1) revert InsufficientLiquidity();
                
                // Pool receives `amount` of outcome0: reserve0 increases
                // Pool burns `amountOut` of both outcomes: both reserves decrease
                reserves[0] = reserve0 + amount - amountOut;
                reserves[1] = reserve1 - amountOut;
            } else {
                // Selling outcome 1
                amountOut = (reserve0 * amount) / (reserve1 + amount);
                
                if (amountOut > reserve0) revert InsufficientLiquidity();
                
                reserves[0] = reserve0 - amountOut;
                reserves[1] = reserve1 + amount - amountOut;
            }
        }

        // Update price
        _updatePrice();

        // Accrue fees to LPs
        if (poolState.liquidity > 0) {
            uint256 feeAmount = (amount * FEE_TIER) / 1000000;
            if (isBuy) {
                // Fee is in collateral
                if (outcomeId == 0) {
                    poolState.feeGrowthGlobal1X128 += (feeAmount << 128) / poolState.liquidity;
                } else {
                    poolState.feeGrowthGlobal0X128 += (feeAmount << 128) / poolState.liquidity;
                }
            } else {
                // Fee is in outcome tokens
                if (outcomeId == 0) {
                    poolState.feeGrowthGlobal0X128 += (feeAmount << 128) / poolState.liquidity;
                } else {
                    poolState.feeGrowthGlobal1X128 += (feeAmount << 128) / poolState.liquidity;
                }
            }
        }
    }

    /**
     * @notice Simulate a swap without executing
     */
    function _simulateSwap(
        uint256 outcomeId,
        uint256 amount,
        bool isBuy
    ) internal view returns (uint256 amountOut) {
        uint256 reserve0 = reserves[0];
        uint256 reserve1 = reserves[1];

        if (reserve0 == 0 || reserve1 == 0) return 0;

        if (isBuy) {
            if (outcomeId == 0) {
                uint256 newReserve1 = reserve1 + amount;
                uint256 newReserve0 = (reserve0 * reserve1) / newReserve1;
                amountOut = reserve0 - newReserve0;
            } else {
                uint256 newReserve0 = reserve0 + amount;
                uint256 newReserve1 = (reserve0 * reserve1) / newReserve0;
                amountOut = reserve1 - newReserve1;
            }
        } else {
            if (outcomeId == 0) {
                uint256 newReserve0 = reserve0 + amount;
                uint256 newReserve1 = (reserve0 * reserve1) / newReserve0;
                amountOut = reserve1 - newReserve1;
            } else {
                uint256 newReserve1 = reserve1 + amount;
                uint256 newReserve0 = (reserve0 * reserve1) / newReserve1;
                amountOut = reserve0 - newReserve0;
            }
        }
    }

    /**
     * @notice Update pool price based on reserves
     */
    function _updatePrice() internal {
        if (reserves[0] == 0 || reserves[1] == 0) return;
        
        // Calculate marginal price as probability: price0 = reserve1 / (reserve0 + reserve1)
        // This maps reserves to [0, 1] range as required for prediction markets
        uint256 price = (reserves[1] * 1e18) / (reserves[0] + reserves[1]);
        uint256 sqrtPrice = _sqrt(price);
        poolState.sqrtPriceX96 = uint160((sqrtPrice * Q96) / 1e9);
        
        // Update tick based on price
        poolState.tick = _getTickAtSqrtRatio(poolState.sqrtPriceX96);
    }

    /**
     * @notice Get position key
     */
    function _getPositionKey(
        address owner,
        int24 tickLower,
        int24 tickUpper
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(owner, tickLower, tickUpper));
    }

    /**
     * @notice Get ERC1155 token ID for an outcome
     */
    function _getTokenId(uint256 outcomeId) internal view returns (uint256) {
        return (marketId << 8) | outcomeId;
    }

    /**
     * @notice Validate tick range
     */
    function _validateTickRange(int24 tickLower, int24 tickUpper) internal pure {
        if (tickLower >= tickUpper) revert InvalidTickRange();
        if (tickLower < MIN_TICK || tickUpper > MAX_TICK) revert InvalidTick();
        if (tickLower % TICK_SPACING != 0 || tickUpper % TICK_SPACING != 0) {
            revert TickSpacingError();
        }
    }

    /**
     * @notice Update tick info when liquidity changes
     */
    function _updateTick(
        int24 tick,
        uint128 liquidityDelta,
        bool upper
    ) internal {
        TickInfo storage tickInfo = ticks[tick];
        
        uint128 liquidityGrossBefore = tickInfo.liquidityGross;
        
        unchecked {
            tickInfo.liquidityGross = liquidityGrossBefore + liquidityDelta;
        }
        
        if (upper) {
            tickInfo.liquidityNet -= int128(liquidityDelta);
        } else {
            tickInfo.liquidityNet += int128(liquidityDelta);
        }
        
        if (liquidityGrossBefore == 0) {
            tickInfo.initialized = true;
        }
    }

    /**
     * @notice Calculate fee growth inside a position's range
     */
    function _getFeeGrowthInside(
        int24 tickLower,
        int24 tickUpper
    ) internal view returns (uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128) {
        TickInfo storage lower = ticks[tickLower];
        TickInfo storage upper = ticks[tickUpper];
        int24 currentTick = poolState.tick;

        uint256 feeGrowthBelow0X128;
        uint256 feeGrowthBelow1X128;
        if (currentTick >= tickLower) {
            feeGrowthBelow0X128 = lower.feeGrowthOutside0X128;
            feeGrowthBelow1X128 = lower.feeGrowthOutside1X128;
        } else {
            feeGrowthBelow0X128 = poolState.feeGrowthGlobal0X128 - lower.feeGrowthOutside0X128;
            feeGrowthBelow1X128 = poolState.feeGrowthGlobal1X128 - lower.feeGrowthOutside1X128;
        }

        uint256 feeGrowthAbove0X128;
        uint256 feeGrowthAbove1X128;
        if (currentTick < tickUpper) {
            feeGrowthAbove0X128 = upper.feeGrowthOutside0X128;
            feeGrowthAbove1X128 = upper.feeGrowthOutside1X128;
        } else {
            feeGrowthAbove0X128 = poolState.feeGrowthGlobal0X128 - upper.feeGrowthOutside0X128;
            feeGrowthAbove1X128 = poolState.feeGrowthGlobal1X128 - upper.feeGrowthOutside1X128;
        }

        feeGrowthInside0X128 = poolState.feeGrowthGlobal0X128 - feeGrowthBelow0X128 - feeGrowthAbove0X128;
        feeGrowthInside1X128 = poolState.feeGrowthGlobal1X128 - feeGrowthBelow1X128 - feeGrowthAbove1X128;
    }

    /**
     * @notice Calculate fees accrued since last update
     */
    function _calculateFeesSinceLastUpdate(
        uint128 liquidity,
        uint256 feeGrowthInsideX128,
        uint256 feeGrowthInsideLastX128
    ) internal pure returns (uint256) {
        return (uint256(liquidity) * (feeGrowthInsideX128 - feeGrowthInsideLastX128)) >> 128;
    }

    /**
     * @notice Calculate amounts for a given liquidity and price range
     */
    function _getAmountsForLiquidity(
        uint160 sqrtPriceX96,
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidity
    ) internal pure returns (uint256 amount0, uint256 amount1) {
        uint160 sqrtRatioAX96 = _getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = _getSqrtRatioAtTick(tickUpper);

        if (sqrtPriceX96 <= sqrtRatioAX96) {
            amount0 = _getAmount0ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity);
        } else if (sqrtPriceX96 < sqrtRatioBX96) {
            amount0 = _getAmount0ForLiquidity(sqrtPriceX96, sqrtRatioBX96, liquidity);
            amount1 = _getAmount1ForLiquidity(sqrtRatioAX96, sqrtPriceX96, liquidity);
        } else {
            amount1 = _getAmount1ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity);
        }
    }

    /**
     * @notice Calculate amount0 for liquidity
     */
    function _getAmount0ForLiquidity(
        uint160 sqrtRatioAX96,
        uint160 sqrtRatioBX96,
        uint128 liquidity
    ) internal pure returns (uint256) {
        if (sqrtRatioAX96 > sqrtRatioBX96) {
            (sqrtRatioAX96, sqrtRatioBX96) = (sqrtRatioBX96, sqrtRatioAX96);
        }
        uint256 numerator1 = uint256(liquidity) << 96;
        uint256 numerator2 = sqrtRatioBX96 - sqrtRatioAX96;
        
        // Avoid overflow by doing division first if numbers are large
        if (numerator1 > type(uint256).max / numerator2) {
            return (numerator1 / sqrtRatioBX96) * numerator2 / sqrtRatioAX96;
        }
        return (numerator1 * numerator2) / sqrtRatioBX96 / sqrtRatioAX96;
    }

    /**
     * @notice Calculate amount1 for liquidity
     */
    function _getAmount1ForLiquidity(
        uint160 sqrtRatioAX96,
        uint160 sqrtRatioBX96,
        uint128 liquidity
    ) internal pure returns (uint256) {
        if (sqrtRatioAX96 > sqrtRatioBX96) {
            (sqrtRatioAX96, sqrtRatioBX96) = (sqrtRatioBX96, sqrtRatioAX96);
        }
        return (uint256(liquidity) * (sqrtRatioBX96 - sqrtRatioAX96)) / Q96;
    }

    /**
     * @notice Get sqrt ratio at tick
     */
    function _getSqrtRatioAtTick(int24 tick) internal pure returns (uint160) {
        // Simplified exponential approximation: price = 1.0001^tick
        // For prediction markets: map tick range to price range [0.01, 0.99]
        // MIN_TICK (-69000) ≈ 1% price, MAX_TICK (69000) ≈ 99% price
        
        uint256 absTick = tick < 0 ? uint256(-int256(tick)) : uint256(int256(tick));
        
        // Use exponential approximation: e^(tick * ln(1.0001))
        // Simplified: price ≈ 0.5 + (tick / 138000) for range [-69000, 69000] → [0.01, 0.99]
        int256 priceDeviation = (int256(tick) * 1e18) / 140000; // Range: [-0.49, 0.49]
        uint256 price = uint256(int256(5e17) + priceDeviation); // Range: [0.01e18, 0.99e18]
        
        // Clamp to valid range
        if (price < 1e16) price = 1e16; // Min 1%
        if (price > 99e16) price = 99e16; // Max 99%
        
        uint256 sqrtPrice = _sqrt(price);
        return uint160((sqrtPrice * Q96) / 1e9);
    }

    /**
     * @notice Get tick at sqrt ratio
     */
    function _getTickAtSqrtRatio(uint160 sqrtPriceX96) internal pure returns (int24) {
        // Inverse of _getSqrtRatioAtTick
        uint256 price = (uint256(sqrtPriceX96) * uint256(sqrtPriceX96) * 1e18) / (Q96 * Q96);
        // price = 0.5 + (tick / 140000)
        // tick = (price - 0.5) * 140000
        int256 deviation = int256(price) - int256(5e17);
        int256 tick = (deviation * 140000) / 1e18;
        
        // Clamp to valid tick range
        if (tick < MIN_TICK) return MIN_TICK;
        if (tick > MAX_TICK) return MAX_TICK;
        
        return int24(tick);
    }

    /**
     * @notice Square root function (Babylonian method)
     */
    function _sqrt(uint256 x) internal pure returns (uint256) {
        if (x == 0) return 0;
        uint256 z = (x + 1) / 2;
        uint256 y = x;
        while (z < y) {
            y = z;
            z = (x / z + z) / 2;
        }
        return y;
    }
}
