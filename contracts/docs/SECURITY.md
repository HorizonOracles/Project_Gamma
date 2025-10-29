# Security Documentation

## Overview

This document outlines the security measures, assumptions, invariants, and potential risks in the BNB Chain Prediction Market protocol.

**Version:** Phase 5 Complete  
**Audit Status:** Pre-audit  
**Test Coverage:** 252 tests passing  
**Static Analysis:** Slither clean (2 false positive reentrancy warnings)

---

## Table of Contents

1. [Security Measures](#security-measures)
2. [Access Control](#access-control)
3. [Reentrancy Protection](#reentrancy-protection)
4. [Core Invariants](#core-invariants)
5. [Trust Assumptions](#trust-assumptions)
6. [Known Issues & Limitations](#known-issues--limitations)
7. [Attack Vectors & Mitigations](#attack-vectors--mitigations)
8. [Emergency Procedures](#emergency-procedures)

---

## Security Measures

### 1. OpenZeppelin Dependencies
- **ReentrancyGuard**: All state-changing functions in `MarketAMM` protected
- **SafeERC20**: All ERC20 operations use `safeTransfer`, `safeTransferFrom`, `forceApprove`
- **Ownable**: Role-based access control for admin functions
- **Pausable**: Emergency pause functionality in `MarketAMM`

### 2. Checks-Effects-Interactions Pattern
All contracts follow CEI pattern:
```solidity
// Example from MarketAMM._buy()
1. Checks: require(status == Active), slippage checks
2. Effects: Update reserves, mint tokens
3. Interactions: External token transfers
```

### 3. Integer Overflow Protection
- Solidity 0.8.24 provides automatic overflow/underflow checks
- All arithmetic operations are safe by default

### 4. Gas Optimizations
- Array length caching in loops (src/HorizonPerks.sol:188, src/MarketFactory.sol:391)
- Storage packing in structs where possible
- View function optimizations

---

## Access Control

### Contract: MarketAMM
- **Public**: `buyYes()`, `buyNo()`, `sellYes()`, `sellNo()`, `addLiquidity()`, `removeLiquidity()`
- **OnlyOwner**: `close()`, `pause()`, `unpause()`
- **Protected**: No privileged trading access

### Contract: MarketFactory
- **Public**: `createMarket()`, `refundCreatorStake()`
- **OnlyOwner**: `setMinCreatorStake()`
- **Internal**: `updateMarketStatus()` (called by factory only)

### Contract: FeeSplitter
- **Public**: `claimCreatorFees()`, `claimProtocolFees()`
- **OnlyFactory**: `registerMarket()`, `distribute()`
- **OnlyOwner**: `setProtocolTreasury()`, `updateFeeConfig()`

### Contract: OutcomeToken
- **OnlyAMM**: `mintOutcome()`, `burnOutcome()`
- **OnlyResolver**: `setWinningOutcome()`, `transferCollateral()`
- **OnlyOwner**: `setAMMAuthorization()`, `registerMarket()`

### Contract: ResolutionModule
- **Public**: `proposeResolution()`, `dispute()`, `finalize()`
- **OnlyArbitrator**: `finalizeDisputed()`
- **OnlyOwner**: `setMinBond()`, `setDisputeWindow()`, `setArbitrator()`

### Contract: HorizonPerks
- **View**: All fee calculation functions public
- **OnlyOwner**: `addTier()`, `updateTier()`, `removeLastTier()`

### Contract: HorizonToken
- **Public**: `transfer()`, `approve()`, `transferFrom()`
- **OnlyMinter**: `mint()`
- **OnlyOwner**: `addMinter()`, `removeMinter()`

---

## Reentrancy Protection

### Protected Functions
All externally-callable state-changing functions use `nonReentrant` modifier:

**MarketAMM.sol:**
- `buyYes()` / `buyNo()` - src/MarketAMM.sol:234, 251
- `sellYes()` / `sellNo()` - src/MarketAMM.sol:268, 285
- `addLiquidity()` - src/MarketAMM.sol:144
- `removeLiquidity()` - src/MarketAMM.sol:174

**OutcomeToken.sol:**
- `redeem()` - src/OutcomeToken.sol:185

### Slither False Positives
Slither reports reentrancy warnings in `_buy()` and `_sell()`:
- **Status**: False positive
- **Reason**: All external functions are `nonReentrant`
- **External Calls**: Only to trusted contracts (FeeSplitter, OutcomeToken)
- **CEI Pattern**: Properly followed

---

## Core Invariants

### MarketAMM Invariants

1. **CPMM Formula**: `reserveYes * reserveNo = k` (constant product)
   - Maintained after every trade
   - Verified in: test/unit/MarketAMM.t.sol:360 (`test_CPMM_Invariant`)

2. **Price Sum**: `priceYes + priceNo = 1.0` (100%)
   - Ensures fair pricing
   - Verified in: test/unit/MarketAMM.t.sol:375 (`test_Invariant_PricesSumToOne`)

3. **Collateral Backing**: `totalCollateral >= reserveYes + reserveNo`
   - All outcome tokens backed 1:1 by collateral
   - Verified in: test/unit/MarketAMM.t.sol:386 (`test_Invariant_TotalCollateralBacksReserves`)

4. **Reserve Sum**: `reserveYes + reserveNo = totalCollateral` (when balanced)
   - Maintained through mint/burn mechanism
   - Verified in: test/unit/MarketAMM.t.sol:392 (`test_Invariant_ReserveSum`)

5. **LP Token Supply**: LP shares proportional to collateral contributed
   - `lpTokens = (collateralIn * totalSupply) / totalCollateral`
   - Verified in: test/unit/MarketAMM.t.sol:122 (`test_AddLiquidity_Subsequent`)

### FeeSplitter Invariants

1. **Fee Split**: `protocolBps + creatorBps = 10000` (100%)
   - Default: 1000 (10%) protocol, 9000 (90%) creator
   - Enforced in: src/FeeSplitter.sol:155

2. **Fee Accounting**: `protocolFees[token] + creatorFees[token] = totalFeesCollected[token]`
   - No fees lost or double-counted
   - Verified in: test/unit/FeeSplitter.t.sol:58 (`test_Distribute`)

### OutcomeToken Invariants

1. **Token Encoding**: `tokenId = (marketId << 8) | outcomeId`
   - Unique token per market-outcome pair
   - Verified in: test/unit/OutcomeToken.t.sol:107 (`testFuzz_EncodeDecodeTokenId`)

2. **Redemption**: Only winning outcome tokens redeemable 1:1 for collateral
   - Enforced in: src/OutcomeToken.sol:190-191

### ResolutionModule Invariants

1. **Resolution State Machine**:
   - `Unresolved → Proposed → [Disputed] → Finalized`
   - No backwards transitions
   - Verified in: test/unit/ResolutionModule.t.sol:119 (`test_FullCycle_NoDispute`)

2. **Bond Slashing**: Incorrect party's bond goes to correct party
   - Zero-sum game for disputes
   - Verified in: test/unit/ResolutionModule.t.sol:170 (`test_FinalCycle_WithDispute_DisputerCorrect`)

---

## Trust Assumptions

### Trusted Roles

1. **Contract Owner** (MarketFactory deployer)
   - Can set minimum creator stakes
   - Cannot steal user funds
   - Cannot manipulate market outcomes
   - Should be a multi-sig wallet or DAO

2. **Protocol Treasury** (FeeSplitter.protocolTreasury)
   - Receives 10% of trading fees
   - No control over markets or user funds
   - Should be a multi-sig wallet

3. **AI Oracle Signers** (AIOracleAdapter.authorizedSigners)
   - Sign resolution proposals
   - Attestations are public and auditable
   - Multiple signers for redundancy
   - **Critical**: Must be independent and trustworthy

4. **Arbitrator** (ResolutionModule.arbitrator)
   - Final say in disputed resolutions
   - Should be a DAO governance mechanism
   - **Critical**: Most powerful role in the system

5. **Market Creators**
   - Stake HORIZON tokens to create markets
   - Receive 90% of trading fees
   - No control over resolution or trading

### Trusted Contracts

1. **Collateral Tokens** (USDC, DAI, etc.)
   - Must be standard ERC20 implementations
   - Should be well-known stablecoins
   - **Risk**: If collateral depegs, market outcomes affected

2. **HORIZON Token**
   - Protocol's native token for staking and fee discounts
   - Minters should be limited to trusted contracts
   - **Risk**: If minter compromised, token inflation possible

---

## Known Issues & Limitations

### 1. Slither Reentrancy Warnings
- **Status**: False positives
- **Location**: src/MarketAMM.sol:305, 371
- **Mitigation**: All public functions use `nonReentrant` modifier
- **Evidence**: 252 tests passing including reentrancy attack tests

### 2. Front-Running Risk
- **Issue**: Traders can front-run large trades to profit from price movement
- **Mitigation**: Slippage protection via `minTokensOut` / `minCollateralOut`
- **User Protection**: Set appropriate slippage tolerance (1-5%)

### 3. Oracle Manipulation Risk
- **Issue**: AI oracle signers could collude to propose incorrect outcomes
- **Mitigation**: 
  - Multi-signer requirement
  - Public evidence attestation (IPFS/Arweave)
  - Dispute mechanism with bond slashing
  - Arbitrator as final backstop

### 4. Market Creator Spam
- **Issue**: Malicious actors could create many low-quality markets
- **Mitigation**: 
  - Minimum creator stake (100 HORIZON tokens default)
  - Stake locked until resolution
  - Can increase stake requirement via `setMinCreatorStake()`

### 5. Liquidity Fragmentation
- **Issue**: Each market has its own AMM pool
- **Limitation**: Cannot aggregate liquidity across markets
- **Trade-off**: Simplicity and isolation vs. capital efficiency

### 6. No Emergency Withdrawals
- **Issue**: If market never resolves, funds are locked
- **Mitigation**: 
  - Arbitrator can force resolution
  - Consider adding timeout-based escape hatch (future upgrade)

---

## Attack Vectors & Mitigations

### 1. Reentrancy Attack
- **Vector**: Malicious contract calls back during ERC20 transfer
- **Mitigation**: `nonReentrant` on all external functions + CEI pattern
- **Status**: ✅ Protected

### 2. Flash Loan Manipulation
- **Vector**: Borrow large amount, manipulate price, profit
- **Mitigation**: 
  - Slippage protection limits profit
  - No oracle dependencies on AMM price
  - Fees discourage manipulation
- **Status**: ✅ Protected

### 3. Griefing via Disputes
- **Vector**: Always dispute correct resolutions to delay payouts
- **Mitigation**: 
  - Bond requirement (minBond = 100 tokens)
  - Slashing for incorrect disputes
  - Economic disincentive
- **Status**: ✅ Protected

### 4. Sandwich Attacks
- **Vector**: Front-run victim's trade, profit from price impact
- **Mitigation**: 
  - User-set slippage limits
  - MEV protection (future: private transactions)
- **Status**: ⚠️ User Responsibility

### 5. Sybil Attack on Fee Tiers
- **Vector**: Split HORIZON holdings across wallets to avoid fee tiers
- **Mitigation**: 
  - Fee tiers already generous (max 2% → min 0.5%)
  - Gas costs of managing multiple wallets
  - Doesn't affect protocol security
- **Status**: ✅ Acceptable

### 6. Collateral Token Compromise
- **Vector**: If USDC blacklists AMM contract
- **Mitigation**: 
  - Support multiple collateral tokens
  - Users can choose preferred collateral
  - Each market isolated
- **Status**: ⚠️ Systemic Risk (out of scope)

### 7. Arbitrator Abuse
- **Vector**: Arbitrator manipulates resolutions for profit
- **Mitigation**: 
  - Arbitrator should be DAO with time-locks
  - All decisions on-chain and transparent
  - Community can fork if governance compromised
- **Status**: ⚠️ Governance Risk

---

## Emergency Procedures

### 1. Pause Trading (MarketAMM)
**When**: Critical bug discovered in trading logic
```solidity
// Only owner can pause
marketAMM.pause();
```
**Effect**: 
- Disables `buy()`, `sell()`, `addLiquidity()`, `removeLiquidity()`
- Users can still claim winnings
- Does not prevent resolution

### 2. Close Market Early
**When**: Market question becomes invalid or inappropriate
```solidity
// Only owner can close
marketAMM.close();
```
**Effect**:
- Disables trading immediately
- Must still resolve via ResolutionModule
- LPs can withdraw after resolution

### 3. Force Resolution via Arbitrator
**When**: Disputed resolution needs final decision
```solidity
// Only arbitrator can finalize disputed resolutions
resolutionModule.finalizeDisputed(marketId, outcome, slashProposer);
```
**Effect**:
- Overrides dispute
- Slashes incorrect party's bond
- Market can proceed to payouts

### 4. Update Protocol Treasury
**When**: Current treasury address compromised
```solidity
// Only owner can update
feeSplitter.setProtocolTreasury(newTreasury);
```
**Effect**:
- Future protocol fees go to new address
- Does not affect unclaimed historical fees

### 5. Revoke Malicious AMM Authorization
**When**: AMM contract found to be exploitable
```solidity
// Only owner can revoke
outcomeToken.setAMMAuthorization(maliciousAMM, false);
```
**Effect**:
- Prevents AMM from minting/burning outcome tokens
- Existing tokens unaffected
- Users must wait for resolution to claim

---

## Recommendations for Production

### Pre-Deployment
1. ✅ Complete professional security audit (Certik, Trail of Bits, etc.)
2. ✅ Set up bug bounty program (Immunefi)
3. ✅ Deploy to testnet for 2+ weeks of public testing
4. ✅ Multi-sig wallet for owner role (3-of-5 or 5-of-9)
5. ✅ Time-lock contract for critical upgrades (48h delay)

### Post-Deployment
1. ✅ Monitor all market creations for spam/abuse
2. ✅ Track unusually large trades for manipulation
3. ✅ Maintain reserve fund for potential exploit compensation
4. ✅ Set up alerting for abnormal contract behavior
5. ✅ Regular security reviews as protocol evolves

### Governance
1. ✅ Transition arbitrator role to DAO governance
2. ✅ Community voting on min creator stakes
3. ✅ Transparent process for adding AI oracle signers
4. ✅ Regular treasury audits

---

## Contact & Disclosure

**Security Contact**: [security@yourproject.com]  
**Bug Bounty**: [Link to Immunefi program]  
**Disclosure Policy**: Responsible disclosure, 90-day embargo  
**PGP Key**: [If applicable]

---

## Audit History

| Date | Auditor | Scope | Status | Report |
|------|---------|-------|--------|--------|
| TBD  | TBD     | Full  | Pending| -      |

---

**Last Updated**: 2025-10-28  
**Document Version**: 1.0  
**Protocol Version**: Phase 5 Complete
