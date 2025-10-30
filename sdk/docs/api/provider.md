# GammaProvider

The `GammaProvider` component provides SDK configuration to all hooks in your application.

## Import

```tsx
import { GammaProvider } from '@project-gamma/react-sdk';
```

## Usage

```tsx
<GammaProvider
  chainId={56}
  oracleApiUrl="https://api.projectgamma.io"
>
  <YourApp />
</GammaProvider>
```

## Props

### Required Props

| Prop | Type | Description |
|------|------|-------------|
| `chainId` | `number` | Chain ID (56 for BNB Chain Mainnet, 97 for Testnet) |
| `children` | `ReactNode` | Your application components |

### Optional Props

| Prop | Type | Description |
|------|------|-------------|
| `oracleApiUrl` | `string` | Oracle API endpoint URL |
| `marketFactoryAddress` | `Address` | Override default MarketFactory address |
| `horizonTokenAddress` | `Address` | Override default HorizonToken address |
| `outcomeTokenAddress` | `Address` | Override default OutcomeToken address |
| `horizonPerksAddress` | `Address` | Override default HorizonPerks address |
| `feeSplitterAddress` | `Address` | Override default FeeSplitter address |
| `resolutionModuleAddress` | `Address` | Override default ResolutionModule address |
| `aiOracleAdapterAddress` | `Address` | Override default AIOracleAdapter address |

## Default Contract Addresses

### BNB Chain Mainnet (Chain ID: 56)

- **MarketFactory**: `0x22Cc806047BB825aa26b766Af737E92B1866E8A6`
- **HorizonToken**: `0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444`
- **OutcomeToken**: `0x17B322784265c105a94e4c3d00aF1E5f46a5F311`
- **HorizonPerks**: `0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C`
- **FeeSplitter**: `0x275017E98adF33051BbF477fe1DD197F681d4eF1`
- **ResolutionModule**: `0xF0CF4C741910cB48AC596F620a0AE892Cd247838`
- **AIOracleAdapter**: `0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93`

### BNB Chain Testnet (Chain ID: 97)

Testnet contracts must be deployed before use. Addresses will be updated once deployed.

## Environment Variables

You can also configure the SDK using environment variables:

```bash
CHAIN_ID=56
ORACLE_API_URL=https://api.projectgamma.io
MARKET_FACTORY_ADDRESS=0x...
# ... other addresses
```

## Accessing Config

Use the `useGammaConfig` hook to access the current configuration:

```tsx
import { useGammaConfig } from '@project-gamma/react-sdk';

function MyComponent() {
  const config = useGammaConfig();
  console.log(config.chainId); // 56
  console.log(config.oracleApiUrl); // https://api.projectgamma.io
}
```

## Examples

### Basic Setup

```tsx
<WagmiProvider config={wagmiConfig}>
  <QueryClientProvider client={queryClient}>
    <GammaProvider chainId={56}>
      <App />
    </GammaProvider>
  </QueryClientProvider>
</WagmiProvider>
```

### With Custom Contract Addresses

```tsx
<GammaProvider
  chainId={56}
  oracleApiUrl="https://api.projectgamma.io"
  marketFactoryAddress="0x..."
  horizonTokenAddress="0x..."
>
  <App />
</GammaProvider>
```

### Testnet Configuration

```tsx
<GammaProvider
  chainId={97}
  oracleApiUrl="https://testnet-api.projectgamma.io"
>
  <App />
</GammaProvider>
```

