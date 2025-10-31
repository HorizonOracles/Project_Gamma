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

## useUploadMetadata

Upload market metadata to IPFS for use when creating markets.

### Import

```tsx
import { useUploadMetadata } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { mutate: uploadMetadata, isLoading, data } = useUploadMetadata();

uploadMetadata({
  question: 'Will BTC hit $100k?',
  description: 'Bitcoin price prediction market',
  category: 'crypto',
});
```

### Parameters

```tsx
interface UseUploadMetadataParams {
  question: string;           // Market question
  description?: string;        // Optional market description
  category: string;            // Market category
  provider?: IPFSProvider;     // Optional: 'pinata' | 'web3.storage' | 'public'
  pinataJwt?: string;          // Optional: Override Pinata JWT (overrides config)
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `mutate` | `(params: UseUploadMetadataParams) => void` | Function to upload metadata |
| `isLoading` | `boolean` | Upload pending state |
| `isSuccess` | `boolean` | Upload success state |
| `data` | `IPFSUploadResult \| undefined` | Upload result with IPFS hash and URL |
| `error` | `Error \| null` | Error object if upload failed |

### IPFSUploadResult

```tsx
interface IPFSUploadResult {
  hash: string;              // IPFS hash
  url: string;              // Full IPFS URL (ipfs://hash)
  gatewayUrl?: string;      // HTTP gateway URL for easy access
}
```

### Pinata JWT Priority

The Pinata JWT is resolved in the following priority order:

1. **Hook parameter** (`pinataJwt` in `UseUploadMetadataParams`) - Highest priority
2. **Provider config** (`pinataJwt` in `GammaProvider`)
3. **Environment variables** (`VITE_PINATA_JWT`, `NEXT_PUBLIC_PINATA_JWT`, or `PINATA_JWT`)

### Examples

#### Basic Upload

```tsx
function CreateMarketForm() {
  const { mutate: uploadMetadata, isLoading, data } = useUploadMetadata();

  const handleUpload = () => {
    uploadMetadata({
      question: 'Will BTC hit $100k by end of 2024?',
      description: 'Bitcoin price prediction market',
      category: 'crypto',
    });
  };

  return (
    <div>
      <button onClick={handleUpload} disabled={isLoading}>
        {isLoading ? 'Uploading...' : 'Upload Metadata'}
      </button>
      {data && (
        <div>
          <p>IPFS Hash: {data.hash}</p>
          <p>IPFS URL: {data.url}</p>
          <a href={data.gatewayUrl} target="_blank" rel="noopener noreferrer">
            View on IPFS Gateway
          </a>
        </div>
      )}
    </div>
  );
}
```

#### Upload with Custom JWT

```tsx
// Override provider config for this specific upload
uploadMetadata({
  question: 'Will ETH hit $5000?',
  category: 'crypto',
  pinataJwt: 'custom-jwt-token',
});
```

#### Complete Market Creation Flow

```tsx
import { useUploadMetadata, useCreateMarket } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
import { useEffect } from 'react';

function CreateMarketFlow() {
  const { mutate: uploadMetadata, data: ipfsData, isLoading: isUploading } = useUploadMetadata();
  const { write: createMarket, isLoading: isCreating } = useCreateMarket();

  const handleCreateMarket = () => {
    // Step 1: Upload metadata to IPFS
    uploadMetadata({
      question: 'Will BTC hit $100k?',
      description: 'Bitcoin price prediction',
      category: 'crypto',
    });
  };

  // Step 2: When upload succeeds, create market with IPFS URI
  useEffect(() => {
    if (ipfsData) {
      createMarket({
        collateralToken: '0x...', // USDC address
        category: 'crypto',
        metadataURI: ipfsData.url, // Use IPFS URL from upload
        endTime: BigInt(Math.floor(Date.now() / 1000) + 86400 * 30),
        creatorStake: parseUnits('1000', 18),
      });
    }
  }, [ipfsData, createMarket]);

  return (
    <button 
      onClick={handleCreateMarket} 
      disabled={isUploading || isCreating}
    >
      {isUploading ? 'Uploading...' : isCreating ? 'Creating...' : 'Create Market'}
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

