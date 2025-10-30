#!/usr/bin/env bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Directory paths
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
CONTRACTS_DIR="$SCRIPT_DIR/contracts"
BACKEND_DIR="$SCRIPT_DIR/ai-resolver"

# Anvil test accounts
DEPLOYER_ADDR="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
DEPLOYER_KEY="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

# Anvil URL
RPC_URL="http://127.0.0.1:8545"
CHAIN_ID=31337

# Cleanup function
cleanup() {
    echo -e "\n${YELLOW}Cleaning up...${NC}"
    if [ ! -z "$ANVIL_PID" ]; then
        kill $ANVIL_PID 2>/dev/null || true
        echo "Stopped Anvil"
    fi
    if [ ! -z "$BACKEND_PID" ]; then
        kill $BACKEND_PID 2>/dev/null || true
        echo "Stopped backend server"
    fi
}

trap cleanup EXIT

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  AI Resolution E2E Test with Anvil${NC}"
echo -e "${GREEN}========================================${NC}\n"

# Step 1: Start Anvil
echo -e "${YELLOW}Step 1: Starting Anvil...${NC}"
cd "$SCRIPT_DIR"
anvil --block-time 1 &
ANVIL_PID=$!
echo "Anvil started (PID: $ANVIL_PID)"
sleep 3

# Step 2: Deploy contracts
echo -e "\n${YELLOW}Step 2: Deploying contracts...${NC}"
cd "$CONTRACTS_DIR"

# Run the deployment script
forge script script/DeployLocal.s.sol:DeployLocal \
    --rpc-url $RPC_URL \
    --broadcast \
    --private-key $DEPLOYER_KEY \
    -vv

# Extract contract addresses from broadcast files
BROADCAST_DIR="$CONTRACTS_DIR/broadcast/DeployLocal.s.sol/$CHAIN_ID"
if [ ! -d "$BROADCAST_DIR" ]; then
    echo -e "${RED}Deployment failed - broadcast directory not found${NC}"
    exit 1
fi

# Parse the latest run for contract addresses
RUN_FILE=$(ls -t "$BROADCAST_DIR"/run-latest.json 2>/dev/null || ls -t "$BROADCAST_DIR"/run-*.json | head -1)

if [ ! -f "$RUN_FILE" ]; then
    echo -e "${RED}Could not find deployment run file${NC}"
    exit 1
fi

echo -e "\n${GREEN}Parsing contract addresses...${NC}"

# Extract addresses using jq
TOKEN_ADDR=$(jq -r '.transactions[] | select(.contractName == "Token") | .contractAddress' "$RUN_FILE" | head -1)
OUTCOME_TOKEN_ADDR=$(jq -r '.transactions[] | select(.contractName == "OutcomeToken") | .contractAddress' "$RUN_FILE" | head -1)
RESOLUTION_MODULE_ADDR=$(jq -r '.transactions[] | select(.contractName == "ResolutionModule") | .contractAddress' "$RUN_FILE" | head -1)
AI_ORACLE_ADAPTER_ADDR=$(jq -r '.transactions[] | select(.contractName == "AIOracleAdapter") | .contractAddress' "$RUN_FILE" | head -1)
MARKET_FACTORY_ADDR=$(jq -r '.transactions[] | select(.contractName == "MarketFactory") | .contractAddress' "$RUN_FILE" | head -1)

# If jq parsing fails, try alternative method
if [ -z "$TOKEN_ADDR" ] || [ "$TOKEN_ADDR" == "null" ]; then
    echo -e "${YELLOW}Trying alternative address extraction...${NC}"
    TOKEN_ADDR=$(grep -oP '"contractName":"Token".*?"contractAddress":"\K0x[a-fA-F0-9]{40}' "$RUN_FILE" | head -1)
    OUTCOME_TOKEN_ADDR=$(grep -oP '"contractName":"OutcomeToken".*?"contractAddress":"\K0x[a-fA-F0-9]{40}' "$RUN_FILE" | head -1)
    RESOLUTION_MODULE_ADDR=$(grep -oP '"contractName":"ResolutionModule".*?"contractAddress":"\K0x[a-fA-F0-9]{40}' "$RUN_FILE" | head -1)
    AI_ORACLE_ADAPTER_ADDR=$(grep -oP '"contractName":"AIOracleAdapter".*?"contractAddress":"\K0x[a-fA-F0-9]{40}' "$RUN_FILE" | head -1)
    MARKET_FACTORY_ADDR=$(grep -oP '"contractName":"MarketFactory".*?"contractAddress":"\K0x[a-fA-F0-9]{40}' "$RUN_FILE" | head -1)
fi

echo "Token: $TOKEN_ADDR"
echo "OutcomeToken: $OUTCOME_TOKEN_ADDR"
echo "ResolutionModule: $RESOLUTION_MODULE_ADDR"
echo "AIOracleAdapter: $AI_ORACLE_ADAPTER_ADDR"
echo "MarketFactory: $MARKET_FACTORY_ADDR"

# Verify addresses
if [ -z "$TOKEN_ADDR" ] || [ -z "$MARKET_FACTORY_ADDR" ]; then
    echo -e "${RED}Failed to extract contract addresses${NC}"
    exit 1
fi

# Step 3: Create a test market
echo -e "\n${YELLOW}Step 3: Creating test market...${NC}"

# Calculate close time (1 hour from now)
CLOSE_TIME=$(($(date +%s) + 3600))

# Create market using cast
echo "Creating market with question: 'Will Bitcoin reach $100k by end of 2025?'"

MARKET_CREATE_DATA=$(cast abi-encode "createMarket((address,uint256,string,string,uint256))" \
    "($TOKEN_ADDR,$CLOSE_TIME,\"crypto\",\"Will Bitcoin reach \$100k by end of 2025?\",10000000000000000000000)")

MARKET_TX=$(cast send $MARKET_FACTORY_ADDR \
    "createMarket((address,uint256,string,string,uint256))" \
    "($TOKEN_ADDR,$CLOSE_TIME,\"crypto\",\"Will Bitcoin reach \$100k by end of 2025?\",10000000000000000000000)" \
    --rpc-url $RPC_URL \
    --private-key $DEPLOYER_KEY \
    --json | jq -r '.transactionHash')

echo "Market creation tx: $MARKET_TX"
sleep 2

# Get market ID from logs (should be 1 for first market)
MARKET_ID=1
echo "Market ID: $MARKET_ID"

# Step 4: Configure and start backend
echo -e "\n${YELLOW}Step 4: Starting AI Resolver backend...${NC}"
cd "$BACKEND_DIR"

# Check if OpenAI API key is set
if [ -z "$OPENAI_API_KEY" ]; then
    echo -e "${RED}Warning: OPENAI_API_KEY not set. Backend will fail without it.${NC}"
    echo "Set it with: export OPENAI_API_KEY=your-key"
    exit 1
fi

# Create .env file for backend
cat > .env <<EOF
# Network Configuration
RPC_ENDPOINT=$RPC_URL
CHAIN_ID=$CHAIN_ID

# Contract Addresses
AI_ORACLE_ADAPTER_ADDR=$AI_ORACLE_ADAPTER_ADDR
RESOLUTION_MODULE_ADDR=$RESOLUTION_MODULE_ADDR
TOKEN_ADDR=$TOKEN_ADDR
MARKET_FACTORY_ADDR=$MARKET_FACTORY_ADDR

# Signer Configuration
SIGNER_PRIVATE_KEY=$DEPLOYER_KEY

# AI Configuration
OPENAI_API_KEY=$OPENAI_API_KEY
OPENAI_MODEL=gpt-4-turbo-preview

# Bond Configuration
DEFAULT_BOND_AMOUNT=1000000000000000000000

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
EOF

echo "Backend .env configured"

# Build and start backend
echo "Building backend..."
go build -o server ./cmd/server/main.go

echo "Starting backend server..."
./server &
BACKEND_PID=$!
echo "Backend started (PID: $BACKEND_PID)"
sleep 3

# Step 5: Test AI resolution
echo -e "\n${YELLOW}Step 5: Testing AI resolution endpoint...${NC}"

# Wait for server to be ready
MAX_RETRIES=10
RETRY_COUNT=0
while [ $RETRY_COUNT -lt $MAX_RETRIES ]; do
    if curl -s http://localhost:8080/healthz > /dev/null 2>&1; then
        echo "Backend is ready"
        break
    fi
    echo "Waiting for backend to start..."
    sleep 2
    RETRY_COUNT=$((RETRY_COUNT + 1))
done

if [ $RETRY_COUNT -eq $MAX_RETRIES ]; then
    echo -e "${RED}Backend failed to start${NC}"
    exit 1
fi

# Check health
echo -e "\n${GREEN}Health check:${NC}"
curl -s http://localhost:8080/healthz | jq .

# Submit proposal
echo -e "\n${GREEN}Submitting AI proposal for market $MARKET_ID:${NC}"
PROPOSAL_RESPONSE=$(curl -s -X POST http://localhost:8080/v1/propose \
    -H "Content-Type: application/json" \
    -d "{
        \"marketId\": $MARKET_ID,
        \"closeTime\": $CLOSE_TIME,
        \"question\": \"Will Bitcoin reach \$100k by end of 2025?\",
        \"outcomeTokens\": [\"NO\", \"YES\"],
        \"metadata\": \"Test market for AI resolution\"
    }")

echo "$PROPOSAL_RESPONSE" | jq .

# Check if proposal was successful
STATUS=$(echo "$PROPOSAL_RESPONSE" | jq -r '.status // empty')
TX_HASH=$(echo "$PROPOSAL_RESPONSE" | jq -r '.txHash // empty')

if [ "$STATUS" == "submitted" ] && [ ! -z "$TX_HASH" ]; then
    echo -e "\n${GREEN}✓ AI proposal submitted successfully!${NC}"
    echo "Transaction: $TX_HASH"
    
    # Wait for transaction to be mined
    sleep 2
    
    # Get transaction receipt
    echo -e "\n${GREEN}Transaction receipt:${NC}"
    cast receipt $TX_HASH --rpc-url $RPC_URL | head -20
    
else
    echo -e "\n${RED}✗ AI proposal failed${NC}"
    echo "$PROPOSAL_RESPONSE"
    exit 1
fi

# Step 6: Verify proposal on-chain
echo -e "\n${YELLOW}Step 6: Verifying proposal on-chain...${NC}"

# Check proposal in ResolutionModule
PROPOSAL_DATA=$(cast call $RESOLUTION_MODULE_ADDR \
    "getProposal(uint256)(uint8,uint256,address,uint256,address,uint256,string,uint256,bool)" \
    $MARKET_ID \
    --rpc-url $RPC_URL)

echo "Proposal data: $PROPOSAL_DATA"

# Print summary
echo -e "\n${GREEN}========================================${NC}"
echo -e "${GREEN}  E2E Test Summary${NC}"
echo -e "${GREEN}========================================${NC}"
echo -e "Market ID: $MARKET_ID"
echo -e "Market Close Time: $(date -r $CLOSE_TIME '+%Y-%m-%d %H:%M:%S')"
echo -e "AI Proposal TX: $TX_HASH"
echo -e "Status: ${GREEN}SUCCESS${NC}"
echo -e "${GREEN}========================================${NC}\n"

# Keep services running for manual testing
echo -e "${YELLOW}Services are running. Press Ctrl+C to stop.${NC}"
echo "Anvil: http://127.0.0.1:8545"
echo "Backend: http://localhost:8080"
echo ""
echo "Test endpoints:"
echo "  curl http://localhost:8080/healthz"
echo "  curl http://localhost:8080/v1/markets"
echo ""

# Wait indefinitely
wait $BACKEND_PID
