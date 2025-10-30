/**
 * Hook to sell outcome tokens
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { MarketOutcome } from '../../types';
import { MarketAMM } from '../../contracts/MarketAMM';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS, MARKET_AMM_ABI } from '../../constants';
import { applySlippageTolerance } from '../../utils';

export interface SellParams {
  outcomeId: number; // 0 for YES, 1 for NO
  amount: bigint; // Amount of outcome tokens to sell
  slippage?: number; // Slippage tolerance percentage (default: 0.5)
  recipient?: string;
}

/**
 * Hook to sell outcome tokens
 * 
 * @example
 * ```tsx
 * const { write: sellYes, isLoading } = useSell(marketId);
 * 
 * sellYes({
 *   outcomeId: 0, // YES
 *   amount: parseUnits('50', 18),
 *   slippage: 0.5,
 * });
 * ```
 */
export function useSell(marketId: number) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const { address } = useAccount();
  const queryClient = useQueryClient();

  const sellMutation = useMutation({
    mutationFn: async (params: SellParams): Promise<string> => {
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
      const quote = await amm.getSellQuote(params.amount, outcome, address);
      const slippageBps = Math.round((params.slippage || 0.5) * 100); // Convert to basis points
      const minAmountOut = applySlippageTolerance(quote.collateralOut, slippageBps);

      // Execute trade via walletClient.writeContract
      const functionName = outcome === 'YES' ? 'sellYes' : 'sellNo';
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
    write: (params: SellParams) => sellMutation.mutate(params),
    isLoading: sellMutation.isPending,
    isSuccess: sellMutation.isSuccess,
    isError: sellMutation.isError,
    hash: sellMutation.data,
    error: sellMutation.error,
  };
}

