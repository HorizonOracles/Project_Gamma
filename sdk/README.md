# Project Gamma React SDK

A simple, React-focused SDK for integrating Project Gamma's decentralized prediction markets into React applications.

## Features

- ✅ **Market Discovery**: Search, filter, and query markets
- ✅ **Trading**: Buy/sell outcome tokens with slippage protection
- ✅ **Liquidity**: Add/remove liquidity, track LP positions
- ✅ **Resolution**: Propose, dispute, and finalize market outcomes
- ✅ **Oracle Integration**: Request AI resolution via public API
- ✅ **Wallet Integration**: Built on Wagmi for wallet connections
- ✅ **Real-Time Updates**: Automatic data refetching and caching
- ✅ **TypeScript**: Full TypeScript support with comprehensive types

## Installation

```bash
npm install @project-gamma/react-sdk
# or
yarn add @project-gamma/react-sdk
# or
pnpm add @project-gamma/react-sdk
```

## Peer Dependencies

The SDK requires the following peer dependencies:

```bash
npm install react react-dom wagmi viem @tanstack/react-query
```

## Quick Start

### 1. Setup Wagmi Provider

The SDK integrates with Wagmi for wallet management. Set up your Wagmi configuration:

```tsx
import { WagmiProvider, createConfig, http } from 'wagmi';
import { bsc, bscTestnet } from 'wagmi/chains';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const config = createConfig({
  chains: [bsc, bscTestnet],
  transports: {
    [bsc.id]: http(),
    [bscTestnet.id]: http(),
  },
});

const queryClient = new QueryClient();

function App() {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        {/* Your app */}
      </QueryClientProvider>
    </WagmiProvider>
  );
}
```

### 2. Setup SDK Provider

Wrap your app with the SDK provider:

```tsx
import { GammaProvider } from '@project-gamma/react-sdk';

function App() {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <GammaProvider
          chainId={56}
          oracleApiUrl="https://api.projectgamma.io"
          pinataJwt="your-pinata-jwt-token" // Optional: for IPFS storage
        >
          {/* Your app */}
        </GammaProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
```

**Note**: The `pinataJwt` prop is optional. If provided, it will be used for IPFS uploads when creating markets. You can also configure it via environment variables or pass it directly to hooks.

### 3. Use SDK Hooks

```tsx
import { useMarkets, useBuy, usePrices } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

function MarketList() {
  const { data: markets, isLoading } = useMarkets({
    status: MarketStatus.Active,
  });

  if (isLoading) return <div>Loading...</div>;

  return (
    <div>
      {markets?.map(market => (
        <MarketCard key={market.id} market={market} />
      ))}
    </div>
  );
}

function MarketCard({ market }: { market: Market }) {
  const { data: prices } = usePrices(market.id);
  const { write: buyYes, isLoading } = useBuy(market.id);

  const handleBuy = () => {
    buyYes({
      outcomeId: 0, // YES
      amount: parseUnits('100', 6), // 100 USDC
      slippage: 0.5, // 0.5%
    });
  };

  return (
    <div>
      <h3>{market.category}</h3>
      {prices && (
        <div>
          YES: {(Number(prices.yesPrice) / 1e18 * 100).toFixed(1)}¢
          NO: {(Number(prices.noPrice) / 1e18 * 100).toFixed(1)}¢
        </div>
      )}
      <button onClick={handleBuy} disabled={isLoading}>
        {isLoading ? 'Buying...' : 'Buy YES'}
      </button>
    </div>
  );
}
```

## API Reference

### React Hooks

#### Market Hooks

##### `useMarkets(filters?)`

Fetch markets with optional filters.

```tsx
import { useMarkets, MarketStatus } from '@project-gamma/react-sdk';

const { data: markets, isLoading } = useMarkets({
  category: 'sports',
  status: MarketStatus.Active,
  limit: 10,
});
```

##### `useMarket(marketId)`

Fetch single market information.

```tsx
import { useMarket } from '@project-gamma/react-sdk';

const { data: market, isLoading } = useMarket(1);
```

##### `useCreateMarket()`

Create a new prediction market.

```tsx
import { useCreateMarket } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: createMarket, isLoading } = useCreateMarket();

createMarket({
  collateralToken: '0x...', // USDC address
  category: 'sports',
  metadataURI: 'ipfs://...',
  closeTime: BigInt(Math.floor(Date.now() / 1000) + 86400 * 30),
  creatorStake: parseUnits('1000', 18), // 1000 HORIZON tokens
});
```

#### Trading Hooks

##### `useBuy(marketId)`

Buy outcome tokens.

```tsx
import { useBuy } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: buyYes, isLoading, hash } = useBuy(1);

buyYes({
  outcomeId: 0, // 0 = YES, 1 = NO
  amount: parseUnits('100', 6), // 100 USDC
  slippage: 0.5, // 0.5% slippage tolerance
});
```

##### `useSell(marketId)`

Sell outcome tokens.

```tsx
import { useSell } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: sellYes, isLoading } = useSell(1);

sellYes({
  outcomeId: 0,
  amount: parseUnits('50', 18), // 50 YES tokens
  slippage: 0.5,
});
```

##### `useQuote(marketId, params)`

Get price quote for a trade.

```tsx
import { useQuote } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { data: quote } = useQuote({
  marketId: 1,
  outcomeId: 0,
  amount: parseUnits('100', 6),
  isBuy: true,
});

// quote: { tokensOut, fee, priceImpact }
```

##### `usePrices(marketId)`

Get current market prices.

```tsx
import { usePrices } from '@project-gamma/react-sdk';

const { data: prices } = usePrices(1);

// prices: { yesPrice, noPrice }
```

#### Liquidity Hooks

##### `useAddLiquidity(marketId)`

Add liquidity to a market.

```tsx
import { useAddLiquidity } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: addLiquidity, isLoading } = useAddLiquidity(1);

addLiquidity({
  amount: parseUnits('1000', 6), // 1000 USDC
});
```

##### `useRemoveLiquidity(marketId)`

Remove liquidity from a market.

```tsx
import { useRemoveLiquidity } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: removeLiquidity, isLoading } = useRemoveLiquidity(1);

removeLiquidity({
  lpTokens: parseUnits('100', 18),
});
```

##### `useLPPosition(marketId)`

Get LP position for a market.

```tsx
import { useLPPosition } from '@project-gamma/react-sdk';

const { data: position } = useLPPosition(1);

// position: { lpTokens, value, yesReserve, noReserve }
```

#### Resolution Hooks

##### `useResolution(marketId)`

Get resolution state for a market.

```tsx
import { useResolution } from '@project-gamma/react-sdk';

const { data: resolution } = useResolution(1);
```

##### `useProposeResolution(marketId)`

Propose a resolution for a market.

```tsx
import { useProposeResolution } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: proposeResolution, isLoading } = useProposeResolution(1);

proposeResolution({
  outcomeId: 0, // YES
  evidenceHash: '0x...',
  signature: '0x...',
  bondAmount: parseUnits('1000', 18),
  evidenceURIs: ['ipfs://...'],
});
```

##### `useDispute(marketId)`

Dispute a proposed resolution.

```tsx
import { useDispute } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: dispute, isLoading } = useDispute(1);

dispute({
  bondAmount: parseUnits('1000', 18),
});
```

##### `useFinalize(marketId)`

Finalize a resolution after dispute window.

```tsx
import { useFinalize } from '@project-gamma/react-sdk';

const { write: finalize, isLoading } = useFinalize(1);

finalize();
```

#### Oracle Hooks

##### `useRequestResolution()`

Request AI resolution via API.

```tsx
import { useRequestResolution } from '@project-gamma/react-sdk';

const { mutate: requestResolution, data: request } = useRequestResolution();

requestResolution({
  marketId: 1,
  metadata: {
    question: 'Market question',
    description: 'Details',
  },
});
```

##### `useOracleStatus(requestId, options?)`

Poll oracle request status.

```tsx
import { useOracleStatus } from '@project-gamma/react-sdk';

const { data: status } = useOracleStatus(request?.requestId, {
  refetchInterval: 5000, // Poll every 5 seconds
});
```

##### `useOracleResult(requestId, options?)`

Get completed oracle result.

```tsx
import { useOracleResult } from '@project-gamma/react-sdk';

const { data: result } = useOracleResult(request?.requestId, {
  enabled: status?.status === 'completed',
});

// result: { outcomeId, confidence, reasoning, sources, evidenceUrl }
```

##### `useOracleHistory(marketId)`

Get oracle request history for a market.

```tsx
import { useOracleHistory } from '@project-gamma/react-sdk';

const { data: history } = useOracleHistory(1);
```

#### Token Hooks

##### `useBalance(tokenAddress, userAddress?)`

Get ERC20 token balance.

```tsx
import { useBalance } from '@project-gamma/react-sdk';

const { data: balance } = useBalance('0x...'); // USDC address
```

##### `useOutcomeBalance(marketId, outcomeId, userAddress?)`

Get outcome token balance.

```tsx
import { useOutcomeBalance } from '@project-gamma/react-sdk';

const { data: yesBalance } = useOutcomeBalance(1, 0); // YES tokens
const { data: noBalance } = useOutcomeBalance(1, 1); // NO tokens
```

##### `useApprove(tokenAddress)`

Approve token spending.

```tsx
import { useApprove } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const { write: approve, isLoading } = useApprove('0x...'); // USDC address

approve({
  spender: '0x...', // Market AMM address
  amount: parseUnits('10000', 6), // Approve 10000 USDC
});
```

##### `useRedeem(marketId)`

Redeem winning outcome tokens.

```tsx
import { useRedeem } from '@project-gamma/react-sdk';

const { write: redeem, isLoading } = useRedeem(1);

redeem();
```

## Configuration

### GammaProvider Props

```tsx
<GammaProvider
  chainId={56} // Required: BNB Chain Mainnet (56) or Testnet (97)
  oracleApiUrl="https://api.projectgamma.io" // Optional: Oracle API URL
  marketFactoryAddress="0x..." // Optional: Override default address
  horizonTokenAddress="0x..." // Optional: Override default address
  outcomeTokenAddress="0x..." // Optional: Override default address
  // ... other contract addresses
>
```

### Default Contract Addresses (BNB Chain Mainnet)

- **MarketFactory**: `0x22Cc806047BB825aa26b766Af737E92B1866E8A6`
- **HorizonToken**: `0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444`
- **OutcomeToken**: `0x17B322784265c105a94e4c3d00aF1E5f46a5F311`
- **HorizonPerks**: `0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C`
- **FeeSplitter**: `0x275017E98adF33051BbF477fe1DD197F681d4eF1`
- **ResolutionModule**: `0xF0CF4C741910cB48AC596F620a0AE892Cd247838`
- **AIOracleAdapter**: `0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93`

## Utility Functions

### EIP-712 Signing (AI Oracle Integration)

The SDK includes comprehensive EIP-712 utilities for creating cryptographically signed proposals that can be verified by the AIOracleAdapter contract.

#### Evidence Hash Computation

```tsx
import { computeEvidenceHash } from '@project-gamma/react-sdk';

// Compute evidence hash from URIs (IPFS, HTTP, etc.)
const evidenceURIs = [
  'ipfs://QmTest1234...',
  'https://example.com/evidence.json'
];
const evidenceHash = computeEvidenceHash(evidenceURIs);
// Result: 0x1234... (32-byte hash)
```

#### EIP-712 Domain and Digest

```tsx
import {
  createAIOracleDomain,
  computeDomainSeparator,
  buildProposedOutcome,
  computeProposalDigest,
} from '@project-gamma/react-sdk';

// 1. Create EIP-712 domain for AIOracleAdapter
const domain = createAIOracleDomain(
  56, // Chain ID (BSC mainnet)
  '0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93' // AIOracleAdapter address
);

// 2. Compute domain separator
const domainSeparator = computeDomainSeparator(domain);

// 3. Build proposal struct
const proposal = buildProposedOutcome({
  marketId: 123n,
  outcomeId: 0n, // YES
  closeTime: 1699999999n,
  evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
  notBefore: 1699000000n, // Optional
  deadline: 1700000000n, // Optional
});

// 4. Compute proposal digest (hash to sign)
const digest = computeProposalDigest(proposal, domainSeparator);
```

#### Signing Proposals

```tsx
import { signProposal, signProposalWithAccount } from '@project-gamma/react-sdk';
import { privateKeyToAccount } from 'viem/accounts';

// Option 1: Sign with private key directly
const privateKey = '0x1234...'; // AI signer private key
const signature = await signProposal(proposal, domainSeparator, privateKey);

// Option 2: Sign with viem account
const account = privateKeyToAccount(privateKey);
const signature = await signProposalWithAccount(proposal, domainSeparator, account);

// Signature is ready to submit to contract
console.log('Signature:', signature);
// Result: 0xabcd... (65-byte signature: r + s + v)
```

#### Verifying Signatures

```tsx
import { verifyProposalSignature, isValidProposalSignature } from '@project-gamma/react-sdk';

// Recover signer address from signature
const signerAddress = await verifyProposalSignature(
  proposal,
  domainSeparator,
  signature
);

// Check if signature matches expected signer
const isValid = await isValidProposalSignature(
  proposal,
  domainSeparator,
  signature,
  expectedSignerAddress
);

console.log('Signer:', signerAddress);
console.log('Valid:', isValid);
```

#### Working with Signature Components

```tsx
import {
  splitSignature,
  joinSignature,
  normalizeSignature,
  isValidSignature,
  formatSignature,
} from '@project-gamma/react-sdk';

// Split signature into r, s, v components
const { r, s, v } = splitSignature(signature);
console.log('r:', r); // 0x... (32 bytes)
console.log('s:', s); // 0x... (32 bytes)
console.log('v:', v); // 27 or 28

// Rejoin components into signature
const rejoined = joinSignature({ r, s, v });

// Normalize signature (convert v=0/1 to v=27/28)
const normalized = normalizeSignature(signature);

// Validate signature format
if (isValidSignature(signature)) {
  console.log('Valid 65-byte signature');
}

// Format for display
console.log(formatSignature(signature));
// Output:
// Signature {
//   r: 0x...
//   s: 0x...
//   v: 27
// }
```

#### Complete AI Oracle Signing Flow

```tsx
import {
  computeEvidenceHash,
  createAIOracleDomain,
  computeDomainSeparator,
  buildProposedOutcome,
  signProposal,
  verifyProposalSignature,
} from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

// 1. Prepare evidence
const evidenceURIs = [
  'ipfs://QmABC123...',
  'https://example.com/evidence.json',
];
const evidenceHash = computeEvidenceHash(evidenceURIs);

// 2. Set up EIP-712 domain
const chainId = 56; // BSC mainnet
const contractAddress = '0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93';
const domain = createAIOracleDomain(chainId, contractAddress);
const domainSeparator = computeDomainSeparator(domain);

// 3. Build proposal
const marketCloseTime = BigInt(Math.floor(Date.now() / 1000) + 86400 * 7);
const proposal = buildProposedOutcome({
  marketId: 123n,
  outcomeId: 0n, // YES
  closeTime: marketCloseTime,
  evidenceHash,
  notBefore: BigInt(Math.floor(Date.now() / 1000)),
  deadline: BigInt(Math.floor(Date.now() / 1000) + 86400),
});

// 4. Sign with AI signer private key
const aiSignerPrivateKey = process.env.AI_SIGNER_PRIVATE_KEY!;
const signature = await signProposal(
  proposal,
  domainSeparator,
  aiSignerPrivateKey
);

// 5. Verify signature (optional)
const recoveredSigner = await verifyProposalSignature(
  proposal,
  domainSeparator,
  signature
);
console.log('Proposal signed by:', recoveredSigner);

// 6. Submit to contract
const { useProposeResolution } = require('@project-gamma/react-sdk');
const { write: proposeResolution } = useProposeResolution(marketId);

proposeResolution({
  outcomeId: Number(proposal.outcomeId),
  evidenceHash: proposal.evidenceHash,
  signature,
  bondAmount: parseUnits('1000', 18),
  evidenceURIs,
});
```

### Token Amount Formatting

```tsx
import { formatTokenAmount, parseTokenAmount } from '@project-gamma/react-sdk';

// Format bigint to readable string
const formatted = formatTokenAmount(1000000000000000000n, 18, 4); // "1.0000"

// Parse string to bigint
const parsed = parseTokenAmount('1.5', 18); // 1500000000000000000n
```

### Price Calculations

```tsx
import { calculatePrice, calculateMarketPrices } from '@project-gamma/react-sdk';

// Calculate price for an outcome
const price = calculatePrice(yesLiquidity, noLiquidity);

// Calculate prices for both outcomes
const prices = calculateMarketPrices(yesLiquidity, noLiquidity);
```

### Address Utilities

```tsx
import { isValidAddress, shortenAddress } from '@project-gamma/react-sdk';

const isValid = isValidAddress('0x...');
const shortened = shortenAddress('0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb'); // "0x742d...bEb"
```

### Token ID Utilities

```tsx
import { getOutcomeTokenId, getMarketIdFromTokenId, getOutcomeFromTokenId } from '@project-gamma/react-sdk';

const tokenId = getOutcomeTokenId(1n, 0); // Market 1, YES outcome (0)
const marketId = getMarketIdFromTokenId(tokenId); // 1n
const outcome = getOutcomeFromTokenId(tokenId); // 0 (YES)
```

## Error Handling

The SDK uses custom error classes for better error handling:

```tsx
import { SDKError, ContractError, TradeError } from '@project-gamma/react-sdk';

try {
  buyYes({ outcomeId: 0, amount: parseUnits('100', 6) });
} catch (error) {
  if (error instanceof SDKError) {
    console.error('SDK Error:', error.message, error.code);
  } else if (error instanceof ContractError) {
    console.error('Contract Error:', error.message, error.contractAddress);
  } else if (error instanceof TradeError) {
    console.error('Trade Error:', error.message, error.marketId);
  }
}
```

## Complete Example

```tsx
import { WagmiProvider, createConfig, http } from 'wagmi';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { bsc } from 'wagmi/chains';
import { GammaProvider, useMarkets, useBuy, usePrices, MarketStatus } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

const config = createConfig({
  chains: [bsc],
  transports: {
    [bsc.id]: http(),
  },
});

const queryClient = new QueryClient();

function App() {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <GammaProvider chainId={56} oracleApiUrl="https://api.projectgamma.io">
          <MarketList />
        </GammaProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

function MarketList() {
  const { data: markets, isLoading } = useMarkets({
    status: MarketStatus.Active,
  });

  if (isLoading) return <div>Loading markets...</div>;

  return (
    <div>
      {markets?.map(market => (
        <MarketCard key={market.id} market={market} />
      ))}
    </div>
  );
}

function MarketCard({ market }: { market: Market }) {
  const { data: prices } = usePrices(market.id);
  const { write: buyYes, isLoading } = useBuy(market.id);

  const handleBuy = () => {
    buyYes({
      outcomeId: 0,
      amount: parseUnits('100', 6),
      slippage: 0.5,
    });
  };

  return (
    <div>
      <h3>{market.category}</h3>
      {prices && (
        <div>
          YES: {(Number(prices.yesPrice) / 1e18 * 100).toFixed(1)}¢
          NO: {(Number(prices.noPrice) / 1e18 * 100).toFixed(1)}¢
        </div>
      )}
      <button onClick={handleBuy} disabled={isLoading}>
        {isLoading ? 'Buying...' : 'Buy YES'}
      </button>
    </div>
  );
}
```

## TypeScript Support

The SDK is fully typed. All types are exported:

```tsx
import type {
  Market,
  MarketOutcome,
  MarketStatus,
  CreateMarketParams,
  TradeQuote,
  MarketPrices,
  ResolutionProposal,
  OracleRequest,
  OracleResult,
  SDKError,
  ContractError,
  TradeError,
} from '@project-gamma/react-sdk';
```

## Testing

The SDK includes comprehensive tests. Run tests with:

```bash
npm test
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](./CONTRIBUTING.md) for details.

## License

MIT

## Support

For support, please open an issue on GitHub or contact the Horizon Oracles team.

## Links

- **Website**: [horizonoracles.com](https://horizonoracles.com)
- **GitHub**: [github.com/HorizonOracles/Project_Gamma](https://github.com/HorizonOracles/Project_Gamma)
- **Contract Address**: `0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`
