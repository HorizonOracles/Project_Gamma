/**
 * Evidence hashing utilities for AIOracleAdapter
 * 
 * These functions ensure evidence hashes computed in the SDK match
 * those computed by the AIOracleAdapter smart contract.
 */

import { encodeAbiParameters, keccak256 } from 'viem';

/**
 * Computes the evidence hash exactly as the AIOracleAdapter contract does
 * 
 * The contract uses: keccak256(abi.encode(evidenceURIs))
 * where evidenceURIs is a string[] array
 * 
 * @param evidenceURIs - Array of evidence URL strings
 * @returns The keccak256 hash as a hex string (0x-prefixed bytes32)
 * 
 * @example
 * ```ts
 * const evidenceURIs = [
 *   'ipfs://QmTest1',
 *   'https://example.com/evidence',
 * ];
 * const hash = computeEvidenceHash(evidenceURIs);
 * // hash === '0x...' (bytes32)
 * ```
 */
export function computeEvidenceHash(evidenceURIs: string[]): `0x${string}` {
  // Empty array edge case
  if (evidenceURIs.length === 0) {
    // Encode empty string array: abi.encode([])
    const encoded = encodeAbiParameters(
      [{ type: 'string[]' }],
      [[]]
    );
    return keccak256(encoded);
  }

  // Encode the string array using ABI encoding (matches Solidity abi.encode)
  const encoded = encodeAbiParameters(
    [{ type: 'string[]' }],
    [evidenceURIs]
  );

  // Return keccak256 hash
  return keccak256(encoded);
}

/**
 * Validates that evidence URIs are properly formatted
 * 
 * @param evidenceURIs - Array of evidence URLs to validate
 * @throws Error if validation fails
 */
export function validateEvidenceURIs(evidenceURIs: string[]): void {
  if (!Array.isArray(evidenceURIs)) {
    throw new Error('evidenceURIs must be an array');
  }

  for (let i = 0; i < evidenceURIs.length; i++) {
    const uri = evidenceURIs[i];
    if (typeof uri !== 'string') {
      throw new Error(`evidenceURIs[${i}] must be a string, got ${typeof uri}`);
    }
    if (uri.length === 0) {
      throw new Error(`evidenceURIs[${i}] cannot be an empty string`);
    }
  }
}

/**
 * Sorts evidence URIs in canonical order (lexicographic)
 * 
 * Note: The AIOracleAdapter contract does NOT automatically sort evidence URIs.
 * The evidence hash is computed from the URIs in the exact order provided.
 * Use this function only if your application requires canonical ordering.
 * 
 * @param evidenceURIs - Array of evidence URLs
 * @returns Sorted copy of the array
 */
export function sortEvidenceURIs(evidenceURIs: string[]): string[] {
  return [...evidenceURIs].sort();
}
