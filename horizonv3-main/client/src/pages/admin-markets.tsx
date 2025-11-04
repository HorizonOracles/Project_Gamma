// Horizon Admin Panel - Market Management
import { useState, useEffect, useMemo } from "react";
import { useQuery, useMutation } from "@tanstack/react-query";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Lock, CheckCircle, Trophy, AlertCircle, ArrowLeft, Search, Grid3x3, Gamepad2, Link as LinkIcon, ExternalLink, Plus } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { useAuth } from "@/hooks/useAuth";
import { apiRequest, queryClient } from "@/lib/queryClient";
import type { Market } from "@shared/schema";
import { useLocation } from "wouter";
import { sportsData } from "@shared/market-categories";
import { useAccount } from 'wagmi';
import { useMarket, useCreateMarket, useMinCreatorStake, useUploadMetadata } from '@project-gamma/sdk';
import { formatEther, parseEther } from 'viem';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { NetworkBackground } from "@/components/NetworkBackground";
import { LiveBettingFeed } from "@/components/LiveBettingFeed";

// Component to display on-chain market stats
function OnChainMarketStats({ marketId }: { marketId: number }) {
  const { data: marketData, isLoading } = useMarket(marketId);

  if (isLoading) {
    return (
      <div className="flex items-center gap-2 text-xs text-muted-foreground">
        <div className="animate-spin h-3 w-3 border-2 border-primary border-t-transparent rounded-full"></div>
        Loading on-chain data...
      </div>
    );
  }

  if (!marketData) {
    return null;
  }

  const totalLiquidity = marketData.totalLiquidity
    ? Number(formatEther(marketData.totalLiquidity.yes + marketData.totalLiquidity.no))
    : 0;
  const totalVolume = marketData.totalVolume
    ? Number(formatEther(marketData.totalVolume))
    : 0;

  return (
    <div className="flex items-center gap-2 text-xs">
      <div className="flex items-center gap-1 text-green-400">
        <span className="text-muted-foreground">Liquidity:</span>
        <span className="font-mono font-bold">{totalLiquidity.toFixed(4)} BNB</span>
      </div>
      <div className="flex items-center gap-1 text-blue-400">
        <span className="text-muted-foreground">Volume:</span>
        <span className="font-mono font-bold">{totalVolume.toFixed(4)} BNB</span>
      </div>
    </div>
  );
}


export default function AdminMarkets() {
  const { toast } = useToast();
  const { isAuthenticated, isLoading } = useAuth();
  const { address, isConnected } = useAccount();
  const [, setLocation] = useLocation();
  const [settleDialogOpen, setSettleDialogOpen] = useState(false);
  const [createMarketDialogOpen, setCreateMarketDialogOpen] = useState(false);
  const [selectedMarket, setSelectedMarket] = useState<Market | null>(null);
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedSport, setSelectedSport] = useState<string>("all");

  // Get minimum creator stake for market creation
  const { data: minStake } = useMinCreatorStake();
  
  // Create market - note: useCreateMarket returns { write, isLoading, ... }
  const createMarketHook = useCreateMarket();

  // Sport icon mapping using existing sport icons from the website
  const getSportIcon = (sport: string) => {
    if (sport === 'all') {
      return <Grid3x3 className="h-5 w-5" />;
    }
    
    // Find matching sport config
    const sportConfig = sportsData.find(s => 
      s.name.toLowerCase() === sport.toLowerCase() || 
      s.id.toLowerCase() === sport.toLowerCase()
    );
    
    if (sportConfig) {
      // Handle lucide icons (e.g., Gamepad2 for ESports)
      if (sportConfig.iconType === 'lucide' || sportConfig.iconName === 'Gamepad2') {
        return <Gamepad2 className="h-5 w-5" />;
      }
      
      // Handle custom (PNG) and SVG icons
      const iconPath = sportConfig.iconType === 'custom'
        ? `/sport-icons/${sportConfig.iconName}.png`
        : `/sport-icons/${sportConfig.iconName}.svg`;
      
      return (
        <img 
          src={iconPath}
          alt={sportConfig.name}
          className="h-5 w-5 object-contain brightness-0 invert"
        />
      );
    }
    
    // Fallback icon
    return <Grid3x3 className="h-5 w-5" />;
  };

  // Redirect if not authenticated
  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      toast({
        title: "Unauthorized",
        description: "You are logged out. Logging in again...",
        variant: "destructive",
      });
      setTimeout(() => {
        window.location.href = "/api/login";
      }, 500);
    }
  }, [isAuthenticated, isLoading, toast]);

  // Fetch all markets (admin endpoint includes locked and settled)
  const { data: markets, isLoading: marketsLoading } = useQuery<Market[]>({
    queryKey: ["/api/admin/markets"],
    refetchInterval: 5000,
  });

  // Get all sports from sportsData config ordered by popularity (left to right)
  const availableSports = useMemo(() => {
    const popularityOrder = [
      'Football',
      'Basketball',
      'American Football',
      'Baseball',
      'Tennis',
      'Counter Strike',
      'Dota 2',
      'Valorant',
      'League of Legends',
      'Ice Hockey',
      'Cricket',
      'Golf',
      'Fighting',
      'Rugby',
      'Motorsport',
      'Table Tennis',
      'Badminton',
      'ESports'
    ];
    
    return sportsData
      .map(s => s.name)
      .sort((a, b) => {
        const indexA = popularityOrder.indexOf(a);
        const indexB = popularityOrder.indexOf(b);
        
        // If both are in the popularity list, sort by that order
        if (indexA !== -1 && indexB !== -1) {
          return indexA - indexB;
        }
        
        // If only one is in the list, prioritize it
        if (indexA !== -1) return -1;
        if (indexB !== -1) return 1;
        
        // If neither is in the list, sort alphabetically
        return a.localeCompare(b);
      });
  }, []);

  // Filter markets based on search and sport selection
  const filteredMarkets = useMemo(() => {
    if (!markets) return [];
    
    return markets.filter(market => {
      // Sport filter
      if (selectedSport !== "all" && market.sport !== selectedSport) {
        return false;
      }
      
      // Search filter
      if (searchQuery) {
        const query = searchQuery.toLowerCase();
        return (
          market.description.toLowerCase().includes(query) ||
          market.teamA.toLowerCase().includes(query) ||
          market.teamB.toLowerCase().includes(query) ||
          market.sport.toLowerCase().includes(query) ||
          market.league.toLowerCase().includes(query)
        );
      }
      
      return true;
    });
  }, [markets, selectedSport, searchQuery]);

  // Lock market mutation
  const lockMarketMutation = useMutation({
    mutationFn: async (marketId: string) => {
      return await apiRequest("POST", `/api/markets/${marketId}/lock`, {});
    },
    onSuccess: () => {
      toast({
        title: "Market Locked",
        description: "Betting has been closed for this market",
      });
      queryClient.invalidateQueries({ queryKey: ["/api/admin/markets"] });
    },
    onError: (error: any) => {
      toast({
        title: "Lock Failed",
        description: error.message || "Failed to lock market",
        variant: "destructive",
      });
    },
  });

  // Settle market mutation
  const settleMarketMutation = useMutation({
    mutationFn: async ({ marketId, winningOutcome }: { marketId: string; winningOutcome: 'A' | 'B' }) => {
      return await apiRequest("POST", `/api/markets/${marketId}/settle`, { winningOutcome });
    },
    onSuccess: (data: any) => {
      toast({
        title: "Market Settled",
        description: `Processed ${data.payoutsProcessed} winning bets`,
      });
      setSettleDialogOpen(false);
      setSelectedMarket(null);
      queryClient.invalidateQueries({ queryKey: ["/api/admin/markets"] });
    },
    onError: (error: any) => {
      toast({
        title: "Settlement Failed",
        description: error.message || "Failed to settle market",
        variant: "destructive",
      });
    },
  });

  const handleLockMarket = (market: Market) => {
    if (confirm(`Lock betting for "${market.description}"?`)) {
      lockMarketMutation.mutate(market.id);
    }
  };

  const handleSettleMarket = (market: Market, winningOutcome: 'A' | 'B') => {
    settleMarketMutation.mutate({ marketId: market.id, winningOutcome });
  };

  const openSettleDialog = (market: Market) => {
    setSelectedMarket(market);
    setSettleDialogOpen(true);
  };

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return <Badge className="bg-primary text-primary-foreground">Active</Badge>;
      case 'locked':
        return <Badge className="bg-accent text-accent-foreground">Locked</Badge>;
      case 'settled':
        return <Badge className="bg-card text-card-foreground border">Settled</Badge>;
      default:
        return <Badge>{status}</Badge>;
    }
  };

  // Helper to get BSCScan URL
  const getBscScanUrl = (marketAddress: string, chainId: number) => {
    const baseUrl = chainId === 56 ? 'https://bscscan.com' : 'https://testnet.bscscan.com';
    return `${baseUrl}/address/${marketAddress}`;
  };

  return (
    <div className="flex-1 overflow-auto relative">
      <LiveBettingFeed />
      <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
        <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
      </div>
      <div className="container mx-auto p-6 max-w-screen-xl relative z-10">
        <div className="mb-4">
          <div className="flex items-center justify-between mb-4">
            <div>
              <h1 className="text-3xl font-bold text-foreground mb-2 font-sohne">
                Markets & <span className="text-primary">Settlements</span>
              </h1>
              <p className="text-muted-foreground">Manage betting markets and process settlements</p>
            </div>
            
            {/* Search Bar */}
            <div className="flex items-center gap-3">
              <Button
                variant="default"
                onClick={() => setCreateMarketDialogOpen(true)}
                disabled={!isConnected}
                data-testid="button-create-onchain-market"
                className="goldify text-primary-foreground border-2"
                style={{ borderColor: '#d5b877' }}
                title={!isConnected ? "Connect wallet to create on-chain markets" : undefined}
              >
                <Plus className="h-4 w-4 mr-2" />
                Create On-Chain Market
              </Button>
              <div className="relative w-96">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                <Input
                  type="text"
                  placeholder="Search markets..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="pl-10"
                  data-testid="input-search-markets"
                />
              </div>
              <Button
                variant="default"
                onClick={() => setLocation('/admin')}
                data-testid="button-back-to-admin"
                className="darkify text-white border-2"
                style={{ borderColor: '#424242' }}
              >
                <ArrowLeft className="h-4 w-4 mr-2" />
                Back to Admin
              </Button>
            </div>
          </div>

          {/* Sport Filter Tabs - Using sport icons from website */}
          <div className="flex items-center gap-1.5 flex-wrap">
            <Button
              variant="default"
              size="sm"
              onClick={() => setSelectedSport("all")}
              data-testid="button-filter-all"
              className={selectedSport === "all" ? "h-9 w-9 p-0 darkify text-white border-2" : "h-9 w-9 p-0"}
              style={selectedSport === "all" ? { borderColor: '#424242' } : undefined}
              title="All Sports"
            >
              {getSportIcon("all")}
            </Button>
            {availableSports.map((sport) => (
              <Button
                key={sport}
                variant="default"
                size="sm"
                onClick={() => setSelectedSport(sport)}
                data-testid={`button-filter-${sport.toLowerCase().replace(/\s+/g, '-')}`}
                className={selectedSport === sport ? "h-9 w-9 p-0 darkify text-white border-2" : "h-9 w-9 p-0"}
                style={selectedSport === sport ? { borderColor: '#424242' } : undefined}
                title={sport}
              >
                {getSportIcon(sport)}
              </Button>
            ))}
          </div>
        </div>

      {marketsLoading ? (
        <div className="space-y-2">
          <Skeleton className="h-20 bg-card" />
          <Skeleton className="h-20 bg-card" />
          <Skeleton className="h-20 bg-card" />
        </div>
      ) : filteredMarkets.length === 0 ? (
        <div className="text-center py-12">
          <p className="text-muted-foreground">No markets found matching your filters.</p>
        </div>
      ) : (
        <div className="space-y-1.5">
          {filteredMarkets.map((market) => (
            <Card key={market.id} className="border-accent/20 bg-card">
              <CardContent className="p-2.5">
                {/* Responsive Two-Row Compact Layout */}
                <div className="flex flex-wrap items-center gap-2">
                  {/* Row 1: Badge, Title, Logos, Actions */}
                  <div className="flex items-center gap-2 flex-1 min-w-[400px]">
                    {/* Sport Badge & Title */}
                    <div className="flex-1 min-w-0">
                      <div className="flex items-center gap-1.5 mb-0.5">
                        <Badge className="bg-primary/20 text-primary border-primary/30 text-xs px-1.5 py-0 shrink-0">
                          {market.sport}
                        </Badge>
                        {market.isLive && (
                          <Badge className="bg-accent text-accent-foreground animate-pulse text-xs px-1.5 py-0 shrink-0">LIVE</Badge>
                        )}
                        {/* Blockchain Indicator */}
                        {market.onChainMarketId && market.marketAddress && (
                          <Badge className="bg-green-500/20 text-green-400 border-green-500/30 text-xs px-1.5 py-0 shrink-0 flex items-center gap-1">
                            <LinkIcon className="h-3 w-3" />
                            On-Chain
                          </Badge>
                        )}
                        {getStatusBadge(market.status)}
                      </div>
                      <h3 className="text-sm font-bold text-foreground truncate">
                        {market.description}
                      </h3>
                      <p className="text-xs text-muted-foreground truncate">
                        {market.teamA} vs {market.teamB}
                        {/* Show market ID for on-chain markets */}
                        {market.onChainMarketId && (
                          <span className="ml-2 text-green-400">
                            (Market #{market.onChainMarketId})
                          </span>
                        )}
                      </p>
                      {/* On-chain stats */}
                      {market.onChainMarketId && (
                        <OnChainMarketStats marketId={parseInt(market.onChainMarketId, 10)} />
                      )}
                    </div>

                    {/* Team/Player Logos */}
                    <div className="flex items-center gap-2 shrink-0">
                      {market.teamALogo && (
                        <div className="w-8 h-8 flex items-center justify-center">
                          <img
                            src={market.teamALogo}
                            alt={market.teamA}
                            className="max-w-full max-h-full object-contain"
                            onError={(e) => { e.currentTarget.style.display = 'none'; }}
                          />
                        </div>
                      )}
                      {market.teamBLogo && (
                        <div className="w-8 h-8 flex items-center justify-center">
                          <img
                            src={market.teamBLogo}
                            alt={market.teamB}
                            className="max-w-full max-h-full object-contain"
                            onError={(e) => { e.currentTarget.style.display = 'none'; }}
                          />
                        </div>
                      )}
                    </div>
                  </div>

                  {/* Pool Info - wraps to second row on smaller screens */}
                  <div className="flex gap-2 shrink-0">
                    <div className="bg-background/50 px-2.5 py-1.5 rounded border border-accent/10 min-w-[100px]">
                      <div className="text-xs text-muted-foreground">Pool A</div>
                      <div className="text-sm font-mono font-bold text-primary">
                        {parseFloat(market.poolATotal).toFixed(2)} BNB
                      </div>
                      <div className="text-xs text-accent">
                        {(() => {
                          const poolA = parseFloat(market.poolATotal || "0");
                          const poolB = parseFloat(market.poolBTotal || "0");
                          const bonus = parseFloat(market.bonusPool || "0");
                          const total = poolA + poolB + bonus;
                          return poolA > 0 ? (total / poolA).toFixed(2) : "2.00";
                        })()}x
                      </div>
                    </div>
                    <div className="bg-background/50 px-2.5 py-1.5 rounded border border-accent/10 min-w-[100px]">
                      <div className="text-xs text-muted-foreground">Pool B</div>
                      <div className="text-sm font-mono font-bold text-primary">
                        {parseFloat(market.poolBTotal).toFixed(2)} BNB
                      </div>
                      <div className="text-xs text-accent">
                        {(() => {
                          const poolA = parseFloat(market.poolATotal || "0");
                          const poolB = parseFloat(market.poolBTotal || "0");
                          const bonus = parseFloat(market.bonusPool || "0");
                          const total = poolA + poolB + bonus;
                          return poolB > 0 ? (total / poolB).toFixed(2) : "2.00";
                        })()}x
                      </div>
                    </div>
                  </div>

                  {/* Action Buttons - Using darkify theme for consistency */}
                  <div className="flex gap-2 shrink-0 ml-auto">
                    {/* BSCScan Link for on-chain markets */}
                    {market.onChainMarketId && market.marketAddress && market.chainId && (
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => window.open(getBscScanUrl(market.marketAddress!, market.chainId!), '_blank')}
                        data-testid={`button-bscscan-${market.id}`}
                        className="border-green-500/30 text-green-400 hover:bg-green-500/10"
                        title="View on BSCScan"
                      >
                        <ExternalLink className="h-4 w-4 mr-1.5" />
                        BSCScan
                      </Button>
                    )}
                    {market.status === 'active' && (
                      <Button
                        variant="default"
                        size="sm"
                        onClick={() => handleLockMarket(market)}
                        disabled={lockMarketMutation.isPending}
                        data-testid={`button-lock-${market.id}`}
                        className="darkify text-white border-2"
                        style={{ borderColor: '#424242' }}
                      >
                        <Lock className="h-4 w-4 mr-1.5" />
                        Lock
                      </Button>
                    )}
                    {(market.status === 'locked' || market.status === 'active') && (
                      <Button
                        variant="default"
                        size="sm"
                        onClick={() => openSettleDialog(market)}
                        disabled={settleMarketMutation.isPending}
                        data-testid={`button-settle-${market.id}`}
                        className="goldify text-primary-foreground border-2"
                        style={{ borderColor: '#d5b877' }}
                      >
                        <CheckCircle className="h-4 w-4 mr-1.5" />
                        Settle
                      </Button>
                    )}
                    {market.status === 'settled' && (
                      <div className="flex items-center gap-1.5 text-sm text-muted-foreground px-2">
                        <Trophy className="h-4 w-4 text-primary" />
                        <span className="truncate max-w-[120px]">
                          {market.winningOutcome === 'A' ? market.teamA : market.teamB}
                        </span>
                      </div>
                    )}
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      {/* Settle Market Dialog */}
      <Dialog open={settleDialogOpen} onOpenChange={setSettleDialogOpen}>
        <DialogContent className="bg-background border-accent">
          <DialogHeader>
            <DialogTitle className="text-foreground">Settle Market</DialogTitle>
            <DialogDescription className="text-muted-foreground">
              Select the winning outcome for: {selectedMarket?.description}
            </DialogDescription>
          </DialogHeader>

          <div className="space-y-3 py-4">
            <Alert className="border-accent/20">
              <AlertCircle className="h-4 w-4 text-primary" />
              <AlertDescription className="text-sm text-muted-foreground">
                This will distribute payouts to all winning bets and mark losing bets.
                This action cannot be undone.
              </AlertDescription>
            </Alert>

            <div className="grid grid-cols-2 gap-3">
              <Button
                variant="outline"
                onClick={() => selectedMarket && handleSettleMarket(selectedMarket, 'A')}
                disabled={settleMarketMutation.isPending}
                className="h-20 flex-col"
                data-testid="button-settle-outcome-a"
              >
                <div className="text-lg font-bold">{selectedMarket?.teamA}</div>
                <div className="text-xs text-muted-foreground">Outcome A Wins</div>
              </Button>
              <Button
                variant="outline"
                onClick={() => selectedMarket && handleSettleMarket(selectedMarket, 'B')}
                disabled={settleMarketMutation.isPending}
                className="h-20 flex-col"
                data-testid="button-settle-outcome-b"
              >
                <div className="text-lg font-bold">{selectedMarket?.teamB}</div>
                <div className="text-xs text-muted-foreground">Outcome B Wins</div>
              </Button>
            </div>
          </div>

          <DialogFooter>
            <Button
              variant="ghost"
              onClick={() => setSettleDialogOpen(false)}
              data-testid="button-cancel-settle"
            >
              Cancel
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Create On-Chain Market Dialog */}
      <Dialog open={createMarketDialogOpen} onOpenChange={setCreateMarketDialogOpen}>
        <DialogContent className="bg-background border-accent max-w-2xl max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle className="text-foreground flex items-center gap-2">
              <LinkIcon className="h-5 w-5 text-green-400" />
              Create On-Chain Market
            </DialogTitle>
            <DialogDescription className="text-muted-foreground">
              Deploy a new prediction market to the blockchain
            </DialogDescription>
          </DialogHeader>

          <div className="space-y-4 py-4">
            <Alert className="border-blue-500/20 bg-blue-500/5">
              <AlertCircle className="h-4 w-4 text-blue-400" />
              <AlertDescription className="text-sm text-muted-foreground">
                This feature is under development. On-chain market creation will allow you to:
                <ul className="list-disc list-inside mt-2 space-y-1">
                  <li>Deploy a new market contract to BNB Chain</li>
                  <li>Set initial liquidity and parameters</li>
                  <li>Automatically sync with the database</li>
                  <li>Enable decentralized trading immediately</li>
                </ul>
              </AlertDescription>
            </Alert>

            {minStake && Number(minStake) > 0 ? (
              <div className="p-3 bg-muted rounded-lg">
                <p className="text-xs text-muted-foreground mb-1">Minimum Creator Stake:</p>
                <code className="text-sm font-mono text-foreground">{formatEther(minStake)} BNB</code>
              </div>
            ) : null}

            <Alert className="border-yellow-500/20 bg-yellow-500/5">
              <AlertCircle className="h-4 w-4 text-yellow-400" />
              <AlertDescription className="text-sm text-muted-foreground">
                For now, please use the existing market creation workflow in Sports Data Integration or Prediction Markets Integration, 
                then manually deploy the market on-chain using the SDK.
              </AlertDescription>
            </Alert>
          </div>

          <DialogFooter>
            <Button
              variant="ghost"
              onClick={() => setCreateMarketDialogOpen(false)}
              data-testid="button-cancel-create-market"
            >
              Close
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
      </div>
    </div>
  );
}
