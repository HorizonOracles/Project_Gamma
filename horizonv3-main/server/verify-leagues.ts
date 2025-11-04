#!/usr/bin/env tsx
import { sportsData } from '../shared/sports-leagues';
import { theSportsDB } from './thesportsdb';
import { writeFileSync } from 'fs';

interface LeagueVerification {
  sport: string;
  leagueName: string;
  leagueId: string;
  eventsCount: number;
  teamsCount: number;
  teamsFromEventsCount: number;
  badgeCoverage: number;
  status: 'ok' | 'no_events' | 'no_teams' | 'error';
  error?: string;
}

async function verifyLeague(sport: string, league: any): Promise<LeagueVerification> {
  const result: LeagueVerification = {
    sport,
    leagueName: league.displayName || league.name,
    leagueId: league.id,
    eventsCount: 0,
    teamsCount: 0,
    teamsFromEventsCount: 0,
    badgeCoverage: 0,
    status: 'ok'
  };

  try {
    // Get upcoming events
    const events = await theSportsDB.getUpcomingEventsByLeague(league.id);
    result.eventsCount = events.length;

    // Get teams from league API
    const teams = await theSportsDB.getTeamsByLeague(league.name);
    result.teamsCount = teams.length;

    // Get teams from events
    const teamsFromEvents = await theSportsDB.getTeamsFromEvents(league.id);
    result.teamsFromEventsCount = teamsFromEvents.length;

    // Calculate badge coverage
    const teamsWithBadges = teamsFromEvents.filter(t => t.badge !== null).length;
    result.badgeCoverage = teamsFromEvents.length > 0 
      ? Math.round((teamsWithBadges / teamsFromEvents.length) * 100) 
      : 0;

    // Determine status
    if (result.eventsCount === 0) {
      result.status = 'no_events';
    } else if (result.teamsCount === 0 && result.teamsFromEventsCount === 0) {
      result.status = 'no_teams';
    }

  } catch (error) {
    result.status = 'error';
    result.error = error instanceof Error ? error.message : String(error);
  }

  return result;
}

async function verifyAllLeagues() {
  console.log('🔍 Starting league verification for all 497 leagues...\n');
  
  const results: LeagueVerification[] = [];
  let processedCount = 0;
  
  for (const sport of sportsData) {
    console.log(`\n📊 Verifying ${sport.name} (${sport.leagues.length} leagues)...`);
    
    for (const league of sport.leagues) {
      processedCount++;
      process.stdout.write(`\r  Progress: ${processedCount}/497 leagues... `);
      
      const result = await verifyLeague(sport.name, league);
      results.push(result);
      
      // Conservative delay to respect TheSportsDB API rate limits
      // With 497 leagues, this ensures we don't exceed limits
      await new Promise(resolve => setTimeout(resolve, 1000)); // 1 second between leagues
    }
  }

  console.log('\n\n✅ Verification complete!\n');
  
  // Generate summary
  const summary = {
    total: results.length,
    ok: results.filter(r => r.status === 'ok').length,
    noEvents: results.filter(r => r.status === 'no_events').length,
    noTeams: results.filter(r => r.status === 'no_teams').length,
    errors: results.filter(r => r.status === 'error').length,
    avgBadgeCoverage: Math.round(
      results.reduce((sum, r) => sum + r.badgeCoverage, 0) / results.length
    )
  };

  // Print summary
  console.log('📈 SUMMARY:');
  console.log(`  Total leagues: ${summary.total}`);
  console.log(`  ✅ OK: ${summary.ok}`);
  console.log(`  ⚠️  No events: ${summary.noEvents}`);
  console.log(`  ⚠️  No teams: ${summary.noTeams}`);
  console.log(`  ❌ Errors: ${summary.errors}`);
  console.log(`  🎯 Avg badge coverage: ${summary.avgBadgeCoverage}%`);

  // Find leagues with issues
  const leaguesWithIssues = results.filter(r => 
    r.status !== 'ok' || r.badgeCoverage < 50
  );

  if (leaguesWithIssues.length > 0) {
    console.log(`\n⚠️  ${leaguesWithIssues.length} leagues need attention:\n`);
    leaguesWithIssues.slice(0, 20).forEach(league => {
      console.log(`  - ${league.sport} / ${league.leagueName}:`);
      console.log(`    Events: ${league.eventsCount}, Teams: ${league.teamsFromEventsCount}, Badges: ${league.badgeCoverage}%`);
      if (league.error) console.log(`    Error: ${league.error}`);
    });
    
    if (leaguesWithIssues.length > 20) {
      console.log(`\n  ... and ${leaguesWithIssues.length - 20} more`);
    }
  }

  // Save detailed report to file
  const report = {
    timestamp: new Date().toISOString(),
    summary,
    results,
    leaguesWithIssues
  };

  const reportPath = './league-verification-report.json';
  writeFileSync(reportPath, JSON.stringify(report, null, 2));
  console.log(`\n📝 Detailed report saved to: ${reportPath}`);

  return report;
}

// Run verification
verifyAllLeagues()
  .then(() => {
    console.log('\n✨ Done!');
    process.exit(0);
  })
  .catch(error => {
    console.error('\n❌ Verification failed:', error);
    process.exit(1);
  });
