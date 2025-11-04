// Blockchain Admin Routes
// Handles blockchain interactions for admin wallet
import type { Express } from "express";
import { storage } from "./storage";
import { requireAdmin } from "./siwe-auth";

// Environment configuration for blockchain
// BNB Chain Testnet: Chain ID 97
// BNB Chain Mainnet: Chain ID 56
const CHAIN_ID = parseInt(process.env.CHAIN_ID || "97"); 
const CHAIN_RPC_URL = process.env.CHAIN_RPC_URL || "https://bsc-dataseed.binance.org/";
const MARKET_FACTORY_ADDRESS = process.env.MARKET_FACTORY_ADDRESS || "0x22Cc806047BB825aa26b766Af737E92B1866E8A6";

/**
 * Register blockchain admin routes
 * These routes handle admin actions that interact with smart contracts
 */
export function registerBlockchainAdminRoutes(app: Express) {
  
  // ============================================================================
  // MARKET CREATION (BLOCKCHAIN)
  // ============================================================================

  /**
   * Simple market creation endpoint for database storage
   * Used by MarketCreationForm component
   */
  app.post('/api/admin/markets', requireAdmin, async (req: any, res) => {
    try {
      const { 
        sport, 
        league, 
        marketType,
        teamA, 
        teamB, 
        teamALogo,
        teamBLogo,
        description, 
        gameTime,
        category = 'sports',
        resolutionSource = 'manual',
      } = req.body;

      // Validate required fields
      if (!sport || !league || !teamA || !teamB || !description || !gameTime) {
        return res.status(400).json({ 
          message: "Missing required fields: sport, league, teamA, teamB, description, gameTime" 
        });
      }

      // Convert gameTime to Date
      const gameTimeDate = new Date(gameTime);
      if (gameTimeDate.getTime() <= Date.now()) {
        return res.status(400).json({ 
          message: "Game time must be in the future" 
        });
      }

      // Store market in database
      const market = await storage.createMarket({
        sport,
        league,
        marketType: marketType || 'match_winner',
        teamA,
        teamB,
        teamALogo: teamALogo || null,
        teamBLogo: teamBLogo || null,
        description,
        gameTime: gameTimeDate,
        status: 'pending', // Will be updated to 'active' after on-chain creation
        resolutionSource,
        isLive: false,
      });

      res.json({
        success: true,
        message: "Market created in database",
        market,
      });
    } catch (error) {
      console.error("Error creating market:", error);
      res.status(500).json({ 
        success: false,
        message: "Failed to create market" 
      });
    }
  });

  /**
   * Create a new market on-chain
   * This route accepts market details and returns transaction data
   * The actual transaction should be signed and sent by the frontend wallet
   */
  app.post('/api/admin/markets/create', requireAdmin, async (req: any, res) => {
    try {
      const { 
        sport, 
        league, 
        marketType,
        teamA, 
        teamB, 
        teamALogo,
        teamBLogo,
        description, 
        gameTime,
        collateralToken,
        resolutionSource = 'manual',
      } = req.body;

      // Validate required fields
      if (!sport || !league || !marketType || !teamA || !teamB || !description || !gameTime) {
        return res.status(400).json({ 
          message: "Missing required fields: sport, league, marketType, teamA, teamB, description, gameTime" 
        });
      }

      // Convert gameTime to timestamp
      const eventTimestamp = Math.floor(new Date(gameTime).getTime() / 1000);
      if (eventTimestamp <= Math.floor(Date.now() / 1000)) {
        return res.status(400).json({ 
          message: "Game time must be in the future" 
        });
      }

      // Store market in database (without blockchain data initially)
      const market = await storage.createMarket({
        sport,
        league,
        marketType,
        teamA,
        teamB,
        teamALogo: teamALogo || null,
        teamBLogo: teamBLogo || null,
        description,
        gameTime: new Date(gameTime),
        status: 'pending', // Will be updated to 'active' after on-chain creation
        resolutionSource,
        isLive: false,
      });

      // Return market data and contract call parameters
      // Frontend will use this to create the transaction
      res.json({
        message: "Market prepared for blockchain creation",
        market,
        contractCallData: {
          contractAddress: MARKET_FACTORY_ADDRESS,
          chainId: CHAIN_ID,
          method: "createMarket",
          params: {
            collateralToken: collateralToken || "0x0000000000000000000000000000000000000000", // Native token (BNB)
            question: description,
            category: `${sport} - ${league}`,
            metadata: JSON.stringify({
              sport,
              league,
              marketType,
              teamA,
              teamB,
              teamALogo,
              teamBLogo,
              databaseId: market.id,
            }),
            eventTimestamp,
          },
        },
      });
    } catch (error) {
      console.error("Error preparing market creation:", error);
      res.status(500).json({ message: "Failed to prepare market creation" });
    }
  });

  /**
   * Update market with on-chain data after transaction is confirmed
   * Called by frontend after market creation transaction succeeds
   */
  app.post('/api/admin/markets/:marketId/confirm-blockchain', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      const { 
        marketAddress, 
        onChainMarketId, 
        yesTokenId, 
        noTokenId,
        transactionHash,
      } = req.body;

      if (!marketAddress || !onChainMarketId || !yesTokenId || !noTokenId) {
        return res.status(400).json({ 
          message: "Missing blockchain data: marketAddress, onChainMarketId, yesTokenId, noTokenId required" 
        });
      }

      // Update market with blockchain data
      const updatedMarket = await storage.updateMarketBlockchainData(marketId, {
        marketAddress: marketAddress.toLowerCase(),
        chainId: CHAIN_ID,
        onChainMarketId: onChainMarketId.toString(),
        yesTokenId: yesTokenId.toString(),
        noTokenId: noTokenId.toString(),
      });

      // Update market status to active now that it's on-chain
      await storage.updateMarketStatus(marketId, 'active');

      res.json({
        message: "Market blockchain data updated successfully",
        market: updatedMarket,
      });
    } catch (error) {
      console.error("Error confirming market blockchain data:", error);
      res.status(500).json({ message: "Failed to update market blockchain data" });
    }
  });

  // ============================================================================
  // LIQUIDITY MANAGEMENT
  // ============================================================================

  /**
   * Prepare liquidity addition transaction
   * Returns transaction data for frontend to sign and send
   */
  app.post('/api/admin/markets/:marketId/add-liquidity', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      const { amount } = req.body;

      if (!amount || parseFloat(amount) <= 0) {
        return res.status(400).json({ message: "Valid liquidity amount required" });
      }

      const market = await storage.getMarketById(marketId);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }

      if (!market.marketAddress) {
        return res.status(400).json({ message: "Market not deployed on-chain" });
      }

      // Return contract call data for frontend
      res.json({
        message: "Liquidity addition prepared",
        contractCallData: {
          contractAddress: market.marketAddress,
          chainId: CHAIN_ID,
          method: "addLiquidity",
          params: {
            amount: amount, // Will be converted to BigInt by frontend
          },
        },
      });
    } catch (error) {
      console.error("Error preparing liquidity addition:", error);
      res.status(500).json({ message: "Failed to prepare liquidity addition" });
    }
  });

  // ============================================================================
  // MARKET RESOLUTION
  // ============================================================================

  /**
   * Resolve market manually (admin action)
   * For markets with resolutionSource = 'manual'
   */
  app.post('/api/admin/markets/:marketId/resolve', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      const { winningOutcome } = req.body;

      if (!winningOutcome || (winningOutcome !== 'A' && winningOutcome !== 'B')) {
        return res.status(400).json({ message: "Invalid winning outcome. Must be 'A' or 'B'" });
      }

      const market = await storage.getMarketById(marketId);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }

      if (!market.marketAddress) {
        return res.status(400).json({ message: "Market not deployed on-chain" });
      }

      if (market.status === 'settled') {
        return res.status(400).json({ message: "Market already settled" });
      }

      // Map A/B to outcome index (0 = No, 1 = Yes)
      const outcomeIndex = winningOutcome === 'A' ? 1 : 0; // Assuming A = Yes, B = No

      // Return contract call data for resolution
      res.json({
        message: "Market resolution prepared",
        market,
        contractCallData: {
          contractAddress: market.marketAddress,
          chainId: CHAIN_ID,
          method: "proposeResolution",
          params: {
            outcomeIndex,
          },
        },
      });
    } catch (error) {
      console.error("Error preparing market resolution:", error);
      res.status(500).json({ message: "Failed to prepare market resolution" });
    }
  });

  /**
   * Confirm market resolution after on-chain transaction
   */
  app.post('/api/admin/markets/:marketId/confirm-resolution', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      const { winningOutcome, transactionHash } = req.body;

      if (!winningOutcome || !transactionHash) {
        return res.status(400).json({ 
          message: "Winning outcome and transaction hash required" 
        });
      }

      // Update market status in database
      const updatedMarket = await storage.settleMarket(marketId, winningOutcome);

      res.json({
        message: "Market resolved successfully",
        market: updatedMarket,
      });
    } catch (error) {
      console.error("Error confirming market resolution:", error);
      res.status(500).json({ message: "Failed to confirm market resolution" });
    }
  });

  // ============================================================================
  // ORACLE RESOLUTION (AI-POWERED)
  // ============================================================================

  /**
   * Request AI oracle resolution for a market
   * This triggers the AI resolver service
   */
  app.post('/api/admin/markets/:marketId/request-oracle', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      const { prompt } = req.body;

      const market = await storage.getMarketById(marketId);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }

      if (!market.marketAddress || !market.onChainMarketId) {
        return res.status(400).json({ message: "Market not deployed on-chain" });
      }

      // Default prompt if not provided
      const oraclePrompt = prompt || `Determine the outcome of: ${market.description}. Event between ${market.teamA} and ${market.teamB} on ${market.gameTime}.`;

      res.json({
        message: "Oracle resolution request prepared",
        market,
        contractCallData: {
          contractAddress: MARKET_FACTORY_ADDRESS,
          chainId: CHAIN_ID,
          method: "requestOracleResolution",
          params: {
            marketId: market.onChainMarketId,
            prompt: oraclePrompt,
          },
        },
      });
    } catch (error) {
      console.error("Error preparing oracle request:", error);
      res.status(500).json({ message: "Failed to prepare oracle request" });
    }
  });

  // ============================================================================
  // MARKET STATUS & INFO
  // ============================================================================

  /**
   * Get blockchain info for a market
   */
  app.get('/api/admin/markets/:marketId/blockchain-info', requireAdmin, async (req: any, res) => {
    try {
      const { marketId } = req.params;
      
      const market = await storage.getMarketById(marketId);
      if (!market) {
        return res.status(404).json({ message: "Market not found" });
      }

      res.json({
        market: {
          id: market.id,
          description: market.description,
          status: market.status,
        },
        blockchain: {
          chainId: market.chainId,
          marketAddress: market.marketAddress,
          onChainMarketId: market.onChainMarketId,
          yesTokenId: market.yesTokenId,
          noTokenId: market.noTokenId,
          resolutionSource: market.resolutionSource,
          oracleRequestId: market.oracleRequestId,
        },
        links: {
          explorer: market.marketAddress 
            ? `https://bscscan.com/address/${market.marketAddress}` 
            : null,
          yesToken: market.yesTokenId 
            ? `https://bscscan.com/token/${market.marketAddress}?a=${market.yesTokenId}` 
            : null,
          noToken: market.noTokenId 
            ? `https://bscscan.com/token/${market.marketAddress}?a=${market.noTokenId}` 
            : null,
        },
      });
    } catch (error) {
      console.error("Error fetching blockchain info:", error);
      res.status(500).json({ message: "Failed to fetch blockchain info" });
    }
  });

  // ============================================================================
  // CONFIGURATION
  // ============================================================================

  /**
   * Get blockchain configuration
   */
  app.get('/api/admin/blockchain/config', requireAdmin, async (req: any, res) => {
    try {
      res.json({
        chainId: CHAIN_ID,
        chainName: CHAIN_ID === 56 ? "BNB Chain" : "BNB Testnet",
        rpcUrl: CHAIN_RPC_URL,
        marketFactory: MARKET_FACTORY_ADDRESS,
        explorer: CHAIN_ID === 56 ? "https://bscscan.com" : "https://testnet.bscscan.com",
      });
    } catch (error) {
      console.error("Error fetching blockchain config:", error);
      res.status(500).json({ message: "Failed to fetch blockchain config" });
    }
  });
}
