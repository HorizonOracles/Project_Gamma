// SIWE Wallet-based Authentication Hook
import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import { useAccount } from "wagmi";
import type { User } from "@shared/schema";

export function useAuth() {
  const { isConnected } = useAccount();
  
  const { data: user, isLoading: queryLoading, refetch } = useQuery<User>({
    queryKey: ["/api/auth/user"],
    retry: false,
    enabled: isConnected, // Only fetch if wallet is connected
  });

  // Add minimum loading time of 2 seconds to show loading animation
  const [isLoading, setIsLoading] = useState(true);
  const [minLoadingComplete, setMinLoadingComplete] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      setMinLoadingComplete(true);
    }, 2000);

    return () => clearTimeout(timer);
  }, []);

  useEffect(() => {
    if (!queryLoading && minLoadingComplete) {
      setIsLoading(false);
    }
  }, [queryLoading, minLoadingComplete]);

  // Logout function
  const logout = async () => {
    try {
      await fetch("/api/auth/logout", {
        method: "POST",
      });
      // Refetch to update auth state
      await refetch();
    } catch (error) {
      console.error("Logout error:", error);
    }
  };

  return {
    user,
    isLoading,
    isAuthenticated: !!user && isConnected,
    logout,
  };
}
