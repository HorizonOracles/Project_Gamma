/**
 * GammaProvider - Main React provider for Project Gamma SDK
 * Simple configuration wrapper around Wagmi
 */

import React, { createContext, useContext, useMemo, ReactNode } from 'react';
import { Address } from 'viem';

interface GammaConfig {
  chainId: number;
  oracleApiUrl?: string;
  marketFactoryAddress?: Address;
  horizonTokenAddress?: Address;
  outcomeTokenAddress?: Address;
  horizonPerksAddress?: Address;
  feeSplitterAddress?: Address;
  resolutionModuleAddress?: Address;
  aiOracleAdapterAddress?: Address;
  pinataJwt?: string; // Optional Pinata JWT for IPFS storage
}

interface GammaContextValue {
  config: GammaConfig;
}

const GammaContext = createContext<GammaContextValue | null>(null);

interface GammaProviderProps {
  children: ReactNode;
  chainId: number;
  oracleApiUrl?: string;
  marketFactoryAddress?: Address;
  horizonTokenAddress?: Address;
  outcomeTokenAddress?: Address;
  horizonPerksAddress?: Address;
  feeSplitterAddress?: Address;
  resolutionModuleAddress?: Address;
  aiOracleAdapterAddress?: Address;
  pinataJwt?: string; // Optional Pinata JWT for IPFS storage
}

/**
 * GammaProvider - Provides SDK configuration to all hooks
 * 
 * @example
 * ```tsx
 * <WagmiConfig config={wagmiConfig}>
 *   <GammaProvider
 *     chainId={56}
 *     oracleApiUrl="https://api.projectgamma.io"
 *     pinataJwt="your-pinata-jwt-token" // Optional: for IPFS storage
 *   >
 *     <YourApp />
 *   </GammaProvider>
 * </WagmiConfig>
 * ```
 */
export function GammaProvider({
  children,
  chainId,
  oracleApiUrl,
  marketFactoryAddress,
  horizonTokenAddress,
  outcomeTokenAddress,
  horizonPerksAddress,
  feeSplitterAddress,
  resolutionModuleAddress,
  aiOracleAdapterAddress,
  pinataJwt,
}: GammaProviderProps) {
  const config = useMemo<GammaConfig>(
    () => ({
      chainId,
      oracleApiUrl,
      marketFactoryAddress,
      horizonTokenAddress,
      outcomeTokenAddress,
      horizonPerksAddress,
      feeSplitterAddress,
      resolutionModuleAddress,
      aiOracleAdapterAddress,
      pinataJwt,
    }),
    [
      chainId,
      oracleApiUrl,
      marketFactoryAddress,
      horizonTokenAddress,
      outcomeTokenAddress,
      horizonPerksAddress,
      feeSplitterAddress,
      resolutionModuleAddress,
      aiOracleAdapterAddress,
      pinataJwt,
    ]
  );
  
  return (
    <GammaContext.Provider value={{ config }}>
      {children}
    </GammaContext.Provider>
  );
}

/**
 * Hook to access Gamma SDK configuration
 * 
 * Must be used within GammaProvider, which should be inside WagmiProvider.
 * Following Wagmi best practices: https://wagmi.sh/react/api/WagmiProvider
 */
export function useGammaConfig(): GammaConfig {
  const context = useContext(GammaContext);

  if (!context) {
    throw new Error('useGammaConfig must be used within GammaProvider');
  }

  return context.config;
}

