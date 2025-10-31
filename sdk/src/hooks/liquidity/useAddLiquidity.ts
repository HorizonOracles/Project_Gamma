/**
 * Hook to add liquidity to a market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, MARKET_AMM_ABI, ERC20_ABI } from '../../constants';

export interface AddLiquidityParams {
  amount: bigint; // Amount of collateral to add
}

/**
 * Hook to add liquidity to a market
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

      // Get market info to find AMM address
      const marketStruct = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [BigInt(marketId)],
      });

      const ammAddress = marketStruct.amm;
      const collateralToken = marketStruct.collateralToken;

      // Check and approve collateral token if needed
      const currentAllowance = await publicClient.readContract({
        address: collateralToken,
        abi: ERC20_ABI,
        functionName: 'allowance',
        args: [address, ammAddress],
      }) as bigint;

      if (currentAllowance < params.amount) {
        const approveHash = await walletClient.writeContract({
          address: collateralToken,
          abi: ERC20_ABI,
          functionName: 'approve',
          args: [ammAddress, params.amount],
        });
        await publicClient.waitForTransactionReceipt({ hash: approveHash });
      }

      // Execute addLiquidity transaction
      const txHash = await walletClient.writeContract({
        address: ammAddress,
        abi: MARKET_AMM_ABI,
        functionName: 'addLiquidity',
        args: [params.amount],
      });

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

