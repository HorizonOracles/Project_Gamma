<div align="center">

# Project Gamma - Smart Contracts

**On-chain prediction markets with AI resolution on BNB Chain**

[![Solidity](https://img.shields.io/badge/Solidity-0.8.23-363636?style=flat-square&logo=solidity)](https://soliditylang.org/)
[![Foundry](https://img.shields.io/badge/Built%20with-Foundry-black?style=flat-square)](https://book.getfoundry.sh/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](../LICENSE)
[![Tests](https://img.shields.io/badge/Tests-Passing-success?style=flat-square)]()

*Decentralized trading and resolution infrastructure for prediction markets*

---

[Overview](#overview) • [Architecture](#architecture) • [Development](#development) • [Testing](#testing) • [Deployment](#deployment)

---

</div>

## Overview

Project Gamma is a decentralized prediction market platform that combines on-chain trading with AI-powered resolution. All critical state including fees, splits, bonds, proposals, disputes, and outcomes lives on-chain, with only the AI reasoning happening off-chain.

### Core Principles

<table>
<tr>
<td width="33%" valign="top">

**On-Chain First**

All trading, fees, state, and resolution data stored permanently on BNB Chain

</td>
<td width="33%" valign="top">

**AI Resolution**

Optimistic AI-powered outcome proposals with community dispute mechanism

</td>
<td width="33%" valign="top">

**Battle-Tested**

Built with OpenZeppelin libraries and comprehensive test coverage

</td>
</tr>
</table>

---

## Architecture

### Smart Contract System

<table>
<tr>
<td width="50%">

**Core Trading Infrastructure**

- **OutcomeToken.sol** - ERC-1155 for Yes/No outcome shares
- **MarketAMM.sol** - Binary constant product AMM (x·y=k)
- **MarketFactory.sol** - Market creation and registry
- **FeeSplitter.sol** - Fee distribution (10% protocol / 90% creator)

</td>
<td width="50%">

**Resolution & Utility**

- **ResolutionModule.sol** - Resolution state machine
- **AIOracleAdapter.sol** - EIP-712 signature verification
- **HorizonPerks.sol** - Fee tier discounts via HORIZON tokens

</td>
</tr>
</table>

### Key Features

**Automated Market Maker**
- Constant product formula for efficient price discovery
- 2% trading fees with dynamic protocol/creator split
- Liquidity provision with LP tokens

**AI-Powered Resolution**
- Optimistic resolution proposals with cryptographic signatures
- 48-hour dispute window for community review
- Bond-based dispute mechanism with arbitrator fallback

**Token Utility**
- HORIZON token for market creation stakes
- Fee tier discounts based on token holdings
- Future governance participation

**Security Architecture**
- Reentrancy guards on all state-changing functions
- Pull-over-push payment pattern
- Pausable functionality for emergencies
- Multi-sig support for critical admin roles

---

## Development

### Prerequisites

<table>
<tr>
<td width="50%">

**Required Tools**

- [Foundry](https://book.getfoundry.sh/getting-started/installation) - Ethereum development toolkit
- Solidity 0.8.23 - Smart contract language

</td>
<td width="50%">

**Recommended Tools**

- [Slither](https://github.com/crytic/slither) - Static analysis
- [Echidna](https://github.com/crytic/echidna) - Fuzzing
- VSCode with Solidity extension

</td>
</tr>
</table>

### Quick Start

**1. Clone and Install**

```bash
git clone https://github.com/HorizonOracles/Project_Gamma.git
cd Project_Gamma/contracts

# Install dependencies
forge install
```

**2. Configure Environment**

```bash
cp .env.example .env
# Edit .env with your configuration
```

Required environment variables:

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

**3. Build and Test**

```bash
# Compile contracts
forge build

# Run tests
forge test
```

### Project Structure

```
contracts/
│
├── src/                    Contract source files
│   ├── OutcomeToken.sol        ERC-1155 outcome tokens
│   ├── MarketAMM.sol           Constant product AMM
│   ├── MarketFactory.sol       Market creation registry
│   ├── FeeSplitter.sol         Fee distribution logic
│   ├── ResolutionModule.sol    Resolution state machine
│   ├── AIOracleAdapter.sol     AI proposal verification
│   └── HorizonPerks.sol        Fee tier management
│
├── test/                   Test files
│   ├── unit/                  Unit tests per contract
│   ├── integration/           Cross-contract workflows
│   └── invariant/             Property-based tests
│
├── script/                 Deployment scripts
│   └── Deploy.s.sol           Main deployment script
│
└── docs/                   Additional documentation
    ├── ARCHITECTURE.md        System design details
    ├── API.md                 Function reference
    └── DEPLOYMENT.md          Deployment guide
```

### Configuration

Key settings in `foundry.toml`:

| Setting | Value | Purpose |
|---------|-------|---------|
| **Solidity Version** | 0.8.23 | Compiler version |
| **Optimizer** | Enabled (200 runs) | Gas optimization |
| **Fuzz Runs** | 256 | Fuzzing iterations |
| **Invariant Runs** | 256 | Property test depth |

---

## Testing

### Test Suite Overview

<table>
<tr>
<td width="33%" align="center">

**Unit Tests**

Test individual contract functionality in isolation

```bash
forge test --match-path "test/unit/*.sol"
```

</td>
<td width="33%" align="center">

**Integration Tests**

Test complete workflows across multiple contracts

```bash
forge test --match-path "test/integration/*.sol"
```

</td>
<td width="33%" align="center">

**Invariant Tests**

Property-based testing for system guarantees

```bash
forge test --match-contract Invariant
```

</td>
</tr>
</table>

### Running Tests

<details>
<summary><strong>Basic Test Commands</strong></summary>

```bash
# Run all tests
forge test

# Run with detailed output
forge test -vvv

# Run specific test
forge test --match-test testBuyYes

# Run specific contract tests
forge test --match-contract MarketAMMTest
```

</details>

<details>
<summary><strong>Advanced Testing</strong></summary>

```bash
# Gas reporting
forge test --gas-report

# Coverage analysis
forge coverage

# Coverage with LCOV report
forge coverage --report lcov

# Run with specific verbosity
forge test -vvvv  # Maximum verbosity for debugging
```

</details>

### Test Coverage Goals

- **Target:** >95% line coverage across all contracts
- **Unit Tests:** 100% function coverage
- **Integration Tests:** All critical user flows
- **Invariant Tests:** Core economic and security properties

---

## Deployment

### Testnet Deployment

**BNB Chain Testnet**

```bash
# Deploy to testnet
forge script script/Deploy.s.sol \
    --rpc-url bsc_testnet \
    --broadcast \
    --verify

# Verify individual contract (if needed)
forge verify-contract <ADDRESS> <CONTRACT_NAME> \
    --chain-id 97 \
    --etherscan-api-key $BSCSCAN_API_KEY
```

### Mainnet Deployment

**Pre-Deployment Checklist**

<table>
<tr>
<td width="50%">

**Code Quality**
- [ ] All tests passing (>95% coverage)
- [ ] Slither analysis clean
- [ ] Echidna fuzzing complete (72h)
- [ ] Code review completed
- [ ] Documentation updated

</td>
<td width="50%">

**Security & Operations**
- [ ] External audit completed
- [ ] Multi-sig wallet configured
- [ ] Emergency pause procedures documented
- [ ] Sufficient BNB for deployment gas
- [ ] Backup deployment key secured

</td>
</tr>
</table>

**Deploy to Mainnet**

```bash
forge script script/Deploy.s.sol \
    --rpc-url bsc \
    --broadcast \
    --verify \
    --slow  # Add delay between transactions
```

**Post-Deployment Steps**

1. Verify all contracts on BscScan
2. Transfer ownership to multi-sig wallet
3. Configure AI oracle signers
4. Set initial protocol parameters
5. Test with small trades before public announcement
6. Update documentation with deployed addresses

### Deployment Addresses

After deployment, update your `.env` and documentation:

**BSC Testnet**
```
HORIZON_TOKEN_ADDRESS=0x...
MARKET_FACTORY_ADDRESS=0x...
OUTCOME_TOKEN_ADDRESS=0x...
RESOLUTION_MODULE_ADDRESS=0x...
```

**BSC Mainnet**
```
HORIZON_TOKEN_ADDRESS=0x...
MARKET_FACTORY_ADDRESS=0x...
OUTCOME_TOKEN_ADDRESS=0x...
RESOLUTION_MODULE_ADDRESS=0x...
```

---

## Gas Optimization

### Target Gas Costs

| Operation | Target Gas | Status |
|-----------|-----------|--------|
| Market Creation | ~TBD | Pending benchmark |
| Add Liquidity | ~TBD | Pending benchmark |
| Buy/Sell Trade | ~TBD | Pending benchmark |
| Propose Resolution | ~TBD | Pending benchmark |
| Redeem Outcome | ~TBD | Pending benchmark |

### Gas Optimization Strategies

- Optimizer enabled at 200 runs (balance between deployment and runtime)
- Minimal storage reads/writes
- Batch operations where possible
- Events for off-chain indexing instead of storage
- Efficient data structures (mappings over arrays)

---

## Security

### Security Measures

<table>
<tr>
<td width="33%">

**Smart Contract Security**
- OpenZeppelin library usage
- Reentrancy guards
- Pull payment pattern
- Pausable contracts
- Access control

</td>
<td width="33%">

**Economic Security**
- Bond-based disputes
- Stake requirements
- Fee distribution checks
- Slippage protection
- Overflow protection

</td>
<td width="33%">

**Operational Security**
- Multi-sig admin control
- Timelock for upgrades
- Emergency pause
- Rate limiting
- Parameter bounds

</td>
</tr>
</table>

### Audit Status

| Phase | Status | Details |
|-------|--------|---------|
| Internal Security Review | ✓ Completed | Core team review |
| Slither Static Analysis | Pending | Automated scanning |
| Echidna Fuzzing | Pending | 72-hour campaign |
| External Audit | Scheduled | Third-party audit firm |
| Bug Bounty | Planned | Post-launch program |

### Reporting Vulnerabilities

**Critical security issues should be reported privately:**

- Use GitHub's security advisory feature
- Email: developers@horizonoracles.com
- Subject: "Security Vulnerability - Project Gamma Contracts"

**Include:**
- Vulnerability description
- Proof of concept
- Potential impact assessment
- Suggested remediation

---

## Development Tools

### Useful Commands

<table>
<tr>
<td width="50%">

**Building & Testing**

```bash
forge build              # Compile contracts
forge clean              # Clean artifacts
forge test               # Run tests
forge coverage           # Coverage report
forge snapshot           # Gas snapshot
```

</td>
<td width="50%">

**Code Quality**

```bash
forge fmt                # Format code
slither .                # Static analysis
echidna .                # Fuzzing
forge doc                # Generate docs
```

</td>
</tr>
</table>

### Debugging

```bash
# Run with maximum verbosity
forge test -vvvv

# Debug specific test with traces
forge test --match-test testFunction --debug

# Generate gas report
forge test --gas-report

# Profile gas usage
forge test --gas-report > gas-report.txt
```

---

## Documentation

### Contract Documentation

<table>
<tr>
<td width="50%">

**Technical Documentation**
- [Architecture Overview](./docs/ARCHITECTURE.md)
- [API Reference](./docs/API.md)
- [Deployment Guide](./docs/DEPLOYMENT.md)
- [Implementation Roadmap](../implementation-roadmap.md)

</td>
<td width="50%">

**Related Repositories**
- [AI Resolver](../ai-resolver) - Go-based resolution service
- [SDK (TypeScript)](../sdk-js) - JavaScript integration
- [Subgraph](../subgraph) - The Graph indexer

</td>
</tr>
</table>

### Generate Documentation

```bash
# Generate contract docs
forge doc

# Serve docs locally
forge doc --serve
# Access at http://localhost:3000
```

---

## Contributing

We welcome contributions to the smart contracts! Please follow these guidelines:

### Development Workflow

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/contract-improvement`
3. Make your changes with tests
4. Ensure all tests pass: `forge test`
5. Format your code: `forge fmt`
6. Commit: `git commit -m "Add: contract improvement"`
7. Push and open a Pull Request

### Code Standards

- Follow Solidity style guide
- Maintain >95% test coverage for new code
- Use NatSpec comments for all public functions
- No compiler warnings allowed
- Gas optimization where reasonable

### Testing Requirements

- Unit tests for all new functions
- Integration tests for new workflows
- Update invariant tests if adding state
- Gas benchmarks for expensive operations

---

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

---

## Support

<div align="center">

**Need Help with Smart Contract Development?**

[![GitHub Issues](https://img.shields.io/badge/GitHub-Issues-181717?style=flat-square&logo=github)](https://github.com/HorizonOracles/Project_Gamma/issues)
[![Discord](https://img.shields.io/badge/Discord-Join%20Us-5865F2?style=flat-square&logo=discord&logoColor=white)](https://discord.com/invite/TuUHwwKjHh)
[![Email](https://img.shields.io/badge/Email-developers%40horizonoracles.com-blue?style=flat-square)](mailto:developers@horizonoracles.com)

For questions about smart contract integration, security, or development, reach out through any of the channels above.

</div>

---

<div align="center">

**Project Gamma Smart Contracts** • Built on BNB Chain with Foundry

[Main Repository](../) • [Website](https://horizonoracles.com/) • [Twitter](https://x.com/HorizonOracles)

</div>
