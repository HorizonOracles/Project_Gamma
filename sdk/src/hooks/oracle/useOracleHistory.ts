/**
 * Hook to get oracle history for a market
 */

import { useQuery } from '@tanstack/react-query';
import { useGammaConfig } from '../../components/GammaProvider';
import { OracleRequest } from '../../types';
import { createOracleApiClient } from '../../utils/api';

/**
 * Hook to get oracle request history for a market
 * 
 * @example
 * ```tsx
 * const { data: history, isLoading } = useOracleHistory(marketId);
 * ```
 */
export function useOracleHistory(marketId: number | undefined) {
  const config = useGammaConfig();

  return useQuery({
    queryKey: ['oracleHistory', marketId],
    queryFn: async (): Promise<OracleRequest[]> => {
      if (!marketId) {
        throw new Error('Market ID is required');
      }

      if (!config.oracleApiUrl) {
        throw new Error('Oracle API URL not configured');
      }

      const apiClient = createOracleApiClient(config.oracleApiUrl);
      return apiClient.getMarketHistory(marketId);
    },
    enabled: !!marketId && !!config.oracleApiUrl,
  });
}

