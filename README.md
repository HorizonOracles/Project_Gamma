# Project Gamma

A decentralized prediction market platform combining on-chain trading with AI-powered resolution on BNB Chain.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Core Features](#core-features)
- [Smart Contracts](#smart-contracts)
- [Getting Started](#getting-started)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)
- [Documentation](#documentation)
- [Security](#security)
- [License](#license)

## Overview

Project Gamma is a decentralized prediction market platform that enables users to create and trade binary outcome markets. The platform leverages AI-powered resolution for automated, evidence-based outcome determination while maintaining full on-chain transparency and decentralization.

### Key Components

1. **Smart Contracts**: Solidity contracts deployed on BNB Chain managing market creation, trading, and resolution
2. **AI Resolver**: Go-based backend service that gathers evidence and proposes resolutions using cryptographic signatures
3. **HORIZON Token**: Platform utility token providing fee discounts and governance participation

## Architecture

### System Design

The platform operates as a two-layer system:

**Layer 1 - On-Chain (BNB Chain)**
- Market creation and lifecycle management
- Automated market maker (AMM) for binary outcome trading
- Resolution state machine with dispute mechanism
- Fee collection and distribution
- Outcome token minting and redemption

**Layer 2 - Off-Chain (AI Resolver)**
- Evidence gathering through web search
- Multi-pass LLM analysis pipeline
- EIP-712 signature generation for proposals
- Automated proposal submission

### Smart Contract Architecture

```
MarketFactory (Central Registry)
    ├── MarketAMM (Per-Market Trading)
    │   └── Uses OutcomeToken (ERC-1155)
    ├── ResolutionModule (Dispute & Resolution)
    │   └── AIOracleAdapter (Signature Verification)
    ├── FeeSplitter (Fee Distribution)
    └── HorizonPerks (Fee Tier Management)
```

## Core Features

### Market Trading

- **Automated Market Maker**: Constant product formula (x*y=k) for price discovery
- **Binary Outcomes**: YES/NO outcome tokens for each market
- **Liquidity Provision**: LP tokens for liquidity providers with proportional fee earnings
- **Slippage Protection**: Minimum output requirements on all trades

### Fee Structure

The platform implements a dynamic fee model based on HORIZON token holdings:

- **User Fee**: Constant 2% on all trades
- **Protocol/Creator Split**: Variable based on trader's HORIZON balance
  - Default (0 HORIZON): 10% protocol / 90% creator
  - Tier 4 (500K+ HORIZON): 2% protocol / 98% creator

This model incentivizes market creators to attract HORIZON token holders while maintaining simple, predictable fees for users.

### Resolution System

**Optimistic Resolution**
- AI proposes outcomes with evidence
- 48-hour dispute window
- Proposer stakes collateral (refunded if accepted)

**Dispute Mechanism**
- Anyone can dispute with counter-evidence
- Disputer stakes higher collateral
- Manual arbitration for disputed outcomes

**Security Features**
- EIP-712 signature verification
- Multi-signer support for redundancy
- Evidence hash validation
- Stake-based incentive alignment

### HORIZON Token Utility

1. **Fee Optimization**: Higher holdings reduce protocol's share of fees
2. **Market Creation**: Required stake for market creation (refunded after resolution)
3. **Governance**: Future governance participation rights
4. **Creator Incentives**: Attracts high-value traders to markets

## Smart Contracts

### Core Contracts

| Contract | Description | Key Functions |
|----------|-------------|---------------|
| `MarketFactory` | Central registry for market creation and management | `createMarket()`, `updateMarketStatus()`, `refundCreatorStake()` |
| `MarketAMM` | Constant product AMM for binary outcome trading | `buyYes()`, `buyNo()`, `addLiquidity()`, `removeLiquidity()` |
| `ResolutionModule` | Resolution lifecycle and dispute management | `proposeResolution()`, `dispute()`, `finalize()` |
| `AIOracleAdapter` | EIP-712 signature verification for AI proposals | `proposeAI()`, `addSigner()`, `removeSigner()` |
| `OutcomeToken` | ERC-1155 for outcome shares and redemption | `mintOutcome()`, `burnOutcome()`, `redeem()` |
| `FeeSplitter` | Fee distribution between protocol and creators | `distribute()`, `claimProtocolFees()`, `claimCreatorFees()` |
| `HorizonPerks` | Fee tier calculation based on HORIZON holdings | `feeBpsFor()`, `protocolBpsFor()`, `addTier()` |
| `HorizonToken` | ERC-20 platform utility token | Standard ERC-20 functions |

### Contract Interactions

**Market Creation Flow**
```
User → MarketFactory.createMarket()
    → Stakes HORIZON tokens
    → Deploys new MarketAMM instance
    → Registers with FeeSplitter
    → Registers with OutcomeToken
    → Returns Market ID
```

**Trading Flow**
```
Trader → MarketAMM.buyYes()
    → Checks HORIZON balance for fee tier
    → Transfers collateral from trader
    → Mints outcome token pairs
    → Executes CPMM swap
    → Distributes fees via FeeSplitter
    → Transfers outcome tokens to trader
```

**Resolution Flow**
```
AI Resolver → AIOracleAdapter.proposeAI()
    → Verifies EIP-712 signature
    → Validates evidence hash
    → Forwards to ResolutionModule
    → Starts dispute window
    → Waits 48 hours
    → Finalizes if no disputes
    → Enables winner redemptions
```

## Getting Started

### Prerequisites

- [Foundry](https://book.getfoundry.sh/) (for smart contract development)
- Go 1.24+ (for AI resolver)
- Node.js 18+ (optional, for additional tooling)
- BNB Chain testnet/mainnet access
- OpenAI API key (for AI resolver)

### Installation

**1. Clone the Repository**

```bash
git clone https://github.com/yourusername/project_gamma.git
cd project_gamma
```

**2. Install Smart Contract Dependencies**

```bash
cd contracts
forge install
```

**3. Configure Environment**

```bash
cp .env.example .env
# Edit .env with your configuration:
# - PRIVATE_KEY: Deployment wallet private key
# - BSC_RPC_URL: BNB Chain RPC endpoint
# - BSCSCAN_API_KEY: Block explorer API key (for verification)
```

**4. Run Tests**

```bash
forge test
```

### Local Development

**Compile Contracts**

```bash
cd contracts
forge build
```

**Run Tests with Gas Report**

```bash
forge test --gas-report
```

**Format Code**

```bash
forge fmt
```

**Generate Documentation**

```bash
forge doc
```

## Development

### Smart Contract Development

The contracts are built using Foundry, a modern Solidity development framework.

**Project Structure**

```
contracts/
├── src/                    # Contract source files
│   ├── MarketFactory.sol
│   ├── MarketAMM.sol
│   ├── ResolutionModule.sol
│   ├── AIOracleAdapter.sol
│   ├── OutcomeToken.sol
│   ├── FeeSplitter.sol
│   ├── HorizonPerks.sol
│   └── HorizonToken.sol
├── test/                   # Test files
│   ├── unit/              # Unit tests for individual contracts
│   └── integration/       # Integration tests for workflows
├── script/                # Deployment and utility scripts
│   └── Deploy.s.sol
└── docs/                  # Additional documentation
```

### AI Resolver Development

The AI resolver is built in Go and provides automated resolution proposals.

**Building the Resolver**

```bash
cd ai-resolver
go build -o bin/ai-resolver ./cmd/server
```

**Running Tests**

```bash
go test ./...
```

**Configuration**

Edit `ai-resolver/.env`:
```env
# Blockchain Configuration
RPC_URL=https://bsc-testnet.publicnode.com
PRIVATE_KEY=your_signer_private_key

# Contract Addresses (from deployment)
MARKET_FACTORY_ADDRESS=0x...
RESOLUTION_MODULE_ADDRESS=0x...
AI_ORACLE_ADAPTER_ADDRESS=0x...

# API Keys
OPENAI_API_KEY=your_openai_key
SEARCH_API_KEY=your_search_api_key

# Server Configuration
PORT=8080
```

### Creating New Smart Contracts

This guide walks you through creating and testing new smart contracts using Foundry's built-in EVM.

**Step 1: Create the Contract**

Create a new contract file in `contracts/src/`:

```bash
cd contracts/src
touch MyNewContract.sol
```

Example contract structure:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract MyNewContract is Ownable, ReentrancyGuard {
    // State variables
    uint256 public myValue;
    
    // Events
    event ValueUpdated(uint256 oldValue, uint256 newValue);
    
    // Constructor
    constructor(uint256 initialValue) Ownable(msg.sender) {
        myValue = initialValue;
    }
    
    // Functions
    function updateValue(uint256 newValue) external onlyOwner {
        uint256 oldValue = myValue;
        myValue = newValue;
        emit ValueUpdated(oldValue, newValue);
    }
}
```

**Step 2: Create the Test File**

Create a test file in `contracts/test/unit/`:

```bash
cd contracts/test/unit
touch MyNewContract.t.sol
```

Example test structure:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../src/MyNewContract.sol";

contract MyNewContractTest is Test {
    MyNewContract public myContract;
    address public owner;
    address public user;
    
    function setUp() public {
        owner = address(this);
        user = makeAddr("user");
        
        // Deploy contract
        myContract = new MyNewContract(100);
    }
    
    function test_InitialValue() public {
        assertEq(myContract.myValue(), 100);
    }
    
    function test_UpdateValue() public {
        // Expect event emission
        vm.expectEmit(true, true, true, true);
        emit MyNewContract.ValueUpdated(100, 200);
        
        // Update value
        myContract.updateValue(200);
        
        // Assert new value
        assertEq(myContract.myValue(), 200);
    }
    
    function test_UpdateValue_OnlyOwner() public {
        // Prank as non-owner
        vm.prank(user);
        
        // Expect revert
        vm.expectRevert();
        myContract.updateValue(300);
    }
    
    function testFuzz_UpdateValue(uint256 randomValue) public {
        myContract.updateValue(randomValue);
        assertEq(myContract.myValue(), randomValue);
    }
}
```

**Step 3: Compile the Contract**

```bash
cd contracts
forge build
```

If compilation fails, check for:
- Correct Solidity version pragma
- Missing imports
- Syntax errors

**Step 4: Run Tests**

Run your specific test file:

```bash
forge test --match-contract MyNewContractTest
```

Run with verbosity to see detailed output:

```bash
forge test --match-contract MyNewContractTest -vvvv
```

Run with gas reporting:

```bash
forge test --match-contract MyNewContractTest --gas-report
```

**Step 5: Test Coverage**

Generate coverage for your contract:

```bash
forge coverage --match-contract MyNewContract
```

Aim for >95% coverage on all new contracts.

**Step 6: Integration Testing**

Create integration tests in `contracts/test/integration/`:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../src/MyNewContract.sol";
import "../../src/MarketFactory.sol";

contract MyNewContractIntegrationTest is Test {
    MyNewContract public myContract;
    MarketFactory public factory;
    
    function setUp() public {
        // Deploy all required contracts
        myContract = new MyNewContract(100);
        factory = new MarketFactory(/* params */);
    }
    
    function test_IntegrationWithFactory() public {
        // Test interaction between contracts
    }
}
```

**Understanding Foundry's EVM**

Foundry uses an isolated EVM instance for each test:

- **Clean State**: Each test starts with a fresh blockchain state
- **No Docker Required**: Foundry's EVM runs natively in your test process
- **Fast Execution**: Tests run in milliseconds without network latency
- **Cheatcodes**: Use `vm.*` functions to manipulate EVM state

**Common Foundry Cheatcodes**

```solidity
// Time manipulation
vm.warp(block.timestamp + 1 days);  // Fast forward time
vm.roll(block.number + 100);        // Fast forward blocks

// Identity manipulation
vm.prank(address);                  // Next call from address
vm.startPrank(address);             // All calls from address
vm.stopPrank();                     // Stop pranking

// Expectation
vm.expectRevert();                  // Expect next call to revert
vm.expectEmit(true, true, true, true);  // Expect event

// Balance manipulation
vm.deal(address, 100 ether);        // Give address ETH

// Storage manipulation
vm.store(address, slot, value);     // Write to storage slot
```

**Best Practices**

1. **Test Structure**
   - Use `setUp()` for common initialization
   - One assertion per test when possible
   - Use descriptive test names: `test_FunctionName_Scenario()`

2. **Coverage Goals**
   - Unit tests: Test each function independently
   - Integration tests: Test contract interactions
   - Edge cases: Test boundary conditions
   - Fuzz tests: Test with random inputs

3. **Gas Optimization**
   - Run `forge test --gas-report` regularly
   - Optimize high-frequency functions
   - Use `forge snapshot` to track gas changes

4. **Documentation**
   - Add NatSpec comments to all public functions
   - Document complex logic inline
   - Keep README updated with new contracts

**Debugging Failed Tests**

If a test fails:

1. Run with maximum verbosity:
   ```bash
   forge test --match-test test_MyFailingTest -vvvv
   ```

2. Check the trace for:
   - Revert reasons
   - Unexpected state changes
   - Gas issues

3. Use `console.log()` for debugging:
   ```solidity
   import "forge-std/console.sol";
   
   function test_Debug() public {
       console.log("Value:", myContract.myValue());
   }
   ```

**Running All Tests**

Before committing, always run the full test suite:

```bash
forge test
```

Expected output:
```
[PASS] test_Function1() (gas: 12345)
[PASS] test_Function2() (gas: 23456)
Test result: ok. 252 passed; 0 failed; finished in 2.34s
```

## Testing

### Unit Tests

Test individual contract functionality:

```bash
cd contracts
forge test --match-path "test/unit/*.sol"
```

### Integration Tests

Test complete workflows:

```bash
forge test --match-path "test/integration/*.sol"
```

### Test Coverage

Generate coverage report:

```bash
forge coverage
```

### Gas Optimization

Analyze gas usage:

```bash
forge test --gas-report
```

### Specific Test Patterns

Run tests matching a pattern:

```bash
forge test --match-test "testBuyYes"
forge test --match-contract "MarketAMMTest"
```

### Verbose Output

Debug test failures:

```bash
forge test -vvvv
```

## Deployment

### Testnet Deployment (BSC Testnet)

**1. Configure Environment**

Ensure `.env` is properly configured with testnet RPC URL and a funded wallet.

**2. Deploy Contracts**

```bash
cd contracts
forge script script/Deploy.s.sol --rpc-url bsc_testnet --broadcast --verify
```

**3. Save Deployment Addresses**

The script will output deployed contract addresses. Save these for AI resolver configuration.

**4. Verify Contracts (if not auto-verified)**

```bash
forge verify-contract <ADDRESS> <CONTRACT_NAME> --chain bsc-testnet
```

### Mainnet Deployment (BSC Mainnet)

**Security Checklist**
- [ ] All tests passing
- [ ] Security audit completed
- [ ] Deployment script reviewed
- [ ] Multi-sig wallet configured (if applicable)
- [ ] Sufficient BNB for deployment gas
- [ ] `.env` points to mainnet RPC
- [ ] Private key for deployment wallet is secure

**Deploy to Mainnet**

```bash
forge script script/Deploy.s.sol --rpc-url bsc_mainnet --broadcast --verify
```

**Post-Deployment**
1. Transfer ownership to multi-sig (if applicable)
2. Configure AI resolver signers
3. Set initial fee parameters
4. Test with small trades before public announcement

### Deployment Script Details

The deployment script (`Deploy.s.sol`) performs the following operations:

1. Deploy `HorizonToken` (if not already deployed)
2. Deploy `OutcomeToken`
3. Deploy `FeeSplitter` with protocol treasury address
4. Deploy `HorizonPerks` with default fee tiers
5. Deploy `AIOracleAdapter` with initial signers
6. Deploy `ResolutionModule`
7. Deploy `MarketFactory`
8. Configure authorizations between contracts
9. Output all contract addresses

## Documentation

### Contract Documentation

- [Contracts README](./contracts/README.md) - Detailed contract specifications
- [Deployment Guide](./contracts/docs/DEPLOYMENT.md) - Step-by-step deployment instructions
- [Security Considerations](./contracts/docs/SECURITY.md) - Security best practices and considerations
- [Roles & Permissions](./contracts/docs/ROLES.md) - Access control documentation

### AI Resolver Documentation

- [AI Resolver README](./ai-resolver/README.md) - Backend service architecture and API
- [Resolution Pipeline](./ai-resolver/docs/PIPELINE.md) - Evidence gathering and analysis process

### API Documentation

Generated documentation available after running:

```bash
cd contracts
forge doc --serve
```

Access at `http://localhost:3000`

## Security

### Best Practices

**Private Key Management**
- Never commit private keys to version control
- Use hardware wallets for mainnet deployments
- Implement multi-sig for admin functions
- Rotate keys regularly

**Environment Configuration**
- All sensitive data in `.env` files (gitignored)
- Use separate keys for testnet and mainnet
- Review `.env.example` for required variables
- Validate environment before deployment

**Smart Contract Security**
- Contracts use OpenZeppelin battle-tested libraries
- ReentrancyGuard on all state-changing functions
- Pausable functionality for emergency stops
- EIP-712 signatures for off-chain/on-chain verification
- Comprehensive test coverage (>95%)

**AI Resolver Security**
- Multi-signer support for redundancy
- Evidence hash validation
- Signature expiration (24-hour window)
- Rate limiting on proposal submissions
- Secure key storage

### Audit Status

**Current Status**: Pre-audit

**Planned Audits**
- Internal security review: Completed
- External audit: Scheduled
- Bug bounty program: Planned post-launch

### Reporting Security Issues

To report security vulnerabilities, please submit an issue on GitHub with the label `security`. 

For sensitive disclosures, you can also reach out directly through GitHub's security advisory feature.

## Project Structure

```
project_gamma/
├── contracts/              # Smart contracts
│   ├── src/               # Contract source files
│   │   ├── MarketFactory.sol
│   │   ├── MarketAMM.sol
│   │   ├── ResolutionModule.sol
│   │   ├── AIOracleAdapter.sol
│   │   ├── OutcomeToken.sol
│   │   ├── FeeSplitter.sol
│   │   ├── HorizonPerks.sol
│   │   └── HorizonToken.sol
│   ├── test/              # Test files
│   │   ├── unit/         # Unit tests
│   │   └── integration/  # Integration tests
│   ├── script/            # Deployment scripts
│   │   └── Deploy.s.sol
│   ├── docs/              # Contract documentation
│   └── foundry.toml       # Foundry configuration
├── ai-resolver/           # AI resolution service
│   ├── cmd/              # Application entry points
│   │   └── server/
│   ├── internal/         # Internal packages
│   │   ├── ai/          # LLM integration
│   │   ├── blockchain/  # Chain interaction
│   │   ├── evidence/    # Evidence gathering
│   │   └── signature/   # EIP-712 signing
│   ├── pkg/              # Public packages
│   └── go.mod
└── README.md             # This file
```

## Contributing

We welcome contributions to Project Gamma. Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

**Smart Contracts**
- Follow Solidity style guide
- Maintain test coverage above 95%
- Document all public functions with NatSpec
- Run `forge fmt` before committing

**Go Code**
- Follow Go best practices
- Add unit tests for new functionality
- Use meaningful variable names
- Run `go fmt` before committing

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) file for details.

## Support

For questions, issues, or contributions:

- Open a [GitHub Issue](https://github.com/yourusername/project_gamma/issues)
- Join our [Discord community](https://discord.gg/projectgamma) (coming soon)
- Read the [Documentation](./contracts/README.md)

## Acknowledgments

Built with:
- [Foundry](https://github.com/foundry-rs/foundry) - Ethereum development toolkit
- [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts) - Secure smart contract library
- [Go Ethereum](https://github.com/ethereum/go-ethereum) - Go implementation of Ethereum
- [OpenAI](https://openai.com/) - AI capabilities

## Roadmap

**Phase 1: Core Platform** (Current)
- Smart contract deployment
- Basic AI resolver
- Testnet launch

**Phase 2: Enhancement**
- X402 payment gateway integration
- Decentralized application (dApp) interface
- Custom AI tool registry using Response API (BSCScan, PancakeSwap TWAP for real-time on-chain data)
- Expanded market mechanics:
  - MultiChoiceMarket - Support for 3-8 discrete outcomes
  - LimitOrderMarket - Professional trading with order matching
  - PooledLiquidityMarket - Enhanced AMM with concentrated liquidity
  - DependentMarket - Outcomes that trigger related market settlements
  - BracketMarket - Predictions within defined value ranges
  - TrendMarket - Time-weighted outcome predictions

**Phase 3: Developer Tools**
- React SDK for market integration
- Boilerplate templates for building markets with full SDK integration
- Developer documentation and API reference

**Phase 4: Governance & Ecosystem**
- HORIZON token governance
- Community proposals and protocol parameter voting
- Third-party integrations and partnerships
