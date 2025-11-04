/**
 * Unit tests for MarketFactory contract interaction
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { MarketFactory } from '../../contracts/MarketFactory';
import { ContractError, CreateMarketParams, MarketStatus } from '../../types';
import {
  mockPublicClient,
  mockWalletClient,
  mockAddress,
  mockMarketId,
  mockTransactionHash,
  createMockReceipt,
  createMockEventLog,
} from '../mocks/viem';
import { decodeEventLog } from 'viem';

vi.mock('viem', async () => {
  const actual = await vi.importActual('viem');
  return {
    ...actual,
    decodeEventLog: vi.fn(),
  };
});

vi.mock('../../utils/markets', () => ({
  getMarketContract: vi.fn(),
}));

describe('MarketFactory', () => {
  let marketFactory: MarketFactory;
  const factoryAddress = '0x22Cc806047BB825aa26b766Af737E92B1866E8A6' as const;

  beforeEach(() => {
    vi.clearAllMocks();
    marketFactory = new MarketFactory(
      mockPublicClient as any,
      factoryAddress,
      mockWalletClient as any
    );
  });

  describe('createMarket', () => {
    const createParams: CreateMarketParams & {
      collateralToken: string;
      category: string;
      metadataURI: string;
      creatorStake: bigint;
    } = {
      question: 'Will Bitcoin reach $100k?',
      description: 'Test market',
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      collateralToken: '0xToken1' as const,
      category: 'crypto',
      metadataURI: 'ipfs://test',
      creatorStake: 0n,
    };

    it('should throw error if wallet client not provided', async () => {
      const factoryWithoutWallet = new MarketFactory(
        mockPublicClient as any,
        factoryAddress
      );

      await expect(
        factoryWithoutWallet.createMarket(createParams)
      ).rejects.toThrow(ContractError);
    });

    it('should create market successfully', async () => {
      // Mock getAddresses to return an array
      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);
      
      const mockEvent = createMockEventLog('MarketCreated', {
        marketId: mockMarketId,
        creator: mockAddress,
        question: createParams.question,
      });

      vi.mocked(decodeEventLog).mockReturnValue({
        eventName: 'MarketCreated',
        args: {
          marketId: mockMarketId,
          creator: mockAddress,
          question: createParams.question,
        },
      } as any);

      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([mockEvent])
      );

      const result = await marketFactory.createMarket(createParams);

      expect(result).toBe(mockMarketId);
      expect(mockWalletClient.writeContract).toHaveBeenCalled();
      const callArgs = mockWalletClient.writeContract.mock.calls[0][0];
      expect(callArgs.address).toBe(factoryAddress);
      expect(callArgs.functionName).toBe('createMarket');
      expect(callArgs.args).toEqual([{
        collateralToken: createParams.collateralToken,
        closeTime: createParams.endTime,
        category: createParams.category,
        metadataURI: createParams.metadataURI,
        creatorStake: createParams.creatorStake,
      }]);
      expect(callArgs.account).toBe(mockAddress);
      expect(Array.isArray(callArgs.abi)).toBe(true);
    });

    it('should throw error if MarketCreated event not found', async () => {
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);
      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([])
      );

      vi.mocked(decodeEventLog).mockImplementation(() => {
        throw new Error('Not the event');
      });

      await expect(
        marketFactory.createMarket(createParams)
      ).rejects.toThrow(ContractError);
    });

    it('should handle transaction errors', async () => {
      mockWalletClient.writeContract.mockRejectedValue(
        new Error('Transaction failed')
      );

      await expect(
        marketFactory.createMarket(createParams)
      ).rejects.toThrow(ContractError);
    });
  });

  describe('getMarket', () => {
    beforeEach(async () => {
      // Mock getMarketContract
      const marketsModule = await import('../../utils/markets');
      const mockMarketContract = {
        getReserves: vi.fn().mockResolvedValue({ yes: 1000000n, no: 1000000n }),
      };
      vi.mocked(marketsModule.getMarketContract).mockResolvedValue(mockMarketContract as any);
    });

    it('should fetch market information correctly', async () => {
      const mockResult = {
        id: mockMarketId,
        creator: mockAddress,
        amm: '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'crypto',
        metadataURI: 'ipfs://test',
        creatorStake: 0n,
        stakeRefunded: false,
        status: 0, // ACTIVE status
      };

      mockPublicClient.readContract
        .mockResolvedValueOnce(mockResult) // getMarket
        .mockResolvedValueOnce({ marketType: 0, outcomeCount: 2n }); // getMarketInfo
      mockPublicClient.getLogs = vi.fn().mockResolvedValue([]);
      mockPublicClient.getBlock = vi.fn().mockResolvedValue({ timestamp: BigInt(Math.floor(Date.now() / 1000)) } as any);

      const result = await marketFactory.getMarket(mockMarketId);

      expect(result.marketId).toBe(mockMarketId);
      expect(result.question).toBe(mockResult.metadataURI);
      expect(result.marketAddress).toBe(mockResult.amm);
      expect(result.status).toBe(MarketStatus.Active);
      expect(result.yesTokenId).toBe((mockMarketId << 8n) | 0n);
      expect(result.noTokenId).toBe((mockMarketId << 8n) | 1n);
    });

    it('should map status codes correctly', async () => {
      const statusTests = [
        [0, MarketStatus.Active],
        [1, MarketStatus.Closed],
        [2, MarketStatus.Resolved],
        [3, MarketStatus.Invalid],
      ];

      for (const [statusCode, expectedStatus] of statusTests) {
        const mockMarketStruct = {
          id: mockMarketId,
          creator: mockAddress,
          amm: '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd' as const,
          collateralToken: '0xToken1' as const,
          closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
          category: 'test',
          metadataURI: 'ipfs://test',
          creatorStake: 0n,
          stakeRefunded: false,
          status: statusCode,
        };

        // Reset mocks for each iteration
        const marketsModule = await import('../../utils/markets');
        const mockMarketContract = {
          getReserves: vi.fn().mockResolvedValue({ yes: 0n, no: 0n }),
        };
        vi.mocked(marketsModule.getMarketContract).mockResolvedValue(mockMarketContract as any);

        mockPublicClient.readContract
          .mockResolvedValueOnce(mockMarketStruct) // getMarket
          .mockResolvedValueOnce({ marketType: 0, outcomeCount: 2n }); // getMarketInfo
        mockPublicClient.getLogs = vi.fn().mockResolvedValue([]);

        const result = await marketFactory.getMarket(mockMarketId);
        expect(result.status).toBe(expectedStatus);
      }
    });

    it('should handle contract read errors', async () => {
      mockPublicClient.readContract.mockRejectedValue(
        new Error('Contract read failed')
      );

      await expect(
        marketFactory.getMarket(mockMarketId)
      ).rejects.toThrow(ContractError);
    });
  });

  describe('getAllMarkets', () => {
    it('should fetch all market IDs', async () => {
      const mockMarkets = [1n, 2n, 3n];
      mockPublicClient.readContract.mockResolvedValue(mockMarkets);

      const result = await marketFactory.getAllMarkets();

      expect(result).toEqual(mockMarkets);
      expect(mockPublicClient.readContract).toHaveBeenCalled();
      const callArgs = mockPublicClient.readContract.mock.calls[0][0];
      expect(callArgs.address).toBe(factoryAddress);
      expect(callArgs.functionName).toBe('getAllMarketIds');
      expect(Array.isArray(callArgs.abi)).toBe(true);
    });

    it('should handle empty markets array', async () => {
      mockPublicClient.readContract.mockResolvedValue([]);

      const result = await marketFactory.getAllMarkets();
      expect(result).toEqual([]);
    });

    it('should handle contract read errors', async () => {
      mockPublicClient.readContract.mockRejectedValue(
        new Error('Contract read failed')
      );

      await expect(
        marketFactory.getAllMarkets()
      ).rejects.toThrow(ContractError);
    });
  });
});

