// Horizon Activity Feed Page
import { useQuery } from "@tanstack/react-query";
import { BettingFeed } from "@/components/betting-feed";
import { Skeleton } from "@/components/ui/skeleton";
import { useEffect } from "react";
import { useAuth } from "@/hooks/useAuth";
import { useToast } from "@/hooks/use-toast";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";

export default function Feed() {
  const { isAuthenticated, isLoading } = useAuth();
  const { toast } = useToast();

  // Redirect if not authenticated
  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      toast({
        title: "Unauthorized",
        description: "You are logged out. Logging in again...",
        variant: "destructive",
      });
      setTimeout(() => {
        window.location.href = "/api/login";
      }, 500);
    }
  }, [isAuthenticated, isLoading, toast]);

  // Fetch betting feed
  const { data: recentBets, isLoading: betsLoading } = useQuery<any[]>({
    queryKey: ["/api/bets/feed"],
    refetchInterval: 3000, // Refetch every 3 seconds
  });

  // Transform bets for feed
  const feedBets = recentBets?.map(bet => ({
    id: bet.id,
    userName: bet.userDisplayName || "Anonymous Player",
    amount: parseFloat(bet.amount || "0").toFixed(4),
    outcome: bet.outcome === 'A' ? bet.teamA : bet.teamB,
    marketDescription: bet.marketDescription || "",
    odds: parseFloat(bet.oddsAtBet || "0").toFixed(2),
    timestamp: new Date(bet.createdAt), // Parse ISO string to Date
    status: bet.status as any,
  })) || [];

  return (
    <div className="flex flex-col h-full">
      {/* Live Betting Feed - Horizontal Ticker */}
      <LiveBettingFeed />
      
      <div className="flex-1 overflow-auto relative">
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
        </div>
        <div className="container mx-auto p-6 max-w-5xl relative z-10">
        <h1 className="text-3xl font-bold text-foreground mb-6 font-sohne">
          Activity <span className="text-primary">Feed</span>
        </h1>

        {betsLoading ? (
          <Skeleton className="h-[600px] bg-card" />
        ) : (
          <BettingFeed bets={feedBets} />
        )}
        </div>
      </div>
    </div>
  );
}
