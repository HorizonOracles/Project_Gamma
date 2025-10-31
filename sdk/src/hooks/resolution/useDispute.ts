/**
 * Hook to dispute a proposed resolution
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS } from '../../constants';

const RESOLUTION_MODULE_ABI = [
  {
    type: 'function',
    name: 'dispute',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'bondAmount', type: 'uint256' },
      { name: 'reason', type: 'string' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const;

export interface DisputeParams {
  bondAmount: bigint;
  reason: string;
}

/**
 * Hook to dispute a proposed resolution
 * 
 * @example
 * ```tsx
 * const { write: dispute, isLoading } = useDispute(marketId);
 * 
 * dispute({
 *   bondAmount: parseUnits('100', 18),
 *   reason: 'Evidence is incorrect',
 * });
 * ```
 */
export function useDispute(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const disputeMutation = useMutation({
    mutationFn: async (params: DisputeParams): Promise<string> => {
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
        functionName: 'dispute',
        args: [BigInt(marketId), params.bondAmount, params.reason],
      });

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['resolution', marketId] });
      queryClient.invalidateQueries({ queryKey: ['market', marketId] });
    },
  });

  return {
    mutate: (params: DisputeParams) => disputeMutation.mutate(params),
    isLoading: disputeMutation.isPending,
    isSuccess: disputeMutation.isSuccess,
    hash: disputeMutation.data,
    error: disputeMutation.error,
  };
}
