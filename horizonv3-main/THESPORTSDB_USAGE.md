# TheSportsDB API Integration Guide

Your DegenArena sports betting platform now has full integration with TheSportsDB API! This allows you to fetch real team logos, upcoming sporting events, and league information.

## 🔑 API Key Setup

Your API key (`164325`) has been securely stored in Replit Secrets as `THESPORTSDB_API_KEY`.

## 📡 Available Backend Endpoints

### Upcoming Events
```
GET /api/sports/events/upcoming/:leagueId
```
Fetch upcoming events for a specific league.

**Example:**
```javascript
// NBA upcoming games (League ID: 4387)
const response = await fetch('/api/sports/events/upcoming/4387');
const events = await response.json();
```

### Past Events
```
GET /api/sports/events/past/:leagueId
```
Fetch past events for a specific league.

### Events by Date
```
GET /api/sports/events/date/:date?sport=SportName&league=LeagueName
```
Fetch events for a specific date.

**Example:**
```javascript
// Get all NBA games on Nov 15, 2024
const response = await fetch('/api/sports/events/date/2024-11-15?league=NBA');
const events = await response.json();
```

### Teams by League
```
GET /api/sports/teams/league/:leagueId
```
Get all teams in a specific league with their logos.

**Example:**
```javascript
// Get all NBA teams with logos
const response = await fetch('/api/sports/teams/league/4387');
const teams = await response.json();
// Each team has: strTeamBadge (logo URL), strTeam (name), etc.
```

### Search Teams
```
GET /api/sports/teams/search?name=TeamName
```
Search for teams by name.

**Example:**
```javascript
const response = await fetch('/api/sports/teams/search?name=Lakers');
const teams = await response.json();
```

### Get Team Details
```
GET /api/sports/teams/:teamId
```
Get detailed information about a specific team.

### Search Leagues
```
GET /api/sports/leagues/search?name=LeagueName
```
Search for leagues.

### Get League Details
```
GET /api/sports/leagues/:leagueId
```
Get detailed information about a specific league.

### Get All Sports
```
GET /api/sports/all
```
Get a list of all available sports.

### Get Event Details
```
GET /api/sports/events/:eventId
```
Get detailed information about a specific event.

## 🎯 Common League IDs

- **NBA (Basketball):** 4387
- **NFL (American Football):** 4391
- **English Premier League (Soccer):** 4328
- **MLB (Baseball):** 4424
- **NHL (Ice Hockey):** 4380

## 💻 Frontend Usage with React Query

### Example: Display Team Logos

```tsx
import { useQuery } from "@tanstack/react-query";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import type { SportsDBTeam } from "@shared/thesportsdb-types";

function NBATeams() {
  const { data: teams, isLoading } = useQuery<SportsDBTeam[]>({
    queryKey: ['/api/sports/teams/league', '4387'],
  });

  if (isLoading) return <div>Loading...</div>;

  return (
    <div className="grid grid-cols-6 gap-4">
      {teams?.map((team) => (
        <div key={team.idTeam} className="text-center">
          <Avatar className="h-16 w-16 mx-auto">
            <AvatarImage src={team.strTeamBadge} alt={team.strTeam} />
            <AvatarFallback>{team.strTeamShort}</AvatarFallback>
          </Avatar>
          <p className="mt-2 font-semibold">{team.strTeam}</p>
        </div>
      ))}
    </div>
  );
}
```

### Example: Display Upcoming Events

```tsx
import { useQuery } from "@tanstack/react-query";
import type { SportsDBEvent } from "@shared/thesportsdb-types";

function UpcomingGames({ leagueId }: { leagueId: string }) {
  const { data: events } = useQuery<SportsDBEvent[]>({
    queryKey: ['/api/sports/events/upcoming', leagueId],
  });

  return (
    <div className="space-y-4">
      {events?.map((event) => (
        <div key={event.idEvent} className="flex items-center justify-between">
          <span>{event.strHomeTeam}</span>
          <span className="font-bold">VS</span>
          <span>{event.strAwayTeam}</span>
          <span className="text-sm text-muted-foreground">
            {new Date(event.dateEvent).toLocaleDateString()}
          </span>
        </div>
      ))}
    </div>
  );
}
```

## 🎨 Live Demo

Visit `/sports-data-demo` in your application to see a live demonstration of:
- Searching for teams
- Viewing all teams in a league with logos
- Displaying upcoming events
- Browsing different leagues (NBA, EPL, NFL, MLB, NHL)

## 🔧 TypeScript Types

All TypeScript types are available in `@shared/thesportsdb-types`:

```typescript
import type { 
  SportsDBTeam, 
  SportsDBEvent, 
  SportsDBLeague 
} from "@shared/thesportsdb-types";
```

**Key Properties:**
- `SportsDBTeam.strTeamBadge` - Team logo URL
- `SportsDBEvent.strHomeTeam` - Home team name
- `SportsDBEvent.strAwayTeam` - Away team name
- `SportsDBEvent.dateEvent` - Event date
- `SportsDBEvent.strTime` - Event time

## 📝 Integration into Your Sports Pages

You can now update your existing sports pages (Basketball, Baseball, Football, etc.) to:

1. **Display Real Team Logos** in your markets and game cards
2. **Fetch Upcoming Events** to create new betting markets
3. **Show Team Information** in game detail pages
4. **Auto-populate Markets** based on real sporting events

### Example: Update Basketball Page

```tsx
// In client/src/pages/sports/basketball.tsx
const { data: upcomingGames } = useQuery<SportsDBEvent[]>({
  queryKey: ['/api/sports/events/upcoming', '4387'], // NBA League ID
});

// Use this data to create markets or display upcoming games
```

## 🚀 Next Steps

1. **Integrate team logos** into your existing MarketCard and GameCard components
2. **Auto-create markets** from upcoming TheSportsDB events
3. **Display real team information** on game detail pages
4. **Add team stats and history** to enhance the betting experience

## 📚 API Documentation

For full API documentation, visit: https://www.thesportsdb.com/api.php

## ⚡ Performance Notes

- The service includes automatic caching (5 minutes for events, 1 hour for teams)
- All API calls are memoized to reduce external requests
- Consider implementing loading states and error handling in your UI
