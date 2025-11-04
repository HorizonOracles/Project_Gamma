/**
 * Hook to buy outcome tokens
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { MarketOutcome } from '../../types';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS } from '../../constants';
import { applySlippageTolerance, getMarketContract } from '../../utils';

export interface BuyParams {
  outcomeId: number; // 0 for YES, 1 for NO
  amount: bigint; // Amount of collateral to spend
  slippage?: number; // Slippage tolerance percentage (default: 0.5)
  recipient?: string;
}

/**
 * Hook to buy outcome tokens
 * 
 * @example
 * ```tsx
 * const { write: buyYes, isLoading } = useBuy(marketId);
 * 
 * buyYes({
 *   outcomeId: 0, // YES
 *   amount: parseUnits('100', 6),
 *   slippage: 0.5,
 * });
 * ```
 */
export function useBuy(marketId: number) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const { address } = useAccount();
  const queryClient = useQueryClient();

  const buyMutation = useMutation({
    mutationFn: async (params: BuyParams): Promise<string> => {
      if (!publicClient) {
        throw new Error('Public client not available');
      }

      if (!walletClient) {
        throw new Error('Wallet not connected');
      }

      if (!address) {
        throw new Error('User address not available');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      // Resolve market address from factory
      const marketFactory = new MarketFactory(publicClient, marketFactoryAddress);
      const marketInfo = await marketFactory.getMarket(BigInt(marketId));
      const marketAddress = marketInfo.marketAddress;

      // Instantiate correct market contract based on type (auto-detects)
      const market = await getMarketContract(publicClient, marketAddress, walletClient);
      const outcome: MarketOutcome = params.outcomeId === 0 ? 'YES' : 'NO';

      // Check liquidity before attempting trade
      const reserves = await market.getReserves();
      if (reserves.yes === 0n || reserves.no === 0n) {
        throw new Error('Market has no liquidity. Please add liquidity before trading.');
      }

      // Get quote to calculate minimum amount out
      const quote = await market.getBuyQuote(params.amount, outcome, address);
      
      // Validate quote is valid
      if (quote.tokensOut === 0n) {
        throw new Error('Invalid quote: tokensOut is zero. The market may have insufficient liquidity for this trade amount.');
      }

      const slippageBps = Math.round((params.slippage || 0.5) * 100); // Convert to basis points
      const minAmountOut = applySlippageTolerance(quote.tokensOut, slippageBps);

      // Validate minAmountOut is not zero
      if (minAmountOut === 0n) {
        throw new Error('Minimum tokens out is zero. Try reducing slippage tolerance or trade amount.');
      }

      // Execute trade using the market contract's buyTokens method
      const result = await market.buyTokens({
        marketId: BigInt(marketId),
        outcome,
        amount: params.amount,
        minAmountOut,
      });

      return result.transactionHash!;
    },
    onSuccess: () => {
      // Invalidate relevant queries
      queryClient.invalidateQueries({ queryKey: ['market', marketId] });
      queryClient.invalidateQueries({ queryKey: ['prices', marketId] });
      queryClient.invalidateQueries({ queryKey: ['quote'] });
      queryClient.invalidateQueries({ queryKey: ['balance'] });
    },
  });

  return {
    write: (params: BuyParams) => buyMutation.mutate(params),
    isLoading: buyMutation.isPending,
    isSuccess: buyMutation.isSuccess,
    isError: buyMutation.isError,
    hash: buyMutation.data,
    error: buyMutation.error,
  };
}

