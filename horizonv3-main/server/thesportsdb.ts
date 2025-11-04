import memoize from 'memoizee';
import type { SportsDBTeam, SportsDBEvent, SportsDBLeague } from '@shared/thesportsdb-types';
import { db } from './db';
import { teamBadgeCache } from '../shared/schema';
import { eq, sql as drizzleSql } from 'drizzle-orm';

const THESPORTSDB_BASE_URL = 'https://www.thesportsdb.com/api/v1/json';
const API_KEY = process.env.THESPORTSDB_API_KEY || '164325';

if (!process.env.THESPORTSDB_API_KEY) {
  console.log('Using TheSportsDB fallback API key (164325)');
}

export type { SportsDBTeam, SportsDBEvent, SportsDBLeague };

class TheSportsDBService {
  private async fetchFromAPI<T>(endpoint: string): Promise<T> {
    const url = `${THESPORTSDB_BASE_URL}/${API_KEY}/${endpoint}`;
    
    try {
      const response = await fetch(url);
      
      if (!response.ok) {
        throw new Error(`TheSportsDB API error: ${response.status} ${response.statusText}`);
      }
      
      return await response.json();
    } catch (error) {
      console.error('TheSportsDB API request failed:', error);
      throw error;
    }
  }

  async getUpcomingEventsByLeague(leagueId: string): Promise<SportsDBEvent[]> {
    const data = await this.fetchFromAPI<{ events: SportsDBEvent[] | null }>(
      `eventsnextleague.php?id=${leagueId}`
    );
    return data.events || [];
  }

  async getPastEventsByLeague(leagueId: string): Promise<SportsDBEvent[]> {
    const data = await this.fetchFromAPI<{ events: SportsDBEvent[] | null }>(
      `eventspastleague.php?id=${leagueId}`
    );
    return data.events || [];
  }

  async getEventsByDate(date: string, sport?: string, league?: string): Promise<SportsDBEvent[]> {
    let endpoint = `eventsday.php?d=${date}`;
    if (sport) endpoint += `&s=${encodeURIComponent(sport)}`;
    if (league) endpoint += `&l=${encodeURIComponent(league)}`;
    
    const data = await this.fetchFromAPI<{ events: SportsDBEvent[] | null }>(endpoint);
    return data.events || [];
  }

  async getTeamsByLeague(leagueName: string): Promise<SportsDBTeam[]> {
    const data = await this.fetchFromAPI<{ teams: SportsDBTeam[] | null }>(
      `search_all_teams.php?l=${encodeURIComponent(leagueName)}`
    );
    
    if (!data.teams || data.teams.length === 0) {
      return [];
    }

    const teamsWithLogos = await Promise.all(
      data.teams.map(async (team) => {
        const fullTeam = await this.getTeamDetails(team.idTeam);
        return fullTeam || team;
      })
    );

    const teamsWithBadges = teamsWithLogos.filter(team => team.strBadge);
    console.log(`✅ Fetched ${teamsWithLogos.length} teams for ${leagueName}, ${teamsWithBadges.length} have badges`);
    
    if (teamsWithBadges.length > 0) {
      console.log(`Sample team with badge:`, {
        name: teamsWithBadges[0].strTeam,
        badge: teamsWithBadges[0].strBadge?.substring(0, 60) + '...',
      });
    } else {
      console.log(`⚠️ No teams in ${leagueName} have logo badges available in TheSportsDB`);
    }

    return teamsWithLogos;
  }

  async searchTeams(teamName: string): Promise<SportsDBTeam[]> {
    const data = await this.fetchFromAPI<{ teams: SportsDBTeam[] | null }>(
      `searchteams.php?t=${encodeURIComponent(teamName)}`
    );
    
    if (!data.teams || data.teams.length === 0) {
      return [];
    }

    // Fetch full details for each team to get logos (same as getTeamsByLeague)
    const teamsWithLogos = await Promise.all(
      data.teams.map(async (team) => {
        const fullTeam = await this.getTeamDetails(team.idTeam);
        return fullTeam || team;
      })
    );

    if (teamsWithLogos.length > 0 && teamsWithLogos[0].strBadge) {
      console.log(`🔍 Search "${teamName}": found team with badge ✅`);
    }
    
    return teamsWithLogos;
  }

  async getTeamDetails(teamId: string): Promise<SportsDBTeam | null> {
    const data = await this.fetchFromAPI<{ teams: SportsDBTeam[] | null }>(
      `lookupteam.php?id=${teamId}`
    );
    return data.teams?.[0] || null;
  }

  async searchLeagues(leagueName: string): Promise<SportsDBLeague[]> {
    const data = await this.fetchFromAPI<{ countries: SportsDBLeague[] | null }>(
      `search_all_leagues.php?l=${encodeURIComponent(leagueName)}`
    );
    return data.countries || [];
  }

  async getLeagueDetails(leagueId: string): Promise<SportsDBLeague | null> {
    const data = await this.fetchFromAPI<{ leagues: SportsDBLeague[] | null }>(
      `lookupleague.php?id=${leagueId}`
    );
    return data.leagues?.[0] || null;
  }

  async getAllSports(): Promise<Array<{ idSport: string; strSport: string; strFormat: string; strSportThumb: string }>> {
    const data = await this.fetchFromAPI<{ sports: Array<{ idSport: string; strSport: string; strFormat: string; strSportThumb: string }> | null }>(
      `all_sports.php`
    );
    return data.sports || [];
  }

  async getEventDetails(eventId: string): Promise<SportsDBEvent | null> {
    const data = await this.fetchFromAPI<{ events: SportsDBEvent[] | null }>(
      `lookupevent.php?id=${eventId}`
    );
    return data.events?.[0] || null;
  }

  async getTeamBadgeByName(teamName: string): Promise<string | null> {
    try {
      // Check database cache first
      const cached = await db.select()
        .from(teamBadgeCache)
        .where(eq(teamBadgeCache.teamName, teamName))
        .limit(1);

      // If cached and not expired, return it
      if (cached.length > 0 && cached[0].expiresAt > new Date()) {
        console.log(`💾 Cache HIT for team: ${teamName}`);
        return cached[0].badgeUrl;
      }

      // Cache miss or expired - fetch from API
      console.log(`🔍 Search "${teamName}": checking API...`);
      const teams = await this.searchTeams(teamName);
      
      const badgeUrl = teams.length > 0 && teams[0].strBadge ? teams[0].strBadge : null;
      const teamId = teams.length > 0 ? teams[0].idTeam : null;

      // Only cache successful results (don't cache NULL from rate limits)
      if (badgeUrl !== null) {
        // Store/update in database cache (30 day expiry)
        const expiresAt = new Date();
        expiresAt.setDate(expiresAt.getDate() + 30);

        if (cached.length > 0) {
          // Update existing cache entry
          await db.update(teamBadgeCache)
            .set({ 
              badgeUrl, 
              teamId,
              fetchedAt: new Date(),
              expiresAt 
            })
            .where(eq(teamBadgeCache.teamName, teamName));
        } else {
          // Insert new cache entry
          await db.insert(teamBadgeCache)
            .values({
              teamName,
              badgeUrl,
              teamId,
              expiresAt
            })
            .onConflictDoUpdate({
              target: teamBadgeCache.teamName,
              set: { badgeUrl, teamId, fetchedAt: new Date(), expiresAt }
            });
        }
      }

      if (badgeUrl) {
        console.log(`✅ Found badge for team: ${teamName}`);
      } else {
        console.log(`⚠️ No badge found for team: ${teamName}`);
      }
      
      return badgeUrl;
    } catch (error) {
      console.error(`Error fetching badge for ${teamName}:`, error);
      
      // On error, try to return any cached value (even if expired)
      try {
        const cached = await db.select()
          .from(teamBadgeCache)
          .where(eq(teamBadgeCache.teamName, teamName))
          .limit(1);
        
        if (cached.length > 0 && cached[0].badgeUrl) {
          console.log(`⚡ Using expired cache for ${teamName} due to API error`);
          return cached[0].badgeUrl;
        }
      } catch (dbError) {
        console.error(`Database cache lookup failed:`, dbError);
      }
      
      return null;
    }
  }

  async getTeamsFromEvents(leagueId: string): Promise<Array<{ name: string; badge: string | null }>> {
    try {
      const events = await this.getUpcomingEventsByLeague(leagueId);
      const uniqueTeamNames = new Set<string>();
      
      events.forEach(event => {
        if (event.strHomeTeam) uniqueTeamNames.add(event.strHomeTeam);
        if (event.strAwayTeam) uniqueTeamNames.add(event.strAwayTeam);
      });

      const teamsArray = Array.from(uniqueTeamNames);
      console.log(`Found ${teamsArray.length} unique teams from events for league ${leagueId}`);

      const teamsWithBadges = await Promise.all(
        teamsArray.map(async (teamName) => {
          const badge = await this.getTeamBadgeByName(teamName);
          return { name: teamName, badge };
        })
      );

      return teamsWithBadges.sort((a, b) => a.name.localeCompare(b.name));
    } catch (error) {
      console.error(`Error extracting teams from events for league ${leagueId}:`, error);
      return [];
    }
  }
}

const memoizedGetUpcomingEventsByLeague = memoize(
  (leagueId: string) => sportsDBService.getUpcomingEventsByLeague(leagueId),
  { maxAge: 5 * 60 * 1000, promise: true }
);

const memoizedGetTeamsByLeague = memoize(
  (leagueId: string) => sportsDBService.getTeamsByLeague(leagueId),
  { maxAge: 60 * 60 * 1000, promise: true }
);

const memoizedSearchTeams = memoize(
  (teamName: string) => sportsDBService.searchTeams(teamName),
  { maxAge: 60 * 60 * 1000, promise: true }
);

const memoizedGetEventsByDate = memoize(
  (date: string, sport?: string, league?: string) => sportsDBService.getEventsByDate(date, sport, league),
  { maxAge: 5 * 60 * 1000, promise: true }
);

const memoizedGetTeamBadgeByName = memoize(
  (teamName: string) => sportsDBService.getTeamBadgeByName(teamName),
  { maxAge: 24 * 60 * 60 * 1000, promise: true } // Cache for 24 hours
);

const memoizedGetTeamsFromEvents = memoize(
  (leagueId: string) => sportsDBService.getTeamsFromEvents(leagueId),
  { maxAge: 60 * 60 * 1000, promise: true } // Cache for 1 hour
);

export const sportsDBService = new TheSportsDBService();

export const theSportsDB = {
  getUpcomingEventsByLeague: memoizedGetUpcomingEventsByLeague,
  getPastEventsByLeague: (leagueId: string) => sportsDBService.getPastEventsByLeague(leagueId),
  getEventsByDate: memoizedGetEventsByDate,
  getTeamsByLeague: memoizedGetTeamsByLeague,
  searchTeams: memoizedSearchTeams,
  getTeamDetails: (teamId: string) => sportsDBService.getTeamDetails(teamId),
  searchLeagues: (leagueName: string) => sportsDBService.searchLeagues(leagueName),
  getLeagueDetails: (leagueId: string) => sportsDBService.getLeagueDetails(leagueId),
  getAllSports: () => sportsDBService.getAllSports(),
  getEventDetails: (eventId: string) => sportsDBService.getEventDetails(eventId),
  getTeamBadgeByName: memoizedGetTeamBadgeByName,
  getTeamsFromEvents: memoizedGetTeamsFromEvents,
};
