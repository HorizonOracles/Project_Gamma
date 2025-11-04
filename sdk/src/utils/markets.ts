/**
 * Market type detection and contract instantiation utilities
 */

import { Address, PublicClient, WalletClient } from 'viem';
import { MarketType, ContractError } from '../types';
import { MARKET_FACTORY_ABI } from '../constants';
import { BinaryMarket } from '../contracts/BinaryMarket';

/**
 * Get market type from MarketFactory
 * @param client - Viem public client
 * @param factoryAddress - MarketFactory contract address
 * @param marketId - Market ID (not the AMM address)
 * @returns MarketType enum value
 */
export async function getMarketType(
  client: PublicClient,
  factoryAddress: Address,
  marketId: bigint
): Promise<MarketType> {
  try {
    // Get market info from factory - includes marketType field
    const result = await client.readContract({
      address: factoryAddress,
      abi: MARKET_FACTORY_ABI,
      functionName: 'getMarket',
      args: [marketId],
    });

    // Result is a Market struct with marketType field
    const marketData = result as {
      id: bigint;
      creator: Address;
      marketType: number;
      amm: Address;
      collateralToken: Address;
      closeTime: bigint;
      category: string;
      metadataURI: string;
      creatorStake: bigint;
      outcomeCount: number;
      stakeRefunded: boolean;
      status: number;
    };
    
    return marketData.marketType as MarketType;
  } catch (error) {
    throw new ContractError(
      `Failed to get market type: ${error instanceof Error ? error.message : 'Unknown error'}`,
      factoryAddress,
      error
    );
  }
}

/**
 * Get market type directly from market contract
 * All market contracts implement IMarket.getMarketType()
 * @param client - Viem public client
 * @param marketAddress - Market contract address
 * @returns MarketType enum value
 */
export async function getMarketTypeFromContract(
  client: PublicClient,
  marketAddress: Address
): Promise<MarketType> {
  try {
    const result = await client.readContract({
      address: marketAddress,
      abi: [
        {
          type: 'function',
          name: 'getMarketType',
          inputs: [],
          outputs: [{ name: '', type: 'uint8' }],
          stateMutability: 'pure',
        },
      ],
      functionName: 'getMarketType',
    });

    return result as MarketType;
  } catch (error) {
    throw new ContractError(
      `Failed to get market type from contract: ${error instanceof Error ? error.message : 'Unknown error'}`,
      marketAddress,
      error
    );
  }
}

/**
 * Get outcome count from market contract
 * All market contracts implement IMarket.getOutcomeCount()
 * @param client - Viem public client
 * @param marketAddress - Market contract address
 * @returns Number of outcomes (2 for binary, 3-8 for multi-choice, etc.)
 */
export async function getOutcomeCount(
  client: PublicClient,
  marketAddress: Address
): Promise<number> {
  try {
    const result = await client.readContract({
      address: marketAddress,
      abi: [
        {
          type: 'function',
          name: 'getOutcomeCount',
          inputs: [],
          outputs: [{ name: '', type: 'uint8' }],
          stateMutability: 'pure',
        },
      ],
      functionName: 'getOutcomeCount',
    });

    return Number(result);
  } catch (error) {
    throw new ContractError(
      `Failed to get outcome count: ${error instanceof Error ? error.message : 'Unknown error'}`,
      marketAddress,
      error
    );
  }
}

/**
 * Type guard to check if market is binary
 */
export function isBinaryMarket(marketType: MarketType): boolean {
  return marketType === MarketType.Binary;
}

/**
 * Type guard to check if market is multi-choice
 */
export function isMultiChoiceMarket(marketType: MarketType): boolean {
  return marketType === MarketType.MultiChoice;
}

/**
 * Type guard to check if market is limit order
 */
export function isLimitOrderMarket(marketType: MarketType): boolean {
  return marketType === MarketType.LimitOrder;
}

/**
 * Type guard to check if market is pooled liquidity
 */
export function isPooledLiquidityMarket(marketType: MarketType): boolean {
  return marketType === MarketType.PooledLiquidity;
}

/**
 * Instantiate correct market contract class based on market type
 * @param client - Viem public client
 * @param marketAddress - Market contract address
 * @param marketType - Market type (if known, otherwise will be fetched)
 * @param walletClient - Optional wallet client for write operations
 * @returns Typed market contract instance
 */
export async function instantiateMarketContract(
  client: PublicClient,
  marketAddress: Address,
  marketType?: MarketType,
  walletClient?: WalletClient
): Promise<BinaryMarket> {
  // If market type not provided, fetch it
  if (marketType === undefined) {
    marketType = await getMarketTypeFromContract(client, marketAddress);
  }

  switch (marketType) {
    case MarketType.Binary:
      return new BinaryMarket(client, marketAddress, walletClient);
    
    // TODO: Add support for other market types
    case MarketType.MultiChoice:
    case MarketType.LimitOrder:
    case MarketType.PooledLiquidity:
    default:
      throw new ContractError(
        `Market type ${marketType} not yet supported. Only BinaryMarket is currently supported.`,
        marketAddress
      );
  }
}

/**
 * Get market contract instance with automatic type detection
 * Convenience wrapper around instantiateMarketContract
 */
export async function getMarketContract(
  client: PublicClient,
  marketAddress: Address,
  walletClient?: WalletClient
): Promise<BinaryMarket> {
  return instantiateMarketContract(client, marketAddress, undefined, walletClient);
}
