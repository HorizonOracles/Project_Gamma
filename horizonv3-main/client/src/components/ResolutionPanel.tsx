// Resolution Panel Component
// Allows admin to resolve markets manually or via oracle
import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/hooks/use-toast";
import { CheckCircle2, Loader2, AlertCircle, Bot, User } from "lucide-react";
import { 
  useProposeResolution, 
  useFinalize, 
  useResolution,
  useRequestResolution,
  useOracleStatus 
} from "@project-gamma/sdk";
import { parseUnits } from "viem";
import { useAccount } from "wagmi";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

interface ResolutionPanelProps {
  marketId: number;
  question?: string;
  teamA?: string;
  teamB?: string;
}

export function ResolutionPanel({ marketId, question, teamA, teamB }: ResolutionPanelProps) {
  const { toast } = useToast();
  const { address } = useAccount();
  const [selectedOutcome, setSelectedOutcome] = useState<string>("0");
  const [evidenceURI, setEvidenceURI] = useState("");
  const [bondAmount, setBondAmount] = useState("100");

  // Fetch resolution state
  const { data: resolution, isLoading: isLoadingResolution } = useResolution(marketId);
  
  // Fetch oracle status (convert marketId to string for oracle API)
  const { data: oracleStatus, isLoading: isLoadingOracle } = useOracleStatus(String(marketId));

  // Resolution hooks
  const { 
    mutate: proposeResolution, 
    isLoading: isProposing,
    error: proposeError 
  } = useProposeResolution(marketId);
  
  const { 
    mutate: finalize, 
    isLoading: isFinalizing,
    error: finalizeError 
  } = useFinalize(marketId);

  // Oracle resolution hook
  const {
    mutate: requestOracleResolution,
    isPending: isRequestingOracle,
    error: oracleError
  } = useRequestResolution();

  const handleManualResolve = () => {
    if (!address) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet to resolve the market",
        variant: "destructive",
      });
      return;
    }

    if (!evidenceURI) {
      toast({
        title: "Evidence required",
        description: "Please provide evidence URI for the resolution",
        variant: "destructive",
      });
      return;
    }

    try {
      const outcomeId = BigInt(selectedOutcome);
      const bond = parseUnits(bondAmount, 18); // Assuming 18 decimals for bond token

      proposeResolution({
        outcomeId,
        bondAmount: bond,
        evidenceURI,
      });

      toast({
        title: "Proposing Resolution",
        description: `Proposing ${outcomeId === 0n ? "YES" : "NO"} outcome for market ${marketId}`,
      });
    } catch (error: any) {
      console.error("Error proposing resolution:", error);
      toast({
        title: "Failed to propose resolution",
        description: error.message || "An error occurred",
        variant: "destructive",
      });
    }
  };

  const handleFinalize = () => {
    if (!address) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet to finalize the market",
        variant: "destructive",
      });
      return;
    }

    finalize();

    toast({
      title: "Finalizing Resolution",
      description: `Finalizing market ${marketId} resolution`,
    });
  };

  const handleOracleResolve = () => {
    if (!address) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet to request oracle resolution",
        variant: "destructive",
      });
      return;
    }

    requestOracleResolution({ 
      marketId: marketId,
      metadata: {
        question: question || "Market resolution request",
        description: `Resolution request for market ${marketId}`
      }
    });

    toast({
      title: "Requesting Oracle Resolution",
      description: "AI Oracle will analyze the market and propose an outcome",
    });
  };

  return (
    <Card className="border-accent/20">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-foreground font-sohne">
          <CheckCircle2 className="h-5 w-5 text-primary" />
          Market Resolution
        </CardTitle>
        <CardDescription className="text-muted-foreground">
          Resolve the market outcome manually or via AI Oracle
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* Market Info */}
        {question && (
          <div className="p-3 bg-muted rounded-lg">
            <p className="text-sm font-medium mb-2">{question}</p>
            {teamA && teamB && (
              <div className="flex items-center gap-4 text-sm text-muted-foreground">
                <span>{teamA} vs {teamB}</span>
              </div>
            )}
          </div>
        )}

        {/* Resolution Status */}
        {isLoadingResolution ? (
          <div className="flex items-center justify-center py-4">
            <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
          </div>
        ) : resolution ? (
          <div className="p-3 bg-primary/5 border border-primary/20 rounded-lg space-y-2">
            <div className="flex items-center gap-2 mb-2">
              <CheckCircle2 className="h-4 w-4 text-primary" />
              <span className="text-sm font-medium">Resolution Status</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">Proposed Outcome:</span>
              <span className="text-sm font-mono">
                {resolution.proposedOutcome === 0n ? "YES" : "NO"}
              </span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">Status:</span>
              <span className="text-sm capitalize">
                {resolution.state === 0 ? "None" : resolution.state === 1 ? "Proposed" : resolution.state === 2 ? "Disputed" : "Finalized"}
              </span>
            </div>
          </div>
        ) : null}

        {/* Resolution Tabs */}
        <Tabs defaultValue="manual" className="w-full">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="manual">
              <User className="h-4 w-4 mr-2" />
              Manual
            </TabsTrigger>
            <TabsTrigger value="oracle">
              <Bot className="h-4 w-4 mr-2" />
              AI Oracle
            </TabsTrigger>
          </TabsList>

          {/* Manual Resolution Tab */}
          <TabsContent value="manual" className="space-y-4">
            {/* Outcome Selection */}
            <div className="space-y-2">
              <Label>Select Winning Outcome</Label>
              <RadioGroup value={selectedOutcome} onValueChange={setSelectedOutcome}>
                <div className="flex items-center space-x-2">
                  <RadioGroupItem value="0" id="outcome-yes" />
                  <Label htmlFor="outcome-yes" className="cursor-pointer">
                    YES / {teamA || "Team A"}
                  </Label>
                </div>
                <div className="flex items-center space-x-2">
                  <RadioGroupItem value="1" id="outcome-no" />
                  <Label htmlFor="outcome-no" className="cursor-pointer">
                    NO / {teamB || "Team B"}
                  </Label>
                </div>
              </RadioGroup>
            </div>

            {/* Evidence URI */}
            <div className="space-y-2">
              <Label htmlFor="evidenceURI">Evidence URI</Label>
              <Textarea
                id="evidenceURI"
                placeholder="ipfs://... or https://..."
                rows={3}
                value={evidenceURI}
                onChange={(e) => setEvidenceURI(e.target.value)}
              />
              <p className="text-xs text-muted-foreground">
                Provide a link to evidence supporting the resolution
              </p>
            </div>

            {/* Bond Amount */}
            <div className="space-y-2">
              <Label htmlFor="bondAmount">Bond Amount (HORIZON Tokens)</Label>
              <input
                id="bondAmount"
                type="number"
                min="0"
                step="10"
                className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
                value={bondAmount}
                onChange={(e) => setBondAmount(e.target.value)}
              />
              <p className="text-xs text-muted-foreground">
                Bond to be locked during dispute period
              </p>
            </div>

            {/* Error Display */}
            {proposeError && (
              <div className="flex items-center gap-2 p-3 bg-destructive/10 border border-destructive/20 rounded-lg">
                <AlertCircle className="h-4 w-4 text-destructive" />
                <span className="text-sm text-destructive">
                  {proposeError.message || "Failed to propose resolution"}
                </span>
              </div>
            )}

            {/* Propose Button */}
            <Button
              onClick={handleManualResolve}
              className="w-full"
              disabled={isProposing || !evidenceURI}
            >
              {isProposing ? (
                <>
                  <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                  Proposing...
                </>
              ) : (
                <>
                  <CheckCircle2 className="h-4 w-4 mr-2" />
                  Propose Resolution
                </>
              )}
            </Button>
          </TabsContent>

          {/* Oracle Resolution Tab */}
          <TabsContent value="oracle" className="space-y-4">
            {/* Oracle Status */}
            {isLoadingOracle ? (
              <div className="flex items-center justify-center py-4">
                <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
              </div>
            ) : oracleStatus ? (
              <div className="p-3 bg-muted rounded-lg space-y-2">
                <div className="flex items-center justify-between">
                  <span className="text-sm text-muted-foreground">Request Status:</span>
                  <span className="text-sm capitalize">{oracleStatus.status}</span>
                </div>
                {oracleStatus.progress !== undefined && (
                  <div className="flex items-center justify-between">
                    <span className="text-sm text-muted-foreground">Progress:</span>
                    <span className="text-sm">{oracleStatus.progress}%</span>
                  </div>
                )}
              </div>
            ) : (
              <div className="p-3 bg-muted/50 rounded-lg">
                <p className="text-sm text-muted-foreground">
                  No oracle resolution requested yet
                </p>
              </div>
            )}

            {/* Info */}
            <div className="p-3 bg-primary/5 border border-primary/20 rounded-lg">
              <p className="text-xs text-muted-foreground">
                <strong>AI Oracle Resolution:</strong> The AI will analyze available data sources, 
                news, and verified information to propose the most accurate market outcome. This 
                process typically takes 1-5 minutes.
              </p>
            </div>

            {/* Error Display */}
            {oracleError && (
              <div className="flex items-center gap-2 p-3 bg-destructive/10 border border-destructive/20 rounded-lg">
                <AlertCircle className="h-4 w-4 text-destructive" />
                <span className="text-sm text-destructive">
                  {oracleError.message || "Failed to request oracle resolution"}
                </span>
              </div>
            )}

            {/* Request Oracle Button */}
            <Button
              onClick={handleOracleResolve}
              className="w-full"
              disabled={isRequestingOracle || oracleStatus?.status === "pending"}
            >
              {isRequestingOracle ? (
                <>
                  <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                  Requesting...
                </>
              ) : (
                <>
                  <Bot className="h-4 w-4 mr-2" />
                  Request AI Oracle Resolution
                </>
              )}
            </Button>
          </TabsContent>
        </Tabs>

        {/* Finalize Button (shown when resolution is proposed) */}
        {resolution && resolution.state === 1 && (
          <div className="pt-4 border-t border-border space-y-4">
            <div className="p-3 bg-muted/50 rounded-lg">
              <p className="text-xs text-muted-foreground">
                <strong>Ready to Finalize:</strong> The dispute period has ended. 
                Finalize the resolution to settle the market and distribute payouts.
              </p>
            </div>

            {finalizeError && (
              <div className="flex items-center gap-2 p-3 bg-destructive/10 border border-destructive/20 rounded-lg">
                <AlertCircle className="h-4 w-4 text-destructive" />
                <span className="text-sm text-destructive">
                  {finalizeError.message || "Failed to finalize"}
                </span>
              </div>
            )}

            <Button
              onClick={handleFinalize}
              variant="default"
              className="w-full"
              disabled={isFinalizing}
            >
              {isFinalizing ? (
                <>
                  <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                  Finalizing...
                </>
              ) : (
                <>
                  <CheckCircle2 className="h-4 w-4 mr-2" />
                  Finalize Resolution
                </>
              )}
            </Button>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
