// LiveBettingFeed - Displays recent betting activity
// TODO: Implement full betting feed functionality
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export function LiveBettingFeed() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Live Betting Feed</CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-muted-foreground">
          Recent betting activity will appear here.
        </p>
      </CardContent>
    </Card>
  );
}
