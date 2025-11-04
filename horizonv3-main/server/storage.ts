// DegenArena Database Storage Implementation
// Reference: javascript_database blueprint and javascript_log_in_with_replit blueprint
import {
  users,
  wallets,
  transactions,
  markets,
  bets,
  chatMessages,
  customTeams,
  customPlayers,
  customLeagues,
  visibilitySettings,
  adminWhitelist,
  type User,
  type UpsertUser,
  type Wallet,
  type InsertWallet,
  type Transaction,
  type InsertTransaction,
  type Market,
  type InsertMarket,
  type Bet,
  type InsertBet,
  type ChatMessage,
  type InsertChatMessage,
  type CustomTeam,
  type InsertCustomTeam,
  type CustomPlayer,
  type InsertCustomPlayer,
  type CustomLeague,
  type InsertCustomLeague,
  type VisibilitySetting,
  type InsertVisibilitySetting,
  type AdminWhitelist,
  type InsertAdminWhitelist,
} from "@shared/schema";
import { db } from "./db";
import { eq, desc, and, sql } from "drizzle-orm";

export interface IStorage {
  // User operations (Required for Replit Auth)
  getUser(id: string): Promise<User | undefined>;
  upsertUser(user: UpsertUser): Promise<User>;
  updateUserProfile(userId: string, updates: { displayName?: string; profileImageUrl?: string; lastNameChange?: Date }): Promise<User>;
  updateUserStats(userId: string, wageredAmount?: string, wonAmount?: string, tx?: any): Promise<User>;
  getLeaderboard(limit: number): Promise<User[]>;
  
  // Wallet operations
  getWalletByUserId(userId: string, tx?: any): Promise<Wallet | undefined>;
  createWallet(wallet: InsertWallet): Promise<Wallet>;
  updateWalletBalance(walletId: string, newBalance: string, tx?: any): Promise<Wallet>;
  
  // Transaction operations
  createTransaction(transaction: InsertTransaction, tx?: any): Promise<Transaction>;
  getTransactionsByWalletId(walletId: string): Promise<Transaction[]>;
  
  // Market operations
  getAllActiveMarkets(): Promise<Market[]>;
  getAllMarkets(): Promise<Market[]>;
  getMarketById(id: string, tx?: any): Promise<Market | undefined>;
  createMarket(market: InsertMarket): Promise<Market>;
  updateMarketPools(marketId: string, poolATotal: string, poolBTotal: string, tx?: any): Promise<Market>;
  updateMarketStatus(marketId: string, status: string, tx?: any): Promise<Market>;
  settleMarket(marketId: string, winningOutcome: 'A' | 'B', tx?: any): Promise<Market>;
  
  // Bet operations
  createBet(bet: InsertBet, tx?: any): Promise<Bet>;
  getBetsByUserId(userId: string): Promise<any[]>;
  getBetsByMarketId(marketId: string, tx?: any): Promise<Bet[]>;
  getRecentBetsForFeed(limit: number): Promise<any[]>;
  updateBetStatus(betId: string, status: string, actualPayout?: string, tx?: any): Promise<Bet>;
  
  // Chat operations
  createChatMessage(message: InsertChatMessage): Promise<ChatMessage>;
  getRecentChatMessages(limit: number): Promise<any[]>;
  
  // Custom sports data operations
  createCustomTeam(team: InsertCustomTeam): Promise<CustomTeam>;
  getCustomTeams(sport?: string, league?: string): Promise<CustomTeam[]>;
  createCustomPlayer(player: InsertCustomPlayer): Promise<CustomPlayer>;
  getCustomPlayers(sport?: string): Promise<CustomPlayer[]>;
  createCustomLeague(league: InsertCustomLeague): Promise<CustomLeague>;
  getCustomLeagues(sport?: string): Promise<CustomLeague[]>;
  
  // Visibility settings operations
  getVisibilitySettings(): Promise<VisibilitySetting[]>;
  toggleSportVisibility(sportId: string, isVisible: boolean, updatedBy: string): Promise<VisibilitySetting>;
  toggleLeagueVisibility(leagueId: string, sportId: string, isVisible: boolean, updatedBy: string): Promise<VisibilitySetting>;
  
  // Admin whitelist operations
  isAdminWhitelisted(walletAddress: string): Promise<boolean>;
  getAdminWhitelist(): Promise<AdminWhitelist[]>;
  getAdminWhitelistEntry(walletAddress: string): Promise<AdminWhitelist | undefined>;
  addAdminToWhitelist(entry: InsertAdminWhitelist): Promise<AdminWhitelist>;
  removeAdminFromWhitelist(walletAddress: string): Promise<void>;
  
  // Market blockchain operations
  updateMarketBlockchainData(marketId: string, data: {
    marketAddress?: string;
    chainId?: number;
    onChainMarketId?: string;
    yesTokenId?: string;
    noTokenId?: string;
    resolutionSource?: string;
    oracleRequestId?: string;
  }): Promise<Market>;
}

export class DatabaseStorage implements IStorage {
  // ============================================================================
  // USER OPERATIONS (Required for Replit Auth)
  // ============================================================================
  
  async getUser(id: string): Promise<User | undefined> {
    const [user] = await db.select().from(users).where(eq(users.id, id));
    return user;
  }

  async upsertUser(userData: UpsertUser): Promise<User> {
    const [user] = await db
      .insert(users)
      .values(userData)
      .onConflictDoUpdate({
        target: users.id,
        set: {
          ...userData,
          updatedAt: new Date(),
        },
      })
      .returning();
    return user;
  }

  async updateUserProfile(userId: string, updates: { displayName?: string; profileImageUrl?: string; lastNameChange?: Date }): Promise<User> {
    const [user] = await db
      .update(users)
      .set({
        ...updates,
        updatedAt: new Date(),
      })
      .where(eq(users.id, userId))
      .returning();
    return user;
  }

  async updateUserStats(userId: string, wageredAmount?: string, wonAmount?: string, tx?: any): Promise<User> {
    const dbClient = tx || db;
    const currentUser = await this.getUser(userId);
    if (!currentUser) throw new Error("User not found");

    const currentWagered = parseFloat(currentUser.totalWagered);
    const currentWon = parseFloat(currentUser.totalWon);
    const currentPoints = currentUser.rankPoints;

    const newWagered = wageredAmount ? currentWagered + parseFloat(wageredAmount) : currentWagered;
    const newWon = wonAmount ? currentWon + parseFloat(wonAmount) : currentWon;
    
    // Calculate rank points: 1 point per BNB wagered, 5 points per BNB won
    const wageredPoints = wageredAmount ? Math.floor(parseFloat(wageredAmount)) : 0;
    const wonPoints = wonAmount ? Math.floor(parseFloat(wonAmount) * 5) : 0;
    const newPoints = currentPoints + wageredPoints + wonPoints;

    // Determine rank based on points - 12 tier system with progressive difficulty
    // Each tier has 100 levels, with XP per level increasing each tier
    let rank = "Bronze";
    if (newPoints >= 660000) rank = "Supernova";      // Levels 1101-1200: 100 levels x 1200 XP/level
    else if (newPoints >= 550000) rank = "Nebula";    // Levels 1001-1100: 100 levels x 1100 XP/level
    else if (newPoints >= 450000) rank = "Stardust";  // Levels 901-1000: 100 levels x 1000 XP/level
    else if (newPoints >= 360000) rank = "Opal";      // Levels 801-900: 100 levels x 900 XP/level
    else if (newPoints >= 280000) rank = "Pearl";     // Levels 701-800: 100 levels x 800 XP/level
    else if (newPoints >= 210000) rank = "Diamond";   // Levels 601-700: 100 levels x 700 XP/level
    else if (newPoints >= 150000) rank = "Ruby";      // Levels 501-600: 100 levels x 600 XP/level
    else if (newPoints >= 100000) rank = "Emerald";   // Levels 401-500: 100 levels x 500 XP/level
    else if (newPoints >= 60000) rank = "Sapphire";   // Levels 301-400: 100 levels x 400 XP/level
    else if (newPoints >= 30000) rank = "Gold";       // Levels 201-300: 100 levels x 300 XP/level
    else if (newPoints >= 10000) rank = "Silver";     // Levels 101-200: 100 levels x 200 XP/level
    // Bronze: Levels 1-100: 100 levels x 100 XP/level = 0-10,000 XP

    const [user] = await dbClient
      .update(users)
      .set({
        totalWagered: newWagered.toString(),
        totalWon: newWon.toString(),
        rankPoints: newPoints,
        rank,
        updatedAt: new Date(),
      })
      .where(eq(users.id, userId))
      .returning();
    return user;
  }

  async getLeaderboard(limit: number): Promise<User[]> {
    const leaders = await db
      .select()
      .from(users)
      .orderBy(sql`(CAST(total_won AS DECIMAL) - CAST(total_wagered AS DECIMAL)) DESC`)
      .limit(limit);
    return leaders;
  }

  // ============================================================================
  // WALLET OPERATIONS
  // ============================================================================

  async getWalletByUserId(userId: string, tx?: any): Promise<Wallet | undefined> {
    const dbClient = tx || db;
    const [wallet] = await dbClient
      .select()
      .from(wallets)
      .where(eq(wallets.userId, userId));
    return wallet;
  }

  async createWallet(walletData: InsertWallet): Promise<Wallet> {
    const [wallet] = await db
      .insert(wallets)
      .values(walletData)
      .returning();
    return wallet;
  }

  async updateWalletBalance(walletId: string, newBalance: string, tx?: any): Promise<Wallet> {
    const dbClient = tx || db;
    const [wallet] = await dbClient
      .update(wallets)
      .set({ 
        balance: newBalance,
        updatedAt: new Date(),
      })
      .where(eq(wallets.id, walletId))
      .returning();
    return wallet;
  }

  // ============================================================================
  // TRANSACTION OPERATIONS
  // ============================================================================

  async createTransaction(txData: InsertTransaction, tx?: any): Promise<Transaction> {
    const dbClient = tx || db;
    const [transaction] = await dbClient
      .insert(transactions)
      .values(txData)
      .returning();
    return transaction;
  }

  async getTransactionsByWalletId(walletId: string): Promise<Transaction[]> {
    return await db
      .select()
      .from(transactions)
      .where(eq(transactions.walletId, walletId))
      .orderBy(desc(transactions.createdAt))
      .limit(50);
  }

  // ============================================================================
  // MARKET OPERATIONS
  // ============================================================================

  async getAllActiveMarkets(): Promise<Market[]> {
    return await db
      .select()
      .from(markets)
      .where(eq(markets.status, 'active'))
      .orderBy(desc(markets.isLive), markets.gameTime);
  }

  async getAllMarkets(): Promise<Market[]> {
    return await db
      .select()
      .from(markets)
      .orderBy(desc(markets.createdAt));
  }

  async getMarketById(id: string, tx?: any): Promise<Market | undefined> {
    const dbClient = tx || db;
    const [market] = await dbClient
      .select()
      .from(markets)
      .where(eq(markets.id, id));
    return market;
  }

  async createMarket(marketData: InsertMarket): Promise<Market> {
    const [market] = await db
      .insert(markets)
      .values(marketData)
      .returning();
    return market;
  }

  async updateMarketPools(marketId: string, poolATotal: string, poolBTotal: string, tx?: any): Promise<Market> {
    const dbClient = tx || db;
    const [market] = await dbClient
      .update(markets)
      .set({ 
        poolATotal,
        poolBTotal,
        updatedAt: new Date(),
      })
      .where(eq(markets.id, marketId))
      .returning();
    return market;
  }

  async updateMarketStatus(marketId: string, status: string, tx?: any): Promise<Market> {
    const dbClient = tx || db;
    const [market] = await dbClient
      .update(markets)
      .set({ 
        status,
        updatedAt: new Date(),
      })
      .where(eq(markets.id, marketId))
      .returning();
    return market;
  }

  async settleMarket(marketId: string, winningOutcome: 'A' | 'B', tx?: any): Promise<Market> {
    const dbClient = tx || db;
    const [market] = await dbClient
      .update(markets)
      .set({ 
        status: 'settled',
        winningOutcome,
        settledAt: new Date(),
        updatedAt: new Date(),
      })
      .where(eq(markets.id, marketId))
      .returning();
    return market;
  }

  // ============================================================================
  // BET OPERATIONS
  // ============================================================================

  async createBet(betData: InsertBet, tx?: any): Promise<Bet> {
    const dbClient = tx || db;
    const [bet] = await dbClient
      .insert(bets)
      .values(betData)
      .returning();
    return bet;
  }

  async getBetsByUserId(userId: string): Promise<any[]> {
    // Join with markets to get team names and market info
    const result = await db
      .select({
        id: bets.id,
        userId: bets.userId,
        marketId: bets.marketId,
        outcome: bets.outcome,
        amount: bets.amount,
        oddsAtBet: bets.oddsAtBet,
        status: bets.status,
        actualPayout: bets.actualPayout,
        createdAt: bets.createdAt,
        settledAt: bets.settledAt,
        teamA: markets.teamA,
        teamB: markets.teamB,
        teamALogo: markets.teamALogo,
        teamBLogo: markets.teamBLogo,
        marketDescription: markets.description,
        sport: markets.sport,
        league: markets.league,
      })
      .from(bets)
      .innerJoin(markets, eq(bets.marketId, markets.id))
      .where(eq(bets.userId, userId))
      .orderBy(desc(bets.createdAt));
    
    return result;
  }

  async getBetsByMarketId(marketId: string, tx?: any): Promise<Bet[]> {
    const dbClient = tx || db;
    return await dbClient
      .select()
      .from(bets)
      .where(eq(bets.marketId, marketId))
      .orderBy(desc(bets.createdAt));
  }

  async getRecentBetsForFeed(limit: number = 20): Promise<any[]> {
    // Join with users and markets to get display info - ONLY WINNING BETS
    const result = await db
      .select({
        id: bets.id,
        userId: bets.userId,
        userEmail: users.email,
        userDisplayName: users.displayName,
        profileImageUrl: users.profileImageUrl,
        userRank: users.rank,
        userRankPoints: users.rankPoints,
        outcome: bets.outcome,
        amount: bets.amount,
        oddsAtBet: bets.oddsAtBet,
        status: bets.status,
        actualPayout: bets.actualPayout,
        createdAt: bets.createdAt,
        teamA: markets.teamA,
        teamB: markets.teamB,
        league: markets.league,
        marketDescription: markets.description,
      })
      .from(bets)
      .innerJoin(users, eq(bets.userId, users.id))
      .innerJoin(markets, eq(bets.marketId, markets.id))
      .where(eq(bets.status, 'won'))
      .orderBy(desc(bets.settledAt))
      .limit(limit);
    
    return result;
  }

  async updateBetStatus(betId: string, status: string, actualPayout?: string, tx?: any): Promise<Bet> {
    const dbClient = tx || db;
    const updateData: any = { 
      status,
      settledAt: new Date(),
    };
    
    if (actualPayout) {
      updateData.actualPayout = actualPayout;
    }
    
    const [bet] = await dbClient
      .update(bets)
      .set(updateData)
      .where(eq(bets.id, betId))
      .returning();
    
    return bet;
  }

  // ============================================================================
  // CHAT OPERATIONS
  // ============================================================================

  async createChatMessage(messageData: InsertChatMessage): Promise<ChatMessage> {
    const [message] = await db
      .insert(chatMessages)
      .values(messageData)
      .returning();
    return message;
  }

  async getRecentChatMessages(limit: number = 50): Promise<any[]> {
    const result = await db
      .select({
        id: chatMessages.id,
        userId: chatMessages.userId,
        message: chatMessages.message,
        createdAt: chatMessages.createdAt,
        userEmail: users.email,
        userDisplayName: users.displayName,
        profileImageUrl: users.profileImageUrl,
        rank: users.rank,
        rankPoints: users.rankPoints,
      })
      .from(chatMessages)
      .innerJoin(users, eq(chatMessages.userId, users.id))
      .orderBy(desc(chatMessages.createdAt))
      .limit(limit);
    
    return result.reverse(); // Return in chronological order (oldest first)
  }

  // ============================================================================
  // CUSTOM SPORTS DATA OPERATIONS
  // ============================================================================

  async createCustomTeam(teamData: InsertCustomTeam): Promise<CustomTeam> {
    const [team] = await db
      .insert(customTeams)
      .values(teamData)
      .returning();
    return team;
  }

  async getCustomTeams(sport?: string, league?: string): Promise<CustomTeam[]> {
    const conditions = [];
    if (sport) {
      conditions.push(eq(customTeams.sport, sport));
    }
    if (league) {
      conditions.push(eq(customTeams.league, league));
    }
    
    if (conditions.length > 0) {
      return await db.select().from(customTeams).where(and(...conditions));
    }
    
    return await db.select().from(customTeams);
  }

  async createCustomPlayer(playerData: InsertCustomPlayer): Promise<CustomPlayer> {
    const [player] = await db
      .insert(customPlayers)
      .values(playerData)
      .returning();
    return player;
  }

  async getCustomPlayers(sport?: string): Promise<CustomPlayer[]> {
    if (sport) {
      return await db
        .select()
        .from(customPlayers)
        .where(eq(customPlayers.sport, sport));
    }
    return await db.select().from(customPlayers);
  }

  async createCustomLeague(leagueData: InsertCustomLeague): Promise<CustomLeague> {
    const [league] = await db
      .insert(customLeagues)
      .values(leagueData)
      .returning();
    return league;
  }

  async getCustomLeagues(sport?: string): Promise<CustomLeague[]> {
    if (sport) {
      return await db
        .select()
        .from(customLeagues)
        .where(eq(customLeagues.sport, sport));
    }
    return await db.select().from(customLeagues);
  }

  // ============================================================================
  // VISIBILITY SETTINGS OPERATIONS
  // ============================================================================
  
  async getVisibilitySettings(): Promise<VisibilitySetting[]> {
    return await db.select().from(visibilitySettings);
  }

  async toggleSportVisibility(sportId: string, isVisible: boolean, updatedBy: string): Promise<VisibilitySetting> {
    // Check if setting exists
    const [existing] = await db
      .select()
      .from(visibilitySettings)
      .where(and(
        eq(visibilitySettings.type, 'sport'),
        eq(visibilitySettings.sportId, sportId)
      ));

    if (existing) {
      // Update existing setting - mark as manual override since admin changed it
      const [updated] = await db
        .update(visibilitySettings)
        .set({
          isVisible,
          manualOverride: true,
          updatedBy,
          updatedAt: new Date(),
        })
        .where(eq(visibilitySettings.id, existing.id))
        .returning();
      return updated;
    } else {
      // Create new setting - mark as manual override since admin explicitly set it
      const [created] = await db
        .insert(visibilitySettings)
        .values({
          type: 'sport',
          sportId,
          isVisible,
          manualOverride: true,
          updatedBy,
        })
        .returning();
      return created;
    }
  }

  async toggleLeagueVisibility(leagueId: string, sportId: string, isVisible: boolean, updatedBy: string): Promise<VisibilitySetting> {
    // Check if setting exists
    const [existing] = await db
      .select()
      .from(visibilitySettings)
      .where(and(
        eq(visibilitySettings.type, 'league'),
        eq(visibilitySettings.leagueId, leagueId)
      ));

    if (existing) {
      // Update existing setting - mark as manual override since admin changed it
      const [updated] = await db
        .update(visibilitySettings)
        .set({
          isVisible,
          manualOverride: true,
          updatedBy,
          updatedAt: new Date(),
        })
        .where(eq(visibilitySettings.id, existing.id))
        .returning();
      return updated;
    } else {
      // Create new setting - mark as manual override since admin explicitly set it
      const [created] = await db
        .insert(visibilitySettings)
        .values({
          type: 'league',
          sportId,
          leagueId,
          isVisible,
          manualOverride: true,
          updatedBy,
        })
        .returning();
      return created;
    }
  }

  // ============================================================================
  // ADMIN WHITELIST OPERATIONS
  // ============================================================================
  
  async isAdminWhitelisted(walletAddress: string): Promise<boolean> {
    const [entry] = await db
      .select()
      .from(adminWhitelist)
      .where(and(
        eq(adminWhitelist.walletAddress, walletAddress.toLowerCase()),
        eq(adminWhitelist.isActive, true)
      ));
    return !!entry;
  }

  async getAdminWhitelist(): Promise<AdminWhitelist[]> {
    return await db
      .select()
      .from(adminWhitelist)
      .orderBy(desc(adminWhitelist.createdAt));
  }

  async getAdminWhitelistEntry(walletAddress: string): Promise<AdminWhitelist | undefined> {
    const [entry] = await db
      .select()
      .from(adminWhitelist)
      .where(eq(adminWhitelist.walletAddress, walletAddress.toLowerCase()));
    return entry;
  }

  async addAdminToWhitelist(entryData: InsertAdminWhitelist): Promise<AdminWhitelist> {
    // Normalize wallet address to lowercase
    const [entry] = await db
      .insert(adminWhitelist)
      .values({
        ...entryData,
        walletAddress: entryData.walletAddress.toLowerCase(),
      })
      .returning();
    return entry;
  }

  async removeAdminFromWhitelist(walletAddress: string): Promise<void> {
    await db
      .update(adminWhitelist)
      .set({
        isActive: false,
        updatedAt: new Date(),
      })
      .where(eq(adminWhitelist.walletAddress, walletAddress.toLowerCase()));
  }

  // ============================================================================
  // MARKET BLOCKCHAIN OPERATIONS
  // ============================================================================
  
  async updateMarketBlockchainData(marketId: string, data: {
    marketAddress?: string;
    chainId?: number;
    onChainMarketId?: string;
    yesTokenId?: string;
    noTokenId?: string;
    resolutionSource?: string;
    oracleRequestId?: string;
  }): Promise<Market> {
    const [market] = await db
      .update(markets)
      .set({
        ...data,
        updatedAt: new Date(),
      })
      .where(eq(markets.id, marketId))
      .returning();
    return market;
  }
}

export const storage = new DatabaseStorage();
