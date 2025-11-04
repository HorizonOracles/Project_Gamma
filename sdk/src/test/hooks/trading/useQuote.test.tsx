/**
 * Tests for useQuote hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useQuote } from '../../../hooks/trading/useQuote';
import { MarketFactory } from '../../../contracts/MarketFactory';
import { BinaryMarket } from '../../../contracts/BinaryMarket';
import { DEFAULT_CONTRACTS, BNB_CHAIN } from '../../../constants';
import { mockPublicClient, mockAddress } from '../../mocks/viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useChainId: vi.fn(),
  useAccount: vi.fn(),
}));

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
}));

vi.mock('../../../contracts/MarketFactory');
vi.mock('../../../contracts/BinaryMarket');
vi.mock('../../../utils', () => ({
  getMarketContract: vi.fn(),
}));

const wrapper = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};

describe('useQuote', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory,
    });
  });

  it('should get buy quote successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: 0,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const mockBuyQuote = {
      tokensOut: 950000000000000000n,
      fee: 50000000000000000n,
    };

    const mockPrices = {
      yesPrice: 500000000000000000n,
      noPrice: 500000000000000000n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const BinaryMarketModule = await import('../../../contracts/BinaryMarket');
    const mockBinaryMarket = {
      getBuyQuote: vi.fn().mockResolvedValue(mockBuyQuote),
      getMarketPrices: vi.fn().mockResolvedValue(mockPrices),
      getPrice: vi.fn().mockResolvedValue(mockPrices.yesPrice),
    };
    vi.mocked(BinaryMarketModule.BinaryMarket).mockImplementation(() => mockBinaryMarket as any);

    const utilsModule = await import('../../../utils');
    vi.mocked(utilsModule.getMarketContract).mockResolvedValue(mockBinaryMarket as any);

    const { result } = renderHook(() => useQuote({
      marketId: 1,
      outcomeId: 0,
      amount: 1000000000000000000n,
      isBuy: true,
    }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.tokensOut).toBe(mockBuyQuote.tokensOut);
    expect(result.current.data?.fee).toBe(mockBuyQuote.fee);
  });

  it('should get sell quote successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: 0,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const mockSellQuote = {
      collateralOut: 950000000000000000n,
      fee: 50000000000000000n,
    };

    const mockPrices = {
      yesPrice: 500000000000000000n,
      noPrice: 500000000000000000n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const BinaryMarketModule = await import('../../../contracts/BinaryMarket');
    const mockBinaryMarket = {
      getSellQuote: vi.fn().mockResolvedValue(mockSellQuote),
      getMarketPrices: vi.fn().mockResolvedValue(mockPrices),
      getPrice: vi.fn().mockResolvedValue(mockPrices.yesPrice),
    };
    vi.mocked(BinaryMarketModule.BinaryMarket).mockImplementation(() => mockBinaryMarket as any);

    const utilsModule = await import('../../../utils');
    vi.mocked(utilsModule.getMarketContract).mockResolvedValue(mockBinaryMarket as any);

    const { result } = renderHook(() => useQuote({
      marketId: 1,
      outcomeId: 0,
      amount: 1000000000000000000n,
      isBuy: false,
    }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.tokensOut).toBe(mockSellQuote.collateralOut);
    expect(result.current.data?.fee).toBe(mockSellQuote.fee);
  });

  it('should not fetch if params are undefined', () => {
    const { result } = renderHook(() => useQuote(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if address is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: undefined } as any);

    const { result } = renderHook(() => useQuote({
      marketId: 1,
      outcomeId: 0,
      amount: 1000000000000000000n,
      isBuy: true,
    }), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should calculate price impact for buy quotes', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: 0,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const mockBuyQuote = {
      tokensOut: 900000000000000000n,
      fee: 100000000000000000n,
    };

    const mockPrices = {
      yesPrice: 1000000000000000000n,
      noPrice: 0n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const BinaryMarketModule = await import('../../../contracts/BinaryMarket');
    const mockBinaryMarket = {
      getBuyQuote: vi.fn().mockResolvedValue(mockBuyQuote),
      getMarketPrices: vi.fn().mockResolvedValue(mockPrices),
      getPrice: vi.fn().mockResolvedValue(mockPrices.yesPrice),
    };
    vi.mocked(BinaryMarketModule.BinaryMarket).mockImplementation(() => mockBinaryMarket as any);

    const utilsModule = await import('../../../utils');
    vi.mocked(utilsModule.getMarketContract).mockResolvedValue(mockBinaryMarket as any);

    const { result } = renderHook(() => useQuote({
      marketId: 1,
      outcomeId: 0,
      amount: 1000000000000000000n,
      isBuy: true,
    }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data?.priceImpact).toBeGreaterThan(0);
  });
});

