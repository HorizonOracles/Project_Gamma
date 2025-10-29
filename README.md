<div align="center">

![Project Gamma Banner](assets/banner1.png)

# Project Gamma

**Decentralized Prediction Markets Powered by AI**

[![Website](https://img.shields.io/badge/Website-horizonoracles.com-blue?style=flat-square)](https://horizonoracles.com/)
[![Twitter](https://img.shields.io/badge/Twitter-@HorizonOracles-1DA1F2?style=flat-square&logo=twitter)](https://x.com/HorizonOracles)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)

[![BNB Chain](assets/bnb-chain-full-binance-smart-chain-logo-1-1024x180.png)](https://www.bnbchain.org/)

*Combining on-chain trading with AI-powered resolution for transparent, automated prediction markets*

---

[Overview](#overview) • [Architecture](#architecture) • [Getting Started](#getting-started) • [Documentation](#documentation) • [Security](#security)

---

</div>

## Overview

Project Gamma is a decentralized prediction market platform that enables users to create and trade binary outcome markets. The platform leverages AI-powered resolution for automated, evidence-based outcome determination while maintaining full on-chain transparency and decentralization on BNB Chain.

### Core Components

<table>
<tr>
<td width="33%" valign="top">

**Smart Contracts**

Solidity contracts deployed on BNB Chain managing market creation, trading, and resolution with full decentralization.

</td>
<td width="33%" valign="top">

**AI Resolver**

Go-based backend service that gathers evidence and proposes resolutions using cryptographic signatures.

</td>
<td width="33%" valign="top">

**HORIZON Token**

Platform utility token providing fee discounts and governance participation rights.

</td>
</tr>
</table>

---

## Architecture

### Two-Layer System Design

<table>
<tr>
<td width="50%" valign="top">

**Layer 1: On-Chain (BNB Chain)**

- Market creation and lifecycle management
- Automated market maker (AMM) for binary outcome trading
- Resolution state machine with dispute mechanism
- Fee collection and distribution
- Outcome token minting and redemption

</td>
<td width="50%" valign="top">

**Layer 2: Off-Chain (AI Resolver)**

- Evidence gathering through web search
- Multi-pass LLM analysis pipeline
- EIP-712 signature generation for proposals
- Automated proposal submission

</td>
</tr>
</table>

### Smart Contract Architecture

```
MarketFactory (Central Registry)
    │
    ├── MarketAMM (Per-Market Trading)
    │   └── Uses OutcomeToken (ERC-1155)
    │
    ├── ResolutionModule (Dispute & Resolution)
    │   └── AIOracleAdapter (Signature Verification)
    │
    ├── FeeSplitter (Fee Distribution)
    │
    └── HorizonPerks (Fee Tier Management)
```

---

## Core Features

### Market Trading

The platform implements a robust automated market maker with the following characteristics:

<table>
<tr>
<td width="50%">

**Constant Product Formula**

Utilizes the x·y=k formula for efficient price discovery and liquidity provision in binary outcome markets.

</td>
<td width="50%">

**Binary Outcomes**

Each market issues YES and NO outcome tokens, enabling traders to take positions on either side of a prediction.

</td>
</tr>
<tr>
<td width="50%">

**Liquidity Provision**

LP token holders earn proportional trading fees while providing liquidity to markets.

</td>
<td width="50%">

**Slippage Protection**

All trades include minimum output requirements to protect against unfavorable price movements.

</td>
</tr>
</table>

### Dynamic Fee Structure

The platform implements a tiered fee model that incentivizes HORIZON token ownership:

| Component | Rate | Description |
|-----------|------|-------------|
| **User Fee** | 2% | Constant across all trades |
| **Protocol/Creator Split** | Variable | Based on trader's HORIZON balance |

**Fee Tier Examples:**

- **Default (0 HORIZON):** 10% protocol / 90% creator
- **Tier 4 (500K+ HORIZON):** 2% protocol / 98% creator

This model incentivizes market creators to attract HORIZON token holders while maintaining simple, predictable fees for users.

### Resolution System

**Three-Stage Resolution Process**

<table>
<tr>
<td width="33%" align="center">

**1. Proposal**

AI analyzes evidence and proposes outcome with cryptographic signature

</td>
<td width="33%" align="center">

**2. Dispute Window**

48-hour period for community review and potential disputes

</td>
<td width="33%" align="center">

**3. Finalization**

Automatic finalization or manual arbitration if disputed

</td>
</tr>
</table>

**Security Features:**
- EIP-712 signature verification for all AI proposals
- Multi-signer support for redundancy and decentralization
- Evidence hash validation to prevent tampering
- Stake-based incentive alignment for proposers and disputers

### HORIZON Token Utility

<table>
<tr>
<td width="25%" align="center">

**Fee Optimization**

Reduce protocol fee share through token holdings

</td>
<td width="25%" align="center">

**Market Creation**

Required stake for creating markets (refunded after resolution)

</td>
<td width="25%" align="center">

**Governance**

Future participation rights in protocol decisions

</td>
<td width="25%" align="center">

**Creator Incentives**

Attract high-value traders to your markets

</td>
</tr>
</table>

---

## Smart Contracts

### Core Contract Registry

| Contract | Purpose | Key Responsibilities |
|----------|---------|---------------------|
| **MarketFactory** | Central Registry | Market creation, status management, creator stake handling |
| **MarketAMM** | Trading Engine | Binary outcome trading via constant product AMM |
| **ResolutionModule** | Resolution Manager | Lifecycle management, dispute handling, finalization |
| **AIOracleAdapter** | Signature Verifier | EIP-712 verification for AI-generated proposals |
| **OutcomeToken** | Token Standard | ERC-1155 implementation for outcome shares |
| **FeeSplitter** | Fee Distribution | Protocol and creator fee allocation |
| **HorizonPerks** | Fee Calculator | Dynamic fee tier computation |
| **HorizonToken** | Utility Token | ERC-20 platform token implementation |

### Contract Interaction Flows

<details>
<summary><strong>Market Creation Flow</strong></summary>

```
User → MarketFactory.createMarket()
    ├── Stakes HORIZON tokens
    ├── Deploys new MarketAMM instance
    ├── Registers with FeeSplitter
    ├── Registers with OutcomeToken
    └── Returns Market ID
```

</details>

<details>
<summary><strong>Trading Flow</strong></summary>

```
Trader → MarketAMM.buyYes()
    ├── Checks HORIZON balance for fee tier
    ├── Transfers collateral from trader
    ├── Mints outcome token pairs
    ├── Executes CPMM swap
    ├── Distributes fees via FeeSplitter
    └── Transfers outcome tokens to trader
```

</details>

<details>
<summary><strong>Resolution Flow</strong></summary>

```
AI Resolver → AIOracleAdapter.proposeAI()
    ├── Verifies EIP-712 signature
    ├── Validates evidence hash
    ├── Forwards to ResolutionModule
    ├── Starts 48-hour dispute window
    ├── Waits for disputes or timeout
    └── Finalizes and enables redemptions
```

</details>

---

## Getting Started

### Prerequisites

<table>
<tr>
<td width="50%">

**Development Environment**
- Foundry (smart contract development)
- Go 1.24+ (AI resolver)
- Node.js 18+ (optional tooling)

</td>
<td width="50%">

**Access & Keys**
- BNB Chain testnet/mainnet access
- OpenAI API key (for AI resolver)
- Private key for deployment

</td>
</tr>
</table>

### Quick Start

**1. Clone and Install**

```bash
git clone https://github.com/yourusername/project_gamma.git
cd project_gamma

# Install smart contract dependencies
cd contracts
forge install
```

**2. Configure Environment**

```bash
cp .env.example .env
# Edit .env with your configuration
```

**3. Run Tests**

```bash
forge test
```

### Local Development Commands

<table>
<tr>
<td width="50%">

**Smart Contracts**

```bash
forge build              # Compile contracts
forge test               # Run tests
forge test --gas-report  # Gas analysis
forge fmt                # Format code
forge doc                # Generate docs
```

</td>
<td width="50%">

**AI Resolver**

```bash
cd ai-resolver
go build -o bin/ai-resolver ./cmd/server
go test ./...            # Run tests
./bin/ai-resolver        # Start server
```

</td>
</tr>
</table>

---

## Development

### Project Structure

```
project_gamma/
│
├── contracts/              Smart contracts and tests
│   ├── src/               Contract source files
│   ├── test/              Unit and integration tests
│   ├── script/            Deployment scripts
│   └── docs/              Contract documentation
│
├── ai-resolver/           AI resolution service
│   ├── cmd/              Application entry points
│   ├── internal/         Internal packages
│   └── pkg/              Public packages
│
└── README.md             This file
```

### Creating New Smart Contracts

<details>
<summary><strong>Step-by-Step Guide</strong></summary>

**Step 1: Create Contract File**

```bash
cd contracts/src
touch MyNewContract.sol
```

**Step 2: Implement Contract**

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract MyNewContract is Ownable, ReentrancyGuard {
    uint256 public myValue;
    
    event ValueUpdated(uint256 oldValue, uint256 newValue);
    
    constructor(uint256 initialValue) Ownable(msg.sender) {
        myValue = initialValue;
    }
    
    function updateValue(uint256 newValue) external onlyOwner {
        uint256 oldValue = myValue;
        myValue = newValue;
        emit ValueUpdated(oldValue, newValue);
    }
}
```

**Step 3: Create Tests**

```bash
cd contracts/test/unit
touch MyNewContract.t.sol
```

**Step 4: Compile and Test**

```bash
forge build
forge test --match-contract MyNewContractTest
forge coverage --match-contract MyNewContract
```

</details>

### Testing Strategy

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

Test complete workflows and contract interactions

```bash
forge test --match-path "test/integration/*.sol"
```

</td>
<td width="33%" align="center">

**Coverage Analysis**

Maintain >95% test coverage

```bash
forge coverage
```

</td>
</tr>
</table>

### Foundry Cheatcodes Reference

<details>
<summary><strong>Common Testing Utilities</strong></summary>

```solidity
// Time manipulation
vm.warp(block.timestamp + 1 days);
vm.roll(block.number + 100);

// Identity manipulation
vm.prank(address);
vm.startPrank(address);
vm.stopPrank();

// Expectations
vm.expectRevert();
vm.expectEmit(true, true, true, true);

// Balance manipulation
vm.deal(address, 100 ether);
```

</details>

---

## Deployment

### Testnet Deployment

**BNB Chain Testnet**

```bash
cd contracts
forge script script/Deploy.s.sol \
    --rpc-url bsc_testnet \
    --broadcast \
    --verify
```

### Mainnet Deployment

**Pre-Deployment Checklist**

- [ ] All tests passing with >95% coverage
- [ ] Security audit completed
- [ ] Deployment script reviewed
- [ ] Multi-sig wallet configured
- [ ] Sufficient BNB for gas fees
- [ ] Environment configured for mainnet
- [ ] Private keys secured

**Deploy to Production**

```bash
forge script script/Deploy.s.sol \
    --rpc-url bsc_mainnet \
    --broadcast \
    --verify
```

**Post-Deployment Steps**

1. Transfer ownership to multi-sig wallet
2. Configure AI resolver signers
3. Set initial fee parameters
4. Verify all contracts on BscScan
5. Test with small trades before public launch

---

## Documentation

### Developer Resources

<table>
<tr>
<td width="50%">

**Smart Contracts**
- [Contract Specifications](./contracts/README.md)
- [Deployment Guide](./contracts/docs/DEPLOYMENT.md)
- [Security Considerations](./contracts/docs/SECURITY.md)
- [Roles & Permissions](./contracts/docs/ROLES.md)

</td>
<td width="50%">

**AI Resolver**
- [Backend Architecture](./ai-resolver/README.md)
- [Resolution Pipeline](./ai-resolver/docs/PIPELINE.md)
- API Documentation (Coming Soon)

</td>
</tr>
</table>

### API Documentation

Generate and serve contract documentation:

```bash
cd contracts
forge doc --serve
```

Access at `http://localhost:3000`

---

## Security

### Security Architecture

<table>
<tr>
<td width="33%">

**Smart Contract Security**
- OpenZeppelin libraries
- ReentrancyGuard protection
- Pausable emergency stops
- Comprehensive test coverage

</td>
<td width="33%">

**Cryptographic Security**
- EIP-712 typed signatures
- Evidence hash validation
- Multi-signer redundancy
- Signature expiration windows

</td>
<td width="33%">

**Operational Security**
- Hardware wallet support
- Multi-sig admin functions
- Environment isolation
- Regular key rotation

</td>
</tr>
</table>

### Best Practices

**Private Key Management**
- Never commit private keys to version control
- Use hardware wallets for mainnet deployments
- Implement multi-sig for admin functions
- Maintain separate keys for testnet and mainnet

**Environment Configuration**
- All sensitive data in `.env` files (gitignored)
- Validate environment before deployment
- Use different keys per network
- Review `.env.example` for required variables

### Audit Status

| Phase | Status |
|-------|--------|
| Internal Security Review | ✓ Completed |
| External Audit | Scheduled |
| Bug Bounty Program | Planned Post-Launch |

### Reporting Vulnerabilities

To report security issues, please use GitHub's security advisory feature or open an issue with the `security` label.

---

## Roadmap

### Phase 1: Core Platform

**Current Focus**

- Smart contract deployment on BNB Chain
- Basic AI resolver implementation
- Testnet launch and initial testing

### Phase 2: Enhancement

**Q2 2025**

- X402 payment gateway integration
- Decentralized application (dApp) interface
- Custom AI tool registry using Response API
- Expanded market mechanics:
  - MultiChoiceMarket (3-8 discrete outcomes)
  - LimitOrderMarket (professional trading)
  - PooledLiquidityMarket (concentrated liquidity)
  - DependentMarket (cascading settlements)
  - BracketMarket (range predictions)
  - TrendMarket (time-weighted outcomes)

### Phase 3: Developer Tools

**Q3 2025**

- React SDK for market integration
- Boilerplate templates with full SDK
- Comprehensive API documentation
- Developer tutorials and guides

### Phase 4: Governance & Ecosystem

**Q4 2025**

- HORIZON token governance launch
- Community proposals system
- Protocol parameter voting
- Third-party integrations and partnerships

---

## Contributing

We welcome contributions from the community. Please follow our development guidelines:

### Contribution Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes with tests
4. Ensure all tests pass (`forge test`)
5. Format your code (`forge fmt` or `go fmt`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Development Standards

<table>
<tr>
<td width="50%">

**Smart Contracts**
- Follow Solidity style guide
- Maintain >95% test coverage
- Document with NatSpec comments
- Run formatter before committing

</td>
<td width="50%">

**Go Code**
- Follow Go best practices
- Add unit tests for new features
- Use meaningful variable names
- Run `go fmt` before committing

</td>
</tr>
</table>

---

## Technology Stack

<div align="center">

**Built With Industry-Leading Tools**

[![Foundry](https://img.shields.io/badge/Foundry-Ethereum%20Toolkit-black?style=for-the-badge)](https://github.com/foundry-rs/foundry)
[![OpenZeppelin](https://img.shields.io/badge/OpenZeppelin-Smart%20Contracts-4E5EE4?style=for-the-badge)](https://github.com/OpenZeppelin/openzeppelin-contracts)
[![Go Ethereum](https://img.shields.io/badge/Go%20Ethereum-Blockchain-00ADD8?style=for-the-badge)](https://github.com/ethereum/go-ethereum)
[![OpenAI](https://img.shields.io/badge/OpenAI-AI%20Resolution-412991?style=for-the-badge)](https://openai.com/)

</div>

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## Support & Community

<div align="center">

**Get Help and Stay Connected**

[![Website](https://img.shields.io/badge/🌐-horizonoracles.com-blue?style=for-the-badge)](https://horizonoracles.com/)
[![Twitter](https://img.shields.io/badge/Twitter-@HorizonOracles-1DA1F2?style=for-the-badge&logo=twitter)](https://x.com/HorizonOracles)
[![GitHub Issues](https://img.shields.io/badge/GitHub-Issues-181717?style=for-the-badge&logo=github)](https://github.com/yourusername/project_gamma/issues)

For questions, feature requests, or bug reports, please open a GitHub issue or reach out on Twitter.

</div>

---

<div align="center">

**Project Gamma** • Decentralized Prediction Markets Powered by AI

Built on BNB Chain with ❤️ by the Horizon Oracles Team

[horizonoracles.com](https://horizonoracles.com/) • [@HorizonOracles](https://x.com/HorizonOracles)

</div>
