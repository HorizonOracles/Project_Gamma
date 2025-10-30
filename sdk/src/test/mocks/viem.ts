/**
 * Mock implementations for viem
 */

import { vi } from 'vitest';
import { Address } from 'viem';

export const mockPublicClient = {
  readContract: vi.fn(),
  writeContract: vi.fn(),
  waitForTransactionReceipt: vi.fn(),
  getLogs: vi.fn(),
  getBlockNumber: vi.fn(),
  getBlock: vi.fn(),
};

export const mockWalletClient = {
  writeContract: vi.fn(),
  getAddresses: vi.fn(),
  getAddress: vi.fn(),
};

export const mockAddress: Address = '0x1234567890123456789012345678901234567890' as Address;
export const mockMarketId = 1n;
export const mockMarketAddress: Address = '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd' as Address;
export const mockTransactionHash = '0xabcd1234abcd1234abcd1234abcd1234abcd1234abcd1234abcd1234abcd1234';

// Helper to create mock transaction receipt
export function createMockReceipt(logs: any[] = []) {
  return {
    transactionHash: mockTransactionHash,
    status: 'success',
    logs,
    blockNumber: 100n,
    blockHash: '0x' + '0'.repeat(64),
  };
}

// Helper to create mock event log
export function createMockEventLog(
  eventName: string,
  args: Record<string, any>,
  address: Address = mockAddress
) {
  return {
    address,
    topics: [`0x${eventName.slice(0, 64)}`],
    data: '0x' + '0'.repeat(64),
    blockNumber: 100n,
    transactionHash: mockTransactionHash,
    logIndex: 0,
  };
}

