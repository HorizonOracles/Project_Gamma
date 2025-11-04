# Horizon V3 - SDK Integration Plan

## Overview
Integrate the `@project-gamma/react-sdk` into the Horizon V3 frontend to enable real prediction market features with wallet connect, on-chain trading, and admin controls.

---

## Phase 1: Database & Backend Setup

### Task 1.1: Add Admin Whitelist Table ✅
**File**: `shared/schema.ts`
- Add `adminWhitelist` table with:
  - `id` (serial primary key)
  - `walletAddress` (varchar, unique, not null) - Ethereum address
  - `addedBy` (varchar, references users.id)
  - `createdAt` (timestamp)
- Export types: `AdminWhitelist`, `InsertAdminWhitelist`

### Task 1.2: Update Markets Schema to Include Chain Data ✅
**File**: `shared/schema.ts`
- Add to `markets` table:
  - `marketId` (bigint, nullable) - On-chain market ID from MarketFactory
  - `marketAddress` (varchar, nullable) - Contract address
  - `chainId` (integer, default 97) - BNB testnet = 97, mainnet = 56
  - `collateralToken` (varchar) - ERC20 token address
  - `yesTokenId` (bigint)
  - `noTokenId` (bigint)

### Task 1.3: Add Server API Routes ✅
**New File**: `server/routes/admin.ts`
- `POST /api/admin/check-whitelist` - Check if address is whitelisted
- `POST /api/admin/markets/create` - Create on-chain market (admin only)
- `POST /api/admin/markets/:id/add-liquidity` - Add liquidity (admin only)
- `POST /api/admin/markets/:id/resolve` - Resolve market (admin only)
- `POST /api/admin/oracle/request-resolution` - Request AI Oracle resolution (admin only)

**New File**: `server/routes/markets.ts`
- `GET /api/markets` - List all markets
- `GET /api/markets/:id` - Get market details
- `GET /api/markets/:id/prices` - Get current prices
- `POST /api/markets/:id/buy` - Buy YES/NO shares (user)
- `POST /api/markets/:id/sell` - Sell YES/NO shares (user)

**New File**: `server/routes/user.ts`
- `GET /api/user/portfolio` - Get user's positions across markets
- `GET /api/user/stats` - Get user wagered/earned stats
- `GET /api/leaderboard` - Get top users by rank points

### Task 1.4: Add Middleware for Wallet Auth ✅
**New File**: `server/middleware/walletAuth.ts`
- Extract wallet address from signed message
- Verify signature using viem
- Attach wallet address to request
- Check admin whitelist for protected routes

### Task 1.5: Integrate Routes into Server ✅
**File**: `server/index.ts`
- Import and mount admin, markets, user routes
- Add wallet auth middleware

---

## Phase 2: Frontend SDK Integration

### Task 2.1: Install SDK Package ✅
**Action**: Link local SDK to frontend
```bash
cd sdk
npm link
cd ../horizonv3-main
npm link @project-gamma/react-sdk
```

### Task 2.2: Setup Wagmi & RainbowKit Config ✅
**New File**: `client/src/config/wagmi.ts`
- Configure chains (BNB testnet, mainnet)
- Setup WalletConnect project ID
- Configure RainbowKit with custom theme
- Export wagmi config

### Task 2.3: Add SDK Providers to App ✅
**File**: `client/src/main.tsx`
```tsx
import { WagmiProvider } from 'wagmi'
import { QueryClientProvider } from '@tanstack/react-query'
import { RainbowKitProvider } from '@rainbow-me/rainbowkit'
import { GammaProvider } from '@project-gamma/react-sdk'
import { wagmiConfig, queryClient } from './config/wagmi'
import '@rainbow-me/rainbowkit/styles.css'

// Wrap App with providers
<WagmiProvider config={wagmiConfig}>
  <QueryClientProvider client={queryClient}>
    <RainbowKitProvider>
      <GammaProvider
        chainId={97}
        marketFactoryAddress="0x..."
      >
        <App />
      </GammaProvider>
    </RainbowKitProvider>
  </QueryClientProvider>
</WagmiProvider>
```

### Task 2.4: Add Wallet Connect Button Component ✅
**New File**: `client/src/components/WalletConnect.tsx`
- Use `ConnectButton` from RainbowKit
- Custom styling to match app theme
- Display user address, balance, avatar
- Dropdown with disconnect option

---

## Phase 3: Admin Features

### Task 3.1: Update Admin Dashboard ✅
**File**: `client/src/pages/admin.tsx`
- Add wallet connect requirement check
- Check admin whitelist via API
- Show admin actions: Create Market, Manage Markets, Oracle
- Link to admin-markets page

### Task 3.2: Create Market Page ✅
**File**: `client/src/pages/admin-markets.tsx`
- **Create Market Section**:
  - Form: question, description, endTime, category
  - Use `useCreateMarket()` hook from SDK
  - Call `/api/admin/markets/create` to save to DB
  - Show transaction status
- **Active Markets List**:
  - Fetch from `/api/markets`
  - Display: question, status, liquidity, volume
  - Actions per market:
    - Add Liquidity button → opens modal
    - Resolve button → opens resolve modal

### Task 3.3: Add Liquidity Modal Component ✅
**New File**: `client/src/components/admin/AddLiquidityModal.tsx`
- Input: initial YES amount, initial NO amount
- Use `useAddLiquidity()` hook from SDK
- Show transaction status
- Refresh market data on success

### Task 3.4: Resolve Market Modal Component ✅
**New File**: `client/src/components/admin/ResolveMarketModal.tsx`
- Radio buttons: YES wins, NO wins, Invalid
- Use `useFinalize()` hook from SDK
- Call `/api/admin/markets/:id/resolve` to update DB
- Show transaction status

### Task 3.5: Oracle Resolution Page ✅
**New File**: `client/src/pages/oracle-resolution.tsx` (admin only)
- List markets ready for resolution
- Button: "Request AI Oracle Resolution"
- Use `useRequestResolution()` hook from SDK
- Display oracle response
- Option to dispute or finalize

---

## Phase 4: User Features

### Task 4.1: Markets List Page ✅
**File**: `client/src/pages/home.tsx` or new `markets.tsx`
- Fetch markets from `/api/markets`
- Display cards with:
  - Question, description
  - Current YES/NO prices (use `usePrices()` hook)
  - Total liquidity, volume
  - Time remaining
- Click → navigate to market detail page

### Task 4.2: Market Detail Page ✅
**New File**: `client/src/pages/market-detail.tsx`
- URL: `/market/:id`
- Display full market info
- **Buy Section**:
  - Tabs: Buy YES | Buy NO
  - Input: amount in collateral token
  - Show estimated shares received (use `useQuote()` hook)
  - Button: "Buy YES" / "Buy NO"
  - Use `useBuy()` hook from SDK
- **Sell Section**:
  - Display user's YES/NO share balances
  - Input: shares to sell
  - Show estimated collateral received
  - Use `useSell()` hook from SDK
- **Market Stats**:
  - Total volume, liquidity, participants
  - Price chart (optional)

### Task 4.3: Portfolio Page ✅
**File**: `client/src/pages/profile.tsx` or new `portfolio.tsx`
- Fetch user positions from `/api/user/portfolio`
- Display:
  - Active positions (market, outcome, shares, current value)
  - Past positions (settled)
- **Stats Section**:
  - Total Wagered (from DB)
  - Total Earned (from DB)
  - Win Rate
  - Rank (Bronze, Silver, Gold, etc.)

### Task 4.4: Leaderboard Page ✅
**File**: `client/src/pages/leaderboard.tsx`
- Fetch from `/api/leaderboard`
- Display table:
  - Rank, Username, Total Wagered, Total Earned, Win Rate
  - Highlight current user
- Filter options: All time, This month, This week

### Task 4.5: Update Profile Page ✅
**File**: `client/src/pages/profile.tsx`
- Add wallet connection status
- Display wallet address
- Show user stats: wagered, earned, rank
- Link to portfolio

---

## Phase 5: Token Approvals & Balance Checks

### Task 5.1: Token Approval Component ✅
**New File**: `client/src/components/TokenApproval.tsx`
- Check if user has approved collateral token for market contract
- Use `useApprove()` hook from SDK
- Button: "Approve Token" (if not approved)
- Show loading state during approval

### Task 5.2: Balance Checks ✅
- Use `useBalance()` hook from SDK
- Display user's collateral token balance
- Warn if insufficient balance before buy/addLiquidity
- Suggest depositing more tokens

---

## Phase 6: Error Handling & UX

### Task 6.1: Transaction Status Toast ✅
**File**: `client/src/hooks/useTransactionToast.ts`
- Custom hook to show toast notifications
- States: pending, success, error
- Link to block explorer for completed transactions

### Task 6.2: Loading States ✅
- Add skeleton loaders for:
  - Market list
  - Market detail
  - Portfolio
  - Leaderboard
- Use `isLoading` from SDK hooks

### Task 6.3: Error Boundaries ✅
- Wrap pages in error boundary
- Display user-friendly error messages
- Log errors to console for debugging

---

## Phase 7: Testing & Cleanup

### Task 7.1: Manual Testing ✅
- Test admin flow: create market, add liquidity, resolve
- Test user flow: connect wallet, buy shares, sell shares
- Test edge cases: insufficient balance, expired markets, etc.
- Test on BNB testnet

### Task 7.2: Remove Old Sports Betting Code ✅
- Delete unused sports league pages (keep as reference if needed)
- Remove hardcoded market data
- Clean up unused components

### Task 7.3: Update README ✅
**File**: `horizonv3-main/README.md`
- Document SDK integration
- List required environment variables
- Add setup instructions
- Add deployment notes

---

## Environment Variables Required

### Backend (`horizonv3-main/.env`)
```env
DATABASE_URL=postgresql://...
NODE_ENV=development
PORT=5000
SESSION_SECRET=...
CHAIN_RPC_URL=https://data-seed-prebsc-1-s1.binance.org:8545/
MARKET_FACTORY_ADDRESS=0x...
```

### Frontend (if needed)
```env
VITE_WALLETCONNECT_PROJECT_ID=...
VITE_MARKET_FACTORY_ADDRESS=0x...
VITE_CHAIN_ID=97
```

---

## Deployment Checklist

- [ ] Deploy contracts to BNB testnet
- [ ] Update `MARKET_FACTORY_ADDRESS` in env
- [ ] Run DB migrations (`npm run db:push`)
- [ ] Add initial admin addresses to whitelist table
- [ ] Test admin features on testnet
- [ ] Test user features on testnet
- [ ] Deploy to production (BNB mainnet)
- [ ] Monitor transactions and error logs

---

## Progress Tracking

### Completed Tasks
- [ ] Phase 1: Database & Backend Setup (0/5)
- [ ] Phase 2: Frontend SDK Integration (0/4)
- [ ] Phase 3: Admin Features (0/5)
- [ ] Phase 4: User Features (0/5)
- [ ] Phase 5: Token Approvals & Balance (0/2)
- [ ] Phase 6: Error Handling & UX (0/3)
- [ ] Phase 7: Testing & Cleanup (0/3)

**Total Progress: 0/27 tasks completed**

---

## Notes
- Keep existing sports data tables for future expansion
- Admin whitelist is a simple security measure; consider more robust auth later
- All on-chain actions require wallet connection
- DB tracks markets for caching and analytics; blockchain is source of truth
- Consider adding WebSocket for live price updates later
