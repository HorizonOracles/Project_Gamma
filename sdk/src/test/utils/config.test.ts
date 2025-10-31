/**
 * Unit tests for configuration utilities
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { loadConfigFromEnv, mergeConfig, getEnvVar } from '../../utils/config';
import { DEFAULT_CONFIG } from '../../constants';
import type { SDKConfig } from '../../types';

describe('getEnvVar', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('should read from process.env in Node.js', () => {
    process.env.TEST_VAR = 'test-value';
    expect(getEnvVar('TEST_VAR')).toBe('test-value');
    delete process.env.TEST_VAR;
  });

  it('should return undefined for non-existent variable', () => {
    expect(getEnvVar('NON_EXISTENT_VAR')).toBeUndefined();
  });
});

describe('loadConfigFromEnv', () => {
  beforeEach(() => {
    // Clear env vars
    delete process.env.CHAIN_ID;
    delete process.env.RPC_URL;
    delete process.env.MARKET_FACTORY_ADDRESS;
  });

  it('should load default config when no env vars set', () => {
    const config = loadConfigFromEnv();
    expect(config.chainId).toBe(56);
    expect(config.rpcUrl).toBe('https://bsc-dataseed.binance.org/');
    expect(config.marketFactoryAddress).toBeDefined();
  });

  it('should override chainId from env', () => {
    process.env.CHAIN_ID = '97';
    const config = loadConfigFromEnv();
    expect(config.chainId).toBe(97);
    delete process.env.CHAIN_ID;
  });

  it('should override RPC URL from env', () => {
    process.env.RPC_URL = 'https://custom-rpc.com';
    const config = loadConfigFromEnv();
    expect(config.rpcUrl).toBe('https://custom-rpc.com');
    delete process.env.RPC_URL;
  });

  it('should override contract addresses from env', () => {
    process.env.MARKET_FACTORY_ADDRESS = '0x1111111111111111111111111111111111111111';
    const config = loadConfigFromEnv();
    expect(config.marketFactoryAddress).toBe('0x1111111111111111111111111111111111111111');
    delete process.env.MARKET_FACTORY_ADDRESS;
  });

  it('should load all contract addresses from env', () => {
    process.env.HORIZON_TOKEN_ADDRESS = '0x2222222222222222222222222222222222222222';
    process.env.OUTCOME_TOKEN_ADDRESS = '0x3333333333333333333333333333333333333333';
    process.env.HORIZON_PERKS_ADDRESS = '0x4444444444444444444444444444444444444444';
    
    const config = loadConfigFromEnv();
    expect(config.horizonTokenAddress).toBe('0x2222222222222222222222222222222222222222');
    expect(config.outcomeTokenAddress).toBe('0x3333333333333333333333333333333333333333');
    expect(config.horizonPerksAddress).toBe('0x4444444444444444444444444444444444444444');
    
    delete process.env.HORIZON_TOKEN_ADDRESS;
    delete process.env.OUTCOME_TOKEN_ADDRESS;
    delete process.env.HORIZON_PERKS_ADDRESS;
  });
});

describe('mergeConfig', () => {
  it('should merge configs with priority: explicit > env > defaults', () => {
    const envConfig: Partial<SDKConfig> = {
      chainId: 97,
      rpcUrl: 'https://testnet-rpc.com',
    };
    
    const explicitConfig: Partial<SDKConfig> = {
      chainId: 56,
    };
    
    const merged = mergeConfig(explicitConfig, envConfig);
    expect(merged.chainId).toBe(56); // Explicit takes priority
    expect(merged.rpcUrl).toBe('https://testnet-rpc.com'); // From env
    expect(merged.marketFactoryAddress).toBeDefined(); // From defaults
  });

  it('should use defaults when no overrides provided', () => {
    const merged = mergeConfig();
    expect(merged.chainId).toBe(DEFAULT_CONFIG.chainId);
    expect(merged.marketFactoryAddress).toBe(DEFAULT_CONFIG.marketFactoryAddress);
  });

  it('should handle partial explicit config', () => {
    const explicitConfig: Partial<SDKConfig> = {
      rpcUrl: 'https://custom.com',
    };
    
    const merged = mergeConfig(explicitConfig);
    expect(merged.rpcUrl).toBe('https://custom.com');
    expect(merged.chainId).toBe(DEFAULT_CONFIG.chainId);
  });
});

