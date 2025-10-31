/**
 * Hook to fetch market metadata from IPFS
 */

import { useQuery } from '@tanstack/react-query';
import { MarketMetadata } from '../../utils/ipfs';

/**
 * Hook to fetch market metadata from IPFS
 * 
 * @example
 * ```tsx
 * const { data: metadata } = useIPFSMetadata('ipfs://Qm...');
 * ```
 */
export function useIPFSMetadata(ipfsUrl: string | undefined) {
  return useQuery({
    queryKey: ['ipfsMetadata', ipfsUrl],
    queryFn: async (): Promise<MarketMetadata | null> => {
      if (!ipfsUrl) {
        return null;
      }

      // Extract IPFS hash from URL (handle both ipfs://hash and direct hash)
      const hash = ipfsUrl.replace(/^ipfs:\/\//, '').trim();
      
      if (!hash) {
        return null;
      }

      // Try multiple IPFS gateways for reliability
      const gateways = [
        `https://gateway.pinata.cloud/ipfs/${hash}`,
        `https://ipfs.io/ipfs/${hash}`,
        `https://cloudflare-ipfs.com/ipfs/${hash}`,
        `https://dweb.link/ipfs/${hash}`,
      ];

      let lastError: Error | null = null;

      for (const gatewayUrl of gateways) {
        try {
          const response = await fetch(gatewayUrl, {
            method: 'GET',
            headers: {
              'Accept': 'application/json',
            },
          });

          if (response.ok) {
            const metadata = await response.json();
            return metadata as MarketMetadata;
          }
        } catch (error) {
          lastError = error instanceof Error ? error : new Error('Failed to fetch IPFS metadata');
          // Try next gateway
          continue;
        }
      }

      // If all gateways failed, throw the last error
      throw lastError || new Error('Failed to fetch metadata from IPFS');
    },
    enabled: !!ipfsUrl && ipfsUrl.startsWith('ipfs://'),
    staleTime: 5 * 60 * 1000, // Cache for 5 minutes
    retry: 2, // Retry 2 times
  });
}

