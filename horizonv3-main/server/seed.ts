// DegenArena Database Seeder
// Populate mock sports markets for testing
import { storage } from "./storage";

const mockMarkets = [
  // Test Market 1: Starting in 30 minutes
  {
    sport: "Tennis",
    league: "ATP World Tour",
    marketType: "match_winner",
    teamA: "Carlos Alcaraz",
    teamB: "Jannik Sinner",
    description: "Carlos Alcaraz vs Jannik Sinner - Match Winner",
    isLive: false,
    gameTime: new Date(Date.now() + 30 * 60 * 1000), // 30 minutes from now
    poolATotal: "2.5",
    poolBTotal: "2.4",
    bonusPool: "0.6",
  },
  
  // Test Market 2: Starting in 2 hours
  {
    sport: "Basketball",
    league: "NBA",
    marketType: "match_winner",
    teamA: "Los Angeles Lakers",
    teamB: "Boston Celtics",
    description: "Los Angeles Lakers vs Boston Celtics - Match Winner",
    isLive: false,
    gameTime: new Date(Date.now() + 2 * 60 * 60 * 1000), // 2 hours from now
    poolATotal: "3.2",
    poolBTotal: "2.8",
    bonusPool: "0.8",
  },
  
  // Test Market 3: Starting in 6 hours
  {
    sport: "American Football",
    league: "NFL",
    marketType: "match_winner",
    teamA: "Kansas City Chiefs",
    teamB: "Buffalo Bills",
    description: "Kansas City Chiefs vs Buffalo Bills - Match Winner",
    isLive: false,
    gameTime: new Date(Date.now() + 6 * 60 * 60 * 1000), // 6 hours from now
    poolATotal: "4.5",
    poolBTotal: "3.8",
    bonusPool: "1.2",
  },
  
  // Test Market 4: Starting in 12 hours
  {
    sport: "Soccer",
    league: "Premier League",
    marketType: "match_winner",
    teamA: "Manchester City",
    teamB: "Arsenal",
    description: "Manchester City vs Arsenal - Match Winner",
    isLive: false,
    gameTime: new Date(Date.now() + 12 * 60 * 60 * 1000), // 12 hours from now
    poolATotal: "5.5",
    poolBTotal: "4.2",
    bonusPool: "1.5",
  },
  
  // Test Market 5: Starting in 24 hours (tomorrow)
  {
    sport: "Baseball",
    league: "MLB",
    marketType: "match_winner",
    teamA: "New York Yankees",
    teamB: "Boston Red Sox",
    description: "Yankees vs Red Sox - Match Winner",
    isLive: false,
    gameTime: new Date(Date.now() + 24 * 60 * 60 * 1000), // 24 hours from now (tomorrow)
    poolATotal: "3.1",
    poolBTotal: "2.9",
    bonusPool: "0.7",
  },
];

export async function seedMarkets() {
  console.log("🌱 Seeding mock sports markets...");
  
  try {
    // Check for existing markets - warn if duplicates might be created
    const existingMarkets = await storage.getAllMarkets();
    if (existingMarkets.length > 0) {
      console.log(`⚠️  Warning: ${existingMarkets.length} markets already exist. New markets will be added (clear database manually to avoid duplicates).`);
    }
    
    // Create new markets
    console.log("📝 Creating new markets...");
    for (const marketData of mockMarkets) {
      await storage.createMarket(marketData);
    }
    console.log(`✅ Successfully seeded ${mockMarkets.length} new markets`);
  } catch (error) {
    console.error("❌ Error seeding markets:", error);
    throw error;
  }
}

// Run seed if this file is executed directly
// Note: ES modules don't have require.main, so we check import.meta
if (import.meta.url === `file://${process.argv[1]}`) {
  seedMarkets()
    .then(() => {
      console.log("✅ Seeding complete");
      process.exit(0);
    })
    .catch((error) => {
      console.error("❌ Seeding failed:", error);
      process.exit(1);
    });
}
