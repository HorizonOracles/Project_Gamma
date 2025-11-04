/**
 * LMSR Pricing utilities for multi-choice markets
 *
 * LMSR (Logarithmic Market Scoring Rule) is used for markets with 3-8 outcomes.
 * The contract handles the complex LMSR math on-chain, but these utilities help
 * with client-side calculations, conversions, and validations.
 */

import { parseUnits } from 'viem';
import { TradeError } from '../types';

/**
 * Basis points scale (10000 = 100%)
 */
export const BASIS_POINTS_SCALE = 10000n;

/**
 * Price precision (18 decimals)
 */
export const PRICE_PRECISION = 18;

/**
 * Convert basis points to percentage (e.g., 2500 -> 25.00%)
 */
export function basisPointsToPercent(bps: bigint): number {
  return Number(bps) / 100;
}

/**
 * Convert percentage to basis points (e.g., 25.5 -> 2550)
 */
export function percentToBasisPoints(percent: number): bigint {
  if (percent < 0 || percent > 100) {
    throw new TradeError(`Invalid percentage: ${percent}. Must be between 0 and 100.`);
  }
  return BigInt(Math.round(percent * 100));
}

/**
 * Convert basis points price to decimal probability (e.g., 2500 -> 0.25)
 */
export function basisPointsToProbability(bps: bigint): number {
  return Number(bps) / 10000;
}

/**
 * Convert decimal probability to basis points (e.g., 0.25 -> 2500)
 */
export function probabilityToBasisPoints(probability: number): bigint {
  if (probability < 0 || probability > 1) {
    throw new TradeError(`Invalid probability: ${probability}. Must be between 0 and 1.`);
  }
  return BigInt(Math.round(probability * 10000));
}

/**
 * Format price in basis points to human-readable percentage string
 * @param bps - Price in basis points
 * @param decimals - Number of decimal places (default: 2)
 * @returns Formatted percentage string (e.g., "25.50%")
 */
export function formatPriceBps(bps: bigint, decimals: number = 2): string {
  const percent = basisPointsToPercent(bps);
  return `${percent.toFixed(decimals)}%`;
}

/**
 * Validate outcome ID is within valid range for multi-choice markets (0-7)
 */
export function validateOutcomeId(outcomeId: bigint | number, outcomeCount: bigint): void {
  const id = typeof outcomeId === 'number' ? BigInt(outcomeId) : outcomeId;

  if (id < 0n) {
    throw new TradeError(`Invalid outcome ID: ${id}. Must be non-negative.`);
  }

  if (id >= outcomeCount) {
    throw new TradeError(
      `Invalid outcome ID: ${id}. Must be less than outcome count (${outcomeCount}).`
    );
  }
}

/**
 * Validate outcome count is within valid range (3-8 for multi-choice)
 */
export function validateOutcomeCount(count: bigint | number): void {
  const c = typeof count === 'number' ? BigInt(count) : count;

  if (c < 3n || c > 8n) {
    throw new TradeError(
      `Invalid outcome count: ${c}. Multi-choice markets must have 3-8 outcomes.`
    );
  }
}

/**
 * Check if prices sum to approximately 100% (10000 basis points)
 * Allows for small rounding errors
 */
export function validatePriceSum(prices: bigint[], tolerance: bigint = 10n): boolean {
  const sum = prices.reduce((acc, price) => acc + price, 0n);
  const diff = sum > BASIS_POINTS_SCALE
    ? sum - BASIS_POINTS_SCALE
    : BASIS_POINTS_SCALE - sum;

  return diff <= tolerance;
}

/**
 * Get the favorite outcome (highest price)
 * Returns the index of the outcome with the highest price
 */
export function getFavoriteOutcome(prices: bigint[]): number {
  if (prices.length === 0) {
    throw new TradeError('Cannot determine favorite from empty prices array');
  }

  let maxPrice = prices[0];
  let maxIndex = 0;

  for (let i = 1; i < prices.length; i++) {
    if (prices[i] > maxPrice) {
      maxPrice = prices[i];
      maxIndex = i;
    }
  }

  return maxIndex;
}

/**
 * Get the underdog outcome (lowest price)
 * Returns the index of the outcome with the lowest price
 */
export function getUnderdogOutcome(prices: bigint[]): number {
  if (prices.length === 0) {
    throw new TradeError('Cannot determine underdog from empty prices array');
  }

  let minPrice = prices[0];
  let minIndex = 0;

  for (let i = 1; i < prices.length; i++) {
    if (prices[i] < minPrice) {
      minPrice = prices[i];
      minIndex = i;
    }
  }

  return minIndex;
}

/**
 * Calculate implied odds from price
 * @param bps - Price in basis points
 * @returns Decimal odds (e.g., 2500 bps (25%) -> 4.0 odds)
 */
export function calculateImpliedOdds(bps: bigint): number {
  if (bps === 0n) {
    return Infinity;
  }

  const probability = basisPointsToProbability(bps);
  return 1 / probability;
}

/**
 * Calculate potential profit for a buy
 * @param collateralIn - Amount of collateral to spend
 * @param tokensOut - Amount of outcome tokens received
 * @param priceBps - Current price in basis points
 * @returns Potential profit if outcome wins (in collateral units)
 */
export function calculatePotentialProfit(
  collateralIn: bigint,
  tokensOut: bigint,
  priceBps: bigint
): bigint {
  // If outcome wins, each token is worth 1 collateral
  // Profit = tokens received - collateral spent
  return tokensOut - collateralIn;
}

/**
 * Calculate ROI (Return on Investment) percentage
 * @param collateralIn - Amount of collateral invested
 * @param tokensOut - Amount of outcome tokens received
 * @returns ROI as a percentage (e.g., 50.5 = 50.5% return)
 */
export function calculateROI(collateralIn: bigint, tokensOut: bigint): number {
  if (collateralIn === 0n) {
    return 0;
  }

  const profit = tokensOut - collateralIn;
  const roiBps = (profit * 10000n) / collateralIn;

  return Number(roiBps) / 100;
}

/**
 * Calculate price impact of a trade
 * @param priceBefore - Price before trade (in basis points)
 * @param priceAfter - Price after trade (in basis points)
 * @returns Price impact as a percentage
 */
export function calculatePriceImpact(priceBefore: bigint, priceAfter: bigint): number {
  if (priceBefore === 0n) {
    return 0;
  }

  const diff = priceAfter > priceBefore
    ? priceAfter - priceBefore
    : priceBefore - priceAfter;

  const impactBps = (diff * 10000n) / priceBefore;
  return Number(impactBps) / 100;
}

/**
 * Format multiple prices as distribution
 * Returns array of outcome probabilities that sum to ~100%
 */
export function formatPriceDistribution(
  prices: bigint[],
  outcomeLabels?: string[]
): Array<{ outcome: string; probability: number; bps: bigint }> {
  if (outcomeLabels && outcomeLabels.length !== prices.length) {
    throw new TradeError(
      `Mismatch between prices length (${prices.length}) and labels length (${outcomeLabels.length})`
    );
  }

  return prices.map((bps, index) => ({
    outcome: outcomeLabels?.[index] || `Outcome ${index}`,
    probability: basisPointsToProbability(bps),
    bps,
  }));
}

/**
 * Sort outcomes by price (highest to lowest)
 */
export function sortOutcomesByPrice(
  prices: bigint[],
  outcomeLabels?: string[]
): Array<{ index: number; outcome: string; price: bigint }> {
  const outcomes = prices.map((price, index) => ({
    index,
    outcome: outcomeLabels?.[index] || `Outcome ${index}`,
    price,
  }));

  return outcomes.sort((a, b) => {
    // Sort descending (highest price first)
    if (a.price > b.price) return -1;
    if (a.price < b.price) return 1;
    return 0;
  });
}

/**
 * Calculate effective price per token
 * @param collateralIn - Total collateral spent
 * @param tokensOut - Tokens received
 * @returns Average price per token in collateral units (18 decimals)
 */
export function calculateEffectivePrice(
  collateralIn: bigint,
  tokensOut: bigint
): bigint {
  if (tokensOut === 0n) {
    return 0n;
  }

  // Return price with 18 decimal precision
  return (collateralIn * parseUnits('1', 18)) / tokensOut;
}

/**
 * Calculate minimum amount out with slippage tolerance for multi-choice
 * @param expectedAmount - Expected amount out
 * @param slippageBps - Slippage tolerance in basis points (e.g., 100 = 1%)
 * @returns Minimum acceptable amount
 */
export function calculateMinAmountOut(
  expectedAmount: bigint,
  slippageBps: number
): bigint {
  if (slippageBps < 0 || slippageBps > 10000) {
    throw new TradeError(`Invalid slippage: ${slippageBps}. Must be between 0 and 10000 bps.`);
  }

  const slippageMultiplier = BigInt(10000 - slippageBps);
  return (expectedAmount * slippageMultiplier) / 10000n;
}

/**
 * Estimate liquidity depth by checking how much the price changes
 * Higher percentage change indicates lower liquidity
 * @param initialPrice - Price before trade
 * @param finalPrice - Price after trade
 * @returns Liquidity depth indicator (0-100, higher is better liquidity)
 */
export function estimateLiquidityDepth(
  initialPrice: bigint,
  finalPrice: bigint
): number {
  const priceImpact = calculatePriceImpact(initialPrice, finalPrice);

  // Invert price impact to get liquidity score
  // 0% impact = 100 score, 10%+ impact = 0 score
  const liquidityScore = Math.max(0, Math.min(100, 100 - priceImpact * 10));

  return liquidityScore;
}

/**
 * Check if a market is balanced (no outcome dominates too heavily)
 * @param prices - Array of outcome prices
 * @param maxDominancePercent - Maximum acceptable price for single outcome (default: 90%)
 * @returns True if market is balanced
 */
export function isMarketBalanced(
  prices: bigint[],
  maxDominancePercent: number = 90
): boolean {
  const maxDominanceBps = percentToBasisPoints(maxDominancePercent);

  for (const price of prices) {
    if (price > maxDominanceBps) {
      return false;
    }
  }

  return true;
}

/**
 * Get outcome token ID from market ID and outcome index
 * For ERC1155 token ID encoding
 * Token ID = (marketId << 8) | outcomeId
 */
export function getMultiChoiceTokenId(
  marketId: bigint,
  outcomeId: bigint | number
): bigint {
  const id = typeof outcomeId === 'number' ? BigInt(outcomeId) : outcomeId;
  return (marketId << 8n) | id;
}

/**
 * Extract outcome ID from token ID
 */
export function getOutcomeIdFromTokenId(tokenId: bigint): bigint {
  return tokenId & 0xFFn; // Lower 8 bits
}

/**
 * Extract market ID from token ID
 */
export function getMarketIdFromTokenId(tokenId: bigint): bigint {
  return tokenId >> 8n; // Upper bits
}
