# Project Gamma - Smart Contracts

On-chain prediction markets with AI resolution on BNB Chain (BSC).

## Overview

Project Gamma is a decentralized prediction market platform that combines on-chain trading with AI-powered resolution. All critical state (fees, splits, bonds, proposals, disputes, outcomes) lives on-chain, with only the AI reasoning happening off-chain.

## Architecture

### Core Contracts

- **OutcomeToken.sol** - ERC-1155 for outcome shares (Yes/No tokens)
- **MarketAMM.sol** - Binary constant product market maker (x*y=k)
- **MarketFactory.sol** - Market creation and registry
- **FeeSplitter.sol** - Fee distribution (10% protocol, 90% creator)
- **ResolutionModule.sol** - Resolution state machine (propose, dispute, finalize)
- **AIOracleAdapter.sol** - EIP-712 signature verification for AI proposals
- **HorizonPerks.sol** - Fee tier discounts based on HORIZON token holdings

## Features

- **On-chain first**: All trading, fees, state, and resolution on-chain
- **AI resolution**: Optimistic AI-powered outcome proposals (disputeable)
- **AMM trading**: Constant product market maker with 2% fees
- **Fee sharing**: 10% protocol / 90% creator split
- **Token utility**: HORIZON token for stakes, bonds, and fee tiers
- **Dispute mechanism**: Bond-based disputes with arbitrator fallback

## Development

### Prerequisites

- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- Solidity 0.8.23

### Setup

```bash
# Clone the repository
git clone <repo-url>
cd contracts

# Install dependencies
forge install

# Copy environment variables
cp .env.example .env
# Edit .env with your configuration
```

### Build

```bash
forge build
```

### Test

```bash
# Run all tests
forge test

# Run with verbosity
forge test -vvv

# Run specific test
forge test --match-test testBuyYes

# Run with gas report
forge test --gas-report

# Run invariant tests
forge test --match-contract Invariant
```

### Format

```bash
forge fmt
```

### Deploy

```bash
# Deploy to BSC testnet
forge script script/Deploy.s.sol --rpc-url bsc_testnet --broadcast --verify

# Deploy to BSC mainnet
forge script script/Deploy.s.sol --rpc-url bsc --broadcast --verify
```

## Project Structure

```
contracts/
├── src/
│   ├── OutcomeToken.sol
│   ├── MarketAMM.sol
│   ├── MarketFactory.sol
│   ├── FeeSplitter.sol
│   ├── ResolutionModule.sol
│   ├── AIOracleAdapter.sol
│   └── HorizonPerks.sol
├── test/
│   ├── unit/
│   ├── integration/
│   └── invariant/
├── script/
│   └── Deploy.s.sol
└── docs/
    ├── ARCHITECTURE.md
    ├── API.md
    └── DEPLOYMENT.md
```

## Configuration

Key parameters (see `foundry.toml`):

- **Solidity**: 0.8.23
- **Optimizer**: Enabled (200 runs)
- **Fuzz runs**: 256
- **Invariant runs**: 256

## Environment Variables

Create a `.env` file with:

```bash
# RPC Endpoints
BSC_RPC_URL=https://bsc-dataseed.binance.org/
BSC_TESTNET_RPC_URL=https://data-seed-prebsc-1-s1.binance.org:8545/

# Deployer
PRIVATE_KEY=your_private_key_here

# Verification
BSCSCAN_API_KEY=your_bscscan_api_key

# Contract Addresses (after deployment)
HORIZON_TOKEN_ADDRESS=
MARKET_FACTORY_ADDRESS=
```

## Security

- All contracts use OpenZeppelin libraries for standard implementations
- Reentrancy guards on all state-changing functions
- Pull-over-push payment pattern for fees
- Pausable functionality for emergency situations
- Multi-sig for critical admin roles

### Audits

- [ ] Internal security review
- [ ] Slither static analysis
- [ ] Echidna fuzzing (72h)
- [ ] External audit (planned)

## Gas Costs

Target gas costs (to be benchmarked):

- Market creation: ~TBD gas
- Add liquidity: ~TBD gas
- Buy/Sell trade: ~TBD gas
- Propose resolution: ~TBD gas
- Redeem: ~TBD gas

## License

MIT

## Documentation

- [Architecture Overview](../docs/ARCHITECTURE.md)
- [API Reference](../docs/API.md)
- [Deployment Guide](../docs/DEPLOYMENT.md)
- [Implementation Roadmap](../implementation-roadmap.md)

## Related Repositories

- [ai-resolver](../ai-resolver) - Go service for AI-powered resolution
- [sdk-js](../sdk-js) - TypeScript SDK (optional)
- [subgraph](../subgraph) - The Graph indexer (optional)

## Support

- GitHub Issues: [Report a bug](https://github.com/...)
- Documentation: [Full docs](../docs/)
