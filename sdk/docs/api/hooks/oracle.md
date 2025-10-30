# Oracle Hooks

Hooks for requesting AI resolution and monitoring oracle status.

## useRequestResolution

Request AI resolution for a market via the oracle API.

### Import

```tsx
import { useRequestResolution } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { mutate: requestResolution, data: request } = useRequestResolution();

requestResolution({
  marketId: 1,
  metadata: {
    question: 'Market question',
    description: 'Details',
  },
});
```

### Parameters

```tsx
interface RequestResolutionMutationParams {
  marketId: number;
  metadata: {
    question: string;
    description?: string;
  };
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `mutate` | `(params: RequestResolutionMutationParams) => void` | Function to request resolution |
| `data` | `RequestResolutionResponse \| undefined` | Request response data |
| `isLoading` | `boolean` | Request pending state |
| `isSuccess` | `boolean` | Request success state |
| `error` | `Error \| null` | Error object if request failed |

### RequestResolutionResponse

```tsx
interface RequestResolutionResponse {
  requestId: string;
  marketId: number;
  status: 'pending';
}
```

### Examples

```tsx
function RequestResolution({ marketId }: { marketId: number }) {
  const { mutate: requestResolution, isLoading } = useRequestResolution();

  const handleRequest = () => {
    requestResolution({
      marketId,
      metadata: {
        question: 'Will Team A win?',
        description: 'Details about the match',
      },
    });
  };

  return (
    <button onClick={handleRequest} disabled={isLoading}>
      {isLoading ? 'Requesting...' : 'Request AI Resolution'}
    </button>
  );
}
```

## useOracleStatus

Poll oracle request status.

### Import

```tsx
import { useOracleStatus } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: status } = useOracleStatus(requestId, {
  refetchInterval: 5000, // Poll every 5 seconds
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `requestId` | `string \| undefined` | Yes | Oracle request ID |
| `options` | `UseOracleStatusOptions` | No | Query options |

### UseOracleStatusOptions

```tsx
interface UseOracleStatusOptions {
  enabled?: boolean;         // Enable/disable query (default: true)
  refetchInterval?: number;  // Polling interval in ms
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `OracleRequest \| undefined` | Status data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### OracleRequest

```tsx
interface OracleRequest {
  requestId: string;
  marketId: number;
  status: 'pending' | 'processing' | 'completed' | 'failed';
  progress?: number;  // Optional progress percentage (0-100)
}
```

### Examples

```tsx
function OracleStatus({ requestId }: { requestId: string }) {
  const { data: status } = useOracleStatus(requestId, {
    refetchInterval: 5000, // Poll every 5 seconds
  });

  if (!status) return <div>Loading status...</div>;

  return (
    <div>
      <p>Status: {status.status}</p>
      {status.progress !== undefined && (
        <p>Progress: {status.progress}%</p>
      )}
    </div>
  );
}
```

## useOracleResult

Get completed oracle result.

### Import

```tsx
import { useOracleResult } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: result } = useOracleResult(requestId, {
  enabled: status?.status === 'completed',
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `requestId` | `string \| undefined` | Yes | Oracle request ID |
| `options` | `UseOracleResultOptions` | No | Query options |

### UseOracleResultOptions

```tsx
interface UseOracleResultOptions {
  enabled?: boolean;  // Enable/disable query (default: true)
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `OracleResult \| undefined` | Result data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### OracleResult

```tsx
interface OracleResult {
  requestId: string;
  marketId: number;
  outcomeId: number;      // 0 for YES, 1 for NO
  confidence: number;     // Confidence percentage (0-100)
  reasoning: string;      // AI reasoning
  sources: string[];      // Source URLs
  evidenceUrl: string;    // Evidence document URL
  timestamp?: number;     // Result timestamp
}
```

### Examples

```tsx
function OracleResult({ requestId }: { requestId: string }) {
  const { data: result } = useOracleResult(requestId);

  if (!result) return <div>No result yet</div>;

  return (
    <div>
      <h3>Oracle Result</h3>
      <p>Outcome: {result.outcomeId === 0 ? 'YES' : 'NO'}</p>
      <p>Confidence: {result.confidence}%</p>
      <p>Reasoning: {result.reasoning}</p>
      <a href={result.evidenceUrl} target="_blank" rel="noopener noreferrer">
        View Evidence
      </a>
    </div>
  );
}
```

## useOracleHistory

Get oracle request history for a market.

### Import

```tsx
import { useOracleHistory } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: history } = useOracleHistory(marketId);
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `OracleRequest[] \| undefined` | History array |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### Examples

```tsx
function OracleHistory({ marketId }: { marketId: number }) {
  const { data: history } = useOracleHistory(marketId);

  if (!history || history.length === 0) {
    return <div>No oracle requests</div>;
  }

  return (
    <div>
      <h3>Oracle History</h3>
      {history.map((request) => (
        <div key={request.requestId}>
          <p>Request ID: {request.requestId}</p>
          <p>Status: {request.status}</p>
        </div>
      ))}
    </div>
  );
}
```

## Complete Oracle Integration Example

```tsx
import { 
  useRequestResolution, 
  useOracleStatus, 
  useOracleResult 
} from '@project-gamma/react-sdk';

function OraclePanel({ marketId }: { marketId: number }) {
  const { mutate: requestResolution, data: request } = useRequestResolution();
  const { data: status } = useOracleStatus(request?.requestId, {
    enabled: !!request,
    refetchInterval: 5000,
  });
  const { data: result } = useOracleResult(request?.requestId, {
    enabled: status?.status === 'completed',
  });

  const handleRequest = () => {
    requestResolution({
      marketId,
      metadata: {
        question: 'Market question',
        description: 'Details',
      },
    });
  };

  return (
    <div>
      <button onClick={handleRequest}>
        Request AI Resolution
      </button>

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
          <h4>Oracle Result</h4>
          <p>Outcome: {result.outcomeId === 0 ? 'YES' : 'NO'}</p>
          <p>Confidence: {result.confidence}%</p>
          <p>Reasoning: {result.reasoning}</p>
          <a href={result.evidenceUrl} target="_blank" rel="noopener noreferrer">
            View Evidence
          </a>
        </div>
      )}
    </div>
  );
}
```

