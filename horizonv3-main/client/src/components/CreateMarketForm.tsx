/**
 * CreateMarketForm - Admin Market Creation Interface
 * Allows admins to create new prediction markets with metadata upload
 * Uses Project Gamma SDK for on-chain market creation
 */

import { useState } from "react";
import { useAccount } from "wagmi";
import { parseUnits, formatUnits, isAddress } from "viem";
import {
  useCreateMarket,
  useUploadMetadata,
  useMinCreatorStake,
  useBalance,
} from "@project-gamma/sdk";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Loader2, PlusCircle, CheckCircle, AlertCircle, Info, Upload } from "lucide-react";
import { useToast } from "@/hooks/use-toast";

// Market type enum matching SDK
enum MarketType {
  Binary = 0,
  MultiChoice = 1,
  LimitOrder = 2,
  PooledLiquidity = 3,
  Dependent = 4,
  Bracket = 5,
  Trend = 6,
}

// Common categories
const MARKET_CATEGORIES = [
  "Sports",
  "Crypto",
  "Politics",
  "Entertainment",
  "Technology",
  "Finance",
  "Science",
  "Other",
];

// Default collateral token (USDC on BNB Chain)
const DEFAULT_COLLATERAL_TOKEN = "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d";

interface CreateMarketFormProps {
  onSuccess?: (marketId: bigint) => void;
}

export function CreateMarketForm({ onSuccess }: CreateMarketFormProps) {
  const { toast } = useToast();
  const { address } = useAccount();

  // SDK hooks
  const { data: minStake, isLoading: minStakeLoading } = useMinCreatorStake();
  const {
    mutate: uploadMetadata,
    isPending: isUploading,
    data: ipfsResult,
  } = useUploadMetadata();
  const {
    write: createMarket,
    isLoading: isCreating,
    isSuccess: createSuccess,
    data: marketId,
  } = useCreateMarket();

  // Form state
  const [question, setQuestion] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("Sports");
  const [customCategory, setCustomCategory] = useState("");
  const [endDate, setEndDate] = useState("");
  const [endTime, setEndTime] = useState("");
  const [collateralToken, setCollateralToken] = useState(DEFAULT_COLLATERAL_TOKEN);
  const [creatorStake, setCreatorStake] = useState("");
  const [marketType, setMarketType] = useState<MarketType>(MarketType.Binary);
  const [outcomeCount, setOutcomeCount] = useState("2");

  // Step tracking
  const [currentStep, setCurrentStep] = useState<"form" | "uploading" | "creating" | "success">("form");

  // Validation
  const isFormValid = () => {
    if (!question || question.length < 10) return false;
    if (!endDate || !endTime) return false;
    if (!isAddress(collateralToken)) return false;
    if (!creatorStake || parseFloat(creatorStake) < 0) return false;
    if (marketType === MarketType.MultiChoice && parseInt(outcomeCount) < 3) return false;
    return true;
  };

  const getEndTimeUnix = () => {
    const dateTime = new Date(`${endDate}T${endTime}`);
    return BigInt(Math.floor(dateTime.getTime() / 1000));
  };

  const getMinEndTime = () => {
    const now = new Date();
    now.setHours(now.getHours() + 1); // At least 1 hour from now
    return now.toISOString().slice(0, 16);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!address) {
      toast({
        title: "Wallet Not Connected",
        description: "Please connect your wallet to create a market",
        variant: "destructive",
      });
      return;
    }

    if (!isFormValid()) {
      toast({
        title: "Invalid Form",
        description: "Please fill all required fields correctly",
        variant: "destructive",
      });
      return;
    }

    try {
      // Step 1: Upload metadata to IPFS
      setCurrentStep("uploading");
      toast({
        title: "Uploading Metadata",
        description: "Uploading market details to IPFS...",
      });

      uploadMetadata(
        {
          question,
          description,
          category: category === "Other" && customCategory ? customCategory : category,
        },
        {
          onSuccess: (result) => {
            // Step 2: Create market on-chain
            setCurrentStep("creating");
            toast({
              title: "Metadata Uploaded",
              description: "Creating market on-chain...",
            });

            const params = {
              question,
              endTime: getEndTimeUnix(),
              collateralToken: collateralToken as `0x${string}`,
              category: category === "Other" && customCategory ? customCategory : category,
              metadataURI: result.uri,
              creatorStake: parseUnits(creatorStake, 18),
              marketType,
              outcomeCount: parseInt(outcomeCount),
              liquidityParameter: 0n, // Not used for Binary markets
            };

            createMarket(params);
          },
          onError: (error: any) => {
            setCurrentStep("form");
            toast({
              title: "Upload Failed",
              description: error.message || "Failed to upload metadata to IPFS",
              variant: "destructive",
            });
          },
        }
      );
    } catch (error: any) {
      setCurrentStep("form");
      toast({
        title: "Error",
        description: error.message || "An unexpected error occurred",
        variant: "destructive",
      });
    }
  };

  // Handle successful market creation
  if (createSuccess && marketId) {
    if (currentStep !== "success") {
      setCurrentStep("success");
      toast({
        title: "Market Created Successfully!",
        description: `Market ID: ${marketId.toString()}`,
      });
      if (onSuccess) {
        onSuccess(marketId);
      }
    }
  }

  // Reset form
  const handleReset = () => {
    setQuestion("");
    setDescription("");
    setCategory("Sports");
    setCustomCategory("");
    setEndDate("");
    setEndTime("");
    setCollateralToken(DEFAULT_COLLATERAL_TOKEN);
    setCreatorStake("");
    setMarketType(MarketType.Binary);
    setOutcomeCount("2");
    setCurrentStep("form");
  };

  if (currentStep === "success" && marketId) {
    return (
      <Card className="w-full max-w-2xl mx-auto border-green-500/20 bg-green-500/5">
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-green-400">
            <CheckCircle className="w-6 h-6" />
            Market Created Successfully
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <Alert className="bg-green-500/10 border-green-500/20">
            <CheckCircle className="h-4 w-4 text-green-400" />
            <AlertDescription className="text-green-100">
              <div className="space-y-2">
                <p className="font-semibold">Market ID: {marketId.toString()}</p>
                <p>Question: {question}</p>
                <p>Category: {category === "Other" && customCategory ? customCategory : category}</p>
                {ipfsResult && (
                  <p className="text-sm break-all">
                    IPFS URI: <span className="text-blue-400">{ipfsResult.uri}</span>
                  </p>
                )}
              </div>
            </AlertDescription>
          </Alert>

          <div className="flex gap-3">
            <Button
              onClick={handleReset}
              className="flex-1 bg-blue-600 hover:bg-blue-700 text-white"
            >
              Create Another Market
            </Button>
            <Button
              onClick={() => window.location.href = "/blockchain-markets"}
              variant="outline"
              className="flex-1"
            >
              View Markets
            </Button>
          </div>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card className="w-full max-w-2xl mx-auto border-blue-500/20 bg-blue-500/5">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-blue-400">
          <PlusCircle className="w-6 h-6" />
          Create New Prediction Market
        </CardTitle>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Info Banner */}
          <Alert className="bg-blue-500/10 border-blue-500/20">
            <Info className="h-4 w-4 text-blue-400" />
            <AlertDescription className="text-blue-100">
              Create a new prediction market on-chain. Metadata will be stored on IPFS.
              {minStake && (
                <p className="mt-1 text-sm">
                  Minimum creator stake: <span className="font-semibold">{formatUnits(minStake, 18)} HORIZON</span>
                </p>
              )}
            </AlertDescription>
          </Alert>

          {/* Question */}
          <div className="space-y-2">
            <Label htmlFor="question" className="text-foreground font-semibold">
              Market Question <span className="text-red-500">*</span>
            </Label>
            <Input
              id="question"
              placeholder="e.g., Will Bitcoin reach $100,000 by end of 2025?"
              value={question}
              onChange={(e) => setQuestion(e.target.value)}
              className="bg-background/50 border-blue-500/20"
              disabled={currentStep !== "form"}
              maxLength={200}
            />
            <p className="text-xs text-muted-foreground">{question.length}/200 characters</p>
          </div>

          {/* Description */}
          <div className="space-y-2">
            <Label htmlFor="description" className="text-foreground font-semibold">
              Description (Optional)
            </Label>
            <Textarea
              id="description"
              placeholder="Provide additional context, resolution criteria, or rules..."
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="bg-background/50 border-blue-500/20 min-h-[100px]"
              disabled={currentStep !== "form"}
              maxLength={1000}
            />
            <p className="text-xs text-muted-foreground">{description.length}/1000 characters</p>
          </div>

          {/* Category */}
          <div className="space-y-2">
            <Label htmlFor="category" className="text-foreground font-semibold">
              Category <span className="text-red-500">*</span>
            </Label>
            <Select value={category} onValueChange={setCategory} disabled={currentStep !== "form"}>
              <SelectTrigger className="bg-background/50 border-blue-500/20">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {MARKET_CATEGORIES.map((cat) => (
                  <SelectItem key={cat} value={cat}>
                    {cat}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {category === "Other" && (
              <Input
                placeholder="Enter custom category"
                value={customCategory}
                onChange={(e) => setCustomCategory(e.target.value)}
                className="bg-background/50 border-blue-500/20 mt-2"
                disabled={currentStep !== "form"}
              />
            )}
          </div>

          {/* End Date & Time */}
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="endDate" className="text-foreground font-semibold">
                End Date <span className="text-red-500">*</span>
              </Label>
              <Input
                id="endDate"
                type="date"
                value={endDate}
                onChange={(e) => setEndDate(e.target.value)}
                min={new Date().toISOString().split("T")[0]}
                className="bg-background/50 border-blue-500/20"
                disabled={currentStep !== "form"}
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="endTime" className="text-foreground font-semibold">
                End Time <span className="text-red-500">*</span>
              </Label>
              <Input
                id="endTime"
                type="time"
                value={endTime}
                onChange={(e) => setEndTime(e.target.value)}
                className="bg-background/50 border-blue-500/20"
                disabled={currentStep !== "form"}
              />
            </div>
          </div>

          {/* Market Type */}
          <div className="space-y-2">
            <Label htmlFor="marketType" className="text-foreground font-semibold">
              Market Type
            </Label>
            <Select
              value={marketType.toString()}
              onValueChange={(value) => setMarketType(parseInt(value) as MarketType)}
              disabled={currentStep !== "form"}
            >
              <SelectTrigger className="bg-background/50 border-blue-500/20">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value={MarketType.Binary.toString()}>Binary (Yes/No)</SelectItem>
                <SelectItem value={MarketType.MultiChoice.toString()}>Multi-Choice (3+ outcomes)</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* Outcome Count (for Multi-Choice) */}
          {marketType === MarketType.MultiChoice && (
            <div className="space-y-2">
              <Label htmlFor="outcomeCount" className="text-foreground font-semibold">
                Number of Outcomes <span className="text-red-500">*</span>
              </Label>
              <Input
                id="outcomeCount"
                type="number"
                min="3"
                max="8"
                value={outcomeCount}
                onChange={(e) => setOutcomeCount(e.target.value)}
                className="bg-background/50 border-blue-500/20"
                disabled={currentStep !== "form"}
              />
              <p className="text-xs text-muted-foreground">Multi-choice markets support 3-8 outcomes</p>
            </div>
          )}

          {/* Collateral Token */}
          <div className="space-y-2">
            <Label htmlFor="collateralToken" className="text-foreground font-semibold">
              Collateral Token Address
            </Label>
            <Input
              id="collateralToken"
              placeholder="0x..."
              value={collateralToken}
              onChange={(e) => setCollateralToken(e.target.value)}
              className="bg-background/50 border-blue-500/20"
              disabled={currentStep !== "form"}
            />
            <p className="text-xs text-muted-foreground">
              Default: USDC on BNB Chain ({DEFAULT_COLLATERAL_TOKEN})
            </p>
          </div>

          {/* Creator Stake */}
          <div className="space-y-2">
            <Label htmlFor="creatorStake" className="text-foreground font-semibold">
              Creator Stake (HORIZON tokens) <span className="text-red-500">*</span>
            </Label>
            <Input
              id="creatorStake"
              type="number"
              placeholder="0.0"
              value={creatorStake}
              onChange={(e) => setCreatorStake(e.target.value)}
              min="0"
              step="0.01"
              className="bg-background/50 border-blue-500/20"
              disabled={currentStep !== "form"}
            />
            {minStake && (
              <p className="text-xs text-muted-foreground">
                Minimum: {formatUnits(minStake, 18)} HORIZON
              </p>
            )}
          </div>

          {/* Submit Button */}
          <div className="flex gap-3 pt-4">
            <Button
              type="submit"
              disabled={!isFormValid() || currentStep !== "form"}
              className="flex-1 bg-blue-600 hover:bg-blue-700 text-white"
            >
              {currentStep === "uploading" && (
                <>
                  <Upload className="w-4 h-4 mr-2 animate-spin" />
                  Uploading to IPFS...
                </>
              )}
              {currentStep === "creating" && (
                <>
                  <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                  Creating Market...
                </>
              )}
              {currentStep === "form" && (
                <>
                  <PlusCircle className="w-4 h-4 mr-2" />
                  Create Market
                </>
              )}
            </Button>
            {currentStep === "form" && (
              <Button type="button" variant="outline" onClick={handleReset}>
                Reset
              </Button>
            )}
          </div>
        </form>
      </CardContent>
    </Card>
  );
}
