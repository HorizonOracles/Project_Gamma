/**
 * Hook to redeem winning tokens after market resolution
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, OUTCOME_TOKEN_ABI } from '../../constants';

export interface RedeemParams {
  amount?: bigint; // Optional: redeem specific amount, otherwise redeem all
}

/**
 * Hook to redeem winning tokens after market resolution
 * 
 * @example
 * ```tsx
 * const { write: redeem, isLoading } = useRedeem(marketId);
 * 
 * redeem(); // Redeem all
 * // or
 * redeem({ amount: parseUnits('100', 18) }); // Redeem specific amount
 * ```
 */
export function useRedeem(marketId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const redeemMutation = useMutation({
    mutationFn: async (params?: RedeemParams): Promise<string> => {
      if (!publicClient) {
        throw new Error('Public client not available');
      }

      if (!walletClient) {
        throw new Error('Wallet not connected');
      }

      if (!address) {
        throw new Error('User address not available');
      }

      const outcomeTokenAddress =
        config.outcomeTokenAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.outcomeToken;

      if (!outcomeTokenAddress || outcomeTokenAddress === '0x0000000000000000000000000000000000000000') {
        throw new Error('OutcomeToken address not configured');
      }

      // Use redeemAmount if amount is specified, otherwise use redeem
      let txHash: string;
      if (params?.amount) {
        txHash = await walletClient.writeContract({
          address: outcomeTokenAddress,
          abi: OUTCOME_TOKEN_ABI,
          functionName: 'redeemAmount',
          args: [BigInt(marketId), params.amount],
        });
      } else {
        txHash = await walletClient.writeContract({
          address: outcomeTokenAddress,
          abi: OUTCOME_TOKEN_ABI,
          functionName: 'redeem',
          args: [BigInt(marketId)],
        });
      }

      return txHash;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['balance'] });
      queryClient.invalidateQueries({ queryKey: ['market', marketId] });
    },
  });

  return {
    mutate: (params?: RedeemParams) => redeemMutation.mutate(params),
    isLoading: redeemMutation.isPending,
    isSuccess: redeemMutation.isSuccess,
    isError: redeemMutation.isError,
    hash: redeemMutation.data,
    error: redeemMutation.error,
  };
}
