# Oracle Integration Example

A complete example of integrating AI oracle resolution into your application.

## Complete Example

```tsx
import { 
  useRequestResolution, 
  useOracleStatus, 
  useOracleResult,
  useOracleHistory 
} from '@project-gamma/react-sdk';
import { useState, useEffect } from 'react';

function OracleIntegration({ marketId }: { marketId: number }) {
  const [requestId, setRequestId] = useState<string | null>(null);

  // Request resolution
  const { mutate: requestResolution, isLoading: isRequesting } = useRequestResolution();

  // Monitor status (polls every 5s)
  const { data: status } = useOracleStatus(requestId || undefined, {
    enabled: !!requestId,
    refetchInterval: 5000,
  });

  // Get result when completed
  const { data: result } = useOracleResult(requestId || undefined, {
    enabled: status?.status === 'completed',
  });

  // Get history
  const { data: history } = useOracleHistory(marketId);

  const handleRequest = () => {
    requestResolution({
      marketId,
      metadata: {
        question: 'Will Team A win the match?',
        description: 'Match details and context',
      },
    }, {
      onSuccess: (data) => {
        setRequestId(data.requestId);
      },
    });
  };

  return (
    <div>
      <h2>AI Oracle Resolution</h2>

      {/* Request Resolution */}
      <div>
        <button onClick={handleRequest} disabled={isRequesting}>
          {isRequesting ? 'Requesting...' : 'Request AI Resolution'}
        </button>
      </div>

      {/* Status Display */}
      {status && (
        <div>
          <h3>Status</h3>
          <p>Request ID: {status.requestId}</p>
          <p>Status: {status.status}</p>
          {status.progress !== undefined && (
            <div>
              <p>Progress: {status.progress}%</p>
              <progress value={status.progress} max={100} />
            </div>
          )}
        </div>
      )}

      {/* Result Display */}
      {result && (
        <div>
          <h3>Oracle Result</h3>
          <div>
            <p><strong>Outcome:</strong> {result.outcomeId === 0 ? 'YES' : 'NO'}</p>
            <p><strong>Confidence:</strong> {result.confidence}%</p>
            <p><strong>Reasoning:</strong></p>
            <p>{result.reasoning}</p>
            <p><strong>Sources:</strong></p>
            <ul>
              {result.sources.map((source, i) => (
                <li key={i}>
                  <a href={source} target="_blank" rel="noopener noreferrer">
                    {source}
                  </a>
                </li>
              ))}
            </ul>
            <a 
              href={result.evidenceUrl} 
              target="_blank" 
              rel="noopener noreferrer"
            >
              View Full Evidence
            </a>
          </div>
        </div>
      )}

      {/* History */}
      {history && history.length > 0 && (
        <div>
          <h3>Request History</h3>
          <table>
            <thead>
              <tr>
                <th>Request ID</th>
                <th>Status</th>
                <th>Progress</th>
              </tr>
            </thead>
            <tbody>
              {history.map((request) => (
                <tr key={request.requestId}>
                  <td>{request.requestId}</td>
                  <td>{request.status}</td>
                  <td>{request.progress ?? '-'}%</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

// Usage with Market Component
function MarketWithOracle({ marketId }: { marketId: number }) {
  return (
    <div>
      <h2>Market {marketId}</h2>
      <OracleIntegration marketId={marketId} />
    </div>
  );
}
```

## Step-by-Step Flow

### 1. Request Resolution

```tsx
const { mutate: requestResolution } = useRequestResolution();

requestResolution({
  marketId: 1,
  metadata: {
    question: 'Market question',
    description: 'Additional context',
  },
}, {
  onSuccess: (data) => {
    console.log('Request ID:', data.requestId);
  },
});
```

### 2. Monitor Status

```tsx
const { data: status } = useOracleStatus(requestId, {
  refetchInterval: 5000, // Poll every 5 seconds
});

// Status will be: 'pending' -> 'processing' -> 'completed' or 'failed'
```

### 3. Get Result

```tsx
const { data: result } = useOracleResult(requestId, {
  enabled: status?.status === 'completed',
});

// Result includes outcome, confidence, reasoning, sources, and evidence
```

## Advanced Example with Auto-Refresh

```tsx
function AutoRefreshOracle({ marketId }: { marketId: number }) {
  const [requestId, setRequestId] = useState<string | null>(null);
  
  const { mutate: requestResolution } = useRequestResolution();
  const { data: status } = useOracleStatus(requestId || undefined, {
    enabled: !!requestId,
    refetchInterval: (data) => {
      // Stop polling if completed or failed
      if (data?.status === 'completed' || data?.status === 'failed') {
        return false;
      }
      return 5000; // Poll every 5 seconds
    },
  });
  const { data: result } = useOracleResult(requestId || undefined, {
    enabled: status?.status === 'completed',
  });

  const handleRequest = () => {
    requestResolution({
      marketId,
      metadata: {
        question: 'Market question',
      },
    }, {
      onSuccess: (data) => {
        setRequestId(data.requestId);
      },
    });
  };

  return (
    <div>
      {!requestId && (
        <button onClick={handleRequest}>
          Request Resolution
        </button>
      )}

      {status && (
        <div>
          <p>Status: {status.status}</p>
          {status.progress !== undefined && (
            <p>Progress: {status.progress}%</p>
          )}
        </div>
      )}

      {result && (
        <div>
          <h3>Result: {result.outcomeId === 0 ? 'YES' : 'NO'}</h3>
          <p>Confidence: {result.confidence}%</p>
          <p>{result.reasoning}</p>
        </div>
      )}
    </div>
  );
}
```

## Error Handling

```tsx
function OracleWithErrorHandling({ marketId }: { marketId: number }) {
  const { mutate: requestResolution, error: requestError } = useRequestResolution();
  const { data: status, error: statusError } = useOracleStatus(requestId);
  const { data: result, error: resultError } = useOracleResult(requestId);

  // Handle errors
  useEffect(() => {
    if (requestError) {
      console.error('Request error:', requestError);
      // Show user-friendly error message
    }
    if (statusError) {
      console.error('Status error:', statusError);
    }
    if (resultError) {
      console.error('Result error:', resultError);
    }
  }, [requestError, statusError, resultError]);

  // ... rest of component
}
```

## Best Practices

1. **Polling**: Use `refetchInterval` to poll status, but disable when completed
2. **Error Handling**: Always handle errors gracefully
3. **Loading States**: Show loading indicators during requests
4. **User Feedback**: Display progress and status updates
5. **Result Validation**: Verify result data before using it

