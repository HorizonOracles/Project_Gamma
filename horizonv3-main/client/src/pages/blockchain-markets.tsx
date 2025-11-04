// Blockchain Markets Page - Fully Decentralized On-chain Prediction Markets
import { useState, useMemo } from "react";
import { useAuth } from "@/hooks/useAuth";
import { useMarkets, MarketStatus } from "@project-gamma/sdk";
import { TradingCard } from "@/components/TradingCard";
import { UserPositions } from "@/components/UserPositions";
import { ResolvedMarkets } from "@/components/ResolvedMarkets";
import { LiquidityPanel } from "@/components/LiquidityPanel";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LoadingSpinner } from "@/components/LoadingSpinner";
import { Skeleton } from "@/components/ui/skeleton";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { AlertCircle, LayoutGrid, List, Search, Filter } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export default function BlockchainMarkets() {
  const { isAuthenticated, isLoading: authLoading } = useAuth();
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');
  const [searchQuery, setSearchQuery] = useState("");
  const [categoryFilter, setCategoryFilter] = useState<string>("all");

  // Fetch blockchain markets from SDK (fully on-chain)
  const {
    data: markets,
    isLoading: marketsLoading,
    error: marketsError,
  } = useMarkets({
    status: MarketStatus.Active,
  });

  // Filter and search markets
  const filteredMarkets = useMemo(() => {
    if (!markets) return [];
    
    return markets.filter((market) => {
      // Category filter
      if (categoryFilter !== "all" && market.category !== categoryFilter) {
        return false;
      }
      
      // Search filter
      if (searchQuery) {
        const query = searchQuery.toLowerCase();
        const question = market.question?.toLowerCase() || "";
        const category = market.category?.toLowerCase() || "";
        
        if (!question.includes(query) && !category.includes(query)) {
          return false;
        }
      }
      
      return true;
    });
  }, [markets, categoryFilter, searchQuery]);

  // Get unique categories from markets
  const categories = useMemo(() => {
    if (!markets) return [];
    const uniqueCategories = new Set(markets.map(m => m.category).filter(Boolean));
    return Array.from(uniqueCategories);
  }, [markets]);

  if (authLoading) {
    return (
      <div className="flex items-center justify-center h-screen bg-background">
        <div className="text-center">
          <LoadingSpinner />
          <p className="text-lg text-muted-foreground">Loading...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      {/* Live Betting Feed */}
      <LiveBettingFeed />

      <div className="flex-1 overflow-auto relative">
        {/* Gold Network Background */}
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground
            color="gold"
            className="w-full h-full opacity-30"
            sizeMultiplier={1.25}
          />
        </div>

        <div className="container mx-auto p-6 max-w-7xl relative z-10">
          <div className="space-y-6">
            {/* Header */}
            <div className="flex flex-col gap-4">
              <div className="flex items-center justify-between">
                <div>
                  <h2 className="text-3xl font-bold font-sohne">
                    <span className="text-white">BLOCKCHAIN</span>{" "}
                    <span className="ml-1 goldify-text">MARKETS</span>
                  </h2>
                  <p className="text-sm text-muted-foreground mt-2">
                    Trade outcome tokens on-chain with full custody of your positions
                  </p>
                </div>

                {/* View Toggle */}
                <div className="flex gap-2">
                  <Button
                    variant={viewMode === 'grid' ? 'default' : 'outline'}
                    size="icon"
                    onClick={() => setViewMode('grid')}
                  >
                    <LayoutGrid className="h-4 w-4" />
                  </Button>
                  <Button
                    variant={viewMode === 'list' ? 'default' : 'outline'}
                    size="icon"
                    onClick={() => setViewMode('list')}
                  >
                    <List className="h-4 w-4" />
                  </Button>
                </div>
              </div>

              {/* Search and Filter Bar */}
              <div className="flex gap-3">
                <div className="relative flex-1">
                  <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Search markets..."
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    className="pl-9"
                  />
                </div>
                <Select value={categoryFilter} onValueChange={setCategoryFilter}>
                  <SelectTrigger className="w-[200px]">
                    <Filter className="h-4 w-4 mr-2" />
                    <SelectValue placeholder="All Categories" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">All Categories</SelectItem>
                    {categories.map((category) => (
                      <SelectItem key={category} value={category}>
                        {category}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </div>

            {/* Info Banner */}
            <Alert className="border-primary/30 bg-primary/10">
              <AlertCircle className="h-4 w-4 text-primary" />
              <AlertDescription className="text-sm">
                <strong>How it works:</strong> Buy YES or NO tokens with your collateral. Winning tokens can be redeemed 1:1 after market resolution. Your tokens are held in your wallet.
              </AlertDescription>
            </Alert>

            {/* Error State */}
            {marketsError && (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>
                  Failed to load blockchain markets. Please check your wallet connection and network.
                </AlertDescription>
              </Alert>
            )}

            {/* Loading State */}
            {marketsLoading && (
              <div
                className={
                  viewMode === 'grid'
                    ? "grid md:grid-cols-2 lg:grid-cols-3 gap-4"
                    : "space-y-3"
                }
              >
                {[1, 2, 3, 4, 5, 6].map((i) => (
                  <Skeleton
                    key={i}
                    className={viewMode === 'grid' ? "h-[500px] bg-card" : "h-48 bg-card"}
                  />
                ))}
              </div>
            )}

            {/* Markets Grid/List */}
            {!marketsLoading && markets && markets.length > 0 ? (
              <Tabs defaultValue="active" className="w-full">
                <TabsList className="grid w-full grid-cols-4">
                  <TabsTrigger value="active">
                    Active{" "}
                    <Badge variant="outline" className="ml-2">
                      {filteredMarkets.length}
                    </Badge>
                  </TabsTrigger>
                  <TabsTrigger value="your-positions">Your Positions</TabsTrigger>
                  <TabsTrigger value="liquidity">Liquidity</TabsTrigger>
                  <TabsTrigger value="resolved">Resolved</TabsTrigger>
                </TabsList>

                <TabsContent value="active" className="mt-6">
                  {filteredMarkets.length > 0 ? (
                    <div
                      className={
                        viewMode === 'grid'
                          ? "grid md:grid-cols-2 lg:grid-cols-3 gap-4"
                          : "space-y-3"
                      }
                    >
                      {filteredMarkets.map((market) => (
                        <TradingCard
                          key={market.id}
                          marketId={market.id}
                          title={market.question || market.metadataURI}
                          description={market.category}
                        />
                      ))}
                    </div>
                  ) : (
                    <div className="text-center py-12 bg-card rounded-lg border border-card-border">
                      <p className="text-muted-foreground text-lg">No markets match your filters</p>
                      <p className="text-sm text-muted-foreground mt-2">
                        Try adjusting your search or category filter
                      </p>
                    </div>
                  )}
                </TabsContent>

                <TabsContent value="your-positions" className="mt-6">
                  <UserPositions />
                </TabsContent>

                <TabsContent value="liquidity" className="mt-6">
                  <div className="space-y-6">
                    <div className="bg-card/50 rounded-lg border border-card-border p-6">
                      <h3 className="text-xl font-semibold mb-2">Provide Liquidity</h3>
                      <p className="text-sm text-muted-foreground mb-4">
                        Add liquidity to markets to earn trading fees. Select a market below to manage your liquidity position.
                      </p>
                    </div>
                    
                    {filteredMarkets.length > 0 ? (
                      <div
                        className={
                          viewMode === 'grid'
                            ? "grid md:grid-cols-2 lg:grid-cols-3 gap-4"
                            : "space-y-3"
                        }
                      >
                        {filteredMarkets.map((market) => (
                          <LiquidityPanel
                            key={market.id}
                            marketId={market.id}
                          />
                        ))}
                      </div>
                    ) : (
                      <div className="text-center py-12 bg-card rounded-lg border border-card-border">
                        <p className="text-muted-foreground text-lg">No markets available</p>
                        <p className="text-sm text-muted-foreground mt-2">
                          Check back soon or adjust your filters
                        </p>
                      </div>
                    )}
                  </div>
                </TabsContent>

                <TabsContent value="resolved" className="mt-6">
                  <ResolvedMarkets />
                </TabsContent>
              </Tabs>
            ) : (
              !marketsLoading && (
                <div className="text-center py-12 bg-card rounded-lg border border-card-border">
                  <p className="text-muted-foreground text-lg">No blockchain markets available</p>
                  <p className="text-sm text-muted-foreground mt-2">
                    Check back soon for new trading opportunities!
                  </p>
                </div>
              )
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
