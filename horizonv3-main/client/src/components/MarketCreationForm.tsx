// Market Creation Form Component
// Allows admin to create new prediction markets on blockchain
import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { useToast } from "@/hooks/use-toast";
import { Plus, Loader2, Calendar, Shield } from "lucide-react";
import { useCreateMarket } from "@project-gamma/sdk";
import { parseUnits } from "viem";
import { useAccount } from "wagmi";

// Form schema for market creation
const marketFormSchema = z.object({
  sport: z.string().min(2, "Sport is required"),
  league: z.string().min(2, "League is required"),
  teamA: z.string().min(2, "Team A is required"),
  teamB: z.string().min(2, "Team B is required"),
  teamALogo: z.string().url("Must be a valid URL").optional().or(z.literal("")),
  teamBLogo: z.string().url("Must be a valid URL").optional().or(z.literal("")),
  description: z.string().min(10, "Description must be at least 10 characters"),
  gameTime: z.string().min(1, "Game time is required"),
  category: z.string().min(2, "Category is required").default("sports"),
  creatorStake: z.string().min(1, "Creator stake is required").default("1000"),
});

type MarketFormData = z.infer<typeof marketFormSchema>;

interface MarketCreationFormProps {
  onSuccess?: (marketId: bigint) => void;
  collateralTokenAddress: string; // Address of the collateral token (e.g., USDC)
}

export function MarketCreationForm({ onSuccess, collateralTokenAddress }: MarketCreationFormProps) {
  const { toast } = useToast();
  const { address } = useAccount();
  const [isCreating, setIsCreating] = useState(false);
  
  const { write: createMarket, isLoading: isWaitingForTx, data: marketId } = useCreateMarket();
  
  const form = useForm<MarketFormData>({
    resolver: zodResolver(marketFormSchema),
    defaultValues: {
      sport: "",
      league: "",
      teamA: "",
      teamB: "",
      teamALogo: "",
      teamBLogo: "",
      description: "",
      gameTime: "",
      category: "sports",
      creatorStake: "1000",
    },
  });

  const onSubmit = async (data: MarketFormData) => {
    if (!address) {
      toast({
        title: "Wallet not connected",
        description: "Please connect your wallet to create a market",
        variant: "destructive",
      });
      return;
    }

    setIsCreating(true);

    try {
      // Convert game time to Unix timestamp
      const gameTimeDate = new Date(data.gameTime);
      const endTime = BigInt(Math.floor(gameTimeDate.getTime() / 1000));

      // Prepare metadata for IPFS (simplified - in production, upload to IPFS first)
      const question = `${data.teamA} vs ${data.teamB} - Match Winner`;
      const metadata = {
        question,
        description: data.description,
        sport: data.sport,
        league: data.league,
        teamA: data.teamA,
        teamB: data.teamB,
        teamALogo: data.teamALogo || undefined,
        teamBLogo: data.teamBLogo || undefined,
        gameTime: data.gameTime,
      };

      // For now, use a temporary metadata URI (in production, upload to IPFS)
      // TODO: Integrate with IPFS upload
      const metadataURI = `temp://${JSON.stringify(metadata)}`;

      // Parse creator stake (assuming 18 decimals for HORIZON token)
      const creatorStake = parseUnits(data.creatorStake, 18);

      // Create market on blockchain
      createMarket({
        question,
        endTime,
        collateralToken: collateralTokenAddress as `0x${string}`,
        category: data.category,
        metadataURI,
        creatorStake,
        marketType: 0, // Binary market
        outcomeCount: 2,
        liquidityParameter: 0n,
      });

      // Save to database
      const response = await fetch("/api/admin/markets", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-wallet-address": address,
        },
        body: JSON.stringify({
          sport: data.sport,
          league: data.league,
          teamA: data.teamA,
          teamB: data.teamB,
          teamALogo: data.teamALogo || null,
          teamBLogo: data.teamBLogo || null,
          description: data.description,
          gameTime: gameTimeDate.toISOString(),
          marketType: "match_winner",
          category: data.category,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to save market to database");
      }

      const result = await response.json();

      toast({
        title: "Market Created!",
        description: `Market "${question}" has been created successfully`,
      });

      form.reset();
      onSuccess?.(marketId!);
    } catch (error: any) {
      console.error("Error creating market:", error);
      toast({
        title: "Failed to create market",
        description: error.message || "An error occurred while creating the market",
        variant: "destructive",
      });
    } finally {
      setIsCreating(false);
    }
  };

  return (
    <Card className="border-accent/20">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-foreground font-sohne">
          <Plus className="h-5 w-5 text-primary" />
          Create New Market
        </CardTitle>
        <CardDescription className="text-muted-foreground">
          Create a new prediction market on the blockchain
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {/* Sport */}
            <div className="space-y-2">
              <Label htmlFor="sport">Sport</Label>
              <Input
                id="sport"
                placeholder="e.g., Basketball, Soccer, Tennis"
                {...form.register("sport")}
              />
              {form.formState.errors.sport && (
                <p className="text-xs text-destructive">{form.formState.errors.sport.message}</p>
              )}
            </div>

            {/* League */}
            <div className="space-y-2">
              <Label htmlFor="league">League</Label>
              <Input
                id="league"
                placeholder="e.g., NBA, Premier League, ATP"
                {...form.register("league")}
              />
              {form.formState.errors.league && (
                <p className="text-xs text-destructive">{form.formState.errors.league.message}</p>
              )}
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {/* Team A */}
            <div className="space-y-2">
              <Label htmlFor="teamA">Team A / Player A</Label>
              <Input
                id="teamA"
                placeholder="e.g., Los Angeles Lakers"
                {...form.register("teamA")}
              />
              {form.formState.errors.teamA && (
                <p className="text-xs text-destructive">{form.formState.errors.teamA.message}</p>
              )}
            </div>

            {/* Team B */}
            <div className="space-y-2">
              <Label htmlFor="teamB">Team B / Player B</Label>
              <Input
                id="teamB"
                placeholder="e.g., Boston Celtics"
                {...form.register("teamB")}
              />
              {form.formState.errors.teamB && (
                <p className="text-xs text-destructive">{form.formState.errors.teamB.message}</p>
              )}
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {/* Team A Logo */}
            <div className="space-y-2">
              <Label htmlFor="teamALogo">Team A Logo URL (Optional)</Label>
              <Input
                id="teamALogo"
                type="url"
                placeholder="https://..."
                {...form.register("teamALogo")}
              />
              {form.formState.errors.teamALogo && (
                <p className="text-xs text-destructive">{form.formState.errors.teamALogo.message}</p>
              )}
            </div>

            {/* Team B Logo */}
            <div className="space-y-2">
              <Label htmlFor="teamBLogo">Team B Logo URL (Optional)</Label>
              <Input
                id="teamBLogo"
                type="url"
                placeholder="https://..."
                {...form.register("teamBLogo")}
              />
              {form.formState.errors.teamBLogo && (
                <p className="text-xs text-destructive">{form.formState.errors.teamBLogo.message}</p>
              )}
            </div>
          </div>

          {/* Description */}
          <div className="space-y-2">
            <Label htmlFor="description">Market Description</Label>
            <Textarea
              id="description"
              placeholder="Describe the market conditions, rules, and resolution criteria..."
              rows={4}
              {...form.register("description")}
            />
            {form.formState.errors.description && (
              <p className="text-xs text-destructive">{form.formState.errors.description.message}</p>
            )}
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {/* Game Time */}
            <div className="space-y-2">
              <Label htmlFor="gameTime" className="flex items-center gap-2">
                <Calendar className="h-4 w-4" />
                Game Time (Market Close Time)
              </Label>
              <Input
                id="gameTime"
                type="datetime-local"
                {...form.register("gameTime")}
              />
              {form.formState.errors.gameTime && (
                <p className="text-xs text-destructive">{form.formState.errors.gameTime.message}</p>
              )}
            </div>

            {/* Creator Stake */}
            <div className="space-y-2">
              <Label htmlFor="creatorStake" className="flex items-center gap-2">
                <Shield className="h-4 w-4" />
                Creator Stake (HORIZON Tokens)
              </Label>
              <Input
                id="creatorStake"
                type="number"
                min="0"
                step="100"
                placeholder="1000"
                {...form.register("creatorStake")}
              />
              {form.formState.errors.creatorStake && (
                <p className="text-xs text-destructive">{form.formState.errors.creatorStake.message}</p>
              )}
            </div>
          </div>

          {/* Category (hidden, defaults to sports) */}
          <Input type="hidden" {...form.register("category")} />

          {/* Submit Button */}
          <Button
            type="submit"
            className="w-full"
            disabled={isCreating || isWaitingForTx}
          >
            {isCreating || isWaitingForTx ? (
              <>
                <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                Creating Market...
              </>
            ) : (
              <>
                <Plus className="h-4 w-4 mr-2" />
                Create Market
              </>
            )}
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
