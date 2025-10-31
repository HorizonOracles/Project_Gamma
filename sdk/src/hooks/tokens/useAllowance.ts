/**
 * Hook to check token allowance
 */

import { useQuery } from '@tanstack/react-query';
import { useAccount, usePublicClient, useChainId } from 'wagmi';
import { Address } from 'viem';
import { ERC20_ABI } from '../../constants';

export interface UseAllowanceParams {
  tokenAddress: Address | undefined;
  spender: Address | undefined;
}

/**
 * Hook to check token allowance
 * 
 * @example
 * ```tsx
 * const { data: allowance } = useAllowance({
 *   tokenAddress: '0x...',
 *   spender: '0x...',
 * });
 * ```
 */
export function useAllowance(params: UseAllowanceParams) {
  const { address } = useAccount();
  const publicClient = usePublicClient();

  return useQuery({
    queryKey: ['allowance', params.tokenAddress, params.spender, address],
    queryFn: async (): Promise<bigint> => {
      if (!publicClient || !address || !params.tokenAddress || !params.spender) {
        throw new Error('Missing required parameters');
      }

      const result = await publicClient.readContract({
        address: params.tokenAddress,
        abi: ERC20_ABI,
        functionName: 'allowance',
        args: [address, params.spender],
      });

      return result as bigint;
    },
    enabled: !!publicClient && !!address && !!params.tokenAddress && !!params.spender,
    staleTime: 10 * 1000, // Refresh every 10 seconds
    refetchInterval: 10 * 1000, // Auto-refetch every 10 seconds
  });
}

