// Seed Admin Wallet
// Adds an admin wallet address to the whitelist
import { db } from "./db";
import { adminWhitelist } from "../shared/schema";
import { eq } from "drizzle-orm";

const ADMIN_WALLET = "0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66";

async function seedAdmin() {
  console.log("🔐 Seeding admin wallet...");
  
  try {
    // Check if admin already exists
    const existing = await db
      .select()
      .from(adminWhitelist)
      .where(eq(adminWhitelist.walletAddress, ADMIN_WALLET));
    
    if (existing.length > 0) {
      console.log(`⚠️  Admin wallet ${ADMIN_WALLET} already exists in whitelist`);
      return;
    }
    
    // Add admin to whitelist
    await db.insert(adminWhitelist).values({
      id: crypto.randomUUID(),
      walletAddress: ADMIN_WALLET,
      notes: "Initial admin wallet",
      isActive: true,
      createdAt: new Date(),
    });
    
    console.log(`✅ Successfully added admin wallet: ${ADMIN_WALLET}`);
  } catch (error) {
    console.error("❌ Error seeding admin:", error);
    throw error;
  }
}

// Run seed if this file is executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  seedAdmin()
    .then(() => {
      console.log("✅ Admin seeding complete");
      process.exit(0);
    })
    .catch((error) => {
      console.error("❌ Admin seeding failed:", error);
      process.exit(1);
    });
}

export { seedAdmin };
