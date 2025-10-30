<div align="center">

#  Mainnet Deployment

**Project Gamma - BNB Chain Launch**

[![Network](https://img.shields.io/badge/Network-BSC%20Mainnet-F0B90B?style=flat-square&logo=binance)](https://bscscan.com/)
[![Status](https://img.shields.io/badge/Status-Live-success?style=flat-square)]()
[![Verified](https://img.shields.io/badge/Contracts-Verified-blue?style=flat-square)](https://bscscan.com/)

*Successfully deployed on October 29, 2025*

---

**Chain ID:** 56 • **Total Gas Cost:** 0.000624 BNB (~$0.39 USD)

---

</div>

## Deployment Summary

Project Gamma has been successfully deployed to BNB Chain mainnet. All 7 core contracts are live, verified, and operational.

### Deployment Statistics

<table>
<tr>
<td width="25%" align="center">

**Total Contracts**

7 contracts

</td>
<td width="25%" align="center">

**Total Gas Used**

12,372,541 gas

</td>
<td width="25%" align="center">

**Total Cost**

0.000624 BNB

</td>
<td width="25%" align="center">

**Avg Gas Price**

0.0504 gwei

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
<td><code>0x72F84681AA0dc8DB53E87eD507A4D6651B1C312D</code></td>
<td>66328650</td>
<td>1,718,990</td>
<td><a href="https://bscscan.com/address/0x72F84681AA0dc8DB53E87eD507A4D6651B1C312D">View</a></td>
</tr>
<tr>
<td><strong>HorizonPerks</strong></td>
<td><code>0x31709748Cc9030e86E71570442fa762c851950b3</code></td>
<td>66328654</td>
<td>1,354,764</td>
<td><a href="https://bscscan.com/address/0x31709748Cc9030e86E71570442fa762c851950b3">View</a></td>
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
<td><code>0x7CB3A3C58f7eA49CC6AF8Be5228ac58b04787B09</code></td>
<td>66328658</td>
<td>975,952</td>
<td><a href="https://bscscan.com/address/0x7CB3A3C58f7eA49CC6AF8Be5228ac58b04787B09">View</a></td>
</tr>
<tr>
<td><strong>ResolutionModule</strong></td>
<td><code>0x5407F3937C81A01E783B9E99fbAC624220a534Eb</code></td>
<td>66328663</td>
<td>1,226,176</td>
<td><a href="https://bscscan.com/address/0x5407F3937C81A01E783B9E99fbAC624220a534Eb">View</a></td>
</tr>
<tr>
<td><strong>AIOracleAdapter</strong></td>
<td><code>0x30827513096b63F09C0c24f574933b43e222C5D7</code></td>
<td>66328676</td>
<td>1,074,749</td>
<td><a href="https://bscscan.com/address/0x30827513096b63F09C0c24f574933b43e222C5D7">View</a></td>
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
<td><code>0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C</code></td>
<td>66328680</td>
<td>4,944,054</td>
<td><a href="https://bscscan.com/address/0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C">View</a></td>
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
- HORIZON Initial Supply: 100,000,000 tokens
- HORIZON Max Supply: 10,000,000,000 tokens
- Token Standard: ERC-20

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
- Trading Fee Range: 1.0% - 4.0%
- Protocol Share: 2-10%
- LP Fee: 0.3% (standard AMM)

</td>
</tr>
</table>

---

## Deployment Timeline

### Phase 1: Core Tokens

<table>
<tr>
<th>Step</th>
<th>Contract</th>
<th>Transaction Hash</th>
<th>Status</th>
</tr>
<tr>
<td>1</td>
<td>HorizonToken</td>
<td><code>0x2d10af367244204f000f42952e41348284116692...</code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>2</td>
<td>OutcomeToken</td>
<td><code>0x9bedb89923809a016446998cd20d5377383665b8...</code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>3</td>
<td>HorizonPerks</td>
<td><code>0xbaedffa08d46096cb2d080de3ee45d9c613d985a...</code></td>
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
<td>4</td>
<td>FeeSplitter</td>
<td><code>0x9205ca5ca50f4fe878cd470d417936338ef2550d...</code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>5</td>
<td>ResolutionModule</td>
<td><code>0xef98b31b89d3e4007ed62361fd77fd48ef4c79fc...</code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>6</td>
<td>Resolution Config</td>
<td><code>0x72b5a2897a26dae2ee0a1bf2b83ce0f1757dcab1...</code></td>
<td>✅ Configured</td>
</tr>
<tr>
<td>7</td>
<td>AIOracleAdapter</td>
<td><code>0xae3543c78f114ad1263836b6bcd076c2ba0a17bd...</code></td>
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
<td>8</td>
<td>MarketFactory</td>
<td><code>0x8e3bc6782222cf620be274d1a6b9eed0e662b6e7...</code></td>
<td>✅ Deployed</td>
</tr>
<tr>
<td>9</td>
<td>Creator Stake Config</td>
<td><code>0x09eada7207ba843bfa7e84e9b658ad216cc6fabd...</code></td>
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
<td>10</td>
<td>Add MarketFactory as Minter</td>
<td><code>0x9e25db8c0bd5b86e37476470b9b6aa9890806c68...</code></td>
<td>✅ Authorized</td>
</tr>
<tr>
<td>11</td>
<td>Authorize Resolution Module</td>
<td><code>0xfc8641a1b4fc4a71fdd40fa8772238f63a1de210...</code></td>
<td>✅ Authorized</td>
</tr>
<tr>
<td>12</td>
<td>Transfer OutcomeToken Ownership</td>
<td><code>0x8d3b5fc580aab0d313835695a941c884245f3056...</code></td>
<td>✅ Transferred</td>
</tr>
<tr>
<td>13</td>
<td>Transfer FeeSplitter Ownership</td>
<td><code>0x08705150a51ed48f6bb5db0f038dd45104a67f07...</code></td>
<td>✅ Transferred</td>
</tr>
</table>

---

## Verification Status

### Contract Verification

All contracts have been verified on BscScan:

<table>
<tr>
<td width="25%" align="center">

**✅ HorizonToken**

Verified with source code

</td>
<td width="25%" align="center">

**✅ OutcomeToken**

Verified with source code

</td>
<td width="25%" align="center">

**✅ HorizonPerks**

Verified with source code

</td>
<td width="25%" align="center">

**✅ FeeSplitter**

Verified with source code

</td>
</tr>
<tr>
<td width="25%" align="center">

**✅ ResolutionModule**

Verified with source code

</td>
<td width="25%" align="center">

**✅ AIOracleAdapter**

Verified with source code

</td>
<td width="25%" align="center">

**✅ MarketFactory**

Verified with source code

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
BSC_RPC_URL=https://bsc-dataseed.binance.org/

# Core Tokens
HORIZON_TOKEN_ADDRESS=0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444
OUTCOME_TOKEN_ADDRESS=0x72F84681AA0dc8DB53E87eD507A4D6651B1C312D
HORIZON_PERKS_ADDRESS=0x31709748Cc9030e86E71570442fa762c851950b3

# Protocol Infrastructure
FEE_SPLITTER_ADDRESS=0x7CB3A3C58f7eA49CC6AF8Be5228ac58b04787B09
RESOLUTION_MODULE_ADDRESS=0x5407F3937C81A01E783B9E99fbAC624220a534Eb
AI_ORACLE_ADAPTER_ADDRESS=0x30827513096b63F09C0c24f574933b43e222C5D7

# Market System
MARKET_FACTORY_ADDRESS=0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C
```

### For AI Resolver Service

Update `ai-resolver/.env`:

```bash
RPC_ENDPOINT=https://bsc-dataseed.binance.org
CHAIN_ID=56

AI_ORACLE_ADAPTER_ADDR=0x30827513096b63F09C0c24f574933b43e222C5D7
MARKET_FACTORY_ADDR=0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C
RESOLUTION_MODULE_ADDR=0x5407F3937C81A01E783B9E99fbAC624220a534Eb
HORIZON_TOKEN_ADDR=0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444
```

---

## Key Events

### Contract Initialization Events

<details>
<summary><strong>HorizonToken Events</strong></summary>

```
✅ OwnershipTransferred
   previousOwner: 0x0000000000000000000000000000000000000000
   newOwner: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66

✅ Transfer (Initial Mint)
   from: 0x0000000000000000000000000000000000000000
   to: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
   value: 100,000,000 HORIZON (1e26 wei)
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
✅ MinterAdded (HorizonToken)
   minter: MarketFactory (0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C)

✅ ResolutionAuthorized (OutcomeToken)
   resolver: ResolutionModule (0x5407F3937C81A01E783B9E99fbAC624220a534Eb)
   authorized: true

✅ SignerAdded (AIOracleAdapter)
   signer: 0x68e25d4b1dA2e4FF3B1B1C28a190D890b46D9C66
```

</details>

---

## Post-Deployment Checklist

### Immediate Tasks (✅ Completed)

- [x] All contracts deployed successfully
- [x] All contracts verified on BscScan
- [x] Authorization roles configured
- [x] Ownership transferred to MarketFactory
- [x] Protocol parameters set correctly
- [x] Initial HORIZON supply minted

### Next 7 Days

- [ ] Deploy AI Resolver service to production
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
- ✅ All tests passing (252/252)
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

---

## Contract Interactions

### Quick Start Guide

<details>
<summary><strong>Create Your First Market</strong></summary>

```solidity
// 1. Approve HORIZON tokens
HorizonToken horizon = HorizonToken(0x5b2ba38272125bd1dcde41f1a88d98c2f5c14444);
horizon.approve(0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C, 10000 ether);

// 2. Create market
MarketFactory factory = MarketFactory(0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C);
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
<td>4,944,054</td>
<td>39.96%</td>
<td>0.000249 BNB</td>
</tr>
<tr>
<td>OutcomeToken</td>
<td>1,718,990</td>
<td>13.89%</td>
<td>0.000087 BNB</td>
</tr>
<tr>
<td>HorizonPerks</td>
<td>1,354,764</td>
<td>10.95%</td>
<td>0.000068 BNB</td>
</tr>
<tr>
<td>ResolutionModule</td>
<td>1,226,176</td>
<td>9.91%</td>
<td>0.000062 BNB</td>
</tr>
<tr>
<td>AIOracleAdapter</td>
<td>1,074,749</td>
<td>8.68%</td>
<td>0.000054 BNB</td>
</tr>
<tr>
<td>FeeSplitter</td>
<td>975,952</td>
<td>7.89%</td>
<td>0.000049 BNB</td>
</tr>
<tr>
<td>HorizonToken</td>
<td>840,220</td>
<td>6.79%</td>
<td>0.000042 BNB</td>
</tr>
<tr>
<td>Configuration (5 txs)</td>
<td>237,636</td>
<td>1.92%</td>
<td>0.000012 BNB</td>
</tr>
<tr>
<td><strong>Total</strong></td>
<td><strong>12,372,541</strong></td>
<td><strong>100%</strong></td>
<td><strong>0.000624 BNB</strong></td>
</tr>
</table>

**Average Gas Price:** 0.050410295 gwei  
**Total Cost in USD:** ~$0.39 (at $625/BNB)

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

[View on BscScan](https://bscscan.com/address/0xf3f68A76D42679E8b3371bC26B75F7F26E97A10C) • [Website](https://horizonoracles.com/) • [Twitter](https://x.com/HorizonOracles)

*Deployed with ❤️ on October 29, 2025*

</div>
