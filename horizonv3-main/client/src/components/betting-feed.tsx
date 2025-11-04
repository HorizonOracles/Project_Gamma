// BettingFeed - Displays list of recent bets
// TODO: Implement full betting feed with real-time updates
import { useQuery } from "@tanstack/react-query";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

interface BetFeedItem {
  id: string;
  userEmail?: string;
  userName?: string; // Alternative display name
  amount: string;
  outcome: 'A' | 'B' | string;
  teamA?: string;
  teamB?: string;
  marketDescription?: string;
  oddsAtBet?: string;
  odds?: string; // Alternative field name
  createdAt?: string;
  timestamp?: Date; // Alternative field name
  status: string;
}

interface BettingFeedProps {
  bets?: BetFeedItem[]; // Optional external bets data
}

export function BettingFeed({ bets: externalBets }: BettingFeedProps = {}) {
  const { data: fetchedBets, isLoading } = useQuery<BetFeedItem[]>({
    queryKey: ['/api/bets/feed'],
    queryFn: async () => {
      const res = await fetch('/api/bets/feed?limit=20', {
        credentials: 'include',
      });
      if (!res.ok) {
        throw new Error('Failed to fetch betting feed');
      }
      return res.json();
    },
    refetchInterval: 5000, // Refresh every 5 seconds
    enabled: !externalBets, // Only fetch if external bets not provided
  });

  // Use external bets if provided, otherwise use fetched bets
  const bets = externalBets ?? fetchedBets;

  if (isLoading) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Recent Bets</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2">
          {[...Array(5)].map((_, i) => (
            <Skeleton key={i} className="h-16 w-full" />
          ))}
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Recent Bets</CardTitle>
      </CardHeader>
      <CardContent>
        {!bets || bets.length === 0 ? (
          <p className="text-sm text-muted-foreground">
            No recent betting activity.
          </p>
        ) : (
          <div className="space-y-2">
            {bets.map((bet) => (
              <div
                key={bet.id}
                className="flex items-center justify-between p-3 rounded-lg border"
              >
                <div className="flex-1">
                  <p className="text-sm font-medium">
                    {bet.userName || bet.userEmail || "Anonymous"} bet on{" "}
                    {typeof bet.outcome === 'string' && (bet.outcome === 'A' || bet.outcome === 'B')
                      ? (bet.outcome === 'A' ? bet.teamA : bet.teamB)
                      : bet.outcome}
                  </p>
                  <p className="text-xs text-muted-foreground">
                    {bet.marketDescription}
                  </p>
                </div>
                <div className="text-right">
                  <p className="text-sm font-bold">{bet.amount} BNB</p>
                  <p className="text-xs text-muted-foreground">
                    @{bet.odds || bet.oddsAtBet}
                  </p>
                </div>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
