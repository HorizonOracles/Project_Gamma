/**
 * MarketAMM contract interaction layer
 */

import { Address, decodeEventLog } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError, TradeParams, TradeResult, MarketOutcome, MarketPrices } from '../types';
import { MARKET_AMM_ABI, OUTCOME_TOKEN_ABI, ERC20_ABI } from '../constants';

export class MarketAMM {
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
        abi: MARKET_AMM_ABI,
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
        abi: MARKET_AMM_ABI,
        functionName,
        args: [params.amount, minAmountOut],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract amountOut from Trade event
      let amountOut = 0n;
      let fee = 0n;
      
      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MARKET_AMM_ABI,
            data: log.data,
            topics: log.topics,
          });
          
          if (decodedLog.eventName === 'Trade') {
            const eventArgs = decodedLog.args as {
              trader: Address;
              buyYes: boolean;
              collateralIn: bigint;
              tokensOut: bigint;
              fee: bigint;
              price: bigint;
            };
            
            // Check if this is our trade (matching trader and outcome)
            if (
              eventArgs.trader.toLowerCase() === account.toLowerCase() &&
              eventArgs.buyYes === (params.outcome === 'YES')
            ) {
              amountOut = eventArgs.tokensOut;
              fee = eventArgs.fee;
              break;
            }
          }
        } catch {
          // Not the event we're looking for, continue
          continue;
        }
      }

      if (amountOut === 0n) {
        throw new ContractError(
          `Trade event not found or amountOut is zero for transaction ${hash}. ` +
          'This may indicate the transaction failed.',
          this.address
        );
      }

      // Normalize outcome to MarketOutcome type
      const normalizedOutcome: MarketOutcome = typeof params.outcome === 'string' 
        ? params.outcome 
        : params.outcome === 0n 
          ? 'YES' 
          : 'NO';

      return {
        success: true,
        transactionHash: hash,
        amountIn: params.amount,
        amountOut,
        outcome: normalizedOutcome,
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
        abi: MARKET_AMM_ABI,
        functionName,
        args: [params.amount, minAmountOut],
        account,
        chain: this.walletClient.chain,
      });

      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract amountOut from Trade event
      let amountOut = 0n;
      let fee = 0n;
      
      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MARKET_AMM_ABI,
            data: log.data,
            topics: log.topics,
          });
          
          if (decodedLog.eventName === 'Trade') {
            const eventArgs = decodedLog.args as {
              trader: Address;
              buyYes: boolean;
              collateralIn: bigint;
              tokensOut: bigint;
              fee: bigint;
              price: bigint;
            };
            
            // For sell operations:
            // - buyYes = false
            // - collateralIn = 0 (no collateral sent in)
            // - tokensOut = collateral amount received
            // Check if this is our sell trade
            if (
              eventArgs.trader.toLowerCase() === account.toLowerCase() &&
              eventArgs.buyYes === false &&
              eventArgs.collateralIn === 0n
            ) {
              // For sell operations, tokensOut represents collateralOut
              amountOut = eventArgs.tokensOut;
              fee = eventArgs.fee;
              break;
            }
          }
        } catch {
          // Not the event we're looking for, continue
          continue;
        }
      }

      if (amountOut === 0n) {
        throw new ContractError(
          `Trade event not found or amountOut is zero for transaction ${hash}. ` +
          'This may indicate the transaction failed or insufficient tokens were sold.',
          this.address
        );
      }

      // Normalize outcome to MarketOutcome type
      const normalizedOutcome: MarketOutcome = typeof params.outcome === 'string' 
        ? params.outcome 
        : params.outcome === 0n 
          ? 'YES' 
          : 'NO';

      return {
        success: true,
        transactionHash: hash,
        amountIn: params.amount,
        amountOut,
        outcome: normalizedOutcome,
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
   */
  async getPrice(marketId: bigint, outcome: MarketOutcome): Promise<bigint> {
    try {
      const functionName = outcome === 'YES' ? 'getYesPrice' : 'getNoPrice';
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_AMM_ABI,
        functionName,
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
   * Get market prices for both outcomes
   */
  async getMarketPrices(marketId: bigint): Promise<MarketPrices> {
    try {
      const [yesPrice, noPrice] = await Promise.all([
        this.client.readContract({
          address: this.address,
          abi: MARKET_AMM_ABI,
          functionName: 'getYesPrice',
        }),
        this.client.readContract({
          address: this.address,
          abi: MARKET_AMM_ABI,
          functionName: 'getNoPrice',
        }),
      ]);

      return {
        yesPrice: yesPrice as bigint,
        noPrice: noPrice as bigint,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get market prices: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get liquidity amounts (reserves)
   */
  async getLiquidity(marketId: bigint): Promise<{ yes: bigint; no: bigint }> {
    try {
      const [yesReserve, noReserve] = await Promise.all([
        this.client.readContract({
          address: this.address,
          abi: MARKET_AMM_ABI,
          functionName: 'reserveYes',
        }),
        this.client.readContract({
          address: this.address,
          abi: MARKET_AMM_ABI,
          functionName: 'reserveNo',
        }),
      ]);

      return {
        yes: yesReserve as bigint,
        no: noReserve as bigint,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get liquidity: ${error instanceof Error ? error.message : 'Unknown error'}`,
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
        abi: MARKET_AMM_ABI,
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
      // Try to read totalSupply, if not available, sum all balances via events
      // For now, we'll use a workaround: if LP tokens follow ERC20, totalSupply should exist
      // Otherwise, we estimate from liquidity values
      const result = await this.client.readContract({
        address: this.address,
        abi: [
          ...MARKET_AMM_ABI,
          {
            type: 'function',
            name: 'totalSupply',
            inputs: [],
            outputs: [{ name: '', type: 'uint256' }],
            stateMutability: 'view',
          },
        ],
        functionName: 'totalSupply',
      }).catch(() => {
        // If totalSupply doesn't exist, estimate from reserves
        // This is an approximation for AMM pools
        return null;
      });

      if (result !== null) {
        return result as bigint;
      }

      // Fallback: estimate from liquidity (square root of product for constant product AMM)
      const liquidity = await this.getLiquidity(0n); // marketId not needed for liquidity
      
      // For constant product AMM: LP tokens ≈ sqrt(x * y)
      // Calculate geometric mean using BigInt arithmetic
      if (liquidity.yes > 0n && liquidity.no > 0n) {
        const product = liquidity.yes * liquidity.no;
        // Approximate square root using Babylonian method
        let sqrt = product;
        if (sqrt > 0n) {
          let x = sqrt;
          while (x * x > product) {
            x = (x + product / x) / 2n;
          }
          sqrt = x;
        }
        return sqrt;
      }
      
      return 0n;
    } catch (error) {
      throw new ContractError(
        `Failed to get total LP supply: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get user balance for a specific outcome token
   * Note: This uses the OutcomeToken contract (ERC-1155), not the AMM contract
   */
  async getUserBalance(
    userAddress: Address,
    marketId: bigint,
    outcome: MarketOutcome,
    outcomeTokenAddress: Address
  ): Promise<bigint> {
    try {
      // Outcome IDs: 0 = YES, 1 = NO
      const outcomeId = outcome === 'YES' ? 0n : 1n;
      
      // Token ID encoding: (marketId << 8) | outcomeId
      const tokenId = (marketId << 8n) | outcomeId;
      
      const result = await this.client.readContract({
        address: outcomeTokenAddress,
        abi: OUTCOME_TOKEN_ABI,
        functionName: 'balanceOf',
        args: [userAddress, tokenId],
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get user balance: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Add liquidity to a market
   * @param amount Amount of collateral tokens to add as liquidity
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
        abi: MARKET_AMM_ABI,
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
   * Remove liquidity from a market
   * @param lpTokens Amount of LP tokens to burn
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
        abi: MARKET_AMM_ABI,
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
   * Get quote for buying tokens
   */
  async getBuyQuote(
    collateralIn: bigint,
    outcome: MarketOutcome,
    userAddress: Address
  ): Promise<{ tokensOut: bigint; fee: bigint }> {
    try {
      const functionName = outcome === 'YES' ? 'getQuoteBuyYes' : 'getQuoteBuyNo';
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_AMM_ABI,
        functionName,
        args: [collateralIn, userAddress],
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
   * Get quote for selling tokens
   */
  async getSellQuote(
    tokensIn: bigint,
    outcome: MarketOutcome,
    userAddress: Address
  ): Promise<{ collateralOut: bigint; fee: bigint }> {
    try {
      const functionName = outcome === 'YES' ? 'getQuoteSellYes' : 'getQuoteSellNo';
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_AMM_ABI,
        functionName,
        args: [tokensIn, userAddress],
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
   * Fund redemptions - transfers collateral to OutcomeToken after market resolution
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
        abi: MARKET_AMM_ABI,
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
}
