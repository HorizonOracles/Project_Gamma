// Horizon Leaderboard Page - Player Rankings
// Shows top players by rank points with profile pictures and colored names
import { useQuery } from "@tanstack/react-query";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Trophy, TrendingUp, Flame, Coins } from "lucide-react";
import type { User } from "@shared/schema";
import { useTheme } from "@/components/ThemeProvider";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";

export default function Leaderboard() {
  const { theme } = useTheme();
  const { data: leaders, isLoading } = useQuery<User[]>({
    queryKey: ["/api/leaderboard"],
  });

  const getRankColor = (rank: string) => {
    switch (rank) {
      case "Bronze": return "#CD7F32";
      case "Silver": return "#C0C0C0";
      case "Gold": return "#fe5a20";
      case "Sapphire": return "#0F52BA";
      case "Emerald": return "#50C878";
      case "Ruby": return "#E0115F";
      case "Diamond": return "#B9F2FF";
      case "Pearl": return "#F0EAD6";
      case "Opal": return "#A8C3BC";
      case "Stardust": return "#FFE4B5";
      case "Nebula": return "#9D4EDD";
      case "Supernova": return "#FF6B35";
      default: return "#CD7F32";
    }
  };

  const getRankIcon = (rank: string) => {
    const topTiers = ["Supernova", "Nebula", "Stardust", "Opal", "Pearl", "Diamond"];
    if (topTiers.includes(rank)) {
      return <Trophy className="h-5 w-5" style={{ color: getRankColor(rank) }} />;
    }
    return <Flame className="h-5 w-5" style={{ color: getRankColor(rank) }} />;
  };

  if (isLoading) {
    return (
      <div className="flex flex-col h-full">
        <LiveBettingFeed />
        <div className="flex-1 overflow-auto relative">
          <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
            <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
          </div>
          <div className="container mx-auto p-6 max-w-5xl relative z-10">
            <div className="space-y-4">
              {[1, 2, 3, 4, 5].map((i) => (
                <Skeleton key={i} className="h-24 bg-card" />
              ))}
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      {/* Live Betting Feed - Horizontal Ticker */}
      <LiveBettingFeed />
      
      <div className="flex-1 overflow-auto relative">
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
        </div>
        <div className="container mx-auto p-6 max-w-5xl relative z-10">
          <div className="space-y-6">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-4xl font-bold text-foreground font-sohne">
              Player <span className="text-primary">Leaderboard</span>
            </h1>
            <p className="text-muted-foreground mt-2">
              Top players ranked by profit and loss (winnings minus wagers)
            </p>
          </div>
          <TrendingUp className="h-12 w-12 text-primary" />
        </div>


        {/* Leaderboard List */}
        <div className="space-y-3">
          {leaders && leaders.length > 0 ? (
            leaders.map((leader, index) => (
              <Card 
                key={leader.id} 
                className="hover-elevate transition-all"
                data-testid={`leaderboard-entry-${index}`}
              >
                <CardContent className="p-6">
                  <div className="flex items-center gap-4">
                    {/* Rank Position */}
                    <div className="flex-shrink-0 w-12 text-center">
                      {index < 3 ? (
                        <div className="text-3xl font-bold" style={{ 
                          color: index === 0 ? "#fe5a20" : index === 1 ? "#C0C0C0" : "#CD7F32" 
                        }}>
                          #{index + 1}
                        </div>
                      ) : (
                        <div className="text-2xl font-semibold text-muted-foreground">
                          #{index + 1}
                        </div>
                      )}
                    </div>

                    {/* Profile Initial */}
                    <div className="flex items-center justify-center h-16 w-16 rounded-full bg-primary/10 border-2 border-accent ring-2 ring-accent/30">
                      <span className="text-2xl font-bold text-primary">
                        {leader.displayName?.[0]?.toUpperCase() || "P"}
                      </span>
                    </div>

                    {/* Player Info */}
                    <div className="flex-1 min-w-0">
                      <div className="flex items-center gap-2 flex-wrap">
                        <h3 
                          className="text-xl font-bold truncate" 
                          style={{ color: getRankColor(leader.rank) }}
                          data-testid={`player-name-${index}`}
                        >
                          {leader.displayName || "Anonymous Player"}
                        </h3>
                      </div>
                      <div className="flex items-center gap-4 mt-2 flex-wrap">
                        <div className="flex items-center gap-1 text-sm">
                          <TrendingUp className="h-4 w-4 text-primary" />
                          <span className={`font-bold text-lg ${
                            parseFloat(leader.totalWon) - parseFloat(leader.totalWagered) >= 0 
                              ? 'text-green-500' 
                              : 'text-red-500'
                          }`}>
                            {parseFloat(leader.totalWon) - parseFloat(leader.totalWagered) >= 0 ? '+' : ''}
                            {(parseFloat(leader.totalWon) - parseFloat(leader.totalWagered)).toFixed(2)}
                          </span>
                          <span className="text-muted-foreground">BNB profit</span>
                        </div>
                        <div className="flex items-center gap-1 text-sm">
                          <Coins className="h-4 w-4 text-accent" />
                          <span className="font-semibold text-accent">{parseFloat(leader.totalWagered).toFixed(2)}</span>
                          <span className="text-muted-foreground">wagered</span>
                        </div>
                        <div className="flex items-center gap-1 text-sm">
                          <Trophy className="h-4 w-4 text-primary" />
                          <span className="font-semibold text-primary">{parseFloat(leader.totalWon).toFixed(2)}</span>
                          <span className="text-muted-foreground">won</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))
          ) : (
            <Card>
              <CardContent className="p-12 text-center">
                <p className="text-muted-foreground">No players ranked yet. Start betting to appear on the leaderboard!</p>
              </CardContent>
            </Card>
          )}
          </div>
        </div>
      </div>
    </div>
    </div>
  );
}
