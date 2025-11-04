// DegenArena Database Schema - SQLite
import { sql } from 'drizzle-orm';
import { relations } from 'drizzle-orm';
import {
  index,
  sqliteTable,
  text,
  integer,
  real,
} from "drizzle-orm/sqlite-core";
import { createInsertSchema } from "drizzle-zod";
import { z } from "zod";

// ============================================================================
// AUTH TABLES (Required for Replit Auth)
// ============================================================================

// Session storage table - Required for Replit Auth
export const sessions = sqliteTable(
  "sessions",
  {
    sid: text("sid").primaryKey(),
    sess: text("sess", { mode: 'json' }).notNull(), // Store JSON as text
    expire: integer("expire", { mode: 'timestamp' }).notNull(),
  },
  (table) => ({
    expireIdx: index("IDX_session_expire").on(table.expire),
  }),
);

// User storage table - Required for Replit Auth
export const users = sqliteTable("users", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  email: text("email").unique(),
  walletAddress: text("wallet_address").unique(), // Ethereum wallet address for Web3 login
  firstName: text("first_name"),
  lastName: text("last_name"),
  displayName: text("display_name"), // User's custom display name
  lastNameChange: integer("last_name_change", { mode: 'timestamp' }), // Last time display name was changed (7 day cooldown)
  profileImageUrl: text("profile_image_url"), // User's profile picture URL
  role: text("role").notNull().default("user"), // 'user' or 'admin' - for access control
  totalWagered: text("total_wagered").notNull().default("0"), // Store as text to maintain precision
  totalWon: text("total_won").notNull().default("0"),
  rankPoints: integer("rank_points").notNull().default(0), // Points for ranking
  rank: text("rank").notNull().default("Bronze"), // Bronze, Silver, Gold, Platinum, Diamond
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  updatedAt: integer("updated_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export type UpsertUser = typeof users.$inferInsert;
export type User = typeof users.$inferSelect;

// ============================================================================
// DEGENARENA CORE TABLES
// ============================================================================

// Wallets - Managed BNB wallets for each user
export const wallets = sqliteTable("wallets", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  userId: text("user_id").notNull().references(() => users.id),
  bnbAddress: text("bnb_address").notNull().unique(), // Mock BNB address for MVP
  balance: text("balance").notNull().default("0"), // Store as text to maintain precision
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  updatedAt: integer("updated_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertWalletSchema = createInsertSchema(wallets).omit({
  id: true,
  createdAt: true,
  updatedAt: true,
});

export type InsertWallet = z.infer<typeof insertWalletSchema>;
export type Wallet = typeof wallets.$inferSelect;

// Transactions - Deposits, Withdrawals, Bets, Settlements
export const transactions = sqliteTable("transactions", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  walletId: text("wallet_id").notNull().references(() => wallets.id),
  userId: text("user_id").notNull().references(() => users.id),
  type: text("type").notNull(), // 'deposit', 'withdrawal', 'bet_placed', 'bet_settled'
  amount: text("amount").notNull(), // Store as text to maintain precision
  status: text("status").notNull().default("pending"), // 'pending', 'completed', 'failed'
  txHash: text("tx_hash"), // Mock blockchain transaction hash
  metadata: text("metadata", { mode: 'json' }), // Additional data (bet details, market info, etc)
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertTransactionSchema = createInsertSchema(transactions).omit({
  id: true,
  createdAt: true,
});

export type InsertTransaction = z.infer<typeof insertTransactionSchema>;
export type Transaction = typeof transactions.$inferSelect;

// Admin Whitelist - Addresses allowed to perform admin actions
export const adminWhitelist = sqliteTable("admin_whitelist", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  walletAddress: text("wallet_address").notNull().unique(), // Ethereum address
  addedBy: text("added_by").references(() => users.id), // Admin who added this address
  notes: text("notes"), // Optional notes about this admin
  isActive: integer("is_active", { mode: 'boolean' }).notNull().default(true), // Can be disabled without deletion
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  updatedAt: integer("updated_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertAdminWhitelistSchema = createInsertSchema(adminWhitelist).omit({
  id: true,
  createdAt: true,
  updatedAt: true,
});

export type InsertAdminWhitelist = z.infer<typeof insertAdminWhitelistSchema>;
export type AdminWhitelist = typeof adminWhitelist.$inferSelect;

// Markets - Sports betting markets (with blockchain integration)
export const markets = sqliteTable("markets", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  sport: text("sport").notNull(), // 'NBA', 'NFL', 'Soccer', 'MLB'
  league: text("league").notNull(), // 'NBA', 'NFL', 'Premier League', etc
  marketType: text("market_type").notNull(), // 'match_winner', 'first_to_score', 'halftime_leader', etc
  teamA: text("team_a").notNull(),
  teamB: text("team_b").notNull(),
  teamALogo: text("team_a_logo"), // URL to team A logo
  teamBLogo: text("team_b_logo"), // URL to team B logo
  description: text("description").notNull(),
  status: text("status").notNull().default("active"), // 'active', 'locked', 'settled', 'voided'
  isLive: integer("is_live", { mode: 'boolean' }).default(false),
  gameTime: integer("game_time", { mode: 'timestamp' }).notNull(),
  poolATotal: text("pool_a_total").notNull().default("0"), // Total BNB in outcome A pool
  poolBTotal: text("pool_b_total").notNull().default("0"), // Total BNB in outcome B pool
  bonusPool: text("bonus_pool").notNull().default("0"), // Platform bonus injection
  winningOutcome: text("winning_outcome"), // 'A', 'B', or null
  platformFee: text("platform_fee").notNull().default("2.00"), // 2% fee
  // Blockchain fields
  marketAddress: text("market_address"), // On-chain market contract address
  chainId: integer("chain_id"), // Chain ID (e.g., 97 for BNB testnet, 56 for BNB mainnet)
  onChainMarketId: text("on_chain_market_id"), // Market ID from MarketFactory
  yesTokenId: text("yes_token_id"), // Token ID for outcome A/Yes
  noTokenId: text("no_token_id"), // Token ID for outcome B/No
  resolutionSource: text("resolution_source"), // 'oracle', 'manual', 'thesportsdb'
  oracleRequestId: text("oracle_request_id"), // AI Oracle request ID if using oracle
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  updatedAt: integer("updated_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  settledAt: integer("settled_at", { mode: 'timestamp' }),
});

export const insertMarketSchema = createInsertSchema(markets).omit({
  id: true,
  createdAt: true,
  updatedAt: true,
  settledAt: true,
});

export type InsertMarket = z.infer<typeof insertMarketSchema>;
export type Market = typeof markets.$inferSelect;

// Bets - User bets on markets
export const bets = sqliteTable("bets", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  userId: text("user_id").notNull().references(() => users.id),
  marketId: text("market_id").notNull().references(() => markets.id),
  walletId: text("wallet_id").notNull().references(() => wallets.id),
  outcome: text("outcome").notNull(), // 'A' or 'B'
  amount: text("amount").notNull(),
  oddsAtBet: text("odds_at_bet").notNull(), // Odds when bet was placed (for display only)
  sharePercentage: text("share_percentage"), // User's % share of winning pool (calculated at settlement)
  potentialPayout: text("potential_payout"), // Calculated at settlement
  actualPayout: text("actual_payout"), // Final payout after settlement
  status: text("status").notNull().default("active"), // 'active', 'won', 'lost', 'voided'
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  settledAt: integer("settled_at", { mode: 'timestamp' }),
});

export const insertBetSchema = createInsertSchema(bets).omit({
  id: true,
  createdAt: true,
  settledAt: true,
  sharePercentage: true,
  potentialPayout: true,
  actualPayout: true,
  status: true,
});

export type InsertBet = z.infer<typeof insertBetSchema>;
export type Bet = typeof bets.$inferSelect;

// TheSportsDB team badge cache - persistent storage to avoid rate limits
export const teamBadgeCache = sqliteTable("team_badge_cache", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  teamName: text("team_name").notNull().unique(),
  badgeUrl: text("badge_url"),
  teamId: text("team_id"),
  fetchedAt: integer("fetched_at", { mode: 'timestamp' }).$defaultFn(() => new Date()).notNull(),
  expiresAt: integer("expires_at", { mode: 'timestamp' }).notNull(),
});

export type TeamBadgeCache = typeof teamBadgeCache.$inferSelect;
export type InsertTeamBadgeCache = typeof teamBadgeCache.$inferInsert;

// ============================================================================
// RELATIONS
// ============================================================================

export const usersRelations = relations(users, ({ one, many }) => ({
  wallet: one(wallets, {
    fields: [users.id],
    references: [wallets.userId],
  }),
  bets: many(bets),
  transactions: many(transactions),
}));

export const walletsRelations = relations(wallets, ({ one, many }) => ({
  user: one(users, {
    fields: [wallets.userId],
    references: [users.id],
  }),
  bets: many(bets),
  transactions: many(transactions),
}));

export const marketsRelations = relations(markets, ({ many }) => ({
  bets: many(bets),
}));

export const betsRelations = relations(bets, ({ one }) => ({
  user: one(users, {
    fields: [bets.userId],
    references: [users.id],
  }),
  market: one(markets, {
    fields: [bets.marketId],
    references: [markets.id],
  }),
  wallet: one(wallets, {
    fields: [bets.walletId],
    references: [wallets.id],
  }),
}));

export const transactionsRelations = relations(transactions, ({ one }) => ({
  user: one(users, {
    fields: [transactions.userId],
    references: [users.id],
  }),
  wallet: one(wallets, {
    fields: [transactions.walletId],
    references: [wallets.id],
  }),
}));

// Chat Messages - Global chatroom for the platform
export const chatMessages = sqliteTable("chat_messages", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  userId: text("user_id").notNull().references(() => users.id),
  message: text("message").notNull(),
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertChatMessageSchema = createInsertSchema(chatMessages).omit({
  id: true,
  createdAt: true,
});

export type InsertChatMessage = z.infer<typeof insertChatMessageSchema>;
export type ChatMessage = typeof chatMessages.$inferSelect;

export const chatMessagesRelations = relations(chatMessages, ({ one }) => ({
  user: one(users, {
    fields: [chatMessages.userId],
    references: [users.id],
  }),
}));

// ============================================================================
// CUSTOM SPORTS DATA TABLES
// ============================================================================

// Custom Teams - User-uploaded teams with logos
export const customTeams = sqliteTable("custom_teams", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  name: text("name").notNull(),
  sport: text("sport").notNull(), // 'Basketball', 'Football', 'Baseball', 'Soccer', etc
  league: text("league").notNull(), // League name this team belongs to
  logoFilename: text("logo_filename").notNull(), // Filename in server/public/ (e.g., 'custom-team-logo-123.png')
  uploadedBy: text("uploaded_by").references(() => users.id),
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertCustomTeamSchema = createInsertSchema(customTeams).omit({
  id: true,
  createdAt: true,
});

export type InsertCustomTeam = z.infer<typeof insertCustomTeamSchema>;
export type CustomTeam = typeof customTeams.$inferSelect;

// Custom Players - User-uploaded players with photos (for individual sports)
export const customPlayers = sqliteTable("custom_players", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  name: text("name").notNull(),
  sport: text("sport").notNull(), // 'Tennis', 'Golf', 'Boxing', 'MMA', etc
  country: text("country"), // Optional country code
  photoFilename: text("photo_filename").notNull(), // Filename in server/public/ (e.g., 'custom-player-123.png')
  uploadedBy: text("uploaded_by").references(() => users.id),
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertCustomPlayerSchema = createInsertSchema(customPlayers).omit({
  id: true,
  createdAt: true,
});

export type InsertCustomPlayer = z.infer<typeof insertCustomPlayerSchema>;
export type CustomPlayer = typeof customPlayers.$inferSelect;

// Custom Leagues - User-uploaded leagues
export const customLeagues = sqliteTable("custom_leagues", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  name: text("name").notNull(),
  sport: text("sport").notNull(), // 'Basketball', 'Football', 'Baseball', 'Soccer', etc
  badgeFilename: text("badge_filename"), // Optional league logo/badge
  uploadedBy: text("uploaded_by").references(() => users.id),
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertCustomLeagueSchema = createInsertSchema(customLeagues).omit({
  id: true,
  createdAt: true,
});

export type InsertCustomLeague = z.infer<typeof insertCustomLeagueSchema>;
export type CustomLeague = typeof customLeagues.$inferSelect;

// Visibility Settings - Controls which sports and leagues are visible on the platform
export const visibilitySettings = sqliteTable("visibility_settings", {
  id: text("id").primaryKey().$defaultFn(() => crypto.randomUUID()),
  type: text("type").notNull(), // 'sport' or 'league'
  sportId: text("sport_id"), // ID from sportsData in sports-leagues.ts (e.g., 'basketball', 'soccer')
  leagueId: text("league_id"), // ID from league in sports-leagues.ts (e.g., '4328')
  isVisible: integer("is_visible", { mode: 'boolean' }).notNull().default(true),
  manualOverride: integer("manual_override", { mode: 'boolean' }).notNull().default(false), // True if admin manually set visibility (overrides auto-hide)
  updatedBy: text("updated_by").references(() => users.id),
  createdAt: integer("created_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
  updatedAt: integer("updated_at", { mode: 'timestamp' }).$defaultFn(() => new Date()),
});

export const insertVisibilitySettingSchema = createInsertSchema(visibilitySettings).omit({
  id: true,
  createdAt: true,
  updatedAt: true,
});

export type InsertVisibilitySetting = z.infer<typeof insertVisibilitySettingSchema>;
export type VisibilitySetting = typeof visibilitySettings.$inferSelect;
