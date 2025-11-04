// Team and Player Utilities
// Helper functions for fetching team images, player headshots, etc.

import type { Market } from "@shared/schema";

/**
 * Extract team or player names from market
 */
export function extractTeamsFromMarket(market: Market): { teamA: string; teamB: string } {
  // For markets with explicit teamA/teamB fields
  if (market.teamA && market.teamB) {
    return { teamA: market.teamA, teamB: market.teamB };
  }

  // Try to parse from description
  // Common patterns: "Team A vs Team B", "Player A vs Player B"
  const vsMatch = market.description.match(/(.+?)\s+vs\.?\s+(.+)/i);
  if (vsMatch) {
    return {
      teamA: vsMatch[1].trim(),
      teamB: vsMatch[2].trim(),
    };
  }

  // Fallback
  return {
    teamA: "Team A",
    teamB: "Team B",
  };
}

/**
 * Get team logo or player image URL
 * Can be extended to fetch from TheSportsDB or other APIs
 */
export async function getTeamOrPlayerImage(
  teamName: string,
  sport: string,
  league?: string
): Promise<string | null> {
  try {
    // TODO: Integrate with TheSportsDB API or custom image service
    // For now, return null to use fallback placeholders
    
    // Example implementation:
    // const response = await fetch(`/api/teams/image?name=${encodeURIComponent(teamName)}&sport=${sport}`);
    // if (response.ok) {
    //   const data = await response.json();
    //   return data.imageUrl;
    // }
    
    return null;
  } catch (error) {
    console.error("Error fetching team/player image:", error);
    return null;
  }
}

/**
 * Generate placeholder image URL for team/player
 */
export function getTeamPlaceholder(teamName: string): string {
  // Generate a placeholder with team initial
  const initial = teamName.charAt(0).toUpperCase();
  return `https://ui-avatars.com/api/?name=${encodeURIComponent(teamName)}&background=random&size=128`;
}
