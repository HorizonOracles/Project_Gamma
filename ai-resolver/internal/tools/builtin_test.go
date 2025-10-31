package tools

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Mock MarketDataClient for testing
type mockMarketDataClient struct {
	market MarketInfo
	err    error
}

func (m *mockMarketDataClient) GetMarket(ctx context.Context, marketID *big.Int) (MarketInfo, error) {
	if m.err != nil {
		return MarketInfo{}, m.err
	}
	return m.market, nil
}

func (m *mockMarketDataClient) GetBalance(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1000000), nil
}

func (m *mockMarketDataClient) GetCurrentBlockTimestamp(ctx context.Context) (int64, error) {
	return time.Now().Unix(), nil
}

func TestWebSearchTool(t *testing.T) {
	tool := NewWebSearchTool()

	if tool.Name() != "web_search" {
		t.Errorf("expected name 'web_search', got %s", tool.Name())
	}

	if tool.Type() != ToolTypeWebSearchPreview {
		t.Errorf("expected type %s, got %s", ToolTypeWebSearchPreview, tool.Type())
	}

	// Execute should work but return delegated message
	ctx := context.Background()
	output, err := tool.Execute(ctx, ToolInput{CallID: "test_1"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if output.Data == nil {
		t.Error("expected data in output")
	}

	// Check OpenAI format
	format := tool.ToOpenAIFormat()
	if format["type"] != "web_search_preview" {
		t.Errorf("expected type 'web_search_preview', got %v", format["type"])
	}
}

func TestMarketDataTool(t *testing.T) {
	mockClient := &mockMarketDataClient{
		market: MarketInfo{
			ID:              big.NewInt(123),
			Creator:         common.HexToAddress("0x1234567890123456789012345678901234567890"),
			AMM:             common.HexToAddress("0x2345678901234567890123456789012345678901"),
			CollateralToken: common.HexToAddress("0x3456789012345678901234567890123456789012"),
			CloseTime:       big.NewInt(1234567890),
			Category:        "sports",
			MetadataURI:     "ipfs://test",
			CreatorStake:    big.NewInt(1000),
			StakeRefunded:   false,
			Status:          0,
		},
	}

	tool := NewMarketDataTool(mockClient)

	if tool.Name() != "get_market_data" {
		t.Errorf("expected name 'get_market_data', got %s", tool.Name())
	}

	if tool.Type() != ToolTypeFunction {
		t.Errorf("expected type %s, got %s", ToolTypeFunction, tool.Type())
	}

	// Test successful execution
	ctx := context.Background()
	input := ToolInput{
		CallID: "test_1",
		Arguments: map[string]any{
			"market_id": "123",
		},
	}

	output, err := tool.Execute(ctx, input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	data, ok := output.Data.(map[string]any)
	if !ok {
		t.Fatal("expected data to be map[string]any")
	}

	if data["market_id"] != "123" {
		t.Errorf("expected market_id '123', got %v", data["market_id"])
	}

	if data["category"] != "sports" {
		t.Errorf("expected category 'sports', got %v", data["category"])
	}

	// Test validation error - missing market_id
	input2 := ToolInput{
		CallID:    "test_2",
		Arguments: map[string]any{},
	}

	_, err = tool.Execute(ctx, input2)
	if err == nil {
		t.Error("expected validation error for missing market_id")
	}

	// Test invalid market_id format
	input3 := ToolInput{
		CallID: "test_3",
		Arguments: map[string]any{
			"market_id": "not_a_number",
		},
	}

	output3, err := tool.Execute(ctx, input3)
	if err == nil {
		t.Error("expected error for invalid market_id format")
	}
	if output3.Error == nil {
		t.Error("expected error in output")
	}

	// Check OpenAI format
	format := tool.ToOpenAIFormat()
	if format["type"] != "function" {
		t.Errorf("expected type 'function', got %v", format["type"])
	}

	fn, ok := format["function"].(map[string]any)
	if !ok {
		t.Fatal("expected function field")
	}

	if fn["name"] != "get_market_data" {
		t.Errorf("expected name 'get_market_data', got %v", fn["name"])
	}
}
