/**
 * Hook to fetch markets with filters
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { Market, MarketStatus } from '../../types';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS } from '../../constants';

export interface UseMarketsFilters {
  category?: string;
  status?: MarketStatus;
  creator?: string;
  limit?: number;
  offset?: number;
}

/**
 * Hook to fetch markets with optional filters
 * 
 * @example
 * ```tsx
 * const { data: markets, isLoading } = useMarkets({
 *   category: 'sports',
 *   status: MarketStatus.Active,
 * });
 * ```
 */
export function useMarkets(filters?: UseMarketsFilters) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['markets', filters, chainId],
    queryFn: async (): Promise<Market[]> => {
      if (!publicClient) {
        throw new Error('Public client not available');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      const marketFactory = new MarketFactory(publicClient, marketFactoryAddress);

      // Get market count
      const marketCount = await marketFactory.getMarketCount();
      const limit = filters?.limit || Number(marketCount);
      const offset = filters?.offset || 0n;

      // Fetch markets
      const marketStructs = await marketFactory.getMarkets(BigInt(offset), BigInt(limit));

      // Convert to Market format
      const markets: Market[] = marketStructs.map((struct, index) => ({
        id: Number(struct.id),
        creator: struct.creator,
        amm: struct.amm,
        collateralToken: struct.collateralToken,
        closeTime: Number(struct.closeTime),
        category: struct.category,
        metadataURI: struct.metadataURI,
        status: struct.status as MarketStatus,
        // Additional fields
        marketId: struct.id,
        marketAddress: struct.amm,
        endTime: struct.closeTime,
      }));

      // Apply filters in memory (can be optimized with subgraph later)
      let filtered = markets;

      if (filters?.category) {
        filtered = filtered.filter((m) => m.category === filters.category);
      }

      if (filters?.status !== undefined) {
        filtered = filtered.filter((m) => m.status === filters.status);
      }

      if (filters?.creator) {
        filtered = filtered.filter((m) => m.creator.toLowerCase() === filters.creator?.toLowerCase());
      }

      return filtered;
    },
    enabled: !!publicClient,
  });
}

