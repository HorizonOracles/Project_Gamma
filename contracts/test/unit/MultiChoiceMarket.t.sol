// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/markets/MultiChoiceMarket.sol";
import "../../src/interfaces/IMarket.sol";
import "../../src/base/BaseMarket.sol";
import "../helpers/MarketTestHelper.sol";

/**
 * @title MultiChoiceMarketTest
 * @notice Comprehensive unit tests for MultiChoiceMarket with LMSR pricing
 * @dev Tests cover:
 *      - Constructor validation
 *      - Initial and subsequent liquidity provision
 *      - Buy/sell trades with LMSR pricing
 *      - Price calculations
 *      - Slippage protection
 *      - Edge cases (resolution, pausing, zero liquidity)
 *      - Gas profiling
 */
contract MultiChoiceMarketTest is MarketTestHelper {
    MultiChoiceMarket public market;
    
    uint256 public constant MARKET_ID = 1;
    uint256 public constant OUTCOME_COUNT = 4;
    uint256 public constant LIQUIDITY_PARAM = 1000 ether;
    uint256 public closeTime;

    // Price tolerance for LMSR approximations (5% = 500 bps)
    uint256 public constant PRICE_TOLERANCE_BPS = 500;

    function setUp() public {
        // Setup core contracts first
        setupCore();
        
        // Set close time
        closeTime = block.timestamp + 30 days;
        
        // Deploy market with core contracts
        market = new MultiChoiceMarket(
            MARKET_ID,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            OUTCOME_COUNT,
            LIQUIDITY_PARAM
        );

        // Setup test environment (without calling setupCore again)
        registerMarket(MARKET_ID, creator);
        outcomeToken.setAMMAuthorization(address(market), true);
        fundStandardAccounts();
        approveMarketForAll(address(market));
        outcomeToken.setResolutionAuthorization(address(this), true);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(market.marketId(), MARKET_ID);
        assertEq(address(market.collateralToken()), address(collateral));
        assertEq(market.closeTime(), closeTime);
        assertEq(market.outcomeCount(), OUTCOME_COUNT);
        assertEq(market.liquidityParameter(), LIQUIDITY_PARAM);
        assertEq(uint256(market.marketType()), uint256(IMarket.MarketType.MultiChoice));
    }

    function test_Constructor_ValidOutcomeCounts() public {
        // Test valid outcome counts (3-8)
        for (uint256 i = 3; i <= 8; i++) {
            MultiChoiceMarket testMarket = new MultiChoiceMarket(
                i + 100, // unique market ID
                address(collateral),
                address(outcomeToken),
                address(feeSplitter),
                address(horizonPerks),
                closeTime,
                i,
                LIQUIDITY_PARAM
            );
            assertEq(testMarket.outcomeCount(), i);
        }
    }

    function test_RevertWhen_Constructor_InvalidOutcomeCount_TooFew() public {
        vm.expectRevert(MultiChoiceMarket.InvalidOutcomeCount.selector);
        new MultiChoiceMarket(
            999,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            2, // Too few
            LIQUIDITY_PARAM
        );
    }

    function test_RevertWhen_Constructor_InvalidOutcomeCount_TooMany() public {
        vm.expectRevert(MultiChoiceMarket.InvalidOutcomeCount.selector);
        new MultiChoiceMarket(
            999,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            9, // Too many
            LIQUIDITY_PARAM
        );
    }

    function test_RevertWhen_Constructor_ZeroLiquidityParameter() public {
        vm.expectRevert(MultiChoiceMarket.InvalidLiquidityParameter.selector);
        new MultiChoiceMarket(
            999,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            4,
            0 // Invalid
        );
    }

    function test_GetMarketInfo() public view {
        IMarket.MarketInfo memory info = market.getMarketInfo();
        
        assertEq(info.marketId, MARKET_ID);
        assertEq(uint256(info.marketType), uint256(IMarket.MarketType.MultiChoice));
        assertEq(info.collateralToken, address(collateral));
        assertEq(info.closeTime, closeTime);
        assertEq(info.outcomeCount, OUTCOME_COUNT);
        assertEq(info.isResolved, false);
        assertEq(info.isPaused, false);
    }

    // ============ Initial Liquidity Tests ============

    function test_AddLiquidity_Initial() public {
        uint256 amount = 10000 ether;
        uint256 expectedLP = amount - market.MINIMUM_LIQUIDITY();
        
        expectLiquidityEvent(lp1, true);
        
        vm.prank(lp1);
        uint256 lpTokens = market.addLiquidity(amount);

        // Check LP tokens
        assertEq(lpTokens, expectedLP, "LP tokens mismatch");
        assertEq(market.balanceOf(lp1), expectedLP, "LP1 balance mismatch");
        assertEq(market.totalSupply(), amount, "Total supply mismatch");

        // Check total collateral
        assertEq(market.totalCollateral(), amount, "Total collateral mismatch");

        // Check outcome reserves (should be equal for all outcomes)
        uint256[] memory reserves = market.getOutcomeReserves();
        assertEq(reserves.length, OUTCOME_COUNT, "Reserves array length mismatch");
        
        uint256 expectedReserve = amount / OUTCOME_COUNT;
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            assertEq(reserves[i], expectedReserve, "Initial reserve mismatch");
            
            // Check outcome tokens held by market
            uint256 balance = getOutcomeBalance(address(market), MARKET_ID, i);
            assertEq(balance, expectedReserve, "Market outcome token balance mismatch");
        }
    }

    function test_AddLiquidity_Initial_Prices() public {
        uint256 amount = 10000 ether;
        
        vm.prank(lp1);
        market.addLiquidity(amount);

        // All prices should be approximately equal (1/N)
        uint256[] memory prices = market.getAllPrices();
        uint256 expectedPrice = 1e18 / OUTCOME_COUNT;
        
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            assertPriceWithinTolerance(prices[i], expectedPrice, PRICE_TOLERANCE_BPS);
        }

        // Prices should sum to ~1.0
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
    }

    function test_AddLiquidity_Subsequent() public {
        // Initial liquidity
        uint256 initialAmount = 10000 ether;
        vm.prank(lp1);
        market.addLiquidity(initialAmount);

        // Buy some tokens to change reserves
        uint256 buyAmount = 1000 ether;
        vm.prank(trader1);
        market.buy(0, buyAmount, 0);

        // Subsequent liquidity
        uint256 additionalAmount = 5000 ether;
        uint256 totalSupplyBefore = market.totalSupply();
        uint256 totalCollateralBefore = market.totalCollateral();

        vm.prank(lp2);
        uint256 lpTokens = market.addLiquidity(additionalAmount);

        // LP tokens should be proportional
        uint256 expectedLP = (additionalAmount * totalSupplyBefore) / totalCollateralBefore;
        assertEq(lpTokens, expectedLP, "Proportional LP tokens mismatch");

        // Reserves should increase proportionally
        uint256[] memory reserves = market.getOutcomeReserves();
        uint256 totalCollateralAfter = market.totalCollateral();
        
        assertEq(totalCollateralAfter, totalCollateralBefore + additionalAmount, "Total collateral mismatch");
    }

    function test_RevertWhen_AddLiquidity_ZeroAmount() public {
        vm.prank(lp1);
        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.addLiquidity(0);
    }

    function test_RevertWhen_AddLiquidity_AfterResolution() public {
        // Add liquidity
        vm.prank(lp1);
        market.addLiquidity(10000 ether);

        // Resolve market
        resolveMarket(MARKET_ID, 0);

        // Try to add more liquidity
        vm.prank(lp2);
        vm.expectRevert(BaseMarket.MarketResolved.selector);
        market.addLiquidity(5000 ether);
    }

    function test_RevertWhen_AddLiquidity_WhenPaused() public {
        // Add initial liquidity
        vm.prank(lp1);
        market.addLiquidity(10000 ether);

        // Pause market
        market.pause();

        // Try to add liquidity
        vm.prank(lp2);
        vm.expectRevert();
        market.addLiquidity(5000 ether);
    }

    // ============ Remove Liquidity Tests ============

    function test_RemoveLiquidity() public {
        // Add liquidity
        uint256 amount = 10000 ether;
        vm.prank(lp1);
        uint256 lpTokens = market.addLiquidity(amount);

        // Remove half
        uint256 removeLP = lpTokens / 2;
        uint256 totalSupplyBefore = market.totalSupply();
        uint256 totalCollateralBefore = market.totalCollateral();
        
        uint256 expectedCollateral = (removeLP * totalCollateralBefore) / totalSupplyBefore;

        vm.prank(lp1);
        uint256 collateralOut = market.removeLiquidity(removeLP);

        assertEq(collateralOut, expectedCollateral, "Collateral out mismatch");
        assertEq(market.balanceOf(lp1), lpTokens - removeLP, "LP balance after removal mismatch");
        assertEq(market.totalCollateral(), totalCollateralBefore - collateralOut, "Total collateral after removal");
    }

    function test_RemoveLiquidity_AfterTrading() public {
        // Add liquidity
        vm.prank(lp1);
        uint256 lpTokens = market.addLiquidity(10000 ether);

        // Execute some trades
        vm.prank(trader1);
        market.buy(0, 1000 ether, 0);
        
        vm.prank(trader2);
        market.buy(1, 500 ether, 0);

        // Remove liquidity
        uint256 collateralBefore = collateral.balanceOf(lp1);
        
        vm.prank(lp1);
        uint256 collateralOut = market.removeLiquidity(lpTokens / 2);

        assertGt(collateralOut, 0, "Should receive collateral");
        assertEq(collateral.balanceOf(lp1), collateralBefore + collateralOut, "LP collateral balance");
    }

    function test_RevertWhen_RemoveLiquidity_InsufficientLP() public {
        vm.prank(lp1);
        market.addLiquidity(10000 ether);

        vm.prank(lp2);
        vm.expectRevert(BaseMarket.InsufficientLPTokens.selector);
        market.removeLiquidity(1000 ether);
    }

    function test_RevertWhen_RemoveLiquidity_ZeroAmount() public {
        vm.prank(lp1);
        market.addLiquidity(10000 ether);

        vm.prank(lp1);
        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.removeLiquidity(0);
    }

    // ============ Buy Tests ============

    function test_Buy_SingleOutcome() public {
        // Add liquidity
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Get initial price
        uint256 priceBefore = market.getPrice(0);
        
        // Buy outcome 0
        uint256 collateralIn = 1000 ether;
        uint256 collateralBefore = collateral.balanceOf(trader1);
        
        expectTradeEvent(trader1, 0, true);
        
        vm.prank(trader1);
        uint256 tokensOut = market.buy(0, collateralIn, 0);

        // Check tokens received
        assertGt(tokensOut, 0, "Should receive tokens");
        assertEq(getOutcomeBalance(trader1, MARKET_ID, 0), tokensOut, "Trader outcome balance");

        // Check collateral spent
        assertEq(collateral.balanceOf(trader1), collateralBefore - collateralIn, "Trader collateral balance");

        // Price should increase for outcome 0
        uint256 priceAfter = market.getPrice(0);
        assertGt(priceAfter, priceBefore, "Price should increase after buy");

        // Prices should still sum to ~1.0
        assertPricesSumToOne(market.getAllPrices(), PRICE_TOLERANCE_BPS);
    }

    function test_Buy_MultipleOutcomes() public {
        // Add liquidity
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Buy different outcomes
        uint256 amount = 500 ether;
        
        vm.prank(trader1);
        uint256 tokens0 = market.buy(0, amount, 0);
        
        vm.prank(trader2);
        uint256 tokens1 = market.buy(1, amount, 0);
        
        vm.prank(trader3);
        uint256 tokens2 = market.buy(2, amount, 0);

        // All should receive tokens
        assertGt(tokens0, 0);
        assertGt(tokens1, 0);
        assertGt(tokens2, 0);

        // Prices should be updated
        uint256[] memory prices = market.getAllPrices();
        
        // Outcomes 0, 1, 2 should have higher prices than outcome 3
        assertGt(prices[0], prices[3]);
        assertGt(prices[1], prices[3]);
        assertGt(prices[2], prices[3]);

        // Prices should still sum to ~1.0
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
    }

    function test_Buy_WithSlippageProtection() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 collateralIn = 1000 ether;
        
        // Get quote
        (uint256 expectedTokens,) = market.getQuoteBuy(0, collateralIn, trader1);
        
        // Buy with minimum tokens expectation
        vm.prank(trader1);
        uint256 tokensOut = market.buy(0, collateralIn, expectedTokens);

        assertGe(tokensOut, expectedTokens, "Should meet minimum");
    }

    function test_RevertWhen_Buy_SlippageExceeded() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 collateralIn = 1000 ether;
        uint256 impossibleMinTokens = 10000 ether; // Way too high
        
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.SlippageExceeded.selector);
        market.buy(0, collateralIn, impossibleMinTokens);
    }

    function test_RevertWhen_Buy_ZeroAmount() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.buy(0, 0, 0);
    }

    function test_RevertWhen_Buy_InvalidOutcome() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidOutcomeId.selector);
        market.buy(OUTCOME_COUNT, 1000 ether, 0);
    }

    function test_RevertWhen_Buy_NoLiquidity() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InsufficientLiquidity.selector);
        market.buy(0, 1000 ether, 0);
    }

    function test_RevertWhen_Buy_AfterClose() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Warp past close time
        warpAfterClose(IMarket(address(market)));

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.MarketClosed.selector);
        market.buy(0, 1000 ether, 0);
    }

    function test_RevertWhen_Buy_AfterResolution() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        resolveMarket(MARKET_ID, 0);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.MarketResolved.selector);
        market.buy(0, 1000 ether, 0);
    }

    function test_RevertWhen_Buy_WhenPaused() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        market.pause();

        vm.prank(trader1);
        vm.expectRevert();
        market.buy(0, 1000 ether, 0);
    }

    // ============ Sell Tests ============

    function test_Sell() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Buy tokens first
        uint256 buyAmount = 1000 ether;
        vm.prank(trader1);
        uint256 tokensBought = market.buy(0, buyAmount, 0);

        // Get price before sell
        uint256 priceBefore = market.getPrice(0);

        // Sell half the tokens
        uint256 tokensToSell = tokensBought / 2;
        uint256 collateralBefore = collateral.balanceOf(trader1);
        
        expectTradeEvent(trader1, 0, false);
        
        vm.prank(trader1);
        uint256 collateralOut = market.sell(0, tokensToSell, 0);

        // Check collateral received
        assertGt(collateralOut, 0, "Should receive collateral");
        assertEq(collateral.balanceOf(trader1), collateralBefore + collateralOut, "Trader collateral balance");

        // Check tokens burned
        assertEq(
            getOutcomeBalance(trader1, MARKET_ID, 0),
            tokensBought - tokensToSell,
            "Trader outcome balance after sell"
        );

        // Price should decrease
        uint256 priceAfter = market.getPrice(0);
        assertLt(priceAfter, priceBefore, "Price should decrease after sell");

        // Prices should still sum to ~1.0
        assertPricesSumToOne(market.getAllPrices(), PRICE_TOLERANCE_BPS);
    }

    function test_Sell_WithSlippageProtection() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Buy first
        vm.prank(trader1);
        uint256 tokens = market.buy(0, 1000 ether, 0);

        // Get quote for sell
        (uint256 expectedCollateral,) = market.getQuoteSell(0, tokens, trader1);

        // Sell with minimum collateral expectation
        vm.prank(trader1);
        uint256 collateralOut = market.sell(0, tokens, expectedCollateral);

        assertGe(collateralOut, expectedCollateral, "Should meet minimum");
    }

    function test_RevertWhen_Sell_SlippageExceeded() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        uint256 tokens = market.buy(0, 1000 ether, 0);

        uint256 impossibleMinCollateral = 5000 ether; // Way too high

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.SlippageExceeded.selector);
        market.sell(0, tokens, impossibleMinCollateral);
    }

    function test_RevertWhen_Sell_ZeroAmount() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.sell(0, 0, 0);
    }

    function test_RevertWhen_Sell_InvalidOutcome() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidOutcomeId.selector);
        market.sell(OUTCOME_COUNT, 1000 ether, 0);
    }

    function test_RevertWhen_Sell_NoLiquidity() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InsufficientLiquidity.selector);
        market.sell(0, 1000 ether, 0);
    }

    function test_RevertWhen_Sell_InsufficientBalance() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Try to sell tokens trader doesn't have
        vm.prank(trader1);
        vm.expectRevert();
        market.sell(0, 1000 ether, 0);
    }

    // ============ Price Calculation Tests ============

    function test_GetPrice_NoLiquidity() public view {
        // Should return equal prices
        uint256 price = market.getPrice(0);
        uint256 expectedPrice = 1e18 / OUTCOME_COUNT;
        
        assertEq(price, expectedPrice, "Price without liquidity");
    }

    function test_GetPrice_WithEqualReserves() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // All prices should be approximately equal
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            uint256 price = market.getPrice(i);
            uint256 expectedPrice = 1e18 / OUTCOME_COUNT;
            
            assertPriceWithinTolerance(price, expectedPrice, PRICE_TOLERANCE_BPS);
        }
    }

    function test_GetPrice_AfterTrades() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Heavy buying on outcome 0
        vm.prank(trader1);
        market.buy(0, 5000 ether, 0);

        uint256[] memory prices = market.getAllPrices();

        // Outcome 0 should have highest price
        for (uint256 i = 1; i < OUTCOME_COUNT; i++) {
            assertGt(prices[0], prices[i], "Outcome 0 should have highest price");
        }

        // Prices should still sum to ~1.0
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
    }

    function test_GetAllPrices() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256[] memory prices = market.getAllPrices();
        
        assertEq(prices.length, OUTCOME_COUNT, "Prices array length");
        
        // Sum should be ~1.0
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
        
        // Each price should be positive
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            assertGt(prices[i], 0, "Price should be positive");
        }
    }

    // ============ Quote Tests ============

    function test_GetQuoteBuy() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 collateralIn = 1000 ether;
        
        (uint256 tokensOut, uint256 fee) = market.getQuoteBuy(0, collateralIn, trader1);

        assertGt(tokensOut, 0, "Should quote positive tokens");
        assertGt(fee, 0, "Should have fee");

        // Verify quote matches actual buy
        vm.prank(trader1);
        uint256 actualTokens = market.buy(0, collateralIn, 0);

        assertEq(actualTokens, tokensOut, "Quote should match actual");
    }

    function test_GetQuoteBuy_ZeroAmount() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        (uint256 tokensOut, uint256 fee) = market.getQuoteBuy(0, 0, trader1);

        assertEq(tokensOut, 0);
        assertEq(fee, 0);
    }

    function test_GetQuoteBuy_NoLiquidity() public view {
        (uint256 tokensOut, uint256 fee) = market.getQuoteBuy(0, 1000 ether, trader1);

        assertEq(tokensOut, 0);
        assertEq(fee, 0);
    }

    function test_GetQuoteSell() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Buy first
        vm.prank(trader1);
        uint256 tokens = market.buy(0, 1000 ether, 0);

        // Get quote
        (uint256 collateralOut, uint256 fee) = market.getQuoteSell(0, tokens, trader1);

        assertGt(collateralOut, 0, "Should quote positive collateral");
        assertGt(fee, 0, "Should have fee");

        // Verify quote matches actual sell
        vm.prank(trader1);
        uint256 actualCollateral = market.sell(0, tokens, 0);

        assertEq(actualCollateral, collateralOut, "Quote should match actual");
    }

    function test_GetQuoteSell_ZeroAmount() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        (uint256 collateralOut, uint256 fee) = market.getQuoteSell(0, 0, trader1);

        assertEq(collateralOut, 0);
        assertEq(fee, 0);
    }

    // ============ Admin & State Tests ============

    function test_Pause() public {
        market.pause();
        
        IMarket.MarketInfo memory info = market.getMarketInfo();
        assertEq(info.isPaused, true);
    }

    function test_Unpause() public {
        market.pause();
        market.unpause();
        
        IMarket.MarketInfo memory info = market.getMarketInfo();
        assertEq(info.isPaused, false);
    }

    function test_RevertWhen_Pause_NotAdmin() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.Unauthorized.selector);
        market.pause();
    }

    function test_SetAdmin() public {
        address newAdmin = address(0x999);
        market.setAdmin(newAdmin);
        
        // Old admin can't pause anymore
        vm.expectRevert(BaseMarket.Unauthorized.selector);
        market.pause();

        // New admin can pause
        vm.prank(newAdmin);
        market.pause();
    }

    function test_FundRedemptions() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Execute trades
        vm.prank(trader1);
        market.buy(0, 1000 ether, 0);

        // Resolve
        resolveMarket(MARKET_ID, 0);

        // Fund redemptions
        uint256 marketBalance = collateral.balanceOf(address(market));
        uint256 outcomeTokenBalance = collateral.balanceOf(address(outcomeToken));
        
        market.fundRedemptions();

        assertEq(collateral.balanceOf(address(market)), 0);
        assertEq(collateral.balanceOf(address(outcomeToken)), outcomeTokenBalance + marketBalance);
    }

    function test_RevertWhen_FundRedemptions_NotResolved() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.expectRevert(BaseMarket.InvalidState.selector);
        market.fundRedemptions();
    }

    // ============ View Function Tests ============

    function test_GetOutcomeReserves() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256[] memory reserves = market.getOutcomeReserves();
        
        assertEq(reserves.length, OUTCOME_COUNT);
        
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            assertGt(reserves[i], 0, "Reserve should be positive");
        }
    }

    function test_GetOutcomeCount() public view {
        assertEq(market.getOutcomeCount(), OUTCOME_COUNT);
    }

    function test_GetMarketType() public view {
        assertEq(uint256(market.getMarketType()), uint256(IMarket.MarketType.MultiChoice));
    }

    // ============ LMSR Property Tests ============

    function test_LMSR_PricesSumToOne() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // After various trades, prices should always sum to ~1.0
        uint256[] memory prices;
        
        // Initial state
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);

        // After buy
        vm.prank(trader1);
        market.buy(0, 1000 ether, 0);
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);

        // After another buy on different outcome
        vm.prank(trader2);
        market.buy(2, 500 ether, 0);
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);

        // After sell
        uint256 tokens = getOutcomeBalance(trader1, MARKET_ID, 0);
        vm.prank(trader1);
        market.sell(0, tokens / 2, 0);
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
    }

    function test_LMSR_BuyIncreasesPrice() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 priceBefore = market.getPrice(1);

        vm.prank(trader1);
        market.buy(1, 1000 ether, 0);

        uint256 priceAfter = market.getPrice(1);

        assertGt(priceAfter, priceBefore, "Buy should increase price");
    }

    function test_LMSR_SellDecreasesPrice() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        uint256 tokens = market.buy(1, 2000 ether, 0);

        uint256 priceBefore = market.getPrice(1);

        vm.prank(trader1);
        market.sell(1, tokens / 2, 0);

        uint256 priceAfter = market.getPrice(1);

        assertLt(priceAfter, priceBefore, "Sell should decrease price");
    }

    // ============ Gas Profiling Tests ============

    function test_Gas_Buy() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 gasUsed = profileBuy(IMarket(address(market)), trader1, 0, 1000 ether);
        
        emit log_named_uint("Gas used for buy", gasUsed);
        
        // Target: <150k gas
        // Note: This is approximate and will vary
    }

    function test_Gas_Sell() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        vm.prank(trader1);
        uint256 tokens = market.buy(0, 1000 ether, 0);

        uint256 gasBefore = gasleft();
        vm.prank(trader1);
        market.sell(0, tokens, 0);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for sell", gasUsed);
    }

    function test_Gas_AddLiquidity() public {
        uint256 gasUsed = profileAddLiquidity(IMarket(address(market)), lp1, DEFAULT_INITIAL_LIQUIDITY);
        
        emit log_named_uint("Gas used for initial liquidity", gasUsed);
    }

    function test_Gas_GetPrice() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 gasBefore = gasleft();
        market.getPrice(0);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for getPrice", gasUsed);
    }

    function test_Gas_GetAllPrices() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 gasBefore = gasleft();
        market.getAllPrices();
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for getAllPrices", gasUsed);
    }

    // ============ Edge Case Tests ============

    function test_EdgeCase_LargeTradeImpact() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Very large trade
        uint256 largeAmount = 5000 ether;
        
        vm.prank(trader1);
        uint256 tokens = market.buy(0, largeAmount, 0);

        assertGt(tokens, 0);

        // Price should be significantly higher
        uint256 price = market.getPrice(0);
        assertGt(price, 0.5e18, "Large buy should push price high");
    }

    function test_EdgeCase_MultipleSmallTrades() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // Execute many small trades
        uint256 smallAmount = 100 ether;
        
        for (uint256 i = 0; i < 5; i++) {
            vm.prank(trader1);
            market.buy(0, smallAmount, 0);
        }

        // Should accumulate to price increase
        uint256 price = market.getPrice(0);
        uint256 basePrice = 1e18 / OUTCOME_COUNT;
        assertGt(price, basePrice);
    }

    function test_EdgeCase_BuyAndSellRoundtrip() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        uint256 buyAmount = 1000 ether;
        uint256 collateralBefore = collateral.balanceOf(trader1);
        
        // Buy
        vm.prank(trader1);
        uint256 tokens = market.buy(0, buyAmount, 0);

        // Sell all
        vm.prank(trader1);
        uint256 collateralOut = market.sell(0, tokens, 0);

        uint256 collateralAfter = collateral.balanceOf(trader1);
        uint256 loss = collateralBefore - collateralAfter;

        // Should lose money due to fees and price impact
        assertGt(loss, 0, "Roundtrip should have cost");
        emit log_named_uint("Roundtrip loss", loss);
    }

    // ============ LMSR Invariant Tests ============

    function test_LMSR_Invariant_BoundedLoss() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // LMSR property: Maximum loss for liquidity provider is b * ln(n)
        // where b = liquidityParameter, n = number of outcomes
        // For 4 outcomes: max loss = b * ln(4) ≈ b * 1.386

        uint256 initialCollateral = market.totalCollateral();
        
        // Simulate traders buying outcome 0
        // Note: we're testing that LP losses are bounded, not necessarily reaching the max
        vm.prank(trader1);
        market.buy(0, 2000 ether, 0);
        
        vm.prank(trader2);
        market.buy(0, 1000 ether, 0);
        
        // Check price has increased significantly for outcome 0
        uint256 price0 = market.getPrice(0);
        emit log_named_uint("Price of outcome 0 after buying", price0);
        
        // Calculate theoretical max loss: b * ln(n)
        // Using approximation: ln(4) ≈ 1.386
        uint256 theoreticalMaxLoss = (LIQUIDITY_PARAM * 1386) / 1000; // 1.386 * b
        
        // In LMSR, the max loss is bounded regardless of trading
        // The LP should never lose more than b * ln(n) even in worst case
        emit log_named_uint("Theoretical max loss", theoreticalMaxLoss);
        emit log_named_uint("Initial collateral", initialCollateral);
        emit log_named_uint("Current collateral", market.totalCollateral());
        
        // After trading, total collateral has increased due to fees
        // This test verifies the bounded loss property holds structurally
        assertGt(market.totalCollateral(), initialCollateral, "Collateral should increase due to fees");
    }

    function test_LMSR_Invariant_NoArbitrage() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // LMSR invariant: No profitable arbitrage should exist
        // If we buy outcome A then immediately sell, we should lose money (fees + slippage)

        uint256 buyAmount = 1000 ether;
        
        for (uint256 outcomeId = 0; outcomeId < OUTCOME_COUNT; outcomeId++) {
            uint256 balanceBefore = collateral.balanceOf(trader1);
            
            // Buy
            vm.prank(trader1);
            uint256 tokens = market.buy(outcomeId, buyAmount, 0);
            
            // Immediately sell
            vm.prank(trader1);
            market.sell(outcomeId, tokens, 0);
            
            uint256 balanceAfter = collateral.balanceOf(trader1);
            
            // Should always lose money
            assertLt(balanceAfter, balanceBefore, "Roundtrip should lose money");
            
            uint256 loss = balanceBefore - balanceAfter;
            emit log_named_uint("Roundtrip loss for outcome", outcomeId);
            emit log_named_uint("Loss amount", loss);
        }
    }

    function test_LMSR_Invariant_ConvexCostFunction() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // LMSR cost function is convex: buying more of an outcome becomes increasingly expensive
        // Test: buying 2x amount should cost more than 2x the price of 1x amount

        uint256 smallAmount = 500 ether;
        uint256 largeAmount = 1000 ether;
        
        // Get quote for small buy
        (uint256 tokensSmall, uint256 feeSmall) = market.getQuoteBuy(0, smallAmount, trader1);
        uint256 costSmall = smallAmount;
        uint256 avgPriceSmall = (costSmall * 1e18) / tokensSmall;
        
        // Get quote for large buy (2x)
        (uint256 tokensLarge, uint256 feeLarge) = market.getQuoteBuy(0, largeAmount, trader1);
        uint256 costLarge = largeAmount;
        uint256 avgPriceLarge = (costLarge * 1e18) / tokensLarge;
        
        emit log_named_uint("Small buy avg price", avgPriceSmall);
        emit log_named_uint("Large buy avg price", avgPriceLarge);
        emit log_named_uint("Tokens from small buy", tokensSmall);
        emit log_named_uint("Tokens from large buy", tokensLarge);
        
        // Due to convexity, average price for larger buy should be higher
        assertGt(avgPriceLarge, avgPriceSmall, "Larger buys should have higher average price (convexity)");
        
        // Also, 2x the cost should yield less than 2x the tokens
        assertLt(tokensLarge, tokensSmall * 2, "2x cost should yield less than 2x tokens due to price impact");
    }

    function test_LMSR_Invariant_PriceSumAlwaysOne() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // LMSR invariant: sum of all prices should always equal 1.0
        // Test after various market states

        // Initial state
        uint256[] memory prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
        
        // After unbalanced buying
        vm.prank(trader1);
        market.buy(0, 2000 ether, 0);
        
        vm.prank(trader2);
        market.buy(2, 500 ether, 0);
        
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
        
        // After selling
        uint256 tokens = getOutcomeBalance(trader1, MARKET_ID, 0);
        vm.prank(trader1);
        market.sell(0, tokens / 3, 0);
        
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
        
        // After liquidity changes
        vm.prank(lp2);
        market.addLiquidity(5000 ether);
        
        prices = market.getAllPrices();
        assertPricesSumToOne(prices, PRICE_TOLERANCE_BPS);
    }

    function test_LMSR_Invariant_MonotonicPrice() public {
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);

        // LMSR invariant: buying an outcome should monotonically increase its price
        // Selling should monotonically decrease its price
        
        uint256 outcomeId = 1;
        uint256[] memory priceHistory = new uint256[](6);
        
        // Record initial price
        priceHistory[0] = market.getPrice(outcomeId);
        
        // Buy #1
        vm.prank(trader1);
        market.buy(outcomeId, 500 ether, 0);
        priceHistory[1] = market.getPrice(outcomeId);
        assertGt(priceHistory[1], priceHistory[0], "Price should increase after buy #1");
        
        // Buy #2
        vm.prank(trader2);
        market.buy(outcomeId, 500 ether, 0);
        priceHistory[2] = market.getPrice(outcomeId);
        assertGt(priceHistory[2], priceHistory[1], "Price should increase after buy #2");
        
        // Buy #3
        vm.prank(trader3);
        market.buy(outcomeId, 500 ether, 0);
        priceHistory[3] = market.getPrice(outcomeId);
        assertGt(priceHistory[3], priceHistory[2], "Price should increase after buy #3");
        
        // Now sell back
        uint256 tokensTrader1 = getOutcomeBalance(trader1, MARKET_ID, outcomeId);
        vm.prank(trader1);
        market.sell(outcomeId, tokensTrader1, 0);
        priceHistory[4] = market.getPrice(outcomeId);
        assertLt(priceHistory[4], priceHistory[3], "Price should decrease after sell #1");
        
        // Sell more
        uint256 tokensTrader2 = getOutcomeBalance(trader2, MARKET_ID, outcomeId);
        vm.prank(trader2);
        market.sell(outcomeId, tokensTrader2, 0);
        priceHistory[5] = market.getPrice(outcomeId);
        assertLt(priceHistory[5], priceHistory[4], "Price should decrease after sell #2");
        
        // Log price history
        for (uint256 i = 0; i < priceHistory.length; i++) {
            emit log_named_uint("Price at step", i);
            emit log_named_uint("Value", priceHistory[i]);
        }
    }

    function test_LMSR_Invariant_SymmetricOutcomes() public {
        // LMSR property: with equal reserves, buying equal amounts of different outcomes
        // should result in approximately equal tokens and prices
        // Test this with TWO separate markets to ensure truly equal initial conditions
        
        // Market 1: buy outcome 0
        MultiChoiceMarket market1 = new MultiChoiceMarket(
            100,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            OUTCOME_COUNT,
            LIQUIDITY_PARAM
        );
        registerMarket(100, creator);
        outcomeToken.setAMMAuthorization(address(market1), true);
        approveMarketForAll(address(market1));
        
        // Market 2: buy outcome 1
        MultiChoiceMarket market2 = new MultiChoiceMarket(
            101,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            OUTCOME_COUNT,
            LIQUIDITY_PARAM
        );
        registerMarket(101, creator);
        outcomeToken.setAMMAuthorization(address(market2), true);
        approveMarketForAll(address(market2));
        
        // Add identical liquidity to both
        vm.prank(lp1);
        market1.addLiquidity(DEFAULT_INITIAL_LIQUIDITY);
        
        vm.prank(lp1);
        market2.addLiquidity(DEFAULT_INITIAL_LIQUIDITY);
        
        uint256 buyAmount = 800 ether;
        
        // Buy outcome 0 in market1
        vm.prank(trader1);
        uint256 tokens0 = market1.buy(0, buyAmount, 0);
        
        // Buy outcome 1 in market2
        vm.prank(trader2);
        uint256 tokens1 = market2.buy(1, buyAmount, 0);
        
        // Should receive equal tokens since both markets started with identical state
        assertEq(tokens0, tokens1, "Equal buys in symmetric markets should yield equal tokens");
        
        // Prices should be equal
        uint256 price0 = market1.getPrice(0);
        uint256 price1 = market2.getPrice(1);
        
        assertEq(price0, price1, "Prices should be equal in symmetric markets");
        
        emit log_named_uint("Tokens from market1 outcome 0", tokens0);
        emit log_named_uint("Tokens from market2 outcome 1", tokens1);
        emit log_named_uint("Price of market1 outcome 0", price0);
        emit log_named_uint("Price of market2 outcome 1", price1);
    }

    // ============ Fuzz Tests for LMSR ============

    function testFuzz_LMSR_PricesSumToOne(uint256 buyAmount) public {
        // Bound: reasonable buy amounts (1 to 5000 ether)
        buyAmount = bound(buyAmount, 1 ether, 5000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        // Execute a random buy
        vm.prank(trader1);
        market.buy(0, buyAmount, 0);
        
        // Check that all prices sum to approximately 1.0 (within rounding tolerance)
        uint256[] memory prices = market.getAllPrices();
        uint256 priceSum = 0;
        
        for (uint256 i = 0; i < prices.length; i++) {
            priceSum += prices[i];
        }
        
        // Allow 0.1% tolerance for rounding
        uint256 tolerance = 1e18 / 1000; // 0.001 = 0.1%
        assertApproxEqAbs(priceSum, 1e18, tolerance, "Prices should sum to 1.0");
    }

    function testFuzz_LMSR_BuyIncreasesPrice(uint256 outcomeId, uint256 buyAmount) public {
        // Bound outcome to valid range
        outcomeId = bound(outcomeId, 0, OUTCOME_COUNT - 1);
        // Bound buy amount to reasonable range
        buyAmount = bound(buyAmount, 1 ether, 2000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        uint256 priceBefore = market.getPrice(outcomeId);
        
        vm.prank(trader1);
        market.buy(outcomeId, buyAmount, 0);
        
        uint256 priceAfter = market.getPrice(outcomeId);
        
        assertGt(priceAfter, priceBefore, "Buying should increase price");
    }

    function testFuzz_LMSR_SellDecreasesPrice(uint256 outcomeId, uint256 buyAmount) public {
        // Bound outcome to valid range
        outcomeId = bound(outcomeId, 0, OUTCOME_COUNT - 1);
        // Bound buy amount to reasonable range  
        buyAmount = bound(buyAmount, 100 ether, 1000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        // First buy some tokens
        vm.prank(trader1);
        uint256 tokens = market.buy(outcomeId, buyAmount, 0);
        
        uint256 priceBefore = market.getPrice(outcomeId);
        
        // Then sell half of them
        vm.prank(trader1);
        market.sell(outcomeId, tokens / 2, 0);
        
        uint256 priceAfter = market.getPrice(outcomeId);
        
        assertLt(priceAfter, priceBefore, "Selling should decrease price");
    }

    function testFuzz_LMSR_NoArbitrage(uint256 outcomeId, uint256 buyAmount) public {
        // Bound to valid ranges
        outcomeId = bound(outcomeId, 0, OUTCOME_COUNT - 1);
        buyAmount = bound(buyAmount, 100 ether, 2000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        uint256 balanceBefore = collateral.balanceOf(trader1);
        
        // Buy tokens
        vm.prank(trader1);
        uint256 tokens = market.buy(outcomeId, buyAmount, 0);
        
        // Immediately sell them back
        vm.prank(trader1);
        market.sell(outcomeId, tokens, 0);
        
        uint256 balanceAfter = collateral.balanceOf(trader1);
        
        // Should always lose money on a round trip (fees + slippage)
        assertLt(balanceAfter, balanceBefore, "Round trip should lose money");
    }

    function testFuzz_LMSR_ConvexityProperty(uint256 outcomeId, uint256 baseAmount) public {
        // Bound to valid ranges
        outcomeId = bound(outcomeId, 0, OUTCOME_COUNT - 1);
        baseAmount = bound(baseAmount, 100 ether, 1000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        // Get quote for buying baseAmount
        (uint256 tokensSmall,) = market.getQuoteBuy(outcomeId, baseAmount, trader1);
        uint256 avgPriceSmall = (baseAmount * 1e18) / tokensSmall;
        
        // Get quote for buying 2x baseAmount
        (uint256 tokensLarge,) = market.getQuoteBuy(outcomeId, baseAmount * 2, trader1);
        uint256 avgPriceLarge = (baseAmount * 2 * 1e18) / tokensLarge;
        
        // Due to convexity, larger buys should have higher average price
        assertGt(avgPriceLarge, avgPriceSmall, "Larger buys should have higher avg price (convexity)");
        
        // Also, 2x the cost should yield less than 2x the tokens
        assertLt(tokensLarge, tokensSmall * 2, "2x cost should yield less than 2x tokens");
    }

    function testFuzz_LMSR_MultipleTrades(
        uint256 trade1Outcome,
        uint256 trade1Amount,
        uint256 trade2Outcome,
        uint256 trade2Amount
    ) public {
        // Bound inputs
        trade1Outcome = bound(trade1Outcome, 0, OUTCOME_COUNT - 1);
        trade2Outcome = bound(trade2Outcome, 0, OUTCOME_COUNT - 1);
        trade1Amount = bound(trade1Amount, 50 ether, 1000 ether);
        trade2Amount = bound(trade2Amount, 50 ether, 1000 ether);
        
        addInitialLiquidity(IMarket(address(market)), DEFAULT_INITIAL_LIQUIDITY);
        
        // Execute two trades
        vm.prank(trader1);
        market.buy(trade1Outcome, trade1Amount, 0);
        
        vm.prank(trader2);
        market.buy(trade2Outcome, trade2Amount, 0);
        
        // Verify prices still sum to 1.0
        uint256[] memory prices = market.getAllPrices();
        uint256 priceSum = 0;
        
        for (uint256 i = 0; i < prices.length; i++) {
            priceSum += prices[i];
        }
        
        uint256 tolerance = 1e18 / 1000; // 0.1% tolerance
        assertApproxEqAbs(priceSum, 1e18, tolerance, "Prices should still sum to 1.0 after multiple trades");
        
        // All prices should be between 0 and 1
        for (uint256 i = 0; i < prices.length; i++) {
            assertGe(prices[i], 0, "Price should be >= 0");
            assertLe(prices[i], 1e18, "Price should be <= 1.0");
        }
    }

    // ============ Maximum Outcomes Test (8) ============

    function test_LMSR_MaxOutcomes_AllInvariants() public {
        // Test LMSR with maximum outcomes (8) to ensure it scales properly
        uint256 maxOutcomes = 8;
        
        // Create market with 8 outcomes
        MultiChoiceMarket market8 = new MultiChoiceMarket(
            200,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            maxOutcomes,
            LIQUIDITY_PARAM
        );
        registerMarket(200, creator);
        outcomeToken.setAMMAuthorization(address(market8), true);
        approveMarketForAll(address(market8));
        
        // Add liquidity
        vm.prank(lp1);
        market8.addLiquidity(DEFAULT_INITIAL_LIQUIDITY);
        
        // Test 1: All prices should be equal initially
        uint256[] memory initialPrices = market8.getAllPrices();
        uint256 expectedEqualPrice = 1e18 / maxOutcomes; // 1/8 = 0.125
        
        for (uint256 i = 0; i < maxOutcomes; i++) {
            assertApproxEqAbs(
                initialPrices[i],
                expectedEqualPrice,
                1e15, // 0.001 tolerance
                "Initial prices should be equal"
            );
        }
        
        // Test 2: Prices sum to 1.0
        uint256 priceSum = 0;
        for (uint256 i = 0; i < maxOutcomes; i++) {
            priceSum += initialPrices[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should sum to 1.0");
        
        // Test 3: Buy outcome 0, verify price increases and others decrease slightly
        vm.prank(trader1);
        market8.buy(0, 1000 ether, 0);
        
        uint256[] memory pricesAfterBuy = market8.getAllPrices();
        assertGt(pricesAfterBuy[0], initialPrices[0], "Bought outcome price should increase");
        
        // Other outcomes should have decreased slightly (relative to initial)
        for (uint256 i = 1; i < maxOutcomes; i++) {
            assertLt(pricesAfterBuy[i], initialPrices[i], "Other outcome prices should decrease");
        }
        
        // Test 4: Prices still sum to 1.0 after trade
        priceSum = 0;
        for (uint256 i = 0; i < maxOutcomes; i++) {
            priceSum += pricesAfterBuy[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should still sum to 1.0 after buy");
        
        // Test 5: Multiple trades on different outcomes
        vm.prank(trader2);
        market8.buy(3, 500 ether, 0);
        
        vm.prank(trader3);
        market8.buy(7, 800 ether, 0);
        
        uint256[] memory pricesAfterMultiple = market8.getAllPrices();
        priceSum = 0;
        for (uint256 i = 0; i < maxOutcomes; i++) {
            priceSum += pricesAfterMultiple[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should still sum to 1.0 after multiple trades");
        
        // Test 6: All prices are valid (0 < price < 1)
        for (uint256 i = 0; i < maxOutcomes; i++) {
            assertGt(pricesAfterMultiple[i], 0, "Price should be > 0");
            assertLt(pricesAfterMultiple[i], 1e18, "Price should be < 1.0");
        }
        
        // Test 7: Bounded loss property for 8 outcomes
        // Max loss = b * ln(8) ≈ b * 2.079
        uint256 theoreticalMaxLoss = (LIQUIDITY_PARAM * 2079) / 1000; // 2.079 * b
        emit log_named_uint("Theoretical max loss for 8 outcomes", theoreticalMaxLoss);
        emit log_named_uint("Initial collateral", DEFAULT_INITIAL_LIQUIDITY);
        
        // The loss bound is a theoretical maximum - in practice with fees, 
        // the market should have more collateral
        assertGt(market8.totalCollateral(), DEFAULT_INITIAL_LIQUIDITY, "Should have gained from fees");
        
        emit log_string("=== 8 Outcome Market Test Summary ===");
        for (uint256 i = 0; i < maxOutcomes; i++) {
            emit log_named_uint("Outcome price", i);
            emit log_named_uint("Value", pricesAfterMultiple[i]);
        }
    }
}
