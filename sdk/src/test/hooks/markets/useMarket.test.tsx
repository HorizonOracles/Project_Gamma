/**
 * Tests for useMarket hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useMarket } from '../../../hooks/markets/useMarket';
import { MarketStatus } from '../../../types';
import { MarketFactory } from '../../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS, BNB_CHAIN } from '../../../constants';
import {
  mockPublicClient,
  mockAddress,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useChainId: vi.fn(),
}));

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
}));

vi.mock('../../../contracts/MarketFactory');

const wrapper = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};

describe('useMarket', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory,
    });
  });

  it('should fetch market data successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test question?',
      description: 'Test description',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: MarketStatus.Active,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 1000000n,
      totalLiquidity: { yes: 500000n, no: 500000n },
      createdAt: BigInt(Math.floor(Date.now() / 1000)),
      collateralToken: '0xToken1' as const,
      category: 'sports',
      metadataURI: 'ipfs://1',
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarket(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.id).toBe(1);
    expect(result.current.data?.question).toBe('Test question?');
    expect(result.current.data?.status).toBe(MarketStatus.Active);
    expect(mockMarketFactory.getMarket).toHaveBeenCalledWith(1n);
  });

  it('should not fetch if marketId is undefined', () => {
    const { result } = renderHook(() => useMarket(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if public client is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(null as any);

    const { result } = renderHook(() => useMarket(1), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle errors gracefully', async () => {
    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockRejectedValue(new Error('Market not found')),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarket(999), { wrapper });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });

  it('should handle missing optional fields', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test question?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: MarketStatus.Active,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarket(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.category).toBe('');
    expect(result.current.data?.metadataURI).toBe('');
  });
});

