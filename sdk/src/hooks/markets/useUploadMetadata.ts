/**
 * Hook to upload market metadata to IPFS
 */

import { useMutation } from '@tanstack/react-query';
import { useAccount } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { uploadMarketMetadata, MarketMetadata, IPFSUploadResult, IPFSProvider } from '../../utils/ipfs';

export interface UseUploadMetadataParams {
  question: string;
  description?: string;
  category: string;
  provider?: IPFSProvider;
  pinataJwt?: string; // Optional: pass JWT directly (overrides config and env vars)
}

/**
 * Hook to upload market metadata to IPFS
 * 
 * Priority for Pinata JWT:
 * 1. pinataJwt parameter (if provided)
 * 2. pinataJwt from GammaProvider config
 * 3. Environment variables (VITE_PINATA_JWT, NEXT_PUBLIC_PINATA_JWT, PINATA_JWT)
 * 
 * @example
 * ```tsx
 * const { mutate: uploadMetadata, isLoading, data } = useUploadMetadata();
 * 
 * uploadMetadata({
 *   question: 'Will BTC hit $100k?',
 *   description: 'Bitcoin price prediction',
 *   category: 'crypto',
 * });
 * ```
 * 
 * @example
 * ```tsx
 * // Override config JWT for this specific upload
 * uploadMetadata({
 *   question: 'Will BTC hit $100k?',
 *   category: 'crypto',
 *   pinataJwt: 'custom-jwt-token',
 * });
 * ```
 */
export function useUploadMetadata() {
  const { address } = useAccount();
  const config = useGammaConfig();

  return useMutation({
    mutationFn: async (params: UseUploadMetadataParams): Promise<IPFSUploadResult> => {
      const metadata: MarketMetadata = {
        question: params.question,
        description: params.description,
        category: params.category,
        createdAt: Math.floor(Date.now() / 1000),
        creator: address,
      };

      // Priority: params.pinataJwt > config.pinataJwt > undefined (will fallback to env vars in uploadMarketMetadata)
      const pinataJwt = params.pinataJwt ?? config.pinataJwt;

      return uploadMarketMetadata(metadata, params.provider, pinataJwt);
    },
  });
}

