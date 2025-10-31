import { defineConfig } from 'tsup';

export default defineConfig({
  entry: ['src/index.ts'],
  format: ['cjs', 'esm'],
  dts: true,
  splitting: false,
  sourcemap: true,
  clean: true,
  // Externalize all dependencies - they should be provided by the consuming app
  // Must use array format (not function) for tsup worker serialization
  external: [
    'react',
    'react-dom',
    'react/jsx-runtime',
    'wagmi',
    'viem',
    '@tanstack/react-query',
  ],
  // Ensure no code evaluation happens during import
  treeshake: {
    preset: 'smallest',
    moduleSideEffects: (id) => {
      // Mark all hook files as having side effects to prevent premature evaluation
      // This ensures wagmi hooks are only called when the SDK hooks are actually invoked
      if (id.includes('/hooks/') || id.includes('/components/')) {
        return true;
      }
      return false;
    },
  },
  // Prevent any code from executing during import
  esbuildOptions(options) {
    options.bundle = true;
    options.format = options.format || 'esm';
    return options;
  },
});

