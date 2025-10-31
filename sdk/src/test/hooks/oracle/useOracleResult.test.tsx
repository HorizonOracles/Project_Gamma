/**
 * Tests for useOracleResult hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useOracleResult } from '../../../hooks/oracle/useOracleResult';
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

describe('useOracleResult', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: 'https://api.example.com',
    });
  });

  it('should fetch oracle result successfully', async () => {
    const mockResult = {
      requestId: 'req-123',
      marketId: 1,
      outcomeId: 0,
      confidence: 95,
      reasoning: 'Strong evidence supports YES outcome',
      sources: ['source1', 'source2'],
      evidenceUrl: 'https://evidence.example.com',
      timestamp: Date.now(),
    };

    const mockApiClient = {
      getRequestResult: vi.fn().mockResolvedValue(mockResult),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useOracleResult('req-123'), { wrapper });

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toEqual(mockResult);
    expect(mockApiClient.getRequestResult).toHaveBeenCalledWith('req-123');
  });

  it('should not fetch if requestId is undefined', () => {
    const { result } = renderHook(() => useOracleResult(undefined), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should not fetch if oracle API URL is not configured', async () => {
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: undefined,
    });

    const { result } = renderHook(() => useOracleResult('req-123'), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should respect enabled option', () => {
    const { result } = renderHook(() => useOracleResult('req-123', {
      enabled: false,
    }), { wrapper });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isFetching).toBe(false);
  });

  it('should handle errors gracefully', async () => {
    const mockApiClient = {
      getRequestResult: vi.fn().mockRejectedValue(new Error('Result not found')),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useOracleResult('invalid-id'), { wrapper });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });
});

