/**
 * Unit tests for MarketAMM contract interaction
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { MarketAMM } from '../../contracts/MarketAMM';
import { ContractError, TradeParams } from '../../types';
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

describe('MarketAMM', () => {
  let marketAMM: MarketAMM;
  const ammAddress = '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd' as const;
  const outcomeTokenAddress = '0x17B322784265c105a94e4c3d00aF1E5f46a5F311' as const;

  beforeEach(() => {
    vi.clearAllMocks();
    marketAMM = new MarketAMM(
      mockPublicClient as any,
      ammAddress,
      mockWalletClient as any
    );
  });

  describe('buyTokens', () => {
    const tradeParams: TradeParams = {
      marketId: mockMarketId,
      outcome: 'YES',
      amount: 1000000000000000000n, // 1 BNB
      minAmountOut: 950000000000000000n,
    };

    it('should throw error if wallet client not provided', async () => {
      const ammWithoutWallet = new MarketAMM(
        mockPublicClient as any,
        ammAddress
      );

      await expect(
        ammWithoutWallet.buyTokens(tradeParams)
      ).rejects.toThrow(ContractError);
    });

    it('should buy tokens successfully', async () => {
      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);

      // Mock collateral token address and allowance
      mockPublicClient.readContract
        .mockResolvedValueOnce('0xCollateralToken' as any) // collateralToken()
        .mockResolvedValueOnce(0n as any); // allowance()

      const mockEvent = createMockEventLog('Trade', {
        trader: mockAddress,
        buyYes: true,
        collateralIn: tradeParams.amount,
        tokensOut: 980000000000000000n,
        fee: 20000000000000000n,
        price: 1020408163265306122n,
      });

      vi.mocked(decodeEventLog).mockReturnValue({
        eventName: 'Trade',
        args: {
          trader: mockAddress,
          buyYes: true,
          collateralIn: tradeParams.amount,
          tokensOut: 980000000000000000n,
          fee: 20000000000000000n,
          price: 1020408163265306122n,
        },
      } as any);

      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([mockEvent])
      );

      const result = await marketAMM.buyTokens(tradeParams);

      expect(result.success).toBe(true);
      expect(result.transactionHash).toBe(mockTransactionHash);
      expect(result.amountIn).toBe(tradeParams.amount);
      expect(result.amountOut).toBe(980000000000000000n);
      expect(result.outcome).toBe('YES');
      expect(result.marketId).toBe(mockMarketId);

      // Should approve first, then buy
      expect(mockWalletClient.writeContract).toHaveBeenCalledTimes(2);
      const buyCall = mockWalletClient.writeContract.mock.calls[1][0];
      expect(buyCall.address).toBe(ammAddress);
      expect(buyCall.functionName).toBe('buyYes');
      expect(buyCall.args).toEqual([tradeParams.amount, tradeParams.minAmountOut]);
      expect(buyCall.account).toBe(mockAddress);
    });

    it('should use zero minAmountOut if not provided', async () => {
      const paramsWithoutMin = { ...tradeParams };
      delete paramsWithoutMin.minAmountOut;

      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);

      // Mock collateral token address and allowance
      mockPublicClient.readContract
        .mockResolvedValueOnce('0xCollateralToken' as any)
        .mockResolvedValueOnce(0n as any);
      
      const mockEvent = createMockEventLog('Trade', {
        trader: mockAddress,
        buyYes: true,
        collateralIn: tradeParams.amount,
        tokensOut: 980000000000000000n,
        fee: 20000000000000000n,
        price: 1020408163265306122n,
      });

      vi.mocked(decodeEventLog).mockReturnValue({
        eventName: 'Trade',
        args: {
          trader: mockAddress,
          buyYes: true,
          collateralIn: tradeParams.amount,
          tokensOut: 980000000000000000n,
          fee: 20000000000000000n,
          price: 1020408163265306122n,
        },
      } as any);

      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([mockEvent])
      );

      await marketAMM.buyTokens(paramsWithoutMin);

      const buyCall = mockWalletClient.writeContract.mock.calls[1][0];
      expect(buyCall.args).toEqual([tradeParams.amount, 0n]);
    });

    it('should throw error if Trade event not found', async () => {
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);
      mockPublicClient.readContract
        .mockResolvedValueOnce('0xCollateralToken' as any)
        .mockResolvedValueOnce(0n as any);
      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([])
      );

      vi.mocked(decodeEventLog).mockImplementation(() => {
        throw new Error('Not the event');
      });

      await expect(
        marketAMM.buyTokens(tradeParams)
      ).rejects.toThrow(ContractError);
    });
  });

  describe('sellTokens', () => {
    const tradeParams: TradeParams = {
      marketId: mockMarketId,
      outcome: 'NO',
      amount: 1000000000000000000n,
      minAmountOut: 950000000000000000n,
    };

    it('should sell tokens successfully', async () => {
      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);

      const mockEvent = createMockEventLog('Trade', {
        trader: mockAddress,
        buyYes: false,
        collateralIn: 0n, // For sell, collateralIn is 0
        tokensOut: 970000000000000000n, // This is collateralOut for sell
        fee: 30000000000000000n,
        price: 9793814432989690721n,
      });

      vi.mocked(decodeEventLog).mockReturnValue({
        eventName: 'Trade',
        args: {
          trader: mockAddress,
          buyYes: false,
          collateralIn: 0n,
          tokensOut: 970000000000000000n,
          fee: 30000000000000000n,
          price: 9793814432989690721n,
        },
      } as any);

      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([mockEvent])
      );

      const result = await marketAMM.sellTokens(tradeParams);

      expect(result.success).toBe(true);
      expect(result.amountOut).toBe(970000000000000000n);
      expect(result.outcome).toBe('NO');

      expect(mockWalletClient.writeContract).toHaveBeenCalledWith(
        expect.objectContaining({
          address: ammAddress,
          abi: expect.any(Array),
          functionName: 'sellNo',
          args: [tradeParams.amount, tradeParams.minAmountOut],
          account: mockAddress,
        })
      );
    });
  });

  describe('getPrice', () => {
    it('should fetch price correctly', async () => {
      mockPublicClient.readContract.mockResolvedValue(500000000000000000n); // 0.5

      const price = await marketAMM.getPrice(mockMarketId, 'YES');

      expect(price).toBe(500000000000000000n);
      expect(mockPublicClient.readContract).toHaveBeenCalled();
      const callArgs = mockPublicClient.readContract.mock.calls[0][0];
      expect(callArgs.address).toBe(ammAddress);
      expect(callArgs.functionName).toBe('getYesPrice');
      // args is optional and may be undefined when no arguments
      if (callArgs.args !== undefined) {
        expect(callArgs.args).toEqual([]);
      }
      expect(Array.isArray(callArgs.abi)).toBe(true);
    });

    it('should handle errors', async () => {
      mockPublicClient.readContract.mockRejectedValue(
        new Error('Contract read failed')
      );

      await expect(
        marketAMM.getPrice(mockMarketId, 'YES')
      ).rejects.toThrow(ContractError);
    });
  });

  describe('getMarketPrices', () => {
    it('should fetch both prices', async () => {
      mockPublicClient.readContract
        .mockResolvedValueOnce(500000000000000000n) // YES price (getYesPrice)
        .mockResolvedValueOnce(500000000000000000n); // NO price (getNoPrice)

      const prices = await marketAMM.getMarketPrices(mockMarketId);

      expect(prices.yesPrice).toBe(500000000000000000n);
      expect(prices.noPrice).toBe(500000000000000000n);
      
      expect(mockPublicClient.readContract).toHaveBeenCalledTimes(2);
      const yesCall = mockPublicClient.readContract.mock.calls[0][0];
      const noCall = mockPublicClient.readContract.mock.calls[1][0];
      
      expect(yesCall.address).toBe(ammAddress);
      expect(yesCall.functionName).toBe('getYesPrice');
      expect(Array.isArray(yesCall.abi)).toBe(true);
      
      expect(noCall.address).toBe(ammAddress);
      expect(noCall.functionName).toBe('getNoPrice');
      expect(Array.isArray(noCall.abi)).toBe(true);
    });
  });

  describe('getLiquidity', () => {
    it('should fetch liquidity correctly', async () => {
      // Mock reserveYes and reserveNo calls separately
      mockPublicClient.readContract
        .mockResolvedValueOnce(2000000000000000000n) // YES liquidity
        .mockResolvedValueOnce(1000000000000000000n); // NO liquidity

      const liquidity = await marketAMM.getLiquidity(mockMarketId);

      expect(liquidity.yes).toBe(2000000000000000000n);
      expect(liquidity.no).toBe(1000000000000000000n);
    });
  });

  describe('getUserBalance', () => {
    it('should fetch user balance correctly', async () => {
      mockPublicClient.readContract.mockResolvedValue(500000000000000000n);

      const balance = await marketAMM.getUserBalance(
        mockAddress,
        mockMarketId,
        'YES',
        outcomeTokenAddress
      );

      expect(balance).toBe(500000000000000000n);
      // Token ID encoding: (marketId << 8) | outcomeId
      // For marketId = 1, YES (outcomeId = 0): (1 << 8) | 0 = 256
      const expectedTokenId = (mockMarketId << 8n) | 0n;
      expect(mockPublicClient.readContract).toHaveBeenCalledWith({
        address: outcomeTokenAddress,
        abi: expect.any(Array),
        functionName: 'balanceOf',
        args: [mockAddress, expectedTokenId],
      });
    });
  });

  describe('addLiquidity', () => {
    it('should add liquidity successfully', async () => {
      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);
      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([])
      );
      // Mock collateral token address query
      mockPublicClient.readContract.mockResolvedValueOnce('0xCollateralToken' as any);
      // Mock allowance check
      mockPublicClient.readContract.mockResolvedValueOnce(0n as any);

      const hash = await marketAMM.addLiquidity(1000000000000000000n);

      expect(hash).toBe(mockTransactionHash);
      expect(mockWalletClient.writeContract).toHaveBeenCalled();
      const callArgs = mockWalletClient.writeContract.mock.calls[1][0]; // Second call (after approve)
      expect(callArgs.address).toBe(ammAddress);
      expect(callArgs.functionName).toBe('addLiquidity');
      expect(callArgs.args).toEqual([1000000000000000000n]);
      expect(callArgs.account).toBe(mockAddress);
    });

    it('should throw error if wallet client not provided', async () => {
      const ammWithoutWallet = new MarketAMM(
        mockPublicClient as any,
        ammAddress
      );

      await expect(
        ammWithoutWallet.addLiquidity(1n)
      ).rejects.toThrow(ContractError);
    });
  });

  describe('removeLiquidity', () => {
    it('should remove liquidity successfully', async () => {
      mockWalletClient.getAddresses.mockResolvedValue([mockAddress]);
      mockWalletClient.writeContract.mockResolvedValue(mockTransactionHash);
      mockPublicClient.waitForTransactionReceipt.mockResolvedValue(
        createMockReceipt([])
      );

      const lpTokens = 500000000000000000n;
      const hash = await marketAMM.removeLiquidity(lpTokens);

      expect(hash).toBe(mockTransactionHash);
      expect(mockWalletClient.writeContract).toHaveBeenCalled();
      const callArgs = mockWalletClient.writeContract.mock.calls[0][0];
      expect(callArgs.address).toBe(ammAddress);
      expect(callArgs.functionName).toBe('removeLiquidity');
      expect(callArgs.args).toEqual([lpTokens]);
      expect(callArgs.account).toBe(mockAddress);
      expect(Array.isArray(callArgs.abi)).toBe(true);
    });

    it('should throw error if wallet client not provided', async () => {
      const ammWithoutWallet = new MarketAMM(
        mockPublicClient as any,
        ammAddress
      );

      await expect(
        ammWithoutWallet.removeLiquidity(1n)
      ).rejects.toThrow(ContractError);
    });
  });

  describe('getLPBalance', () => {
    it('should fetch LP balance correctly', async () => {
      mockPublicClient.readContract.mockResolvedValue(1000000000000000000n);

      const balance = await marketAMM.getLPBalance(mockAddress);

      expect(balance).toBe(1000000000000000000n);
      expect(mockPublicClient.readContract).toHaveBeenCalledWith({
        address: ammAddress,
        abi: expect.any(Array),
        functionName: 'balanceOf',
        args: [mockAddress],
      });
    });
  });
});

