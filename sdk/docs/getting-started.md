# Getting Started

This guide will help you get started with the Project Gamma React SDK in just a few minutes.

## Prerequisites

- Node.js 18+ 
- React 18+
- A React project (Create React App, Vite, Next.js, etc.)

## Installation

Install the SDK and its peer dependencies:

```bash
npm install @project-gamma/react-sdk wagmi viem @tanstack/react-query
# or
yarn add @project-gamma/react-sdk wagmi viem @tanstack/react-query
# or
pnpm add @project-gamma/react-sdk wagmi viem @tanstack/react-query
```

## Setup

### 1. Configure Wagmi

First, set up Wagmi with your desired chains:

```tsx
// config/wagmi.ts
import { createConfig, http } from 'wagmi';
import { bsc, bscTestnet } from 'wagmi/chains';
import { injected, metaMask, walletConnect } from 'wagmi/connectors';

export const wagmiConfig = createConfig({
  chains: [bsc, bscTestnet],
  connectors: [
    injected(),
    metaMask(),
    walletConnect({ projectId: 'your-project-id' }),
  ],
  transports: {
    [bsc.id]: http(),
    [bscTestnet.id]: http(),
  },
});
```

### 2. Setup Providers

Wrap your app with the required providers:

```tsx
// App.tsx
import { WagmiProvider } from 'wagmi';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { GammaProvider } from '@project-gamma/react-sdk';
import { wagmiConfig } from './config/wagmi';

const queryClient = new QueryClient();

function App() {
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <GammaProvider
          chainId={56} // BNB Chain Mainnet
          oracleApiUrl="https://api.projectgamma.io"
          pinataJwt="your-pinata-jwt-token" // Optional: for IPFS storage when creating markets
        >
          <YourApp />
        </GammaProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
```

**Note**: The `pinataJwt` prop is optional. If provided, it will be used for IPFS uploads when creating markets. You can also configure it via environment variables (`VITE_PINATA_JWT`, `NEXT_PUBLIC_PINATA_JWT`, or `PINATA_JWT`) or pass it directly to the `useUploadMetadata` hook.

### 3. Use SDK Hooks

Now you can use SDK hooks in your components:

```tsx
// components/MarketList.tsx
import { useMarkets, MarketStatus } from '@project-gamma/react-sdk';

export function MarketList() {
  const { data: markets, isLoading } = useMarkets({
    status: MarketStatus.Active,
  });

  if (isLoading) return <div>Loading markets...</div>;

  return (
    <div>
      {markets?.map(market => (
        <div key={market.id}>
          <h3>{market.category}</h3>
          <p>Market ID: {market.id}</p>
        </div>
      ))}
    </div>
  );
}
```

## Next Steps

- Read the [API Reference](./api/) for detailed hook documentation
- Check out [Examples](./examples/) for complete code samples
- Learn about [Error Handling](../README.md#error-handling) in the main README

