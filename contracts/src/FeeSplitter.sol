// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title FeeSplitter
 * @notice Splits trading fees between protocol and market creators
 * @dev Fee distribution:
 *      - Protocol share varies based on trader's HORIZON holdings (2-10%)
 *      - Creator receives the remainder (90-98%)
 *      - Incentivizes creators to attract HORIZON whale traders
 *
 *      Uses pull-payment pattern for security
 */
contract FeeSplitter is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error InvalidFeeConfig();
    error NoFeesToClaim();
    error MarketNotRegistered();

    // ============ Events ============

    event FeeDistributed(
        uint256 indexed marketId,
        address indexed token,
        uint256 totalAmount,
        uint256 protocolAmount,
        uint256 creatorAmount,
        uint16 protocolBps
    );
    event FeeClaimed(uint256 indexed marketId, address indexed creator, address indexed token, uint256 amount);
    event ProtocolFeeClaimed(address indexed token, uint256 amount);
    event MarketRegistered(uint256 indexed marketId, address indexed creator);
    event FeeConfigUpdated(uint256 indexed marketId, uint16 protocolBps, uint16 creatorBps);

    // ============ Structs ============

    /**
     * @notice Fee configuration for a market
     * @param protocolBps Basis points allocated to protocol (1000 = 10%)
     * @param creatorBps Basis points allocated to creator (9000 = 90%)
     */
    struct FeeConfig {
        uint16 protocolBps; // 1000 = 10%
        uint16 creatorBps;  // 9000 = 90%
    }

    // ============ Constants ============

    /// @notice 100% in basis points
    uint256 public constant BPS_DENOMINATOR = 10000;

    /// @notice Default protocol fee (10%)
    uint16 public constant DEFAULT_PROTOCOL_BPS = 1000;

    /// @notice Default creator fee (90%)
    uint16 public constant DEFAULT_CREATOR_BPS = 9000;

    // ============ State Variables ============

    /// @notice Mapping of marketId to creator address
    mapping(uint256 => address) public marketCreator;

    /// @notice Mapping of marketId to fee configuration
    mapping(uint256 => FeeConfig) public feeConfigs;

    /// @notice Pending fees for creators: marketId => token => amount
    mapping(uint256 => mapping(address => uint256)) public creatorPendingFees;

    /// @notice Pending fees for protocol: token => amount
    mapping(address => uint256) public protocolPendingFees;

    /// @notice Protocol treasury address
    address public protocolTreasury;

    // ============ Constructor ============

    /**
     * @notice Initializes the FeeSplitter
     * @param _protocolTreasury Address of the protocol treasury
     */
    constructor(address _protocolTreasury) Ownable(msg.sender) {
        require(_protocolTreasury != address(0), "Invalid treasury");
        protocolTreasury = _protocolTreasury;
    }

    // ============ Admin Functions ============

    /**
     * @notice Registers a new market with its creator
     * @param marketId The market identifier
     * @param creator The market creator address
     */
    function registerMarket(uint256 marketId, address creator) external onlyOwner {
        require(creator != address(0), "Invalid creator");
        require(marketCreator[marketId] == address(0), "Market already registered");

        marketCreator[marketId] = creator;
        feeConfigs[marketId] = FeeConfig({
            protocolBps: DEFAULT_PROTOCOL_BPS,
            creatorBps: DEFAULT_CREATOR_BPS
        });

        emit MarketRegistered(marketId, creator);
    }

    /**
     * @notice Updates the fee configuration for a market
     * @param marketId The market identifier
     * @param protocolBps New protocol basis points
     * @param creatorBps New creator basis points
     */
    function updateFeeConfig(uint256 marketId, uint16 protocolBps, uint16 creatorBps) external onlyOwner {
        if (marketCreator[marketId] == address(0)) revert MarketNotRegistered();
        if (protocolBps + creatorBps != BPS_DENOMINATOR) revert InvalidFeeConfig();

        feeConfigs[marketId] = FeeConfig({
            protocolBps: protocolBps,
            creatorBps: creatorBps
        });

        emit FeeConfigUpdated(marketId, protocolBps, creatorBps);
    }

    /**
     * @notice Updates the protocol treasury address
     * @param newTreasury New treasury address
     */
    function setProtocolTreasury(address newTreasury) external onlyOwner {
        require(newTreasury != address(0), "Invalid treasury");
        protocolTreasury = newTreasury;
    }

    // ============ Fee Distribution Functions ============

    /**
     * @notice Distributes fees for a market between protocol and creator
     * @dev Called by the AMM contract when fees are collected
     * @param marketId The market identifier
     * @param token The token address (collateral token)
     * @param amount Total fee amount to distribute
     * @param protocolBps Protocol's share in basis points (e.g., 1000 = 10%)
     */
    function distribute(uint256 marketId, address token, uint256 amount, uint16 protocolBps) external {
        if (marketCreator[marketId] == address(0)) revert MarketNotRegistered();
        if (amount == 0) return;
        if (protocolBps > 10000) revert InvalidFeeConfig();

        // Calculate splits
        uint256 protocolAmount = (amount * protocolBps) / BPS_DENOMINATOR;
        uint256 creatorAmount = amount - protocolAmount; // Give any rounding dust to creator

        // Update pending balances
        protocolPendingFees[token] += protocolAmount;
        creatorPendingFees[marketId][token] += creatorAmount;

        // Transfer tokens from sender (AMM) to this contract
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

        emit FeeDistributed(marketId, token, amount, protocolAmount, creatorAmount, protocolBps);
    }

    // ============ Claim Functions ============

    /**
     * @notice Allows a creator to claim their pending fees for a market
     * @param marketId The market identifier
     * @param token The token address
     */
    function claimCreatorFees(uint256 marketId, address token) external nonReentrant {
        address creator = marketCreator[marketId];
        require(msg.sender == creator, "Not the creator");

        uint256 amount = creatorPendingFees[marketId][token];
        if (amount == 0) revert NoFeesToClaim();

        // Reset pending fees before transfer (CEI pattern)
        creatorPendingFees[marketId][token] = 0;

        // Transfer fees to creator
        IERC20(token).safeTransfer(creator, amount);

        emit FeeClaimed(marketId, creator, token, amount);
    }

    /**
     * @notice Allows a creator to claim fees from multiple markets
     * @param marketIds Array of market identifiers
     * @param tokens Array of token addresses (must match marketIds length)
     */
    function claimCreatorFeesMultiple(uint256[] calldata marketIds, address[] calldata tokens)
        external
        nonReentrant
    {
        require(marketIds.length == tokens.length, "Array length mismatch");

        for (uint256 i = 0; i < marketIds.length; i++) {
            uint256 marketId = marketIds[i];
            address token = tokens[i];
            address creator = marketCreator[marketId];

            if (msg.sender != creator) continue; // Skip if not creator
            if (creatorPendingFees[marketId][token] == 0) continue; // Skip if no fees

            uint256 amount = creatorPendingFees[marketId][token];
            creatorPendingFees[marketId][token] = 0;

            IERC20(token).safeTransfer(creator, amount);
            emit FeeClaimed(marketId, creator, token, amount);
        }
    }

    /**
     * @notice Allows protocol to claim their pending fees
     * @param token The token address
     */
    function claimProtocolFees(address token) external nonReentrant {
        require(msg.sender == protocolTreasury || msg.sender == owner(), "Not authorized");

        uint256 amount = protocolPendingFees[token];
        if (amount == 0) revert NoFeesToClaim();

        // Reset pending fees before transfer (CEI pattern)
        protocolPendingFees[token] = 0;

        // Transfer fees to protocol treasury
        IERC20(token).safeTransfer(protocolTreasury, amount);

        emit ProtocolFeeClaimed(token, amount);
    }

    /**
     * @notice Allows protocol to claim fees from multiple tokens
     * @param tokens Array of token addresses
     */
    function claimProtocolFeesMultiple(address[] calldata tokens) external nonReentrant {
        require(msg.sender == protocolTreasury || msg.sender == owner(), "Not authorized");

        for (uint256 i = 0; i < tokens.length; i++) {
            address token = tokens[i];
            uint256 amount = protocolPendingFees[token];

            if (amount == 0) continue; // Skip if no fees

            protocolPendingFees[token] = 0;
            IERC20(token).safeTransfer(protocolTreasury, amount);
            emit ProtocolFeeClaimed(token, amount);
        }
    }

    // ============ View Functions ============

    /**
     * @notice Gets the pending fees for a creator for a specific market and token
     * @param marketId The market identifier
     * @param token The token address
     * @return The pending fee amount
     */
    function getCreatorPendingFees(uint256 marketId, address token) external view returns (uint256) {
        return creatorPendingFees[marketId][token];
    }

    /**
     * @notice Gets the pending protocol fees for a specific token
     * @param token The token address
     * @return The pending fee amount
     */
    function getProtocolPendingFees(address token) external view returns (uint256) {
        return protocolPendingFees[token];
    }

    /**
     * @notice Gets the fee configuration for a market
     * @param marketId The market identifier
     * @return protocolBps The protocol basis points
     * @return creatorBps The creator basis points
     */
    function getFeeConfig(uint256 marketId) external view returns (uint16 protocolBps, uint16 creatorBps) {
        FeeConfig memory config = feeConfigs[marketId];
        return (config.protocolBps, config.creatorBps);
    }

    /**
     * @notice Previews the fee split for a given amount and protocol share
     * @param amount The total fee amount
     * @param protocolBps The protocol's share in basis points
     * @return protocolAmount Amount that goes to protocol
     * @return creatorAmount Amount that goes to creator
     */
    function previewSplit(uint256 amount, uint16 protocolBps)
        external
        pure
        returns (uint256 protocolAmount, uint256 creatorAmount)
    {
        protocolAmount = (amount * protocolBps) / BPS_DENOMINATOR;
        creatorAmount = amount - protocolAmount;
    }
}
