/**
 * AIOracleAdapter contract interaction layer
 */

import { Address } from 'viem';
import { PublicClient, WalletClient } from 'viem';
import { ContractError } from '../types';
import { AI_ORACLE_ADAPTER_ABI, ERC20_ABI } from '../constants';
import { keccak256 } from 'viem';

/**
 * ProposedOutcome struct as expected by the contract
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
 * Parameters for proposing an AI resolution
 */
export interface ProposeAIParams {
  proposal: ProposedOutcome;
  signature: `0x${string}`;
  bondAmount: bigint;
  evidenceURIs: string[];
}

export class AIOracleAdapter {
  private client: PublicClient;
  private walletClient?: WalletClient;
  private address: Address;

  constructor(
    client: PublicClient,
    address: Address,
    walletClient?: WalletClient
  ) {
    this.client = client;
    this.address = address;
    this.walletClient = walletClient;
  }

  /**
   * Submit an AI-signed resolution proposal
   */
  async proposeAI(params: ProposeAIParams): Promise<string> {
    if (!this.walletClient) {
      throw new ContractError('Wallet client required for proposing AI resolution', this.address);
    }

    try {
      const [account] = await this.walletClient.getAddresses();
      if (!account) {
        throw new ContractError('No account found in wallet client', this.address);
      }

      // Ensure bond token is approved
      // Note: bond token address should be retrieved from ResolutionModule
      // For now, assuming it's HORIZON token - this should be configurable
      
      const hash = await this.walletClient.writeContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'proposeAI',
        args: [
          {
            marketId: params.proposal.marketId,
            outcomeId: params.proposal.outcomeId,
            closeTime: params.proposal.closeTime,
            evidenceHash: params.proposal.evidenceHash,
            notBefore: params.proposal.notBefore,
            deadline: params.proposal.deadline,
          },
          params.signature,
          params.bondAmount,
          params.evidenceURIs,
        ] as const,
        account,
        chain: this.walletClient.chain,
      });

      await this.client.waitForTransactionReceipt({ hash });
      return hash;
    } catch (error) {
      throw new ContractError(
        `Failed to propose AI resolution: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get proposal hash for signing
   */
  async getProposalHash(proposal: ProposedOutcome): Promise<`0x${string}`> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'getProposalHash',
        args: [{
          marketId: proposal.marketId,
          outcomeId: proposal.outcomeId,
          closeTime: proposal.closeTime,
          evidenceHash: proposal.evidenceHash,
          notBefore: proposal.notBefore,
          deadline: proposal.deadline,
        }],
      });

      return result as `0x${string}`;
    } catch (error) {
      throw new ContractError(
        `Failed to get proposal hash: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Check if a signature has been used
   */
  async isSignatureUsed(signature: `0x${string}`): Promise<boolean> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'isSignatureUsed',
        args: [signature],
      });

      return result as boolean;
    } catch (error) {
      throw new ContractError(
        `Failed to check signature usage: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Hash evidence URIs (client-side helper)
   * Note: This can also be called on-chain using hashEvidence
   */
  async hashEvidence(evidenceURIs: string[]): Promise<`0x${string}`> {
    try {
      // Try on-chain first (pure function)
      const result = await this.client.readContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'hashEvidence',
        args: [evidenceURIs],
      });

      return result as `0x${string}`;
    } catch (error) {
      // Fallback: Contract uses abi.encode which includes array length
      // For client-side, we'll use keccak256 of abi.encode format
      // Note: This should match the contract's _hashEvidence implementation
      // The contract uses: keccak256(abi.encode(evidenceURIs))
      // We'll rely on on-chain call for accuracy, but provide basic fallback
      // In production, always use the on-chain method for accuracy
      const { encodeAbiParameters } = await import('viem');
      const encoded = encodeAbiParameters(
        [{ type: 'string[]' }],
        [evidenceURIs]
      );
      return keccak256(encoded);
    }
  }

  /**
   * Check if an address is an allowed signer
   */
  async isAllowedSigner(signer: Address): Promise<boolean> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'allowedSigners',
        args: [signer],
      });

      return result as boolean;
    } catch (error) {
      throw new ContractError(
        `Failed to check allowed signer: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }

  /**
   * Get domain separator (for EIP-712 signing)
   */
  async getDomainSeparator(): Promise<`0x${string}`> {
    try {
      const result = await this.client.readContract({
        address: this.address,
        abi: AI_ORACLE_ADAPTER_ABI,
        functionName: 'DOMAIN_SEPARATOR',
      });

      return result as `0x${string}`;
    } catch (error) {
      throw new ContractError(
        `Failed to get domain separator: ${error instanceof Error ? error.message : 'Unknown error'}`,
        this.address,
        error
      );
    }
  }
}

