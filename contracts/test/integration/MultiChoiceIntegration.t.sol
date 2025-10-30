// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/markets/MultiChoiceMarket.sol";
import "../../src/OutcomeToken.sol";

import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";
import "../../src/ResolutionModule.sol";
import "../mocks/MockERC20.sol";

/**
 * @title MultiChoiceIntegrationTest
 * @notice Comprehensive integration tests for MultiChoiceMarket with full system
 * @dev Tests the complete lifecycle: Creation → Liquidity → Trading → Resolution → Redemption
 */
contract MultiChoiceIntegrationTest is Test {
    // Contracts
    MultiChoiceMarket public market;
    OutcomeToken public outcomeToken;
    MockERC20 public horizonToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    ResolutionModule public resolution;
    MockERC20 public collateral;

    // Actors
    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator = address(0x2);
    address public lp1 = address(0x3);
    address public lp2 = address(0x4);
    address public trader1 = address(0x5);
    address public trader2 = address(0x6);
    address public trader3 = address(0x7);
    address public arbitrator = address(0x8);

    // Market parameters
    uint256 public constant MARKET_ID = 1;
    uint256 public constant OUTCOME_COUNT = 4;
    uint256 public constant LIQUIDITY_PARAM = 1000 ether;
    uint256 public closeTime;

    // Test constants
    uint256 public constant INITIAL_LP = 10_000 ether;
    uint256 public constant TRADER_BALANCE = 5_000 ether;

    function setUp() public {
        // Deploy core system contracts
        collateral = new MockERC20("USDC", "USDC");
        horizonToken = new MockERC20("Horizon Token", "HORIZON"); horizonToken.mint(address(this), 1_000_000_000 * 10 ** 18);
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        feeSplitter = new FeeSplitter(treasury);
        horizonPerks = new HorizonPerks(address(horizonToken));
        resolution = new ResolutionModule(
            address(outcomeToken),
            address(horizonToken),
            arbitrator
        );

        // Register market in outcome token
        outcomeToken.registerMarket(MARKET_ID, collateral);
        feeSplitter.registerMarket(MARKET_ID, creator);

        // Deploy MultiChoiceMarket
        closeTime = block.timestamp + 30 days;
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

        // Authorize contracts
        outcomeToken.setAMMAuthorization(address(market), true);
        outcomeToken.setResolutionAuthorization(address(resolution), true);

        // Fund actors with collateral
        collateral.mint(lp1, 50_000 ether);
        collateral.mint(lp2, 50_000 ether);
        collateral.mint(trader1, TRADER_BALANCE);
        collateral.mint(trader2, TRADER_BALANCE);
        collateral.mint(trader3, TRADER_BALANCE);

        // Fund creator with HORIZON for resolution
        horizonToken.transfer(creator, 10_000 ether);

        // Approve market for all actors
        vm.prank(lp1);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(lp2);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(trader1);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(trader2);
        collateral.approve(address(market), type(uint256).max);
        vm.prank(trader3);
        collateral.approve(address(market), type(uint256).max);

        // Approve resolution module
        vm.prank(creator);
        horizonToken.approve(address(resolution), type(uint256).max);
    }

    // ============ Full Lifecycle Tests ============

    /**
     * @notice Test complete market lifecycle with 4 outcomes
     */
    function test_Integration_FullLifecycle_Outcome0Wins() public {
        console.log("\n=== MULTICHOICE MARKET INTEGRATION TEST ===");
        console.log("Market ID: 1");
        console.log("Outcomes: 4");
        
        // ===== PHASE 1: LIQUIDITY PROVISION =====
        console.log("\n=== PHASE 1: LIQUIDITY PROVISION ===");
        
        vm.prank(lp1);
        uint256 lpTokens = market.addLiquidity(INITIAL_LP);
        console.log("LP1 provided: 10000 tokens");
        console.log("LP1 received LP tokens:");
        console.log(lpTokens / 1e18);
        assertGt(lpTokens, 0, "Should receive LP tokens");
        
        // Check initial prices (should be equal)
        uint256[] memory initialPrices = market.getAllPrices();
        console.log("\nInitial prices (should be equal ~0.25):");
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            console.log("  Outcome price / 1000:");
            console.log(initialPrices[i] / 1e15);
        }
        
        // ===== PHASE 2: TRADING =====
        console.log("\n=== PHASE 2: TRADING ===");
        
        // Trader1 buys outcome 0
        vm.prank(trader1);
        uint256 tokens0 = market.buy(0, 1000 ether, 0);
        console.log("Trader1 bought tokens of outcome 0:");
        console.log(tokens0 / 1e18);
        
        // Trader2 buys outcome 1
        vm.prank(trader2);
        uint256 tokens1 = market.buy(1, 800 ether, 0);
        console.log("Trader2 bought tokens of outcome 1:");
        console.log(tokens1 / 1e18);
        
        // Trader3 buys outcome 0 (betting same as trader1)
        vm.prank(trader3);
        uint256 tokens0_2 = market.buy(0, 1200 ether, 0);
        console.log("Trader3 bought tokens of outcome 0:");
        console.log(tokens0_2 / 1e18);
        
        // Check prices after trading
        uint256[] memory tradingPrices = market.getAllPrices();
        console.log("\nPrices after trading / 1000:");
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            console.log(tradingPrices[i] / 1e15);
        }
        
        // Outcome 0 should have highest price (most bought)
        assertGt(tradingPrices[0], tradingPrices[1], "Outcome 0 should be highest");
        assertGt(tradingPrices[0], tradingPrices[2], "Outcome 0 should be highest");
        assertGt(tradingPrices[0], tradingPrices[3], "Outcome 0 should be highest");
        
        // ===== PHASE 3: FEE ACCUMULATION =====
        console.log("\n=== PHASE 3: FEE VERIFICATION ===");
        
        // Claim fees
        vm.prank(treasury);
        feeSplitter.claimProtocolFees(address(collateral));
        
        vm.prank(creator);
        feeSplitter.claimCreatorFees(MARKET_ID, address(collateral));
        
        uint256 treasuryBalance = collateral.balanceOf(treasury);
        uint256 creatorBalance = collateral.balanceOf(creator);
        console.log("Treasury fees:");
        console.log(treasuryBalance / 1e18);
        console.log("Creator fees:");
        console.log(creatorBalance / 1e18);
        assertGt(treasuryBalance, 0, "Treasury should have fees");
        assertGt(creatorBalance, 0, "Creator should have fees");
        
        // ===== PHASE 4: MARKET CLOSE =====
        console.log("\n=== PHASE 4: MARKET CLOSE ===");
        
        vm.warp(closeTime + 1);
        console.log("Market closed");
        
        // Trading should revert after close
        vm.expectRevert();
        vm.prank(trader1);
        market.buy(0, 100 ether, 0);
        
        // ===== PHASE 5: RESOLUTION =====
        console.log("\n=== PHASE 5: RESOLUTION ===");
        
        // Propose outcome 0 as winner
        vm.prank(creator);
        resolution.proposeResolution(MARKET_ID, 0, 1000 ether, "Outcome 0 won as predicted");
        console.log("Resolution proposed: Outcome 0 wins");
        
        // Fast forward past arbitration period
        vm.warp(block.timestamp + 3 days);
        
        // Finalize resolution
        vm.prank(creator);
        resolution.finalize(MARKET_ID);
        console.log("Resolution finalized");
        
        // Verify resolution
        assertTrue(outcomeToken.isResolved(MARKET_ID), "Market should be resolved");
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0, "Outcome 0 should be winner");
        
        // ===== PHASE 6: FUND REDEMPTIONS =====
        console.log("\n=== PHASE 6: FUND REDEMPTIONS ===");
        
        uint256 marketCollateral = market.totalCollateral();
        console.log("Market collateral before funding:");
        console.log(marketCollateral / 1e18);
        
        vm.prank(lp1);
        market.fundRedemptions();
        console.log("Redemptions funded by LP1");
        
        uint256 redemptionPool = collateral.balanceOf(address(outcomeToken));
        console.log("Redemption pool:");
        console.log(redemptionPool / 1e18);
        assertGt(redemptionPool, 0, "Should have redemption pool");
        
        // ===== PHASE 7: REDEEM WINNING TOKENS =====
        console.log("\n=== PHASE 7: REDEEM WINNING TOKENS ===");
        
        // Trader1 redeems (has outcome 0 tokens - winning outcome)
        uint256 trader1BalanceBefore = collateral.balanceOf(trader1);
        uint256 tokenId0 = outcomeToken.encodeTokenId(MARKET_ID, 0);
        uint256 trader1Tokens = outcomeToken.balanceOf(trader1, tokenId0);
        console.log("Trader1 outcome 0 tokens:");
        console.log(trader1Tokens / 1e18);
        
        vm.prank(trader1);
        outcomeToken.redeem(MARKET_ID);
        
        uint256 trader1BalanceAfter = collateral.balanceOf(trader1);
        uint256 trader1Profit = trader1BalanceAfter - trader1BalanceBefore;
        console.log("Trader1 redeemed for:");
        console.log(trader1Profit / 1e18);
        assertGt(trader1Profit, 0, "Winner should receive collateral");
        
        // Trader3 also redeems (also has outcome 0)
        uint256 trader3BalanceBefore = collateral.balanceOf(trader3);
        uint256 trader3Tokens = outcomeToken.balanceOf(trader3, tokenId0);
        
        vm.prank(trader3);
        outcomeToken.redeem(MARKET_ID);
        
        uint256 trader3BalanceAfter = collateral.balanceOf(trader3);
        uint256 trader3Profit = trader3BalanceAfter - trader3BalanceBefore;
        console.log("Trader3 redeemed for:");
        console.log(trader3Profit / 1e18);
        assertGt(trader3Profit, 0, "Winner should receive collateral");
        
        // Trader2 tries to redeem losing outcome (should revert with NoTokensToRedeem)
        uint256 tokenId1 = outcomeToken.encodeTokenId(MARKET_ID, 1);
        uint256 trader2Tokens = outcomeToken.balanceOf(trader2, tokenId1);
        console.log("Trader2 outcome 1 tokens (losing):");
        console.log(trader2Tokens / 1e18);
        
        // Trader2 has losing tokens, so redeem should fail
        vm.expectRevert();
        vm.prank(trader2);
        outcomeToken.redeem(MARKET_ID);
        
        console.log("\n=== TEST COMPLETE ===");
        console.log("Note: After fundRedemptions(), all collateral goes to redemption pool.");
        console.log("LP cannot withdraw after funding - collateral goes to winning token holders.");
    }

    /**
     * @notice Test multiple LPs providing liquidity at different times
     */
    function test_Integration_MultipleLPs() public {
        console.log("\n=== MULTIPLE LPs TEST ===");
        
        // LP1 adds initial liquidity
        vm.prank(lp1);
        uint256 lp1Tokens = market.addLiquidity(INITIAL_LP);
        console.log("LP1 added 10000 collateral");
        
        // Some trading happens
        vm.prank(trader1);
        market.buy(0, 1000 ether, 0);
        
        vm.prank(trader2);
        market.buy(1, 800 ether, 0);
        
        // LP2 adds liquidity later (should get proportional LP tokens)
        vm.prank(lp2);
        uint256 lp2Tokens = market.addLiquidity(INITIAL_LP);
        console.log("LP2 added 10000 collateral");
        
        // More trading
        vm.prank(trader3);
        market.buy(2, 1200 ether, 0);
        
        // Check LP token balances
        uint256 lp1Balance = market.balanceOf(lp1);
        uint256 lp2Balance = market.balanceOf(lp2);
        console.log("LP1 tokens:");
        console.log(lp1Balance / 1e18);
        console.log("LP2 tokens:");
        console.log(lp2Balance / 1e18);
        
        // LP1 should have more tokens (provided first)
        assertGt(lp1Balance, lp2Balance, "LP1 should have more LP tokens");
        
        // Both should be able to withdraw
        vm.prank(lp1);
        uint256 lp1Out = market.removeLiquidity(lp1Balance);
        console.log("LP1 withdrew:");
        console.log(lp1Out / 1e18);
        
        vm.prank(lp2);
        uint256 lp2Out = market.removeLiquidity(lp2Balance);
        console.log("LP2 withdrew:");
        console.log(lp2Out / 1e18);
        
        assertGt(lp1Out, 0, "LP1 should withdraw successfully");
        assertGt(lp2Out, 0, "LP2 should withdraw successfully");
    }

    /**
     * @notice Test fee distribution across treasury and creator
     */
    function test_Integration_FeeDistribution() public {
        console.log("\n=== FEE DISTRIBUTION TEST ===");
        
        // Add liquidity
        vm.prank(lp1);
        market.addLiquidity(INITIAL_LP);
        
        uint256 treasuryBefore = collateral.balanceOf(treasury);
        uint256 creatorBefore = collateral.balanceOf(creator);
        
        // Execute multiple trades to generate fees
        vm.prank(trader1);
        market.buy(0, 1000 ether, 0);
        
        vm.prank(trader2);
        market.buy(1, 1000 ether, 0);
        
        vm.prank(trader3);
        market.buy(2, 1000 ether, 0);
        
        uint256 feeSplitterBalance = collateral.balanceOf(address(feeSplitter));
        console.log("FeeSplitter balance before claims:");
        console.log(feeSplitterBalance / 1e18);
        
        // Claim fees
        vm.prank(treasury);
        feeSplitter.claimProtocolFees(address(collateral));
        
        vm.prank(creator);
        feeSplitter.claimCreatorFees(MARKET_ID, address(collateral));
        
        uint256 treasuryAfter = collateral.balanceOf(treasury);
        uint256 creatorAfter = collateral.balanceOf(creator);
        
        uint256 treasuryFees = treasuryAfter - treasuryBefore;
        uint256 creatorFees = creatorAfter - creatorBefore;
        
        console.log("Treasury fees:");
        console.log(treasuryFees / 1e18);
        console.log("Creator fees:");
        console.log(creatorFees / 1e18);
        
        assertGt(treasuryFees, 0, "Treasury should collect fees");
        assertGt(creatorFees, 0, "Creator should collect fees");
        
        // Creator should get more than treasury (default split is 10% protocol / 90% creator)
        assertGt(creatorFees, treasuryFees, "Creator should get larger share");
    }

    /**
     * @notice Test LMSR invariants hold throughout full lifecycle
     */
    function test_Integration_LMSRInvariants() public {
        console.log("\n=== LMSR INVARIANTS TEST ===");
        
        // Add liquidity
        vm.prank(lp1);
        market.addLiquidity(INITIAL_LP);
        
        // Execute various trades and check invariants after each
        uint256[] memory prices;
        uint256 priceSum;
        
        // Initial state
        prices = market.getAllPrices();
        priceSum = 0;
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            priceSum += prices[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should sum to 1.0");
        
        // After trade 1
        vm.prank(trader1);
        market.buy(0, 500 ether, 0);
        
        prices = market.getAllPrices();
        priceSum = 0;
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            priceSum += prices[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should still sum to 1.0");
        
        // After trade 2
        vm.prank(trader2);
        market.buy(1, 800 ether, 0);
        
        prices = market.getAllPrices();
        priceSum = 0;
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            priceSum += prices[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should still sum to 1.0");
        
        // After trade 3 (different outcome)
        vm.prank(trader3);
        market.buy(3, 600 ether, 0);
        
        prices = market.getAllPrices();
        priceSum = 0;
        for (uint256 i = 0; i < OUTCOME_COUNT; i++) {
            priceSum += prices[i];
        }
        assertApproxEqAbs(priceSum, 1e18, 1e15, "Prices should still sum to 1.0");
        
        console.log("All LMSR invariants maintained throughout trading");
    }

    /**
     * @notice Test pausing and unpausing functionality
     */
    function test_Integration_PauseUnpause() public {
        console.log("\n=== PAUSE/UNPAUSE TEST ===");
        
        // Add liquidity
        vm.prank(lp1);
        market.addLiquidity(INITIAL_LP);
        
        // Normal trade should work
        vm.prank(trader1);
        market.buy(0, 100 ether, 0);
        console.log("Trade 1 succeeded");
        
        // Pause market
        market.pause();
        console.log("Market paused");
        
        // Trades should revert
        vm.expectRevert();
        vm.prank(trader2);
        market.buy(1, 100 ether, 0);
        
        // Liquidity provision should also revert
        vm.expectRevert();
        vm.prank(lp2);
        market.addLiquidity(1000 ether);
        
        // Unpause
        market.unpause();
        console.log("Market unpaused");
        
        // Trading should work again
        vm.prank(trader2);
        market.buy(1, 100 ether, 0);
        console.log("Trade 2 succeeded after unpause");
    }

    /**
     * @notice Test selling functionality in full context
     */
    function test_Integration_BuyAndSell() public {
        console.log("\n=== BUY AND SELL TEST ===");
        
        // Add liquidity
        vm.prank(lp1);
        market.addLiquidity(INITIAL_LP);
        
        // Trader buys tokens
        vm.prank(trader1);
        uint256 tokensBought = market.buy(0, 1000 ether, 0);
        console.log("Bought tokens:");
        console.log(tokensBought / 1e18);
        
        uint256 priceBefore = market.getPrice(0);
        console.log("Price before sell / 1000:");
        console.log(priceBefore / 1e15);
        
        // Trader sells half
        vm.prank(trader1);
        uint256 collateralReceived = market.sell(0, tokensBought / 2, 0);
        console.log("Sold half for collateral:");
        console.log(collateralReceived / 1e18);
        
        uint256 priceAfter = market.getPrice(0);
        console.log("Price after sell / 1000:");
        console.log(priceAfter / 1e15);
        
        // Price should have decreased
        assertLt(priceAfter, priceBefore, "Selling should decrease price");
        
        // Trader should have received less than they paid (fees + slippage)
        assertLt(collateralReceived, 500 ether, "Should lose money on round trip");
    }
}
