/**
 * Hook to create a new market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { Address, PublicClient, WalletClient, TransactionReceipt } from 'viem';
import { useGammaConfig } from '../../components/GammaProvider';
import { CreateMarketParams, ContractError } from '../../types';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, ERC20_ABI } from '../../constants';
import { decodeEventLog, maxUint256 } from 'viem';

/**
 * Transaction timeout constants
 */
const TRANSACTION_TIMEOUT_MS = 120_000; // 2 minutes
const STATE_UPDATE_DELAY_MS = 1000; // 1 second

/**
 * Helper function to get contract address with fallback
 */
function getContractAddress(
  configAddress: Address | undefined,
  chainId: number,
  contractKey: 'marketFactory' | 'horizonToken'
): Address {
  if (configAddress) {
    return configAddress;
  }
  
  const contracts = DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS];
  const address = contracts?.[contractKey];
  
  if (!address || address === '0x0000000000000000000000000000000000000000') {
    throw new ContractError(
      `${contractKey} address not configured for chain ${chainId}`,
      address || '0x0000000000000000000000000000000000000000'
    );
  }
  
  return address;
}

/**
 * Helper function to ensure token allowance
 */
async function ensureTokenAllowance(
  publicClient: PublicClient,
  walletClient: WalletClient,
  tokenAddress: Address,
  spenderAddress: Address,
  requiredAmount: bigint,
  userAddress: Address
): Promise<void> {
  // Check balance
  const balance = await publicClient.readContract({
    address: tokenAddress,
    abi: ERC20_ABI,
    functionName: 'balanceOf',
    args: [userAddress],
  }) as bigint;

  if (balance < requiredAmount) {
    throw new ContractError(
      `Insufficient token balance. Required: ${requiredAmount.toString()}, Have: ${balance.toString()}`,
      tokenAddress
    );
  }

  // Check current allowance
  const currentAllowance = await publicClient.readContract({
    address: tokenAddress,
    abi: ERC20_ABI,
    functionName: 'allowance',
    args: [userAddress, spenderAddress],
  }) as bigint;

  // Approve if needed
  if (currentAllowance < requiredAmount) {
    // Get account from wallet client
    const [account] = await walletClient.getAddresses();
    if (!account) {
      throw new ContractError('No account found in wallet client', tokenAddress);
    }

    const approveHash = await walletClient.writeContract({
      address: tokenAddress,
      abi: ERC20_ABI,
      functionName: 'approve',
      args: [spenderAddress, maxUint256],
      account,
      chain: walletClient.chain,
    });

    // Wait for approval transaction
    const approveReceipt = await publicClient.waitForTransactionReceipt({
      hash: approveHash,
      timeout: TRANSACTION_TIMEOUT_MS,
    });

    if (approveReceipt.status !== 'success') {
      throw new ContractError(
        `Token approval transaction failed. Status: ${approveReceipt.status}`,
        tokenAddress
      );
    }

    // Verify allowance after approval
    const newAllowance = await publicClient.readContract({
      address: tokenAddress,
      abi: ERC20_ABI,
      functionName: 'allowance',
      args: [userAddress, spenderAddress],
    }) as bigint;

    if (newAllowance < requiredAmount) {
      throw new ContractError(
        `Insufficient allowance after approval. Required: ${requiredAmount.toString()}, Got: ${newAllowance.toString()}`,
        tokenAddress
      );
    }

    // Small delay to ensure state is fully updated
    await new Promise(resolve => setTimeout(resolve, STATE_UPDATE_DELAY_MS));
  }
}

/**
 * Helper function to extract market ID from transaction receipt
 */
async function extractMarketIdFromReceipt(
  receipt: TransactionReceipt,
  publicClient: PublicClient,
  marketFactoryAddress: Address
): Promise<bigint> {
  // Try to find MarketCreated event
  const marketCreatedEvent = receipt.logs.find((log) => {
    try {
      const decoded = decodeEventLog({
        abi: MARKET_FACTORY_ABI,
        data: log.data,
        topics: log.topics,
      });
      return decoded.eventName === 'MarketCreated';
    } catch {
      return false;
    }
  });

  if (marketCreatedEvent) {
    const decoded = decodeEventLog({
      abi: MARKET_FACTORY_ABI,
      data: marketCreatedEvent.data,
      topics: marketCreatedEvent.topics,
    });
    
    if (decoded.eventName === 'MarketCreated' && 'marketId' in decoded.args) {
      return decoded.args.marketId as bigint;
    }
  }

  // Fallback: get nextMarketId - 1
  // This is less reliable but provides a fallback if event parsing fails
  try {
    const nextMarketId = await publicClient.readContract({
      address: marketFactoryAddress,
      abi: MARKET_FACTORY_ABI,
      functionName: 'nextMarketId',
    }) as bigint;

    return nextMarketId - 1n;
  } catch (error) {
    throw new ContractError(
      'MarketCreated event not found in transaction receipt and fallback method failed. Unable to determine market ID.',
      marketFactoryAddress,
      error
    );
  }
}

/**
 * Hook to create a new market
 * 
 * @example
 * ```tsx
 * const { write: createMarket, isLoading } = useCreateMarket();
 * 
 * createMarket({
 *   question: 'Will BTC hit $100k?',
 *   endTime: BigInt(Math.floor(Date.now() / 1000) + 86400),
 *   collateralToken: '0x...',
 *   category: 'crypto',
 *   metadataURI: 'ipfs://...',
 *   creatorStake: parseUnits('1000', 18),
 * });
 * ```
 */
export function useCreateMarket() {
  const config = useGammaConfig();
  const { address } = useAccount();
  const publicClient = usePublicClient();
  const { data: walletClient } = useWalletClient();
  const chainId = useChainId();
  const queryClient = useQueryClient();

  const createMarketMutation = useMutation({
    mutationFn: async (params: CreateMarketParams): Promise<bigint> => {
      if (!publicClient) {
        throw new ContractError('Public client not available');
      }

      if (!address) {
        throw new ContractError('Wallet not connected');
      }

      if (!walletClient) {
        throw new ContractError('Wallet client not available');
      }

      const marketFactoryAddress = getContractAddress(
        config.marketFactoryAddress,
        chainId,
        'marketFactory'
      );

      // Handle token approval if creator stake is required
      if (params.creatorStake > 0n) {
        const horizonTokenAddress = getContractAddress(
          config.horizonTokenAddress,
          chainId,
          'horizonToken'
        );

        await ensureTokenAllowance(
          publicClient,
          walletClient,
          horizonTokenAddress,
          marketFactoryAddress,
          params.creatorStake,
          address
        );
      }

      // Create market with new MarketParams struct
      const txHash = await walletClient.writeContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'createMarket',
        args: [{
          marketType: params.marketType ?? 0, // Default to Binary (0)
          collateralToken: params.collateralToken,
          closeTime: params.endTime,
          category: params.category,
          metadataURI: params.metadataURI,
          creatorStake: params.creatorStake,
          outcomeCount: params.outcomeCount ?? 2, // Default to 2 outcomes
          liquidityParameter: params.liquidityParameter ?? 0n, // Default to 0 (not used for Binary)
        }],
      });

      // Wait for transaction and extract market ID from event
      const receipt = await publicClient.waitForTransactionReceipt({ hash: txHash });
      
      return await extractMarketIdFromReceipt(receipt, publicClient, marketFactoryAddress);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['markets'] });
    },
  });

  return {
    write: (params: CreateMarketParams) => createMarketMutation.mutate(params),
    isLoading: createMarketMutation.isPending,
    isSuccess: createMarketMutation.isSuccess,
    isError: createMarketMutation.isError,
    data: createMarketMutation.data,
    hash: createMarketMutation.data, // Alias for backwards compatibility
    error: createMarketMutation.error,
  };
}

