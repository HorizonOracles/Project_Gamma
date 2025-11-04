// Horizon User Profile Page
// Reference: design_guidelines.md - Profile with betting history and stats
import { useEffect, useState } from "react";
import { useQuery, useMutation } from "@tanstack/react-query";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { 
  User as UserIcon, 
  Trophy, 
  TrendingUp, 
  TrendingDown,
  Activity,
  Image as ImageIcon
} from "lucide-react";
import { useAuth } from "@/hooks/useAuth";
import { useToast } from "@/hooks/use-toast";
import { apiRequest, queryClient } from "@/lib/queryClient";
import type { Bet, Wallet } from "@shared/schema";
import { useTheme } from "@/components/ThemeProvider";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";

// Tier data for level calculation
const tiers = [
  { name: "Bronze", cumulative: 0, xpPerLevel: 100 },
  { name: "Silver", cumulative: 10000, xpPerLevel: 200 },
  { name: "Gold", cumulative: 30000, xpPerLevel: 300 },
  { name: "Sapphire", cumulative: 60000, xpPerLevel: 400 },
  { name: "Emerald", cumulative: 100000, xpPerLevel: 500 },
  { name: "Ruby", cumulative: 150000, xpPerLevel: 600 },
  { name: "Diamond", cumulative: 210000, xpPerLevel: 700 },
  { name: "Pearl", cumulative: 280000, xpPerLevel: 800 },
  { name: "Opal", cumulative: 360000, xpPerLevel: 900 },
  { name: "Stardust", cumulative: 450000, xpPerLevel: 1000 },
  { name: "Nebula", cumulative: 550000, xpPerLevel: 1100 },
  { name: "Supernova", cumulative: 660000, xpPerLevel: 1200 },
];

// Calculate level from XP
const calculateLevel = (xp: number) => {
  let totalLevels = 0;
  
  for (const tier of tiers) {
    if (xp >= tier.cumulative) {
      const xpInTier = xp - tier.cumulative;
      const levelsInTier = Math.floor(xpInTier / tier.xpPerLevel);
      
      if (levelsInTier >= 100) {
        totalLevels += 100;
      } else {
        totalLevels += levelsInTier;
        break;
      }
    } else {
      break;
    }
  }
  
  return Math.min(totalLevels + 1, 1200);
};

export default function Profile() {
  const { user, isAuthenticated, isLoading } = useAuth();
  const { toast } = useToast();
  const { theme } = useTheme();
  
  // Display name and profile image state
  const [displayName, setDisplayName] = useState(user?.displayName || "");
  const [profileImageUrl, setProfileImageUrl] = useState(user?.profileImageUrl || "");

  // Update local state when user data loads
  useEffect(() => {
    if (user) {
      setDisplayName(user.displayName || "");
      setProfileImageUrl(user.profileImageUrl || "");
    }
  }, [user]);

  // Calculate days until next name change
  const getDaysUntilNextChange = () => {
    if (!user?.lastNameChange) return 0;
    const daysSinceLastChange = (Date.now() - new Date(user.lastNameChange).getTime()) / (1000 * 60 * 60 * 24);
    return Math.max(0, Math.ceil(7 - daysSinceLastChange));
  };

  // Profile update mutation
  const updateProfileMutation = useMutation({
    mutationFn: async (updates: { displayName: string; profileImageUrl?: string }) => {
      const res = await apiRequest("PATCH", "/api/profile", updates);
      if (!res.ok) {
        const error = await res.json();
        throw new Error(error.message || "Failed to update profile");
      }
      return await res.json();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["/api/auth/user"] });
      toast({
        title: "Display Name Updated",
        description: "Your display name has been successfully updated!",
      });
    },
    onError: (error: Error) => {
      toast({
        title: "Update Failed",
        description: error.message,
        variant: "destructive",
      });
    },
  });

  const handleUpdateProfile = () => {
    if (!displayName.trim()) {
      toast({
        title: "Invalid Name",
        description: "Display name cannot be empty",
        variant: "destructive",
      });
      return;
    }
    updateProfileMutation.mutate({
      displayName: displayName.trim(),
      profileImageUrl: profileImageUrl.trim() || undefined,
    });
  };

  // Helper function for rank colors - 12 tier system (Bronze first, no Obsidian)
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

  // Fetch user's bets
  const { data: bets, isLoading: betsLoading } = useQuery<any[]>({
    queryKey: ["/api/bets/my-bets"],
    enabled: isAuthenticated,
  });

  // Fetch wallet for total wagered
  const { data: wallet } = useQuery<Wallet>({
    queryKey: ["/api/wallet"],
    enabled: isAuthenticated,
  });

  // Calculate stats
  const totalBets = bets?.length || 0;
  const wonBets = bets?.filter(b => b.status === 'won').length || 0;
  const lostBets = bets?.filter(b => b.status === 'lost').length || 0;
  const activeBets = bets?.filter(b => b.status === 'active').length || 0;
  const winRate = totalBets > 0 ? ((wonBets / totalBets) * 100).toFixed(1) : "0.0";
  const totalWagered = bets?.reduce((sum, bet) => sum + parseFloat(bet.amount || "0"), 0) || 0;
  const totalWon = bets?.filter(b => b.status === 'won')
    .reduce((sum, bet) => sum + parseFloat(bet.actualPayout || "0"), 0) || 0;

  if (isLoading || betsLoading) {
    return (
      <div className="flex flex-col h-full">
        <LiveBettingFeed />
        <div className="flex-1 overflow-auto relative">
          <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
            <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
          </div>
          <div className="container mx-auto p-6 max-w-5xl space-y-6 relative z-10">
            <Skeleton className="h-12 w-64" />
            <Skeleton className="h-48" />
            <Skeleton className="h-96" />
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
        <div className="container mx-auto p-6 max-w-5xl space-y-6 relative z-10">
        <h1 className="text-3xl font-bold text-foreground font-sohne">
          My <span className="text-primary">Profile</span>
        </h1>

      {/* Profile Header */}
      <Card className="bg-gradient-to-br from-card to-accent/5 border-card-border">
        <CardContent className="p-6">
          <div className="flex items-start gap-6">
            <Avatar className="h-24 w-24 border-4 border-accent ring-4 ring-accent/30">
              <AvatarImage src={user?.profileImageUrl || ""} alt={user?.displayName || "User"} />
              <AvatarFallback className="text-5xl font-bold text-primary bg-primary/10">
                {(user?.displayName || user?.email)?.[0]?.toUpperCase() || "U"}
              </AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <h2 className="text-2xl font-bold" style={{ color: getRankColor(user?.rank || "Bronze") }}>
                {user?.displayName || user?.email?.split('@')[0] || "User"}
              </h2>
              <p className="text-muted-foreground">{user?.email}</p>
              <div className="flex gap-2 mt-3 flex-wrap">
                <Badge variant="outline" className="border-primary text-primary">
                  <TrendingUp className="h-3 w-3 mr-1" />
                  {user?.rankPoints || 0} Points
                </Badge>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Display Name Customization */}
      <Card className="bg-card border-card-border">
        <CardHeader className="border-b border-accent">
          <CardTitle className="text-foreground flex items-center gap-2 font-sohne">
            <UserIcon className="h-5 w-5 text-primary" />
            Display Name
          </CardTitle>
          <CardDescription>
            Change your display name. You can update it once every 7 days.
          </CardDescription>
        </CardHeader>
        <CardContent className="p-6 space-y-6">
          {/* Display Name Input */}
          <div className="space-y-2">
            <Label htmlFor="display-name" className="flex items-center gap-2">
              <UserIcon className="h-4 w-4 text-primary" />
              Display Name
            </Label>
            <Input
              id="display-name"
              type="text"
              placeholder="Enter display name"
              value={displayName}
              onChange={(e) => setDisplayName(e.target.value)}
              className="bg-background border-accent"
              data-testid="input-display-name"
              maxLength={20}
            />
            <p className="text-xs text-muted-foreground">
              {displayName.length}/20 characters • Your name color matches your rank tier
            </p>
          </div>

          {/* Profile Image URL Input */}
          <div className="space-y-2">
            <Label htmlFor="profile-image" className="flex items-center gap-2">
              <ImageIcon className="h-4 w-4 text-primary" />
              Profile Picture URL
            </Label>
            <Input
              id="profile-image"
              type="url"
              placeholder="https://example.com/your-image.jpg"
              value={profileImageUrl}
              onChange={(e) => setProfileImageUrl(e.target.value)}
              className="bg-background border-accent"
              data-testid="input-profile-image"
            />
            <p className="text-xs text-muted-foreground">
              Enter a URL to your profile picture (optional)
            </p>
          </div>

          {/* Cooldown Info */}
          {user?.lastNameChange && getDaysUntilNextChange() > 0 && (
            <div className="p-4 rounded-md bg-destructive/10 border border-destructive/30 flex items-center gap-3">
              <Activity className="h-5 w-5 text-destructive" />
              <div>
                <p className="text-sm font-semibold text-destructive">Cooldown Active</p>
                <p className="text-xs text-muted-foreground">
                  You can change your name again in {getDaysUntilNextChange()} day{getDaysUntilNextChange() !== 1 ? 's' : ''}
                </p>
              </div>
            </div>
          )}

          {/* Save Button */}
          <Button
            onClick={handleUpdateProfile}
            disabled={updateProfileMutation.isPending || !!(user?.lastNameChange && getDaysUntilNextChange() > 0)}
            className="w-full"
            data-testid="button-save-display-name"
          >
            {updateProfileMutation.isPending ? "Saving..." : "Update Display Name"}
          </Button>
        </CardContent>
      </Card>

      {/* Stats Grid */}
      <div className="grid md:grid-cols-4 gap-4">
        <Card className="bg-card border-card-border">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Win Rate
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-mono font-bold text-primary">{winRate}%</p>
            <p className="text-xs text-muted-foreground mt-1">
              {wonBets} wins / {totalBets} bets
            </p>
          </CardContent>
        </Card>

        <Card className="bg-card border-card-border">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Total Wagered
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-mono font-bold text-accent">
              {totalWagered.toFixed(2)}
            </p>
            <p className="text-xs text-muted-foreground mt-1">BNB</p>
          </CardContent>
        </Card>

        <Card className="bg-card border-card-border">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-muted-foreground flex items-center gap-1">
              <Trophy className="h-4 w-4 text-primary" />
              Total Won
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-mono font-bold text-primary">
              {totalWon.toFixed(2)}
            </p>
            <p className="text-xs text-muted-foreground mt-1">BNB</p>
          </CardContent>
        </Card>

        <Card className="bg-card border-card-border">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Active Bets
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-mono font-bold text-foreground">{activeBets}</p>
            <p className="text-xs text-muted-foreground mt-1">In progress</p>
          </CardContent>
        </Card>
      </div>

      {/* Betting History */}
      <Card className="bg-card border-card-border">
        <CardHeader className="border-b border-accent">
          <CardTitle className="text-foreground">Betting History</CardTitle>
        </CardHeader>
        <CardContent className="p-0">
          {bets && bets.length > 0 ? (
            <div className="divide-y divide-border">
              {bets.map((bet) => (
                <div
                  key={bet.id}
                  className="p-4 hover-elevate flex items-center justify-between"
                  data-testid={`bet-${bet.id}`}
                >
                  <div className="flex-1">
                    <div className="flex items-center gap-3 mb-2">
                      <Badge
                        variant={
                          bet.status === 'won' 
                            ? 'default' 
                            : bet.status === 'lost'
                            ? 'outline'
                            : 'secondary'
                        }
                        className={
                          bet.status === 'won'
                            ? 'bg-primary text-primary-foreground'
                            : bet.status === 'lost'
                            ? 'border-destructive text-destructive'
                            : 'border-accent text-accent'
                        }
                      >
                        {bet.status === 'won' && <TrendingUp className="h-3 w-3 mr-1" />}
                        {bet.status === 'lost' && <TrendingDown className="h-3 w-3 mr-1" />}
                        {bet.status.toUpperCase()}
                      </Badge>
                      <span className="text-sm text-muted-foreground">
                        {new Date(bet.createdAt).toLocaleDateString()}
                      </span>
                    </div>
                    <p className="text-foreground font-medium">
                      {bet.marketDescription}
                    </p>
                    <p className="text-sm text-muted-foreground mt-1">
                      Bet <span className="font-mono font-bold text-primary">{parseFloat(bet.amount).toFixed(4)} BNB</span>
                      {' '}on{' '}
                      <span className="font-semibold text-primary">
                        {bet.outcome === 'A' ? bet.teamA : bet.teamB}
                      </span>
                      {' '}@ <span className="font-mono text-accent">{parseFloat(bet.oddsAtBet).toFixed(2)}</span> odds
                    </p>
                  </div>
                  {bet.status === 'won' && bet.actualPayout && (
                    <div className="text-right">
                      <p className="text-sm text-muted-foreground">Won</p>
                      <p className="font-mono font-bold text-primary">
                        +{parseFloat(bet.actualPayout).toFixed(4)} BNB
                      </p>
                    </div>
                  )}
                </div>
              ))}
            </div>
          ) : (
            <div className="p-12 text-center text-muted-foreground">
              <Trophy className="h-12 w-12 mx-auto mb-3 opacity-50" />
              <p>No bets placed yet</p>
              <p className="text-sm mt-1">Start betting to build your history!</p>
            </div>
          )}
        </CardContent>
      </Card>
        </div>
      </div>
    </div>
  );
}
