// Sign-In With Ethereum (SIWE) Authentication
import type { Express, RequestHandler } from "express";
import { SiweMessage } from "siwe";
import { getIronSession } from "iron-session";
import { storage } from "./storage";

// Session configuration
export const sessionOptions = {
  password: process.env.SESSION_SECRET || 'dev-secret-change-in-production-must-be-at-least-32-chars-long',
  cookieName: "siwe-session",
  cookieOptions: {
    secure: process.env.NODE_ENV === "production",
    httpOnly: true,
    maxAge: 7 * 24 * 60 * 60, // 1 week in seconds
  },
};

// Session data type
export interface SessionData {
  nonce?: string;
  siwe?: {
    address: string;
    chainId: number;
  };
  // Legacy fields for backward compatibility
  address?: string;
  chainId?: number;
}

// Get session from request
async function getSession(req: any, res: any) {
  return await getIronSession<SessionData>(req, res, sessionOptions);
}

// Setup SIWE authentication routes
export async function setupSIWEAuth(app: Express) {
  // Generate nonce for signing
  app.get("/api/auth/nonce", async (req, res) => {
    try {
      const session = await getSession(req, res);
      
      // Generate random nonce
      const nonce = Math.random().toString(36).substring(2);
      session.nonce = nonce;
      await session.save();
      
      res.json({ nonce });
    } catch (error) {
      console.error("Error generating nonce:", error);
      res.status(500).json({ error: "Failed to generate nonce" });
    }
  });

  // Verify SIWE message and establish session
  app.post("/api/auth/verify", async (req, res) => {
    try {
      const { message, signature } = req.body;
      
      if (!message || !signature) {
        return res.status(400).json({ error: "Missing message or signature" });
      }

      const session = await getSession(req, res);
      
      if (!session.nonce) {
        return res.status(400).json({ error: "No nonce found in session" });
      }

      // Parse and verify the SIWE message
      const siweMessage = new SiweMessage(message);
      const fields = await siweMessage.verify({ signature, nonce: session.nonce });

      if (!fields.success) {
        return res.status(401).json({ error: "Invalid signature" });
      }

      // Store address in session (both formats for compatibility)
      session.siwe = {
        address: fields.data.address,
        chainId: fields.data.chainId,
      };
      session.address = fields.data.address;
      session.chainId = fields.data.chainId;
      delete session.nonce; // Clear nonce after use
      await session.save();

      // Upsert user in database using wallet address as ID
      await storage.upsertUser({
        id: fields.data.address.toLowerCase(),
        email: `${fields.data.address.toLowerCase()}@wallet.address`, // Placeholder email
        firstName: fields.data.address.substring(0, 6),
        lastName: fields.data.address.substring(38),
        profileImageUrl: `https://api.dicebear.com/7.x/identicon/svg?seed=${fields.data.address}`,
      });

      res.json({ 
        success: true, 
        address: fields.data.address,
        chainId: fields.data.chainId,
      });
    } catch (error) {
      console.error("Error verifying SIWE message:", error);
      res.status(500).json({ error: "Failed to verify signature" });
    }
  });

  // Get current user session
  app.get("/api/auth/me", async (req, res) => {
    try {
      const session = await getSession(req, res);
      
      if (!session.address) {
        return res.status(401).json({ error: "Not authenticated" });
      }

      // Get user from database
      const user = await storage.getUser(session.address.toLowerCase());
      
      if (!user) {
        return res.status(404).json({ error: "User not found" });
      }

      res.json({
        address: session.address,
        chainId: session.chainId,
        user,
      });
    } catch (error) {
      console.error("Error getting user session:", error);
      res.status(500).json({ error: "Failed to get user session" });
    }
  });

  // Logout (clear session)
  app.post("/api/auth/logout", async (req, res) => {
    try {
      const session = await getSession(req, res);
      session.destroy();
      res.json({ success: true });
    } catch (error) {
      console.error("Error logging out:", error);
      res.status(500).json({ error: "Failed to logout" });
    }
  });

  // Get user by wallet address endpoint (for querying user data)
  app.get("/api/auth/user", async (req, res) => {
    try {
      const session = await getSession(req, res);
      
      if (!session.address) {
        return res.status(401).json({ error: "Not authenticated" });
      }

      const user = await storage.getUser(session.address.toLowerCase());
      
      if (!user) {
        return res.status(404).json({ error: "User not found" });
      }

      res.json(user);
    } catch (error) {
      console.error("Error fetching user:", error);
      res.status(500).json({ error: "Failed to fetch user" });
    }
  });
}

// Middleware to check if user is authenticated
export const requireAuth: RequestHandler = async (req: any, res, next) => {
  try {
    const session = await getSession(req, res);
    
    if (!session.address) {
      return res.status(401).json({ error: "Not authenticated" });
    }

    // Attach user address to request for downstream use
    req.userAddress = session.address.toLowerCase();
    next();
  } catch (error) {
    console.error("Error checking authentication:", error);
    res.status(500).json({ error: "Authentication check failed" });
  }
};

// Middleware to check if user is admin (specific wallet address)
export const requireAdmin: RequestHandler = async (req: any, res, next) => {
  try {
    const session = await getSession(req, res);
    
    if (!session.address) {
      return res.status(401).json({ error: "Not authenticated" });
    }

    // Admin address from environment (0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444)
    const adminAddress = (process.env.ADMIN_WALLET_ADDRESS || process.env.VITE_ADMIN_WALLET_ADDRESS)?.toLowerCase();
    const userAddress = session.address.toLowerCase();

    if (!adminAddress) {
      console.error("ADMIN_WALLET_ADDRESS not configured in environment");
      return res.status(500).json({ error: "Admin configuration error" });
    }

    if (userAddress !== adminAddress) {
      console.warn(`Unauthorized admin access attempt from: ${userAddress}`);
      return res.status(403).json({ error: "Forbidden - Admin access required" });
    }

    req.userAddress = userAddress;
    next();
  } catch (error) {
    console.error("Error checking admin authorization:", error);
    res.status(500).json({ error: "Authorization check failed" });
  }
};
