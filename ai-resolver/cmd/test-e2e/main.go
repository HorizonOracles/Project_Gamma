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
	"github.com/project-gamma/ai-resolver/internal/llm"
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
	return big.NewInt(5000000000000000000), nil
}

func (m *mockMarketDataClient) GetCurrentBlockTimestamp(ctx context.Context) (int64, error) {
	return time.Now().Unix(), nil
}

// Simple adapter that implements llm.ToolRegistry
type simpleToolRegistry struct {
	tools map[string]llm.Tool
}

func newSimpleToolRegistry() *simpleToolRegistry {
	return &simpleToolRegistry{
		tools: make(map[string]llm.Tool),
	}
}

func (r *simpleToolRegistry) Register(tool llm.Tool) error {
	r.tools[tool.Name()] = tool
	return nil
}

func (r *simpleToolRegistry) Get(name string) (llm.Tool, bool) {
	tool, ok := r.tools[name]
	return tool, ok
}

func (r *simpleToolRegistry) List() []llm.Tool {
	result := make([]llm.Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		result = append(result, tool)
	}
	return result
}

// Simple tool adapter
type simpleTool struct {
	tool tools.Tool
}

func (t *simpleTool) Name() string {
	return t.tool.Name()
}

func (t *simpleTool) Description() string {
	return t.tool.Description()
}

func (t *simpleTool) ToOpenAIFormat() map[string]any {
	if bt, ok := t.tool.(*tools.BaseTool); ok {
		return bt.ToOpenAIFormat()
	}
	return map[string]any{
		"type": string(t.tool.Type()),
		"name": t.tool.Name(),
	}
}

func (t *simpleTool) Execute(ctx context.Context, arguments map[string]any) (map[string]any, error) {
	input := tools.ToolInput{
		CallID:    "test",
		Arguments: arguments,
	}
	output, err := t.tool.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	if dataMap, ok := output.Data.(map[string]any); ok {
		return dataMap, nil
	}
	return map[string]any{"result": output.Data}, nil
}

func main() {
	fmt.Println("=== AI Resolver End-to-End Test ===\n")

	// Load config
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if cfg.OpenAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY not set in environment")
	}

	// Create tool registry
	registry := newSimpleToolRegistry()

	// Register calculator tool
	calcTool := tools.NewCalculatorTool()
	registry.Register(&simpleTool{tool: calcTool})
	fmt.Println("✓ Registered calculator tool")

	// Register datetime tool
	datetimeTool := tools.NewDateTimeTool()
	registry.Register(&simpleTool{tool: datetimeTool})
	fmt.Println("✓ Registered datetime tool")

	// Register market data tool with mock client
	mockClient := &mockMarketDataClient{}
	marketDataTool := tools.NewMarketDataTool(mockClient)
	registry.Register(&simpleTool{tool: marketDataTool})
	fmt.Println("✓ Registered market data tool")

	// Register web search tool
	webSearchTool := tools.NewWebSearchTool()
	registry.Register(&simpleTool{tool: webSearchTool})
	fmt.Println("✓ Registered web search tool")

	fmt.Printf("\nTotal tools registered: %d\n\n", len(registry.List()))

	// Create LLM pipeline
	pipeline := llm.NewOpenAIPipeline(cfg.OpenAIAPIKey, "gpt-4o")
	pipeline.SetToolRegistry(registry)
	fmt.Println("✓ Created LLM pipeline with tool registry")

	// Test 1: Simple market question that requires calculation
	fmt.Println("\n=== Test 1: Market Question Requiring Calculation ===")
	market1 := llm.MarketInfo{
		MarketID:     1,
		Question:     "What is the probability that 0.7 * 0.8 equals approximately 0.56?",
		Description:  "Calculate the product of 0.7 and 0.8 to verify if it equals approximately 0.56",
		Category:     "math",
		CloseTime:    time.Now().Add(24 * time.Hour).Unix(),
		OutcomeCount: 2,
	}

	fmt.Println("\nQuestion:", market1.Question)
	fmt.Println("Analyzing market...")

	ctx := context.Background()
	decision1, err := pipeline.AnalyzeMarket(ctx, market1)
	if err != nil {
		fmt.Printf("✗ Analysis failed: %v\n", err)
	} else {
		fmt.Println("✓ Analysis completed")
		decisionJSON, _ := json.MarshalIndent(decision1, "", "  ")
		fmt.Printf("\nDecision:\n%s\n", string(decisionJSON))
	}

	// Test 2: Date-based question
	fmt.Println("\n=== Test 2: Market Question Requiring Date Comparison ===")
	market2 := llm.MarketInfo{
		MarketID:     2,
		Question:     "Is the current Unix timestamp greater than 1700000000?",
		Description:  "Check if the current time is after November 14, 2023",
		Category:     "time",
		CloseTime:    time.Now().Add(24 * time.Hour).Unix(),
		OutcomeCount: 2,
	}

	fmt.Println("\nQuestion:", market2.Question)
	fmt.Println("Analyzing market...")

	decision2, err := pipeline.AnalyzeMarket(ctx, market2)
	if err != nil {
		fmt.Printf("✗ Analysis failed: %v\n", err)
	} else {
		fmt.Println("✓ Analysis completed")
		decisionJSON, _ := json.MarshalIndent(decision2, "", "  ")
		fmt.Printf("\nDecision:\n%s\n", string(decisionJSON))
	}

	fmt.Println("\n=== Tests Completed ===")
	fmt.Println("\nNote: These tests verify that tools are properly integrated.")
	fmt.Println("Check the output above to confirm the LLM called the appropriate tools.")
}
