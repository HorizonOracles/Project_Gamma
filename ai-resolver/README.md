<div align="center">

# AI Resolver Service

**Automated AI-Powered Market Resolution**

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![OpenAI](https://img.shields.io/badge/OpenAI-GPT--4-412991?style=flat-square)](https://openai.com/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)
[![EIP-712](https://img.shields.io/badge/EIP--712-Signatures-green?style=flat-square)]()

*Backend service for evidence-based outcome determination on Gamma prediction markets*

---

[Architecture](#architecture) • [Setup](#setup) • [API Reference](#api-endpoints) • [Deployment](#deployment) • [Troubleshooting](#troubleshooting)

---

</div>

## Overview

The AI Resolver implements a complete automated pipeline for resolving binary prediction markets using AI analysis, web search, and cryptographic verification. All resolution proposals are signed using EIP-712 and verified on-chain.

### Resolution Pipeline

<table>
<tr>
<td width="20%" align="center">

**1. Web Search**

Query multiple search providers for evidence

</td>
<td width="20%" align="center">

**2. LLM Analysis**

4-step multi-pass pipeline using GPT-4

</td>
<td width="20%" align="center">

**3. EIP-712 Signing**

Cryptographic signature generation

</td>
<td width="20%" align="center">

**4. On-Chain Submission**

Submit to AIOracleAdapter contract

</td>
<td width="20%" align="center">

**5. Dispute Window**

48-hour community review period

</td>
</tr>
</table>

### Key Features

<table>
<tr>
<td width="33%">

**Multi-Pass Analysis**
- Extract facts from sources
- Check contradictions
- Determine confidence
- Build citations

</td>
<td width="33%">

**Multiple Search Providers**
- Serper.dev (recommended)
- Google Search API
- Brave Search API
- Configurable priority

</td>
<td width="33%">

**Cryptographic Security**
- EIP-712 typed signatures
- On-chain verification
- Replay protection
- Time-bounded validity

</td>
</tr>
</table>

---

## Architecture

### System Components

```
AI Resolver Service
│
├── HTTP API Server (Port 8080)
│   ├── Health check endpoints
│   ├── Proposal submission
│   └── Market listing
│
├── Resolution Pipeline
│   ├── Web Search Module
│   │   ├── Serper.dev provider
│   │   ├── Google Search provider
│   │   └── Brave Search provider
│   │
│   ├── LLM Analysis (4-step)
│   │   ├── 1. Fact extraction
│   │   ├── 2. Contradiction checking
│   │   ├── 3. Outcome decision
│   │   └── 4. Citation building
│   │
│   └── EIP-712 Signing
│       └── ECDSA signature generation
│
└── Blockchain Integration
    ├── Contract clients (abigen)
    ├── Transaction submission
    └── Event monitoring
```

### Project Structure

```
ai-resolver/
│
├── cmd/
│   └── server/              HTTP server entry point
│
├── internal/
│   ├── config/             Configuration management
│   ├── llm/                OpenAI multi-pass pipeline
│   ├── search/             Web search providers
│   ├── eip712/             EIP-712 signing utilities
│   └── adapter/            Ethereum contract client
│
├── pkg/
│   └── abi/                Generated Go bindings (abigen)
│
├── configs/                Configuration files
├── bin/                    Compiled binaries
└── docker-compose.yml      Local development stack
```

---

## Setup

### Prerequisites

<table>
<tr>
<td width="50%">

**Development Tools**

- **Go 1.24+** - Programming language
- **Docker** - Container runtime (optional)
- **Foundry** - For contract bindings generation

</td>
<td width="50%">

**External Services**

- **Ethereum RPC** - BSC, Ethereum, or Base node
- **OpenAI API Key** - GPT-4 access
- **Search API Key** - Serper.dev recommended
- **HORIZON Tokens** - For proposal bonding

</td>
</tr>
</table>

### Installation

**Step 1: Clone and Configure**

```bash
# Clone repository
git clone https://github.com/HorizonOracles/Project_Gamma.git
cd Project_Gamma/ai-resolver

# Copy environment template
cp .env.example .env
```

**Step 2: Set Environment Variables**

Edit `.env` with your configuration:

<details>
<summary><strong>Blockchain Configuration</strong></summary>

```env
# Network Settings
RPC_ENDPOINT=https://bsc-dataseed.binance.org
CHAIN_ID=56

# Signer Account (AI Oracle Signer)
SIGNER_PRIVATE_KEY=your_private_key_here
```

**Supported Networks:**
- BSC Mainnet: Chain ID 56
- BSC Testnet: Chain ID 97
- Ethereum: Chain ID 1
- Sepolia: Chain ID 11155111
- Base: Chain ID 8453

</details>

<details>
<summary><strong>Contract Addresses</strong></summary>

```env
# Deployed Contract Addresses (from deployment)
AI_ORACLE_ADAPTER_ADDR=0x...
MARKET_FACTORY_ADDR=0x...
RESOLUTION_MODULE_ADDR=0x...
HORIZON_TOKEN_ADDR=0x...
```

Get these addresses from your deployment output or `DEPLOYMENT.md`.

</details>

<details>
<summary><strong>AI Services</strong></summary>

```env
# OpenAI Configuration
OPENAI_API_KEY=sk-...
OPENAI_MODEL=gpt-4-turbo-preview

# Search Provider
SEARCH_PROVIDER=serper
SEARCH_API_KEY=your_serper_api_key
MAX_SEARCH_RESULTS=10
```

**Get API Keys:**
- OpenAI: https://platform.openai.com/api-keys
- Serper: https://serper.dev/

</details>

**Step 3: Build the Binary**

```bash
go build -o bin/ai-resolver ./cmd/server
```

**Expected Output:**
```
Successfully built ai-resolver
Binary location: bin/ai-resolver
```

---

## Usage

### Start the Server

```bash
# Run directly
./bin/ai-resolver

# Or with custom port
SERVER_PORT=3000 ./bin/ai-resolver
```

**Expected Output:**
```
AI Resolver Service v1.0.0
Server listening on :8080
Signer: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb
Chain ID: 56
```

### API Endpoints

#### Health Check

<table>
<tr>
<td width="30%">

**Endpoint**

```
GET /healthz
GET /v1/healthz
```

</td>
<td width="70%">

**Response**

```json
{
  "status": "healthy",
  "version": "1.0.0",
  "time": 1698765432,
  "signer": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chainId": 56
}
```

</td>
</tr>
</table>

**Status Codes:**
- `200 OK` - Service healthy
- `503 Service Unavailable` - Service unhealthy

#### Propose Market Resolution

<table>
<tr>
<td width="30%">

**Endpoint**

```
POST /v1/propose
Content-Type: application/json
```

</td>
<td width="70%">

**Request Body**

```json
{
  "marketId": 123
}
```

</td>
</tr>
</table>

**Response (Success):**

```json
{
  "status": "submitted",
  "marketId": 123,
  "outcomeId": 1,
  "confidence": 0.87,
  "reasoning": "Based on multiple credible sources...",
  "txHash": "0xabc123...",
  "evidenceHash": "0xdef456...",
  "citations": 5,
  "facts": 8
}
```

**Status Codes:**
- `200 OK` - Proposal submitted successfully
- `400 Bad Request` - Invalid market ID
- `409 Conflict` - Market already resolved
- `500 Internal Server Error` - Processing failed

#### List Pending Markets

<table>
<tr>
<td width="30%">

**Endpoint**

```
GET /v1/markets
```

</td>
<td width="70%">

**Response**

```json
{
  "markets": [],
  "count": 0,
  "message": "Market listing not yet implemented"
}
```

</td>
</tr>
</table>

**Coming Soon:** This endpoint will return markets eligible for resolution.

---

## Configuration

### Environment Variables Reference

<table>
<tr>
<th>Variable</th>
<th>Description</th>
<th>Default</th>
<th>Required</th>
</tr>
<tr>
<td><strong>SERVER_PORT</strong></td>
<td>HTTP server port</td>
<td>8080</td>
<td>No</td>
</tr>
<tr>
<td><strong>LOG_LEVEL</strong></td>
<td>Logging verbosity (debug/info/warn/error)</td>
<td>info</td>
<td>No</td>
</tr>
<tr>
<td><strong>CHAIN_ID</strong></td>
<td>Blockchain network ID</td>
<td>56 (BSC)</td>
<td>Yes</td>
</tr>
<tr>
<td><strong>RPC_ENDPOINT</strong></td>
<td>Ethereum JSON-RPC URL</td>
<td>-</td>
<td>Yes</td>
</tr>
<tr>
<td><strong>SIGNER_PRIVATE_KEY</strong></td>
<td>Private key for EIP-712 signing</td>
<td>-</td>
<td>Yes</td>
</tr>
<tr>
<td><strong>OPENAI_API_KEY</strong></td>
<td>OpenAI API authentication key</td>
<td>-</td>
<td>Yes</td>
</tr>
<tr>
<td><strong>OPENAI_MODEL</strong></td>
<td>GPT model to use</td>
<td>gpt-4-turbo-preview</td>
<td>No</td>
</tr>
<tr>
<td><strong>SEARCH_PROVIDER</strong></td>
<td>Search provider (serper/google/brave)</td>
<td>serper</td>
<td>No</td>
</tr>
<tr>
<td><strong>SEARCH_API_KEY</strong></td>
<td>Search provider API key</td>
<td>-</td>
<td>Yes</td>
</tr>
<tr>
<td><strong>MAX_SEARCH_RESULTS</strong></td>
<td>Maximum search results to analyze</td>
<td>10</td>
<td>No</td>
</tr>
<tr>
<td><strong>DEFAULT_BOND_AMOUNT</strong></td>
<td>Bond amount in wei (1000 HORIZON)</td>
<td>1000000000000000000000</td>
<td>No</td>
</tr>
<tr>
<td><strong>PROPOSAL_TIMEOUT</strong></td>
<td>Timeout for proposal processing</td>
<td>5m</td>
<td>No</td>
</tr>
</table>

### Contract Addresses

<table>
<tr>
<th>Variable</th>
<th>Contract</th>
<th>Purpose</th>
</tr>
<tr>
<td><strong>AI_ORACLE_ADAPTER_ADDR</strong></td>
<td>AIOracleAdapter</td>
<td>Receives signed proposals</td>
</tr>
<tr>
<td><strong>MARKET_FACTORY_ADDR</strong></td>
<td>MarketFactory</td>
<td>Query market information</td>
</tr>
<tr>
<td><strong>RESOLUTION_MODULE_ADDR</strong></td>
<td>ResolutionModule</td>
<td>Check resolution status</td>
</tr>
<tr>
<td><strong>HORIZON_TOKEN_ADDR</strong></td>
<td>HorizonToken</td>
<td>Bond token transfers</td>
</tr>
</table>

---

## Development

### Generate Contract Bindings

After contract updates, regenerate Go bindings using `abigen`:

<details>
<summary><strong>Generate AIOracleAdapter Bindings</strong></summary>

```bash
cd contracts

# Extract ABI
jq -r '.abi' out/AIOracleAdapter.sol/AIOracleAdapter.json > /tmp/AIOracleAdapter.abi

# Generate Go bindings
~/go/bin/abigen \
  --abi /tmp/AIOracleAdapter.abi \
  --pkg abi \
  --type AIOracleAdapter \
  --out ../ai-resolver/pkg/abi/AIOracleAdapter.go
```

</details>

<details>
<summary><strong>Generate All Contract Bindings</strong></summary>

```bash
# Generate bindings for all contracts
for contract in AIOracleAdapter MarketFactory ResolutionModule HorizonToken; do
  jq -r '.abi' out/${contract}.sol/${contract}.json > /tmp/${contract}.abi
  ~/go/bin/abigen \
    --abi /tmp/${contract}.abi \
    --pkg abi \
    --type ${contract} \
    --out ../ai-resolver/pkg/abi/${contract}.go
done
```

</details>

### Run Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./internal/llm

# Verbose output
go test -v ./...
```

### Local Development with Anvil

Use Foundry's Anvil for local testing:

<details>
<summary><strong>Setup Local Testnet</strong></summary>

```bash
# Terminal 1: Start Anvil
cd contracts
anvil --fork-url https://bsc-dataseed.binance.org --fork-block-number 43000000

# Terminal 2: Update .env for local testing
RPC_ENDPOINT=http://127.0.0.1:8545
CHAIN_ID=31337
SIGNER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Terminal 3: Deploy contracts locally
forge script script/Deploy.s.sol:Deploy --rpc-url http://127.0.0.1:8545 --broadcast

# Terminal 4: Start AI Resolver
./bin/ai-resolver
```

</details>

---

## Deployment

### Docker Deployment

#### Quick Start with Docker Compose

**Easiest way to run the full stack:**

```bash
# 1. Configure environment
cp .env.example .env
# Edit .env with your OPENAI_API_KEY

# 2. Start services (Anvil + AI Resolver)
docker-compose up -d

# 3. Check health
curl http://localhost:8080/healthz

# 4. View logs
docker-compose logs -f ai-resolver

# 5. Stop services
docker-compose down
```

**Services Started:**

<table>
<tr>
<th>Service</th>
<th>Port</th>
<th>Description</th>
</tr>
<tr>
<td><strong>anvil</strong></td>
<td>8545</td>
<td>Local Ethereum testnet (chain ID: 31337)</td>
</tr>
<tr>
<td><strong>ai-resolver</strong></td>
<td>8080</td>
<td>AI Resolver API server</td>
</tr>
</table>

#### Manual Docker Build

<details>
<summary><strong>Build and Run Manually</strong></summary>

```bash
# Build the image
docker build -t ai-resolver:latest .

# Run the container
docker run -p 8080:8080 \
  --env-file .env \
  --name ai-resolver \
  ai-resolver:latest

# View logs
docker logs -f ai-resolver

# Stop container
docker stop ai-resolver
docker rm ai-resolver
```

</details>

#### Docker Compose Configuration

The `docker-compose.yml` provides a complete development environment:

**Anvil Service:**
- Pre-funded test accounts
- 2-second block time
- Chain ID: 31337
- Accessible at `http://localhost:8545`

**AI Resolver Service:**
- Connects to Anvil automatically
- Health check monitoring
- Auto-restart on failure
- Volume-mounted for logs

**Environment Variables:**

```env
# Anvil default account (TESTING ONLY!)
SIGNER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# OpenAI (required - add your key)
OPENAI_API_KEY=sk-...

# Network (auto-configured for anvil)
CHAIN_ID=31337
RPC_ENDPOINT=http://anvil:8545
```

**⚠️ WARNING**: The default signer key is Anvil's first account. NEVER use this in production!

### Production Deployment

<table>
<tr>
<td width="33%">

**AWS Lambda**

Serverless deployment with API Gateway

*Coming Soon*

</td>
<td width="33%">

**Kubernetes**

See `k8s/deployment.yaml`

*Coming Soon*

</td>
<td width="33%">

**VPS/Cloud**

Traditional server deployment with systemd

*Recommended for MVP*

</td>
</tr>
</table>

---

## Security Considerations

### Critical Security Practices

<table>
<tr>
<td width="50%">

**Private Key Management**

- ❌ Never commit private keys to Git
- ✅ Use environment variables only
- ✅ AWS KMS for production keys
- ✅ Rotate keys regularly
- ✅ Use dedicated signer accounts

</td>
<td width="50%">

**Operational Security**

- ✅ Rate limit API endpoints
- ✅ Update CORS settings for production
- ✅ Monitor signer account balance
- ✅ Set up alerting for failures
- ✅ Implement request logging

</td>
</tr>
</table>

### Bond Requirements

The signer account must maintain sufficient HORIZON tokens:

```bash
# Check HORIZON balance
cast balance --erc20 $HORIZON_TOKEN_ADDR $SIGNER_ADDRESS --rpc-url <network>

# Minimum required: DEFAULT_BOND_AMOUNT (default: 1000 HORIZON)
# Recommended: 10,000+ HORIZON for multiple proposals
```

### Signature Verification

All proposals are verified on-chain:
- EIP-712 typed data structure
- ECDSA signature validation
- Signer authorization check
- Replay attack prevention
- Time-bounded validity

---

## Monitoring

### Logging

Structured logs output to stdout:

```json
{
  "level": "info",
  "timestamp": "2025-10-29T12:34:56Z",
  "message": "Proposal submitted successfully",
  "marketId": 123,
  "txHash": "0xabc...",
  "confidence": 0.87
}
```

### Integration Options

<table>
<tr>
<td width="33%">

**AWS CloudWatch**

- Log aggregation
- Metric dashboards
- Alarms and alerts

</td>
<td width="33%">

**GCP Stackdriver**

- Centralized logging
- Trace analysis
- Error reporting

</td>
<td width="33%">

**Elastic Stack**

- Self-hosted option
- Full-text search
- Kibana dashboards

</td>
</tr>
</table>

### Key Metrics to Monitor

<table>
<tr>
<th>Metric</th>
<th>Description</th>
<th>Alert Threshold</th>
</tr>
<tr>
<td><strong>Proposal Success Rate</strong></td>
<td>Percentage of proposals accepted</td>
<td>< 90%</td>
</tr>
<tr>
<td><strong>LLM Latency</strong></td>
<td>Time for AI analysis completion</td>
<td>> 30 seconds</td>
</tr>
<tr>
<td><strong>Transaction Gas Costs</strong></td>
<td>Average gas per proposal</td>
<td>> 2x expected</td>
</tr>
<tr>
<td><strong>API Response Time</strong></td>
<td>95th percentile response time</td>
<td>> 5 seconds</td>
</tr>
<tr>
<td><strong>Bond Balance</strong></td>
<td>HORIZON tokens available</td>
<td>< 5000 tokens</td>
</tr>
</table>

---

## Troubleshooting

### Common Issues

<table>
<tr>
<th>Issue</th>
<th>Cause</th>
<th>Solution</th>
</tr>
<tr>
<td><strong>Failed to connect to Ethereum node</strong></td>
<td>
RPC endpoint unreachable or rate limited
</td>
<td>
• Check <code>RPC_ENDPOINT</code> accessibility<br>
• Verify network connectivity<br>
• Try alternative RPC endpoints<br>
• Check firewall settings
</td>
</tr>
<tr>
<td><strong>Failed to parse private key</strong></td>
<td>
Private key format incorrect
</td>
<td>
• Remove <code>0x</code> prefix if present<br>
• Verify key is 64 hex characters<br>
• Ensure no spaces or newlines<br>
• Check key validity with <code>cast wallet address</code>
</td>
</tr>
<tr>
<td><strong>Insufficient bond balance</strong></td>
<td>
Signer account lacks HORIZON tokens
</td>
<td>
• Check balance: <code>cast balance --erc20 $HORIZON_TOKEN $SIGNER</code><br>
• Acquire HORIZON from DEX or faucet<br>
• Verify token contract address<br>
• Ensure sufficient approval
</td>
</tr>
<tr>
<td><strong>Transaction failed</strong></td>
<td>
Various blockchain-related issues
</td>
<td>
• Check gas price and limits<br>
• Verify contract addresses<br>
• Ensure market is in correct state<br>
• Check signer is authorized<br>
• Review transaction revert reason
</td>
</tr>
<tr>
<td><strong>OpenAI API error</strong></td>
<td>
API key invalid or rate limited
</td>
<td>
• Verify API key is active<br>
• Check billing status<br>
• Reduce request rate<br>
• Use rate limiting
</td>
</tr>
</table>

### Debug Commands

<details>
<summary><strong>Check Signer Configuration</strong></summary>

```bash
# Get signer address from private key
cast wallet address $SIGNER_PRIVATE_KEY

# Check if signer is authorized
cast call $AI_ORACLE_ADAPTER_ADDR "authorizedSigners(address)" $SIGNER_ADDRESS \
  --rpc-url $RPC_ENDPOINT

# Expected output: true (0x0000000000000000000000000000000000000000000000000000000000000001)
```

</details>

<details>
<summary><strong>Check HORIZON Balance</strong></summary>

```bash
# Check signer's HORIZON balance
cast balance --erc20 $HORIZON_TOKEN_ADDR $SIGNER_ADDRESS \
  --rpc-url $RPC_ENDPOINT

# Check allowance to ResolutionModule
cast call $HORIZON_TOKEN_ADDR "allowance(address,address)" \
  $SIGNER_ADDRESS $RESOLUTION_MODULE_ADDR \
  --rpc-url $RPC_ENDPOINT
```

</details>

<details>
<summary><strong>Test RPC Connectivity</strong></summary>

```bash
# Test RPC endpoint
cast block-number --rpc-url $RPC_ENDPOINT

# Check chain ID
cast chain-id --rpc-url $RPC_ENDPOINT

# Get gas price
cast gas-price --rpc-url $RPC_ENDPOINT
```

</details>

### Enable Debug Logging

```bash
# Set log level to debug
LOG_LEVEL=debug ./bin/ai-resolver

# Or in .env
LOG_LEVEL=debug
```

---

## Additional Resources

<div align="center">

<table>
<tr>
<td align="center" width="25%">

**Smart Contracts**

[Contracts README](../contracts/)

Contract specifications

</td>
<td align="center" width="25%">

**Deployment**

[Deployment Guide](../contracts/docs/DEPLOYMENT.md)

Contract addresses

</td>
<td align="center" width="25%">

**Security**

[Security Docs](../contracts/docs/SECURITY.md)

Security measures

</td>
<td align="center" width="25%">

**API Reference**

[OpenAI Docs](https://platform.openai.com/docs)

GPT-4 API

</td>
</tr>
</table>

</div>

---

## License

MIT License - See [LICENSE](../LICENSE) file for details

---

## Support

<div align="center">

**Need Help with AI Resolver?**

<table>
<tr>
<td align="center" width="33%">

**GitHub Issues**

Report bugs or request features

[Open Issue](https://github.com/HorizonOracles/Project_Gamma/issues)

</td>
<td align="center" width="33%">

**Discord Community**

[![Discord](https://img.shields.io/badge/Discord-Join%20Us-5865F2?style=flat-square&logo=discord&logoColor=white)](https://discord.com/invite/TuUHwwKjHh)

</td>
<td align="center" width="33%">

**Developer Support**

developers@horizonoracles.com

</td>
</tr>
</table>

</div>

---

<div align="center">

**AI Resolver Service** • Automated Evidence-Based Resolution

[Main Documentation](../) • [Smart Contracts](../contracts/) • [Website](https://horizonoracles.com/)

*Built with Go, OpenAI GPT-4, and EIP-712 signatures*

</div>
