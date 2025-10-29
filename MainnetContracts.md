<div align="center">

![Project Gamma Banner](assets/banner1.png)

# Project Gamma

**Decentralized Prediction Markets Powered by AI**

[![Website](https://img.shields.io/badge/Website-horizonoracles.com-blue?style=flat-square)](https://horizonoracles.com/)
[![Twitter](https://img.shields.io/badge/Twitter-@HorizonOracles-1DA1F2?style=flat-square&logo=twitter)](https://x.com/HorizonOracles)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Discord](https://img.shields.io/badge/Discord-Join%20Us-5865F2?style=flat-square&logo=discord&logoColor=white)](https://discord.com/invite/TuUHwwKjHh)

**Built on** <a href="https://www.bnbchain.org/"><img src="assets/bnb-chain-full-binance-smart-chain-logo-1-1024x180.png" alt="BNB Chain" height="30"></a>

*Combining on-chain trading with AI-powered resolution for transparent, automated prediction markets*

---

[Overview](#overview) • [Live Deployment](#live-deployment) • [Architecture](#architecture) • [Getting Started](#getting-started) • [Documentation](#documentation)

---

</div>

## 🌟 Live Deployment

**Network:** BNB Smart Chain Mainnet (Chain ID: 56)

### Core Contract Addresses

<table>
<tr>
<th>Contract</th>
<th>Address</th>
<th>BscScan</th>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td><code>0xf3f68a76d42679e8b3371bc26b75f7f26e97a10c</code></td>
<td><a href="https://bscscan.com/address/0xf3f68a76d42679e8b3371bc26b75f7f26e97a10c">View Contract</a></td>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td><code>0x29cfa909515cece341f6e8149026c9d7861a04bf</code></td>
<td><a href="https://bscscan.com/address/0x29cfa909515cece341f6e8149026c9d7861a04bf">View Contract</a></td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td><code>0x72f84681aa0dc8db53e87ed507a4d6651b1c312d</code></td>
<td><a href="https://bscscan.com/address/0x72f84681aa0dc8db53e87ed507a4d6651b1c312d">View Contract</a></td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>0x31709748cc9030e86e71570442fa762c851950b3</code></td>
<td><a href="https://bscscan.com/address/0x31709748cc9030e86e71570442fa762c851950b3">View Contract</a></td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>0x7cb3a3c58f7ea49cc6af8be5228ac58b04787b09</code></td>
<td><a href="https://bscscan.com/address/0x7cb3a3c58f7ea49cc6af8be5228ac58b04787b09">View Contract</a></td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>0x5407f3937c81a01e783b9e99fbac624220a534eb</code></td>
<td><a href="https://bscscan.com/address/0x5407f3937c81a01e783b9e99fbac624220a534eb">View Contract</a></td>
</tr>
<tr>
<td><strong>AIOracleAdapter</strong></td>
<td><code>0x30827513096b63f09c0c24f574933b43e222c5d7</code></td>
<td><a href="https://bscscan.com/address/0x30827513096b63f09c0c24f574933b43e222c5d7">View Contract</a></td>
</tr>
</table>

**Deployment Cost:** ~0.0008 BNB (~$0.50 USD)

---

## Overview

Project Gamma is a decentralized prediction market platform that enables users to create and trade binary outcome markets with automated, evidence-based resolution. The platform leverages AI-powered resolution for automated outcome determination while maintaining full on-chain transparency and decentralization on BNB Chain.

### What You Can Do

<table>
<tr>
<td width="20%" align="center">

**Create Markets**

Stake HORIZON tokens to launch prediction markets on any topic

</td>
<td width="20%" align="center">

**Trade Positions**

Buy and sell YES/NO outcome tokens using our AMM

</td>
<td width="20%" align="center">

**Provide Liquidity**

Earn fees by providing liquidity to market pools

</td>
<td width="20%" align="center">

**Resolve Markets**

Participate in decentralized resolution with disputes

</td>
<td width="20%" align="center">

**Earn Rewards**

Reduced fees by holding HORIZON tokens

</td>
</tr>
</table>

---

## Core Features

### AI-Powered Resolution

<table>
<tr>
<td width="50%">

**Automated Analysis**
- Multi-signature verification system
- Evidence-based outcome determination
- Public attestation on IPFS/Arweave
- 48-hour dispute window

</td>
<td width="50%">

**Community Oversight**
- Dispute mechanism with bond requirements
- Economic incentives for accuracy
- Arbitrator as final backstop
- Transparent on-chain process

</td>
</tr>
</table>

### Dynamic Fee Structure

The platform implements a tiered fee model that incentivizes HORIZON token ownership:

| Tier | HORIZON Holdings | Trading Fee | Protocol Share | Creator Share |
|------|------------------|-------------|----------------|---------------|
| 0 | 0 | 4.0% | 10% | 90% |
| 1 | 10,000+ | 3.0% | 8% | 92% |
| 2 | 50,000+ | 2.0% | 6% | 94% |
| 3 | 100,000+ | 1.5% | 4% | 96% |
| 4 | 500,000+ | 1.25% | 3% | 97% |
| 5 | 1,000,000+ | 1.0% | 2% | 98% |

**Key Benefits:**
- Simple UX: Clear fee structure
- Creator Incentive: Earn more from HORIZON holders
- Token Utility: Strong incentive to hold HORIZON
- Growth Model: Protocol subsidizes whale retention

### Automated Market Maker (AMM)

<table>
<tr>
<td width="33%">

**Constant Product Formula**

Efficient x·y=k price discovery for binary outcomes

</td>
<td width="33%">

**Slippage Protection**

User-configurable minimum output requirements

</td>
<td width="33%">

**LP Token System**

Proportional fee earnings for liquidity providers

</td>
</tr>
</table>

### Security First

<table>
<tr>
<td width="25%" align="center">

**Battle-Tested**

OpenZeppelin libraries and patterns

</td>
<td width="25%" align="center">

**Comprehensive Tests**

252 passing tests with >95% coverage

</td>
<td width="25%" align="center">

**Emergency Controls**

Pausable functions and access control

</td>
<td width="25%" align="center">

**Reentrancy Protected**

Guards on all critical functions

</td>
</tr>
</table>

---

## Architecture

### System Design

```
┌─────────────────────────────────────────────────────────────┐
│                     USER INTERFACE                          │
│              (Web3 Frontend / Smart Wallets)                │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    MARKET FACTORY                           │
│         (Central Hub for Market Creation)                   │
│  • Create Markets  • Manage Stakes  • Deploy AMMs          │
└─────────────────────────────────────────────────────────────┘
         │                    │                    │
         ▼                    ▼                    ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│  MARKET AMM  │    │  MARKET AMM  │    │  MARKET AMM  │
│              │    │              │    │              │
│ • Trading    │    │ • Liquidity  │    │ • Pricing    │
│ • Fees       │    │ • LP Tokens  │    │ • Slippage   │
└──────────────┘    └──────────────┘    └──────────────┘
         │                    │                    │
         └────────────────────┴────────────────────┘
                              │
                              ▼
        ┌─────────────────────────────────────────┐
        │          SHARED COMPONENTS              │
        ├─────────────────────────────────────────┤
        │ OutcomeToken  │ ERC-1155 YES/NO tokens │
        │ FeeSplitter   │ Revenue distribution   │
        │ HorizonPerks  │ Fee tier management    │
        │ HorizonToken  │ Governance & staking   │
        └─────────────────────────────────────────┘
                              │
                              ▼
        ┌─────────────────────────────────────────┐
        │      RESOLUTION LAYER                   │
        ├─────────────────────────────────────────┤
        │ ResolutionModule │ Dispute handling    │
        │ AIOracleAdapter  │ AI oracle integration│
        └─────────────────────────────────────────┘
```

### Contract Overview

<table>
<tr>
<th>Contract</th>
<th>Purpose</th>
<th>Key Responsibilities</th>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td>Central Registry</td>
<td>Market creation, status management, creator stake handling</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td>Trading Engine</td>
<td>Binary outcome trading via constant product AMM</td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td>Token Standard</td>
<td>ERC-1155 implementation for outcome shares</td>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td>Utility Token</td>
<td>ERC-20 platform token for staking and fees</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td>Resolution Manager</td>
<td>Lifecycle management, dispute handling, finalization</td>
</tr>
<tr>
<td><strong>AIOracleAdapter</strong></td>
<td>Signature Verifier</td>
<td>EIP-712 verification for AI-generated proposals</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td>Fee Distribution</td>
<td>Protocol and creator fee allocation</td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td>Fee Calculator</td>
<td>Dynamic fee tier computation</td>
</tr>
</table>

---

## Getting Started

### Prerequisites

<table>
<tr>
<td width="50%">

**Development Tools**
- Node.js v18+
- Foundry toolkit
- Git version control

</td>
<td width="50%">

**For Users**
- Web3 wallet (MetaMask, etc.)
- BNB for gas fees
- Collateral tokens (USDC, etc.)
- HORIZON tokens (optional)

</td>
</tr>
</table>

### Installation

**1. Clone Repository**

```bash
git clone https://github.com/HorizonOracles/Project_Gamma.git
cd Project_Gamma/contracts
```

**2. Install Dependencies**

```bash
# Foundry dependencies
forge install

# Node dependencies (optional)
npm install
```

**3. Configure Environment**

```bash
cp .env.example .env
# Edit .env with your configuration
```

**4. Compile Contracts**

```bash
forge build
```

**5. Run Tests**

```bash
forge test
```

---

## Smart Contract Interactions

### Creating a Market

<details>
<summary><strong>Example: Create Prediction Market</strong></summary>

```solidity
// 1. Approve HORIZON tokens for staking
horizonToken.approve(marketFactory, 10000 ether);

// 2. Create market
uint256 marketId = marketFactory.createMarket(
    collateralToken,                    // ERC20 token (e.g., USDC)
    "Will BTC reach $100k by 2025?",   // Market question
    "Crypto",                           // Category
    "ipfs://...",                       // Metadata URI
    block.timestamp + 30 days           // Close time
);
```

</details>

### Trading on a Market

<details>
<summary><strong>Example: Buy and Sell Outcome Tokens</strong></summary>

```solidity
// Get market AMM address
address ammAddress = marketFactory.getMarketAMM(marketId);
MarketAMM amm = MarketAMM(ammAddress);

// Approve collateral
collateralToken.approve(ammAddress, 1000 ether);

// Buy YES tokens
uint256 yesTokens = amm.buyYes(
    1000 ether,  // Collateral amount
    0            // Minimum tokens (slippage protection)
);

// Sell YES tokens
uint256 collateralOut = amm.sellYes(
    yesTokens,   // Token amount to sell
    0            // Minimum collateral out
);
```

</details>

### Providing Liquidity

<details>
<summary><strong>Example: Add and Remove Liquidity</strong></summary>

```solidity
// Approve collateral
collateralToken.approve(ammAddress, 10000 ether);

// Add liquidity
uint256 lpTokens = amm.addLiquidity(
    10000 ether,  // Collateral amount
    0             // Minimum LP tokens
);

// Remove liquidity later
uint256 collateralReturned = amm.removeLiquidity(
    lpTokens,     // LP tokens to burn
    0             // Minimum collateral out
);
```

</details>

### Resolving a Market

<details>
<summary><strong>Example: Propose and Finalize Resolution</strong></summary>

```solidity
// 1. Approve bond tokens
horizonToken.approve(resolutionModule, 1000 ether);

// 2. Propose resolution
resolutionModule.proposeResolution(
    marketId,
    0,                  // 0 = YES, 1 = NO
    "ipfs://evidence"   // Evidence URI
);

// 3. Wait 48-hour dispute window...

// 4. Finalize (if no disputes)
resolutionModule.finalize(marketId);
```

</details>

### Claiming Winnings

<details>
<summary><strong>Example: Redeem Winning Tokens</strong></summary>

```solidity
// Redeem winning tokens for collateral (1:1 ratio)
uint256 payout = outcomeToken.redeem(
    marketId,
    0,          // Outcome ID (0 = YES, 1 = NO)
    amount      // Amount of tokens to redeem
);
```

</details>

---

## Development

### Test Suite

The protocol includes comprehensive test coverage:

<table>
<tr>
<td width="25%" align="center">

**252 Tests**

All passing with >95% coverage

</td>
<td width="25%" align="center">

**Unit Tests**

Individual contract functionality

</td>
<td width="25%" align="center">

**Integration Tests**

End-to-end workflows

</td>
<td width="25%" align="center">

**Fuzz Tests**

Edge case discovery

</td>
</tr>
</table>

**Run Tests:**

```bash
# All tests
forge test

# With verbosity
forge test -vvv

# Specific test
forge test --match-test testMarketCreation

# Coverage report
forge coverage
```

### Local Development

**Start Local Node:**

```bash
anvil
```

**Deploy Locally:**

```bash
forge script script/Deploy.s.sol \
  --rpc-url http://localhost:8545 \
  --broadcast
```

---

## Protocol Parameters

### Market Creation

<table>
<tr>
<td width="50%">

**Minimum Creator Stake**
- Default: 10,000 HORIZON
- Refundable after resolution
- Configurable by governance

</td>
<td width="50%">

**Stake Lockup**
- Locked until market resolution
- Prevents spam markets
- Incentivizes proper resolution

</td>
</tr>
</table>

### Resolution System

<table>
<tr>
<td width="33%">

**Minimum Bond**

1,000 HORIZON tokens

</td>
<td width="33%">

**Dispute Window**

48 hours (172,800 seconds)

</td>
<td width="33%">

**Arbitrator**

Multi-sig governance

</td>
</tr>
</table>

### Token Economics

<table>
<tr>
<td width="50%">

**HORIZON Token**
- Max Supply: 10 billion
- Initial Supply: 100 million
- Standard: ERC-20

</td>
<td width="50%">

**Fee Distribution**
- Creator: 90-98%
- Protocol: 2-10%
- LP Fee: 0.3%

</td>
</tr>
</table>

---

## Security

### Security Measures

<table>
<tr>
<td width="33%">

**Smart Contract Security**
- OpenZeppelin libraries
- ReentrancyGuard protection
- Pausable emergency stops
- Access control

</td>
<td width="33%">

**Cryptographic Security**
- EIP-712 typed signatures
- Evidence hash validation
- Multi-signer redundancy
- Replay protection

</td>
<td width="33%">

**Testing & Audits**
- 252 passing tests
- >95% coverage
- Internal review complete
- External audit scheduled

</td>
</tr>
</table>

### Audit Status

| Phase | Status | Details |
|-------|--------|---------|
| Internal Security Review | ✓ Completed | Core team review |
| External Audit | Scheduled | Third-party firm |
| Bug Bounty Program | Planned | Post-launch |

### Reporting Security Issues

Please use GitHub's security advisory feature or email developers@horizonoracles.com

---

## Documentation

### Developer Resources

<table>
<tr>
<td width="50%">

**Smart Contracts**
- [Contract Specifications](./contracts/README.md)
- [Security Documentation](./contracts/docs/SECURITY.md)
- [Roles & Permissions](./contracts/docs/ROLES.md)
- [Deployment Guide](./contracts/docs/DEPLOYMENT.md)

</td>
<td width="50%">

**AI Resolver**
- [Backend Architecture](./ai-resolver/README.md)
- [API Documentation](./ai-resolver/README.md#api-endpoints)
- [Docker Setup](./ai-resolver/README.md#deployment)

</td>
</tr>
</table>

---

## Roadmap

### Phase 1: Core Platform (✓ Completed)

- [x] Smart contract development
- [x] Comprehensive testing suite
- [x] BSC mainnet deployment
- [x] Contract verification

### Phase 2: Enhancement (Q1 2025)

- [ ] Web interface launch
- [ ] External security audit
- [ ] Community testing program
- [ ] Bug bounty program

### Phase 3: Growth (Q2 2025)

- [ ] Mobile app release
- [ ] Advanced analytics dashboard
- [ ] DAO governance implementation
- [ ] Cross-chain expansion

### Phase 4: Scale (Q3-Q4 2025)

- [ ] API for third-party integrations
- [ ] Market maker partnerships
- [ ] Institutional features
- [ ] Global market categories

---

## Supported Networks

<table>
<tr>
<td width="20%" align="center">

**BNB Chain**

✅ Mainnet Live

</td>
<td width="20%" align="center">

**Ethereum**

⏳ Coming Soon

</td>
<td width="20%" align="center">

**Polygon**

⏳ Coming Soon

</td>
<td width="20%" align="center">

**Arbitrum**

⏳ Coming Soon

</td>
<td width="20%" align="center">

**Optimism**

⏳ Coming Soon

</td>
</tr>
</table>

---

## Contributing

We welcome contributions from the community! Please follow our development guidelines:

### Contribution Workflow

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes with tests
4. Ensure all tests pass: `forge test`
5. Format your code: `forge fmt`
6. Commit: `git commit -m 'Add amazing feature'`
7. Push and open a Pull Request

### Development Standards

- Follow Solidity style guide
- Maintain >95% test coverage
- Document with NatSpec comments
- No compiler warnings
- Gas optimization where reasonable

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

## Disclaimer

**⚠️ Important Notice**

This protocol is experimental software provided "as is" without warranties. Users should:

- Understand the risks of smart contract interactions
- Never invest more than they can afford to lose
- Conduct their own research before participating
- Be aware that markets can be volatile and unpredictable
- Know that past performance doesn't guarantee future results

The developers are not responsible for any losses incurred through the use of this protocol.

---

## Support & Community

<div align="center">

**Get Help and Stay Connected**

<table>
<tr>
<td align="center" width="25%">

**Business Inquiries**

Enterprise integrations and strategic collaborations

partnerships@horizonoracles.com

</td>
<td align="center" width="25%">

**Developer Support**

Technical support and integration assistance

developers@horizonoracles.com

</td>
<td align="center" width="25%">

**Media**

Press inquiries and media requests

press@horizonoracles.com

</td>
<td align="center" width="25%">

**Community**

Join our Discord server

[discord.gg/TuUHwwKjHh](https://discord.com/invite/TuUHwwKjHh)

</td>
</tr>
</table>

</div>

---

<div align="center">

**Project Gamma** • Decentralized Prediction Markets Powered by AI

Built on BNB Chain with ❤️ by the Horizon Oracles Team

[horizonoracles.com](https://horizonoracles.com/) • [@HorizonOracles](https://x.com/HorizonOracles)

*If you find this project interesting, please star the repository!*

</div>
