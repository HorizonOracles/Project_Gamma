// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title OutcomeToken
 * @notice ERC-1155 token contract for prediction market outcome shares
 * @dev Each market has multiple outcome token IDs (typically 2 for binary markets: Yes=0, No=1)
 *      Token IDs are encoded as: (marketId << 8) | outcomeId
 *      Only authorized AMM contracts can mint/burn outcome tokens
 */
contract OutcomeToken is ERC1155, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error Unauthorized();
    error MarketNotResolved();
    error MarketAlreadyResolved();
    error InvalidOutcome();
    error NoTokensToRedeem();
    error InvalidTokenId();

    // ============ Events ============

    event OutcomeMinted(uint256 indexed marketId, uint256 indexed outcomeId, address indexed to, uint256 amount);
    event OutcomeBurned(uint256 indexed marketId, uint256 indexed outcomeId, address indexed from, uint256 amount);
    event WinningOutcomeSet(uint256 indexed marketId, uint256 indexed winningOutcomeId);
    event Redeemed(
        uint256 indexed marketId, address indexed user, uint256 winningTokens, uint256 collateralPaid
    );
    event AMMAuthorized(address indexed amm, bool authorized);
    event ResolutionAuthorized(address indexed resolver, bool authorized);

    // ============ State Variables ============

    /// @notice Mapping of market ID to its collateral token
    mapping(uint256 => IERC20) public marketCollateral;

    /// @notice Mapping of market ID to winning outcome ID (MAX_UINT256 = not resolved)
    mapping(uint256 => uint256) public winningOutcome;

    /// @notice Mapping of AMM addresses authorized to mint/burn
    mapping(address => bool) public authorizedAMMs;

    /// @notice Mapping of Resolution module addresses authorized to set winning outcomes
    mapping(address => bool) public authorizedResolvers;

    /// @notice Constant representing unresolved market
    uint256 public constant UNRESOLVED = type(uint256).max;

    // ============ Constructor ============

    /**
     * @notice Initializes the OutcomeToken contract
     * @param _uri Base URI for token metadata
     */
    constructor(string memory _uri) ERC1155(_uri) Ownable(msg.sender) {}

    // ============ Modifiers ============

    modifier onlyAuthorizedAMM() {
        if (!authorizedAMMs[msg.sender]) revert Unauthorized();
        _;
    }

    modifier onlyAuthorizedResolver() {
        if (!authorizedResolvers[msg.sender]) revert Unauthorized();
        _;
    }

    // ============ Admin Functions ============

    /**
     * @notice Authorizes or deauthorizes an AMM contract
     * @param amm Address of the AMM contract
     * @param authorized Whether the AMM is authorized
     */
    function setAMMAuthorization(address amm, bool authorized) external onlyOwner {
        authorizedAMMs[amm] = authorized;
        emit AMMAuthorized(amm, authorized);
    }

    /**
     * @notice Authorizes or deauthorizes a Resolution module
     * @param resolver Address of the Resolution module
     * @param authorized Whether the resolver is authorized
     */
    function setResolutionAuthorization(address resolver, bool authorized) external onlyOwner {
        authorizedResolvers[resolver] = authorized;
        emit ResolutionAuthorized(resolver, authorized);
    }

    /**
     * @notice Registers a new market with its collateral token
     * @param marketId The market identifier
     * @param collateral The ERC20 token used as collateral for this market
     */
    function registerMarket(uint256 marketId, IERC20 collateral) external onlyOwner {
        marketCollateral[marketId] = collateral;
        winningOutcome[marketId] = UNRESOLVED;
    }

    /**
     * @notice Sets the winning outcome for a resolved market
     * @dev Can only be called by authorized Resolution modules
     * @param marketId The market identifier
     * @param outcomeId The winning outcome ID
     */
    function setWinningOutcome(uint256 marketId, uint256 outcomeId) external onlyAuthorizedResolver {
        if (winningOutcome[marketId] != UNRESOLVED) revert MarketAlreadyResolved();
        winningOutcome[marketId] = outcomeId;
        emit WinningOutcomeSet(marketId, outcomeId);
    }

    // ============ AMM Functions ============

    /**
     * @notice Mints outcome tokens for a market
     * @dev Only callable by authorized AMM contracts
     * @param marketId The market identifier
     * @param outcomeId The outcome identifier (e.g., 0=Yes, 1=No)
     * @param to The address to mint tokens to
     * @param amount The amount of tokens to mint
     */
    function mintOutcome(uint256 marketId, uint256 outcomeId, address to, uint256 amount)
        external
        onlyAuthorizedAMM
    {
        uint256 tokenId = encodeTokenId(marketId, outcomeId);
        _mint(to, tokenId, amount, "");
        emit OutcomeMinted(marketId, outcomeId, to, amount);
    }

    /**
     * @notice Burns outcome tokens from a market
     * @dev Only callable by authorized AMM contracts
     * @param marketId The market identifier
     * @param outcomeId The outcome identifier
     * @param from The address to burn tokens from
     * @param amount The amount of tokens to burn
     */
    function burnOutcome(uint256 marketId, uint256 outcomeId, address from, uint256 amount)
        external
        onlyAuthorizedAMM
    {
        uint256 tokenId = encodeTokenId(marketId, outcomeId);
        _burn(from, tokenId, amount);
        emit OutcomeBurned(marketId, outcomeId, from, amount);
    }

    // ============ Redemption Functions ============

    /**
     * @notice Redeems all outcome tokens for a resolved market
     * @dev Burns winning outcome tokens and pays out collateral 1:1
     * @param marketId The market identifier
     * @return payout The amount of collateral paid out
     */
    function redeem(uint256 marketId) external nonReentrant returns (uint256 payout) {
        uint256 winning = winningOutcome[marketId];
        if (winning == UNRESOLVED) revert MarketNotResolved();

        uint256 winningTokenId = encodeTokenId(marketId, winning);
        uint256 winningBalance = balanceOf(msg.sender, winningTokenId);

        if (winningBalance == 0) revert NoTokensToRedeem();

        // Burn the winning tokens
        _burn(msg.sender, winningTokenId, winningBalance);

        // Pay out collateral 1:1
        payout = winningBalance;
        IERC20 collateral = marketCollateral[marketId];
        collateral.safeTransfer(msg.sender, payout);

        emit Redeemed(marketId, msg.sender, winningBalance, payout);
    }

    /**
     * @notice Redeems outcome tokens for a specific outcome
     * @dev Allows partial redemption by specifying amount
     * @param marketId The market identifier
     * @param amount The amount of tokens to redeem
     * @return payout The amount of collateral paid out
     */
    function redeemAmount(uint256 marketId, uint256 amount) external nonReentrant returns (uint256 payout) {
        uint256 winning = winningOutcome[marketId];
        if (winning == UNRESOLVED) revert MarketNotResolved();

        uint256 winningTokenId = encodeTokenId(marketId, winning);
        uint256 winningBalance = balanceOf(msg.sender, winningTokenId);

        if (winningBalance == 0 || amount == 0 || amount > winningBalance) revert NoTokensToRedeem();

        // Burn the winning tokens
        _burn(msg.sender, winningTokenId, amount);

        // Pay out collateral 1:1
        payout = amount;
        IERC20 collateral = marketCollateral[marketId];
        collateral.safeTransfer(msg.sender, payout);

        emit Redeemed(marketId, msg.sender, amount, payout);
    }

    // ============ View Functions ============

    /**
     * @notice Encodes a market ID and outcome ID into a single token ID
     * @param marketId The market identifier
     * @param outcomeId The outcome identifier
     * @return tokenId The encoded token ID
     */
    function encodeTokenId(uint256 marketId, uint256 outcomeId) public pure returns (uint256 tokenId) {
        // Use upper bits for marketId, lower 8 bits for outcomeId
        // This allows up to 256 outcomes per market
        tokenId = (marketId << 8) | outcomeId;
    }

    /**
     * @notice Decodes a token ID into market ID and outcome ID
     * @param tokenId The encoded token ID
     * @return marketId The market identifier
     * @return outcomeId The outcome identifier
     */
    function decodeTokenId(uint256 tokenId) public pure returns (uint256 marketId, uint256 outcomeId) {
        outcomeId = tokenId & 0xFF; // Lower 8 bits
        marketId = tokenId >> 8; // Upper bits
    }

    /**
     * @notice Checks if a market has been resolved
     * @param marketId The market identifier
     * @return True if the market is resolved
     */
    function isResolved(uint256 marketId) public view returns (bool) {
        return winningOutcome[marketId] != UNRESOLVED;
    }

    /**
     * @notice Gets the balance of a specific outcome for a user
     * @param user The user address
     * @param marketId The market identifier
     * @param outcomeId The outcome identifier
     * @return The balance of outcome tokens
     */
    function balanceOfOutcome(address user, uint256 marketId, uint256 outcomeId) public view returns (uint256) {
        uint256 tokenId = encodeTokenId(marketId, outcomeId);
        return balanceOf(user, tokenId);
    }

    /**
     * @notice Gets the total supply of a specific outcome token
     * @dev Implemented by tracking via events in subgraph, or can be added with storage
     * @param marketId The market identifier
     * @param outcomeId The outcome identifier
     * @return Always returns 0 (not tracked on-chain to save gas)
     */
    function totalSupply(uint256 marketId, uint256 outcomeId) public pure returns (uint256) {
        // Not tracking supply on-chain to save gas
        // Use events and subgraph for this
        return 0;
    }
}
