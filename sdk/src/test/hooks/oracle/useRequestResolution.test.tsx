/**
 * Tests for useRequestResolution hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useRequestResolution } from '../../../hooks/oracle/useRequestResolution';
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

describe('useRequestResolution', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: 'https://api.example.com',
    });
  });

  it('should request resolution successfully', async () => {
    const mockResponse = {
      requestId: 'req-123',
      marketId: 1,
      status: 'pending' as const,
    };

    const mockApiClient = {
      requestResolution: vi.fn().mockResolvedValue(mockResponse),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useRequestResolution(), { wrapper });

    const params = {
      marketId: 1,
      metadata: {
        question: 'Test question?',
        description: 'Test description',
      },
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toEqual(mockResponse);
    expect(mockApiClient.requestResolution).toHaveBeenCalledWith(params);
  });

  it('should handle errors gracefully', async () => {
    const mockApiClient = {
      requestResolution: vi.fn().mockRejectedValue(new Error('API error')),
    };

    vi.mocked(createOracleApiClient).mockReturnValue(mockApiClient as any);

    const { result } = renderHook(() => useRequestResolution(), { wrapper });

    const params = {
      marketId: 1,
      metadata: {
        question: 'Test question?',
      },
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });

  it('should throw error if oracle API URL is not configured', async () => {
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: 56,
      oracleApiUrl: undefined,
    });

    const { result } = renderHook(() => useRequestResolution(), { wrapper });

    const params = {
      marketId: 1,
      metadata: {
        question: 'Test question?',
      },
    };

    result.current.mutate(params);

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    if (result.current.error instanceof Error) {
      expect(result.current.error.message).toContain('Oracle API URL not configured');
    }
  });
});

