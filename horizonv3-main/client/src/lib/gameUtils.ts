// Utility functions for grouping markets by games
import type { Market } from "@shared/schema";

export interface Game {
  teamA: string;
  teamB: string;
  league: string;
  gameTime: Date;
  marketCount: number;
  isLive: boolean;
  slug: string;
  sport: string;
}

export function buildGameSlug(teamA: string, teamB: string, gameTime: Date): string {
  const teamASlug = teamA.toLowerCase().replace(/\s+/g, '-');
  const teamBSlug = teamB.toLowerCase().replace(/\s+/g, '-');
  const timestamp = new Date(gameTime).getTime();
  return `${teamASlug}-vs-${teamBSlug}-${timestamp}`;
}

export function groupMarketsByGame(markets: Market[]): Game[] {
  const gameMap = new Map<string, Game>();

  markets.forEach(market => {
    const gameTime = new Date(market.gameTime);
    // Create unique key for each game
    const gameKey = `${market.teamA}-${market.teamB}-${gameTime.getTime()}`;
    
    if (!gameMap.has(gameKey)) {
      gameMap.set(gameKey, {
        teamA: market.teamA,
        teamB: market.teamB,
        league: market.league,
        gameTime,
        marketCount: 1,
        isLive: market.isLive || false,
        slug: buildGameSlug(market.teamA, market.teamB, gameTime),
        sport: market.sport,
      });
    } else {
      const game = gameMap.get(gameKey)!;
      game.marketCount++;
      if (market.isLive) game.isLive = true;
    }
  });

  return Array.from(gameMap.values()).sort((a, b) => a.gameTime.getTime() - b.gameTime.getTime());
}
