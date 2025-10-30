/**
 * OutcomeToken contract interaction layer
 */

import { Address, decodeEventLog } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError } from '../types';
import { OUTCOME_TOKEN_ABI } from '../constants';

export class OutcomeToken {
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
   * Check if a market is resolved
   */
  async isResolved(marketId: bigint): Promise<boolean> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'isResolved',
        args: [marketId],
      });

      return result as boolean;
    } catch (error) {
      throw new ContractError(
        `Failed to check if market is resolved: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get winning outcome for a market
   * Returns MAX_UINT256 if not resolved
   */
  async getWinningOutcome(marketId: bigint): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'winningOutcome',
        args: [marketId],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get winning outcome: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get collateral token address for a market
   */
  async getMarketCollateral(marketId: bigint): Promise<Address> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'marketCollateral',
        args: [marketId],
      });

      return result as Address;
    } catch (error) {
      throw new ContractError(
        `Failed to get market collateral: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Redeem all winning tokens for a market
   */
  async redeem(marketId: bigint): Promise<{ hash: string; payout: bigint }> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for redeeming tokens', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Simulate the contract call to get the expected payout before executing
      let payout = 0n;
      try {
        const simulatedResult = await this.client.simulateContract({
          address: this.address,
          abi: OUTCOME_TOKEN_ABI,
          functionName: 'redeem',
          args: [marketId],
          account,
        });
        payout = simulatedResult.result as bigint;
      } catch {
        // If simulation fails, payout will remain 0n
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'redeem',
        args: [marketId],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      
      return {
        hash,
        payout,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to redeem tokens: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Redeem a specific amount of winning tokens
   */
  async redeemAmount(marketId: bigint, amount: bigint): Promise<{ hash: string; payout: bigint }> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for redeeming tokens', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'redeemAmount',
        args: [marketId, amount],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });
      
      return {
        hash,
        payout: amount, // Payout equals amount redeemed (1:1)
      };
    } catch (error) {
      throw new ContractError(
        `Failed to redeem tokens: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Encode token ID from market ID and outcome ID
   */
  async encodeTokenId(marketId: bigint, outcomeId: bigint): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'encodeTokenId',
        args: [marketId, outcomeId],
      });

      return result as bigint;
    } catch (error) {
      // Fallback to client-side encoding: (marketId << 8) | outcomeId
      return (marketId << 8n) | outcomeId;
    }
  }

  /**
   * Get balance of outcome tokens for a user
   */
  async getBalance(userAddress: Address, marketId: bigint, outcomeId: bigint): Promise<bigint> {
    try {
      const tokenId = await this.encodeTokenId(marketId, outcomeId);
      
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'balanceOf',
        args: [userAddress, tokenId],
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
   * Get balances for multiple token IDs
   */
  async getBalancesBatch(userAddress: Address, tokenIds: bigint[]): Promise<bigint[]> {
    try {
      const accounts = new Array(tokenIds.length).fill(userAddress);
      
      const result = await this.client.readContract({
        address: this.address,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'balanceOfBatch',
        args: [accounts, tokenIds],
      });

      return result as bigint[];
    } catch (error) {
      throw new ContractError(
        `Failed to get balances batch: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

