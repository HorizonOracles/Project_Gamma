/**
 * Test setup file for Vitest
 */

import { afterEach, vi } from 'vitest';
import { cleanup } from '@testing-library/react';
import '@testing-library/jest-dom';

// BigInt serialization support for JSON.stringify in tests
(BigInt.prototype as any).toJSON = function() {
  return this.toString();
};

// Cleanup after each test
afterEach(() => {
  cleanup();
});

// Mock window.ethereum
Object.defineProperty(window, 'ethereum', {
  value: {
    request: vi.fn(),
    on: vi.fn(),
    removeListener: vi.fn(),
  },
  writable: true,
});

// Mock process.env
if (typeof process !== 'undefined') {
  process.env = {
    ...process.env,
    CHAIN_ID: '56',
    RPC_URL: 'https://bsc-dataseed.binance.org/',
  };
}

