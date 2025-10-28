// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/OutcomeToken.sol";
import "../mocks/MockERC20.sol";

contract OutcomeTokenTest is Test {
    OutcomeToken public outcomeToken;
    MockERC20 public collateral;

    address public owner = address(this);
    address public amm = address(0x1);
    address public user1 = address(0x2);
    address public user2 = address(0x3);
    address public unauthorized = address(0x4);

    uint256 public constant MARKET_ID = 1;
    uint256 public constant OUTCOME_YES = 0;
    uint256 public constant OUTCOME_NO = 1;

    event OutcomeMinted(uint256 indexed marketId, uint256 indexed outcomeId, address indexed to, uint256 amount);
    event OutcomeBurned(uint256 indexed marketId, uint256 indexed outcomeId, address indexed from, uint256 amount);
    event WinningOutcomeSet(uint256 indexed marketId, uint256 indexed winningOutcomeId);
    event Redeemed(uint256 indexed marketId, address indexed user, uint256 winningTokens, uint256 collateralPaid);
    event AMMAuthorized(address indexed amm, bool authorized);

    function setUp() public {
        // Deploy contracts
        outcomeToken = new OutcomeToken("https://api.example.com/metadata/{id}.json");
        collateral = new MockERC20("USDC", "USDC");

        // Setup market
        outcomeToken.registerMarket(MARKET_ID, collateral);

        // Authorize AMM
        outcomeToken.setAMMAuthorization(amm, true);

        // Authorize test contract as resolver (for testing setWinningOutcome)
        outcomeToken.setResolutionAuthorization(address(this), true);

        // Fund contract with collateral for redemptions
        collateral.transfer(address(outcomeToken), 1000 ether);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public {
        assertEq(outcomeToken.owner(), owner);
        assertTrue(outcomeToken.authorizedAMMs(amm));
        assertEq(address(outcomeToken.marketCollateral(MARKET_ID)), address(collateral));
    }

    // ============ Authorization Tests ============

    function test_SetAMMAuthorization() public {
        address newAMM = address(0x5);

        vm.expectEmit(true, true, false, true);
        emit AMMAuthorized(newAMM, true);

        outcomeToken.setAMMAuthorization(newAMM, true);
        assertTrue(outcomeToken.authorizedAMMs(newAMM));
    }

    function test_SetAMMAuthorization_Revoke() public {
        vm.expectEmit(true, true, false, true);
        emit AMMAuthorized(amm, false);

        outcomeToken.setAMMAuthorization(amm, false);
        assertFalse(outcomeToken.authorizedAMMs(amm));
    }

    function test_RevertWhen_SetAMMAuthorization_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        outcomeToken.setAMMAuthorization(address(0x5), true);
    }

    // ============ Market Registration Tests ============

    function test_RegisterMarket() public {
        uint256 newMarketId = 2;
        MockERC20 newCollateral = new MockERC20("DAI", "DAI");

        outcomeToken.registerMarket(newMarketId, newCollateral);

        assertEq(address(outcomeToken.marketCollateral(newMarketId)), address(newCollateral));
        assertEq(outcomeToken.winningOutcome(newMarketId), outcomeToken.UNRESOLVED());
        assertFalse(outcomeToken.isResolved(newMarketId));
    }

    // ============ Token ID Encoding Tests ============

    function test_EncodeTokenId() public {
        uint256 tokenId = outcomeToken.encodeTokenId(MARKET_ID, OUTCOME_YES);
        assertEq(tokenId, (MARKET_ID << 8) | OUTCOME_YES);
    }

    function test_DecodeTokenId() public {
        uint256 tokenId = outcomeToken.encodeTokenId(MARKET_ID, OUTCOME_YES);
        (uint256 marketId, uint256 outcomeId) = outcomeToken.decodeTokenId(tokenId);
        assertEq(marketId, MARKET_ID);
        assertEq(outcomeId, OUTCOME_YES);
    }

    function testFuzz_EncodeDecodeTokenId(uint248 marketId, uint8 outcomeId) public {
        uint256 tokenId = outcomeToken.encodeTokenId(marketId, outcomeId);
        (uint256 decodedMarketId, uint256 decodedOutcomeId) = outcomeToken.decodeTokenId(tokenId);
        assertEq(decodedMarketId, marketId);
        assertEq(decodedOutcomeId, outcomeId);
    }

    // ============ Minting Tests ============

    function test_MintOutcome() public {
        uint256 amount = 100 ether;

        vm.expectEmit(true, true, true, true);
        emit OutcomeMinted(MARKET_ID, OUTCOME_YES, user1, amount);

        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount);
    }

    function test_MintOutcome_Multiple() public {
        uint256 amount1 = 100 ether;
        uint256 amount2 = 50 ether;

        vm.startPrank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount1);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_NO, user1, amount2);
        vm.stopPrank();

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount1);
        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_NO), amount2);
    }

    function test_RevertWhen_MintOutcome_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert(OutcomeToken.Unauthorized.selector);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, 100 ether);
    }

    function testFuzz_MintOutcome(uint256 amount) public {
        vm.assume(amount > 0 && amount < type(uint128).max);

        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount);
    }

    // ============ Burning Tests ============

    function test_BurnOutcome() public {
        uint256 amount = 100 ether;

        // First mint
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        // Then burn
        vm.expectEmit(true, true, true, true);
        emit OutcomeBurned(MARKET_ID, OUTCOME_YES, user1, amount);

        vm.prank(amm);
        outcomeToken.burnOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), 0);
    }

    function test_BurnOutcome_Partial() public {
        uint256 amount = 100 ether;
        uint256 burnAmount = 30 ether;

        vm.startPrank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);
        outcomeToken.burnOutcome(MARKET_ID, OUTCOME_YES, user1, burnAmount);
        vm.stopPrank();

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount - burnAmount);
    }

    function test_RevertWhen_BurnOutcome_Unauthorized() public {
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, 100 ether);

        vm.prank(unauthorized);
        vm.expectRevert(OutcomeToken.Unauthorized.selector);
        outcomeToken.burnOutcome(MARKET_ID, OUTCOME_YES, user1, 50 ether);
    }

    // ============ Resolution Tests ============

    function test_SetWinningOutcome() public {
        vm.expectEmit(true, true, false, true);
        emit WinningOutcomeSet(MARKET_ID, OUTCOME_YES);

        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        assertEq(outcomeToken.winningOutcome(MARKET_ID), OUTCOME_YES);
        assertTrue(outcomeToken.isResolved(MARKET_ID));
    }

    function test_RevertWhen_SetWinningOutcome_AlreadyResolved() public {
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);
        vm.expectRevert(OutcomeToken.MarketAlreadyResolved.selector);
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_NO); // Should fail
    }

    function test_RevertWhen_SetWinningOutcome_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);
    }

    // ============ Redemption Tests ============

    function test_Redeem() public {
        uint256 amount = 100 ether;

        // Mint winning tokens
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        // Resolve market
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        // Get initial balance
        uint256 initialBalance = collateral.balanceOf(user1);

        // Redeem
        vm.expectEmit(true, true, false, true);
        emit Redeemed(MARKET_ID, user1, amount, amount);

        vm.prank(user1);
        uint256 payout = outcomeToken.redeem(MARKET_ID);

        assertEq(payout, amount);
        assertEq(collateral.balanceOf(user1), initialBalance + amount);
        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), 0);
    }

    function test_RedeemAmount() public {
        uint256 amount = 100 ether;
        uint256 redeemAmount = 60 ether;

        // Mint winning tokens
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        // Resolve market
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        // Get initial balance
        uint256 initialBalance = collateral.balanceOf(user1);

        // Partial redeem
        vm.prank(user1);
        uint256 payout = outcomeToken.redeemAmount(MARKET_ID, redeemAmount);

        assertEq(payout, redeemAmount);
        assertEq(collateral.balanceOf(user1), initialBalance + redeemAmount);
        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount - redeemAmount);
    }

    function test_RevertWhen_Redeem_NotResolved() public {
        // Mint tokens but don't resolve
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, 100 ether);

        vm.prank(user1);
        vm.expectRevert(OutcomeToken.MarketNotResolved.selector);
        outcomeToken.redeem(MARKET_ID); // Should fail
    }

    function test_RevertWhen_Redeem_NoTokens() public {
        // Resolve but user has no tokens
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        vm.prank(user1);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(MARKET_ID); // Should fail
    }

    function test_Redeem_LosingOutcome() public {
        // Mint losing tokens
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_NO, user1, 100 ether);

        // Resolve with YES as winner
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        // Try to redeem (should fail - no winning tokens)
        vm.prank(user1);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(MARKET_ID);
    }

    function testFuzz_Redeem(uint128 amount) public {
        vm.assume(amount > 0 && amount <= 1000 ether);

        // Ensure contract has enough collateral for redemption
        collateral.mint(address(outcomeToken), amount);

        // Mint winning tokens
        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        // Resolve market
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        // Redeem
        uint256 initialBalance = collateral.balanceOf(user1);
        vm.prank(user1);
        uint256 payout = outcomeToken.redeem(MARKET_ID);

        assertEq(payout, amount);
        assertEq(collateral.balanceOf(user1), initialBalance + amount);
    }

    // ============ View Function Tests ============

    function test_BalanceOfOutcome() public {
        uint256 amount = 100 ether;

        vm.prank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, amount);

        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), amount);
        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_NO), 0);
    }

    function test_IsResolved() public {
        assertFalse(outcomeToken.isResolved(MARKET_ID));

        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);
        assertTrue(outcomeToken.isResolved(MARKET_ID));
    }

    // ============ Integration Tests ============

    function test_FullLifecycle() public {
        // 1. Mint tokens for both users
        vm.startPrank(amm);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_YES, user1, 100 ether);
        outcomeToken.mintOutcome(MARKET_ID, OUTCOME_NO, user2, 50 ether);
        vm.stopPrank();

        // 2. Verify balances
        assertEq(outcomeToken.balanceOfOutcome(user1, MARKET_ID, OUTCOME_YES), 100 ether);
        assertEq(outcomeToken.balanceOfOutcome(user2, MARKET_ID, OUTCOME_NO), 50 ether);

        // 3. Resolve market with YES winning
        outcomeToken.setWinningOutcome(MARKET_ID, OUTCOME_YES);

        // 4. User1 redeems winning tokens
        uint256 user1InitialBalance = collateral.balanceOf(user1);
        vm.prank(user1);
        outcomeToken.redeem(MARKET_ID);
        assertEq(collateral.balanceOf(user1), user1InitialBalance + 100 ether);

        // 5. User2 cannot redeem losing tokens
        vm.prank(user2);
        vm.expectRevert(OutcomeToken.NoTokensToRedeem.selector);
        outcomeToken.redeem(MARKET_ID);
    }
}
