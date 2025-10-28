// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./ResolutionModule.sol";

/**
 * @title AIOracleAdapter
 * @notice Adapter contract for AI-powered resolution proposals using EIP-712 signatures
 * @dev Verifies cryptographic signatures from authorized AI signers before submitting proposals
 *
 *      Flow:
 *      1. AI agent analyzes market outcome off-chain
 *      2. AI agent signs ProposedOutcome with private key (EIP-712)
 *      3. Anyone calls proposeAI() with signature + evidence
 *      4. Contract verifies signature is from allowedSigner
 *      5. Contract validates timestamps and evidence
 *      6. Contract transfers bond from msg.sender
 *      7. Contract calls ResolutionModule.proposeResolution()
 */
contract AIOracleAdapter is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ EIP-712 Constants ============

    /// @notice EIP-712 domain type hash
    bytes32 public constant EIP712_DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");

    /// @notice ProposedOutcome type hash for EIP-712
    bytes32 public constant PROPOSED_OUTCOME_TYPEHASH = keccak256(
        "ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)"
    );

    /// @notice EIP-712 domain separator (computed in constructor)
    bytes32 public immutable DOMAIN_SEPARATOR;

    // ============ Errors ============

    error InvalidSigner();
    error InvalidSignature();
    error SignatureExpired();
    error SignatureNotYetValid();
    error InvalidEvidenceHash();
    error SignatureAlreadyUsed();
    error MarketNotClosed();
    error InvalidAddress();

    // ============ Events ============

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

    // ============ Structs ============

    /**
     * @notice Proposed outcome signed by AI
     * @param marketId Market identifier
     * @param outcomeId Proposed winning outcome (0=Yes, 1=No)
     * @param closeTime Market close timestamp (must match actual market)
     * @param evidenceHash Keccak256 hash of sorted evidence URIs
     * @param notBefore Signature not valid before this timestamp
     * @param deadline Signature expires after this timestamp
     */
    struct ProposedOutcome {
        uint256 marketId;
        uint256 outcomeId;
        uint256 closeTime;
        bytes32 evidenceHash;
        uint256 notBefore;
        uint256 deadline;
    }

    // ============ State Variables ============

    /// @notice Resolution module to submit proposals to
    ResolutionModule public immutable resolutionModule;

    /// @notice Bond token (HORIZON)
    IERC20 public immutable bondToken;

    /// @notice Mapping of allowed AI signer addresses
    mapping(address => bool) public allowedSigners;

    /// @notice Mapping to prevent signature replay attacks
    mapping(bytes32 => bool) public usedSignatures;

    // ============ Constructor ============

    /**
     * @notice Initializes the AIOracleAdapter
     * @param _resolutionModule Address of ResolutionModule contract
     * @param _bondToken Address of bond token (HORIZON)
     * @param _initialSigner Initial allowed AI signer address
     */
    constructor(address _resolutionModule, address _bondToken, address _initialSigner) Ownable(msg.sender) {
        if (_resolutionModule == address(0) || _bondToken == address(0)) revert InvalidAddress();

        resolutionModule = ResolutionModule(_resolutionModule);
        bondToken = IERC20(_bondToken);

        // Compute EIP-712 domain separator
        DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                EIP712_DOMAIN_TYPEHASH,
                keccak256(bytes("AIOracleAdapter")),
                keccak256(bytes("1")),
                block.chainid,
                address(this)
            )
        );

        // Add initial signer
        if (_initialSigner != address(0)) {
            allowedSigners[_initialSigner] = true;
            emit SignerAdded(_initialSigner);
        }
    }

    // ============ AI Proposal Functions ============

    /**
     * @notice Submits an AI-signed resolution proposal
     * @param proposal ProposedOutcome struct with all parameters
     * @param signature EIP-712 signature from allowed AI signer (65 bytes: r, s, v)
     * @param bondAmount Amount of HORIZON tokens to bond
     * @param evidenceURIs Array of evidence URLs (must hash to proposal.evidenceHash)
     */
    function proposeAI(
        ProposedOutcome calldata proposal,
        bytes calldata signature,
        uint256 bondAmount,
        string[] calldata evidenceURIs
    ) external nonReentrant {
        // 1. Validate timestamps
        if (block.timestamp < proposal.notBefore) revert SignatureNotYetValid();
        if (block.timestamp > proposal.deadline) revert SignatureExpired();

        // 2. Verify signature hasn't been used
        bytes32 signatureHash = keccak256(signature);
        if (usedSignatures[signatureHash]) revert SignatureAlreadyUsed();

        // 3. Verify evidence hash
        bytes32 computedHash = _hashEvidence(evidenceURIs);
        if (computedHash != proposal.evidenceHash) revert InvalidEvidenceHash();

        // 4. Verify EIP-712 signature
        address signer = _verifySignature(proposal, signature);
        if (!allowedSigners[signer]) revert InvalidSigner();

        // 5. Mark signature as used (replay protection)
        usedSignatures[signatureHash] = true;

        // 6. Transfer bond from proposer to this contract
        bondToken.safeTransferFrom(msg.sender, address(this), bondAmount);

        // 7. Approve ResolutionModule to spend bond
        bondToken.safeIncreaseAllowance(address(resolutionModule), bondAmount);

        // 8. Build evidence URI string (concatenate with newlines)
        string memory evidenceURI = _concatenateEvidence(evidenceURIs);

        // 9. Submit proposal to ResolutionModule
        resolutionModule.proposeResolution(proposal.marketId, proposal.outcomeId, bondAmount, evidenceURI);

        emit AIProposalSubmitted(
            proposal.marketId, proposal.outcomeId, msg.sender, signer, bondAmount, signatureHash
        );
    }

    // ============ Admin Functions ============

    /**
     * @notice Adds an allowed AI signer
     * @param signer Address to add as allowed signer
     */
    function addSigner(address signer) external onlyOwner {
        if (signer == address(0)) revert InvalidAddress();
        allowedSigners[signer] = true;
        emit SignerAdded(signer);
    }

    /**
     * @notice Removes an allowed AI signer
     * @param signer Address to remove from allowed signers
     */
    function removeSigner(address signer) external onlyOwner {
        allowedSigners[signer] = false;
        emit SignerRemoved(signer);
    }

    // ============ Internal Functions ============

    /**
     * @notice Verifies EIP-712 signature for a proposed outcome
     * @param proposal ProposedOutcome to verify
     * @param signature 65-byte signature (r, s, v)
     * @return signer Address that signed the message
     */
    function _verifySignature(ProposedOutcome calldata proposal, bytes calldata signature)
        internal
        view
        returns (address signer)
    {
        // Require 65-byte signature
        if (signature.length != 65) revert InvalidSignature();

        // Split signature into r, s, v
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 32))
            v := byte(0, calldataload(add(signature.offset, 64)))
        }

        // Compute EIP-712 struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PROPOSED_OUTCOME_TYPEHASH,
                proposal.marketId,
                proposal.outcomeId,
                proposal.closeTime,
                proposal.evidenceHash,
                proposal.notBefore,
                proposal.deadline
            )
        );

        // Compute EIP-712 digest
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));

        // Recover signer from signature
        signer = ecrecover(digest, v, r, s);
        if (signer == address(0)) revert InvalidSignature();
    }

    /**
     * @notice Hashes evidence URIs for verification
     * @param evidenceURIs Array of evidence URL strings
     * @return Hash of encoded evidence URIs
     */
    function _hashEvidence(string[] calldata evidenceURIs) internal pure returns (bytes32) {
        // Use abi.encode for arrays (abi.encodePacked doesn't support string arrays)
        return keccak256(abi.encode(evidenceURIs));
    }

    /**
     * @notice Concatenates evidence URIs into single string
     * @param evidenceURIs Array of evidence URL strings
     * @return Concatenated string with newlines
     */
    function _concatenateEvidence(string[] calldata evidenceURIs) internal pure returns (string memory) {
        if (evidenceURIs.length == 0) return "";
        if (evidenceURIs.length == 1) return evidenceURIs[0];

        // Concatenate with newlines
        string memory result = evidenceURIs[0];
        for (uint256 i = 1; i < evidenceURIs.length; i++) {
            result = string(abi.encodePacked(result, "\n", evidenceURIs[i]));
        }
        return result;
    }

    // ============ View Functions ============

    /**
     * @notice Computes the EIP-712 hash for a proposed outcome
     * @param proposal ProposedOutcome to hash
     * @return Hash of the proposal for signing
     */
    function getProposalHash(ProposedOutcome calldata proposal) external view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                PROPOSED_OUTCOME_TYPEHASH,
                proposal.marketId,
                proposal.outcomeId,
                proposal.closeTime,
                proposal.evidenceHash,
                proposal.notBefore,
                proposal.deadline
            )
        );
        return keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));
    }

    /**
     * @notice Checks if a signature has been used
     * @param signature Signature bytes to check
     * @return True if signature has been used
     */
    function isSignatureUsed(bytes calldata signature) external view returns (bool) {
        return usedSignatures[keccak256(signature)];
    }

    /**
     * @notice Computes evidence hash from URIs
     * @param evidenceURIs Array of evidence URL strings
     * @return Hash of the evidence
     */
    function hashEvidence(string[] calldata evidenceURIs) external pure returns (bytes32) {
        return _hashEvidence(evidenceURIs);
    }
}
