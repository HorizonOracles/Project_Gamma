// Liquidity Manager Component
// Allows admin to add liquidity to markets
import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/hooks/use-toast";
import { Droplets, Loader2, TrendingUp, AlertCircle } from "lucide-react";
import { useAddLiquidity, useLPPosition, useMarket } from "@project-gamma/sdk";
import { parseUnits, formatUnits } from "viem";
import { useAccount } from "wagmi";

interface LiquidityManagerProps {
  marketId: number;
  collateralDecimals?: number; // Decimals of the collateral token (default: 6 for USDC)
}

export function LiquidityManager({ marketId, collateralDecimals = 6 }: LiquidityManagerProps) {
  const { toast } = useToast();
  const { address } = useAccount();
  const [liquidityAmount, setLiquidityAmount] = useState("");
  
  // Fetch market data
  const { data: market, isLoading: isLoadingMarket } = useMarket(marketId);
  
  // Fetch LP position
  const { data: lpPosition, isLoading: isLoadingLP } = useLPPosition(marketId);
  
  // Add liquidity hook
  const { 
    write: addLiquidity, 
    isLoading: isAddingLiquidity,
    isSuccess,
    error
  } = useAddLiquidity(marketId);

  const handleAddLiquidity = () => {
    if (!address) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet to add liquidity",
        variant: "destructive",
      });
      return;
    }

    if (!liquidityAmount || parseFloat(liquidityAmount) <= 0) {
      toast({
        title: "Invalid amount",
        description: "Please enter a valid liquidity amount",
        variant: "destructive",
      });
      return;
    }

    try {
      const amount = parseUnits(liquidityAmount, collateralDecimals);
      
      addLiquidity({ amount });
      
      toast({
        title: "Adding Liquidity",
        description: `Adding ${liquidityAmount} tokens of liquidity to the market`,
      });
    } catch (error: any) {
      console.error("Error adding liquidity:", error);
      toast({
        title: "Failed to add liquidity",
        description: error.message || "An error occurred while adding liquidity",
        variant: "destructive",
      });
    }
  };

  // Handle success
  if (isSuccess) {
    setTimeout(() => {
      toast({
        title: "Liquidity Added!",
        description: `Successfully added ${liquidityAmount} tokens to the market`,
      });
      setLiquidityAmount("");
    }, 1000);
  }

  return (
    <Card className="border-accent/20">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-foreground font-sohne">
          <Droplets className="h-5 w-5 text-primary" />
          Liquidity Manager
        </CardTitle>
        <CardDescription className="text-muted-foreground">
          Add liquidity to provide trading depth and earn fees
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* Market Info */}
        {isLoadingMarket ? (
          <div className="flex items-center justify-center py-4">
            <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
          </div>
        ) : market ? (
          <div className="p-3 bg-muted rounded-lg space-y-2">
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">Market ID:</span>
              <span className="text-sm font-mono">{marketId}</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">Status:</span>
              <span className="text-sm capitalize">{market.status}</span>
            </div>
            {market.question && (
              <div className="pt-2 border-t border-border">
                <p className="text-sm font-medium">{market.question}</p>
              </div>
            )}
          </div>
        ) : (
          <div className="flex items-center gap-2 p-3 bg-destructive/10 border border-destructive/20 rounded-lg">
            <AlertCircle className="h-4 w-4 text-destructive" />
            <span className="text-sm text-destructive">Market not found</span>
          </div>
        )}

        {/* LP Position */}
        {!isLoadingLP && lpPosition && (
          <div className="p-3 bg-primary/5 border border-primary/20 rounded-lg space-y-2">
            <div className="flex items-center gap-2 mb-2">
              <TrendingUp className="h-4 w-4 text-primary" />
              <span className="text-sm font-medium">Your LP Position</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">LP Tokens:</span>
              <span className="text-sm font-mono">
                {formatUnits(lpPosition.lpTokens, 18)}
              </span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">Share:</span>
              <span className="text-sm font-mono">
                {lpPosition.share ? `${(lpPosition.share * 100).toFixed(2)}%` : "0%"}
              </span>
            </div>
          </div>
        )}

        {/* Add Liquidity Form */}
        <div className="space-y-2">
          <Label htmlFor="liquidityAmount">Liquidity Amount</Label>
          <Input
            id="liquidityAmount"
            type="number"
            min="0"
            step="0.01"
            placeholder="Enter amount (e.g., 1000)"
            value={liquidityAmount}
            onChange={(e) => setLiquidityAmount(e.target.value)}
            disabled={isAddingLiquidity || !market}
          />
          <p className="text-xs text-muted-foreground">
            Amount of collateral tokens to add as liquidity
          </p>
        </div>

        {/* Error Display */}
        {error && (
          <div className="flex items-center gap-2 p-3 bg-destructive/10 border border-destructive/20 rounded-lg">
            <AlertCircle className="h-4 w-4 text-destructive" />
            <span className="text-sm text-destructive">
              {error.message || "Failed to add liquidity"}
            </span>
          </div>
        )}

        {/* Submit Button */}
        <Button
          onClick={handleAddLiquidity}
          className="w-full"
          disabled={isAddingLiquidity || !market || !liquidityAmount}
        >
          {isAddingLiquidity ? (
            <>
              <Loader2 className="h-4 w-4 mr-2 animate-spin" />
              Adding Liquidity...
            </>
          ) : (
            <>
              <Droplets className="h-4 w-4 mr-2" />
              Add Liquidity
            </>
          )}
        </Button>

        {/* Info Note */}
        <div className="p-3 bg-muted/50 rounded-lg">
          <p className="text-xs text-muted-foreground">
            <strong>Note:</strong> Adding liquidity provides trading depth for users and earns you 
            a share of trading fees. LP tokens represent your proportional share of the liquidity pool.
          </p>
        </div>
      </CardContent>
    </Card>
  );
}
