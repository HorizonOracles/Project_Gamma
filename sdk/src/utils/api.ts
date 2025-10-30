/**
 * API client for Project Gamma Oracle backend
 * Handles all HTTP requests to the public API
 */

import { OracleRequest, OracleResult } from '../types';

export interface RequestResolutionParams {
  marketId: number;
  metadata: {
    question: string;
    description?: string;
  };
}

export interface RequestResolutionResponse {
  requestId: string;
  marketId: number;
  status: 'pending';
}

/**
 * Creates an API client for oracle requests
 */
export function createOracleApiClient(baseUrl: string) {
  if (!baseUrl) {
    throw new Error('Oracle API URL is required');
  }

  // Remove trailing slash
  const apiUrl = baseUrl.replace(/\/$/, '');

  return {
    /**
     * Request AI resolution for a market
     */
    async requestResolution(
      params: RequestResolutionParams
    ): Promise<RequestResolutionResponse> {
      const response = await fetch(`${apiUrl}/api/v1/oracle/request`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          marketId: params.marketId,
          metadata: params.metadata,
        }),
      });

      if (!response.ok) {
        const error = await response.json().catch(() => ({
          message: `HTTP ${response.status}: ${response.statusText}`,
        }));
        throw new Error(error.message || 'Failed to request resolution');
      }

      return response.json();
    },

    /**
     * Get oracle request status
     */
    async getRequestStatus(requestId: string): Promise<OracleRequest> {
      const response = await fetch(`${apiUrl}/api/v1/oracle/request/${requestId}`);

      if (!response.ok) {
        const error = await response.json().catch(() => ({
          message: `HTTP ${response.status}: ${response.statusText}`,
        }));
        throw new Error(error.message || 'Failed to get request status');
      }

      return response.json();
    },

    /**
     * Get oracle result for a completed request
     */
    async getRequestResult(requestId: string): Promise<OracleResult> {
      const response = await fetch(`${apiUrl}/api/v1/oracle/result/${requestId}`);

      if (!response.ok) {
        const error = await response.json().catch(() => ({
          message: `HTTP ${response.status}: ${response.statusText}`,
        }));
        throw new Error(error.message || 'Failed to get result');
      }

      return response.json();
    },

    /**
     * Get oracle history for a market
     */
    async getMarketHistory(marketId: number): Promise<OracleRequest[]> {
      const response = await fetch(`${apiUrl}/api/v1/oracle/market/${marketId}/history`);

      if (!response.ok) {
        const error = await response.json().catch(() => ({
          message: `HTTP ${response.status}: ${response.statusText}`,
        }));
        throw new Error(error.message || 'Failed to get market history');
      }

      return response.json();
    },
  };
}

/**
 * Default API client instance
 * Should be initialized with useGammaConfig()
 */
export type OracleApiClient = ReturnType<typeof createOracleApiClient>;

