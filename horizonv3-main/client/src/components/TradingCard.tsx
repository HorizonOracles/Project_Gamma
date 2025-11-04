/**
 * TradingCard - Fully Decentralized Market Trading Component
 * Uses Project Gamma SDK for on-chain market interactions
 * All data fetched from blockchain/IPFS - no centralized database dependencies
 */

import { useState } from "react";
import { useAccount } from "wagmi";
import { parseUnits, formatUnits } from "viem";
import {
  useMarket,
  usePrices,
  useBuy,
  useSell,
  useQuote,
  useOutcomeBalance,
  useAllowance,
  useApprove,
  useBalance,
  useIPFSMetadata,
} from "@project-gamma/sdk";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Loader2, TrendingUp, TrendingDown, CheckCircle, AlertCircle } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { useGammaConfig } from "@project-gamma/sdk";

interface TradingCardProps {
  marketId: number;
  title?: string; // Optional override for market question
  description?: string; // Optional override for category/description
}

export function TradingCard({
  marketId,
  title,
  description,
}: TradingCardProps) {
  const { toast } = useToast();
  const { address } = useAccount();
  const config = useGammaConfig();

  // Market data from blockchain
  const { data: market, isLoading: marketLoading } = useMarket(marketId);
  const { data: prices, isLoading: pricesLoading } = usePrices(marketId);
  
  // Fetch IPFS metadata for additional details (optional)
  const { data: ipfsMetadata, isLoading: ipfsLoading } = useIPFSMetadata(market?.metadataURI);

  // User balances
  const { data: yesBalance } = useOutcomeBalance(marketId, 0);
  const { data: noBalance } = useOutcomeBalance(marketId, 1);
  const { data: collateralBalance } = useBalance(
    market?.collateralToken as `0x${string}`
  );

  // Token approval
  const { data: allowance } = useAllowance({
    tokenAddress: market?.collateralToken as `0x${string}`,
    spenderAddress: market?.marketAddress as `0x${string}`,
  });

  const { write: approve, isLoading: isApproving } = useApprove();

  // Trading state
  const [tradeMode, setTradeMode] = useState<"buy" | "sell">("buy");
  const [selectedOutcome, setSelectedOutcome] = useState<"YES" | "NO">("YES");
  const [amount, setAmount] = useState("");
  const { write: buy, isLoading: isBuying, isSuccess: buySuccess, hash: buyHash } = useBuy(marketId);
  const { write: sell, isLoading: isSelling, isSuccess: sellSuccess, hash: sellHash } = useSell(marketId);

  // Get quote for sell (to show expected proceeds)
  const sellQuoteParams = tradeMode === "sell" && amount && parseFloat(amount) > 0
    ? {
        marketId,
        outcomeId: selectedOutcome === "YES" ? 0 : 1,
        amount: parseUnits(amount, 18),
        isBuy: false,
      }
    : undefined;
  const { data: sellQuote } = useQuote(sellQuoteParams);

  const isSuccess = buySuccess || sellSuccess;
  const hash = buyHash || sellHash;

  // Check if approval is needed
  const needsApproval = () => {
    if (!amount || !allowance || !market?.collateralToken) return false;
    try {
      const amountBigInt = parseUnits(amount, 18); // Assuming 18 decimals for collateral
      return allowance < amountBigInt;
    } catch {
      return false;
    }
  };

  const handleApprove = async () => {
    if (!market?.collateralToken || !market?.marketAddress) {
      toast({
        title: "Error",
        description: "Market data not loaded",
        variant: "destructive",
      });
      return;
    }

    try {
      const amountToApprove = parseUnits(amount || "0", 18);
      approve({
        tokenAddress: market.collateralToken as `0x${string}`,
        spenderAddress: market.marketAddress as `0x${string}`,
        amount: amountToApprove,
      });

      toast({
        title: "Approval Submitted",
        description: "Approving tokens for trading...",
      });
    } catch (error: any) {
      toast({
        title: "Approval Failed",
        description: error.message || "Failed to approve tokens",
        variant: "destructive",
      });
    }
  };

  const handleBuy = async () => {
    if (!amount || parseFloat(amount) <= 0) {
      toast({
        title: "Invalid Amount",
        description: "Please enter a valid amount",
        variant: "destructive",
      });
      return;
    }

    try {
      const amountBigInt = parseUnits(amount, 18); // Assuming 18 decimals
      const outcomeId = selectedOutcome === "YES" ? 0 : 1;

      buy({
        outcomeId,
        amount: amountBigInt,
        slippage: 0.5, // 0.5% slippage tolerance
      });

      toast({
        title: "Trade Submitted",
        description: `Buying ${selectedOutcome} tokens...`,
      });
    } catch (error: any) {
      toast({
        title: "Trade Failed",
        description: error.message || "Failed to execute trade",
        variant: "destructive",
      });
    }
  };

  const handleSell = async () => {
    if (!amount || parseFloat(amount) <= 0) {
      toast({
        title: "Invalid Amount",
        description: "Please enter a valid amount",
        variant: "destructive",
      });
      return;
    }

    try {
      const amountBigInt = parseUnits(amount, 18);
      const outcomeId = selectedOutcome === "YES" ? 0 : 1;

      sell({
        outcomeId,
        amount: amountBigInt,
        slippage: 0.5, // 0.5% slippage tolerance
      });

      toast({
        title: "Sell Order Submitted",
        description: `Selling ${selectedOutcome} tokens...`,
      });
    } catch (error: any) {
      toast({
        title: "Sell Failed",
        description: error.message || "Failed to execute sell",
        variant: "destructive",
      });
    }
  };

  if (marketLoading || pricesLoading) {
    return (
      <Card className="hover-elevate">
        <CardContent className="flex items-center justify-center p-8">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
        </CardContent>
      </Card>
    );
  }

  if (!market) {
    return (
      <Card className="hover-elevate">
        <CardContent className="p-6">
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertDescription>Market not found or not deployed on-chain</AlertDescription>
          </Alert>
        </CardContent>
      </Card>
    );
  }

  const yesPrice = prices?.yesPrice ? Number(formatUnits(prices.yesPrice, 18)) : 0.5;
  const noPrice = prices?.noPrice ? Number(formatUnits(prices.noPrice, 18)) : 0.5;
  const yesBalanceFormatted = yesBalance ? formatUnits(yesBalance, 18) : "0";
  const noBalanceFormatted = noBalance ? formatUnits(noBalance, 18) : "0";
  const collateralBalanceFormatted = collateralBalance
    ? formatUnits(collateralBalance, 18)
    : "0";

  return (
    <Card className="hover-elevate border-primary/30 bg-card/80">
      <CardHeader>
        <div className="flex items-start justify-between gap-4">
          <div className="flex-1">
            {/* IPFS Image if available */}
            {ipfsMetadata?.image && (
              <div className="mb-3 rounded-lg overflow-hidden border border-card-border">
                <img 
                  src={ipfsMetadata.image} 
                  alt={title || market.question}
                  className="w-full h-40 object-cover"
                  onError={(e) => {
                    // Hide image on load error
                    e.currentTarget.style.display = 'none';
                  }}
                />
              </div>
            )}
            
            <CardTitle className="text-lg font-sohne">
              {title || ipfsMetadata?.question || market.question || `Market #${marketId}`}
              {ipfsLoading && !ipfsMetadata && (
                <Loader2 className="inline-block ml-2 h-4 w-4 animate-spin text-muted-foreground" />
              )}
            </CardTitle>
            {(description || ipfsMetadata?.description || market.category) && (
              <p className="text-sm text-muted-foreground mt-2">
                {description || ipfsMetadata?.description || market.category}
              </p>
            )}
          </div>
          <Badge variant="outline" className="text-sm font-mono">
            #{marketId}
          </Badge>
        </div>
      </CardHeader>

      <CardContent className="space-y-4">
        {/* Price Display */}
        <div className="grid grid-cols-2 gap-3">
          <div className="p-3 rounded-lg bg-green-500/10 border border-green-500/30">
            <div className="flex items-center gap-2 mb-1">
              <TrendingUp className="h-4 w-4 text-green-400" />
              <span className="text-xs text-muted-foreground">YES Price</span>
            </div>
            <p className="text-2xl font-bold font-mono text-green-400">
              ${yesPrice.toFixed(3)}
            </p>
          </div>
          <div className="p-3 rounded-lg bg-red-500/10 border border-red-500/30">
            <div className="flex items-center gap-2 mb-1">
              <TrendingDown className="h-4 w-4 text-red-400" />
              <span className="text-xs text-muted-foreground">NO Price</span>
            </div>
            <p className="text-2xl font-bold font-mono text-red-400">
              ${noPrice.toFixed(3)}
            </p>
          </div>
        </div>

        {/* User Positions */}
        {address && (
          <div className="p-3 rounded-lg bg-card border border-card-border">
            <p className="text-xs text-muted-foreground mb-2">Your Position</p>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <span className="text-muted-foreground">YES: </span>
                <span className="font-mono font-bold text-green-400">
                  {parseFloat(yesBalanceFormatted).toFixed(2)}
                </span>
              </div>
              <div>
                <span className="text-muted-foreground">NO: </span>
                <span className="font-mono font-bold text-red-400">
                  {parseFloat(noBalanceFormatted).toFixed(2)}
                </span>
              </div>
            </div>
          </div>
        )}

        {/* Trading Interface */}
        {address ? (
          <div className="space-y-3">
            {/* Trade Mode Tabs (Buy/Sell) */}
            <Tabs
              value={tradeMode}
              onValueChange={(v) => setTradeMode(v as "buy" | "sell")}
            >
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="buy">Buy</TabsTrigger>
                <TabsTrigger value="sell">Sell</TabsTrigger>
              </TabsList>

              <TabsContent value={tradeMode} className="space-y-3 mt-4">
                {/* Outcome Selection */}
                <Tabs
                  value={selectedOutcome}
                  onValueChange={(v) => setSelectedOutcome(v as "YES" | "NO")}
                >
                  <TabsList className="grid w-full grid-cols-2">
                    <TabsTrigger value="YES" className="data-[state=active]:bg-green-500/20">
                      YES
                    </TabsTrigger>
                    <TabsTrigger value="NO" className="data-[state=active]:bg-red-500/20">
                      NO
                    </TabsTrigger>
                  </TabsList>

                  <TabsContent value={selectedOutcome} className="space-y-3 mt-4">
                    <div>
                      <label className="text-sm text-muted-foreground mb-2 block">
                        {tradeMode === "buy" ? "Amount (Collateral)" : "Amount (Tokens)"}
                      </label>
                      <Input
                        type="number"
                        placeholder="0.0"
                        value={amount}
                        onChange={(e) => setAmount(e.target.value)}
                        className="font-mono"
                        step="0.01"
                        min="0"
                      />
                      {tradeMode === "buy" ? (
                        <p className="text-xs text-muted-foreground mt-1">
                          Balance: {parseFloat(collateralBalanceFormatted).toFixed(4)}
                        </p>
                      ) : (
                        <p className="text-xs text-muted-foreground mt-1">
                          Balance: {selectedOutcome === "YES" 
                            ? parseFloat(yesBalanceFormatted).toFixed(4)
                            : parseFloat(noBalanceFormatted).toFixed(4)} {selectedOutcome}
                        </p>
                      )}
                    </div>

                    {/* Sell Quote Display */}
                    {tradeMode === "sell" && sellQuote && (
                      <div className="p-3 rounded-lg bg-card border border-card-border">
                        <p className="text-xs text-muted-foreground mb-1">Expected Proceeds</p>
                        <p className="text-lg font-mono font-bold">
                          {formatUnits(sellQuote.tokensOut, 18)} collateral
                        </p>
                        {sellQuote.priceImpact > 1 && (
                          <Alert className="mt-2 border-yellow-500/50 bg-yellow-500/10">
                            <AlertCircle className="h-4 w-4 text-yellow-400" />
                            <AlertDescription className="text-xs text-yellow-400">
                              Price Impact: {sellQuote.priceImpact.toFixed(2)}%
                            </AlertDescription>
                          </Alert>
                        )}
                      </div>
                    )}

                    {tradeMode === "buy" && needsApproval() ? (
                      <Button
                        onClick={handleApprove}
                        disabled={isApproving}
                        className="w-full"
                        variant="outline"
                      >
                        {isApproving ? (
                          <>
                            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                            Approving...
                          </>
                        ) : (
                          "Approve Tokens"
                        )}
                      </Button>
                    ) : (
                      <Button
                        onClick={tradeMode === "buy" ? handleBuy : handleSell}
                        disabled={
                          (tradeMode === "buy" ? isBuying : isSelling) || 
                          !amount || 
                          parseFloat(amount) <= 0
                        }
                        className={
                          selectedOutcome === "YES"
                            ? "w-full bg-green-500 hover:bg-green-600"
                            : "w-full bg-red-500 hover:bg-red-600"
                        }
                      >
                        {(tradeMode === "buy" ? isBuying : isSelling) ? (
                          <>
                            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                            {tradeMode === "buy" ? "Buying..." : "Selling..."}
                          </>
                        ) : (
                          `${tradeMode === "buy" ? "Buy" : "Sell"} ${selectedOutcome}`
                        )}
                      </Button>
                    )}

                    {isSuccess && hash && (
                      <Alert className="border-green-500/50 bg-green-500/10">
                        <CheckCircle className="h-4 w-4 text-green-400" />
                        <AlertDescription className="text-green-400">
                          Trade successful! Transaction: {hash.slice(0, 10)}...
                        </AlertDescription>
                      </Alert>
                    )}
                  </TabsContent>
                </Tabs>
              </TabsContent>
            </Tabs>
          </div>
        ) : (
          <Alert>
            <AlertCircle className="h-4 w-4" />
            <AlertDescription>Connect wallet to trade</AlertDescription>
          </Alert>
        )}
      </CardContent>
    </Card>
  );
}
