/**
 * Hook to create a new market
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAccount, usePublicClient, useWalletClient, useChainId } from 'wagmi';
import { useGammaConfig } from '../../components/GammaProvider';
import { CreateMarketParams } from '../../types';
import { DEFAULT_CONTRACTS, MARKET_FACTORY_ABI, ERC20_ABI } from '../../constants';
import { decodeEventLog } from 'viem';

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
        throw new Error('Public client not available');
      }

      if (!address) {
        throw new Error('Wallet not connected');
      }

      const marketFactoryAddress =
        config.marketFactoryAddress ||
        DEFAULT_CONTRACTS[chainId as keyof typeof DEFAULT_CONTRACTS]?.marketFactory;

      if (!marketFactoryAddress) {
        throw new Error('MarketFactory address not configured');
      }

      // Check and approve HORIZON token if creatorStake is required
      if (params.creatorStake > 0n && config.horizonTokenAddress) {
        const currentAllowance = await publicClient.readContract({
          address: config.horizonTokenAddress,
          abi: ERC20_ABI,
          functionName: 'allowance',
          args: [address, marketFactoryAddress],
        }) as bigint;

        if (currentAllowance < params.creatorStake && config.horizonTokenAddress && walletClient) {
          // Approve transaction
          const approveHash = await walletClient.writeContract({
            address: config.horizonTokenAddress,
            abi: ERC20_ABI,
            functionName: 'approve',
            args: [marketFactoryAddress, params.creatorStake],
          });
          await publicClient.waitForTransactionReceipt({ hash: approveHash });
        }
      }

      // Create market
      if (!walletClient) {
        throw new Error('Wallet not connected');
      }

      const txHash = await walletClient.writeContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'createMarket',
        args: [{
          collateralToken: params.collateralToken,
          closeTime: params.endTime,
          category: params.category,
          metadataURI: params.metadataURI,
          creatorStake: params.creatorStake,
        }],
      });

      // Wait for transaction and extract market ID from event
      const receipt = await publicClient.waitForTransactionReceipt({ hash: txHash });
      
      // Extract market ID from MarketCreated event
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
        return (decoded.args as any).marketId as bigint;
      }

      // Fallback: get nextMarketId - 1
      const nextMarketId = await publicClient.readContract({
        address: marketFactoryAddress,
        abi: MARKET_FACTORY_ABI,
        functionName: 'nextMarketId',
      }) as bigint;

      return nextMarketId - 1n;
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

