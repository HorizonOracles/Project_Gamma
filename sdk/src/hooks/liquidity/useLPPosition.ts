/**
 * Hook to get user's LP position in a market
 */

import { useQuery } from '@tanstack/react-query';
import { useAccount, usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, MARKET_AMM_ABI } from '../../constants';

export interface LPPosition {
  lpTokens: bigint;
  value: bigint; // Value in collateral tokens
  share: number; // Percentage share of pool (0-1)
}

/**
 * Hook to get user's LP position in a market
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

      // Get market info to find AMM address
      const marketStruct = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'getMarket',
        args: [BigInt(marketId)],
      });

      const ammAddress = marketStruct.amm;

      // Get LP balance, reserves, and total LP supply
      const [lpBalance, reserveYes, reserveNo, totalCollateral] = await Promise.all([
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'balanceOf',
          args: [address],
        }),
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'reserveYes',
        }),
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'reserveNo',
        }),
        publicClient.readContract({
          address: ammAddress,
          abi: MARKET_AMM_ABI,
          functionName: 'totalCollateral',
        }),
      ]);

      const lpBalanceBigInt = lpBalance as bigint;
      const reserveYesBigInt = reserveYes as bigint;
      const reserveNoBigInt = reserveNo as bigint;
      const totalCollateralBigInt = totalCollateral as bigint;

      // Calculate share percentage
      // Note: totalCollateral represents total LP supply in this AMM
      const share = totalCollateralBigInt > 0n 
        ? Number((lpBalanceBigInt * 10000n) / totalCollateralBigInt) / 100 
        : 0;

      // Calculate value: share of total liquidity in collateral tokens
      // Total liquidity value = yesReserve + noReserve (they should be equal in balanced pool)
      const totalLiquidityValue = reserveYesBigInt + reserveNoBigInt;
      const value = totalCollateralBigInt > 0n
        ? (lpBalanceBigInt * totalLiquidityValue) / totalCollateralBigInt
        : 0n;

      return {
        lpTokens: lpBalanceBigInt,
        value,
        share,
      };
    },
    enabled: !!marketId && !!address && !!publicClient,
  });
}

