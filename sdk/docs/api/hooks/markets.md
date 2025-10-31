# Market Hooks

Hooks for discovering, fetching, and creating prediction markets.

## useMarkets

Fetch markets with optional filters.

### Import

```tsx
import { useMarkets, MarketStatus } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: markets, isLoading, error } = useMarkets({
  category: 'sports',
  status: MarketStatus.Active,
  limit: 10,
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `filters` | `UseMarketsFilters` | No | Filter options |

### UseMarketsFilters

```tsx
interface UseMarketsFilters {
  category?: string;      // Filter by category
  status?: MarketStatus;  // Filter by status
  creator?: string;       // Filter by creator address
  limit?: number;         // Maximum number of markets to return
  offset?: number;        // Pagination offset
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `Market[] \| undefined` | Array of markets |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |
| `refetch` | `() => void` | Function to refetch markets |

### Examples

#### Fetch All Active Markets

```tsx
const { data: markets } = useMarkets({
  status: MarketStatus.Active,
});
```

#### Fetch Markets by Category

```tsx
const { data: markets } = useMarkets({
  category: 'sports',
  limit: 20,
});
```

#### Fetch Markets by Creator

```tsx
const { data: markets } = useMarkets({
  creator: '0x...',
});
```

## useMarket

Fetch a single market by ID.

### Import

```tsx
import { useMarket } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: market, isLoading } = useMarket(marketId);
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number \| undefined` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `Market \| undefined` | Market data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### Examples

```tsx
function MarketDetail({ marketId }: { marketId: number }) {
  const { data: market, isLoading } = useMarket(marketId);

  if (isLoading) return <div>Loading...</div>;
  if (!market) return <div>Market not found</div>;

  return (
    <div>
      <h2>{market.category}</h2>
      <p>Creator: {market.creator}</p>
      <p>Status: {market.status}</p>
    </div>
  );
}
```

## useCreateMarket

Create a new prediction market.

### Import

```tsx
import { useCreateMarket } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: createMarket, isLoading } = useCreateMarket();

createMarket({
  collateralToken: '0x...', // USDC address
  category: 'sports',
  metadataURI: 'ipfs://...',
  endTime: BigInt(Math.floor(Date.now() / 1000) + 86400 * 30),
  creatorStake: parseUnits('1000', 18), // 1000 HORIZON tokens
});
```

### Parameters

```tsx
interface CreateMarketParams {
  collateralToken: Address;    // ERC20 token address
  category: string;             // Market category
  metadataURI: string;          // IPFS URI with market metadata
  endTime: bigint;              // Unix timestamp for market close
  creatorStake: bigint;         // Amount of HORIZON tokens to stake
  question?: string;             // Market question (optional)
  description?: string;          // Market description (optional)
  initialLiquidity?: {          // Optional initial liquidity
    yesAmount: bigint;
    noAmount: bigint;
  };
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `write` | `(params: CreateMarketParams) => void` | Function to create market |
| `isLoading` | `boolean` | Transaction pending state |
| `isSuccess` | `boolean` | Transaction success state |
| `hash` | `string \| undefined` | Transaction hash |
| `error` | `Error \| null` | Error object if transaction failed |

### Examples

```tsx
function CreateMarketForm() {
  const { write: createMarket, isLoading } = useCreateMarket();

  const handleSubmit = () => {
    createMarket({
      collateralToken: '0x...',
      category: 'sports',
      metadataURI: 'ipfs://Qm...',
      endTime: BigInt(Math.floor(Date.now() / 1000) + 86400 * 30),
      creatorStake: parseUnits('1000', 18),
    });
  };

  return (
    <button onClick={handleSubmit} disabled={isLoading}>
      {isLoading ? 'Creating...' : 'Create Market'}
    </button>
  );
}
```

## Market Types

### Market

```tsx
interface Market {
  id: number;
  creator: Address;
  amm: Address;
  collateralToken: Address;
  closeTime: number;
  category: string;
  metadataURI: string;
  status: MarketStatus;
  // Additional fields
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
```

### MarketStatus

```tsx
enum MarketStatus {
  Active = 0,
  Closed = 1,
  Resolved = 2,
  Invalid = 3,
}
```

