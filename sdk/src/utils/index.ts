/**
 * Utility functions for the SDK
 */

import { Address, formatUnits, parseUnits } from 'viem';
import { MarketOutcome, MarketPrices, TradeError } from '../types';

/**
 * Format a bigint value to a human-readable string
 */
export function formatTokenAmount(
  amount: bigint,
  decimals: number = 18,
  precision: number = 4
): string {
  // Handle zero case
  if (amount === 0n) {
    return '0';
  }
  
  const formatted = formatUnits(amount, decimals);
  const [integer, decimal] = formatted.split('.');
  
  if (!decimal) {
    // If no decimal part and precision is 0, return just the integer
    if (precision === 0) {
      return integer;
    }
    // Pad with zeros if precision is specified
    return `${integer}.${'0'.repeat(precision)}`;
  }
  
  // If precision is 0, return just the integer part
  if (precision === 0) {
    return integer;
  }
  
  const truncatedDecimal = decimal.slice(0, precision).padEnd(precision, '0');
  return `${integer}.${truncatedDecimal}`;
}

/**
 * Parse a human-readable string to bigint
 */
export function parseTokenAmount(
  amount: string,
  decimals: number = 18
): bigint {
  // Trim whitespace
  const trimmed = amount.trim();
  
  if (!trimmed || trimmed === '') {
    throw new TradeError(`Invalid amount format: ${amount}`);
  }
  
  // Check for negative numbers
  if (trimmed.startsWith('-')) {
    throw new TradeError(`Invalid amount format: negative numbers not allowed: ${amount}`);
  }
  
  // Check for comma separators (not supported)
  if (trimmed.includes(',')) {
    throw new TradeError(`Invalid amount format: comma separators not supported: ${amount}`);
  }
  
  try {
    // parseUnits will throw if the number is too large or invalid
    const result = parseUnits(trimmed, decimals);
    
    // Additional check for very large numbers that might have passed parseUnits
    // This is a safety check for numbers that exceed practical limits
    if (result > 2n ** 128n) {
      throw new TradeError(`Invalid amount format: amount too large: ${amount}`);
    }
    
    return result;
  } catch (error) {
    if (error instanceof TradeError) {
      throw error;
    }
    throw new TradeError(`Invalid amount format: ${amount}`, undefined, error);
  }
}

/**
 * Calculate the price of an outcome token
 * Uses constant product formula: price = (liquidityOther / liquidityOutcome)
 */
export function calculatePrice(
  outcomeLiquidity: bigint,
  otherLiquidity: bigint
): bigint {
  if (outcomeLiquidity === 0n || otherLiquidity === 0n) {
    return 0n;
  }
  
  // Price = (otherLiquidity * 1e18) / (outcomeLiquidity + otherLiquidity)
  const totalLiquidity = outcomeLiquidity + otherLiquidity;
  return (otherLiquidity * parseUnits('1', 18)) / totalLiquidity;
}

/**
 * Calculate market prices for both outcomes
 */
export function calculateMarketPrices(
  yesLiquidity: bigint,
  noLiquidity: bigint
): MarketPrices {
  return {
    yesPrice: calculatePrice(yesLiquidity, noLiquidity),
    noPrice: calculatePrice(noLiquidity, yesLiquidity),
  };
}

/**
 * Calculate amount out using constant product formula
 * amountOut = (amountIn * otherLiquidity) / (outcomeLiquidity + amountIn)
 */
export function calculateAmountOut(
  amountIn: bigint,
  outcomeLiquidity: bigint,
  otherLiquidity: bigint
): bigint {
  if (outcomeLiquidity === 0n || otherLiquidity === 0n) {
    return 0n;
  }
  
  const numerator = amountIn * otherLiquidity;
  const denominator = outcomeLiquidity + amountIn;
  
  return numerator / denominator;
}

/**
 * Calculate amount in required to get desired amount out
 * amountIn = (amountOut * outcomeLiquidity) / (otherLiquidity - amountOut)
 */
export function calculateAmountIn(
  amountOut: bigint,
  outcomeLiquidity: bigint,
  otherLiquidity: bigint
): bigint {
  if (otherLiquidity <= amountOut) {
    throw new TradeError('Insufficient liquidity for desired amount out');
  }
  
  const numerator = amountOut * outcomeLiquidity;
  const denominator = otherLiquidity - amountOut;
  
  return numerator / denominator;
}

/**
 * Calculate slippage percentage
 */
export function calculateSlippage(
  expectedAmount: bigint,
  actualAmount: bigint
): number {
  if (expectedAmount === 0n) {
    return 0;
  }
  
  const diff = expectedAmount - actualAmount;
  return Number((diff * 10000n) / expectedAmount) / 100; // Convert to percentage
}

/**
 * Validate address format
 */
export function isValidAddress(address: string): boolean {
  try {
    return /^0x[a-fA-F0-9]{40}$/.test(address);
  } catch {
    return false;
  }
}

/**
 * Shorten address for display
 */
export function shortenAddress(address: Address, chars: number = 4): string {
  return `${address.slice(0, chars + 2)}...${address.slice(-chars)}`;
}

/**
 * Get outcome token ID from market ID and outcome
 * Token ID encoding: (marketId << 8) | outcomeId
 * outcomeId: 0 = YES, 1 = NO
 */
export function getOutcomeTokenId(
  marketId: bigint,
  outcome: MarketOutcome
): bigint {
  const outcomeId = outcome === 'YES' ? 0n : 1n;
  return (marketId << 8n) | outcomeId;
}

/**
 * Get market ID from token ID
 */
export function getMarketIdFromTokenId(tokenId: bigint): bigint {
  return tokenId >> 8n;
}

/**
 * Get outcome from token ID
 */
export function getOutcomeFromTokenId(tokenId: bigint): MarketOutcome {
  const outcomeId = tokenId & 0xFFn; // Lower 8 bits
  return outcomeId === 0n ? 'YES' : 'NO';
}

/**
 * Calculate minimum amount out with slippage tolerance
 */
export function applySlippageTolerance(
  amount: bigint,
  slippageBps: number // Basis points (e.g., 100 = 1%)
): bigint {
  const slippageMultiplier = BigInt(10000 - slippageBps);
  return (amount * slippageMultiplier) / 10000n;
}

