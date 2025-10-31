# AI Resolver Tool Registry - Development Plan

## Executive Summary

This document outlines a plan to extend the `ai-resolver` with a **Tool Registry System** that enables the AI to use custom tools for data aggregation. The system will allow developers to create specialized tools that the AI can invoke dynamically during market resolution analysis.

---

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Base Tool Interface](#base-tool-interface)
4. [Built-in Tools](#built-in-tools)
5. [Custom Tool Development](#custom-tool-development)
6. [Integration with LLM Pipeline](#integration-with-llm-pipeline)
7. [Implementation Roadmap](#implementation-roadmap)
8. [Example Tools](#example-tools)

---

## 1. Overview

### 1.1 Goals

- **Extensibility**: Allow developers to create custom data aggregation tools
- **Flexibility**: Support both JSON schema tools and custom text-based tools
- **Performance**: Efficient tool execution with caching and rate limiting
- **Security**: Sandboxed execution environment for untrusted tools
- **Observability**: Comprehensive logging and monitoring of tool usage

### 1.2 Use Cases

- **Sports Results**: Fetch game scores, statistics, player performance
- **Financial Data**: Get stock prices, market caps, trading volumes
- **Weather Data**: Query weather conditions, forecasts, historical data
- **Social Media**: Analyze trends, sentiment, engagement metrics
- **Blockchain Data**: Query on-chain metrics, transaction data, token prices
- **News Aggregation**: Fetch and summarize recent news articles
- **Custom APIs**: Integrate domain-specific data sources

---

## 2. Architecture

### 2.1 System Design

```
┌─────────────────────────────────────────────────────────────┐
│                    AI Resolver Server                        │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────────┐        ┌─────────────────────┐       │
│  │  LLM Pipeline    │───────▶│   Tool Registry     │       │
│  │  (OpenAI)        │        │  - Register tools   │       │
│  └──────────────────┘        │  - Discover tools   │       │
│           │                   │  - Execute tools    │       │
│           │                   │  - Cache results    │       │
│           ▼                   └─────────────────────┘       │
│  ┌──────────────────┐                 │                     │
│  │  Tool Executor   │◀────────────────┘                     │
│  │  - Validation    │                                       │
│  │  - Execution     │                                       │
│  │  - Error handling│                                       │
│  └──────────────────┘                                       │
│           │                                                  │
│           ▼                                                  │
│  ┌─────────────────────────────────────────────┐           │
│  │           Tool Implementations                │           │
│  ├─────────────────────────────────────────────┤           │
│  │ • WebSearchTool (built-in)                  │           │
│  │ • SportsDataTool                            │           │
│  │ • FinancialDataTool                         │           │
│  │ • BlockchainDataTool                        │           │
│  │ • CustomTool (user-defined)                 │           │
│  └─────────────────────────────────────────────┘           │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Package Structure

```
ai-resolver/
├── internal/
│   ├── tools/
│   │   ├── registry.go          # Tool registry implementation
│   │   ├── base.go              # Base tool interface
│   │   ├── executor.go          # Tool execution engine
│   │   ├── types.go             # Tool types and schemas
│   │   ├── cache.go             # Result caching layer
│   │   ├── builtin/             # Built-in tools
│   │   │   ├── web_search.go    # Web search tool
│   │   │   ├── sports.go        # Sports data tool
│   │   │   ├── financial.go     # Financial data tool
│   │   │   ├── blockchain.go    # Blockchain data tool
│   │   │   └── weather.go       # Weather data tool
│   │   └── custom/              # Custom user tools
│   │       └── example.go       # Example custom tool
│   ├── llm/
│   │   ├── pipeline.go          # Updated with tool support
│   │   ├── openai.go            # Updated with tool calls
│   │   └── tool_integration.go  # Tool integration logic
│   └── ...
└── ...
```

---

## 3. Base Tool Interface

### 3.1 Tool Interface Definition

```go
// File: internal/tools/base.go
package tools

import (
	"context"
)

// Tool represents a data aggregation tool that the AI can use
type Tool interface {
	// Name returns the unique identifier for this tool
	Name() string

	// Description returns what this tool does (for AI context)
	Description() string

	// Type returns whether this is a JSON schema or custom tool
	Type() ToolType

	// Schema returns the JSON schema for this tool (nil for custom tools)
	Schema() *ToolSchema

	// Execute runs the tool with the given input
	// For JSON schema tools, input is parsed JSON
	// For custom tools, input is a raw string
	Execute(ctx context.Context, input string) (*ToolResult, error)

	// Validate checks if the input is valid for this tool
	Validate(input string) error

	// CacheKey returns a unique cache key for this input (optional)
	CacheKey(input string) string

	// CacheDuration returns how long results should be cached (0 = no cache)
	CacheDuration() time.Duration
}

// ToolType defines the type of tool
type ToolType string

const (
	ToolTypeJSONSchema ToolType = "json_schema" // OpenAI function tool
	ToolTypeCustom     ToolType = "custom"      // Custom string-based tool
)

// ToolSchema defines the JSON schema for function tools
type ToolSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}

// ToolResult represents the result of a tool execution
type ToolResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Cached  bool        `json:"cached"`
	
	// Metadata
	ExecutionTime time.Duration      `json:"executionTime"`
	Sources       []string           `json:"sources,omitempty"`
	Metadata      map[string]string  `json:"metadata,omitempty"`
}
```

### 3.2 Base Tool Implementation

```go
// BaseTool provides common functionality for all tools
type BaseTool struct {
	name          string
	description   string
	toolType      ToolType
	schema        *ToolSchema
	cacheDuration time.Duration
}

func (b *BaseTool) Name() string {
	return b.name
}

func (b *BaseTool) Description() string {
	return b.description
}

func (b *BaseTool) Type() ToolType {
	return b.toolType
}

func (b *BaseTool) Schema() *ToolSchema {
	return b.schema
}

func (b *BaseTool) CacheDuration() time.Duration {
	return b.cacheDuration
}

func (b *BaseTool) CacheKey(input string) string {
	h := sha256.New()
	h.Write([]byte(b.name + ":" + input))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *BaseTool) Validate(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}
```

---

## 4. Built-in Tools

### 4.1 Web Search Tool (Already Integrated)

```go
// File: internal/tools/builtin/web_search.go
package builtin

import (
	"context"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

type WebSearchTool struct {
	*tools.BaseTool
	apiKey string
}

func NewWebSearchTool(apiKey string) *WebSearchTool {
	return &WebSearchTool{
		BaseTool: &tools.BaseTool{
			name:          "web_search",
			description:   "Search the web for current information about any topic",
			toolType:      tools.ToolTypeCustom,
			cacheDuration: 30 * time.Minute,
		},
		apiKey: apiKey,
	}
}

func (t *WebSearchTool) Execute(ctx context.Context, query string) (*tools.ToolResult, error) {
	start := time.Now()
	
	// Use OpenAI Responses API with web search
	// (Current implementation in openai.go)
	
	return &tools.ToolResult{
		Success:       true,
		Data:          searchResults,
		ExecutionTime: time.Since(start),
		Sources:       extractedURLs,
	}, nil
}
```

### 4.2 Sports Data Tool

```go
// File: internal/tools/builtin/sports.go
package builtin

type SportsDataTool struct {
	*tools.BaseTool
	apiKey string
}

func NewSportsDataTool(apiKey string) *SportsDataTool {
	return &SportsDataTool{
		BaseTool: &tools.BaseTool{
			name:        "sports_data",
			description: "Get real-time sports scores, statistics, and game results",
			toolType:    tools.ToolTypeJSONSchema,
			schema: &tools.ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"sport": map[string]interface{}{
						"type":        "string",
						"description": "Sport type (nfl, nba, mlb, soccer, etc.)",
					},
					"query": map[string]interface{}{
						"type":        "string",
						"description": "What to search for (team name, player, game, etc.)",
					},
					"date": map[string]interface{}{
						"type":        "string",
						"description": "Date for the query (YYYY-MM-DD, optional)",
					},
				},
				Required: []string{"sport", "query"},
			},
			cacheDuration: 5 * time.Minute,
		},
		apiKey: apiKey,
	}
}

func (t *SportsDataTool) Execute(ctx context.Context, input string) (*tools.ToolResult, error) {
	start := time.Now()
	
	// Parse JSON input
	var params struct {
		Sport string `json:"sport"`
		Query string `json:"query"`
		Date  string `json:"date"`
	}
	if err := json.Unmarshal([]byte(input), &params); err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   fmt.Sprintf("invalid input: %v", err),
		}, nil
	}
	
	// Call sports data API (e.g., ESPN, The Odds API, SportsData.io)
	results, err := t.querySportsAPI(ctx, params)
	if err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	
	return &tools.ToolResult{
		Success:       true,
		Data:          results,
		ExecutionTime: time.Since(start),
		Sources:       []string{"https://api.sportsdata.io"},
	}, nil
}

func (t *SportsDataTool) querySportsAPI(ctx context.Context, params struct {
	Sport string
	Query string
	Date  string
}) (interface{}, error) {
	// Implementation depends on sports API provider
	// Example: ESPN API, The Odds API, SportsData.io
	// Return structured data about games, scores, stats
	return nil, nil
}
```

### 4.3 Financial Data Tool

```go
// File: internal/tools/builtin/financial.go
package builtin

type FinancialDataTool struct {
	*tools.BaseTool
	apiKey string
}

func NewFinancialDataTool(apiKey string) *FinancialDataTool {
	return &FinancialDataTool{
		BaseTool: &tools.BaseTool{
			name:        "financial_data",
			description: "Get stock prices, market data, and financial metrics",
			toolType:    tools.ToolTypeJSONSchema,
			schema: &tools.ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"symbol": map[string]interface{}{
						"type":        "string",
						"description": "Stock ticker symbol (e.g., AAPL, TSLA)",
					},
					"metric": map[string]interface{}{
						"type":        "string",
						"description": "Metric to fetch (price, volume, market_cap, etc.)",
					},
					"date": map[string]interface{}{
						"type":        "string",
						"description": "Date for historical data (YYYY-MM-DD, optional)",
					},
				},
				Required: []string{"symbol"},
			},
			cacheDuration: 1 * time.Minute,
		},
		apiKey: apiKey,
	}
}

func (t *FinancialDataTool) Execute(ctx context.Context, input string) (*tools.ToolResult, error) {
	// Parse input and query financial API
	// Examples: Alpha Vantage, Yahoo Finance, Polygon.io
	return nil, nil
}
```

### 4.4 Blockchain Data Tool

```go
// File: internal/tools/builtin/blockchain.go
package builtin

type BlockchainDataTool struct {
	*tools.BaseTool
	ethClient *ethclient.Client
}

func NewBlockchainDataTool(rpcURL string) *BlockchainDataTool {
	return &BlockchainDataTool{
		BaseTool: &tools.BaseTool{
			name:        "blockchain_data",
			description: "Query blockchain data including token prices, transactions, and on-chain metrics",
			toolType:    tools.ToolTypeJSONSchema,
			schema: &tools.ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"chain": map[string]interface{}{
						"type":        "string",
						"description": "Blockchain (ethereum, bsc, polygon, etc.)",
					},
					"query_type": map[string]interface{}{
						"type":        "string",
						"description": "Type of query (token_price, balance, transaction, etc.)",
					},
					"address": map[string]interface{}{
						"type":        "string",
						"description": "Contract or wallet address",
					},
				},
				Required: []string{"chain", "query_type"},
			},
			cacheDuration: 30 * time.Second,
		},
	}
}

func (t *BlockchainDataTool) Execute(ctx context.Context, input string) (*tools.ToolResult, error) {
	// Query blockchain data using ethclient or API services
	// Examples: Etherscan, DexScreener, CoinGecko
	return nil, nil
}
```

### 4.5 Weather Data Tool

```go
// File: internal/tools/builtin/weather.go
package builtin

type WeatherDataTool struct {
	*tools.BaseTool
	apiKey string
}

func NewWeatherDataTool(apiKey string) *WeatherDataTool {
	return &WeatherDataTool{
		BaseTool: &tools.BaseTool{
			name:        "weather_data",
			description: "Get weather conditions, forecasts, and historical weather data",
			toolType:    tools.ToolTypeJSONSchema,
			schema: &tools.ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"location": map[string]interface{}{
						"type":        "string",
						"description": "City name or coordinates",
					},
					"date": map[string]interface{}{
						"type":        "string",
						"description": "Date for weather query (YYYY-MM-DD)",
					},
					"metric": map[string]interface{}{
						"type":        "string",
						"description": "Metric (temperature, precipitation, wind, etc.)",
					},
				},
				Required: []string{"location", "date"},
			},
			cacheDuration: 10 * time.Minute,
		},
		apiKey: apiKey,
	}
}

func (t *WeatherDataTool) Execute(ctx context.Context, input string) (*tools.ToolResult, error) {
	// Query weather API (OpenWeatherMap, Weather.gov, etc.)
	return nil, nil
}
```

---

## 5. Custom Tool Development

### 5.1 Tool Registry

```go
// File: internal/tools/registry.go
package tools

import (
	"context"
	"fmt"
	"sync"
)

// Registry manages available tools
type Registry struct {
	mu    sync.RWMutex
	tools map[string]Tool
	cache *Cache
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
		cache: NewCache(),
	}
}

// Register adds a tool to the registry
func (r *Registry) Register(tool Tool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, exists := r.tools[tool.Name()]; exists {
		return fmt.Errorf("tool %s already registered", tool.Name())
	}
	
	r.tools[tool.Name()] = tool
	return nil
}

// Get retrieves a tool by name
func (r *Registry) Get(name string) (Tool, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	tool, ok := r.tools[name]
	return tool, ok
}

// List returns all registered tools
func (r *Registry) List() []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// Execute runs a tool with caching
func (r *Registry) Execute(ctx context.Context, toolName, input string) (*ToolResult, error) {
	tool, ok := r.Get(toolName)
	if !ok {
		return nil, fmt.Errorf("tool not found: %s", toolName)
	}
	
	// Check cache
	if tool.CacheDuration() > 0 {
		cacheKey := tool.CacheKey(input)
		if cached, found := r.cache.Get(cacheKey); found {
			result := cached.(*ToolResult)
			result.Cached = true
			return result, nil
		}
	}
	
	// Validate input
	if err := tool.Validate(input); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	
	// Execute tool
	result, err := tool.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	
	// Cache result
	if tool.CacheDuration() > 0 && result.Success {
		cacheKey := tool.CacheKey(input)
		r.cache.Set(cacheKey, result, tool.CacheDuration())
	}
	
	return result, nil
}

// GetOpenAITools returns tools formatted for OpenAI API
func (r *Registry) GetOpenAITools() []interface{} {
	tools := r.List()
	openAITools := make([]interface{}, 0, len(tools))
	
	for _, tool := range tools {
		switch tool.Type() {
		case ToolTypeJSONSchema:
			openAITools = append(openAITools, map[string]interface{}{
				"type": "function",
				"function": map[string]interface{}{
					"name":        tool.Name(),
					"description": tool.Description(),
					"parameters":  tool.Schema(),
				},
			})
		case ToolTypeCustom:
			openAITools = append(openAITools, map[string]interface{}{
				"type":        "custom",
				"name":        tool.Name(),
				"description": tool.Description(),
			})
		}
	}
	
	return openAITools
}
```

### 5.2 Cache Implementation

```go
// File: internal/tools/cache.go
package tools

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]*cacheItem),
	}
	
	// Start cleanup goroutine
	go c.cleanup()
	
	return c
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.items[key] = &cacheItem{
		value:      value,
		expiration: time.Now().Add(duration),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, exists := c.items[key]
	if !exists || time.Now().After(item.expiration) {
		return nil, false
	}
	
	return item.value, true
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, item := range c.items {
			if now.After(item.expiration) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
```

---

## 6. Integration with LLM Pipeline

### 6.1 Updated Pipeline Interface

```go
// File: internal/llm/pipeline.go (updated)
package llm

import (
	"context"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

// Pipeline defines the multi-pass LLM analysis pipeline with tool support
type Pipeline interface {
	// AnalyzeMarket performs the complete analysis pipeline
	AnalyzeMarket(ctx context.Context, market MarketInfo) (*Decision, error)
	
	// SetToolRegistry configures the tool registry for this pipeline
	SetToolRegistry(registry *tools.Registry)
	
	// GetToolRegistry returns the current tool registry
	GetToolRegistry() *tools.Registry
}
```

### 6.2 Tool-Enabled OpenAI Pipeline

```go
// File: internal/llm/tool_integration.go
package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

// Updated OpenAIPipeline with tool support
func (p *OpenAIPipeline) SetToolRegistry(registry *tools.Registry) {
	p.toolRegistry = registry
}

func (p *OpenAIPipeline) GetToolRegistry() *tools.Registry {
	return p.toolRegistry
}

// callOpenAIWithTools makes a request with tool support
func (p *OpenAIPipeline) callOpenAIWithTools(ctx context.Context, prompt string) (string, []ToolCall, error) {
	// Get available tools from registry
	openAITools := p.toolRegistry.GetOpenAITools()
	
	reqBody := map[string]interface{}{
		"model": p.model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are an AI assistant with access to tools for data aggregation."},
			{"role": "user", "content": prompt},
		},
		"tools":       openAITools,
		"tool_choice": "auto",
	}
	
	// Make API call
	response, err := p.makeAPICall(ctx, reqBody)
	if err != nil {
		return "", nil, err
	}
	
	// Parse response and extract tool calls
	toolCalls := p.extractToolCalls(response)
	
	return response.Content, toolCalls, nil
}

// executeToolCalls processes tool calls from the AI
func (p *OpenAIPipeline) executeToolCalls(ctx context.Context, toolCalls []ToolCall) ([]ToolCallResult, error) {
	results := make([]ToolCallResult, 0, len(toolCalls))
	
	for _, call := range toolCalls {
		result, err := p.toolRegistry.Execute(ctx, call.Name, call.Input)
		if err != nil {
			results = append(results, ToolCallResult{
				CallID:  call.ID,
				Success: false,
				Error:   err.Error(),
			})
			continue
		}
		
		results = append(results, ToolCallResult{
			CallID:  call.ID,
			Success: result.Success,
			Data:    result.Data,
			Sources: result.Sources,
		})
	}
	
	return results, nil
}

type ToolCall struct {
	ID    string
	Name  string
	Input string
}

type ToolCallResult struct {
	CallID  string
	Success bool
	Data    interface{}
	Error   string
	Sources []string
}
```

---

## 7. Implementation Roadmap

### Phase 1: Foundation (Week 1)
- [ ] Create base tool interface and types
- [ ] Implement tool registry
- [ ] Implement caching layer
- [ ] Create tool executor with validation
- [ ] Write unit tests for core components

### Phase 2: Built-in Tools (Week 2)
- [ ] Refactor existing web search into WebSearchTool
- [ ] Implement SportsDataTool
- [ ] Implement FinancialDataTool
- [ ] Implement BlockchainDataTool
- [ ] Implement WeatherDataTool
- [ ] Add API key management for tools

### Phase 3: LLM Integration (Week 3)
- [ ] Update Pipeline interface with tool support
- [ ] Integrate tool registry with OpenAI pipeline
- [ ] Implement tool call handling (JSON schema & custom)
- [ ] Add multi-turn conversations for tool results
- [ ] Update fact extraction to use tool data

### Phase 4: Server Integration (Week 4)
- [ ] Initialize tool registry in server startup
- [ ] Register built-in tools
- [ ] Add configuration for tool API keys
- [ ] Add tool usage metrics and logging
- [ ] Update API responses with tool metadata

### Phase 5: Documentation & Testing (Week 5)
- [ ] Write comprehensive documentation
- [ ] Create example custom tools
- [ ] Write integration tests
- [ ] Add performance benchmarks
- [ ] Create developer guide for custom tools

---

## 8. Example Tools

### 8.1 Example: Crypto Price Tool

```go
// File: internal/tools/custom/crypto_price.go
package custom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

type CryptoPriceTool struct {
	*tools.BaseTool
	httpClient *http.Client
}

func NewCryptoPriceTool() *CryptoPriceTool {
	return &CryptoPriceTool{
		BaseTool: &tools.BaseTool{
			name:        "crypto_price",
			description: "Get current cryptocurrency prices from CoinGecko",
			toolType:    tools.ToolTypeJSONSchema,
			schema: &tools.ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"symbol": map[string]interface{}{
						"type":        "string",
						"description": "Cryptocurrency symbol (e.g., BTC, ETH, BNB)",
					},
					"currency": map[string]interface{}{
						"type":        "string",
						"description": "Fiat currency (default: usd)",
					},
				},
				Required: []string{"symbol"},
			},
			cacheDuration: 1 * time.Minute,
		},
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (t *CryptoPriceTool) Execute(ctx context.Context, input string) (*tools.ToolResult, error) {
	start := time.Now()
	
	var params struct {
		Symbol   string `json:"symbol"`
		Currency string `json:"currency"`
	}
	if err := json.Unmarshal([]byte(input), &params); err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   fmt.Sprintf("invalid input: %v", err),
		}, nil
	}
	
	if params.Currency == "" {
		params.Currency = "usd"
	}
	
	// Query CoinGecko API
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
		params.Symbol, params.Currency)
	
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	defer resp.Body.Close()
	
	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   fmt.Sprintf("failed to parse response: %v", err),
		}, nil
	}
	
	return &tools.ToolResult{
		Success:       true,
		Data:          result,
		ExecutionTime: time.Since(start),
		Sources:       []string{"https://www.coingecko.com"},
	}, nil
}
```

### 8.2 Example: Custom Code Execution Tool

```go
// File: internal/tools/custom/code_exec.go
package custom

import (
	"context"
	"fmt"
	"os/exec"
	"time"
	"github.com/project-gamma/ai-resolver/internal/tools"
)

type CodeExecTool struct {
	*tools.BaseTool
}

func NewCodeExecTool() *CodeExecTool {
	return &CodeExecTool{
		BaseTool: &tools.BaseTool{
			name:          "code_exec",
			description:   "Executes arbitrary Python code for data processing",
			toolType:      tools.ToolTypeCustom,
			cacheDuration: 0, // No caching for code execution
		},
	}
}

func (t *CodeExecTool) Execute(ctx context.Context, code string) (*tools.ToolResult, error) {
	start := time.Now()
	
	// SECURITY WARNING: This is dangerous in production!
	// Should use sandboxed environment (Docker, gVisor, etc.)
	
	// Create timeout context
	execCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	// Execute Python code
	cmd := exec.CommandContext(execCtx, "python3", "-c", code)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return &tools.ToolResult{
			Success: false,
			Error:   fmt.Sprintf("execution failed: %v\nOutput: %s", err, string(output)),
		}, nil
	}
	
	return &tools.ToolResult{
		Success:       true,
		Data:          string(output),
		ExecutionTime: time.Since(start),
	}, nil
}

func (t *CodeExecTool) Validate(code string) error {
	// Basic validation
	if len(code) > 10000 {
		return fmt.Errorf("code too long (max 10000 chars)")
	}
	
	// Could add more security checks here
	// - Blacklist dangerous imports (os, subprocess, etc.)
	// - AST parsing to detect suspicious patterns
	
	return nil
}
```

---

## 9. Configuration

### 9.1 Environment Variables

```bash
# Tool API Keys
SPORTS_API_KEY=your_sports_api_key
FINANCIAL_API_KEY=your_financial_api_key
WEATHER_API_KEY=your_weather_api_key

# Tool Settings
TOOL_CACHE_ENABLED=true
TOOL_CACHE_MAX_SIZE=1000
TOOL_EXECUTION_TIMEOUT=30s
TOOL_MAX_RETRIES=3

# Feature Flags
ENABLE_SPORTS_TOOL=true
ENABLE_FINANCIAL_TOOL=true
ENABLE_BLOCKCHAIN_TOOL=true
ENABLE_WEATHER_TOOL=true
ENABLE_CUSTOM_TOOLS=false  # Disable in production for security
```

### 9.2 Server Initialization

```go
// File: cmd/server/main.go (updated)
func main() {
	// ... existing setup ...
	
	// Initialize tool registry
	toolRegistry := tools.NewRegistry()
	
	// Register built-in tools
	if cfg.EnableWebSearch {
		toolRegistry.Register(builtin.NewWebSearchTool(cfg.OpenAIAPIKey))
	}
	
	if cfg.EnableSportsTool && cfg.SportsAPIKey != "" {
		toolRegistry.Register(builtin.NewSportsDataTool(cfg.SportsAPIKey))
	}
	
	if cfg.EnableFinancialTool && cfg.FinancialAPIKey != "" {
		toolRegistry.Register(builtin.NewFinancialDataTool(cfg.FinancialAPIKey))
	}
	
	if cfg.EnableBlockchainTool {
		toolRegistry.Register(builtin.NewBlockchainDataTool(cfg.RPCEndpoint))
	}
	
	if cfg.EnableWeatherTool && cfg.WeatherAPIKey != "" {
		toolRegistry.Register(builtin.NewWeatherDataTool(cfg.WeatherAPIKey))
	}
	
	// Set tool registry in LLM pipeline
	llmPipeline.SetToolRegistry(toolRegistry)
	
	// ... rest of server setup ...
}
```

---

## 10. Security Considerations

### 10.1 Sandboxing

- **Code Execution**: Use Docker containers or gVisor for isolated execution
- **API Keys**: Store securely, rotate regularly
- **Input Validation**: Sanitize all tool inputs
- **Rate Limiting**: Prevent abuse of external APIs

### 10.2 Access Control

```go
// Tool permissions
type ToolPermissions struct {
	AllowedTools []string
	RateLimit    int
	MaxCost      float64  // Max API cost per request
}

// Check permissions before tool execution
func (r *Registry) ExecuteWithPermissions(ctx context.Context, toolName, input string, perms *ToolPermissions) (*ToolResult, error) {
	// Check if tool is allowed
	if !contains(perms.AllowedTools, toolName) {
		return nil, fmt.Errorf("tool not permitted: %s", toolName)
	}
	
	// Check rate limit
	// Implement rate limiting logic
	
	return r.Execute(ctx, toolName, input)
}
```

---

## 11. Monitoring & Observability

### 11.1 Metrics

```go
type ToolMetrics struct {
	TotalCalls       int64
	SuccessfulCalls  int64
	FailedCalls      int64
	CacheHits        int64
	CacheMisses      int64
	AvgExecutionTime time.Duration
	TotalAPIcost     float64
}

// Track metrics for each tool
func (r *Registry) GetMetrics(toolName string) *ToolMetrics {
	// Return metrics for the tool
	return nil
}
```

### 11.2 Logging

```go
// Log tool execution
log.Printf("Tool: %s, Input: %s, Duration: %v, Success: %v, Cached: %v",
	toolName, truncate(input, 100), result.ExecutionTime, result.Success, result.Cached)
```

---

## 12. Success Criteria

### Technical
- [ ] Tool registry supports 5+ built-in tools
- [ ] Tool execution time < 5 seconds (p95)
- [ ] Cache hit rate > 30%
- [ ] Zero security vulnerabilities
- [ ] Comprehensive test coverage (>80%)

### Developer Experience
- [ ] Clear documentation for creating custom tools
- [ ] Example tools in multiple categories
- [ ] Easy tool registration (3 lines of code)
- [ ] Helpful error messages

### Production Readiness
- [ ] Graceful error handling
- [ ] Rate limiting and quotas
- [ ] Monitoring and alerting
- [ ] Performance benchmarks

---

## Conclusion

This tool registry system will significantly enhance the AI resolver's capabilities by:

1. **Enabling specialized data sources** - Access sports, financial, blockchain, and other domain-specific data
2. **Improving accuracy** - Use authoritative APIs instead of just web search
3. **Extensibility** - Easy to add new tools for specific use cases
4. **Performance** - Caching and optimized execution
5. **Reliability** - Structured error handling and validation

The system is designed to be production-ready with security, monitoring, and scalability in mind.

---

**Document Version**: 1.0  
**Last Updated**: 2025-10-31  
**Status**: Ready for Implementation
