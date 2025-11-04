import { useEffect } from "react";
import { useLocation } from "wouter";
import { useAccount, useSignMessage } from "wagmi";
import { SiweMessage } from "siwe";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";
import { Loader2 } from "lucide-react";
import { useState } from "react";

export default function LoginPage() {
  const [, setLocation] = useLocation();
  const { toast } = useToast();
  const { address, isConnected, chainId } = useAccount();
  const { signMessageAsync } = useSignMessage();
  const [isLoading, setIsLoading] = useState(false);

  // If already connected and authenticated, redirect to home
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await fetch("/api/auth/me");
        if (response.ok) {
          setLocation("/");
        }
      } catch (error) {
        // Not authenticated, stay on login page
      }
    };
    checkAuth();
  }, [setLocation]);

  const handleSignIn = async () => {
    if (!address || !chainId) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet first",
        variant: "destructive",
      });
      return;
    }

    setIsLoading(true);

    try {
      // Step 1: Get nonce from server
      const nonceResponse = await fetch("/api/auth/nonce");
      if (!nonceResponse.ok) {
        throw new Error("Failed to get nonce");
      }
      const { nonce } = await nonceResponse.json();

      // Step 2: Create SIWE message
      const message = new SiweMessage({
        domain: window.location.host,
        address,
        statement: "Sign in to DegenArena with Ethereum",
        uri: window.location.origin,
        version: "1",
        chainId,
        nonce,
      });

      const messageString = message.prepareMessage();

      // Step 3: Sign message with wallet
      const signature = await signMessageAsync({
        message: messageString,
      });

      // Step 4: Verify signature with server
      const verifyResponse = await fetch("/api/auth/verify", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          message: messageString,
          signature,
        }),
      });

      if (!verifyResponse.ok) {
        throw new Error("Signature verification failed");
      }

      const { success } = await verifyResponse.json();

      if (success) {
        toast({
          title: "Success!",
          description: "You have been signed in",
        });
        setLocation("/");
      } else {
        throw new Error("Authentication failed");
      }
    } catch (error: any) {
      console.error("Sign-in error:", error);
      toast({
        title: "Sign-in failed",
        description: error.message || "Failed to sign in with wallet",
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-background to-muted p-4">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl font-bold text-center">
            Welcome to DegenArena
          </CardTitle>
          <CardDescription className="text-center">
            Connect your wallet and sign in to start trading prediction markets
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex justify-center">
            <ConnectButton />
          </div>

          {isConnected && address && (
            <div className="space-y-4">
              <div className="text-sm text-muted-foreground text-center">
                Connected as {address.substring(0, 6)}...{address.substring(38)}
              </div>
              <Button
                onClick={handleSignIn}
                disabled={isLoading}
                className="w-full"
                size="lg"
              >
                {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Sign In with Ethereum
              </Button>
            </div>
          )}

          <div className="text-xs text-center text-muted-foreground">
            By signing in, you agree to our Terms of Service and Privacy Policy
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
