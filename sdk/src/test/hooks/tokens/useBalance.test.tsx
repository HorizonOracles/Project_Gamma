/**
 * Tests for useBalance and useOutcomeBalance hooks
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useBalance, useOutcomeBalance } from '../../../hooks/tokens/useBalance';
import { DEFAULT_CONTRACTS, BNB_CHAIN } from '../../../constants';
import {
  mockAddress,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  useAccount: vi.fn(),
  useChainId: vi.fn(),
  useReadContract: vi.fn(),
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

describe('useBalance', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    vi.mocked(wagmiModule.useReadContract).mockReturnValue({
      data: 1000000000000000000n,
      isLoading: false,
      isSuccess: true,
    } as any);
  });

  it('should fetch ERC20 balance successfully', async () => {
    const tokenAddress = '0xToken1' as const;

    const { result } = renderHook(() => useBalance(tokenAddress), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBe(1000000000000000000n);
  });

  it('should not fetch if tokenAddress is undefined', () => {
    const { result } = renderHook(() => useBalance(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
  });

  it('should not fetch if address is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: undefined } as any);

    const tokenAddress = '0xToken1' as const;
    const { result } = renderHook(() => useBalance(tokenAddress), { wrapper });

    expect(result.current.isLoading).toBe(false);
  });
});

describe('useOutcomeBalance', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    vi.mocked(wagmiModule.useReadContract).mockReturnValue({
      data: 500000000000000000n,
      isLoading: false,
      isSuccess: true,
    } as any);
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      outcomeTokenAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].outcomeToken,
    });
  });

  it('should fetch outcome token balance successfully', async () => {
    const { result } = renderHook(() => useOutcomeBalance(1, 0), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBe(500000000000000000n);
  });

  it('should encode token ID correctly', async () => {
    const wagmiModule = await import('wagmi');
    const readContractMock = vi.fn().mockReturnValue({
      data: 500000000000000000n,
      isLoading: false,
      isSuccess: true,
    });
    vi.mocked(wagmiModule.useReadContract).mockReturnValue(readContractMock() as any);

    const { result } = renderHook(() => useOutcomeBalance(1, 0), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    const expectedTokenId = (1n << 8n) | 0n;
    expect(readContractMock).toHaveBeenCalled();
  });

  it('should not fetch if marketId is undefined', () => {
    const { result } = renderHook(() => useOutcomeBalance(undefined, 0), { wrapper });

    expect(result.current.isLoading).toBe(false);
  });

  it('should not fetch if address is not available', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: undefined } as any);

    const { result } = renderHook(() => useOutcomeBalance(1, 0), { wrapper });

    expect(result.current.isLoading).toBe(false);
  });
});

