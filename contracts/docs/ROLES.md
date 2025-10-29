# Access Control & Roles Documentation

## Overview

This document defines all access control roles, permissions, and responsibilities in the BNB Chain Prediction Market protocol.

**Version:** Phase 5 Complete  
**Last Updated:** 2025-10-28

---

## Role Hierarchy

```
Protocol Owner (Multi-sig)
    ├── MarketFactory Owner
    ├── FeeSplitter Owner
    ├── OutcomeToken Owner
    ├── HorizonPerks Owner
    ├── HorizonToken Owner
    └── ResolutionModule Owner
        └── Arbitrator (DAO/Multi-sig)

Market Creators (Stakers)
    └── Create Markets, Claim Fees

AI Oracle Signers (Independent Validators)
    └── Sign Resolution Proposals

Protocol Treasury (Multi-sig)
    └── Claim Protocol Fees

Users (Anyone)
    └── Trade, Add Liquidity, Claim Winnings
```

---

## Role Definitions

### 1. Protocol Owner

**Description**: Administrator of all core contracts  
**Recommended**: 5-of-9 multi-sig wallet or DAO timelock  
**Transition Plan**: Start with deployer, transition to DAO within 3-6 months

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **MarketFactory** | `setMinCreatorStake(uint256)` | Adjust minimum stake for market creation | Low |
| **FeeSplitter** | `setProtocolTreasury(address)` | Update protocol treasury address | Medium |
| **FeeSplitter** | `updateFeeConfig(uint256, uint16, uint16)` | Change fee split per market | Medium |
| **OutcomeToken** | `setAMMAuthorization(address, bool)` | Authorize/revoke AMM contracts | High |
| **OutcomeToken** | `setResolutionAuthorization(address, bool)` | Authorize/revoke resolver contracts | High |
| **HorizonPerks** | `addTier(uint256, uint16)` | Add new fee tier | Low |
| **HorizonPerks** | `updateTier(uint256, uint256, uint16)` | Update existing tier | Low |
| **HorizonPerks** | `removeLastTier()` | Remove last tier | Low |
| **HorizonToken** | `addMinter(address)` | Authorize new minter | High |
| **HorizonToken** | `removeMinter(address)` | Revoke minter | Medium |
| **ResolutionModule** | `setMinBond(uint256)` | Set minimum bond for proposals | Low |
| **ResolutionModule** | `setDisputeWindow(uint256)` | Set dispute period duration | Medium |
| **ResolutionModule** | `setArbitrator(address)` | Set arbitrator address | **Critical** |
| **MarketAMM** | `pause()` / `unpause()` | Emergency pause trading | **Critical** |
| **MarketAMM** | `close()` | Close market early | High |

#### Responsibilities
- Monitor protocol health and security
- Respond to critical bugs or exploits
- Coordinate upgrades and migrations
- Manage multi-sig operations
- Transition to DAO governance

#### Limitations
- **Cannot** steal user funds
- **Cannot** manipulate market outcomes
- **Cannot** change resolved market results
- **Cannot** mint HORIZON tokens (only authorized minters can)

---

### 2. Arbitrator

**Description**: Final authority for disputed resolutions  
**Recommended**: DAO governance contract with time-lock (48h minimum)  
**Accountability**: All decisions on-chain and transparent

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **ResolutionModule** | `finalizeDisputed(uint256, uint8, bool)` | Resolve disputed markets | **Critical** |

#### Responsibilities
- Review evidence from disputers and proposers
- Make fair, unbiased resolution decisions
- Slash bonds of incorrect parties
- Maintain community trust

#### Decision Process
1. Review market question and criteria
2. Examine proposer's evidence (IPFS/Arweave)
3. Examine disputer's counter-evidence
4. Make determination based on facts
5. Execute `finalizeDisputed()` with outcome + slashing decision

#### Limitations
- **Cannot** resolve non-disputed markets (only ResolutionModule can finalize)
- **Cannot** change finalized results
- **Cannot** access user funds directly

#### Slashing Rules
```solidity
// Slash proposer if they were incorrect
finalizeDisputed(marketId, correctOutcome, slashProposer: true);
// Result: Proposer loses bond, Disputer gains bond + refund

// Slash disputer if they were incorrect
finalizeDisputed(marketId, proposerOutcome, slashProposer: false);
// Result: Disputer loses bond, Proposer gains bond + refund
```

---

### 3. Market Creators

**Description**: Users who stake HORIZON tokens to create prediction markets  
**Requirements**: Hold minimum stake (default: 100 HORIZON tokens)  
**Incentive**: Receive 90% of trading fees from their markets

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **MarketFactory** | `createMarket(...)` | Create new prediction market | Low |
| **MarketFactory** | `refundCreatorStake(uint256)` | Claim stake back after resolution | None |
| **FeeSplitter** | `claimCreatorFees(uint256, address)` | Claim earned trading fees | None |
| **FeeSplitter** | `claimCreatorFeesMultiple(...)` | Batch claim fees | None |

#### Market Creation Parameters
```solidity
function createMarket(
    string memory question,       // e.g., "Will BTC reach $100k by Dec 31, 2024?"
    string memory category,       // e.g., "Crypto", "Sports", "Politics"
    string[] memory aiEvidenceURIs, // IPFS/Arweave links for AI resolution
    uint256 closeTime,            // Unix timestamp when trading closes
    address collateralToken,      // USDC, DAI, or other ERC20
    uint256 minLiquidityAmount    // Minimum for initial LP
) external returns (uint256 marketId);
```

#### Responsibilities
- Write clear, unambiguous market questions
- Set appropriate close times (before event occurs)
- Provide evidence sources for AI resolution
- Add initial liquidity to markets (optional but recommended)
- Ensure question is resolvable objectively

#### Fee Structure
- **Earn**: 90-98% of all trading fees (higher % when attracting HORIZON holders)
- **Claim**: Anytime after fees accrue via `claimCreatorFees()`
- **Stake**: Locked until market resolves, then refundable

**Example Revenue Per $1000 Trade:**
- Trader with 0 HORIZON: Creator earns $18.00 (90% of $20 fee)
- Trader with 500K+ HORIZON: Creator earns $19.60 (98% of $20 fee)

This incentivizes creators to market their markets to HORIZON token holders!

#### Limitations
- **Cannot** cancel markets once created
- **Cannot** change market parameters after creation
- **Cannot** influence resolution outcome
- **Cannot** withdraw stake before resolution

---

### 4. AI Oracle Signers

**Description**: Independent entities that sign resolution proposals  
**Recommended**: 3-5 independent AI agents or trusted validators  
**Selection**: Appointed by Protocol Owner, should be diverse and reputable

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **AIOracleAdapter** | `proposeAI(...)` | Submit AI-signed resolution proposal | Medium |

#### Signature Requirements
```solidity
struct AIProposal {
    uint256 marketId;
    uint8 outcome;           // 0 = YES, 1 = NO
    string[] evidenceURIs;   // Public evidence (IPFS/Arweave)
    uint256 validFrom;       // Timestamp when signature becomes valid
    uint256 validUntil;      // Expiration timestamp
}
```

#### Signing Process
1. AI agent analyzes market question and gathers evidence
2. Determines outcome based on facts
3. Uploads evidence to IPFS/Arweave
4. Signs EIP-712 structured data with private key
5. Anyone can submit signed proposal to `AIOracleAdapter`

#### Responsibilities
- Provide accurate, unbiased resolutions
- Publish evidence publicly for transparency
- Use reliable data sources
- Respond to community concerns about resolutions
- Maintain operational security of signing keys

#### Adding/Removing Signers
```solidity
// Only Protocol Owner can manage signers
aiOracleAdapter.addSigner(newSignerAddress);    // Add new signer
aiOracleAdapter.removeSigner(oldSignerAddress); // Revoke signer
```

#### Limitations
- **Cannot** propose outcomes without valid signature
- **Cannot** override disputes (Arbitrator decides)
- **Cannot** reuse signatures (replay protection)
- **Cannot** access user funds

---

### 5. Protocol Treasury

**Description**: Multi-sig wallet that receives protocol fees  
**Recommended**: 3-of-5 multi-sig or DAO treasury  
**Purpose**: Fund development, marketing, security audits, and protocol growth

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **FeeSplitter** | `claimProtocolFees(address)` | Claim protocol fees for one token | None |
| **FeeSplitter** | `claimProtocolFeesMultiple(address[])` | Batch claim fees | None |

#### Fee Structure
- **Earn**: 2-10% of all trading fees (varies by trader's HORIZON holdings)
- **Tokens**: Can accumulate in any collateral token (USDC, DAI, etc.)
- **Claim**: Anytime via `claimProtocolFees()`

**Example Revenue Per $1000 Trade:**
- Trader with 0 HORIZON: Protocol earns $2.00 (10% of $20 fee)
- Trader with 500K+ HORIZON: Protocol earns $0.40 (2% of $20 fee)

The protocol subsidizes whale retention to grow the ecosystem.

#### Fund Usage (Recommended)
- 40%: Development and engineering
- 25%: Security audits and bug bounties
- 20%: Marketing and growth
- 10%: Liquidity incentives
- 5%: Operations and legal

#### Limitations
- **Cannot** claim creator fees
- **Cannot** change fee split (only Protocol Owner can)
- **Cannot** access market funds directly

---

### 6. Users (Traders & Liquidity Providers)

**Description**: Anyone interacting with the protocol  
**Requirements**: Hold collateral tokens for trading, HORIZON tokens for fee discounts (optional)

#### Permissions

| Contract | Function | Purpose | Risk Level |
|----------|----------|---------|------------|
| **MarketAMM** | `buyYes(uint256, uint256)` | Buy YES outcome tokens | None |
| **MarketAMM** | `buyNo(uint256, uint256)` | Buy NO outcome tokens | None |
| **MarketAMM** | `sellYes(uint256, uint256)` | Sell YES outcome tokens | None |
| **MarketAMM** | `sellNo(uint256, uint256)` | Sell NO outcome tokens | None |
| **MarketAMM** | `addLiquidity(uint256)` | Provide liquidity, earn LP tokens | None |
| **MarketAMM** | `removeLiquidity(uint256)` | Withdraw liquidity, burn LP tokens | None |
| **OutcomeToken** | `redeem(uint256, uint256)` | Claim winnings after resolution | None |
| **ResolutionModule** | `proposeResolution(uint256, uint8)` | Propose outcome with bond | Low |
| **ResolutionModule** | `dispute(uint256, uint8)` | Dispute proposal with bond | Medium |

#### Fee Tiers (Based on HORIZON Holdings)

**All users pay 2% trading fee. HORIZON holdings determine the protocol/creator split:**

| Tier | HORIZON Balance | User Fee | Protocol Share | Creator Share | Creator Gets* | Protocol Gets* |
|------|-----------------|----------|----------------|---------------|---------------|----------------|
| 0    | 0               | 2.00%    | 10%            | 90%           | $18.00        | $2.00          |
| 1    | 10,000          | 2.00%    | 8%             | 92%           | $18.40        | $1.60          |
| 2    | 50,000          | 2.00%    | 6%             | 94%           | $18.80        | $1.20          |
| 3    | 100,000         | 2.00%    | 4%             | 96%           | $19.20        | $0.80          |
| 4    | 500,000+        | 2.00%    | 2%             | 98%           | $19.60        | $0.40          |

_*Per $1,000 trade volume_

**Key Benefits:**
1. **Simple UX**: Users always pay 2% - easy to understand
2. **Creator Incentive**: Market creators earn MORE from HORIZON whale traders
3. **Token Utility**: Strong incentive for traders to hold HORIZON tokens
4. **Growth Model**: Protocol subsidizes whale retention to grow ecosystem

#### Responsibilities
- Set appropriate slippage tolerance (1-5% recommended)
- Verify market questions before trading
- Hold HORIZON tokens for fee discounts (optional)
- Claim winnings after markets resolve

#### Limitations
- **Cannot** trade after market closes
- **Cannot** redeem losing outcome tokens
- **Cannot** remove liquidity before resolution (frozen during resolution)
- **Cannot** cancel trades (slippage protection only)

---

## Access Control Patterns

### 1. Owner-Based (OpenZeppelin Ownable)

Used in: All core contracts

```solidity
modifier onlyOwner() {
    if (msg.sender != owner()) revert Unauthorized();
    _;
}
```

**Transition**: `transferOwnership(newOwner)` - should be 2-step with `acceptOwnership()`

---

### 2. Role-Based (Custom)

Used in: OutcomeToken, FeeSplitter

```solidity
// OutcomeToken
modifier onlyAuthorizedAMM() {
    if (!authorizedAMMs[msg.sender]) revert Unauthorized();
    _;
}

modifier onlyAuthorizedResolver() {
    if (!authorizedResolvers[msg.sender]) revert Unauthorized();
    _;
}

// FeeSplitter
modifier onlyFactory() {
    if (msg.sender != factory) revert Unauthorized();
    _;
}
```

---

### 3. Signer-Based (EIP-712)

Used in: AIOracleAdapter

```solidity
modifier validSigner(address signer) {
    if (!authorizedSigners[signer]) revert UnauthorizedSigner();
    _;
}
```

---

### 4. Bond-Based (Economic)

Used in: ResolutionModule

```solidity
// Anyone can propose if they post bond
function proposeResolution(uint256 marketId, uint8 outcome) external {
    bondToken.transferFrom(msg.sender, address(this), minBond);
    // ...
}
```

---

## Security Checklist

### Before Mainnet Deployment

- [ ] Transfer all contract ownership to multi-sig (3-of-5 minimum)
- [ ] Set up arbitrator as DAO governance with time-lock (48h)
- [ ] Configure protocol treasury as multi-sig
- [ ] Add 3-5 independent AI oracle signers
- [ ] Test all role transitions on testnet
- [ ] Document emergency procedures for all roles
- [ ] Set up monitoring for privileged function calls
- [ ] Create runbook for multi-sig operations

### During Operations

- [ ] Regular rotation of multi-sig signers (annually)
- [ ] Audit all privileged function calls (weekly)
- [ ] Monitor arbitrator decisions for fairness
- [ ] Track AI oracle signer performance
- [ ] Review fee configurations quarterly
- [ ] Test pause/unpause mechanisms monthly

---

## Role Transition Guide

### 1. Initial Deployment (Deployer Control)

```solidity
// All contracts owned by deployer
deployer = msg.sender;
```

**Duration**: 2-4 weeks  
**Purpose**: Bug fixes, configuration tuning

---

### 2. Interim Multi-sig (3-of-5)

```solidity
// Transfer ownership to interim multi-sig
marketFactory.transferOwnership(interimMultisig);
feeSplitter.transferOwnership(interimMultisig);
// ... repeat for all contracts
```

**Duration**: 3-6 months  
**Purpose**: Stability, community building

---

### 3. DAO Governance (Final State)

```solidity
// Transfer to DAO timelock
marketFactory.transferOwnership(daoTimelock);
resolutionModule.setArbitrator(daoGovernance);
```

**Duration**: Permanent  
**Purpose**: Full decentralization

---

## Emergency Scenarios

### Scenario 1: Critical Bug in MarketAMM

**Action**: Protocol Owner calls `pause()`
```solidity
marketAMM.pause(); // Disables trading
```

**Effect**: Trading stops, users can still claim winnings  
**Resolution**: Fix bug, deploy new AMM, unpause

---

### Scenario 2: Compromised AI Oracle Signer

**Action**: Protocol Owner removes signer
```solidity
aiOracleAdapter.removeSigner(compromisedSigner);
```

**Effect**: Future signatures from that signer rejected  
**Resolution**: Add new reputable signer

---

### Scenario 3: Disputed Arbitrator Decision

**Action**: Community governance vote to replace arbitrator
```solidity
resolutionModule.setArbitrator(newArbitrator);
```

**Effect**: Future disputes go to new arbitrator  
**Resolution**: Existing disputes unaffected

---

### Scenario 4: Protocol Treasury Compromise

**Action**: Protocol Owner updates treasury address
```solidity
feeSplitter.setProtocolTreasury(newTreasury);
```

**Effect**: Future fees go to new address  
**Resolution**: Historical unclaimed fees may be lost

---

## Contact Information

**Protocol Owner**: [multi-sig address or DAO contact]  
**Arbitrator**: [DAO governance portal]  
**AI Oracle Signers**: [List of validator contacts]  
**Protocol Treasury**: [Treasury multi-sig address]  

**For Security Issues**: See docs/SECURITY.md

---

**Last Updated**: 2025-10-28  
**Document Version**: 1.0  
**Protocol Version**: Phase 5 Complete
