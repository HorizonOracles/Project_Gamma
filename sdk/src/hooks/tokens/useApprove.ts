/**
 * Hook to approve token spending
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useWriteContract, useWaitForTransactionReceipt } from 'wagmi';
import { Address } from 'viem';

export interface ApproveParams {
  tokenAddress: Address;
  spender: Address;
  amount: bigint;
}

/**
 * Hook to approve token spending
 * 
 * @example
 * ```tsx
 * const { write: approve, isLoading } = useApprove();
 * 
 * approve({
 *   tokenAddress: '0x...',
 *   spender: '0x...',
 *   amount: parseUnits('1000', 6),
 * });
 * ```
 */
export function useApprove() {
  const queryClient = useQueryClient();
  const { writeContract, data: hash } = useWriteContract();
  const { isLoading, isSuccess } = useWaitForTransactionReceipt({
    hash,
  });

  return {
    write: (params: ApproveParams) => {
      writeContract({
        address: params.tokenAddress,
        abi: [
          {
            type: 'function',
            name: 'approve',
            inputs: [
              { name: 'spender', type: 'address' },
              { name: 'amount', type: 'uint256' },
            ],
            outputs: [{ name: '', type: 'bool' }],
            stateMutability: 'nonpayable',
          },
        ],
        functionName: 'approve',
        args: [params.spender, params.amount],
      });
    },
    isLoading,
    isSuccess,
    hash,
  };
}

