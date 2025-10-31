/**
 * IPFS upload utilities
 * Supports multiple IPFS providers: Pinata, web3.storage, NFT.storage
 */

export interface MarketMetadata {
  question: string;
  description?: string;
  category: string;
  createdAt: number;
  creator?: string;
}

export interface IPFSUploadResult {
  hash: string;
  url: string; // Full IPFS URL (ipfs://hash)
  gatewayUrl?: string; // HTTP gateway URL for easy access
}

/**
 * Type-safe environment variable accessor
 * Handles both Vite (import.meta.env) and Node.js (process.env) environments
 */
function getEnvVariable(key: string): string | undefined {
  // Check for Vite environment (browser)
  if (typeof window !== 'undefined') {
    try {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const viteEnv = (globalThis as any).import?.meta?.env;
      if (viteEnv?.[key]) {
        return viteEnv[key];
      }
    } catch {
      // Ignore errors accessing import.meta.env
    }
  }

  // Check for Node.js environment
  if (typeof process !== 'undefined' && process.env) {
    return process.env[key];
  }

  return undefined;
}

/**
 * Upload market metadata to IPFS using Pinata
 * Requires PINATA_JWT environment variable
 */
export async function uploadToPinata(
  metadata: MarketMetadata,
  jwt?: string // Optional: pass JWT directly
): Promise<IPFSUploadResult> {
  // Priority: passed JWT > VITE_PINATA_JWT > NEXT_PUBLIC_PINATA_JWT > PINATA_JWT
  const pinataJwt = 
    jwt ||
    getEnvVariable('VITE_PINATA_JWT') ||
    getEnvVariable('NEXT_PUBLIC_PINATA_JWT') ||
    getEnvVariable('PINATA_JWT');
  
  if (!pinataJwt) {
    throw new Error(
      'Pinata JWT not configured. Set VITE_PINATA_JWT, NEXT_PUBLIC_PINATA_JWT, or PINATA_JWT environment variable, or pass jwt parameter.'
    );
  }

  const metadataJson = JSON.stringify(metadata, null, 2);
  const blob = new Blob([metadataJson], { type: 'application/json' });
  const formData = new FormData();
  formData.append('file', blob, 'market-metadata.json');

  // Pinata pinFileToIPFS options
  const pinataMetadata = JSON.stringify({
    name: `Market: ${metadata.question.substring(0, 50)}...`,
    keyvalues: {
      category: metadata.category,
      type: 'prediction-market',
    },
  });
  formData.append('pinataMetadata', pinataMetadata);

  const pinataOptions = JSON.stringify({
    cidVersion: 1,
  });
  formData.append('pinataOptions', pinataOptions);

  try {
    const response = await fetch('https://api.pinata.cloud/pinning/pinFileToIPFS', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${pinataJwt}`,
      },
      body: formData,
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: response.statusText }));
      throw new Error(`Pinata upload failed: ${error.error || response.statusText}`);
    }

    const result = await response.json();
    const ipfsHash = result.IpfsHash;

    return {
      hash: ipfsHash,
      url: `ipfs://${ipfsHash}`,
      gatewayUrl: `https://gateway.pinata.cloud/ipfs/${ipfsHash}`,
    };
  } catch (error) {
    if (error instanceof Error) {
      throw error;
    }
    throw new Error('Failed to upload to Pinata');
  }
}

/**
 * Upload market metadata to IPFS using web3.storage
 * Requires WEB3_STORAGE_TOKEN environment variable
 */
export async function uploadToWeb3Storage(
  metadata: MarketMetadata
): Promise<IPFSUploadResult> {
  const token = 
    getEnvVariable('VITE_WEB3_STORAGE_TOKEN') ||
    getEnvVariable('NEXT_PUBLIC_WEB3_STORAGE_TOKEN') ||
    getEnvVariable('WEB3_STORAGE_TOKEN');
  
  if (!token) {
    throw new Error(
      'web3.storage token not configured. Set VITE_WEB3_STORAGE_TOKEN, NEXT_PUBLIC_WEB3_STORAGE_TOKEN, or WEB3_STORAGE_TOKEN environment variable.'
    );
  }

  const metadataJson = JSON.stringify(metadata, null, 2);
  const blob = new Blob([metadataJson], { type: 'application/json' });
  const file = new File([blob], 'market-metadata.json', { type: 'application/json' });

  const formData = new FormData();
  formData.append('file', file);

  try {
    const response = await fetch('https://api.web3.storage/upload', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: formData,
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: response.statusText }));
      throw new Error(`web3.storage upload failed: ${error.error || response.statusText}`);
    }

    const result = await response.json();
    const cid = result.cid || result; // web3.storage returns CID directly or in result.cid

    return {
      hash: cid,
      url: `ipfs://${cid}`,
      gatewayUrl: `https://${cid}.ipfs.w3s.link`,
    };
  } catch (error) {
    if (error instanceof Error) {
      throw error;
    }
    throw new Error('Failed to upload to web3.storage');
  }
}

/**
 * Upload market metadata to IPFS using a public IPFS node
 * This is a fallback option that uses a public gateway
 * Note: This doesn't guarantee persistence - use Pinata or web3.storage for production
 */
export async function uploadToPublicIPFS(
  metadata: MarketMetadata
): Promise<IPFSUploadResult> {
  const metadataJson = JSON.stringify(metadata, null, 2);
  const blob = new Blob([metadataJson], { type: 'application/json' });

  // Use a public IPFS gateway for upload
  const formData = new FormData();
  formData.append('file', blob, 'market-metadata.json');

  try {
    // Try using ipfs.io public gateway
    const response = await fetch('https://ipfs.infura.io:5001/api/v0/add', {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      // Fallback to alternative public gateway
      const fallbackResponse = await fetch('https://ipfs.io/api/v0/add', {
        method: 'POST',
        body: formData,
      });

      if (!fallbackResponse.ok) {
        throw new Error('Public IPFS upload failed. Consider using Pinata or web3.storage for reliable uploads.');
      }

      const result = await fallbackResponse.json();
      const hash = result.Hash;

      return {
        hash,
        url: `ipfs://${hash}`,
        gatewayUrl: `https://ipfs.io/ipfs/${hash}`,
      };
    }

    const result = await response.json();
    const hash = result.Hash;

    return {
      hash,
      url: `ipfs://${hash}`,
      gatewayUrl: `https://ipfs.io/ipfs/${hash}`,
    };
  } catch (error) {
    if (error instanceof Error) {
      throw error;
    }
    throw new Error('Failed to upload to public IPFS');
  }
}

export type IPFSProvider = 'pinata' | 'web3.storage' | 'public';

/**
 * Upload market metadata to IPFS
 * Automatically selects provider based on available environment variables
 * Priority: Pinata > web3.storage > public IPFS
 */
export async function uploadMarketMetadata(
  metadata: MarketMetadata,
  provider?: IPFSProvider,
  pinataJwt?: string // Optional: pass JWT directly
): Promise<IPFSUploadResult> {
  // Auto-detect provider if not specified
  if (!provider) {
    const detectedPinataJwt = 
      pinataJwt ||
      getEnvVariable('VITE_PINATA_JWT') ||
      getEnvVariable('NEXT_PUBLIC_PINATA_JWT') ||
      getEnvVariable('PINATA_JWT');
    const web3Token = 
      getEnvVariable('VITE_WEB3_STORAGE_TOKEN') ||
      getEnvVariable('NEXT_PUBLIC_WEB3_STORAGE_TOKEN') ||
      getEnvVariable('WEB3_STORAGE_TOKEN');

    if (detectedPinataJwt) {
      provider = 'pinata';
    } else if (web3Token) {
      provider = 'web3.storage';
    } else {
      provider = 'public';
    }
  }

  switch (provider) {
    case 'pinata':
      return uploadToPinata(metadata, pinataJwt);
    case 'web3.storage':
      return uploadToWeb3Storage(metadata);
    case 'public':
      return uploadToPublicIPFS(metadata);
    default:
      throw new Error(`Unknown IPFS provider: ${provider}`);
  }
}

