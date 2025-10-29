// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/markets/PooledLiquidityMarket.sol";
import "../../src/interfaces/IMarket.sol";
import "../../src/base/BaseMarket.sol";
import "../helpers/MarketTestHelper.sol";

contract PooledLiquidityMarketTest is MarketTestHelper {
    PooledLiquidityMarket public market;

    uint256 public constant MARKET_ID = 3;
    uint256 public constant OUTCOME_COUNT = 2; // Binary market
    uint256 public closeTime;

    // Tick constants
    int24 public constant MIN_TICK = -69000;
    int24 public constant MAX_TICK = 69000;
    int24 public constant TICK_SPACING = 10;

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

    function setUp() public {
        // Setup core contracts first
        setupCore();
        
        // Set close time
        closeTime = block.timestamp + 30 days;
        
        // Deploy market with core contracts
        market = new PooledLiquidityMarket(
            MARKET_ID,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            "Pooled LP Token",
            "PLP"
        );

        // Setup test environment
        registerMarket(MARKET_ID, creator);
        outcomeToken.setAMMAuthorization(address(market), true);

        // Fund users
        collateral.mint(trader1, DEFAULT_COLLATERAL_AMOUNT);
        collateral.mint(trader2, DEFAULT_COLLATERAL_AMOUNT);
        collateral.mint(trader3, DEFAULT_COLLATERAL_AMOUNT);

        // Approve market
        vm.prank(trader1);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(trader2);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(trader3);
        collateral.approve(address(market), type(uint256).max);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public {
        assertEq(market.marketId(), MARKET_ID);
        assertEq(address(market.collateralToken()), address(collateral));
        assertEq(market.closeTime(), closeTime);
        assertEq(market.getOutcomeCount(), 2);
        
        // Check initial pool state
        (uint160 sqrtPriceX96, int24 tick, uint128 liquidity,,) = market.poolState();
        assertEq(uint256(sqrtPriceX96), 56022770974670905984299832681); // sqrt(0.5) for tick 0
        assertEq(tick, 0);
        assertEq(liquidity, 0);
    }

    // ============ Position Minting Tests ============

    function test_MintPosition_FullRange() public {
        int24 tickLower = (MIN_TICK / TICK_SPACING) * TICK_SPACING;
        int24 tickUpper = (MAX_TICK / TICK_SPACING) * TICK_SPACING;
        uint128 liquidityDesired = 100_000e18;

        vm.startPrank(trader1);

        vm.expectEmit(true, true, true, false);
        emit PositionMinted(trader1, tickLower, tickUpper, liquidityDesired, 0, 0);

        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityDesired,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertEq(liquidity, liquidityDesired);
        assertGt(amount0, 0);
        assertGt(amount1, 0);

        // Check position was created
        bytes32 positionKey = keccak256(abi.encodePacked(trader1, tickLower, tickUpper));
        (uint128 posLiquidity,,,,) = market.positions(positionKey);
        assertEq(posLiquidity, liquidityDesired);
    }

    function test_MintPosition_NarrowRange() public {
        // Create position around current price (tick 0)
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityDesired,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertEq(liquidity, liquidityDesired);
        
        // Narrow range should require both tokens
        assertGt(amount0, 0);
        assertGt(amount1, 0);

        // Check active liquidity increased (position in range)
        (,, uint128 activeLiquidity,,) = market.poolState();
        assertEq(activeLiquidity, liquidityDesired);
    }

    function test_MintPosition_OutOfRange_Lower() public {
        // Position below current price (only token0)
        int24 tickLower = -1000;
        int24 tickUpper = -500;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityDesired,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertEq(liquidity, liquidityDesired);
        assertEq(amount0, 0); // No token0 needed (below current price)
        assertGt(amount1, 0); // Only token1 needed

        // Active liquidity should not increase (out of range)
        (,, uint128 activeLiquidity,,) = market.poolState();
        assertEq(activeLiquidity, 0);
    }

    function test_MintPosition_OutOfRange_Upper() public {
        // Position above current price (only token1)
        int24 tickLower = 500;
        int24 tickUpper = 1000;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityDesired,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertEq(liquidity, liquidityDesired);
        assertGt(amount0, 0); // Only token0 needed (above current price)
        assertEq(amount1, 0); // No token1 needed
    }

    function test_MintPosition_MultiplePositions_SameUser() public {
        int24 tickLower1 = -200;
        int24 tickUpper1 = 0;
        int24 tickLower2 = 0;
        int24 tickUpper2 = 200;
        uint128 liquidityDesired = 30_000e18;

        vm.startPrank(trader1);

        market.mintPosition(tickLower1, tickUpper1, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);
        market.mintPosition(tickLower2, tickUpper2, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.stopPrank();

        // Verify both positions exist
        bytes32 key1 = keccak256(abi.encodePacked(trader1, tickLower1, tickUpper1));
        bytes32 key2 = keccak256(abi.encodePacked(trader1, tickLower2, tickUpper2));
        
        (uint128 liq1,,,,) = market.positions(key1);
        (uint128 liq2,,,,) = market.positions(key2);
        
        assertEq(liq1, liquidityDesired);
        assertEq(liq2, liquidityDesired);
    }

    function test_RevertWhen_MintPosition_InvalidTickRange() public {
        vm.startPrank(trader1);

        // tickLower >= tickUpper
        vm.expectRevert(PooledLiquidityMarket.InvalidTickRange.selector);
        market.mintPosition(100, 100, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.expectRevert(PooledLiquidityMarket.InvalidTickRange.selector);
        market.mintPosition(200, 100, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.stopPrank();
    }

    function test_RevertWhen_MintPosition_InvalidTick() public {
        vm.startPrank(trader1);

        // Ticks out of bounds
        vm.expectRevert(PooledLiquidityMarket.InvalidTick.selector);
        market.mintPosition(MIN_TICK - 100, 0, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.expectRevert(PooledLiquidityMarket.InvalidTick.selector);
        market.mintPosition(0, MAX_TICK + 100, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.stopPrank();
    }

    function test_RevertWhen_MintPosition_InvalidTickSpacing() public {
        vm.startPrank(trader1);

        // Ticks not aligned to spacing
        vm.expectRevert(PooledLiquidityMarket.TickSpacingError.selector);
        market.mintPosition(-105, 105, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.expectRevert(PooledLiquidityMarket.TickSpacingError.selector);
        market.mintPosition(-100, 107, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.stopPrank();
    }

    function test_RevertWhen_MintPosition_SlippageExceeded() public {
        vm.startPrank(trader1);

        // Set max amounts too low
        vm.expectRevert(BaseMarket.SlippageExceeded.selector);
        market.mintPosition(-100, 100, 100_000e18, 1, 1);

        vm.stopPrank();
    }

    function test_RevertWhen_MintPosition_ZeroLiquidity() public {
        vm.startPrank(trader1);

        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.mintPosition(-100, 100, 0, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.stopPrank();
    }

    // ============ Position Burning Tests ============

    function test_BurnPosition() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        (uint128 liquidity,,) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityDesired,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.expectEmit(true, true, true, false);
        emit PositionBurned(trader1, tickLower, tickUpper, liquidity, 0, 0);

        (uint256 amount0, uint256 amount1) = market.burnPosition(
            tickLower,
            tickUpper,
            liquidity
        );

        vm.stopPrank();

        assertGt(amount0, 0);
        assertGt(amount1, 0);

        // Position should be zeroed
        bytes32 positionKey = keccak256(abi.encodePacked(trader1, tickLower, tickUpper));
        (uint128 posLiquidity,,,,) = market.positions(positionKey);
        assertEq(posLiquidity, 0);
    }

    function test_BurnPosition_Partial() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 100_000e18;

        vm.startPrank(trader1);

        market.mintPosition(tickLower, tickUpper, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        uint128 burnAmount = 30_000e18;
        market.burnPosition(tickLower, tickUpper, burnAmount);

        vm.stopPrank();

        // Check remaining liquidity
        bytes32 positionKey = keccak256(abi.encodePacked(trader1, tickLower, tickUpper));
        (uint128 posLiquidity,,,,) = market.positions(positionKey);
        assertEq(posLiquidity, liquidityDesired - burnAmount);
    }

    function test_RevertWhen_BurnPosition_InsufficientPosition() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        market.mintPosition(tickLower, tickUpper, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.expectRevert(PooledLiquidityMarket.InsufficientPosition.selector);
        market.burnPosition(tickLower, tickUpper, liquidityDesired + 1);

        vm.stopPrank();
    }

    function test_RevertWhen_BurnPosition_ZeroAmount() public {
        vm.startPrank(trader1);

        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.burnPosition(-100, 100, 0);

        vm.stopPrank();
    }

    // ============ Fee Collection Tests ============

    function test_CollectFees_AfterSwaps() public {
        // Use full range to ensure position stays in range during swaps
        int24 tickLower = market.MIN_TICK();
        int24 tickUpper = market.MAX_TICK();
        uint128 liquidityDesired = 100_000e18;

        // User1 provides liquidity
        vm.startPrank(trader1);
        market.mintPosition(tickLower, tickUpper, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);
        vm.stopPrank();

        // User2 performs smaller swaps to generate fees while keeping position in range
        vm.startPrank(trader2);
        market.buy(0, 100e18, 0);  // Even smaller swap amounts
        market.buy(1, 100e18, 0);
        vm.stopPrank();

        // User1 collects fees
        vm.startPrank(trader1);
        
        // Don't check exact amounts, just expect fees > 0
        (uint128 fees0, uint128 fees1) = market.collectFees(tickLower, tickUpper);
        
        vm.stopPrank();

        // Should have earned some fees
        assertGt(fees0, 0);
        assertGt(fees1, 0);
    }

    function test_CollectFees_NoFeesAccrued() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;

        vm.startPrank(trader1);
        market.mintPosition(tickLower, tickUpper, 50_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Collect immediately (no swaps yet)
        (uint128 fees0, uint128 fees1) = market.collectFees(tickLower, tickUpper);
        
        vm.stopPrank();

        assertEq(fees0, 0);
        assertEq(fees1, 0);
    }

    // ============ Swap Tests ============

    function test_Buy_Outcome0() public {
        // Setup liquidity
        int24 tickLower = -1000;
        int24 tickUpper = 1000;
        
        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        uint256 collateralIn = 10_000e18;
        
        vm.startPrank(trader2);
        
        uint256 balanceBefore = outcomeToken.balanceOfOutcome(trader2, MARKET_ID, 0);
        
        vm.expectEmit(true, true, false, false);
        emit Swap(trader2, 0, 0, 0, 0, 0);

        uint256 tokensOut = market.buy(0, collateralIn, 0);
        
        vm.stopPrank();

        assertGt(tokensOut, 0);
        
        uint256 balanceAfter = outcomeToken.balanceOfOutcome(trader2, MARKET_ID, 0);
        assertEq(balanceAfter - balanceBefore, tokensOut);
    }

    function test_Buy_Outcome1() public {
        // Setup liquidity
        int24 tickLower = -1000;
        int24 tickUpper = 1000;
        
        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        uint256 collateralIn = 10_000e18;
        
        vm.startPrank(trader2);
        uint256 tokensOut = market.buy(1, collateralIn, 0);
        vm.stopPrank();

        assertGt(tokensOut, 0);
        assertEq(outcomeToken.balanceOfOutcome(trader2, MARKET_ID, 1), tokensOut);
    }

    function test_Sell_Outcome0() public {
        // Setup liquidity
        int24 tickLower = -1000;
        int24 tickUpper = 1000;
        
        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // First buy some tokens
        vm.startPrank(trader2);
        uint256 tokensBought = market.buy(0, 10_000e18, 0);
        
        // Approve market to burn tokens
        outcomeToken.setApprovalForAll(address(market), true);
        
        // Then sell them
        uint256 collateralOut = market.sell(0, tokensBought, 0);
        vm.stopPrank();

        assertGt(collateralOut, 0);
    }

    function test_Sell_Outcome1() public {
        // Setup liquidity
        int24 tickLower = -1000;
        int24 tickUpper = 1000;
        
        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.startPrank(trader2);
        uint256 tokensBought = market.buy(1, 10_000e18, 0);
        
        outcomeToken.setApprovalForAll(address(market), true);
        uint256 collateralOut = market.sell(1, tokensBought, 0);
        vm.stopPrank();

        assertGt(collateralOut, 0);
    }

    function test_PriceMovement_AfterBuys() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        uint256 priceBefore = market.getPrice(0);

        // Buy outcome 0 - should increase its price
        vm.prank(trader2);
        market.buy(0, 10_000e18, 0);

        uint256 priceAfter = market.getPrice(0);

        assertGt(priceAfter, priceBefore);
    }

    function test_RevertWhen_Buy_NoLiquidity() public {
        vm.startPrank(trader2);

        vm.expectRevert(BaseMarket.InsufficientLiquidity.selector);
        market.buy(0, 10_000e18, 0);

        vm.stopPrank();
    }

    function test_RevertWhen_Buy_SlippageExceeded() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.startPrank(trader2);

        // Set minTokensOut too high
        vm.expectRevert(BaseMarket.SlippageExceeded.selector);
        market.buy(0, 10_000e18, 1_000_000e18);

        vm.stopPrank();
    }

    function test_RevertWhen_Sell_SlippageExceeded() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.startPrank(trader2);
        uint256 tokensBought = market.buy(0, 10_000e18, 0);
        
        outcomeToken.setApprovalForAll(address(market), true);

        // Set minCollateralOut too high
        vm.expectRevert(BaseMarket.SlippageExceeded.selector);
        market.sell(0, tokensBought, 1_000_000e18);

        vm.stopPrank();
    }

    // ============ Quote Tests ============

    function test_GetQuoteBuy() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        (uint256 tokensOut, uint256 fee) = market.getQuoteBuy(0, 10_000e18, trader2);

        assertGt(tokensOut, 0);
        assertGt(fee, 0);
    }

    function test_GetQuoteSell() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        (uint256 collateralOut, uint256 fee) = market.getQuoteSell(0, 10_000e18, trader2);

        assertGt(collateralOut, 0);
        assertGe(fee, 0);
    }

    // ============ Simplified Liquidity Functions Tests ============

    function test_AddLiquidity() public {
        vm.startPrank(trader1);

        uint256 amount = 100_000e18;
        uint256 lpTokens = market.addLiquidity(amount);

        vm.stopPrank();

        assertGt(lpTokens, 0);
        assertEq(market.balanceOf(trader1), lpTokens);
    }

    function test_RemoveLiquidity() public {
        vm.startPrank(trader1);

        uint256 amount = 100_000e18;
        uint256 lpTokens = market.addLiquidity(amount);

        uint256 balanceBefore = collateral.balanceOf(trader1);
        uint256 collateralOut = market.removeLiquidity(lpTokens);
        uint256 balanceAfter = collateral.balanceOf(trader1);

        vm.stopPrank();

        assertGt(collateralOut, 0);
        assertEq(balanceAfter - balanceBefore, collateralOut);
        assertEq(market.balanceOf(trader1), 0);
    }

    function test_RevertWhen_RemoveLiquidity_InsufficientLPTokens() public {
        vm.startPrank(trader1);

        vm.expectRevert(BaseMarket.InsufficientLPTokens.selector);
        market.removeLiquidity(100_000e18);

        vm.stopPrank();
    }

    // ============ Admin Tests ============

    function test_Pause() public {
        market.pause();
        assertTrue(market.paused());
    }

    function test_Unpause() public {
        market.pause();
        market.unpause();
        assertFalse(market.paused());
    }

    function test_RevertWhen_MintPosition_WhenPaused() public {
        market.pause();

        vm.startPrank(trader1);
        vm.expectRevert();
        market.mintPosition(-100, 100, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);
        vm.stopPrank();
    }

    function test_RevertWhen_Buy_WhenPaused() public {
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        market.pause();

        vm.startPrank(trader2);
        vm.expectRevert();
        market.buy(0, 10_000e18, 0);
        vm.stopPrank();
    }

    // ============ Gas Profiling Tests ============

    function test_Gas_MintPosition() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        uint256 gasBefore = gasleft();
        market.mintPosition(tickLower, tickUpper, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);
        uint256 gasUsed = gasBefore - gasleft();

        vm.stopPrank();

        console.log("Gas used for minting position:", gasUsed);
    }

    function test_Gas_BurnPosition() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityDesired = 50_000e18;

        vm.startPrank(trader1);

        market.mintPosition(tickLower, tickUpper, liquidityDesired, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        uint256 gasBefore = gasleft();
        market.burnPosition(tickLower, tickUpper, liquidityDesired);
        uint256 gasUsed = gasBefore - gasleft();

        vm.stopPrank();

        console.log("Gas used for burning position:", gasUsed);
    }

    function test_Gas_Swap() public {
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.startPrank(trader2);

        uint256 gasBefore = gasleft();
        market.buy(0, 10_000e18, 0);
        uint256 gasUsed = gasBefore - gasleft();

        vm.stopPrank();

        console.log("Gas used for swap:", gasUsed);
    }

    function test_Gas_CollectFees() public {
        int24 tickLower = -200;
        int24 tickUpper = 200;

        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        vm.prank(trader2);
        market.buy(0, 10_000e18, 0);

        vm.startPrank(trader1);

        uint256 gasBefore = gasleft();
        market.collectFees(tickLower, tickUpper);
        uint256 gasUsed = gasBefore - gasleft();

        vm.stopPrank();

        console.log("Gas used for collecting fees:", gasUsed);
    }

    // ============ Edge Case Tests ============

    function test_EdgeCase_ExtremePriceMovement_NearMinTick() public {
        // Setup: Add liquidity at full range
        vm.prank(trader1);
        market.mintPosition(MIN_TICK, MAX_TICK, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Perform large series of swaps to push price toward minimum
        vm.startPrank(trader2);
        for (uint256 i = 0; i < 5; i++) {
            market.buy(1, 10_000e18, 0); // Buy outcome 1 to push price of outcome 0 down
        }
        vm.stopPrank();

        // Verify market still functions
        (uint160 sqrtPriceX96, int24 tick,,, ) = market.poolState();
        assertGt(sqrtPriceX96, 0, "Price should be valid");
        assertGe(tick, MIN_TICK, "Tick should not go below MIN_TICK");
        
        // Verify we can still trade
        vm.prank(trader3);
        uint256 tokensOut = market.buy(0, 1000e18, 0);
        assertGt(tokensOut, 0, "Should still be able to buy");
    }

    function test_EdgeCase_ExtremePriceMovement_NearMaxTick() public {
        // Setup: Add liquidity at full range
        vm.prank(trader1);
        market.mintPosition(MIN_TICK, MAX_TICK, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Perform large series of swaps to push price toward maximum
        vm.startPrank(trader2);
        for (uint256 i = 0; i < 5; i++) {
            market.buy(0, 10_000e18, 0); // Buy outcome 0 to push price up
        }
        vm.stopPrank();

        // Verify market still functions
        (uint160 sqrtPriceX96, int24 tick,,, ) = market.poolState();
        assertGt(sqrtPriceX96, 0, "Price should be valid");
        assertLe(tick, MAX_TICK, "Tick should not exceed MAX_TICK");
        
        // Verify we can still trade
        vm.prank(trader3);
        uint256 tokensOut = market.buy(1, 1000e18, 0);
        assertGt(tokensOut, 0, "Should still be able to buy");
    }

    function test_EdgeCase_BoundaryTick_NearMinTick() public {
        // Test position near minimum boundary
        int24 tickLower = MIN_TICK + (TICK_SPACING * 100); // Start 100 ticks above MIN
        int24 tickUpper = MIN_TICK + (TICK_SPACING * 200); // 100 ticks range

        vm.startPrank(trader1);
        
        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            100_000e18, // Larger amount for extreme tick range
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertGt(liquidity, 0, "Should mint liquidity near MIN_TICK");
        // Position near MIN_TICK requires tokens based on whether it's in range
        assertTrue(amount0 > 0 || amount1 > 0, "Should require at least one token");
    }

    function test_EdgeCase_BoundaryTick_NearMaxTick() public {
        // Test position near maximum boundary
        int24 tickLower = MAX_TICK - (TICK_SPACING * 200); // 100 ticks range
        int24 tickUpper = MAX_TICK - (TICK_SPACING * 100); // End 100 ticks below MAX

        vm.startPrank(trader1);
        
        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            100_000e18, // Larger amount for extreme tick range
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertGt(liquidity, 0, "Should mint liquidity near MAX_TICK");
        // Position near MAX_TICK requires tokens based on whether it's in range
        assertTrue(amount0 > 0 || amount1 > 0, "Should require at least one token");
    }

    function test_EdgeCase_BoundaryTick_FullRangeExact() public {
        // Test full range position with exact boundaries
        int24 tickLower = MIN_TICK;
        int24 tickUpper = MAX_TICK;

        vm.startPrank(trader1);
        
        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            tickLower,
            tickUpper,
            50_000e18,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertGt(liquidity, 0, "Should mint liquidity");
        assertGt(amount0, 0, "Should require amount0");
        assertGt(amount1, 0, "Should require amount1");
    }

    function test_EdgeCase_MinimumLiquidityAmount() public {
        // Test with minimal liquidity amount (1 wei)
        vm.startPrank(trader1);
        
        (uint128 liquidity, uint256 amount0, uint256 amount1) = market.mintPosition(
            -100,
            100,
            1, // Minimum possible liquidity
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        vm.stopPrank();

        assertEq(liquidity, 1, "Should mint exactly 1 wei of liquidity");
        // With 1 wei liquidity, amounts should be 0 or very small
        assertTrue(amount0 == 0 || amount1 == 0, "One amount should be 0 with minimal liquidity");
    }

    function test_EdgeCase_MultiplePositionsSameRange() public {
        // Test multiple users with positions in the exact same range
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityAmount = 10_000e18;

        // User 1 creates position
        vm.prank(trader1);
        (uint128 liquidity1,,) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityAmount,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        // User 2 creates position in same range
        vm.prank(trader2);
        (uint128 liquidity2,,) = market.mintPosition(
            tickLower,
            tickUpper,
            liquidityAmount,
            DEFAULT_COLLATERAL_AMOUNT,
            DEFAULT_COLLATERAL_AMOUNT
        );

        assertEq(liquidity1, liquidity2, "Should mint same liquidity for same range");

        // Perform swap to generate fees
        vm.prank(trader3);
        market.buy(0, 5_000e18, 0);

        // Both should be able to collect fees independently
        vm.prank(trader1);
        (uint128 fees1_0, uint128 fees1_1) = market.collectFees(tickLower, tickUpper);

        vm.prank(trader2);
        (uint128 fees2_0, uint128 fees2_1) = market.collectFees(tickLower, tickUpper);

        // Both should have collected fees
        assertTrue(fees1_0 > 0 || fees1_1 > 0, "Trader1 should collect fees");
        assertTrue(fees2_0 > 0 || fees2_1 > 0, "Trader2 should collect fees");
    }

    function test_EdgeCase_PartialBurnWithAccruedFees() public {
        int24 tickLower = -100;
        int24 tickUpper = 100;
        uint128 liquidityAmount = 20_000e18;

        // Create position
        vm.prank(trader1);
        market.mintPosition(tickLower, tickUpper, liquidityAmount, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Generate fees
        vm.prank(trader2);
        market.buy(0, 10_000e18, 0);

        // Burn only half the position
        vm.prank(trader1);
        (uint256 amount0, uint256 amount1) = market.burnPosition(tickLower, tickUpper, liquidityAmount / 2);

        assertGt(amount0 + amount1, 0, "Should return liquidity");

        // Should still be able to collect fees on remaining position
        vm.prank(trader1);
        (uint128 fees0, uint128 fees1) = market.collectFees(tickLower, tickUpper);

        assertTrue(fees0 > 0 || fees1 > 0, "Should still collect fees on remaining liquidity");
    }

    function test_EdgeCase_ZeroReservesAtInit() public {
        // Verify initial reserves are zero
        assertEq(market.reserves(0), 0, "Reserve 0 should start at 0");
        assertEq(market.reserves(1), 0, "Reserve 1 should start at 0");

        // First liquidity addition should initialize reserves
        vm.prank(trader1);
        market.mintPosition(-100, 100, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        assertGt(market.reserves(0), 0, "Reserve 0 should be initialized");
        assertGt(market.reserves(1), 0, "Reserve 1 should be initialized");
    }

    function test_EdgeCase_VerySmallSwapAmount() public {
        // Setup liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 100_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Perform very small swap (1 wei)
        vm.prank(trader2);
        uint256 tokensOut = market.buy(0, 1, 0); // Buy with 1 wei

        // Should handle small amounts gracefully
        assertTrue(tokensOut == 0 || tokensOut == 1, "Small swap should return 0 or 1 wei");
    }

    function test_EdgeCase_LargeSwapRelativeToLiquidity() public {
        // Setup limited liquidity
        vm.prank(trader1);
        market.mintPosition(-1000, 1000, 10_000e18, DEFAULT_COLLATERAL_AMOUNT, DEFAULT_COLLATERAL_AMOUNT);

        // Attempt large swap (should not drain pool completely)
        vm.prank(trader2);
        uint256 tokensOut = market.buy(0, 50_000e18, 0);

        assertGt(tokensOut, 0, "Should execute swap");
        
        // Reserves should never hit zero
        assertGt(market.reserves(0), 0, "Reserve 0 should not be zero");
        assertGt(market.reserves(1), 0, "Reserve 1 should not be zero");
    }
}

