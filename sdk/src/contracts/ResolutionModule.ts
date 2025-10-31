/**
 * ResolutionModule contract interaction layer
 */

import { Address } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError } from '../types';

export enum ResolutionState {
  None = 0,
  Proposed = 1,
  Disputed = 2,
  Finalized = 3,
}

export interface Resolution {
  state: ResolutionState;
  proposedOutcome: bigint;
  proposalTime: bigint;
  proposer: Address;
  proposerBond: bigint;
  disputer: Address;
  disputerBond: bigint;
  evidenceURI: string;
}

const RESOLUTION_MODULE_ABI = [
  {
    type: 'function',
    name: 'proposeResolution',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'outcomeId', type: 'uint256' },
      { name: 'bondAmount', type: 'uint256' },
      { name: 'evidenceURI', type: 'string' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'dispute',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'bondAmount', type: 'uint256' },
      { name: 'reason', type: 'string' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'finalize',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'finalizeDisputed',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'outcomeId', type: 'uint256' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'resolutions',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [
      { name: 'state', type: 'uint8' },
      { name: 'proposedOutcome', type: 'uint256' },
      { name: 'proposalTime', type: 'uint256' },
      { name: 'proposer', type: 'address' },
      { name: 'proposerBond', type: 'uint256' },
      { name: 'disputer', type: 'address' },
      { name: 'disputerBond', type: 'uint256' },
      { name: 'evidenceURI', type: 'string' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'disputeWindow',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'minBond',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'arbitrator',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
] as const;

export class ResolutionModule {
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
   * Get resolution data for a market
   */
  async getResolution(marketId: bigint): Promise<Resolution> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'resolutions',
        args: [marketId],
      });

      const [
        state,
        proposedOutcome,
        proposalTime,
        proposer,
        proposerBond,
        disputer,
        disputerBond,
        evidenceURI,
      ] = result as [
        number,
        bigint,
        bigint,
        Address,
        bigint,
        Address,
        bigint,
        string,
      ];

      return {
        state: state as ResolutionState,
        proposedOutcome,
        proposalTime,
        proposer,
        proposerBond,
        disputer,
        disputerBond,
        evidenceURI,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Propose a resolution for a market
   */
  async proposeResolution(
    marketId: bigint,
    outcomeId: bigint,
    bondAmount: bigint,
    evidenceURI: string
  ): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for proposing resolution', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'proposeResolution',
        args: [marketId, outcomeId, bondAmount, evidenceURI],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to propose resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Dispute a proposed resolution
   */
  async dispute(
    marketId: bigint,
    bondAmount: bigint,
    reason: string
  ): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for disputing', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'dispute',
        args: [marketId, bondAmount, reason],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to dispute resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Finalize a resolution (after dispute window)
   */
  async finalize(marketId: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for finalizing', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'finalize',
        args: [marketId],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to finalize resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Finalize a disputed resolution (arbitrator only)
   */
  async finalizeDisputed(marketId: bigint, outcomeId: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for finalizing disputed resolution', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'finalizeDisputed',
        args: [marketId, outcomeId],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to finalize disputed resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get dispute window duration
   */
  async getDisputeWindow(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'disputeWindow',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get dispute window: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get minimum bond amount
   */
  async getMinBond(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'minBond',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get min bond: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get arbitrator address
   */
  async getArbitrator(): Promise<Address> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'arbitrator',
      });

      return result as Address;
    } catch (error) {
      throw new ContractError(
        `Failed to get arbitrator: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

