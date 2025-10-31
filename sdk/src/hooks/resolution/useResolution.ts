/**
 * Hook to get resolution state for a market
 */

import { useQuery } from '@tanstack/react-query';
import { usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS } from '../../constants';
import { Resolution, ResolutionState } from '../../contracts/ResolutionModule';

const RESOLUTION_MODULE_ABI = [
  {
    type: 'function',
    name: 'resolutions',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [
      { name: 'state', type: 'uint8' },
      { name: 'proposedOutcome', type: 'uint256' },
      { name: 'proposalTime', type: 'uint256' },
      { name: 'proposer', type: 'address' },
      { name: 'proposerBond', type: 'uint256' },
      { name: 'disputer', type: 'address' },
      { name: 'disputerBond', type: 'uint256' },
      { name: 'evidenceURI', type: 'string' },
    ],
    stateMutability: 'view',
  },
] as const;

/**
 * Hook to get resolution state for a market
 * 
 * @example
 * ```tsx
 * const { data: resolution } = useResolution(marketId);
 * ```
 */
export function useResolution(marketId: number | undefined) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['resolution', marketId, chainId],
    queryFn: async (): Promise<Resolution | null> => {
      if (!marketId) {
        throw new Error('Market ID is required');
      }

      if (!publicClient) {
        throw new Error('Public client not available');
      }

      const resolutionModuleAddress =
        config.resolutionModuleAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.resolutionModule;

      if (!resolutionModuleAddress || resolutionModuleAddress === '0x0000000000000000000000000000000000000000') {
        return null;
      }

      try {
        const result = await publicClient.readContract({
          address: resolutionModuleAddress,
          abi: RESOLUTION_MODULE_ABI,
          functionName: 'resolutions',
          args: [BigInt(marketId)],
        });

        const [
          state,
          proposedOutcome,
          proposalTime,
          proposer,
          proposerBond,
          disputer,
          disputerBond,
          evidenceURI,
        ] = result as [
          number,
          bigint,
          bigint,
          string,
          bigint,
          string,
          bigint,
          string,
        ];

        return {
          state: state as ResolutionState,
          proposedOutcome,
          proposalTime,
          proposer: proposer as `0x${string}`,
          proposerBond,
          disputer: disputer as `0x${string}`,
          disputerBond,
          evidenceURI,
        };
      } catch (error) {
        // If resolution doesn't exist, return null
        return null;
      }
    },
    enabled: !!marketId && !!publicClient && !!config.resolutionModuleAddress,
  });
}

