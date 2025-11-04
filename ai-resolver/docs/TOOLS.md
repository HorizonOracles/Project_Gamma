# AI Resolver Tool System

## Table of Contents

- [Overview](#overview)
- [Tool Types](#tool-types)
- [Built-in Tools](#built-in-tools)
- [Tool Orchestration](#tool-orchestration)
- [Creating Custom Tools](#creating-custom-tools)
- [Tool Registry](#tool-registry)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

## Overview

The AI Resolver Tool System enables the LLM to access external data sources, perform computations, and interact with blockchain systems. Tools extend the AI's capabilities beyond pure language understanding to include:

- Real-time blockchain data queries
- Mathematical calculations
- Date/time operations
- External API calls
- Custom business logic

### How It Works

```
┌─────────────────────────────────────────────────────────┐
│  User submits market question                            │
└────────────────────┬────────────────────────────────────┘
                     │
         ┌───────────▼──────────┐
         │  LLM receives:        │
         │  - Question           │
         │  - Tool definitions   │
         └───────────┬───────────┘
                     │
         ┌───────────▼──────────┐
         │  LLM decides:         │
         │  - Use tools? YES     │
         │  - Which tools?       │
         │  - With what params?  │
         └───────────┬───────────┘
                     │
         ┌───────────▼──────────┐
         │  Execute tools        │
         │  sequentially         │
         └───────────┬───────────┘
                     │
         ┌───────────▼──────────┐
         │  Append results       │
         │  to conversation      │
         └───────────┬───────────┘
                     │
         ┌───────────▼──────────┐
         │  LLM synthesizes      │
         │  final answer         │
         └───────────────────────┘
```

## Tool Types

### 1. Function Tools

**Description**: Structured tools with defined JSON schemas for input validation.

**Use Case**: When you need strict input validation and type checking.

**Schema Format**:
```json
{
  "type": "function",
  "function": {
    "name": "tool_name",
    "description": "What the tool does",
    "parameters": {
      "type": "object",
      "properties": {
        "param1": {
          "type": "string",
          "description": "Parameter description"
        }
      },
      "required": ["param1"]
    }
  }
}
```

**Example**: `get_market_data`, `calculate`, `datetime`

### 2. Custom Tools

**Description**: Free-form tools that accept raw text input.

**Use Case**: When input structure is flexible or unpredictable.

**Format**:
```json
{
  "type": "custom",
  "name": "tool_name",
  "description": "What the tool does"
}
```

**Example**: Text analysis, sentiment detection

### 3. Web Search Tools

**Description**: Native OpenAI web search integration.

**Use Case**: When the LLM needs to search the web for information.

**Format**:
```json
{
  "type": "web_search_preview"
}
```

**Note**: This is delegated to OpenAI's native search capability.

## Built-in Tools

### 1. get_market_data

**Purpose**: Fetch on-chain market information.

**Type**: Function Tool

**Parameters**:
- `market_id` (string, required): Market identifier

**Returns**:
```json
{
  "market_id": "123",
  "creator": "0x742d35...",
  "amm": "0x123abc...",
  "collateral_token": "0x456def...",
  "close_time": 1762172000,
  "category": "sports",
  "metadata_uri": "ipfs://...",
  "creator_stake": "1000000000000000000000",
  "status": 0
}
```

**Example Usage**:
```
Question: "What is the close time of market 1?"

LLM calls: get_market_data(market_id: "1")
Response: {"close_time": 1762172000, ...}
LLM answers: "Market 1 closes at timestamp 1762172000"
```

**Error Handling**:
- Market not found → Returns error message
- Invalid market_id → Validation error
- RPC failure → Returns descriptive error

---

### 2. calculate

**Purpose**: Perform mathematical operations.

**Type**: Function Tool

**Parameters**:
- `operation` (string, required): Operation name
- `values` (array, required): Operands

**Supported Operations**:
- `add`: Addition (a + b + c + ...)
- `subtract`: Subtraction (a - b)
- `multiply`: Multiplication (a × b × c × ...)
- `divide`: Division (a ÷ b)
- `modulo`: Modulo (a % b)
- `power`: Exponentiation (a ^ b)
- `sqrt`: Square root (√a)
- `abs`: Absolute value (|a|)
- `min`: Minimum value
- `max`: Maximum value
- `equals`: Equality check (a == b)
- `greater_than`: Greater than (a > b)
- `less_than`: Less than (a < b)

**Returns**:
```json
{
  "result": 42,
  "operation": "add",
  "values": [10, 20, 12]
}
```

**Example Usage**:
```
Question: "If market 1 closes at 1762172000 and market 2 closes 86400 seconds later, when does market 2 close?"

LLM calls:
1. get_market_data(market_id: "1")
   → {close_time: 1762172000}
2. calculate(operation: "add", values: [1762172000, 86400])
   → {result: 1762258400}
LLM answers: "Market 2 closes at 1762258400"
```

**Error Handling**:
- Division by zero → Returns error
- Invalid operation → Validation error
- Non-numeric values → Type error

---

### 3. datetime

**Purpose**: Convert Unix timestamps to human-readable dates.

**Type**: Function Tool

**Parameters**:
- `timestamp` (number, required): Unix timestamp in seconds

**Returns**:
```json
{
  "timestamp": 1762172000,
  "datetime": "2025-11-03 12:13:20",
  "date": "2025-11-03",
  "time": "12:13:20",
  "year": 2025,
  "month": 11,
  "day": 3,
  "hour": 12,
  "minute": 13,
  "second": 20
}
```

**Example Usage**:
```
Question: "Convert timestamp 1762172000 to a date"

LLM calls: datetime(timestamp: 1762172000)
Response: {datetime: "2025-11-03 12:13:20", ...}
LLM answers: "The timestamp 1762172000 corresponds to November 3, 2025 at 12:13:20 UTC"
```

**Error Handling**:
- Negative timestamp → Returns error
- Invalid format → Validation error

---

### 4. bscscan

**Purpose**: Query BSCScan API for blockchain data.

**Type**: Function Tool

**Parameters**:
- `module` (string, required): API module (account, contract, transaction, etc.)
- `action` (string, required): API action (balance, txlist, etc.)
- `params` (object, required): Additional parameters

**Supported Queries**:
- Account balance
- Transaction history
- Token balances
- Contract ABI
- Block information

**Returns**: Raw BSCScan API response

**Example Usage**:
```
Question: "What is the BNB balance of address 0x742d35...?"

LLM calls: bscscan(
  module: "account",
  action: "balance",
  params: {address: "0x742d35..."}
)
Response: {result: "1234567890000000000"}
LLM answers: "The address has 1.234567890 BNB"
```

**Error Handling**:
- API key missing → Returns error
- Rate limit exceeded → Returns error message
- Invalid address → BSCScan error

---

### 5. pancakeswap

**Purpose**: Fetch DEX data from PancakeSwap.

**Type**: Function Tool

**Parameters**:
- `pair` (string, required): Trading pair address
- `action` (string, required): Data to fetch (price, reserves, volume)

**Returns**:
```json
{
  "pair": "0x123...",
  "token0": "0xabc...",
  "token1": "0xdef...",
  "reserve0": "1000000000000000000000",
  "reserve1": "2000000000000000000000",
  "price": "2.0"
}
```

**Example Usage**:
```
Question: "What is the current price of CAKE/BNB?"

LLM calls: pancakeswap(
  pair: "0x0eD7e52944161450477ee417DE9Cd3a859b14fD0",
  action: "price"
)
Response: {price: "0.0125"}
LLM answers: "The current CAKE/BNB price is 0.0125"
```

## Tool Orchestration

### Multi-Tool Chaining

The LLM can chain multiple tools together to answer complex questions.

**Example 1: Temporal Calculation**
```
Question: "Is market 1's close time exactly 1 day after its creation?"

Execution Flow:
1. get_market_data(market_id: "1")
   → {created_at: 1762000000, close_time: 1762086400}

2. calculate(operation: "subtract", values: [1762086400, 1762000000])
   → {result: 86400}

3. calculate(operation: "equals", values: [86400, 86400])
   → {result: true}

4. LLM synthesizes: "YES, market 1 closes exactly 1 day (86400 seconds) after creation"
```

**Example 2: Price Comparison**
```
Question: "Compare CAKE price from 2 different sources"

Execution Flow:
1. pancakeswap(pair: "CAKE/BNB", action: "price")
   → {price: "0.0125"}

2. bscscan(module: "stats", action: "tokenprice", params: {contractaddress: "CAKE"})
   → {price: "0.0124"}

3. calculate(operation: "subtract", values: [0.0125, 0.0124])
   → {result: 0.0001}

4. LLM synthesizes: "CAKE price is 0.0125 on PancakeSwap and 0.0124 on BSCScan, a difference of 0.0001 (0.8%)"
```

### Iteration Limits

**Maximum Iterations**: 10

**Why**: Prevents infinite loops and excessive API costs.

**What Happens After 10 Iterations**:
- System returns partial results
- LLM makes best-effort decision
- Warning logged for monitoring

### Stateless Context

The system uses **prompt appending** instead of native conversation history:

```
Iteration 1:
System: "Question: [X]. Tools: [list]"
LLM: tool_calls = [T1]
System: Executes T1

Iteration 2:
System: "Question: [X]. Tools: [list]
         Previous: T1(args) → result"
LLM: tool_calls = [T2]
System: Executes T2

Iteration 3:
System: "Question: [X]. Tools: [list]
         Previous: T1(args) → result
         Previous: T2(args) → result"
LLM: "Based on results, answer is [Y]"
```

## Creating Custom Tools

See [`internal/tools/HOW_TO_ADD_TOOLS.md`](../internal/tools/HOW_TO_ADD_TOOLS.md) for detailed guide.

### Quick Start

**1. Create Tool File** (`internal/tools/builtin_mytool.go`):
```go
package tools

import "context"

type MyTool struct {
    *BaseTool
}

func NewMyTool() *MyTool {
    schema := NewToolSchema()
    schema.AddProperty("param1", PropertyTypeString, "Description", true)
    
    tool := &MyTool{
        BaseTool: NewBaseTool(
            "my_tool",
            "Tool description for LLM",
            ToolTypeFunction,
            schema,
        ),
    }
    
    tool.SetExecutor(tool.execute)
    return tool
}

func (t *MyTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    param1 := input.Arguments["param1"].(string)
    
    // Do work here
    result := processData(param1)
    
    return ToolOutput{
        Content: result,
        CallID:  input.CallID,
    }, nil
}
```

**2. Register Tool** (`cmd/server/main.go`):
```go
myTool := tools.NewMyTool()
if err := toolRegistry.Register(myTool); err != nil {
    log.Fatalf("Failed to register tool: %v", err)
}
```

**3. Test**:
```bash
go test ./internal/tools -v -run TestMyTool
```

## Tool Registry

### Registry Operations

**Registration**:
```go
registry := tools.NewRegistry()
tool := tools.NewMyTool()
err := registry.Register(tool)
```

**Lookup**:
```go
tool, exists := registry.Get("my_tool")
if !exists {
    // Tool not found
}
```

**List All**:
```go
toolList := registry.List()
for _, tool := range toolList {
    fmt.Println(tool.Name())
}
```

**Execute**:
```go
output, err := registry.Execute(ctx, "my_tool", input)
```

### Registry Validation

The registry validates tools during registration:

✅ **Checks**:
- Tool name is non-empty
- Tool name is unique
- Tool implements Tool interface
- Schema is valid (for function tools)

❌ **Rejects**:
- Duplicate tool names
- Empty descriptions
- Invalid schemas
- Nil executors

## Best Practices

### 1. Tool Design

**Do**:
- Keep tools focused (single responsibility)
- Use descriptive names (verb_noun pattern)
- Provide clear descriptions for the LLM
- Validate all inputs
- Return structured data when possible

**Don't**:
- Create overly complex tools
- Use ambiguous names
- Assume valid input
- Return raw error messages to LLM

### 2. Error Handling

**Do**:
```go
func (t *MyTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // Validate input
    if input.Arguments["param"] == nil {
        return ToolOutput{}, fmt.Errorf("param is required")
    }
    
    // Handle external errors gracefully
    result, err := externalAPI.Call(ctx)
    if err != nil {
        return ToolOutput{
            Content: fmt.Sprintf("Failed to fetch data: %v", err),
            Error:   err,
        }, nil // Return error in output, not as function error
    }
    
    return ToolOutput{Content: result}, nil
}
```

**Don't**:
```go
// Bad: Unvalidated input
param := input.Arguments["param"].(string) // May panic!

// Bad: Raw error return
return ToolOutput{}, err // LLM can't interpret

// Bad: Exposing sensitive info
return ToolOutput{Content: err.Error()}, nil // May leak secrets
```

### 3. Performance

**Timeouts**:
```go
func (t *MyTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // Respect context timeout
    select {
    case <-ctx.Done():
        return ToolOutput{}, ctx.Err()
    default:
    }
    
    // Pass context to external calls
    result, err := client.FetchWithContext(ctx, data)
    return ToolOutput{Content: result}, err
}
```

**Caching**:
```go
type CachedTool struct {
    *BaseTool
    cache map[string]string
    mu    sync.RWMutex
}

func (t *CachedTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    key := input.Arguments["key"].(string)
    
    // Check cache
    t.mu.RLock()
    if cached, ok := t.cache[key]; ok {
        t.mu.RUnlock()
        return ToolOutput{Content: cached}, nil
    }
    t.mu.RUnlock()
    
    // Fetch and cache
    result := fetchData(key)
    t.mu.Lock()
    t.cache[key] = result
    t.mu.Unlock()
    
    return ToolOutput{Content: result}, nil
}
```

### 4. Testing

**Comprehensive Tests**:
```go
func TestMyTool(t *testing.T) {
    tests := []struct {
        name      string
        input     ToolInput
        wantErr   bool
        wantMatch string
    }{
        {
            name: "valid input",
            input: ToolInput{
                CallID: "test-1",
                Arguments: map[string]any{"param": "value"},
            },
            wantMatch: "expected result",
        },
        {
            name: "missing param",
            input: ToolInput{
                CallID: "test-2",
                Arguments: map[string]any{},
            },
            wantErr: true,
        },
        {
            name: "invalid param type",
            input: ToolInput{
                CallID: "test-3",
                Arguments: map[string]any{"param": 123},
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

## Troubleshooting

### Common Issues

**Issue**: Tool not being called by LLM
**Solutions**:
- Check tool description is clear
- Verify tool is registered
- Ensure schema matches LLM expectations
- Add examples in tool description

**Issue**: Tool arguments validation failing
**Solutions**:
- Log incoming arguments for debugging
- Use safe type assertions (comma-ok pattern)
- Check schema property types match LLM output
- Handle optional parameters gracefully

**Issue**: Tool execution timeouts
**Solutions**:
- Respect context timeouts
- Implement early cancellation checks
- Optimize slow operations
- Cache frequently accessed data

**Issue**: Tool errors not reaching LLM
**Solutions**:
- Return errors in ToolOutput.Content, not as function error
- Format errors in user-friendly way
- Don't include stack traces
- Provide actionable error messages

### Debugging

**Enable Debug Logging**:
```bash
LOG_LEVEL=debug ./bin/ai-resolver
```

**Debug Output Example**:
```
DEBUG: Tool call received: get_market_data
DEBUG: Arguments: {"market_id": "123"}
DEBUG: Executing tool...
DEBUG: Tool result: {"close_time": 1762172000, ...}
DEBUG: Appending to context: "Previous call: get_market_data(market_id: '123') returned ..."
```

**Inspect Tool Registry**:
```go
// In main.go
log.Printf("Registered tools: %v", toolRegistry.List())
```

### Performance Monitoring

**Track Tool Usage**:
```go
type MetricsTool struct {
    *BaseTool
    execCount int64
    execTime  time.Duration
}

func (t *MetricsTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    start := time.Now()
    defer func() {
        atomic.AddInt64(&t.execCount, 1)
        t.execTime += time.Since(start)
    }()
    
    // Execute tool logic
    return t.actualExecute(ctx, input)
}
```

## Advanced Topics

### Tool Dependencies

Some tools may depend on others:

```go
type CompositeTool struct {
    *BaseTool
    registry *Registry
}

func (t *CompositeTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    // Use another tool
    output1, _ := t.registry.Execute(ctx, "tool1", ToolInput{...})
    
    // Process output1 and call another tool
    output2, _ := t.registry.Execute(ctx, "tool2", ToolInput{...})
    
    // Combine results
    return ToolOutput{Content: combine(output1, output2)}, nil
}
```

### Conditional Tool Execution

```go
func (t *SmartTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
    mode := input.Arguments["mode"].(string)
    
    switch mode {
    case "fast":
        return t.fastPath(ctx, input)
    case "accurate":
        return t.accuratePath(ctx, input)
    default:
        return ToolOutput{}, fmt.Errorf("invalid mode: %s", mode)
    }
}
```

### Tool Versioning

```go
func NewMyToolV2() *MyTool {
    tool := &MyTool{
        BaseTool: NewBaseTool(
            "my_tool_v2",  // Version in name
            "Enhanced version with new features",
            ToolTypeFunction,
            schemaV2,
        ),
        version: 2,
    }
    
    tool.SetExecutor(tool.executeV2)
    return tool
}
```

## Future Enhancements

### Planned Features

**Q1 2025**:
- [ ] Tool result caching system
- [ ] Parallel tool execution
- [ ] Tool execution middleware
- [ ] Built-in retry logic

**Q2 2025**:
- [ ] Dynamic tool loading from plugins
- [ ] Tool permission system
- [ ] Rate limiting per tool
- [ ] Tool analytics dashboard

**Q3+ 2025**:
- [ ] Tool marketplace
- [ ] Community tool repository
- [ ] Tool composition framework
- [ ] AI-generated tools

## Resources

- **Tool Interface**: [`internal/tools/tool.go`](../internal/tools/tool.go)
- **Schema System**: [`internal/tools/schema.go`](../internal/tools/schema.go)
- **Registry**: [`internal/tools/registry.go`](../internal/tools/registry.go)
- **Examples**: [`internal/tools/builtin_*.go`](../internal/tools/)
- **How-to Guide**: [`internal/tools/HOW_TO_ADD_TOOLS.md`](../internal/tools/HOW_TO_ADD_TOOLS.md)
- **Architecture**: [`docs/ARCHITECTURE.md`](./ARCHITECTURE.md)
