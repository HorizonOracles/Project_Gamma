/**
 * Hook to remove liquidity from a market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI } from '../../constants';
import { getMarketContract } from '../../utils/markets';

export interface RemoveLiquidityParams {
  lpTokens: bigint; // Amount of LP tokens to remove
}

/**
 * Hook to remove liquidity from a market
 * Works with all market types (Binary, MultiChoice, etc.)
 * 
 * @example
 * ```tsx
 * const { write: removeLiquidity, isLoading } = useRemoveLiquidity(marketId);
 * 
 * removeLiquidity({
 *   lpTokens: parseUnits('100', 18),
 * });
 * ```
 */
export function useRemoveLiquidity(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const removeLiquidityMutation = useMutation({
    mutationFn: async (params: RemoveLiquidityParams): Promise<string> => {
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

      // Get market info to find market address
      const marketStruct = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [BigInt(marketId)],
      });

      const marketAddress = marketStruct.amm;

      // Get the correct market contract based on type
      const marketContract = await getMarketContract(publicClient, marketAddress, walletClient);

      // Execute removeLiquidity using contract class method
      const txHash = await marketContract.removeLiquidity(params.lpTokens);

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['liquidity', marketId] });
      queryClient.invalidateQueries({ queryKey: ['lpPosition', marketId] });
      queryClient.invalidateQueries({ queryKey: ['prices', marketId] });
    },
  });

  return {
    write: (params: RemoveLiquidityParams) => removeLiquidityMutation.mutate(params),
    isLoading: removeLiquidityMutation.isPending,
    isSuccess: removeLiquidityMutation.isSuccess,
    hash: removeLiquidityMutation.data,
    error: removeLiquidityMutation.error,
  };
}

