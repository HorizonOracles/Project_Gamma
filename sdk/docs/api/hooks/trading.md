# Trading Hooks

Hooks for buying and selling outcome tokens, getting quotes, and fetching prices.

## useBuy

Buy outcome tokens for a market.

### Import

```tsx
import { useBuy } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: buyYes, isLoading, hash } = useBuy(marketId);

buyYes({
  outcomeId: 0, // 0 = YES, 1 = NO
  amount: parseUnits('100', 6), // 100 USDC
  slippage: 0.5, // 0.5% slippage tolerance
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### BuyParams

```tsx
interface BuyParams {
  outcomeId: number;    // 0 for YES, 1 for NO
  amount: bigint;       // Amount of collateral to spend
  slippage?: number;    // Slippage tolerance percentage (default: 0.5)
  recipient?: string;   // Optional recipient address
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `write` | `(params: BuyParams) => void` | Function to execute buy |
| `isLoading` | `boolean` | Transaction pending state |
| `isSuccess` | `boolean` | Transaction success state |
| `hash` | `string \| undefined` | Transaction hash |
| `error` | `Error \| null` | Error object if transaction failed |

### Examples

```tsx
function BuyPanel({ marketId }: { marketId: number }) {
  const [amount, setAmount] = useState('100');
  const { write: buyYes, isLoading } = useBuy(marketId);

  const handleBuy = () => {
    buyYes({
      outcomeId: 0,
      amount: parseUnits(amount, 6),
      slippage: 0.5,
    });
  };

  return (
    <div>
      <input
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
      />
      <button onClick={handleBuy} disabled={isLoading}>
        {isLoading ? 'Buying...' : 'Buy YES'}
      </button>
    </div>
  );
}
```

## useSell

Sell outcome tokens for a market.

### Import

```tsx
import { useSell } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: sellYes, isLoading } = useSell(marketId);

sellYes({
  outcomeId: 0,
  amount: parseUnits('50', 18), // 50 YES tokens
  slippage: 0.5,
});
```

### Parameters

Same as `useBuy`.

### SellParams

```tsx
interface SellParams {
  outcomeId: number;    // 0 for YES, 1 for NO
  amount: bigint;       // Amount of outcome tokens to sell
  slippage?: number;    // Slippage tolerance percentage (default: 0.5)
  recipient?: string;   // Optional recipient address
}
```

### Returns

Same as `useBuy`.

### Examples

```tsx
function SellPanel({ marketId }: { marketId: number }) {
  const { write: sellYes, isLoading } = useSell(marketId);

  const handleSell = () => {
    sellYes({
      outcomeId: 0,
      amount: parseUnits('50', 18),
      slippage: 0.5,
    });
  };

  return (
    <button onClick={handleSell} disabled={isLoading}>
      {isLoading ? 'Selling...' : 'Sell YES'}
    </button>
  );
}
```

## useQuote

Get a price quote for a trade.

### Import

```tsx
import { useQuote } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { data: quote } = useQuote({
  marketId: 1,
  outcomeId: 0,
  amount: parseUnits('100', 6),
  isBuy: true,
});
```

### Parameters

```tsx
interface QuoteParams {
  marketId: number;
  outcomeId: number;  // 0 for YES, 1 for NO
  amount: bigint;
  isBuy: boolean;     // true for buy, false for sell
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `TradeQuote \| undefined` | Quote data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### TradeQuote

```tsx
interface TradeQuote {
  tokensOut: bigint;      // Tokens you'll receive
  fee: bigint;            // Trading fee
  priceImpact: number;    // Price impact percentage
}
```

### Examples

```tsx
function TradeQuote({ marketId }: { marketId: number }) {
  const [amount, setAmount] = useState('100');
  const { data: quote } = useQuote({
    marketId,
    outcomeId: 0,
    amount: parseUnits(amount, 6),
    isBuy: true,
  });

  return (
    <div>
      <input
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
      />
      {quote && (
        <div>
          <p>You'll receive: {formatUnits(quote.tokensOut, 18)} YES</p>
          <p>Fee: {formatUnits(quote.fee, 6)} USDC</p>
          <p>Price impact: {quote.priceImpact.toFixed(2)}%</p>
        </div>
      )}
    </div>
  );
}
```

## usePrices

Get current market prices for YES and NO outcomes.

### Import

```tsx
import { usePrices } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: prices } = usePrices(marketId);
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `MarketPrices \| undefined` | Price data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### MarketPrices

```tsx
interface MarketPrices {
  yesPrice: bigint;  // YES price in wei (1e18 = 1.0)
  noPrice: bigint;   // NO price in wei (1e18 = 1.0)
  yes?: number;      // YES price as decimal (0-1)
  no?: number;       // NO price as decimal (0-1)
}
```

### Examples

```tsx
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
```

## Complete Trading Example

```tsx
import { useBuy, useSell, useQuote, usePrices } from '@project-gamma/react-sdk';
import { parseUnits, formatUnits } from 'viem';

function TradingPanel({ marketId }: { marketId: number }) {
  const [amount, setAmount] = useState('100');
  const [outcome, setOutcome] = useState<0 | 1>(0);

  const { data: prices } = usePrices(marketId);
  const { data: quote } = useQuote({
    marketId,
    outcomeId: outcome,
    amount: parseUnits(amount, 6),
    isBuy: true,
  });

  const { write: buy, isLoading: isBuying } = useBuy(marketId);
  const { write: sell, isLoading: isSelling } = useSell(marketId);

  const handleBuy = () => {
    buy({
      outcomeId: outcome,
      amount: parseUnits(amount, 6),
      slippage: 0.5,
    });
  };

  const handleSell = () => {
    sell({
      outcomeId: outcome,
      amount: parseUnits(amount, 18),
      slippage: 0.5,
    });
  };

  return (
    <div>
      <select value={outcome} onChange={(e) => setOutcome(Number(e.target.value) as 0 | 1)}>
        <option value={0}>YES</option>
        <option value={1}>NO</option>
      </select>
      <input
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
      />
      {quote && (
        <div>
          <p>You'll receive: {formatUnits(quote.tokensOut, 18)} tokens</p>
          <p>Price impact: {quote.priceImpact.toFixed(2)}%</p>
        </div>
      )}
      <button onClick={handleBuy} disabled={isBuying}>
        Buy
      </button>
      <button onClick={handleSell} disabled={isSelling}>
        Sell
      </button>
    </div>
  );
}
```

