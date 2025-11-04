/**
 * LiquidityPanel - Liquidity Provider Interface
 * Allows users to add/remove liquidity from prediction markets
 * Uses Project Gamma SDK for on-chain liquidity operations
 */

import { useState } from "react";
import { useAccount } from "wagmi";
import { parseUnits, formatUnits } from "viem";
import {
  useMarket,
  useAddLiquidity,
  useRemoveLiquidity,
  useLPPosition,
  useAllowance,
  useApprove,
  useBalance,
} from "@project-gamma/sdk";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Loader2, Droplets, TrendingUp, CheckCircle, AlertCircle, Info } from "lucide-react";
import { useToast } from "@/hooks/use-toast";

interface LiquidityPanelProps {
  marketId: number;
}

export function LiquidityPanel({ marketId }: LiquidityPanelProps) {
  const { toast } = useToast();
  const { address } = useAccount();

  // Market data
  const { data: market, isLoading: marketLoading } = useMarket(marketId);
  const { data: lpPosition } = useLPPosition(marketId);
  const { data: collateralBalance } = useBalance(
    market?.collateralToken as `0x${string}`
  );

  // Token approval for adding liquidity
  const { data: allowance } = useAllowance({
    tokenAddress: market?.collateralToken as `0x${string}`,
    spenderAddress: market?.marketAddress as `0x${string}`,
  });

  const { write: approve, isLoading: isApproving } = useApprove();

  // Liquidity operations
  const {
    write: addLiquidity,
    isLoading: isAdding,
    isSuccess: addSuccess,
    hash: addHash,
  } = useAddLiquidity(marketId);

  const {
    write: removeLiquidity,
    isLoading: isRemoving,
    isSuccess: removeSuccess,
    hash: removeHash,
  } = useRemoveLiquidity(marketId);

  // UI state
  const [mode, setMode] = useState<"add" | "remove">("add");
  const [amount, setAmount] = useState("");

  const isSuccess = addSuccess || removeSuccess;
  const hash = addHash || removeHash;

  // Check if approval is needed for adding liquidity
  const needsApproval = () => {
    if (mode !== "add" || !amount || !allowance || !market?.collateralToken) return false;
    try {
      const amountBigInt = parseUnits(amount, 18);
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
        description: "Approving collateral tokens...",
      });
    } catch (error: any) {
      toast({
        title: "Approval Failed",
        description: error.message || "Failed to approve tokens",
        variant: "destructive",
      });
    }
  };

  const handleAddLiquidity = async () => {
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

      addLiquidity({ amount: amountBigInt });

      toast({
        title: "Adding Liquidity",
        description: "Transaction submitted...",
      });
    } catch (error: any) {
      toast({
        title: "Add Liquidity Failed",
        description: error.message || "Failed to add liquidity",
        variant: "destructive",
      });
    }
  };

  const handleRemoveLiquidity = async () => {
    if (!amount || parseFloat(amount) <= 0) {
      toast({
        title: "Invalid Amount",
        description: "Please enter a valid amount",
        variant: "destructive",
      });
      return;
    }

    try {
      const lpTokensBigInt = parseUnits(amount, 18);

      removeLiquidity({ lpTokens: lpTokensBigInt });

      toast({
        title: "Removing Liquidity",
        description: "Transaction submitted...",
      });
    } catch (error: any) {
      toast({
        title: "Remove Liquidity Failed",
        description: error.message || "Failed to remove liquidity",
        variant: "destructive",
      });
    }
  };

  const setMaxAmount = () => {
    if (mode === "add" && collateralBalance) {
      setAmount(formatUnits(collateralBalance, 18));
    } else if (mode === "remove" && lpPosition?.lpTokens) {
      setAmount(formatUnits(lpPosition.lpTokens, 18));
    }
  };

  if (marketLoading) {
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
            <AlertDescription>Market not found</AlertDescription>
          </Alert>
        </CardContent>
      </Card>
    );
  }

  const collateralBalanceFormatted = collateralBalance
    ? formatUnits(collateralBalance, 18)
    : "0";
  const lpTokensFormatted = lpPosition?.lpTokens
    ? formatUnits(lpPosition.lpTokens, 18)
    : "0";
  const lpValueFormatted = lpPosition?.value
    ? formatUnits(lpPosition.value, 18)
    : "0";
  const lpShareFormatted = lpPosition?.share
    ? (lpPosition.share * 100).toFixed(4)
    : "0";

  return (
    <Card className="hover-elevate border-blue-500/30 bg-card/80">
      <CardHeader>
        <div className="flex items-start justify-between gap-4">
          <div className="flex items-center gap-2">
            <Droplets className="h-5 w-5 text-blue-400" />
            <CardTitle className="text-lg font-sohne">
              Liquidity Pool
            </CardTitle>
          </div>
          <Badge variant="outline" className="text-sm font-mono">
            #{marketId}
          </Badge>
        </div>
      </CardHeader>

      <CardContent className="space-y-4">
        {/* LP Position Display */}
        {address && lpPosition && lpPosition.lpTokens > 0n && (
          <div className="p-4 rounded-lg bg-blue-500/10 border border-blue-500/30">
            <div className="flex items-center gap-2 mb-3">
              <TrendingUp className="h-4 w-4 text-blue-400" />
              <p className="text-sm font-semibold text-blue-400">Your LP Position</p>
            </div>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <span className="text-muted-foreground">LP Tokens: </span>
                <span className="font-mono font-bold text-foreground">
                  {parseFloat(lpTokensFormatted).toFixed(2)}
                </span>
              </div>
              <div>
                <span className="text-muted-foreground">Value: </span>
                <span className="font-mono font-bold text-foreground">
                  {parseFloat(lpValueFormatted).toFixed(2)}
                </span>
              </div>
              <div className="col-span-2">
                <span className="text-muted-foreground">Pool Share: </span>
                <span className="font-mono font-bold text-blue-400">
                  {lpShareFormatted}%
                </span>
              </div>
            </div>
          </div>
        )}

        {/* Information Banner */}
        <Alert className="border-blue-500/50 bg-blue-500/5">
          <Info className="h-4 w-4 text-blue-400" />
          <AlertDescription className="text-xs text-muted-foreground">
            {mode === "add" 
              ? "Add collateral to the pool to earn trading fees. LP tokens represent your share."
              : "Remove your LP tokens to withdraw collateral plus accumulated fees."}
          </AlertDescription>
        </Alert>

        {/* Liquidity Interface */}
        {address ? (
          <div className="space-y-3">
            <Tabs
              value={mode}
              onValueChange={(v) => {
                setMode(v as "add" | "remove");
                setAmount(""); // Clear amount when switching modes
              }}
            >
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="add">Add Liquidity</TabsTrigger>
                <TabsTrigger value="remove">Remove Liquidity</TabsTrigger>
              </TabsList>

              <TabsContent value={mode} className="space-y-3 mt-4">
                <div>
                  <div className="flex items-center justify-between mb-2">
                    <label className="text-sm text-muted-foreground">
                      {mode === "add" ? "Amount (Collateral)" : "Amount (LP Tokens)"}
                    </label>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={setMaxAmount}
                      className="h-6 text-xs text-blue-400 hover:text-blue-300"
                    >
                      MAX
                    </Button>
                  </div>
                  <Input
                    type="number"
                    placeholder="0.0"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    className="font-mono"
                    step="0.01"
                    min="0"
                  />
                  {mode === "add" ? (
                    <p className="text-xs text-muted-foreground mt-1">
                      Balance: {parseFloat(collateralBalanceFormatted).toFixed(4)} collateral
                    </p>
                  ) : (
                    <p className="text-xs text-muted-foreground mt-1">
                      Balance: {parseFloat(lpTokensFormatted).toFixed(4)} LP tokens
                    </p>
                  )}
                </div>

                {mode === "add" && needsApproval() ? (
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
                      "Approve Collateral"
                    )}
                  </Button>
                ) : (
                  <Button
                    onClick={mode === "add" ? handleAddLiquidity : handleRemoveLiquidity}
                    disabled={
                      (mode === "add" ? isAdding : isRemoving) ||
                      !amount ||
                      parseFloat(amount) <= 0
                    }
                    className="w-full bg-blue-500 hover:bg-blue-600"
                  >
                    {(mode === "add" ? isAdding : isRemoving) ? (
                      <>
                        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                        {mode === "add" ? "Adding..." : "Removing..."}
                      </>
                    ) : (
                      <>
                        <Droplets className="mr-2 h-4 w-4" />
                        {mode === "add" ? "Add Liquidity" : "Remove Liquidity"}
                      </>
                    )}
                  </Button>
                )}

                {isSuccess && hash && (
                  <Alert className="border-green-500/50 bg-green-500/10">
                    <CheckCircle className="h-4 w-4 text-green-400" />
                    <AlertDescription className="text-green-400">
                      Transaction successful! Hash: {hash.slice(0, 10)}...
                    </AlertDescription>
                  </Alert>
                )}
              </TabsContent>
            </Tabs>
          </div>
        ) : (
          <Alert>
            <AlertCircle className="h-4 w-4" />
            <AlertDescription>Connect wallet to manage liquidity</AlertDescription>
          </Alert>
        )}
      </CardContent>
    </Card>
  );
}
