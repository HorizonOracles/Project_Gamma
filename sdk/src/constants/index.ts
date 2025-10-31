/**
 * SDK constants and default configurations
 */

import { Address } from 'viem';
import { SDKConfig } from '../types';

/**
 * BNB Chain network IDs
 */
export const BNB_CHAIN = {
  MAINNET: 56,
  TESTNET: 97,
} as const;

/**
 * Default contract addresses (BNB Chain Mainnet)
 * These are the actual deployed contract addresses
 */
export const DEFAULT_CONTRACTS = {
  [BNB_CHAIN.MAINNET]: {
    marketFactory: '0x22Cc806047BB825aa26b766Af737E92B1866E8A6' as Address,
    horizonToken: '0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444' as Address,
    outcomeToken: '0x17B322784265c105a94e4c3d00aF1E5f46a5F311' as Address,
    horizonPerks: '0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C' as Address,
    feeSplitter: '0x275017E98adF33051BbF477fe1DD197F681d4eF1' as Address,
    resolutionModule: '0xF0CF4C741910cB48AC596F620a0AE892Cd247838' as Address,
    aiOracleAdapter: '0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93' as Address,
    // USDC on BNB Chain Mainnet (commonly used as collateral token)
    usdc: '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d' as Address,
  },
  [BNB_CHAIN.TESTNET]: {
    // Placeholder addresses - update with actual testnet contract addresses once deployed
    // Contracts must be deployed to BNB Chain Testnet before use
    marketFactory: '0x0000000000000000000000000000000000000000' as Address,
    horizonToken: '0x0000000000000000000000000000000000000000' as Address,
    outcomeToken: '0x0000000000000000000000000000000000000000' as Address,
    horizonPerks: '0x0000000000000000000000000000000000000000' as Address,
    feeSplitter: '0x0000000000000000000000000000000000000000' as Address,
    resolutionModule: '0x0000000000000000000000000000000000000000' as Address,
    aiOracleAdapter: '0x0000000000000000000000000000000000000000' as Address,
    // USDC on BNB Chain Testnet
    usdc: '0x64544969ed7EBf5f083679233325356EbE738930' as Address,
  },
} as const;

/**
 * Default SDK configuration for mainnet
 * Automatically selects mainnet or testnet based on chainId
 */
export const DEFAULT_CONFIG: SDKConfig = {
  chainId: BNB_CHAIN.MAINNET,
  rpcUrl: 'https://bsc-dataseed.binance.org/',
  marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].marketFactory,
  horizonTokenAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].horizonToken,
  outcomeTokenAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].outcomeToken,
  horizonPerksAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].horizonPerks,
  feeSplitterAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].feeSplitter,
  resolutionModuleAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].resolutionModule,
  aiOracleAdapterAddress: DEFAULT_CONTRACTS[BNB_CHAIN.MAINNET].aiOracleAdapter,
  explorerUrl: 'https://bscscan.com',
};

/**
 * Default testnet configuration
 */
export const DEFAULT_TESTNET_CONFIG: SDKConfig = {
  chainId: BNB_CHAIN.TESTNET,
  rpcUrl: 'https://data-seed-prebsc-1-s1.binance.org:8545/',
  marketFactoryAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].marketFactory,
  horizonTokenAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].horizonToken,
  outcomeTokenAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].outcomeToken,
  horizonPerksAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].horizonPerks,
  feeSplitterAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].feeSplitter,
  resolutionModuleAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].resolutionModule,
  aiOracleAdapterAddress: DEFAULT_CONTRACTS[BNB_CHAIN.TESTNET].aiOracleAdapter,
  explorerUrl: 'https://testnet.bscscan.com',
};

/**
 * Contract ABIs will be imported here
 * For now, we'll define minimal interfaces
 */
export const MARKET_FACTORY_ABI = [
  // Market Creation
  {
    type: 'function',
    name: 'createMarket',
    inputs: [
      {
        name: 'params',
        type: 'tuple',
        components: [
          { name: 'collateralToken', type: 'address' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'category', type: 'string' },
          { name: 'metadataURI', type: 'string' },
          { name: 'creatorStake', type: 'uint256' },
        ],
      },
    ],
    outputs: [{ name: 'marketId', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  // Market Queries
  {
    type: 'function',
    name: 'getMarket',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'id', type: 'uint256' },
          { name: 'creator', type: 'address' },
          { name: 'amm', type: 'address' },
          { name: 'collateralToken', type: 'address' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'category', type: 'string' },
          { name: 'metadataURI', type: 'string' },
          { name: 'creatorStake', type: 'uint256' },
          { name: 'stakeRefunded', type: 'bool' },
          { name: 'status', type: 'uint8' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getAllMarketIds',
    inputs: [],
    outputs: [{ name: '', type: 'uint256[]' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getMarketCount',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getMarkets',
    inputs: [
      { name: 'offset', type: 'uint256' },
      { name: 'limit', type: 'uint256' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple[]',
        components: [
          { name: 'id', type: 'uint256' },
          { name: 'creator', type: 'address' },
          { name: 'amm', type: 'address' },
          { name: 'collateralToken', type: 'address' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'category', type: 'string' },
          { name: 'metadataURI', type: 'string' },
          { name: 'creatorStake', type: 'uint256' },
          { name: 'stakeRefunded', type: 'bool' },
          { name: 'status', type: 'uint8' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getActiveMarkets',
    inputs: [
      { name: 'offset', type: 'uint256' },
      { name: 'limit', type: 'uint256' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple[]',
        components: [
          { name: 'id', type: 'uint256' },
          { name: 'creator', type: 'address' },
          { name: 'amm', type: 'address' },
          { name: 'collateralToken', type: 'address' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'category', type: 'string' },
          { name: 'metadataURI', type: 'string' },
          { name: 'creatorStake', type: 'uint256' },
          { name: 'stakeRefunded', type: 'bool' },
          { name: 'status', type: 'uint8' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getMarketIdsByCategory',
    inputs: [{ name: 'category', type: 'string' }],
    outputs: [{ name: '', type: 'uint256[]' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getMarketIdsByCreator',
    inputs: [{ name: 'creator', type: 'address' }],
    outputs: [{ name: '', type: 'uint256[]' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'marketExists',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'nextMarketId',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'minCreatorStake',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  // Stake Management
  {
    type: 'function',
    name: 'refundCreatorStake',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  // Status Management
  {
    type: 'function',
    name: 'updateMarketStatus',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  // Events
  {
    type: 'event',
    name: 'MarketCreated',
    inputs: [
      { name: 'marketId', type: 'uint256', indexed: true },
      { name: 'creator', type: 'address', indexed: true },
      { name: 'ammAddress', type: 'address', indexed: true },
      { name: 'collateralToken', type: 'address', indexed: false },
      { name: 'closeTime', type: 'uint256', indexed: false },
      { name: 'category', type: 'string', indexed: false },
      { name: 'metadataURI', type: 'string', indexed: false },
      { name: 'creatorStake', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'CreatorStakeRefunded',
    inputs: [
      { name: 'marketId', type: 'uint256', indexed: true },
      { name: 'creator', type: 'address', indexed: true },
      { name: 'amount', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'MarketStatusUpdated',
    inputs: [
      { name: 'marketId', type: 'uint256', indexed: true },
      { name: 'oldStatus', type: 'uint8', indexed: false },
      { name: 'newStatus', type: 'uint8', indexed: false },
    ],
  },
] as const;

export const MARKET_AMM_ABI = [
  // Trading Functions
  {
    type: 'function',
    name: 'buyYes',
    inputs: [
      { name: 'collateralIn', type: 'uint256' },
      { name: 'minTokensOut', type: 'uint256' },
    ],
    outputs: [{ name: 'tokensOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'buyNo',
    inputs: [
      { name: 'collateralIn', type: 'uint256' },
      { name: 'minTokensOut', type: 'uint256' },
    ],
    outputs: [{ name: 'tokensOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'sellYes',
    inputs: [
      { name: 'tokensIn', type: 'uint256' },
      { name: 'minCollateralOut', type: 'uint256' },
    ],
    outputs: [{ name: 'collateralOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'sellNo',
    inputs: [
      { name: 'tokensIn', type: 'uint256' },
      { name: 'minCollateralOut', type: 'uint256' },
    ],
    outputs: [{ name: 'collateralOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  // Liquidity Functions
  {
    type: 'function',
    name: 'addLiquidity',
    inputs: [{ name: 'amount', type: 'uint256' }],
    outputs: [{ name: 'lpTokens', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'removeLiquidity',
    inputs: [{ name: 'lpTokens', type: 'uint256' }],
    outputs: [{ name: 'collateralOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  // Price Functions
  {
    type: 'function',
    name: 'getYesPrice',
    inputs: [],
    outputs: [{ name: 'price', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getNoPrice',
    inputs: [],
    outputs: [{ name: 'price', type: 'uint256' }],
    stateMutability: 'view',
  },
  // Quote Functions
  {
    type: 'function',
    name: 'getQuoteBuyYes',
    inputs: [
      { name: 'collateralIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'tokensOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getQuoteBuyNo',
    inputs: [
      { name: 'collateralIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'tokensOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getQuoteSellYes',
    inputs: [
      { name: 'tokensIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'collateralOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getQuoteSellNo',
    inputs: [
      { name: 'tokensIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'collateralOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  // View Functions
  {
    type: 'function',
    name: 'reserveYes',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'reserveNo',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'totalCollateral',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'collateralToken',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'balanceOf',
    inputs: [{ name: 'account', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  // Resolution Functions
  {
    type: 'function',
    name: 'fundRedemptions',
    inputs: [],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  // Events
  {
    type: 'event',
    name: 'Trade',
    inputs: [
      { name: 'trader', type: 'address', indexed: true },
      { name: 'buyYes', type: 'bool', indexed: true },
      { name: 'collateralIn', type: 'uint256', indexed: false },
      { name: 'tokensOut', type: 'uint256', indexed: false },
      { name: 'fee', type: 'uint256', indexed: false },
      { name: 'price', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'LiquidityAdded',
    inputs: [
      { name: 'provider', type: 'address', indexed: true },
      { name: 'collateralAmount', type: 'uint256', indexed: false },
      { name: 'lpTokens', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'LiquidityRemoved',
    inputs: [
      { name: 'provider', type: 'address', indexed: true },
      { name: 'lpTokens', type: 'uint256', indexed: false },
      { name: 'collateralAmount', type: 'uint256', indexed: false },
    ],
  },
] as const;

export const ERC20_ABI = [
  {
    type: 'function',
    name: 'balanceOf',
    inputs: [{ name: 'account', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
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
  {
    type: 'function',
    name: 'allowance',
    inputs: [
      { name: 'owner', type: 'address' },
      { name: 'spender', type: 'address' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'transfer',
    inputs: [
      { name: 'to', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'transferFrom',
    inputs: [
      { name: 'from', type: 'address' },
      { name: 'to', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
] as const;

export const AI_ORACLE_ADAPTER_ABI = [
  {
    type: 'function',
    name: 'proposeAI',
    inputs: [
      {
        name: 'proposal',
        type: 'tuple',
        components: [
          { name: 'marketId', type: 'uint256' },
          { name: 'outcomeId', type: 'uint256' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'evidenceHash', type: 'bytes32' },
          { name: 'notBefore', type: 'uint256' },
          { name: 'deadline', type: 'uint256' },
        ],
      },
      { name: 'signature', type: 'bytes' },
      { name: 'bondAmount', type: 'uint256' },
      { name: 'evidenceURIs', type: 'string[]' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'getProposalHash',
    inputs: [
      {
        name: 'proposal',
        type: 'tuple',
        components: [
          { name: 'marketId', type: 'uint256' },
          { name: 'outcomeId', type: 'uint256' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'evidenceHash', type: 'bytes32' },
          { name: 'notBefore', type: 'uint256' },
          { name: 'deadline', type: 'uint256' },
        ],
      },
    ],
    outputs: [{ name: '', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'isSignatureUsed',
    inputs: [{ name: 'signature', type: 'bytes' }],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'hashEvidence',
    inputs: [{ name: 'evidenceURIs', type: 'string[]' }],
    outputs: [{ name: '', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    name: 'allowedSigners',
    inputs: [{ name: '', type: 'address' }],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'DOMAIN_SEPARATOR',
    inputs: [],
    outputs: [{ name: '', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'event',
    name: 'AIProposalSubmitted',
    inputs: [
      { name: 'marketId', type: 'uint256', indexed: true },
      { name: 'outcomeId', type: 'uint256', indexed: true },
      { name: 'proposer', type: 'address', indexed: true },
      { name: 'aiSigner', type: 'address', indexed: false },
      { name: 'bondAmount', type: 'uint256', indexed: false },
      { name: 'signatureHash', type: 'bytes32', indexed: false },
    ],
  },
] as const;

export const OUTCOME_TOKEN_ABI = [
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
  {
    type: 'function',
    name: 'balanceOfBatch',
    inputs: [
      { name: 'accounts', type: 'address[]' },
      { name: 'ids', type: 'uint256[]' },
    ],
    outputs: [{ name: '', type: 'uint256[]' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'safeTransferFrom',
    inputs: [
      { name: 'from', type: 'address' },
      { name: 'to', type: 'address' },
      { name: 'id', type: 'uint256' },
      { name: 'amount', type: 'uint256' },
      { name: 'data', type: 'bytes' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'isResolved',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'winningOutcome',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'marketCollateral',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'redeem',
    inputs: [{ name: 'marketId', type: 'uint256' }],
    outputs: [{ name: 'payout', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'redeemAmount',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: 'payout', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'encodeTokenId',
    inputs: [
      { name: 'marketId', type: 'uint256' },
      { name: 'outcomeId', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'pure',
  },
] as const;

/**
 * IMarket interface ABI - Common interface for all market types
 * This allows us to interact with any market type (MarketAMM, LimitOrderMarket, etc.)
 */
export const I_MARKET_ABI = [
  {
    type: 'function',
    name: 'getMarketType',
    inputs: [],
    outputs: [{ name: '', type: 'uint8' }], // Returns MarketType enum
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getMarketInfo',
    inputs: [],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'marketId', type: 'uint256' },
          { name: 'marketType', type: 'uint8' },
          { name: 'collateralToken', type: 'address' },
          { name: 'closeTime', type: 'uint256' },
          { name: 'outcomeCount', type: 'uint256' },
          { name: 'isResolved', type: 'bool' },
          { name: 'isPaused', type: 'bool' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'marketId',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'collateralToken',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'closeTime',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getOutcomeCount',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'buy',
    inputs: [
      { name: 'outcomeId', type: 'uint256' },
      { name: 'collateralIn', type: 'uint256' },
      { name: 'minTokensOut', type: 'uint256' },
    ],
    outputs: [{ name: 'tokensOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'sell',
    inputs: [
      { name: 'outcomeId', type: 'uint256' },
      { name: 'tokensIn', type: 'uint256' },
      { name: 'minCollateralOut', type: 'uint256' },
    ],
    outputs: [{ name: 'collateralOut', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'getPrice',
    inputs: [{ name: 'outcomeId', type: 'uint256' }],
    outputs: [{ name: 'price', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getQuoteBuy',
    inputs: [
      { name: 'outcomeId', type: 'uint256' },
      { name: 'collateralIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'tokensOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'getQuoteSell',
    inputs: [
      { name: 'outcomeId', type: 'uint256' },
      { name: 'tokensIn', type: 'uint256' },
      { name: 'user', type: 'address' },
    ],
    outputs: [
      { name: 'collateralOut', type: 'uint256' },
      { name: 'fee', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'pause',
    inputs: [],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'unpause',
    inputs: [],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'fundRedemptions',
    inputs: [],
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const;

