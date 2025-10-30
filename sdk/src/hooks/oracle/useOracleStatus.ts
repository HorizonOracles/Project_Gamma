/**
 * Hook to get oracle request status
 */

import { useQuery } from '@tanstack/react-query';
import { useGammaConfig } from '../../components/GammaProvider';
import { OracleRequest } from '../../types';
import { createOracleApiClient } from '../../utils/api';

export interface UseOracleStatusOptions {
  enabled?: boolean;
  refetchInterval?: number;
}

/**
 * Hook to get oracle request status (with polling support)
 * 
 * @example
 * ```tsx
 * const { data: status } = useOracleStatus(requestId, {
 *   refetchInterval: 5000, // Poll every 5 seconds
 * });
 * ```
 */
export function useOracleStatus(
  requestId: string | undefined,
  options?: UseOracleStatusOptions
) {
  const config = useGammaConfig();
  const { enabled = true, refetchInterval } = options || {};

  return useQuery({
    queryKey: ['oracleStatus', requestId],
    queryFn: async (): Promise<OracleRequest> => {
      if (!requestId) {
        throw new Error('Request ID is required');
      }

      if (!config.oracleApiUrl) {
        throw new Error('Oracle API URL not configured');
      }

      const apiClient = createOracleApiClient(config.oracleApiUrl);
      return apiClient.getRequestStatus(requestId);
    },
    enabled: !!requestId && enabled && !!config.oracleApiUrl,
    refetchInterval: refetchInterval || (requestId ? 5000 : false),
  });
}

