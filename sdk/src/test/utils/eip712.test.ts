/**
 * Unit tests for EIP-712 utilities
 * 
 * These tests verify that our EIP-712 implementation matches the contract's behavior.
 */

import { describe, it, expect } from 'vitest';
import { keccak256, encodePacked } from 'viem';
import {
  computeDomainSeparator,
  createAIOracleDomain,
  computeStructHash,
  computeProposalDigest,
  computeProposalHash,
  buildProposedOutcome,
  validateProposedOutcome,
  EIP712_DOMAIN_TYPEHASH,
  PROPOSED_OUTCOME_TYPEHASH,
} from '../../utils/eip712';
import { computeEvidenceHash } from '../../utils/evidence';
import type { ProposedOutcome } from '../../types';

// Test constants
const TEST_CONTRACT_ADDRESS = '0x1234567890123456789012345678901234567890' as const;
const TEST_CHAIN_ID = 56; // BNB Chain mainnet

describe('EIP-712 Constants', () => {
  it('should have correct EIP712_DOMAIN_TYPEHASH', () => {
    const expected = keccak256(
      encodePacked(
        ['string'],
        ['EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)']
      )
    );
    expect(EIP712_DOMAIN_TYPEHASH).toBe(expected);
  });

  it('should have correct PROPOSED_OUTCOME_TYPEHASH', () => {
    const expected = keccak256(
      encodePacked(
        ['string'],
        ['ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)']
      )
    );
    expect(PROPOSED_OUTCOME_TYPEHASH).toBe(expected);
  });
});

describe('createAIOracleDomain', () => {
  it('should create domain with correct fields', () => {
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);

    expect(domain.name).toBe('AIOracleAdapter');
    expect(domain.version).toBe('1');
    expect(domain.chainId).toBe(TEST_CHAIN_ID);
    expect(domain.verifyingContract).toBe(TEST_CONTRACT_ADDRESS);
  });

  it('should work with testnet chain ID', () => {
    const domain = createAIOracleDomain(97, TEST_CONTRACT_ADDRESS);
    expect(domain.chainId).toBe(97);
  });
});

describe('computeDomainSeparator', () => {
  it('should compute domain separator correctly', () => {
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const separator = computeDomainSeparator(domain);

    // Should be a valid bytes32 hex string
    expect(separator).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should produce consistent results', () => {
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const separator1 = computeDomainSeparator(domain);
    const separator2 = computeDomainSeparator(domain);

    expect(separator1).toBe(separator2);
  });

  it('should produce different results for different chain IDs', () => {
    const domain1 = createAIOracleDomain(56, TEST_CONTRACT_ADDRESS);
    const domain2 = createAIOracleDomain(97, TEST_CONTRACT_ADDRESS);

    const separator1 = computeDomainSeparator(domain1);
    const separator2 = computeDomainSeparator(domain2);

    expect(separator1).not.toBe(separator2);
  });

  it('should produce different results for different contract addresses', () => {
    const domain1 = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domain2 = createAIOracleDomain(TEST_CHAIN_ID, '0x9876543210987654321098765432109876543210');

    const separator1 = computeDomainSeparator(domain1);
    const separator2 = computeDomainSeparator(domain2);

    expect(separator1).not.toBe(separator2);
  });
});

describe('buildProposedOutcome', () => {
  it('should build proposal with required fields', () => {
    const evidenceHash = computeEvidenceHash(['ipfs://QmTest']);
    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash,
    });

    expect(proposal.marketId).toBe(123n);
    expect(proposal.outcomeId).toBe(0n);
    expect(proposal.closeTime).toBe(1699999999n);
    expect(proposal.evidenceHash).toBe(evidenceHash);
    expect(proposal.notBefore).toBeDefined();
    expect(proposal.deadline).toBeDefined();
  });

  it('should use provided notBefore and deadline', () => {
    const evidenceHash = computeEvidenceHash(['ipfs://QmTest']);
    const notBefore = 1699000000n;
    const deadline = 1700000000n;

    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash,
      notBefore,
      deadline,
    });

    expect(proposal.notBefore).toBe(notBefore);
    expect(proposal.deadline).toBe(deadline);
  });

  it('should default deadline to ~1 hour from now', () => {
    const evidenceHash = computeEvidenceHash(['ipfs://QmTest']);
    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash,
    });

    const oneHour = 3600n;
    const expectedDeadline = proposal.notBefore + oneHour;

    expect(proposal.deadline).toBe(expectedDeadline);
  });
});

describe('validateProposedOutcome', () => {
  const validProposal: ProposedOutcome = {
    marketId: 123n,
    outcomeId: 0n,
    closeTime: 1699999999n,
    evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
    notBefore: 1699000000n,
    deadline: 1700000000n,
  };

  it('should accept valid proposal', () => {
    expect(() => validateProposedOutcome(validProposal)).not.toThrow();
  });

  it('should reject negative marketId', () => {
    const invalid = { ...validProposal, marketId: -1n };
    expect(() => validateProposedOutcome(invalid)).toThrow('marketId must be non-negative');
  });

  it('should reject negative outcomeId', () => {
    const invalid = { ...validProposal, outcomeId: -1n };
    expect(() => validateProposedOutcome(invalid)).toThrow('outcomeId must be non-negative');
  });

  it('should reject non-positive closeTime', () => {
    const invalid = { ...validProposal, closeTime: 0n };
    expect(() => validateProposedOutcome(invalid)).toThrow('closeTime must be positive');
  });

  it('should reject invalid evidenceHash', () => {
    const invalid = { ...validProposal, evidenceHash: '0xinvalid' as `0x${string}` };
    expect(() => validateProposedOutcome(invalid)).toThrow('evidenceHash must be a valid bytes32');
  });

  it('should reject deadline before notBefore', () => {
    const invalid = { ...validProposal, deadline: 1698999999n };
    expect(() => validateProposedOutcome(invalid)).toThrow('deadline must be after notBefore');
  });
});

describe('computeStructHash', () => {
  it('should compute struct hash correctly', () => {
    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const structHash = computeStructHash(proposal);

    // Should be a valid bytes32 hex string
    expect(structHash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should produce consistent results', () => {
    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const hash1 = computeStructHash(proposal);
    const hash2 = computeStructHash(proposal);

    expect(hash1).toBe(hash2);
  });

  it('should produce different hashes for different marketIds', () => {
    const evidenceHash = computeEvidenceHash(['ipfs://QmTest']);
    const proposal1: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash,
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const proposal2: ProposedOutcome = {
      ...proposal1,
      marketId: 456n,
    };

    const hash1 = computeStructHash(proposal1);
    const hash2 = computeStructHash(proposal2);

    expect(hash1).not.toBe(hash2);
  });
});

describe('computeProposalDigest', () => {
  it('should compute proposal digest correctly', () => {
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const digest = computeProposalDigest(proposal, domainSeparator);

    // Should be a valid bytes32 hex string
    expect(digest).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should produce consistent results', () => {
    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);

    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const digest1 = computeProposalDigest(proposal, domainSeparator);
    const digest2 = computeProposalDigest(proposal, domainSeparator);

    expect(digest1).toBe(digest2);
  });

  it('should produce different digests for different chain IDs', () => {
    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const domain1 = createAIOracleDomain(56, TEST_CONTRACT_ADDRESS);
    const separator1 = computeDomainSeparator(domain1);
    const digest1 = computeProposalDigest(proposal, separator1);

    const domain2 = createAIOracleDomain(97, TEST_CONTRACT_ADDRESS);
    const separator2 = computeDomainSeparator(domain2);
    const digest2 = computeProposalDigest(proposal, separator2);

    expect(digest1).not.toBe(digest2);
  });
});

describe('computeProposalHash', () => {
  it('should compute proposal hash (convenience function)', () => {
    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const hash = computeProposalHash(proposal, TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);

    // Should be a valid bytes32 hex string
    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('should match computeProposalDigest result', () => {
    const proposal: ProposedOutcome = {
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash: computeEvidenceHash(['ipfs://QmTest']),
      notBefore: 1699000000n,
      deadline: 1700000000n,
    };

    const hash1 = computeProposalHash(proposal, TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);

    const domain = createAIOracleDomain(TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);
    const domainSeparator = computeDomainSeparator(domain);
    const hash2 = computeProposalDigest(proposal, domainSeparator);

    expect(hash1).toBe(hash2);
  });
});

describe('Integration: Evidence Hash + EIP-712', () => {
  it('should create complete signed proposal structure', () => {
    // Step 1: Create evidence and compute hash
    const evidenceURIs = [
      'ipfs://QmTest1',
      'https://example.com/evidence',
    ];
    const evidenceHash = computeEvidenceHash(evidenceURIs);

    // Step 2: Build proposal
    const proposal = buildProposedOutcome({
      marketId: 123n,
      outcomeId: 0n,
      closeTime: 1699999999n,
      evidenceHash,
      notBefore: 1699000000n,
      deadline: 1700000000n,
    });

    // Step 3: Validate proposal
    expect(() => validateProposedOutcome(proposal)).not.toThrow();

    // Step 4: Compute digest to sign
    const digest = computeProposalHash(proposal, TEST_CHAIN_ID, TEST_CONTRACT_ADDRESS);

    expect(digest).toMatch(/^0x[0-9a-f]{64}$/i);
  });
});
