/**
 * Hook to get user's LP position in a market
 */

import { useQuery } from '@tanstack/react-query';
import { useAccount, usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI } from '../../constants';
import { getMarketContract } from '../../utils/markets';

export interface LPPosition {
  lpTokens: bigint;
  value: bigint; // Value in collateral tokens
  share: number; // Percentage share of pool (0-1)
}

/**
 * Hook to get user's LP position in a market
 * Works with all market types (Binary, MultiChoice, etc.)
 * 
 * @example
 * ```tsx
 * const { data: position } = useLPPosition(marketId);
 * 
 * // position.lpTokens - LP tokens owned
 * // position.value - Value in collateral
 * // position.share - Pool share percentage
 * ```
 */
export function useLPPosition(marketId: number | undefined) {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const chainId = useChainId();

  return useQuery({
    queryKey: ['lpPosition', marketId, address, chainId],
    queryFn: async (): Promise<LPPosition> => {
      if (!marketId) {
        throw new Error('Market ID is required');
      }

      if (!address) {
        throw new Error('User address is required');
      }

      if (!publicClient) {
        throw new Error('Public client not available');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      // Get market info to find market address
      const marketStruct = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [BigInt(marketId)],
      });

      const marketAddress = marketStruct.amm;

      // Get the correct market contract based on type
      const marketContract = await getMarketContract(publicClient, marketAddress);

      // Get LP balance and reserves using contract class methods
      const [lpBalance, reserves, totalCollateral] = await Promise.all([
        marketContract.getLPBalance(address),
        marketContract.getReserves(),
        marketContract.getTotalCollateral(),
      ]);

      // Calculate share percentage
      // totalCollateral represents total LP supply
      const share = totalCollateral > 0n 
        ? Number((lpBalance * 10000n) / totalCollateral) / 100 
        : 0;

      // Calculate value: share of total liquidity in collateral tokens
      // Total liquidity value = yesReserve + noReserve (they should be equal in balanced pool)
      const totalLiquidityValue = reserves.yes + reserves.no;
      const value = totalCollateral > 0n
        ? (lpBalance * totalLiquidityValue) / totalCollateral
        : 0n;

      return {
        lpTokens: lpBalance,
        value,
        share,
      };
    },
    enabled: !!marketId && !!address && !!publicClient,
  });
}

