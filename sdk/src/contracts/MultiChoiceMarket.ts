/**
 * MultiChoiceMarket contract interaction layer
 * Multi-outcome prediction market with LMSR pricing (3-8 outcomes)
 */

import { Address, decodeEventLog } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError, TradeParams, TradeResult } from '../types';
import { MULTI_CHOICE_MARKET_ABI, ERC20_ABI } from '../constants';

export class MultiChoiceMarket {
  private client: PublicClient;
  private walletClient?: WalletClient;
  private address: Address;
  private collateralTokenAddress?: Address;

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
   * Get collateral token address for this market
   */
  async getCollateralTokenAddress(): Promise<Address> {
    if (this.collateralTokenAddress) {
      return this.collateralTokenAddress;
    }

    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'collateralToken',
      });

      this.collateralTokenAddress = result as Address;
      return this.collateralTokenAddress;
    } catch (error) {
      throw new ContractError(
        `Failed to get collateral token address: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Ensure collateral token is approved for trading
   */
  private async ensureApproval(amount: bigint): Promise<void> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for approval', this.address);
    }

    const [account] = await this.walletClient.getAddresses();
    if (!account) {
      throw new ContractError('No account found in wallet client', this.address);
    }

    const collateralToken = await this.getCollateralTokenAddress();

    // Check current allowance
    const currentAllowance = await this.client.readContract({
      address: collateralToken,
      abi: ERC20_ABI,
      functionName: 'allowance',
      args: [account, this.address],
    }) as bigint;

    // If allowance is insufficient, approve
    if (currentAllowance < amount) {
      const hash = await this.walletClient.writeContract({
        address: collateralToken,
        abi: ERC20_ABI,
        functionName: 'approve',
        args: [this.address, amount],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
    }
  }

  /**
   * Buy outcome tokens for a specific outcome (0-7)
   * Uses LMSR pricing
   */
  async buyTokens(params: TradeParams): Promise<TradeResult> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for trading', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Validate outcomeId
      const outcomeId = typeof params.outcome === 'number' ? BigInt(params.outcome) : params.outcome;
      if (typeof outcomeId !== 'bigint') {
        throw new ContractError(`Invalid outcome: ${params.outcome}`, this.address);
      }

      // Ensure collateral token is approved
      await this.ensureApproval(params.amount);

      const minAmountOut = params.minAmountOut || 0n;

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'buy',
        args: [outcomeId, params.amount, minAmountOut],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract amountOut from SharesPurchased event
      let amountOut = 0n;
      let fee = 0n;

      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MULTI_CHOICE_MARKET_ABI,
            data: log.data,
            topics: log.topics,
          });

          if (decodedLog.eventName === 'SharesPurchased') {
            const eventArgs = decodedLog.args as {
              buyer: Address;
              outcomeId: bigint;
              shares: bigint;
              collateralPaid: bigint;
              fee: bigint;
            };

            if (
              eventArgs.buyer.toLowerCase() === account.toLowerCase() &&
              eventArgs.outcomeId === outcomeId
            ) {
              amountOut = eventArgs.shares;
              fee = eventArgs.fee;
              break;
            }
          }
        } catch {
          continue;
        }
      }

      if (amountOut === 0n) {
        throw new ContractError(
          `SharesPurchased event not found for transaction ${hash}`,
          this.address
        );
      }

      return {
        success: true,
        transactionHash: hash,
        amountIn: params.amount,
        amountOut,
        outcome: params.outcome,
        marketId: params.marketId,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to buy tokens: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Sell outcome tokens for a specific outcome (0-7)
   * Uses LMSR pricing
   * Only allowed before market close
   */
  async sellTokens(params: TradeParams): Promise<TradeResult> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for trading', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Validate outcomeId
      const outcomeId = typeof params.outcome === 'number' ? BigInt(params.outcome) : params.outcome;
      if (typeof outcomeId !== 'bigint') {
        throw new ContractError(`Invalid outcome: ${params.outcome}`, this.address);
      }

      const minAmountOut = params.minAmountOut || 0n;

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'sell',
        args: [outcomeId, params.amount, minAmountOut],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract amountOut from SharesSold event
      let amountOut = 0n;
      let fee = 0n;

      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MULTI_CHOICE_MARKET_ABI,
            data: log.data,
            topics: log.topics,
          });

          if (decodedLog.eventName === 'SharesSold') {
            const eventArgs = decodedLog.args as {
              seller: Address;
              outcomeId: bigint;
              shares: bigint;
              collateralReceived: bigint;
              fee: bigint;
            };

            if (
              eventArgs.seller.toLowerCase() === account.toLowerCase() &&
              eventArgs.outcomeId === outcomeId
            ) {
              amountOut = eventArgs.collateralReceived;
              fee = eventArgs.fee;
              break;
            }
          }
        } catch {
          continue;
        }
      }

      if (amountOut === 0n) {
        throw new ContractError(
          `SharesSold event not found for transaction ${hash}`,
          this.address
        );
      }

      return {
        success: true,
        transactionHash: hash,
        amountIn: params.amount,
        amountOut,
        outcome: params.outcome,
        marketId: params.marketId,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to sell tokens: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get price for a specific outcome (LMSR price)
   * Returns price as basis points (10000 = 100%)
   */
  async getPrice(outcomeId: bigint): Promise<bigint> {
    try {
      const price = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getPrice',
        args: [outcomeId],
      }) as bigint;

      return price;
    } catch (error) {
      throw new ContractError(
        `Failed to get price for outcome ${outcomeId}: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get prices for all outcomes
   * Returns array of prices in basis points
   */
  async getAllPrices(): Promise<bigint[]> {
    try {
      const prices = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getAllPrices',
      }) as bigint[];

      return prices;
    } catch (error) {
      throw new ContractError(
        `Failed to get all prices: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get reserves for all outcomes
   * Used for LMSR calculations
   */
  async getOutcomeReserves(): Promise<bigint[]> {
    try {
      const reserves = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getOutcomeReserves',
      }) as bigint[];

      return reserves;
    } catch (error) {
      throw new ContractError(
        `Failed to get outcome reserves: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get quote for buying outcome tokens
   * Returns expected tokens out and fee
   */
  async getQuoteBuy(
    outcomeId: bigint,
    collateralIn: bigint,
    user: Address
  ): Promise<{ tokensOut: bigint; fee: bigint }> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getQuoteBuy',
        args: [outcomeId, collateralIn, user],
      }) as [bigint, bigint];

      return {
        tokensOut: result[0],
        fee: result[1],
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get buy quote: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get quote for selling outcome tokens
   * Returns expected collateral out and fee
   */
  async getQuoteSell(
    outcomeId: bigint,
    tokensIn: bigint,
    user: Address
  ): Promise<{ collateralOut: bigint; fee: bigint }> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getQuoteSell',
        args: [outcomeId, tokensIn, user],
      }) as [bigint, bigint];

      return {
        collateralOut: result[0],
        fee: result[1],
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get sell quote: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get outcome count for this market
   */
  async getOutcomeCount(): Promise<bigint> {
    try {
      const count = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getOutcomeCount',
      }) as bigint;

      return count;
    } catch (error) {
      throw new ContractError(
        `Failed to get outcome count: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Add liquidity to the market
   * Returns LP tokens received
   */
  async addLiquidity(amount: bigint): Promise<{ hash: Address; lpTokens: bigint }> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for adding liquidity', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Ensure collateral token is approved
      await this.ensureApproval(amount);

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'addLiquidity',
        args: [amount],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract LP tokens from LiquidityAdded event
      let lpTokens = 0n;

      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MULTI_CHOICE_MARKET_ABI,
            data: log.data,
            topics: log.topics,
          });

          if (decodedLog.eventName === 'LiquidityAdded') {
            const eventArgs = decodedLog.args as {
              provider: Address;
              amount: bigint;
              lpTokens: bigint;
            };

            if (eventArgs.provider.toLowerCase() === account.toLowerCase()) {
              lpTokens = eventArgs.lpTokens;
              break;
            }
          }
        } catch {
          continue;
        }
      }

      return { hash, lpTokens };
    } catch (error) {
      throw new ContractError(
        `Failed to add liquidity: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Remove liquidity from the market
   * Returns collateral received
   */
  async removeLiquidity(lpTokens: bigint): Promise<{ hash: Address; collateralOut: bigint }> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for removing liquidity', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'removeLiquidity',
        args: [lpTokens],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract collateral from LiquidityRemoved event
      let collateralOut = 0n;

      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MULTI_CHOICE_MARKET_ABI,
            data: log.data,
            topics: log.topics,
          });

          if (decodedLog.eventName === 'LiquidityRemoved') {
            const eventArgs = decodedLog.args as {
              provider: Address;
              lpTokens: bigint;
              collateralAmount: bigint;
            };

            if (eventArgs.provider.toLowerCase() === account.toLowerCase()) {
              collateralOut = eventArgs.collateralAmount;
              break;
            }
          }
        } catch {
          continue;
        }
      }

      return { hash, collateralOut };
    } catch (error) {
      throw new ContractError(
        `Failed to remove liquidity: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market information
   */
  async getMarketInfo(): Promise<{
    marketId: bigint;
    marketType: number;
    collateralToken: Address;
    closeTime: bigint;
    outcomeCount: bigint;
    isResolved: boolean;
    isPaused: boolean;
  }> {
    try {
      const info = await this.client.readContract({
        address: this.address,
        abi: MULTI_CHOICE_MARKET_ABI,
        functionName: 'getMarketInfo',
      }) as {
        marketId: bigint;
        marketType: number;
        collateralToken: Address;
        closeTime: bigint;
        outcomeCount: bigint;
        isResolved: boolean;
        isPaused: boolean;
      };

      return info;
    } catch (error) {
      throw new ContractError(
        `Failed to get market info: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}
