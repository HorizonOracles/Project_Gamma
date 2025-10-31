/**
 * Tests for useBuy hook
 */

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import { useBuy } from '../../../hooks/trading/useBuy';
import { MarketFactory } from '../../../contracts/MarketFactory';
import { MarketAMM } from '../../../contracts/MarketAMM';
import { DEFAULT_CONTRACTS, BNB_CHAIN } from '../../../constants';
import {
  mockPublicClient,
  mockWalletClient,
  mockAddress,
  mockTransactionHash,
} from '../../mocks/viem';

vi.mock('wagmi', () => ({
  usePublicClient: vi.fn(),
  useWalletClient: vi.fn(),
  useChainId: vi.fn(),
  useAccount: vi.fn(),
  useWriteContract: vi.fn(),
  useWaitForTransactionReceipt: vi.fn(),
}));

vi.mock('../../../components/GammaProvider', () => ({
  useGammaConfig: vi.fn(),
}));

vi.mock('../../../contracts/MarketFactory');
vi.mock('../../../contracts/MarketAMM');

const wrapper = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};

describe('useBuy', () => {
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
    
    const gammaModule = await import('../../../components/GammaProvider');
    vi.mocked(gammaModule.useGammaConfig).mockReturnValue({
      chainId: BNB_CHAIN.MAINNET,
      marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory,
    });
  });

  it('should execute buy transaction successfully', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: 0,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const mockQuote = {
      tokensOut: 950000000000000000n,
      fee: 50000000000000000n,
    };

    const mockLiquidity = {
      yes: 1000000000000000000n,
      no: 1000000000000000000n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const MarketAMMModule = await import('../../../contracts/MarketAMM');
    const mockMarketAMM = {
      getBuyQuote: vi.fn().mockResolvedValue(mockQuote),
      getLiquidity: vi.fn().mockResolvedValue(mockLiquidity),
    };
    vi.mocked(MarketAMMModule.MarketAMM).mockImplementation(() => mockMarketAMM as any);

    const { result } = renderHook(() => useBuy(1), { wrapper });

    const params = {
      outcomeId: 0,
      amount: 1000000000000000000n,
      slippage: 0.5,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.hash).toBe(mockTransactionHash);
  });

  it('should handle errors gracefully', async () => {
    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockRejectedValue(new Error('Market not found')),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const { result } = renderHook(() => useBuy(999), { wrapper });

    const params = {
      outcomeId: 0,
      amount: 1000000000000000000n,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.error).toBeDefined();
    });
  });

  it('should not execute if wallet is not connected', async () => {
    const wagmiModule = await import('wagmi');
    vi.mocked(wagmiModule.useWalletClient).mockReturnValue({ data: null } as any);

    const { result } = renderHook(() => useBuy(1), { wrapper });

    const params = {
      outcomeId: 0,
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

  it('should apply slippage tolerance correctly', async () => {
    const mockMarketInfo = {
      marketId: 1n,
      marketAddress: '0xMarket1' as const,
      question: 'Test?',
      creator: mockAddress,
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
      status: 0,
      yesTokenId: 256n,
      noTokenId: 257n,
      totalVolume: 0n,
      totalLiquidity: { yes: 0n, no: 0n },
      createdAt: 0n,
    };

    const mockQuote = {
      tokensOut: 1000000000000000000n,
      fee: 0n,
    };

    const mockLiquidity = {
      yes: 1000000000000000000n,
      no: 1000000000000000000n,
    };

    const MarketFactoryModule = await import('../../../contracts/MarketFactory');
    const mockMarketFactory = {
      getMarket: vi.fn().mockResolvedValue(mockMarketInfo),
    };
    vi.mocked(MarketFactoryModule.MarketFactory).mockImplementation(() => mockMarketFactory as any);

    const MarketAMMModule = await import('../../../contracts/MarketAMM');
    const mockMarketAMM = {
      getBuyQuote: vi.fn().mockResolvedValue(mockQuote),
      getLiquidity: vi.fn().mockResolvedValue(mockLiquidity),
    };
    vi.mocked(MarketAMMModule.MarketAMM).mockImplementation(() => mockMarketAMM as any);

    if (mockWalletClient.writeContract) {
      vi.mocked(mockWalletClient.writeContract).mockResolvedValue(mockTransactionHash);
    }

    const { result } = renderHook(() => useBuy(1), { wrapper });

    const params = {
      outcomeId: 0,
      amount: 1000000000000000000n,
      slippage: 1.0,
    };

    result.current.write(params);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    if (mockWalletClient.writeContract) {
      expect(mockWalletClient.writeContract).toHaveBeenCalled();
      const callArgs = mockWalletClient.writeContract.mock.calls[0][0];
      expect(callArgs.args[1]).toBeLessThan(mockQuote.tokensOut);
    }
  });
});

