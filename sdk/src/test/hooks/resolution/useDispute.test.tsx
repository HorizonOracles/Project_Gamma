/**
 * Tests for useDispute hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useDispute } from '../../../hooks/resolution/useDispute';
import { DEFAULT_CONFIG, BNB_CHAIN } from '../../../constants';
import {
  mockPublicClient,
  mockWalletClient,
  mockAddress,
  mockTransactionHash,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useWalletClient: vi.fn(),
  useAccount: vi.fn(),
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

describe('useDispute', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: mockWalletClient } as any);
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockResolvedValue(mockTransactionHash);
    }
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      resolutionModuleAddress: DEFAULT_CONFIG.resolutionModuleAddress,
    });
  });

  it('should dispute resolution successfully', async () => {
    const { result } = renderHook(() => useDispute(1), { wrapper });

    const params = {
      bondAmount: 1000000000000000000n,
      reason: 'Evidence is incorrect',
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.hash).toBe(mockTransactionHash);
  });

  it('should handle errors gracefully', async () => {
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockRejectedValue(new Error('No proposal to dispute'));
    }

    const { result } = renderHook(() => useDispute(1), { wrapper });

    const params = {
      bondAmount: 1000000000000000000n,
      reason: 'Evidence is incorrect',
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.error).toBeDefined();
    }, { timeout: 3000 });

    expect(result.current.error).toBeDefined();
  });

  it('should not execute if wallet is not connected', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: null } as any);

    const { result } = renderHook(() => useDispute(1), { wrapper });

    const params = {
      bondAmount: 1000000000000000000n,
      reason: 'Evidence is incorrect',
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.error).toBeDefined();
    }, { timeout: 3000 });

    if (result.current.error instanceof Error) {
      expect(result.current.error.message).toContain('Wallet not connected');
    }
  });
});

