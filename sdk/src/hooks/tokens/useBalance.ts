/**
 * Hook to get token balance
 */

import { useQuery } from '@tanstack/react-query';
import { useAccount, usePublicClient, useChainId } from 'wagmi';
import { useReadContract } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { Address } from 'viem';
import { DEFAULT_CONTRACTS } from '../../constants';

/**
 * Hook to get ERC20 token balance
 * 
 * @example
 * ```tsx
 * const { data: balance } = useBalance(tokenAddress);
 * ```
 */
export function useBalance(tokenAddress: Address | undefined) {
  const { address } = useAccount();
  const chainId = useChainId();

  return useReadContract({
    address: tokenAddress,
    abi: [
      {
        type: 'function',
        name: 'balanceOf',
        inputs: [{ name: 'account', type: 'address' }],
        outputs: [{ name: '', type: 'uint256' }],
        stateMutability: 'view',
      },
    ],
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: {
      enabled: !!tokenAddress && !!address,
    },
  });
}

/**
 * Hook to get outcome token balance for a specific market
 * 
 * @example
 * ```tsx
 * const { data: yesBalance } = useOutcomeBalance(marketId, 0); // YES
 * const { data: noBalance } = useOutcomeBalance(marketId, 1); // NO
 * ```
 */
export function useOutcomeBalance(marketId: number | undefined, outcomeId: number) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const chainId = useChainId();

  const outcomeTokenAddress =
    config.outcomeTokenAddress ||
    DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.outcomeToken;

  // Get token ID from market ID and outcome ID
  // Token ID encoding: (marketId << 8) | outcomeId
  const tokenId = marketId !== undefined && outcomeTokenAddress
    ? (BigInt(marketId) << 8n) | BigInt(outcomeId)
    : undefined;

  return useReadContract({
    address: outcomeTokenAddress,
    abi: [
      {
        type: 'function',
        name: 'balanceOf',
        inputs: [
          { name: 'account', type: 'address' },
          { name: 'id', type: 'uint256' },
        ],
        outputs: [{ name: '', type: 'uint256' }],
        stateMutability: 'view',
      },
    ],
    functionName: 'balanceOf',
    args: address && tokenId !== undefined ? [address, tokenId] : undefined,
    query: {
      enabled: !!outcomeTokenAddress && !!address && tokenId !== undefined,
    },
  });
}

