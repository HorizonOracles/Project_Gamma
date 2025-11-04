// Admin Markets Page - Market Creation, Liquidity & Resolution Management
import { useState, useMemo } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeft, List, Plus, ExternalLink, Info, Filter, X } from "lucide-react";
import { useLocation } from "wouter";
import { useQuery } from "@tanstack/react-query";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";
import { MarketCreationForm } from "@/components/MarketCreationForm";
import { LiquidityManager } from "@/components/LiquidityManager";
import { ResolutionPanel } from "@/components/ResolutionPanel";
import { useAdmin } from "@/hooks/useAdmin";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";

// USDC address on Avalanche Fuji Testnet
// TODO: Move this to environment variables
const COLLATERAL_TOKEN_ADDRESS = "0x5425890298aed601595a70AB815c96711a31Bc65"; // USDC on Fuji

export default function AdminMarketsPage() {
  const [, setLocation] = useLocation();
  const { isAdmin, isConnected } = useAdmin();
  const [selectedMarketId, setSelectedMarketId] = useState<number | null>(null);
  const [activeTab, setActiveTab] = useState<string>("manage");
  
  // Filter state
  const [statusFilter, setStatusFilter] = useState<string>("all");
  const [sportFilter, setSportFilter] = useState<string>("all");
  const [searchQuery, setSearchQuery] = useState<string>("");
  
  // Fetch all markets
  const { data: markets, isLoading: isLoadingMarkets } = useQuery<any[]>({
    queryKey: ["/api/markets"],
    retry: false,
  });

  // Filter markets based on filters
  const filteredMarkets = useMemo(() => {
    if (!markets) return [];
    
    return markets.filter((market: any) => {
      // Status filter
      if (statusFilter !== "all" && market.status !== statusFilter) {
        return false;
      }
      
      // Sport filter
      if (sportFilter !== "all" && market.sport !== sportFilter) {
        return false;
      }
      
      // Search query filter
      if (searchQuery) {
        const query = searchQuery.toLowerCase();
        const matchesTeam = market.teamA?.toLowerCase().includes(query) || 
                           market.teamB?.toLowerCase().includes(query);
        const matchesLeague = market.league?.toLowerCase().includes(query);
        const matchesDescription = market.description?.toLowerCase().includes(query);
        
        if (!matchesTeam && !matchesLeague && !matchesDescription) {
          return false;
        }
      }
      
      return true;
    });
  }, [markets, statusFilter, sportFilter, searchQuery]);

  // Get unique sports from markets
  const availableSports = useMemo(() => {
    if (!markets) return [];
    const sports = new Set(markets.map((m: any) => m.sport).filter(Boolean));
    return Array.from(sports).sort();
  }, [markets]);

  // Redirect if not admin
  if (!isConnected || !isAdmin) {
    setLocation('/admin');
    return null;
  }

  const handleMarketCreated = (marketId: bigint) => {
    // When a market is created, select it and switch to manage tab
    setSelectedMarketId(Number(marketId));
    setActiveTab("manage");
  };

  return (
    <div className="flex flex-col h-full">
      {/* Live Betting Feed */}
      <LiveBettingFeed />
      
      <div className="flex-1 overflow-auto relative">
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
        </div>
        
        <div className="container mx-auto p-6 max-w-screen-xl relative z-10">
          {/* Header */}
          <div className="mb-8 flex items-center gap-4">
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setLocation('/admin')}
              className="hover:bg-accent/20"
            >
              <ArrowLeft className="h-5 w-5" />
            </Button>
            <div>
              <h1 className="text-3xl font-bold text-foreground font-sohne">
                Markets & <span className="text-primary">Settlement</span>
              </h1>
              <p className="text-muted-foreground">
                Create markets, manage liquidity, and resolve outcomes
              </p>
            </div>
          </div>

          {/* Tabs for Create Market and Manage Markets */}
          <Tabs value={activeTab} onValueChange={setActiveTab} className="w-full">
            <TabsList className="grid w-full max-w-md mx-auto grid-cols-2 mb-6">
              <TabsTrigger value="create" className="flex items-center gap-2">
                <Plus className="h-4 w-4" />
                Create Market
              </TabsTrigger>
              <TabsTrigger value="manage" className="flex items-center gap-2">
                <List className="h-4 w-4" />
                Manage Markets
              </TabsTrigger>
            </TabsList>

            {/* Create Market Tab */}
            <TabsContent value="create">
              <div className="mb-6">
                <MarketCreationForm
                  collateralTokenAddress={COLLATERAL_TOKEN_ADDRESS}
                  onSuccess={handleMarketCreated}
                />
              </div>
            </TabsContent>

            {/* Manage Markets Tab */}
            <TabsContent value="manage">
              {isLoadingMarkets ? (
                <Card className="border-accent/20 p-8 text-center">
                  <p className="text-muted-foreground">Loading markets...</p>
                </Card>
              ) : !markets || markets.length === 0 ? (
                <Card className="border-accent/20 p-8 text-center">
                  <p className="text-muted-foreground">
                    No markets found. Create your first market to get started.
                  </p>
                </Card>
              ) : (
                <div className="space-y-6">
                  {/* Filters */}
                  <Card className="border-accent/20">
                    <CardContent className="p-4">
                      <div className="flex items-center gap-2 mb-3">
                        <Filter className="h-4 w-4 text-muted-foreground" />
                        <h3 className="font-semibold text-sm">Filters</h3>
                        {(statusFilter !== "all" || sportFilter !== "all" || searchQuery) && (
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => {
                              setStatusFilter("all");
                              setSportFilter("all");
                              setSearchQuery("");
                            }}
                            className="ml-auto h-7 px-2 text-xs"
                          >
                            <X className="h-3 w-3 mr-1" />
                            Clear
                          </Button>
                        )}
                      </div>
                      <div className="grid grid-cols-1 md:grid-cols-3 gap-3">
                        {/* Search */}
                        <Input
                          placeholder="Search teams, league..."
                          value={searchQuery}
                          onChange={(e) => setSearchQuery(e.target.value)}
                          className="h-9"
                        />
                        
                        {/* Status Filter */}
                        <Select value={statusFilter} onValueChange={setStatusFilter}>
                          <SelectTrigger className="h-9">
                            <SelectValue placeholder="All Statuses" />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectItem value="all">All Statuses</SelectItem>
                            <SelectItem value="pending">Pending</SelectItem>
                            <SelectItem value="active">Active</SelectItem>
                            <SelectItem value="locked">Locked</SelectItem>
                            <SelectItem value="resolved">Resolved</SelectItem>
                            <SelectItem value="cancelled">Cancelled</SelectItem>
                          </SelectContent>
                        </Select>
                        
                        {/* Sport Filter */}
                        <Select value={sportFilter} onValueChange={setSportFilter}>
                          <SelectTrigger className="h-9">
                            <SelectValue placeholder="All Sports" />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectItem value="all">All Sports</SelectItem>
                            {availableSports.map((sport) => (
                              <SelectItem key={sport} value={sport}>
                                {sport}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </div>
                      
                      {/* Filter Results Count */}
                      <div className="mt-3 text-xs text-muted-foreground">
                        Showing {filteredMarkets.length} of {markets.length} markets
                      </div>
                    </CardContent>
                  </Card>
                  
                  {filteredMarkets.length === 0 ? (
                    <Card className="border-accent/20 p-8 text-center">
                      <p className="text-muted-foreground">
                        No markets match your filters. Try adjusting your search criteria.
                      </p>
                    </Card>
                  ) : (
                    <>
                  {/* Market List */}
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {filteredMarkets.map((market: any) => (
                      <Card
                        key={market.id}
                        className={`border-accent/20 cursor-pointer transition-all hover:border-primary/50 ${
                          selectedMarketId === market.id ? 'border-primary ring-2 ring-primary/20' : ''
                        }`}
                        onClick={() => setSelectedMarketId(market.id)}
                      >
                        <CardContent className="p-4">
                          <div className="flex items-start gap-3 mb-3">
                            {market.teamALogo && (
                              <img src={market.teamALogo} alt={market.teamA} className="w-10 h-10 rounded" />
                            )}
                            <div className="flex-1 min-w-0">
                              <h3 className="font-semibold text-sm truncate">
                                {market.teamA} vs {market.teamB}
                              </h3>
                              <p className="text-xs text-muted-foreground">
                                {market.sport} • {market.league}
                              </p>
                            </div>
                            {market.teamBLogo && (
                              <img src={market.teamBLogo} alt={market.teamB} className="w-10 h-10 rounded" />
                            )}
                          </div>
                          
                          <div className="space-y-2 text-xs mb-3">
                            <div className="flex justify-between">
                              <span className="text-muted-foreground">Status:</span>
                              <span className={`font-medium ${
                                market.status === 'active' ? 'text-green-500' :
                                market.status === 'resolved' ? 'text-blue-500' :
                                'text-yellow-500'
                              }`}>
                                {market.status}
                              </span>
                            </div>
                            {market.gameTime && (
                              <div className="flex justify-between">
                                <span className="text-muted-foreground">Game:</span>
                                <span>{new Date(market.gameTime).toLocaleDateString()}</span>
                              </div>
                            )}
                            {market.marketAddress && (
                              <div className="flex justify-between">
                                <span className="text-muted-foreground">On-Chain:</span>
                                <span className="text-green-500">✓</span>
                              </div>
                            )}
                          </div>

                          {/* Blockchain Details - Collapsible */}
                          {market.marketAddress && (
                            <Collapsible>
                              <CollapsibleTrigger className="flex items-center gap-1 text-xs text-primary hover:text-primary/80 w-full justify-center py-2 border-t border-accent/20">
                                <Info className="h-3 w-3" />
                                <span>Blockchain Info</span>
                              </CollapsibleTrigger>
                              <CollapsibleContent className="pt-2 space-y-2 text-xs">
                                <div className="flex justify-between">
                                  <span className="text-muted-foreground">Chain ID:</span>
                                  <span className="font-mono">{market.chainId || 56}</span>
                                </div>
                                {market.onChainMarketId !== null && market.onChainMarketId !== undefined && (
                                  <div className="flex justify-between">
                                    <span className="text-muted-foreground">Market ID:</span>
                                    <span className="font-mono">{market.onChainMarketId}</span>
                                  </div>
                                )}
                                <div className="flex justify-between items-center">
                                  <span className="text-muted-foreground">Address:</span>
                                  <a
                                    href={`https://bscscan.com/address/${market.marketAddress}`}
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    className="flex items-center gap-1 text-primary hover:text-primary/80 font-mono text-xs"
                                    onClick={(e) => e.stopPropagation()}
                                  >
                                    {market.marketAddress.slice(0, 6)}...{market.marketAddress.slice(-4)}
                                    <ExternalLink className="h-3 w-3" />
                                  </a>
                                </div>
                                {market.yesTokenId !== null && market.yesTokenId !== undefined && (
                                  <div className="flex justify-between">
                                    <span className="text-muted-foreground">Yes Token:</span>
                                    <span className="font-mono">{market.yesTokenId}</span>
                                  </div>
                                )}
                                {market.noTokenId !== null && market.noTokenId !== undefined && (
                                  <div className="flex justify-between">
                                    <span className="text-muted-foreground">No Token:</span>
                                    <span className="font-mono">{market.noTokenId}</span>
                                  </div>
                                )}
                                {market.resolutionSource && (
                                  <div className="flex justify-between">
                                    <span className="text-muted-foreground">Resolution:</span>
                                    <span className="capitalize">{market.resolutionSource}</span>
                                  </div>
                                )}
                              </CollapsibleContent>
                            </Collapsible>
                          )}
                        </CardContent>
                      </Card>
                    ))}
                  </div>

                  {/* Liquidity & Resolution Management for Selected Market */}
                  {selectedMarketId && (
                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 mt-6">
                      <LiquidityManager
                        marketId={selectedMarketId}
                        collateralDecimals={6} // USDC has 6 decimals
                      />
                      <ResolutionPanel
                        marketId={selectedMarketId}
                      />
                    </div>
                  )}
                  </>
                )}
                </div>
              )}
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </div>
  );
}
