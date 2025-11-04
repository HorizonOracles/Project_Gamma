import { useAccount } from "wagmi";

/**
 * Hook to check if the connected wallet is the admin address
 * Uses fixed admin address from environment variable
 */
export function useAdmin() {
  const { address, isConnected } = useAccount();

  // Get admin address from environment
  const adminAddress = import.meta.env.VITE_ADMIN_WALLET_ADDRESS?.toLowerCase();

  // Check if current address matches admin address
  const isAdmin = isConnected && !!address && !!adminAddress && 
                  address.toLowerCase() === adminAddress;

  return {
    isAdmin,
    isConnected,
    address,
    adminAddress,
  };
}
