/**
 * Hook to get current market prices
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { MarketPrices } from '../../types';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, MARKET_AMM_ABI } from '../../constants';

/**
 * Hook to get current prices for YES and NO outcomes
 * 
 * @example
 * ```tsx
 * const { data: prices } = usePrices(marketId);
 * 
 * // prices.yes - YES price (0-1)
 * // prices.no - NO price (0-1)
 * ```
 */
export function usePrices(marketId: number | undefined) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['prices', marketId, chainId],
    queryFn: async (): Promise<MarketPrices> => {
      if (!marketId) {
        throw new Error('Market ID is required');
      }

      if (!publicClient) {
        throw new Error('Public client not available');
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

      // Get prices from AMM contract
      const [yesPrice, noPrice] = await Promise.all([
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'getYesPrice',
        }),
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'getNoPrice',
        }),
      ]);

      // Convert to simple 0-1 format
      const yesDecimal = Number(yesPrice) / 1e18;
      const noDecimal = Number(noPrice) / 1e18;

      return {
        yesPrice: yesPrice as bigint,
        noPrice: noPrice as bigint,
        yes: yesDecimal,
        no: noDecimal,
      };
    },
    enabled: !!marketId && !!publicClient,
    refetchInterval: 5000, // Refetch every 5 seconds
  });
}

