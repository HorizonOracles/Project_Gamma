/**
 * Main entry point for Project Gamma React SDK
 * Simple, React-focused SDK for decentralized prediction markets
 */

// React Components
export { GammaProvider, useGammaConfig } from './components/GammaProvider';

// Market Hooks
export { useMarkets } from './hooks/markets/useMarkets';
export type { UseMarketsFilters } from './hooks/markets/useMarkets';
export { useMarket } from './hooks/markets/useMarket';
export { useCreateMarket } from './hooks/markets/useCreateMarket';
export { useMinCreatorStake } from './hooks/markets/useMinCreatorStake';
export { useUploadMetadata } from './hooks/markets/useUploadMetadata';
export type { UseUploadMetadataParams } from './hooks/markets/useUploadMetadata';
export { useIPFSMetadata } from './hooks/markets/useIPFSMetadata';
export { useHasLiquidity } from './hooks/markets/useHasLiquidity';

// Trading Hooks
export { useBuy } from './hooks/trading/useBuy';
export type { BuyParams } from './hooks/trading/useBuy';
export { useSell } from './hooks/trading/useSell';
export type { SellParams } from './hooks/trading/useSell';
export { useQuote } from './hooks/trading/useQuote';
export type { QuoteParams } from './hooks/trading/useQuote';
export { usePrices } from './hooks/trading/usePrices';

// Liquidity Hooks
export { useAddLiquidity } from './hooks/liquidity/useAddLiquidity';
export type { AddLiquidityParams } from './hooks/liquidity/useAddLiquidity';
export { useRemoveLiquidity } from './hooks/liquidity/useRemoveLiquidity';
export type { RemoveLiquidityParams } from './hooks/liquidity/useRemoveLiquidity';
export { useLPPosition } from './hooks/liquidity/useLPPosition';
export type { LPPosition } from './hooks/liquidity/useLPPosition';

// Resolution Hooks
export { useResolution } from './hooks/resolution/useResolution';
export { useProposeResolution } from './hooks/resolution/useProposeResolution';
export type { ProposeResolutionParams } from './hooks/resolution/useProposeResolution';
export { useDispute } from './hooks/resolution/useDispute';
export type { DisputeParams } from './hooks/resolution/useDispute';
export { useFinalize } from './hooks/resolution/useFinalize';

// Oracle Hooks
export { useRequestResolution } from './hooks/oracle/useRequestResolution';
export type { RequestResolutionMutationParams } from './hooks/oracle/useRequestResolution';
export { useOracleStatus } from './hooks/oracle/useOracleStatus';
export type { UseOracleStatusOptions } from './hooks/oracle/useOracleStatus';
export { useOracleResult } from './hooks/oracle/useOracleResult';
export type { UseOracleResultOptions } from './hooks/oracle/useOracleResult';
export { useOracleHistory } from './hooks/oracle/useOracleHistory';

// Token Hooks
export { useBalance, useOutcomeBalance } from './hooks/tokens/useBalance';
export { useApprove } from './hooks/tokens/useApprove';
export type { ApproveParams } from './hooks/tokens/useApprove';
export { useAllowance } from './hooks/tokens/useAllowance';
export { useRedeem } from './hooks/tokens/useRedeem';
export type { RedeemParams } from './hooks/tokens/useRedeem';

// Types
export type {
  Market,
  MarketOutcome,
  CreateMarketParams,
  TradeParams,
  TradeResult,
  TradeQuote,
  UserPosition,
  MarketPrices,
  ResolutionProposal,
  FeeTier,
  OracleRequest,
  OracleResult,
  SDKError,
  ContractError,
  TradeError,
} from './types';

export { MarketType, MarketStatus } from './types';

// Utilities
export {
  formatTokenAmount,
  parseTokenAmount,
  calculatePrice,
  calculateMarketPrices,
  calculateAmountOut,
  calculateAmountIn,
  calculateSlippage,
  isValidAddress,
  shortenAddress,
  getOutcomeTokenId,
  getMarketIdFromTokenId,
  getOutcomeFromTokenId,
  applySlippageTolerance,
} from './utils';

// IPFS utilities
export {
  uploadMarketMetadata,
  uploadToPinata,
  uploadToWeb3Storage,
  uploadToPublicIPFS,
} from './utils/ipfs';
export type {
  MarketMetadata,
  IPFSUploadResult,
  IPFSProvider,
} from './utils/ipfs';

// API Client
export { createOracleApiClient } from './utils/api';
export type { OracleApiClient, RequestResolutionParams } from './utils/api';

// Constants
export {
  BNB_CHAIN,
  DEFAULT_CONTRACTS,
  DEFAULT_CONFIG,
  DEFAULT_TESTNET_CONFIG,
} from './constants';

