<div align="center">

# Security Documentation

**Project Gamma - Smart Contract Security Overview**

[![Tests](https://img.shields.io/badge/Tests-252%20Passing-success?style=flat-square)]()
[![Coverage](https://img.shields.io/badge/Coverage-%3E95%25-green?style=flat-square)]()
[![Audit](https://img.shields.io/badge/Audit-Pre--Audit-orange?style=flat-square)]()
[![Slither](https://img.shields.io/badge/Slither-Clean-success?style=flat-square)]()

*Comprehensive security analysis and risk assessment for on-chain prediction markets*

---

**Version:** Phase 5 Complete • **Protocol Version:** 1.0 • **Last Updated:** 2025-10-28

---

[Security Measures](#security-measures) • [Access Control](#access-control) • [Invariants](#core-invariants) • [Attack Vectors](#attack-vectors--mitigations) • [Emergency Procedures](#emergency-procedures)

---

</div>

## Overview

This document outlines the security measures, assumptions, invariants, and potential risks in the BNB Chain Prediction Market protocol. All smart contracts have undergone internal security review with 252 passing tests and clean static analysis results.

### Security Status Summary

<table>
<tr>
<td width="25%" align="center">

**Test Coverage**

252 tests passing
>95% line coverage

</td>
<td width="25%" align="center">

**Static Analysis**

Slither clean
2 false positives

</td>
<td width="25%" align="center">

**Audit Status**

Internal review complete
External audit scheduled

</td>
<td width="25%" align="center">

**Known Issues**

0 critical
0 high severity

</td>
</tr>
</table>

---

## Security Measures

### Battle-Tested Dependencies

<table>
<tr>
<td width="50%">

**OpenZeppelin Libraries**

- **ReentrancyGuard** - All state-changing functions protected
- **SafeERC20** - Safe token operations throughout
- **Ownable** - Role-based access control
- **Pausable** - Emergency pause functionality

</td>
<td width="50%">

**Solidity 0.8.24 Safety**

- Automatic overflow/underflow protection
- Built-in arithmetic safety
- Type-safe operations
- Modern compiler optimizations

</td>
</tr>
</table>

### Design Patterns

**1. Checks-Effects-Interactions (CEI)**

All contracts follow the CEI pattern to prevent reentrancy:

```solidity
// Example from MarketAMM._buy()
1. Checks:    require(status == Active), slippage validation
2. Effects:   Update reserves, mint outcome tokens
3. Interactions: External token transfers to trusted contracts
```

**2. Pull Payment Pattern**

Fee distribution uses pull-over-push:
- Winners claim their own payouts
- No automatic transfers to potentially malicious contracts
- Prevents denial-of-service attacks

**3. Access Control Hierarchy**

Clear separation of privileges:
- Public functions for user operations
- OnlyOwner for admin configuration
- Role-specific permissions (AMM, Resolver, Factory)

---

## Access Control

### Permission Matrix

<table>
<tr>
<th>Contract</th>
<th>Public Functions</th>
<th>Privileged Functions</th>
<th>Owner Controls</th>
</tr>
<tr>
<td><strong>MarketAMM</strong></td>
<td>
• buyYes / buyNo<br>
• sellYes / sellNo<br>
• addLiquidity<br>
• removeLiquidity
</td>
<td>None - No privileged trading access</td>
<td>
• close()<br>
• pause()<br>
• unpause()
</td>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td>
• createMarket()<br>
• refundCreatorStake()
</td>
<td>
• updateMarketStatus() (internal)
</td>
<td>
• setMinCreatorStake()
</td>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td>
• claimCreatorFees()<br>
• claimProtocolFees()
</td>
<td>
• registerMarket() (OnlyFactory)<br>
• distribute() (OnlyFactory)
</td>
<td>
• setProtocolTreasury()<br>
• updateFeeConfig()
</td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td>Standard ERC-1155</td>
<td>
• mintOutcome() (OnlyAMM)<br>
• burnOutcome() (OnlyAMM)<br>
• setWinningOutcome() (OnlyResolver)
</td>
<td>
• setAMMAuthorization()<br>
• registerMarket()
</td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td>
• proposeResolution()<br>
• dispute()<br>
• finalize()
</td>
<td>
• finalizeDisputed() (OnlyArbitrator)
</td>
<td>
• setMinBond()<br>
• setDisputeWindow()<br>
• setArbitrator()
</td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td>All fee calculations (view)</td>
<td>None</td>
<td>
• addTier()<br>
• updateTier()<br>
• removeLastTier()
</td>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td>Standard ERC-20</td>
<td>
• mint() (OnlyMinter)
</td>
<td>
• addMinter()<br>
• removeMinter()
</td>
</tr>
</table>

### Critical Roles

<table>
<tr>
<td width="33%">

**Contract Owner**
- Admin configuration
- Cannot steal funds
- Cannot manipulate outcomes
- Should be multi-sig

</td>
<td width="33%">

**Arbitrator**
- Final dispute resolution
- Most powerful role
- Should be DAO governance
- Transparent on-chain

</td>
<td width="33%">

**AI Oracle Signers**
- Sign resolution proposals
- Multiple signers required
- Public attestations
- Regular rotation

</td>
</tr>
</table>

---

## Reentrancy Protection

### Protected Functions

All state-changing external functions use the `nonReentrant` modifier from OpenZeppelin's ReentrancyGuard:

<details>
<summary><strong>MarketAMM.sol</strong></summary>

- `buyYes()` - Line 234
- `buyNo()` - Line 251
- `sellYes()` - Line 268
- `sellNo()` - Line 285
- `addLiquidity()` - Line 144
- `removeLiquidity()` - Line 174

</details>

<details>
<summary><strong>OutcomeToken.sol</strong></summary>

- `redeem()` - Line 185

</details>

### Static Analysis Results

**Slither Reentrancy Warnings**

| Location | Type | Status | Reason |
|----------|------|--------|--------|
| src/MarketAMM.sol:305 | Reentrancy | ✅ False Positive | All callers use `nonReentrant` |
| src/MarketAMM.sol:371 | Reentrancy | ✅ False Positive | CEI pattern + trusted contracts only |

**Why These Are Safe:**
- All external entry points protected with `nonReentrant`
- External calls only to trusted contracts (FeeSplitter, OutcomeToken)
- CEI pattern strictly followed
- 252 tests passing including reentrancy attack scenarios

---

## Core Invariants

### Market AMM Invariants

<table>
<tr>
<th>Invariant</th>
<th>Description</th>
<th>Test Verification</th>
</tr>
<tr>
<td><strong>CPMM Formula</strong></td>
<td><code>reserveYes * reserveNo = k</code><br>Constant product maintained after every trade</td>
<td>test/unit/MarketAMM.t.sol:360<br><code>test_CPMM_Invariant</code></td>
</tr>
<tr>
<td><strong>Price Sum</strong></td>
<td><code>priceYes + priceNo = 1.0</code><br>Prices always sum to 100%</td>
<td>test/unit/MarketAMM.t.sol:375<br><code>test_Invariant_PricesSumToOne</code></td>
</tr>
<tr>
<td><strong>Collateral Backing</strong></td>
<td><code>totalCollateral ≥ reserveYes + reserveNo</code><br>All tokens backed 1:1 by collateral</td>
<td>test/unit/MarketAMM.t.sol:386<br><code>test_Invariant_TotalCollateralBacksReserves</code></td>
</tr>
<tr>
<td><strong>Reserve Sum</strong></td>
<td><code>reserveYes + reserveNo = totalCollateral</code><br>When market is balanced</td>
<td>test/unit/MarketAMM.t.sol:392<br><code>test_Invariant_ReserveSum</code></td>
</tr>
<tr>
<td><strong>LP Token Supply</strong></td>
<td><code>lpTokens ∝ collateralContributed</code><br>Proportional to contribution</td>
<td>test/unit/MarketAMM.t.sol:122<br><code>test_AddLiquidity_Subsequent</code></td>
</tr>
</table>

### Fee Splitter Invariants

<table>
<tr>
<th>Invariant</th>
<th>Formula</th>
<th>Enforcement</th>
</tr>
<tr>
<td><strong>Fee Split</strong></td>
<td><code>protocolBps + creatorBps = 10000</code> (100%)</td>
<td>src/FeeSplitter.sol:155</td>
</tr>
<tr>
<td><strong>Fee Accounting</strong></td>
<td><code>protocolFees + creatorFees = totalFeesCollected</code></td>
<td>test/unit/FeeSplitter.t.sol:58</td>
</tr>
</table>

Default split: 1000 bps (10%) protocol / 9000 bps (90%) creator

### Outcome Token Invariants

<table>
<tr>
<th>Invariant</th>
<th>Description</th>
<th>Verification</th>
</tr>
<tr>
<td><strong>Token Encoding</strong></td>
<td><code>tokenId = (marketId << 8) | outcomeId</code><br>Unique token per market-outcome</td>
<td>test/unit/OutcomeToken.t.sol:107<br><code>testFuzz_EncodeDecodeTokenId</code></td>
</tr>
<tr>
<td><strong>Redemption</strong></td>
<td>Only winning outcome tokens redeemable 1:1</td>
<td>src/OutcomeToken.sol:190-191</td>
</tr>
</table>

### Resolution Module Invariants

<table>
<tr>
<th>Invariant</th>
<th>State Machine</th>
<th>Test Coverage</th>
</tr>
<tr>
<td><strong>Resolution Flow</strong></td>
<td><code>Unresolved → Proposed → [Disputed] → Finalized</code><br>No backwards transitions allowed</td>
<td>test/unit/ResolutionModule.t.sol:119<br><code>test_FullCycle_NoDispute</code></td>
</tr>
<tr>
<td><strong>Bond Slashing</strong></td>
<td>Incorrect party's bond goes to correct party<br>Zero-sum game for disputes</td>
<td>test/unit/ResolutionModule.t.sol:170<br><code>test_FinalCycle_WithDispute_DisputerCorrect</code></td>
</tr>
</table>

---

## Trust Assumptions

### Trusted Roles & Responsibilities

<table>
<tr>
<td width="33%" valign="top">

**Contract Owner**

*Power Level: Medium*

**Can:**
- Set minimum creator stakes
- Configure protocol parameters
- Pause trading in emergencies

**Cannot:**
- Steal user funds
- Manipulate market outcomes
- Override resolutions

**Recommendation:** Multi-sig wallet (3-of-5 or 5-of-9)

</td>
<td width="33%" valign="top">

**AI Oracle Signers**

*Power Level: Medium-High*

**Can:**
- Sign resolution proposals
- Provide evidence attestations

**Cannot:**
- Force outcomes (disputeable)
- Access user funds

**Critical Requirements:**
- Multiple independent signers
- Public evidence on IPFS/Arweave
- Regular signer rotation

**Risk:** Collusion to propose incorrect outcomes

</td>
<td width="33%" valign="top">

**Arbitrator**

*Power Level: High*

**Can:**
- Final say on disputed resolutions
- Override dispute outcomes
- Slash bonds

**Cannot:**
- Create resolutions directly
- Steal non-disputed funds

**Recommendation:** DAO governance with timelock

**Risk:** Most powerful role; could manipulate outcomes if compromised

</td>
</tr>
</table>

### Trusted Contracts

<table>
<tr>
<td width="50%">

**Collateral Tokens (USDC, DAI, etc.)**

**Requirements:**
- Standard ERC20 implementation
- Well-known stablecoins
- No blacklist of AMM contracts

**Risk:** If collateral depegs or blacklists AMM, market affected

</td>
<td width="50%">

**HORIZON Token**

**Requirements:**
- Limited minter access
- Trusted minting contracts only

**Risk:** If minter compromised, token inflation possible

</td>
</tr>
</table>

---

## Known Issues & Limitations

### 1. Slither Reentrancy Warnings

<table>
<tr>
<td width="70%">

**Status:** ✅ False Positives

**Location:** src/MarketAMM.sol:305, 371

**Analysis:**
- All public functions use `nonReentrant` modifier
- External calls only to trusted contracts
- CEI pattern properly followed
- 252 tests passing including attack scenarios

</td>
<td width="30%">

**Severity:** None

**Action Required:** None

**Documentation:** This section

</td>
</tr>
</table>

### 2. Front-Running Risk

<table>
<tr>
<td width="70%">

**Issue:** Traders can front-run large trades to profit from price movement

**Mitigation:**
- Slippage protection via `minTokensOut` / `minCollateralOut`
- User-configurable tolerance (recommended 1-5%)
- Mempool monitoring services available

</td>
<td width="30%">

**Severity:** ⚠️ Low

**Action:** User responsibility

**Status:** Mitigated

</td>
</tr>
</table>

### 3. Oracle Manipulation Risk

<table>
<tr>
<td width="70%">

**Issue:** AI oracle signers could collude to propose incorrect outcomes

**Mitigation:**
- Multi-signer requirement (minimum 2-of-3)
- Public evidence attestation (IPFS/Arweave)
- 48-hour dispute window
- Bond-based dispute mechanism
- Arbitrator as final backstop
- Economic disincentive (bond slashing)

</td>
<td width="30%">

**Severity:** ⚠️ Medium

**Action:** Governance oversight

**Status:** Mitigated

</td>
</tr>
</table>

### 4. Market Creator Spam

<table>
<tr>
<td width="70%">

**Issue:** Malicious actors could create many low-quality markets

**Mitigation:**
- Minimum creator stake (100 HORIZON tokens default)
- Stake locked until resolution
- Configurable via `setMinCreatorStake()`
- UI filtering for quality markets

</td>
<td width="30%">

**Severity:** ⚠️ Low

**Action:** Parameter tuning

**Status:** Mitigated

</td>
</tr>
</table>

### 5. Liquidity Fragmentation

<table>
<tr>
<td width="70%">

**Issue:** Each market has its own isolated AMM pool

**Limitation:** Cannot aggregate liquidity across markets

**Trade-off:** Simplicity and isolation vs. capital efficiency

**Note:** Design decision, not a bug

</td>
<td width="30%">

**Severity:** Info

**Action:** Future optimization

**Status:** Accepted

</td>
</tr>
</table>

### 6. No Emergency Withdrawals

<table>
<tr>
<td width="70%">

**Issue:** If market never resolves, funds are locked indefinitely

**Mitigation:**
- Arbitrator can force resolution
- Consider timeout-based escape hatch (future upgrade)
- Market creator stake incentivizes proper resolution

</td>
<td width="30%">

**Severity:** ⚠️ Low

**Action:** Future feature

**Status:** Under consideration

</td>
</tr>
</table>

---

## Attack Vectors & Mitigations

### Attack Surface Analysis

<table>
<tr>
<th>Attack Vector</th>
<th>Description</th>
<th>Mitigation</th>
<th>Status</th>
</tr>
<tr>
<td><strong>1. Reentrancy</strong></td>
<td>Malicious contract calls back during ERC20 transfer</td>
<td>
• <code>nonReentrant</code> on all external functions<br>
• CEI pattern followed<br>
• Only trusted contract interactions
</td>
<td>✅ Protected</td>
</tr>
<tr>
<td><strong>2. Flash Loan Manipulation</strong></td>
<td>Borrow large amount, manipulate price, repay and profit</td>
<td>
• Slippage protection limits profit potential<br>
• No oracle price dependencies<br>
• Trading fees discourage manipulation<br>
• Economic disincentive
</td>
<td>✅ Protected</td>
</tr>
<tr>
<td><strong>3. Griefing via Disputes</strong></td>
<td>Always dispute correct resolutions to delay payouts</td>
<td>
• Bond requirement (minBond = 100 tokens)<br>
• Bond slashing for incorrect disputes<br>
• Economic disincentive
</td>
<td>✅ Protected</td>
</tr>
<tr>
<td><strong>4. Sandwich Attacks</strong></td>
<td>Front-run victim's trade, profit from price impact</td>
<td>
• User-set slippage limits<br>
• MEV protection (future: private txs)<br>
• Flashbots integration possible
</td>
<td>⚠️ User Responsibility</td>
</tr>
<tr>
<td><strong>5. Sybil Attack on Fee Tiers</strong></td>
<td>Split HORIZON holdings across wallets to avoid higher tiers</td>
<td>
• Fee tiers already generous (2% → 0.5%)<br>
• Gas costs of managing multiple wallets<br>
• Doesn't affect protocol security
</td>
<td>✅ Acceptable</td>
</tr>
<tr>
<td><strong>6. Collateral Token Compromise</strong></td>
<td>USDC blacklists AMM contract</td>
<td>
• Support multiple collateral tokens<br>
• Users choose preferred collateral<br>
• Each market isolated
</td>
<td>⚠️ Systemic Risk</td>
</tr>
<tr>
<td><strong>7. Arbitrator Abuse</strong></td>
<td>Arbitrator manipulates resolutions for profit</td>
<td>
• Arbitrator should be DAO with timelocks<br>
• All decisions on-chain and transparent<br>
• Community can fork if governance compromised
</td>
<td>⚠️ Governance Risk</td>
</tr>
</table>

---

## Emergency Procedures

### Emergency Response Matrix

<table>
<tr>
<th>Emergency Type</th>
<th>Action</th>
<th>Command</th>
<th>Effect</th>
</tr>
<tr>
<td><strong>1. Critical Bug in Trading</strong></td>
<td>Pause MarketAMM</td>
<td><code>marketAMM.pause()</code><br>(Only owner)</td>
<td>
• Disables buy/sell/liquidity ops<br>
• Claiming winnings still works<br>
• Resolution not affected
</td>
</tr>
<tr>
<td><strong>2. Invalid Market Question</strong></td>
<td>Close Market Early</td>
<td><code>marketAMM.close()</code><br>(Only owner)</td>
<td>
• Disables trading immediately<br>
• Must still resolve via ResolutionModule<br>
• LPs withdraw after resolution
</td>
</tr>
<tr>
<td><strong>3. Disputed Resolution</strong></td>
<td>Force Resolution</td>
<td><code>resolutionModule.finalizeDisputed()</code><br>(Only arbitrator)</td>
<td>
• Overrides dispute<br>
• Slashes incorrect party's bond<br>
• Market proceeds to payouts
</td>
</tr>
<tr>
<td><strong>4. Treasury Compromise</strong></td>
<td>Update Protocol Treasury</td>
<td><code>feeSplitter.setProtocolTreasury()</code><br>(Only owner)</td>
<td>
• Future fees go to new address<br>
• Historical unclaimed fees unaffected
</td>
</tr>
<tr>
<td><strong>5. Exploitable AMM</strong></td>
<td>Revoke AMM Authorization</td>
<td><code>outcomeToken.setAMMAuthorization()</code><br>(Only owner)</td>
<td>
• Prevents AMM from mint/burn<br>
• Existing tokens unaffected<br>
• Users wait for resolution
</td>
</tr>
</table>

### Emergency Contact Chain

1. **Immediate:** Pause affected contracts
2. **Within 1 hour:** Notify community via Discord/Twitter
3. **Within 6 hours:** Publish post-mortem analysis
4. **Within 24 hours:** Deploy fix or migration plan
5. **Within 48 hours:** Resume operations or provide compensation plan

---

## Recommendations for Production

### Pre-Deployment Checklist

<table>
<tr>
<td width="50%">

**Security Hardening**
- [ ] Complete professional security audit
- [ ] Set up bug bounty program (Immunefi)
- [ ] Deploy to testnet for 2+ weeks
- [ ] Public beta with limited funds
- [ ] Multi-sig wallet for owner role (3-of-5 minimum)
- [ ] Timelock contract for critical upgrades (48h)

</td>
<td width="50%">

**Monitoring & Response**
- [ ] Set up contract monitoring (Tenderly/Defender)
- [ ] Alert system for unusual activity
- [ ] Reserve fund for potential exploits
- [ ] Incident response playbook
- [ ] Emergency pause procedures documented
- [ ] Community communication channels

</td>
</tr>
</table>

### Post-Deployment Operations

<table>
<tr>
<td width="33%">

**Monitoring**
- Monitor market creations
- Track large trades
- Alert on abnormal patterns
- Regular security reviews
- Quarterly audits

</td>
<td width="33%">

**Governance**
- Transition to DAO
- Community voting
- Transparent processes
- Regular treasury audits
- Signer rotation policy

</td>
<td width="33%">

**Communication**
- Public incident reports
- Regular security updates
- Bug bounty highlights
- Audit publication
- Community engagement

</td>
</tr>
</table>

---

## Audit History

<table>
<tr>
<th>Date</th>
<th>Auditor</th>
<th>Scope</th>
<th>Findings</th>
<th>Status</th>
<th>Report</th>
</tr>
<tr>
<td>2025-10-28</td>
<td>Internal Team</td>
<td>Full Protocol</td>
<td>0 Critical, 0 High<br>2 False Positives</td>
<td>✅ Complete</td>
<td>This Document</td>
</tr>
<tr>
<td>TBD</td>
<td>External Auditor</td>
<td>Full Protocol</td>
<td>-</td>
<td>⏳ Scheduled</td>
<td>-</td>
</tr>
<tr>
<td>TBD</td>
<td>Bug Bounty (Immunefi)</td>
<td>Full Protocol</td>
<td>-</td>
<td>📋 Planned</td>
<td>-</td>
</tr>
</table>

---

## Contact & Disclosure

<div align="center">

### Security Contact Information

<table>
<tr>
<td align="center" width="33%">

**Critical Issues**

GitHub Security Advisory
or
developers@horizonoracles.com

</td>
<td align="center" width="33%">

**Bug Bounty**

Program launching post-audit
Details: TBD

</td>
<td align="center" width="33%">

**Disclosure Policy**

Responsible disclosure
90-day embargo
Coordinated release

</td>
</tr>
</table>

</div>

---

<div align="center">

**Project Gamma Security Documentation**

Built with security-first principles on BNB Chain

[Main Documentation](../) • [Smart Contracts](../contracts/) • [Website](https://horizonoracles.com/)

*Last Updated: 2025-10-28 • Document Version: 1.0 • Protocol Version: Phase 5 Complete*

</div>
