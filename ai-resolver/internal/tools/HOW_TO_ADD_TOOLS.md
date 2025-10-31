# How to Add New Tools to the AI Resolver

This guide explains how to create and register new tools for the AI resolver system.

## Quick Start

### 1. Choose a Tool Type

- **ToolTypeFunction** - Structured JSON input with defined schema
- **ToolTypeCustom** - Raw string/text input
- **ToolTypeWebSearchPreview** - OpenAI's built-in web search (already provided)

### 2. Create Your Tool File

Create a new file in `internal/tools/` following the naming pattern `builtin_*.go`:

```go
package tools

import (
    "context"
    "fmt"
    "time"
)

// YourTool implements Tool interface for [describe purpose]
type YourTool struct {
    *BaseTool
    // Add any dependencies here
    client SomeClient
}

// NewYourTool creates a new instance of YourTool
func NewYourTool(client SomeClient) *YourTool {
    // Define schema for function tools
    schema := NewToolSchema()
    schema.AddProperty("param1", PropertyTypeString, "Description", true)
    schema.AddProperty("param2", PropertyTypeNumber, "Description", false)
    
    tool := &YourTool{
        BaseTool: NewBaseTool(
            "your_tool_name",           // Unique identifier
            "Description for the LLM",  // What this tool does
            ToolTypeFunction,            // or ToolTypeCustom
            schema,                      // or nil for custom tools
        ),
        client: client,
    }
    
    tool.SetExecutor(tool.execute)
    return tool
}

// execute implements the tool logic
func (t *YourTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // For function tools, parse arguments
    param1, ok := input.Arguments["param1"].(string)
    if !ok {
        return ToolOutput{}, fmt.Errorf("invalid param1")
    }
    
    // Do your work here
    result := doSomething(param1)
    
    // Return result
    return ToolOutput{
        Content: fmt.Sprintf("Result: %v", result),
        CallID:  input.CallID,
    }, nil
}
```

### 3. Add Tests

Create tests in the same file or in `builtin_test.go`:

```go
func TestYourTool(t *testing.T) {
    // Create mock dependencies if needed
    mockClient := &MockClient{}
    
    // Create tool
    tool := NewYourTool(mockClient)
    
    // Test execution
    ctx := context.Background()
    input := ToolInput{
        CallID: "test-call",
        Arguments: map[string]any{
            "param1": "test value",
        },
    }
    
    output, err := tool.Execute(ctx, input)
    assert.NoError(t, err)
    assert.Contains(t, output.Content, "expected result")
}
```

### 4. Register in Server

Update `cmd/server/main.go` to register your tool:

```go
// In main() function after creating toolRegistry:

// Create any adapters needed
yourAdapter := &YourAdapter{client: someClient}

// Create and register tool
yourTool := tools.NewYourTool(yourAdapter)
if err := toolRegistry.Register(yourTool); err != nil {
    log.Fatalf("Failed to register your tool: %v", err)
}
```

### 5. Test End-to-End

Build and run:
```bash
go test ./internal/tools -v
go build ./cmd/server
./server
```

Check logs for: `Registered N tools: [...your_tool_name...]`

## Examples

### Example 1: Simple Function Tool (No Dependencies)

```go
type DateParserTool struct {
    *BaseTool
}

func NewDateParserTool() *DateParserTool {
    schema := NewToolSchema()
    schema.AddProperty("date_string", PropertyTypeString, "Date string to parse", true)
    schema.AddProperty("format", PropertyTypeString, "Expected date format", false)
    
    tool := &DateParserTool{
        BaseTool: NewBaseTool(
            "parse_date",
            "Parse date strings into timestamps",
            ToolTypeFunction,
            schema,
        ),
    }
    
    tool.SetExecutor(tool.execute)
    return tool
}

func (t *DateParserTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    dateStr := input.Arguments["date_string"].(string)
    format := "2006-01-02" // default
    if f, ok := input.Arguments["format"].(string); ok {
        format = f
    }
    
    parsed, err := time.Parse(format, dateStr)
    if err != nil {
        return ToolOutput{}, fmt.Errorf("parse error: %w", err)
    }
    
    return ToolOutput{
        Content: fmt.Sprintf("Timestamp: %d", parsed.Unix()),
        CallID:  input.CallID,
    }, nil
}
```

Registration:
```go
dateParser := tools.NewDateParserTool()
toolRegistry.Register(dateParser)
```

### Example 2: Custom Tool (Raw Input)

```go
type TextAnalyzerTool struct {
    *BaseTool
}

func NewTextAnalyzerTool() *TextAnalyzerTool {
    tool := &TextAnalyzerTool{
        BaseTool: NewBaseTool(
            "analyze_text",
            "Analyze text for sentiment, keywords, etc.",
            ToolTypeCustom,
            nil, // No schema for custom tools
        ),
    }
    
    tool.SetExecutor(tool.execute)
    return tool
}

func (t *TextAnalyzerTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    text := input.RawInput // Custom tools use RawInput
    
    // Perform analysis
    sentiment := analyzeSentiment(text)
    keywords := extractKeywords(text)
    
    result := fmt.Sprintf("Sentiment: %s\nKeywords: %v", sentiment, keywords)
    
    return ToolOutput{
        Content: result,
        CallID:  input.CallID,
    }, nil
}
```

### Example 3: Tool with External Dependencies

```go
type TokenPriceTool struct {
    *BaseTool
    priceOracle PriceOracle // External dependency
}

func NewTokenPriceTool(oracle PriceOracle) *TokenPriceTool {
    schema := NewToolSchema()
    schema.AddProperty("token_address", PropertyTypeString, "Token contract address", true)
    schema.AddProperty("chain_id", PropertyTypeNumber, "Blockchain chain ID", false)
    
    tool := &TokenPriceTool{
        BaseTool: NewBaseTool(
            "get_token_price",
            "Get current token price from oracle",
            ToolTypeFunction,
            schema,
        ),
        priceOracle: oracle,
    }
    
    tool.SetExecutor(tool.execute)
    return tool
}

func (t *TokenPriceTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    address := input.Arguments["token_address"].(string)
    chainID := 1 // default
    if cid, ok := input.Arguments["chain_id"].(float64); ok {
        chainID = int(cid)
    }
    
    price, err := t.priceOracle.GetPrice(ctx, address, chainID)
    if err != nil {
        return ToolOutput{}, fmt.Errorf("price fetch failed: %w", err)
    }
    
    return ToolOutput{
        Content: fmt.Sprintf("Price: $%.2f", price),
        CallID:  input.CallID,
    }, nil
}
```

Registration with adapter:
```go
// Create adapter if needed
type priceOracleAdapter struct {
    client *adapter.Client
}

func (a *priceOracleAdapter) GetPrice(ctx context.Context, token string, chain int) (float64, error) {
    // Implementation
}

// Register
oracle := &priceOracleAdapter{client: client}
priceTool := tools.NewTokenPriceTool(oracle)
toolRegistry.Register(priceTool)
```

## Schema Definition Reference

### Property Types

```go
PropertyTypeString   // "hello"
PropertyTypeNumber   // 42, 3.14
PropertyTypeInteger  // 42 (no decimals)
PropertyTypeBoolean  // true/false
PropertyTypeArray    // [1, 2, 3]
PropertyTypeObject   // {"key": "value"}
```

### Adding Properties

```go
schema := NewToolSchema()

// Basic property
schema.AddProperty("name", PropertyTypeString, "Description", true) // required

// Optional property
schema.AddProperty("age", PropertyTypeNumber, "Optional age", false)

// Property with enum
prop := &ToolProperty{
    Type:        PropertyTypeString,
    Description: "Status",
    Enum:        []any{"pending", "active", "completed"},
}
schema.Properties["status"] = prop
schema.Required = append(schema.Required, "status")

// Array property
arrayProp := &ToolProperty{
    Type:        PropertyTypeArray,
    Description: "List of IDs",
    Items: &ToolProperty{
        Type: PropertyTypeNumber,
    },
}
schema.Properties["ids"] = arrayProp
```

## Best Practices

### 1. Tool Naming
- Use lowercase with underscores: `get_market_data`
- Be descriptive but concise
- Follow verb_noun pattern when possible

### 2. Descriptions
Write clear descriptions for the LLM:
```go
// ❌ Bad
"Gets data"

// ✅ Good
"Fetches on-chain market data including creator, AMM address, status, and metadata URI"
```

### 3. Error Handling
```go
func (t *MyTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // Validate input
    if input.Arguments["param"] == nil {
        return ToolOutput{}, NewValidationError("param", "required parameter missing", nil)
    }
    
    // Handle external errors
    result, err := t.client.DoSomething(ctx)
    if err != nil {
        return ToolOutput{}, fmt.Errorf("failed to execute: %w", err)
    }
    
    // Return success
    return ToolOutput{
        Content: result,
        CallID:  input.CallID,
    }, nil
}
```

### 4. Timeouts
Always respect context timeouts:
```go
func (t *MyTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // Check context before expensive operations
    select {
    case <-ctx.Done():
        return ToolOutput{}, ctx.Err()
    default:
    }
    
    // Use context in external calls
    result, err := t.client.FetchData(ctx)
    // ...
}
```

### 5. Testing
```go
func TestMyTool(t *testing.T) {
    tests := []struct{
        name      string
        input     ToolInput
        wantErr   bool
        wantMatch string
    }{
        {
            name: "valid input",
            input: ToolInput{
                CallID: "test",
                Arguments: map[string]any{"param": "value"},
            },
            wantErr: false,
            wantMatch: "expected result",
        },
        {
            name: "missing parameter",
            input: ToolInput{
                CallID: "test",
                Arguments: map[string]any{},
            },
            wantErr: true,
        },
    }
    
    tool := NewMyTool()
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output, err := tool.Execute(context.Background(), tt.input)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Contains(t, output.Content, tt.wantMatch)
            }
        })
    }
}
```

## Checklist

Before submitting your new tool:

- [ ] Tool file created in `internal/tools/`
- [ ] Appropriate tool type chosen
- [ ] Schema defined (for function tools)
- [ ] Execute function implemented
- [ ] Error handling added
- [ ] Context timeout respected
- [ ] Tests written and passing
- [ ] Server registration added
- [ ] Documentation updated
- [ ] Tested end-to-end

## Common Issues

### Issue: "Tool not found"
**Solution:** Make sure tool is registered in `main.go` and name matches exactly.

### Issue: "Invalid arguments"
**Solution:** Check schema definition matches what LLM is sending. Add logging to see incoming arguments.

### Issue: "Type assertion panic"
**Solution:** Use safe type assertions with comma-ok pattern:
```go
if val, ok := input.Arguments["key"].(string); ok {
    // use val
}
```

### Issue: "Context deadline exceeded"
**Solution:** Check for context cancellation and use it in all external calls.

## References

- Tool interface: `internal/tools/tool.go`
- Schema system: `internal/tools/schema.go`
- Registry: `internal/tools/registry.go`
- Examples: `internal/tools/builtin_*.go`
- Tests: `internal/tools/builtin_test.go`

## Support

For questions or issues:
1. Check existing tool implementations in `internal/tools/builtin_*.go`
2. Review test cases in `builtin_test.go`
3. Consult `TOOL_INTEGRATION_COMPLETE.md` for architecture overview
