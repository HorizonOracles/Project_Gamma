/**
 * Hook to propose a resolution for a market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS } from '../../constants';

const RESOLUTION_MODULE_ABI = [
  {
    type: 'function',
    name: 'proposeResolution',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'outcomeId', type: 'uint256' },
      { name: 'bondAmount', type: 'uint256' },
      { name: 'evidenceURI', type: 'string' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const;

export interface ProposeResolutionParams {
  outcomeId: bigint;
  bondAmount: bigint;
  evidenceURI: string;
}

/**
 * Hook to propose a resolution for a market
 * 
 * @example
 * ```tsx
 * const { write: proposeResolution, isLoading } = useProposeResolution(marketId);
 * 
 * proposeResolution({
 *   outcomeId: 0n, // YES
 *   bondAmount: parseUnits('100', 18),
 *   evidenceURI: 'ipfs://...',
 * });
 * ```
 */
export function useProposeResolution(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const proposeResolutionMutation = useMutation({
    mutationFn: async (params: ProposeResolutionParams): Promise<string> => {
      if (!publicClient) {
        throw new Error('Public client not available');
      }

      if (!walletClient) {
        throw new Error('Wallet not connected');
      }

      if (!address) {
        throw new Error('User address not available');
      }

      const resolutionModuleAddress =
        config.resolutionModuleAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.resolutionModule;

      if (!resolutionModuleAddress || resolutionModuleAddress === '0x0000000000000000000000000000000000000000') {
        throw new Error('ResolutionModule address not configured');
      }

      const txHash = await walletClient.writeContract({
        address: resolutionModuleAddress,
        abi: RESOLUTION_MODULE_ABI,
        functionName: 'proposeResolution',
        args: [BigInt(marketId), params.outcomeId, params.bondAmount, params.evidenceURI],
      });

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['resolution', marketId] });
      queryClient.invalidateQueries({ queryKey: ['market', marketId] });
    },
  });

  return {
    mutate: (params: ProposeResolutionParams) => proposeResolutionMutation.mutate(params),
    isLoading: proposeResolutionMutation.isPending,
    isSuccess: proposeResolutionMutation.isSuccess,
    isError: proposeResolutionMutation.isError,
    hash: proposeResolutionMutation.data,
    error: proposeResolutionMutation.error,
  };
}
