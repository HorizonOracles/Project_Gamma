#!/bin/bash

# Load environment variables
if [ -f .env ]; then
    source .env
else
    echo "Error: .env file not found"
    exit 1
fi

# Check required variables
if [ -z "$HORIZON_TOKEN_ADDRESS" ]; then
    echo "Error: HORIZON_TOKEN_ADDRESS not set in .env"
    exit 1
fi

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set in .env"
    exit 1
fi

# Prompt for private key (don't store in .env)
echo "Enter your private key (input hidden):"
read -s PRIVATE_KEY

if [ -z "$PRIVATE_KEY" ]; then
    echo "Error: Private key cannot be empty"
    exit 1
fi

echo ""
echo "=========================================="
echo "Mainnet Deployment Configuration"
echo "=========================================="
echo "RPC URL: $RPC_URL"
echo "Token Address: $HORIZON_TOKEN_ADDRESS"
echo "Chain ID: $CHAIN_ID"
echo "=========================================="
echo ""
echo "⚠️  This will deploy to MAINNET and spend real funds!"
echo ""
read -p "Continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo "Deployment cancelled"
    exit 0
fi

echo ""
echo "Starting deployment..."
echo ""

# Run deployment
forge script script/DeployMainnet.s.sol:DeployMainnet \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    --verify \
    --etherscan-api-key $ETHERSCAN_API_KEY \
    -vvv

echo ""
echo "Deployment complete!"
