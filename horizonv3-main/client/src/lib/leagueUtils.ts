import { sportsData } from "@shared/market-categories";

export function getLeagueBadge(leagueName: string): string | undefined {
  const normalizedSearchName = leagueName.toLowerCase().trim();
  
  for (const sport of sportsData) {
    for (const league of sport.leagues) {
      const normalizedName = league.name.toLowerCase();
      const normalizedDisplayName = league.displayName.toLowerCase();
      
      if (normalizedName === normalizedSearchName || 
          normalizedDisplayName === normalizedSearchName ||
          normalizedName.includes(normalizedSearchName) ||
          normalizedSearchName.includes(normalizedName)) {
        return league.badge;
      }
    }
  }
  return undefined;
}
