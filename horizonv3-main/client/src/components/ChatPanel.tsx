import { useState, useEffect, useRef, createContext, useContext } from "react";
import { useQuery, useMutation } from "@tanstack/react-query";
import { apiRequest, queryClient } from "@/lib/queryClient";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { MessageCircle, X, Send } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { formatDistanceToNow } from "date-fns";
import { useTheme } from "@/components/ThemeProvider";

// Create context for chat panel state
const ChatContext = createContext<{
  isOpen: boolean;
  setIsOpen: (open: boolean) => void;
} | null>(null);

export const useChatPanel = () => {
  const context = useContext(ChatContext);
  if (!context) {
    throw new Error("useChatPanel must be used within ChatPanelProvider");
  }
  return context;
};

interface ChatMessage {
  id: string;
  userId: string;
  message: string;
  createdAt: Date;
  userEmail?: string;
  userDisplayName?: string;
  profileImageUrl?: string;
  rank?: string;
  rankPoints?: number;
}

export function ChatPanelProvider({ children }: { children: React.ReactNode }) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <ChatContext.Provider value={{ isOpen, setIsOpen }}>
      {children}
    </ChatContext.Provider>
  );
}

export function ChatPanelTrigger() {
  const { isOpen, setIsOpen } = useChatPanel();

  if (isOpen) return null;

  return (
    <Button
      onClick={() => setIsOpen(true)}
      size="icon"
      variant="ghost"
      className="text-primary hover:text-primary/80"
      data-testid="button-open-chat"
    >
      <MessageCircle className="h-5 w-5" />
    </Button>
  );
}

export function ChatPanel() {
  const { isOpen, setIsOpen } = useChatPanel();
  const { theme } = useTheme();
  const [message, setMessage] = useState("");
  const { toast } = useToast();
  const scrollRef = useRef<HTMLDivElement>(null);

  // Fetch chat messages
  const { data: messages = [] } = useQuery<ChatMessage[]>({
    queryKey: ["/api/chat/messages"],
    refetchInterval: isOpen ? 5000 : false,
  });

  // Send message mutation
  const sendMessageMutation = useMutation({
    mutationFn: async (text: string) => {
      const res = await apiRequest("POST", "/api/chat/messages", { message: text });
      return await res.json();
    },
    onSuccess: () => {
      setMessage("");
      queryClient.invalidateQueries({ queryKey: ["/api/chat/messages"] });
    },
    onError: (error: any) => {
      toast({
        title: "Send Failed",
        description: error.message || "Failed to send message",
        variant: "destructive",
      });
    },
  });

  // WebSocket real-time updates
  useEffect(() => {
    if (!isOpen) return;

    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const wsUrl = `${protocol}//${window.location.host}/ws`;
    
    const ws = new WebSocket(wsUrl);

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'chat_message') {
        queryClient.setQueryData<ChatMessage[]>(["/api/chat/messages"], (old = []) => {
          return [...old, data.data];
        });
      }
    };

    return () => {
      ws.close();
    };
  }, [isOpen]);

  // Auto-scroll to bottom on new messages
  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [messages]);

  const handleSendMessage = (e: React.FormEvent) => {
    e.preventDefault();
    if (message.trim() && !sendMessageMutation.isPending) {
      sendMessageMutation.mutate(message.trim());
    }
  };

  const getUserDisplayName = (msg: ChatMessage) => {
    // Never show email - use display name or anonymous
    return msg.userDisplayName || 'Anonymous Player';
  };

  // Helper to calculate level from rank points (same logic as LiveBettingFeed)
  const calculateLevel = (rankPoints: number): number => {
    if (rankPoints < 10000) return Math.floor(rankPoints / 100) + 1;
    if (rankPoints < 30000) return Math.floor((rankPoints - 10000) / 200) + 101;
    if (rankPoints < 60000) return Math.floor((rankPoints - 30000) / 300) + 201;
    if (rankPoints < 100000) return Math.floor((rankPoints - 60000) / 400) + 301;
    if (rankPoints < 150000) return Math.floor((rankPoints - 100000) / 500) + 401;
    if (rankPoints < 210000) return Math.floor((rankPoints - 150000) / 600) + 501;
    if (rankPoints < 280000) return Math.floor((rankPoints - 210000) / 700) + 601;
    if (rankPoints < 360000) return Math.floor((rankPoints - 280000) / 800) + 701;
    if (rankPoints < 450000) return Math.floor((rankPoints - 360000) / 900) + 801;
    if (rankPoints < 550000) return Math.floor((rankPoints - 450000) / 1000) + 901;
    if (rankPoints < 660000) return Math.floor((rankPoints - 550000) / 1100) + 1001;
    return Math.floor((rankPoints - 660000) / 1200) + 1101;
  };

  if (!isOpen) return null;

  return (
    <div className="h-screen w-80 bg-background border-l border-accent flex flex-col">
      {/* Header */}
      <div className="flex items-center justify-between px-4 h-28 border-b border-accent bg-background" style={{ zIndex: 101, position: 'relative' }}>
        <div className="flex items-center gap-2">
          <MessageCircle className="h-5 w-5 text-primary" />
          <h2 className="text-lg font-bold text-foreground">Live Chat</h2>
        </div>
        <Button
          onClick={() => setIsOpen(false)}
          size="icon"
          variant="ghost"
          className="h-8 w-8"
          data-testid="button-close-chat"
        >
          <X className="h-4 w-4" />
        </Button>
      </div>

      {/* Messages */}
      <ScrollArea className="flex-1 p-4" ref={scrollRef}>
        <div className="space-y-3">
          {messages.map((msg) => (
            <div key={msg.id} className="space-y-1" data-testid={`chat-message-${msg.id}`}>
              <div className="flex items-start gap-2">
                <Avatar className="h-6 w-6 border border-primary/30">
                  <AvatarImage src={msg.profileImageUrl || ""} alt={getUserDisplayName(msg)} />
                  <AvatarFallback className="text-[10px] font-bold text-primary bg-primary/10">
                    {getUserDisplayName(msg)[0].toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 flex-wrap">
                    <span 
                      className="text-sm font-semibold text-foreground" 
                      data-testid={`chat-username-${msg.id}`}
                    >
                      {getUserDisplayName(msg)}
                    </span>
                    <span className="text-xs text-muted-foreground">
                      {formatDistanceToNow(new Date(msg.createdAt), { addSuffix: true })}
                    </span>
                  </div>
                  <p className="text-sm text-foreground bg-background/50 rounded-lg p-2 break-words mt-1">
                    {msg.message}
                  </p>
                </div>
              </div>
            </div>
          ))}
          {messages.length === 0 && (
            <p className="text-center text-muted-foreground text-sm py-8">
              No messages yet. Start the conversation!
            </p>
          )}
        </div>
      </ScrollArea>

      {/* Input */}
      <form onSubmit={handleSendMessage} className="p-4 border-t border-accent bg-background">
        <div className="flex gap-2">
          <Input
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            placeholder="Type a message..."
            maxLength={500}
            disabled={sendMessageMutation.isPending}
            className="flex-1 bg-card border-accent focus:border-primary"
            data-testid="input-chat-message"
          />
          <Button
            type="submit"
            size="icon"
            disabled={!message.trim() || sendMessageMutation.isPending}
            className="bg-primary hover:bg-primary/90 text-background"
            data-testid="button-send-message"
          >
            <Send className="h-4 w-4" />
          </Button>
        </div>
        <p className="text-xs text-muted-foreground mt-1">
          {message.length}/500 characters
        </p>
      </form>
    </div>
  );
}
