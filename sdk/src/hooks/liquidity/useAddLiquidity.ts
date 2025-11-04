/**
 * Hook to add liquidity to a market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI } from '../../constants';
import { getMarketContract } from '../../utils/markets';

export interface AddLiquidityParams {
  amount: bigint; // Amount of collateral to add
}

/**
 * Hook to add liquidity to a market
 * Works with all market types (Binary, MultiChoice, etc.)
 * 
 * @example
 * ```tsx
 * const { write: addLiquidity, isLoading } = useAddLiquidity(marketId);
 * 
 * addLiquidity({
 *   amount: parseUnits('1000', 6), // 1000 USDC
 * });
 * ```
 */
export function useAddLiquidity(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const addLiquidityMutation = useMutation({
    mutationFn: async (params: AddLiquidityParams): Promise<string> => {
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

      // Execute addLiquidity using contract class method
      // BinaryMarket.addLiquidity handles approval internally
      const txHash = await marketContract.addLiquidity(params.amount);

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['liquidity', marketId] });
      queryClient.invalidateQueries({ queryKey: ['lpPosition', marketId] });
      queryClient.invalidateQueries({ queryKey: ['prices', marketId] });
    },
  });

  return {
    write: (params: AddLiquidityParams) => addLiquidityMutation.mutate(params),
    isLoading: addLiquidityMutation.isPending,
    isSuccess: addLiquidityMutation.isSuccess,
    hash: addLiquidityMutation.data,
    error: addLiquidityMutation.error,
  };
}

