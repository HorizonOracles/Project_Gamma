/**
 * ResolvedMarkets - Display resolved markets and allow claiming winnings
 */

import { useState } from "react";
import { useAccount } from "wagmi";
import { formatUnits } from "viem";
import {
  useMarkets,
  useMarket,
  useOutcomeBalance,
  useRedeem,
  MarketStatus,
} from "@project-gamma/sdk";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  Trophy, 
  Wallet, 
  AlertCircle, 
  CheckCircle, 
  Loader2,
  TrendingUp,
  TrendingDown 
} from "lucide-react";
import { useToast } from "@/hooks/use-toast";

interface ResolvedMarketCardProps {
  marketId: number;
}

function ResolvedMarketCard({ marketId }: ResolvedMarketCardProps) {
  const { toast } = useToast();
  const { data: market } = useMarket(marketId);
  const { data: yesBalance } = useOutcomeBalance(marketId, 0);
  const { data: noBalance } = useOutcomeBalance(marketId, 1);
  const { write: redeem, isLoading: isRedeeming, isSuccess } = useRedeem();

  if (!market) {
    return <Skeleton className="h-48 bg-card" />;
  }

  const yesBalanceNum = yesBalance ? Number(formatUnits(yesBalance, 18)) : 0;
  const noBalanceNum = noBalance ? Number(formatUnits(noBalance, 18)) : 0;
  
  // Determine if user holds winning tokens
  // Assuming winningOutcome is 0 for YES, 1 for NO
  const winningOutcome = market.winningOutcome;
  const hasWinningTokens = 
    (winningOutcome === 0 && yesBalanceNum > 0) || 
    (winningOutcome === 1 && noBalanceNum > 0);
  
  const winningAmount = winningOutcome === 0 ? yesBalanceNum : noBalanceNum;
  const losingAmount = winningOutcome === 0 ? noBalanceNum : yesBalanceNum;

  // Skip if user has no position in this market
  if (yesBalanceNum === 0 && noBalanceNum === 0) {
    return null;
  }

  const handleClaim = async () => {
    if (!hasWinningTokens) {
      toast({
        title: "No Winnings",
        description: "You don't have any winning tokens to claim",
        variant: "destructive",
      });
      return;
    }

    try {
      redeem({
        marketId,
        outcomeId: winningOutcome as number,
      });

      toast({
        title: "Claim Submitted",
        description: `Claiming ${winningAmount.toFixed(2)} winning tokens...`,
      });
    } catch (error: any) {
      toast({
        title: "Claim Failed",
        description: error.message || "Failed to claim winnings",
        variant: "destructive",
      });
    }
  };

  return (
    <Card className={`border-card-border ${hasWinningTokens ? 'bg-green-500/5' : 'bg-card/80'}`}>
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between gap-3">
          <div className="flex-1">
            <CardTitle className="text-base font-sohne">
              {market.question || `Market #${marketId}`}
            </CardTitle>
            <p className="text-xs text-muted-foreground mt-1">
              {market.category}
            </p>
          </div>
          <Badge variant="secondary" className="text-xs">
            Resolved
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent className="space-y-3">
        {/* Winning Outcome */}
        <div className="p-3 rounded-lg bg-card border border-card-border">
          <div className="flex items-center gap-2 mb-2">
            <Trophy className="h-4 w-4 text-yellow-400" />
            <span className="text-sm font-bold">Winning Outcome</span>
          </div>
          <Badge 
            variant={winningOutcome === 0 ? "default" : "destructive"}
            className="text-sm"
          >
            {winningOutcome === 0 ? (
              <>
                <TrendingUp className="h-3 w-3 mr-1" />
                YES
              </>
            ) : (
              <>
                <TrendingDown className="h-3 w-3 mr-1" />
                NO
              </>
            )}
          </Badge>
        </div>

        {/* User Position */}
        <div className="p-3 rounded-lg bg-card border border-card-border">
          <p className="text-xs text-muted-foreground mb-2">Your Position</p>
          <div className="space-y-2">
            {yesBalanceNum > 0 && (
              <div className="flex items-center justify-between">
                <span className="text-sm">
                  <TrendingUp className="inline h-3 w-3 mr-1 text-green-400" />
                  YES Tokens
                </span>
                <span className={`text-sm font-mono font-bold ${winningOutcome === 0 ? 'text-green-400' : 'text-muted-foreground'}`}>
                  {yesBalanceNum.toFixed(2)}
                  {winningOutcome === 0 && <CheckCircle className="inline h-3 w-3 ml-1" />}
                </span>
              </div>
            )}
            {noBalanceNum > 0 && (
              <div className="flex items-center justify-between">
                <span className="text-sm">
                  <TrendingDown className="inline h-3 w-3 mr-1 text-red-400" />
                  NO Tokens
                </span>
                <span className={`text-sm font-mono font-bold ${winningOutcome === 1 ? 'text-green-400' : 'text-muted-foreground'}`}>
                  {noBalanceNum.toFixed(2)}
                  {winningOutcome === 1 && <CheckCircle className="inline h-3 w-3 ml-1" />}
                </span>
              </div>
            )}
          </div>
        </div>

        {/* Claim Button */}
        {hasWinningTokens ? (
          <div>
            <Button
              onClick={handleClaim}
              disabled={isRedeeming || isSuccess}
              className="w-full bg-green-500 hover:bg-green-600"
            >
              {isRedeeming ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Claiming...
                </>
              ) : isSuccess ? (
                <>
                  <CheckCircle className="mr-2 h-4 w-4" />
                  Claimed!
                </>
              ) : (
                <>
                  <Trophy className="mr-2 h-4 w-4" />
                  Claim {winningAmount.toFixed(2)} Tokens
                </>
              )}
            </Button>
            <p className="text-xs text-center text-muted-foreground mt-2">
              Winning tokens redeem 1:1 for collateral
            </p>
          </div>
        ) : (
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertDescription className="text-xs">
              You held losing tokens. No winnings to claim.
            </AlertDescription>
          </Alert>
        )}
      </CardContent>
    </Card>
  );
}

export function ResolvedMarkets() {
  const { address } = useAccount();
  const { data: markets, isLoading: marketsLoading } = useMarkets({
    status: MarketStatus.Resolved,
  });

  if (!address) {
    return (
      <Alert>
        <Wallet className="h-4 w-4" />
        <AlertDescription>
          Connect your wallet to view resolved markets and claim winnings
        </AlertDescription>
      </Alert>
    );
  }

  if (marketsLoading) {
    return (
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        {[1, 2, 3, 4, 5, 6].map((i) => (
          <Skeleton key={i} className="h-48 bg-card" />
        ))}
      </div>
    );
  }

  if (!markets || markets.length === 0) {
    return (
      <Alert>
        <AlertCircle className="h-4 w-4" />
        <AlertDescription>
          No resolved markets yet. Check back after markets are finalized!
        </AlertDescription>
      </Alert>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-lg font-bold font-sohne">Resolved Markets</h3>
          <p className="text-sm text-muted-foreground">
            Claim your winnings from settled prediction markets
          </p>
        </div>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        {markets.map((market) => (
          <ResolvedMarketCard key={market.id} marketId={market.id} />
        ))}
      </div>
    </div>
  );
}
