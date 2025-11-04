// Horizon Main App Component
// Reference: javascript_log_in_with_replit blueprint for auth routing
// Reference: design_guidelines.md - Shadcn sidebar implementation
import { useEffect } from "react";
import { Switch, Route } from "wouter";
import { queryClient } from "./lib/queryClient";
import { QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider } from "wagmi";
import { RainbowKitProvider } from "@rainbow-me/rainbowkit";
import "@rainbow-me/rainbowkit/styles.css";
import { GammaProvider } from "@project-gamma/sdk";
import { config as wagmiConfig } from "./lib/wagmi-config";
import { Toaster } from "@/components/ui/toaster";
import { TooltipProvider } from "@/components/ui/tooltip";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import { ChatPanel, ChatPanelProvider, ChatPanelTrigger } from "@/components/ChatPanel";
import { NetworkBackground } from "@/components/NetworkBackground";
import { ThemeProvider } from "@/components/ThemeProvider";
import { ThemeToggle } from "@/components/ThemeToggle";
import { Button } from "@/components/ui/button";
import { Flame, LogIn } from "lucide-react";
import { useAuth } from "@/hooks/useAuth";
import { useAdmin } from "@/hooks/useAdmin";
import NotFound from "@/pages/not-found";
import Home from "@/pages/home";
import Wallet from "@/pages/wallet";
import Profile from "@/pages/profile";
import Feed from "@/pages/feed";
import Admin from "@/pages/admin";
import AdminMarkets from "@/pages/admin-markets";
import Leaderboard from "@/pages/leaderboard";
import MyBets from "@/pages/my-bets";
import LoginPage from "@/pages/login";
import BlockchainMarkets from "@/pages/blockchain-markets";
import CreateMarket from "@/pages/create-market";

function AdminRoute({ component: Component }: { component: React.ComponentType }) {
  const { isAdmin } = useAdmin();
  const { isAuthenticated } = useAuth();
  
  if (!isAuthenticated) {
    return <LoginPage />;
  }
  
  if (!isAdmin) {
    return <NotFound />;
  }
  
  return <Component />;
}

function Router() {
  return (
    <Switch>
      <Route path="/login" component={LoginPage} />
      <Route path="/" component={Home} />
      <Route path="/blockchain-markets" component={BlockchainMarkets} />
      <Route path="/my-bets" component={MyBets} />
      <Route path="/leaderboard" component={Leaderboard} />
      <Route path="/wallet" component={Wallet} />
      <Route path="/profile" component={Profile} />
      <Route path="/feed" component={Feed} />
      <Route path="/create-market">
        {() => <AdminRoute component={CreateMarket} />}
      </Route>
      <Route path="/admin/markets">
        {() => <AdminRoute component={AdminMarkets} />}
      </Route>
      <Route path="/admin">
        {() => <AdminRoute component={Admin} />}
      </Route>
      <Route component={NotFound} />
    </Switch>
  );
}

function AuthenticatedLayout() {
  const { isAuthenticated, isLoading } = useAuth();
  
  // Custom sidebar width for prediction markets application - matches chat panel width
  const style = {
    "--sidebar-width": "20rem",       // 320px to match chat panel (w-80)
    "--sidebar-width-icon": "4rem",   // default icon width
  };

  // Show same layout for both logged-in and logged-out users
  // Logged-out: sidebar with markets only, no chat, no account section
  // Logged-in: sidebar with markets + account, chat enabled
  return (
    <ChatPanelProvider>
      <SidebarProvider style={style as React.CSSProperties}>
        <div className="flex h-screen w-full bg-background">
          {/* Unified Network Background spanning all three headers */}
          <div className="absolute top-0 left-0 right-0 h-28 pointer-events-none" style={{ zIndex: 100 }}>
            <NetworkBackground 
              className="w-full h-full opacity-50" 
              color="gold"
            />
          </div>
          
          <AppSidebar />
          <div className="flex flex-col flex-1 overflow-hidden">
            <header className="flex items-center justify-between px-4 h-28 border-b border-accent bg-transparent relative" style={{ zIndex: 101 }}>
              <SidebarTrigger data-testid="button-sidebar-toggle" className="text-primary hover:text-primary/80" />
              <div className="flex items-center gap-4">
                <ThemeToggle />
                {isAuthenticated && <ChatPanelTrigger />}
              </div>
            </header>
            <main className="flex-1 overflow-auto bg-background">
              <Router />
            </main>
          </div>
          {isAuthenticated && <ChatPanel />}
        </div>
      </SidebarProvider>
    </ChatPanelProvider>
  );
}

function App() {
  // Disable right-click on all images
  useEffect(() => {
    const handleContextMenu = (e: MouseEvent) => {
      const target = e.target as HTMLElement;
      if (target.tagName === 'IMG') {
        e.preventDefault();
        return false;
      }
    };

    document.addEventListener('contextmenu', handleContextMenu);
    return () => {
      document.removeEventListener('contextmenu', handleContextMenu);
    };
  }, []);

  // SDK Configuration from environment variables
  const sdkConfig = {
    chainId: import.meta.env.VITE_CHAIN_ID ? parseInt(import.meta.env.VITE_CHAIN_ID) : 56,
    oracleApiUrl: import.meta.env.VITE_ORACLE_API_URL || 'https://api.oraprotocol.ai',
    marketFactoryAddress: (import.meta.env.VITE_MARKET_FACTORY_ADDRESS || '0x22Cc806047BB825aa26b766Af737E92B1866E8A6') as `0x${string}`,
    horizonTokenAddress: (import.meta.env.VITE_HORIZON_TOKEN_ADDRESS || '0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444') as `0x${string}`,
    outcomeTokenAddress: (import.meta.env.VITE_OUTCOME_TOKEN_ADDRESS || '0x17B322784265c105a94e4c3d00aF1E5f46a5F311') as `0x${string}`,
    horizonPerksAddress: (import.meta.env.VITE_HORIZON_PERKS_ADDRESS || '0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C') as `0x${string}`,
    feeSplitterAddress: (import.meta.env.VITE_FEE_SPLITTER_ADDRESS || '0x275017E98adF33051BbF477fe1DD197F681d4eF1') as `0x${string}`,
    resolutionModuleAddress: (import.meta.env.VITE_RESOLUTION_MODULE_ADDRESS || '0xF0CF4C741910cB48AC596F620a0AE892Cd247838') as `0x${string}`,
    aiOracleAdapterAddress: (import.meta.env.VITE_AI_ORACLE_ADAPTER_ADDRESS || '0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93') as `0x${string}`,
    pinataJwt: import.meta.env.VITE_PINATA_JWT,
  };

  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider>
          <GammaProvider
            chainId={sdkConfig.chainId}
            oracleApiUrl={sdkConfig.oracleApiUrl}
            marketFactoryAddress={sdkConfig.marketFactoryAddress}
            horizonTokenAddress={sdkConfig.horizonTokenAddress}
            outcomeTokenAddress={sdkConfig.outcomeTokenAddress}
            horizonPerksAddress={sdkConfig.horizonPerksAddress}
            feeSplitterAddress={sdkConfig.feeSplitterAddress}
            resolutionModuleAddress={sdkConfig.resolutionModuleAddress}
            aiOracleAdapterAddress={sdkConfig.aiOracleAdapterAddress}
            pinataJwt={sdkConfig.pinataJwt}
          >
            <ThemeProvider>
              <TooltipProvider>
                <AuthenticatedLayout />
                <Toaster />
              </TooltipProvider>
            </ThemeProvider>
          </GammaProvider>
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

export default App;
