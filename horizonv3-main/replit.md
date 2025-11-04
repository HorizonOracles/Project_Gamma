# DegenArena - Blockchain Sports Betting Platform

## Overview
DegenArena is a pool-based sports betting platform built on Binance Smart Chain, utilizing a parimutuel betting system where users bet against each other. This creates dynamic odds based on real-time activity. The platform features managed wallets for accessibility, social betting feeds, and supports a wide range of sports including NBA, NFL, Soccer, Baseball, and various esports. The business vision is to provide a decentralized, engaging, and transparent sports betting experience.

## User Preferences
Preferred communication style: Simple, everyday language.

## System Architecture

### Frontend
- **Technology Stack**: React with TypeScript, Vite, Wouter for routing, TanStack Query for state management.
- **Design System**: Shadcn/ui with Radix UI primitives, Tailwind CSS (dark mode with yellow/gold accents), custom typography (Inter, JetBrains Mono). Features a three-column dashboard layout.
- **Button Themes**: Two primary animated gradient themes defined in `index.css`:
  - **darkify**: Dark gray animated gradient background (#292929 → #545454) with 4s animation cycle. Used for admin buttons, active filter buttons, and Pool B betting buttons. Apply with `className="darkify text-white border-2"` and `style={{ borderColor: '#424242' }}`.
  - **goldify**: Gold animated gradient background (#d5b877 → #f0e6c8) with 4s animation cycle. Used for Pool A betting buttons and premium features. Apply with `className="goldify text-primary-foreground border-2"` and `style={{ borderColor: '#d5b877' }}`.
  - Both themes include smooth background position animations and can be combined with NetworkBackground component for enhanced visual effects.
- **Component Architecture**: Modular and reusable components, with feature-specific components like `MarketCard` and `BetSlipDialog`.
- **Real-time Features**: 
  - WebSocket connection established on all authenticated pages for live market updates
  - Server broadcasts market updates every 5 seconds and immediately after bet placements or market status changes
  - Client-side query invalidation with `staleTime: 0` ensures instant refetching when WebSocket broadcasts updates
  - Users see odds, pool totals, and bet feed update in real-time without page refresh
  - Admin page polls every 5 seconds with `refetchInterval: 5000` for reliable settlement interface
- **UI/UX Decisions**: Responsive grid layout and dedicated sport category pages with consistent layouts.
- **User Features**: My Bets page for comprehensive betting history and display name system with cooldowns.
- **Leaderboard**: Players are ranked by profit and loss (total won minus total wagered). Top performers with positive P&L are highlighted in green, while losses are shown in red.
- **Privacy Protection**: Complete email privacy - no user emails are displayed anywhere on the platform. All user identification uses display names with "Anonymous Player" fallback. Applied across winning bet details, leaderboard, activity feed, and chat panel.
- **Individual Sports (Players vs Teams)**: The platform distinguishes between team sports and individual sports. Individual sports display "Players" instead of "Teams" and include: Tennis (ATP/WTA), Table Tennis, Darts, Motorsport, Fighting (MMA/Boxing), and Golf. Team sports include Basketball, American Football, Soccer, Baseball, etc.
- **Tennis Player Database**: Expanded ATP player database with 116 players total. Player headshots stored locally in `server/public/` directory using `firstname-lastname.png` naming convention. 114 player headshot images currently available.
- **Sport Icon Template Pattern**: ALL sports market pages must display the sport icon before the page title. Icons are loaded from `/sport-icons/` directory (SVG for traditional sports, PNG for esports). Use `sportConfig?.iconName` to reference the icon file, with 10x10 (h-10 w-10) sizing with `object-contain` in a flex container with gap-4 alignment. All icons (both SVG and PNG) must use `brightness-0 invert` classes to display as white. This pattern applies to both league selection pages and individual league match pages.
- **Sports Dropdown Navigation**: Basketball, American Football, Baseball, and Tennis feature collapsible dropdown menus in the sidebar showing top leagues with individual market counts. Dropdowns display total sport market count, nested leagues with league badges (h-5 w-5), and dynamic sorting by market count (descending). League logos maintain original colors on hover without white filter effect. Each league navigates directly to `/sports/{sport}/league/{leagueName}` pages.
- **Admin Settlement Page**: Compact card layout with search functionality and icon-based sport filtering. Markets display in dense rows showing badges, logos, pool data, and action buttons. Space-efficient design allows viewing 5+ markets simultaneously. Sport filter buttons use actual sport icon images (same as used throughout site) in square icon-only buttons with darkify theme when active.
- **Auto-Hide Inactive Leagues**: Leagues without events scheduled within the next 2 weeks are automatically hidden from the sidebar to reduce clutter. This behavior can be overridden by admins through manual visibility controls. The system checks for upcoming events via `/api/leagues/activity-status` endpoint (refreshes every minute). Admin panel shows visual badges indicating whether leagues are manually controlled, auto-hidden, or active with upcoming events.

### Backend
- **Server Framework**: Express.js with TypeScript, ESM modules, custom logging.
- **API Design**: RESTful, authenticated endpoints, standardized error handling, JSON format.
- **Authentication & Authorization**: OpenID Connect with Replit Auth, Passport.js, session storage in PostgreSQL, protected routes.
- **Business Logic**: Pool-based betting with dynamic odds calculation, automatic wallet creation, transaction tracking, market status management, and bet settlement with payout calculation.
- **Real-time Features**: WebSocket server attached to HTTP server for live updates.

### Data Storage
- **Database**: PostgreSQL (Neon serverless) with Drizzle ORM.
- **Schema Design**:
    - `sessions`: Replit Auth session data.
    - `users`: User profiles and display names.
    - `wallets`: User wallets with BNB addresses and balances.
    - `transactions`: Audit trail of all financial operations.
    - `markets`: Sports betting markets with pool details.
    - `bets`: Individual bets linked to users and markets.
- **Data Relationships**: One-to-one (User → Wallet), One-to-many (User → Bets, User → Transactions, Market → Bets).
- **Query Patterns**: Optimized for dashboard views, transaction history, and aggregated bet displays.

## External Dependencies

- **Development & Build Tools**: Vite, TypeScript, ESBuild, Drizzle Kit.
- **Third-Party Services**: Replit Auth, Neon PostgreSQL, Binance Smart Chain (future on-chain wallet management).
- **UI Component Libraries**: Radix UI, Shadcn/ui, Lucide React, React Icons.
- **Utility Libraries**: date-fns, Zod, class-variance-authority, clsx, tailwind-merge, memoizee.
- **WebSocket**: `ws` library for server implementation.
- **Session Management**: `express-session`, `connect-pg-simple`.