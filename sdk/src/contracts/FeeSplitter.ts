/**
 * FeeSplitter contract interaction layer
 */

import { Address } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError } from '../types';

export interface FeeConfig {
  protocolBps: number;
  creatorBps: number;
}

const FEE_SPLITTER_ABI = [
  {
    type: 'function',
    name: 'claimCreatorFees',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'token', type: 'address' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'claimProtocolFees',
    inputs: [{ name: 'token', type: 'address' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'marketCreator',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'feeConfigs',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [
      { name: 'protocolBps', type: 'uint16' },
      { name: 'creatorBps', type: 'uint16' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'creatorPendingFees',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'token', type: 'address' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'protocolPendingFees',
    inputs: [{ name: 'token', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'protocolTreasury',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
] as const;

export class FeeSplitter {
  private client: PublicClient;
  private walletClient?: WalletClient;
  private address: Address;

  constructor(
    client: PublicClient,
    address: Address,
    walletClient?: WalletClient
  ) {
    this.client = client;
    this.address = address;
    this.walletClient = walletClient;
  }

  /**
   * Get market creator address
   */
  async getMarketCreator(marketId: bigint): Promise<Address> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'marketCreator',
        args: [marketId],
      });

      return result as Address;
    } catch (error) {
      throw new ContractError(
        `Failed to get market creator: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get fee configuration for a market
   */
  async getFeeConfig(marketId: bigint): Promise<FeeConfig> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'feeConfigs',
        args: [marketId],
      });

      const [protocolBps, creatorBps] = result as [number, number];

      return {
        protocolBps,
        creatorBps,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get fee config: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get pending creator fees for a market
   */
  async getCreatorPendingFees(marketId: bigint, token: Address): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'creatorPendingFees',
        args: [marketId, token],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get creator pending fees: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get pending protocol fees
   */
  async getProtocolPendingFees(token: Address): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'protocolPendingFees',
        args: [token],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get protocol pending fees: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Claim creator fees for a market
   */
  async claimCreatorFees(marketId: bigint, token: Address): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for claiming fees', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'claimCreatorFees',
        args: [marketId, token],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to claim creator fees: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Claim protocol fees
   */
  async claimProtocolFees(token: Address): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for claiming fees', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'claimProtocolFees',
        args: [token],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to claim protocol fees: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get protocol treasury address
   */
  async getProtocolTreasury(): Promise<Address> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: FEE_SPLITTER_ABI,
        functionName: 'protocolTreasury',
      });

      return result as Address;
    } catch (error) {
      throw new ContractError(
        `Failed to get protocol treasury: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

