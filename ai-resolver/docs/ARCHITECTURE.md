# AI Resolver Architecture

## Table of Contents

- [Overview](#overview)
- [System Components](#system-components)
- [Resolution Pipeline](#resolution-pipeline)
- [Tool Orchestration](#tool-orchestration)
- [Data Flow](#data-flow)
- [Security Model](#security-model)
- [Performance Considerations](#performance-considerations)

## Overview

The AI Resolver is a sophisticated backend service that automates the resolution of binary prediction markets using AI analysis, blockchain data, and cryptographic verification. It combines OpenAI's GPT-4 language model with a custom tool orchestration system to make evidence-based decisions about market outcomes.

### Key Design Principles

1. **Evidence-Based Resolution**: All decisions must be backed by verifiable data sources
2. **Transparent Reasoning**: Every resolution includes detailed reasoning and citations
3. **Cryptographic Verification**: All proposals are signed using EIP-712 for on-chain verification
4. **Extensible Architecture**: New tools and data sources can be easily added
5. **Fault Tolerance**: Graceful handling of errors and edge cases

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         AI Resolver Service                      │
│                                                                  │
│  ┌────────────────┐      ┌──────────────┐    ┌───────────────┐ │
│  │   HTTP Server  │──────│   Pipeline   │────│  Tool Registry│ │
│  │   (Port 8080)  │      │   Handler    │    │               │ │
│  └────────────────┘      └──────────────┘    └───────────────┘ │
│                                  │                      │        │
│                          ┌───────┴──────────┐         │        │
│                          │                  │          │        │
│                    ┌─────▼────┐      ┌─────▼────┐    │        │
│                    │   LLM    │      │  Tools   │◄───┘        │
│                    │ Provider │      │          │              │
│                    └─────┬────┘      └─────┬────┘              │
│                          │                  │                   │
│                    ┌─────▼────────────────┬─▼────┐             │
│                    │   OpenAI API         │ Data │             │
│                    │   (GPT-4)            │Source│             │
│                    └──────────────────────┴──────┘             │
│                                  │                              │
│                          ┌───────▼────────┐                    │
│                          │   EIP-712      │                    │
│                          │   Signer       │                    │
│                          └───────┬────────┘                    │
└──────────────────────────────────┼─────────────────────────────┘
                                   │
                           ┌───────▼────────┐
                           │   Blockchain   │
                           │  (BSC/Ethereum)│
                           └────────────────┘
```

## System Components

### 1. HTTP Server (`cmd/server/main.go`)

**Purpose**: Handles incoming HTTP requests and routes them to appropriate handlers.

**Responsibilities**:
- Accept POST requests to `/v1/propose` endpoint
- Validate request payloads
- Initialize and coordinate pipeline execution
- Return formatted responses to clients
- Health check endpoints

**Key Features**:
- CORS handling for cross-origin requests
- Request timeout management (default: 5 minutes)
- Structured error responses
- Request/response logging

### 2. Configuration Management (`internal/config/`)

**Purpose**: Centralized configuration loading and validation.

**Configuration Sources**:
1. Environment variables (primary)
2. `.env` file (development)
3. Default values (fallback)

**Configuration Categories**:
- **Server**: Port, host, logging level
- **Blockchain**: RPC endpoint, contract addresses, chain ID
- **AI**: OpenAI API key, model selection
- **Security**: Signer configuration (local key or KMS)
- **Operational**: Timeouts, concurrency limits

**Validation**:
- Required fields are checked at startup
- Invalid configurations cause immediate failure
- Clear error messages for missing/invalid values

### 3. LLM Provider (`internal/llm/`)

**Purpose**: Interface with OpenAI's API for intelligent decision-making.

**Components**:

#### `openai.go` - OpenAI API Client
- Handles communication with OpenAI's Responses API
- Manages API authentication and rate limiting
- Supports tool orchestration through function calling
- Implements retry logic for transient failures

#### `pipeline.go` - Resolution Pipeline
- Multi-pass analysis pipeline (future enhancement)
- Fact extraction from sources
- Contradiction detection
- Confidence scoring
- Citation generation

**Tool Integration**:
The LLM provider integrates with the tool registry to enable:
- Dynamic tool discovery
- Automatic tool schema injection
- Function call parsing and execution
- Stateless context management through prompt appending

**Request Flow**:
```
1. Receive market question + tool definitions
2. Send to OpenAI with tools array
3. Parse response (text or tool_calls)
4. If tool_calls: execute tools and append results
5. Repeat until final answer or max iterations
6. Return final decision with reasoning
```

### 4. Tool Registry (`internal/tools/`)

**Purpose**: Manage the collection of tools available to the LLM.

**Architecture**:

```
┌─────────────────────────────────────────────────┐
│              Tool Registry                       │
│                                                  │
│  ┌──────────────────────────────────────────┐  │
│  │        Tool Interface                     │  │
│  │  - Name() string                          │  │
│  │  - Description() string                   │  │
│  │  - Execute(ctx, input) (output, error)   │  │
│  │  - ToOpenAIFormat() map[string]any       │  │
│  └──────────────────────────────────────────┘  │
│                     △                            │
│      ┌──────────────┼──────────────┐           │
│      │              │               │            │
│  ┌───▼────┐   ┌────▼─────┐   ┌────▼────┐      │
│  │ Market │   │Calculator│   │DateTime │      │
│  │  Data  │   │          │   │         │      │
│  └────────┘   └──────────┘   └─────────┘      │
│                                                  │
│  ┌─────────┐   ┌───────────┐                   │
│  │BSCScan  │   │PancakeSwap│                   │
│  │         │   │           │                   │
│  └─────────┘   └───────────┘                   │
└─────────────────────────────────────────────────┘
```

**Tool Types**:
1. **Function Tools**: Structured JSON input/output
2. **Custom Tools**: Raw text input/output
3. **Web Search Tools**: OpenAI's native web search

**Built-in Tools**:
- `get_market_data`: Fetch on-chain market information
- `calculate`: Mathematical operations (add, subtract, multiply, divide, etc.)
- `datetime`: Unix timestamp conversions
- `bscscan`: Query BSCScan API for blockchain data
- `pancakeswap`: Fetch DEX prices and liquidity data

**Tool Lifecycle**:
```
1. Tool Registration (startup)
   ├── Validate tool interface
   ├── Check for name conflicts
   └── Add to registry

2. Tool Discovery (request)
   ├── List all available tools
   └── Format for OpenAI API

3. Tool Execution (runtime)
   ├── Validate input against schema
   ├── Execute with timeout
   ├── Handle errors gracefully
   └── Return formatted output

4. Result Processing
   ├── Append to conversation context
   └── Continue LLM iteration
```

### 5. EIP-712 Signer (`internal/eip712/`)

**Purpose**: Generate cryptographic signatures for market resolution proposals.

**Signing Process**:
```
1. Construct typed data structure
   ├── Domain separator (verifying contract, chain ID)
   ├── Market data (ID, outcome, timestamp)
   └── Evidence hash (IPFS CID)

2. Hash the structured data
   ├── Type hash (keccak256 of type string)
   ├── Encode values (abi.encode)
   └── EIP-712 hash (domain + message)

3. Sign with private key
   ├── ECDSA signature generation
   ├── v, r, s components
   └── Recoverable signature

4. Return signature
   └── 65-byte signature (r + s + v)
```

**Security Considerations**:
- Private keys never leave the process
- Signatures are single-use (nonce + timestamp)
- Domain separation prevents cross-contract replay
- KMS support for production environments

### 6. Blockchain Adapter (`internal/adapter/`)

**Purpose**: Interface with Ethereum-compatible blockchains.

**Responsibilities**:
- Query contract state (market data, balances)
- Submit transactions (resolution proposals)
- Monitor events (resolution status changes)
- Gas price estimation and management

**Contract Interactions**:
```
AIOracleAdapter.proposeResolution(
    marketId,
    outcomeId,
    evidenceHash,
    signature,
    bondAmount
)
```

**State Queries**:
- Market information from MarketFactory
- Resolution status from ResolutionModule
- Token balances and allowances
- Signer authorization status

## Resolution Pipeline

### Pipeline Stages

#### Stage 1: Request Validation
```
Input: POST /v1/propose
├── Validate JSON structure
├── Check required fields
├── Verify data types
└── Return 400 if invalid
```

#### Stage 2: Context Preparation
```
Prepare:
├── Market metadata
├── Question text
├── Outcome options
├── Close time
└── Tool definitions
```

#### Stage 3: LLM Analysis
```
For i = 1 to MAX_ITERATIONS (10):
├── Send prompt + tools to OpenAI
├── Parse response
│   ├── If tool_calls: execute tools
│   │   ├── Validate arguments
│   │   ├── Execute with timeout
│   │   ├── Append results to context
│   │   └── Continue to next iteration
│   └── If text response: extract decision
│       ├── Parse outcome choice
│       ├── Extract confidence score
│       ├── Collect reasoning
│       └── Exit loop
└── Return final decision
```

#### Stage 4: Signature Generation
```
Create EIP-712 signature:
├── Hash evidence (reasoning + citations)
├── Upload to IPFS (future)
├── Construct typed data
├── Sign with private key
└── Verify signature locally
```

#### Stage 5: On-Chain Submission
```
Submit to blockchain:
├── Estimate gas
├── Check token balance/allowance
├── Send transaction
├── Wait for confirmation
└── Return transaction hash
```

### Error Handling

**Error Categories**:
1. **Validation Errors** (400): Invalid request format
2. **Authentication Errors** (401): Invalid API keys
3. **Resource Errors** (404): Market not found
4. **Conflict Errors** (409): Market already resolved
5. **Processing Errors** (500): Internal failures
6. **Timeout Errors** (504): Exceeded time limit

**Recovery Strategies**:
- **Transient Failures**: Retry with exponential backoff
- **Tool Failures**: Skip tool, continue with available data
- **LLM Failures**: Fall back to simpler prompt
- **Blockchain Failures**: Queue for later retry

## Tool Orchestration

### Orchestration Flow

```
┌──────────────────────────────────────────────────────┐
│  LLM receives question + tool definitions             │
└─────────────────┬────────────────────────────────────┘
                  │
     ┌────────────▼─────────────┐
     │ Needs tools?              │
     └────┬─────────────┬────────┘
          │ Yes         │ No
    ┌─────▼──────┐      │
    │ tool_calls │      │
    └─────┬──────┘      │
          │             │
    ┌─────▼──────────┐  │
    │ Execute tools  │  │
    │ in sequence    │  │
    └─────┬──────────┘  │
          │             │
    ┌─────▼──────────┐  │
    │ Append results │  │
    │ to context     │  │
    └─────┬──────────┘  │
          │             │
          └─────┬───────┘
                │
     ┌──────────▼─────────────┐
     │ Send updated context   │
     │ back to LLM            │
     └────────────────────────┘
```

### Multi-Tool Chaining

**Example: Complex Temporal Query**
```
Question: "Is the close time exactly 24 hours after market creation?"

Step 1: get_market_data(market_id: 1)
→ Returns: {created_at: 1700000000, close_time: 1700086400}

Step 2: calculate(subtract([1700086400, 1700000000]))
→ Returns: 86400

Step 3: calculate(equals([86400, 86400]))
→ Returns: true

Step 4: LLM synthesizes final answer
→ "YES, the close time is exactly 24 hours (86400 seconds) after creation"
```

### Stateless Context Management

Unlike native tool calling frameworks, the AI Resolver uses **stateless context passing**:

```
Iteration 1:
Prompt: "Question: [X]. Tools: [T1, T2, T3]"
Response: tool_calls = [call_T1]

Iteration 2:
Prompt: "Question: [X]. Tools: [T1, T2, T3]
         Previous call: T1(args) = result"
Response: tool_calls = [call_T2]

Iteration 3:
Prompt: "Question: [X]. Tools: [T1, T2, T3]
         Previous call: T1(args) = result
         Previous call: T2(args) = result"
Response: "Based on the results, the answer is [Y]"
```

This approach:
- ✅ Works with OpenAI's Responses API
- ✅ Provides full conversation history
- ✅ Enables complex multi-step reasoning
- ✅ Allows error recovery and retries

## Data Flow

### Request Processing Flow

```
1. HTTP Request arrives
   └─> JSON payload parsed

2. Request validation
   ├─> Check required fields
   ├─> Validate data types
   └─> Return 400 if invalid

3. Initialize pipeline
   ├─> Load configuration
   ├─> Create LLM provider
   ├─> Load tool registry
   └─> Initialize signer

4. Execute resolution
   ├─> Send to LLM with tools
   ├─> Handle tool calls
   ├─> Collect reasoning
   └─> Make decision

5. Generate signature
   ├─> Hash evidence
   ├─> Create typed data
   └─> Sign with private key

6. Submit to blockchain
   ├─> Call proposeResolution()
   ├─> Wait for confirmation
   └─> Return transaction hash

7. Return response
   └─> JSON response with result
```

### Data Structures

**Request Structure**:
```json
{
  "marketId": 123,
  "closeTime": 1762172000,
  "question": "Will Bitcoin reach $100k by end of 2024?",
  "outcomeTokens": ["YES", "NO"],
  "metadata": "{}"
}
```

**Response Structure**:
```json
{
  "status": "submitted",
  "marketId": 123,
  "outcomeId": 1,
  "confidence": 0.87,
  "reasoning": "Based on current market data...",
  "txHash": "0xabc123...",
  "evidenceHash": "0xdef456...",
  "citations": 5,
  "facts": 8
}
```

**Tool Input Structure**:
```go
type ToolInput struct {
    CallID     string         // Unique call identifier
    Name       string         // Tool name
    Arguments  map[string]any // Structured arguments (function tools)
    RawInput   string         // Raw input (custom tools)
}
```

**Tool Output Structure**:
```go
type ToolOutput struct {
    CallID  string // Matching call identifier
    Content string // Human-readable output
    Data    any    // Structured data (optional)
    Error   error  // Error if execution failed
}
```

## Security Model

### Authentication & Authorization

**API Level**:
- No authentication required for `/v1/propose` (public service)
- Rate limiting recommended for production
- CORS policies configurable per environment

**Blockchain Level**:
- Signer must be authorized in AIOracleAdapter contract
- Only authorized signers can propose resolutions
- Bond requirement prevents spam proposals

### Cryptographic Security

**EIP-712 Typed Signatures**:
```
Domain Separator:
├── name: "HorizonOracles"
├── version: "1"
├── chainId: 56 (or current chain)
└── verifyingContract: AIOracleAdapter address

Message Types:
├── Resolution(uint256 marketId, uint256 outcomeId, ...)
└── Evidence(string evidenceHash, uint256 timestamp, ...)

Signature Components:
├── r: 32 bytes (signature R)
├── s: 32 bytes (signature S)
└── v: 1 byte (recovery ID)
```

**Key Management**:
- **Development**: Private key in environment variable
- **Production**: AWS KMS or similar HSM
- **Rotation**: Automated key rotation support
- **Backup**: Multi-sig backup keys

### Attack Vectors & Mitigations

| Attack Vector | Mitigation |
|---------------|------------|
| **Signature Replay** | Nonce + timestamp in typed data |
| **Cross-Contract Replay** | Domain separator with contract address |
| **Unauthorized Proposer** | On-chain authorization check |
| **Insufficient Bond** | Balance verification before proposal |
| **Gas Price Manipulation** | Gas estimation + 20% buffer |
| **Front-Running** | Private mempool (future) |
| **Data Poisoning** | Multiple data source verification |
| **Prompt Injection** | Input sanitization, structured prompts |

## Performance Considerations

### Latency Optimization

**Expected Latencies**:
- Request validation: < 10ms
- Tool execution: 100-500ms per tool
- LLM inference: 2-10s per call
- EIP-712 signing: < 5ms
- Transaction submission: 1-30s (depending on network)

**Total Expected Time**: 10-60 seconds per resolution

**Optimization Strategies**:
1. **Parallel Tool Execution**: Execute independent tools concurrently
2. **Caching**: Cache blockchain data for repeated queries
3. **Streaming Responses**: Stream LLM responses as they arrive
4. **Connection Pooling**: Reuse HTTP connections to OpenAI/RPC
5. **Batch Processing**: Process multiple markets in parallel

### Resource Management

**Memory**:
- LLM context: ~4KB per request
- Tool registry: ~100KB
- HTTP server: ~10MB base
- Total: < 50MB for typical workload

**CPU**:
- Mostly I/O bound (waiting on APIs)
- Signature generation: < 1% CPU
- JSON parsing: < 5% CPU

**Network**:
- OpenAI API: 10-100KB per request
- Blockchain RPC: 1-10KB per query
- Total: < 1MB per resolution

### Scalability

**Horizontal Scaling**:
- Stateless service (scales linearly)
- Load balancer distributes requests
- Shared database for coordination (future)

**Vertical Scaling**:
- CPU: 2-4 cores sufficient for 100s of resolutions/day
- Memory: 512MB-1GB sufficient
- Network: Standard bandwidth sufficient

**Bottlenecks**:
1. **OpenAI API**: Rate limits (3000 RPM for GPT-4)
2. **Blockchain RPC**: Rate limits (varies by provider)
3. **Signer Account**: Bond token availability

## Monitoring & Observability

### Key Metrics

**Service Health**:
- Request rate (requests/second)
- Success rate (%)
- Error rate (%)
- P50/P95/P99 latency (seconds)

**LLM Metrics**:
- Token usage (input/output)
- Average iterations per resolution
- Tool usage frequency
- Confidence score distribution

**Blockchain Metrics**:
- Transaction success rate
- Average gas used
- Bond balance remaining
- Proposal acceptance rate

### Logging

**Log Levels**:
- **DEBUG**: Tool execution details, API payloads
- **INFO**: Request processing, successful resolutions
- **WARN**: Retries, degraded performance
- **ERROR**: Failed resolutions, critical errors

**Structured Logging Format**:
```json
{
  "timestamp": "2025-11-02T12:34:56Z",
  "level": "info",
  "message": "Resolution submitted",
  "marketId": 123,
  "outcomeId": 1,
  "confidence": 0.87,
  "txHash": "0xabc...",
  "duration_ms": 12340
}
```

## Future Enhancements

### Short-term (Q1 2025)
- [ ] Health endpoint implementation
- [ ] Prometheus metrics export
- [ ] Request queuing system
- [ ] Automated market discovery
- [ ] Multi-pass analysis pipeline

### Medium-term (Q2 2025)
- [ ] Support for multi-outcome markets
- [ ] Advanced evidence aggregation
- [ ] Dispute resolution participation
- [ ] Real-time market monitoring
- [ ] Confidence threshold configuration

### Long-term (Q3+ 2025)
- [ ] Custom LLM fine-tuning
- [ ] Decentralized oracle network
- [ ] Multi-chain support
- [ ] Advanced cryptoeconomic mechanisms
- [ ] DAO-governed resolution rules
