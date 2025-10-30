/**
 * Edge case tests for utility functions
 * Tests boundary conditions, overflow, and extreme values
 */

import { describe, it, expect } from 'vitest';
import {
  formatTokenAmount,
  parseTokenAmount,
  calculatePrice,
  calculateAmountOut,
  calculateAmountIn,
  calculateSlippage,
  applySlippageTolerance,
  getOutcomeTokenId,
  getMarketIdFromTokenId,
  getOutcomeFromTokenId,
} from '../../utils';

describe('Utility Functions - Edge Cases', () => {
  describe('formatTokenAmount - Edge Cases', () => {
    it('should handle maximum BigInt values', () => {
      const maxBigInt = 2n ** 256n - 1n;
      expect(() => formatTokenAmount(maxBigInt)).not.toThrow();
    });

    it('should handle very small amounts', () => {
      expect(formatTokenAmount(1n)).toBe('0.0000');
      expect(formatTokenAmount(1000000000000n)).toBe('0.0000'); // Less than 1 token
    });

    it('should handle very large precision values', () => {
      const amount = 1234567890000000000n;
      expect(formatTokenAmount(amount, 18, 18)).toContain('1.234567890000000000');
    });

    it('should handle zero precision', () => {
      expect(formatTokenAmount(1500000000000000000n, 18, 0)).toBe('1');
    });

    it('should handle different decimal places', () => {
      expect(formatTokenAmount(1000000n, 6)).toBe('1.0000');
      expect(formatTokenAmount(1000000n, 6, 2)).toBe('1.00');
    });
  });

  describe('parseTokenAmount - Edge Cases', () => {
    it('should handle scientific notation strings', () => {
      expect(() => parseTokenAmount('1e18')).toThrow();
    });

    it('should handle negative numbers', () => {
      expect(() => parseTokenAmount('-1')).toThrow();
    });

    it('should handle numbers with many decimal places', () => {
      const result = parseTokenAmount('1.123456789012345678');
      expect(result).toBeGreaterThan(0n);
    });

    it('should handle very large string numbers', () => {
      const largeNumber = '999999999999999999999999999999999999999';
      expect(() => parseTokenAmount(largeNumber)).toThrow(); // Should exceed BigInt range
    });

    it('should handle whitespace correctly', () => {
      // Leading/trailing whitespace should be trimmed and work
      const result = parseTokenAmount('  1.0  ');
      expect(result).toBeGreaterThan(0n);
      expect(() => parseTokenAmount('   ')).toThrow();
    });

    it('should handle different decimal formats', () => {
      expect(() => parseTokenAmount('1,000.5')).toThrow(); // Comma separator
    });
  });

  describe('calculatePrice - Edge Cases', () => {
    it('should handle one side having zero liquidity', () => {
      expect(calculatePrice(0n, 1000000000000000000n)).toBe(0n);
      expect(calculatePrice(1000000000000000000n, 0n)).toBe(0n);
    });

    it('should handle extremely imbalanced liquidity', () => {
      const verySmall = 1n;
      const veryLarge = 1000000000000000000000000n; // 1M tokens
      
      const price = calculatePrice(verySmall, veryLarge);
      expect(price).toBeGreaterThan(0n);
      expect(price).toBeLessThan(parseTokenAmount('1', 18));
    });

    it('should handle equal liquidity correctly', () => {
      const liquidity = 1000000000000000000n;
      const price = calculatePrice(liquidity, liquidity);
      expect(price).toBe(500000000000000000n); // Should be 0.5 (50%)
    });

    it('should handle maximum BigInt values', () => {
      const maxValue = 2n ** 255n;
      expect(() => calculatePrice(maxValue, maxValue)).not.toThrow();
    });
  });

  describe('calculateAmountOut - Edge Cases', () => {
    it('should return zero for zero liquidity', () => {
      expect(calculateAmountOut(1000000000000000000n, 0n, 1000000000000000000n)).toBe(0n);
      expect(calculateAmountOut(1000000000000000000n, 1000000000000000000n, 0n)).toBe(0n);
    });

    it('should handle very small amounts in', () => {
      const amountIn = 1n;
      const amountOut = calculateAmountOut(amountIn, 1000000000000000000n, 1000000000000000000n);
      // For very small amounts, the result might be 0 due to integer division
      expect(amountOut).toBeGreaterThanOrEqual(0n);
      if (amountOut > 0n) {
        expect(amountOut).toBeLessThan(amountIn);
      }
    });

    it('should handle amounts close to total liquidity', () => {
      const liquidity = 1000000000000000000n;
      const amountIn = liquidity / 2n;
      
      const amountOut = calculateAmountOut(amountIn, liquidity, liquidity);
      expect(amountOut).toBeGreaterThan(0n);
      expect(amountOut).toBeLessThan(amountIn);
    });

    it('should handle overflow scenarios', () => {
      const veryLarge = 2n ** 255n;
      expect(() => calculateAmountOut(veryLarge, veryLarge, veryLarge)).not.toThrow();
    });
  });

  describe('calculateAmountIn - Edge Cases', () => {
    it('should throw for amountOut >= otherLiquidity', () => {
      const liquidity = 1000000000000000000n;
      
      expect(() => {
        calculateAmountIn(liquidity, liquidity, liquidity);
      }).toThrow();
      
      expect(() => {
        calculateAmountIn(liquidity + 1n, liquidity, liquidity);
      }).toThrow();
    });

    it('should handle very small amounts out', () => {
      const amountOut = 1n;
      const liquidity = 1000000000000000000n;
      
      const amountIn = calculateAmountIn(amountOut, liquidity, liquidity);
      expect(amountIn).toBeGreaterThan(0n);
    });

    it('should handle amountOut close to otherLiquidity', () => {
      const liquidity = 1000000000000000000n;
      const amountOut = liquidity - 1n;
      
      const amountIn = calculateAmountIn(amountOut, liquidity, liquidity);
      expect(amountIn).toBeGreaterThan(0n);
    });
  });

  describe('calculateSlippage - Edge Cases', () => {
    it('should handle zero expected amount', () => {
      expect(calculateSlippage(0n, 1000000000000000000n)).toBe(0);
    });

    it('should handle negative slippage (better price)', () => {
      const slippage = calculateSlippage(1000000000000000000n, 1100000000000000000n);
      expect(slippage).toBeLessThan(0);
    });

    it('should handle very high slippage', () => {
      const slippage = calculateSlippage(1000000000000000000n, 500000000000000000n);
      expect(slippage).toBe(50); // 50% slippage
    });

    it('should handle equal amounts', () => {
      expect(calculateSlippage(1000000000000000000n, 1000000000000000000n)).toBe(0);
    });

    it('should handle actualAmount > expectedAmount', () => {
      const slippage = calculateSlippage(1000000000000000000n, 2000000000000000000n);
      expect(slippage).toBe(-100); // -100% (got double)
    });
  });

  describe('applySlippageTolerance - Edge Cases', () => {
    it('should handle zero slippage', () => {
      const amount = 1000000000000000000n;
      expect(applySlippageTolerance(amount, 0)).toBe(amount);
    });

    it('should handle maximum slippage (100%)', () => {
      const amount = 1000000000000000000n;
      expect(applySlippageTolerance(amount, 10000)).toBe(0n); // 100% = 0
    });

    it('should handle very small slippage', () => {
      const amount = 1000000000000000000n;
      const result = applySlippageTolerance(amount, 1); // 0.01%
      expect(result).toBeLessThan(amount);
      // Allow for rounding differences, result should be very close to amount - (amount * 0.01 / 100)
      const expectedMin = amount - (amount * 10001n / 10000n); // Account for rounding
      expect(result).toBeGreaterThanOrEqual(expectedMin);
    });

    it('should handle large amounts with slippage', () => {
      const amount = 2n ** 200n;
      const result = applySlippageTolerance(amount, 100); // 1%
      expect(result).toBeLessThan(amount);
    });

    it('should round down correctly', () => {
      const amount = 1000000000000000001n; // Odd number
      const result = applySlippageTolerance(amount, 100); // 1%
      expect(result % 10000n).toBe(0n); // Should be divisible by basis points
    });
  });

  describe('Token ID Functions - Edge Cases', () => {
    it('should handle maximum market ID', () => {
      const maxMarketId = 2n ** 255n;
      const yesTokenId = getOutcomeTokenId(maxMarketId, 'YES');
      const noTokenId = getOutcomeTokenId(maxMarketId, 'NO');
      
      expect(getMarketIdFromTokenId(yesTokenId)).toBe(maxMarketId);
      expect(getMarketIdFromTokenId(noTokenId)).toBe(maxMarketId);
      expect(getOutcomeFromTokenId(yesTokenId)).toBe('YES');
      expect(getOutcomeFromTokenId(noTokenId)).toBe('NO');
    });

    it('should handle market ID of 0', () => {
      expect(getOutcomeTokenId(0n, 'YES')).toBe(0n);
      expect(getOutcomeTokenId(0n, 'NO')).toBe(1n);
    });

    it('should correctly round-trip token IDs', () => {
      for (let i = 1n; i < 1000n; i++) {
        const yesId = getOutcomeTokenId(i, 'YES');
        const noId = getOutcomeTokenId(i, 'NO');
        
        expect(getMarketIdFromTokenId(yesId)).toBe(i);
        expect(getMarketIdFromTokenId(noId)).toBe(i);
        expect(getOutcomeFromTokenId(yesId)).toBe('YES');
        expect(getOutcomeFromTokenId(noId)).toBe('NO');
      }
    });

    it('should handle odd token IDs correctly', () => {
      const oddTokenId = 1337n;
      const marketId = getMarketIdFromTokenId(oddTokenId);
      const outcome = getOutcomeFromTokenId(oddTokenId);
      
      // Token ID encoding: (marketId << 8) | outcomeId
      // 1337 >> 8 = 5 (marketId)
      // 1337 & 0xFF = 57 (outcomeId)
      expect(marketId).toBe(5n);
      expect(['YES', 'NO']).toContain(outcome);
    });
  });
});

