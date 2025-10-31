/**
 * Hook to check if a market has liquidity
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, MARKET_AMM_ABI } from '../../constants';

/**
 * Hook to check if a market has liquidity
 * 
 * @example
 * ```tsx
 * const { data: hasLiquidity } = useHasLiquidity(marketId);
 * 
 * if (!hasLiquidity) {
 *   return <div>No liquidity available</div>;
 * }
 * ```
 */
export function useHasLiquidity(marketId: number | undefined) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['hasLiquidity', marketId, chainId],
    queryFn: async (): Promise<boolean> => {
      if (!marketId) {
        return false;
      }

      if (!publicClient) {
        return false;
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        return false;
      }

      // Get market info to find AMM address
      const marketStruct = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [BigInt(marketId)],
      });

      const ammAddress = marketStruct.amm;

      // Check if there's liquidity by checking totalCollateral
      // A market has liquidity if totalCollateral > 0
      const totalCollateral = await publicClient.readContract({
        address: ammAddress,
        abi: MARKET_AMM_ABI,
        functionName: 'totalCollateral',
      });

      return (totalCollateral as bigint) > 0n;
    },
    enabled: !!marketId && !!publicClient,
    refetchInterval: 10000, // Refetch every 10 seconds
  });
}

