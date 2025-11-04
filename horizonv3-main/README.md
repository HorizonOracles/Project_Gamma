# DegenArena (Horizon Market) 🎯

A decentralized sports betting platform built on Binance Smart Chain, featuring pool-based parimutuel betting across 497 leagues and 20+ sports.

## 🌟 Features

- **Pool-Based Betting System**: Parimutuel betting where users bet against each other with dynamic odds
- **497+ Leagues**: Comprehensive coverage across 20+ sports including NBA, NFL, Soccer, Tennis, MMA, and esports
- **Real-Time Updates**: WebSocket-powered live odds and market updates
- **Managed Wallets**: Built-in wallet system for easy accessibility
- **Social Betting Feed**: Live activity feed showing all betting action
- **Individual & Team Sports**: Support for both team sports and individual competitions (Tennis, Golf, MMA, etc.)
- **Auto-Hide Inactive Leagues**: Smart league visibility based on upcoming events (2-week window)
- **Admin Controls**: Comprehensive admin panel for market management and visibility settings
- **Leaderboard System**: Player rankings by profit/loss with visual indicators
- **Complete Privacy**: No user emails displayed anywhere on the platform

## 🏗️ Tech Stack

### Frontend
- **React** with TypeScript
- **Vite** for blazing-fast development
- **Wouter** for routing
- **TanStack Query** for state management
- **Shadcn/ui** + Radix UI for component library
- **Tailwind CSS** for styling with custom dark theme
- **WebSocket** for real-time updates

### Backend
- **Express.js** with TypeScript
- **PostgreSQL** (Neon serverless) with Drizzle ORM
- **Passport.js** + OpenID Connect for authentication
- **WebSocket Server** for live market updates
- **TheSportsDB API** for comprehensive sports data

### Blockchain
- **Binance Smart Chain** integration (planned)
- Managed wallets for user accessibility

## 🚀 Getting Started

### Prerequisites
- Node.js 20+
- PostgreSQL database
- TheSportsDB API key (Premium)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/malicexmaple/horizonmarket.git
cd horizonmarket
```

2. Install dependencies:
```bash
npm install
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env and fill in your values
```

4. Push database schema:
```bash
npm run db:push
```

5. Start the development server:
```bash
npm run dev
```

The application will be available at `http://localhost:5000`

## 📁 Project Structure

```
├── client/              # Frontend React application
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── pages/       # Page components
│   │   ├── hooks/       # Custom React hooks
│   │   └── lib/         # Utility functions
├── server/              # Backend Express server
│   ├── routes.ts        # API routes
│   ├── storage.ts       # Data access layer
│   ├── thesportsdb.ts   # TheSportsDB API integration
│   └── index.ts         # Server entry point
├── shared/              # Shared types and schemas
│   ├── schema.ts        # Drizzle database schema
│   └── sports-leagues.ts # Sports and league configuration
└── attached_assets/     # Static assets
```

## 🎨 Design System

- **Dark Theme**: Black backgrounds with yellow/gold accents
- **Typography**: Inter, JetBrains Mono
- **Button Themes**:
  - **darkify**: Dark gray animated gradient (admin, filters, Pool B)
  - **goldify**: Gold animated gradient (Pool A, premium features)
- **Three-Column Dashboard**: Optimal betting experience layout

## 🔧 Key Features Explained

### Pool-Based Betting
Users bet into pools (Pool A vs Pool B) with dynamic odds calculated based on pool ratios. This creates a parimutuel system where bettors compete against each other rather than the house.

### Real-Time Updates
- WebSocket broadcasts every 5 seconds
- Immediate updates after bet placements
- Live odds calculation
- Activity feed updates

### League Visibility
- Automatic hiding of leagues with no events in the next 2 weeks
- Admin manual override capability
- Visual badges showing league status (Manual, Auto-hidden, Active)

### Privacy Protection
- Complete email privacy throughout the platform
- Display name system with "Anonymous Player" fallback
- Applied across betting feed, leaderboard, and winning details

## 🔐 Environment Variables

See `.env.example` for required environment variables:
- `DATABASE_URL`: PostgreSQL connection string
- `SESSION_SECRET`: Random string for session encryption
- `THESPORTSDB_API_KEY`: Premium API key for sports data
- `NODE_ENV`: development or production

## 📊 Database Schema

Key tables:
- `users`: User profiles and display names
- `wallets`: User wallets with BNB addresses
- `markets`: Sports betting markets with pool details
- `bets`: Individual bets linked to users and markets
- `transactions`: Financial operation audit trail
- `visibility_settings`: League/sport visibility controls

## 🛠️ Development

```bash
# Start development server
npm run dev

# Push database schema changes
npm run db:push

# Force push schema (data loss warning)
npm run db:push -- --force
```

## 📈 Future Enhancements

- On-chain wallet integration with Binance Smart Chain
- Live betting during events
- Mobile app (React Native)
- Additional sports and leagues
- Advanced analytics and statistics
- Multi-language support

## 🤝 Contributing

This is a private project, but suggestions and feedback are welcome!

## 📄 License

All rights reserved.

## 🙏 Acknowledgments

- TheSportsDB for comprehensive sports data
- Shadcn/ui for beautiful component library
- Replit for development environment
- Neon for serverless PostgreSQL

---

Built with ❤️ for the degen community
