/**
 * MarketFactory contract interaction layer
 */

import { Address, decodeEventLog } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError, CreateMarketParams, MarketInfo, MarketStatus, MarketType } from '../types';
import { MARKET_FACTORY_ABI, MARKET_AMM_ABI, I_MARKET_ABI } from '../constants';

/**
 * MarketParams as expected by the MarketFactory contract
 */
export interface MarketParams {
  collateralToken: Address;
  closeTime: bigint;
  category: string;
  metadataURI: string;
  creatorStake: bigint;
}

/**
 * Market struct as returned by the contract
 */
export interface MarketStruct {
  id: bigint;
  creator: Address;
  amm: Address;
  collateralToken: Address;
  closeTime: bigint;
  category: string;
  metadataURI: string;
  creatorStake: bigint;
  stakeRefunded: boolean;
  status: number; // 0=Active, 1=Closed, 2=Resolved, 3=Invalid
}

export class MarketFactory {
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
   * Create a new prediction market
   * Note: SDK's CreateMarketParams needs to be mapped to contract's MarketParams
   */
  async createMarket(params: CreateMarketParams & {
    collateralToken: Address;
    category: string;
    metadataURI: string;
    creatorStake: bigint;
  }): Promise<bigint> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for creating markets', this.address);
    }

    try {
      // Get account from wallet client
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Map SDK params to contract MarketParams tuple
      // viem expects tuple parameters as an array
      const marketParams: [
        Address, // collateralToken
        bigint,  // closeTime
        string,  // category
        string,  // metadataURI
        bigint   // creatorStake
      ] = [
        params.collateralToken,
        params.endTime,
        params.category,
        params.metadataURI,
        params.creatorStake,
      ];

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'createMarket',
        args: [{ 
          collateralToken: marketParams[0],
          closeTime: marketParams[1],
          category: marketParams[2],
          metadataURI: marketParams[3],
          creatorStake: marketParams[4],
        }],
        account,
        chain: this.walletClient.chain,
      });

      // Wait for transaction receipt
      const receipt = await this.client.waitForTransactionReceipt({ hash });

      // Extract market ID from MarketCreated event
      let marketId: bigint | undefined;
      
      for (const log of receipt.logs) {
        try {
          const decodedLog = decodeEventLog({
            abi: MARKET_FACTORY_ABI,
            data: log.data,
            topics: log.topics,
          });
          
          if (decodedLog.eventName === 'MarketCreated') {
            marketId = decodedLog.args.marketId as bigint;
            break;
          }
        } catch {
          // Not the event we're looking for, continue
          continue;
        }
      }

      if (marketId === undefined) {
        throw new ContractError(
          `MarketCreated event not found in transaction receipt for hash ${hash}. ` +
          'The transaction may have failed or emitted events in a different format.',
          this.address
        );
      }

      return marketId;
    } catch (error) {
      throw new ContractError(
        `Failed to create market: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market information
   */
  async getMarket(marketId: bigint): Promise<MarketInfo> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [marketId],
      }) as MarketStruct;
      
      // Map numeric status to MarketStatus enum
      let marketStatus: MarketStatus;
      switch (result.status) {
        case 0:
          marketStatus = MarketStatus.Active;
          break;
        case 1:
          marketStatus = MarketStatus.Closed; // Closed
          break;
        case 2:
          marketStatus = MarketStatus.Resolved;
          break;
        case 3:
          marketStatus = MarketStatus.Invalid; // Invalid
          break;
        default:
          marketStatus = MarketStatus.Active;
      }
      
      // Fetch liquidity from MarketAMM contract
      let liquidity = { yes: 0n, no: 0n };
      try {
        const [yesReserve, noReserve] = await Promise.all([
          this.client.readContract({
            address: result.amm,
            abi: MARKET_AMM_ABI,
            functionName: 'reserveYes',
          }),
          this.client.readContract({
            address: result.amm,
            abi: MARKET_AMM_ABI,
            functionName: 'reserveNo',
          }),
        ]);
        liquidity = {
          yes: yesReserve as bigint,
          no: noReserve as bigint,
        };
      } catch {
        // If liquidity fetch fails, keep defaults
      }
      
      // Fetch createdAt from MarketCreated event logs
      let createdAt = 0n;
      try {
        const marketCreatedEvent = MARKET_FACTORY_ABI.find(
          (item) => item.type === 'event' && item.name === 'MarketCreated'
        );
        if (marketCreatedEvent) {
          const logs = await this.client.getLogs({
            address: this.address,
            event: marketCreatedEvent as any,
            args: {
              marketId: marketId,
            },
            fromBlock: 0n,
          });
          
          if (logs.length > 0) {
            const block = await this.client.getBlock({ blockNumber: logs[0].blockNumber });
            createdAt = BigInt(block.timestamp);
          }
        }
      } catch {
        // If event log fetch fails, keep default
      }
      
      // Calculate total volume from Trade events
      let totalVolume = 0n;
      try {
        const tradeEvent = MARKET_AMM_ABI.find(
          (item) => item.type === 'event' && item.name === 'Trade'
        );
        if (tradeEvent) {
          const tradeLogs = await this.client.getLogs({
            address: result.amm,
            event: tradeEvent as any,
            fromBlock: 0n,
          });
          
          // Sum up all collateralIn amounts
          for (const log of tradeLogs) {
            try {
              const decodedLog = decodeEventLog({
                abi: MARKET_AMM_ABI,
                data: log.data,
                topics: log.topics,
              });
              
              if (decodedLog.eventName === 'Trade') {
                const eventArgs = decodedLog.args as {
                  collateralIn: bigint;
                };
                totalVolume += eventArgs.collateralIn || 0n;
              }
            } catch {
              // Skip invalid logs
              continue;
            }
          }
        }
      } catch {
        // If volume calculation fails, keep default
      }

      // Get market type and outcome count from IMarket interface
      let marketType: MarketType | undefined;
      let outcomeCount: bigint | undefined;
      try {
        const marketInfo = await this.client.readContract({
          address: result.amm,
          abi: I_MARKET_ABI,
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
        
        marketType = marketInfo.marketType as MarketType;
        outcomeCount = marketInfo.outcomeCount;
      } catch {
        // If IMarket interface not available, default to Binary (MarketAMM)
        marketType = MarketType.Binary;
        outcomeCount = 2n;
      }
      
      return {
        marketId: result.id,
        marketAddress: result.amm,
        question: result.metadataURI,
        description: result.metadataURI,
        creator: result.creator,
        endTime: result.closeTime,
        status: marketStatus,
        yesTokenId: (marketId << 8n) | 0n,
        noTokenId: (marketId << 8n) | 1n,
        totalVolume,
        totalLiquidity: liquidity,
        createdAt,
        marketType,
        outcomeCount,
        // Additional fields from MarketStruct
        collateralToken: result.collateralToken,
        category: result.category,
        metadataURI: result.metadataURI,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get market: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get all market IDs
   */
  async getAllMarkets(): Promise<bigint[]> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getAllMarketIds',
      });

      return result as bigint[];
    } catch (error) {
      throw new ContractError(
        `Failed to get all markets: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market count
   */
  async getMarketCount(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarketCount',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get market count: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get markets with pagination
   */
  async getMarkets(offset: bigint, limit: bigint): Promise<MarketStruct[]> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarkets',
        args: [offset, limit],
      });

      return result as MarketStruct[];
    } catch (error) {
      throw new ContractError(
        `Failed to get markets: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get active markets with pagination
   */
  async getActiveMarkets(offset: bigint, limit: bigint): Promise<MarketStruct[]> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getActiveMarkets',
        args: [offset, limit],
      });

      return result as MarketStruct[];
    } catch (error) {
      throw new ContractError(
        `Failed to get active markets: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market IDs by category
   */
  async getMarketIdsByCategory(category: string): Promise<bigint[]> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarketIdsByCategory',
        args: [category],
      });

      return result as bigint[];
    } catch (error) {
      throw new ContractError(
        `Failed to get market IDs by category: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get market type for a specific market
   * @param marketAddress The market contract address
   * @returns MarketType enum value
   */
  async getMarketType(marketAddress: Address): Promise<MarketType> {
    try {
      const marketType = await this.client.readContract({
        address: marketAddress,
        abi: I_MARKET_ABI,
        functionName: 'getMarketType',
      }) as number;
      
      return marketType as MarketType;
    } catch (error) {
      // Default to Binary if IMarket interface not available
      return MarketType.Binary;
    }
  }

  /**
   * Get market info including type and outcome count
   * @param marketAddress The market contract address
   * @returns Market info with type and outcome count
   */
  async getMarketInfo(marketAddress: Address): Promise<{
    marketId: bigint;
    marketType: MarketType;
    collateralToken: Address;
    closeTime: bigint;
    outcomeCount: bigint;
    isResolved: boolean;
    isPaused: boolean;
  }> {
    try {
      const info = await this.client.readContract({
        address: marketAddress,
        abi: I_MARKET_ABI,
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
      
      return {
        ...info,
        marketType: info.marketType as MarketType,
      };
    } catch (error) {
      throw new ContractError(
        `Failed to get market info: ${error instanceof Error ? error.message : 'Unknown error'}`,
        marketAddress,
        error
      );
    }
  }

  /**
   * Get market IDs by creator
   */
  async getMarketIdsByCreator(creator: Address): Promise<bigint[]> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarketIdsByCreator',
        args: [creator],
      });

      return result as bigint[];
    } catch (error) {
      throw new ContractError(
        `Failed to get market IDs by creator: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Check if market exists
   */
  async marketExists(marketId: bigint): Promise<boolean> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'marketExists',
        args: [marketId],
      });

      return result as boolean;
    } catch (error) {
      throw new ContractError(
        `Failed to check market existence: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get minimum creator stake
   */
  async getMinCreatorStake(): Promise<bigint> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'minCreatorStake',
      });

      return result as bigint;
    } catch (error) {
      throw new ContractError(
        `Failed to get min creator stake: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Refund creator stake
   */
  async refundCreatorStake(marketId: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for refunding stake', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'refundCreatorStake',
        args: [marketId],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to refund creator stake: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Update market status
   */
  async updateMarketStatus(marketId: bigint): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for updating market status', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: MARKET_FACTORY_ABI,
        functionName: 'updateMarketStatus',
        args: [marketId],
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to update market status: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

