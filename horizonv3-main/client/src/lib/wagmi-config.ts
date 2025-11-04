// Wagmi Configuration for BNB Chain (BSC)
import { getDefaultConfig } from '@rainbow-me/rainbowkit';
import { bsc, bscTestnet } from 'wagmi/chains';

export const config = getDefaultConfig({
  appName: 'Horizon Prediction Market',
  projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID || 'YOUR_PROJECT_ID', // Get from https://cloud.walletconnect.com
  chains: [
    bscTestnet, // BNB Chain Testnet (for development/testing) - Chain ID 97
    bsc,        // BNB Chain Mainnet (for production) - Chain ID 56
  ],
  ssr: false,
});
