<div align="center">

# Access Control & Roles Documentation

**Project Gamma - Roles, Permissions & Governance**

[![Multi-sig](https://img.shields.io/badge/Governance-Multi--sig-blue?style=flat-square)]()
[![DAO Ready](https://img.shields.io/badge/Future-DAO%20Governance-purple?style=flat-square)]()
[![Decentralized](https://img.shields.io/badge/Architecture-Decentralized-green?style=flat-square)]()

*Complete access control hierarchy and role-based permissions for prediction markets*

---

**Version:** Phase 5 Complete • **Last Updated:** 2025-10-28

---

[Role Hierarchy](#role-hierarchy) • [Permissions](#role-definitions) • [Security](#security-checklist) • [Transitions](#role-transition-guide) • [Emergency](#emergency-scenarios)

---

</div>

## Overview

This document defines all access control roles, permissions, and responsibilities in the BNB Chain Prediction Market protocol. The system is designed for progressive decentralization, starting with deployer control and transitioning to DAO governance.

### Role Hierarchy

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

<table>
<tr>
<td width="50%">

**Description**

Administrator of all core contracts with configuration authority

**Recommended Structure**
- 5-of-9 multi-sig wallet
- Or DAO timelock contract
- Progressive decentralization path

</td>
<td width="50%">

**Transition Plan**

1. **Weeks 1-4:** Deployer control (testing)
2. **Months 1-6:** Interim multi-sig (3-of-5)
3. **6+ Months:** DAO governance (final)

</td>
</tr>
</table>

#### Permissions Matrix

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td><code>setMinCreatorStake(uint256)</code></td>
<td>Adjust minimum stake for market creation</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>setProtocolTreasury(address)</code></td>
<td>Update protocol treasury address</td>
<td>🟡 Medium</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>updateFeeConfig(uint256, uint16, uint16)</code></td>
<td>Change fee split per market</td>
<td>🟡 Medium</td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td><code>setAMMAuthorization(address, bool)</code></td>
<td>Authorize/revoke AMM contracts</td>
<td>🔴 High</td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td><code>setResolutionAuthorization(address, bool)</code></td>
<td>Authorize/revoke resolver contracts</td>
<td>🔴 High</td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>addTier(uint256, uint16)</code></td>
<td>Add new fee tier</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>updateTier(uint256, uint256, uint16)</code></td>
<td>Update existing tier</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>removeLastTier()</code></td>
<td>Remove last tier</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td><code>addMinter(address)</code></td>
<td>Authorize new minter</td>
<td>🔴 High</td>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td><code>removeMinter(address)</code></td>
<td>Revoke minter</td>
<td>🟡 Medium</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>setMinBond(uint256)</code></td>
<td>Set minimum bond for proposals</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>setDisputeWindow(uint256)</code></td>
<td>Set dispute period duration</td>
<td>🟡 Medium</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>setArbitrator(address)</code></td>
<td>Set arbitrator address</td>
<td>⚠️ <strong>CRITICAL</strong></td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>pause() / unpause()</code></td>
<td>Emergency pause trading</td>
<td>⚠️ <strong>CRITICAL</strong></td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>close()</code></td>
<td>Close market early</td>
<td>🔴 High</td>
</tr>
</table>

#### Responsibilities

<table>
<tr>
<td width="50%">

**Day-to-Day Operations**
- Monitor protocol health and security
- Respond to critical bugs or exploits
- Coordinate upgrades and migrations
- Manage multi-sig operations

</td>
<td width="50%">

**Long-Term Goals**
- Transition to DAO governance
- Document all procedures
- Train community governance
- Establish transparent processes

</td>
</tr>
</table>

#### Limitations

**What Protocol Owner CANNOT Do:**
- ❌ Steal user funds
- ❌ Manipulate market outcomes
- ❌ Change resolved market results
- ❌ Mint HORIZON tokens directly (only authorized minters)
- ❌ Access user wallets or private keys
- ❌ Reverse finalized resolutions

---

### 2. Arbitrator

<table>
<tr>
<td width="50%">

**Description**

Final authority for disputed market resolutions

**Recommended Structure**
- DAO governance contract
- Minimum 48-hour timelock
- Community voting mechanism

</td>
<td width="50%">

**Accountability**

All decisions are:
- ✅ On-chain and permanent
- ✅ Transparent and auditable
- ✅ Subject to community review
- ✅ Economically incentivized (bond slashing)

</td>
</tr>
</table>

#### Permissions

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>finalizeDisputed(uint256, uint8, bool)</code></td>
<td>Resolve disputed markets and slash incorrect parties</td>
<td>⚠️ <strong>CRITICAL</strong></td>
</tr>
</table>

#### Decision Process

<table>
<tr>
<td width="20%" align="center">

**1. Review**

Examine market question and resolution criteria

</td>
<td width="20%" align="center">

**2. Evidence**

Study proposer's evidence (IPFS/Arweave)

</td>
<td width="20%" align="center">

**3. Counter-Evidence**

Review disputer's objections

</td>
<td width="20%" align="center">

**4. Determine**

Make fact-based decision

</td>
<td width="20%" align="center">

**5. Execute**

Call contract with outcome

</td>
</tr>
</table>

#### Slashing Rules

<details>
<summary><strong>Slash Proposer (Incorrect Proposal)</strong></summary>

```solidity
finalizeDisputed(marketId, correctOutcome, slashProposer: true);
```

**Result:**
- Proposer loses their bond
- Disputer receives proposer's bond + their own bond refunded
- Market resolves with correct outcome

</details>

<details>
<summary><strong>Slash Disputer (Incorrect Dispute)</strong></summary>

```solidity
finalizeDisputed(marketId, proposerOutcome, slashProposer: false);
```

**Result:**
- Disputer loses their bond
- Proposer receives disputer's bond + their own bond refunded
- Market resolves with proposer's outcome

</details>

#### Limitations

**What Arbitrator CANNOT Do:**
- ❌ Resolve non-disputed markets (automatic after 48h)
- ❌ Change finalized results
- ❌ Access user funds directly
- ❌ Create resolution proposals
- ❌ Override automatic finalization

---

### 3. Market Creators

<table>
<tr>
<td width="50%">

**Description**

Users who stake HORIZON tokens to create prediction markets

**Requirements**
- Hold minimum stake (default: 100 HORIZON)
- Stake locked until market resolution
- Refundable after resolution

</td>
<td width="50%">

**Incentive**

Receive 90-98% of trading fees based on traders' HORIZON holdings

**Revenue Example:**
- $1000 trade with 0 HORIZON trader: Earn $18.00
- $1000 trade with 500K+ HORIZON trader: Earn $19.60

</td>
</tr>
</table>

#### Permissions

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td><code>createMarket(...)</code></td>
<td>Create new prediction market</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td><code>refundCreatorStake(uint256)</code></td>
<td>Claim stake back after resolution</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>claimCreatorFees(uint256, address)</code></td>
<td>Claim earned trading fees</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>claimCreatorFeesMultiple(...)</code></td>
<td>Batch claim fees from multiple markets</td>
<td>🟢 None</td>
</tr>
</table>

#### Market Creation Parameters

```solidity
function createMarket(
    string memory question,           // "Will BTC reach $100k by Dec 31, 2024?"
    string memory category,           // "Crypto", "Sports", "Politics"
    string[] memory aiEvidenceURIs,   // IPFS/Arweave links for AI resolution
    uint256 closeTime,                // Unix timestamp when trading closes
    address collateralToken,          // USDC, DAI, or other ERC20
    uint256 minLiquidityAmount        // Minimum for initial LP
) external returns (uint256 marketId);
```

#### Fee Structure & Revenue

<table>
<tr>
<th>Trader's HORIZON Balance</th>
<th>User Pays</th>
<th>Creator Receives (per $1000 trade)</th>
<th>Protocol Receives (per $1000 trade)</th>
</tr>
<tr>
<td>0 (Tier 0)</td>
<td>2%</td>
<td>$18.00 (90%)</td>
<td>$2.00 (10%)</td>
</tr>
<tr>
<td>10,000 (Tier 1)</td>
<td>2%</td>
<td>$18.40 (92%)</td>
<td>$1.60 (8%)</td>
</tr>
<tr>
<td>50,000 (Tier 2)</td>
<td>2%</td>
<td>$18.80 (94%)</td>
<td>$1.20 (6%)</td>
</tr>
<tr>
<td>100,000 (Tier 3)</td>
<td>2%</td>
<td>$19.20 (96%)</td>
<td>$0.80 (4%)</td>
</tr>
<tr>
<td>500,000+ (Tier 4)</td>
<td>2%</td>
<td>$19.60 (98%)</td>
<td>$0.40 (2%)</td>
</tr>
</table>

**Why This Matters:**

This creates a powerful incentive for market creators to attract HORIZON token holders to their markets through marketing and community building!

#### Responsibilities

<table>
<tr>
<td width="50%">

**Market Quality**
- Write clear, unambiguous questions
- Set appropriate close times
- Provide evidence sources for AI
- Ensure objective resolvability

</td>
<td width="50%">

**Market Success**
- Add initial liquidity (recommended)
- Promote market to traders
- Attract HORIZON holders
- Monitor market activity

</td>
</tr>
</table>

#### Limitations

**What Market Creators CANNOT Do:**
- ❌ Cancel markets once created
- ❌ Change parameters after creation
- ❌ Influence resolution outcome
- ❌ Withdraw stake before resolution
- ❌ Claim fees before they're earned

---

### 4. AI Oracle Signers

<table>
<tr>
<td width="50%">

**Description**

Independent entities that sign resolution proposals using AI analysis

**Recommended**
- 3-5 independent AI agents
- Or trusted validators
- Geographic diversity
- Operational redundancy

</td>
<td width="50%">

**Selection Criteria**

- Appointed by Protocol Owner
- Should be diverse and reputable
- Regular performance reviews
- Subject to removal if compromised
- Public accountability

</td>
</tr>
</table>

#### Permissions

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>AIOracleAdapter</strong></td>
<td><code>proposeAI(...)</code></td>
<td>Submit AI-signed resolution proposal</td>
<td>🟡 Medium</td>
</tr>
</table>

#### Signature Requirements

```solidity
struct AIProposal {
    uint256 marketId;
    uint8 outcome;              // 0 = YES, 1 = NO
    string[] evidenceURIs;      // Public evidence (IPFS/Arweave)
    uint256 validFrom;          // Timestamp when signature becomes valid
    uint256 validUntil;         // Expiration timestamp (replay protection)
}
```

#### Signing Process

<table>
<tr>
<td width="20%" align="center">

**1. Analyze**

AI agent examines market question

</td>
<td width="20%" align="center">

**2. Gather**

Collect evidence from sources

</td>
<td width="20%" align="center">

**3. Determine**

Make fact-based outcome decision

</td>
<td width="20%" align="center">

**4. Publish**

Upload evidence to IPFS/Arweave

</td>
<td width="20%" align="center">

**5. Sign**

Create EIP-712 signature

</td>
</tr>
</table>

#### Managing Signers

<details>
<summary><strong>Add New Signer</strong></summary>

```solidity
// Only Protocol Owner can add signers
aiOracleAdapter.addSigner(newSignerAddress);
```

**Use Cases:**
- Onboarding new AI service providers
- Adding redundancy
- Geographic expansion

</details>

<details>
<summary><strong>Remove Signer</strong></summary>

```solidity
// Only Protocol Owner can remove signers
aiOracleAdapter.removeSigner(oldSignerAddress);
```

**Use Cases:**
- Signer compromised
- Poor performance
- Service discontinued

</details>

#### Responsibilities

- Provide accurate, unbiased resolutions
- Publish evidence publicly for transparency
- Use reliable data sources
- Respond to community concerns
- Maintain operational security of signing keys
- Report suspicious activities

#### Limitations

**What AI Signers CANNOT Do:**
- ❌ Propose without valid signature
- ❌ Override disputes (Arbitrator decides)
- ❌ Reuse signatures (replay protection)
- ❌ Access user funds
- ❌ Force immediate finalization

---

### 5. Protocol Treasury

<table>
<tr>
<td width="50%">

**Description**

Multi-sig wallet that receives protocol's share of trading fees

**Recommended Structure**
- 3-of-5 multi-sig wallet
- Or DAO treasury contract
- Transparent fund management

</td>
<td width="50%">

**Purpose**

Fund protocol growth and sustainability:
- Development & engineering
- Security audits & bounties
- Marketing & growth
- Liquidity incentives
- Operations & legal

</td>
</tr>
</table>

#### Permissions

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>claimProtocolFees(address)</code></td>
<td>Claim protocol fees for one token</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>claimProtocolFeesMultiple(address[])</code></td>
<td>Batch claim fees from multiple tokens</td>
<td>🟢 None</td>
</tr>
</table>

#### Fee Structure

<table>
<tr>
<th>Trader's HORIZON Balance</th>
<th>Protocol Share</th>
<th>Revenue per $1000 Trade</th>
</tr>
<tr>
<td>0 (Tier 0)</td>
<td>10%</td>
<td>$2.00</td>
</tr>
<tr>
<td>10,000 (Tier 1)</td>
<td>8%</td>
<td>$1.60</td>
</tr>
<tr>
<td>50,000 (Tier 2)</td>
<td>6%</td>
<td>$1.20</td>
</tr>
<tr>
<td>100,000 (Tier 3)</td>
<td>4%</td>
<td>$0.80</td>
</tr>
<tr>
<td>500,000+ (Tier 4)</td>
<td>2%</td>
<td>$0.40</td>
</tr>
</table>

**Economic Design:** The protocol intentionally subsidizes whale retention (lower protocol share for HORIZON holders) to grow the ecosystem and increase total volume.

#### Recommended Fund Allocation

<table>
<tr>
<td width="20%" align="center">

**40%**

Development & Engineering

</td>
<td width="20%" align="center">

**25%**

Security & Audits

</td>
<td width="20%" align="center">

**20%**

Marketing & Growth

</td>
<td width="20%" align="center">

**10%**

Liquidity Incentives

</td>
<td width="20%" align="center">

**5%**

Operations & Legal

</td>
</tr>
</table>

#### Limitations

**What Treasury CANNOT Do:**
- ❌ Claim creator fees
- ❌ Change fee split (only Protocol Owner can)
- ❌ Access market funds directly
- ❌ Manipulate market outcomes

---

### 6. Users (Traders & Liquidity Providers)

<table>
<tr>
<td width="50%">

**Description**

Anyone interacting with the protocol

**Requirements**
- Collateral tokens for trading
- HORIZON tokens for fee discounts (optional)
- No KYC or permissions needed

</td>
<td width="50%">

**Benefits**

- Trade on any market
- Provide liquidity and earn fees
- Propose resolutions
- Dispute incorrect outcomes
- Full self-custody

</td>
</tr>
</table>

#### Permissions

<table>
<tr>
<th>Contract</th>
<th>Function</th>
<th>Purpose</th>
<th>Risk Level</th>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>buyYes(uint256, uint256)</code></td>
<td>Buy YES outcome tokens</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>buyNo(uint256, uint256)</code></td>
<td>Buy NO outcome tokens</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>sellYes(uint256, uint256)</code></td>
<td>Sell YES outcome tokens</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>sellNo(uint256, uint256)</code></td>
<td>Sell NO outcome tokens</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>addLiquidity(uint256)</code></td>
<td>Provide liquidity, earn LP tokens and fees</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td><code>removeLiquidity(uint256)</code></td>
<td>Withdraw liquidity, burn LP tokens</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td><code>redeem(uint256, uint256)</code></td>
<td>Claim winnings after resolution (1:1)</td>
<td>🟢 None</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>proposeResolution(uint256, uint8)</code></td>
<td>Propose outcome with bond stake</td>
<td>🟢 Low</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>dispute(uint256, uint8)</code></td>
<td>Dispute incorrect proposal with bond</td>
<td>🟡 Medium</td>
</tr>
</table>

#### Fee Tiers (Based on HORIZON Holdings)

**Simple User Experience: All users pay 2% trading fee**

The HORIZON balance only determines how the fee is split between protocol and market creator:

<table>
<tr>
<th>Tier</th>
<th>HORIZON Balance</th>
<th>User Fee</th>
<th>Protocol Share</th>
<th>Creator Share</th>
</tr>
<tr>
<td>0</td>
<td>0</td>
<td>2.00%</td>
<td>10%</td>
<td>90%</td>
</tr>
<tr>
<td>1</td>
<td>10,000</td>
<td>2.00%</td>
<td>8%</td>
<td>92%</td>
</tr>
<tr>
<td>2</td>
<td>50,000</td>
<td>2.00%</td>
<td>6%</td>
<td>94%</td>
</tr>
<tr>
<td>3</td>
<td>100,000</td>
<td>2.00%</td>
<td>4%</td>
<td>96%</td>
</tr>
<tr>
<td>4</td>
<td>500,000+</td>
<td>2.00%</td>
<td>2%</td>
<td>98%</td>
</tr>
</table>

**Key Benefits:**
1. **Simple UX** - Users always pay 2%, easy to understand
2. **Creator Incentive** - Market creators earn MORE from HORIZON whale traders
3. **Token Utility** - Strong incentive for active traders to hold HORIZON
4. **Growth Model** - Protocol subsidizes whale retention for ecosystem growth

#### Responsibilities

- Set appropriate slippage tolerance (1-5% recommended)
- Verify market questions before trading
- Hold HORIZON tokens for better creator revenue (optional)
- Claim winnings after markets resolve
- Report suspicious activity or bugs

#### Limitations

**What Users CANNOT Do:**
- ❌ Trade after market closes
- ❌ Redeem losing outcome tokens
- ❌ Remove liquidity during resolution period
- ❌ Cancel trades (slippage protection only)
- ❌ Force market resolution

---

## Access Control Patterns

### Implementation Details

<table>
<tr>
<td width="33%" valign="top">

**1. Owner-Based**

Uses OpenZeppelin Ownable

```solidity
modifier onlyOwner() {
    if (msg.sender != owner()) {
        revert Unauthorized();
    }
    _;
}
```

**Transition:**
- `transferOwnership(newOwner)`
- Should use 2-step with `acceptOwnership()`

</td>
<td width="33%" valign="top">

**2. Role-Based**

Custom authorization checks

```solidity
// OutcomeToken
modifier onlyAuthorizedAMM() {
    if (!authorizedAMMs[msg.sender]) {
        revert Unauthorized();
    }
    _;
}

// FeeSplitter
modifier onlyFactory() {
    if (msg.sender != factory) {
        revert Unauthorized();
    }
    _;
}
```

</td>
<td width="33%" valign="top">

**3. Economic-Based**

Bond requirements for participation

```solidity
// Anyone can propose if they post bond
function proposeResolution(
    uint256 marketId,
    uint8 outcome
) external {
    bondToken.transferFrom(
        msg.sender,
        address(this),
        minBond
    );
    // Process proposal
}
```

</td>
</tr>
</table>

### Signer-Based (EIP-712)

Used in AIOracleAdapter for cryptographic verification:

```solidity
modifier validSigner(address signer) {
    if (!authorizedSigners[signer]) revert UnauthorizedSigner();
    _;
}
```

---

## Security Checklist

### Before Mainnet Deployment

<table>
<tr>
<td width="50%">

**Access Control Setup**
- [ ] Transfer all contract ownership to multi-sig (3-of-5 minimum)
- [ ] Set up arbitrator as DAO governance with timelock (48h)
- [ ] Configure protocol treasury as multi-sig
- [ ] Add 3-5 independent AI oracle signers
- [ ] Test all role transitions on testnet

</td>
<td width="50%">

**Documentation & Monitoring**
- [ ] Document emergency procedures for all roles
- [ ] Set up monitoring for privileged function calls
- [ ] Create runbook for multi-sig operations
- [ ] Establish communication channels
- [ ] Test pause/unpause mechanisms

</td>
</tr>
</table>

### During Operations

<table>
<tr>
<td width="33%">

**Regular Tasks**
- [ ] Rotation of multi-sig signers (annually)
- [ ] Audit privileged function calls (weekly)
- [ ] Monitor arbitrator decisions
- [ ] Track AI signer performance

</td>
<td width="33%">

**Quarterly Reviews**
- [ ] Review fee configurations
- [ ] Assess governance effectiveness
- [ ] Security audit updates
- [ ] Community feedback

</td>
<td width="33%">

**Emergency Drills**
- [ ] Test pause mechanisms (monthly)
- [ ] Practice multi-sig coordination
- [ ] Run incident response simulations
- [ ] Update contact lists

</td>
</tr>
</table>

---

## Role Transition Guide

### Progressive Decentralization Path

<table>
<tr>
<th>Phase</th>
<th>Duration</th>
<th>Control Structure</th>
<th>Purpose</th>
</tr>
<tr>
<td><strong>1. Initial Deployment</strong></td>
<td>2-4 weeks</td>
<td>Deployer control (single address)</td>
<td>Bug fixes, configuration tuning, rapid iteration</td>
</tr>
<tr>
<td><strong>2. Interim Multi-sig</strong></td>
<td>3-6 months</td>
<td>3-of-5 multi-sig</td>
<td>Stability, community building, operational testing</td>
</tr>
<tr>
<td><strong>3. DAO Governance</strong></td>
<td>Permanent</td>
<td>DAO with timelock (48h+)</td>
<td>Full decentralization, community control</td>
</tr>
</table>

### Transition Steps

<details>
<summary><strong>Phase 1 → Phase 2: Deployer to Multi-sig</strong></summary>

```solidity
// Transfer ownership of all contracts
marketFactory.transferOwnership(interimMultisig);
feeSplitter.transferOwnership(interimMultisig);
outcomeToken.transferOwnership(interimMultisig);
horizonPerks.transferOwnership(interimMultisig);
horizonToken.transferOwnership(interimMultisig);
resolutionModule.transferOwnership(interimMultisig);
aiOracleAdapter.transferOwnership(interimMultisig);

// Set arbitrator
resolutionModule.setArbitrator(interimMultisig);

// Set protocol treasury
feeSplitter.setProtocolTreasury(interimMultisig);
```

**Verify:** All `owner()` calls return multi-sig address

</details>

<details>
<summary><strong>Phase 2 → Phase 3: Multi-sig to DAO</strong></summary>

```solidity
// Deploy DAO contracts
address daoTimelock = deploy DAOTimelock(48 hours);
address daoGovernance = deploy DAOGovernance();

// Transfer ownership to DAO timelock
marketFactory.transferOwnership(daoTimelock);
feeSplitter.transferOwnership(daoTimelock);
// ... repeat for all contracts

// Set arbitrator to DAO governance
resolutionModule.setArbitrator(daoGovernance);

// Set treasury to DAO treasury
feeSplitter.setProtocolTreasury(daoTreasury);
```

**Verify:** Test governance proposal flow on testnet first

</details>

---

## Emergency Scenarios

### Emergency Response Playbook

<table>
<tr>
<th>Scenario</th>
<th>Action</th>
<th>Command</th>
<th>Effect</th>
</tr>
<tr>
<td><strong>1. Critical Bug in Trading</strong></td>
<td>Pause MarketAMM</td>
<td><code>marketAMM.pause()</code><br>(Only owner)</td>
<td>
• Disables buy/sell/liquidity operations<br>
• Users can still claim winnings<br>
• Resolution process continues normally
</td>
</tr>
<tr>
<td><strong>2. Invalid Market Question</strong></td>
<td>Close Market Early</td>
<td><code>marketAMM.close()</code><br>(Only owner)</td>
<td>
• Trading stops immediately<br>
• Must still resolve via ResolutionModule<br>
• LPs can withdraw after resolution
</td>
</tr>
<tr>
<td><strong>3. Disputed Resolution</strong></td>
<td>Arbitrator Decides</td>
<td><code>resolutionModule.finalizeDisputed()</code><br>(Only arbitrator)</td>
<td>
• Overrides dispute with final decision<br>
• Slashes incorrect party's bond<br>
• Market proceeds to payouts
</td>
</tr>
<tr>
<td><strong>4. Treasury Compromise</strong></td>
<td>Update Treasury Address</td>
<td><code>feeSplitter.setProtocolTreasury()</code><br>(Only owner)</td>
<td>
• Future protocol fees → new address<br>
• Historical unclaimed fees unaffected
</td>
</tr>
<tr>
<td><strong>5. Compromised AI Signer</strong></td>
<td>Remove Signer</td>
<td><code>aiOracleAdapter.removeSigner()</code><br>(Only owner)</td>
<td>
• Future signatures from signer rejected<br>
• Existing proposals unaffected<br>
• Add replacement signer
</td>
</tr>
<tr>
<td><strong>6. Exploitable AMM Contract</strong></td>
<td>Revoke Authorization</td>
<td><code>outcomeToken.setAMMAuthorization()</code><br>(Only owner)</td>
<td>
• Prevents AMM from mint/burn operations<br>
• Existing outcome tokens unaffected<br>
• Users must wait for resolution
</td>
</tr>
</table>

### Emergency Contact Chain

1. **Immediate (0-15 min):** Pause affected contracts, notify core team
2. **Within 1 hour:** Notify community via Discord/Twitter
3. **Within 6 hours:** Publish detailed post-mortem and analysis
4. **Within 24 hours:** Deploy fix or provide migration plan
5. **Within 48 hours:** Resume operations or implement compensation plan

---

## Contact Information

<div align="center">

<table>
<tr>
<td align="center" width="25%">

**Protocol Owner**

Multi-sig address or DAO contact

[Link TBD]

</td>
<td align="center" width="25%">

**Arbitrator**

DAO governance portal

[Link TBD]

</td>
<td align="center" width="25%">

**Protocol Treasury**

Treasury multi-sig address

[Link TBD]

</td>
<td align="center" width="25%">

**Security Issues**

See SECURITY.md

developers@horizonoracles.com

</td>
</tr>
</table>

</div>

---

<div align="center">

**Project Gamma - Access Control Documentation**

Progressive decentralization • Community governance • Transparent operations

[Main Documentation](../) • [Security](./SECURITY.md) • [Smart Contracts](../contracts/) • [Website](https://horizonoracles.com/)

*Last Updated: 2025-10-28 • Document Version: 1.0 • Protocol Version: Phase 5 Complete*

</div>
