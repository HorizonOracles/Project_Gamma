// DegenArena API Routes
import type { Express } from "express";
import { createServer, type Server } from "http";
import { WebSocketServer, WebSocket } from "ws";
import { storage } from "./storage";
import { registerBlockchainAdminRoutes } from "./blockchain-admin";
import { setupSIWEAuth, requireAuth, requireAdmin } from "./siwe-auth";
import { randomBytes } from "crypto";
import multer from "multer";
import path from "path";
import { sportsData } from "@shared/market-categories";

export async function registerRoutes(app: Express): Promise<Server> {
  // ============================================================================
  // SIWE AUTHENTICATION
  // ============================================================================
  
  // Setup Sign-In With Ethereum authentication
  await setupSIWEAuth(app);

  // ============================================================================
  // BLOCKCHAIN ADMIN ROUTES
  // ============================================================================
  
  // Register blockchain admin routes for whitelisted wallet addresses
  registerBlockchainAdminRoutes(app);

  // ============================================================================
  // MARKET ROUTES (Public - no auth required)
  // ============================================================================

  // Get all active markets (filtered by visibility settings)
  app.get('/api/markets', async (req, res) => {
    try {
      const markets = await storage.getAllActiveMarkets();
      const visibilitySettings = await storage.getVisibilitySettings();
      
      // Create a map of sport names to sport IDs (handles all sports including esports)
      const sportNameToId = new Map<string, string>();
      sportsData.forEach(sport => {
        sportNameToId.set(sport.name, sport.id);
      });
      
      // Create a map of league names to league IDs
      const leagueNameToId = new Map<string, string>();
      sportsData.forEach(sport => {
        sport.leagues.forEach(league => {
          leagueNameToId.set(league.name, league.id);
        });
      });
      
      // Filter markets based on visibility settings
      const filteredMarkets = markets.filter(market => {
        // Get sport ID from market using the sportNameToId map
        const sportId = sportNameToId.get(market.sport);
        
        if (sportId) {
          // Check sport visibility (default to visible if no setting)
          const sportSetting = visibilitySettings.find(s => s.type === 'sport' && s.sportId === sportId);
          const isSportVisible = sportSetting ? sportSetting.isVisible : true;
          
          if (!isSportVisible) {
            return false; // Hide entire sport
          }
        }
        
        // Check league visibility
        const leagueId = leagueNameToId.get(market.league);
        if (leagueId) {
          const leagueSetting = visibilitySettings.find(s => s.type === 'league' && s.leagueId === leagueId);
          const isLeagueVisible = leagueSetting ? leagueSetting.isVisible : true;
          
          if (!isLeagueVisible) {
            return false; // Hide markets from this league
          }
        }
        
        return true;
      });
      
      res.json(filteredMarkets);
    } catch (error) {
      console.error("Error fetching markets:", error);
      res.status(500).json({ message: "Failed to fetch markets" });
    }
  });

  // Note: Admin authentication is handled by blockchain-admin.ts using wallet whitelist

  // Get market by ID
  app.get('/api/markets/:id', async (req, res) => {
    try {
      const market = await storage.getMarketById(req.params.id);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }
      res.json(market);
    } catch (error) {
      console.error("Error fetching market:", error);
      res.status(500).json({ message: "Failed to fetch market" });
    }
  });

  // Settle market with winning outcome (public for now - should be protected by blockchain)
  app.post('/api/markets/:id/settle', async (req: any, res) => {
    const { db } = await import('./db');
    
    try {
      const { winningOutcome } = req.body;
      
      if (!winningOutcome || (winningOutcome !== 'A' && winningOutcome !== 'B')) {
        return res.status(400).json({ message: "Invalid winning outcome. Must be 'A' or 'B'" });
      }
      
      // Execute settlement in atomic transaction
      const result = await db.transaction(async (tx) => {
        // Fetch market inside transaction for consistent read
        const market = await storage.getMarketById(req.params.id, tx);
        if (!market) {
          throw new Error("Market not found");
        }
        
        if (market.status === 'settled') {
          throw new Error("Market already settled");
        }
        
        // Calculate pool totals from transactional read
        const totalPool = parseFloat(market.poolATotal) + parseFloat(market.poolBTotal);
        const winningPool = winningOutcome === 'A' ? parseFloat(market.poolATotal) : parseFloat(market.poolBTotal);
        
        // Check for zero winning pool (no bets on winning outcome)
        if (winningPool === 0) {
          throw new Error("Cannot settle market - no bets placed on winning outcome");
        }
        
        // Settle the market (use tx parameter)
        const settledMarket = await storage.settleMarket(req.params.id, winningOutcome, tx);
        
        // Get all bets on this market (use tx parameter)
        const bets = await storage.getBetsByMarketId(req.params.id, tx);
        
        // Process each bet
        const payouts = [];
        for (const bet of bets) {
          if (bet.outcome === winningOutcome) {
            // Calculate payout: (bet amount / winning pool) * total pool
            const betAmount = parseFloat(bet.amount);
            const payout = (betAmount / winningPool) * totalPool;
            const actualPayout = payout.toString();
            
            // Update bet status to won (use tx parameter)
            await storage.updateBetStatus(bet.id, 'won', actualPayout, tx);
            
            // Get user's wallet and add payout (use tx parameter)
            const wallet = await storage.getWalletByUserId(bet.userId, tx);
            if (wallet) {
              const newBalance = (parseFloat(wallet.balance) + payout).toString();
              await storage.updateWalletBalance(wallet.id, newBalance, tx);
              
              // Create payout transaction (use tx parameter)
              await storage.createTransaction({
                walletId: wallet.id,
                userId: bet.userId,
                type: 'bet_won',
                amount: actualPayout,
                status: 'completed',
                metadata: {
                  betId: bet.id,
                  marketId: market.id,
                },
              }, tx);
              
              // Update user stats (won amount) - use tx parameter
              await storage.updateUserStats(bet.userId, undefined, actualPayout, tx);
              
              payouts.push({ userId: bet.userId, amount: payout });
            }
          } else {
            // Losing bet (use tx parameter)
            await storage.updateBetStatus(bet.id, 'lost', '0', tx);
          }
        }
        
        return { settledMarket, payouts };
      });
      
      // Broadcast market settlement
      if (app.locals.broadcastMarketUpdate) {
        app.locals.broadcastMarketUpdate({
          id: result.settledMarket.id,
          status: 'settled',
          winningOutcome,
        });
      }
      
      res.json({ 
        message: "Market settled successfully", 
        market: result.settledMarket,
        payoutsProcessed: result.payouts.length,
      });
    } catch (error: any) {
      console.error("Error settling market:", error);
      // Return specific error messages from transaction
      if (error.message === "Market not found") {
        return res.status(404).json({ message: error.message });
      }
      if (error.message === "Market already settled" || 
          error.message === "Cannot settle market - no bets placed on winning outcome") {
        return res.status(400).json({ message: error.message });
      }
      res.status(500).json({ message: "Failed to settle market" });
    }
  });

  // ============================================================================
  // VISIBILITY SETTINGS ROUTES
  // ============================================================================

  // Get all visibility settings (public - needed for sidebar filtering for all users)
  app.get('/api/visibility-settings', async (req, res) => {
    try {
      const settings = await storage.getVisibilitySettings();
      res.json(settings);
    } catch (error) {
      console.error("Error fetching visibility settings:", error);
      res.status(500).json({ message: "Failed to fetch visibility settings" });
    }
  });

  // Toggle sport visibility (wallet auth handled by blockchain-admin)
  app.post('/api/admin/visibility/sport/:sportId', requireAdmin, async (req: any, res) => {
    try {
      const { sportId } = req.params;
      const { isVisible, userId } = req.body;

      if (typeof isVisible !== 'boolean') {
        return res.status(400).json({ message: "isVisible must be a boolean" });
      }

      const setting = await storage.toggleSportVisibility(sportId, isVisible, userId || 'admin');
      res.json(setting);
    } catch (error) {
      console.error("Error toggling sport visibility:", error);
      res.status(500).json({ message: "Failed to toggle sport visibility" });
    }
  });

  // Toggle league visibility (wallet auth handled by blockchain-admin)
  app.post('/api/admin/visibility/league/:leagueId', requireAdmin, async (req: any, res) => {
    try {
      const { leagueId } = req.params;
      const { sportId, isVisible, userId } = req.body;

      if (!sportId) {
        return res.status(400).json({ message: "sportId is required" });
      }

      if (typeof isVisible !== 'boolean') {
        return res.status(400).json({ message: "isVisible must be a boolean" });
      }

      const setting = await storage.toggleLeagueVisibility(leagueId, sportId, isVisible, userId || 'admin');
      res.json(setting);
    } catch (error) {
      console.error("Error toggling league visibility:", error);
      res.status(500).json({ message: "Failed to toggle league visibility" });
    }
  });

  // ============================================================================
  // BETTING ROUTES (DISABLED - Requires wallet authentication to be implemented)
  // ============================================================================

  // Place a bet (DISABLED - requires wallet auth)
  // TODO: Re-implement with wallet-based authentication
  /*
  app.post('/api/bets', async (req: any, res) => {
    const { db } = await import('./db');
    
    try {
      const userId = req.user.claims.sub;
      const { marketId, outcome, amount, oddsAtBet } = req.body;
      
      if (!marketId || !outcome || !amount || !oddsAtBet) {
        return res.status(400).json({ message: "Missing required fields" });
      }
      
      if (outcome !== 'A' && outcome !== 'B') {
        return res.status(400).json({ message: "Invalid outcome" });
      }
      
      if (parseFloat(amount) <= 0) {
        return res.status(400).json({ message: "Invalid bet amount" });
      }
      
      // Get user's wallet
      const wallet = await storage.getWalletByUserId(userId);
      if (!wallet) {
        return res.status(404).json({ message: "Wallet not found" });
      }
      
      const currentBalance = parseFloat(wallet.balance);
      const betAmount = parseFloat(amount);
      
      if (currentBalance < betAmount) {
        return res.status(400).json({ message: "Insufficient balance" });
      }
      
      // Get market
      const market = await storage.getMarketById(marketId);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }
      
      if (market.status !== 'active') {
        return res.status(400).json({ message: "Market is not accepting bets" });
      }
      
      // Execute all operations in a database transaction for atomicity
      const result = await db.transaction(async (tx) => {
        // Deduct from wallet
        const newBalance = (currentBalance - betAmount).toString();
        await storage.updateWalletBalance(wallet.id, newBalance, tx);
        
        // Update market pools
        const poolATotal = parseFloat(market.poolATotal);
        const poolBTotal = parseFloat(market.poolBTotal);
        
        const newPoolATotal = outcome === 'A' 
          ? (poolATotal + betAmount).toString()
          : poolATotal.toString();
        const newPoolBTotal = outcome === 'B' 
          ? (poolBTotal + betAmount).toString()
          : poolBTotal.toString();
        
        await storage.updateMarketPools(marketId, newPoolATotal, newPoolBTotal, tx);
        
        // Create bet
        const bet = await storage.createBet({
          userId,
          marketId,
          walletId: wallet.id,
          outcome,
          amount,
          oddsAtBet,
        }, tx);
        
        // Create transaction record
        await storage.createTransaction({
          walletId: wallet.id,
          userId,
          type: 'bet_placed',
          amount,
          status: 'completed',
          metadata: { 
            betId: bet.id,
            marketId,
            outcome,
          },
        }, tx);
        
        // Update user stats (wagered amount)
        await storage.updateUserStats(userId, amount, undefined, tx);
        
        return { bet, updatedMarket: { ...market, poolATotal: newPoolATotal, poolBTotal: newPoolBTotal } };
      });
      
      // Recalculate odds after pool updates
      const totalPool = parseFloat(result.updatedMarket.poolATotal) + parseFloat(result.updatedMarket.poolBTotal);
      const oddsA = totalPool > 0 ? (totalPool / parseFloat(result.updatedMarket.poolATotal)).toFixed(2) : "0.00";
      const oddsB = totalPool > 0 ? (totalPool / parseFloat(result.updatedMarket.poolBTotal)).toFixed(2) : "0.00";
      
      // Broadcast new bet to all WebSocket clients for live feed updates
      if (app.locals.broadcastBet) {
        const betPayload = {
          id: result.bet.id,
          userEmail: req.user?.email || "Anonymous",
          amount: result.bet.amount,
          outcome: result.bet.outcome,
          teamA: market.teamA,
          teamB: market.teamB,
          marketDescription: market.description,
          oddsAtBet: result.bet.oddsAtBet,
          createdAt: result.bet.createdAt,
          status: result.bet.status,
        };
        app.locals.broadcastBet(betPayload);
      }
      
      // Broadcast updated market state for real-time odds updates
      if (app.locals.broadcastMarketUpdate) {
        const marketUpdate = {
          id: marketId,
          poolATotal: result.updatedMarket.poolATotal,
          poolBTotal: result.updatedMarket.poolBTotal,
          oddsA,
          oddsB,
        };
        app.locals.broadcastMarketUpdate(marketUpdate);
      }
      
      res.json(result.bet);
    } catch (error) {
      console.error("Error placing bet:", error);
      res.status(500).json({ message: "Failed to place bet" });
    }
  });
  */

  // Get user's bets (DISABLED - requires wallet auth)
  /*
  app.get('/api/bets/my-bets', async (req: any, res) => {
    try {
      const userId = req.user.claims.sub;
      const bets = await storage.getBetsByUserId(userId);
      res.json(bets);
    } catch (error) {
      console.error("Error fetching bets:", error);
      res.status(500).json({ message: "Failed to fetch bets" });
    }
  });
  */

  // Get betting feed (recent bets from all users)
  app.get('/api/bets/feed', async (req, res) => {
    try {
      const limit = parseInt(req.query.limit as string) || 20;
      const feed = await storage.getRecentBetsForFeed(limit);
      res.json(feed);
    } catch (error) {
      console.error("Error fetching betting feed:", error);
      res.status(500).json({ message: "Failed to fetch betting feed" });
    }
  });

  // ============================================================================
  // CRYPTO PRICE API (For currency conversion)
  // ============================================================================

  // Get crypto prices from CoinGecko (BNB to multiple currencies)
  app.get('/api/crypto/prices', async (req: any, res) => {
    try {
      const currencies = req.query.currencies || 'usd,eur,gbp,jpy,cad,aud';
      const response = await fetch(
        `https://api.coingecko.com/api/v3/simple/price?ids=binancecoin&vs_currencies=${currencies}&include_24hr_change=true`
      );
      
      if (!response.ok) {
        throw new Error('Failed to fetch crypto prices');
      }
      
      const data = await response.json();
      res.json(data.binancecoin || {});
    } catch (error) {
      console.error("Error fetching crypto prices:", error);
      // Return fallback prices if API fails
      res.json({
        usd: 600,
        eur: 550,
        gbp: 475,
        jpy: 90000,
        cad: 820,
        aud: 920,
      });
    }
  });

  // Get user stats (wallet + rank + XP all in one) (DISABLED - requires wallet auth)
  /*
  app.get('/api/user/stats', async (req: any, res) => {
    try {
      const userId = req.user.claims.sub;
      const user = await storage.getUser(userId);
      
      if (!user) {
        return res.status(404).json({ message: "User not found" });
      }
      
      let wallet = await storage.getWalletByUserId(userId);
      
      // Auto-create wallet if it doesn't exist
      if (!wallet) {
        const bnbAddress = `0x${randomBytes(20).toString('hex')}`;
        wallet = await storage.createWallet({
          userId,
          bnbAddress,
          balance: "0",
        });
      }
      
      // Calculate XP progress to next rank - 12 tier progressive system
      const rankThresholds = {
        Bronze: { min: 0, max: 10000 },
        Silver: { min: 10000, max: 30000 },
        Gold: { min: 30000, max: 60000 },
        Sapphire: { min: 60000, max: 100000 },
        Emerald: { min: 100000, max: 150000 },
        Ruby: { min: 150000, max: 210000 },
        Diamond: { min: 210000, max: 280000 },
        Pearl: { min: 280000, max: 360000 },
        Opal: { min: 360000, max: 450000 },
        Stardust: { min: 450000, max: 550000 },
        Nebula: { min: 550000, max: 660000 },
        Supernova: { min: 660000, max: Infinity },
      };
      
      // Map legacy ranks (Platinum, Obsidian) to new tier equivalents based on actual XP
      let mappedRank = user.rank;
      
      // For Platinum and Obsidian users (legacy ranks), recalculate rank based on their actual XP
      if (user.rank === "Platinum" || user.rank === "Obsidian") {
        const xp = user.rankPoints;
        if (xp >= 660000) mappedRank = "Supernova";
        else if (xp >= 550000) mappedRank = "Nebula";
        else if (xp >= 450000) mappedRank = "Stardust";
        else if (xp >= 360000) mappedRank = "Opal";
        else if (xp >= 280000) mappedRank = "Pearl";
        else if (xp >= 210000) mappedRank = "Diamond";
        else if (xp >= 150000) mappedRank = "Ruby";
        else if (xp >= 100000) mappedRank = "Emerald";
        else if (xp >= 60000) mappedRank = "Sapphire";
        else if (xp >= 30000) mappedRank = "Gold";
        else if (xp >= 10000) mappedRank = "Silver";
        else mappedRank = "Bronze";
      }
      
      const currentRank = mappedRank as keyof typeof rankThresholds;
      const threshold = rankThresholds[currentRank] || rankThresholds.Bronze; // Fallback to Bronze for old ranks
      const progress = threshold.max === Infinity 
        ? 100 
        : Math.min(100, ((user.rankPoints - threshold.min) / (threshold.max - threshold.min)) * 100);
      
      res.json({
        user: {
          id: user.id,
          email: user.email,
          firstName: user.firstName,
          lastName: user.lastName,
          profileImageUrl: user.profileImageUrl,
          nameColor: user.nameColor,
          rank: mappedRank, // Return mapped rank for UI consistency
          rankPoints: user.rankPoints,
          totalWagered: user.totalWagered,
          totalWon: user.totalWon,
        },
        wallet: {
          id: wallet.id,
          balance: wallet.balance,
          bnbAddress: wallet.bnbAddress,
        },
        rankProgress: {
          current: user.rankPoints,
          min: threshold.min,
          max: threshold.max === Infinity ? user.rankPoints : threshold.max,
          percentage: progress,
          nextRank: threshold.max === Infinity ? 'Max Rank' : Object.keys(rankThresholds)[Object.keys(rankThresholds).indexOf(currentRank) + 1],
        },
      });
    } catch (error) {
      console.error("Error fetching user stats:", error);
      res.status(500).json({ message: "Failed to fetch user stats" });
    }
  });
  */

  // ============================================================================
  // CRYPTO PRICE API (For currency conversion)
  // ============================================================================

  // Note: Custom sports upload routes removed (depended on Replit auth)

  // Get custom teams by sport/league
  app.get('/api/sports/custom/teams', async (req, res) => {
    try {
      const { sport, league } = req.query;
      const teams = await storage.getCustomTeams(
        sport as string | undefined,
        league as string | undefined
      );
      res.json(teams);
    } catch (error) {
      console.error("Error fetching custom teams:", error);
      res.status(500).json({ message: "Failed to fetch custom teams" });
    }
  });

  // Get custom players by sport
  app.get('/api/sports/custom/players', async (req, res) => {
    try {
      const { sport } = req.query;
      const players = await storage.getCustomPlayers(sport as string | undefined);
      res.json(players);
    } catch (error) {
      console.error("Error fetching custom players:", error);
      res.status(500).json({ message: "Failed to fetch custom players" });
    }
  });

  // Get custom leagues by sport
  app.get('/api/sports/custom/leagues', async (req, res) => {
    try {
      const { sport } = req.query;
      const leagues = await storage.getCustomLeagues(sport as string | undefined);
      res.json(leagues);
    } catch (error) {
      console.error("Error fetching custom leagues:", error);
      res.status(500).json({ message: "Failed to fetch custom leagues" });
    }
  });

  // ============================================================================
  // WEBSOCKET SERVER FOR REAL-TIME UPDATES
  // ============================================================================
  // Reference: javascript_websocket blueprint
  
  const httpServer = createServer(app);
  const wss = new WebSocketServer({ server: httpServer, path: '/ws' });

  wss.on('connection', (ws: WebSocket) => {
    console.log('New WebSocket client connected');

    // Send initial connection message
    ws.send(JSON.stringify({ 
      type: 'connected',
      message: 'Connected to DegenArena live updates',
    }));

    // Broadcast market updates every 5 seconds
    const updateInterval = setInterval(async () => {
      if (ws.readyState === WebSocket.OPEN) {
        try {
          const markets = await storage.getAllActiveMarkets();
          ws.send(JSON.stringify({
            type: 'markets_update',
            data: markets,
          }));
        } catch (error) {
          console.error('Error sending market updates:', error);
        }
      }
    }, 5000);

    ws.on('close', () => {
      console.log('WebSocket client disconnected');
      clearInterval(updateInterval);
    });

    ws.on('error', (error: Error) => {
      console.error('WebSocket error:', error);
      clearInterval(updateInterval);
    });
  });

  // Broadcast function for bet placement events
  app.locals.broadcastBet = (bet: any) => {
    wss.clients.forEach((client: WebSocket) => {
      if (client.readyState === WebSocket.OPEN) {
        client.send(JSON.stringify({
          type: 'new_bet',
          data: bet,
        }));
      }
    });
  };

  // Broadcast function for market updates (real-time odds)
  app.locals.broadcastMarketUpdate = (marketUpdate: any) => {
    wss.clients.forEach((client: WebSocket) => {
      if (client.readyState === WebSocket.OPEN) {
        client.send(JSON.stringify({
          type: 'market_update',
          data: marketUpdate,
        }));
      }
    });
  };

  // Broadcast function for chat messages
  app.locals.broadcastChatMessage = (chatMessage: any) => {
    wss.clients.forEach((client: WebSocket) => {
      if (client.readyState === WebSocket.OPEN) {
        client.send(JSON.stringify({
          type: 'chat_message',
          data: chatMessage,
        }));
      }
    });
  };

  return httpServer;
}
