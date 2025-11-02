// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./markets/BinaryMarket.sol";
import "./markets/MultiChoiceMarket.sol";
import "./markets/LimitOrderMarket.sol";
import "./markets/PooledLiquidityMarket.sol";
import "./OutcomeToken.sol";
import "./FeeSplitter.sol";
import "./HorizonPerks.sol";

/**
 * @title MarketFactory
 * @notice Factory contract for creating and managing prediction markets
 * @dev Orchestrates deployment of MarketAMM instances and registration across all system contracts
 *
     *      Market Lifecycle:
     *      1. Creator calls createMarket() with parameters + HORIZON stake
 *      2. Factory deploys new MarketAMM instance
 *      3. Factory registers market in OutcomeToken and FeeSplitter
 *      4. Market is Active (can trade)
 *      5. Market closes at closeTime (no more trading)
 *      6. Market is resolved via ResolutionModule
 *      7. Winners claim via OutcomeToken
 *      8. If resolution is fair, creator stake is refunded
 */
contract MarketFactory is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ Enums ============

    enum MarketStatus {
        Active,     // Market is open for trading
        Closed,     // Past closeTime, awaiting resolution
        Resolved,   // Resolution finalized
        Invalid     // Market was invalidated
    }

    enum MarketType {
        Binary,            // 2 outcomes, MarketAMM (constant product)
        MultiChoice,       // 3-8 outcomes, LMSR pricing
        LimitOrder,        // 2+ outcomes, order book
        PooledLiquidity    // 2 outcomes, Uniswap V3-style
    }

    // ============ Errors ============

    error InvalidCloseTime();
    error InvalidCollateral();
    error InvalidCreatorStake();
    error MarketDoesNotExist();
    error NotMarketCreator();
    error MarketNotResolved();
    error StakeAlreadyClaimed();
    error InvalidAddress();
    error InvalidCategory();
    error InvalidMarketType();
    error InvalidOutcomeCount();
    error InvalidLiquidityParameter();

    // ============ Events ============

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
    event MarketStatusUpdated(uint256 indexed marketId, MarketStatus oldStatus, MarketStatus newStatus);
    event MinCreatorStakeUpdated(uint256 oldStake, uint256 newStake);

    // ============ Structs ============

    /**
     * @notice Parameters for creating a new market
     * @param marketType Type of market (Binary, MultiChoice, LimitOrder, PooledLiquidity)
     * @param collateralToken ERC20 token used for trading (e.g., USDC)
     * @param closeTime Timestamp when market closes for trading
     * @param category Market category (e.g., "politics", "sports", "crypto")
     * @param metadataURI IPFS CID containing market question and details
     * @param creatorStake Amount of HORIZON tokens to stake (must be >= minCreatorStake)
     * @param outcomeCount Number of outcomes (2 for Binary/PooledLiquidity, 3-8 for MultiChoice, 2+ for LimitOrder)
     * @param liquidityParameter LMSR liquidity parameter for MultiChoice markets (ignored for other types)
     */
    struct MarketParams {
        uint8 marketType;
        address collateralToken;
        uint256 closeTime;
        string category;
        string metadataURI;
        uint256 creatorStake;
        uint8 outcomeCount;
        uint256 liquidityParameter;
    }

    /**
     * @notice Market information stored on-chain
     * @param id Unique market identifier
     * @param creator Address that created the market
     * @param marketType Type of market contract deployed
     * @param amm MarketAMM contract address
     * @param collateralToken ERC20 token for trading
     * @param closeTime When market closes
     * @param category Market category
     * @param metadataURI IPFS CID
     * @param creatorStake HORIZON tokens staked
     * @param outcomeCount Number of outcomes in this market
     * @param stakeRefunded Whether creator stake was refunded
     * @param status Current market status
     */
    struct Market {
        uint256 id;
        address creator;
        uint8 marketType;
        address amm;
        address collateralToken;
        uint256 closeTime;
        string category;
        string metadataURI;
        uint256 creatorStake;
        uint8 outcomeCount;
        bool stakeRefunded;
        MarketStatus status;
    }

    // ============ State Variables ============

    /// @notice Core system contracts
    OutcomeToken public immutable outcomeToken;
    FeeSplitter public immutable feeSplitter;
    HorizonPerks public immutable horizonPerks;
    IERC20 public immutable horizonToken;

    /// @notice Counter for generating unique market IDs
    uint256 public nextMarketId = 1;

    /// @notice Minimum HORIZON stake required to create market
    uint256 public minCreatorStake = 100 ether;

    /// @notice Mapping of market ID to Market struct
    mapping(uint256 => Market) public markets;

    /// @notice Array of all market IDs for enumeration
    uint256[] public allMarketIds;

    /// @notice Mapping of category to array of market IDs
    mapping(string => uint256[]) public marketsByCategory;

    /// @notice Mapping of creator address to array of their market IDs
    mapping(address => uint256[]) public marketsByCreator;

    // ============ Constructor ============

    /**
     * @notice Initializes the MarketFactory
     * @param _outcomeToken OutcomeToken contract address
     * @param _feeSplitter FeeSplitter contract address
     * @param _horizonPerks HorizonPerks contract address
     * @param _horizonToken HORIZON token contract address
     */
    constructor(address _outcomeToken, address _feeSplitter, address _horizonPerks, address _horizonToken)
        Ownable(msg.sender)
    {
        if (_outcomeToken == address(0) || _feeSplitter == address(0) || _horizonPerks == address(0) || _horizonToken == address(0)) {
            revert InvalidAddress();
        }

        outcomeToken = OutcomeToken(_outcomeToken);
        feeSplitter = FeeSplitter(_feeSplitter);
        horizonPerks = HorizonPerks(_horizonPerks);
        horizonToken = IERC20(_horizonToken);
    }

    // ============ Market Creation ============

    /**
     * @notice Internal function to deploy the appropriate market contract based on type
     * @param marketId Market identifier
     * @param params Market parameters
     * @return ammAddress Address of the deployed market contract
     */
    function _deployMarket(uint256 marketId, MarketParams calldata params) internal returns (address ammAddress) {
        if (params.marketType == uint8(MarketType.Binary)) {
            // Binary market using BinaryMarket (2 outcomes, static 1:1 pricing with 2% fee)
            if (params.outcomeCount != 2) revert InvalidOutcomeCount();
            
            BinaryMarket market = new BinaryMarket(
                marketId,
                params.collateralToken,
                address(outcomeToken),
                address(feeSplitter),
                address(horizonPerks),
                params.closeTime
            );
            return address(market);
            
        } else if (params.marketType == uint8(MarketType.MultiChoice)) {
            // Multi-choice market using LMSR (3-8 outcomes)
            if (params.outcomeCount < 3 || params.outcomeCount > 8) revert InvalidOutcomeCount();
            if (params.liquidityParameter == 0) revert InvalidLiquidityParameter();
            
            MultiChoiceMarket market = new MultiChoiceMarket(
                marketId,
                params.collateralToken,
                address(outcomeToken),
                address(feeSplitter),
                address(horizonPerks),
                params.closeTime,
                params.outcomeCount,
                params.liquidityParameter
            );
            return address(market);
            
        } else if (params.marketType == uint8(MarketType.LimitOrder)) {
            // Limit order market (2+ outcomes, order book)
            if (params.outcomeCount < 2) revert InvalidOutcomeCount();
            
            LimitOrderMarket market = new LimitOrderMarket(
                marketId,
                params.collateralToken,
                address(outcomeToken),
                address(feeSplitter),
                address(horizonPerks),
                params.closeTime,
                params.outcomeCount
            );
            return address(market);
            
        } else if (params.marketType == uint8(MarketType.PooledLiquidity)) {
            // Pooled liquidity market (binary only, Uniswap V3-style)
            if (params.outcomeCount != 2) revert InvalidOutcomeCount();
            
            PooledLiquidityMarket market = new PooledLiquidityMarket(
                marketId,
                params.collateralToken,
                address(outcomeToken),
                address(feeSplitter),
                address(horizonPerks),
                params.closeTime,
                "Horizon LP Token",
                "HZN-LP"
            );
            return address(market);
            
        } else {
            revert InvalidMarketType();
        }
    }

    /**
     * @notice Creates a new prediction market
     * @param params Market parameters
     * @return marketId The ID of the newly created market
     */
    function createMarket(MarketParams calldata params) external nonReentrant returns (uint256 marketId) {
        // Validate parameters
        if (params.collateralToken == address(0)) revert InvalidCollateral();
        if (params.closeTime <= block.timestamp) revert InvalidCloseTime();
        if (params.creatorStake < minCreatorStake) revert InvalidCreatorStake();
        if (bytes(params.category).length == 0) revert InvalidCategory();
        if (params.marketType > uint8(MarketType.PooledLiquidity)) revert InvalidMarketType();

        // Generate market ID
        marketId = nextMarketId++;

        // Transfer creator stake
        IERC20(address(horizonToken)).safeTransferFrom(msg.sender, address(this), params.creatorStake);

        // Register market in OutcomeToken
        outcomeToken.registerMarket(marketId, IERC20(params.collateralToken));

        // Register market in FeeSplitter
        feeSplitter.registerMarket(marketId, msg.sender);

        // Deploy appropriate market contract based on type
        address ammAddress = _deployMarket(marketId, params);

        // Authorize AMM to mint/burn outcome tokens
        outcomeToken.setAMMAuthorization(ammAddress, true);

        // Store market info
        markets[marketId] = Market({
            id: marketId,
            creator: msg.sender,
            marketType: params.marketType,
            amm: ammAddress,
            collateralToken: params.collateralToken,
            closeTime: params.closeTime,
            category: params.category,
            metadataURI: params.metadataURI,
            creatorStake: params.creatorStake,
            outcomeCount: params.outcomeCount,
            stakeRefunded: false,
            status: MarketStatus.Active
        });

        // Add to registries
        allMarketIds.push(marketId);
        marketsByCategory[params.category].push(marketId);
        marketsByCreator[msg.sender].push(marketId);

        emit MarketCreated(
            marketId,
            msg.sender,
            ammAddress,
            params.collateralToken,
            params.closeTime,
            params.category,
            params.metadataURI,
            params.creatorStake
        );
    }

    // ============ Creator Stake Management ============

    /**
     * @notice Refunds creator stake after successful market resolution
     * @dev Can be called by creator or anyone after market is resolved
     * @param marketId Market identifier
     */
    function refundCreatorStake(uint256 marketId) external nonReentrant {
        Market storage market = markets[marketId];
        if (market.id == 0) revert MarketDoesNotExist();
        if (market.stakeRefunded) revert StakeAlreadyClaimed();
        if (!outcomeToken.isResolved(marketId)) revert MarketNotResolved();

        // Mark as refunded before transfer (CEI pattern)
        market.stakeRefunded = true;

        // Transfer stake back to creator
        IERC20(address(horizonToken)).safeTransfer(market.creator, market.creatorStake);

        emit CreatorStakeRefunded(marketId, market.creator, market.creatorStake);
    }

    // ============ Market Status Management ============

    /**
     * @notice Updates market status based on current state
     * @dev Anyone can call to update status (no privileges required)
     * @param marketId Market identifier
     */
    function updateMarketStatus(uint256 marketId) external {
        Market storage market = markets[marketId];
        if (market.id == 0) revert MarketDoesNotExist();

        MarketStatus oldStatus = market.status;
        MarketStatus newStatus = _computeMarketStatus(market);

        if (newStatus != oldStatus) {
            market.status = newStatus;
            emit MarketStatusUpdated(marketId, oldStatus, newStatus);
        }
    }

    /**
     * @notice Computes the current status of a market
     * @param market Market struct
     * @return Current market status
     */
    function _computeMarketStatus(Market memory market) internal view returns (MarketStatus) {
        // Check if resolved
        if (outcomeToken.isResolved(market.id)) {
            return MarketStatus.Resolved;
        }

        // Check if closed
        if (block.timestamp > market.closeTime) {
            return MarketStatus.Closed;
        }

        // Otherwise active
        return MarketStatus.Active;
    }

    // ============ Admin Functions ============

    /**
     * @notice Updates minimum creator stake requirement
     * @param newMinStake New minimum stake in HORIZON tokens
     */
    function setMinCreatorStake(uint256 newMinStake) external onlyOwner {
        uint256 oldStake = minCreatorStake;
        minCreatorStake = newMinStake;
        emit MinCreatorStakeUpdated(oldStake, newMinStake);
    }

    // ============ View Functions ============

    /**
     * @notice Gets the total number of markets
     * @return Total market count
     */
    function getMarketCount() external view returns (uint256) {
        return allMarketIds.length;
    }

    /**
     * @notice Gets market details by ID
     * @param marketId Market identifier
     * @return Market struct with all details
     */
    function getMarket(uint256 marketId) external view returns (Market memory) {
        Market memory market = markets[marketId];
        if (market.id == 0) revert MarketDoesNotExist();

        // Update status in the returned copy
        market.status = _computeMarketStatus(market);
        return market;
    }

    /**
     * @notice Gets all market IDs
     * @return Array of all market IDs
     */
    function getAllMarketIds() external view returns (uint256[] memory) {
        return allMarketIds;
    }

    /**
     * @notice Gets market IDs by category
     * @param category Category string
     * @return Array of market IDs in that category
     */
    function getMarketIdsByCategory(string calldata category) external view returns (uint256[] memory) {
        return marketsByCategory[category];
    }

    /**
     * @notice Gets market IDs created by a specific address
     * @param creator Creator address
     * @return Array of market IDs created by that address
     */
    function getMarketIdsByCreator(address creator) external view returns (uint256[] memory) {
        return marketsByCreator[creator];
    }

    /**
     * @notice Gets multiple markets with pagination
     * @param offset Starting index
     * @param limit Maximum number of markets to return
     * @return markets Array of Market structs
     */
    function getMarkets(uint256 offset, uint256 limit) external view returns (Market[] memory) {
        uint256 total = allMarketIds.length;
        if (offset >= total) return new Market[](0);

        uint256 end = offset + limit;
        if (end > total) end = total;

        uint256 count = end - offset;
        Market[] memory result = new Market[](count);

        for (uint256 i = 0; i < count; i++) {
            uint256 marketId = allMarketIds[offset + i];
            result[i] = markets[marketId];
            // Update status in the returned copy
            result[i].status = _computeMarketStatus(result[i]);
        }

        return result;
    }

    /**
     * @notice Gets active markets with pagination
     * @param offset Starting index in active markets array
     * @param limit Maximum number to return
     * @return markets Array of active Market structs
     */
    function getActiveMarkets(uint256 offset, uint256 limit) external view returns (Market[] memory) {
        // First pass: count active markets
        uint256 activeCount = 0;
        uint256 marketCount = allMarketIds.length;
        for (uint256 i = 0; i < marketCount; i++) {
            Market memory market = markets[allMarketIds[i]];
            if (_computeMarketStatus(market) == MarketStatus.Active) {
                activeCount++;
            }
        }

        if (offset >= activeCount) return new Market[](0);

        uint256 end = offset + limit;
        if (end > activeCount) end = activeCount;
        uint256 count = end - offset;

        Market[] memory result = new Market[](count);
        uint256 resultIndex = 0;
        uint256 activeIndex = 0;

        // Second pass: collect active markets
        for (uint256 i = 0; i < allMarketIds.length && resultIndex < count; i++) {
            Market memory market = markets[allMarketIds[i]];
            if (_computeMarketStatus(market) == MarketStatus.Active) {
                if (activeIndex >= offset) {
                    result[resultIndex] = market;
                    result[resultIndex].status = MarketStatus.Active;
                    resultIndex++;
                }
                activeIndex++;
            }
        }

        return result;
    }

    /**
     * @notice Checks if a market exists
     * @param marketId Market identifier
     * @return True if market exists
     */
    function marketExists(uint256 marketId) external view returns (bool) {
        return markets[marketId].id != 0;
    }
}
