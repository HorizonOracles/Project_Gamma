/**
 * Unit tests for type definitions and error classes
 */

import { describe, it, expect } from 'vitest';
import { SDKError, ContractError, TradeError } from '../../types';

describe('Error Classes', () => {
  describe('SDKError', () => {
    it('should create SDKError with message', () => {
      const error = new SDKError('Test error');
      expect(error.message).toBe('Test error');
      expect(error.name).toBe('SDKError');
      expect(error.code).toBeUndefined();
    });

    it('should create SDKError with code', () => {
      const error = new SDKError('Test error', 'TEST_CODE');
      expect(error.code).toBe('TEST_CODE');
    });

    it('should create SDKError with data', () => {
      const data = { key: 'value' };
      const error = new SDKError('Test error', 'TEST_CODE', data);
      expect(error.data).toEqual(data);
    });
  });

  describe('ContractError', () => {
    it('should create ContractError with contract address', () => {
      const address = '0x1234567890123456789012345678901234567890' as const;
      const error = new ContractError('Contract error', address);
      
      expect(error.message).toBe('Contract error');
      expect(error.name).toBe('ContractError');
      expect(error.code).toBe('CONTRACT_ERROR');
      expect(error.contractAddress).toBe(address);
    });

    it('should be instance of SDKError', () => {
      const error = new ContractError('Test', '0x123' as const);
      expect(error).toBeInstanceOf(SDKError);
    });
  });

  describe('TradeError', () => {
    it('should create TradeError with market ID', () => {
      const marketId = 1n;
      const error = new TradeError('Trade error', marketId);
      
      expect(error.message).toBe('Trade error');
      expect(error.name).toBe('TradeError');
      expect(error.code).toBe('TRADE_ERROR');
      expect(error.marketId).toBe(marketId);
    });

    it('should be instance of SDKError', () => {
      const error = new TradeError('Test', 1n);
      expect(error).toBeInstanceOf(SDKError);
    });
  });
});

