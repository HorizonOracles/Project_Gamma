// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../base/BaseMarket.sol";
import {UD60x18, ud, intoUint256} from "prb-math/UD60x18.sol";

/**
 * @title MultiChoiceMarket
 * @notice Prediction market with 3-8 discrete outcomes using multi-way LMSR pricing
 * @dev Implements Logarithmic Market Scoring Rule (LMSR) for multiple outcomes
 *      
 *      ═══════════════════════════════════════════════════════════════════
 *      KEY FEATURES
 *      ═══════════════════════════════════════════════════════════════════
 *      - 3-8 discrete outcomes (e.g., "Which team wins?": A, B, C, D)
 *      - LMSR pricing ensures prices sum to 1.0
 *      - Bounded loss for liquidity providers (controlled by parameter b)
 *      - Efficient price discovery for multi-outcome scenarios
 * 
 *      ═══════════════════════════════════════════════════════════════════
 *      LMSR MATHEMATICS
 *      ═══════════════════════════════════════════════════════════════════
 *      
 *      Cost Function:
 *        C(q) = b × ln(Σᵢ exp(qᵢ / b))
 *        where qᵢ = quantity of outcome i, b = liquidity parameter
 *      
 *      Price Formula:
 *        pᵢ = exp(qᵢ / b) / Σⱼ exp(qⱼ / b)
 *      
 *      Trade Cost:
 *        Buying tokens:  cost = C(q') - C(q)  where q' has increased reserves
 *        Selling tokens: gain = C(q) - C(q')  where q' has decreased reserves
 *      
 *      Key Properties:
 *        1. Σ pᵢ = 1  (prices always sum to 1.0)
 *        2. Max LP loss = b × ln(n)  where n = number of outcomes
 *        3. Convex cost function (larger buys = higher marginal price)
 *        4. No arbitrage (round-trip trades always lose money due to fees)
 *      
 *      ═══════════════════════════════════════════════════════════════════
 *      IMPLEMENTATION DETAILS
 *      ═══════════════════════════════════════════════════════════════════
 *      
 *      Mathematical Library:
 *        - Uses PRBMath for accurate exp() and ln() calculations
 *        - Prevents precision loss and overflow in exponential operations
 *        - Critical for maintaining LMSR invariants across all trades
 *      
 *      Linearized Buy Approximation:
 *        - _calculateBuyTokens uses one-shot price estimation + adjustment
 *        - More gas-efficient than iterative binary search
 *        - Accurate for normal trade sizes relative to liquidity parameter
 *        - See function documentation for detailed algorithm
 *      
 *      Gas Optimization:
 *        - Caches storage variables (liquidityParameter, outcomeCount)
 *        - Uses unchecked arithmetic in bounded loops
 *        - Target: <150k gas per trade (currently ~248k for buy, ~108k for sell)
 *        - Higher gas cost is the tradeoff for mathematical accuracy
 *      
 *      ═══════════════════════════════════════════════════════════════════
 *      ECONOMIC TRADEOFFS
 *      ═══════════════════════════════════════════════════════════════════
 *      
 *      Liquidity Parameter (b):
 *        Higher b → More liquidity, lower price impact, higher LP risk
 *        Lower b  → Less liquidity, higher price impact, lower LP risk
 *      
 *      LP Risk Management:
 *        - Maximum loss is bounded: max_loss = b × ln(n)
 *        - For 4 outcomes: max_loss ≈ 1.386 × b
 *        - For 8 outcomes: max_loss ≈ 2.079 × b
 *        - LPs collect fees on all trades to compensate for risk
 *      
 *      ═══════════════════════════════════════════════════════════════════
 *      REFERENCES
 *      ═══════════════════════════════════════════════════════════════════
 *      - Hanson, R. (2003). "Combinatorial Information Market Design"
 *      - Chen & Pennock (2007). "A Utility Framework for Bounded-Loss 
 *        Market Makers"
 */
contract MultiChoiceMarket is BaseMarket {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    error InvalidOutcomeCount();
    error InvalidLiquidityParameter();
    error PriceCalculationOverflow();
    error InsufficientReserves();

    // ============ Events ============

    event OutcomeReservesUpdated(uint256[] reserves);
    event LiquidityParameterUpdated(uint256 oldB, uint256 newB);

    // ============ State Variables ============

    /// @notice Reserves for each outcome (q_i in LMSR formula)
    uint256[] public outcomeReserves;

    /// @notice Liquidity parameter (b in LMSR formula) - controls market depth
    /// @dev Higher b = more liquidity, lower price impact, higher LP exposure
    ///      Lower b = less liquidity, higher price impact, lower LP exposure
    ///      Max LP loss = b × ln(outcomeCount)
    ///      Typical values: 1000-10000 tokens for standard markets
    uint256 public liquidityParameter;

    /// @notice Scale factor for LMSR calculations to prevent underflow
    /// We scale reserves up before division to maintain precision
    uint256 private constant LMSR_SCALE = 1e18;

    // ============ Constructor ============

    /**
     * @notice Initializes a multi-choice market
     * @dev Constructor sets up the LMSR market with specified parameters
     * 
     * @param _marketId Market identifier
     * @param _collateralToken Collateral token address (e.g., USDC, DAI)
     * @param _outcomeToken Outcome token contract (ERC1155)
     * @param _feeSplitter Fee splitter contract for distributing fees
     * @param _horizonPerks Horizon perks contract for fee discounts
     * @param _closeTime Market close timestamp (no trades after this)
     * @param _outcomeCount Number of outcomes (3-8)
     *        Examples:
     *        - 3: Low/Medium/High
     *        - 4: Q1/Q2/Q3/Q4
     *        - 8: Team A/B/C/D/E/F/G/H
     * @param _liquidityParameter Liquidity parameter (b in LMSR formula)
     *        Controls market depth and LP risk:
     *        - Higher b: deeper liquidity, lower price impact, higher LP exposure
     *        - Lower b: shallower liquidity, higher price impact, lower LP exposure
     *        - Recommended: 1000-10000 for typical markets
     *        - Max LP loss: b × ln(_outcomeCount)
     */
    constructor(
        uint256 _marketId,
        address _collateralToken,
        address _outcomeToken,
        address _feeSplitter,
        address _horizonPerks,
        uint256 _closeTime,
        uint256 _outcomeCount,
        uint256 _liquidityParameter
    )
        BaseMarket(
            _marketId,
            MarketType.MultiChoice,
            _collateralToken,
            _outcomeToken,
            _feeSplitter,
            _horizonPerks,
            _closeTime,
            _outcomeCount,
            "Multi-Choice LP Token",
            "MC-LP"
        )
    {
        if (_outcomeCount < 3 || _outcomeCount > 8) revert InvalidOutcomeCount();
        if (_liquidityParameter == 0) revert InvalidLiquidityParameter();

        liquidityParameter = _liquidityParameter;

        // Initialize outcome reserves to zero
        outcomeReserves = new uint256[](_outcomeCount);
    }

    // ============ Liquidity Functions ============

    /**
     * @notice Adds liquidity to the market
     * @param amount Amount of collateral to add
     * @return lpTokens Amount of LP tokens minted
     */
    function addLiquidity(uint256 amount)
        external
        override
        nonReentrant
        whenNotPaused
        notResolved
        returns (uint256 lpTokens)
    {
        _validateAmount(amount);

        // Transfer collateral from user
        collateralToken.safeTransferFrom(msg.sender, address(this), amount);

        uint256 totalSupply = totalSupply();

        if (totalSupply == 0) {
            // Initial liquidity provision
            // Initialize all outcomes with equal quantities
            // Note: outcomeReserves track LMSR quantities (q_i), representing shares sold to traders
            // Initially q_i = 0 for all i (no sales yet), but we need starting capital
            // We initialize with equal amounts to set initial equal prices
            uint256 initialReserve = amount / outcomeCount;

            for (uint256 i = 0; i < outcomeCount; i++) {
                // Set initial LMSR state (this determines initial prices)
                outcomeReserves[i] = initialReserve;
                // Mint tokens to market as inventory for future sales
                _mintOutcome(i, address(this), initialReserve);
            }

            totalCollateral = amount;

            // Mint LP tokens (lock minimum liquidity)
            lpTokens = amount - MINIMUM_LIQUIDITY;
            _mint(address(0xdead), MINIMUM_LIQUIDITY);
            _mint(msg.sender, lpTokens);

            emit LiquidityChanged(msg.sender, amount, true);
            emit OutcomeReservesUpdated(outcomeReserves);
        } else {
            // Subsequent liquidity provision
            // Mint proportional to existing liquidity
            lpTokens = (amount * totalSupply) / totalCollateral;

            // Mint outcome tokens proportional to current market holdings
            // This provides inventory for future sales WITHOUT changing LMSR state
            // The LMSR state (outcomeReserves) only changes during buy/sell trades
            for (uint256 i = 0; i < outcomeCount; i++) {
                uint256 tokenId = outcomeToken.encodeTokenId(marketId, i);
                uint256 marketBalance = outcomeToken.balanceOf(address(this), tokenId);
                
                // Mint proportional to existing holdings
                uint256 additionalTokens = (amount * marketBalance) / totalCollateral;
                
                if (additionalTokens > 0) {
                    _mintOutcome(i, address(this), additionalTokens);
                }
                
                // NOTE: We do NOT update outcomeReserves here!
                // outcomeReserves tracks LMSR quantities (cumulative sales to traders)
                // Adding LP inventory shouldn't change prices or LMSR state
            }

            totalCollateral += amount;
            _mint(msg.sender, lpTokens);

            emit LiquidityChanged(msg.sender, amount, true);
            // Don't emit OutcomeReservesUpdated since reserves didn't change
        }
    }

    /**
     * @notice Removes liquidity from the market
     * @param lpTokens Amount of LP tokens to burn
     * @return collateralOut Amount of collateral returned
     */
    function removeLiquidity(uint256 lpTokens)
        external
        override
        nonReentrant
        returns (uint256 collateralOut)
    {
        _validateAmount(lpTokens);

        uint256 totalSupply = totalSupply();
        if (lpTokens > balanceOf(msg.sender)) revert InsufficientLPTokens();

        // Calculate proportional collateral
        collateralOut = (lpTokens * totalCollateral) / totalSupply;

        // Burn proportional outcome tokens based on actual market holdings (not outcomeReserves)
        // outcomeReserves tracks LMSR quantities (shares sold to traders), but we need to burn
        // based on actual ERC1155 balances held by the market contract
        for (uint256 i = 0; i < outcomeCount; i++) {
            // Get the market's actual ERC1155 balance for this outcome
            uint256 tokenId = outcomeToken.encodeTokenId(marketId, i);
            uint256 marketBalance = outcomeToken.balanceOf(address(this), tokenId);
            
            // Burn proportional to LP tokens being redeemed
            uint256 tokensToBurn = (lpTokens * marketBalance) / totalSupply;
            
            if (tokensToBurn > 0) {
                _burnOutcome(i, address(this), tokensToBurn);
            }
        }

        totalCollateral -= collateralOut;
        _burn(msg.sender, lpTokens);

        // Transfer collateral to user
        collateralToken.safeTransfer(msg.sender, collateralOut);

        emit LiquidityChanged(msg.sender, collateralOut, false);
        emit OutcomeReservesUpdated(outcomeReserves);
    }

    // ============ Trading Functions ============

    /**
     * @notice Buys outcome tokens
     * @param outcomeId ID of outcome to buy
     * @param collateralIn Amount of collateral to spend
     * @param minTokensOut Minimum tokens expected (slippage protection)
     * @return tokensOut Amount of outcome tokens received
     */
    function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut)
        external
        override
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        validOutcome(outcomeId)
        returns (uint256 tokensOut)
    {
        _validateAmount(collateralIn);
        if (totalCollateral == 0) revert InsufficientLiquidity();

        // Calculate fee
        (uint256 fee, uint16 protocolBps) = _calculateFee(collateralIn, msg.sender);
        uint256 collateralAfterFee = collateralIn - fee;

        // Transfer collateral from user
        collateralToken.safeTransferFrom(msg.sender, address(this), collateralIn);

        // Distribute fee
        _distributeFee(fee, protocolBps);

        // Calculate tokens out using LMSR
        tokensOut = _calculateBuyTokens(outcomeId, collateralAfterFee);

        // Slippage check
        _validateSlippage(tokensOut, minTokensOut);

        // Update reserve for purchased outcome (increase quantity in circulation)
        outcomeReserves[outcomeId] += tokensOut;

        // Mint and transfer outcome tokens to user
        _mintOutcome(outcomeId, msg.sender, tokensOut);

        totalCollateral += collateralAfterFee;

        emit Trade(msg.sender, outcomeId, collateralIn, tokensOut, fee, true);
        emit OutcomeReservesUpdated(outcomeReserves);
    }

    /**
     * @notice Sells outcome tokens
     * @param outcomeId ID of outcome to sell
     * @param tokensIn Amount of tokens to sell
     * @param minCollateralOut Minimum collateral expected (slippage protection)
     * @return collateralOut Amount of collateral received
     */
    function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut)
        external
        override
        nonReentrant
        whenNotPaused
        beforeClose
        notResolved
        validOutcome(outcomeId)
        returns (uint256 collateralOut)
    {
        _validateAmount(tokensIn);
        if (totalCollateral == 0) revert InsufficientLiquidity();

        // Calculate collateral out using LMSR
        collateralOut = _calculateSellCollateral(outcomeId, tokensIn);

        // Calculate fee
        (uint256 fee, uint16 protocolBps) = _calculateFee(collateralOut, msg.sender);
        collateralOut -= fee;

        // Slippage check
        _validateSlippage(collateralOut, minCollateralOut);

        // Burn tokens from user
        _burnOutcome(outcomeId, msg.sender, tokensIn);

        // Update reserve for sold outcome (decrease quantity in circulation)
        if (outcomeReserves[outcomeId] < tokensIn) revert InsufficientReserves();
        outcomeReserves[outcomeId] -= tokensIn;

        totalCollateral -= collateralOut;

        // Distribute fee
        _distributeFee(fee, protocolBps);

        // Transfer collateral to user (after fee)
        collateralToken.safeTransfer(msg.sender, collateralOut);

        emit Trade(msg.sender, outcomeId, tokensIn, collateralOut, fee, false);
        emit OutcomeReservesUpdated(outcomeReserves);
    }

    // ============ Price & Quote Functions ============

    /**
     * @notice Gets current price for an outcome using LMSR
     * @param outcomeId Outcome identifier
     * @return price Price in 1e18 precision (e.g., 0.25e18 = 25%)
     */
    function getPrice(uint256 outcomeId)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 price)
    {
        if (totalCollateral == 0) return PRICE_PRECISION / outcomeCount; // Equal prices if no liquidity

        price = _calculatePrice(outcomeId);
    }

    /**
     * @notice Gets quote for buying outcome tokens
     * @param outcomeId Outcome to buy
     * @param collateralIn Amount of collateral to spend
     * @param user User address (for fee tier calculation)
     * @return tokensOut Expected tokens received
     * @return fee Fee amount
     */
    function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 tokensOut, uint256 fee)
    {
        if (collateralIn == 0 || totalCollateral == 0) return (0, 0);

        (fee,) = _calculateFee(collateralIn, user);
        uint256 collateralAfterFee = collateralIn - fee;

        tokensOut = _calculateBuyTokens(outcomeId, collateralAfterFee);
    }

    /**
     * @notice Gets quote for selling outcome tokens
     * @param outcomeId Outcome to sell
     * @param tokensIn Amount of tokens to sell
     * @param user User address (for fee tier calculation)
     * @return collateralOut Expected collateral received
     * @return fee Fee amount
     */
    function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user)
        external
        view
        override
        validOutcome(outcomeId)
        returns (uint256 collateralOut, uint256 fee)
    {
        if (tokensIn == 0 || totalCollateral == 0) return (0, 0);

        collateralOut = _calculateSellCollateral(outcomeId, tokensIn);
        (fee,) = _calculateFee(collateralOut, user);
        collateralOut -= fee;
    }

    // ============ Internal LMSR Calculation Functions ============

    /**
     * @notice Calculates price for an outcome using LMSR formula
     * @dev Price formula: pᵢ = exp(qᵢ / b) / Σⱼ exp(qⱼ / b)
     * 
     *      This implements the standard LMSR price calculation where each
     *      outcome's price is proportional to exp(qᵢ / b). The prices are
     *      normalized so they sum to 1.0, reflecting probability distribution.
     * 
     *      Uses PRBMath for accurate exponential calculations to prevent
     *      precision loss that could violate LMSR invariants.
     * 
     * @param outcomeId Outcome to price
     * @return price Price in 1e18 precision (e.g., 0.25e18 = 25% probability)
     */
    function _calculatePrice(uint256 outcomeId) internal view returns (uint256 price) {
        // Cache liquidityParameter to save SLOAD operations
        uint256 b = liquidityParameter;
        
        // Calculate exp(q_i / b) for the target outcome
        UD60x18 scaledRatio = ud((outcomeReserves[outcomeId] * LMSR_SCALE) / b);
        UD60x18 numerator = scaledRatio.exp();
        
        // Calculate Σ exp(q_j / b) for all outcomes
        UD60x18 denominator = ud(0);
        uint256 count = outcomeCount; // Cache to save SLOAD
        for (uint256 i = 0; i < count;) {
            UD60x18 iScaledRatio = ud((outcomeReserves[i] * LMSR_SCALE) / b);
            denominator = denominator.add(iScaledRatio.exp());
            unchecked { ++i; } // Safe: i < outcomeCount (max 8)
        }
        
        if (denominator.isZero()) revert PriceCalculationOverflow();
        
        // Price = numerator / denominator
        price = numerator.div(denominator).mul(ud(PRICE_PRECISION)).intoUint256();
    }

    /**
     * @notice Calculates the LMSR cost function C(q) = b × ln(Σ exp(qᵢ / b))
     * @dev The cost function represents the total cost to reach a given state.
     *      Trade costs are calculated as the difference: C(q') - C(q)
     * 
     *      Key properties:
     *      - Strictly increasing (more tokens = higher cost)
     *      - Convex (marginal cost increases with quantity)
     *      - Smooth and differentiable (enables efficient pricing)
     * 
     *      Uses PRBMath for accurate exp() and ln() to maintain precision
     *      across all possible reserve states.
     * 
     * @param reserves Array of outcome quantities (q₀, q₁, ..., qₙ₋₁)
     * @return cost The cost function value in collateral units
     */
    function _calculateCostFunction(uint256[] memory reserves) internal view returns (UD60x18 cost) {
        // Cache liquidityParameter to save SLOAD operations
        uint256 b = liquidityParameter;
        
        // Calculate Σ exp(q_i / b)
        UD60x18 sumExp = ud(0);
        uint256 length = reserves.length;
        
        for (uint256 i = 0; i < length;) {
            UD60x18 scaledRatio = ud((reserves[i] * LMSR_SCALE) / b);
            sumExp = sumExp.add(scaledRatio.exp());
            unchecked { ++i; } // Safe: i < reserves.length (max 8)
        }
        
        // C(q) = b * ln(Σ exp(q_i / b))
        // Convert b to UD60x18, multiply by ln(sumExp)
        UD60x18 bUD = ud(b);
        cost = bUD.mul(sumExp.ln()).div(ud(LMSR_SCALE));
    }

    /**
     * @notice Calculates tokens received when buying
     * @dev Uses proper LMSR with linearized approximation for efficiency
     *      
     *      Algorithm:
     *      1. Start with linear estimate: tokens ≈ collateral / current_price
     *      2. Check actual LMSR cost with those tokens
     *      3. Adjust proportionally if needed (one-shot correction)
     *      
     *      This linearized approach is more gas-efficient than binary search
     *      and accurate for normal trade sizes relative to liquidity parameter.
     *      
     * @param outcomeId Outcome being bought
     * @param collateralAmount Amount of collateral spent (after fees)
     * @return tokensOut Tokens received
     */
    function _calculateBuyTokens(uint256 outcomeId, uint256 collateralAmount)
        internal
        view
        returns (uint256 tokensOut)
    {
        // Create a copy of current reserves
        uint256 count = outcomeCount; // Cache to save SLOAD
        uint256[] memory newReserves = new uint256[](count);
        for (uint256 i = 0; i < count;) {
            newReserves[i] = outcomeReserves[i];
            unchecked { ++i; } // Safe: i < outcomeCount (max 8)
        }
        
        // Calculate current cost
        UD60x18 costBefore = _calculateCostFunction(outcomeReserves);
        
        // Start with an estimate based on current price
        uint256 currentPrice = _calculatePrice(outcomeId);
        if (currentPrice == 0) revert PriceCalculationOverflow();
        
        uint256 estimatedTokens = (collateralAmount * PRICE_PRECISION) / currentPrice;
        
        // Use the estimate as starting point, check if it's close enough
        newReserves[outcomeId] = outcomeReserves[outcomeId] + estimatedTokens;
        UD60x18 costAfter = _calculateCostFunction(newReserves);
        uint256 actualCost = costAfter.sub(costBefore).intoUint256();
        
        // Simple adjustment: if actual cost is too high/low, scale proportionally
        // This is a linearized approximation that works well for small trades
        if (actualCost > collateralAmount) {
            // We overshot, reduce tokens proportionally
            estimatedTokens = (estimatedTokens * collateralAmount) / actualCost;
        } else if (actualCost < collateralAmount) {
            // We undershot, increase tokens proportionally
            uint256 ratio = (collateralAmount * PRICE_PRECISION) / actualCost;
            if (ratio > PRICE_PRECISION * 2) {
                ratio = PRICE_PRECISION * 2; // Cap at 2x to prevent overshooting
            }
            estimatedTokens = (estimatedTokens * ratio) / PRICE_PRECISION;
        }
        
        tokensOut = estimatedTokens;
    }

    /**
     * @notice Calculates collateral received when selling
     * @dev Uses proper LMSR: cost difference between states determines collateral received
     *      Selling reduces the cost function, and the difference is returned to the seller
     * @param outcomeId Outcome being sold
     * @param tokensAmount Amount of tokens sold
     * @return collateralOut Collateral received (before fees)
     */
    function _calculateSellCollateral(uint256 outcomeId, uint256 tokensAmount)
        internal
        view
        returns (uint256 collateralOut)
    {
        // Create a copy of current reserves
        uint256 count = outcomeCount; // Cache to save SLOAD
        uint256[] memory newReserves = new uint256[](count);
        for (uint256 i = 0; i < count;) {
            newReserves[i] = outcomeReserves[i];
            unchecked { ++i; } // Safe: i < outcomeCount (max 8)
        }
        
        // Calculate cost before selling
        UD60x18 costBefore = _calculateCostFunction(outcomeReserves);
        
        // Calculate cost after selling (reduce reserves for this outcome)
        if (outcomeReserves[outcomeId] < tokensAmount) revert InsufficientReserves();
        newReserves[outcomeId] = outcomeReserves[outcomeId] - tokensAmount;
        UD60x18 costAfter = _calculateCostFunction(newReserves);
        
        // Collateral received = cost reduction
        collateralOut = costBefore.sub(costAfter).intoUint256();
    }

    // ============ View Functions ============

    /**
     * @notice Gets all outcome reserves
     * @return Array of reserves for each outcome
     */
    function getOutcomeReserves() external view returns (uint256[] memory) {
        return outcomeReserves;
    }

    /**
     * @notice Gets all current prices
     * @return prices Array of prices for each outcome
     */
    function getAllPrices() external view returns (uint256[] memory prices) {
        prices = new uint256[](outcomeCount);
        for (uint256 i = 0; i < outcomeCount; i++) {
            prices[i] = _calculatePrice(i);
        }
    }
}
