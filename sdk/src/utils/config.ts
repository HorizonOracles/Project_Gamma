/**
 * Configuration utilities for loading environment variables
 */

import { Address } from 'viem';
import { SDKConfig } from '../types';
import { DEFAULT_CONFIG, DEFAULT_TESTNET_CONFIG, BNB_CHAIN } from '../constants';
import { isValidAddress } from './index';

/**
 * Safely get environment variable
 * Works in both Node.js and browser environments
 */
export function getEnvVar(key: string): string | undefined {
  if (typeof process !== 'undefined' && process.env) {
    return process.env[key];
  }
  // In browser environments, env vars should be passed via config
  // Try to access via window if available (for bundlers like Vite)
  if (typeof window !== 'undefined' && (window as any).__ENV__) {
    return (window as any).__ENV__[key];
  }
  return undefined;
}

/**
 * Validate and normalize an address
 * @throws Error if address is invalid
 */
function validateAddress(address: string | undefined, defaultAddress: Address, fieldName: string): Address {
  if (!address) {
    return defaultAddress;
  }
  
  if (!isValidAddress(address)) {
    throw new Error(`Invalid ${fieldName} address: ${address}. Must be a valid Ethereum address.`);
  }
  
  return address as Address;
}

/**
 * Load configuration from environment variables
 * Environment variables take precedence over defaults
 * Automatically selects testnet config when chainId is 97
 */
export function loadConfigFromEnv(): Partial<SDKConfig> {
  const chainIdStr = getEnvVar('CHAIN_ID');
  const chainId = chainIdStr ? parseInt(chainIdStr, 10) : BNB_CHAIN.MAINNET;
  
  // Select appropriate default config based on chainId
  const baseConfig = chainId === BNB_CHAIN.TESTNET ? DEFAULT_TESTNET_CONFIG : DEFAULT_CONFIG;

  const marketFactoryAddr = getEnvVar('MARKET_FACTORY_ADDRESS');
  const horizonTokenAddr = getEnvVar('HORIZON_TOKEN_ADDRESS');
  const outcomeTokenAddr = getEnvVar('OUTCOME_TOKEN_ADDRESS');
  const horizonPerksAddr = getEnvVar('HORIZON_PERKS_ADDRESS');
  const feeSplitterAddr = getEnvVar('FEE_SPLITTER_ADDRESS');
  const resolutionModuleAddr = getEnvVar('RESOLUTION_MODULE_ADDRESS');
  const aiOracleAdapterAddr = getEnvVar('AI_ORACLE_ADAPTER_ADDRESS');

  return {
    chainId,
    rpcUrl: getEnvVar('RPC_URL') || baseConfig.rpcUrl,
    marketFactoryAddress: validateAddress(marketFactoryAddr, baseConfig.marketFactoryAddress, 'MarketFactory'),
    horizonTokenAddress: horizonTokenAddr ? validateAddress(horizonTokenAddr, baseConfig.horizonTokenAddress!, 'HorizonToken') : baseConfig.horizonTokenAddress,
    outcomeTokenAddress: outcomeTokenAddr ? validateAddress(outcomeTokenAddr, baseConfig.outcomeTokenAddress!, 'OutcomeToken') : baseConfig.outcomeTokenAddress,
    horizonPerksAddress: horizonPerksAddr ? validateAddress(horizonPerksAddr, baseConfig.horizonPerksAddress!, 'HorizonPerks') : baseConfig.horizonPerksAddress,
    feeSplitterAddress: feeSplitterAddr ? validateAddress(feeSplitterAddr, baseConfig.feeSplitterAddress!, 'FeeSplitter') : baseConfig.feeSplitterAddress,
    resolutionModuleAddress: resolutionModuleAddr ? validateAddress(resolutionModuleAddr, baseConfig.resolutionModuleAddress!, 'ResolutionModule') : baseConfig.resolutionModuleAddress,
    aiOracleAdapterAddress: aiOracleAdapterAddr ? validateAddress(aiOracleAdapterAddr, baseConfig.aiOracleAdapterAddress!, 'AIOracleAdapter') : baseConfig.aiOracleAdapterAddress,
    explorerUrl: chainId === BNB_CHAIN.MAINNET ? 'https://bscscan.com' : 'https://testnet.bscscan.com',
  };
}

/**
 * Merge configurations with priority: explicit config > env vars > defaults
 * Automatically selects appropriate base config (mainnet/testnet) based on chainId
 */
export function mergeConfig(
  explicitConfig?: Partial<SDKConfig>,
  envConfig?: Partial<SDKConfig>
): SDKConfig {
  const env = envConfig || loadConfigFromEnv();
  
  // Determine base config based on final chainId (explicit > env > default)
  const finalChainId = explicitConfig?.chainId ?? env.chainId ?? DEFAULT_CONFIG.chainId;
  const baseConfig = finalChainId === BNB_CHAIN.TESTNET ? DEFAULT_TESTNET_CONFIG : DEFAULT_CONFIG;
  
  return {
    ...baseConfig,
    ...env,
    ...explicitConfig,
  };
}

