package tools

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MarketDataClient defines the interface for fetching market data from blockchain
// This allows the tool to work with the adapter.Client without direct coupling
type MarketDataClient interface {
	GetMarket(ctx context.Context, marketID *big.Int) (MarketInfo, error)
	GetBalance(ctx context.Context) (*big.Int, error)
	GetCurrentBlockTimestamp(ctx context.Context) (int64, error)
}

// MarketInfo represents market data returned from the blockchain
type MarketInfo struct {
	ID              *big.Int
	Creator         common.Address
	AMM             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}

// MarketDataTool fetches blockchain market data
type MarketDataTool struct {
	*BaseTool
	client MarketDataClient
}

// NewMarketDataTool creates a new market data tool
func NewMarketDataTool(client MarketDataClient) *MarketDataTool {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"market_id": {
				Type:        "string",
				Description: "The ID of the market to fetch (as a decimal string)",
			},
		},
		Required: []string{"market_id"},
	}

	base := NewBaseTool(
		"get_market_data",
		"Fetches detailed information about a prediction market from the blockchain, including creator, category, close time, and current status.",
		ToolTypeFunction,
		schema,
	)

	tool := &MarketDataTool{
		BaseTool: base,
		client:   client,
	}

	base.SetExecutor(tool.execute)

	return tool
}

// execute fetches market data from blockchain
func (t *MarketDataTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	// Extract market_id from arguments
	marketIDStr, ok := input.Arguments["market_id"].(string)
	if !ok {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("market_id must be a string"),
		}, fmt.Errorf("market_id must be a string")
	}

	// Parse market ID
	marketID, ok := new(big.Int).SetString(marketIDStr, 10)
	if !ok {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("invalid market_id format: %s", marketIDStr),
		}, fmt.Errorf("invalid market_id format: %s", marketIDStr)
	}

	// Fetch market data from blockchain
	market, err := t.client.GetMarket(ctx, marketID)
	if err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("failed to fetch market data: %w", err),
		}, fmt.Errorf("failed to fetch market data: %w", err)
	}

	// Build response data
	data := map[string]any{
		"market_id":        market.ID.String(),
		"creator":          market.Creator.Hex(),
		"amm":              market.AMM.Hex(),
		"collateral_token": market.CollateralToken.Hex(),
		"close_time":       market.CloseTime.Int64(),
		"category":         market.Category,
		"metadata_uri":     market.MetadataURI,
		"creator_stake":    market.CreatorStake.String(),
		"stake_refunded":   market.StakeRefunded,
		"status":           market.Status,
	}

	return ToolOutput{
		CallID: input.CallID,
		Data:   data,
	}, nil
}
