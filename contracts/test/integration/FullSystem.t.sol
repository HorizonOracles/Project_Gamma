// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/MarketFactory.sol";
import "../../src/OutcomeToken.sol";
import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";

import "../../src/ResolutionModule.sol";
import "../../src/markets/BinaryMarket.sol";
import "../mocks/MockERC20.sol";

/**
 * @title FullSystemTest
 * @notice Comprehensive integration tests using MarketFactory for complete system orchestration
 * @dev Tests the full lifecycle: Factory → AMM → Resolution → Redemption with fees
 */
contract FullSystemTest is Test {
    // Core contracts
    MarketFactory public factory;
    OutcomeToken public outcomeToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    MockERC20 public horizonToken;
    ResolutionModule public resolution;
    MockERC20 public usdc;
    MockERC20 public dai;

    // Actors
    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator1 = address(0x2);
    address public creator2 = address(0x3);
    address public liquidityProvider = address(0x4);
    address public trader1 = address(0x5);
    address public trader2 = address(0x6);
    address public trader3 = address(0x7);
    address public arbitrator = address(0x8);

    uint256 public constant MIN_STAKE = 100 ether;

    event MarketCreated(
        uint256 indexed marketId,
        address indexed creator,
        address indexed ammAddress,
        address collateralToken,
        uint256 closeTime,
        string category,
        string metadataURI,
        uint256 creatorStake
    );

    function setUp() public {
        // Deploy base contracts
        usdc = new MockERC20("USDC", "USDC");
        dai = new MockERC20("DAI", "DAI");
        horizonToken = new MockERC20("Horizon Token", "HORIZON"); horizonToken.mint(address(this), 1_000_000_000 * 10 ** 18);
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        feeSplitter = new FeeSplitter(treasury);
        horizonPerks = new HorizonPerks(address(horizonToken));
        resolution = new ResolutionModule(address(outcomeToken), address(horizonToken), arbitrator);

        // Set resolution authorization
        outcomeToken.setResolutionAuthorization(address(resolution), true);

        // Deploy factory
        factory = new MarketFactory(
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            address(horizonToken)
        );

        // Transfer ownership to factory
        outcomeToken.transferOwnership(address(factory));
        feeSplitter.transferOwnership(address(factory));

        // Fund actors with HORIZON tokens
        horizonToken.transfer(creator1, 10_000 ether);
        horizonToken.transfer(creator2, 10_000 ether);

        // Fund actors with collateral
        usdc.mint(liquidityProvider, 100_000 ether);
        usdc.mint(trader1, 10_000 ether);
        usdc.mint(trader2, 10_000 ether);
        usdc.mint(trader3, 10_000 ether);

        dai.mint(liquidityProvider, 100_000 ether);
        dai.mint(trader1, 10_000 ether);

        // Approvals
        vm.prank(creator1);
        horizonToken.approve(address(factory), type(uint256).max);
        vm.prank(creator1);
        horizonToken.approve(address(resolution), type(uint256).max);

        vm.prank(creator2);
        horizonToken.approve(address(factory), type(uint256).max);
        vm.prank(creator2);
        horizonToken.approve(address(resolution), type(uint256).max);

        vm.prank(liquidityProvider);
        usdc.approve(address(0xdead), type(uint256).max); // Will approve actual AMM later

        vm.prank(trader1);
        usdc.approve(address(0xdead), type(uint256).max);

        vm.prank(trader2);
        usdc.approve(address(0xdead), type(uint256).max);

        vm.prank(trader3);
        usdc.approve(address(0xdead), type(uint256).max);
    }

    /**
     * @notice Full system test: Factory creates market → LP adds liquidity → Trading → Resolution → Claiming
     */
    function test_FullSystemLifecycle() public {
        console.log("\n=== PHASE 1: MARKET CREATION ===");

        // Create market via factory
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0, // Binary
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmExample123",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        console.log("Market created with ID:", marketId);
        assertEq(marketId, 1);

        // Get market details
        MarketFactory.Market memory market = factory.getMarket(marketId);
        address ammAddress = market.amm;
        assertNotEq(ammAddress, address(0));
        console.log("AMM deployed at:", ammAddress);

        BinaryMarket amm = BinaryMarket(ammAddress);

        // Approve AMM
        vm.prank(liquidityProvider);
        usdc.approve(ammAddress, type(uint256).max);
        vm.prank(trader1);
        usdc.approve(ammAddress, type(uint256).max);
        vm.prank(trader2);
        usdc.approve(ammAddress, type(uint256).max);

        console.log("\n=== PHASE 2: LIQUIDITY PROVISION ===");

        // Add liquidity
        vm.prank(liquidityProvider);
        uint256 lpTokens = amm.addLiquidity(10_000 ether);
        console.log("LP tokens received:", lpTokens / 1e18);
        assertGt(lpTokens, 0);

        console.log("\n=== PHASE 3: TRADING ===");

        // Trader1 buys YES
        uint256 trader1BalanceBefore = usdc.balanceOf(trader1);
        vm.prank(trader1);
        uint256 yesTokens = amm.buyYes(500 ether, 0);
        console.log("Trader1 bought YES:", yesTokens / 1e18, "for 500 USDC");

        // Trader2 buys NO
        vm.prank(trader2);
        uint256 noTokens = amm.buyNo(300 ether, 0);
        console.log("Trader2 bought NO:", noTokens / 1e18, "for 300 USDC");

        // Verify balances
        assertEq(outcomeToken.balanceOfOutcome(trader1, marketId, 0), yesTokens);
        assertEq(outcomeToken.balanceOfOutcome(trader2, marketId, 1), noTokens);

        // Verify fees were collected
        uint256 treasuryFees = feeSplitter.protocolPendingFees(address(usdc));
        uint256 creatorFees = feeSplitter.creatorPendingFees(marketId, address(usdc));
        console.log("Treasury fees collected:", treasuryFees / 1e18, "USDC");
        console.log("Creator fees collected:", creatorFees / 1e18, "USDC");
        assertGt(treasuryFees, 0);
        assertGt(creatorFees, 0);

        console.log("\n=== PHASE 4: MARKET CLOSES ===");

        vm.warp(params.closeTime + 1);
        factory.updateMarketStatus(marketId);
        market = factory.getMarket(marketId);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Closed));
        console.log("Market status: Closed");

        console.log("\n=== PHASE 5: RESOLUTION ===");

        // Propose resolution: YES wins
        vm.prank(creator1);
        resolution.proposeResolution(marketId, 0, 1000 ether, "ipfs://evidence-yes-wins");
        console.log("Resolution proposed: YES wins");

        // Wait for dispute window
        vm.warp(block.timestamp + 48 hours + 1);

        // Finalize
        resolution.finalize(marketId);
        console.log("Resolution finalized");

        factory.updateMarketStatus(marketId);
        market = factory.getMarket(marketId);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Resolved));

        // Fund redemptions
        amm.fundRedemptions();
        console.log("AMM collateral transferred for redemptions");

        console.log("\n=== PHASE 6: CLAIMING ===");

        // Winner (Trader1) claims
        uint256 trader1Spent = trader1BalanceBefore - usdc.balanceOf(trader1);
        vm.prank(trader1);
        uint256 payout = outcomeToken.redeem(marketId);
        console.log("Trader1 payout:", payout / 1e18, "USDC");
        console.log("Trader1 spent:", trader1Spent / 1e18, "USDC");
        assertEq(payout, yesTokens);

        // Loser cannot claim
        vm.prank(trader2);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(marketId);
        console.log("Trader2 (loser) correctly blocked from claiming");

        console.log("\n=== PHASE 7: CREATOR STAKE REFUND ===");

        // Refund creator stake
        uint256 creatorBalanceBefore = horizonToken.balanceOf(creator1);
        factory.refundCreatorStake(marketId);
        assertEq(horizonToken.balanceOf(creator1), creatorBalanceBefore + MIN_STAKE);
        console.log("Creator stake refunded:", MIN_STAKE / 1e18, "HORIZON");

        console.log("\n=== PHASE 8: FEE CLAIMING ===");

        // Creator claims fees
        uint256 creatorUsdcBefore = usdc.balanceOf(creator1);
        vm.prank(creator1);
        feeSplitter.claimCreatorFees(marketId, address(usdc));
        uint256 creatorUsdcAfter = usdc.balanceOf(creator1);
        console.log("Creator claimed fees:", (creatorUsdcAfter - creatorUsdcBefore) / 1e18, "USDC");
        assertGt(creatorUsdcAfter, creatorUsdcBefore);

        // Treasury claims fees
        uint256 treasuryUsdcBefore = usdc.balanceOf(treasury);
        vm.prank(treasury);
        feeSplitter.claimProtocolFees(address(usdc));
        uint256 treasuryUsdcAfter = usdc.balanceOf(treasury);
        console.log("Treasury claimed fees:", (treasuryUsdcAfter - treasuryUsdcBefore) / 1e18, "USDC");
        assertGt(treasuryUsdcAfter, treasuryUsdcBefore);

        console.log("\n=== TEST COMPLETE ===");
    }

    /**
     * @notice Test multiple concurrent markets
     */
    function test_MultipleMarketsSimultaneous() public {
        console.log("\n=== CREATING MULTIPLE MARKETS ===");

        uint256[] memory marketIds = new uint256[](3);

        // Create 3 markets
        for (uint256 i = 0; i < 3; i++) {
            MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
                marketType: 0, // Binary
                collateralToken: address(usdc),
                closeTime: block.timestamp + (i + 1) * 7 days,
                category: i == 0 ? "politics" : i == 1 ? "sports" : "crypto",
                metadataURI: string(abi.encodePacked("ipfs://QmMarket", vm.toString(i))),
                creatorStake: MIN_STAKE,
                outcomeCount: 2,
                liquidityParameter: 0
            });

            vm.prank(i < 2 ? creator1 : creator2);
            marketIds[i] = factory.createMarket(params);
            console.log("Created market", i + 1, "with ID:", marketIds[i]);
        }

        assertEq(factory.getMarketCount(), 3);

        console.log("\n=== TRADING ON MULTIPLE MARKETS ===");

        // Trade on all markets
        for (uint256 i = 0; i < 3; i++) {
            MarketFactory.Market memory market = factory.getMarket(marketIds[i]);
            BinaryMarket amm = BinaryMarket(market.amm);

            // Approve
            vm.prank(liquidityProvider);
            usdc.approve(address(amm), type(uint256).max);
            vm.prank(trader1);
            usdc.approve(address(amm), type(uint256).max);

            // LP provides liquidity
            vm.prank(liquidityProvider);
            amm.addLiquidity(5_000 ether);

            // Trader1 buys YES
            vm.prank(trader1);
            uint256 yesTokens = amm.buyYes(100 ether, 0);
            console.log("Market %d - Trader1 bought %d YES tokens", i + 1, yesTokens / 1e18);
        }

        console.log("\n=== RESOLVING MARKETS INDEPENDENTLY ===");

        // Resolve market 1 immediately
        vm.warp(block.timestamp + 7 days + 1);
        vm.prank(creator1);
        resolution.proposeResolution(marketIds[0], 0, 1000 ether, "ipfs://evidence1");
        vm.warp(block.timestamp + 48 hours + 1);
        resolution.finalize(marketIds[0]);
        console.log("Market 1 resolved");

        // Market 2 and 3 still active
        MarketFactory.Market memory market2 = factory.getMarket(marketIds[1]);
        assertEq(uint8(market2.status), uint8(MarketFactory.MarketStatus.Active));

        // Resolve market 2
        vm.warp(block.timestamp + 7 days);
        vm.prank(creator1);
        resolution.proposeResolution(marketIds[1], 1, 1000 ether, "ipfs://evidence2");
        vm.warp(block.timestamp + 48 hours + 1);
        resolution.finalize(marketIds[1]);
        console.log("Market 2 resolved");

        // Verify independent states
        assertTrue(outcomeToken.isResolved(marketIds[0]));
        assertTrue(outcomeToken.isResolved(marketIds[1]));
        assertFalse(outcomeToken.isResolved(marketIds[2]));

        console.log("\n=== CLAIMING FROM RESOLVED MARKETS ===");

        // Fund and claim from market 1
        BinaryMarket(market2.amm).fundRedemptions();
        vm.prank(trader1);
        uint256 payout1 = outcomeToken.redeem(marketIds[0]);
        console.log("Claimed from market 1:", payout1 / 1e18, "USDC");

        // Market 2: Trader1 loses (NO won)
        MarketFactory.Market memory market1 = factory.getMarket(marketIds[1]);
        BinaryMarket(market1.amm).fundRedemptions();
        vm.prank(trader1);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(marketIds[1]);
        console.log("Market 2: Trader1 correctly cannot claim (lost)");

        console.log("\n=== TEST COMPLETE ===");
        console.log("Total markets created:", factory.getMarketCount());
        console.log("Creator1 markets:", factory.getMarketIdsByCreator(creator1).length);
        console.log("Creator2 markets:", factory.getMarketIdsByCreator(creator2).length);
    }

    /**
     * @notice Test fee flow across all components
     */
    function test_CompleteFeeFlow() public {
        console.log("\n=== TESTING COMPLETE FEE FLOW ===");

        // Create market
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0, // Binary
            collateralToken: address(usdc),
            closeTime: block.timestamp + 7 days,
            category: "test",
            metadataURI: "ipfs://QmFeeTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        BinaryMarket amm = BinaryMarket(market.amm);

        // Give trader1 some HORIZON tokens for better protocol/creator split
        horizonToken.transfer(trader1, 50_000 ether); // Tier 2: 2% fee, 6% protocol share
        uint256 trader1FeeBps = horizonPerks.feeBpsFor(trader1);
        assertEq(trader1FeeBps, 200); // 2.0%

        // Trader 2 has no HORIZON (tier 0: 2.0% fee, 10% protocol share)
        uint256 trader2FeeBps = horizonPerks.feeBpsFor(trader2);
        console.log("Trader2 fee tier (bps):", trader2FeeBps);
        assertEq(trader2FeeBps, 200); // 2.0%

        // Setup
        vm.prank(liquidityProvider);
        usdc.approve(address(amm), type(uint256).max);
        vm.prank(trader1);
        usdc.approve(address(amm), type(uint256).max);
        vm.prank(trader2);
        usdc.approve(address(amm), type(uint256).max);

        vm.prank(liquidityProvider);
        amm.addLiquidity(10_000 ether);

        console.log("\n=== TRADING WITH DIFFERENT FEE TIERS ===");

        // Trader1 trades (2% fee, lower protocol share)
        vm.prank(trader1);
        amm.buyYes(1000 ether, 0);

        uint256 treasuryFeesAfterT1 = feeSplitter.protocolPendingFees(address(usdc));
        uint256 creatorFeesAfterT1 = feeSplitter.creatorPendingFees(marketId, address(usdc));
        console.log("After Trader1 trade:");
        console.log("  Treasury fees:", treasuryFeesAfterT1 / 1e18, "USDC");
        console.log("  Creator fees:", creatorFeesAfterT1 / 1e18, "USDC");

        // Trader2 trades (2.0% fee)
        vm.prank(trader2);
        amm.buyNo(1000 ether, 0);

        uint256 treasuryFeesAfterT2 = feeSplitter.protocolPendingFees(address(usdc));
        uint256 creatorFeesAfterT2 = feeSplitter.creatorPendingFees(marketId, address(usdc));
        console.log("After Trader2 trade:");
        console.log("  Treasury fees:", treasuryFeesAfterT2 / 1e18, "USDC");
        console.log("  Creator fees:", creatorFeesAfterT2 / 1e18, "USDC");

        // Both traders pay same total fees (2%)
        uint256 trader2FeeIncrease = (treasuryFeesAfterT2 - treasuryFeesAfterT1) + (creatorFeesAfterT2 - creatorFeesAfterT1);
        uint256 trader1FeeTotal = treasuryFeesAfterT1 + creatorFeesAfterT1;
        console.log("Trader1 total fees paid:", trader1FeeTotal / 1e18, "USDC");
        console.log("Trader2 total fees paid:", trader2FeeIncrease / 1e18, "USDC");
        assertEq(trader2FeeIncrease, trader1FeeTotal); // Same 2% fee for both

        console.log("\n=== VERIFYING DIFFERENT PROTOCOL/CREATOR SPLITS ===");

        // Trader1 (with HORIZON) should result in lower protocol fees, higher creator fees
        // Trader2 (no HORIZON) should result in higher protocol fees, lower creator fees
        uint256 trader1ProtocolFees = treasuryFeesAfterT1;
        uint256 trader2ProtocolFees = treasuryFeesAfterT2 - treasuryFeesAfterT1;
        assertGt(trader2ProtocolFees, trader1ProtocolFees); // Trader2 generates more protocol fees
        
        uint256 trader1CreatorFees = creatorFeesAfterT1;
        uint256 trader2CreatorFees = creatorFeesAfterT2 - creatorFeesAfterT1;
        assertGt(trader1CreatorFees, trader2CreatorFees); // Trader1 generates more creator fees

        console.log("Protocol/creator split verified: HORIZON holders benefit creators");

        console.log("\n=== CLAIMING FEES ===");

        // Creator claims
        uint256 creatorBalanceBefore = usdc.balanceOf(creator1);
        vm.prank(creator1);
        feeSplitter.claimCreatorFees(marketId, address(usdc));
        uint256 creatorClaimed = usdc.balanceOf(creator1) - creatorBalanceBefore;
        console.log("Creator claimed:", creatorClaimed / 1e18, "USDC");
        assertEq(creatorClaimed, creatorFeesAfterT2);

        // Treasury claims
        uint256 treasuryBalanceBefore = usdc.balanceOf(treasury);
        vm.prank(treasury);
        feeSplitter.claimProtocolFees(address(usdc));
        uint256 treasuryClaimed = usdc.balanceOf(treasury) - treasuryBalanceBefore;
        console.log("Treasury claimed:", treasuryClaimed / 1e18, "USDC");
        assertEq(treasuryClaimed, treasuryFeesAfterT2);

        console.log("\n=== FEE FLOW TEST COMPLETE ===");
    }

    /**
     * @notice Test market creation with different collateral tokens
     */
    function test_MultipleCollateralTokens() public {
        console.log("\n=== TESTING MULTIPLE COLLATERAL TOKENS ===");

        // Create USDC market
        MarketFactory.MarketParams memory params1 = MarketFactory.MarketParams({
            marketType: 0, // Binary
            collateralToken: address(usdc),
            closeTime: block.timestamp + 7 days,
            category: "test",
            metadataURI: "ipfs://QmUSDC",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId1 = factory.createMarket(params1);
        console.log("USDC market created:", marketId1);

        // Create DAI market
        MarketFactory.MarketParams memory params2 = MarketFactory.MarketParams({
            marketType: 0, // Binary
            collateralToken: address(dai),
            closeTime: block.timestamp + 7 days,
            category: "test",
            metadataURI: "ipfs://QmDAI",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator2);
        uint256 marketId2 = factory.createMarket(params2);
        console.log("DAI market created:", marketId2);

        // Verify different collateral
        MarketFactory.Market memory market1 = factory.getMarket(marketId1);
        MarketFactory.Market memory market2 = factory.getMarket(marketId2);
        assertEq(market1.collateralToken, address(usdc));
        assertEq(market2.collateralToken, address(dai));

        console.log("\n=== TRADING ON DIFFERENT COLLATERAL MARKETS ===");

        // Trade on USDC market
        BinaryMarket amm1 = BinaryMarket(market1.amm);
        vm.prank(liquidityProvider);
        usdc.approve(address(amm1), type(uint256).max);
        vm.prank(liquidityProvider);
        amm1.addLiquidity(5_000 ether);

        vm.prank(trader1);
        usdc.approve(address(amm1), type(uint256).max);
        vm.prank(trader1);
        amm1.buyYes(100 ether, 0);
        console.log("Traded on USDC market");

        // Trade on DAI market
        BinaryMarket amm2 = BinaryMarket(market2.amm);
        vm.prank(liquidityProvider);
        dai.approve(address(amm2), type(uint256).max);
        vm.prank(liquidityProvider);
        amm2.addLiquidity(5_000 ether);

        vm.prank(trader1);
        dai.approve(address(amm2), type(uint256).max);
        vm.prank(trader1);
        amm2.buyYes(100 ether, 0);
        console.log("Traded on DAI market");

        // Verify fees collected in different tokens
        uint256 usdcFees = feeSplitter.creatorPendingFees(marketId1, address(usdc));
        uint256 daiFees = feeSplitter.creatorPendingFees(marketId2, address(dai));
        assertGt(usdcFees, 0);
        assertGt(daiFees, 0);
        console.log("Fees collected in USDC:", usdcFees / 1e18);
        console.log("Fees collected in DAI:", daiFees / 1e18);

        console.log("\n=== MULTIPLE COLLATERAL TEST COMPLETE ===");
    }
}
