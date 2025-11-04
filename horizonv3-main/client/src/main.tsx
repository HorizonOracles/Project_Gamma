import { createRoot } from "react-dom/client";
import '@rainbow-me/rainbowkit/styles.css';
import { RainbowKitProvider } from '@rainbow-me/rainbowkit';
import { WagmiProvider } from 'wagmi';
import { QueryClientProvider } from '@tanstack/react-query';
import { GammaProvider, BNB_CHAIN, DEFAULT_CONTRACTS } from '@project-gamma/sdk';
import { queryClient } from './lib/queryClient';
import { config } from './lib/wagmi-config';
import App from "./App";
import "./index.css";

// Use BNB Chain Mainnet contracts
const CHAIN_ID = BNB_CHAIN.MAINNET;
const contracts = DEFAULT_CONTRACTS[CHAIN_ID];

createRoot(document.getElementById("root")!).render(
  <WagmiProvider config={config}>
    <QueryClientProvider client={queryClient}>
      <GammaProvider
        chainId={CHAIN_ID}
        marketFactoryAddress={contracts.marketFactory}
        horizonTokenAddress={contracts.horizonToken}
        outcomeTokenAddress={contracts.outcomeToken}
        horizonPerksAddress={contracts.horizonPerks}
        feeSplitterAddress={contracts.feeSplitter}
        resolutionModuleAddress={contracts.resolutionModule}
        aiOracleAdapterAddress={contracts.aiOracleAdapter}
      >
        <RainbowKitProvider>
          <App />
        </RainbowKitProvider>
      </GammaProvider>
    </QueryClientProvider>
  </WagmiProvider>
);
