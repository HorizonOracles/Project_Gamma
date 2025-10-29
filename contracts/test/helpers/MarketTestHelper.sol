// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/interfaces/IMarket.sol";
import "../../src/OutcomeToken.sol";
import "../../src/FeeSplitter.sol";
import "../../src/HorizonPerks.sol";
import "../../src/HorizonToken.sol";
import "../mocks/MockERC20.sol";

/**
 * @title MarketTestHelper
 * @notice Reusable test utilities for all market types
 * @dev Provides common setup, helper functions, and assertions for market testing
 */
contract MarketTestHelper is Test {
    // ============ Core Contracts ============
    
    OutcomeToken public outcomeToken;
    FeeSplitter public feeSplitter;
    HorizonPerks public horizonPerks;
    HorizonToken public horizonToken;
    MockERC20 public collateral;

    // ============ Test Accounts ============
    
    address public owner = address(this);
    address public treasury = address(0x1);
    address public creator = address(0x2);
    address public lp1 = address(0x3);
    address public lp2 = address(0x4);
    address public trader1 = address(0x5);
    address public trader2 = address(0x6);
    address public trader3 = address(0x7);
    address public trader4 = address(0x8);

    // ============ Test Configuration ============
    
    uint256 public constant DEFAULT_INITIAL_LIQUIDITY = 10000 ether;
    uint256 public constant DEFAULT_COLLATERAL_AMOUNT = 1_000_000 ether;
    
    // ============ Setup Functions ============

    /**
     * @notice Deploys core system contracts
     * @dev Call this in your test's setUp() function before deploying market
     */
    function setupCore() internal {
        // Deploy tokens
        collateral = new MockERC20("USDC", "USDC");
        horizonToken = new HorizonToken(1_000_000_000 * 10 ** 18);

        // Deploy core contracts
        outcomeToken = new OutcomeToken("https://api.horizonoracles.io/metadata/{id}");
        feeSplitter = new FeeSplitter(treasury);
        horizonPerks = new HorizonPerks(address(horizonToken));
    }

    /**
     * @notice Registers a market with the core contracts
     * @param marketId Market identifier
     * @param marketCreator Creator address
     */
    function registerMarket(uint256 marketId, address marketCreator) internal {
        outcomeToken.registerMarket(marketId, collateral);
        feeSplitter.registerMarket(marketId, marketCreator);
    }

    /**
     * @notice Funds test accounts with collateral
     * @param accounts Array of addresses to fund
     * @param amount Amount to fund each account
     */
    function fundAccounts(address[] memory accounts, uint256 amount) internal {
        for (uint256 i = 0; i < accounts.length; i++) {
            collateral.mint(accounts[i], amount);
        }
    }

    /**
     * @notice Funds standard test accounts with default amounts
     */
    function fundStandardAccounts() internal {
        address[] memory accounts = new address[](6);
        accounts[0] = lp1;
        accounts[1] = lp2;
        accounts[2] = trader1;
        accounts[3] = trader2;
        accounts[4] = trader3;
        accounts[5] = trader4;
        
        fundAccounts(accounts, DEFAULT_COLLATERAL_AMOUNT);
    }

    /**
     * @notice Approves a market contract for all test accounts
     * @param market Market address
     */
    function approveMarketForAll(address market) internal {
        address[] memory accounts = new address[](6);
        accounts[0] = lp1;
        accounts[1] = lp2;
        accounts[2] = trader1;
        accounts[3] = trader2;
        accounts[4] = trader3;
        accounts[5] = trader4;
        
        for (uint256 i = 0; i < accounts.length; i++) {
            vm.prank(accounts[i]);
            collateral.approve(market, type(uint256).max);
        }
    }

    /**
     * @notice Complete setup for a market test
     * @param marketId Market identifier
     * @param market Market contract address
     */
    function setupMarketTest(uint256 marketId, address market) internal {
        setupCore();
        registerMarket(marketId, creator);
        outcomeToken.setAMMAuthorization(market, true);
        fundStandardAccounts();
        approveMarketForAll(market);
        
        // Authorize test contract as resolver (for testing resolution scenarios)
        outcomeToken.setResolutionAuthorization(address(this), true);
    }

    // ============ Helper Functions ============

    /**
     * @notice Adds initial liquidity from lp1
     * @param market Market contract
     * @param amount Liquidity amount
     * @return lpTokens LP tokens received
     */
    function addInitialLiquidity(IMarket market, uint256 amount) internal returns (uint256 lpTokens) {
        vm.prank(lp1);
        lpTokens = market.addLiquidity(amount);
    }

    /**
     * @notice Gets user's outcome token balance
     * @param user User address
     * @param marketId Market ID
     * @param outcomeId Outcome ID
     * @return Balance of outcome tokens
     */
    function getOutcomeBalance(address user, uint256 marketId, uint256 outcomeId) 
        internal 
        view 
        returns (uint256) 
    {
        return outcomeToken.balanceOfOutcome(user, marketId, outcomeId);
    }

    /**
     * @notice Warps to just before market close
     * @param market Market contract
     */
    function warpBeforeClose(IMarket market) internal {
        vm.warp(market.closeTime() - 1 hours);
    }

    /**
     * @notice Warps to just after market close
     * @param market Market contract
     */
    function warpAfterClose(IMarket market) internal {
        vm.warp(market.closeTime() + 1 hours);
    }

    /**
     * @notice Resolves a market with specified outcome
     * @param marketId Market identifier
     * @param winningOutcome Winning outcome ID
     */
    function resolveMarket(uint256 marketId, uint256 winningOutcome) internal {
        outcomeToken.setWinningOutcome(marketId, winningOutcome);
    }

    /**
     * @notice Calculates expected price impact for a trade
     * @param reserveIn Reserve of token being bought
     * @param reserveOut Reserve of token being sold
     * @param amountIn Amount being traded
     * @return amountOut Expected amount out
     */
    function calculateCPMMOutput(uint256 reserveIn, uint256 reserveOut, uint256 amountIn)
        internal
        pure
        returns (uint256 amountOut)
    {
        uint256 k = reserveIn * reserveOut;
        uint256 newReserveIn = reserveIn + amountIn;
        uint256 newReserveOut = k / newReserveIn;
        amountOut = reserveOut - newReserveOut;
    }

    // ============ Assertion Helpers ============

    /**
     * @notice Asserts that a market is in the expected state
     * @param market Market contract
     * @param expectedType Expected market type
     * @param expectedOutcomes Expected number of outcomes
     * @param shouldBePaused Expected pause state
     * @param shouldBeResolved Expected resolution state
     */
    function assertMarketState(
        IMarket market,
        IMarket.MarketType expectedType,
        uint256 expectedOutcomes,
        bool shouldBePaused,
        bool shouldBeResolved
    ) internal view {
        IMarket.MarketInfo memory info = market.getMarketInfo();
        
        assertEq(uint256(info.marketType), uint256(expectedType), "Market type mismatch");
        assertEq(info.outcomeCount, expectedOutcomes, "Outcome count mismatch");
        assertEq(info.isPaused, shouldBePaused, "Pause state mismatch");
        assertEq(info.isResolved, shouldBeResolved, "Resolution state mismatch");
    }

    /**
     * @notice Asserts price is within expected bounds
     * @param actualPrice Actual price
     * @param expectedPrice Expected price
     * @param toleranceBps Tolerance in basis points (100 = 1%)
     */
    function assertPriceWithinTolerance(
        uint256 actualPrice,
        uint256 expectedPrice,
        uint256 toleranceBps
    ) internal pure {
        uint256 diff = actualPrice > expectedPrice 
            ? actualPrice - expectedPrice 
            : expectedPrice - actualPrice;
        
        uint256 maxDiff = (expectedPrice * toleranceBps) / 10000;
        
        require(diff <= maxDiff, "Price outside tolerance");
    }

    /**
     * @notice Asserts prices sum to approximately 1.0 (for probability checks)
     * @param prices Array of prices
     * @param toleranceBps Tolerance in basis points
     */
    function assertPricesSumToOne(uint256[] memory prices, uint256 toleranceBps) internal pure {
        uint256 sum = 0;
        for (uint256 i = 0; i < prices.length; i++) {
            sum += prices[i];
        }
        
        assertPriceWithinTolerance(sum, 1e18, toleranceBps);
    }

    /**
     * @notice Asserts slippage protection works
     * @param market Market contract
     * @param outcomeId Outcome to buy
     * @param collateralIn Amount to spend
     * @param minTokensOut Minimum expected (will cause revert)
     */
    function assertSlippageReverts(
        IMarket market,
        uint256 outcomeId,
        uint256 collateralIn,
        uint256 minTokensOut
    ) internal {
        vm.expectRevert();
        market.buy(outcomeId, collateralIn, minTokensOut);
    }

    // ============ Gas Profiling Helpers ============

    /**
     * @notice Profiles gas for a buy trade
     * @param market Market contract
     * @param user User address
     * @param outcomeId Outcome to buy
     * @param amount Amount to spend
     * @return gasUsed Gas consumed
     */
    function profileBuy(IMarket market, address user, uint256 outcomeId, uint256 amount)
        internal
        returns (uint256 gasUsed)
    {
        uint256 gasBefore = gasleft();
        vm.prank(user);
        market.buy(outcomeId, amount, 0);
        gasUsed = gasBefore - gasleft();
    }

    /**
     * @notice Profiles gas for adding liquidity
     * @param market Market contract
     * @param user User address
     * @param amount Amount to add
     * @return gasUsed Gas consumed
     */
    function profileAddLiquidity(IMarket market, address user, uint256 amount)
        internal
        returns (uint256 gasUsed)
    {
        uint256 gasBefore = gasleft();
        vm.prank(user);
        market.addLiquidity(amount);
        gasUsed = gasBefore - gasleft();
    }

    // ============ Event Testing Helpers ============

    /**
     * @notice Expects a Trade event
     * @param trader Trader address
     * @param outcomeId Outcome ID
     * @param isBuy True for buy, false for sell
     */
    function expectTradeEvent(address trader, uint256 outcomeId, bool isBuy) internal {
        vm.expectEmit(true, true, false, false);
        emit IMarket.Trade(trader, outcomeId, 0, 0, 0, isBuy);
    }

    /**
     * @notice Expects a LiquidityChanged event
     * @param provider LP provider address
     * @param isAddition True for add, false for remove
     */
    function expectLiquidityEvent(address provider, bool isAddition) internal {
        vm.expectEmit(true, false, false, false);
        emit IMarket.LiquidityChanged(provider, 0, isAddition);
    }

    // ============ Utility Functions ============

    /**
     * @notice Creates an array of sequential outcome IDs
     * @param count Number of outcomes
     * @return Array [0, 1, 2, ..., count-1]
     */
    function createOutcomeArray(uint256 count) internal pure returns (uint256[] memory) {
        uint256[] memory outcomes = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            outcomes[i] = i;
        }
        return outcomes;
    }

    /**
     * @notice Calculates percentage difference between two values
     * @param a First value
     * @param b Second value
     * @return Percentage difference in basis points
     */
    function percentDiff(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 diff = a > b ? a - b : b - a;
        return (diff * 10000) / (a > b ? a : b);
    }
}
