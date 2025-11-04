/**
 * Create Market Page - Admin-only market creation interface
 * Allows administrators to create new prediction markets on-chain
 */

import { useLocation } from "wouter";
import { useAccount } from "wagmi";
import { useAdmin } from "@/hooks/useAdmin";
import { CreateMarketForm } from "@/components/CreateMarketForm";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { AlertCircle, ArrowLeft } from "lucide-react";

export default function CreateMarket() {
  const [, setLocation] = useLocation();
  const { isConnected } = useAccount();
  const { isAdmin } = useAdmin();

  // Redirect if not admin
  if (!isConnected) {
    return (
      <div className="container mx-auto px-4 py-8 max-w-2xl">
        <Alert className="bg-yellow-500/10 border-yellow-500/20">
          <AlertCircle className="h-4 w-4 text-yellow-400" />
          <AlertDescription className="text-yellow-100">
            <p className="font-semibold mb-2">Wallet Not Connected</p>
            <p>Please connect your wallet to access this page.</p>
          </AlertDescription>
        </Alert>
        <Button
          onClick={() => setLocation("/blockchain-markets")}
          variant="outline"
          className="mt-4"
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Markets
        </Button>
      </div>
    );
  }

  if (!isAdmin) {
    return (
      <div className="container mx-auto px-4 py-8 max-w-2xl">
        <Alert className="bg-red-500/10 border-red-500/20">
          <AlertCircle className="h-4 w-4 text-red-400" />
          <AlertDescription className="text-red-100">
            <p className="font-semibold mb-2">Access Denied</p>
            <p>
              Only administrators can create markets. Your wallet address is not
              authorized.
            </p>
          </AlertDescription>
        </Alert>
        <Button
          onClick={() => setLocation("/blockchain-markets")}
          variant="outline"
          className="mt-4"
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Markets
        </Button>
      </div>
    );
  }

  // Admin access granted
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-6">
        <Button
          onClick={() => setLocation("/blockchain-markets")}
          variant="ghost"
          className="mb-4"
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Markets
        </Button>
      </div>

      <CreateMarketForm
        onSuccess={(marketId) => {
          // Optionally redirect to market page after creation
          setTimeout(() => {
            setLocation("/blockchain-markets");
          }, 3000);
        }}
      />
    </div>
  );
}
