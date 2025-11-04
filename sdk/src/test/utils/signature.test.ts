/**
 * Unit tests for signature utilities
 * 
 * These tests verify signature creation, normalization, and verification.
 */

import { describe, it, expect } from 'vitest';
import { privateKeyToAccount, generatePrivateKey } from 'viem/accounts';
import {
  signProposal,
  signProposalWithAccount,
  normalizeSignature,
  splitSignature,
  joinSignature,
  verifyProposalSignature,
  isValidProposalSignature,
  isValidSignature,
  hashSignature,
  createAccountFromPrivateKey,
  formatSignature,
} from '../../utils/signature';
import {
  computeDomainSeparator,
  createAIOracleDomain,
  buildProposedOutcome,
} from '../../utils/eip712';
import { computeEvidenceHash } from '../../utils/evidence';
import type { ProposedOutcome } from '../../types';

// Test constants
const TEST_CONTRACT_ADDRESS = '0x1234567890123456789012345678901234567890' as const;
const TEST_CHAIN_ID = 56;

describe('isValidSignature', () => {
  it('should accept valid 65-byte signature', () => {
    const validSig = '0x' + 'a'.repeat(130);
    expect(isValidSignature(validSig)).toBe(true);
  });

  it('should reject non-hex string', () => {
    expect(isValidSignature('not a hex string')).toBe(false);
  });

  it('should reject signature with wrong length', () => {
    const tooShort = '0x' + 'a'.repeat(128);
    const tooLong = '0x' + 'a'.repeat(132);
    expect(isValidSignature(tooShort)).toBe(false);
    expect(isValidSignature(tooLong)).toBe(false);
  });

  it('should reject signature without 0x prefix', () => {
    const noPrefix = 'a'.repeat(130);
    expect(isValidSignature(noPrefix)).toBe(false);
  });
});

describe('splitSignature and joinSignature', () => {
  it('should split and rejoin signature correctly', () => {
    const testSig = '0x' + 
      'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa' + // r (32 bytes)
      'bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb' + // s (32 bytes)
      '1b'; // v (1 byte, 27 in hex)

    const components = splitSignature(testSig);
    expect(components.r).toBe('0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa');
    expect(components.s).toBe('0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb');
    expect(components.v).toBe(27);

    const rejoined = joinSignature(components);
    expect(rejoined.toLowerCase()).toBe(testSig.toLowerCase());
  });

  it('should throw on invalid signature', () => {
    const invalidSig = '0xinvalid';
    expect(() => splitSignature(invalidSig)).toThrow('Invalid signature format');
  });
});

describe('normalizeSignature', () => {
  it('should keep v=27 unchanged', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '1b'; // v = 27

    const normalized = normalizeSignature(sig);
    const { v } = splitSignature(normalized);
    expect(v).toBe(27);
  });

  it('should keep v=28 unchanged', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '1c'; // v = 28

    const normalized = normalizeSignature(sig);
    const { v } = splitSignature(normalized);
    expect(v).toBe(28);
  });

  it('should convert v=0 to v=27', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '00'; // v = 0

    const normalized = normalizeSignature(sig);
    const { v } = splitSignature(normalized);
    expect(v).toBe(27);
  });

  it('should convert v=1 to v=28', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '01'; // v = 1

    const normalized = normalizeSignature(sig);
    const { v } = splitSignature(normalized);
    expect(v).toBe(28);
  });

  it('should throw on invalid v value', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '1f'; // v = 31 (invalid)

    // Viem will throw before we can validate
    expect(() => normalizeSignature(sig)).toThrow();
  });
});

describe('createAccountFromPrivateKey', () => {
  it('should create account from valid private key', () => {
    const privateKey = generatePrivateKey();
    const account = createAccountFromPrivateKey(privateKey);

    expect(account).toBeDefined();
    expect(account.address).toMatch(/^0x[0-9a-f]{40}$/i);
  });
});

describe('hashSignature', () => {
  it('should hash signature correctly', () => {
    const sig = '0x' + 'a'.repeat(130);
    const hash = hashSignature(sig);

    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should produce consistent hashes', () => {
    const sig = '0x' + 'a'.repeat(130);
    const hash1 = hashSignature(sig);
    const hash2 = hashSignature(sig);

    expect(hash1).toBe(hash2);
  });

  it('should produce different hashes for different signatures', () => {
    const sig1 = '0x' + 'a'.repeat(130);
    const sig2 = '0x' + 'b'.repeat(130);

    const hash1 = hashSignature(sig1);
    const hash2 = hashSignature(sig2);

    expect(hash1).not.toBe(hash2);
  });

  it('should throw on invalid signature', () => {
    const invalidSig = '0xinvalid';
    expect(() => hashSignature(invalidSig)).toThrow('Invalid signature format');
  });
});

describe('formatSignature', () => {
  it('should format signature for display', () => {
    const sig = '0x' + 
      'a'.repeat(64) + 
      'b'.repeat(64) + 
      '1b';

    const formatted = formatSignature(sig);
    expect(formatted).toContain('Signature');
    expect(formatted).toContain('r:');
    expect(formatted).toContain('s:');
    expect(formatted).toContain('v:');
  });
});

describe('signProposal', () => {
  it('should sign proposal correctly', async () => {
    const privateKey = generatePrivateKey();
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposal(proposal, domainSeparator, privateKey);

    // Should be a valid 65-byte signature
    expect(isValidSignature(signature)).toBe(true);
  });

  it('should produce normalized signature (v=27 or v=28)', async () => {
    const privateKey = generatePrivateKey();
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposal(proposal, domainSeparator, privateKey);
    const { v } = splitSignature(signature);

    expect(v === 27 || v === 28).toBe(true);
  });
});

describe('signProposalWithAccount', () => {
  it('should sign proposal with account', async () => {
    const privateKey = generatePrivateKey();
    const account = privateKeyToAccount(privateKey);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposalWithAccount(proposal, domainSeparator, account);

    expect(isValidSignature(signature)).toBe(true);
  });
});

describe('verifyProposalSignature', () => {
  it('should recover correct signer address', async () => {
    const privateKey = generatePrivateKey();
    const account = privateKeyToAccount(privateKey);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposal(proposal, domainSeparator, privateKey);
    const recoveredSigner = await verifyProposalSignature(proposal, domainSeparator, signature);

    expect(recoveredSigner.toLowerCase()).toBe(account.address.toLowerCase());
  });

  it('should work with different proposals', async () => {
    const privateKey = generatePrivateKey();
    const account = privateKeyToAccount(privateKey);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal1 = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest1']),
    });

    const proposal2 = buildProposedOutcome({
      marketId: 456n,
      outcomeId: 1n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest2']),
    });

    const signature1 = await signProposal(proposal1, domainSeparator, privateKey);
    const signature2 = await signProposal(proposal2, domainSeparator, privateKey);

    const signer1 = await verifyProposalSignature(proposal1, domainSeparator, signature1);
    const signer2 = await verifyProposalSignature(proposal2, domainSeparator, signature2);

    expect(signer1.toLowerCase()).toBe(account.address.toLowerCase());
    expect(signer2.toLowerCase()).toBe(account.address.toLowerCase());
    expect(signature1).not.toBe(signature2); // Different proposals = different signatures
  });
});

describe('isValidProposalSignature', () => {
  it('should return true for valid signature', async () => {
    const privateKey = generatePrivateKey();
    const account = privateKeyToAccount(privateKey);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposal(proposal, domainSeparator, privateKey);
    const isValid = await isValidProposalSignature(
      proposal,
      domainSeparator,
      signature,
      account.address
    );

    expect(isValid).toBe(true);
  });

  it('should return false for wrong signer', async () => {
    const privateKey1 = generatePrivateKey();
    const privateKey2 = generatePrivateKey();
    const account2 = privateKeyToAccount(privateKey2);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    });

    const signature = await signProposal(proposal, domainSeparator, privateKey1);
    const isValid = await isValidProposalSignature(
      proposal,
      domainSeparator,
      signature,
      account2.address // Wrong signer
    );

    expect(isValid).toBe(false);
  });

  it('should return false for wrong proposal', async () => {
    const privateKey = generatePrivateKey();
    const account = privateKeyToAccount(privateKey);
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal1 = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest1']),
    });

    const proposal2 = buildProposedOutcome({
      marketId: 456n,
      outcomeId: 1n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest2']),
    });

    const signature = await signProposal(proposal1, domainSeparator, privateKey);
    const isValid = await isValidProposalSignature(
      proposal2, // Wrong proposal
      domainSeparator,
      signature,
      account.address
    );

    expect(isValid).toBe(false);
  });
});

describe('Integration: Complete Signing Flow', () => {
  it('should complete full signing and verification flow', async () => {
    // Step 1: Create AI signer account
    const aiPrivateKey = generatePrivateKey();
    const aiAccount = privateKeyToAccount(aiPrivateKey);

    // Step 2: Set up EIP-712 domain
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    // Step 3: Create proposal
    const evidenceURIs = ['ipfs://QmTest1', 'https://example.com/evidence'];
    const evidenceHash = computeEvidenceHash(evidenceURIs);

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n, // YES
      closeTime: 1699999999n,
      evidenceHash,
      notBefore: 1699000000n,
      deadline: 1700000000n,
    });

    // Step 4: Sign proposal
    const signature = await signProposal(proposal, domainSeparator, aiPrivateKey);

    // Step 5: Verify signature
    const recoveredSigner = await verifyProposalSignature(proposal, domainSeparator, signature);
    expect(recoveredSigner.toLowerCase()).toBe(aiAccount.address.toLowerCase());

    // Step 6: Validate signature is correct format
    expect(isValidSignature(signature)).toBe(true);

    // Step 7: Check signature components
    const { v } = splitSignature(signature);
    expect(v === 27 || v === 28).toBe(true);

    // This signature can now be used to call the AIOracleAdapter.proposeAI() function
    // The contract will verify the signature and recover the same signer address
  });
});
