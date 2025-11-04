// BetSlipSDK - SDK-based betting component for blockchain transactions
import { useState, useEffect } from "react";
import { useAccount } from 'wagmi';
import { parseUnits, formatUnits } from 'viem';
import { useBuy, useQuote } from '@project-gamma/sdk';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { AlertCircle, TrendingUp, Loader2 } from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";

interface BetSlipSDKProps {
  open: boolean;
  onClose: () => void;
  marketId: number | null; // On-chain market ID
  marketAddress: string | null;
  outcome: 'YES' | 'NO' | null; // YES = Team A, NO = Team B
  teamName: string;
  onSuccess?: (txHash: string) => void;
}

export function BetSlipSDK({
  open,
  onClose,
  marketId,
  marketAddress,
  outcome,
  teamName,
  onSuccess,
}: BetSlipSDKProps) {
  const [amount, setAmount] = useState("");
  const { address, isConnected } = useAccount();

  // Parse amount to bigint for SDK (assuming 18 decimals for BNB)
  const amountBigInt = amount && parseFloat(amount) > 0 
    ? parseUnits(amount, 18) 
    : 0n;

  // Get quote for the trade
  const outcomeId = outcome === 'YES' ? 0 : 1;
  const { data: quote, isLoading: quoteLoading } = useQuote(
    marketId && outcome && amountBigInt > 0n
      ? {
          marketId,
          outcomeId,
          amount: amountBigInt,
          isBuy: true,
        }
      : undefined
  );

  // Buy hook
  const { write: buyTokens, isLoading: isBuying, isSuccess, hash, error } = useBuy(marketId || 0);

  const potentialPayout = quote?.tokensOut 
    ? formatUnits(quote.tokensOut, 18) 
    : "0.0000";

  const fee = quote?.fee 
    ? formatUnits(quote.fee, 18) 
    : "0.0000";

  const priceImpact = quote?.priceImpact || 0;

  useEffect(() => {
    if (!open) {
      setAmount("");
    }
  }, [open]);

  useEffect(() => {
    if (isSuccess && hash) {
      onSuccess?.(hash);
      onClose();
    }
  }, [isSuccess, hash, onSuccess, onClose]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!isConnected) {
      return;
    }

    if (!amount || parseFloat(amount) <= 0) {
      return;
    }

    if (!marketId || !outcome) {
      return;
    }

    // Execute buy transaction
    buyTokens({
      outcomeId,
      amount: amountBigInt,
      slippage: 0.5, // 0.5% slippage tolerance
    });
  };

  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent className="bg-card border-card-border sm:max-w-md">
        <DialogHeader className="border-b border-accent pb-4">
          <DialogTitle className="text-2xl font-bold text-foreground">
            Place Your Bet
          </DialogTitle>
          <DialogDescription className="text-muted-foreground">
            Betting on <span className="text-primary font-semibold">{teamName}</span>
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-6 pt-4">
          {/* Wallet Connection Check */}
          {!isConnected && (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertDescription>
                Please connect your wallet to place bets
              </AlertDescription>
            </Alert>
          )}

          {/* Bet Amount Input */}
          <div className="space-y-2">
            <Label htmlFor="bet-amount" className="text-foreground">
              Bet Amount (BNB)
            </Label>
            <div className="relative">
              <Input
                id="bet-amount"
                type="number"
                step="0.0001"
                min="0"
                placeholder="0.0000"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                className="font-mono text-lg pr-16 bg-input border-accent text-foreground focus:border-primary"
                data-testid="input-bet-amount"
                disabled={!isConnected || isBuying}
              />
              <span className="absolute right-3 top-1/2 -translate-y-1/2 text-sm font-medium text-accent">
                BNB
              </span>
            </div>
          </div>

          {/* Quote Information */}
          {quoteLoading && amount && parseFloat(amount) > 0 && (
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <Loader2 className="h-4 w-4 animate-spin" />
              Getting quote...
            </div>
          )}

          {quote && !quoteLoading && (
            <div className="space-y-3">
              {/* Potential Payout Display */}
              <div className="bg-primary/10 border border-primary rounded-lg p-4">
                <div className="flex items-center justify-between mb-1">
                  <span className="text-sm font-medium text-foreground flex items-center gap-1">
                    <TrendingUp className="h-4 w-4 text-primary" />
                    Tokens You'll Receive
                  </span>
                </div>
                <p className="text-3xl font-mono font-bold text-primary">
                  {parseFloat(potentialPayout).toFixed(4)} Tokens
                </p>
                <p className="text-xs text-muted-foreground mt-1">
                  If outcome wins, redeem for 1:1 collateral
                </p>
              </div>

              {/* Trade Details */}
              <div className="bg-muted rounded-lg p-3 space-y-2 text-sm">
                <div className="flex items-center justify-between">
                  <span className="text-muted-foreground">Trading Fee</span>
                  <span className="font-mono text-foreground">{parseFloat(fee).toFixed(4)} BNB</span>
                </div>
                {priceImpact > 0 && (
                  <div className="flex items-center justify-between">
                    <span className="text-muted-foreground">Price Impact</span>
                    <span className={`font-mono ${priceImpact > 5 ? 'text-red-500' : 'text-foreground'}`}>
                      {priceImpact.toFixed(2)}%
                    </span>
                  </div>
                )}
              </div>

              {priceImpact > 5 && (
                <Alert className="border-yellow-500 bg-yellow-500/10">
                  <AlertCircle className="h-4 w-4 text-yellow-500" />
                  <AlertDescription className="text-xs text-muted-foreground">
                    High price impact! Consider reducing bet amount.
                  </AlertDescription>
                </Alert>
              )}
            </div>
          )}

          {/* Error Display */}
          {error && (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertDescription>
                {error instanceof Error ? error.message : 'Failed to place bet'}
              </AlertDescription>
            </Alert>
          )}

          {/* Market Info */}
          {marketId && (
            <div className="text-xs text-muted-foreground">
              On-chain Market ID: {marketId}
            </div>
          )}

          {/* Action Buttons */}
          <div className="flex gap-3 pt-2">
            <Button
              type="button"
              variant="outline"
              onClick={onClose}
              className="flex-1 border-accent text-accent hover:bg-accent hover:text-accent-foreground"
              disabled={isBuying}
              data-testid="button-cancel-bet"
            >
              Cancel
            </Button>
            <Button
              type="submit"
              className="flex-1 bg-primary hover:bg-primary/90 text-primary-foreground font-bold"
              disabled={!isConnected || isBuying || !amount || parseFloat(amount) <= 0 || !quote}
              data-testid="button-confirm-bet"
            >
              {isBuying ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Confirming...
                </>
              ) : (
                'Confirm Bet'
              )}
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
}
