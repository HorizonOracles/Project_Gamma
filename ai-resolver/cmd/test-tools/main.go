package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/project-gamma/ai-resolver/internal/config"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

// Mock market data client for testing
type mockMarketDataClient struct{}

func (m *mockMarketDataClient) GetMarket(ctx context.Context, marketID *big.Int) (tools.MarketInfo, error) {
	return tools.MarketInfo{
		ID:              marketID,
		Creator:         common.HexToAddress("0x1234567890123456789012345678901234567890"),
		AMM:             common.HexToAddress("0x2345678901234567890123456789012345678901"),
		CollateralToken: common.HexToAddress("0x3456789012345678901234567890123456789012"),
		CloseTime:       big.NewInt(time.Now().Add(24 * time.Hour).Unix()),
		Category:        "sports",
		MetadataURI:     "ipfs://QmTest123",
		CreatorStake:    big.NewInt(1000000000000000000),
		StakeRefunded:   false,
		Status:          0,
	}, nil
}

func (m *mockMarketDataClient) GetBalance(ctx context.Context) (*big.Int, error) {
	return big.NewInt(5000000000000000000), nil // 5 ETH
}

func (m *mockMarketDataClient) GetCurrentBlockTimestamp(ctx context.Context) (int64, error) {
	return time.Now().Unix(), nil
}

func main() {
	fmt.Println("=== AI Resolver Tool Integration Test ===\n")

	// Initialize tool registry
	registry := tools.NewRegistry()

	// Test 1: Register and verify Web Search Tool
	fmt.Println("Test 1: Web Search Tool")
	fmt.Println("------------------------")
	webSearchTool := tools.NewWebSearchTool()
	if err := registry.Register(webSearchTool); err != nil {
		log.Fatalf("Failed to register web search tool: %v", err)
	}
	fmt.Printf("✓ Registered: %s\n", webSearchTool.Name())
	fmt.Printf("  Type: %s\n", webSearchTool.Type())
	fmt.Printf("  Description: %s\n", webSearchTool.Description())

	// Test web search tool format
	format := webSearchTool.ToOpenAIFormat()
	formatJSON, _ := json.MarshalIndent(format, "  ", "  ")
	fmt.Printf("  OpenAI Format:\n  %s\n\n", string(formatJSON))

	// Test 2: Register and test Calculator Tool
	fmt.Println("Test 2: Calculator Tool")
	fmt.Println("------------------------")
	calculatorTool := tools.NewCalculatorTool()
	if err := registry.Register(calculatorTool); err != nil {
		log.Fatalf("Failed to register calculator tool: %v", err)
	}
	fmt.Printf("✓ Registered: %s\n", calculatorTool.Name())

	// Test calculator operations
	ctx := context.Background()

	// Test addition
	addInput := tools.ToolInput{
		CallID: "test_add",
		Arguments: map[string]any{
			"operation": "add",
			"values":    []any{5.0, 3.0},
		},
	}
	addOutput, err := calculatorTool.Execute(ctx, addInput)
	if err != nil {
		fmt.Printf("✗ Addition failed: %v\n", err)
	} else {
		fmt.Printf("✓ Addition: 5 + 3 = %v\n", addOutput.Data.(map[string]any)["result"])
	}

	// Test probability_multiply
	probInput := tools.ToolInput{
		CallID: "test_prob",
		Arguments: map[string]any{
			"operation": "probability_multiply",
			"values":    []any{0.8, 0.9},
		},
	}
	probOutput, err := calculatorTool.Execute(ctx, probInput)
	if err != nil {
		fmt.Printf("✗ Probability multiply failed: %v\n", err)
	} else {
		fmt.Printf("✓ Probability multiply: 0.8 * 0.9 = %v\n", probOutput.Data.(map[string]any)["result"])
	}

	// Test calculator format
	calcFormat := calculatorTool.ToOpenAIFormat()
	calcJSON, _ := json.MarshalIndent(calcFormat, "  ", "  ")
	fmt.Printf("  OpenAI Format:\n  %s\n\n", string(calcJSON))

	// Test 3: Register and test DateTime Tool
	fmt.Println("Test 3: DateTime Tool")
	fmt.Println("---------------------")
	datetimeTool := tools.NewDateTimeTool()
	if err := registry.Register(datetimeTool); err != nil {
		log.Fatalf("Failed to register datetime tool: %v", err)
	}
	fmt.Printf("✓ Registered: %s\n", datetimeTool.Name())

	// Test current timestamp
	nowInput := tools.ToolInput{
		CallID: "test_now",
		Arguments: map[string]any{
			"operation": "current_timestamp",
		},
	}
	nowOutput, err := datetimeTool.Execute(ctx, nowInput)
	if err != nil {
		fmt.Printf("✗ Current timestamp failed: %v\n", err)
	} else {
		fmt.Printf("✓ Current timestamp: %v\n", nowOutput.Data.(map[string]any)["timestamp"])
	}

	// Test time comparison
	// Parse dates first to get timestamps
	time1, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	time2, _ := time.Parse(time.RFC3339, "2024-12-31T23:59:59Z")
	compareInput := tools.ToolInput{
		CallID: "test_compare",
		Arguments: map[string]any{
			"operation":  "compare",
			"timestamp1": time1.Unix(),
			"timestamp2": time2.Unix(),
		},
	}
	compareOutput, err := datetimeTool.Execute(ctx, compareInput)
	if err != nil {
		fmt.Printf("✗ Time comparison failed: %v\n", err)
	} else {
		fmt.Printf("✓ Time comparison: 2024-01-01 vs 2024-12-31 = %v seconds difference\n", compareOutput.Data.(map[string]any)["difference_seconds"])
	}

	// Test datetime format
	dtFormat := datetimeTool.ToOpenAIFormat()
	dtJSON, _ := json.MarshalIndent(dtFormat, "  ", "  ")
	fmt.Printf("  OpenAI Format:\n  %s\n\n", string(dtJSON))

	// Test 4: Register and test Market Data Tool
	fmt.Println("Test 4: Market Data Tool")
	fmt.Println("------------------------")
	mockClient := &mockMarketDataClient{}
	marketDataTool := tools.NewMarketDataTool(mockClient)
	if err := registry.Register(marketDataTool); err != nil {
		log.Fatalf("Failed to register market data tool: %v", err)
	}
	fmt.Printf("✓ Registered: %s\n", marketDataTool.Name())

	// Test market data query
	marketInput := tools.ToolInput{
		CallID: "test_market",
		Arguments: map[string]any{
			"market_id": "123",
		},
	}
	marketOutput, err := marketDataTool.Execute(ctx, marketInput)
	if err != nil {
		fmt.Printf("✗ Market data query failed: %v\n", err)
	} else {
		data := marketOutput.Data.(map[string]any)
		fmt.Printf("✓ Market data retrieved:\n")
		fmt.Printf("  Market ID: %v\n", data["market_id"])
		fmt.Printf("  Category: %v\n", data["category"])
		fmt.Printf("  Close Time: %v\n", data["close_time"])
	}

	// Test market data format
	mdFormat := marketDataTool.ToOpenAIFormat()
	mdJSON, _ := json.MarshalIndent(mdFormat, "  ", "  ")
	fmt.Printf("  OpenAI Format:\n  %s\n\n", string(mdJSON))

	// Test 5: BSCScan Tool (if API key provided)
	fmt.Println("Test 5: BSCScan Tool")
	fmt.Println("--------------------")
	cfg, _ := config.LoadFromEnv()
	if cfg != nil && cfg.BSCScanAPIKey != "" {
		bscscanTool := tools.NewBSCScanTool(cfg.BSCScanAPIKey)
		if err := registry.Register(bscscanTool); err != nil {
			log.Fatalf("Failed to register bscscan tool: %v", err)
		}
		fmt.Printf("✓ Registered: %s\n", bscscanTool.Name())

		// Test BNB price query
		priceInput := tools.ToolInput{
			CallID: "test_price",
			Arguments: map[string]any{
				"action": "price",
			},
		}
		priceOutput, err := bscscanTool.Execute(ctx, priceInput)
		if err != nil {
			fmt.Printf("✗ BNB price query failed: %v\n", err)
		} else {
			data := priceOutput.Data.(map[string]any)
			fmt.Printf("✓ BNB Price: $%v USD\n", data["bnb_usd"])
		}
	} else {
		fmt.Println("⊘ Skipped: BSCSCAN_API_KEY not set")
	}
	fmt.Println()

	// Test 6: Registry Summary
	fmt.Println("Test 6: Registry Summary")
	fmt.Println("------------------------")
	allTools := registry.List()
	fmt.Printf("Total tools registered: %d\n", registry.Count())
	for i, tool := range allTools {
		fmt.Printf("%d. %s (%s)\n", i+1, tool.Name(), tool.Type())
	}
	fmt.Println()

	// Test 7: Tool Format for OpenAI API
	fmt.Println("Test 7: OpenAI API Tool Definitions")
	fmt.Println("-----------------------------------")
	fmt.Println("Tools array for OpenAI API:")
	toolsArray := registry.ToOpenAIFormat()
	toolsJSON, _ := json.MarshalIndent(toolsArray, "", "  ")
	fmt.Println(string(toolsJSON))
	fmt.Println()

	fmt.Println("=== All Tests Completed Successfully ===")
}
