/**
 * Signature utilities for EIP-712 proposals
 * 
 * These functions handle signing ProposedOutcome structs and verifying signatures
 * to match the AIOracleAdapter contract's signature verification.
 */

import {
  type Hex,
  type PrivateKeyAccount,
  type SignableMessage,
  hexToSignature,
  signatureToHex,
  recoverAddress,
  isHex,
} from 'viem';
import { privateKeyToAccount, sign } from 'viem/accounts';
import type { ProposedOutcome } from './eip712';
import { computeProposalDigest } from './eip712';

/**
 * Signature result with r, s, v components
 */
export interface SignatureComponents {
  r: Hex;
  s: Hex;
  v: number;
}

/**
 * Complete signature (65 bytes)
 */
export type Signature = Hex; // 0x + 130 hex chars (65 bytes)

// ============ Signature Creation ============

/**
 * Signs a ProposedOutcome using a private key
 * 
 * This produces an EIP-712 signature that can be verified by the AIOracleAdapter contract.
 * 
 * @param proposal - ProposedOutcome struct to sign
 * @param domainSeparator - EIP-712 domain separator
 * @param privateKey - Private key as hex string (0x-prefixed, 32 bytes)
 * @returns 65-byte signature (r + s + v format)
 * 
 * @example
 * ```ts
 * const proposal = buildProposedOutcome({
 *   marketId: 123n,
 *   outcomeId: 0n,
 *   closeTime: market.closeTime,
 *   evidenceHash: computeEvidenceHash(evidenceURIs),
 * });
 * 
 * const domain = createAIOracleDomain(56, contractAddress);
 * const domainSeparator = computeDomainSeparator(domain);
 * 
 * const signature = await signProposal(
 *   proposal,
 *   domainSeparator,
 *   '0x1234...' // AI signer private key
 * );
 * 
 * // Submit to contract
 * await aiOracleAdapter.write.proposeAI([proposal, signature, bondAmount, evidenceURIs]);
 * ```
 */
export async function signProposal(
  proposal: ProposedOutcome,
  domainSeparator: Hex,
  privateKey: Hex
): Promise<Signature> {
  // Compute the digest to sign
  const digest = computeProposalDigest(proposal, domainSeparator);

  // Sign the digest directly (without Ethereum signed message prefix)
  // This is correct for EIP-712 signatures
  const signature = await sign({
    hash: digest,
    privateKey,
  });

  // Convert { r, s, v, yParity } to hex string and normalize
  const signatureHex = signatureToHex(signature);
  return normalizeSignature(signatureHex);
}

/**
 * Signs a ProposedOutcome using a viem account (wallet client)
 * 
 * @param proposal - ProposedOutcome struct to sign
 * @param domainSeparator - EIP-712 domain separator
 * @param account - Viem PrivateKeyAccount or compatible signer
 * @returns 65-byte signature (r + s + v format)
 */
export async function signProposalWithAccount(
  proposal: ProposedOutcome,
  domainSeparator: Hex,
  account: PrivateKeyAccount
): Promise<Signature> {
  const digest = computeProposalDigest(proposal, domainSeparator);

  // Use the account's sign method which handles the signing internally
  const signature = await account.sign({
    hash: digest,
  });

  // Normalize v value (signature is already a hex string)
  return normalizeSignature(signature);
}

// ============ Signature Normalization ============

/**
 * Normalizes a signature to ensure v is 27 or 28
 * 
 * Some wallets produce signatures with v = 0 or 1, but the contract expects v = 27 or 28.
 * This function converts v values to the correct range.
 * 
 * @param signature - 65-byte signature
 * @returns Normalized 65-byte signature
 */
export function normalizeSignature(signature: Hex): Signature {
  if (!isValidSignature(signature)) {
    throw new Error('Invalid signature format: must be 0x + 130 hex characters (65 bytes)');
  }

  // Parse signature components
  const parsed = hexToSignature(signature);
  const { r, s } = parsed;
  
  // Viem returns different formats depending on v value:
  // - v=27/28: returns { r, s, v: bigint, yParity: number }
  // - v=0/1: returns { r, s, yParity: number } (no v field!)
  let normalizedV: bigint;
  
  if ('v' in parsed && parsed.v !== undefined) {
    // v is present (27 or 28)
    const vNum = Number(parsed.v);
    if (vNum === 27 || vNum === 28) {
      normalizedV = parsed.v;
    } else {
      throw new Error(`Invalid v value: ${vNum}. Expected 27 or 28.`);
    }
  } else if ('yParity' in parsed && parsed.yParity !== undefined) {
    // Only yParity is present (v was 0 or 1)
    // Convert yParity (0 or 1) to v (27 or 28)
    if (parsed.yParity === 0 || parsed.yParity === 1) {
      normalizedV = BigInt(parsed.yParity + 27);
    } else {
      throw new Error(`Invalid yParity value: ${parsed.yParity}. Expected 0 or 1.`);
    }
  } else {
    throw new Error('Signature must have either v or yParity field');
  }

  // Reconstruct signature with normalized v (as bigint)
  return signatureToHex({ r, s, v: normalizedV });
}

/**
 * Splits a signature into r, s, v components
 * 
 * @param signature - 65-byte signature
 * @returns Signature components
 */
export function splitSignature(signature: Hex): SignatureComponents {
  if (!isValidSignature(signature)) {
    throw new Error('Invalid signature format: must be 0x + 130 hex characters (65 bytes)');
  }

  const parsed = hexToSignature(signature);
  const { r, s } = parsed;
  
  // Viem returns v as bigint when present, or only yParity when v=0/1
  let v: number;
  if ('v' in parsed && parsed.v !== undefined) {
    v = Number(parsed.v);
  } else if ('yParity' in parsed && parsed.yParity !== undefined) {
    // Convert yParity (0 or 1) to v (27 or 28)
    v = parsed.yParity + 27;
  } else {
    throw new Error('Signature must have either v or yParity field');
  }
  
  return { r, s, v };
}

/**
 * Combines r, s, v components into a 65-byte signature
 * 
 * @param components - Signature components
 * @returns 65-byte signature
 */
export function joinSignature(components: SignatureComponents): Signature {
  // Viem's signatureToHex requires v to be a bigint
  return signatureToHex({
    r: components.r,
    s: components.s,
    v: BigInt(components.v),
  });
}

// ============ Signature Verification ============

/**
 * Verifies a signature and recovers the signer address
 * 
 * This matches the contract's ecrecover logic.
 * 
 * @param proposal - ProposedOutcome that was signed
 * @param domainSeparator - EIP-712 domain separator
 * @param signature - 65-byte signature to verify
 * @returns The address that signed the proposal
 * 
 * @example
 * ```ts
 * const signer = await verifyProposalSignature(proposal, domainSeparator, signature);
 * 
 * // Check if signer is an allowed AI signer
 * const isAllowed = await aiOracleAdapter.read.allowedSigners([signer]);
 * if (!isAllowed) {
 *   throw new Error('Signature is not from an allowed AI signer');
 * }
 * ```
 */
export async function verifyProposalSignature(
  proposal: ProposedOutcome,
  domainSeparator: Hex,
  signature: Hex
): Promise<`0x${string}`> {
  // Compute the digest that was signed
  const digest = computeProposalDigest(proposal, domainSeparator);

  // Recover the signer address from the signature
  const signer = await recoverAddress({
    hash: digest,
    signature: normalizeSignature(signature),
  });

  return signer;
}

/**
 * Checks if a signature is valid for a given proposal and expected signer
 * 
 * @param proposal - ProposedOutcome that was signed
 * @param domainSeparator - EIP-712 domain separator
 * @param signature - 65-byte signature to verify
 * @param expectedSigner - Expected signer address
 * @returns True if signature is valid
 */
export async function isValidProposalSignature(
  proposal: ProposedOutcome,
  domainSeparator: Hex,
  signature: Hex,
  expectedSigner: `0x${string}`
): Promise<boolean> {
  try {
    const recoveredSigner = await verifyProposalSignature(proposal, domainSeparator, signature);
    return recoveredSigner.toLowerCase() === expectedSigner.toLowerCase();
  } catch {
    return false;
  }
}

// ============ Validation ============

/**
 * Validates that a signature is properly formatted
 * 
 * @param signature - Signature to validate
 * @returns True if signature is valid format
 */
export function isValidSignature(signature: string): signature is Hex {
  // Must be hex string
  if (!isHex(signature)) {
    return false;
  }

  // Must be exactly 65 bytes (0x + 130 hex chars)
  if (signature.length !== 132) {
    return false;
  }

  return true;
}

/**
 * Computes the keccak256 hash of a signature (for replay protection)
 * 
 * This matches the contract's signature hashing for the usedSignatures mapping.
 * 
 * @param signature - 65-byte signature
 * @returns Keccak256 hash of the signature
 */
export function hashSignature(signature: Hex): Hex {
  if (!isValidSignature(signature)) {
    throw new Error('Invalid signature format');
  }

  // The contract uses keccak256(signature) for replay protection
  const { keccak256 } = require('viem');
  return keccak256(signature);
}

// ============ Utilities ============

/**
 * Creates a private key account from a hex private key
 * 
 * @param privateKey - Private key as hex string (0x-prefixed, 32 bytes)
 * @returns Viem PrivateKeyAccount
 */
export function createAccountFromPrivateKey(privateKey: Hex): PrivateKeyAccount {
  return privateKeyToAccount(privateKey);
}

/**
 * Formats a signature for display or logging
 * 
 * @param signature - 65-byte signature
 * @returns Formatted signature string with components
 */
export function formatSignature(signature: Hex): string {
  const { r, s, v } = splitSignature(signature);
  return `Signature {\n  r: ${r}\n  s: ${s}\n  v: ${v}\n}`;
}
