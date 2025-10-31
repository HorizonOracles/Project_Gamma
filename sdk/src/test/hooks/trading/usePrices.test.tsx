/**
 * Tests for usePrices hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { usePrices } from '../../../hooks/trading/usePrices';
import { DEFAULT_CONFIG, BNB_CHAIN } from '../../../constants';
import {
  mockPublicClient,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useChainId: vi.fn(),
}));

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
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

describe('usePrices', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONFIG.marketFactoryAddress,
    });
  });

  it('should fetch market prices successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      amm: '0xAMM' as const,
      collateralToken: '0xToken' as const,
      closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      category: 'test',
      metadataURI: 'ipfs://test',
      status: 0,
    };

    mockPublicClient.readContract = vi.fn()
      .mockResolvedValueOnce(mockMarketInfo) // getMarket
      .mockResolvedValueOnce(600000000000000000n) // getYesPrice
      .mockResolvedValueOnce(400000000000000000n); // getNoPrice

    const { result } = renderHook(() => usePrices(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.yesPrice).toBe(600000000000000000n);
    expect(result.current.data?.noPrice).toBe(400000000000000000n);
    expect(result.current.data?.yes).toBeCloseTo(0.6, 2);
    expect(result.current.data?.no).toBeCloseTo(0.4, 2);
  });

  it('should not fetch if marketId is undefined', () => {
    const { result } = renderHook(() => usePrices(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if public client is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(null as any);

    const { result } = renderHook(() => usePrices(1), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle errors gracefully', async () => {
    mockPublicClient.readContract = vi.fn().mockRejectedValue(new Error('Market not found'));

    const { result } = renderHook(() => usePrices(999), { wrapper });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });
});
