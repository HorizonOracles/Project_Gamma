# Production Implementation Plan - Horizon Prediction Market
## From DegenArena to Full Blockchain Prediction Market

**Admin Address:** `0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`

---

## Overview

This plan transforms the existing DegenArena betting platform into a production-ready, blockchain-powered prediction market using the Project Gamma SDK. The platform will feature real on-chain markets with SIWE authentication, comprehensive admin controls, and user-friendly trading interfaces.

---

## Phase 1: Setup & Infrastructure

### 1.1 Update Blockchain Configuration
**Current:** Avalanche configuration  
**Target:** BNB Chain (BSC) Mainnet & Testnet

#### Actions:
- [ ] Update `client/src/lib/wagmi-config.ts` to use BNB Chain
  - Replace Avalanche chains with `bsc` (56) and `bscTestnet` (97)
  - Configure RPC endpoints for BNB Chain
  - Update WalletConnect configuration

- [ ] Update `.env.example` and environment variables:
  ```env
  # Blockchain - BNB Chain
  CHAIN_ID=56  # Mainnet (use 97 for testnet)
  CHAIN_RPC_URL=https://bsc-dataseed1.binance.org/
  
  # Contract Addresses (BNB Chain Mainnet)
  MARKET_FACTORY_ADDRESS=0x22Cc806047BB825aa26b766Af737E92B1866E8A6
  HORIZON_TOKEN_ADDRESS=0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444
  OUTCOME_TOKEN_ADDRESS=0x17B322784265c105a94e4c3d00aF1E5f46a5F311
  AI_ORACLE_ADAPTER_ADDRESS=0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93
  RESOLUTION_MODULE_ADDRESS=0xF0CF4C741910cB48AC596F620a0AE892Cd247838
  FEE_SPLITTER_ADDRESS=0x275017E98adF33051BbF477fe1DD197F681d4eF1
  
  # Admin Configuration
  ADMIN_ADDRESS=0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444
  
  # WalletConnect
  VITE_WALLETCONNECT_PROJECT_ID=your-project-id
  
  # Database
  DATABASE_PATH=./local.db
  SESSION_SECRET=generate-random-secret
  
  # Oracle API (for AI resolution)
  ORACLE_API_URL=https://api.projectgamma.io
  
  # IPFS (for market metadata)
  PINATA_JWT=your-pinata-jwt-token
  ```

### 1.2 SDK Integration
- [ ] Verify `@project-gamma/sdk` is properly installed (already in package.json)
- [ ] Create SDK configuration wrapper in `client/src/lib/sdk-config.ts`
- [ ] Set up GammaProvider in App.tsx

### 1.3 Database Schema Updates
- [ ] Add blockchain-specific fields to existing schema:
  ```typescript
  // Additional fields for markets table (already present):
  - marketAddress: on-chain contract address
  - chainId: chain ID
  - onChainMarketId: market ID from MarketFactory
  - yesTokenId: outcome token ID for YES
  - noTokenId: outcome token ID for NO
  - resolutionSource: 'oracle' | 'manual' | 'thesportsdb'
  - oracleRequestId: AI oracle request ID
  ```

- [ ] Add admin whitelist validation against `0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`

---

## Phase 2: Authentication - SIWE (Sign-In With Ethereum)

### 2.1 Remove Old Auth System
- [ ] Remove Replit Auth dependencies from `server/replitAuth.ts`
- [ ] Remove OpenID Client auth flow
- [ ] Keep SQLite database but update auth tables

### 2.2 Implement SIWE Authentication
- [ ] Install dependencies:
  ```bash
  npm install siwe iron-session
  ```

- [ ] Create `server/auth/siwe.ts`:
  ```typescript
  // Server-side SIWE verification
  // - Generate nonce
  // - Verify signature
  // - Create/update user session
  // - Store wallet address as primary identifier
  ```

- [ ] Update `shared/schema.ts`:
  ```typescript
  // Update users table to make walletAddress primary auth method
  export const users = sqliteTable("users", {
    id: text("id").primaryKey(),
    walletAddress: text("wallet_address").notNull().unique(),
    displayName: text("display_name"),
    profileImageUrl: text("profile_image_url"),
    role: text("role").notNull().default("user"), // 'user' or 'admin'
    // Remove email, firstName, lastName (not needed for wallet auth)
    createdAt: integer("created_at", { mode: 'timestamp' }),
    updatedAt: integer("updated_at", { mode: 'timestamp' }),
  });
  ```

- [ ] Create `client/src/hooks/useAuth.ts`:
  ```typescript
  // Client-side SIWE flow
  // - Connect wallet (via wagmi)
  // - Sign message
  // - Send to server for verification
  // - Store session
  ```

- [ ] Create login page at `client/src/pages/login.tsx`:
  - Wallet connection button
  - SIWE sign message prompt
  - Redirect to home after successful auth

### 2.3 Session Management
- [ ] Update session storage to use wallet addresses
- [ ] Implement session expiry (7 days default)
- [ ] Add logout functionality

---

## Phase 3: Admin System

### 3.1 Admin Address Configuration
**Fixed Admin Address:** `0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`

- [ ] Update `server/routes.ts` to add admin middleware:
  ```typescript
  function requireAdmin(req, res, next) {
    const ADMIN_ADDRESS = process.env.ADMIN_ADDRESS;
    if (req.user?.walletAddress?.toLowerCase() !== ADMIN_ADDRESS.toLowerCase()) {
      return res.status(403).json({ error: "Admin access required" });
    }
    next();
  }
  ```

- [ ] Update `client/src/hooks/useAdmin.ts`:
  ```typescript
  export function useAdmin() {
    const { address } = useAccount();
    const ADMIN_ADDRESS = "0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444";
    
    return {
      isAdmin: address?.toLowerCase() === ADMIN_ADDRESS.toLowerCase(),
      adminAddress: ADMIN_ADDRESS
    };
  }
  ```

### 3.2 Admin-Only Endpoints
- [ ] `/api/admin/markets/create` - Create market (on-chain + DB)
- [ ] `/api/admin/markets/:id/resolve` - Resolve market outcome
- [ ] `/api/admin/markets/:id/add-liquidity` - Add initial liquidity
- [ ] `/api/admin/markets/:id/invalidate` - Invalidate market
- [ ] `/api/admin/oracle/request` - Request AI oracle resolution

### 3.3 Admin UI Protection
- [ ] Add admin route guard in App.tsx
- [ ] Hide admin menu items for non-admin users
- [ ] Show admin badge in UI when logged in as admin

---

## Phase 4: SDK Integration

### 4.1 Wrap App with GammaProvider
Update `client/src/App.tsx`:

```typescript
import { GammaProvider } from '@project-gamma/sdk';
import { config } from './lib/wagmi-config';

function App() {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <GammaProvider
          chainId={Number(import.meta.env.VITE_CHAIN_ID) || 56}
          oracleApiUrl={import.meta.env.VITE_ORACLE_API_URL}
          pinataJwt={import.meta.env.VITE_PINATA_JWT}
          marketFactoryAddress={import.meta.env.VITE_MARKET_FACTORY_ADDRESS}
          horizonTokenAddress={import.meta.env.VITE_HORIZON_TOKEN_ADDRESS}
          outcomeTokenAddress={import.meta.env.VITE_OUTCOME_TOKEN_ADDRESS}
        >
          <RainbowKitProvider>
            {/* App routes */}
          </RainbowKitProvider>
        </GammaProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
```

### 4.2 Create Custom Hooks for Common Operations

- [ ] `client/src/hooks/useMarketOperations.ts`:
  ```typescript
  // Wrapper around SDK hooks with error handling
  // - useCreateMarketWithMetadata
  // - useBuyWithApproval
  // - useSellWithApproval
  ```

- [ ] `client/src/hooks/useUserPositions.ts`:
  ```typescript
  // Track user positions across all markets
  // - Fetch outcome token balances
  // - Calculate unrealized P&L
  // - Track total wagered
  ```

---

## Phase 5: Market Management (Admin)

### 5.1 Admin Market Creation Page
Location: `client/src/pages/admin/markets.tsx`

#### Features:
- [ ] Market creation form:
  - Market question (title)
  - Description
  - Category (Sports, Crypto, Politics, Weather, etc.)
  - Outcome options (Binary: YES/NO, Multi-choice: 3-8 options)
  - Close time (when market closes for trading)
  - Resolution source (Manual, Oracle, TheSportsDB)
  - Collateral token (USDC/BUSD)
  - Creator stake amount (HORIZON tokens)

- [ ] Upload market metadata to IPFS (using SDK's `useUploadMetadata`)

- [ ] Create on-chain market (using SDK's `useCreateMarket`)

- [ ] Store market in database with on-chain reference:
  ```typescript
  await db.insert(markets).values({
    onChainMarketId: receipt.marketId,
    marketAddress: marketContract.address,
    chainId: 56,
    yesTokenId: getOutcomeTokenId(marketId, 0),
    noTokenId: getOutcomeTokenId(marketId, 1),
    // ... other fields
  });
  ```

- [ ] Add initial liquidity (optional, using SDK's `useAddLiquidity`)

### 5.2 Admin Market List
- [ ] Display all markets (active, closed, resolved)
- [ ] Filter by status, category, date
- [ ] Quick actions: Add liquidity, Resolve, View details

### 5.3 Market Categories
Categories to support:
- **Sports**: NBA, NFL, MLB, NHL, Soccer, Tennis, MMA, Boxing, Golf, etc.
- **Crypto**: Bitcoin price predictions, ETH trends, altcoin movements
- **Politics**: Elections, policy outcomes, approval ratings
- **Weather**: Temperature ranges, rainfall predictions, natural events
- **Entertainment**: Awards shows, box office predictions, TV ratings
- **Economics**: Market indices, inflation rates, GDP forecasts
- **Science & Tech**: Product launches, scientific discoveries, tech trends

---

## Phase 6: Trading Interface (Users)

### 6.1 Market Discovery Page
Location: `client/src/pages/feed.tsx` (or create `markets.tsx`)

#### Features:
- [ ] Market cards showing:
  - Title/question
  - Category badge
  - Current odds (YES: 65¢, NO: 35¢)
  - Total volume (pool size)
  - Time remaining until close
  - Quick trade buttons

- [ ] Filter sidebar:
  - By category
  - By status (Active, Closed, Resolved)
  - By closing date
  - Search by keyword

- [ ] Sort options:
  - Most popular (by volume)
  - Closing soon
  - Recently created
  - Highest odds

### 6.2 Market Detail Page
Location: `client/src/pages/market/[id].tsx` (create new)

#### Features:
- [ ] Market information:
  - Full question and description
  - Category
  - Creator info
  - Close time
  - Resolution source
  - Total volume

- [ ] Current odds display:
  - Large odds display (YES: 65%, NO: 35%)
  - Price chart over time (optional for MVP)

- [ ] Trading interface:
  - Input amount (USDC/BUSD)
  - Select outcome (YES/NO)
  - Show quote (expected tokens, fees, price impact)
  - Approve collateral button (if needed)
  - Buy/Sell button
  - Slippage settings (default 0.5%)

- [ ] User position section:
  - Current holdings (YES tokens, NO tokens)
  - Potential payout
  - Unrealized P&L
  - Sell buttons

- [ ] Activity feed:
  - Recent trades
  - Large positions
  - Resolution updates

### 6.3 Trading Flow

**Buy Flow:**
1. User enters amount (e.g., 100 USDC)
2. Select outcome (YES)
3. SDK `useQuote` calculates expected tokens out
4. Check allowance with `useAllowance`
5. If insufficient, prompt approval with `useApprove`
6. Execute buy with `useBuy`
7. Show transaction status
8. Update user balance and positions

**Sell Flow:**
1. User sees current position (e.g., 150 YES tokens)
2. Enter amount to sell
3. SDK `useQuote` calculates USDC out
4. Execute sell with `useSell`
5. Show transaction status
6. Update user balance and positions

### 6.4 Transaction Notifications
- [ ] Use `useToast` for transaction updates:
  - Pending: "Transaction submitted..."
  - Success: "Trade successful! You received X tokens"
  - Error: "Transaction failed: [reason]"

---

## Phase 7: Dashboard & Tracking

### 7.1 User Dashboard
Location: `client/src/pages/profile.tsx` (update existing)

#### Wallet Connection
- [ ] RainbowKit Connect Wallet button (top right)
- [ ] Display connected address (shortened)
- [ ] Display balance (BNB, USDC, HORIZON)

#### Winnings Dashboard
- [ ] **Total Winnings Card**:
  - Sum of all resolved winning positions
  - Percentage return on total wagered
  - Compare to initial capital

- [ ] **Total Wagered Card**:
  - Sum of all collateral spent on trades
  - Number of markets participated in
  - Average wager size

- [ ] **Active Positions Card**:
  - Total value of current holdings
  - Unrealized P&L
  - Number of active markets

- [ ] **Resolved Markets Card**:
  - Win/loss record
  - Best prediction (highest return)
  - Worst prediction

### 7.2 Wager Tracking
Location: `client/src/pages/my-bets.tsx` (update existing)

#### Features:
- [ ] **Position List** (tabs: Active, History, Wins, Losses):
  - Market name
  - Outcome bet on (YES/NO)
  - Amount wagered
  - Current value (if active)
  - Profit/Loss
  - Status (Active, Won, Lost, Voided)
  - Action buttons (Sell, Redeem)

- [ ] **Filters**:
  - By category
  - By outcome
  - By date range
  - By P&L (winners/losers)

- [ ] **Detailed Position View**:
  - Click to expand
  - Show transaction history
  - Show current holdings
  - Calculate break-even price
  - Redemption status

### 7.3 Data Fetching
- [ ] Create `useUserStats` hook:
  ```typescript
  // Aggregate user statistics
  // - Query all user bets from DB
  // - Fetch on-chain positions
  // - Calculate P&L
  ```

- [ ] Create `useUserPositions` hook:
  ```typescript
  // Fetch user's outcome token balances
  // - Loop through markets
  // - Get YES/NO token balances
  // - Calculate current value and P&L
  ```

---

## Phase 8: Market Discovery & Filtering

### 8.1 Market Sidebar
Location: `client/src/components/MarketSidebar.tsx` (create new)

#### Features:
- [ ] Category navigation:
  ```typescript
  Categories:
  - All Markets
  - 🏀 Sports
    - Basketball
    - Football
    - Soccer
    - Tennis
    - Combat Sports
    - Other Sports
  - 💰 Crypto
  - 🏛️ Politics
  - 🌤️ Weather
  - 🎬 Entertainment
  - 📊 Economics
  - 🔬 Science & Tech
  ```

- [ ] Market count badges (e.g., "Sports (24)")
- [ ] Active/Closed/Resolved filters
- [ ] Collapsible sections
- [ ] Sticky sidebar (scrolls with page)

### 8.2 Market Filtering
- [ ] Create `useMarketFilters` hook:
  ```typescript
  interface MarketFilters {
    category?: string;
    status?: 'Active' | 'Closed' | 'Resolved';
    search?: string;
    sortBy?: 'volume' | 'closing' | 'created' | 'odds';
  }
  ```

- [ ] Apply filters to market list:
  - Client-side filtering for fast UX
  - Backend filtering for large datasets

- [ ] Search functionality:
  - Search by market question
  - Search by category
  - Fuzzy matching

### 8.3 Category Pages
- [ ] Create category-specific pages (e.g., `/markets/sports`)
- [ ] Show category description and image
- [ ] Filter markets by category
- [ ] Show popular markets in category

---

## Phase 9: Leaderboard

### 9.1 Leaderboard Page
Location: `client/src/pages/leaderboard.tsx` (update existing)

#### Features:
- [ ] **Top Winners Table**:
  - Rank (#1, #2, #3, ...)
  - User wallet address (shortened)
  - Display name (if set)
  - Total winnings (in USDC)
  - Total wagered
  - ROI percentage
  - Win/loss record
  - Number of markets participated

- [ ] **Leaderboard Tabs**:
  - All-Time
  - This Month
  - This Week
  - Today

- [ ] **Ranking Criteria**:
  - Primary: Total profit (winnings - wagered)
  - Secondary: ROI percentage
  - Tertiary: Number of wins

- [ ] **User Highlighting**:
  - Highlight current user's row
  - Show user's rank if not in top 10

### 9.2 Leaderboard Data
- [ ] Create `server/routes.ts` endpoint:
  ```typescript
  GET /api/leaderboard?period=all|month|week|day&limit=100
  
  // Query resolved bets from database
  // Calculate profit/loss per user
  // Sort by total profit
  // Return top N users
  ```

- [ ] Create `useLeaderboard` hook:
  ```typescript
  // Fetch and cache leaderboard data
  // Refresh every 5 minutes
  // Support pagination
  ```

### 9.3 Profile Integration
- [ ] Add "View on Leaderboard" link from profile
- [ ] Show user's current rank in profile header
- [ ] Badge system (optional):
  - 🥇 Gold: Top 10
  - 🥈 Silver: Top 50
  - 🥉 Bronze: Top 100

---

## Phase 10: Resolution System (Admin)

### 10.1 Market Resolution Interface
Location: `client/src/pages/admin/resolution.tsx` (or add to markets page)

#### Features:
- [ ] **Markets Awaiting Resolution**:
  - List of closed markets
  - Market details
  - Current pool totals
  - Resolution source

- [ ] **Manual Resolution**:
  - Select outcome (YES/NO or specific option)
  - Provide evidence/reasoning
  - Upload evidence to IPFS
  - Sign EIP-712 message
  - Submit resolution proposal on-chain
  - Bond HORIZON tokens (using `useProposeResolution`)

- [ ] **Oracle Resolution** (AI-powered):
  - Request resolution from AI Oracle
  - Show oracle request status
  - Display confidence score and reasoning
  - Review and submit oracle result
  - Automatic EIP-712 signing by oracle

- [ ] **TheSportsDB Resolution** (Sports markets):
  - Fetch game results from API
  - Parse winner/loser
  - Auto-propose resolution

### 10.2 Resolution Flow

**Manual Resolution:**
1. Admin selects closed market
2. Chooses outcome (YES/NO)
3. Uploads evidence (news articles, APIs, etc.)
4. Signs EIP-712 message using SDK `signProposal`
5. Calls `useProposeResolution` with signature
6. Transaction submitted, resolution enters dispute period
7. After dispute period (e.g., 24 hours), finalize with `useFinalize`

**Oracle Resolution:**
1. Admin clicks "Request Oracle Resolution"
2. SDK `useRequestResolution` sends market data to AI
3. Poll oracle status with `useOracleStatus`
4. When complete, fetch result with `useOracleResult`
5. Review AI's decision (outcome + confidence + reasoning)
6. Approve or reject AI's proposal
7. If approved, submit with `useProposeResolution`
8. Oracle signature already included, no manual signing

### 10.3 Resolution Panel Component
Location: `client/src/components/ResolutionPanel.tsx` (already exists, update)

- [ ] Show market details
- [ ] Resolution options (Manual, Oracle, TheSportsDB)
- [ ] Evidence upload form
- [ ] EIP-712 signature preview
- [ ] Transaction status
- [ ] Dispute period countdown

### 10.4 Finalization
- [ ] After dispute period ends:
  - Admin (or anyone) calls `useFinalize`
  - Market status updated to Resolved
  - Winners can redeem with `useRedeem`
  - Creator stake refunded

### 10.5 Database Sync
- [ ] Listen for on-chain resolution events
- [ ] Update market status in database
- [ ] Update bet statuses (won/lost)
- [ ] Calculate payouts
- [ ] Store resolution evidence

---

## Phase 11: Testing & Polish

### 11.1 Unit Testing
- [ ] Test SDK hooks with mock data
- [ ] Test authentication flow
- [ ] Test admin access control
- [ ] Test market creation
- [ ] Test trading calculations

### 11.2 Integration Testing
- [ ] Test end-to-end market creation (testnet)
- [ ] Test trading flow (buy/sell)
- [ ] Test resolution flow
- [ ] Test redemption
- [ ] Test multi-user scenarios

### 11.3 Testnet Deployment
- [ ] Deploy to BNB Chain Testnet (chain ID 97)
- [ ] Create test markets
- [ ] Invite users to test
- [ ] Monitor for bugs

### 11.4 UI/UX Polish
- [ ] Loading states for all async operations
- [ ] Error handling and user-friendly messages
- [ ] Responsive design (mobile, tablet, desktop)
- [ ] Dark mode consistency
- [ ] Accessibility (ARIA labels, keyboard navigation)
- [ ] Transaction confirmations
- [ ] Success/error toasts

### 11.5 Performance Optimization
- [ ] Lazy load pages and components
- [ ] Optimize database queries
- [ ] Cache market data (React Query)
- [ ] Minimize bundle size
- [ ] Image optimization

### 11.6 Security Audit
- [ ] Review admin access controls
- [ ] Validate all user inputs
- [ ] Sanitize database queries
- [ ] Check for XSS vulnerabilities
- [ ] Rate limiting on API endpoints
- [ ] CSRF protection

---

## Phase 12: Production Deployment

### 12.1 Environment Setup
- [ ] Set up production environment variables
- [ ] Configure production RPC endpoints
- [ ] Set up production database (SQLite or PostgreSQL)
- [ ] Configure domain and SSL certificate

### 12.2 Smart Contract Verification
- [ ] Verify all contracts on BSCScan
- [ ] Test all admin functions on mainnet (with small amounts)
- [ ] Verify contract addresses match environment variables

### 12.3 Launch Checklist
- [ ] Database backups configured
- [ ] Monitoring and logging set up
- [ ] Error tracking (e.g., Sentry)
- [ ] Analytics (e.g., Google Analytics, Mixpanel)
- [ ] Documentation updated
- [ ] User guide created
- [ ] Terms of service and privacy policy
- [ ] Contact/support channel

### 12.4 Mainnet Deployment
- [ ] Update .env to use mainnet (chain ID 56)
- [ ] Update contract addresses to mainnet
- [ ] Deploy frontend to hosting (Vercel, Netlify, etc.)
- [ ] Deploy backend (if separate)
- [ ] Run smoke tests on production

### 12.5 Post-Launch
- [ ] Monitor user activity
- [ ] Track transaction success/failure rates
- [ ] Monitor gas usage
- [ ] Collect user feedback
- [ ] Plan feature iterations

---

## Key Features Summary

### For Admin (`0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`)
✅ Create markets on-chain via MarketFactory  
✅ Add initial liquidity to markets  
✅ Resolve market outcomes (manual or oracle)  
✅ Manage market categories  
✅ View all markets and positions  
✅ Invalidate markets if needed  

### For Users
✅ Connect wallet via SIWE (Sign-In With Ethereum)  
✅ Browse markets by category  
✅ Filter and search markets  
✅ View real-time odds and prices  
✅ Buy outcome tokens (YES/NO)  
✅ Sell outcome tokens  
✅ Track positions and P&L  
✅ View winnings dashboard  
✅ View wager history  
✅ Redeem winning tokens  
✅ View leaderboard  
✅ View profile statistics  

### Technical Features
✅ Full blockchain integration via SDK  
✅ BNB Chain (BSC) support  
✅ ERC-20 collateral (USDC/BUSD)  
✅ Outcome tokens (ERC-1155)  
✅ AI Oracle integration  
✅ IPFS metadata storage  
✅ EIP-712 signatures  
✅ SQLite database for caching  
✅ Real-time updates  
✅ Transaction tracking  

---

## Technology Stack

### Frontend
- **React 18** with TypeScript
- **Vite** for building
- **Wagmi 2.x** for Ethereum interactions
- **RainbowKit** for wallet connection
- **@project-gamma/sdk** for market operations
- **TanStack Query** for data fetching
- **Wouter** for routing
- **Tailwind CSS** for styling
- **Shadcn/ui** for components

### Backend
- **Express.js** with TypeScript
- **SQLite** for local database
- **Drizzle ORM** for database management
- **iron-session** for SIWE sessions
- **siwe** for signature verification

### Blockchain
- **BNB Chain (BSC)** - Mainnet & Testnet
- **MarketFactory** contract
- **BinaryMarket** contracts (primary)
- **OutcomeToken** (ERC-1155)
- **HORIZON** governance token
- **AI Oracle** for resolution

### DevOps
- **Git** for version control
- **npm/pnpm** for package management
- **Environment variables** for configuration
- **Vercel/Netlify** for frontend hosting
- **Railway/Render** for backend hosting (if needed)

---

## Database Schema Changes

### Updated `users` Table
```typescript
export const users = sqliteTable("users", {
  id: text("id").primaryKey(),
  walletAddress: text("wallet_address").notNull().unique(),
  displayName: text("display_name"),
  profileImageUrl: text("profile_image_url"),
  role: text("role").notNull().default("user"),
  createdAt: integer("created_at", { mode: 'timestamp' }),
  updatedAt: integer("updated_at", { mode: 'timestamp' }),
});
```

### `markets` Table (Existing)
Already has blockchain fields:
- `marketAddress`
- `chainId`
- `onChainMarketId`
- `yesTokenId`
- `noTokenId`
- `resolutionSource`
- `oracleRequestId`

### `bets` Table (Existing)
Tracks user positions:
- `userId`
- `marketId`
- `outcome` (A or B)
- `amount`
- `status` (active, won, lost, voided)

---

## Contract Addresses (BNB Chain Mainnet)

```
MarketFactory:      0x22Cc806047BB825aa26b766Af737E92B1866E8A6
HorizonToken:       0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444
OutcomeToken:       0x17B322784265c105a94e4c3d00aF1E5f46a5F311
HorizonPerks:       0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C
FeeSplitter:        0x275017E98adF33051BbF477fe1DD197F681d4eF1
ResolutionModule:   0xF0CF4C741910cB48AC596F620a0AE892Cd247838
AIOracleAdapter:    0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93
```

**Admin Address:** `0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444`

---

## Estimated Timeline

**Phase 1-3** (Setup, Auth, Admin): 3-5 days  
**Phase 4** (SDK Integration): 1-2 days  
**Phase 5** (Admin Market Management): 2-3 days  
**Phase 6** (Trading Interface): 3-4 days  
**Phase 7** (Dashboard & Tracking): 2-3 days  
**Phase 8** (Market Discovery): 2 days  
**Phase 9** (Leaderboard): 1-2 days  
**Phase 10** (Resolution System): 2-3 days  
**Phase 11** (Testing & Polish): 3-5 days  
**Phase 12** (Production Deployment): 1-2 days  

**Total Estimated Time:** 20-30 days for a single developer  
**With 2-3 developers:** 10-15 days  

---

## Success Criteria

### MVP Launch Checklist
- [ ] Admin can create markets on-chain
- [ ] Admin can resolve markets
- [ ] Users can connect wallet via SIWE
- [ ] Users can buy/sell outcome tokens
- [ ] Users can view their positions and P&L
- [ ] Users can browse markets by category
- [ ] Leaderboard shows top winners
- [ ] All data syncs between blockchain and database
- [ ] Mobile responsive
- [ ] No critical bugs

### Post-Launch Goals
- 100+ active markets
- 1,000+ registered users
- $100,000+ total volume
- <1% transaction failure rate
- <2s average page load time
- 95%+ uptime

---

## Support & Resources

**Project Gamma SDK Docs:** See `/sdk/docs/`  
**Smart Contracts:** See `/contracts/src/`  
**BNB Chain Docs:** https://docs.bnbchain.org/  
**Wagmi Docs:** https://wagmi.sh/  
**RainbowKit Docs:** https://www.rainbowkit.com/  

---

## Next Steps

1. Review this plan with the team
2. Set up development environment
3. Begin Phase 1: Infrastructure setup
4. Start building! 🚀
