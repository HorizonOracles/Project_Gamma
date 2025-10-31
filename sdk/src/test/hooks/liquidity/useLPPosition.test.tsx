/**
 * Tests for useLPPosition hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useLPPosition } from '../../../hooks/liquidity/useLPPosition';
import { DEFAULT_CONFIG, BNB_CHAIN } from '../../../constants';
import {
  mockPublicClient,
  mockAddress,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  useAccount: vi.fn(),
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

describe('useLPPosition', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONFIG.marketFactoryAddress,
    });
  });

  it('should fetch LP position successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      amm: '0xAMM' as const,
      collateralToken: '0xToken' as const,
      closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      category: 'test',
      metadataURI: 'ipfs://test',
      status: 0,
    };

    // Mock readContract calls: getMarket, balanceOf, reserveYes, reserveNo, totalCollateral
    mockPublicClient.readContract = vi.fn()
      .mockResolvedValueOnce(mockMarketInfo) // getMarket
      .mockResolvedValueOnce(5000000000000000000n) // balanceOf (LP tokens)
      .mockResolvedValueOnce(10000000000000000000n) // reserveYes
      .mockResolvedValueOnce(10000000000000000000n) // reserveNo
      .mockResolvedValueOnce(100000000000000000000n); // totalCollateral

    const { result } = renderHook(() => useLPPosition(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.lpTokens).toBe(5000000000000000000n);
    // Share calculation: (lpBalance * 10000) / totalCollateral / 100
    // (5000000000000000000n * 10000n) / 100000000000000000000n = 500n
    // Number(500n) / 100 = 5 (which represents 5%)
    expect(result.current.data?.share).toBe(5);
  });

  it('should not fetch if marketId is undefined', () => {
    const { result } = renderHook(() => useLPPosition(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if address is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: undefined } as any);

    const { result } = renderHook(() => useLPPosition(1), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle zero LP supply correctly', async () => {
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
      .mockResolvedValueOnce(0n) // balanceOf (LP tokens)
      .mockResolvedValueOnce(0n) // reserveYes
      .mockResolvedValueOnce(0n) // reserveNo
      .mockResolvedValueOnce(0n); // totalCollateral

    const { result } = renderHook(() => useLPPosition(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data?.share).toBe(0);
    expect(result.current.data?.value).toBe(0n);
  });
});
