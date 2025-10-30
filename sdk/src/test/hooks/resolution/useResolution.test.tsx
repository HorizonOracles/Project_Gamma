/**
 * Tests for useResolution hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useResolution } from '../../../hooks/resolution/useResolution';
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

describe('useResolution', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      resolutionModuleAddress: DEFAULT_CONFIG.resolutionModuleAddress,
    });
  });

  it('should fetch resolution successfully', async () => {
    const mockResolution = [
      1, // state
      0n, // proposedOutcome
      BigInt(Math.floor(Date.now() / 1000)), // proposalTime
      '0xProposer' as const, // proposer
      1000000000000000000n, // proposerBond
      '0xDisputer' as const, // disputer
      0n, // disputerBond
      'ipfs://evidence', // evidenceURI
    ];

    mockPublicClient.readContract = vi.fn().mockResolvedValue(mockResolution);

    const { result } = renderHook(() => useResolution(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
    expect(result.current.data?.state).toBe(1);
    expect(result.current.data?.proposedOutcome).toBe(0n);
  });

  it('should return null if resolution module is not configured', async () => {
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      resolutionModuleAddress: '0x0000000000000000000000000000000000000000' as const,
    });

    const { result } = renderHook(() => useResolution(1), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeNull();
  });

  it('should not fetch if marketId is undefined', () => {
    const { result } = renderHook(() => useResolution(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle errors gracefully and return null', async () => {
    mockPublicClient.readContract = vi.fn().mockRejectedValue(new Error('Resolution not found'));

    const { result } = renderHook(() => useResolution(999), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    // Hook returns null on error, not an error state
    expect(result.current.data).toBeNull();
  });
});
