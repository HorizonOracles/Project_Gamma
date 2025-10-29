<div align="center">

# Deployment Guide

**Project Gamma - Smart Contract Deployment**

[![Foundry](https://img.shields.io/badge/Foundry-Deploy-black?style=flat-square)](https://book.getfoundry.sh/)
[![Networks](https://img.shields.io/badge/Networks-BSC%20%7C%20Ethereum%20%7C%20Base-blue?style=flat-square)]()
[![Verification](https://img.shields.io/badge/Auto-Verification-green?style=flat-square)]()

*Comprehensive deployment instructions for mainnet and testnet environments*

---

**Version:** 1.0.0 • **Script:** script/Deploy.s.sol • **Phase:** 7 Complete

---

[Prerequisites](#prerequisites) • [Configuration](#configuration) • [Deployment](#deployment-process) • [Verification](#post-deployment-verification) • [Troubleshooting](#troubleshooting)

---

</div>

## Overview

This guide provides step-by-step instructions for deploying the Horizon prediction market protocol to various networks. The deployment script handles all contract deployments, authorizations, and initial configuration in a single transaction sequence.

### Deployment Features

<table>
<tr>
<td width="25%" align="center">

**Automated**

Single command deploys all 7 contracts with proper configuration

</td>
<td width="25%" align="center">

**Verified**

Automatic contract verification on block explorers

</td>
<td width="25%" align="center">

**Secure**

Built-in verification checks and safety validations

</td>
<td width="25%" align="center">

**Multi-Network**

Supports BSC, Ethereum, Base, and more

</td>
</tr>
</table>

---

## Prerequisites

### Required Tools

<table>
<tr>
<td width="50%">

**Development Tools**

**Foundry** - Ethereum development toolkit
```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
```

**Git** - Version control
```bash
# Already installed on most systems
git --version
```

</td>
<td width="50%">

**Optional Tools**

**Node.js** - v18+ for additional tooling
```bash
node --version
npm --version
```

**Cast** - Command-line tool (included with Foundry)
```bash
cast --version
```

</td>
</tr>
</table>

### Required Information

Before deployment, gather the following information:

<table>
<tr>
<th>Category</th>
<th>Requirements</th>
<th>Notes</th>
</tr>
<tr>
<td><strong>Wallet</strong></td>
<td>
• Deployer private key<br>
• Sufficient native tokens for gas
</td>
<td>
Keep private key secure<br>
Never commit to Git
</td>
</tr>
<tr>
<td><strong>Network</strong></td>
<td>
• RPC endpoint URL<br>
• Block explorer API key
</td>
<td>
Use reliable RPC providers<br>
API key enables auto-verification
</td>
</tr>
<tr>
<td><strong>Admin Addresses</strong></td>
<td>
• Protocol Owner<br>
• Protocol Treasury<br>
• Arbitrator<br>
• AI Signer
</td>
<td>
Can use multi-sig addresses<br>
Defaults to deployer if not set
</td>
</tr>
</table>

---

## Pre-Deployment Checklist

### Security Review

<table>
<tr>
<td width="50%">

**Code Quality**
- [ ] All contracts audited or reviewed
- [ ] Phase 6 security hardening complete
- [ ] Slither analysis clean (no critical issues)
- [ ] All 252 tests passing
- [ ] Gas optimization reviewed

</td>
<td width="50%">

**Documentation**
- [ ] SECURITY.md reviewed
- [ ] ROLES.md reviewed
- [ ] Deployment parameters validated
- [ ] Admin addresses confirmed
- [ ] Emergency procedures documented

</td>
</tr>
</table>

### Configuration Review

**Protocol Parameters:**

- [ ] HORIZON initial supply appropriate for network
- [ ] Minimum creator stake set correctly
- [ ] Minimum resolution bond configured
- [ ] Dispute window set (48 hours default)
- [ ] Fee tiers reviewed

**Admin Addresses:**

- [ ] Protocol owner address secured
- [ ] Multi-sig wallet addresses prepared
- [ ] Treasury address confirmed
- [ ] Arbitrator address assigned
- [ ] AI signer addresses ready

### Infrastructure

<table>
<tr>
<td width="33%">

**Network Access**
- [ ] RPC endpoint tested
- [ ] Backup RPC identified
- [ ] Network stable
- [ ] Gas prices reasonable

</td>
<td width="33%">

**Wallet Setup**
- [ ] Deployer wallet funded
- [ ] Gas buffer included
- [ ] Private key secured
- [ ] Backup wallet ready

</td>
<td width="33%">

**Verification**
- [ ] Block explorer API key
- [ ] Verification tested
- [ ] Network supported
- [ ] API rate limits checked

</td>
</tr>
</table>

### Testing

**Local Testing:**

```bash
# Test deployment on local Anvil
forge script script/Deploy.s.sol:Deploy

# Run simulation
forge script script/Deploy.s.sol:Deploy --sig "simulate()"
```

**Expected Results:**
- All contracts deploy successfully
- All authorizations set correctly
- No reverts or failures
- Gas estimates reasonable

---

## Configuration

### 1. Environment Setup

**Create Environment File**

```bash
# Copy example configuration
cp .env.example .env

# Edit with your preferred editor
vim .env
# or
nano .env
```

### 2. RPC Endpoints

Configure network RPC URLs in `.env`:

<details>
<summary><strong>BNB Chain Networks</strong></summary>

```bash
# BSC Mainnet
BSC_RPC_URL=https://bsc-dataseed.binance.org/

# BSC Testnet
BSC_TESTNET_RPC_URL=https://data-seed-prebsc-1-s1.binance.org:8545/
```

</details>

<details>
<summary><strong>Ethereum Networks</strong></summary>

```bash
# Ethereum Mainnet
MAINNET_RPC_URL=https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY

# Sepolia Testnet
SEPOLIA_RPC_URL=https://rpc.sepolia.org
```

</details>

<details>
<summary><strong>Base Networks</strong></summary>

```bash
# Base Mainnet
BASE_RPC_URL=https://mainnet.base.org

# Base Sepolia Testnet
BASE_SEPOLIA_RPC_URL=https://sepolia.base.org
```

</details>

### 3. Deployer Private Key

**⚠️ SECURITY CRITICAL**: Never commit your private key!

```bash
# Option 1: Generate new wallet
cast wallet new

# Option 2: Use existing wallet
# Add to .env (never commit!)
PRIVATE_KEY=0x...
```

**Security Best Practices:**
- Use a dedicated deployment wallet
- Fund with only necessary gas + small buffer
- Transfer ownership immediately after deployment
- Consider hardware wallet for mainnet

### 4. Admin Addresses

Configure administrative addresses in `.env`:

```bash
# Protocol Owner (manages all contracts)
PROTOCOL_OWNER=0x...

# Protocol Treasury (receives protocol fees)
PROTOCOL_TREASURY=0x...

# Arbitrator (dispute resolution)
ARBITRATOR_ADDRESS=0x...

# AI Signer (AI-powered resolutions)
AI_SIGNER_ADDRESS=0x...
```

**Note:** If not specified, all addresses default to deployer address.

### 5. Protocol Parameters

Adjust protocol parameters as needed:

<details>
<summary><strong>Token Configuration</strong></summary>

```bash
# HORIZON Token
HORIZON_INITIAL_SUPPLY=100000000000000000000000000  # 100M HORIZON
HORIZON_NAME="Horizon Token"
HORIZON_SYMBOL="HORIZON"
```

</details>

<details>
<summary><strong>Market Creation</strong></summary>

```bash
# Minimum stake to create markets
MIN_CREATOR_STAKE=10000000000000000000000  # 10,000 HORIZON
```

</details>

<details>
<summary><strong>Resolution System</strong></summary>

```bash
# Minimum bond for resolution proposals
MIN_RESOLUTION_BOND=1000000000000000000000  # 1,000 HORIZON

# Dispute window (seconds)
DISPUTE_WINDOW=172800  # 48 hours
```

</details>

### 6. Block Explorer API Keys

Configure for automatic contract verification:

```bash
# BNB Chain
BSCSCAN_API_KEY=your_api_key_here

# Ethereum
ETHERSCAN_API_KEY=your_api_key_here

# Base
BASESCAN_API_KEY=your_api_key_here
```

**Get API Keys:**
- BscScan: https://bscscan.com/myapikey
- Etherscan: https://etherscan.io/myapikey
- Basescan: https://basescan.org/myapikey

---

## Deployment Process

### Local Testing (Anvil)

**Step 1: Start Local Node**

```bash
# Terminal 1: Start Anvil
anvil
```

**Step 2: Deploy to Local Network**

```bash
# Terminal 2: Deploy contracts
forge script script/Deploy.s.sol:Deploy \
  --rpc-url http://localhost:8545 \
  --broadcast
```

**Expected Output:**
- 7 contracts deployed
- All addresses logged
- Verification checks passed

### Testnet Deployment

**Always deploy to testnet first!**

<table>
<tr>
<th>Network</th>
<th>Command</th>
<th>Notes</th>
</tr>
<tr>
<td><strong>BSC Testnet</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url bsc_testnet \
  --broadcast \
  --verify \
  -vvvv
```

</td>
<td>
Recommended first testnet<br>
Fast finality<br>
Free testnet BNB
</td>
</tr>
<tr>
<td><strong>Sepolia</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url sepolia \
  --broadcast \
  --verify \
  -vvvv
```

</td>
<td>
Ethereum testnet<br>
Slower finality<br>
Use faucets for ETH
</td>
</tr>
<tr>
<td><strong>Base Sepolia</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url base_sepolia \
  --broadcast \
  --verify \
  -vvvv
```

</td>
<td>
Base testnet<br>
L2 gas savings<br>
Fast and cheap
</td>
</tr>
</table>

### Mainnet Deployment

**⚠️ CRITICAL: Only deploy after thorough testnet validation!**

<table>
<tr>
<th>Network</th>
<th>Command</th>
<th>Estimated Gas Cost</th>
</tr>
<tr>
<td><strong>BSC Mainnet</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url bsc \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

</td>
<td>~0.1 BNB<br>(varies with gas price)</td>
</tr>
<tr>
<td><strong>Ethereum Mainnet</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url mainnet \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

</td>
<td>~0.5-1 ETH<br>(highly variable)</td>
</tr>
<tr>
<td><strong>Base Mainnet</strong></td>
<td>

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url base \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

</td>
<td>~0.01 ETH<br>(L2 savings)</td>
</tr>
</table>

### Deployment Flags Reference

<table>
<tr>
<th>Flag</th>
<th>Purpose</th>
<th>When to Use</th>
</tr>
<tr>
<td><code>--rpc-url</code></td>
<td>Network to deploy to</td>
<td>Always required (defined in foundry.toml)</td>
</tr>
<tr>
<td><code>--broadcast</code></td>
<td>Actually send transactions</td>
<td>Omit for dry run / simulation</td>
</tr>
<tr>
<td><code>--verify</code></td>
<td>Auto-verify on block explorer</td>
<td>Recommended for all deployments</td>
</tr>
<tr>
<td><code>--slow</code></td>
<td>Add delays between transactions</td>
<td>Recommended for mainnet (prevents nonce issues)</td>
</tr>
<tr>
<td><code>-vvvv</code></td>
<td>Maximum verbosity</td>
<td>For debugging or first-time deployments</td>
</tr>
</table>

---

## Post-Deployment Verification

### Automated Verification

The deployment script includes automatic checks:

<table>
<tr>
<td width="25%" align="center">

**Contract Deployment**

All 7 contracts deployed successfully

</td>
<td width="25%" align="center">

**Constructor Params**

All parameters set correctly

</td>
<td width="25%" align="center">

**Authorizations**

All roles and permissions configured

</td>
<td width="25%" align="center">

**Ownership**

Ownership transferred to protocol owner

</td>
</tr>
</table>

### Save Contract Addresses

**CRITICAL: Save these addresses immediately!**

The deployment script outputs addresses in this format:

```
=== Deployment Summary ===
HorizonToken: 0x...
OutcomeToken: 0x...
FeeSplitter: 0x...
HorizonPerks: 0x...
MarketFactory: 0x...
ResolutionModule: 0x...
AIOracleAdapter: 0x...
```

**Save To:**
1. `.env` file (gitignored)
2. Secure password manager
3. Team documentation
4. Deployment log file

### Manual Verification Checklist

#### 1. Contract Verification

<details>
<summary><strong>Verify on Block Explorer</strong></summary>

- [ ] All contracts show verified checkmark
- [ ] Source code matches deployment
- [ ] Constructor arguments visible
- [ ] Compiler version correct (0.8.23)

**Manual Verification (if needed):**

```bash
forge verify-contract \
  --chain-id <chain_id> \
  --compiler-version 0.8.23 \
  <contract_address> \
  <contract_path>:<contract_name>
```

</details>

#### 2. Token Verification

<details>
<summary><strong>HorizonToken Checks</strong></summary>

```bash
# Check total supply (should be initial supply)
cast call $HORIZON_TOKEN_ADDRESS "totalSupply()" --rpc-url <network>

# Check max supply (should be 1 billion)
cast call $HORIZON_TOKEN_ADDRESS "MAX_SUPPLY()" --rpc-url <network>

# Check owner (should be protocol owner)
cast call $HORIZON_TOKEN_ADDRESS "owner()" --rpc-url <network>

# Check deployer is minter
cast call $HORIZON_TOKEN_ADDRESS "minters(address)" $DEPLOYER_ADDRESS --rpc-url <network>
```

**Expected Values:**
- Total Supply: 100,000,000 * 10^18
- Max Supply: 1,000,000,000 * 10^18
- Owner: Protocol Owner address
- Deployer is minter: true

</details>

#### 3. Authorization Checks

<details>
<summary><strong>OutcomeToken Authorizations</strong></summary>

```bash
# Check MarketAMM is authorized
cast call $OUTCOME_TOKEN_ADDRESS "authorizedAMMs(address)" $MARKET_AMM_ADDRESS --rpc-url <network>

# Check ResolutionModule is authorized
cast call $OUTCOME_TOKEN_ADDRESS "authorizedResolvers(address)" $RESOLUTION_MODULE_ADDRESS --rpc-url <network>
```

**Expected:** Both should return `true`

</details>

<details>
<summary><strong>FeeSplitter Configuration</strong></summary>

```bash
# Check factory address
cast call $FEE_SPLITTER_ADDRESS "factory()" --rpc-url <network>

# Check protocol treasury
cast call $FEE_SPLITTER_ADDRESS "protocolTreasury()" --rpc-url <network>
```

**Expected:**
- Factory: MarketFactory address
- Treasury: Protocol Treasury address

</details>

#### 4. Parameter Verification

<details>
<summary><strong>MarketFactory Parameters</strong></summary>

```bash
# Check minimum creator stake
cast call $MARKET_FACTORY_ADDRESS "minCreatorStake()" --rpc-url <network>
```

**Expected:** 10,000 * 10^18 (or configured value)

</details>

<details>
<summary><strong>ResolutionModule Parameters</strong></summary>

```bash
# Check minimum bond
cast call $RESOLUTION_MODULE_ADDRESS "minBond()" --rpc-url <network>

# Check dispute window
cast call $RESOLUTION_MODULE_ADDRESS "disputeWindow()" --rpc-url <network>

# Check arbitrator
cast call $RESOLUTION_MODULE_ADDRESS "arbitrator()" --rpc-url <network>
```

**Expected:**
- Min Bond: 1,000 * 10^18 (or configured value)
- Dispute Window: 172,800 seconds (48 hours)
- Arbitrator: Arbitrator address

</details>

---

## Multi-Sig Setup

### Transferring to Multi-Sig

After successful deployment and initial testing, transfer ownership to a multi-sig wallet:

<details>
<summary><strong>Transfer Ownership Commands</strong></summary>

```bash
# HorizonToken
cast send $HORIZON_TOKEN_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# OutcomeToken
cast send $OUTCOME_TOKEN_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# FeeSplitter
cast send $FEE_SPLITTER_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# HorizonPerks
cast send $HORIZON_PERKS_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# MarketFactory
cast send $MARKET_FACTORY_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# ResolutionModule
cast send $RESOLUTION_MODULE_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY

# AIOracleAdapter
cast send $AI_ORACLE_ADAPTER_ADDRESS \
  "transferOwnership(address)" $MULTISIG_ADDRESS \
  --rpc-url <network> \
  --private-key $PRIVATE_KEY
```

</details>

### Multi-Sig Operations

**Future Admin Operations:**

All privileged functions must now be called through the multi-sig:
- Setting protocol parameters
- Adding/removing minters and resolvers
- Transferring ownership
- Emergency pause/unpause

**Recommended Multi-Sig Providers:**
- Gnosis Safe (most popular)
- Safe{Wallet} (rebrand of Gnosis Safe)
- Multi-sig wallet of your choice

---

## Troubleshooting

### Common Issues

<table>
<tr>
<th>Issue</th>
<th>Possible Cause</th>
<th>Solution</th>
</tr>
<tr>
<td><strong>Insufficient Gas</strong></td>
<td>Deployer wallet has insufficient funds</td>
<td>
Fund wallet with more native tokens<br>
Add 20% buffer for gas spikes
</td>
</tr>
<tr>
<td><strong>RPC Connection Failed</strong></td>
<td>RPC endpoint down or rate limited</td>
<td>
Switch to backup RPC<br>
Check RPC provider status<br>
Add API key if required
</td>
</tr>
<tr>
<td><strong>Verification Failed</strong></td>
<td>Block explorer API issues or rate limit</td>
<td>
Wait and retry verification<br>
Verify manually if needed<br>
Check API key validity
</td>
</tr>
<tr>
<td><strong>Nonce Too Low</strong></td>
<td>Pending transaction conflict</td>
<td>
Wait for pending tx to confirm<br>
Use <code>--slow</code> flag<br>
Check mempool status
</td>
</tr>
<tr>
<td><strong>Simulation Reverted</strong></td>
<td>Configuration error or parameter issue</td>
<td>
Review .env configuration<br>
Check parameter values<br>
Run local tests first
</td>
</tr>
</table>

### Debug Commands

**Check Wallet Balance:**

```bash
cast balance $DEPLOYER_ADDRESS --rpc-url <network>
```

**Check Pending Transactions:**

```bash
cast nonce $DEPLOYER_ADDRESS --rpc-url <network>
```

**Estimate Gas:**

```bash
forge script script/Deploy.s.sol:Deploy --rpc-url <network> --estimate-gas
```

**Get Gas Price:**

```bash
cast gas-price --rpc-url <network>
```

---

## Post-Deployment Checklist

### Immediate Actions

<table>
<tr>
<td width="50%">

**Documentation**
- [ ] Save all contract addresses
- [ ] Update `.env` file
- [ ] Document deployment in changelog
- [ ] Create release notes
- [ ] Update frontend/backend configs

</td>
<td width="50%">

**Verification**
- [ ] All contracts verified on explorer
- [ ] Authorization checks passed
- [ ] Parameter checks passed
- [ ] Ownership transferred (if applicable)
- [ ] Multi-sig setup complete

</td>
</tr>
</table>

### Operations Setup

<table>
<tr>
<td width="33%">

**Monitoring**
- [ ] Block explorer bookmarks
- [ ] Alert system configured
- [ ] Dashboard setup
- [ ] Analytics tracking
- [ ] Error monitoring

</td>
<td width="33%">

**Communication**
- [ ] Team notified
- [ ] Admin access confirmed
- [ ] Support team trained
- [ ] Community announcement
- [ ] Social media updates

</td>
<td width="33%">

**Security**
- [ ] Private keys secured
- [ ] Backup keys stored
- [ ] Emergency contacts set
- [ ] Incident response ready
- [ ] Audit scheduled

</td>
</tr>
</table>

---

## Emergency Procedures

### Post-Deployment Issues

**If you discover critical issues:**

<table>
<tr>
<td width="50%">

**Immediate Actions**
1. Do NOT create any markets yet
2. Notify all admin addresses
3. Document the issue thoroughly
4. Assess severity and impact
5. Determine if redeploy is needed

</td>
<td width="50%">

**Communication**
1. Alert core team immediately
2. Notify community if public
3. Prepare incident report
4. Document timeline
5. Plan resolution strategy

</td>
</tr>
</table>

### Recovery Options

The protocol includes safety mechanisms:

- **Ownership Control**: Owner can update critical parameters
- **Arbitrator Role**: Can intervene in disputes
- **Dispute System**: Community can challenge incorrect resolutions
- **Time Locks**: Windows provide intervention opportunities
- **Pause Functions**: Emergency pause for AMM operations

---

## Additional Resources

<div align="center">

<table>
<tr>
<td align="center" width="25%">

**Architecture**

[Main README](../README.md)

System design and overview

</td>
<td align="center" width="25%">

**Security**

[SECURITY.md](./SECURITY.md)

Security measures and audits

</td>
<td align="center" width="25%">

**Roles**

[ROLES.md](./ROLES.md)

Access control and permissions

</td>
<td align="center" width="25%">

**Foundry**

[Foundry Book](https://book.getfoundry.sh)

Development toolkit docs

</td>
</tr>
</table>

</div>

---

## Support

<div align="center">

**Need Deployment Help?**

<table>
<tr>
<td align="center" width="33%">

**Documentation**

Check troubleshooting section and test suite reference

</td>
<td align="center" width="33%">

**Community**

[![Discord](https://img.shields.io/badge/Discord-Join%20Us-5865F2?style=flat-square&logo=discord&logoColor=white)](https://discord.com/invite/TuUHwwKjHh)

</td>
<td align="center" width="33%">

**Developer Support**

developers@horizonoracles.com

</td>
</tr>
</table>

</div>

---

<div align="center">

**Project Gamma - Deployment Guide**

Secure • Automated • Verified

[Main Documentation](../) • [Smart Contracts](../contracts/) • [Website](https://horizonoracles.com/)

*Last Updated: Phase 7 Complete • Version: 1.0.0 • Deployment Script: script/Deploy.s.sol*

</div>
