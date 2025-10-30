/**
 * Tests for useMarkets hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useMarkets } from '../../../hooks/markets/useMarkets';
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

describe('useMarkets', () => {
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

  it('should fetch all markets without filters', async () => {
    const mockMarkets = [
      {
        id: 1n,
        creator: mockAddress,
        amm: '0xMarket1' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'sports',
        metadataURI: 'ipfs://1',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
      {
        id: 2n,
        creator: mockAddress,
        amm: '0xMarket2' as const,
        collateralToken: '0xToken2' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'politics',
        metadataURI: 'ipfs://2',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
    ];

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockResolvedValue(2n),
      getMarkets: vi.fn().mockResolvedValue(mockMarkets),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets(), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toHaveLength(2);
    expect(result.current.data?.[0].id).toBe(1);
    expect(result.current.data?.[1].id).toBe(2);
  });

  it('should filter markets by category', async () => {
    const mockMarkets = [
      {
        id: 1n,
        creator: mockAddress,
        amm: '0xMarket1' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'sports',
        metadataURI: 'ipfs://1',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
      {
        id: 2n,
        creator: mockAddress,
        amm: '0xMarket2' as const,
        collateralToken: '0xToken2' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'politics',
        metadataURI: 'ipfs://2',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
    ];

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockResolvedValue(2n),
      getMarkets: vi.fn().mockResolvedValue(mockMarkets),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets({ category: 'sports' }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toHaveLength(1);
    expect(result.current.data?.[0].category).toBe('sports');
  });

  it('should filter markets by status', async () => {
    const mockMarkets = [
      {
        id: 1n,
        creator: mockAddress,
        amm: '0xMarket1' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'sports',
        metadataURI: 'ipfs://1',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
      {
        id: 2n,
        creator: mockAddress,
        amm: '0xMarket2' as const,
        collateralToken: '0xToken2' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) - 86400),
        category: 'politics',
        metadataURI: 'ipfs://2',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Closed,
      },
    ];

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockResolvedValue(2n),
      getMarkets: vi.fn().mockResolvedValue(mockMarkets),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets({ status: MarketStatus.Active }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toHaveLength(1);
    expect(result.current.data?.[0].status).toBe(MarketStatus.Active);
  });

  it('should filter markets by creator', async () => {
    const otherAddress = '0x9999999999999999999999999999999999999999' as const;
    const mockMarkets = [
      {
        id: 1n,
        creator: mockAddress,
        amm: '0xMarket1' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'sports',
        metadataURI: 'ipfs://1',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
      {
        id: 2n,
        creator: otherAddress,
        amm: '0xMarket2' as const,
        collateralToken: '0xToken2' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'politics',
        metadataURI: 'ipfs://2',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
    ];

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockResolvedValue(2n),
      getMarkets: vi.fn().mockResolvedValue(mockMarkets),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets({ creator: mockAddress }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toHaveLength(1);
    expect(result.current.data?.[0].creator.toLowerCase()).toBe(mockAddress.toLowerCase());
  });

  it('should handle pagination with limit and offset', async () => {
    const mockMarkets = [
      {
        id: 1n,
        creator: mockAddress,
        amm: '0xMarket1' as const,
        collateralToken: '0xToken1' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'sports',
        metadataURI: 'ipfs://1',
        creatorStake: 1000n,
        stakeRefunded: false,
        status: MarketStatus.Active,
      },
    ];

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockResolvedValue(10n),
      getMarkets: vi.fn().mockResolvedValue(mockMarkets),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets({ limit: 5, offset: 0 }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(mockMarketFactory.getMarkets).toHaveBeenCalledWith(0n, 5n);
  });

  it('should not fetch if public client is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(null as any);

    const { result } = renderHook(() => useMarkets(), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle errors gracefully', async () => {
    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarketCount: vi.fn().mockRejectedValue(new Error('Network error')),
      getMarkets: vi.fn(),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useMarkets(), { wrapper });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });
});

