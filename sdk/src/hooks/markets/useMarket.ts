/**
 * Hook to fetch a single market by ID
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { Address } from 'viem';
import { useGammaConfig } from '../../components/GammaProvider';
import { Market, MarketInfo } from '../../types';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS } from '../../constants';

/**
 * Hook to fetch a single market
 * 
 * @example
 * ```tsx
 * const { data: market } = useMarket(1);
 * ```
 */
export function useMarket(marketId: number | undefined) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['market', marketId, chainId],
    queryFn: async (): Promise<Market> => {
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

      const marketFactory = new MarketFactory(publicClient, marketFactoryAddress);
      const marketInfo: MarketInfo = await marketFactory.getMarket(BigInt(marketId));

      return {
        id: Number(marketInfo.marketId),
        creator: marketInfo.creator,
        amm: marketInfo.marketAddress,
        collateralToken: marketInfo.collateralToken || ('0x0000000000000000000000000000000000000000' as Address),
        closeTime: Number(marketInfo.endTime),
        category: marketInfo.category || '',
        metadataURI: marketInfo.metadataURI || '',
        status: marketInfo.status,
        // Additional fields
        marketId: marketInfo.marketId,
        marketAddress: marketInfo.marketAddress,
        question: marketInfo.question,
        description: marketInfo.description,
        endTime: marketInfo.endTime,
        yesTokenId: marketInfo.yesTokenId,
        noTokenId: marketInfo.noTokenId,
        totalVolume: marketInfo.totalVolume,
        totalLiquidity: marketInfo.totalLiquidity,
        createdAt: marketInfo.createdAt,
      };
    },
    enabled: !!marketId && !!publicClient,
  });
}

