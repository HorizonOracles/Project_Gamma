// Prediction Markets Categories and Configuration
// Non-sports betting markets for real-world events

export interface PredictionCategory {
  id: string;
  name: string;
  description: string;
  iconName: string; // Lucide icon name
  color: string; // Tailwind color class
}

export const predictionCategories: PredictionCategory[] = [
  {
    id: 'finance',
    name: 'Finance',
    description: 'Stock markets, IPOs, and financial events',
    iconName: 'DollarSign',
    color: 'text-green-500',
  },
  {
    id: 'crypto',
    name: 'Crypto',
    description: 'Cryptocurrency prices and blockchain events',
    iconName: 'Bitcoin',
    color: 'text-orange-500',
  },
  {
    id: 'geopolitics',
    name: 'Geopolitics',
    description: 'International relations and political events',
    iconName: 'Globe',
    color: 'text-blue-500',
  },
  {
    id: 'earnings',
    name: 'Earnings',
    description: 'Company earnings reports and forecasts',
    iconName: 'TrendingUp',
    color: 'text-purple-500',
  },
  {
    id: 'tech',
    name: 'Tech',
    description: 'Technology launches and industry news',
    iconName: 'Cpu',
    color: 'text-cyan-500',
  },
  {
    id: 'culture',
    name: 'Culture',
    description: 'Entertainment, awards, and cultural events',
    iconName: 'Star',
    color: 'text-pink-500',
  },
  {
    id: 'world',
    name: 'World',
    description: 'Global events and current affairs',
    iconName: 'MapPin',
    color: 'text-indigo-500',
  },
  {
    id: 'economy',
    name: 'Economy',
    description: 'Economic indicators and market trends',
    iconName: 'LineChart',
    color: 'text-emerald-500',
  },
  {
    id: 'elections',
    name: 'Elections',
    description: 'Political elections and voting outcomes',
    iconName: 'Vote',
    color: 'text-red-500',
  },
  {
    id: 'mentions',
    name: 'Mentions',
    description: 'Media mentions and trending topics',
    iconName: 'MessageSquare',
    color: 'text-yellow-500',
  },
];

export interface PredictionEvent {
  id: string;
  categoryId: string;
  title: string;
  description: string;
  outcomes: string[];
  deadline: string; // ISO date string
  status: 'open' | 'locked' | 'settled';
  createdAt?: string;
  settledOutcome?: string;
}

export function getCategoryById(categoryId: string): PredictionCategory | undefined {
  return predictionCategories.find(cat => cat.id === categoryId);
}
