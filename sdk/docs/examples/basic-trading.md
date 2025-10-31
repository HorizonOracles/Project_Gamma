# Basic Trading Example

A complete example of building a simple trading interface with the Project Gamma SDK.

## Complete Example

```tsx
import { WagmiProvider, createConfig, http } from 'wagmi';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { bsc } from 'wagmi/chains';
import { 
  GammaProvider, 
  useMarkets, 
  useBuy, 
  useSell, 
  usePrices, 
  useQuote,
  MarketStatus 
} from '@project-gamma/react-sdk';
import { parseUnits, formatUnits } from 'viem';
import { useState } from 'react';

// Setup Wagmi config
const wagmiConfig = createConfig({
  chains: [bsc],
  transports: {
    [bsc.id]: http(),
  },
});

const queryClient = new QueryClient();

function App() {
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <GammaProvider chainId={56} oracleApiUrl="https://api.projectgamma.io">
          <TradingApp />
        </GammaProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

// Market List Component
function MarketList() {
  const { data: markets, isLoading } = useMarkets({
    status: MarketStatus.Active,
    limit: 10,
  });

  if (isLoading) return <div>Loading markets...</div>;

  return (
    <div>
      <h2>Active Markets</h2>
      {markets?.map(market => (
        <MarketCard key={market.id} market={market} />
      ))}
    </div>
  );
}

// Market Card Component
function MarketCard({ market }: { market: Market }) {
  const [showTrading, setShowTrading] = useState(false);

  return (
    <div style={{ border: '1px solid #ccc', padding: '1rem', margin: '1rem' }}>
      <h3>{market.category}</h3>
      <p>Market ID: {market.id}</p>
      <PriceDisplay marketId={market.id} />
      <button onClick={() => setShowTrading(!showTrading)}>
        {showTrading ? 'Hide' : 'Show'} Trading
      </button>
      {showTrading && <TradingPanel marketId={market.id} />}
    </div>
  );
}

// Price Display Component
function PriceDisplay({ marketId }: { marketId: number }) {
  const { data: prices } = usePrices(marketId);

  if (!prices) return <div>Loading prices...</div>;

  const yesPercent = Number(prices.yesPrice) / 1e18 * 100;
  const noPercent = Number(prices.noPrice) / 1e18 * 100;

  return (
    <div>
      <div>YES: {yesPercent.toFixed(1)}¢</div>
      <div>NO: {noPercent.toFixed(1)}¢</div>
    </div>
  );
}

// Trading Panel Component
function TradingPanel({ marketId }: { marketId: number }) {
  const [amount, setAmount] = useState('100');
  const [outcome, setOutcome] = useState<0 | 1>(0);
  const [tradeType, setTradeType] = useState<'buy' | 'sell'>('buy');

  const { data: quote } = useQuote({
    marketId,
    outcomeId: outcome,
    amount: parseUnits(amount, tradeType === 'buy' ? 6 : 18),
    isBuy: tradeType === 'buy',
  });

  const { write: buy, isLoading: isBuying } = useBuy(marketId);
  const { write: sell, isLoading: isSelling } = useSell(marketId);

  const handleTrade = () => {
    if (tradeType === 'buy') {
      buy({
        outcomeId: outcome,
        amount: parseUnits(amount, 6),
        slippage: 0.5,
      });
    } else {
      sell({
        outcomeId: outcome,
        amount: parseUnits(amount, 18),
        slippage: 0.5,
      });
    }
  };

  return (
    <div style={{ marginTop: '1rem', padding: '1rem', background: '#f5f5f5' }}>
      <div>
        <label>
          <input
            type="radio"
            checked={tradeType === 'buy'}
            onChange={() => setTradeType('buy')}
          />
          Buy
        </label>
        <label>
          <input
            type="radio"
            checked={tradeType === 'sell'}
            onChange={() => setTradeType('sell')}
          />
          Sell
        </label>
      </div>

      <div>
        <label>
          <input
            type="radio"
            checked={outcome === 0}
            onChange={() => setOutcome(0)}
          />
          YES
        </label>
        <label>
          <input
            type="radio"
            checked={outcome === 1}
            onChange={() => setOutcome(1)}
          />
          NO
        </label>
      </div>

      <div>
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          placeholder={tradeType === 'buy' ? 'Amount (USDC)' : 'Amount (tokens)'}
        />
      </div>

      {quote && (
        <div>
          <p>
            {tradeType === 'buy' ? 'You\'ll receive' : 'You\'ll get'}:{' '}
            {formatUnits(
              tradeType === 'buy' ? quote.tokensOut : quote.collateralOut,
              tradeType === 'buy' ? 18 : 6
            )}
          </p>
          <p>Fee: {formatUnits(quote.fee, 6)} USDC</p>
          <p>Price impact: {quote.priceImpact.toFixed(2)}%</p>
        </div>
      )}

      <button
        onClick={handleTrade}
        disabled={isBuying || isSelling}
      >
        {isBuying || isSelling ? 'Trading...' : `${tradeType === 'buy' ? 'Buy' : 'Sell'} ${outcome === 0 ? 'YES' : 'NO'}`}
      </button>
    </div>
  );
}

// Main Trading App
function TradingApp() {
  return (
    <div>
      <h1>Project Gamma Trading</h1>
      <MarketList />
    </div>
  );
}

export default App;
```

## Key Features

1. **Market Discovery**: Lists all active markets
2. **Price Display**: Shows current YES/NO prices
3. **Trading Interface**: Buy/sell outcome tokens
4. **Quote Preview**: Shows quote before executing trade
5. **Slippage Protection**: Built-in slippage tolerance

## Step-by-Step Breakdown

### 1. Setup Providers

Wrap your app with Wagmi, React Query, and Gamma providers.

### 2. Fetch Markets

Use `useMarkets` to fetch active markets with filters.

### 3. Display Prices

Use `usePrices` to show current market prices.

### 4. Get Quotes

Use `useQuote` to preview trade outcomes before executing.

### 5. Execute Trades

Use `useBuy` and `useSell` to execute trades with slippage protection.

## Customization

- Add wallet connection UI
- Implement transaction status tracking
- Add error handling and user feedback
- Customize styling
- Add market filtering and search

