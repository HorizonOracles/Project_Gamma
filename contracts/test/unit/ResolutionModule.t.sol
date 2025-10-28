// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/ResolutionModule.sol";
import "../../src/OutcomeToken.sol";
import "../../src/HorizonToken.sol";

contract ResolutionModuleTest is Test {
    ResolutionModule public resolution;
    OutcomeToken public outcomeToken;
    HorizonToken public bondToken;

    address public owner = address(this);
    address public arbitrator = address(0x1);
    address public proposer = address(0x2);
    address public disputer = address(0x3);
    address public creator = address(0x4);

    uint256 public constant MARKET_ID = 1;
    uint256 public constant MIN_BOND = 1000 ether;

    event ResolutionProposed(
        uint256 indexed marketId,
        uint256 indexed outcomeId,
        address indexed proposer,
        uint256 bond,
        string evidenceURI,
        uint256 deadline
    );

    event Disputed(
        uint256 indexed marketId,
        address indexed disputer,
        uint256 bond,
        string reason
    );

    event Finalized(
        uint256 indexed marketId,
        uint256 indexed outcomeId,
        bool wasDisputed
    );

    event BondRefunded(address indexed recipient, uint256 amount);
    event BondSlashed(address indexed slashedAddress, uint256 amount, address indexed recipient);

    function setUp() public {
        // Deploy contracts
        bondToken = new HorizonToken(1_000_000_000 * 10 ** 18);
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        resolution = new ResolutionModule(address(outcomeToken), address(bondToken), arbitrator);

        // Register market
        address mockCollateral = address(0x999);
        outcomeToken.registerMarket(MARKET_ID, IERC20(mockCollateral));

        // Set resolution module as authorized
        outcomeToken.setResolutionAuthorization(address(resolution), true);

        // Fund test accounts
        bondToken.transfer(proposer, 10000 ether);
        bondToken.transfer(disputer, 10000 ether);

        // Approve bonds
        vm.prank(proposer);
        bondToken.approve(address(resolution), type(uint256).max);

        vm.prank(disputer);
        bondToken.approve(address(resolution), type(uint256).max);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(address(resolution.outcomeToken()), address(outcomeToken));
        assertEq(address(resolution.bondToken()), address(bondToken));
        assertEq(resolution.arbitrator(), arbitrator);
        assertEq(resolution.disputeWindow(), 48 hours);
        assertEq(resolution.minBond(), MIN_BOND);
    }

    // ============ Propose Resolution Tests ============

    function test_ProposeResolution() public {
        uint256 bondAmount = MIN_BOND;
        string memory evidence = "ipfs://QmExample";

        vm.expectEmit(true, true, true, true);
        emit ResolutionProposed(
            MARKET_ID,
            0, // outcomeId
            proposer,
            bondAmount,
            evidence,
            block.timestamp + 48 hours
        );

        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, bondAmount, evidence);

        // Check state
        (
            ResolutionModule.ResolutionState state,
            uint256 proposedOutcome,
            uint256 proposalTime,
            address storedProposer,
            uint256 proposerBond,
            ,
            ,
            string memory storedEvidence
        ) = resolution.resolutions(MARKET_ID);

        assertEq(uint8(state), uint8(ResolutionModule.ResolutionState.Proposed));
        assertEq(proposedOutcome, 0);
        assertEq(proposalTime, block.timestamp);
        assertEq(storedProposer, proposer);
        assertEq(proposerBond, bondAmount);
        assertEq(storedEvidence, evidence);
    }

    function test_RevertWhen_ProposeResolution_InsufficientBond() public {
        vm.prank(proposer);
        vm.expectRevert(ResolutionModule.InsufficientBond.selector);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND - 1, "evidence");
    }

    function test_RevertWhen_ProposeResolution_AlreadyProposed() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(proposer);
        vm.expectRevert(ResolutionModule.InvalidState.selector);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence2");
    }

    // ============ Dispute Tests ============

    function test_Dispute() public {
        // Propose resolution
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        uint256 disputeBond = MIN_BOND * 2;
        string memory reason = "Incorrect outcome";

        vm.expectEmit(true, true, false, true);
        emit Disputed(MARKET_ID, disputer, disputeBond, reason);

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, disputeBond, reason);

        // Check state
        (
            ResolutionModule.ResolutionState state,
            ,
            ,
            ,
            ,
            address storedDisputer,
            uint256 disputerBond,

        ) = resolution.resolutions(MARKET_ID);

        assertEq(uint8(state), uint8(ResolutionModule.ResolutionState.Disputed));
        assertEq(storedDisputer, disputer);
        assertEq(disputerBond, disputeBond);
    }

    function test_RevertWhen_Dispute_NotProposed() public {
        vm.prank(disputer);
        vm.expectRevert(ResolutionModule.InvalidState.selector);
        resolution.dispute(MARKET_ID, MIN_BOND, "reason");
    }

    function test_RevertWhen_Dispute_WindowClosed() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        // Warp past dispute window
        vm.warp(block.timestamp + 48 hours + 1);

        vm.prank(disputer);
        vm.expectRevert(ResolutionModule.DisputeWindowClosed.selector);
        resolution.dispute(MARKET_ID, MIN_BOND, "reason");
    }

    function test_RevertWhen_Dispute_InsufficientBond() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(disputer);
        vm.expectRevert(ResolutionModule.InsufficientBond.selector);
        resolution.dispute(MARKET_ID, MIN_BOND - 1, "reason");
    }

    // ============ Finalize Tests ============

    function test_Finalize() public {
        // Propose resolution
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        uint256 proposerBalanceBefore = bondToken.balanceOf(proposer);

        // Warp past dispute window
        vm.warp(block.timestamp + 48 hours + 1);

        vm.expectEmit(true, true, false, false);
        emit BondRefunded(proposer, MIN_BOND);

        vm.expectEmit(true, true, false, false);
        emit Finalized(MARKET_ID, 0, false);

        resolution.finalize(MARKET_ID);

        // Check state
        (ResolutionModule.ResolutionState state,,,,,,, ) = resolution.resolutions(MARKET_ID);
        assertEq(uint8(state), uint8(ResolutionModule.ResolutionState.Finalized));

        // Check outcome set
        assertTrue(outcomeToken.isResolved(MARKET_ID));
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0);

        // Check bond refunded
        assertEq(bondToken.balanceOf(proposer), proposerBalanceBefore + MIN_BOND);
    }

    function test_RevertWhen_Finalize_WindowOpen() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.expectRevert(ResolutionModule.DisputeWindowOpen.selector);
        resolution.finalize(MARKET_ID);
    }

    function test_RevertWhen_Finalize_NotProposed() public {
        vm.expectRevert(ResolutionModule.InvalidState.selector);
        resolution.finalize(MARKET_ID);
    }

    // ============ Finalize Disputed Tests ============

    function test_FinalizeDisputed_SlashProposer() public {
        // Propose and dispute
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND * 2, "Wrong outcome");

        uint256 disputerBalanceBefore = bondToken.balanceOf(disputer);

        // Arbitrator finalizes, slashing proposer
        vm.prank(arbitrator);
        resolution.finalizeDisputed(MARKET_ID, 1, true); // true = slash proposer

        // Check disputer received both bonds
        uint256 expectedReward = MIN_BOND + (MIN_BOND * 2); // Their bond + slashed bond
        assertEq(bondToken.balanceOf(disputer), disputerBalanceBefore + expectedReward);

        // Check outcome set to arbitrator's choice
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 1);
    }

    function test_FinalizeDisputed_SlashDisputer() public {
        // Propose and dispute
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND * 2, "Wrong outcome");

        uint256 proposerBalanceBefore = bondToken.balanceOf(proposer);

        // Arbitrator finalizes, slashing disputer
        vm.prank(arbitrator);
        resolution.finalizeDisputed(MARKET_ID, 0, false); // false = slash disputer

        // Check proposer received both bonds
        uint256 expectedReward = MIN_BOND + (MIN_BOND * 2); // Their bond + slashed bond
        assertEq(bondToken.balanceOf(proposer), proposerBalanceBefore + expectedReward);

        // Check outcome set to original proposal
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0);
    }

    function test_FinalizeDisputed_ByOwner() public {
        // Propose and dispute
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND, "Wrong outcome");

        // Owner can also finalize
        vm.prank(owner);
        resolution.finalizeDisputed(MARKET_ID, 1, true);

        assertEq(outcomeToken.winningOutcome(MARKET_ID), 1);
    }

    function test_RevertWhen_FinalizeDisputed_Unauthorized() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND, "reason");

        vm.prank(proposer);
        vm.expectRevert(ResolutionModule.Unauthorized.selector);
        resolution.finalizeDisputed(MARKET_ID, 0, false);
    }

    function test_RevertWhen_FinalizeDisputed_NotDisputed() public {
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        vm.prank(arbitrator);
        vm.expectRevert(ResolutionModule.InvalidState.selector);
        resolution.finalizeDisputed(MARKET_ID, 0, false);
    }

    // ============ Admin Functions Tests ============

    function test_SetDisputeWindow() public {
        uint256 newWindow = 24 hours;
        resolution.setDisputeWindow(newWindow);
        assertEq(resolution.disputeWindow(), newWindow);
    }

    function test_RevertWhen_SetDisputeWindow_Unauthorized() public {
        vm.prank(proposer);
        vm.expectRevert();
        resolution.setDisputeWindow(24 hours);
    }

    function test_SetMinBond() public {
        uint256 newBond = 5000 ether;
        resolution.setMinBond(newBond);
        assertEq(resolution.minBond(), newBond);
    }

    function test_RevertWhen_SetMinBond_Zero() public {
        vm.expectRevert(ResolutionModule.InvalidBondAmount.selector);
        resolution.setMinBond(0);
    }

    function test_SetArbitrator() public {
        address newArbitrator = address(0x99);
        resolution.setArbitrator(newArbitrator);
        assertEq(resolution.arbitrator(), newArbitrator);
    }

    // ============ View Functions Tests ============

    function test_CanDispute() public {
        // No proposal yet
        assertFalse(resolution.canDispute(MARKET_ID));

        // After proposal, within window
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");
        assertTrue(resolution.canDispute(MARKET_ID));

        // After window
        vm.warp(block.timestamp + 48 hours + 1);
        assertFalse(resolution.canDispute(MARKET_ID));
    }

    function test_CanFinalize() public {
        // No proposal yet
        assertFalse(resolution.canFinalize(MARKET_ID));

        // After proposal, within window
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");
        assertFalse(resolution.canFinalize(MARKET_ID));

        // After window
        vm.warp(block.timestamp + 48 hours + 1);
        assertTrue(resolution.canFinalize(MARKET_ID));
    }

    function test_GetDisputeTimeRemaining() public {
        // No proposal
        assertEq(resolution.getDisputeTimeRemaining(MARKET_ID), 0);

        // After proposal
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");
        assertEq(resolution.getDisputeTimeRemaining(MARKET_ID), 48 hours);

        // After some time
        vm.warp(block.timestamp + 24 hours);
        assertEq(resolution.getDisputeTimeRemaining(MARKET_ID), 24 hours);

        // After window
        vm.warp(block.timestamp + 24 hours + 1);
        assertEq(resolution.getDisputeTimeRemaining(MARKET_ID), 0);
    }

    // ============ Integration Tests ============

    function test_FullCycle_NoDispute() public {
        // Propose
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        // Wait
        vm.warp(block.timestamp + 48 hours + 1);

        // Finalize
        resolution.finalize(MARKET_ID);

        // Verify
        assertTrue(outcomeToken.isResolved(MARKET_ID));
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0);
    }

    function test_FullCycle_WithDispute_ProposerCorrect() public {
        // Propose
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        // Dispute
        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND, "I disagree");

        // Arbitrator sides with proposer
        vm.prank(arbitrator);
        resolution.finalizeDisputed(MARKET_ID, 0, false); // slash disputer

        // Verify
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 0);
    }

    function test_FullCycle_WithDispute_DisputerCorrect() public {
        // Propose
        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, MIN_BOND, "evidence");

        // Dispute
        vm.prank(disputer);
        resolution.dispute(MARKET_ID, MIN_BOND, "Wrong outcome!");

        // Arbitrator sides with disputer
        vm.prank(arbitrator);
        resolution.finalizeDisputed(MARKET_ID, 1, true); // slash proposer

        // Verify
        assertEq(outcomeToken.winningOutcome(MARKET_ID), 1);
    }

    function testFuzz_ProposeResolution(uint128 bondAmount) public {
        vm.assume(bondAmount >= MIN_BOND && bondAmount <= 10000 ether);

        // Fund proposer
        bondToken.transfer(proposer, bondAmount);

        vm.prank(proposer);
        resolution.proposeResolution(MARKET_ID, 0, bondAmount, "evidence");

        (, , , , uint256 storedBond, , , ) = resolution.resolutions(MARKET_ID);
        assertEq(storedBond, bondAmount);
    }
}
