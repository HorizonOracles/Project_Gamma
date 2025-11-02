// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/MarketFactory.sol";
import "../../src/OutcomeToken.sol";
import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";

import "../../src/ResolutionModule.sol";
import "../mocks/MockERC20.sol";

contract MarketFactoryTest is Test {
    MarketFactory public factory;
    OutcomeToken public outcomeToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    MockERC20 public horizonToken;
    ResolutionModule public resolution;
    MockERC20 public usdc;

    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator1 = address(0x2);
    address public creator2 = address(0x3);
    address public arbitrator = address(0x4);

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

    event CreatorStakeRefunded(uint256 indexed marketId, address indexed creator, uint256 amount);
    event MarketStatusUpdated(uint256 indexed marketId, MarketFactory.MarketStatus oldStatus, MarketFactory.MarketStatus newStatus);

    function setUp() public {
        // Deploy system contracts
        usdc = new MockERC20("USDC", "USDC");
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

        // Give factory permissions
        outcomeToken.transferOwnership(address(factory));
        feeSplitter.transferOwnership(address(factory));

        // Fund creators
        horizonToken.transfer(creator1, 10_000 ether);
        horizonToken.transfer(creator2, 10_000 ether);

        // Approve factory
        vm.prank(creator1);
        horizonToken.approve(address(factory), type(uint256).max);
        vm.prank(creator2);
        horizonToken.approve(address(factory), type(uint256).max);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(address(factory.outcomeToken()), address(outcomeToken));
        assertEq(address(factory.feeSplitter()), address(feeSplitter));
        assertEq(address(factory.horizonPerks()), address(horizonPerks));
        assertEq(address(factory.horizonToken()), address(horizonToken));
        assertEq(factory.nextMarketId(), 1);
        assertEq(factory.minCreatorStake(), MIN_STAKE);
    }

    function test_RevertWhen_Constructor_InvalidAddress() public {
        vm.expectRevert(MarketFactory.InvalidAddress.selector);
        new MarketFactory(address(0), address(feeSplitter), address(horizonPerks), address(horizonToken));

        vm.expectRevert(MarketFactory.InvalidAddress.selector);
        new MarketFactory(address(outcomeToken), address(0), address(horizonPerks), address(horizonToken));

        vm.expectRevert(MarketFactory.InvalidAddress.selector);
        new MarketFactory(address(outcomeToken), address(feeSplitter), address(0), address(horizonToken));

        vm.expectRevert(MarketFactory.InvalidAddress.selector);
        new MarketFactory(address(outcomeToken), address(feeSplitter), address(horizonPerks), address(0));
    }

    // ============ Market Creation Tests ============

    function test_CreateMarket() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0, // Binary market
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmExample123",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        uint256 creator1BalanceBefore = horizonToken.balanceOf(creator1);

        vm.expectEmit(true, true, false, true);
        emit MarketCreated(
            1, // marketId
            creator1,
            address(0), // amm address (we don't know it yet)
            address(usdc),
            params.closeTime,
            params.category,
            params.metadataURI,
            params.creatorStake
        );

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        // Verify market ID
        assertEq(marketId, 1);
        assertEq(factory.nextMarketId(), 2);

        // Verify stake was transferred
        assertEq(horizonToken.balanceOf(creator1), creator1BalanceBefore - MIN_STAKE);
        assertEq(horizonToken.balanceOf(address(factory)), MIN_STAKE);

        // Verify market was stored
        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(market.id, marketId);
        assertEq(market.creator, creator1);
        assertEq(market.marketType, 0); // Binary
        assertNotEq(market.amm, address(0));
        assertEq(market.collateralToken, address(usdc));
        assertEq(market.closeTime, params.closeTime);
        assertEq(market.category, params.category);
        assertEq(market.metadataURI, params.metadataURI);
        assertEq(market.creatorStake, params.creatorStake);
        assertEq(market.outcomeCount, 2);
        assertFalse(market.stakeRefunded);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Active));

        // Verify registries updated
        assertEq(factory.getMarketCount(), 1);
        assertEq(factory.getAllMarketIds().length, 1);
        assertEq(factory.getMarketIdsByCategory("politics").length, 1);
        assertEq(factory.getMarketIdsByCreator(creator1).length, 1);

        // Verify market registered in OutcomeToken
        assertTrue(outcomeToken.marketCollateral(marketId) == IERC20(address(usdc)));

        // Verify AMM was authorized
        assertTrue(outcomeToken.authorizedAMMs(market.amm));
    }

    function test_CreateMarket_LargerStake() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 7 days,
            category: "crypto",
            metadataURI: "ipfs://QmTest",
            creatorStake: 500 ether, // Larger stake
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(market.creatorStake, 500 ether);
        assertEq(horizonToken.balanceOf(address(factory)), 500 ether);
    }

    function test_CreateMarket_MultipleMarkets() public {
        // Create first market
        MarketFactory.MarketParams memory params1 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmMarket1",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId1 = factory.createMarket(params1);

        // Create second market
        MarketFactory.MarketParams memory params2 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 60 days,
            category: "sports",
            metadataURI: "ipfs://QmMarket2",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator2);
        uint256 marketId2 = factory.createMarket(params2);

        // Verify IDs
        assertEq(marketId1, 1);
        assertEq(marketId2, 2);
        assertEq(factory.nextMarketId(), 3);

        // Verify count
        assertEq(factory.getMarketCount(), 2);

        // Verify creator registries
        assertEq(factory.getMarketIdsByCreator(creator1).length, 1);
        assertEq(factory.getMarketIdsByCreator(creator2).length, 1);
    }

    function test_RevertWhen_CreateMarket_InvalidCollateral() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(0),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmExample",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        vm.expectRevert(MarketFactory.InvalidCollateral.selector);
        factory.createMarket(params);
    }

    function test_RevertWhen_CreateMarket_InvalidCloseTime() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0, // Binary
            collateralToken: address(usdc),
            closeTime: block.timestamp - 1, // In the past
            category: "politics",
            metadataURI: "ipfs://QmExample",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        vm.expectRevert(MarketFactory.InvalidCloseTime.selector);
        factory.createMarket(params);
    }

    function test_RevertWhen_CreateMarket_InsufficientStake() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmExample",
            creatorStake: MIN_STAKE - 1,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        vm.expectRevert(MarketFactory.InvalidCreatorStake.selector);
        factory.createMarket(params);
    }

    function test_RevertWhen_CreateMarket_EmptyCategory() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "",
            metadataURI: "ipfs://QmExample",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        vm.expectRevert(MarketFactory.InvalidCategory.selector);
        factory.createMarket(params);
    }

    // ============ Creator Stake Tests ============

    function test_RefundCreatorStake() public {
        // Create market
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmExample",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        // Warp past close time
        vm.warp(params.closeTime + 1);

        // Resolve market (need to get owner back first for this test)
        vm.prank(address(factory));
        outcomeToken.transferOwnership(owner);

        outcomeToken.setResolutionAuthorization(owner, true);
        outcomeToken.setWinningOutcome(marketId, 0);

        // Refund stake
        uint256 creator1BalanceBefore = horizonToken.balanceOf(creator1);

        vm.expectEmit(true, true, false, true);
        emit CreatorStakeRefunded(marketId, creator1, MIN_STAKE);

        factory.refundCreatorStake(marketId);

        // Verify refund
        assertEq(horizonToken.balanceOf(creator1), creator1BalanceBefore + MIN_STAKE);
        assertEq(horizonToken.balanceOf(address(factory)), 0);

        // Verify market updated
        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertTrue(market.stakeRefunded);
    }

    function test_RefundCreatorStake_ByCreator() public {
        // Create and resolve market
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 1 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        vm.warp(params.closeTime + 1);

        // Resolve
        vm.prank(address(factory));
        outcomeToken.transferOwnership(owner);
        outcomeToken.setResolutionAuthorization(owner, true);
        outcomeToken.setWinningOutcome(marketId, 1);

        // Creator refunds themselves
        vm.prank(creator1);
        factory.refundCreatorStake(marketId);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertTrue(market.stakeRefunded);
    }

    function test_RevertWhen_RefundCreatorStake_MarketDoesNotExist() public {
        vm.expectRevert(MarketFactory.MarketDoesNotExist.selector);
        factory.refundCreatorStake(999);
    }

    function test_RevertWhen_RefundCreatorStake_NotResolved() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        vm.expectRevert(MarketFactory.MarketNotResolved.selector);
        factory.refundCreatorStake(marketId);
    }

    function test_RevertWhen_RefundCreatorStake_AlreadyClaimed() public {
        // Create and resolve market
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 1 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        vm.warp(params.closeTime + 1);

        vm.prank(address(factory));
        outcomeToken.transferOwnership(owner);
        outcomeToken.setResolutionAuthorization(owner, true);
        outcomeToken.setWinningOutcome(marketId, 0);

        // First refund succeeds
        factory.refundCreatorStake(marketId);

        // Second refund fails
        vm.expectRevert(MarketFactory.StakeAlreadyClaimed.selector);
        factory.refundCreatorStake(marketId);
    }

    // ============ Market Status Tests ============

    function test_UpdateMarketStatus_Active() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Active));
    }

    function test_UpdateMarketStatus_Closed() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 1 hours,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        // Warp past close time
        vm.warp(params.closeTime + 1);

        // Update status
        vm.expectEmit(true, false, false, true);
        emit MarketStatusUpdated(marketId, MarketFactory.MarketStatus.Active, MarketFactory.MarketStatus.Closed);

        factory.updateMarketStatus(marketId);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Closed));
    }

    function test_UpdateMarketStatus_Resolved() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 1 hours,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        vm.warp(params.closeTime + 1);

        // Resolve market
        vm.prank(address(factory));
        outcomeToken.transferOwnership(owner);
        outcomeToken.setResolutionAuthorization(owner, true);
        outcomeToken.setWinningOutcome(marketId, 0);

        // Update status
        factory.updateMarketStatus(marketId);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(uint8(market.status), uint8(MarketFactory.MarketStatus.Resolved));
    }

    // ============ Query Functions Tests ============

    function test_GetMarket() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "sports",
            metadataURI: "ipfs://QmSports",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        MarketFactory.Market memory market = factory.getMarket(marketId);
        assertEq(market.id, marketId);
        assertEq(market.creator, creator1);
        assertEq(market.category, "sports");
    }

    function test_RevertWhen_GetMarket_DoesNotExist() public {
        vm.expectRevert(MarketFactory.MarketDoesNotExist.selector);
        factory.getMarket(999);
    }

    function test_GetMarkets_Pagination() public {
        // Create 5 markets
        for (uint256 i = 0; i < 5; i++) {
            MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

            vm.prank(creator1);
            factory.createMarket(params);
        }

        // Get first 3
        MarketFactory.Market[] memory markets = factory.getMarkets(0, 3);
        assertEq(markets.length, 3);
        assertEq(markets[0].id, 1);
        assertEq(markets[2].id, 3);

        // Get next 2
        markets = factory.getMarkets(3, 3);
        assertEq(markets.length, 2);
        assertEq(markets[0].id, 4);
        assertEq(markets[1].id, 5);

        // Get beyond end
        markets = factory.getMarkets(10, 5);
        assertEq(markets.length, 0);
    }

    function test_GetActiveMarkets() public {
        // Create market that will be active
        MarketFactory.MarketParams memory params1 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmActive",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        factory.createMarket(params1);

        // Create market that will be closed
        MarketFactory.MarketParams memory params2 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 1 seconds,
            category: "test",
            metadataURI: "ipfs://QmClosed",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        factory.createMarket(params2);

        // Warp to close second market
        vm.warp(block.timestamp + 2 seconds);

        // Get active markets
        MarketFactory.Market[] memory activeMarkets = factory.getActiveMarkets(0, 10);
        assertEq(activeMarkets.length, 1);
        assertEq(activeMarkets[0].id, 1);
    }

    function test_GetMarketIdsByCategory() public {
        // Create markets in different categories
        MarketFactory.MarketParams memory params1 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmPolitics1",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        factory.createMarket(params1);

        MarketFactory.MarketParams memory params2 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "politics",
            metadataURI: "ipfs://QmPolitics2",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        factory.createMarket(params2);

        MarketFactory.MarketParams memory params3 = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "sports",
            metadataURI: "ipfs://QmSports",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator2);
        factory.createMarket(params3);

        // Query by category
        uint256[] memory politicsMarkets = factory.getMarketIdsByCategory("politics");
        assertEq(politicsMarkets.length, 2);
        assertEq(politicsMarkets[0], 1);
        assertEq(politicsMarkets[1], 2);

        uint256[] memory sportsMarkets = factory.getMarketIdsByCategory("sports");
        assertEq(sportsMarkets.length, 1);
        assertEq(sportsMarkets[0], 3);
    }

    function test_GetMarketIdsByCreator() public {
        // Create markets
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmTest1",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        factory.createMarket(params);

        params.metadataURI = "ipfs://QmTest2";
        vm.prank(creator1);
        factory.createMarket(params);

        params.metadataURI = "ipfs://QmTest3";
        vm.prank(creator2);
        factory.createMarket(params);

        // Query by creator
        uint256[] memory creator1Markets = factory.getMarketIdsByCreator(creator1);
        assertEq(creator1Markets.length, 2);

        uint256[] memory creator2Markets = factory.getMarketIdsByCreator(creator2);
        assertEq(creator2Markets.length, 1);
    }

    function test_MarketExists() public {
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            marketType: 0,
            collateralToken: address(usdc),
            closeTime: block.timestamp + 30 days,
            category: "test",
            metadataURI: "ipfs://QmTest",
            creatorStake: MIN_STAKE,
            outcomeCount: 2,
            liquidityParameter: 0
        });

        vm.prank(creator1);
        uint256 marketId = factory.createMarket(params);

        assertTrue(factory.marketExists(marketId));
        assertFalse(factory.marketExists(999));
    }

    // ============ Admin Functions Tests ============

    function test_SetMinCreatorStake() public {
        uint256 newStake = 500 ether;
        factory.setMinCreatorStake(newStake);
        assertEq(factory.minCreatorStake(), newStake);
    }

    function test_RevertWhen_SetMinCreatorStake_Unauthorized() public {
        vm.prank(creator1);
        vm.expectRevert();
        factory.setMinCreatorStake(500 ether);
    }
}
