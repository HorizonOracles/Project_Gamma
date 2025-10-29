# Horizon Deployment Guide

This document provides comprehensive instructions for deploying the Horizon prediction market protocol to various networks.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Pre-Deployment Checklist](#pre-deployment-checklist)
3. [Configuration](#configuration)
4. [Deployment Process](#deployment-process)
5. [Post-Deployment Verification](#post-deployment-verification)
6. [Network-Specific Instructions](#network-specific-instructions)
7. [Troubleshooting](#troubleshooting)
8. [Multi-Sig Setup](#multi-sig-setup)

---

## Prerequisites

### Required Tools

- **Foundry**: Install from [getfoundry.sh](https://getfoundry.sh)
  ```bash
  curl -L https://foundry.paradigm.xyz | bash
  foundryup
  ```

- **Git**: For version control
- **Node.js**: v18+ (optional, for additional tooling)

### Required Information

Before deployment, gather the following:

1. **Private Key**: Deployer wallet with sufficient native tokens for gas
2. **RPC Endpoint**: Reliable RPC URL for target network
3. **Block Explorer API Key**: For contract verification (optional but recommended)
4. **Admin Addresses**: 
   - Protocol Owner (can manage all contracts)
   - Protocol Treasury (receives protocol fees)
   - Arbitrator (can finalize disputed resolutions)
   - AI Signer (authorized for AI-powered resolutions)

---

## Pre-Deployment Checklist

### Security Review

- [ ] All contracts have been audited or reviewed
- [ ] Phase 6 security hardening completed
- [ ] Slither analysis run with no critical issues
- [ ] All tests passing (252/252)
- [ ] SECURITY.md and ROLES.md documentation reviewed

### Configuration Review

- [ ] Admin addresses confirmed and secured
- [ ] Multi-sig wallet addresses prepared (if applicable)
- [ ] Protocol parameters validated:
  - [ ] HORIZON initial supply appropriate for network
  - [ ] Min creator stake set appropriately
  - [ ] Min resolution bond set appropriately
  - [ ] Dispute window configured (48 hours default)

### Infrastructure

- [ ] RPC endpoint tested and reliable
- [ ] Deployer wallet funded with sufficient gas
- [ ] Block explorer API key obtained
- [ ] Backup RPC endpoints identified

### Testing

- [ ] Deployment script tested on local Anvil:
  ```bash
  forge script script/Deploy.s.sol:Deploy
  ```
- [ ] Simulation run successfully:
  ```bash
  forge script script/Deploy.s.sol:Deploy --sig "simulate()"
  ```

---

## Configuration

### 1. Environment Setup

Copy the example environment file:

```bash
cp .env.example .env
```

### 2. Configure RPC Endpoints

Edit `.env` and add your RPC URLs:

```bash
# Example for BSC Testnet
BSC_TESTNET_RPC_URL=https://data-seed-prebsc-1-s1.binance.org:8545/

# Example for Sepolia
SEPOLIA_RPC_URL=https://rpc.sepolia.org
```

### 3. Configure Deployer Private Key

**⚠️ SECURITY WARNING**: Never commit your private key to version control!

```bash
# Generate a new wallet (optional)
cast wallet new

# Add to .env
PRIVATE_KEY=0x...
```

### 4. Configure Admin Addresses

Set the admin addresses in `.env`:

```bash
# Protocol Owner (manages contracts)
PROTOCOL_OWNER=0x...

# Protocol Treasury (receives fees)
PROTOCOL_TREASURY=0x...

# Arbitrator (dispute resolution)
ARBITRATOR_ADDRESS=0x...

# AI Signer (AI-powered resolutions)
AI_SIGNER_ADDRESS=0x...
```

**Note**: If not set, all addresses default to the deployer address.

### 5. Configure Protocol Parameters

Adjust parameters in `.env` as needed:

```bash
# Token Configuration
HORIZON_INITIAL_SUPPLY=100000000000000000000000000  # 100M HORIZON

# Market Creation
MIN_CREATOR_STAKE=10000000000000000000000  # 10,000 HORIZON

# Resolution System
MIN_RESOLUTION_BOND=1000000000000000000000  # 1,000 HORIZON
DISPUTE_WINDOW=172800  # 48 hours
```

### 6. Configure Block Explorer API Keys

For contract verification:

```bash
BSCSCAN_API_KEY=your_api_key_here
ETHERSCAN_API_KEY=your_api_key_here
BASESCAN_API_KEY=your_api_key_here
```

---

## Deployment Process

### Local Testing (Anvil)

**Step 1**: Start local Anvil node

```bash
anvil
```

**Step 2**: Deploy to Anvil (in new terminal)

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url http://localhost:8545 \
  --broadcast
```

### Testnet Deployment

**Recommended for first deployment**: Always deploy to testnet first!

#### BSC Testnet

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url bsc_testnet \
  --broadcast \
  --verify \
  -vvvv
```

#### Sepolia

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url sepolia \
  --broadcast \
  --verify \
  -vvvv
```

#### Base Sepolia

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url base_sepolia \
  --broadcast \
  --verify \
  -vvvv
```

### Mainnet Deployment

**⚠️ CRITICAL**: Only deploy to mainnet after thorough testnet validation!

#### BSC Mainnet

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url bsc \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

#### Ethereum Mainnet

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url mainnet \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

#### Base Mainnet

```bash
forge script script/Deploy.s.sol:Deploy \
  --rpc-url base \
  --broadcast \
  --verify \
  --slow \
  -vvvv
```

### Deployment Flags Explained

- `--rpc-url`: Network to deploy to (defined in foundry.toml)
- `--broadcast`: Actually send transactions (omit for dry run)
- `--verify`: Automatically verify contracts on block explorer
- `--slow`: Add delays between transactions (recommended for mainnet)
- `-vvvv`: Maximum verbosity for debugging

---

## Post-Deployment Verification

After deployment, the script will output all contract addresses. **Save these immediately!**

### Automated Verification

The deployment script includes automatic verification checks:

1. **Contract Deployment**: All 7 contracts deployed
2. **Constructor Parameters**: All parameters set correctly
3. **Authorizations**: All roles and permissions configured
4. **Ownership**: Ownership transferred to protocol owner (if specified)

### Manual Verification Checklist

#### 1. Contract Addresses

- [ ] Copy all addresses from deployment output
- [ ] Save addresses to `.env` file
- [ ] Verify addresses on block explorer
- [ ] Confirm all contracts verified on block explorer

#### 2. Token Verification

**HorizonToken** (src/HorizonToken.sol)
```bash
# Check total supply
cast call $HORIZON_TOKEN_ADDRESS "totalSupply()" --rpc-url <network>

# Check max supply
cast call $HORIZON_TOKEN_ADDRESS "MAX_SUPPLY()" --rpc-url <network>

# Check owner
cast call $HORIZON_TOKEN_ADDRESS "owner()" --rpc-url <network>
```

### Multi-Sig Operations

All admin operations must now be initiated through the multi-sig:

- Setting protocol parameters
- Adding/removing minters and resolvers
- Transferring ownership
- Emergency actions

---

## Post-Deployment Checklist

After successful deployment and verification:

- [ ] All contract addresses saved to `.env` and secure location
- [ ] All contracts verified on block explorer
- [ ] All authorization checks passed
- [ ] All parameter checks passed
- [ ] Ownership transferred to protocol owner / multi-sig
- [ ] Admin addresses notified and access confirmed
- [ ] Deployment documented in changelog / release notes
- [ ] Monitoring and alerting configured
- [ ] Frontend/backend updated with new contract addresses
- [ ] Community announcement prepared (if applicable)

---

## Emergency Procedures

### Pause Operations

If issues are discovered post-deployment:

1. **Do not create markets**: Wait for fix
2. **Contact arbitrator**: Can intervene in disputes
3. **Monitor contracts**: Set up alerts for unusual activity

### Recovery Options

The protocol includes several safety mechanisms:

- **Ownership control**: Owner can update critical parameters
- **Arbitrator role**: Can resolve disputes fairly
- **Resolution disputes**: Community can dispute incorrect resolutions
- **Time locks**: Dispute windows provide time for intervention

---

## Additional Resources

- **Architecture Documentation**: See main README.md
- **Security Documentation**: See docs/SECURITY.md
- **Roles Documentation**: See docs/ROLES.md
- **Foundry Documentation**: https://book.getfoundry.sh

---

## Support

For deployment support:

1. Check [Troubleshooting](#troubleshooting) section
2. Review test suite for reference: `test/integration/FullSystem.t.sol`
3. Consult Foundry documentation: https://book.getfoundry.sh

---

**Last Updated**: Phase 7 - Deployment Scripts & Infrastructure  
**Version**: 1.0.0  
**Deployment Script**: script/Deploy.s.sol
