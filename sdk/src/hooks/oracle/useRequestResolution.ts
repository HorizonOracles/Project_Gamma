/**
 * Hook to request AI resolution via API
 */

import { useMutation } from '@tanstack/react-query';
import { useGammaConfig } from '../../components/GammaProvider';
import { createOracleApiClient, RequestResolutionParams } from '../../utils/api';

export interface RequestResolutionMutationParams {
  marketId: number;
  metadata: {
    question: string;
    description?: string;
  };
}

/**
 * Hook to request AI resolution for a market
 * 
 * @example
 * ```tsx
 * const { mutate: requestResolution, data: request } = useRequestResolution();
 * 
 * requestResolution({
 *   marketId: 1,
 *   metadata: {
 *     question: 'Market question',
 *     description: 'Details',
 *   },
 * });
 * ```
 */
export function useRequestResolution() {
  const config = useGammaConfig();

  return useMutation({
    mutationFn: async (params: RequestResolutionMutationParams) => {
      if (!config.oracleApiUrl) {
        throw new Error('Oracle API URL not configured');
      }

      const apiClient = createOracleApiClient(config.oracleApiUrl);
      return apiClient.requestResolution(params);
    },
  });
}

