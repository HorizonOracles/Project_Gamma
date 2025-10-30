/**
 * Tests for useOracleStatus hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useOracleStatus } from '../../../hooks/oracle/useOracleStatus';
import { createOracleApiClient } from '../../../utils/api';

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
}));

vi.mock('../../../utils/api', () => ({
  createOracleApiClient: vi.fn(),
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

describe('useOracleStatus', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: 'https://api.example.com',
    });
  });

  it('should fetch oracle status successfully', async () => {
    const mockStatus = {
      requestId: 'req-123',
      marketId: 1,
      status: 'processing' as const,
      progress: 50,
    };

    const mockApiClient = {
      getRequestStatus: vi.fn().mockResolvedValue(mockStatus),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useOracleStatus('req-123'), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toEqual(mockStatus);
    expect(mockApiClient.getRequestStatus).toHaveBeenCalledWith('req-123');
  });

  it('should not fetch if requestId is undefined', () => {
    const { result } = renderHook(() => useOracleStatus(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if oracle API URL is not configured', async () => {
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: undefined,
    });

    const { result } = renderHook(() => useOracleStatus('req-123'), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should poll status at intervals when enabled', async () => {
    const mockStatus = {
      requestId: 'req-123',
      marketId: 1,
      status: 'processing' as const,
    };

    const mockApiClient = {
      getRequestStatus: vi.fn().mockResolvedValue(mockStatus),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useOracleStatus('req-123', {
      refetchInterval: 1000,
    }), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBeDefined();
  });

  it('should handle errors gracefully', async () => {
    const mockApiClient = {
      getRequestStatus: vi.fn().mockRejectedValue(new Error('Request not found')),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useOracleStatus('invalid-id'), { wrapper });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });
});

