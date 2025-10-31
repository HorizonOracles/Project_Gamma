# Tool Registry Base Class Implementation Plan

## Overview

This document outlines the design and implementation plan for a flexible tool registry system that enables extending the AI resolver with custom tools. The system follows OpenAI's tool calling patterns and provides a base class architecture for registering, managing, and executing both function-based and custom tools.

## Architecture Goals

1. **Extensibility**: Easy registration of new tools without modifying core code
2. **Type Safety**: Strong typing for tool definitions and parameters
3. **Error Handling**: Robust error handling and validation
4. **Observability**: Logging and monitoring of tool execution
5. **Testability**: Easy to mock and test individual tools
6. **OpenAI Integration**: Native support for OpenAI's tool calling API patterns

## Core Components

### 1. Tool Interface

```go
// Tool represents a callable tool that can be used by the LLM
type Tool interface {
    // Name returns the unique identifier for this tool
    Name() string
    
    // Description returns a human-readable description for the LLM
    Description() string
    
    // Type returns the tool type (function, custom, web_search_preview, etc.)
    Type() ToolType
    
    // Schema returns the JSON schema for tool parameters (nil for custom tools)
    Schema() *ToolSchema
    
    // Execute runs the tool with the provided input
    Execute(ctx context.Context, input ToolInput) (ToolOutput, error)
    
    // Validate checks if the input is valid for this tool
    Validate(input ToolInput) error
}
```

### 2. Tool Types

```go
type ToolType string

const (
    ToolTypeFunction         ToolType = "function"
    ToolTypeCustom           ToolType = "custom"
    ToolTypeWebSearchPreview ToolType = "web_search_preview"
)
```

### 3. Tool Registry

```go
// Registry manages the collection of available tools
type Registry interface {
    // Register adds a new tool to the registry
    Register(tool Tool) error
    
    // Unregister removes a tool from the registry
    Unregister(name string) error
    
    // Get retrieves a tool by name
    Get(name string) (Tool, bool)
    
    // List returns all registered tools
    List() []Tool
    
    // ListByType returns tools of a specific type
    ListByType(toolType ToolType) []Tool
    
    // ExecuteTool runs a tool by name with the provided input
    ExecuteTool(ctx context.Context, name string, input ToolInput) (ToolOutput, error)
    
    // ToOpenAIFormat converts registered tools to OpenAI API format
    ToOpenAIFormat() []map[string]any
}
```

### 4. Base Tool Implementation

```go
// BaseTool provides common functionality for all tools
type BaseTool struct {
    name        string
    description string
    toolType    ToolType
    schema      *ToolSchema
    executor    ToolExecutor
    validator   ToolValidator
    middleware  []ToolMiddleware
}

// ToolExecutor is the function signature for tool execution logic
type ToolExecutor func(ctx context.Context, input ToolInput) (ToolOutput, error)

// ToolValidator validates tool input
type ToolValidator func(input ToolInput) error

// ToolMiddleware wraps tool execution with additional functionality
type ToolMiddleware func(next ToolExecutor) ToolExecutor
```

### 5. Tool Schema Definition

```go
// ToolSchema defines the JSON schema for tool parameters
type ToolSchema struct {
    Type       string                 `json:"type"`
    Properties map[string]Property    `json:"properties"`
    Required   []string               `json:"required"`
}

type Property struct {
    Type        string   `json:"type"`
    Description string   `json:"description"`
    Enum        []string `json:"enum,omitempty"`
    Items       *Property `json:"items,omitempty"`
}
```

### 6. Tool Input/Output

```go
// ToolInput represents input to a tool
type ToolInput struct {
    // For function tools: JSON object with parameters
    Arguments map[string]any
    
    // For custom tools: raw string input
    RawInput string
    
    // Metadata
    CallID    string
    Timestamp int64
}

// ToolOutput represents output from a tool
type ToolOutput struct {
    // Result data (will be JSON encoded for function tools)
    Data any
    
    // Error if execution failed
    Error error
    
    // Metadata
    CallID       string
    ExecutionTime time.Duration
    Logs         []string
}
```

## Implementation Phases

### Phase 1: Core Infrastructure (Week 1)

#### 1.1 Create Base Package Structure
```
ai-resolver/internal/tools/
├── registry.go          # Registry implementation
├── tool.go             # Tool interface and base implementation
├── types.go            # Common types (ToolInput, ToolOutput, etc.)
├── schema.go           # Schema definitions and validation
├── middleware.go       # Middleware support
├── errors.go           # Tool-specific errors
└── registry_test.go    # Unit tests
```

#### 1.2 Implement Core Types
- Define `Tool` interface
- Implement `BaseTool` struct
- Create `ToolSchema` with JSON schema support
- Implement `ToolInput` and `ToolOutput` structures

#### 1.3 Implement Registry
- Create thread-safe registry with `sync.RWMutex`
- Implement registration/unregistration
- Add tool lookup by name and type
- Implement OpenAI format conversion

#### 1.4 Add Validation
- JSON schema validation for function tools
- Input type checking
- Parameter requirement validation

### Phase 2: Built-in Tools (Week 2)

#### 2.1 Web Search Tool
```go
// WebSearchTool wraps the existing web search functionality
type WebSearchTool struct {
    *BaseTool
    client *http.Client
    apiKey string
}

// NewWebSearchTool creates a web search tool
func NewWebSearchTool() Tool {
    return &WebSearchTool{
        BaseTool: NewBaseTool(
            "web_search",
            "Search the web for current information",
            ToolTypeWebSearchPreview,
            nil, // No schema for web_search_preview
        ),
    }
}
```

#### 2.2 Market Data Tool
```go
// MarketDataTool fetches market information from the blockchain
type MarketDataTool struct {
    *BaseTool
    client *adapter.Client
}

func NewMarketDataTool(client *adapter.Client) Tool {
    schema := &ToolSchema{
        Type: "object",
        Properties: map[string]Property{
            "marketId": {
                Type:        "number",
                Description: "The ID of the market to fetch",
            },
        },
        Required: []string{"marketId"},
    }
    
    return &MarketDataTool{
        BaseTool: NewBaseTool(
            "get_market_data",
            "Fetch market details from the blockchain",
            ToolTypeFunction,
            schema,
        ),
        client: client,
    }
}
```

#### 2.3 Code Execution Tool (Custom)
```go
// CodeExecTool executes arbitrary code (with safety constraints)
type CodeExecTool struct {
    *BaseTool
    timeout time.Duration
}

func NewCodeExecTool() Tool {
    return &CodeExecTool{
        BaseTool: NewBaseTool(
            "code_exec",
            "Executes arbitrary Python code for calculations and analysis",
            ToolTypeCustom,
            nil, // Custom tools don't use JSON schema
        ),
        timeout: 5 * time.Second,
    }
}
```

### Phase 3: LLM Integration (Week 2)

#### 3.1 Update Pipeline Interface
```go
// Update Pipeline interface to accept tool registry
type Pipeline interface {
    // SetToolRegistry configures available tools
    SetToolRegistry(registry Registry)
    
    // AnalyzeMarket performs analysis with tool support
    AnalyzeMarket(ctx context.Context, market MarketInfo) (*Decision, error)
}
```

#### 3.2 Modify OpenAI Pipeline
- Add tool registry field to `OpenAIPipeline`
- Update request building to include tools from registry
- Implement tool call handling and execution loop
- Add tool output formatting for subsequent requests

#### 3.3 Tool Execution Flow
```go
func (p *OpenAIPipeline) executeWithTools(ctx context.Context, prompt string) (string, error) {
    inputList := []map[string]any{
        {"role": "user", "content": prompt},
    }
    
    for {
        // Make request with tools
        response, err := p.callOpenAI(ctx, inputList)
        if err != nil {
            return "", err
        }
        
        // Check for tool calls
        toolCalls := extractToolCalls(response)
        if len(toolCalls) == 0 {
            // No more tool calls, return final response
            return extractTextResponse(response), nil
        }
        
        // Execute tool calls
        for _, call := range toolCalls {
            output, err := p.registry.ExecuteTool(ctx, call.Name, call.Input)
            if err != nil {
                // Handle error
                continue
            }
            
            // Add tool output to input list
            inputList = append(inputList, map[string]any{
                "type":    "function_call_output",
                "call_id": call.CallID,
                "output":  output.Data,
            })
        }
    }
}
```

### Phase 4: Middleware & Extensions (Week 3)

#### 4.1 Implement Middleware System
```go
// Logging middleware
func LoggingMiddleware(logger *log.Logger) ToolMiddleware {
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            start := time.Now()
            logger.Printf("Executing tool with input: %+v", input)
            
            output, err := next(ctx, input)
            
            logger.Printf("Tool completed in %v (error: %v)", time.Since(start), err)
            return output, err
        }
    }
}

// Rate limiting middleware
func RateLimitMiddleware(limiter *rate.Limiter) ToolMiddleware {
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            if err := limiter.Wait(ctx); err != nil {
                return ToolOutput{}, fmt.Errorf("rate limit: %w", err)
            }
            return next(ctx, input)
        }
    }
}

// Caching middleware
func CachingMiddleware(cache Cache) ToolMiddleware {
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            key := generateCacheKey(input)
            
            if cached, found := cache.Get(key); found {
                return cached.(ToolOutput), nil
            }
            
            output, err := next(ctx, input)
            if err == nil {
                cache.Set(key, output, 5*time.Minute)
            }
            
            return output, err
        }
    }
}
```

#### 4.2 Add Metrics & Observability
```go
// MetricsMiddleware tracks tool execution metrics
func MetricsMiddleware(metrics *prometheus.Registry) ToolMiddleware {
    executionCount := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "tool_executions_total",
            Help: "Total number of tool executions",
        },
        []string{"tool_name", "status"},
    )
    
    executionDuration := prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "tool_execution_duration_seconds",
            Help: "Tool execution duration in seconds",
        },
        []string{"tool_name"},
    )
    
    metrics.MustRegister(executionCount, executionDuration)
    
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            start := time.Now()
            output, err := next(ctx, input)
            duration := time.Since(start).Seconds()
            
            toolName := input.CallID // Or extract from context
            status := "success"
            if err != nil {
                status = "error"
            }
            
            executionCount.WithLabelValues(toolName, status).Inc()
            executionDuration.WithLabelValues(toolName).Observe(duration)
            
            return output, err
        }
    }
}
```

### Phase 5: Advanced Features (Week 4)

#### 5.1 Tool Composition
```go
// CompositeTool combines multiple tools
type CompositeTool struct {
    *BaseTool
    tools []Tool
    flow  ExecutionFlow
}

// ExecutionFlow defines how tools are composed
type ExecutionFlow interface {
    Execute(ctx context.Context, tools []Tool, input ToolInput) (ToolOutput, error)
}

// SequentialFlow executes tools in sequence
type SequentialFlow struct{}

func (f *SequentialFlow) Execute(ctx context.Context, tools []Tool, input ToolInput) (ToolOutput, error) {
    currentInput := input
    var finalOutput ToolOutput
    
    for _, tool := range tools {
        output, err := tool.Execute(ctx, currentInput)
        if err != nil {
            return ToolOutput{}, err
        }
        
        // Pass output as input to next tool
        currentInput = ToolInput{
            Arguments: map[string]any{"data": output.Data},
        }
        finalOutput = output
    }
    
    return finalOutput, nil
}
```

#### 5.2 Dynamic Tool Loading
```go
// ToolLoader loads tools from external sources
type ToolLoader interface {
    Load(source string) ([]Tool, error)
}

// PluginLoader loads Go plugins
type PluginLoader struct {
    pluginDir string
}

func (l *PluginLoader) Load(source string) ([]Tool, error) {
    // Load Go plugin
    p, err := plugin.Open(filepath.Join(l.pluginDir, source))
    if err != nil {
        return nil, err
    }
    
    // Look for NewTool symbol
    symbol, err := p.Lookup("NewTool")
    if err != nil {
        return nil, err
    }
    
    // Cast and execute
    newTool := symbol.(func() Tool)
    return []Tool{newTool()}, nil
}
```

#### 5.3 Tool Authorization & Security
```go
// AuthorizationMiddleware enforces tool access control
func AuthorizationMiddleware(authorizer Authorizer) ToolMiddleware {
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            // Extract user/role from context
            user := ctx.Value("user")
            
            if !authorizer.CanExecute(user, input) {
                return ToolOutput{}, ErrUnauthorized
            }
            
            return next(ctx, input)
        }
    }
}

// SandboxMiddleware runs tools in isolated environment
func SandboxMiddleware(sandbox Sandbox) ToolMiddleware {
    return func(next ToolExecutor) ToolExecutor {
        return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
            return sandbox.Execute(ctx, func() (ToolOutput, error) {
                return next(ctx, input)
            })
        }
    }
}
```

## Usage Examples

### Example 1: Basic Tool Registration

```go
// Initialize registry
registry := tools.NewRegistry()

// Register built-in tools
registry.Register(tools.NewWebSearchTool())
registry.Register(tools.NewMarketDataTool(client))

// Register custom tool
registry.Register(tools.NewBaseTool(
    "calculate",
    "Performs mathematical calculations",
    tools.ToolTypeFunction,
    &tools.ToolSchema{
        Type: "object",
        Properties: map[string]tools.Property{
            "expression": {
                Type:        "string",
                Description: "Mathematical expression to evaluate",
            },
        },
        Required: []string{"expression"},
    },
).WithExecutor(func(ctx context.Context, input tools.ToolInput) (tools.ToolOutput, error) {
    expr := input.Arguments["expression"].(string)
    result := evaluateExpression(expr)
    return tools.ToolOutput{Data: result}, nil
}))

// Initialize pipeline with registry
pipeline := llm.NewOpenAIPipeline(apiKey, model)
pipeline.SetToolRegistry(registry)
```

### Example 2: Creating a Custom Tool

```go
// Define custom tool
type PriceOracleTool struct {
    *tools.BaseTool
    httpClient *http.Client
}

func NewPriceOracleTool() tools.Tool {
    schema := &tools.ToolSchema{
        Type: "object",
        Properties: map[string]tools.Property{
            "asset": {
                Type:        "string",
                Description: "Asset symbol (e.g., BTC, ETH)",
            },
            "currency": {
                Type:        "string",
                Description: "Quote currency (e.g., USD)",
            },
        },
        Required: []string{"asset", "currency"},
    }
    
    tool := &PriceOracleTool{
        BaseTool: tools.NewBaseTool(
            "get_asset_price",
            "Get current price of an asset",
            tools.ToolTypeFunction,
            schema,
        ),
        httpClient: &http.Client{Timeout: 10 * time.Second},
    }
    
    // Set executor
    tool.SetExecutor(tool.execute)
    
    return tool
}

func (t *PriceOracleTool) execute(ctx context.Context, input tools.ToolInput) (tools.ToolOutput, error) {
    asset := input.Arguments["asset"].(string)
    currency := input.Arguments["currency"].(string)
    
    // Fetch price from external API
    price, err := t.fetchPrice(ctx, asset, currency)
    if err != nil {
        return tools.ToolOutput{}, err
    }
    
    return tools.ToolOutput{
        Data: map[string]any{
            "asset":    asset,
            "currency": currency,
            "price":    price,
            "timestamp": time.Now().Unix(),
        },
    }, nil
}
```

### Example 3: Using Middleware

```go
// Create tool with middleware
tool := tools.NewMarketDataTool(client)

// Add middleware layers
tool = tool.
    WithMiddleware(tools.LoggingMiddleware(logger)).
    WithMiddleware(tools.RateLimitMiddleware(rateLimiter)).
    WithMiddleware(tools.CachingMiddleware(cache)).
    WithMiddleware(tools.MetricsMiddleware(metrics))

// Register with middleware applied
registry.Register(tool)
```

## Testing Strategy

### Unit Tests
```go
func TestToolRegistry(t *testing.T) {
    t.Run("Register and retrieve tool", func(t *testing.T) {
        registry := tools.NewRegistry()
        tool := tools.NewMockTool("test", "Test tool")
        
        err := registry.Register(tool)
        assert.NoError(t, err)
        
        retrieved, found := registry.Get("test")
        assert.True(t, found)
        assert.Equal(t, tool.Name(), retrieved.Name())
    })
    
    t.Run("Duplicate registration fails", func(t *testing.T) {
        registry := tools.NewRegistry()
        tool := tools.NewMockTool("test", "Test tool")
        
        registry.Register(tool)
        err := registry.Register(tool)
        assert.Error(t, err)
    })
}
```

### Integration Tests
```go
func TestToolExecution(t *testing.T) {
    // Create test context
    ctx := context.Background()
    
    // Create registry with real tools
    registry := tools.NewRegistry()
    registry.Register(tools.NewWebSearchTool())
    
    // Execute tool
    output, err := registry.ExecuteTool(ctx, "web_search", tools.ToolInput{
        RawInput: "current weather in San Francisco",
    })
    
    assert.NoError(t, err)
    assert.NotNil(t, output.Data)
}
```

## Configuration

### Registry Configuration
```yaml
tools:
  enabled:
    - web_search
    - market_data
    - code_exec
    - price_oracle
  
  web_search:
    rate_limit: 10/minute
    timeout: 30s
  
  market_data:
    cache_ttl: 5m
    rate_limit: 100/minute
  
  code_exec:
    enabled: false  # Disabled by default for security
    timeout: 5s
    max_memory: 128MB
```

## Security Considerations

1. **Input Validation**: All tool inputs must be validated against their schema
2. **Execution Timeouts**: All tools must have configurable timeouts
3. **Resource Limits**: Memory and CPU limits for code execution tools
4. **Authorization**: Role-based access control for sensitive tools
5. **Audit Logging**: All tool executions must be logged
6. **Sandboxing**: Isolate tool execution in containers/VMs where possible

## Performance Considerations

1. **Tool Caching**: Cache idempotent tool results
2. **Parallel Execution**: Execute independent tool calls in parallel
3. **Connection Pooling**: Reuse HTTP connections for external APIs
4. **Lazy Loading**: Load tool plugins on-demand
5. **Metrics**: Monitor tool execution time and resource usage

## Migration Path

### Phase 1: Parallel Implementation
- Implement tool registry alongside existing code
- Keep existing web search integration unchanged
- Test new system independently

### Phase 2: Gradual Migration
- Migrate web search to tool-based approach
- Update pipeline to use tool registry
- Maintain backward compatibility

### Phase 3: Full Adoption
- Remove legacy tool implementations
- Make tool registry the primary interface
- Update documentation and examples

## Future Enhancements

1. **Tool Chaining**: Declarative tool chains with dependency management
2. **Multi-Modal Tools**: Support for image, audio, video processing
3. **Distributed Tools**: Execute tools across multiple nodes
4. **Tool Marketplace**: Registry of community-contributed tools
5. **Visual Tool Builder**: UI for creating and configuring tools
6. **Smart Tool Selection**: AI-powered tool recommendation based on task
7. **Tool Versioning**: Support for multiple versions of the same tool
8. **Streaming Tools**: Support for streaming tool outputs

## Success Metrics

1. **Developer Experience**: Time to add new tool < 30 minutes
2. **Performance**: Tool execution overhead < 10ms
3. **Reliability**: Tool execution success rate > 99%
4. **Coverage**: Unit test coverage > 90%
5. **Documentation**: All public APIs fully documented

## References

- [OpenAI Tool Calling Documentation](https://platform.openai.com/docs/guides/function-calling)
- [JSON Schema Specification](https://json-schema.org/)
- [Go Plugin Package](https://pkg.go.dev/plugin)
- [OpenAI Responses API](https://platform.openai.com/docs/api-reference/responses)

## Team & Timeline

**Team**: 1-2 backend engineers
**Duration**: 4 weeks
**Review Cycles**: Weekly

### Week 1: Foundation
- Core types and interfaces
- Registry implementation
- Basic validation

### Week 2: Integration
- Built-in tools
- LLM pipeline integration
- End-to-end testing

### Week 3: Enhancement
- Middleware system
- Metrics and logging
- Performance optimization

### Week 4: Polish
- Documentation
- Advanced features
- Production readiness

---

**Last Updated**: 2025-10-31  
**Version**: 1.0  
**Status**: Planning
