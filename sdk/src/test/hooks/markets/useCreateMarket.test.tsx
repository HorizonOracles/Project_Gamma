/**
 * Tests for useCreateMarket hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useCreateMarket } from '../../../hooks/markets/useCreateMarket';
import { DEFAULT_CONTRACTS, BNB_CHAIN, MARKET_FACTORY_ABI } from '../../../constants';
import {
  mockPublicClient,
  mockWalletClient,
  mockAddress,
  mockTransactionHash,
  createMockReceipt,
  createMockEventLog,
} from '../../mocks/viem';
import { decodeEventLog } from 'viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useWalletClient: vi.fn(),
  useChainId: vi.fn(),
  useAccount: vi.fn(),
}));

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
}));

vi.mock('viem', async () => {
  const actual = await vi.importActual('viem');
  return {
    ...actual,
    decodeEventLog: vi.fn(),
  };
});

const wrapper = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};

describe('useCreateMarket', () => {
  beforeEach(async () => {
    vi.clearAllMocks();
    
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.usePublicClient).mockReturnValue(mockPublicClient as any);
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: mockWalletClient } as any);
    vi.mocked(wagmiModule.useChainId).mockReturnValue(BNB_CHAIN.MAINNET);
    vi.mocked(wagmiModule.useAccount).mockReturnValue({ address: mockAddress } as any);
    
    // Mock walletClient.writeContract
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockResolvedValue(mockTransactionHash);
    }
    
    // Mock publicClient.readContract (for allowance check)
    mockPublicClient.readContract = vi.fn().mockResolvedValue(0n);
    
    // Mock publicClient.waitForTransactionReceipt
    const mockEvent = createMockEventLog('MarketCreated', {
      marketId: 1n,
      creator: mockAddress,
    });
    
    vi.mocked(decodeEventLog).mockReturnValue({
      eventName: 'MarketCreated',
      args: {
        marketId: 1n,
        creator: mockAddress,
      },
    } as any);
    
    mockPublicClient.waitForTransactionReceipt = vi.fn().mockResolvedValue(
      createMockReceipt([mockEvent])
    );
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory,
    });
  });

  it('should create market successfully', async () => {
    const { result } = renderHook(() => useCreateMarket(), { wrapper });

    const params = {
      question: 'New market?',
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      collateralToken: '0xToken1' as const,
      category: 'sports',
      metadataURI: 'ipfs://test',
      creatorStake: 0n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data).toBe(1n);
    if (mockWalletClient.writeContract) {
      expect(mockWalletClient.writeContract).toHaveBeenCalled();
    }
  });

  it('should handle creation errors', async () => {
    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockRejectedValue(new Error('Creation failed'));
    }

    const { result } = renderHook(() => useCreateMarket(), { wrapper });

    const params = {
      question: 'Fail market?',
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      collateralToken: '0xToken1' as const,
      category: 'sports',
      metadataURI: 'ipfs://test',
      creatorStake: 0n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
  });

  it('should not execute if wallet is not connected', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: null } as any);

    const { result } = renderHook(() => useCreateMarket(), { wrapper });

    const params = {
      question: 'Test market?',
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      collateralToken: '0xToken1' as const,
      category: 'sports',
      metadataURI: 'ipfs://test',
      creatorStake: 0n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });

    expect(result.current.error).toBeDefined();
    if (result.current.error instanceof Error) {
      expect(result.current.error.message).toContain('Wallet client not available');
    }
  });
});
