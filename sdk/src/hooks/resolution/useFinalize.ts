/**
 * Hook to finalize a resolution
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS } from '../../constants';

const RESOLUTION_MODULE_ABI = [
  {
    type: 'function',
    name: 'finalize',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const;

/**
 * Hook to finalize a resolution
 * 
 * @example
 * ```tsx
 * const { write: finalize, isLoading } = useFinalize(marketId);
 * 
 * finalize();
 * ```
 */
export function useFinalize(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const finalizeMutation = useMutation({
    mutationFn: async (): Promise<string> => {
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
        functionName: 'finalize',
        args: [BigInt(marketId)],
      });

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['resolution', marketId] });
      queryClient.invalidateQueries({ queryKey: ['market', marketId] });
    },
  });

  return {
    mutate: () => finalizeMutation.mutate(),
    isLoading: finalizeMutation.isPending,
    isSuccess: finalizeMutation.isSuccess,
    hash: finalizeMutation.data,
    error: finalizeMutation.error,
  };
}
