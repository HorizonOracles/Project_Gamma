# AI Resolver Service

Backend service for automated AI-powered market resolution on the Gamma prediction market platform.

## Architecture

The AI Resolver implements a complete pipeline for resolving binary prediction markets:

1. **Web Search** - Queries search providers (Serper.dev, Google, Brave) for evidence
2. **LLM Multi-Pass Analysis** - 4-step pipeline using OpenAI GPT-4:
   - Extract facts from search results
   - Check for contradictions
   - Decide outcome with confidence
   - Build citations with evidence
3. **EIP-712 Signing** - Cryptographic signature for on-chain proposal verification
4. **Blockchain Submission** - Submit proposal to AIOracleAdapter smart contract

## Project Structure

```
ai-resolver/
├── cmd/server/          # HTTP server entry point
├── internal/
│   ├── config/         # Configuration management
│   ├── llm/            # OpenAI multi-pass pipeline
│   ├── search/         # Web search providers
│   ├── eip712/         # EIP-712 signing
│   └── adapter/        # Ethereum contract client
├── pkg/abi/            # Generated Go bindings (abigen)
├── configs/            # Configuration files
└── bin/                # Compiled binaries
```

## Setup

### Prerequisites

- Go 1.24+
- Ethereum node access (BSC, Ethereum, or Base)
- OpenAI API key
- Search API key (Serper.dev)
- HORIZON token balance for bonding

### Installation

1. Clone the repository
2. Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

3. Set required environment variables:

```env
# Blockchain
RPC_ENDPOINT=https://bsc-dataseed.binance.org
CHAIN_ID=56
SIGNER_PRIVATE_KEY=your_private_key_here

# Contracts (from DEPLOYMENT.md)
AI_ORACLE_ADAPTER_ADDR=0x...
MARKET_FACTORY_ADDR=0x...
RESOLUTION_MODULE_ADDR=0x...
HORIZON_TOKEN_ADDR=0x...

# AI Services
OPENAI_API_KEY=sk-...
SEARCH_API_KEY=your_serper_api_key
```

4. Build the binary:

```bash
go build -o bin/ai-resolver ./cmd/server
```

## Usage

### Start the Server

```bash
./bin/ai-resolver
```

Server will start on `http://0.0.0.0:8080` by default.

### API Endpoints

#### Health Check
```bash
GET /healthz
GET /v1/healthz

Response:
{
  "status": "healthy",
  "version": "1.0.0",
  "time": 1698765432,
  "signer": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chainId": 56
}
```

#### Propose Market Resolution
```bash
POST /v1/propose
Content-Type: application/json

{
  "marketId": 123
}

Response:
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

#### List Pending Markets
```bash
GET /v1/markets

Response:
{
  "markets": [],
  "count": 0,
  "message": "Market listing not yet implemented"
}
```

## Configuration

All configuration is via environment variables. See `.env.example` for complete options.

### Key Settings

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | HTTP server port | `8080` |
| `CHAIN_ID` | Blockchain network | `56` (BSC) |
| `RPC_ENDPOINT` | Ethereum JSON-RPC URL | Required |
| `SIGNER_PRIVATE_KEY` | Private key for signing | Required |
| `OPENAI_API_KEY` | OpenAI API key | Required |
| `OPENAI_MODEL` | Model to use | `gpt-4-turbo-preview` |
| `SEARCH_PROVIDER` | Search provider | `serper` |
| `MAX_SEARCH_RESULTS` | Max search results | `10` |
| `DEFAULT_BOND_AMOUNT` | Bond in wei | `1000000000000000000000` |
| `PROPOSAL_TIMEOUT` | Timeout for proposals | `5m` |

## Development

### Generate Contract Bindings

After contract updates, regenerate Go bindings:

```bash
cd contracts
jq -r '.abi' out/AIOracleAdapter.sol/AIOracleAdapter.json > /tmp/AIOracleAdapter.abi
~/go/bin/abigen --abi /tmp/AIOracleAdapter.abi --pkg abi --type AIOracleAdapter \
  --out ../ai-resolver/pkg/abi/AIOracleAdapter.go
```

Repeat for `MarketFactory`, `ResolutionModule`, and `HorizonToken`.

### Run Tests

```bash
go test ./...
```

### Local Development

Use a local testnet or fork:

```bash
# In contracts directory
anvil --fork-url https://bsc-dataseed.binance.org --fork-block-number 43000000

# Update .env
RPC_ENDPOINT=http://127.0.0.1:8545
CHAIN_ID=31337
```

## Security Considerations

- **Private Key Management**: Never commit private keys. Use environment variables or AWS KMS in production.
- **Bond Requirements**: Ensure signer account has sufficient HORIZON tokens for bonding.
- **Rate Limiting**: Consider rate limiting API endpoints in production.
- **CORS**: Update CORS settings for production deployment.
- **Signature Verification**: Smart contract verifies EIP-712 signatures on-chain.

## Deployment

### Docker

#### Quick Start with Docker Compose

The easiest way to run the AI Resolver with a local Ethereum testnet:

```bash
# 1. Configure environment variables
cp .env.example .env
# Edit .env with your OPENAI_API_KEY and other settings

# 2. Start the full stack (Anvil + AI Resolver)
docker-compose up -d

# 3. Check service health
curl http://localhost:8080/healthz

# 4. View logs
docker-compose logs -f ai-resolver

# 5. Stop services
docker-compose down
```

This will start:
- **Anvil**: Local Ethereum node on `http://localhost:8545` (chain ID: 31337)
- **AI Resolver**: API server on `http://localhost:8080`

#### Manual Docker Build

```bash
# Build the image
docker build -t ai-resolver:latest .

# Run the container
docker run -p 8080:8080 --env-file .env ai-resolver:latest
```

#### Docker Compose Services

The `docker-compose.yml` includes:

- **anvil**: Foundry's local Ethereum node
  - Accessible at `http://localhost:8545`
  - Pre-funded accounts for testing
  - 2-second block time
  - Chain ID: 31337

- **ai-resolver**: AI Resolver service
  - Accessible at `http://localhost:8080`
  - Health check endpoint: `/healthz`
  - Auto-restarts on failure
  - Connects to anvil for blockchain operations

#### Environment Variables

Key variables for Docker deployment:

```env
# Server
SERVER_PORT=8080
LOG_LEVEL=debug

# Blockchain (automatically configured for anvil)
CHAIN_ID=31337
RPC_ENDPOINT=http://anvil:8545

# OpenAI (required)
OPENAI_API_KEY=sk-...
OPENAI_MODEL=gpt-4-turbo-preview

# Signer (Anvil default account - for testing only!)
SIGNER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Contract addresses (set after deployment)
AI_ORACLE_ADAPTER_ADDR=0x...
RESOLUTION_MODULE_ADDR=0x...
HORIZON_TOKEN_ADDR=0x...
MARKET_FACTORY_ADDR=0x...
```

**Important**: The default `SIGNER_PRIVATE_KEY` is Anvil's first account. Never use this key in production!

### Kubernetes

See `k8s/deployment.yaml` for Kubernetes manifests (coming soon).

### AWS Lambda

Can be adapted for serverless deployment with API Gateway.

## Monitoring

Logs are output to stdout in structured format. Integrate with:

- **CloudWatch** (AWS)
- **Stackdriver** (GCP)
- **Elastic** (self-hosted)

Key metrics to monitor:
- Proposal success rate
- LLM latency
- Transaction gas costs
- API response times

## Troubleshooting

### Common Issues

**"Failed to connect to Ethereum node"**
- Check `RPC_ENDPOINT` is accessible
- Verify network connectivity
- Try alternative RPC endpoints

**"Failed to parse private key"**
- Ensure `SIGNER_PRIVATE_KEY` is hex format **without** 0x prefix
- Example: `ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` ✓
- Not: `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` ✗
- Verify key is valid ECDSA private key

**"Insufficient bond balance"**
- Check HORIZON token balance: `cast balance --erc20 $HORIZON_TOKEN $SIGNER_ADDRESS`
- Acquire HORIZON tokens from DEX or faucet

**"Transaction failed"**
- Check gas price and limits
- Verify contract addresses are correct
- Ensure market is in correct state (closed, not yet resolved)

## License

MIT

## Contact

For issues and questions, please open a GitHub issue.
