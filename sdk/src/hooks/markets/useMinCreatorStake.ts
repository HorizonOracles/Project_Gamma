/**
 * Hook to fetch the minimum creator stake required to create a market
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI } from '../../constants';

/**
 * Hook to get minimum creator stake
 * 
 * @example
 * ```tsx
 * const { data: minStake } = useMinCreatorStake();
 * ```
 */
export function useMinCreatorStake() {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['minCreatorStake', chainId],
    queryFn: async (): Promise<bigint> => {
      if (!publicClient) {
        throw new Error('Public client not available');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      const result = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'minCreatorStake',
      });

      return result as bigint;
    },
    enabled: !!publicClient,
    staleTime: 5 * 60 * 1000, // Cache for 5 minutes
  });
}

