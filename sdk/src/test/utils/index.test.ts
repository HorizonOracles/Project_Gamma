/**
 * Unit tests for utility functions
 */

import { describe, it, expect } from 'vitest';
import {
  formatTokenAmount,
  parseTokenAmount,
  calculatePrice,
  calculateMarketPrices,
  calculateAmountOut,
  calculateAmountIn,
  calculateSlippage,
  isValidAddress,
  shortenAddress,
  getOutcomeTokenId,
  getMarketIdFromTokenId,
  getOutcomeFromTokenId,
  applySlippageTolerance,
} from '../../utils';

describe('formatTokenAmount', () => {
  it('should format token amount correctly', () => {
    expect(formatTokenAmount(1000000000000000000n)).toBe('1.0000');
    expect(formatTokenAmount(1500000000000000000n)).toBe('1.5000');
    expect(formatTokenAmount(100000000000000000n)).toBe('0.1000');
  });

  it('should handle zero amount', () => {
    expect(formatTokenAmount(0n)).toBe('0');
  });

  it('should respect precision parameter', () => {
    expect(formatTokenAmount(1234567890000000000n, 18, 2)).toBe('1.23');
    expect(formatTokenAmount(1234567890000000000n, 18, 6)).toBe('1.234567');
  });
});

describe('parseTokenAmount', () => {
  it('should parse token amount correctly', () => {
    expect(parseTokenAmount('1.0')).toBe(1000000000000000000n);
    expect(parseTokenAmount('1.5')).toBe(1500000000000000000n);
    expect(parseTokenAmount('0.1')).toBe(100000000000000000n);
  });

  it('should handle integer amounts', () => {
    expect(parseTokenAmount('1')).toBe(1000000000000000000n);
    expect(parseTokenAmount('10')).toBe(10000000000000000000n);
  });

  it('should respect decimals parameter', () => {
    expect(parseTokenAmount('1.0', 6)).toBe(1000000n);
    expect(parseTokenAmount('1.5', 6)).toBe(1500000n);
  });

  it('should throw error for invalid amount', () => {
    expect(() => parseTokenAmount('invalid')).toThrow();
    expect(() => parseTokenAmount('')).toThrow();
  });
});

describe('calculatePrice', () => {
  it('should calculate price correctly', () => {
    const price = calculatePrice(1000000000000000000n, 1000000000000000000n);
    expect(price).toBe(500000000000000000n); // 0.5 (50%)
  });

  it('should handle zero liquidity', () => {
    expect(calculatePrice(0n, 1000000000000000000n)).toBe(0n);
    expect(calculatePrice(1000000000000000000n, 0n)).toBe(0n);
  });

  it('should handle unequal liquidity', () => {
    const price = calculatePrice(2000000000000000000n, 1000000000000000000n);
    // price = (1e18 * 1e18) / (2e18 + 1e18) = 1e36 / 3e18 = 1e18 / 3
    expect(price).toBeGreaterThan(0n);
  });
});

describe('calculateMarketPrices', () => {
  it('should calculate both market prices', () => {
    const prices = calculateMarketPrices(
      1000000000000000000n,
      1000000000000000000n
    );
    expect(prices.yesPrice).toBeGreaterThan(0n);
    expect(prices.noPrice).toBeGreaterThan(0n);
    // Prices should sum to approximately 1 (1e18), allowing for rounding
    const totalPrice = prices.yesPrice + prices.noPrice;
    expect(totalPrice).toBeGreaterThan(990000000000000000n);
    expect(totalPrice).toBeLessThan(1010000000000000000n);
  });
});

describe('calculateAmountOut', () => {
  it('should calculate amount out correctly', () => {
    const amountOut = calculateAmountOut(
      1000000000000000000n, // 1 token in
      2000000000000000000n, // 2 tokens liquidity
      1000000000000000000n  // 1 token other liquidity
    );
    expect(amountOut).toBeGreaterThan(0n);
    expect(amountOut).toBeLessThan(1000000000000000000n);
  });

  it('should handle zero liquidity', () => {
    expect(calculateAmountOut(1000000000000000000n, 0n, 1000000000000000000n)).toBe(0n);
  });
});

describe('calculateAmountIn', () => {
  it('should calculate amount in correctly', () => {
    const amountIn = calculateAmountIn(
      500000000000000000n, // 0.5 tokens out
      2000000000000000000n, // 2 tokens liquidity
      1000000000000000000n  // 1 token other liquidity
    );
    expect(amountIn).toBeGreaterThan(0n);
  });

  it('should throw error for insufficient liquidity', () => {
    expect(() => {
      calculateAmountIn(
        2000000000000000000n, // Want more than available
        1000000000000000000n,
        1000000000000000000n
      );
    }).toThrow();
  });
});

describe('calculateSlippage', () => {
  it('should calculate slippage correctly', () => {
    const slippage = calculateSlippage(1000000000000000000n, 950000000000000000n);
    expect(slippage).toBe(5); // 5% slippage
  });

  it('should handle zero expected amount', () => {
    expect(calculateSlippage(0n, 1000000000000000000n)).toBe(0);
  });

  it('should handle negative slippage (better price)', () => {
    const slippage = calculateSlippage(1000000000000000000n, 1050000000000000000n);
    expect(slippage).toBeLessThan(0);
  });
});

describe('isValidAddress', () => {
  it('should validate correct addresses', () => {
    expect(isValidAddress('0x1234567890123456789012345678901234567890')).toBe(true);
    expect(isValidAddress('0xABCDEFABCDEFABCDEFABCDEFABCDEFABCDEFABCD')).toBe(true);
  });

  it('should reject invalid addresses', () => {
    expect(isValidAddress('0x123')).toBe(false);
    expect(isValidAddress('not an address')).toBe(false);
    expect(isValidAddress('')).toBe(false);
  });
});

describe('shortenAddress', () => {
  it('should shorten address correctly', () => {
    const address = '0x1234567890123456789012345678901234567890' as Address;
    const shortened = shortenAddress(address);
    expect(shortened).toContain('...');
    expect(shortened).toHaveLength(13); // 0x (2) + 4 chars + ... (3) + 4 chars = 13
  });

  it('should respect chars parameter', () => {
    const address = '0x1234567890123456789012345678901234567890' as Address;
    const shortened = shortenAddress(address, 6);
    expect(shortened).toHaveLength(17); // 0x + 6 + ... + 6
  });
});

describe('getOutcomeTokenId', () => {
  it('should return correct token ID for YES', () => {
    // Token ID encoding: (marketId << 8) | outcomeId
    // For marketId=1, YES (outcomeId=0): (1 << 8) | 0 = 256
    expect(getOutcomeTokenId(1n, 'YES')).toBe(256n);
    // For marketId=5, YES (outcomeId=0): (5 << 8) | 0 = 1280
    expect(getOutcomeTokenId(5n, 'YES')).toBe(1280n);
  });

  it('should return correct token ID for NO', () => {
    // Token ID encoding: (marketId << 8) | outcomeId
    // For marketId=1, NO (outcomeId=1): (1 << 8) | 1 = 257
    expect(getOutcomeTokenId(1n, 'NO')).toBe(257n);
    // For marketId=5, NO (outcomeId=1): (5 << 8) | 1 = 1281
    expect(getOutcomeTokenId(5n, 'NO')).toBe(1281n);
  });
});

describe('getMarketIdFromTokenId', () => {
  it('should extract market ID correctly', () => {
    // Token ID 256 = (1 << 8) | 0, marketId = 256 >> 8 = 1
    expect(getMarketIdFromTokenId(256n)).toBe(1n);
    // Token ID 257 = (1 << 8) | 1, marketId = 257 >> 8 = 1
    expect(getMarketIdFromTokenId(257n)).toBe(1n);
    // Token ID 1280 = (5 << 8) | 0, marketId = 1280 >> 8 = 5
    expect(getMarketIdFromTokenId(1280n)).toBe(5n);
  });
});

describe('getOutcomeFromTokenId', () => {
  it('should extract YES outcome correctly', () => {
    // Token ID 256 = (1 << 8) | 0, outcomeId = 256 & 0xFF = 0 (YES)
    expect(getOutcomeFromTokenId(256n)).toBe('YES');
    // Token ID 1280 = (5 << 8) | 0, outcomeId = 1280 & 0xFF = 0 (YES)
    expect(getOutcomeFromTokenId(1280n)).toBe('YES');
  });

  it('should extract NO outcome correctly', () => {
    // Token ID 257 = (1 << 8) | 1, outcomeId = 257 & 0xFF = 1 (NO)
    expect(getOutcomeFromTokenId(257n)).toBe('NO');
    // Token ID 1281 = (5 << 8) | 1, outcomeId = 1281 & 0xFF = 1 (NO)
    expect(getOutcomeFromTokenId(1281n)).toBe('NO');
  });
});

describe('applySlippageTolerance', () => {
  it('should apply slippage tolerance correctly', () => {
    const amount = 1000000000000000000n; // 1 token
    const withSlippage = applySlippageTolerance(amount, 100); // 1% slippage
    expect(withSlippage).toBe(990000000000000000n); // 0.99 tokens
  });

  it('should handle zero slippage', () => {
    const amount = 1000000000000000000n;
    const withSlippage = applySlippageTolerance(amount, 0);
    expect(withSlippage).toBe(amount);
  });

  it('should handle high slippage', () => {
    const amount = 1000000000000000000n;
    const withSlippage = applySlippageTolerance(amount, 500); // 5% slippage
    expect(withSlippage).toBe(950000000000000000n);
  });
});

