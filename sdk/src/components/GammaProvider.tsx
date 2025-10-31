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
 */
export function useGammaConfig(): GammaConfig {
  const context = useContext(GammaContext);

  if (!context) {
    throw new Error('useGammaConfig must be used within GammaProvider');
  }

  return context.config;
}

