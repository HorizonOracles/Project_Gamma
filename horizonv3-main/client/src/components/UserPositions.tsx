/**
 * UserPositions - Display user's holdings across all markets
 * Calculates unrealized P&L based on current market prices
 */

import { useMemo } from "react";
import { useAccount } from "wagmi";
import { formatUnits } from "viem";
import {
  useMarkets,
  useMarket,
  usePrices,
  useOutcomeBalance,
  MarketStatus,
} from "@project-gamma/sdk";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Skeleton } from "@/components/ui/skeleton";
import { TrendingUp, TrendingDown, Wallet, AlertCircle } from "lucide-react";

interface PositionCardProps {
  marketId: number;
}

function PositionCard({ marketId }: PositionCardProps) {
  const { data: market } = useMarket(marketId);
  const { data: prices } = usePrices(marketId);
  const { data: yesBalance } = useOutcomeBalance(marketId, 0);
  const { data: noBalance } = useOutcomeBalance(marketId, 1);

  if (!market || !prices) {
    return (
      <Skeleton className="h-32 bg-card" />
    );
  }

  const yesBalanceNum = yesBalance ? Number(formatUnits(yesBalance, 18)) : 0;
  const noBalanceNum = noBalance ? Number(formatUnits(noBalance, 18)) : 0;
  const yesPrice = Number(formatUnits(prices.yesPrice, 18));
  const noPrice = Number(formatUnits(prices.noPrice, 18));

  // Calculate position value
  const yesValue = yesBalanceNum * yesPrice;
  const noValue = noBalanceNum * noPrice;
  const totalValue = yesValue + noValue;

  // Skip if no position
  if (yesBalanceNum === 0 && noBalanceNum === 0) {
    return null;
  }

  return (
    <Card className="border-card-border bg-card/80">
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between">
          <div className="flex-1">
            <CardTitle className="text-base font-sohne">
              {market.question || `Market #${marketId}`}
            </CardTitle>
            <p className="text-xs text-muted-foreground mt-1">
              {market.category}
            </p>
          </div>
          <Badge 
            variant={market.status === MarketStatus.Active ? "default" : "outline"}
            className="text-xs"
          >
            {market.status === MarketStatus.Active ? "Active" : "Resolved"}
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent className="space-y-3">
        {/* Position Details */}
        <div className="grid grid-cols-2 gap-3">
          {yesBalanceNum > 0 && (
            <div className="p-2 rounded-lg bg-green-500/10 border border-green-500/30">
              <div className="flex items-center gap-1 mb-1">
                <TrendingUp className="h-3 w-3 text-green-400" />
                <span className="text-xs text-muted-foreground">YES</span>
              </div>
              <p className="text-sm font-bold font-mono text-green-400">
                {yesBalanceNum.toFixed(2)} tokens
              </p>
              <p className="text-xs text-muted-foreground">
                @ ${yesPrice.toFixed(3)} = ${yesValue.toFixed(2)}
              </p>
            </div>
          )}
          
          {noBalanceNum > 0 && (
            <div className="p-2 rounded-lg bg-red-500/10 border border-red-500/30">
              <div className="flex items-center gap-1 mb-1">
                <TrendingDown className="h-3 w-3 text-red-400" />
                <span className="text-xs text-muted-foreground">NO</span>
              </div>
              <p className="text-sm font-bold font-mono text-red-400">
                {noBalanceNum.toFixed(2)} tokens
              </p>
              <p className="text-xs text-muted-foreground">
                @ ${noPrice.toFixed(3)} = ${noValue.toFixed(2)}
              </p>
            </div>
          )}
        </div>

        {/* Total Value */}
        <div className="pt-2 border-t border-card-border">
          <div className="flex items-center justify-between">
            <span className="text-sm text-muted-foreground">Position Value</span>
            <span className="text-sm font-bold font-mono">
              ${totalValue.toFixed(2)}
            </span>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}

export function UserPositions() {
  const { address } = useAccount();
  const { data: markets, isLoading: marketsLoading } = useMarkets({
    status: MarketStatus.Active,
  });

  if (!address) {
    return (
      <Alert>
        <Wallet className="h-4 w-4" />
        <AlertDescription>
          Connect your wallet to view your positions
        </AlertDescription>
      </Alert>
    );
  }

  if (marketsLoading) {
    return (
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        {[1, 2, 3, 4, 5, 6].map((i) => (
          <Skeleton key={i} className="h-32 bg-card" />
        ))}
      </div>
    );
  }

  if (!markets || markets.length === 0) {
    return (
      <Alert>
        <AlertCircle className="h-4 w-4" />
        <AlertDescription>
          No active markets found. Start trading to build your positions!
        </AlertDescription>
      </Alert>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-lg font-bold font-sohne">Your Positions</h3>
          <p className="text-sm text-muted-foreground">
            Track your holdings and unrealized P&L across all markets
          </p>
        </div>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        {markets.map((market) => (
          <PositionCard key={market.id} marketId={market.id} />
        ))}
      </div>
    </div>
  );
}
