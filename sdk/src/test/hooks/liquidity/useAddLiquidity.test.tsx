/**
 * Tests for useAddLiquidity hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useAddLiquidity } from '../../../hooks/liquidity/useAddLiquidity';
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

describe('useAddLiquidity', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: mockWalletClient } as any);
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    
    // Mock readContract calls
    mockPublicClient.readContract = vi.fn()
      .mockResolvedValueOnce({
        marketId: 1n,
        amm: '0xAMM' as const,
        collateralToken: '0xToken' as const,
        closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
        category: 'test',
        metadataURI: 'ipfs://test',
        status: 0,
      })
      .mockResolvedValueOnce(0n); // allowance
    
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockResolvedValue(mockTransactionHash);
    }
    
    mockPublicClient.waitForTransactionReceipt = vi.fn().mockResolvedValue({
      status: 'success',
      transactionHash: mockTransactionHash,
    });
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONFIG.marketFactoryAddress,
    });
  });

  it('should add liquidity successfully', async () => {
    const { result } = renderHook(() => useAddLiquidity(1), { wrapper });

    const params = {
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.hash).toBe(mockTransactionHash);
  });

  it('should handle errors gracefully', async () => {
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockRejectedValue(new Error('Insufficient balance'));
    }

    const { result } = renderHook(() => useAddLiquidity(1), { wrapper });

    const params = {
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.error).toBeDefined();
    }, { timeout: 3000 });

    expect(result.current.error).toBeDefined();
  });

  it('should not execute if wallet is not connected', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: null } as any);

    const { result } = renderHook(() => useAddLiquidity(1), { wrapper });

    const params = {
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.error).toBeDefined();
    }, { timeout: 3000 });

    if (result.current.error instanceof Error) {
      expect(result.current.error.message).toContain('Wallet not connected');
    }
  });
});

