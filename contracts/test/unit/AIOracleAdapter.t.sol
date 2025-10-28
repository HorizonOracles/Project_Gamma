// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/AIOracleAdapter.sol";
import "../../src/ResolutionModule.sol";
import "../../src/OutcomeToken.sol";
import "../../src/HorizonToken.sol";

contract AIOracleAdapterTest is Test {
    AIOracleAdapter public adapter;
    ResolutionModule public resolution;
    OutcomeToken public outcomeToken;
    HorizonToken public bondToken;

    address public owner = address(this);
    address public aiSignerPrivateKeyHolder; // We'll use vm.sign with a known key
    uint256 public aiSignerPrivateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
    address public aiSigner; // Derived from private key
    address public proposer = address(0x3);
    address public arbitrator = address(0x4);

    uint256 public constant MARKET_ID = 1;
    uint256 public constant MIN_BOND = 1000 ether;

    event SignerAdded(address indexed signer);
    event SignerRemoved(address indexed signer);
    event AIProposalSubmitted(
        uint256 indexed marketId,
        uint256 indexed outcomeId,
        address indexed proposer,
        address aiSigner,
        uint256 bondAmount,
        bytes32 signatureHash
    );

    function setUp() public {
        // Derive AI signer address from private key
        aiSigner = vm.addr(aiSignerPrivateKey);

        // Deploy contracts
        bondToken = new HorizonToken(1_000_000_000 * 10 ** 18);
        outcomeToken = new OutcomeToken("https://api.example.com/{id}");
        resolution = new ResolutionModule(address(outcomeToken), address(bondToken), arbitrator);
        adapter = new AIOracleAdapter(address(resolution), address(bondToken), aiSigner);

        // Register market
        address mockCollateral = address(0x999);
        outcomeToken.registerMarket(MARKET_ID, IERC20(mockCollateral));

        // Set resolution module as authorized
        outcomeToken.setResolutionAuthorization(address(resolution), true);

        // Authorize adapter to call resolution module (adapter will call proposeResolution)
        // The adapter needs to have HORIZON tokens approved

        // Fund proposer with HORIZON
        bondToken.transfer(proposer, 10000 ether);

        // Proposer approves adapter
        vm.prank(proposer);
        bondToken.approve(address(adapter), type(uint256).max);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(address(adapter.resolutionModule()), address(resolution));
        assertEq(address(adapter.bondToken()), address(bondToken));
        assertTrue(adapter.allowedSigners(aiSigner));
        assertNotEq(adapter.DOMAIN_SEPARATOR(), bytes32(0));
    }

    function test_RevertWhen_Constructor_InvalidAddress() public {
        vm.expectRevert(AIOracleAdapter.InvalidAddress.selector);
        new AIOracleAdapter(address(0), address(bondToken), aiSigner);

        vm.expectRevert(AIOracleAdapter.InvalidAddress.selector);
        new AIOracleAdapter(address(resolution), address(0), aiSigner);
    }

    // ============ Signer Management Tests ============

    function test_AddSigner() public {
        address newSigner = address(0x5);

        vm.expectEmit(true, false, false, false);
        emit SignerAdded(newSigner);

        adapter.addSigner(newSigner);
        assertTrue(adapter.allowedSigners(newSigner));
    }

    function test_RemoveSigner() public {
        vm.expectEmit(true, false, false, false);
        emit SignerRemoved(aiSigner);

        adapter.removeSigner(aiSigner);
        assertFalse(adapter.allowedSigners(aiSigner));
    }

    function test_RevertWhen_AddSigner_Unauthorized() public {
        vm.prank(proposer);
        vm.expectRevert();
        adapter.addSigner(address(0x5));
    }

    function test_RevertWhen_AddSigner_ZeroAddress() public {
        vm.expectRevert(AIOracleAdapter.InvalidAddress.selector);
        adapter.addSigner(address(0));
    }

    // ============ EIP-712 Signature Tests ============

    function test_ProposeAI_ValidSignature() public {
        // Create proposal
        string[] memory evidence = new string[](2);
        evidence[0] = "https://example.com/evidence1";
        evidence[1] = "https://example.com/evidence2";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        // Sign proposal
        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Submit proposal
        vm.expectEmit(true, true, true, false);
        emit AIProposalSubmitted(MARKET_ID, 0, proposer, aiSigner, MIN_BOND, keccak256(signature));

        vm.prank(proposer);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);

        // Verify proposal was submitted to ResolutionModule
        (ResolutionModule.ResolutionState state,,,,,,, ) = resolution.resolutions(MARKET_ID);
        assertEq(uint8(state), uint8(ResolutionModule.ResolutionState.Proposed));
    }

    function test_RevertWhen_ProposeAI_InvalidSigner() public {
        // Create proposal
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        // Sign with wrong private key
        uint256 wrongKey = 0xabcdef;
        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(wrongKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Should revert because signer is not allowed
        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.InvalidSigner.selector);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    function test_RevertWhen_ProposeAI_InvalidSignatureLength() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        // Invalid signature (too short)
        bytes memory invalidSig = new bytes(64);

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.InvalidSignature.selector);
        adapter.proposeAI(proposal, invalidSig, MIN_BOND, evidence);
    }

    function test_RevertWhen_ProposeAI_MalformedSignature() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        // Malformed signature (all zeros, ecrecover will return address(0))
        bytes memory malformedSig = new bytes(65);

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.InvalidSignature.selector);
        adapter.proposeAI(proposal, malformedSig, MIN_BOND, evidence);
    }

    // ============ Timestamp Validation Tests ============

    function test_RevertWhen_ProposeAI_NotYetValid() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        // Signature not valid until 1 hour from now
        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp + 1 hours,
            deadline: block.timestamp + 2 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.SignatureNotYetValid.selector);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    function test_RevertWhen_ProposeAI_Expired() public {
        // Set timestamp to a reasonable value to avoid underflow
        vm.warp(block.timestamp + 10 hours);

        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        // Signature expired 1 hour ago
        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp - 2 hours,
            deadline: block.timestamp - 1 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.SignatureExpired.selector);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    function test_ProposeAI_AtDeadline() public {
        // Set timestamp to a reasonable value to avoid underflow
        vm.warp(block.timestamp + 10 hours);

        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        // Signature expires exactly now
        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp - 1 hours,
            deadline: block.timestamp
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Should succeed (deadline is inclusive: block.timestamp <= deadline)
        vm.prank(proposer);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    // ============ Evidence Hash Tests ============

    function test_RevertWhen_ProposeAI_InvalidEvidenceHash() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        // Proposal with wrong evidence hash
        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: keccak256("wrong"),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.InvalidEvidenceHash.selector);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    function test_HashEvidence_SingleURI() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        bytes32 hash = adapter.hashEvidence(evidence);
        assertEq(hash, keccak256(abi.encode(evidence)));
    }

    function test_HashEvidence_MultipleURIs() public {
        string[] memory evidence = new string[](3);
        evidence[0] = "https://example.com/1";
        evidence[1] = "https://example.com/2";
        evidence[2] = "https://example.com/3";

        bytes32 hash = adapter.hashEvidence(evidence);
        assertEq(hash, keccak256(abi.encode(evidence)));
    }

    function test_HashEvidence_EmptyArray() public {
        string[] memory evidence = new string[](0);
        bytes32 hash = adapter.hashEvidence(evidence);
        assertEq(hash, keccak256(abi.encode(evidence)));
    }

    // ============ Replay Protection Tests ============

    function test_RevertWhen_ProposeAI_ReplayAttack() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // First submission should succeed
        vm.prank(proposer);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);

        // Verify signature is marked as used
        assertTrue(adapter.isSignatureUsed(signature));

        // Second submission with same signature should fail
        // Need to use a different market to avoid "InvalidState" from ResolutionModule
        uint256 marketId2 = 2;
        outcomeToken.registerMarket(marketId2, IERC20(address(0x999)));

        proposal.marketId = marketId2;
        // Note: We're using the SAME signature, which should fail even though marketId changed
        // because signature hash stays the same

        vm.prank(proposer);
        vm.expectRevert(AIOracleAdapter.SignatureAlreadyUsed.selector);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);
    }

    function test_IsSignatureUsed() public {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Initially not used
        assertFalse(adapter.isSignatureUsed(signature));

        // After submission, should be used
        vm.prank(proposer);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);

        assertTrue(adapter.isSignatureUsed(signature));
    }

    // ============ Integration Tests ============

    function test_FullFlow_AIProposal() public {
        // Set timestamp to a reasonable value to avoid underflow
        vm.warp(block.timestamp + 10 hours);

        // 1. AI signs proposal
        string[] memory evidence = new string[](3);
        evidence[0] = "https://news.com/article1";
        evidence[1] = "https://news.com/article2";
        evidence[2] = "https://wikipedia.org/fact";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 1, // NO wins
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp - 1 hours,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest = adapter.getProposalHash(proposal);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(aiSignerPrivateKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // 2. Anyone can submit with bond
        uint256 proposerBalanceBefore = bondToken.balanceOf(proposer);

        vm.prank(proposer);
        adapter.proposeAI(proposal, signature, MIN_BOND, evidence);

        // 3. Verify bond was transferred
        assertEq(bondToken.balanceOf(proposer), proposerBalanceBefore - MIN_BOND);

        // 4. Verify resolution was created
        (
            ResolutionModule.ResolutionState state,
            uint256 proposedOutcome,
            ,
            address resolutionProposer,
            uint256 proposerBond,
            ,
            ,

        ) = resolution.resolutions(MARKET_ID);

        assertEq(uint8(state), uint8(ResolutionModule.ResolutionState.Proposed));
        assertEq(proposedOutcome, 1);
        assertEq(resolutionProposer, address(adapter)); // Adapter is the proposer
        assertEq(proposerBond, MIN_BOND);
    }

    function test_MultipleSignersCanSubmit() public {
        // Add second signer
        uint256 signer2Key = 0xfedcba;
        address signer2 = vm.addr(signer2Key);
        adapter.addSigner(signer2);

        // First signer submits market 1
        uint256 market1 = MARKET_ID;
        string[] memory evidence1 = new string[](1);
        evidence1[0] = "https://example.com/1";

        AIOracleAdapter.ProposedOutcome memory proposal1 = AIOracleAdapter.ProposedOutcome({
            marketId: market1,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence1),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest1 = adapter.getProposalHash(proposal1);
        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(aiSignerPrivateKey, digest1);
        bytes memory sig1 = abi.encodePacked(r1, s1, v1);

        vm.prank(proposer);
        adapter.proposeAI(proposal1, sig1, MIN_BOND, evidence1);

        // Second signer submits market 2
        uint256 market2 = 2;
        outcomeToken.registerMarket(market2, IERC20(address(0x999)));

        string[] memory evidence2 = new string[](1);
        evidence2[0] = "https://example.com/2";

        AIOracleAdapter.ProposedOutcome memory proposal2 = AIOracleAdapter.ProposedOutcome({
            marketId: market2,
            outcomeId: 1,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence2),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 digest2 = adapter.getProposalHash(proposal2);
        (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(signer2Key, digest2);
        bytes memory sig2 = abi.encodePacked(r2, s2, v2);

        vm.prank(proposer);
        adapter.proposeAI(proposal2, sig2, MIN_BOND, evidence2);

        // Both should succeed
        (ResolutionModule.ResolutionState state1,,,,,,, ) = resolution.resolutions(market1);
        (ResolutionModule.ResolutionState state2,,,,,,, ) = resolution.resolutions(market2);

        assertEq(uint8(state1), uint8(ResolutionModule.ResolutionState.Proposed));
        assertEq(uint8(state2), uint8(ResolutionModule.ResolutionState.Proposed));
    }

    // ============ View Function Tests ============

    function test_GetProposalHash() public view {
        string[] memory evidence = new string[](1);
        evidence[0] = "https://example.com/evidence";

        AIOracleAdapter.ProposedOutcome memory proposal = AIOracleAdapter.ProposedOutcome({
            marketId: MARKET_ID,
            outcomeId: 0,
            closeTime: block.timestamp,
            evidenceHash: adapter.hashEvidence(evidence),
            notBefore: block.timestamp,
            deadline: block.timestamp + 1 hours
        });

        bytes32 hash = adapter.getProposalHash(proposal);
        assertNotEq(hash, bytes32(0));
    }

    function test_DomainSeparator() public view {
        bytes32 domainSep = adapter.DOMAIN_SEPARATOR();
        assertNotEq(domainSep, bytes32(0));

        // Domain separator should include contract address and chain ID
        // This makes signatures non-transferable across chains/deployments
    }

    function testFuzz_HashEvidence(string[] calldata evidence) public view {
        vm.assume(evidence.length < 100); // Reasonable limit
        bytes32 hash = adapter.hashEvidence(evidence);
        assertEq(hash, keccak256(abi.encode(evidence)));
    }
}
