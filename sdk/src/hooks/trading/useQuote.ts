/**
 * Hook to get price quotes for trading
 */

import { useQuery } from '@tanstack/react-query';
import { useAccount, usePublicClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { TradeQuote, MarketOutcome } from '../../types';
import { MarketAMM } from '../../contracts/MarketAMM';
import { MarketFactory } from '../../contracts/MarketFactory';
import { DEFAULT_CONTRACTS } from '../../constants';

export interface QuoteParams {
  marketId: number;
  outcomeId: number; // 0 for YES, 1 for NO
  amount: bigint;
  isBuy: boolean; // true for buy, false for sell
}

/**
 * Hook to get a trade quote
 * 
 * @example
 * ```tsx
 * const { data: quote } = useQuote({
 *   marketId: 1,
 *   outcomeId: 0,
 *   amount: parseUnits('100', 6),
 *   isBuy: true,
 * });
 * 
 * // quote.tokensOut - tokens you'll receive
 * // quote.fee - fee amount
 * // quote.priceImpact - price impact percentage
 * ```
 */
export function useQuote(params: QuoteParams | undefined) {
  const config = useGammaConfig();
  const publicClient = usePublicClient();
  const chainId = useChainId();
  const { address } = useAccount();

  return useQuery({
    queryKey: ['quote', params, chainId, address],
    queryFn: async (): Promise<TradeQuote> => {
      if (!params) {
        throw new Error('Quote parameters are required');
      }

      if (!publicClient) {
        throw new Error('Public client not available');
      }

      if (!address) {
        throw new Error('User address is required for quotes');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      // Resolve market AMM address
      const marketFactory = new MarketFactory(publicClient, marketFactoryAddress);
      const marketInfo = await marketFactory.getMarket(BigInt(params.marketId));
      const ammAddress = marketInfo.marketAddress;

      // Create MarketAMM instance
      const amm = new MarketAMM(publicClient, ammAddress);
      const outcome: MarketOutcome = params.outcomeId === 0 ? 'YES' : 'NO';

      // Get current prices for price impact calculation
      const currentPrices = await amm.getMarketPrices(BigInt(params.marketId));

      let tokensOut: bigint;
      let fee: bigint;
      let priceImpact = 0;

      if (params.isBuy) {
        // Get buy quote
        const quote = await amm.getBuyQuote(params.amount, outcome, address);
        tokensOut = quote.tokensOut;
        fee = quote.fee;

        // Calculate price impact: ((currentPrice - executionPrice) / currentPrice) * 100
        const currentPrice = outcome === 'YES' ? currentPrices.yesPrice : currentPrices.noPrice;
        const executionPrice = params.amount > 0n 
          ? (tokensOut * 10n**18n) / params.amount 
          : currentPrice;
        
        if (currentPrice > 0n && executionPrice < currentPrice) {
          const priceDiff = currentPrice - executionPrice;
          priceImpact = Number((priceDiff * 10000n) / currentPrice) / 100;
        }
      } else {
        // Get sell quote
        const quote = await amm.getSellQuote(params.amount, outcome, address);
        tokensOut = quote.collateralOut;
        fee = quote.fee;

        // Calculate price impact for sell
        const currentPrice = outcome === 'YES' ? currentPrices.yesPrice : currentPrices.noPrice;
        const executionPrice = params.amount > 0n
          ? (tokensOut * 10n**18n) / params.amount
          : currentPrice;
        
        if (currentPrice > 0n && executionPrice < currentPrice) {
          const priceDiff = currentPrice - executionPrice;
          priceImpact = Number((priceDiff * 10000n) / currentPrice) / 100;
        }
      }

      return {
        tokensOut,
        fee,
        priceImpact,
      };
    },
    enabled: !!params && !!publicClient && !!address,
    refetchInterval: 10000, // Refetch every 10 seconds for accurate quotes
  });
}

