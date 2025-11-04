/**
 * BinaryMarket contract interaction layer
 * Binary prediction market with static 1:1 pricing and 2% fixed fee
 */

import { Address, decodeEventLog } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError, TradeParams, TradeResult, MarketOutcome } from '../types';
import { BINARY_MARKET_ABI, ERC20_ABI } from '../constants';

export class BinaryMarket {
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
        abi: BINARY_MARKET_ABI,
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
   * Buy outcome tokens (YES or NO)
   * Static 1:1 pricing with 2% fee
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

      // Ensure collateral token is approved
      await this.ensureApproval(params.amount);

      const minAmountOut = params.minAmountOut || 0n;
      const functionName = params.outcome === 'YES' ? 'buyYes' : 'buyNo';

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName,
        args: [params.amount, minAmountOut],
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
            abi: BINARY_MARKET_ABI,
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
            
            const expectedOutcomeId = params.outcome === 'YES' ? 0n : 1n;
            
            if (
              eventArgs.buyer.toLowerCase() === account.toLowerCase() &&
              eventArgs.outcomeId === expectedOutcomeId
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
        outcome: params.outcome as MarketOutcome,
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
   * Sell outcome tokens (YES or NO)
   * Static 1:1 pricing with 2% fee
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

      const minAmountOut = params.minAmountOut || 0n;
      const functionName = params.outcome === 'YES' ? 'sellYes' : 'sellNo';

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName,
        args: [params.amount, minAmountOut],
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
            abi: BINARY_MARKET_ABI,
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
            
            const expectedOutcomeId = params.outcome === 'YES' ? 0n : 1n;
            
            if (
              eventArgs.seller.toLowerCase() === account.toLowerCase() &&
              eventArgs.outcomeId === expectedOutcomeId
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
        outcome: params.outcome as MarketOutcome,
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
   * Get current price for an outcome
   * BinaryMarket uses static 50/50 pricing (always returns 0.5)
   */
  async getPrice(outcomeId: bigint): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'getPrice',
        args: [outcomeId],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get price: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get buy quote with fee calculation
   */
  async getBuyQuote(
    collateralIn: bigint,
    outcome: MarketOutcome,
    userAddress: Address
  ): Promise<{ tokensOut: bigint; fee: bigint }> {
    try {
      const outcomeId = outcome === 'YES' ? 0n : 1n;
      
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'getQuoteBuy',
        args: [outcomeId, collateralIn, userAddress],
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
   * Get sell quote with fee calculation
   */
  async getSellQuote(
    tokensIn: bigint,
    outcome: MarketOutcome,
    userAddress: Address
  ): Promise<{ collateralOut: bigint; fee: bigint }> {
    try {
      const outcomeId = outcome === 'YES' ? 0n : 1n;
      
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'getQuoteSell',
        args: [outcomeId, tokensIn, userAddress],
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
   * Get liquidity reserves (pools)
   */
  async getReserves(): Promise<{ yes: bigint; no: bigint }> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'getReserves',
      }) as [bigint, bigint];

      return {
        yes: result[0],
        no: result[1],
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get reserves: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get user's LP token balance
   */
  async getLPBalance(userAddress: Address): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'balanceOf',
        args: [userAddress],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get LP balance: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get total LP token supply
   */
  async getTotalLPSupply(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'totalSupply',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get total LP supply: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Add liquidity to the market
   */
  async addLiquidity(amount: bigint): Promise<string> {
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
        abi: BINARY_MARKET_ABI,
        functionName: 'addLiquidity',
        args: [amount],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
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
   */
  async removeLiquidity(lpTokens: bigint): Promise<string> {
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
        abi: BINARY_MARKET_ABI,
        functionName: 'removeLiquidity',
        args: [lpTokens],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to remove liquidity: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Fund redemptions after market resolution
   * Transfers collateral to OutcomeToken for 1:1 redemptions
   */
  async fundRedemptions(): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for funding redemptions', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'fundRedemptions',
        args: [],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to fund redemptions: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market ID
   */
  async getMarketId(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'marketId',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get market ID: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market close time
   */
  async getCloseTime(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'closeTime',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get close time: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get total collateral in market
   */
  async getTotalCollateral(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: BINARY_MARKET_ABI,
        functionName: 'totalCollateral',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get total collateral: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}
