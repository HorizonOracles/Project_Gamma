/**
 * Hook to get oracle result for a completed request
 */

import { useQuery } from '@tanstack/react-query';
import { useGammaConfig } from '../../components/GammaProvider';
import { OracleResult } from '../../types';
import { createOracleApiClient } from '../../utils/api';

export interface UseOracleResultOptions {
  enabled?: boolean;
}

/**
 * Hook to get oracle result for a completed request
 * 
 * @example
 * ```tsx
 * const { data: result } = useOracleResult(requestId, {
 *   enabled: status?.status === 'completed',
 * });
 * ```
 */
export function useOracleResult(
  requestId: string | undefined,
  options?: UseOracleResultOptions
) {
  const config = useGammaConfig();
  const { enabled = true } = options || {};

  return useQuery({
    queryKey: ['oracleResult', requestId],
    queryFn: async (): Promise<OracleResult> => {
      if (!requestId) {
        throw new Error('Request ID is required');
      }

      if (!config.oracleApiUrl) {
        throw new Error('Oracle API URL not configured');
      }

      const apiClient = createOracleApiClient(config.oracleApiUrl);
      return apiClient.getRequestResult(requestId);
    },
    enabled: !!requestId && enabled && !!config.oracleApiUrl,
  });
}

