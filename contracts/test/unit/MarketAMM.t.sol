// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/MarketAMM.sol";
import "../../src/OutcomeToken.sol";
import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";

import "../mocks/MockERC20.sol";

contract MarketAMMTest is Test {
    MarketAMM public amm;
    OutcomeToken public outcomeToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    MockERC20 public horizonToken;
    MockERC20 public collateral;

    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator = address(0x2);
    address public lp1 = address(0x3);
    address public lp2 = address(0x4);
    address public trader1 = address(0x5);
    address public trader2 = address(0x6);

    uint256 public constant MARKET_ID = 1;
    uint256 public closeTime;

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

    function setUp() public {
        // Deploy tokens
        collateral = new MockERC20("USDC", "USDC");
        horizonToken = new MockERC20("Horizon Token", "HORIZON"); horizonToken.mint(address(this), 1_000_000_000 * 10 ** 18);

        // Deploy core contracts
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        feeSplitter = new FeeSplitter(treasury);
        horizonPerks = new HorizonPerks(address(horizonToken));

        // Register market
        outcomeToken.registerMarket(MARKET_ID, collateral);
        feeSplitter.registerMarket(MARKET_ID, creator);

        // Deploy AMM
        closeTime = block.timestamp + 30 days;
        amm = new MarketAMM(
            MARKET_ID, address(collateral), address(outcomeToken), address(feeSplitter), address(horizonPerks), closeTime
        );

        // Authorize AMM
        outcomeToken.setAMMAuthorization(address(amm), true);

        // Authorize test contract as resolver (for testing resolution scenarios)
        outcomeToken.setResolutionAuthorization(address(this), true);

        // Fund test accounts
        collateral.mint(lp1, 1_000_000 ether);
        collateral.mint(lp2, 1_000_000 ether);
        collateral.mint(trader1, 1_000_000 ether);
        collateral.mint(trader2, 1_000_000 ether);

        // Approve AMM
        vm.prank(lp1);
        collateral.approve(address(amm), type(uint256).max);
        vm.prank(lp2);
        collateral.approve(address(amm), type(uint256).max);
        vm.prank(trader1);
        collateral.approve(address(amm), type(uint256).max);
        vm.prank(trader2);
        collateral.approve(address(amm), type(uint256).max);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(amm.marketId(), MARKET_ID);
        assertEq(address(amm.collateralToken()), address(collateral));
        assertEq(address(amm.outcomeToken()), address(outcomeToken));
        assertEq(address(amm.feeSplitter()), address(feeSplitter));
        assertEq(address(amm.horizonPerks()), address(horizonPerks));
        assertEq(amm.closeTime(), closeTime);
    }

    // ============ Initial Liquidity Tests ============

    function test_AddLiquidity_Initial() public {
        uint256 amount = 10000 ether;

        vm.expectEmit(true, false, false, true);
        emit LiquidityAdded(lp1, amount, amount - amm.MINIMUM_LIQUIDITY());

        vm.prank(lp1);
        uint256 lpTokens = amm.addLiquidity(amount);

        // Check LP tokens
        assertEq(lpTokens, amount - amm.MINIMUM_LIQUIDITY());
        assertEq(amm.balanceOf(lp1), lpTokens);
        assertEq(amm.totalSupply(), amount);

        // Check reserves
        assertEq(amm.reserveYes(), amount);
        assertEq(amm.reserveNo(), amount);
        assertEq(amm.totalCollateral(), amount);
    }

    function test_AddLiquidity_Subsequent() public {
        // Initial liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Subsequent liquidity
        uint256 amount = 5000 ether;
        uint256 totalSupplyBefore = amm.totalSupply();
        uint256 totalCollateralBefore = amm.totalCollateral();

        vm.prank(lp2);
        uint256 lpTokens = amm.addLiquidity(amount);

        // LP tokens should be proportional
        uint256 expectedLP = (amount * totalSupplyBefore) / totalCollateralBefore;
        assertEq(lpTokens, expectedLP);
    }

    function test_RevertWhen_AddLiquidity_ZeroAmount() public {
        vm.prank(lp1);
        vm.expectRevert(MarketAMM.InvalidAmount.selector);
        amm.addLiquidity(0);
    }

    function test_RevertWhen_AddLiquidity_AfterResolution() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Resolve market
        outcomeToken.setWinningOutcome(MARKET_ID, 0);

        // Try to add more liquidity
        vm.prank(lp2);
        vm.expectRevert(MarketAMM.MarketResolved.selector);
        amm.addLiquidity(5000 ether);
    }

    // ============ Remove Liquidity Tests ============

    function test_RemoveLiquidity() public {
        // Add liquidity
        uint256 amount = 10000 ether;
        vm.prank(lp1);
        uint256 lpTokens = amm.addLiquidity(amount);

        // Remove half
        uint256 removeAmount = lpTokens / 2;
        uint256 expectedCollateral = (removeAmount * amount) / amm.totalSupply();

        vm.expectEmit(true, false, false, true);
        emit LiquidityRemoved(lp1, removeAmount, expectedCollateral);

        vm.prank(lp1);
        uint256 collateralOut = amm.removeLiquidity(removeAmount);

        assertApproxEqRel(collateralOut, expectedCollateral, 0.01e18); // 1% tolerance
        assertEq(amm.balanceOf(lp1), lpTokens - removeAmount);
    }

    function test_RevertWhen_RemoveLiquidity_ZeroAmount() public {
        vm.prank(lp1);
        vm.expectRevert(MarketAMM.InvalidAmount.selector);
        amm.removeLiquidity(0);
    }

    function test_RevertWhen_RemoveLiquidity_InsufficientBalance() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        vm.prank(lp2);
        vm.expectRevert(MarketAMM.InsufficientLPTokens.selector);
        amm.removeLiquidity(1000 ether);
    }

    // ============ Buy Tests ============

    function test_BuyYes() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Buy Yes tokens
        uint256 collateralIn = 100 ether;
        (uint256 expectedOut, uint256 expectedFee) = amm.getQuoteBuyYes(collateralIn, trader1);

        vm.prank(trader1);
        uint256 tokensOut = amm.buyYes(collateralIn, 0);

        assertEq(tokensOut, expectedOut);
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 0), tokensOut);

        // Check reserves updated
        assertTrue(amm.reserveYes() < 10000 ether); // Yes reserve decreased
        assertTrue(amm.reserveNo() > 10000 ether); // No reserve increased
    }

    function test_BuyNo() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Buy No tokens
        uint256 collateralIn = 100 ether;
        (uint256 expectedOut,) = amm.getQuoteBuyNo(collateralIn, trader1);

        vm.prank(trader1);
        uint256 tokensOut = amm.buyNo(collateralIn, 0);

        assertEq(tokensOut, expectedOut);
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 1), tokensOut);

        // Check reserves updated
        assertTrue(amm.reserveNo() < 10000 ether); // No reserve decreased
        assertTrue(amm.reserveYes() > 10000 ether); // Yes reserve increased
    }

    function test_BuyYes_MovesPrice() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 initialPrice = amm.getYesPrice();
        assertEq(initialPrice, 0.5e18); // 50% initially

        // Buy Yes tokens (should increase Yes price)
        vm.prank(trader1);
        amm.buyYes(1000 ether, 0);

        uint256 newPrice = amm.getYesPrice();
        assertTrue(newPrice > initialPrice); // Price increased
    }

    function test_RevertWhen_Buy_SlippageExceeded() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 collateralIn = 100 ether;
        (uint256 expectedOut,) = amm.getQuoteBuyYes(collateralIn, trader1);

        vm.prank(trader1);
        vm.expectRevert(MarketAMM.SlippageExceeded.selector);
        amm.buyYes(collateralIn, expectedOut + 1); // Request more than possible
    }

    function test_RevertWhen_Buy_AfterClose() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Warp past close time
        vm.warp(closeTime + 1);

        vm.prank(trader1);
        vm.expectRevert(MarketAMM.MarketClosed.selector);
        amm.buyYes(100 ether, 0);
    }

    function test_RevertWhen_Buy_NoLiquidity() public {
        vm.prank(trader1);
        vm.expectRevert(MarketAMM.InsufficientLiquidity.selector);
        amm.buyYes(100 ether, 0);
    }

    // ============ Sell Tests ============

    function test_SellYes() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Buy Yes tokens first
        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(100 ether, 0);

        // Approve outcome tokens for burning
        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        // Sell Yes tokens
        uint256 sellAmount = yesTokens / 2;
        (uint256 expectedOut,) = amm.getQuoteSellYes(sellAmount, trader1);

        vm.prank(trader1);
        uint256 collateralOut = amm.sellYes(sellAmount, 0);

        assertEq(collateralOut, expectedOut);
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 0), yesTokens - sellAmount);
    }

    function test_SellNo() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Buy No tokens first
        vm.prank(trader1);
        uint256 noTokens = amm.buyNo(100 ether, 0);

        // Approve outcome tokens
        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        // Sell No tokens
        uint256 sellAmount = noTokens / 2;
        (uint256 expectedOut,) = amm.getQuoteSellNo(sellAmount, trader1);

        vm.prank(trader1);
        uint256 collateralOut = amm.sellNo(sellAmount, 0);

        assertEq(collateralOut, expectedOut);
    }

    function test_RevertWhen_Sell_SlippageExceeded() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(100 ether, 0);

        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        (uint256 expectedOut,) = amm.getQuoteSellYes(yesTokens, trader1);

        vm.prank(trader1);
        vm.expectRevert(MarketAMM.SlippageExceeded.selector);
        amm.sellYes(yesTokens, expectedOut + 1);
    }

    // ============ Fee Tests ============

    function test_Fees_DefaultTier() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 collateralIn = 100 ether;
        (uint256 tokensOut, uint256 fee) = amm.getQuoteBuyYes(collateralIn, trader1);

        // Default fee is 200 bps (2%)
        uint256 expectedFee = (collateralIn * 200) / 10000;
        assertEq(fee, expectedFee);

        vm.prank(trader1);
        amm.buyYes(collateralIn, 0);

        // Check fee was distributed
        assertTrue(feeSplitter.getProtocolPendingFees(address(collateral)) > 0);
    }

    function test_Fees_WithHorizonDiscount() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Give trader HORIZON tokens for tier 1 (10K HORIZON)
        horizonToken.transfer(trader1, 10_000 * 10 ** 18);

        uint256 collateralIn = 100 ether;
        (uint256 tokensOut, uint256 fee) = amm.getQuoteBuyYes(collateralIn, trader1);

        // In new model, user fee is constant at 200 bps (2%) regardless of HORIZON holdings
        // What changes is protocol/creator split, not the user-facing fee
        uint256 expectedFee = (collateralIn * 200) / 10000;
        assertEq(fee, expectedFee);
    }

    // ============ Price Tests ============

    function test_GetYesPrice_Initial() public view {
        uint256 price = amm.getYesPrice();
        assertEq(price, 0.5e18); // 50% when no liquidity
    }

    function test_GetYesPrice_AfterLiquidity() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 price = amm.getYesPrice();
        assertEq(price, 0.5e18); // 50% with equal reserves
    }

    function test_GetYesPrice_AfterTrade() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        vm.prank(trader1);
        amm.buyYes(1000 ether, 0);

        uint256 yesPrice = amm.getYesPrice();
        uint256 noPrice = amm.getNoPrice();

        // Prices should sum to ~1.0 (with small rounding)
        assertApproxEqRel(yesPrice + noPrice, 1e18, 0.01e18);
        assertTrue(yesPrice > 0.5e18); // Yes price increased
        assertTrue(noPrice < 0.5e18); // No price decreased
    }

    // ============ Quote Accuracy Tests ============

    function test_Quote_MatchesActual_Buy() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 collateralIn = 100 ether;
        (uint256 quotedOut, uint256 quotedFee) = amm.getQuoteBuyYes(collateralIn, trader1);

        vm.prank(trader1);
        uint256 actualOut = amm.buyYes(collateralIn, 0);

        assertEq(actualOut, quotedOut);
    }

    function test_Quote_MatchesActual_Sell() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(100 ether, 0);

        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        (uint256 quotedOut, uint256 quotedFee) = amm.getQuoteSellYes(yesTokens, trader1);

        vm.prank(trader1);
        uint256 actualOut = amm.sellYes(yesTokens, 0);

        assertEq(actualOut, quotedOut);
    }

    // ============ Integration Tests ============

    function test_FullCycle_BuyAndSell() public {
        // Add liquidity
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 initialBalance = collateral.balanceOf(trader1);

        // Buy Yes
        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(100 ether, 0);

        // Sell Yes
        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        vm.prank(trader1);
        uint256 collateralOut = amm.sellYes(yesTokens, 0);

        // Should get back less than input due to fees and slippage
        uint256 finalBalance = collateral.balanceOf(trader1);
        assertTrue(finalBalance < initialBalance);
    }

    function test_MultipleLPs_ProportionalShares() public {
        // LP1 adds 10K
        vm.prank(lp1);
        uint256 lp1Tokens = amm.addLiquidity(10000 ether);

        // LP2 adds 5K (should get half the LP tokens)
        vm.prank(lp2);
        uint256 lp2Tokens = amm.addLiquidity(5000 ether);

        assertApproxEqRel(lp2Tokens, lp1Tokens / 2, 0.01e18);
    }

    function test_CPMM_Invariant() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 k1 = amm.reserveYes() * amm.reserveNo();

        // Trade
        vm.prank(trader1);
        amm.buyYes(100 ether, 0);

        uint256 k2 = amm.reserveYes() * amm.reserveNo();

        // k should stay approximately constant (fees are taken before swap, not added to pool)
        // Allow small decrease due to rounding in integer division (< 0.01%)
        assertApproxEqRel(k2, k1, 0.0001e18); // 0.01% tolerance
    }

    function testFuzz_AddLiquidity(uint128 amount) public {
        vm.assume(amount > 1000 && amount < 1_000_000 ether);

        vm.prank(lp1);
        uint256 lpTokens = amm.addLiquidity(amount);

        assertGt(lpTokens, 0);
        assertEq(amm.totalCollateral(), amount);
    }

    function testFuzz_BuyYes(uint64 amount) public {
        vm.assume(amount > 0.1 ether && amount < 1000 ether);

        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 k1 = amm.reserveYes() * amm.reserveNo();

        vm.prank(trader1);
        uint256 tokensOut = amm.buyYes(amount, 0);

        assertGt(tokensOut, 0);

        uint256 k2 = amm.reserveYes() * amm.reserveNo();
        // Invariant should hold approximately (allow small rounding error)
        assertApproxEqRel(k2, k1, 0.0001e18); // 0.01% tolerance
    }

    // ============ Additional Invariant Tests ============

    function test_Invariant_ReserveSum() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 initialSum = amm.reserveYes() + amm.reserveNo();

        // Trade multiple times
        vm.prank(trader1);
        amm.buyYes(100 ether, 0);

        vm.prank(trader2);
        amm.buyNo(50 ether, 0);

        uint256 finalSum = amm.reserveYes() + amm.reserveNo();

        // Reserve sum should increase with new collateral added (minus fees)
        assertTrue(finalSum > initialSum);
    }

    function test_Invariant_TotalCollateralBacksReserves() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Initial state: collateral = 10K, reserves = 10K Yes + 10K No (20K outcome tokens backed by 10K collateral)
        assertEq(amm.totalCollateral(), 10000 ether);
        assertEq(amm.reserveYes(), 10000 ether);
        assertEq(amm.reserveNo(), 10000 ether);

        // After trade: collateral increases (minus fees), reserves adjust
        vm.prank(trader1);
        amm.buyYes(100 ether, 0);

        // Collateral should increase by collateralAfterFee (98 ether with 2% fee)
        assertGt(amm.totalCollateral(), 10000 ether);
        assertLt(amm.totalCollateral(), 10100 ether); // Less than 10100 due to fees
    }

    function test_Invariant_PricesSumToOne() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 yesPrice = amm.getYesPrice();
        uint256 noPrice = amm.getNoPrice();

        // Prices should always sum to ~1.0
        assertApproxEqAbs(yesPrice + noPrice, 1e18, 1e15); // 0.001 tolerance

        // After imbalanced trade
        vm.prank(trader1);
        amm.buyYes(1000 ether, 0);

        yesPrice = amm.getYesPrice();
        noPrice = amm.getNoPrice();
        assertApproxEqAbs(yesPrice + noPrice, 1e18, 1e15);
    }

    function test_Invariant_NoArbitrage() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 startBalance = collateral.balanceOf(trader1);

        // Buy and immediately sell back
        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(100 ether, 0);

        vm.prank(trader1);
        outcomeToken.setApprovalForAll(address(amm), true);

        vm.prank(trader1);
        amm.sellYes(yesTokens, 0);

        uint256 endBalance = collateral.balanceOf(trader1);

        // Should lose money due to fees + slippage (no arbitrage)
        assertTrue(endBalance < startBalance);
    }

    // ============ Edge Case Tests ============

    function test_EdgeCase_VerySmallTrade() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Trade 0.01 tokens
        vm.prank(trader1);
        uint256 tokensOut = amm.buyYes(0.01 ether, 0);

        assertGt(tokensOut, 0);
    }

    function test_EdgeCase_VeryLargeTrade() public {
        vm.prank(lp1);
        amm.addLiquidity(100000 ether);

        // Trade 50% of pool
        vm.prank(trader1);
        uint256 tokensOut = amm.buyYes(50000 ether, 0);

        assertGt(tokensOut, 0);
        // Reserve should be significantly depleted but not empty
        assertGt(amm.reserveYes(), 0);
    }

    function test_EdgeCase_ExtremelyImbalancedPool() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Buy heavily in one direction (multiple trades to get extreme imbalance)
        vm.prank(trader1);
        amm.buyYes(2000 ether, 0);

        vm.prank(trader1);
        amm.buyYes(2000 ether, 0);

        vm.prank(trader1);
        amm.buyYes(2000 ether, 0);

        // Yes price should be high, No price low
        uint256 yesPrice = amm.getYesPrice();
        uint256 noPrice = amm.getNoPrice();

        assertTrue(yesPrice > 0.6e18); // > 60%
        assertTrue(noPrice < 0.4e18); // < 40%
        assertApproxEqAbs(yesPrice + noPrice, 1e18, 1e15); // Prices still sum to 1
    }

    function test_EdgeCase_MultipleSequentialTrades() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        uint256 k1 = amm.reserveYes() * amm.reserveNo();

        // 10 sequential trades
        for (uint256 i = 0; i < 10; i++) {
            vm.prank(trader1);
            amm.buyYes(10 ether, 0);
        }

        uint256 k2 = amm.reserveYes() * amm.reserveNo();

        // Invariant should still hold
        assertApproxEqRel(k2, k1, 0.001e18); // 0.1% tolerance for multiple trades
    }

    function test_EdgeCase_AddLiquidityToImbalancedPool() public {
        vm.prank(lp1);
        amm.addLiquidity(10000 ether);

        // Imbalance the pool
        vm.prank(trader1);
        amm.buyYes(2000 ether, 0);

        uint256 reserveYesBefore = amm.reserveYes();
        uint256 reserveNoBefore = amm.reserveNo();

        // Add liquidity to imbalanced pool
        vm.prank(lp2);
        amm.addLiquidity(5000 ether);

        // Reserves should scale proportionally
        uint256 reserveYesAfter = amm.reserveYes();
        uint256 reserveNoAfter = amm.reserveNo();

        assertApproxEqRel(reserveYesAfter * reserveNoBefore, reserveNoAfter * reserveYesBefore, 0.01e18);
    }

    function test_EdgeCase_RemoveLiquidityFromImbalancedPool() public {
        vm.prank(lp1);
        uint256 lpTokens = amm.addLiquidity(10000 ether);

        // Imbalance the pool
        vm.prank(trader1);
        amm.buyYes(2000 ether, 0);

        uint256 totalCollateralBeforeRemove = amm.totalCollateral();

        // Remove half of liquidity
        vm.prank(lp1);
        uint256 collateralOut = amm.removeLiquidity(lpTokens / 2);

        // Should get back roughly half of the current total collateral
        // The pool now has more collateral than initial (10K + 2K trade minus fees)
        assertGt(collateralOut, 0);
        assertApproxEqRel(collateralOut, totalCollateralBeforeRemove / 2, 0.02e18); // 2% tolerance
    }
}
