// Horizon Prediction Market Card Component
// Reference: design_guidelines.md - Card styling with yellow/gold odds buttons
import { useState, useEffect } from "react";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Clock, TrendingUp } from "lucide-react";
import { cn } from "@/lib/utils";
import { useTheme } from "@/components/ThemeProvider";
import { NetworkBackground } from "@/components/NetworkBackground";
import { getLeagueBadge } from "@/lib/leagueUtils";
import { getTeamOrPlayerImage, extractTeamsFromMarket } from "@/lib/teamUtils";
import type { Market } from "@shared/schema";

interface MarketCardProps {
  market: Market;
  onPlaceBet: (marketId: string, outcome: 'A' | 'B', odds: number) => void;
}

export function MarketCard({ market, onPlaceBet }: MarketCardProps) {
  const [hoveredOutcome, setHoveredOutcome] = useState<'A' | 'B' | null>(null);
  const [teamABadge, setTeamABadge] = useState<string | null>(null);
  const [teamBBadge, setTeamBBadge] = useState<string | null>(null);
  const [timeRemaining, setTimeRemaining] = useState({ hours: 0, minutes: 0, seconds: 0 });
  const { theme } = useTheme();
  const leagueBadge = getLeagueBadge(market.league);

  // Fetch team badges or player headshots depending on sport
  useEffect(() => {
    const fetchTeamBadges = async () => {
      const teams = extractTeamsFromMarket(market);
      
      const [badgeA, badgeB] = await Promise.all([
        teams.teamA ? getTeamOrPlayerImage(teams.teamA, market.sport) : null,
        teams.teamB ? getTeamOrPlayerImage(teams.teamB, market.sport) : null
      ]);
      
      console.log(`[${market.sport}] ${teams.teamA} -> ${badgeA}`);
      console.log(`[${market.sport}] ${teams.teamB} -> ${badgeB}`);
      
      setTeamABadge(badgeA);
      setTeamBBadge(badgeB);
    };
    fetchTeamBadges();
  }, [market.teamA, market.teamB, market.description, market.sport]);

  // Live countdown timer that updates every second
  useEffect(() => {
    const updateCountdown = () => {
      const gameTime = new Date(market.gameTime);
      const now = new Date();
      const timeDiff = gameTime.getTime() - now.getTime();
      
      if (timeDiff > 0) {
        const hours = Math.floor(timeDiff / (1000 * 60 * 60));
        const minutes = Math.floor((timeDiff / (1000 * 60)) % 60);
        const seconds = Math.floor((timeDiff / 1000) % 60);
        setTimeRemaining({ hours, minutes, seconds });
      } else {
        setTimeRemaining({ hours: 0, minutes: 0, seconds: 0 });
      }
    };

    // Update immediately
    updateCountdown();

    // Then update every second
    const interval = setInterval(updateCountdown, 1000);

    return () => clearInterval(interval);
  }, [market.gameTime]);

  // Calculate dynamic odds
  const poolATotal = parseFloat(market.poolATotal || "0");
  const poolBTotal = parseFloat(market.poolBTotal || "0");
  const bonusPool = parseFloat(market.bonusPool || "0");
  const totalPool = poolATotal + poolBTotal + bonusPool;

  const oddsA = poolATotal > 0 ? (totalPool / poolATotal) : 2.00;
  const oddsB = poolBTotal > 0 ? (totalPool / poolBTotal) : 2.00;

  const formatOdds = (odds: number) => odds.toFixed(2);
  const formatPool = (amount: number) => amount.toFixed(4);

  const gameTime = new Date(market.gameTime);
  const now = new Date();
  const isUpcoming = gameTime > now;
  
  // Determine market state: Upcoming, LIVE, or Ended
  const isEnded = market.status === 'settled' || market.status === 'voided';
  const isLive = !isUpcoming && !isEnded;
  const isLocked = market.status === 'locked';
  
  // Format game time in user's local timezone
  const formatGameTime = (date: Date) => {
    return date.toLocaleString('en-US', {
      month: 'short',
      day: 'numeric',
      hour: 'numeric',
      minute: '2-digit',
      hour12: true
    });
  };

  // Dynamic font sizing based on team name lengths
  const totalNameLength = market.teamA.length + market.teamB.length;
  const getTeamNameSize = () => {
    if (totalNameLength > 35) return "text-sm"; // Very long names
    if (totalNameLength > 25) return "text-base"; // Long names
    return "text-lg"; // Normal/short names
  };

  return (
    <Card className="bg-card border-card-border hover-elevate transition-all duration-200">
      <CardHeader className="flex flex-row items-center justify-between gap-2 space-y-0 pb-3">
        <div className="flex items-center gap-2">
          {leagueBadge ? (
            <img 
              src={leagueBadge} 
              alt={market.league}
              className="h-12 w-12 object-contain"
              data-testid={`img-league-logo-${market.id}`}
            />
          ) : (
            <Badge variant="outline" className="border-accent text-accent text-xs font-mono">
              {market.league}
            </Badge>
          )}
        </div>
        
        {/* Team Logos in Header */}
        <div className="flex items-center gap-3 flex-1 justify-center">
          {teamABadge && (
            <img 
              src={teamABadge} 
              alt={market.teamA}
              className="h-11 w-11 object-contain"
              data-testid={`img-team-a-logo-${market.id}`}
            />
          )}
          {teamBBadge && (
            <img 
              src={teamBBadge} 
              alt={market.teamB}
              className="h-11 w-11 object-contain"
              data-testid={`img-team-b-logo-${market.id}`}
            />
          )}
        </div>
        
        <div className="flex flex-col items-end gap-1">
          {isEnded && (
            <div 
              className="inline-flex items-center rounded-md px-2.5 py-0.5 text-xs font-bold border-0"
              style={{ backgroundColor: '#666666', color: '#ffffff', boxShadow: 'none', textShadow: 'none' }}
              data-testid={`badge-ended-${market.id}`}
            >
              ENDED
            </div>
          )}
          {isLive && (
            <div className="flex items-center gap-1">
              <div 
                className="inline-flex items-center rounded-md px-2.5 py-0.5 text-xs font-bold border-0"
                style={{ backgroundColor: '#ff0000', color: '#ffffff', boxShadow: 'none', textShadow: 'none' }}
                data-testid={`badge-live-${market.id}`}
              >
                LIVE
              </div>
              {isLocked && (
                <div 
                  className="inline-flex items-center justify-center rounded-md px-1.5 py-0.5 text-xs font-bold"
                  style={{ backgroundColor: '#ffa500', color: '#ffffff' }}
                  data-testid={`badge-locked-${market.id}`}
                  title="Betting locked"
                >
                  🔒
                </div>
              )}
            </div>
          )}
          {isUpcoming && (
            <div className="flex items-center gap-1" data-testid={`time-remaining-${market.id}`}>
              <Clock className="h-3.5 w-3.5 text-muted-foreground" />
              <span className="font-mono text-sm font-semibold text-muted-foreground">
                {timeRemaining.hours}h {timeRemaining.minutes}m {timeRemaining.seconds}s
              </span>
            </div>
          )}
        </div>
      </CardHeader>
      
      <CardContent className="space-y-4">

        {/* Matchup */}
        <div className="text-center space-y-2">
          <div className="flex items-center justify-center gap-4">
            <h3 className="text-base font-bold text-white leading-none">{market.teamA}</h3>
            <span className="text-3xl font-bold goldify-text leading-none">VS</span>
            <h3 className="text-base font-bold text-white leading-none">{market.teamB}</h3>
          </div>
          <p className="text-xs text-white/80 line-clamp-1">{market.description}</p>
        </div>

        {/* Odds Buttons */}
        <div className="grid grid-cols-2 gap-3">
          <Button
            onClick={() => onPlaceBet(market.id, 'A', oddsA)}
            onMouseEnter={() => setHoveredOutcome('A')}
            onMouseLeave={() => setHoveredOutcome(null)}
            disabled={market.status !== 'active'}
            className={cn(
              "h-[100px] flex-col justify-center py-3 goldify text-primary-foreground relative overflow-hidden border-2 rounded-lg",
              market.status !== 'active' && "opacity-50 cursor-not-allowed"
            )}
            style={{ borderColor: '#d5b877' }}
            data-testid={`button-bet-${market.id}-a`}
          >
            {/* Network Background Animation */}
            <div className="absolute inset-0 opacity-60">
              <NetworkBackground color="gold" className="w-full h-full" />
            </div>
            <div className="flex flex-col items-center justify-center gap-1 relative z-10">
              <span className="text-sm font-semibold">{market.teamA}</span>
              <span className="text-3xl font-mono font-bold">{formatOdds(oddsA)}</span>
              <span className="text-xs opacity-90">
                Pool: {formatPool(poolATotal)} BNB
              </span>
            </div>
          </Button>

          <Button
            onClick={() => onPlaceBet(market.id, 'B', oddsB)}
            onMouseEnter={() => setHoveredOutcome('B')}
            onMouseLeave={() => setHoveredOutcome(null)}
            disabled={market.status !== 'active'}
            className={cn(
              "h-[100px] flex-col justify-center py-3 darkify text-accent-foreground relative overflow-hidden border-2 rounded-lg",
              market.status !== 'active' && "opacity-50 cursor-not-allowed"
            )}
            style={{ borderColor: '#424242' }}
            data-testid={`button-bet-${market.id}-b`}
          >
            {/* Network Background Animation */}
            <div className="absolute inset-0 opacity-60">
              <NetworkBackground color="gray" className="w-full h-full" />
            </div>
            <div className="flex flex-col items-center justify-center gap-1 relative z-10">
              <span className="text-sm font-semibold" style={{ textShadow: '0 0 8px rgba(0,0,0,0.9), 0 0 4px rgba(0,0,0,0.8)' }}>{market.teamB}</span>
              <span className="text-3xl font-mono font-bold" style={{ textShadow: '0 0 8px rgba(0,0,0,0.9), 0 0 4px rgba(0,0,0,0.8)' }}>{formatOdds(oddsB)}</span>
              <span className="text-xs opacity-90" style={{ textShadow: '0 0 8px rgba(0,0,0,0.9), 0 0 4px rgba(0,0,0,0.8)' }}>
                Pool: {formatPool(poolBTotal)} BNB
              </span>
            </div>
          </Button>
        </div>

        {/* Bonus Pool Info */}
        {bonusPool > 0 && (
          <div className="flex items-center justify-center gap-1.5 text-xs text-white/90">
            <TrendingUp className="h-3 w-3" />
            <span className="font-medium">
              Includes {formatPool(bonusPool)} BNB Bonus Pool!
            </span>
          </div>
        )}

        {market.status !== 'active' && (
          <Badge 
            variant="outline" 
            className="w-full justify-center border-accent text-accent"
          >
            {market.status.toUpperCase()}
          </Badge>
        )}
      </CardContent>
    </Card>
  );
}
