// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../base/BaseMarket.sol";

/**
 * @title LimitOrderMarket
 * @notice Professional trading interface with limit orders and on-chain order matching
 * @dev Implements a central limit order book (CLOB) for prediction markets
 *      
 *      Key Features:
 *      - Limit buy/sell orders at custom prices
 *      - Price-time priority matching
 *      - Partial order fills
 *      - Order cancellation
 *      - Market orders (immediate execution)
 *      - Post-only orders (no immediate matching)
 *      - IOC (Immediate-or-Cancel) orders
 * 
 *      Gas Target: <200k per order placement
 */
contract LimitOrderMarket is BaseMarket {
    using SafeERC20 for IERC20;

    // ============ Types ============

    enum OrderType {
        LIMIT,      // Standard limit order
        MARKET,     // Market order (immediate execution)
        POST_ONLY,  // Only add to book, don't match
        IOC         // Immediate-or-cancel (match or cancel)
    }

    struct Order {
        address trader;
        uint8 outcomeId;
        uint256 price;          // Price in PRICE_PRECISION (1e18 = 1.0)
        uint256 amount;         // Total amount
        uint256 filled;         // Amount filled so far
        uint256 totalCost;      // Total cost/proceeds (for calculating avg price)
        uint256 timestamp;      // Order creation time
        bool isBuy;             // True for buy, false for sell
        OrderType orderType;
        bool isActive;          // False when cancelled or fully filled
    }

    // ============ Errors ============

    error InvalidPrice();
    error OrderNotFound();
    error OrderAlreadyFilled();
    error OrderNotOwned();
    error NoMatchingOrders();
    error PostOnlyWouldMatch();
    error MaxOrdersReached();

    // ============ Events ============

    event OrderPlaced(
        bytes32 indexed orderId,
        address indexed trader,
        uint8 indexed outcomeId,
        bool isBuy,
        uint256 price,
        uint256 amount,
        OrderType orderType
    );

    event OrderMatched(
        bytes32 indexed buyOrderId,
        bytes32 indexed sellOrderId,
        uint8 indexed outcomeId,
        uint256 price,
        uint256 amount,
        address buyer,
        address seller
    );

    event OrderCancelled(
        bytes32 indexed orderId,
        address indexed trader,
        uint256 amountRemaining
    );

    event OrderPartiallyFilled(
        bytes32 indexed orderId,
        uint256 filledAmount,
        uint256 remainingAmount
    );

    // ============ State Variables ============

    /// @notice Order storage
    mapping(bytes32 => Order) public orders;

    /// @notice Buy order IDs per outcome, sorted by price (highest first)
    mapping(uint8 => bytes32[]) public buyOrdersByOutcome;

    /// @notice Sell order IDs per outcome, sorted by price (lowest first)
    mapping(uint8 => bytes32[]) public sellOrdersByOutcome;

    /// @notice User's active order IDs
    mapping(address => bytes32[]) public userOrders;

    /// @notice Order counter for ID generation
    uint256 private orderNonce;

    /// @notice Maximum orders per outcome per side (gas limit protection)
    uint256 public constant MAX_ORDERS_PER_SIDE = 100;

    // ============ Constructor ============

    /**
     * @notice Initializes a limit order market
     * @param _marketId Market identifier
     * @param _collateralToken Collateral token address
     * @param _outcomeToken Outcome token contract
     * @param _feeSplitter Fee splitter contract
     * @param _horizonPerks Horizon perks contract
     * @param _closeTime Market close timestamp
     * @param _outcomeCount Number of outcomes (2 for binary, 3+ for multi-choice)
     */
    constructor(
        uint256 _marketId,
        address _collateralToken,
        address _outcomeToken,
        address _feeSplitter,
        address _horizonPerks,
        uint256 _closeTime,
        uint256 _outcomeCount
    )
        BaseMarket(
            _marketId,
            MarketType.LimitOrder,
            _collateralToken,
            _outcomeToken,
            _feeSplitter,
            _horizonPerks,
            _closeTime,
            _outcomeCount,
            "Limit Order Market LP",
            "LOM-LP"
        )
    {
        admin = msg.sender;
    }

    // ============ Order Placement ============

    /**
     * @notice Places a limit order
     * @param outcomeId Outcome to trade
     * @param isBuy True for buy order, false for sell
     * @param price Price per outcome token (in PRICE_PRECISION)
     * @param amount Amount of outcome tokens
     * @param orderType Type of order (LIMIT, POST_ONLY, IOC)
     * @return orderId The ID of the created order
     */
    function placeOrder(
        uint8 outcomeId,
        bool isBuy,
        uint256 price,
        uint256 amount,
        OrderType orderType
    ) external nonReentrant whenNotPaused returns (bytes32 orderId) {
        if (block.timestamp >= closeTime) revert MarketClosed();
        if (outcomeToken.isResolved(marketId)) revert MarketResolved();
        if (outcomeId >= outcomeCount) revert InvalidOutcomeId();
        if (amount == 0) revert InvalidAmount();
        if (price == 0 || price > PRICE_PRECISION) revert InvalidPrice();
        if (orderType == OrderType.MARKET) revert InvalidState(); // Use placeMarketOrder instead

        // Check if POST_ONLY order would match
        if (orderType == OrderType.POST_ONLY) {
            if (isBuy) {
                // Check if there are sell orders at or below our buy price
                bytes32[] storage sellOrders = sellOrdersByOutcome[outcomeId];
                if (sellOrders.length > 0 && orders[sellOrders[0]].price <= price) {
                    revert PostOnlyWouldMatch();
                }
            } else {
                // Check if there are buy orders at or above our sell price
                bytes32[] storage buyOrders = buyOrdersByOutcome[outcomeId];
                if (buyOrders.length > 0 && orders[buyOrders[0]].price >= price) {
                    revert PostOnlyWouldMatch();
                }
            }
        }

        // Generate order ID
        orderId = keccak256(abi.encodePacked(msg.sender, orderNonce++, block.timestamp));

        // Create order
        orders[orderId] = Order({
            trader: msg.sender,
            outcomeId: outcomeId,
            price: price,
            amount: amount,
            filled: 0,
            totalCost: 0,
            timestamp: block.timestamp,
            isBuy: isBuy,
            orderType: orderType,
            isActive: true
        });

        if (isBuy) {
            // Escrow collateral for buy orders
            uint256 collateralRequired = (amount * price) / PRICE_PRECISION;
            collateralToken.safeTransferFrom(msg.sender, address(this), collateralRequired);
            totalCollateral += collateralRequired;
        } else {
            // Escrow outcome tokens for sell orders
            outcomeToken.safeTransferFrom(
                msg.sender,
                address(this),
                outcomeToken.encodeTokenId(marketId, outcomeId),
                amount,
                ""
            );
        }

        // Add to user's orders
        userOrders[msg.sender].push(orderId);

        emit OrderPlaced(orderId, msg.sender, outcomeId, isBuy, price, amount, orderType);

        // Try to match order (except for POST_ONLY)
        if (orderType != OrderType.POST_ONLY) {
            _matchOrder(orderId);
        }

        // Handle partially filled orders
        if (orders[orderId].filled < orders[orderId].amount && orders[orderId].isActive) {
            if (orderType == OrderType.IOC) {
                // Cancel IOC order if not fully filled
                _cancelOrder(orderId);
            } else {
                // Add to order book if not fully filled
                _addToOrderBook(orderId);
            }
        }

        return orderId;
    }

    /**
     * @notice Places a market order (immediate execution at best available price)
     * @param outcomeId Outcome to trade
     * @param isBuy True for buy, false for sell
     * @param amount Amount of outcome tokens
     * @param maxSlippage Maximum acceptable slippage (in PRICE_PRECISION)
     * @return filledAmount Amount successfully filled
     */
    function placeMarketOrder(
        uint8 outcomeId,
        bool isBuy,
        uint256 amount,
        uint256 maxSlippage
    ) external nonReentrant whenNotPaused returns (uint256 filledAmount) {
        if (block.timestamp >= closeTime) revert MarketClosed();
        if (outcomeToken.isResolved(marketId)) revert MarketResolved();
        if (outcomeId >= outcomeCount) revert InvalidOutcomeId();
        if (amount == 0) revert InvalidAmount();

        // Generate temporary order ID
        bytes32 orderId = keccak256(abi.encodePacked(msg.sender, orderNonce++, block.timestamp));

        // Create temporary order
        orders[orderId] = Order({
            trader: msg.sender,
            outcomeId: outcomeId,
            price: isBuy ? PRICE_PRECISION : 0, // Buy at any price up to 1.0, sell at any price
            amount: amount,
            filled: 0,
            totalCost: 0,
            timestamp: block.timestamp,
            isBuy: isBuy,
            orderType: OrderType.MARKET,
            isActive: true
        });

        if (isBuy) {
            // Escrow max collateral needed
            uint256 maxCollateral = amount; // Max 1:1 ratio
            collateralToken.safeTransferFrom(msg.sender, address(this), maxCollateral);
            totalCollateral += maxCollateral;
        } else {
            // Escrow outcome tokens
            outcomeToken.safeTransferFrom(
                msg.sender,
                address(this),
                outcomeToken.encodeTokenId(marketId, outcomeId),
                amount,
                ""
            );
        }

        // Get best price for slippage reference
        uint256 bestPrice;
        if (isBuy) {
            // For buy orders, check best ask (sell orders)
            bytes32[] storage sellOrders = sellOrdersByOutcome[outcomeId];
            if (sellOrders.length == 0) revert NoMatchingOrders();
            bestPrice = orders[sellOrders[0]].price;
        } else {
            // For sell orders, check best bid (buy orders)
            bytes32[] storage buyOrders = buyOrdersByOutcome[outcomeId];
            if (buyOrders.length == 0) revert NoMatchingOrders();
            bestPrice = orders[buyOrders[0]].price;
        }

        // Match the order
        filledAmount = _matchOrder(orderId);

        if (filledAmount == 0) revert NoMatchingOrders();

        // Check slippage against best available price
        uint256 avgPrice = _calculateAveragePrice(orderId);
        if (isBuy) {
            // For buy: avgPrice should not exceed bestPrice * (1 + maxSlippage)
            uint256 maxAcceptablePrice = (bestPrice * (10000 + maxSlippage)) / 10000;
            if (avgPrice > maxAcceptablePrice) {
                revert SlippageExceeded();
            }
        } else {
            // For sell: avgPrice should not be below bestPrice * (1 - maxSlippage)
            uint256 minAcceptablePrice = (bestPrice * (10000 - maxSlippage)) / 10000;
            if (avgPrice < minAcceptablePrice) {
                revert SlippageExceeded();
            }
        }

        // Refund unused collateral for buy orders
        if (isBuy) {
            uint256 collateralUsed = (filledAmount * avgPrice) / PRICE_PRECISION;
            uint256 collateralToRefund = amount - collateralUsed;
            if (collateralToRefund > 0) {
                totalCollateral -= collateralToRefund;
                collateralToken.safeTransfer(msg.sender, collateralToRefund);
            }
        }

        // Clean up temporary order
        delete orders[orderId];

        return filledAmount;
    }

    /**
     * @notice Cancels an active order and refunds escrowed funds
     * @param orderId Order to cancel
     */
    function cancelOrder(bytes32 orderId) external nonReentrant {
        Order storage order = orders[orderId];
        
        if (!order.isActive) revert OrderNotFound();
        if (order.trader != msg.sender) revert OrderNotOwned();

        _cancelOrder(orderId);
    }

    // ============ View Functions ============

    /**
     * @notice Gets the best bid price for an outcome
     * @param outcomeId Outcome ID
     * @return Best bid price (0 if no bids)
     */
    function getBestBid(uint8 outcomeId) external view returns (uint256) {
        bytes32[] storage buyOrders = buyOrdersByOutcome[outcomeId];
        if (buyOrders.length == 0) return 0;
        return orders[buyOrders[0]].price;
    }

    /**
     * @notice Gets the best ask price for an outcome
     * @param outcomeId Outcome ID
     * @return Best ask price (0 if no asks)
     */
    function getBestAsk(uint8 outcomeId) external view returns (uint256) {
        bytes32[] storage sellOrders = sellOrdersByOutcome[outcomeId];
        if (sellOrders.length == 0) return 0;
        return orders[sellOrders[0]].price;
    }

    /**
     * @notice Gets the bid-ask spread for an outcome
     * @param outcomeId Outcome ID
     * @return spread The difference between best ask and best bid
     */
    function getSpread(uint8 outcomeId) external view returns (uint256 spread) {
        bytes32[] storage buyOrders = buyOrdersByOutcome[outcomeId];
        bytes32[] storage sellOrders = sellOrdersByOutcome[outcomeId];
        
        if (buyOrders.length == 0 || sellOrders.length == 0) return PRICE_PRECISION;
        
        uint256 bestBid = orders[buyOrders[0]].price;
        uint256 bestAsk = orders[sellOrders[0]].price;
        
        return bestAsk > bestBid ? bestAsk - bestBid : 0;
    }

    /**
     * @notice Gets all buy orders for an outcome
     * @param outcomeId Outcome ID
     * @return orderIds Array of buy order IDs
     */
    function getBuyOrders(uint8 outcomeId) external view returns (bytes32[] memory) {
        return buyOrdersByOutcome[outcomeId];
    }

    /**
     * @notice Gets all sell orders for an outcome
     * @param outcomeId Outcome ID
     * @return orderIds Array of sell order IDs
     */
    function getSellOrders(uint8 outcomeId) external view returns (bytes32[] memory) {
        return sellOrdersByOutcome[outcomeId];
    }

    /**
     * @notice Gets a user's active orders
     * @param user User address
     * @return orderIds Array of order IDs
     */
    function getUserOrders(address user) external view returns (bytes32[] memory) {
        return userOrders[user];
    }

    /**
     * @notice Gets order details
     * @param orderId Order ID
     * @return Order struct
     */
    function getOrder(bytes32 orderId) external view returns (Order memory) {
        return orders[orderId];
    }

    // ============ BaseMarket Implementation ============

    function getPrice(uint256) external pure override returns (uint256) {
        revert("Use getBestBid/getBestAsk for order book pricing");
    }

    function getQuoteBuy(uint256, uint256, address) external pure override returns (uint256, uint256) {
        revert("Use order book for quotes");
    }

    function getQuoteSell(uint256, uint256, address) external pure override returns (uint256, uint256) {
        revert("Use order book for quotes");
    }

    function buy(uint256, uint256, uint256) external pure override returns (uint256) {
        revert("Use placeOrder or placeMarketOrder");
    }

    function sell(uint256, uint256, uint256) external pure override returns (uint256) {
        revert("Use placeOrder or placeMarketOrder");
    }

    function addLiquidity(uint256) external pure override returns (uint256) {
        revert("Liquidity provided via limit orders");
    }

    function removeLiquidity(uint256) external pure override returns (uint256) {
        revert("Liquidity removed via order cancellation");
    }

    // ============ Internal Functions ============

    /**
     * @notice Matches an order against the order book
     * @param orderId Order to match
     * @return filledAmount Total amount filled
     */
    function _matchOrder(bytes32 orderId) internal returns (uint256 filledAmount) {
        Order storage order = orders[orderId];
        if (!order.isActive) return 0;

        bytes32[] storage oppositeOrders = order.isBuy
            ? sellOrdersByOutcome[order.outcomeId]
            : buyOrdersByOutcome[order.outcomeId];

        uint256 i = 0;
        while (i < oppositeOrders.length && order.filled < order.amount) {
            bytes32 matchOrderId = oppositeOrders[i];
            Order storage matchOrder = orders[matchOrderId];

            // Check if prices cross
            bool pricesCross = order.isBuy
                ? order.price >= matchOrder.price
                : order.price <= matchOrder.price;

            if (!pricesCross) break;

            // Calculate match amount
            uint256 remainingOrder = order.amount - order.filled;
            uint256 remainingMatch = matchOrder.amount - matchOrder.filled;
            uint256 matchAmount = remainingOrder < remainingMatch ? remainingOrder : remainingMatch;

            // Execute the trade at the maker's price
            uint256 tradePrice = matchOrder.price;
            _executeTrade(orderId, matchOrderId, matchAmount, tradePrice);

            filledAmount += matchAmount;

            // Update filled amounts
            order.filled += matchAmount;
            matchOrder.filled += matchAmount;

            // Mark as inactive if fully filled
            if (matchOrder.filled == matchOrder.amount) {
                matchOrder.isActive = false;
                // Remove from order book
                _removeFromOrderBook(matchOrderId, i, !order.isBuy);
            } else {
                i++;
                emit OrderPartiallyFilled(
                    matchOrderId,
                    matchAmount,
                    matchOrder.amount - matchOrder.filled
                );
            }
        }

        if (order.filled == order.amount) {
            order.isActive = false;
        }

        return filledAmount;
    }

    /**
     * @notice Executes a trade between two orders
     * @param buyOrderId Buy order ID
     * @param sellOrderId Sell order ID
     * @param amount Amount to trade
     * @param price Execution price
     */
    function _executeTrade(
        bytes32 buyOrderId,
        bytes32 sellOrderId,
        uint256 amount,
        uint256 price
    ) internal {
        Order storage buyOrder = orders[buyOrderId];
        Order storage sellOrder = orders[sellOrderId];

        // Determine buyer and seller
        address buyer = buyOrder.isBuy ? buyOrder.trader : sellOrder.trader;
        address seller = buyOrder.isBuy ? sellOrder.trader : buyOrder.trader;

        // Calculate collateral amount
        uint256 collateralAmount = (amount * price) / PRICE_PRECISION;

        // Track total cost for average price calculation (for market orders)
        buyOrder.totalCost += collateralAmount;
        sellOrder.totalCost += collateralAmount;

        // Apply trading fee
        (uint256 feeAmount, uint16 protocolBps) = _calculateFee(collateralAmount, buyer);
        uint256 amountAfterFee = collateralAmount - feeAmount;

        // Transfer outcome tokens to buyer
        outcomeToken.safeTransferFrom(
            address(this),
            buyer,
            outcomeToken.encodeTokenId(marketId, buyOrder.isBuy ? buyOrder.outcomeId : sellOrder.outcomeId),
            amount,
            ""
        );

        // Transfer collateral to seller
        totalCollateral -= collateralAmount;
        collateralToken.safeTransfer(seller, amountAfterFee);

        // Send fee to fee splitter
        _distributeFee(feeAmount, protocolBps);

        emit OrderMatched(
            buyOrder.isBuy ? buyOrderId : sellOrderId,
            buyOrder.isBuy ? sellOrderId : buyOrderId,
            buyOrder.isBuy ? buyOrder.outcomeId : sellOrder.outcomeId,
            price,
            amount,
            buyer,
            seller
        );
    }

    /**
     * @notice Adds an order to the order book maintaining sort order
     * @param orderId Order to add
     */
    function _addToOrderBook(bytes32 orderId) internal {
        Order storage order = orders[orderId];
        bytes32[] storage orderList = order.isBuy
            ? buyOrdersByOutcome[order.outcomeId]
            : sellOrdersByOutcome[order.outcomeId];

        if (orderList.length >= MAX_ORDERS_PER_SIDE) revert MaxOrdersReached();

        // Find insertion point (maintain sorted order)
        uint256 insertIndex = orderList.length;
        for (uint256 i = 0; i < orderList.length; i++) {
            Order storage existingOrder = orders[orderList[i]];
            bool shouldInsertHere = order.isBuy
                ? order.price > existingOrder.price || (order.price == existingOrder.price && order.timestamp < existingOrder.timestamp)
                : order.price < existingOrder.price || (order.price == existingOrder.price && order.timestamp < existingOrder.timestamp);

            if (shouldInsertHere) {
                insertIndex = i;
                break;
            }
        }

        // Insert at position
        orderList.push(orderId);
        for (uint256 i = orderList.length - 1; i > insertIndex; i--) {
            orderList[i] = orderList[i - 1];
        }
        orderList[insertIndex] = orderId;
    }

    /**
     * @notice Removes an order from the order book
     * @param orderId Order to remove
     * @param index Index in order list
     * @param isBuyBook True if removing from buy book
     */
    function _removeFromOrderBook(bytes32 orderId, uint256 index, bool isBuyBook) internal {
        Order storage order = orders[orderId];
        bytes32[] storage orderList = isBuyBook
            ? buyOrdersByOutcome[order.outcomeId]
            : sellOrdersByOutcome[order.outcomeId];

        // Remove by shifting
        for (uint256 i = index; i < orderList.length - 1; i++) {
            orderList[i] = orderList[i + 1];
        }
        orderList.pop();
    }

    /**
     * @notice Cancels an order and refunds escrowed funds
     * @param orderId Order to cancel
     */
    function _cancelOrder(bytes32 orderId) internal {
        Order storage order = orders[orderId];
        uint256 remainingAmount = order.amount - order.filled;

        // Refund escrowed funds
        if (order.isBuy) {
            uint256 collateralToRefund = (remainingAmount * order.price) / PRICE_PRECISION;
            totalCollateral -= collateralToRefund;
            collateralToken.safeTransfer(order.trader, collateralToRefund);
        } else {
            outcomeToken.safeTransferFrom(
                address(this),
                order.trader,
                outcomeToken.encodeTokenId(marketId, order.outcomeId),
                remainingAmount,
                ""
            );
        }

        // Mark as inactive
        order.isActive = false;

        emit OrderCancelled(orderId, order.trader, remainingAmount);

        // Remove from order book
        bytes32[] storage orderList = order.isBuy
            ? buyOrdersByOutcome[order.outcomeId]
            : sellOrdersByOutcome[order.outcomeId];

        for (uint256 i = 0; i < orderList.length; i++) {
            if (orderList[i] == orderId) {
                _removeFromOrderBook(orderId, i, order.isBuy);
                break;
            }
        }
    }

    /**
     * @notice Calculates average execution price for an order
     * @param orderId Order ID
     * @return Average price
     */
    function _calculateAveragePrice(bytes32 orderId) internal view returns (uint256) {
        Order storage order = orders[orderId];
        if (order.filled == 0) return 0;
        
        // Calculate weighted average price from total cost
        return (order.totalCost * PRICE_PRECISION) / order.filled;
    }
}
