/**
 * EIP-712 utilities for AIOracleAdapter
 * 
 * These functions compute EIP-712 hashes that match the AIOracleAdapter contract's
 * signature verification logic.
 */

import { encodeAbiParameters, keccak256, encodePacked } from 'viem';
import type { Address } from 'viem';

/**
 * ProposedOutcome struct matching the AIOracleAdapter contract
 */
export interface ProposedOutcome {
  marketId: bigint;
  outcomeId: bigint;
  closeTime: bigint;
  evidenceHash: `0x${string}`;
  notBefore: bigint;
  deadline: bigint;
}

/**
 * EIP-712 domain fields for AIOracleAdapter
 */
export interface EIP712Domain {
  name: string;
  version: string;
  chainId: number;
  verifyingContract: Address;
}

// ============ EIP-712 Constants ============

/**
 * EIP-712 domain type hash
 * keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")
 */
export const EIP712_DOMAIN_TYPEHASH = keccak256(
  encodePacked(
    ['string'],
    ['EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)']
  )
);

/**
 * ProposedOutcome type hash for EIP-712
 * keccak256("ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)")
 */
export const PROPOSED_OUTCOME_TYPEHASH = keccak256(
  encodePacked(
    ['string'],
    ['ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)']
  )
);

// ============ Domain Separator Computation ============

/**
 * Computes the EIP-712 domain separator for AIOracleAdapter
 * 
 * This MUST match the domain separator computed by the contract in its constructor.
 * 
 * @param domain - EIP-712 domain parameters
 * @returns The domain separator hash (bytes32)
 * 
 * @example
 * ```ts
 * const domainSeparator = computeDomainSeparator({
 *   name: 'AIOracleAdapter',
 *   version: '1',
 *   chainId: 56, // BNB Chain mainnet
 *   verifyingContract: '0x1234...' // AIOracleAdapter contract address
 * });
 * ```
 */
export function computeDomainSeparator(domain: EIP712Domain): `0x${string}` {
  // Encode domain struct: abi.encode(typeHash, nameHash, versionHash, chainId, verifyingContract)
  const encoded = encodeAbiParameters(
    [
      { type: 'bytes32' }, // EIP712_DOMAIN_TYPEHASH
      { type: 'bytes32' }, // keccak256(bytes(name))
      { type: 'bytes32' }, // keccak256(bytes(version))
      { type: 'uint256' }, // chainId
      { type: 'address' }, // verifyingContract
    ],
    [
      EIP712_DOMAIN_TYPEHASH,
      keccak256(encodePacked(['string'], [domain.name])),
      keccak256(encodePacked(['string'], [domain.version])),
      BigInt(domain.chainId),
      domain.verifyingContract,
    ]
  );

  return keccak256(encoded);
}

/**
 * Creates the standard EIP-712 domain for AIOracleAdapter
 * 
 * @param chainId - Chain ID (56 for BNB mainnet, 97 for testnet)
 * @param verifyingContract - AIOracleAdapter contract address
 * @returns EIP-712 domain object
 */
export function createAIOracleDomain(chainId: number, verifyingContract: Address): EIP712Domain {
  return {
    name: 'AIOracleAdapter',
    version: '1',
    chainId,
    verifyingContract,
  };
}

// ============ Proposal Hash Computation ============

/**
 * Computes the EIP-712 struct hash for a ProposedOutcome
 * 
 * This is the intermediate hash before adding the domain separator.
 * 
 * @param proposal - ProposedOutcome struct
 * @returns The struct hash (bytes32)
 */
export function computeStructHash(proposal: ProposedOutcome): `0x${string}` {
  // Encode: abi.encode(PROPOSED_OUTCOME_TYPEHASH, marketId, outcomeId, closeTime, evidenceHash, notBefore, deadline)
  const encoded = encodeAbiParameters(
    [
      { type: 'bytes32' }, // PROPOSED_OUTCOME_TYPEHASH
      { type: 'uint256' }, // marketId
      { type: 'uint256' }, // outcomeId
      { type: 'uint256' }, // closeTime
      { type: 'bytes32' }, // evidenceHash
      { type: 'uint256' }, // notBefore
      { type: 'uint256' }, // deadline
    ],
    [
      PROPOSED_OUTCOME_TYPEHASH,
      proposal.marketId,
      proposal.outcomeId,
      proposal.closeTime,
      proposal.evidenceHash,
      proposal.notBefore,
      proposal.deadline,
    ]
  );

  return keccak256(encoded);
}

/**
 * Computes the EIP-712 digest (final hash to sign) for a ProposedOutcome
 * 
 * This matches the contract's getProposalHash() function.
 * The digest is computed as: keccak256("\x19\x01" || domainSeparator || structHash)
 * 
 * @param proposal - ProposedOutcome struct
 * @param domainSeparator - EIP-712 domain separator (from computeDomainSeparator)
 * @returns The digest to sign (bytes32)
 * 
 * @example
 * ```ts
 * const proposal = {
 *   marketId: 123n,
 *   outcomeId: 0n,
 *   closeTime: 1699999999n,
 *   evidenceHash: computeEvidenceHash(['ipfs://...']),
 *   notBefore: 1699000000n,
 *   deadline: 1700000000n,
 * };
 * 
 * const domain = createAIOracleDomain(56, contractAddress);
 * const domainSeparator = computeDomainSeparator(domain);
 * const digest = computeProposalDigest(proposal, domainSeparator);
 * 
 * // Now sign the digest with a private key
 * const signature = await signMessage({ message: { raw: digest }, privateKey });
 * ```
 */
export function computeProposalDigest(
  proposal: ProposedOutcome,
  domainSeparator: `0x${string}`
): `0x${string}` {
  const structHash = computeStructHash(proposal);

  // EIP-712 digest: keccak256("\x19\x01" || domainSeparator || structHash)
  const digest = keccak256(
    encodePacked(
      ['bytes1', 'bytes1', 'bytes32', 'bytes32'],
      ['0x19', '0x01', domainSeparator, structHash]
    )
  );

  return digest;
}

/**
 * Convenience function to compute proposal digest with inline domain construction
 * 
 * @param proposal - ProposedOutcome struct
 * @param chainId - Chain ID (56 for BNB mainnet, 97 for testnet)
 * @param verifyingContract - AIOracleAdapter contract address
 * @returns The digest to sign (bytes32)
 */
export function computeProposalHash(
  proposal: ProposedOutcome,
  chainId: number,
  verifyingContract: Address
): `0x${string}` {
  const domain = createAIOracleDomain(chainId, verifyingContract);
  const domainSeparator = computeDomainSeparator(domain);
  return computeProposalDigest(proposal, domainSeparator);
}

// ============ Helper Functions ============

/**
 * Builds a ProposedOutcome struct with sensible defaults
 * 
 * @param params - Partial proposal parameters
 * @returns Complete ProposedOutcome struct
 * 
 * @example
 * ```ts
 * const proposal = buildProposedOutcome({
 *   marketId: 123n,
 *   outcomeId: 0n, // YES
 *   closeTime: market.closeTime,
 *   evidenceHash: computeEvidenceHash(evidenceURIs),
 *   // notBefore defaults to now
 *   // deadline defaults to now + 1 hour
 * });
 * ```
 */
export function buildProposedOutcome(
  params: Partial<ProposedOutcome> & {
    marketId: bigint;
    outcomeId: bigint;
    closeTime: bigint;
    evidenceHash: `0x${string}`;
  }
): ProposedOutcome {
  const now = BigInt(Math.floor(Date.now() / 1000));
  const oneHour = 3600n;

  return {
    marketId: params.marketId,
    outcomeId: params.outcomeId,
    closeTime: params.closeTime,
    evidenceHash: params.evidenceHash,
    notBefore: params.notBefore ?? now,
    deadline: params.deadline ?? now + oneHour,
  };
}

/**
 * Validates a ProposedOutcome struct
 * 
 * @param proposal - ProposedOutcome to validate
 * @throws Error if validation fails
 */
export function validateProposedOutcome(proposal: ProposedOutcome): void {
  if (proposal.marketId < 0n) {
    throw new Error('marketId must be non-negative');
  }
  if (proposal.outcomeId < 0n) {
    throw new Error('outcomeId must be non-negative');
  }
  if (proposal.closeTime <= 0n) {
    throw new Error('closeTime must be positive');
  }
  if (!proposal.evidenceHash || proposal.evidenceHash.length !== 66) {
    throw new Error('evidenceHash must be a valid bytes32 hex string (0x + 64 chars)');
  }
  if (proposal.notBefore < 0n) {
    throw new Error('notBefore must be non-negative');
  }
  if (proposal.deadline <= proposal.notBefore) {
    throw new Error('deadline must be after notBefore');
  }
}
