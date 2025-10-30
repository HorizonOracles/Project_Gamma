// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/HorizonPerks.sol";
import "../mocks/MockERC20.sol";

contract HorizonPerksTest is Test {
    HorizonPerks public perks;
    MockERC20 public horizonToken;

    address public owner = address(this);
    address public user1 = address(0x1);
    address public user2 = address(0x2);
    address public user3 = address(0x3);
    address public unauthorized = address(0x4);

    event TierUpdated(uint256 indexed tierId, uint256 minBalance, uint16 feeBps, uint16 protocolBps);
    event TierAdded(uint256 indexed tierId, uint256 minBalance, uint16 feeBps, uint16 protocolBps);
    event TierRemoved(uint256 indexed tierId);

    function setUp() public {
        // Deploy HORIZON token
        horizonToken = new MockERC20("Horizon Token", "HORIZON"); horizonToken.mint(address(this), 1_000_000_000 * 10 ** 18);

        // Deploy HorizonPerks
        perks = new HorizonPerks(address(horizonToken));
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(address(perks.horizonToken()), address(horizonToken));
        assertEq(perks.owner(), owner);
        assertEq(perks.getTierCount(), 5); // 5 default tiers
    }

    function test_Constructor_DefaultTiers() public view {
        // Tier 0: 0 HORIZON → 200 bps, 10% protocol
        (uint256 minBalance0, uint16 feeBps0, uint16 protocolBps0) = perks.feeTiers(0);
        assertEq(minBalance0, 0);
        assertEq(feeBps0, 200);
        assertEq(protocolBps0, 1000); // 10%

        // Tier 1: 10K HORIZON → 200 bps, 8% protocol
        (uint256 minBalance1, uint16 feeBps1, uint16 protocolBps1) = perks.feeTiers(1);
        assertEq(minBalance1, 10_000 * 10 ** 18);
        assertEq(feeBps1, 200);
        assertEq(protocolBps1, 800); // 8%

        // Tier 2: 50K HORIZON → 200 bps, 6% protocol
        (uint256 minBalance2, uint16 feeBps2, uint16 protocolBps2) = perks.feeTiers(2);
        assertEq(minBalance2, 50_000 * 10 ** 18);
        assertEq(feeBps2, 200);
        assertEq(protocolBps2, 600); // 6%

        // Tier 3: 100K HORIZON → 200 bps, 4% protocol
        (uint256 minBalance3, uint16 feeBps3, uint16 protocolBps3) = perks.feeTiers(3);
        assertEq(minBalance3, 100_000 * 10 ** 18);
        assertEq(feeBps3, 200);
        assertEq(protocolBps3, 400); // 4%

        // Tier 4: 500K HORIZON → 200 bps, 2% protocol
        (uint256 minBalance4, uint16 feeBps4, uint16 protocolBps4) = perks.feeTiers(4);
        assertEq(minBalance4, 500_000 * 10 ** 18);
        assertEq(feeBps4, 200);
        assertEq(protocolBps4, 200); // 2%
    }

    function test_RevertWhen_Constructor_InvalidToken() public {
        vm.expectRevert("Invalid HORIZON token");
        new HorizonPerks(address(0));
    }

    // ============ Fee Calculation Tests ============

    function test_FeeBpsFor_Tier0() public {
        // User with 0 HORIZON should get tier 0 (200 bps)
        assertEq(perks.feeBpsFor(user1), 200);
    }

    function test_FeeBpsFor_Tier1() public {
        // Give user exactly 10K HORIZON (tier 1 threshold)
        horizonToken.transfer(user1, 10_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200);
    }

    function test_FeeBpsFor_Tier2() public {
        // Give user 50K HORIZON (tier 2 threshold)
        horizonToken.transfer(user1, 50_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200);
    }

    function test_FeeBpsFor_Tier3() public {
        // Give user 100K HORIZON (tier 3 threshold)
        horizonToken.transfer(user1, 100_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200);
    }

    function test_FeeBpsFor_Tier4() public {
        // Give user 500K HORIZON (tier 4 threshold)
        horizonToken.transfer(user1, 500_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200);
    }

    function test_FeeBpsFor_AboveTier4() public {
        // Give user 1M HORIZON (above tier 4)
        horizonToken.transfer(user1, 1_000_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200); // Should still get tier 4
    }

    function test_FeeBpsFor_BetweenTiers() public {
        // Give user 25K HORIZON (between tier 1 and tier 2)
        horizonToken.transfer(user1, 25_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 200); // Should get tier 1
    }

    function test_FeeBpsForBalance() public view {
        assertEq(perks.feeBpsForBalance(0), 200); // Tier 0
        assertEq(perks.feeBpsForBalance(5_000 * 10 ** 18), 200); // Still tier 0
        assertEq(perks.feeBpsForBalance(10_000 * 10 ** 18), 200); // Tier 1
        assertEq(perks.feeBpsForBalance(50_000 * 10 ** 18), 200); // Tier 2
        assertEq(perks.feeBpsForBalance(100_000 * 10 ** 18), 200); // Tier 3
        assertEq(perks.feeBpsForBalance(500_000 * 10 ** 18), 200); // Tier 4
        assertEq(perks.feeBpsForBalance(1_000_000 * 10 ** 18), 200); // Still tier 4
    }

    // ============ Tier Level Tests ============

    function test_TierLevelFor() public {
        horizonToken.transfer(user1, 0);
        assertEq(perks.tierLevelFor(user1), 0);

        horizonToken.transfer(user2, 10_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user2), 1);

        horizonToken.transfer(user3, 100_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user3), 3);
    }

    function test_TierLevelForBalance() public view {
        assertEq(perks.tierLevelForBalance(0), 0);
        assertEq(perks.tierLevelForBalance(10_000 * 10 ** 18), 1);
        assertEq(perks.tierLevelForBalance(50_000 * 10 ** 18), 2);
        assertEq(perks.tierLevelForBalance(100_000 * 10 ** 18), 3);
        assertEq(perks.tierLevelForBalance(500_000 * 10 ** 18), 4);
    }

    // ============ Tier Management Tests ============

    function test_UpdateTier() public {
        vm.expectEmit(true, false, false, true);
        emit TierUpdated(1, 15_000 * 10 ** 18, 180, 900);

        perks.updateTier(1, 15_000 * 10 ** 18, 180, 900);

        (uint256 minBalance, uint16 feeBps, uint16 protocolBps) = perks.feeTiers(1);
        assertEq(minBalance, 15_000 * 10 ** 18);
        assertEq(feeBps, 180);
        assertEq(protocolBps, 900);
    }

    function test_RevertWhen_UpdateTier_InvalidTierId() public {
        vm.expectRevert("Invalid tier ID");
        perks.updateTier(999, 1000, 200, 1000);
    }

    function test_RevertWhen_UpdateTier_InvalidFeeBps_Zero() public {
        vm.expectRevert("Invalid fee bps");
        perks.updateTier(1, 15_000 * 10 ** 18, 0, 900);
    }

    function test_RevertWhen_UpdateTier_InvalidFeeBps_TooHigh() public {
        vm.expectRevert("Invalid fee bps");
        perks.updateTier(1, 15_000 * 10 ** 18, 1001, 900);
    }

    function test_RevertWhen_UpdateTier_InvalidOrdering_LowerThanPrevious() public {
        vm.expectRevert("Must be higher than previous tier");
        perks.updateTier(2, 5_000 * 10 ** 18, 200, 600); // Lower than tier 1
    }

    function test_RevertWhen_UpdateTier_InvalidOrdering_HigherThanNext() public {
        vm.expectRevert("Must be lower than next tier");
        perks.updateTier(2, 200_000 * 10 ** 18, 200, 600); // Higher than tier 3
    }

    function test_RevertWhen_UpdateTier_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        perks.updateTier(1, 15_000 * 10 ** 18, 180, 900);
    }

    function test_AddTier() public {
        vm.expectEmit(true, false, false, true);
        emit TierAdded(5, 1_000_000 * 10 ** 18, 200, 100);

        perks.addTier(1_000_000 * 10 ** 18, 200, 100);

        assertEq(perks.getTierCount(), 6);
        (uint256 minBalance, uint16 feeBps, uint16 protocolBps) = perks.feeTiers(5);
        assertEq(minBalance, 1_000_000 * 10 ** 18);
        assertEq(feeBps, 200);
        assertEq(protocolBps, 100);
    }

    function test_RevertWhen_AddTier_InvalidFeeBps() public {
        vm.expectRevert("Invalid fee bps");
        perks.addTier(1_000_000 * 10 ** 18, 0, 100);
    }

    function test_RevertWhen_AddTier_NotHigherThanLast() public {
        vm.expectRevert("Must be higher than last tier");
        perks.addTier(100_000 * 10 ** 18, 200, 100); // Lower than tier 4 (500K)
    }

    function test_RemoveLastTier() public {
        vm.expectEmit(true, false, false, false);
        emit TierRemoved(4);

        perks.removeLastTier();

        assertEq(perks.getTierCount(), 4);
    }

    function test_RemoveMultipleTiers() public {
        perks.removeLastTier();
        perks.removeLastTier();
        perks.removeLastTier();

        assertEq(perks.getTierCount(), 2);
    }

    function test_RevertWhen_RemoveLastTier_OnlyOneTier() public {
        // Remove all but one
        perks.removeLastTier();
        perks.removeLastTier();
        perks.removeLastTier();
        perks.removeLastTier();

        vm.expectRevert("Cannot remove all tiers");
        perks.removeLastTier();
    }

    // ============ View Functions Tests ============

    function test_GetAllTiers() public view {
        HorizonPerks.FeeTier[] memory tiers = perks.getAllTiers();
        assertEq(tiers.length, 5);
        assertEq(tiers[0].feeBps, 200);
        assertEq(tiers[4].feeBps, 200);
    }

    function test_GetTier() public view {
        HorizonPerks.FeeTier memory tier = perks.getTier(2);
        assertEq(tier.minBalance, 50_000 * 10 ** 18);
        assertEq(tier.feeBps, 200);
    }

    function test_RevertWhen_GetTier_InvalidId() public {
        vm.expectRevert("Invalid tier ID");
        perks.getTier(999);
    }

    function test_GetTierCount() public view {
        assertEq(perks.getTierCount(), 5);
    }

    function test_GetNextTierInfo_HasNext() public {
        // User with 25K HORIZON (tier 1)
        horizonToken.transfer(user1, 25_000 * 10 ** 18);

        (bool hasNextTier, uint256 nextTierMinBalance, uint16 nextTierFeeBps, uint16 nextTierProtocolBps, uint256 tokensNeeded) =
            perks.getNextTierInfo(user1);

        assertTrue(hasNextTier);
        assertEq(nextTierMinBalance, 50_000 * 10 ** 18); // Tier 2
        assertEq(nextTierFeeBps, 200);
        assertEq(nextTierProtocolBps, 600); // 6%
        assertEq(tokensNeeded, 25_000 * 10 ** 18); // Need 25K more
    }

    function test_GetNextTierInfo_NoNext() public {
        // User with 500K HORIZON (tier 4 - highest)
        horizonToken.transfer(user1, 500_000 * 10 ** 18);

        (bool hasNextTier, uint256 nextTierMinBalance, uint16 nextTierFeeBps, uint16 nextTierProtocolBps, uint256 tokensNeeded) =
            perks.getNextTierInfo(user1);

        assertFalse(hasNextTier);
        assertEq(nextTierMinBalance, 0);
        assertEq(nextTierFeeBps, 0);
        assertEq(nextTierProtocolBps, 0);
        assertEq(tokensNeeded, 0);
    }

    function test_GetNextTierInfo_AtExactThreshold() public {
        // User with exactly 50K HORIZON (tier 2 threshold)
        horizonToken.transfer(user1, 50_000 * 10 ** 18);

        (bool hasNextTier, uint256 nextTierMinBalance, uint16 nextTierFeeBps, uint16 nextTierProtocolBps, uint256 tokensNeeded) =
            perks.getNextTierInfo(user1);

        assertTrue(hasNextTier);
        assertEq(nextTierMinBalance, 100_000 * 10 ** 18); // Tier 3
        assertEq(nextTierFeeBps, 200);
        assertEq(nextTierProtocolBps, 400); // 4%
        assertEq(tokensNeeded, 50_000 * 10 ** 18);
    }

    // ============ Fee Calculation Helpers Tests ============

    function test_CalculateFee() public {
        uint256 amount = 1000 ether;

        // User with no HORIZON (tier 0: 200 bps = 2%)
        uint256 fee0 = perks.calculateFee(user1, amount);
        assertEq(fee0, 20 ether); // 2% of 1000

        // Give user 10K HORIZON (tier 1: 200 bps = 2%)
        horizonToken.transfer(user1, 10_000 * 10 ** 18);
        uint256 fee1 = perks.calculateFee(user1, amount);
        assertEq(fee1, 20 ether); // 2% of 1000

        // Give user more to reach tier 4 (200 bps = 2%)
        horizonToken.transfer(user1, 490_000 * 10 ** 18);
        uint256 fee4 = perks.calculateFee(user1, amount);
        assertEq(fee4, 20 ether); // 2% of 1000
    }

    function test_CalculateAmountAfterFee() public {
        uint256 amount = 1000 ether;

        // User with no HORIZON (tier 0: 2%)
        uint256 after0 = perks.calculateAmountAfterFee(user1, amount);
        assertEq(after0, 980 ether); // 1000 - 20

        // Give user 500K HORIZON (tier 4: 2%)
        horizonToken.transfer(user1, 500_000 * 10 ** 18);
        uint256 after4 = perks.calculateAmountAfterFee(user1, amount);
        assertEq(after4, 980 ether); // 1000 - 20
    }

    function testFuzz_CalculateFee(uint128 amount) public {
        vm.assume(amount > 0);

        uint256 fee = perks.calculateFee(user1, amount);
        uint256 amountAfterFee = perks.calculateAmountAfterFee(user1, amount);

        // Verify: fee + amountAfterFee = original amount
        assertEq(fee + amountAfterFee, amount);
    }

    // ============ Integration Tests ============

    function test_TierProgression() public {
        uint256 amount = 1000 ether;

        // Start at tier 0
        assertEq(perks.tierLevelFor(user1), 0);
        assertEq(perks.calculateFee(user1, amount), 20 ether);

        // Progress to tier 1
        horizonToken.transfer(user1, 10_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user1), 1);
        assertEq(perks.calculateFee(user1, amount), 20 ether);

        // Progress to tier 2
        horizonToken.transfer(user1, 40_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user1), 2);
        assertEq(perks.calculateFee(user1, amount), 20 ether);

        // Progress to tier 3
        horizonToken.transfer(user1, 50_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user1), 3);
        assertEq(perks.calculateFee(user1, amount), 20 ether);

        // Progress to tier 4
        horizonToken.transfer(user1, 400_000 * 10 ** 18);
        assertEq(perks.tierLevelFor(user1), 4);
        assertEq(perks.calculateFee(user1, amount), 20 ether);
    }

    function test_CustomTierConfiguration() public {
        // Create new perks with modified tiers
        perks.updateTier(1, 20_000 * 10 ** 18, 190, 900);
        perks.addTier(1_000_000 * 10 ** 18, 175, 50);

        // Test updated tier 1
        horizonToken.transfer(user1, 20_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user1), 190);

        // Test new tier 5
        horizonToken.transfer(user2, 1_000_000 * 10 ** 18);
        assertEq(perks.feeBpsFor(user2), 175);
        assertEq(perks.tierLevelFor(user2), 5);
    }

    function test_EdgeCases_ExactBoundaries() public {
        // Test exact threshold boundaries
        horizonToken.transfer(user1, 10_000 * 10 ** 18 - 1); // 1 wei below threshold
        assertEq(perks.feeBpsFor(user1), 200); // Still tier 0

        horizonToken.transfer(user1, 1); // Exactly at threshold
        assertEq(perks.feeBpsFor(user1), 200); // Now tier 1 (but same fee)
    }
}
