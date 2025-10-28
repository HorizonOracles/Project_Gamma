// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/FeeSplitter.sol";
import "../mocks/MockERC20.sol";

contract FeeSplitterTest is Test {
    FeeSplitter public splitter;
    MockERC20 public token;

    address public owner = address(this);
    address public protocolTreasury = address(0x1);
    address public creator1 = address(0x2);
    address public creator2 = address(0x3);
    address public amm = address(0x4);
    address public unauthorized = address(0x5);

    uint256 public constant MARKET_ID_1 = 1;
    uint256 public constant MARKET_ID_2 = 2;

    event FeeDistributed(
        uint256 indexed marketId, address indexed token, uint256 totalAmount, uint256 protocolAmount, uint256 creatorAmount, uint16 protocolBps
    );
    event FeeClaimed(uint256 indexed marketId, address indexed creator, address indexed token, uint256 amount);
    event ProtocolFeeClaimed(address indexed token, uint256 amount);
    event MarketRegistered(uint256 indexed marketId, address indexed creator);
    event FeeConfigUpdated(uint256 indexed marketId, uint16 protocolBps, uint16 creatorBps);

    function setUp() public {
        splitter = new FeeSplitter(protocolTreasury);
        token = new MockERC20("USDC", "USDC");

        // Register markets
        splitter.registerMarket(MARKET_ID_1, creator1);
        splitter.registerMarket(MARKET_ID_2, creator2);

        // Fund AMM with tokens
        token.mint(amm, 1_000_000 ether);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(splitter.owner(), owner);
        assertEq(splitter.protocolTreasury(), protocolTreasury);
    }

    function test_RevertWhen_Constructor_InvalidTreasury() public {
        vm.expectRevert("Invalid treasury");
        new FeeSplitter(address(0));
    }

    // ============ Market Registration Tests ============

    function test_RegisterMarket() public {
        uint256 newMarketId = 3;
        address newCreator = address(0x6);

        vm.expectEmit(true, true, false, false);
        emit MarketRegistered(newMarketId, newCreator);

        splitter.registerMarket(newMarketId, newCreator);

        assertEq(splitter.marketCreator(newMarketId), newCreator);
        (uint16 protocolBps, uint16 creatorBps) = splitter.getFeeConfig(newMarketId);
        assertEq(protocolBps, 1000); // 10%
        assertEq(creatorBps, 9000); // 90%
    }

    function test_RevertWhen_RegisterMarket_InvalidCreator() public {
        vm.expectRevert("Invalid creator");
        splitter.registerMarket(3, address(0));
    }

    function test_RevertWhen_RegisterMarket_AlreadyRegistered() public {
        vm.expectRevert("Market already registered");
        splitter.registerMarket(MARKET_ID_1, creator1);
    }

    function test_RevertWhen_RegisterMarket_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        splitter.registerMarket(3, creator1);
    }

    // ============ Fee Configuration Tests ============

    function test_UpdateFeeConfig() public {
        uint16 newProtocolBps = 2000; // 20%
        uint16 newCreatorBps = 8000; // 80%

        vm.expectEmit(true, false, false, true);
        emit FeeConfigUpdated(MARKET_ID_1, newProtocolBps, newCreatorBps);

        splitter.updateFeeConfig(MARKET_ID_1, newProtocolBps, newCreatorBps);

        (uint16 protocolBps, uint16 creatorBps) = splitter.getFeeConfig(MARKET_ID_1);
        assertEq(protocolBps, newProtocolBps);
        assertEq(creatorBps, newCreatorBps);
    }

    function test_RevertWhen_UpdateFeeConfig_InvalidSum() public {
        vm.expectRevert(FeeSplitter.InvalidFeeConfig.selector);
        splitter.updateFeeConfig(MARKET_ID_1, 2000, 7000); // Doesn't sum to 10000
    }

    function test_RevertWhen_UpdateFeeConfig_MarketNotRegistered() public {
        vm.expectRevert(FeeSplitter.MarketNotRegistered.selector);
        splitter.updateFeeConfig(999, 1000, 9000);
    }

    function test_RevertWhen_UpdateFeeConfig_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        splitter.updateFeeConfig(MARKET_ID_1, 1000, 9000);
    }

    // ============ Treasury Update Tests ============

    function test_SetProtocolTreasury() public {
        address newTreasury = address(0x7);
        splitter.setProtocolTreasury(newTreasury);
        assertEq(splitter.protocolTreasury(), newTreasury);
    }

    function test_RevertWhen_SetProtocolTreasury_InvalidAddress() public {
        vm.expectRevert("Invalid treasury");
        splitter.setProtocolTreasury(address(0));
    }

    // ============ Fee Distribution Tests ============

    function test_Distribute() public {
        uint256 feeAmount = 1000 ether;

        // AMM approves splitter
        vm.prank(amm);
        token.approve(address(splitter), feeAmount);

        // Calculate expected splits
        uint256 expectedProtocol = (feeAmount * 1000) / 10000; // 10% = 100 ether
        uint256 expectedCreator = (feeAmount * 9000) / 10000; // 90% = 900 ether

        vm.expectEmit(true, true, false, true);
        emit FeeDistributed(MARKET_ID_1, address(token), feeAmount, expectedProtocol, expectedCreator, 1000);

        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);

        // Verify pending fees
        assertEq(splitter.getProtocolPendingFees(address(token)), expectedProtocol);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), expectedCreator);
        assertEq(token.balanceOf(address(splitter)), feeAmount);
    }

    function test_Distribute_MultipleMarkets() public {
        uint256 feeAmount = 1000 ether;

        vm.startPrank(amm);
        token.approve(address(splitter), feeAmount * 2);

        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_2, address(token), feeAmount, 1000);
        vm.stopPrank();

        // Each market should have 900 ether pending for creator
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), 900 ether);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_2, address(token)), 900 ether);

        // Protocol should have 200 ether total (100 from each)
        assertEq(splitter.getProtocolPendingFees(address(token)), 200 ether);
    }

    function test_Distribute_ZeroAmount() public {
        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), 0, 1000);

        // Should not change balances
        assertEq(splitter.getProtocolPendingFees(address(token)), 0);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), 0);
    }

    function test_RevertWhen_Distribute_MarketNotRegistered() public {
        vm.prank(amm);
        vm.expectRevert(FeeSplitter.MarketNotRegistered.selector);
        splitter.distribute(999, address(token), 1000 ether, 1000);
    }

    function testFuzz_Distribute(uint128 amount) public {
        vm.assume(amount > 0 && amount <= 1_000_000 ether);

        token.mint(amm, amount);
        vm.startPrank(amm);
        token.approve(address(splitter), amount);
        splitter.distribute(MARKET_ID_1, address(token), amount, 1000);
        vm.stopPrank();

        uint256 protocolFees = splitter.getProtocolPendingFees(address(token));
        uint256 creatorFees = splitter.getCreatorPendingFees(MARKET_ID_1, address(token));

        // Verify split is approximately correct (accounting for rounding)
        assertGe(protocolFees + creatorFees, amount);
        assertLe(protocolFees + creatorFees, amount + 1); // Max 1 wei rounding error
    }

    // ============ Creator Claim Tests ============

    function test_ClaimCreatorFees() public {
        uint256 feeAmount = 1000 ether;

        // Distribute fees
        vm.prank(amm);
        token.approve(address(splitter), feeAmount);
        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);

        uint256 creatorBalanceBefore = token.balanceOf(creator1);
        uint256 expectedClaim = 900 ether; // 90% of 1000

        vm.expectEmit(true, true, true, true);
        emit FeeClaimed(MARKET_ID_1, creator1, address(token), expectedClaim);

        // Creator claims fees
        vm.prank(creator1);
        splitter.claimCreatorFees(MARKET_ID_1, address(token));

        assertEq(token.balanceOf(creator1), creatorBalanceBefore + expectedClaim);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), 0);
    }

    function test_RevertWhen_ClaimCreatorFees_NotCreator() public {
        vm.prank(unauthorized);
        vm.expectRevert("Not the creator");
        splitter.claimCreatorFees(MARKET_ID_1, address(token));
    }

    function test_RevertWhen_ClaimCreatorFees_NoFees() public {
        vm.prank(creator1);
        vm.expectRevert(FeeSplitter.NoFeesToClaim.selector);
        splitter.claimCreatorFees(MARKET_ID_1, address(token));
    }

    function test_ClaimCreatorFees_Multiple() public {
        uint256 feeAmount = 1000 ether;

        // Distribute fees twice
        vm.startPrank(amm);
        token.approve(address(splitter), feeAmount * 2);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        vm.stopPrank();

        // Creator should have 1800 ether (900 * 2)
        uint256 expectedClaim = 1800 ether;

        vm.prank(creator1);
        splitter.claimCreatorFees(MARKET_ID_1, address(token));

        assertEq(token.balanceOf(creator1), expectedClaim);
    }

    // ============ Batch Creator Claim Tests ============

    function test_ClaimCreatorFeesMultiple() public {
        uint256 feeAmount = 1000 ether;
        MockERC20 token2 = new MockERC20("DAI", "DAI");

        // Distribute fees in both tokens
        token.mint(amm, feeAmount);
        token2.mint(amm, feeAmount);

        vm.startPrank(amm);
        token.approve(address(splitter), feeAmount);
        token2.approve(address(splitter), feeAmount);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_1, address(token2), feeAmount, 1000);
        vm.stopPrank();

        // Creator claims from both
        uint256[] memory marketIds = new uint256[](2);
        address[] memory tokens = new address[](2);
        marketIds[0] = MARKET_ID_1;
        marketIds[1] = MARKET_ID_1;
        tokens[0] = address(token);
        tokens[1] = address(token2);

        vm.prank(creator1);
        splitter.claimCreatorFeesMultiple(marketIds, tokens);

        assertEq(token.balanceOf(creator1), 900 ether);
        assertEq(token2.balanceOf(creator1), 900 ether);
    }

    function test_ClaimCreatorFeesMultiple_SkipsInvalid() public {
        uint256[] memory marketIds = new uint256[](2);
        address[] memory tokens = new address[](2);
        marketIds[0] = MARKET_ID_1;
        marketIds[1] = MARKET_ID_2; // Different creator
        tokens[0] = address(token);
        tokens[1] = address(token);

        // Should not revert, just skip MARKET_ID_2
        vm.prank(creator1);
        splitter.claimCreatorFeesMultiple(marketIds, tokens);
    }

    function test_RevertWhen_ClaimCreatorFeesMultiple_ArrayMismatch() public {
        uint256[] memory marketIds = new uint256[](2);
        address[] memory tokens = new address[](1);

        vm.prank(creator1);
        vm.expectRevert("Array length mismatch");
        splitter.claimCreatorFeesMultiple(marketIds, tokens);
    }

    // ============ Protocol Claim Tests ============

    function test_ClaimProtocolFees() public {
        uint256 feeAmount = 1000 ether;

        // Distribute fees
        vm.prank(amm);
        token.approve(address(splitter), feeAmount);
        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);

        uint256 treasuryBalanceBefore = token.balanceOf(protocolTreasury);
        uint256 expectedClaim = 100 ether; // 10% of 1000

        vm.expectEmit(true, false, false, true);
        emit ProtocolFeeClaimed(address(token), expectedClaim);

        // Protocol treasury claims
        vm.prank(protocolTreasury);
        splitter.claimProtocolFees(address(token));

        assertEq(token.balanceOf(protocolTreasury), treasuryBalanceBefore + expectedClaim);
        assertEq(splitter.getProtocolPendingFees(address(token)), 0);
    }

    function test_ClaimProtocolFees_ByOwner() public {
        uint256 feeAmount = 1000 ether;

        vm.prank(amm);
        token.approve(address(splitter), feeAmount);
        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);

        // Owner can also claim
        vm.prank(owner);
        splitter.claimProtocolFees(address(token));

        assertEq(token.balanceOf(protocolTreasury), 100 ether);
    }

    function test_RevertWhen_ClaimProtocolFees_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert("Not authorized");
        splitter.claimProtocolFees(address(token));
    }

    function test_RevertWhen_ClaimProtocolFees_NoFees() public {
        vm.prank(protocolTreasury);
        vm.expectRevert(FeeSplitter.NoFeesToClaim.selector);
        splitter.claimProtocolFees(address(token));
    }

    // ============ Batch Protocol Claim Tests ============

    function test_ClaimProtocolFeesMultiple() public {
        MockERC20 token2 = new MockERC20("DAI", "DAI");
        uint256 feeAmount = 1000 ether;

        // Distribute in both tokens
        token.mint(amm, feeAmount);
        token2.mint(amm, feeAmount);

        vm.startPrank(amm);
        token.approve(address(splitter), feeAmount);
        token2.approve(address(splitter), feeAmount);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_1, address(token2), feeAmount, 1000);
        vm.stopPrank();

        address[] memory tokens = new address[](2);
        tokens[0] = address(token);
        tokens[1] = address(token2);

        vm.prank(protocolTreasury);
        splitter.claimProtocolFeesMultiple(tokens);

        assertEq(token.balanceOf(protocolTreasury), 100 ether);
        assertEq(token2.balanceOf(protocolTreasury), 100 ether);
    }

    // ============ Preview Split Tests ============

    function test_PreviewSplit() public view {
        uint256 amount = 1000 ether;
        (uint256 protocolAmount, uint256 creatorAmount) = splitter.previewSplit(amount, 1000);

        assertEq(protocolAmount, 100 ether);
        assertEq(creatorAmount, 900 ether);
        assertEq(protocolAmount + creatorAmount, amount);
    }

    function test_PreviewSplit_CustomConfig() public view {
        uint256 amount = 1000 ether;
        (uint256 protocolAmount, uint256 creatorAmount) = splitter.previewSplit(amount, 2000);

        assertEq(protocolAmount, 200 ether);
        assertEq(creatorAmount, 800 ether);
    }

    function testFuzz_PreviewSplit(uint128 amount) public view {
        vm.assume(amount > 0);
        (uint256 protocolAmount, uint256 creatorAmount) = splitter.previewSplit(amount, 1000);

        // Verify total equals original amount (accounting for rounding dust to creator)
        assertEq(protocolAmount + creatorAmount, amount);
    }

    // ============ Integration Tests ============

    function test_FullLifecycle() public {
        uint256 feeAmount = 10000 ether;

        // 1. AMM distributes fees
        vm.prank(amm);
        token.approve(address(splitter), feeAmount);
        vm.prank(amm);
        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);

        // 2. Verify pending fees
        assertEq(splitter.getProtocolPendingFees(address(token)), 1000 ether);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), 9000 ether);

        // 3. Creator claims
        vm.prank(creator1);
        splitter.claimCreatorFees(MARKET_ID_1, address(token));
        assertEq(token.balanceOf(creator1), 9000 ether);

        // 4. Protocol claims
        vm.prank(protocolTreasury);
        splitter.claimProtocolFees(address(token));
        assertEq(token.balanceOf(protocolTreasury), 1000 ether);

        // 5. Verify all pending fees cleared
        assertEq(splitter.getProtocolPendingFees(address(token)), 0);
        assertEq(splitter.getCreatorPendingFees(MARKET_ID_1, address(token)), 0);
    }

    function test_MultipleMarkets_MultipleTokens() public {
        MockERC20 token2 = new MockERC20("DAI", "DAI");
        uint256 feeAmount = 1000 ether;

        // Fund AMM with both tokens
        token.mint(amm, feeAmount * 2);
        token2.mint(amm, feeAmount * 2);

        // Distribute fees across markets and tokens
        vm.startPrank(amm);
        token.approve(address(splitter), feeAmount * 2);
        token2.approve(address(splitter), feeAmount * 2);

        splitter.distribute(MARKET_ID_1, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_1, address(token2), feeAmount, 1000);
        splitter.distribute(MARKET_ID_2, address(token), feeAmount, 1000);
        splitter.distribute(MARKET_ID_2, address(token2), feeAmount, 1000);
        vm.stopPrank();

        // Creator1 claims from market 1
        vm.prank(creator1);
        uint256[] memory marketIds = new uint256[](2);
        address[] memory tokens = new address[](2);
        marketIds[0] = MARKET_ID_1;
        marketIds[1] = MARKET_ID_1;
        tokens[0] = address(token);
        tokens[1] = address(token2);
        splitter.claimCreatorFeesMultiple(marketIds, tokens);

        // Creator2 claims from market 2
        vm.prank(creator2);
        marketIds[0] = MARKET_ID_2;
        marketIds[1] = MARKET_ID_2;
        splitter.claimCreatorFeesMultiple(marketIds, tokens);

        // Verify balances
        assertEq(token.balanceOf(creator1), 900 ether);
        assertEq(token2.balanceOf(creator1), 900 ether);
        assertEq(token.balanceOf(creator2), 900 ether);
        assertEq(token2.balanceOf(creator2), 900 ether);

        // Protocol claims both tokens
        address[] memory protocolTokens = new address[](2);
        protocolTokens[0] = address(token);
        protocolTokens[1] = address(token2);

        vm.prank(protocolTreasury);
        splitter.claimProtocolFeesMultiple(protocolTokens);

        // Protocol should have 200 ether in each token (100 * 2 markets)
        assertEq(token.balanceOf(protocolTreasury), 200 ether);
        assertEq(token2.balanceOf(protocolTreasury), 200 ether);
    }
}
