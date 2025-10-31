/**
 * Tests for useApprove hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useApprove } from '../../../hooks/tokens/useApprove';
import {
  mockTransactionHash,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  useWriteContract: vi.fn(),
  useWaitForTransactionReceipt: vi.fn(),
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

describe('useApprove', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useWriteContract).mockReturnValue({
      writeContract: vi.fn().mockResolvedValue(mockTransactionHash),
      data: mockTransactionHash,
    } as any);
    vi.mocked(wagmiModule.useWaitForTransactionReceipt).mockReturnValue({
      isLoading: false,
      isSuccess: true,
    } as any);
  });

  it('should approve token spending successfully', async () => {
    const wagmiModule = await import('wagmi');
    const writeContractMock = vi.fn().mockResolvedValue(mockTransactionHash);
    vi.mocked(wagmiModule.useWriteContract).mockReturnValue({
      writeContract: writeContractMock,
      data: mockTransactionHash,
    } as any);

    const { result } = renderHook(() => useApprove(), { wrapper });

    const params = {
      tokenAddress: '0xToken1' as const,
      spender: '0xSpender1' as const,
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(writeContractMock).toHaveBeenCalled();
    });

    expect(writeContractMock).toHaveBeenCalledWith({
      address: params.tokenAddress,
      abi: expect.any(Array),
      functionName: 'approve',
      args: [params.spender, params.amount],
    });
  });

  it('should track transaction status', async () => {
    const { result } = renderHook(() => useApprove(), { wrapper });

    const params = {
      tokenAddress: '0xToken1' as const,
      spender: '0xSpender1' as const,
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.hash).toBe(mockTransactionHash);
  });
});

