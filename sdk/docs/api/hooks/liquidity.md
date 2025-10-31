# Liquidity Hooks

Hooks for adding and removing liquidity, and tracking LP positions.

## useAddLiquidity

Add liquidity to a market's AMM pool.

### Import

```tsx
import { useAddLiquidity } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: addLiquidity, isLoading } = useAddLiquidity(marketId);

addLiquidity({
  amount: parseUnits('1000', 6), // 1000 USDC
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### AddLiquidityParams

```tsx
interface AddLiquidityParams {
  amount: bigint;  // Amount of collateral to add
}
```

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `write` | `(params: AddLiquidityParams) => void` | Function to add liquidity |
| `isLoading` | `boolean` | Transaction pending state |
| `isSuccess` | `boolean` | Transaction success state |
| `hash` | `string \| undefined` | Transaction hash |
| `error` | `Error \| null` | Error object if transaction failed |

### Examples

```tsx
function AddLiquidityPanel({ marketId }: { marketId: number }) {
  const [amount, setAmount] = useState('1000');
  const { write: addLiquidity, isLoading } = useAddLiquidity(marketId);

  const handleAdd = () => {
    addLiquidity({
      amount: parseUnits(amount, 6),
    });
  };

  return (
    <div>
      <input
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
      />
      <button onClick={handleAdd} disabled={isLoading}>
        {isLoading ? 'Adding...' : 'Add Liquidity'}
      </button>
    </div>
  );
}
```

## useRemoveLiquidity

Remove liquidity from a market's AMM pool.

### Import

```tsx
import { useRemoveLiquidity } from '@project-gamma/react-sdk';
import { parseUnits } from 'viem';
```

### Usage

```tsx
const { write: removeLiquidity, isLoading } = useRemoveLiquidity(marketId);

removeLiquidity({
  lpTokens: parseUnits('100', 18),
});
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### RemoveLiquidityParams

```tsx
interface RemoveLiquidityParams {
  lpTokens: bigint;  // Amount of LP tokens to remove
}
```

### Returns

Same as `useAddLiquidity`.

### Examples

```tsx
function RemoveLiquidityPanel({ marketId }: { marketId: number }) {
  const { data: position } = useLPPosition(marketId);
  const { write: removeLiquidity, isLoading } = useRemoveLiquidity(marketId);

  const handleRemove = () => {
    if (position) {
      removeLiquidity({
        lpTokens: position.lpTokens, // Remove all
      });
    }
  };

  return (
    <button onClick={handleRemove} disabled={isLoading || !position}>
      {isLoading ? 'Removing...' : 'Remove Liquidity'}
    </button>
  );
}
```

## useLPPosition

Get LP position information for a market.

### Import

```tsx
import { useLPPosition } from '@project-gamma/react-sdk';
```

### Usage

```tsx
const { data: position } = useLPPosition(marketId);
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `marketId` | `number` | Yes | Market ID |

### Returns

| Property | Type | Description |
|----------|------|-------------|
| `data` | `LPPosition \| undefined` | LP position data |
| `isLoading` | `boolean` | Loading state |
| `error` | `Error \| null` | Error object if query failed |

### LPPosition

```tsx
interface LPPosition {
  lpTokens: bigint;      // LP tokens owned
  value: bigint;         // Current value in collateral
  yesReserve: bigint;    // YES reserve in pool
  noReserve: bigint;     // NO reserve in pool
}
```

### Examples

```tsx
import { useLPPosition } from '@project-gamma/react-sdk';
import { formatUnits } from 'viem';

function LPPositionDisplay({ marketId }: { marketId: number }) {
  const { data: position, isLoading } = useLPPosition(marketId);

  if (isLoading) return <div>Loading position...</div>;
  if (!position) return <div>No position found</div>;

  return (
    <div>
      <h3>Your LP Position</h3>
      <p>LP Tokens: {formatUnits(position.lpTokens, 18)}</p>
      <p>Value: ${formatUnits(position.value, 6)}</p>
      <p>YES Reserve: {formatUnits(position.yesReserve, 18)}</p>
      <p>NO Reserve: {formatUnits(position.noReserve, 18)}</p>
    </div>
  );
}
```

## Complete Liquidity Example

```tsx
import { useAddLiquidity, useRemoveLiquidity, useLPPosition } from '@project-gamma/react-sdk';
import { parseUnits, formatUnits } from 'viem';

function LiquidityPanel({ marketId }: { marketId: number }) {
  const [addAmount, setAddAmount] = useState('1000');
  const [removeAmount, setRemoveAmount] = useState('100');

  const { data: position } = useLPPosition(marketId);
  const { write: addLiquidity, isLoading: isAdding } = useAddLiquidity(marketId);
  const { write: removeLiquidity, isLoading: isRemoving } = useRemoveLiquidity(marketId);

  const handleAdd = () => {
    addLiquidity({
      amount: parseUnits(addAmount, 6),
    });
  };

  const handleRemove = () => {
    removeLiquidity({
      lpTokens: parseUnits(removeAmount, 18),
    });
  };

  return (
    <div>
      <h3>Liquidity Management</h3>
      
      {position && (
        <div>
          <p>Your LP Tokens: {formatUnits(position.lpTokens, 18)}</p>
          <p>Current Value: ${formatUnits(position.value, 6)}</p>
        </div>
      )}

      <div>
        <h4>Add Liquidity</h4>
        <input
          value={addAmount}
          onChange={(e) => setAddAmount(e.target.value)}
          placeholder="Amount (USDC)"
        />
        <button onClick={handleAdd} disabled={isAdding}>
          {isAdding ? 'Adding...' : 'Add'}
        </button>
      </div>

      <div>
        <h4>Remove Liquidity</h4>
        <input
          value={removeAmount}
          onChange={(e) => setRemoveAmount(e.target.value)}
          placeholder="LP Tokens"
        />
        <button onClick={handleRemove} disabled={isRemoving}>
          {isRemoving ? 'Removing...' : 'Remove'}
        </button>
      </div>
    </div>
  );
}
```

