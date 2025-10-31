/**
 * Test helpers for React hooks testing
 */

import React from 'react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { vi } from 'vitest';
import { Address } from 'viem';
import { DEFAULT_CONTRACTS, BNB_CHAIN } from '../../constants';
import {
  mockPublicClient,
  mockWalletClient,
  mockAddress,
  mockTransactionHash,
} from '../mocks/viem';

/**
 * Wrapper component for React Query hooks
 */
export function createWrapper() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });

  return ({ children }: { children: React.ReactNode }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}

/**
 * Setup default mocks for wagmi hooks
 */
export async function setupWagmiMocks(options?: {
  publicClient?: any;
  walletClient?: any;
  chainId?: number;
  address?: string;
}) {
  const wagmiModule = await import('wagmi');
  
  vi.mocked(wagmiModule.usePublicClient).mockReturnValue(
    options?.publicClient || mockPublicClient as any
  );
  
  vi.mocked(wagmiModule.useWalletClient).mockReturnValue({
    data: options?.walletClient || mockWalletClient,
  } as any);
  
  vi.mocked(wagmiModule.useChainId).mockReturnValue(
    options?.chainId || BNB_CHAIN.MAINNET
  );
  
  vi.mocked(wagmiModule.useAccount).mockReturnValue({
    address: options?.address || mockAddress,
  } as any);
  
  vi.mocked(wagmiModule.useWriteContract).mockReturnValue({
    writeContract: vi.fn().mockResolvedValue(mockTransactionHash),
    data: mockTransactionHash,
  } as any);
  
  vi.mocked(wagmiModule.useWaitForTransactionReceipt).mockReturnValue({
    isLoading: false,
    isSuccess: true,
    hash: mockTransactionHash,
  } as any);
  
  vi.mocked(wagmiModule.useReadContract).mockReturnValue({
    data: 0n,
    isLoading: false,
    isSuccess: true,
  } as any);
}

/**
 * Setup default mocks for GammaProvider
 */
export async function setupGammaMocks(options?: {
  chainId?: number;
  marketFactoryAddress?: string;
  oracleApiUrl?: string;
}) {
  const gammaModule = await import('../../components/GammaProvider');
  
  vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
    chainId: options?.chainId || BNB_CHAIN.MAINNET,
    marketFactoryAddress: (options?.marketFactoryAddress || 
      DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory) as Address,
    oracleApiUrl: options?.oracleApiUrl,
  });
}

/**
 * Mock fetch for API tests
 */
export function mockFetch(response: any, ok: boolean = true) {
  global.fetch = vi.fn().mockResolvedValue({
    ok,
    json: async () => response,
    status: ok ? 200 : 400,
    statusText: ok ? 'OK' : 'Bad Request',
  } as Response);
}

