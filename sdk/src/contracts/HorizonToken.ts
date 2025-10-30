/**
 * HorizonToken (ERC-20) contract interaction layer
 */

import { Address } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError } from '../types';

const ERC20_ABI = [
  {
    type: 'function',
    name: 'balanceOf',
    inputs: [{ name: 'account', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'allowance',
    inputs: [
      { name: 'owner', type: 'address' },
      { name: 'spender', type: 'address' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'approve',
    inputs: [
      { name: 'spender', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'transfer',
    inputs: [
      { name: 'to', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'transferFrom',
    inputs: [
      { name: 'from', type: 'address' },
      { name: 'to', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'decimals',
    inputs: [],
    outputs: [{ name: '', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'symbol',
    inputs: [],
    outputs: [{ name: '', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'name',
    inputs: [],
    outputs: [{ name: '', type: 'string' }],
    stateMutability: 'view',
  },
] as const;

export class HorizonToken {
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
   * Get balance of HORIZON tokens for an address
   */
  async balanceOf(account: Address): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'balanceOf',
        args: [account],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get balance: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get allowance for a spender
   */
  async allowance(owner: Address, spender: Address): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'allowance',
        args: [owner, spender],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get allowance: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Approve spender to spend tokens
   */
  async approve(spender: Address, amount: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for approving', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'approve',
        args: [spender, amount],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to approve: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Transfer tokens
   */
  async transfer(to: Address, amount: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for transferring', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'transfer',
        args: [to, amount],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to transfer: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get token decimals
   */
  async decimals(): Promise<number> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'decimals',
      });

      return result as number;
    } catch (error) {
      throw new ContractError(
        `Failed to get decimals: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get token symbol
   */
  async symbol(): Promise<string> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'symbol',
      });

      return result as string;
    } catch (error) {
      throw new ContractError(
        `Failed to get symbol: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get token name
   */
  async name(): Promise<string> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: ERC20_ABI,
        functionName: 'name',
      });

      return result as string;
    } catch (error) {
      throw new ContractError(
        `Failed to get name: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

