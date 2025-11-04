// Horizon Wallet Dashboard
// Reference: design_guidelines.md - Wallet interface with deposit/withdrawal
import { useState, useEffect } from "react";
import { useQuery, useMutation } from "@tanstack/react-query";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  Wallet as WalletIcon, 
  ArrowDownToLine, 
  ArrowUpFromLine, 
  Copy, 
  ExternalLink,
  AlertCircle,
  CheckCircle
} from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { useAuth } from "@/hooks/useAuth";
import { apiRequest, queryClient } from "@/lib/queryClient";
import { isUnauthorizedError } from "@/lib/authUtils";
import type { Wallet, Transaction } from "@shared/schema";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";

export default function WalletDashboard() {
  const { toast } = useToast();
  const { user, isAuthenticated, isLoading } = useAuth();
  const [depositAmount, setDepositAmount] = useState("");
  const [withdrawAmount, setWithdrawAmount] = useState("");
  const [withdrawAddress, setWithdrawAddress] = useState("");
  const [copied, setCopied] = useState(false);

  // Redirect if not authenticated
  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      toast({
        title: "Unauthorized",
        description: "You are logged out. Logging in again...",
        variant: "destructive",
      });
      setTimeout(() => {
        window.location.href = "/api/login";
      }, 500);
    }
  }, [isAuthenticated, isLoading, toast]);

  // Fetch wallet data
  const { data: wallet, isLoading: walletLoading } = useQuery<Wallet>({
    queryKey: ["/api/wallet"],
    enabled: isAuthenticated,
  });

  // Fetch transactions
  const { data: transactions, isLoading: txLoading } = useQuery<Transaction[]>({
    queryKey: ["/api/wallet/transactions"],
    enabled: isAuthenticated,
  });

  // Fetch BNB price from CoinGecko
  const { data: cryptoPrices } = useQuery<{ usd: number; usd_24h_change: number }>({
    queryKey: ["/api/crypto/prices"],
    refetchInterval: 30000, // Refresh every 30 seconds
  });

  // Deposit mutation
  const depositMutation = useMutation({
    mutationFn: async (amount: string) => {
      return await apiRequest("POST", "/api/wallet/deposit", { amount });
    },
    onSuccess: () => {
      toast({
        title: "Deposit Successful",
        description: `Added ${depositAmount} BNB to your wallet`,
      });
      setDepositAmount("");
      queryClient.invalidateQueries({ queryKey: ["/api/wallet"] });
      queryClient.invalidateQueries({ queryKey: ["/api/wallet/transactions"] });
    },
    onError: (error: any) => {
      if (isUnauthorizedError(error)) {
        toast({
          title: "Unauthorized",
          description: "You are logged out. Logging in again...",
          variant: "destructive",
        });
        setTimeout(() => {
          window.location.href = "/api/login";
        }, 500);
        return;
      }
      toast({
        title: "Deposit Failed",
        description: error.message || "Please try again",
        variant: "destructive",
      });
    },
  });

  // Withdraw mutation
  const withdrawMutation = useMutation({
    mutationFn: async ({ amount, address }: { amount: string; address: string }) => {
      return await apiRequest("POST", "/api/wallet/withdraw", { amount, toAddress: address });
    },
    onSuccess: () => {
      toast({
        title: "Withdrawal Initiated",
        description: `Withdrawing ${withdrawAmount} BNB`,
      });
      setWithdrawAmount("");
      setWithdrawAddress("");
      queryClient.invalidateQueries({ queryKey: ["/api/wallet"] });
      queryClient.invalidateQueries({ queryKey: ["/api/wallet/transactions"] });
    },
    onError: (error: any) => {
      if (isUnauthorizedError(error)) {
        toast({
          title: "Unauthorized",
          description: "You are logged out. Logging in again...",
          variant: "destructive",
        });
        setTimeout(() => {
          window.location.href = "/api/login";
        }, 500);
        return;
      }
      toast({
        title: "Withdrawal Failed",
        description: error.message || "Please try again",
        variant: "destructive",
      });
    },
  });

  const handleCopyAddress = () => {
    if (wallet?.bnbAddress) {
      navigator.clipboard.writeText(wallet.bnbAddress);
      setCopied(true);
      toast({
        title: "Address Copied",
        description: "BNB address copied to clipboard",
      });
      setTimeout(() => setCopied(false), 2000);
    }
  };

  const handleDeposit = (e: React.FormEvent) => {
    e.preventDefault();
    if (depositAmount && parseFloat(depositAmount) > 0) {
      depositMutation.mutate(depositAmount);
    }
  };

  const handleWithdraw = (e: React.FormEvent) => {
    e.preventDefault();
    if (withdrawAmount && parseFloat(withdrawAmount) > 0 && withdrawAddress) {
      withdrawMutation.mutate({ amount: withdrawAmount, address: withdrawAddress });
    }
  };

  if (walletLoading) {
    return (
      <div className="flex flex-col h-full">
        <LiveBettingFeed />
        <div className="flex-1 overflow-auto relative">
          <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
            <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
          </div>
          <div className="container mx-auto p-6 max-w-5xl space-y-6 relative z-10">
            <Skeleton className="h-12 w-64" />
            <Skeleton className="h-48" />
            <div className="grid md:grid-cols-2 gap-6">
              <Skeleton className="h-96" />
              <Skeleton className="h-96" />
            </div>
          </div>
        </div>
      </div>
    );
  }

  const balance = parseFloat(wallet?.balance || "0");
  const bnbPrice = cryptoPrices?.usd || 600; // Fallback to $600 if API fails
  const usdValue = balance * bnbPrice;

  return (
    <div className="flex flex-col h-full">
      {/* Live Betting Feed - Horizontal Ticker */}
      <LiveBettingFeed />
      
      <div className="flex-1 overflow-auto relative">
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
      </div>
      <div className="container mx-auto p-6 max-w-5xl space-y-6 relative z-10">
        <h1 className="text-3xl font-bold text-foreground font-sohne">
          My <span className="text-primary">Wallet</span>
        </h1>

      {/* Balance Card */}
      <Card className="bg-gradient-to-br from-card to-accent/5 border-card-border">
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-foreground">
            <WalletIcon className="h-6 w-6 text-primary" />
            Wallet Balance
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <p className="text-5xl font-mono font-bold text-primary">
              {balance.toFixed(4)} <span className="text-2xl text-accent">BNB</span>
            </p>
            <p className="text-sm text-muted-foreground mt-1">
              ≈ ${usdValue.toFixed(2)} USD
              {cryptoPrices?.usd_24h_change && (
                <span className={`ml-2 text-xs ${cryptoPrices.usd_24h_change >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                  {cryptoPrices.usd_24h_change >= 0 ? '+' : ''}{cryptoPrices.usd_24h_change.toFixed(2)}%
                </span>
              )}
            </p>
          </div>

          <div className="bg-muted rounded-lg p-3">
            <Label className="text-xs text-muted-foreground">Your BNB Address</Label>
            <div className="flex items-center gap-2 mt-1">
              <code className="flex-1 font-mono text-sm text-foreground truncate">
                {wallet?.bnbAddress}
              </code>
              <Button
                size="sm"
                variant="outline"
                onClick={handleCopyAddress}
                className="border-accent text-accent hover:bg-accent hover:text-accent-foreground"
                data-testid="button-copy-address"
              >
                {copied ? <CheckCircle className="h-4 w-4" /> : <Copy className="h-4 w-4" />}
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Deposit & Withdraw */}
      <div className="grid md:grid-cols-2 gap-6">
        {/* Deposit */}
        <Card className="bg-card border-card-border">
          <CardHeader className="border-b border-accent">
            <CardTitle className="flex items-center gap-2 text-foreground">
              <ArrowDownToLine className="h-5 w-5 text-primary" />
              Deposit
            </CardTitle>
            <CardDescription>Add BNB to your wallet</CardDescription>
          </CardHeader>
          <CardContent className="pt-6">
            <form onSubmit={handleDeposit} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="deposit-amount">Amount (BNB)</Label>
                <Input
                  id="deposit-amount"
                  type="number"
                  step="0.0001"
                  min="0"
                  placeholder="0.0000"
                  value={depositAmount}
                  onChange={(e) => setDepositAmount(e.target.value)}
                  className="font-mono bg-input border-accent"
                  data-testid="input-deposit-amount"
                />
              </div>
              <Alert className="border-accent bg-accent/10">
                <AlertCircle className="h-4 w-4 text-accent" />
                <AlertDescription className="text-xs">
                  In production, you would send BNB from an exchange to your deposit address.
                  For MVP, this simulates the deposit.
                </AlertDescription>
              </Alert>
              <Button
                type="submit"
                className="w-full bg-primary hover:bg-primary/90 text-primary-foreground font-bold"
                disabled={depositMutation.isPending || !depositAmount}
                data-testid="button-deposit"
              >
                {depositMutation.isPending ? "Processing..." : "Deposit BNB"}
              </Button>
            </form>
          </CardContent>
        </Card>

        {/* Withdraw */}
        <Card className="bg-card border-card-border">
          <CardHeader className="border-b border-accent">
            <CardTitle className="flex items-center gap-2 text-foreground">
              <ArrowUpFromLine className="h-5 w-5 text-accent" />
              Withdraw
            </CardTitle>
            <CardDescription>Send BNB to external address</CardDescription>
          </CardHeader>
          <CardContent className="pt-6">
            <form onSubmit={handleWithdraw} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="withdraw-amount">Amount (BNB)</Label>
                <Input
                  id="withdraw-amount"
                  type="number"
                  step="0.0001"
                  min="0"
                  max={balance}
                  placeholder="0.0000"
                  value={withdrawAmount}
                  onChange={(e) => setWithdrawAmount(e.target.value)}
                  className="font-mono bg-input border-accent"
                  data-testid="input-withdraw-amount"
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="withdraw-address">To Address</Label>
                <Input
                  id="withdraw-address"
                  type="text"
                  placeholder="0x..."
                  value={withdrawAddress}
                  onChange={(e) => setWithdrawAddress(e.target.value)}
                  className="font-mono bg-input border-accent"
                  data-testid="input-withdraw-address"
                />
              </div>
              <Button
                type="submit"
                variant="outline"
                className="w-full border-accent text-accent hover:bg-accent hover:text-accent-foreground font-bold"
                disabled={withdrawMutation.isPending || !withdrawAmount || !withdrawAddress}
                data-testid="button-withdraw"
              >
                {withdrawMutation.isPending ? "Processing..." : "Withdraw BNB"}
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>

      {/* Transaction History */}
      <Card className="bg-card border-card-border">
        <CardHeader className="border-b border-accent">
          <CardTitle className="text-foreground">Transaction History</CardTitle>
        </CardHeader>
        <CardContent className="p-0">
          {txLoading ? (
            <div className="p-6 space-y-3">
              {[1, 2, 3].map((i) => (
                <Skeleton key={i} className="h-16" />
              ))}
            </div>
          ) : transactions && transactions.length > 0 ? (
            <div className="divide-y divide-border">
              {transactions.slice(0, 10).map((tx) => (
                <div
                  key={tx.id}
                  className="p-4 hover-elevate flex items-center justify-between"
                  data-testid={`tx-${tx.id}`}
                >
                  <div className="flex items-center gap-3">
                    <Badge
                      variant={tx.type.includes('deposit') ? 'default' : 'outline'}
                      className={
                        tx.type.includes('deposit')
                          ? 'bg-primary text-primary-foreground'
                          : 'border-accent text-accent'
                      }
                    >
                      {tx.type}
                    </Badge>
                    <div>
                      <p className="font-mono font-medium text-foreground">
                        {parseFloat(tx.amount).toFixed(4)} BNB
                      </p>
                      <p className="text-xs text-muted-foreground">
                        {new Date(tx.createdAt!).toLocaleString()}
                      </p>
                    </div>
                  </div>
                  {tx.txHash && (
                    <Button
                      size="sm"
                      variant="ghost"
                      className="text-accent hover:text-accent-foreground"
                      data-testid={`button-view-tx-${tx.id}`}
                    >
                      <ExternalLink className="h-4 w-4" />
                    </Button>
                  )}
                </div>
              ))}
            </div>
          ) : (
            <div className="p-12 text-center text-muted-foreground">
              No transactions yet
            </div>
          )}
        </CardContent>
      </Card>
        </div>
      </div>
    </div>
  );
}
