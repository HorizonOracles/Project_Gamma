/**
 * Unit tests for evidence hashing utilities
 * 
 * These tests verify that evidence hashes computed in the SDK match
 * those computed by the AIOracleAdapter smart contract.
 */

import { describe, it, expect } from 'vitest';
import { encodeAbiParameters, keccak256 } from 'viem';
import {
  computeEvidenceHash,
  validateEvidenceURIs,
  sortEvidenceURIs,
} from '../../utils/evidence';

describe('computeEvidenceHash', () => {
  it('should compute hash for single URI', () => {
    const uris = ['ipfs://QmTest'];
    const hash = computeEvidenceHash(uris);

    // Should be a valid bytes32 hex string
    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should compute hash for multiple URIs', () => {
    const uris = [
      'ipfs://QmTest1',
      'https://example.com/evidence',
      'https://twitter.com/status/123',
    ];
    const hash = computeEvidenceHash(uris);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should compute hash for empty array', () => {
    const uris: string[] = [];
    const hash = computeEvidenceHash(uris);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should produce consistent hashes', () => {
    const uris = ['ipfs://QmTest1', 'ipfs://QmTest2'];
    const hash1 = computeEvidenceHash(uris);
    const hash2 = computeEvidenceHash(uris);

    expect(hash1).toBe(hash2);
  });

  it('should produce different hashes for different URIs', () => {
    const uris1 = ['ipfs://QmTest1'];
    const uris2 = ['ipfs://QmTest2'];

    const hash1 = computeEvidenceHash(uris1);
    const hash2 = computeEvidenceHash(uris2);

    expect(hash1).not.toBe(hash2);
  });

  it('should be sensitive to URI order', () => {
    const uris1 = ['ipfs://QmTest1', 'ipfs://QmTest2'];
    const uris2 = ['ipfs://QmTest2', 'ipfs://QmTest1'];

    const hash1 = computeEvidenceHash(uris1);
    const hash2 = computeEvidenceHash(uris2);

    // Different order = different hash (contract does NOT sort)
    expect(hash1).not.toBe(hash2);
  });

  it('should match manual ABI encoding', () => {
    const uris = ['ipfs://QmTest'];
    
    // Compute hash using our function
    const hash1 = computeEvidenceHash(uris);

    // Compute hash manually (simulating contract behavior)
    const encoded = encodeAbiParameters([{ type: 'string[]' }], [uris]);
    const hash2 = keccak256(encoded);

    expect(hash1).toBe(hash2);
  });

  it('should handle long URIs', () => {
    const longUri = 'https://example.com/' + 'a'.repeat(1000);
    const uris = [longUri];
    const hash = computeEvidenceHash(uris);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should handle special characters in URIs', () => {
    const uris = [
      'ipfs://QmTest?param=value',
      'https://example.com/path/to/file.json',
      'ar://abc123-xyz789',
    ];
    const hash = computeEvidenceHash(uris);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should handle Unicode in URIs', () => {
    const uris = ['https://example.com/证据', 'ipfs://QmTest-日本語'];
    const hash = computeEvidenceHash(uris);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });
});

describe('validateEvidenceURIs', () => {
  it('should accept valid URI array', () => {
    const uris = ['ipfs://QmTest', 'https://example.com'];
    expect(() => validateEvidenceURIs(uris)).not.toThrow();
  });

  it('should accept empty array', () => {
    const uris: string[] = [];
    expect(() => validateEvidenceURIs(uris)).not.toThrow();
  });

  it('should reject non-array input', () => {
    expect(() => validateEvidenceURIs('not an array' as any)).toThrow('must be an array');
  });

  it('should reject array with non-string elements', () => {
    const uris = ['valid', 123, 'valid'] as any;
    expect(() => validateEvidenceURIs(uris)).toThrow('must be a string');
  });

  it('should reject array with empty string', () => {
    const uris = ['valid', '', 'valid'];
    expect(() => validateEvidenceURIs(uris)).toThrow('cannot be an empty string');
  });

  it('should reject array with null', () => {
    const uris = ['valid', null, 'valid'] as any;
    expect(() => validateEvidenceURIs(uris)).toThrow('must be a string');
  });

  it('should reject array with undefined', () => {
    const uris = ['valid', undefined, 'valid'] as any;
    expect(() => validateEvidenceURIs(uris)).toThrow('must be a string');
  });
});

describe('sortEvidenceURIs', () => {
  it('should sort URIs lexicographically', () => {
    const uris = ['zebra', 'apple', 'banana'];
    const sorted = sortEvidenceURIs(uris);

    expect(sorted).toEqual(['apple', 'banana', 'zebra']);
  });

  it('should not mutate original array', () => {
    const original = ['zebra', 'apple', 'banana'];
    const sorted = sortEvidenceURIs(original);

    expect(original).toEqual(['zebra', 'apple', 'banana']); // Unchanged
    expect(sorted).toEqual(['apple', 'banana', 'zebra']);
  });

  it('should handle already sorted array', () => {
    const uris = ['apple', 'banana', 'zebra'];
    const sorted = sortEvidenceURIs(uris);

    expect(sorted).toEqual(['apple', 'banana', 'zebra']);
  });

  it('should handle empty array', () => {
    const uris: string[] = [];
    const sorted = sortEvidenceURIs(uris);

    expect(sorted).toEqual([]);
  });

  it('should handle single element', () => {
    const uris = ['single'];
    const sorted = sortEvidenceURIs(uris);

    expect(sorted).toEqual(['single']);
  });

  it('should sort case-sensitively', () => {
    const uris = ['Zebra', 'apple', 'Banana'];
    const sorted = sortEvidenceURIs(uris);

    // Uppercase letters come before lowercase in lexicographic order
    expect(sorted).toEqual(['Banana', 'Zebra', 'apple']);
  });
});

describe('Integration: Evidence hash order sensitivity', () => {
  it('should demonstrate that order matters', () => {
    const uris = ['ipfs://QmTest1', 'ipfs://QmTest2', 'ipfs://QmTest3'];
    
    // Hash in original order
    const hash1 = computeEvidenceHash(uris);

    // Hash in reverse order
    const reversedUris = [...uris].reverse();
    const hash2 = computeEvidenceHash(reversedUris);

    // Hash in sorted order (same as original in this case, so use different order)
    const reorderedUris = ['ipfs://QmTest2', 'ipfs://QmTest1', 'ipfs://QmTest3'];
    const hash3 = computeEvidenceHash(reorderedUris);

    // All three should be different
    expect(hash1).not.toBe(hash2);
    expect(hash1).not.toBe(hash3);
    expect(hash2).not.toBe(hash3);
  });

  it('should show canonical ordering use case', () => {
    // User provides evidence in arbitrary order
    const userProvidedURIs = [
      'https://twitter.com/status/456',
      'ipfs://QmAbc',
      'https://example.com/doc.pdf',
      'ipfs://QmXyz',
    ];

    // For canonical ordering, sort before hashing
    const sortedURIs = sortEvidenceURIs(userProvidedURIs);
    const canonicalHash = computeEvidenceHash(sortedURIs);

    // Any user providing the same set of URIs will get the same canonical hash
    const anotherUserURIs = [
      'ipfs://QmXyz',
      'https://example.com/doc.pdf',
      'ipfs://QmAbc',
      'https://twitter.com/status/456',
    ];
    const anotherSortedURIs = sortEvidenceURIs(anotherUserURIs);
    const anotherCanonicalHash = computeEvidenceHash(anotherSortedURIs);

    expect(canonicalHash).toBe(anotherCanonicalHash);
  });

  it('should show contract behavior (order preserved)', () => {
    // The contract does NOT sort evidence URIs
    // It hashes them in the exact order provided
    const uris1 = ['ipfs://QmTest1', 'ipfs://QmTest2'];
    const uris2 = ['ipfs://QmTest2', 'ipfs://QmTest1'];

    const hash1 = computeEvidenceHash(uris1);
    const hash2 = computeEvidenceHash(uris2);

    // Different order = different hash
    expect(hash1).not.toBe(hash2);

    // To ensure consistency, users should either:
    // 1. Maintain the same order when creating proposals
    // 2. Use sortEvidenceURIs() for canonical ordering
  });
});
