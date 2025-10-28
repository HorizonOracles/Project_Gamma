// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/OutcomeToken.sol";
import "../../src/HorizonToken.sol";
import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";
import "../../src/MarketAMM.sol";
import "../../src/ResolutionModule.sol";
import "../mocks/MockERC20.sol";

/**
 * @title EndToEndTest
 * @notice Comprehensive integration test covering the full market lifecycle
 */
contract EndToEndTest is Test {
    // Contracts
    OutcomeToken public outcomeToken;
    HorizonToken public horizonToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    MarketAMM public amm;
    ResolutionModule public resolution;
    MockERC20 public collateral;

    // Actors
    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator = address(0x2);
    address public liquidityProvider = address(0x3);
    address public trader1 = address(0x4); // Will buy YES
    address public trader2 = address(0x5); // Will buy NO
    address public arbitrator = address(0x6);

    uint256 public constant MARKET_ID = 1;
    uint256 public closeTime;

    function setUp() public {
        // Deploy all contracts
        collateral = new MockERC20("USDC", "USDC");
        horizonToken = new HorizonToken(1_000_000_000 * 10 ** 18);
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        feeSplitter = new FeeSplitter(treasury);
        horizonPerks = new HorizonPerks(address(horizonToken));
        resolution = new ResolutionModule(address(outcomeToken), address(horizonToken), arbitrator);

        // Register market
        outcomeToken.registerMarket(MARKET_ID, collateral);
        feeSplitter.registerMarket(MARKET_ID, creator);

        // Deploy AMM
        closeTime = block.timestamp + 30 days;
        amm = new MarketAMM(
            MARKET_ID, address(collateral), address(outcomeToken), address(feeSplitter), address(horizonPerks), closeTime
        );

        // Authorize contracts
        outcomeToken.setAMMAuthorization(address(amm), true);
        outcomeToken.setResolutionAuthorization(address(resolution), true);

        // Fund actors with collateral
        collateral.mint(liquidityProvider, 100_000 ether);
        collateral.mint(trader1, 10_000 ether);
        collateral.mint(trader2, 10_000 ether);

        // Fund resolution proposer with HORIZON
        horizonToken.transfer(creator, 10_000 ether);

        // Approve contracts
        vm.prank(liquidityProvider);
        collateral.approve(address(amm), type(uint256).max);

        vm.prank(trader1);
        collateral.approve(address(amm), type(uint256).max);

        vm.prank(trader2);
        collateral.approve(address(amm), type(uint256).max);

        vm.prank(creator);
        horizonToken.approve(address(resolution), type(uint256).max);
    }

    /**
     * @notice Full lifecycle test: Liquidity → Trading → Resolution → Claiming
     */
    function test_FullMarketLifecycle_YesWins() public {
        // ===== PHASE 1: ADD LIQUIDITY =====
        console.log("=== PHASE 1: ADD LIQUIDITY ===");

        vm.prank(liquidityProvider);
        uint256 lpTokens = amm.addLiquidity(10_000 ether);

        console.log("LP tokens received:", lpTokens / 1e18);
        assertGt(lpTokens, 0);

        // ===== PHASE 2: TRADING =====
        console.log("\n=== PHASE 2: TRADING ===");

        // Trader1 buys YES (500 collateral)
        uint256 trader1CollateralBefore = collateral.balanceOf(trader1);
        vm.prank(trader1);
        uint256 trader1YesTokens = amm.buyYes(500 ether, 0);
        console.log("Trader1 bought YES tokens:", trader1YesTokens / 1e18);

        // Trader2 buys NO (300 collateral)
        vm.prank(trader2);
        uint256 trader2NoTokens = amm.buyNo(300 ether, 0);
        console.log("Trader2 bought NO tokens:", trader2NoTokens / 1e18);

        // Check outcome token balances
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 0), trader1YesTokens);
        assertEq(outcomeToken.balanceOfOutcome(trader2, MARKET_ID, 1), trader2NoTokens);

        // ===== PHASE 3: MARKET CLOSES =====
        console.log("\n=== PHASE 3: MARKET CLOSES ===");
        vm.warp(closeTime + 1);

        // ===== PHASE 4: RESOLUTION =====
        console.log("\n=== PHASE 4: RESOLUTION ===");

        // Creator proposes YES (outcome 0) wins
        vm.prank(creator);
        resolution.proposeResolution(MARKET_ID, 0, 1000 ether, "ipfs://evidence");
        console.log("Resolution proposed: YES wins");

        // Wait for dispute window
        vm.warp(block.timestamp + 48 hours + 1);

        // Finalize resolution
        resolution.finalize(MARKET_ID);
        console.log("Resolution finalized");

        // Verify market is resolved
        assertTrue(outcomeToken.isResolved(MARKET_ID));
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0); // YES

        // Fund OutcomeToken with collateral for redemptions
        amm.fundRedemptions();
        console.log("AMM collateral transferred to OutcomeToken for redemptions");

        // ===== PHASE 5: WINNERS CLAIM =====
        console.log("\n=== PHASE 5: WINNERS CLAIM ===");

        // Trader1 (YES holder) claims
        uint256 trader1BalanceBefore = collateral.balanceOf(trader1);
        uint256 trader1Spent = trader1CollateralBefore - trader1BalanceBefore; // How much was spent on trading

        vm.prank(trader1);
        uint256 payout = outcomeToken.redeem(MARKET_ID);
        uint256 trader1BalanceAfter = collateral.balanceOf(trader1);

        console.log("Trader1 payout:", payout / 1e18);
        console.log("Trader1 spent:", trader1Spent / 1e18);
        int256 profit = int256(payout) - int256(trader1Spent);
        console.log("Trader1 profit:", profit > 0 ? uint256(profit) / 1e18 : 0);

        // Verify payout
        assertEq(payout, trader1YesTokens); // 1:1 redemption
        assertEq(trader1BalanceAfter, trader1BalanceBefore + payout);

        // Trader1's tokens should be burned
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 0), 0);

        // ===== PHASE 6: LOSERS CAN'T CLAIM =====
        console.log("\n=== PHASE 6: LOSERS CAN'T CLAIM ===");

        // Trader2 (NO holder) cannot redeem
        vm.prank(trader2);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(MARKET_ID);

        console.log("Trader2 (loser) correctly cannot claim");

        // ===== VERIFICATION =====
        console.log("\n=== FINAL VERIFICATION ===");
        console.log("Market resolved:", outcomeToken.isResolved(MARKET_ID));
        console.log("Winning outcome: YES (0)");
        console.log("Winner claimed successfully");
    }

    /**
     * @notice Test resolution with dispute - disputer wins
     */
    function test_FullMarketLifecycle_WithDispute() public {
        // Setup: LP and trading
        vm.prank(liquidityProvider);
        amm.addLiquidity(10_000 ether);

        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(500 ether, 0);

        vm.prank(trader2);
        uint256 noTokens = amm.buyNo(300 ether, 0);

        // Close market
        vm.warp(closeTime + 1);

        // Creator proposes YES wins (incorrectly)
        horizonToken.transfer(creator, 10_000 ether); // More for dispute
        vm.prank(creator);
        resolution.proposeResolution(MARKET_ID, 0, 1000 ether, "ipfs://wrong-evidence");

        // Someone disputes
        address disputer = address(0x99);
        horizonToken.transfer(disputer, 10_000 ether);
        vm.prank(disputer);
        horizonToken.approve(address(resolution), type(uint256).max);

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, 1000 ether, "This is wrong!");

        // Arbitrator sides with disputer - NO actually won
        vm.prank(arbitrator);
        resolution.finalizeDisputed(MARKET_ID, 1, true); // outcome 1 (NO), slash proposer

        // Verify correct outcome
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 1); // NO

        // Fund redemptions
        amm.fundRedemptions();

        // Trader2 (NO holder) can now claim
        uint256 trader2BalanceBefore = collateral.balanceOf(trader2);
        vm.prank(trader2);
        uint256 payout = outcomeToken.redeem(MARKET_ID);

        assertEq(payout, noTokens);
        assertEq(collateral.balanceOf(trader2), trader2BalanceBefore + payout);

        // Trader1 (YES holder) cannot claim
        vm.prank(trader1);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(MARKET_ID);
    }

    /**
     * @notice Test multiple winners claiming proportionally
     */
    function test_MultipleWinnersClaim() public {
        // Setup liquidity
        vm.prank(liquidityProvider);
        amm.addLiquidity(20_000 ether);

        // Multiple traders buy YES
        address trader3 = address(0x7);
        address trader4 = address(0x8);

        collateral.mint(trader3, 10_000 ether);
        collateral.mint(trader4, 10_000 ether);

        vm.prank(trader3);
        collateral.approve(address(amm), type(uint256).max);
        vm.prank(trader4);
        collateral.approve(address(amm), type(uint256).max);

        // All buy YES
        vm.prank(trader1);
        uint256 yes1 = amm.buyYes(500 ether, 0);

        vm.prank(trader3);
        uint256 yes3 = amm.buyYes(1000 ether, 0);

        vm.prank(trader4);
        uint256 yes4 = amm.buyYes(200 ether, 0);

        // Resolve to YES
        vm.warp(closeTime + 1);
        vm.prank(creator);
        resolution.proposeResolution(MARKET_ID, 0, 1000 ether, "ipfs://evidence");
        vm.warp(block.timestamp + 48 hours + 1);
        resolution.finalize(MARKET_ID);

        // Fund redemptions
        amm.fundRedemptions();

        // All winners claim their proportional share
        vm.prank(trader1);
        uint256 payout1 = outcomeToken.redeem(MARKET_ID);
        assertEq(payout1, yes1);

        vm.prank(trader3);
        uint256 payout3 = outcomeToken.redeem(MARKET_ID);
        assertEq(payout3, yes3);

        vm.prank(trader4);
        uint256 payout4 = outcomeToken.redeem(MARKET_ID);
        assertEq(payout4, yes4);

        console.log("All winners claimed proportionally");
    }

    /**
     * @notice Test partial redemption
     */
    function test_PartialRedemption() public {
        // Setup and trading
        vm.prank(liquidityProvider);
        amm.addLiquidity(10_000 ether);

        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(1000 ether, 0);

        // Resolve to YES
        vm.warp(closeTime + 1);
        vm.prank(creator);
        resolution.proposeResolution(MARKET_ID, 0, 1000 ether, "ipfs://evidence");
        vm.warp(block.timestamp + 48 hours + 1);
        resolution.finalize(MARKET_ID);

        // Fund redemptions
        amm.fundRedemptions();

        // Trader1 redeems in 3 parts
        vm.prank(trader1);
        uint256 payout1 = outcomeToken.redeemAmount(MARKET_ID, yesTokens / 3);

        vm.prank(trader1);
        uint256 payout2 = outcomeToken.redeemAmount(MARKET_ID, yesTokens / 3);

        vm.prank(trader1);
        uint256 payout3 = outcomeToken.redeem(MARKET_ID); // Redeem rest

        assertEq(payout1 + payout2 + payout3, yesTokens);
        assertEq(outcomeToken.balanceOfOutcome(trader1, MARKET_ID, 0), 0);
    }
}
