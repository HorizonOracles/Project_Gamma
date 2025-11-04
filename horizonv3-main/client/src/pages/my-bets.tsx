// My Bets Page - User's Betting History
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "@/hooks/useAuth";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";
import { Skeleton } from "@/components/ui/skeleton";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle, Trophy, TrendingUp, XCircle, Ban } from "lucide-react";
import { LoadingSpinner } from "@/components/LoadingSpinner";
import { NetworkBackground } from "@/components/NetworkBackground";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { format } from "date-fns";
import { getLeagueBadge } from "@/lib/leagueUtils";

interface UserBet {
  id: string;
  userId: string;
  marketId: string;
  outcome: 'A' | 'B';
  amount: string;
  oddsAtBet: string;
  status: 'active' | 'won' | 'lost' | 'voided';
  actualPayout: string | null;
  createdAt: string;
  settledAt: string | null;
  teamA: string;
  teamB: string;
  teamALogo: string | null;
  teamBLogo: string | null;
  marketDescription: string;
  sport: string;
  league: string;
}

export default function MyBets() {
  const { isAuthenticated, isLoading: authLoading } = useAuth();

  const { data: bets, isLoading: betsLoading, error: betsError } = useQuery<UserBet[]>({
    queryKey: ["/api/bets/my-bets"],
    enabled: isAuthenticated,
    retry: false,
  });

  if (authLoading) {
    return (
      <div className="flex items-center justify-center h-screen bg-background">
        <div className="text-center">
          <LoadingSpinner />
          <p className="text-lg text-muted-foreground font-sohne">Loading...</p>
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return (
      <div className="flex items-center justify-center h-screen bg-background">
        <div className="text-center">
          <AlertCircle className="h-16 w-16 text-muted-foreground mx-auto mb-4" />
          <p className="text-lg text-foreground mb-2">Login Required</p>
          <p className="text-sm text-muted-foreground">Please login to view your bets</p>
        </div>
      </div>
    );
  }

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'won':
        return <Badge className="bg-green-500/20 text-green-400 border-green-500" data-testid="badge-status-won">Won</Badge>;
      case 'lost':
        return <Badge className="bg-red-500/20 text-red-400 border-red-500" data-testid="badge-status-lost">Lost</Badge>;
      case 'active':
        return <Badge className="bg-blue-500/20 text-blue-400 border-blue-500" data-testid="badge-status-active">Active</Badge>;
      case 'voided':
        return <Badge className="bg-gray-500/20 text-gray-400 border-gray-500" data-testid="badge-status-voided">Voided</Badge>;
      default:
        return <Badge variant="outline">{status}</Badge>;
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'won':
        return <Trophy className="h-5 w-5 text-green-400" />;
      case 'lost':
        return <XCircle className="h-5 w-5 text-red-400" />;
      case 'active':
        return <TrendingUp className="h-5 w-5 text-blue-400" />;
      case 'voided':
        return <Ban className="h-5 w-5 text-gray-400" />;
      default:
        return null;
    }
  };

  return (
    <div className="flex flex-col h-full">
      <LiveBettingFeed />
      
      <div className="flex-1 overflow-auto relative">
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
        </div>
        <div className="container mx-auto p-6 max-w-7xl relative z-10">
          <div className="space-y-6">
            <div className="flex items-center justify-between">
              <h2 className="text-3xl font-bold text-foreground font-sohne">
                MY <span className="text-primary ml-1">BETS</span>
              </h2>
            </div>

            {betsError && (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>Failed to load your bets. Please try again later.</AlertDescription>
              </Alert>
            )}

            {betsLoading ? (
              <div className="space-y-4">
                {[1, 2, 3, 4].map((i) => (
                  <Skeleton key={i} className="h-48 bg-card" />
                ))}
              </div>
            ) : bets && bets.length > 0 ? (
              <>
                {/* Active Bets Section */}
                {bets.filter(b => b.status === 'active').length > 0 && (
                  <div className="space-y-4">
                    <div className="flex items-center gap-3">
                      <TrendingUp className="h-6 w-6 text-primary" />
                      <h3 className="text-xl font-bold text-foreground font-sohne">
                        LIVE <span className="text-primary ml-1">BETS</span>
                      </h3>
                      <Badge className="bg-primary/20 text-primary border-primary" data-testid="badge-live-count">
                        {bets.filter(b => b.status === 'active').length} Active
                      </Badge>
                    </div>
                    {bets.filter(b => b.status === 'active').map((bet) => {
                      const teamName = bet.outcome === 'A' ? bet.teamA : bet.teamB;
                      const potentialPayout = parseFloat(bet.amount) * parseFloat(bet.oddsAtBet);

                      const leagueBadge = getLeagueBadge(bet.league);
                      
                      return (
                        <Card key={bet.id} className="hover-elevate border-primary/30 bg-card/80" data-testid={`card-bet-${bet.id}`}>
                          <CardHeader className="pb-3">
                            <div className="flex items-start justify-between gap-4">
                              <div className="flex-1">
                                <div className="flex items-center gap-2 mb-3">
                                  <Badge variant="outline" className="text-sm font-mono">
                                    {bet.sport}
                                  </Badge>
                                  <div className="flex items-center gap-2">
                                    {leagueBadge && (
                                      <img 
                                        src={leagueBadge} 
                                        alt={bet.league}
                                        className="h-10 w-10 object-contain"
                                      />
                                    )}
                                    <Badge variant="outline" className="text-sm font-mono text-muted-foreground">
                                      {bet.league}
                                    </Badge>
                                  </div>
                                </div>
                                <div className="flex items-center gap-4 mb-2">
                                  <div className="flex items-center gap-2.5">
                                    {bet.teamALogo && (
                                      <img 
                                        src={bet.teamALogo} 
                                        alt={bet.teamA}
                                        className="h-10 w-10 object-contain"
                                      />
                                    )}
                                    <span className="font-bold text-foreground text-base">{bet.teamA}</span>
                                  </div>
                                  <span className="text-muted-foreground font-bold">vs</span>
                                  <div className="flex items-center gap-2.5">
                                    {bet.teamBLogo && (
                                      <img 
                                        src={bet.teamBLogo} 
                                        alt={bet.teamB}
                                        className="h-10 w-10 object-contain"
                                      />
                                    )}
                                    <span className="font-bold text-foreground text-base">{bet.teamB}</span>
                                  </div>
                                </div>
                                <p className="text-sm text-muted-foreground line-clamp-2">
                                  {bet.marketDescription}
                                </p>
                              </div>
                              <div className="flex items-center gap-2">
                                <TrendingUp className="h-5 w-5 text-blue-400" />
                                <Badge className="bg-blue-500/20 text-blue-400 border-blue-500" data-testid="badge-status-active">Active</Badge>
                              </div>
                            </div>
                          </CardHeader>
                          <CardContent>
                            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                              <div>
                                <p className="text-xs text-muted-foreground mb-1">Bet On</p>
                                <p className="text-sm font-bold text-accent" data-testid={`text-bet-team-${bet.id}`}>
                                  {teamName}
                                </p>
                              </div>
                              <div>
                                <p className="text-xs text-muted-foreground mb-1">Amount</p>
                                <p className="text-sm font-bold font-mono text-foreground" data-testid={`text-bet-amount-${bet.id}`}>
                                  {parseFloat(bet.amount).toFixed(4)} BNB
                                </p>
                              </div>
                              <div>
                                <p className="text-xs text-muted-foreground mb-1">Odds</p>
                                <p className="text-sm font-bold font-mono text-foreground" data-testid={`text-bet-odds-${bet.id}`}>
                                  {parseFloat(bet.oddsAtBet).toFixed(2)}x
                                </p>
                              </div>
                              <div>
                                <p className="text-xs text-muted-foreground mb-1">Potential Win</p>
                                <p className="text-sm font-bold font-mono text-primary" data-testid={`text-bet-payout-${bet.id}`}>
                                  {potentialPayout.toFixed(4)} BNB
                                </p>
                              </div>
                            </div>
                            <div className="mt-4 pt-4 border-t border-card-border flex items-center justify-between">
                              <div>
                                <p className="text-xs text-muted-foreground">
                                  Placed: {format(new Date(bet.createdAt), 'MMM dd, yyyy HH:mm')}
                                </p>
                              </div>
                              <div className="text-right">
                                <p className="text-xs text-muted-foreground">Potential Profit</p>
                                <p className="text-sm font-bold text-primary font-mono">
                                  +{(potentialPayout - parseFloat(bet.amount)).toFixed(4)} BNB
                                </p>
                              </div>
                            </div>
                          </CardContent>
                        </Card>
                      );
                    })}
                  </div>
                )}

                {/* Settled Bets Section */}
                {bets.filter(b => b.status !== 'active').length > 0 && (
                  <div className="space-y-4">
                    <div className="flex items-center gap-3 mt-8">
                      <Trophy className="h-6 w-6 text-muted-foreground" />
                      <h3 className="text-xl font-bold text-foreground font-sohne">
                        BET <span className="text-muted-foreground ml-1">HISTORY</span>
                      </h3>
                    </div>
                    {bets.filter(b => b.status !== 'active').map((bet) => {
                  const teamName = bet.outcome === 'A' ? bet.teamA : bet.teamB;
                  const potentialPayout = bet.status === 'won' && bet.actualPayout 
                    ? parseFloat(bet.actualPayout)
                    : parseFloat(bet.amount) * parseFloat(bet.oddsAtBet);
                  const leagueBadge = getLeagueBadge(bet.league);

                  return (
                    <Card key={bet.id} className="hover-elevate" data-testid={`card-bet-${bet.id}`}>
                      <CardHeader className="pb-3">
                        <div className="flex items-start justify-between gap-4">
                          <div className="flex-1">
                            <div className="flex items-center gap-2 mb-3">
                              <Badge variant="outline" className="text-sm font-mono">
                                {bet.sport}
                              </Badge>
                              <div className="flex items-center gap-2">
                                {leagueBadge && (
                                  <img 
                                    src={leagueBadge} 
                                    alt={bet.league}
                                    className="h-6 w-6 object-contain"
                                  />
                                )}
                                <Badge variant="outline" className="text-sm font-mono text-muted-foreground">
                                  {bet.league}
                                </Badge>
                              </div>
                            </div>
                            <div className="flex items-center gap-4 mb-2">
                              <div className="flex items-center gap-2.5">
                                {bet.teamALogo && (
                                  <img 
                                    src={bet.teamALogo} 
                                    alt={bet.teamA}
                                    className="h-10 w-10 object-contain"
                                  />
                                )}
                                <span className="font-bold text-foreground text-base">{bet.teamA}</span>
                              </div>
                              <span className="text-muted-foreground font-bold">vs</span>
                              <div className="flex items-center gap-2.5">
                                {bet.teamBLogo && (
                                  <img 
                                    src={bet.teamBLogo} 
                                    alt={bet.teamB}
                                    className="h-10 w-10 object-contain"
                                  />
                                )}
                                <span className="font-bold text-foreground text-base">{bet.teamB}</span>
                              </div>
                            </div>
                            <p className="text-sm text-muted-foreground line-clamp-2">
                              {bet.marketDescription}
                            </p>
                          </div>
                          <div className="flex items-center gap-2">
                            {getStatusIcon(bet.status)}
                            {getStatusBadge(bet.status)}
                          </div>
                        </div>
                      </CardHeader>
                      <CardContent>
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                          <div>
                            <p className="text-xs text-muted-foreground mb-1">Bet On</p>
                            <p className="text-sm font-bold text-accent" data-testid={`text-bet-team-${bet.id}`}>
                              {teamName}
                            </p>
                          </div>
                          <div>
                            <p className="text-xs text-muted-foreground mb-1">Amount</p>
                            <p className="text-sm font-bold font-mono text-foreground" data-testid={`text-bet-amount-${bet.id}`}>
                              {parseFloat(bet.amount).toFixed(4)} BNB
                            </p>
                          </div>
                          <div>
                            <p className="text-xs text-muted-foreground mb-1">Odds</p>
                            <p className="text-sm font-bold font-mono text-foreground" data-testid={`text-bet-odds-${bet.id}`}>
                              {parseFloat(bet.oddsAtBet).toFixed(2)}x
                            </p>
                          </div>
                          <div>
                            <p className="text-xs text-muted-foreground mb-1">
                              {bet.status === 'won' ? 'Won' : bet.status === 'active' ? 'Potential Win' : 'Payout'}
                            </p>
                            <p className={`text-sm font-bold font-mono ${
                              bet.status === 'won' ? 'text-green-400' : 
                              bet.status === 'lost' ? 'text-red-400' : 
                              'text-foreground'
                            }`} data-testid={`text-bet-payout-${bet.id}`}>
                              {potentialPayout.toFixed(4)} BNB
                            </p>
                          </div>
                        </div>
                        <div className="mt-4 pt-4 border-t border-card-border flex items-center justify-between">
                          <div>
                            <p className="text-xs text-muted-foreground">
                              Placed: {format(new Date(bet.createdAt), 'MMM dd, yyyy HH:mm')}
                            </p>
                            {bet.settledAt && (
                              <p className="text-xs text-muted-foreground mt-1">
                                Settled: {format(new Date(bet.settledAt), 'MMM dd, yyyy HH:mm')}
                              </p>
                            )}
                          </div>
                          {bet.status === 'won' && bet.actualPayout && (
                            <div className="text-right">
                              <p className="text-xs text-muted-foreground">Profit</p>
                              <p className="text-sm font-bold text-green-400 font-mono">
                                +{(parseFloat(bet.actualPayout) - parseFloat(bet.amount)).toFixed(4)} BNB
                              </p>
                            </div>
                          )}
                          {bet.status === 'lost' && (
                            <div className="text-right">
                              <p className="text-xs text-muted-foreground">Loss</p>
                              <p className="text-sm font-bold text-red-400 font-mono">
                                -{parseFloat(bet.amount).toFixed(4)} BNB
                              </p>
                            </div>
                          )}
                        </div>
                      </CardContent>
                    </Card>
                      );
                    })}
                  </div>
                )}
              </>
            ) : (
              <div className="text-center py-12 bg-card rounded-lg border border-card-border">
                <Trophy className="h-16 w-16 text-muted-foreground mx-auto mb-4 opacity-50" />
                <p className="text-muted-foreground text-lg">No bets yet</p>
                <p className="text-sm text-muted-foreground mt-2">
                  Start betting on your favorite markets to see your bet history here!
                </p>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
