/**
 * HorizonPerks contract interaction layer
 */

import { Address } from 'viem';
import { PublicClient } from 'viem';
import { ContractError, FeeTier } from '../types';

const HORIZON_PERKS_ABI = [
  {
    type: 'function',
    name: 'feeBpsFor',
    inputs: [{ name: 'user', type: 'address' }],
    outputs: [{ name: 'feeBps', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'protocolBpsFor',
    inputs: [{ name: 'user', type: 'address' }],
    outputs: [{ name: 'protocolBps', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'feeBpsForBalance',
    inputs: [{ name: 'balance', type: 'uint256' }],
    outputs: [{ name: 'feeBps', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'protocolBpsForBalance',
    inputs: [{ name: 'balance', type: 'uint256' }],
    outputs: [{ name: 'protocolBps', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'feeTiers',
    inputs: [{ name: 'tierId', type: 'uint256' }],
    outputs: [
      { name: 'minBalance', type: 'uint256' },
      { name: 'feeBps', type: 'uint16' },
      { name: 'protocolBps', type: 'uint16' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'horizonToken',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
] as const;

export class HorizonPerks {
  private client: PublicClient;
  private address: Address;

  constructor(client: PublicClient, address: Address) {
    this.client = client;
    this.address = address;
  }

  /**
   * Get fee tier for a user
   */
  async getFeeTierForUser(userAddress: Address): Promise<FeeTier> {
    try {
      const [feeBps, protocolBps, horizonTokenAddress] = await Promise.all([
        this.client.readContract({
          address: this.address,
          abi: HORIZON_PERKS_ABI,
          functionName: 'feeBpsFor',
          args: [userAddress],
        }),
        this.client.readContract({
          address: this.address,
          abi: HORIZON_PERKS_ABI,
          functionName: 'protocolBpsFor',
          args: [userAddress],
        }),
        this.getHorizonTokenAddress(),
      ]);

      // Get user's HORIZON token balance
      let balance = 0n;
      try {
        const ERC20_ABI = [
          {
            type: 'function',
            name: 'balanceOf',
            inputs: [{ name: 'account', type: 'address' }],
            outputs: [{ name: '', type: 'uint256' }],
            stateMutability: 'view',
          },
        ] as const;
        
        balance = await this.client.readContract({
          address: horizonTokenAddress,
          abi: ERC20_ABI,
          functionName: 'balanceOf',
          args: [userAddress],
        }) as bigint;
      } catch {
        // If balance fetch fails, balance remains 0n
      }

      // Calculate tier from balance by comparing with all tiers
      let tier = 0;
      try {
        const allTiers = await this.getAllFeeTiers();
        // Find the highest tier that the balance qualifies for
        for (let i = allTiers.length - 1; i >= 0; i--) {
          if (balance >= allTiers[i].requiresHorizonTokens) {
            tier = allTiers[i].tier;
            break;
          }
        }
      } catch {
        // If tier calculation fails, tier remains 0
      }

      return {
        tier,
        feeRate: BigInt(feeBps as number),
        requiresHorizonTokens: balance,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get fee tier: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get fee tier for a specific balance
   */
  async getFeeTierForBalance(balance: bigint): Promise<FeeTier> {
    try {
      const [feeBps, protocolBps] = await Promise.all([
        this.client.readContract({
          address: this.address,
          abi: HORIZON_PERKS_ABI,
          functionName: 'feeBpsForBalance',
          args: [balance],
        }),
        this.client.readContract({
          address: this.address,
          abi: HORIZON_PERKS_ABI,
          functionName: 'protocolBpsForBalance',
          args: [balance],
        }),
      ]);

      // Calculate tier from balance by comparing with all tiers
      let tier = 0;
      try {
        const allTiers = await this.getAllFeeTiers();
        // Find the highest tier that the balance qualifies for
        for (let i = allTiers.length - 1; i >= 0; i--) {
          if (balance >= allTiers[i].requiresHorizonTokens) {
            tier = allTiers[i].tier;
            break;
          }
        }
      } catch {
        // If tier calculation fails, tier remains 0
      }

      return {
        tier,
        feeRate: BigInt(feeBps as number),
        requiresHorizonTokens: balance,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get fee tier for balance: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get all fee tiers
   */
  async getAllFeeTiers(): Promise<FeeTier[]> {
    try {
      // Try to read tiers by index until we get an error
      const tiers: FeeTier[] = [];
      let tierId = 0;

      while (true) {
        try {
          const result = await this.client.readContract({
            address: this.address,
            abi: HORIZON_PERKS_ABI,
            functionName: 'feeTiers',
            args: [BigInt(tierId)],
          });

          const [minBalance, feeBps, protocolBps] = result as [bigint, number, number];

          tiers.push({
            tier: tierId,
            feeRate: BigInt(feeBps),
            requiresHorizonTokens: minBalance,
          });

          tierId++;
        } catch {
          // Stop when we've read all tiers
          break;
        }
      }

      return tiers;
    } catch (error) {
      throw new ContractError(
        `Failed to get all fee tiers: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get HORIZON token address
   */
  async getHorizonTokenAddress(): Promise<Address> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: HORIZON_PERKS_ABI,
        functionName: 'horizonToken',
      });

      return result as Address;
    } catch (error) {
      throw new ContractError(
        `Failed to get HORIZON token address: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

