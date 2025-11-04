# AI Resolver API Reference

## Table of Contents

- [Overview](#overview)
- [Base URL](#base-url)
- [Authentication](#authentication)
- [Common Response Codes](#common-response-codes)
- [Endpoints](#endpoints)
  - [Health Check](#health-check)
  - [Propose Market Resolution](#propose-market-resolution)
  - [List Pending Markets](#list-pending-markets)
- [Error Handling](#error-handling)
- [Rate Limiting](#rate-limiting)
- [Examples](#examples)

## Overview

The AI Resolver API provides HTTP endpoints for automated market resolution. All requests and responses use JSON format. The API is RESTful and follows standard HTTP conventions.

### API Version

Current version: **v1**

### Content Type

All requests must include:
```
Content-Type: application/json
```

All responses return:
```
Content-Type: application/json
```

## Base URL

**Development/Local**:
```
http://localhost:8080
```

**Production** (example):
```
https://api.horizonoracles.com/ai-resolver
```

## Authentication

Currently, the AI Resolver API does not require authentication for public endpoints. However, production deployments should implement:

- API key authentication
- Rate limiting per IP/key
- Request signing for sensitive operations

**Future Authentication Header**:
```
Authorization: Bearer <api_key>
```

## Common Response Codes

| Status Code | Meaning | Description |
|-------------|---------|-------------|
| **200** | OK | Request succeeded |
| **400** | Bad Request | Invalid request format or parameters |
| **401** | Unauthorized | Authentication required or failed |
| **404** | Not Found | Resource not found |
| **409** | Conflict | Resource state conflict (e.g., already resolved) |
| **429** | Too Many Requests | Rate limit exceeded |
| **500** | Internal Server Error | Server error occurred |
| **503** | Service Unavailable | Service temporarily unavailable |
| **504** | Gateway Timeout | Request took too long to process |

## Endpoints

### Health Check

Check the health status of the AI Resolver service.

#### Endpoint

```
GET /healthz
GET /v1/healthz
```

#### Request

No parameters required.

#### Response

**Success (200 OK)**:
```json
{
  "status": "healthy",
  "version": "1.0.0",
  "time": 1698765432,
  "signer": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chainId": 56
}
```

**Field Descriptions**:
- `status` (string): Service health status (`"healthy"` or `"unhealthy"`)
- `version` (string): Service version number
- `time` (number): Current Unix timestamp
- `signer` (string): Address of the signer account
- `chainId` (number): Blockchain network ID

**Error (503 Service Unavailable)**:
```json
{
  "status": "unhealthy",
  "error": "Unable to connect to blockchain",
  "time": 1698765432
}
```

#### Example

**cURL**:
```bash
curl http://localhost:8080/v1/healthz
```

**Response**:
```json
{
  "status": "healthy",
  "version": "1.0.0",
  "time": 1730563200,
  "signer": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chainId": 56
}
```

---

### Propose Market Resolution

Submit a market for AI-powered resolution. The service will analyze the question, gather evidence, and submit a resolution proposal to the blockchain.

#### Endpoint

```
POST /v1/propose
```

#### Request Headers

```
Content-Type: application/json
```

#### Request Body

```json
{
  "marketId": number,
  "closeTime": number,
  "question": string,
  "outcomeTokens": string[],
  "metadata": string (optional)
}
```

**Field Descriptions**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `marketId` | number | Yes | Unique market identifier | `123` |
| `closeTime` | number | Yes | Unix timestamp when market closes | `1762172000` |
| `question` | string | Yes | Market question to resolve | `"Will Bitcoin reach $100k by end of 2024?"` |
| `outcomeTokens` | string[] | Yes | Array of possible outcomes | `["YES", "NO"]` |
| `metadata` | string | No | Additional metadata as JSON string | `"{\"category\":\"crypto\"}"` |

**Constraints**:
- `marketId`: Must be a positive integer
- `closeTime`: Must be in the past (market must be closed)
- `question`: Non-empty string, max 500 characters
- `outcomeTokens`: Array of 2-10 strings, each max 50 characters
- `metadata`: Valid JSON string, max 1000 characters

#### Response

**Success (200 OK)**:
```json
{
  "status": "submitted",
  "marketId": 123,
  "outcomeId": 1,
  "confidence": 0.87,
  "reasoning": "Based on multiple credible sources including CoinDesk and Bloomberg, Bitcoin reached a peak of $98,500 in December 2024 but did not reach $100,000. The question asks if Bitcoin will reach $100k by end of 2024, which did not occur.",
  "txHash": "0xabc123def456...",
  "evidenceHash": "0xdef456abc789...",
  "citations": 5,
  "facts": 8,
  "toolCalls": 3,
  "iterations": 2
}
```

**Field Descriptions**:
- `status` (string): Proposal status (`"submitted"`, `"pending"`, `"failed"`)
- `marketId` (number): Market identifier (echoed from request)
- `outcomeId` (number): Proposed outcome index (0-based)
- `confidence` (number): Confidence score (0.0 to 1.0)
- `reasoning` (string): Detailed explanation of the decision
- `txHash` (string): Blockchain transaction hash
- `evidenceHash` (string): Hash of evidence data (IPFS CID in future)
- `citations` (number): Number of sources cited
- `facts` (number): Number of facts extracted
- `toolCalls` (number): Number of tools executed
- `iterations` (number): Number of LLM iterations

**Error Responses**:

**400 Bad Request** - Invalid request format:
```json
{
  "error": "Invalid request",
  "message": "marketId is required",
  "field": "marketId"
}
```

**409 Conflict** - Market already resolved:
```json
{
  "error": "Market already resolved",
  "message": "Market 123 has already been resolved with outcome 1",
  "marketId": 123,
  "existingOutcome": 1
}
```

**500 Internal Server Error** - Processing failed:
```json
{
  "error": "Resolution failed",
  "message": "Failed to process resolution: insufficient bond balance",
  "marketId": 123
}
```

**504 Gateway Timeout** - Processing took too long:
```json
{
  "error": "Request timeout",
  "message": "Resolution processing exceeded 5 minute timeout",
  "marketId": 123
}
```

#### Examples

**cURL - Basic Request**:
```bash
curl -X POST http://localhost:8080/v1/propose \
  -H "Content-Type: application/json" \
  -d '{
    "marketId": 123,
    "closeTime": 1762172000,
    "question": "Will Bitcoin reach $100k by end of 2024?",
    "outcomeTokens": ["YES", "NO"]
  }'
```

**cURL - With Metadata**:
```bash
curl -X POST http://localhost:8080/v1/propose \
  -H "Content-Type: application/json" \
  -d '{
    "marketId": 456,
    "closeTime": 1762172000,
    "question": "Will Ethereum switch to proof-of-stake in 2024?",
    "outcomeTokens": ["YES", "NO"],
    "metadata": "{\"category\":\"crypto\",\"tags\":[\"ethereum\",\"pos\"]}"
  }'
```

**JavaScript (fetch)**:
```javascript
const response = await fetch('http://localhost:8080/v1/propose', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    marketId: 123,
    closeTime: 1762172000,
    question: 'Will Bitcoin reach $100k by end of 2024?',
    outcomeTokens: ['YES', 'NO'],
  }),
});

const result = await response.json();
console.log('Resolution:', result);
```

**Python (requests)**:
```python
import requests

response = requests.post(
    'http://localhost:8080/v1/propose',
    json={
        'marketId': 123,
        'closeTime': 1762172000,
        'question': 'Will Bitcoin reach $100k by end of 2024?',
        'outcomeTokens': ['YES', 'NO'],
    }
)

result = response.json()
print(f"Outcome: {result['outcomeId']}")
print(f"Confidence: {result['confidence']}")
```

**Go**:
```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type ProposeRequest struct {
    MarketID      int      `json:"marketId"`
    CloseTime     int64    `json:"closeTime"`
    Question      string   `json:"question"`
    OutcomeTokens []string `json:"outcomeTokens"`
}

func main() {
    req := ProposeRequest{
        MarketID:      123,
        CloseTime:     1762172000,
        Question:      "Will Bitcoin reach $100k by end of 2024?",
        OutcomeTokens: []string{"YES", "NO"},
    }
    
    body, _ := json.Marshal(req)
    resp, _ := http.Post(
        "http://localhost:8080/v1/propose",
        "application/json",
        bytes.NewBuffer(body),
    )
    defer resp.Body.Close()
    
    var result map[string]any
    json.NewDecoder(resp.Body).Decode(&result)
}
```

---

### List Pending Markets

Retrieve a list of markets eligible for resolution.

#### Endpoint

```
GET /v1/markets
```

#### Query Parameters

| Parameter | Type | Required | Description | Example |
|-----------|------|----------|-------------|---------|
| `status` | string | No | Filter by market status | `"pending"` |
| `limit` | number | No | Maximum number of results | `10` |
| `offset` | number | No | Pagination offset | `0` |

#### Request

```
GET /v1/markets?status=pending&limit=10&offset=0
```

#### Response

**Success (200 OK)**:
```json
{
  "markets": [],
  "count": 0,
  "message": "Market listing not yet implemented"
}
```

**Future Response Format**:
```json
{
  "markets": [
    {
      "marketId": 123,
      "question": "Will Bitcoin reach $100k by end of 2024?",
      "closeTime": 1762172000,
      "outcomeTokens": ["YES", "NO"],
      "status": "closed",
      "eligible": true
    },
    {
      "marketId": 456,
      "question": "Will Ethereum switch to PoS?",
      "closeTime": 1762258400,
      "outcomeTokens": ["YES", "NO"],
      "status": "closed",
      "eligible": true
    }
  ],
  "count": 2,
  "total": 25,
  "limit": 10,
  "offset": 0
}
```

#### Example

**cURL**:
```bash
curl "http://localhost:8080/v1/markets?limit=10"
```

**JavaScript**:
```javascript
const response = await fetch('http://localhost:8080/v1/markets?limit=10');
const data = await response.json();
console.log(`Found ${data.count} markets`);
```

---

## Error Handling

### Error Response Format

All errors follow a consistent JSON format:

```json
{
  "error": "Error Type",
  "message": "Detailed error message",
  "field": "fieldName",
  "code": "ERROR_CODE"
}
```

### Error Codes

| Code | Description | Action |
|------|-------------|--------|
| `INVALID_REQUEST` | Request format is invalid | Check request body format |
| `MISSING_FIELD` | Required field is missing | Add the missing field |
| `INVALID_FIELD` | Field value is invalid | Correct the field value |
| `MARKET_NOT_FOUND` | Market does not exist | Verify market ID |
| `MARKET_OPEN` | Market is still open | Wait for market to close |
| `ALREADY_RESOLVED` | Market already resolved | Check current resolution |
| `INSUFFICIENT_BOND` | Not enough bond tokens | Add bond tokens to signer |
| `UNAUTHORIZED_SIGNER` | Signer not authorized | Check signer authorization |
| `LLM_ERROR` | AI processing failed | Retry request |
| `BLOCKCHAIN_ERROR` | Blockchain interaction failed | Check network status |
| `TIMEOUT` | Request timed out | Retry with simpler question |

### Example Error Responses

**Missing Required Field**:
```json
{
  "error": "Invalid request",
  "message": "marketId is required",
  "field": "marketId",
  "code": "MISSING_FIELD"
}
```

**Invalid Field Type**:
```json
{
  "error": "Invalid request",
  "message": "closeTime must be a number",
  "field": "closeTime",
  "code": "INVALID_FIELD"
}
```

**Market Not Closed**:
```json
{
  "error": "Market not eligible",
  "message": "Market 123 is still open (closes at 1762172000)",
  "field": "marketId",
  "code": "MARKET_OPEN",
  "closeTime": 1762172000,
  "currentTime": 1762000000
}
```

**Insufficient Bond**:
```json
{
  "error": "Insufficient funds",
  "message": "Signer has 500 HORIZON but needs 1000 HORIZON for bond",
  "code": "INSUFFICIENT_BOND",
  "required": "1000000000000000000000",
  "available": "500000000000000000000"
}
```

## Rate Limiting

### Current Limits

**Development**:
- No rate limits

**Production** (recommended):
- 100 requests per hour per IP
- 1000 requests per day per IP
- Burst allowance: 10 requests

### Rate Limit Headers

```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1698768000
```

### Rate Limit Exceeded Response

```json
{
  "error": "Rate limit exceeded",
  "message": "You have exceeded 100 requests per hour",
  "code": "RATE_LIMIT_EXCEEDED",
  "retryAfter": 3600
}
```

## Request/Response Examples

### Complete Request Flow

**1. Check Service Health**:
```bash
curl http://localhost:8080/v1/healthz
```

Response:
```json
{
  "status": "healthy",
  "version": "1.0.0",
  "time": 1730563200,
  "signer": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chainId": 56
}
```

**2. Submit Resolution Request**:
```bash
curl -X POST http://localhost:8080/v1/propose \
  -H "Content-Type: application/json" \
  -d '{
    "marketId": 1,
    "closeTime": 1762172000,
    "question": "Will Bitcoin reach $100k by December 31, 2024?",
    "outcomeTokens": ["YES", "NO"]
  }'
```

Response:
```json
{
  "status": "submitted",
  "marketId": 1,
  "outcomeId": 0,
  "confidence": 0.92,
  "reasoning": "Based on comprehensive analysis of market data from multiple sources including CoinMarketCap, CoinGecko, and major exchanges, Bitcoin's price on December 31, 2024 was $98,750. While Bitcoin came close to the $100k milestone, reaching a peak of $98,950 on December 28, it did not reach or exceed $100,000 by the specified deadline. The question specifically asks if Bitcoin will reach $100k by December 31, 2024, which did not occur. Therefore, the answer is NO.",
  "txHash": "0xabc123def456789...",
  "evidenceHash": "0xdef456abc789012...",
  "citations": 7,
  "facts": 12,
  "toolCalls": 2,
  "iterations": 1
}
```

**3. Verify Transaction on Blockchain**:
```bash
# Using cast (Foundry)
cast tx 0xabc123def456789... --rpc-url https://bsc-dataseed.binance.org

# Or view in block explorer
# https://bscscan.com/tx/0xabc123def456789...
```

### Advanced Examples

**Multi-Tool Resolution**:
```bash
curl -X POST http://localhost:8080/v1/propose \
  -H "Content-Type: application/json" \
  -d '{
    "marketId": 2,
    "closeTime": 1762172000,
    "question": "Use get_market_data to get market 1 close_time, then use datetime to convert it. What date did market 1 close?",
    "outcomeTokens": ["2024-11-03", "2024-11-04", "2024-11-05"]
  }'
```

Response shows tool execution:
```json
{
  "status": "submitted",
  "marketId": 2,
  "outcomeId": 0,
  "confidence": 1.0,
  "reasoning": "Used get_market_data tool to fetch market 1 data, which returned close_time: 1762172000. Then used datetime tool to convert this timestamp to human-readable date: 2024-11-03 12:13:20 UTC. Therefore, market 1 closed on 2024-11-03.",
  "txHash": "0x123abc...",
  "evidenceHash": "0x456def...",
  "toolCalls": 2,
  "iterations": 2
}
```

## Best Practices

### 1. Request Formatting

✅ **Do**:
- Use proper JSON formatting
- Include all required fields
- Validate data types before sending
- Use descriptive, clear questions

❌ **Don't**:
- Send malformed JSON
- Include unnecessary fields
- Use ambiguous questions
- Send duplicate requests

### 2. Error Handling

✅ **Do**:
- Check HTTP status codes
- Parse error messages
- Implement retry logic for 5xx errors
- Log errors for debugging

❌ **Don't**:
- Ignore error responses
- Retry on 4xx errors without fixing the request
- Assume success without checking status
- Expose error details to end users

### 3. Performance

✅ **Do**:
- Batch multiple markets if possible (future)
- Cache health check results
- Use connection pooling
- Set appropriate timeouts

❌ **Don't**:
- Poll the API excessively
- Send concurrent requests for the same market
- Set unrealistic timeouts (< 30s)

### 4. Security

✅ **Do**:
- Use HTTPS in production
- Validate all input on client side
- Implement API key authentication
- Rate limit your own requests

❌ **Don't**:
- Send sensitive data in URLs
- Trust user input without validation
- Hardcode API URLs
- Expose API keys in client-side code

## Changelog

### v1.0.0 (2024-11-02)
- Initial API release
- `/v1/propose` endpoint for market resolution
- `/v1/healthz` health check endpoint
- `/v1/markets` endpoint (placeholder)
- Support for binary markets
- Tool orchestration system
- EIP-712 signature generation

### Upcoming Features

**v1.1.0** (Planned):
- `/v1/markets` full implementation
- Pagination support
- Market filtering by category
- Bulk resolution submission

**v1.2.0** (Planned):
- Multi-outcome market support
- Confidence threshold configuration
- Custom bond amounts
- Webhook notifications

**v2.0.0** (Future):
- WebSocket support for real-time updates
- GraphQL API
- Advanced filtering and search
- Historical resolution data
