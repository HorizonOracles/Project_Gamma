<div align="center">

# BNB Liquid Staking

**Horizon Oracles Liquid Staking Platform**

[![Network](https://img.shields.io/badge/Network-BNB%20Chain-F0B90B?style=flat-square&logo=binance)](https://bscscan.com/)
[![aBNBc](https://img.shields.io/badge/Token-aBNBc-blue?style=flat-square)](https://bscscan.com/address/0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A)
[![APY](https://img.shields.io/badge/APY-~1.05%25-success?style=flat-square)]()

*Liquid staking interface with real-time blockchain integration and instant liquidity*

---

[Overview](#overview) • [How It Works](#how-it-works) • [Technical Details](#technical-architecture) • [API Reference](#api-endpoints)

---

</div>

## Overview

Horizon Oracles provides BNB liquid staking, enabling users to stake BNB and receive aBNBc (BNB Reward-Bearing Certificate) tokens. This implementation provides instant liquidity, automatic rewards, and seamless DeFi composability.

### What is Liquid Staking?

<table>
<tr>
<td width="25%" align="center">

**Instant Liquidity**

Receive aBNBc tokens immediately upon staking

</td>
<td width="25%" align="center">

**Automatic Rewards**

Rewards accrue within the aBNBc token value

</td>
<td width="25%" align="center">

**DeFi Composability**

Use aBNBc in protocols while earning rewards

</td>
<td width="25%" align="center">

**No Lock-up**

Unstake anytime with no waiting period

</td>
</tr>
</table>

---

## aBNBc Token

### Contract Details

<table>
<tr>
<th>Property</th>
<th>Value</th>
</tr>
<tr>
<td><strong>Contract Address</strong></td>
<td><code>0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A</code></td>
</tr>
<tr>
<td><strong>Token Standard</strong></td>
<td>ERC-20</td>
</tr>
<tr>
<td><strong>Decimals</strong></td>
<td>18</td>
</tr>
<tr>
<td><strong>Network</strong></td>
<td>BNB Chain (Chain ID: 56)</td>
</tr>
<tr>
<td><strong>Block Explorer</strong></td>
<td><a href="https://bscscan.com/address/0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A">View on BscScan</a></td>
</tr>
</table>

### Reward Mechanism

aBNBc uses a **rebasing exchange rate** model where your token quantity stays constant but becomes redeemable for more BNB over time.

<table>
<tr>
<td width="50%">

**How It Works**

1. Token quantity never changes
2. Exchange rate increases daily
3. Value appreciates automatically
4. No manual claiming needed

</td>
<td width="50%">

**Example**

```
Day 1:  Stake 100 BNB
        → Receive 110.11 aBNBc
        (ratio: 0.9084)

Day 30: Hold 110.11 aBNBc
        → Redeemable for 100.5 BNB
        (ratio: 0.9129)
```

</td>
</tr>
</table>

### Exchange Rate Calculation

**Contract Function:**
```solidity
function ratio() view returns (uint256)
```

**Conversion Formulas:**

<table>
<tr>
<td width="50%">

**BNB to aBNBc**

```
aBNBc Amount = BNB Amount × Exchange Rate
```

Example:
```
100 BNB × 0.9084 = 90.84 aBNBc
```

</td>
<td width="50%">

**aBNBc to BNB**

```
BNB Amount = aBNBc Amount ÷ Exchange Rate
```

Example:
```
90.84 aBNBc ÷ 0.9084 = 100 BNB
```

</td>
</tr>
</table>

**Current Typical Values:**
- Exchange Rate: ~0.9084 (updates continuously)
- APY: ~1.05% (annual percentage yield)

---

## How It Works

### Staking Process

<details>
<summary><strong>Step-by-Step User Flow</strong></summary>

**1. Connect Wallet**
- User connects Web3 wallet (MetaMask, Trust Wallet, etc.)
- System validates BNB Chain connection

**2. Enter Amount**
- User specifies BNB amount to stake
- Interface shows expected aBNBc tokens

**3. Validation**
- Minimum stake: 0.1 BNB
- Sufficient BNB balance check
- Network verification (Chain ID: 56)

**4. Transaction Submission**
- User approves transaction in wallet
- Gas fee displayed (~0.0005 BNB)

**5. Blockchain Confirmation**
- Transaction processes on-chain
- Block confirmation received

**6. Receive aBNBc**
- aBNBc tokens deposited to wallet
- Balance updates immediately

**7. Event Recording**
- Platform logs staking event
- Database records transaction details

</details>

**Smart Contract Interaction:**

```typescript
// Function called
writeContract({
  address: '0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A',
  abi: ABNBC_ABI,
  functionName: 'stakeCerts',
  value: parseEther(amount), // BNB amount in wei
})
```

**Requirements:**

| Parameter | Value |
|-----------|-------|
| Minimum Amount | 0.1 BNB |
| Network | BNB Chain (56) |
| Gas Fee | ~0.0005 BNB |

### Unstaking Process

<details>
<summary><strong>Step-by-Step User Flow</strong></summary>

**1. Select Unstake**
- User switches to unstake interface
- System displays aBNBc balance

**2. Enter Amount**
- User specifies BNB equivalent to unstake
- Interface calculates aBNBc to burn

**3. Validation**
- Minimum unstake: 0.5 BNB equivalent
- Sufficient aBNBc balance check
- Valid approval verification

**4. Transaction Submission**
- User approves transaction
- Gas fee displayed

**5. Blockchain Confirmation**
- aBNBc tokens burned
- BNB returned to wallet

**6. Event Recording**
- Platform logs unstaking event
- Database updated

</details>

**Smart Contract Interaction:**

```typescript
// Function called
writeContract({
  address: '0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A',
  abi: ABNBC_ABI,
  functionName: 'unstakeCerts',
  args: [parseEther(amount)], // aBNBc amount in wei
})
```

**Requirements:**

| Parameter | Value |
|-----------|-------|
| Minimum Amount | 0.5 BNB equivalent |
| Required Balance | Sufficient aBNBc |
| Gas Fee | ~0.0005 BNB |

---

## Technical Architecture

### Frontend Components

**File:** `client/src/pages/Staking.tsx`

<table>
<tr>
<td width="50%">

**Key Features**
- Dual-tab interface (Stake/Unstake)
- Real-time balance display
- Transaction state management
- Wallet connection handling
- Form validation
- Confirmation tracking

</td>
<td width="50%">

**Core Hooks**
- `useStakeBNB()` - Staking transactions
- `useUnstakeBNB()` - Unstaking transactions
- `useABNBcBalance()` - Balance fetching
- Validation functions

</td>
</tr>
</table>

**File:** `client/src/lib/ankr-staking.ts`

### Backend Services

**File:** `server/services/stakingMetrics.ts`

#### Real-Time Data Providers

<table>
<tr>
<th>Data Type</th>
<th>Cache Duration</th>
<th>Source</th>
<th>Fallback</th>
</tr>
<tr>
<td><strong>Exchange Rate</strong></td>
<td>60 seconds</td>
<td>Smart Contract <code>ratio()</code></td>
<td>0.9084</td>
</tr>
<tr>
<td><strong>Total Staked</strong></td>
<td>60 seconds</td>
<td>Smart Contract <code>totalSupply()</code></td>
<td>1600 BNB</td>
</tr>
<tr>
<td><strong>BNB Price</strong></td>
<td>120 seconds</td>
<td>CoinGecko API</td>
<td>$600</td>
</tr>
<tr>
<td><strong>APY</strong></td>
<td>1 hour</td>
<td>Staking Protocol API</td>
<td>1.05%</td>
</tr>
<tr>
<td><strong>Gas Fees</strong></td>
<td>Live</td>
<td>Blockchain RPC</td>
<td>0.0005 BNB</td>
</tr>
<tr>
<td><strong>Staker Count</strong></td>
<td>5 minutes</td>
<td>Database</td>
<td>0</td>
</tr>
</table>

**Cache Implementation:**

```typescript
private cache = new Map<string, CachedData<any>>();

private getCached<T>(key: string): T | null {
  const cached = this.cache.get(key);
  if (!cached) return null;
  
  const now = Date.now();
  if (now - cached.timestamp > cached.ttl) {
    this.cache.delete(key);
    return null;
  }
  
  return cached.data as T;
}
```

### Database Schema

**Table:** `staking_events`

```typescript
{
  id: varchar,              // UUID
  walletAddress: varchar,   // User's wallet address
  eventType: varchar,       // 'stake' or 'unstake'
  amount: varchar,          // BNB amount
  txHash: varchar,          // Transaction hash (unique)
  timestamp: timestamp,     // Event timestamp
  blockNumber: varchar      // Block number
}
```

**Unique Index:** `tx_hash` (prevents duplicate records)

### Event Recording

Events are recorded using a **fire-and-forget** pattern:

```typescript
// Backend API endpoint
POST /api/staking/events

// Request body
{
  walletAddress: "0x...",
  eventType: "stake",
  amount: "1.5",
  txHash: "0x...",
  blockNumber: "12345678"
}
```

**Deduplication:** Database enforces unique `tx_hash` constraint with `onConflictDoNothing()` strategy.

---

## Wallet Attribution System

### Problem & Solution

**Problem:** When users switch wallets mid-transaction, events could be attributed incorrectly.

<table>
<tr>
<td width="50%">

**Without Attribution Tracking**

1. User connects Wallet A
2. Initiates stake transaction
3. Switches to Wallet B
4. Transaction from Wallet A confirms
5. ❌ Event recorded for Wallet B

</td>
<td width="50%">

**With Attribution Tracking**

1. User connects Wallet A
2. System captures Wallet A address
3. Switches to Wallet B
4. Transaction from Wallet A confirms
5. ✅ Event correctly recorded for Wallet A

</td>
</tr>
</table>

**Implementation:**

```typescript
const [originatingAddress, setOriginatingAddress] = useState<string | null>(null);

const stake = (amount: string, address: string) => {
  setOriginatingAddress(address); // Capture initiating wallet
  writeContract({...});
};

// When transaction confirms:
recordEvent({
  walletAddress: originatingAddress, // Use captured address
  txHash: hash,
  ...
});
```

### Hash Tracking Reset

When wallet address changes, pending transaction hashes are cleared:

```typescript
useEffect(() => {
  if (address !== previousAddress) {
    setStakeHash(null); // Clear when wallet changes
  }
}, [address]);
```

This prevents attempting to record events for transactions from different wallets.

---

## Security

### Transaction Security

<table>
<tr>
<td width="50%">

**Validation Checks**
- Network verification (Chain ID: 56)
- Amount validation (minimums enforced)
- Balance verification before transaction
- Gas estimation displayed upfront

</td>
<td width="50%">

**Smart Contract Security**
- Audited production contract
- No approvals for staking (payable)
- Standard ERC-20 for aBNBc
- Established token patterns

</td>
</tr>
</table>

### Event Integrity

- **Unique Transaction Hashes:** Database constraint prevents duplicates
- **Wallet Attribution:** Events correctly attributed to originating wallet
- **Error Handling:** Failed transactions don't create database records
- **Idempotent Operations:** Safe to retry failed requests

---

## API Endpoints

### Get Staking Metrics

<table>
<tr>
<td width="30%">

**Endpoint**

```
GET /api/staking/metrics
```

</td>
<td width="70%">

**Response**

```json
{
  "totalStaked": "1600.5",
  "marketCap": "960300.00",
  "stakerCount": 3388,
  "apy": "1.05",
  "exchangeRate": "0.9084",
  "bnbPrice": "600.00",
  "lastUpdated": "2025-10-31T13:45:00.000Z"
}
```

</td>
</tr>
</table>

### Get Gas Fees

<table>
<tr>
<td width="30%">

**Endpoint**

```
GET /api/staking/fees
```

</td>
<td width="70%">

**Response**

```json
{
  "bnb": "0.000500",
  "usd": "0.30"
}
```

</td>
</tr>
</table>

### Record Staking Event

<table>
<tr>
<td width="30%">

**Endpoint**

```
POST /api/staking/events
```

</td>
<td width="70%">

**Request**

```json
{
  "walletAddress": "0x742d35...",
  "eventType": "stake",
  "amount": "1.5",
  "txHash": "0x...",
  "blockNumber": "12345678"
}
```

**Response**

```json
{
  "success": true,
  "message": "Event recorded successfully"
}
```

</td>
</tr>
</table>

---

## Error Handling

### Common Errors

<table>
<tr>
<th>Error</th>
<th>Message</th>
<th>Solution</th>
</tr>
<tr>
<td><strong>Insufficient Balance</strong></td>
<td>Insufficient BNB balance</td>
<td>Add more BNB to wallet</td>
</tr>
<tr>
<td><strong>Below Minimum</strong></td>
<td>Minimum stake amount is 0.1 BNB</td>
<td>Stake at least 0.1 BNB</td>
</tr>
<tr>
<td><strong>Wrong Network</strong></td>
<td>Please switch to BNB Chain</td>
<td>Connect to Chain ID 56</td>
</tr>
<tr>
<td><strong>Transaction Rejected</strong></td>
<td>User rejected the transaction</td>
<td>Retry and approve in wallet</td>
</tr>
<tr>
<td><strong>Insufficient Gas</strong></td>
<td>Insufficient funds for gas</td>
<td>Add more BNB for gas fees</td>
</tr>
</table>

### Error Logging

All errors are logged for debugging:

```typescript
try {
  // ... operation
} catch (error) {
  console.error('Error in staking operation:', error);
  throw error; // Re-throw for UI handling
}
```

---

## Future Enhancements

### Planned Features

<table>
<tr>
<td width="25%" align="center">

**Transaction History**

Display past transactions and rewards earned

</td>
<td width="25%" align="center">

**Portfolio Analytics**

Charts and performance tracking

</td>
<td width="25%" align="center">

**Multi-Chain Support**

Extend to other assets and chains

</td>
<td width="25%" align="center">

**Notifications**

Email/push alerts for confirmations

</td>
</tr>
</table>

### Roadmap

**Phase 1: Current** (✅ Completed)
- Basic staking/unstaking interface
- Real-time metrics display
- Event tracking and attribution

**Phase 2: Q1 2025**
- Transaction history dashboard
- Reward tracking and analytics
- Portfolio value charts

**Phase 3: Q2 2025**
- Multi-asset liquid staking
- Enhanced analytics
- Mobile app integration

**Phase 4: Q3 2025**
- Cross-chain staking support
- Advanced notifications
- API for third-party integrations

---

## Resources

### Smart Contracts

<table>
<tr>
<td width="50%">

**aBNBc Contract**

Address: `0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A`

[View on BscScan](https://bscscan.com/address/0xE85aFCcDaFBE7F2B096f268e31ccE3da8dA2990A)

</td>
<td width="50%">

**Network Details**

- Network: BNB Chain
- Chain ID: 56
- Currency: BNB
- Explorer: https://bscscan.com

</td>
</tr>
</table>

### External APIs

- **CoinGecko:** BNB price data
- **Staking Protocol API:** APY and staking data
- **BNB Chain RPC:** Public blockchain nodes

### Documentation

- [BNB Chain Developer Docs](https://docs.bnbchain.org/)
- [Viem Documentation](https://viem.sh/)
- [Wagmi Documentation](https://wagmi.sh/)
- [ERC-20 Token Standard](https://eips.ethereum.org/EIPS/eip-20)

---

## Support

<div align="center">

**Need Help with Liquid Staking?**

<table>
<tr>
<td align="center" width="33%">

**GitHub**

[Open Issue](https://github.com/HorizonOracles/Project_Gamma/issues)

Report bugs or request features

</td>
<td align="center" width="33%">

**Discord**

[Join Community](https://discord.com/invite/TuUHwwKjHh)

Get help from the community

</td>
<td align="center" width="33%">

**Email**

developers@horizonoracles.com

Technical support

</td>
</tr>
</table>

</div>

---

<div align="center">

**BNB Liquid Staking** • Powered by Horizon Oracles

[Website](https://horizonoracles.com/) • [Twitter](https://x.com/HorizonOracles) • [Documentation](https://github.com/HorizonOracles/Project_Gamma)

*Last Updated: October 31, 2025*

</div>
