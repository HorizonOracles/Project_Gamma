import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Search, Calendar, TrendingUp, AlertCircle, Plus } from "lucide-react";
import * as LucideIcons from "lucide-react";
import { predictionCategories, type PredictionCategory, type PredictionEvent } from "@shared/prediction-markets";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface PredictionMarketsDemoProps {
  categoryId?: string;
  isAdmin?: boolean;
}

export function PredictionMarketsDemo({ categoryId = "finance", isAdmin = false }: PredictionMarketsDemoProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const [eventsToShow, setEventsToShow] = useState(20);
  const [dialogOpen, setDialogOpen] = useState(false);
  
  const category = predictionCategories.find(c => c.id === categoryId);

  // Mock events for demonstration - in production, these would come from API
  const mockEvents: PredictionEvent[] = [
    {
      id: '1',
      categoryId: 'finance',
      title: 'Will S&P 500 close above 6000 by end of Q1 2025?',
      description: 'Predict whether the S&P 500 index will close above 6000 points by March 31, 2025',
      outcomes: ['Yes', 'No'],
      deadline: '2025-03-31T23:59:59Z',
      status: 'open' as const,
    },
    {
      id: '2',
      categoryId: 'crypto',
      title: 'Bitcoin to reach $100,000 in 2025?',
      description: 'Will Bitcoin price exceed $100,000 at any point during 2025?',
      outcomes: ['Yes', 'No'],
      deadline: '2025-12-31T23:59:59Z',
      status: 'open' as const,
    },
    {
      id: '3',
      categoryId: 'elections',
      title: '2024 US Presidential Election Winner',
      description: 'Who will win the 2024 United States Presidential Election?',
      outcomes: ['Democrat', 'Republican', 'Other'],
      deadline: '2024-11-05T23:59:59Z',
      status: 'locked' as const,
    },
  ].filter(event => event.categoryId === categoryId);

  const filteredEvents = mockEvents.filter(event => 
    event.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
    event.description.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const displayedEvents = filteredEvents.slice(0, eventsToShow);
  const hasMoreEvents = filteredEvents.length > eventsToShow;

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'open':
        return <Badge variant="default" className="bg-green-500">Open</Badge>;
      case 'locked':
        return <Badge variant="secondary">Locked</Badge>;
      case 'settled':
        return <Badge variant="outline">Settled</Badge>;
      default:
        return <Badge>{status}</Badge>;
    }
  };

  if (!category) {
    return (
      <Alert variant="destructive">
        <AlertCircle className="h-4 w-4" />
        <AlertDescription>Category not found</AlertDescription>
      </Alert>
    );
  }

  const IconComponent = (LucideIcons as any)[category.iconName] || LucideIcons.Circle;

  return (
    <div className="space-y-6">
      {/* Category Header */}
      <Card className="border-accent/20">
        <CardHeader>
          <CardTitle className="flex items-center gap-3">
            <IconComponent className={`h-6 w-6 ${category.color}`} />
            <span className="font-sohne">{category.name} Markets</span>
          </CardTitle>
          <CardDescription>{category.description}</CardDescription>
        </CardHeader>
      </Card>

      {/* Search and Actions */}
      <div className="flex items-center gap-4">
        <div className="relative flex-1 max-w-md">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Search prediction markets..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-10"
            data-testid="input-search-predictions"
          />
        </div>
        
        {isAdmin && (
          <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
            <DialogTrigger asChild>
              <Button className="darkify" style={{ borderColor: '#424242' }} data-testid="button-add-prediction">
                <Plus className="h-4 w-4 mr-2" />
                Add Market
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Create Prediction Market</DialogTitle>
                <DialogDescription>
                  Add a new prediction market for users to bet on
                </DialogDescription>
              </DialogHeader>
              <div className="space-y-4 py-4">
                <div className="space-y-2">
                  <Label htmlFor="title">Market Title</Label>
                  <Input id="title" placeholder="Enter market question..." />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Description</Label>
                  <Textarea id="description" placeholder="Provide details about the market..." />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="deadline">Deadline</Label>
                  <Input id="deadline" type="datetime-local" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="outcomes">Number of Outcomes</Label>
                  <Select defaultValue="2">
                    <SelectTrigger id="outcomes">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="2">2 (Yes/No)</SelectItem>
                      <SelectItem value="3">3 Options</SelectItem>
                      <SelectItem value="4">4 Options</SelectItem>
                      <SelectItem value="5">5+ Options</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <DialogFooter>
                <Button variant="outline" onClick={() => setDialogOpen(false)}>Cancel</Button>
                <Button className="goldify" style={{ borderColor: '#d5b877' }}>Create Market</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        )}
      </div>

      {/* Events List */}
      {filteredEvents.length === 0 ? (
        <Card className="border-accent/20">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <AlertCircle className="h-12 w-12 text-muted-foreground mb-4" />
            <p className="text-lg font-medium text-muted-foreground">No markets found</p>
            <p className="text-sm text-muted-foreground">
              {searchTerm ? 'Try adjusting your search terms' : 'Check back later for new markets'}
            </p>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-3">
          <div className="flex items-center justify-between">
            <p className="text-sm text-muted-foreground">
              Showing {displayedEvents.length} of {filteredEvents.length} markets
            </p>
          </div>
          
          {displayedEvents.map((event) => (
            <Card key={event.id} className="border-accent/20 hover-elevate">
              <CardHeader>
                <div className="flex items-start justify-between gap-4">
                  <div className="flex-1">
                    <CardTitle className="text-lg font-sohne mb-2">{event.title}</CardTitle>
                    <CardDescription>{event.description}</CardDescription>
                  </div>
                  {getStatusBadge(event.status)}
                </div>
              </CardHeader>
              <CardContent>
                <div className="flex items-center gap-6 text-sm text-muted-foreground">
                  <div className="flex items-center gap-2">
                    <Calendar className="h-4 w-4" />
                    <span>Deadline: {new Date(event.deadline).toLocaleDateString()}</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <TrendingUp className="h-4 w-4" />
                    <span>{event.outcomes.length} Outcomes</span>
                  </div>
                </div>
                <div className="flex gap-2 mt-4">
                  {event.outcomes.map((outcome, idx) => (
                    <Badge key={idx} variant="outline">
                      {outcome}
                    </Badge>
                  ))}
                </div>
              </CardContent>
            </Card>
          ))}

          {hasMoreEvents && (
            <div className="flex justify-center pt-4">
              <Button
                variant="outline"
                onClick={() => setEventsToShow(eventsToShow + 20)}
                data-testid="button-show-more-predictions"
              >
                Show More Markets
              </Button>
            </div>
          )}
          
          {eventsToShow > 20 && (
            <div className="flex justify-center">
              <Button
                variant="outline"
                onClick={() => setEventsToShow(20)}
                data-testid="button-show-less-predictions"
              >
                Show Less
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
