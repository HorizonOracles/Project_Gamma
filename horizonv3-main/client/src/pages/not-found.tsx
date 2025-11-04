import { Card, CardContent } from "@/components/ui/card";
import { AlertCircle } from "lucide-react";
import { NetworkBackground } from "@/components/NetworkBackground";

export default function NotFound() {
  return (
    <div className="min-h-screen w-full flex items-center justify-center bg-background relative">
      <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
        <NetworkBackground color="gold" className="w-full h-full opacity-30" sizeMultiplier={1.25} />
      </div>
      <Card className="w-full max-w-md mx-4 relative z-10">
        <CardContent className="pt-6">
          <div className="flex mb-4 gap-2">
            <AlertCircle className="h-8 w-8 text-primary" />
            <h1 className="text-2xl font-bold text-foreground">404 Page Not Found</h1>
          </div>

          <p className="mt-4 text-sm text-muted-foreground">
            Did you forget to add the page to the router?
          </p>
        </CardContent>
      </Card>
    </div>
  );
}
