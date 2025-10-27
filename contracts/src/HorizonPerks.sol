// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title HorizonPerks
 * @notice Provides fee tier benefits based on HORIZON token holdings
 * @dev View-only contract that calculates trading fees and protocol/creator splits
 *      Higher HORIZON balance = more fees go to market creators (incentivizes creator marketing)
 *
 *      Default Tiers (All charge 2% user fee, but split changes):
 *      Tier 0: 0 HORIZON      → 2.0% fee, 10% protocol / 90% creator
 *      Tier 1: 10K HORIZON    → 2.0% fee, 8% protocol / 92% creator
 *      Tier 2: 50K HORIZON    → 2.0% fee, 6% protocol / 94% creator
 *      Tier 3: 100K HORIZON   → 2.0% fee, 4% protocol / 96% creator
 *      Tier 4: 500K HORIZON   → 2.0% fee, 2% protocol / 98% creator
 */
contract HorizonPerks is Ownable {
    // ============ Events ============

    event TierUpdated(uint256 indexed tierId, uint256 minBalance, uint16 feeBps, uint16 protocolBps);
    event TierAdded(uint256 indexed tierId, uint256 minBalance, uint16 feeBps, uint16 protocolBps);
    event TierRemoved(uint256 indexed tierId);

    // ============ Structs ============

    /**
     * @notice Fee tier configuration
     * @param minBalance Minimum HORIZON balance required for this tier
     * @param feeBps Trading fee in basis points (e.g., 200 = 2%)
     * @param protocolBps Protocol's share of the fee in basis points (e.g., 1000 = 10%)
     */
    struct FeeTier {
        uint256 minBalance;
        uint16 feeBps;
        uint16 protocolBps;
    }

    // ============ Constants ============

    /// @notice Default trading fee (no HORIZON holdings)
    uint16 public constant DEFAULT_FEE_BPS = 200; // 2%
    
    /// @notice Default protocol share (10% of fees)
    uint16 public constant DEFAULT_PROTOCOL_BPS = 1000; // 10%

    // ============ State Variables ============

    /// @notice HORIZON token contract
    IERC20 public immutable horizonToken;

    /// @notice Array of fee tiers (sorted by minBalance ascending)
    FeeTier[] public feeTiers;

    // ============ Constructor ============

    /**
     * @notice Initializes HorizonPerks with default tier structure
     * @param _horizonToken Address of the HORIZON token contract
     */
    constructor(address _horizonToken) Ownable(msg.sender) {
        require(_horizonToken != address(0), "Invalid HORIZON token");
        horizonToken = IERC20(_horizonToken);

        // Initialize default tiers (sorted by minBalance ascending)
        // All charge 2% user fee, but protocol share decreases with higher HORIZON holdings
        feeTiers.push(FeeTier({minBalance: 0, feeBps: 200, protocolBps: 1000})); // Tier 0: 10% protocol / 90% creator
        feeTiers.push(FeeTier({minBalance: 10_000 * 10 ** 18, feeBps: 200, protocolBps: 800})); // Tier 1: 8% protocol / 92% creator
        feeTiers.push(FeeTier({minBalance: 50_000 * 10 ** 18, feeBps: 200, protocolBps: 600})); // Tier 2: 6% protocol / 94% creator
        feeTiers.push(FeeTier({minBalance: 100_000 * 10 ** 18, feeBps: 200, protocolBps: 400})); // Tier 3: 4% protocol / 96% creator
        feeTiers.push(FeeTier({minBalance: 500_000 * 10 ** 18, feeBps: 200, protocolBps: 200})); // Tier 4: 2% protocol / 98% creator
    }

    // ============ Admin Functions ============

    /**
     * @notice Updates an existing fee tier
     * @param tierId Index of the tier to update
     * @param minBalance New minimum balance
     * @param feeBps New fee in basis points
     * @param protocolBps New protocol share in basis points
     */
    function updateTier(uint256 tierId, uint256 minBalance, uint16 feeBps, uint16 protocolBps) external onlyOwner {
        require(tierId < feeTiers.length, "Invalid tier ID");
        require(feeBps > 0 && feeBps <= 1000, "Invalid fee bps"); // Max 10%
        require(protocolBps <= 10000, "Invalid protocol bps"); // Max 100%

        // Validate ordering: each tier should have higher minBalance than previous
        if (tierId > 0) {
            require(minBalance > feeTiers[tierId - 1].minBalance, "Must be higher than previous tier");
        }
        if (tierId < feeTiers.length - 1) {
            require(minBalance < feeTiers[tierId + 1].minBalance, "Must be lower than next tier");
        }

        feeTiers[tierId] = FeeTier({minBalance: minBalance, feeBps: feeBps, protocolBps: protocolBps});

        emit TierUpdated(tierId, minBalance, feeBps, protocolBps);
    }

    /**
     * @notice Adds a new fee tier
     * @param minBalance Minimum balance for this tier
     * @param feeBps Fee in basis points
     * @param protocolBps Protocol share in basis points
     */
    function addTier(uint256 minBalance, uint16 feeBps, uint16 protocolBps) external onlyOwner {
        require(feeBps > 0 && feeBps <= 1000, "Invalid fee bps");
        require(protocolBps <= 10000, "Invalid protocol bps");

        // New tier should have higher minBalance than last tier
        if (feeTiers.length > 0) {
            require(minBalance > feeTiers[feeTiers.length - 1].minBalance, "Must be higher than last tier");
        }

        feeTiers.push(FeeTier({minBalance: minBalance, feeBps: feeBps, protocolBps: protocolBps}));

        emit TierAdded(feeTiers.length - 1, minBalance, feeBps, protocolBps);
    }

    /**
     * @notice Removes the last fee tier
     * @dev Can only remove from the end to maintain ordering
     */
    function removeLastTier() external onlyOwner {
        require(feeTiers.length > 1, "Cannot remove all tiers");

        uint256 removedTierId = feeTiers.length - 1;
        feeTiers.pop();

        emit TierRemoved(removedTierId);
    }

    // ============ View Functions ============

    /**
     * @notice Gets the trading fee for a user based on their HORIZON balance
     * @param user Address of the user
     * @return feeBps Fee in basis points
     */
    function feeBpsFor(address user) external view returns (uint16 feeBps) {
        uint256 balance = horizonToken.balanceOf(user);
        return feeBpsForBalance(balance);
    }

    /**
     * @notice Gets the protocol share for a user based on their HORIZON balance
     * @param user Address of the user
     * @return protocolBps Protocol share in basis points (e.g., 1000 = 10%)
     */
    function protocolBpsFor(address user) external view returns (uint16 protocolBps) {
        uint256 balance = horizonToken.balanceOf(user);
        return protocolBpsForBalance(balance);
    }

    /**
     * @notice Gets both fee and protocol share for a user
     * @param user Address of the user
     * @return feeBps Fee in basis points
     * @return protocolBps Protocol share in basis points
     */
    function getTierInfoFor(address user) external view returns (uint16 feeBps, uint16 protocolBps) {
        uint256 balance = horizonToken.balanceOf(user);
        FeeTier memory tier = getTierForBalance(balance);
        return (tier.feeBps, tier.protocolBps);
    }

    /**
     * @notice Gets the trading fee for a given HORIZON balance
     * @param balance HORIZON token balance
     * @return feeBps Fee in basis points
     */
    function feeBpsForBalance(uint256 balance) public view returns (uint16 feeBps) {
        FeeTier memory tier = getTierForBalance(balance);
        return tier.feeBps;
    }

    /**
     * @notice Gets the protocol share for a given HORIZON balance
     * @param balance HORIZON token balance
     * @return protocolBps Protocol share in basis points
     */
    function protocolBpsForBalance(uint256 balance) public view returns (uint16 protocolBps) {
        FeeTier memory tier = getTierForBalance(balance);
        return tier.protocolBps;
    }

    /**
     * @notice Gets the fee tier for a given balance
     * @param balance HORIZON token balance
     * @return tier The applicable FeeTier struct
     */
    function getTierForBalance(uint256 balance) public view returns (FeeTier memory tier) {
        // Start from the end (highest tier) and work backwards
        for (uint256 i = feeTiers.length; i > 0; i--) {
            uint256 index = i - 1;
            if (balance >= feeTiers[index].minBalance) {
                return feeTiers[index];
            }
        }

        // Fallback to default (should never reach here if tier 0 has minBalance = 0)
        return FeeTier({minBalance: 0, feeBps: DEFAULT_FEE_BPS, protocolBps: DEFAULT_PROTOCOL_BPS});
    }

    /**
     * @notice Gets the tier level for a user
     * @param user Address of the user
     * @return tierLevel The tier level (0 = lowest)
     */
    function tierLevelFor(address user) external view returns (uint256 tierLevel) {
        uint256 balance = horizonToken.balanceOf(user);
        return tierLevelForBalance(balance);
    }

    /**
     * @notice Gets the tier level for a given balance
     * @param balance HORIZON token balance
     * @return tierLevel The tier level (0 = lowest)
     */
    function tierLevelForBalance(uint256 balance) public view returns (uint256 tierLevel) {
        // Start from the end (highest tier) and work backwards
        for (uint256 i = feeTiers.length; i > 0; i--) {
            uint256 index = i - 1;
            if (balance >= feeTiers[index].minBalance) {
                return index;
            }
        }
        return 0; // Default to tier 0
    }

    /**
     * @notice Gets all fee tiers
     * @return tiers Array of all fee tiers
     */
    function getAllTiers() external view returns (FeeTier[] memory tiers) {
        uint256 length = feeTiers.length;
        tiers = new FeeTier[](length);
        for (uint256 i = 0; i < length; i++) {
            tiers[i] = feeTiers[i];
        }
    }

    /**
     * @notice Gets a specific fee tier
     * @param tierId Index of the tier
     * @return tier The fee tier struct
     */
    function getTier(uint256 tierId) external view returns (FeeTier memory tier) {
        require(tierId < feeTiers.length, "Invalid tier ID");
        return feeTiers[tierId];
    }

    /**
     * @notice Gets the number of tiers
     * @return count Number of fee tiers
     */
    function getTierCount() external view returns (uint256 count) {
        return feeTiers.length;
    }

    /**
     * @notice Gets the next tier info for a user
     * @param user Address of the user
     * @return hasNextTier Whether there is a next tier
     * @return nextTierMinBalance Minimum balance needed for next tier
     * @return nextTierFeeBps Fee bps of next tier
     * @return nextTierProtocolBps Protocol bps of next tier
     * @return tokensNeeded Tokens needed to reach next tier
     */
    function getNextTierInfo(address user)
        external
        view
        returns (
            bool hasNextTier,
            uint256 nextTierMinBalance,
            uint16 nextTierFeeBps,
            uint16 nextTierProtocolBps,
            uint256 tokensNeeded
        )
    {
        uint256 balance = horizonToken.balanceOf(user);
        uint256 currentTier = tierLevelForBalance(balance);

        if (currentTier < feeTiers.length - 1) {
            hasNextTier = true;
            uint256 nextTier = currentTier + 1;
            nextTierMinBalance = feeTiers[nextTier].minBalance;
            nextTierFeeBps = feeTiers[nextTier].feeBps;
            nextTierProtocolBps = feeTiers[nextTier].protocolBps;
            tokensNeeded = nextTierMinBalance > balance ? nextTierMinBalance - balance : 0;
        } else {
            hasNextTier = false;
            nextTierMinBalance = 0;
            nextTierFeeBps = 0;
            nextTierProtocolBps = 0;
            tokensNeeded = 0;
        }
    }

    /**
     * @notice Calculates fee amount for a given input amount and user
     * @param user Address of the user
     * @param amount Input amount
     * @return feeAmount The calculated fee amount
     */
    function calculateFee(address user, uint256 amount) external view returns (uint256 feeAmount) {
        uint16 feeBps = this.feeBpsFor(user);
        feeAmount = (amount * feeBps) / 10000;
    }

    /**
     * @notice Calculates the amount after fees for a given input and user
     * @param user Address of the user
     * @param amount Input amount
     * @return amountAfterFee Amount remaining after fee deduction
     */
    function calculateAmountAfterFee(address user, uint256 amount) external view returns (uint256 amountAfterFee) {
        uint16 feeBps = this.feeBpsFor(user);
        uint256 feeAmount = (amount * feeBps) / 10000;
        amountAfterFee = amount - feeAmount;
    }
}
