/**
 * Core types and interfaces for Project Gamma SDK
 */

import { Address } from 'viem';

/**
 * Binary market outcome type
 */
export type MarketOutcome = 'YES' | 'NO';

/**
 * Market type enumeration matching IMarket.MarketType
 */
export enum MarketType {
  Binary = 0, // Traditional Yes/No market (MarketAMM)
  MultiChoice = 1, // 3-8 discrete outcomes (MultiChoiceMarket)
  LimitOrder = 2, // Order book based market (LimitOrderMarket)
  PooledLiquidity = 3, // Concentrated liquidity AMM (PooledLiquidityMarket)
  Dependent = 4, // Conditional market with parent dependency
  Bracket = 5, // Value range prediction
  Trend = 6, // Time-weighted outcome tracking
}

/**
 * Market status enumeration - simplified to match contract state
 */
export enum MarketStatus {
  Active = 0,
  Closed = 1,
  Resolved = 2,
  Invalid = 3,
}

/**
 * Market creation parameters
 * Note: For actual contract calls, these fields are required:
 * - collateralToken: Address of ERC20 token to use as collateral
 * - category: Market category string
 * - metadataURI: IPFS URI containing market question and details
 * - creatorStake: Amount of HORIZON tokens to stake (must be >= minCreatorStake)
 */
export interface CreateMarketParams {
  question: string;
  description?: string;
  endTime: bigint; // Unix timestamp (maps to closeTime in contract)
  resolutionData?: string; // Additional resolution metadata
  initialLiquidity?: {
    yesAmount: bigint;
    noAmount: bigint;
  };
  // Required contract parameters
  collateralToken: Address;
  category: string;
  metadataURI: string;
  creatorStake: bigint;
}

/**
 * Market information structure - matches plan API
 */
export interface Market {
  id: number;
  creator: Address;
  amm: Address;
  collateralToken: Address;
  closeTime: number;
  category: string;
  metadataURI: string;
  status: MarketStatus;
  // Additional fields for compatibility
  marketId?: bigint;
  marketAddress?: Address;
  question?: string;
  description?: string;
  endTime?: bigint;
  yesTokenId?: bigint;
  noTokenId?: bigint;
  totalVolume?: bigint;
  totalLiquidity?: {
    yes: bigint;
    no: bigint;
  };
  createdAt?: bigint;
  marketType?: MarketType;
  outcomeCount?: bigint;
}

/**
 * Market information structure - legacy compatibility
 */
export interface MarketInfo {
  marketId: bigint;
  marketAddress: Address;
  question: string;
  description?: string;
  creator: Address;
  endTime: bigint;
  status: MarketStatus;
  resolutionData?: string;
  yesTokenId: bigint;
  noTokenId: bigint;
  totalVolume: bigint;
  totalLiquidity: {
    yes: bigint;
    no: bigint;
  };
  createdAt: bigint;
  marketType?: MarketType; // Market type from IMarket.getMarketType()
  outcomeCount?: bigint; // Number of outcomes (2 for binary, 3-8 for multi-choice)
  // Additional fields from MarketStruct
  collateralToken?: Address;
  category?: string;
  metadataURI?: string;
}

/**
 * Market prices for YES and NO outcomes
 */
export interface MarketPrices {
  yesPrice: bigint; // Price in wei (1e18 = 1.0)
  noPrice: bigint; // Price in wei (1e18 = 1.0)
  // Simplified price format (0-1 range)
  yes?: number; // Price as decimal (0-1)
  no?: number; // Price as decimal (0-1)
}

/**
 * Trade quote for slippage protection
 */
export interface TradeQuote {
  tokensOut: bigint;
  fee: bigint;
  priceImpact: number; // Percentage (e.g., 0.5 = 0.5%)
}

/**
 * Trade parameters
 */
export interface TradeParams {
  marketId: bigint;
  outcome: MarketOutcome | bigint; // For binary: 'YES' | 'NO', for multi-choice: outcomeId (bigint)
  amount: bigint; // Amount of outcome tokens to buy/sell
  minAmountOut?: bigint; // Minimum amount out (slippage protection)
  recipient?: Address; // Optional recipient address
}

/**
 * Trade result
 */
export interface TradeResult {
  success: boolean;
  transactionHash?: string;
  amountIn: bigint;
  amountOut: bigint;
  outcome: MarketOutcome;
  marketId: bigint;
}

/**
 * User position in a market
 */
export interface UserPosition {
  marketId: bigint;
  yesBalance: bigint;
  noBalance: bigint;
  totalValue: bigint; // Current value of position
}

/**
 * Resolution proposal from AI resolver
 */
export interface ResolutionProposal {
  marketId: bigint;
  outcome: MarketOutcome;
  evidenceHash: string;
  signature: string;
  timestamp: bigint;
}

/**
 * Fee tier information
 */
export interface FeeTier {
  tier: number;
  feeRate: bigint; // Fee rate in basis points (e.g., 200 = 2%)
  requiresHorizonTokens: bigint; // Minimum HORIZON tokens required
}

/**
 * Fee configuration for a market
 */
export interface FeeConfig {
  protocolBps: number; // Protocol share in basis points (e.g., 1000 = 10%)
  creatorBps: number; // Creator share in basis points (e.g., 9000 = 90%)
}

/**
 * SDK configuration
 */
export interface SDKConfig {
  chainId: number; // BNB Chain: 56 (mainnet), 97 (testnet)
  rpcUrl?: string;
  marketFactoryAddress: Address;
  horizonTokenAddress?: Address;
  outcomeTokenAddress?: Address;
  horizonPerksAddress?: Address;
  feeSplitterAddress?: Address;
  resolutionModuleAddress?: Address;
  aiOracleAdapterAddress?: Address;
  explorerUrl?: string;
}

/**
 * Error types
 */
export class SDKError extends Error {
  constructor(
    message: string,
    public code?: string,
    public data?: unknown
  ) {
    super(message);
    this.name = 'SDKError';
  }
}

export class ContractError extends SDKError {
  constructor(message: string, public contractAddress?: Address, data?: unknown) {
    super(message, 'CONTRACT_ERROR', data);
    this.name = 'ContractError';
  }
}

export class TradeError extends SDKError {
  constructor(message: string, public marketId?: bigint, data?: unknown) {
    super(message, 'TRADE_ERROR', data);
    this.name = 'TradeError';
  }
}

/**
 * Oracle request from API
 */
export interface OracleRequest {
  requestId: string;
  marketId: number;
  status: 'pending' | 'processing' | 'completed' | 'failed';
  progress?: number; // Optional progress percentage (0-100)
}

/**
 * Oracle result from API
 */
export interface OracleResult {
  requestId: string;
  marketId: number;
  outcomeId: number;
  confidence: number; // Percentage (0-100)
  reasoning: string;
  sources: string[];
  evidenceUrl: string;
  timestamp?: number;
}

