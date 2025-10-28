// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./OutcomeToken.sol";

/**
 * @title ResolutionModule
 * @notice Manages the resolution lifecycle for prediction markets
 * @dev Implements a state machine for market resolution with dispute mechanism
 *
 *      State Flow:
 *      None → Proposed (via proposeResolution)
 *      Proposed → Disputed (via dispute, within dispute window)
 *      Proposed → Finalized (via finalize, after dispute window)
 *      Disputed → Finalized (via finalizeDisputed, only by arbitrator)
 */
contract ResolutionModule is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ Enums ============

    enum ResolutionState {
        None,       // No resolution proposed
        Proposed,   // Resolution proposed, dispute window open
        Disputed,   // Resolution disputed, awaiting arbitration
        Finalized   // Resolution finalized, outcome set
    }

    // ============ Errors ============

    error MarketAlreadyResolved();
    error InvalidOutcome();
    error InvalidState();
    error DisputeWindowClosed();
    error DisputeWindowOpen();
    error InsufficientBond();
    error Unauthorized();
    error InvalidBondAmount();

    // ============ Events ============

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

    event BondRefunded(
        address indexed recipient,
        uint256 amount
    );

    event BondSlashed(
        address indexed slashedAddress,
        uint256 amount,
        address indexed recipient
    );

    event DisputeWindowUpdated(uint256 oldWindow, uint256 newWindow);
    event MinBondUpdated(uint256 oldBond, uint256 newBond);
    event ArbitratorUpdated(address indexed oldArbitrator, address indexed newArbitrator);

    // ============ Structs ============

    struct Resolution {
        ResolutionState state;
        uint256 proposedOutcome;
        uint256 proposalTime;
        address proposer;
        uint256 proposerBond;
        address disputer;
        uint256 disputerBond;
        string evidenceURI;
    }

    // ============ Immutable State ============

    /// @notice Outcome token contract
    OutcomeToken public immutable outcomeToken;

    /// @notice Bond token (e.g., HORIZON)
    IERC20 public immutable bondToken;

    // ============ Mutable State ============

    /// @notice Resolution data for each market
    mapping(uint256 => Resolution) public resolutions;

    /// @notice Dispute window duration (default 48 hours)
    uint256 public disputeWindow = 48 hours;

    /// @notice Minimum bond required for proposing or disputing
    uint256 public minBond = 1000 ether; // 1000 HORIZON tokens

    /// @notice Arbitrator address (can resolve disputed markets)
    address public arbitrator;

    // ============ Constructor ============

    /**
     * @notice Initializes the resolution module
     * @param _outcomeToken Outcome token contract
     * @param _bondToken Bond token (HORIZON)
     * @param _arbitrator Initial arbitrator address
     */
    constructor(address _outcomeToken, address _bondToken, address _arbitrator) Ownable(msg.sender) {
        outcomeToken = OutcomeToken(_outcomeToken);
        bondToken = IERC20(_bondToken);
        arbitrator = _arbitrator;
    }

    // ============ Resolution Functions ============

    /**
     * @notice Proposes a resolution for a market
     * @param marketId Market identifier
     * @param outcomeId Proposed winning outcome
     * @param bondAmount Bond amount (must be >= minBond)
     * @param evidenceURI IPFS URI with evidence
     */
    function proposeResolution(
        uint256 marketId,
        uint256 outcomeId,
        uint256 bondAmount,
        string calldata evidenceURI
    ) external nonReentrant {
        Resolution storage resolution = resolutions[marketId];

        // Validate state
        if (resolution.state != ResolutionState.None) revert InvalidState();
        if (outcomeToken.isResolved(marketId)) revert MarketAlreadyResolved();
        if (bondAmount < minBond) revert InsufficientBond();

        // Transfer bond from proposer
        bondToken.safeTransferFrom(msg.sender, address(this), bondAmount);

        // Store resolution proposal
        resolution.state = ResolutionState.Proposed;
        resolution.proposedOutcome = outcomeId;
        resolution.proposalTime = block.timestamp;
        resolution.proposer = msg.sender;
        resolution.proposerBond = bondAmount;
        resolution.evidenceURI = evidenceURI;

        emit ResolutionProposed(
            marketId,
            outcomeId,
            msg.sender,
            bondAmount,
            evidenceURI,
            block.timestamp + disputeWindow
        );
    }

    /**
     * @notice Disputes a proposed resolution
     * @param marketId Market identifier
     * @param bondAmount Bond amount (must be >= minBond)
     * @param reason Reason for dispute
     */
    function dispute(
        uint256 marketId,
        uint256 bondAmount,
        string calldata reason
    ) external nonReentrant {
        Resolution storage resolution = resolutions[marketId];

        // Validate state
        if (resolution.state != ResolutionState.Proposed) revert InvalidState();
        if (block.timestamp > resolution.proposalTime + disputeWindow) revert DisputeWindowClosed();
        if (bondAmount < minBond) revert InsufficientBond();

        // Transfer bond from disputer
        bondToken.safeTransferFrom(msg.sender, address(this), bondAmount);

        // Update resolution state
        resolution.state = ResolutionState.Disputed;
        resolution.disputer = msg.sender;
        resolution.disputerBond = bondAmount;

        emit Disputed(marketId, msg.sender, bondAmount, reason);
    }

    /**
     * @notice Finalizes a non-disputed resolution after dispute window
     * @param marketId Market identifier
     */
    function finalize(uint256 marketId) external nonReentrant {
        Resolution storage resolution = resolutions[marketId];

        // Validate state
        if (resolution.state != ResolutionState.Proposed) revert InvalidState();
        if (block.timestamp <= resolution.proposalTime + disputeWindow) revert DisputeWindowOpen();

        // Finalize resolution
        resolution.state = ResolutionState.Finalized;
        outcomeToken.setWinningOutcome(marketId, resolution.proposedOutcome);

        // Refund proposer bond
        bondToken.safeTransfer(resolution.proposer, resolution.proposerBond);
        emit BondRefunded(resolution.proposer, resolution.proposerBond);

        emit Finalized(marketId, resolution.proposedOutcome, false);
    }

    /**
     * @notice Finalizes a disputed resolution (arbitrator only)
     * @param marketId Market identifier
     * @param outcomeId Final winning outcome
     * @param slashProposer True to slash proposer, false to slash disputer
     */
    function finalizeDisputed(
        uint256 marketId,
        uint256 outcomeId,
        bool slashProposer
    ) external nonReentrant {
        if (msg.sender != arbitrator && msg.sender != owner()) revert Unauthorized();

        Resolution storage resolution = resolutions[marketId];

        // Validate state
        if (resolution.state != ResolutionState.Disputed) revert InvalidState();

        // Finalize resolution
        resolution.state = ResolutionState.Finalized;
        outcomeToken.setWinningOutcome(marketId, outcomeId);

        // Handle bonds
        if (slashProposer) {
            // Slash proposer, refund disputer
            bondToken.safeTransfer(resolution.disputer, resolution.disputerBond);
            emit BondRefunded(resolution.disputer, resolution.disputerBond);

            bondToken.safeTransfer(resolution.disputer, resolution.proposerBond);
            emit BondSlashed(resolution.proposer, resolution.proposerBond, resolution.disputer);
        } else {
            // Slash disputer, refund proposer
            bondToken.safeTransfer(resolution.proposer, resolution.proposerBond);
            emit BondRefunded(resolution.proposer, resolution.proposerBond);

            bondToken.safeTransfer(resolution.proposer, resolution.disputerBond);
            emit BondSlashed(resolution.disputer, resolution.disputerBond, resolution.proposer);
        }

        emit Finalized(marketId, outcomeId, true);
    }

    // ============ Admin Functions ============

    /**
     * @notice Updates the dispute window duration
     * @param newWindow New dispute window in seconds
     */
    function setDisputeWindow(uint256 newWindow) external onlyOwner {
        uint256 oldWindow = disputeWindow;
        disputeWindow = newWindow;
        emit DisputeWindowUpdated(oldWindow, newWindow);
    }

    /**
     * @notice Updates the minimum bond amount
     * @param newMinBond New minimum bond amount
     */
    function setMinBond(uint256 newMinBond) external onlyOwner {
        if (newMinBond == 0) revert InvalidBondAmount();
        uint256 oldBond = minBond;
        minBond = newMinBond;
        emit MinBondUpdated(oldBond, newMinBond);
    }

    /**
     * @notice Updates the arbitrator address
     * @param newArbitrator New arbitrator address
     */
    function setArbitrator(address newArbitrator) external onlyOwner {
        address oldArbitrator = arbitrator;
        arbitrator = newArbitrator;
        emit ArbitratorUpdated(oldArbitrator, newArbitrator);
    }

    // ============ View Functions ============

    /**
     * @notice Gets the resolution state for a market
     * @param marketId Market identifier
     * @return state Current resolution state
     */
    function getResolutionState(uint256 marketId) external view returns (ResolutionState) {
        return resolutions[marketId].state;
    }

    /**
     * @notice Checks if a resolution can be disputed
     * @param marketId Market identifier
     * @return canDispute True if within dispute window
     */
    function canDispute(uint256 marketId) external view returns (bool) {
        Resolution storage resolution = resolutions[marketId];
        return resolution.state == ResolutionState.Proposed
            && block.timestamp <= resolution.proposalTime + disputeWindow;
    }

    /**
     * @notice Checks if a resolution can be finalized
     * @param marketId Market identifier
     * @return canFinalize True if dispute window passed
     */
    function canFinalize(uint256 marketId) external view returns (bool) {
        Resolution storage resolution = resolutions[marketId];
        return resolution.state == ResolutionState.Proposed
            && block.timestamp > resolution.proposalTime + disputeWindow;
    }

    /**
     * @notice Gets time remaining in dispute window
     * @param marketId Market identifier
     * @return timeRemaining Seconds remaining (0 if window closed)
     */
    function getDisputeTimeRemaining(uint256 marketId) external view returns (uint256) {
        Resolution storage resolution = resolutions[marketId];
        if (resolution.state != ResolutionState.Proposed) return 0;

        uint256 deadline = resolution.proposalTime + disputeWindow;
        if (block.timestamp >= deadline) return 0;

        return deadline - block.timestamp;
    }
}
