# Installation

Installation instructions for the Project Gamma React SDK.

## Install Package

```bash
npm install @project-gamma/react-sdk
# or
yarn add @project-gamma/react-sdk
# or
pnpm add @project-gamma/react-sdk
```

## Peer Dependencies

The SDK requires the following peer dependencies. Install them if not already present:

```bash
npm install react react-dom wagmi viem @tanstack/react-query
```

### Version Requirements

- **React**: ^18.0.0
- **Wagmi**: ^2.0.0
- **Viem**: ^2.0.0
- **@tanstack/react-query**: ^5.0.0

## TypeScript Support

The SDK is written in TypeScript and includes full type definitions. No additional type packages are required.

## Verification

After installation, verify it works:

```tsx
import { GammaProvider } from '@project-gamma/react-sdk';

// If this imports without errors, installation was successful
```

## Build Tools

The SDK works with all modern build tools:

- **Vite** ✅
- **Create React App** ✅
- **Next.js** ✅
- **Webpack** ✅
- **Rollup** ✅

## Bundle Size

The SDK is optimized for size:
- Core bundle: ~50kb gzipped
- Tree-shakeable exports
- No unnecessary dependencies

## Troubleshooting

If you encounter installation issues:

1. **Clear node_modules and reinstall**:
   ```bash
   rm -rf node_modules package-lock.json
   npm install
   ```

2. **Check Node.js version**: Requires Node.js 18+

3. **Verify peer dependencies**: Ensure all peer dependencies are installed

4. See [Troubleshooting](./troubleshooting.md) for more help

