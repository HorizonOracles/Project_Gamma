# Resolution Hooks

Hooks for proposing, disputing, and finalizing market resolutions.

## useResolution

Get resolution state for a market.

### Import

```tsx
import { useResolution } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: resolution } = useResolution(marketId);
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `ResolutionProposal \| undefined` | Resolution data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### Examples

```tsx
function ResolutionStatus({ marketId }: { marketId: number }) {
  const { data: resolution } = useResolution(marketId);

  if (!resolution) return <div>No resolution proposed</div>;

  return (
    <div>
      <p>Outcome: {resolution.outcome === 'YES' ? 'YES' : 'NO'}</p>
      <p>Evidence Hash: {resolution.evidenceHash}</p>
    </div>
  );
}
```

## useProposeResolution

Propose a resolution for a market.

### Import

```tsx
import { useProposeResolution } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: proposeResolution, isLoading } = useProposeResolution(marketId);

proposeResolution({
  outcomeId: 0, // YES
  evidenceHash: '0x...',
  signature: '0x...',
  bondAmount: parseUnits('1000', 18),
  evidenceURIs: ['ipfs://...'],
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### ProposeResolutionParams

```tsx
interface ProposeResolutionParams {
  outcomeId: number;        // 0 for YES, 1 for NO
  evidenceHash: string;     // Hash of evidence
  signature: string;         // EIP-712 signature
  bondAmount: bigint;        // Bond amount in HORIZON tokens
  evidenceURIs: string[];    // Array of evidence URIs
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `write` | `(params: ProposeResolutionParams) => void` | Function to propose resolution |
| `isLoading` | `boolean` | Transaction pending state |
| `isSuccess` | `boolean` | Transaction success state |
| `hash` | `string \| undefined` | Transaction hash |
| `error` | `Error \| null` | Error object if transaction failed |

### Examples

```tsx
function ProposeResolution({ marketId }: { marketId: number }) {
  const { write: proposeResolution, isLoading } = useProposeResolution(marketId);

  const handlePropose = () => {
    proposeResolution({
      outcomeId: 0,
      evidenceHash: '0x...',
      signature: '0x...',
      bondAmount: parseUnits('1000', 18),
      evidenceURIs: ['ipfs://Qm...'],
    });
  };

  return (
    <button onClick={handlePropose} disabled={isLoading}>
      {isLoading ? 'Proposing...' : 'Propose Resolution'}
    </button>
  );
}
```

## useDispute

Dispute a proposed resolution.

### Import

```tsx
import { useDispute } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: dispute, isLoading } = useDispute(marketId);

dispute({
  bondAmount: parseUnits('1000', 18),
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### DisputeParams

```tsx
interface DisputeParams {
  bondAmount: bigint;  // Bond amount in HORIZON tokens
}
```

### Returns

Same as `useProposeResolution`.

### Examples

```tsx
function DisputeResolution({ marketId }: { marketId: number }) {
  const { write: dispute, isLoading } = useDispute(marketId);

  const handleDispute = () => {
    dispute({
      bondAmount: parseUnits('1000', 18),
    });
  };

  return (
    <button onClick={handleDispute} disabled={isLoading}>
      {isLoading ? 'Disputing...' : 'Dispute Resolution'}
    </button>
  );
}
```

## useFinalize

Finalize a resolution after the dispute window has passed.

### Import

```tsx
import { useFinalize } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { write: finalize, isLoading } = useFinalize(marketId);

finalize();
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `write` | `() => void` | Function to finalize resolution |
| `isLoading` | `boolean` | Transaction pending state |
| `isSuccess` | `boolean` | Transaction success state |
| `hash` | `string \| undefined` | Transaction hash |
| `error` | `Error \| null` | Error object if transaction failed |

### Examples

```tsx
function FinalizeResolution({ marketId }: { marketId: number }) {
  const { write: finalize, isLoading } = useFinalize(marketId);

  return (
    <button onClick={() => finalize()} disabled={isLoading}>
      {isLoading ? 'Finalizing...' : 'Finalize Resolution'}
    </button>
  );
}
```

## Complete Resolution Flow Example

```tsx
import { 
  useResolution, 
  useProposeResolution, 
  useDispute, 
  useFinalize 
} from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';

function ResolutionPanel({ marketId }: { marketId: number }) {
  const { data: resolution } = useResolution(marketId);
  const { write: proposeResolution, isLoading: isProposing } = useProposeResolution(marketId);
  const { write: dispute, isLoading: isDisputing } = useDispute(marketId);
  const { write: finalize, isLoading: isFinalizing } = useFinalize(marketId);

  if (!resolution) {
    return (
      <div>
        <h3>No Resolution Proposed</h3>
        <button 
          onClick={() => proposeResolution({
            outcomeId: 0,
            evidenceHash: '0x...',
            signature: '0x...',
            bondAmount: parseUnits('1000', 18),
            evidenceURIs: ['ipfs://...'],
          })}
          disabled={isProposing}
        >
          {isProposing ? 'Proposing...' : 'Propose Resolution'}
        </button>
      </div>
    );
  }

  return (
    <div>
      <h3>Resolution Status</h3>
      <p>Outcome: {resolution.outcome}</p>
      <p>Evidence Hash: {resolution.evidenceHash}</p>
      
      <button 
        onClick={() => dispute({ bondAmount: parseUnits('1000', 18) })}
        disabled={isDisputing}
      >
        {isDisputing ? 'Disputing...' : 'Dispute'}
      </button>
      
      <button 
        onClick={() => finalize()}
        disabled={isFinalizing}
      >
        {isFinalizing ? 'Finalizing...' : 'Finalize'}
      </button>
    </div>
  );
}
```

