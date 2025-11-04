// Horizon Market Pane Component - Compact List View
import { useState, useEffect } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Clock } from "lucide-react";
import { cn } from "@/lib/utils";
import { useTheme } from "@/components/ThemeProvider";
import { NetworkBackground } from "@/components/NetworkBackground";
import { getLeagueBadge } from "@/lib/leagueUtils";
import { getTeamOrPlayerImage } from "@/lib/teamUtils";
import type { Market } from "@shared/schema";

interface MarketPaneProps {
  market: Market;
  onPlaceBet: (marketId: string, outcome: 'A' | 'B', odds: number) => void;
}

export function MarketPane({ market, onPlaceBet }: MarketPaneProps) {
  const [teamABadge, setTeamABadge] = useState<string | null>(null);
  const [teamBBadge, setTeamBBadge] = useState<string | null>(null);
  const [timeRemaining, setTimeRemaining] = useState({ hours: 0, minutes: 0, seconds: 0 });
  const { theme } = useTheme();
  const leagueBadge = getLeagueBadge(market.league);

  // Fetch team badges or player headshots depending on sport
  useEffect(() => {
    const fetchTeamBadges = async () => {
      // Extract actual team names from description
      const { extractTeamsFromMarket } = await import('@/lib/teamUtils');
      const teams = extractTeamsFromMarket(market);
      
      const [badgeA, badgeB] = await Promise.all([
        teams.teamA ? getTeamOrPlayerImage(teams.teamA, market.sport) : null,
        teams.teamB ? getTeamOrPlayerImage(teams.teamB, market.sport) : null
      ]);
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
  const formatPool = (amount: number) => amount.toFixed(3);

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

  const formatDate = (date: Date) => {
    const today = new Date();
    const tomorrow = new Date(today);
    tomorrow.setDate(tomorrow.getDate() + 1);

    if (date.toDateString() === today.toDateString()) return "Today";
    if (date.toDateString() === tomorrow.toDateString()) return "Tomorrow";
    
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
  };

  // Dynamic font sizing based on team name lengths (for pane view)
  const totalNameLength = market.teamA.length + market.teamB.length;
  const getTeamNameSize = () => {
    if (totalNameLength > 35) return "text-xs"; // Very long names
    if (totalNameLength > 25) return "text-sm"; // Long names
    return "text-base"; // Normal/short names
  };

  return (
    <div 
      className="bg-card border border-card-border hover-elevate transition-all duration-200 p-4 rounded-md"
      data-testid={`pane-market-${market.id}`}
    >
      <div className="flex items-center justify-between gap-4">
        {/* League Logo - Left Aligned and Vertically Centered */}
        <div className="flex items-center justify-center min-w-[60px]">
          {leagueBadge ? (
            <img 
              src={leagueBadge} 
              alt={market.league}
              className="h-10 w-10 object-contain"
              data-testid={`img-league-logo-pane-${market.id}`}
            />
          ) : (
            <Badge variant="outline" className="border-accent text-accent text-xs font-mono">
              {market.league}
            </Badge>
          )}
        </div>

        {/* Time/Date */}
        <div className="flex flex-col min-w-[120px] gap-1">
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

        {/* Teams and Description */}
        <div className="flex-1 min-w-0">
          <div className="relative flex items-center justify-center min-h-[2rem] gap-2">
            <div className="flex items-center justify-end gap-2 flex-1">
              {teamABadge ? (
                <img 
                  src={teamABadge} 
                  alt={market.teamA}
                  className="w-12 h-12 object-contain flex-shrink-0"
                  data-testid={`img-team-a-logo-pane-${market.id}`}
                />
              ) : (
                <div className="w-12 h-12 rounded-full bg-white/10 backdrop-blur-sm border border-white/20 flex items-center justify-center flex-shrink-0">
                  <div className="w-10 h-10 rounded-full bg-gradient-to-br from-primary/20 to-primary/5 flex items-center justify-center">
                    <span className="text-xs font-bold text-primary">A</span>
                  </div>
                </div>
              )}
              <span className={cn(getTeamNameSize(), "font-bold text-white leading-none text-right")}>{market.teamA}</span>
            </div>
            <span className="text-2xl font-bold goldify-text leading-none">VS</span>
            <div className="flex items-center justify-start gap-2 flex-1">
              <span className={cn(getTeamNameSize(), "font-bold text-white leading-none text-left")}>{market.teamB}</span>
              {teamBBadge ? (
                <img 
                  src={teamBBadge} 
                  alt={market.teamB}
                  className="w-12 h-12 object-contain flex-shrink-0"
                  data-testid={`img-team-b-logo-pane-${market.id}`}
                />
              ) : (
                <div className="w-12 h-12 rounded-full bg-white/10 backdrop-blur-sm border border-white/20 flex items-center justify-center flex-shrink-0">
                  <div className="w-10 h-10 rounded-full bg-gradient-to-br from-accent/20 to-accent/5 flex items-center justify-center">
                    <span className="text-xs font-bold text-accent">B</span>
                  </div>
                </div>
              )}
            </div>
          </div>
          <p className="text-xs text-white mt-1 line-clamp-1 text-center">{market.description}</p>
        </div>

        {/* Liquidity */}
        <div className="text-right min-w-[100px]">
          <div className="text-xs text-muted-foreground">Liquidity</div>
          <div className="font-mono text-sm text-foreground">{formatPool(totalPool)} BNB</div>
        </div>

        {/* Odds Buttons */}
        <div className="flex items-center gap-2">
          <Button
            onClick={() => onPlaceBet(market.id, 'A', oddsA)}
            disabled={market.status !== 'active'}
            className={cn(
              "h-[60px] w-[80px] flex flex-col items-center justify-center gap-0.5 goldify text-primary-foreground relative overflow-hidden border-2 p-2",
              market.status !== 'active' && "opacity-50 cursor-not-allowed"
            )}
            style={{ borderColor: '#d5b877' }}
            data-testid={`button-bet-${market.id}-a`}
          >
            {/* Network Background Animation */}
            <div className="absolute inset-0 opacity-60">
              <NetworkBackground color="gold" className="w-full h-full" />
            </div>
            <span className="text-xl font-mono font-bold relative z-10">{formatOdds(oddsA)}</span>
            <span className="text-[10px] opacity-80 relative z-10">
              {formatPool(poolATotal)} BNB
            </span>
          </Button>

          <Button
            onClick={() => onPlaceBet(market.id, 'B', oddsB)}
            disabled={market.status !== 'active'}
            className={cn(
              "h-[60px] w-[80px] flex flex-col items-center justify-center gap-0.5 darkify text-accent-foreground relative overflow-hidden border-2 p-2",
              market.status !== 'active' && "opacity-50 cursor-not-allowed"
            )}
            style={{ borderColor: '#424242' }}
            data-testid={`button-bet-${market.id}-b`}
          >
            {/* Network Background Animation */}
            <div className="absolute inset-0 opacity-60">
              <NetworkBackground color="gray" className="w-full h-full" />
            </div>
            <span className="text-xl font-mono font-bold relative z-10" style={{ textShadow: '0 0 8px rgba(0,0,0,0.9), 0 0 4px rgba(0,0,0,0.8)' }}>{formatOdds(oddsB)}</span>
            <span className="text-[10px] opacity-80 relative z-10" style={{ textShadow: '0 0 8px rgba(0,0,0,0.9), 0 0 4px rgba(0,0,0,0.8)' }}>
              {formatPool(poolBTotal)} BNB
            </span>
          </Button>
        </div>
      </div>
    </div>
  );
}
