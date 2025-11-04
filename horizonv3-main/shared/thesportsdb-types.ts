// TheSportsDB API Type Definitions
export interface SportsDBTeam {
  idTeam: string;
  strTeam: string;
  strTeamBadge?: string;
  strBadge?: string; // Alternative field name
  strTeamLogo?: string;
  strTeamShort?: string;
  strSport?: string;
  strLeague?: string;
}

export interface SportsDBEvent {
  idEvent: string;
  strEvent: string;
  strEventAlternate?: string;
  strHomeTeam: string;
  strAwayTeam: string;
  intHomeScore?: string;
  intAwayScore?: string;
  dateEvent: string;
  strTime?: string;
  idHomeTeam: string;
  idAwayTeam: string;
  strLeague: string;
  strSeason?: string;
  strStatus?: string;
}

export interface SportsDBLeague {
  idLeague: string;
  strLeague: string;
  strSport: string;
  strLeagueAlternate?: string;
}
