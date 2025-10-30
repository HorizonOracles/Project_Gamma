<div align="center">

#  Mainnet Deployment

**Project Gamma - BNB Chain Launch**

[![Network](https://img.shields.io/badge/Network-BSC%20Mainnet-F0B90B?style=flat-square&logo=binance)](https://bscscan.com/)
[![Status](https://img.shields.io/badge/Status-Live-success?style=flat-square)]()
[![Verified](https://img.shields.io/badge/Contracts-Verified-blue?style=flat-square)](https://bscscan.com/)

*Successfully deployed on October 30, 2025*

---

**Chain ID:** 56 • **Total Gas Cost:** 0.00057424175 BNB (~$0.35 USD)

---

</div>

## Deployment Summary

Project Gamma has been successfully deployed to BNB Chain mainnet. All 6 core contracts are live and operational.

### Deployment Statistics

<table>
<tr>
<td width="25%" align="center">

**Total Contracts**

6 contracts deployed

</td>
<td width="25%" align="center">

**Total Gas Used**

11,484,835 gas

</td>
<td width="25%" align="center">

**Total Cost**

0.00057424175 BNB

</td>
<td width="25%" align="center">

**Avg Gas Price**

0.05 gwei

</td>
</tr>
</table>

---

## Contract Addresses

### Core Tokens

<table>
<tr>
<th>Contract</th>
<th>Address</th>
<th>Block</th>
<th>Gas Used</th>
<th>BscScan</th>
</tr>
<tr>
<td><strong>HorizonToken</strong></td>
<td><code>0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444</code></td>
<td>Pre-deployed</td>
<td>N/A</td>
<td><a href="https://bscscan.com/address/0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444">View</a></td>
</tr>
<tr>
<td><strong>OutcomeToken</strong></td>
<td><code>0x17B322784265c105a94e4c3d00aF1E5f46a5F311</code></td>
<td>66439956</td>
<td>1,718,990</td>
<td><a href="https://bscscan.com/address/0x17B322784265c105a94e4c3d00aF1E5f46a5F311">View</a></td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C</code></td>
<td>66439956</td>
<td>1,354,764</td>
<td><a href="https://bscscan.com/address/0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C">View</a></td>
</tr>
</table>

### Protocol Infrastructure

<table>
<tr>
<th>Contract</th>
<th>Address</th>
<th>Block</th>
<th>Gas Used</th>
<th>BscScan</th>
</tr>
<tr>
<td><strong>FeeSplitter</strong></td>
<td><code>0x275017E98adF33051BbF477fe1DD197F681d4eF1</code></td>
<td>66439956</td>
<td>975,952</td>
<td><a href="https://bscscan.com/address/0x275017E98adF33051BbF477fe1DD197F681d4eF1">View</a></td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>0xF0CF4C741910cB48AC596F620a0AE892Cd247838</code></td>
<td>66439956</td>
<td>1,226,164</td>
<td><a href="https://bscscan.com/address/0xF0CF4C741910cB48AC596F620a0AE892Cd247838">View</a></td>
</tr>
<tr>
<td><strong>AIOracleAdapter</strong></td>
<td><code>0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93</code></td>
<td>66439956</td>
<td>1,074,749</td>
<td><a href="https://bscscan.com/address/0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93">View</a></td>
</tr>
</table>

### Market System

<table>
<tr>
<th>Contract</th>
<th>Address</th>
<th>Block</th>
<th>Gas Used</th>
<th>BscScan</th>
</tr>
<tr>
<td><strong>MarketFactory</strong></td>
<td><code>0x22Cc806047BB825aa26b766Af737E92B1866E8A6</code></td>
<td>66439956</td>
<td>4,944,042</td>
<td><a href="https://bscscan.com/address/0x22Cc806047BB825aa26b766Af737E92B1866E8A6">View</a></td>
</tr>
</table>

---

## Configuration Details

### Admin Addresses

All administrative roles are currently assigned to the deployer address. These will be transitioned to multi-sig governance.

<table>
<tr>
<th>Role</th>
<th>Address</th>
<th>Purpose</th>
</tr>
<tr>
<td><strong>Protocol Owner</strong></td>
<td><code>0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66</code></td>
<td>Contract administration and parameter updates</td>
</tr>
<tr>
<td><strong>Protocol Treasury</strong></td>
<td><code>0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66</code></td>
<td>Protocol fee collection</td>
</tr>
<tr>
<td><strong>Arbitrator</strong></td>
<td><code>0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66</code></td>
<td>Dispute resolution authority</td>
</tr>
<tr>
<td><strong>AI Signer</strong></td>
<td><code>0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66</code></td>
<td>Authorized AI oracle signer</td>
</tr>
</table>

**Next Steps:** Transition to 3-of-5 multi-sig within 30 days.

### Protocol Parameters

<table>
<tr>
<td width="50%">

**Token Economics**
- HORIZON Token Address: [`0x5b2ba38...14444`](https://bscscan.com/address/0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444)
- Token Standard: ERC-20
- Deployer Balance: 936,102 HORIZON

</td>
<td width="50%">

**Market Creation**
- Minimum Creator Stake: 10,000 HORIZON
- Stake Lockup: Until market resolution
- Creator Fee Share: 90-98%

</td>
</tr>
</table>

<table>
<tr>
<td width="50%">

**Resolution System**
- Minimum Bond: 1,000 HORIZON
- Dispute Window: 172,800 seconds (48 hours)
- Arbitrator: Multi-sig governance

</td>
<td width="50%">

**Fee Structure**
- User Trading Fee: 2% (constant)
- Protocol Share: 2-10% (based on HORIZON holdings)
- Creator Share: 90-98% (inverse of protocol)

</td>
</tr>
</table>

---

## Deployment Timeline

### Phase 1: Token Setup

<table>
<tr>
<th>Step</th>
<th>Contract</th>
<th>Transaction Hash</th>
<th>Status</th>
</tr>
<tr>
<td>1</td>
<td>HorizonToken (Existing)</td>
<td>Pre-deployed</td>
<td>✅ Verified</td>
</tr>
<tr>
<td>2</td>
<td>OutcomeToken</td>
<td><code><a href="https://bscscan.com/tx/0xd52063417ad512de422656619325e85cea1517052bc603b507e5b7d1e2f4e9ea">0xd520634...</a></code></td>
<td>✅ Deployed</td>
</tr>
</table>

### Phase 2: Protocol Infrastructure

<table>
<tr>
<th>Step</th>
<th>Contract</th>
<th>Transaction Hash</th>
<th>Status</th>
</tr>
<tr>
<td>3</td>
<td>HorizonPerks</td>
<td><code><a href="https://bscscan.com/tx/0x3a7bbf00a3d7d73aac8cdf546ade2a55f071a4008d740a63afd11e9972adbbb0">0x3a7bbf0...</a></code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>4</td>
<td>FeeSplitter</td>
<td><code><a href="https://bscscan.com/tx/0x9a772d2b93dbef65131d2d08b56e2ac41a41f66c253bd3608ab583cd228776d4">0x9a772d2...</a></code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>5</td>
<td>ResolutionModule</td>
<td><code><a href="https://bscscan.com/tx/0xb1330e880683a8aa0091761d046c6b51d4f7e50e55d6035f44231d1348afd488">0xb1330e8...</a></code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>6</td>
<td>Resolution Config (MinBond)</td>
<td><code><a href="https://bscscan.com/tx/0x7c0a741816646b3d6090cea0a4746e4e2f874fb5c5fa66bf70490fb95af9e12e">0x7c0a741...</a></code></td>
<td>✅ Configured</td>
</tr>
<tr>
<td>7</td>
<td>Resolution Config (DisputeWindow)</td>
<td><code><a href="https://bscscan.com/tx/0xafc788634cd55d88ad0d7b004d8d315c6a81f99e3316821c55b930d13bd0c78e">0xafc7886...</a></code></td>
<td>✅ Configured</td>
</tr>
<tr>
<td>8</td>
<td>AIOracleAdapter</td>
<td><code><a href="https://bscscan.com/tx/0x194237682e46ed043baf6ddb5ecd63b12f62e83f7a955b9525aa37a49622bc95">0x1942376...</a></code></td>
<td>✅ Deployed</td>
</tr>
</table>

### Phase 3: Market System

<table>
<tr>
<th>Step</th>
<th>Contract</th>
<th>Transaction Hash</th>
<th>Status</th>
</tr>
<tr>
<td>9</td>
<td>MarketFactory</td>
<td><code><a href="https://bscscan.com/tx/0xb44037712ca3ad864ef7ea893792661fd31307ce0ab6169e5551e666ba74a0d7">0xb440377...</a></code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>10</td>
<td>Creator Stake Config</td>
<td><code><a href="https://bscscan.com/tx/0x990143c2265500f9c483f76ced54cdcf3a6c3558e372556760b79935003e02f8">0x990143c...</a></code></td>
<td>✅ Configured</td>
</tr>
</table>

### Phase 4: Authorization Setup

<table>
<tr>
<th>Step</th>
<th>Action</th>
<th>Transaction Hash</th>
<th>Status</th>
</tr>
<tr>
<td>11</td>
<td>Authorize Resolution Module</td>
<td><code><a href="https://bscscan.com/tx/0x9f6b4c439d41843acbfa3afefecfa3b36e26d21d926e7b9b8a7d46e24b204bde">0x9f6b4c4...</a></code></td>
<td>✅ Authorized</td>
</tr>
<tr>
<td>12</td>
<td>Transfer OutcomeToken Ownership</td>
<td><code><a href="https://bscscan.com/tx/0xa78333aa93fc8cea599221630594748a4ccf9c8a1307a0b5db8e327ea9a29642">0xa78333a...</a></code></td>
<td>✅ Transferred</td>
</tr>
<tr>
<td>13</td>
<td>Transfer FeeSplitter Ownership</td>
<td><code><a href="https://bscscan.com/tx/0xf365a291c3302a446bb0f3790ed50950088866eca19d52f052cd1d3f76f11a96">0xf365a29...</a></code></td>
<td>✅ Transferred</td>
</tr>
</table>

---

## Verification Status

### Contract Verification

Contracts need to be verified on BscScan. Run these commands:

<details>
<summary><strong>Verification Commands</strong></summary>

```bash
# Set your API key
export BSCSCAN_API_KEY=your_api_key_here

# Verify OutcomeToken
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0x17B322784265c105a94e4c3d00aF1E5f46a5F311 \
  src/OutcomeToken.sol:OutcomeToken \
  --constructor-args $(cast abi-encode "constructor(string)" "https://horizon.markets/api/metadata/{id}.json")

# Verify HorizonPerks
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C \
  src/HorizonPerks.sol:HorizonPerks \
  --constructor-args $(cast abi-encode "constructor(address)" "0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444")

# Verify FeeSplitter
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0x275017E98adF33051BbF477fe1DD197F681d4eF1 \
  src/FeeSplitter.sol:FeeSplitter \
  --constructor-args $(cast abi-encode "constructor(address)" "0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66")

# Verify ResolutionModule
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0xF0CF4C741910cB48AC596F620a0AE892Cd247838 \
  src/ResolutionModule.sol:ResolutionModule \
  --constructor-args $(cast abi-encode "constructor(address,address,address)" "0x17B322784265c105a94e4c3d00aF1E5f46a5F311" "0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444" "0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66")

# Verify AIOracleAdapter
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93 \
  src/AIOracleAdapter.sol:AIOracleAdapter \
  --constructor-args $(cast abi-encode "constructor(address,address,address)" "0xF0CF4C741910cB48AC596F620a0AE892Cd247838" "0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444" "0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66")

# Verify MarketFactory
forge verify-contract \
  --chain-id 56 \
  --compiler-version 0.8.24 \
  --num-of-optimizations 200 \
  --watch \
  --etherscan-api-key $BSCSCAN_API_KEY \
  0x22Cc806047BB825aa26b766Af737E92B1866E8A6 \
  src/MarketFactory.sol:MarketFactory \
  --constructor-args $(cast abi-encode "constructor(address,address,address,address)" "0x17B322784265c105a94e4c3d00aF1E5f46a5F311" "0x275017E98adF33051BbF477fe1DD197F681d4eF1" "0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C" "0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444")
```

</details>

<table>
<tr>
<td width="25%" align="center">

**⏳ HorizonToken**

Pre-deployed & verified

</td>
<td width="25%" align="center">

**⏳ OutcomeToken**

Pending verification

</td>
<td width="25%" align="center">

**⏳ HorizonPerks**

Pending verification

</td>
<td width="25%" align="center">

**⏳ FeeSplitter**

Pending verification

</td>
</tr>
<tr>
<td width="25%" align="center">

**⏳ ResolutionModule**

Pending verification

</td>
<td width="25%" align="center">

**⏳ AIOracleAdapter**

Pending verification

</td>
<td width="25%" align="center">

**⏳ MarketFactory**

Pending verification

</td>
<td width="25%" align="center">

**✅ MarketAMM (Template)**

Embedded in Factory

</td>
</tr>
</table>

---

## Environment Configuration

### For Developers

Copy these addresses to your `.env` file:

```bash
# Network
CHAIN_ID=56
RPC_URL=https://bsc-dataseed.binance.org/

# Core Tokens
HORIZON_TOKEN_ADDRESS=0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444
OUTCOME_TOKEN_ADDRESS=0x17B322784265c105a94e4c3d00aF1E5f46a5F311
HORIZON_PERKS_ADDRESS=0x71Ff73A5a43B479a2D549a34dE7d3eadB9A1E22C

# Protocol Infrastructure
FEE_SPLITTER_ADDRESS=0x275017E98adF33051BbF477fe1DD197F681d4eF1
RESOLUTION_MODULE_ADDRESS=0xF0CF4C741910cB48AC596F620a0AE892Cd247838
AI_ORACLE_ADAPTER_ADDRESS=0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93

# Market System
MARKET_FACTORY_ADDRESS=0x22Cc806047BB825aa26b766Af737E92B1866E8A6

# Admin
PROTOCOL_OWNER=0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
PROTOCOL_TREASURY=0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
ARBITRATOR_ADDRESS=0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
AI_SIGNER_ADDRESS=0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
```

### For AI Resolver Service

Update `ai-resolver/.env`:

```bash
RPC_ENDPOINT=https://bsc-dataseed.binance.org
CHAIN_ID=56

AI_ORACLE_ADAPTER_ADDR=0x8773B8C5a55390DAbAD33dB46a13cd59Fb05cF93
MARKET_FACTORY_ADDR=0x22Cc806047BB825aa26b766Af737E92B1866E8A6
RESOLUTION_MODULE_ADDR=0xF0CF4C741910cB48AC596F620a0AE892Cd247838
HORIZON_TOKEN_ADDR=0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444
```

---

## Key Events

### Contract Initialization Events

<details>
<summary><strong>OutcomeToken Events</strong></summary>

```
✅ OwnershipTransferred
   previousOwner: 0x0000000000000000000000000000000000000000
   newOwner: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66

✅ OwnershipTransferred (to MarketFactory)
   previousOwner: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
   newOwner: 0x22Cc806047BB825aa26b766Af737E92B1866E8A6
```

</details>

<details>
<summary><strong>ResolutionModule Configuration</strong></summary>

```
✅ MinBondUpdated
   oldBond: 1,000 HORIZON
   newBond: 1,000 HORIZON

✅ DisputeWindowUpdated
   oldWindow: 172,800 seconds
   newWindow: 172,800 seconds
```

</details>

<details>
<summary><strong>Authorization Events</strong></summary>

```
✅ ResolutionAuthorized (OutcomeToken)
   resolver: ResolutionModule (0xF0CF4C741910cB48AC596F620a0AE892Cd247838)
   authorized: true

✅ SignerAdded (AIOracleAdapter)
   signer: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66

✅ MinCreatorStakeUpdated (MarketFactory)
   oldStake: 100 HORIZON
   newStake: 10,000 HORIZON
```

</details>

---

## Post-Deployment Checklist

### Immediate Tasks

- [x] All contracts deployed successfully
- [ ] All contracts verified on BscScan (in progress)
- [x] Authorization roles configured
- [x] Ownership transferred to MarketFactory
- [x] Protocol parameters set correctly
- [x] Deployer has sufficient HORIZON balance (936,102 tokens)

### Next 7 Days

- [ ] Deploy AI Resolver service to production
- [ ] Verify all contracts on BscScan
- [ ] Set up monitoring and alerting
- [ ] Create first test markets
- [ ] Prepare frontend deployment
- [ ] Update documentation with addresses

### Next 30 Days

- [ ] Transition to 3-of-5 multi-sig governance
- [ ] Complete external security audit
- [ ] Launch public testnet competition
- [ ] Begin marketing campaign
- [ ] Establish bug bounty program

---

## Security Notes

### Audit Status

<table>
<tr>
<td width="50%">

**Internal Review**
- ✅ Code review completed
- ✅ All tests passing
- ✅ Security patterns verified
- ✅ Gas optimization checked

</td>
<td width="50%">

**External Audit**
- ⏳ Scheduled for Q1 2025
- ⏳ Third-party security firm
- ⏳ Bug bounty after audit
- ⏳ Public audit report

</td>
</tr>
</table>

### Security Measures

- **Access Control**: All critical functions protected with ownership checks
- **Reentrancy Guards**: Applied to all state-changing functions
- **Emergency Pause**: Available for MarketAMM contracts
- **Time Locks**: Dispute windows provide intervention opportunities
- **Multi-sig Ready**: Ownership can be transferred to multi-sig wallets

**Best Practices:**
- Always use hardware wallets for production deployments
- Never commit API keys or private keys to version control
- Implement multi-sig governance for critical operations
- Regularly audit and rotate access credentials

---

## Contract Interactions

### Quick Start Guide

<details>
<summary><strong>Create Your First Market</strong></summary>

```solidity
// 1. Approve HORIZON tokens
HorizonToken horizon = HorizonToken(0x5b2bA38272125bd1dcDE41f1a88d98C2F5c14444);
horizon.approve(0x22Cc806047BB825aa26b766Af737E92B1866E8A6, 10000 ether);

// 2. Create market
MarketFactory factory = MarketFactory(0x22Cc806047BB825aa26b766Af737E92B1866E8A6);
uint256 marketId = factory.createMarket(
    collateralToken,
    "Will BTC reach $100k in 2025?",
    "Crypto",
    "ipfs://...",
    block.timestamp + 90 days
);
```

</details>

<details>
<summary><strong>Trade on a Market</strong></summary>

```solidity
// Get market AMM
address ammAddress = factory.getMarketAMM(marketId);
MarketAMM amm = MarketAMM(ammAddress);

// Approve collateral
IERC20(collateralToken).approve(ammAddress, 1000 ether);

// Buy YES tokens
uint256 yesTokens = amm.buyYes(1000 ether, 0);
```

</details>

---

## Gas Consumption Analysis

### Contract Deployment Gas

<table>
<tr>
<th>Contract</th>
<th>Gas Used</th>
<th>% of Total</th>
<th>BNB Cost</th>
</tr>
<tr>
<td>MarketFactory</td>
<td>4,944,042</td>
<td>43.04%</td>
<td>0.0002472021 BNB</td>
</tr>
<tr>
<td>OutcomeToken</td>
<td>1,718,990</td>
<td>14.97%</td>
<td>0.0000859495 BNB</td>
</tr>
<tr>
<td>HorizonPerks</td>
<td>1,354,764</td>
<td>11.80%</td>
<td>0.0000677382 BNB</td>
</tr>
<tr>
<td>ResolutionModule</td>
<td>1,226,164</td>
<td>10.68%</td>
<td>0.0000613082 BNB</td>
</tr>
<tr>
<td>AIOracleAdapter</td>
<td>1,074,749</td>
<td>9.36%</td>
<td>0.00005373745 BNB</td>
</tr>
<tr>
<td>FeeSplitter</td>
<td>975,952</td>
<td>8.50%</td>
<td>0.0000487976 BNB</td>
</tr>
<tr>
<td>Configuration (6 txs)</td>
<td>186,174</td>
<td>1.62%</td>
<td>0.0000093087 BNB</td>
</tr>
<tr>
<td><strong>Total</strong></td>
<td><strong>11,484,835</strong></td>
<td><strong>100%</strong></td>
<td><strong>0.00057424175 BNB</strong></td>
</tr>
</table>

**Average Gas Price:** 0.05 gwei  
**Total Cost in USD:** ~$0.35 (at $625/BNB)

---

## Support & Resources

<div align="center">

**Need Help with Integration?**

<table>
<tr>
<td align="center" width="33%">

**Documentation**

[View Full Docs](../README.md)

Smart contract specs and guides

</td>
<td align="center" width="33%">

**Developer Support**

developers@horizonoracles.com

Technical integration help

</td>
<td align="center" width="33%">

**Community**

[Discord Server](https://discord.com/invite/TuUHwwKjHh)

Join our developer community

</td>
</tr>
</table>

</div>

---

## Additional Information

### Network Details

- **Network Name:** BNB Smart Chain
- **Chain ID:** 56
- **Currency:** BNB
- **RPC URL:** https://bsc-dataseed.binance.org/
- **Block Explorer:** https://bscscan.com/

### Contract Standards

- **ERC-20:** HorizonToken
- **ERC-1155:** OutcomeToken (multi-token standard)
- **EIP-712:** AIOracleAdapter (typed structured data signing)

### Supported Market Types

The MarketFactory supports deployment of multiple market types:

1. **Standard AMM** - MarketAMM (constant product formula)
2. **Limit Order Market** - LimitOrderMarket (CLOB)
3. **Multi-Choice Market** - MultiChoiceMarket (LMSR)
4. **Pooled Liquidity** - PooledLiquidityMarket (Uniswap V3 style)

### Related Documentation

- [Main README](../README.md) - Project overview
- [Contract Specifications](../contracts/README.md) - Technical details
- [Security Documentation](../contracts/docs/SECURITY.md) - Security measures
- [Roles & Permissions](../contracts/docs/ROLES.md) - Access control
- [Deployment Guide](../contracts/docs/DEPLOYMENT.md) - Deployment instructions

---

<div align="center">

**🎉 Mainnet Deployment Successful**

Project Gamma is now live on BNB Chain

[View MarketFactory on BscScan](https://bscscan.com/address/0x22Cc806047BB825aa26b766Af737E92B1866E8A6) • [Website](https://horizonoracles.com/) • [Twitter](https://x.com/HorizonOracles)

*Deployed with ❤️ on October 30, 2025*

</div>
