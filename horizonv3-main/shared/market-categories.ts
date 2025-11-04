export interface League {
  id: string;
  name: string;
  displayName: string;
  badge?: string;
}

export interface Sport {
  id: string;
  name: string;
  iconName: string; // For lucide icons, SVG filename, or PNG filename
  iconType?: 'lucide' | 'svg' | 'custom'; // Determines how to render the icon (custom = PNG)
  leagues: League[];
}

export const sportsData: Sport[] = [
  {
    id: "real-world-events",
    name: "Real World Events",
    iconName: "MapPin",
    iconType: "lucide",
    leagues: [
      { id: "politics", name: "Politics", displayName: "Politics" },
      { id: "elections", name: "Elections", displayName: "Elections" },
      { id: "awards", name: "Awards & Entertainment", displayName: "Awards" },
      { id: "weather", name: "Weather & Climate", displayName: "Weather" },
      { id: "technology", name: "Technology", displayName: "Technology" },
      { id: "space", name: "Space & Science", displayName: "Space" },
    ],
  },
  {
    id: "crypto",
    name: "Crypto",
    iconName: "Coins",
    iconType: "lucide",
    leagues: [
      { id: "bitcoin", name: "Bitcoin", displayName: "Bitcoin" },
      { id: "ethereum", name: "Ethereum", displayName: "Ethereum" },
      { id: "altcoins", name: "Altcoins", displayName: "Altcoins" },
      { id: "defi", name: "DeFi", displayName: "DeFi" },
      { id: "nft", name: "NFTs", displayName: "NFTs" },
      { id: "crypto-events", name: "Crypto Events", displayName: "Crypto Events" },
    ],
  },
  {
    id: "finance",
    name: "Finance",
    iconName: "TrendingUp",
    iconType: "lucide",
    leagues: [
      { id: "stocks", name: "Stocks", displayName: "Stocks" },
      { id: "commodities", name: "Commodities", displayName: "Commodities" },
      { id: "forex", name: "Forex", displayName: "Forex" },
      { id: "indices", name: "Market Indices", displayName: "Indices" },
      { id: "economics", name: "Economic Indicators", displayName: "Economics" },
      { id: "corporate", name: "Corporate Events", displayName: "Corporate" },
    ],
  },
];
