// BetSlipDialog - Dialog for placing bets
// This is a wrapper around BetSlipSDK for backward compatibility
import { BetSlipSDK } from "./bet-slip-sdk";

interface BetSlipDialogProps {
  open: boolean;
  onClose: () => void;
  // Accept either a market object or direct marketId string
  market?: {
    id: string;
    chainMarketId?: number;
    marketAddress?: string;
  } | null;
  marketId?: string;
  outcome: 'A' | 'B' | null;
  teamName: string;
  currentOdds?: number; // Optional odds display
  onSuccess?: (txHash: string) => void;
  onConfirm?: (amount: string) => Promise<void>; // Legacy callback
}

export function BetSlipDialog({
  open,
  onClose,
  market,
  marketId,
  outcome,
  teamName,
  currentOdds,
  onSuccess,
  onConfirm,
}: BetSlipDialogProps) {
  // Use marketId from either prop
  const finalMarketId = market?.id ?? marketId;
  
  // Convert outcome 'A'/'B' to 'YES'/'NO' for SDK
  const sdkOutcome = outcome === 'A' ? 'YES' : outcome === 'B' ? 'NO' : null;
  
  return (
    <BetSlipSDK
      open={open}
      onClose={onClose}
      marketId={market?.chainMarketId ?? null}
      marketAddress={market?.marketAddress ?? null}
      outcome={sdkOutcome}
      teamName={teamName}
      onSuccess={onSuccess}
    />
  );
}
