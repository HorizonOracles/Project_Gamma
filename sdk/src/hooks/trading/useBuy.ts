/**
 * Hook to buy outcome tokens
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { MarketOutcome } from '../../types';
import { MarketAMM } from '../../contracts/MarketAMM';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS, MARKET_AMM_ABI } from '../../constants';
import { applySlippageTolerance } from '../../utils';

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

      // Resolve market AMM address
      const marketFactory = new MarketFactory(publicClient, marketFactoryAddress);
      const marketInfo = await marketFactory.getMarket(BigInt(marketId));
      const ammAddress = marketInfo.marketAddress;

      // Create MarketAMM instance
      const amm = new MarketAMM(publicClient, ammAddress);
      const outcome: MarketOutcome = params.outcomeId === 0 ? 'YES' : 'NO';

      // Get quote to calculate minimum amount out
      const quote = await amm.getBuyQuote(params.amount, outcome, address);
      const slippageBps = Math.round((params.slippage || 0.5) * 100); // Convert to basis points
      const minAmountOut = applySlippageTolerance(quote.tokensOut, slippageBps);

      // Execute trade via walletClient.writeContract
      const functionName = outcome === 'YES' ? 'buyYes' : 'buyNo';
      const txHash = await walletClient.writeContract({
        address: ammAddress,
        abi: MARKET_AMM_ABI,
        functionName,
        args: [params.amount, minAmountOut],
      });

      return txHash;
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

