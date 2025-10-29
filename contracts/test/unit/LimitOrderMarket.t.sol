// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/markets/LimitOrderMarket.sol";
import "../../src/interfaces/IMarket.sol";
import "../../src/base/BaseMarket.sol";
import "../helpers/MarketTestHelper.sol";

/**
 * @title LimitOrderMarketTest
 * @notice Comprehensive unit tests for LimitOrderMarket with order book
 * @dev Tests cover:
 *      - Order placement (limit, market, post-only, IOC)
 *      - Order matching and execution
 *      - Order cancellation
 *      - Order book management
 *      - Partial fills
 *      - Price-time priority
 *      - Edge cases
 *      - Gas profiling
 */
contract LimitOrderMarketTest is MarketTestHelper {
    LimitOrderMarket public market;
    
    uint256 public constant MARKET_ID = 2;
    uint256 public constant OUTCOME_COUNT = 2; // Binary market
    uint256 public closeTime;

    function setUp() public {
        // Setup core contracts first
        setupCore();
        
        // Set close time
        closeTime = block.timestamp + 30 days;
        
        // Deploy market with core contracts
        market = new LimitOrderMarket(
            MARKET_ID,
            address(collateral),
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            closeTime,
            OUTCOME_COUNT
        );

        // Setup test environment
        registerMarket(MARKET_ID, creator);
        outcomeToken.setAMMAuthorization(address(market), true);
        outcomeToken.setAMMAuthorization(creator, true); // Allow creator to mint for testing
        fundStandardAccounts();
        approveMarketForAll(address(market));
        approveOutcomeTokenForAll(address(market)); // Approve ERC1155 transfers
        outcomeToken.setResolutionAuthorization(address(this), true);
    }

    // Helper to approve outcome tokens for all test accounts
    function approveOutcomeTokenForAll(address operator) internal {
        address[] memory accounts = new address[](6);
        accounts[0] = lp1;
        accounts[1] = lp2;
        accounts[2] = trader1;
        accounts[3] = trader2;
        accounts[4] = trader3;
        accounts[5] = trader4;
        
        for (uint256 i = 0; i < accounts.length; i++) {
            vm.prank(accounts[i]);
            outcomeToken.setApprovalForAll(operator, true);
        }
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(market.marketId(), MARKET_ID);
        assertEq(address(market.collateralToken()), address(collateral));
        assertEq(market.closeTime(), closeTime);
        assertEq(market.outcomeCount(), OUTCOME_COUNT);
        assertEq(uint256(market.marketType()), uint256(IMarket.MarketType.LimitOrder));
    }

    // ============ Limit Order Placement Tests ============

    function test_PlaceOrder_BuyLimit() public {
        uint256 price = 0.6e18; // 60%
        uint256 amount = 100 ether;
        uint256 collateralRequired = (amount * price) / 1e18;

        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            price,
            amount,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Check order was created
        LimitOrderMarket.Order memory order = market.getOrder(orderId);
        assertEq(order.trader, trader1);
        assertEq(order.outcomeId, 0);
        assertEq(order.price, price);
        assertEq(order.amount, amount);
        assertEq(order.filled, 0);
        assertTrue(order.isBuy);
        assertTrue(order.isActive);

        // Check collateral was escrowed
        assertEq(market.totalCollateral(), collateralRequired);

        // Check order is in book
        bytes32[] memory buyOrders = market.getBuyOrders(0);
        assertEq(buyOrders.length, 1);
        assertEq(buyOrders[0], orderId);
    }

    function test_PlaceOrder_SellLimit() public {
        // First get some outcome tokens
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        uint256 price = 0.55e18; // 55%
        uint256 amount = 50 ether;

        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            false,
            price,
            amount,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Check order was created
        LimitOrderMarket.Order memory order = market.getOrder(orderId);
        assertEq(order.trader, trader1);
        assertEq(order.outcomeId, 0);
        assertEq(order.price, price);
        assertEq(order.amount, amount);
        assertFalse(order.isBuy);

        // Check order is in book
        bytes32[] memory sellOrders = market.getSellOrders(0);
        assertEq(sellOrders.length, 1);
    }

    function test_PlaceOrder_MultipleOrders_SortedByPrice() public {
        // Place multiple buy orders with different prices
        vm.startPrank(trader1);
        bytes32 order1 = market.placeOrder(0, true, 0.5e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        bytes32 order2 = market.placeOrder(0, true, 0.7e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        bytes32 order3 = market.placeOrder(0, true, 0.6e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        vm.stopPrank();

        // Check orders are sorted by price (highest first)
        bytes32[] memory buyOrders = market.getBuyOrders(0);
        assertEq(buyOrders.length, 3);
        assertEq(buyOrders[0], order2); // 0.7
        assertEq(buyOrders[1], order3); // 0.6
        assertEq(buyOrders[2], order1); // 0.5
    }

    function test_PlaceOrder_PriceTimePriority() public {
        // Place two orders at same price
        vm.prank(trader1);
        bytes32 order1 = market.placeOrder(0, true, 0.6e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);

        vm.warp(block.timestamp + 1);

        vm.prank(trader2);
        bytes32 order2 = market.placeOrder(0, true, 0.6e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);

        // Check time priority (earlier order first)
        bytes32[] memory buyOrders = market.getBuyOrders(0);
        assertEq(buyOrders[0], order1);
        assertEq(buyOrders[1], order2);
    }

    // ============ Order Matching Tests ============

    function test_OrderMatching_FullFill() public {
        // Place sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        bytes32 sellOrderId = market.placeOrder(
            0,
            false,
            0.6e18,
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Place matching buy order
        vm.prank(trader2);
        bytes32 buyOrderId = market.placeOrder(
            0,
            true,
            0.65e18, // Higher price, should match
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Check both orders are fully filled
        LimitOrderMarket.Order memory sellOrder = market.getOrder(sellOrderId);
        LimitOrderMarket.Order memory buyOrder = market.getOrder(buyOrderId);

        assertEq(sellOrder.filled, 50 ether);
        assertEq(buyOrder.filled, 50 ether);
        assertFalse(sellOrder.isActive);
        assertFalse(buyOrder.isActive);

        // Check trader2 received outcome tokens
        uint256 balance = getOutcomeBalance(trader2, MARKET_ID, 0);
        assertEq(balance, 50 ether);

        // Check order books are empty
        assertEq(market.getBuyOrders(0).length, 0);
        assertEq(market.getSellOrders(0).length, 0);
    }

    function test_OrderMatching_PartialFill() public {
        // Place large sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        bytes32 sellOrderId = market.placeOrder(
            0,
            false,
            0.6e18,
            100 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Place smaller buy order
        vm.prank(trader2);
        bytes32 buyOrderId = market.placeOrder(
            0,
            true,
            0.65e18,
            40 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Check partial fill
        LimitOrderMarket.Order memory sellOrder = market.getOrder(sellOrderId);
        LimitOrderMarket.Order memory buyOrder = market.getOrder(buyOrderId);

        assertEq(sellOrder.filled, 40 ether);
        assertEq(buyOrder.filled, 40 ether);
        assertTrue(sellOrder.isActive); // Still has remaining amount
        assertFalse(buyOrder.isActive); // Fully filled

        // Check sell order still in book with reduced amount
        bytes32[] memory sellOrders = market.getSellOrders(0);
        assertEq(sellOrders.length, 1);
    }

    function test_OrderMatching_MultipleOrders() public {
        // Place multiple sell orders
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 200 ether);

        vm.startPrank(trader1);
        market.placeOrder(0, false, 0.55e18, 30 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, false, 0.60e18, 40 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, false, 0.65e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
        vm.stopPrank();

        // Place large buy order that matches multiple sells
        vm.prank(trader2);
        market.placeOrder(
            0,
            true,
            0.70e18,
            100 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Check all orders matched correctly
        // Should match: 30 @ 0.55 + 40 @ 0.60 + 30 @ 0.65 = 100 total
        assertEq(market.getSellOrders(0).length, 1); // One order partially filled
        assertEq(market.getBuyOrders(0).length, 0); // Buy fully matched
    }

    function test_OrderMatching_NoCross_DifferentPrices() public {
        // Place sell order at high price
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.70e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Place buy order at lower price (shouldn't match)
        vm.prank(trader2);
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Check both orders remain in book
        assertEq(market.getSellOrders(0).length, 1);
        assertEq(market.getBuyOrders(0).length, 1);
    }

    // ============ Market Order Tests ============

    function test_PlaceMarketOrder_Buy() public {
        // Setup: Place sell order in book
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Place market buy order
        vm.prank(trader2);
        uint256 filled = market.placeMarketOrder(
            0,
            true,
            30 ether,
            1000 // 10% max slippage
        );

        assertEq(filled, 30 ether);

        // Check trader2 received tokens
        uint256 balance = getOutcomeBalance(trader2, MARKET_ID, 0);
        assertEq(balance, 30 ether);
    }

    function test_PlaceMarketOrder_Sell() public {
        // Setup: Place buy order in book
        vm.prank(trader1);
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Give trader2 outcome tokens
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader2, 100 ether);

        // Place market sell order
        vm.prank(trader2);
        uint256 filled = market.placeMarketOrder(
            0,
            false,
            30 ether,
            1000
        );

        assertEq(filled, 30 ether);
    }

    function test_RevertWhen_PlaceMarketOrder_NoLiquidity() public {
        // Try to place market order with empty book
        vm.prank(trader1);
        vm.expectRevert(LimitOrderMarket.NoMatchingOrders.selector);
        market.placeMarketOrder(0, true, 50 ether, 1000);
    }

    // ============ Post-Only Order Tests ============

    function test_PlaceOrder_PostOnly_NoMatch() public {
        // Place post-only order (should add to book)
        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.60e18,
            50 ether,
            LimitOrderMarket.OrderType.POST_ONLY
        );

        // Check order is in book
        assertEq(market.getBuyOrders(0).length, 1);
        assertTrue(market.getOrder(orderId).isActive);
    }

    function test_RevertWhen_PlaceOrder_PostOnly_WouldMatch() public {
        // Place sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Try to place post-only buy at higher price (would match)
        vm.prank(trader2);
        vm.expectRevert(LimitOrderMarket.PostOnlyWouldMatch.selector);
        market.placeOrder(
            0,
            true,
            0.65e18,
            50 ether,
            LimitOrderMarket.OrderType.POST_ONLY
        );
    }

    // ============ IOC Order Tests ============

    function test_PlaceOrder_IOC_FullMatch() public {
        // Place sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Place IOC order that fully matches
        vm.prank(trader2);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.65e18,
            50 ether,
            LimitOrderMarket.OrderType.IOC
        );

        // Check order is fully filled and not in book
        assertFalse(market.getOrder(orderId).isActive);
        assertEq(market.getBuyOrders(0).length, 0);
    }

    function test_PlaceOrder_IOC_PartialMatch_Cancelled() public {
        // Place small sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 30 ether, LimitOrderMarket.OrderType.LIMIT);

        // Place larger IOC order (should match 30, cancel 20)
        vm.prank(trader2);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.65e18,
            50 ether,
            LimitOrderMarket.OrderType.IOC
        );

        // Check order is cancelled (partial fill + cancel for remaining)
        assertFalse(market.getOrder(orderId).isActive);
        assertEq(market.getOrder(orderId).filled, 30 ether);
    }

    // ============ Order Cancellation Tests ============

    function test_CancelOrder() public {
        // Place order
        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.60e18,
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        uint256 balanceBefore = collateral.balanceOf(trader1);

        // Cancel order
        vm.prank(trader1);
        market.cancelOrder(orderId);

        // Check order is inactive
        assertFalse(market.getOrder(orderId).isActive);

        // Check collateral refunded
        uint256 balanceAfter = collateral.balanceOf(trader1);
        uint256 refunded = balanceAfter - balanceBefore;
        assertEq(refunded, (50 ether * 0.60e18) / 1e18);

        // Check order removed from book
        assertEq(market.getBuyOrders(0).length, 0);
    }

    function test_CancelOrder_PartiallyFilled() public {
        // Place sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        bytes32 sellOrderId = market.placeOrder(
            0,
            false,
            0.60e18,
            100 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        // Partially fill
        vm.prank(trader2);
        market.placeOrder(0, true, 0.65e18, 40 ether, LimitOrderMarket.OrderType.LIMIT);

        uint256 balanceBefore = getOutcomeBalance(trader1, MARKET_ID, 0);

        // Cancel remaining
        vm.prank(trader1);
        market.cancelOrder(sellOrderId);

        // Check remaining tokens refunded (60 out of 100)
        uint256 balanceAfter = getOutcomeBalance(trader1, MARKET_ID, 0);
        assertEq(balanceAfter - balanceBefore, 60 ether);
    }

    function test_RevertWhen_CancelOrder_NotOwner() public {
        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.60e18,
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        vm.prank(trader2);
        vm.expectRevert(LimitOrderMarket.OrderNotOwned.selector);
        market.cancelOrder(orderId);
    }

    function test_RevertWhen_CancelOrder_AlreadyCancelled() public {
        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.60e18,
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        vm.prank(trader1);
        market.cancelOrder(orderId);

        vm.prank(trader1);
        vm.expectRevert(LimitOrderMarket.OrderNotFound.selector);
        market.cancelOrder(orderId);
    }

    // ============ View Function Tests ============

    function test_GetBestBid() public {
        vm.startPrank(trader1);
        market.placeOrder(0, true, 0.50e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, true, 0.60e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, true, 0.55e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        vm.stopPrank();

        assertEq(market.getBestBid(0), 0.60e18);
    }

    function test_GetBestAsk() public {
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.startPrank(trader1);
        market.placeOrder(0, false, 0.65e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, false, 0.70e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        market.placeOrder(0, false, 0.60e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        vm.stopPrank();

        assertEq(market.getBestAsk(0), 0.60e18);
    }

    function test_GetSpread() public {
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.65e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);

        vm.prank(trader2);
        market.placeOrder(0, true, 0.60e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);

        assertEq(market.getSpread(0), 0.05e18); // 5% spread
    }

    function test_GetUserOrders() public {
        vm.startPrank(trader1);
        bytes32 order1 = market.placeOrder(0, true, 0.60e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        bytes32 order2 = market.placeOrder(0, true, 0.55e18, 10 ether, LimitOrderMarket.OrderType.LIMIT);
        vm.stopPrank();

        bytes32[] memory userOrders = market.getUserOrders(trader1);
        assertEq(userOrders.length, 2);
        assertEq(userOrders[0], order1);
        assertEq(userOrders[1], order2);
    }

    // ============ Validation Tests ============

    function test_RevertWhen_PlaceOrder_InvalidOutcome() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidOutcomeId.selector);
        market.placeOrder(5, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_ZeroAmount() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.InvalidAmount.selector);
        market.placeOrder(0, true, 0.60e18, 0, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_InvalidPrice_Zero() public {
        vm.prank(trader1);
        vm.expectRevert(LimitOrderMarket.InvalidPrice.selector);
        market.placeOrder(0, true, 0, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_InvalidPrice_TooHigh() public {
        vm.prank(trader1);
        vm.expectRevert(LimitOrderMarket.InvalidPrice.selector);
        market.placeOrder(0, true, 1.5e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_AfterClose() public {
        vm.warp(closeTime + 1);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.MarketClosed.selector);
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_AfterResolution() public {
        outcomeToken.setWinningOutcome(MARKET_ID, 0);

        vm.prank(trader1);
        vm.expectRevert(BaseMarket.MarketResolved.selector);
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    function test_RevertWhen_PlaceOrder_WhenPaused() public {
        market.pause();

        vm.prank(trader1);
        vm.expectRevert();
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
    }

    // ============ Admin Tests ============

    function test_Pause() public {
        market.pause();
        assertTrue(market.paused());
    }

    function test_Unpause() public {
        market.pause();
        market.unpause();
        assertFalse(market.paused());
    }

    function test_SetAdmin() public {
        market.setAdmin(trader1);
        assertEq(market.admin(), trader1);
    }

    function test_RevertWhen_Pause_NotAdmin() public {
        vm.prank(trader1);
        vm.expectRevert(BaseMarket.Unauthorized.selector);
        market.pause();
    }

    // ============ Gas Profiling Tests ============

    function test_Gas_PlaceOrder() public {
        vm.prank(trader1);
        uint256 gasBefore = gasleft();
        market.placeOrder(0, true, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for limit order placement", gasUsed);
        
        // Target: <200k gas
    }

    function test_Gas_OrderMatching() public {
        // Setup sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Measure gas for matching buy order
        vm.prank(trader2);
        uint256 gasBefore = gasleft();
        market.placeOrder(0, true, 0.65e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for order matching", gasUsed);
    }

    function test_Gas_MarketOrder() public {
        // Setup sell order
        vm.prank(creator);
        outcomeToken.mintOutcome(MARKET_ID, 0, trader1, 100 ether);

        vm.prank(trader1);
        market.placeOrder(0, false, 0.60e18, 50 ether, LimitOrderMarket.OrderType.LIMIT);

        // Measure gas for market order
        vm.prank(trader2);
        uint256 gasBefore = gasleft();
        market.placeMarketOrder(0, true, 30 ether, 1000);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for market order", gasUsed);
    }

    function test_Gas_CancelOrder() public {
        vm.prank(trader1);
        bytes32 orderId = market.placeOrder(
            0,
            true,
            0.60e18,
            50 ether,
            LimitOrderMarket.OrderType.LIMIT
        );

        vm.prank(trader1);
        uint256 gasBefore = gasleft();
        market.cancelOrder(orderId);
        uint256 gasUsed = gasBefore - gasleft();
        
        emit log_named_uint("Gas used for order cancellation", gasUsed);
    }
}
