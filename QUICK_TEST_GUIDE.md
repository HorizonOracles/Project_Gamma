# Quick Test Guide - Token3 Integration

## Prerequisites
- Anvil running on `http://127.0.0.1:8545`
- OpenAI API key (if testing AI resolution)

## 1. Start Anvil
```bash
cd contracts
anvil --block-time 1
```

## 2. Deploy Token3
```bash
cd contracts
forge script script/DeployToken3.s.sol \
  --fork-url http://localhost:8545 \
  --broadcast \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

Save the Token3 address from the output.

## 3. Deploy Utility Contracts
```bash
cd contracts
TOKEN_ADDRESS=<TOKEN3_ADDRESS> forge script script/DeployLocal.s.sol \
  --fork-url http://localhost:8545 \
  --broadcast \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

## 4. Run Solidity Integration Tests
```bash
cd contracts
forge script script/TestToken3Integration.s.sol \
  --fork-url http://localhost:8545 \
  --broadcast \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

## 5. Run Go Backend Integration Tests
```bash
cd ai-resolver
go run test_token_integration.go
```

## 6. Regenerate Go Bindings (if needed)
```bash
cd contracts
forge build
cd ../ai-resolver
abigen --abi ../contracts/out/Token3.sol/Token.abi.json \
  --pkg abi \
  --type Token \
  --out pkg/abi/Token.go
```

## Deployed Contract Addresses (Anvil)

- **Token3**: `0x5FbDB2315678afecb367f032d93F642f64180aa3`
- **MarketFactory**: `0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e`
- **ResolutionModule**: `0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6`
- **HorizonPerks**: `0x0165878A594ca255338adfa4d48449f69242Eb8F`

## Key Methods in Go Bindings

### Read Operations
```go
// Get balance
balance, err := token.BalanceOf(&bind.CallOpts{}, address)

// Check allowance
allowance, err := token.Allowance(&bind.CallOpts{}, owner, spender)

// Get mode constants
modeNormal, err := token.MODENORMAL(&bind.CallOpts{})
```

### Write Operations
```go
// Approve spending
tx, err := token.Approve(auth, spender, amount)

// Transfer tokens
tx, err := token.Transfer(auth, recipient, amount)

// Initialize token (owner only)
tx, err := token.Init(auth, "Name", "SYMBOL", totalSupply)

// Set mode (owner only)
tx, err := token.SetMode(auth, modeValue)
```

## Troubleshooting

### Issue: "Identifier already declared"
- This happens when importing Token3.sol alongside OpenZeppelin contracts
- Solution: Use separate deployment scripts or use IERC20 interface

### Issue: "Transfer is restricted"
- Token is in MODE_TRANSFER_RESTRICTED
- Solution: Call `SetMode(MODE_NORMAL)` from owner account

### Issue: Go binding missing methods
- ABI generation may have failed
- Solution: Recompile with `forge build` and regenerate bindings

## Environment Variables (.env.local)

```bash
RPC_URL=http://127.0.0.1:8545
CHAIN_ID=31337
SIGNER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
TOKEN_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3
FACTORY_ADDRESS=0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e
RESOLUTION_ADDRESS=0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6
OPENAI_API_KEY=your_key_here
```
