# Troubleshooting

Common issues and solutions when using the Project Gamma React SDK.

## Installation Issues

### Peer Dependencies Not Found

**Error**: `Cannot find module 'wagmi'` or similar

**Solution**: Install all peer dependencies:

```bash
npm install wagmi viem @tanstack/react-query
```

### TypeScript Errors

**Error**: Type errors after installation

**Solution**: 
1. Ensure TypeScript version is 5.0+
2. Clear TypeScript cache: `rm -rf node_modules/.cache`
3. Restart your IDE/TypeScript server

## Provider Setup Issues

### "useGammaConfig must be used within GammaProvider"

**Error**: Hooks are called outside of GammaProvider

**Solution**: Ensure your component tree is wrapped:

```tsx
<WagmiProvider config={wagmiConfig}>
  <QueryClientProvider client={queryClient}>
    <GammaProvider chainId={56}>
      <YourComponent /> {/* Hooks work here */}
    </GammaProvider>
  </QueryClientProvider>
</WagmiProvider>
```

### Wallet Not Connected

**Error**: "Wallet not connected" when calling hooks

**Solution**: 
1. Ensure Wagmi is properly configured
2. Connect wallet before using trading hooks
3. Check wallet connection status:

```tsx
import { useAccount } from 'wagmi';

function MyComponent() {
  const { isConnected } = useAccount();
  if (!isConnected) return <div>Please connect wallet</div>;
  // Use hooks here
}
```

## Contract Interaction Issues

### "Public client not available"

**Error**: Public client is undefined

**Solution**: 
1. Ensure Wagmi config includes public client setup
2. Check chain configuration matches SDK chainId
3. Verify RPC endpoint is accessible

### Transaction Failures

**Error**: Transactions fail silently or with cryptic errors

**Solution**:
1. Check wallet has sufficient balance
2. Verify token approvals (for trading)
3. Check slippage tolerance is appropriate
4. Ensure market is in correct state (active, not resolved)

### Slippage Errors

**Error**: "Slippage tolerance exceeded"

**Solution**:
1. Increase slippage tolerance (e.g., from 0.5% to 1%)
2. Reduce trade amount
3. Check market liquidity

## API Issues

### Oracle API Not Responding

**Error**: Oracle API requests fail

**Solution**:
1. Verify `oracleApiUrl` is correct in GammaProvider
2. Check API endpoint is accessible
3. Verify network connectivity
4. Check API rate limits

### Request Status Stuck

**Error**: Oracle request status doesn't update

**Solution**:
1. Verify polling is enabled: `refetchInterval: 5000`
2. Check request ID is valid
3. Manually refetch: `refetch()`
4. Check API logs for errors

## Data Fetching Issues

### Markets Not Loading

**Error**: `useMarkets` returns empty array or undefined

**Solution**:
1. Check MarketFactory address is correct
2. Verify chain ID matches deployed contracts
3. Check RPC endpoint is working
4. Verify markets exist on chain

### Prices Not Updating

**Error**: Prices show stale data

**Solution**:
1. Manually refetch: `refetch()`
2. Check query cache invalidation
3. Verify market AMM address is correct
4. Check market is active

## TypeScript Issues

### Type Errors in Hooks

**Error**: TypeScript errors in hook usage

**Solution**:
1. Ensure all types are imported correctly
2. Check TypeScript version (5.0+)
3. Verify types are exported from SDK
4. Clear TypeScript cache

### Missing Type Definitions

**Error**: Cannot find type definitions

**Solution**:
1. Ensure SDK is installed: `npm install @project-gamma/react-sdk`
2. Check `dist/index.d.ts` exists
3. Verify TypeScript can resolve types
4. Restart TypeScript server

## Performance Issues

### Slow Queries

**Error**: Queries take too long

**Solution**:
1. Use appropriate RPC endpoint (premium providers are faster)
2. Implement caching strategies
3. Reduce unnecessary refetches
4. Use query filters to limit data

### Too Many Re-renders

**Error**: Components re-render excessively

**Solution**:
1. Use React.memo for expensive components
2. Memoize callback functions
3. Avoid creating new objects in render
4. Use query selectors to transform data

## Build Issues

### Bundle Size Too Large

**Error**: Bundle size exceeds expectations

**Solution**:
1. Use tree-shaking (import specific hooks)
2. Exclude unused dependencies
3. Use dynamic imports for large features
4. Check for duplicate dependencies

### Build Fails

**Error**: Build process fails

**Solution**:
1. Check Node.js version (18+)
2. Clear build cache
3. Verify all dependencies are installed
4. Check build tool configuration

## Getting Help

If you're still experiencing issues:

1. **Check Documentation**: Review the [API Reference](./api/)
2. **Search Issues**: Check [GitHub Issues](https://github.com/HorizonOracles/Project_Gamma/issues)
3. **Create Issue**: Open a new issue with:
   - Error message
   - Code snippet
   - Environment details
   - Steps to reproduce

## Common Patterns

### Checking Market State

```tsx
const { data: market } = useMarket(marketId);
if (market?.status !== MarketStatus.Active) {
  return <div>Market is not active</div>;
}
```

### Handling Loading States

```tsx
const { data, isLoading, error } = useMarkets();

if (isLoading) return <LoadingSpinner />;
if (error) return <ErrorMessage error={error} />;
if (!data) return <div>No data</div>;

return <MarketList markets={data} />;
```

### Error Boundaries

```tsx
import { ErrorBoundary } from 'react-error-boundary';

<ErrorBoundary fallback={<ErrorFallback />}>
  <YourComponent />
</ErrorBoundary>
```

