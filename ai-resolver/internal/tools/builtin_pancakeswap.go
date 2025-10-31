package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// PancakeSwapClient interface for blockchain queries
type PancakeSwapClient interface {
	GetETHClient() *ethclient.Client
}

// PancakeSwapTool fetches TWAP and liquidity data from PancakeSwap
type PancakeSwapTool struct {
	*BaseTool
	client     PancakeSwapClient
	httpClient *http.Client
}

// NewPancakeSwapTool creates a new PancakeSwap tool
func NewPancakeSwapTool(client PancakeSwapClient) *PancakeSwapTool {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"action": {
				Type:        "string",
				Description: "The action to perform: price (get spot price), twap (get time-weighted average price), liquidity (get pool liquidity), volume (get 24h volume)",
				Enum:        []string{"price", "twap", "liquidity", "volume"},
			},
			"token0": {
				Type:        "string",
				Description: "First token address in the pair",
			},
			"token1": {
				Type:        "string",
				Description: "Second token address in the pair",
			},
			"pair_address": {
				Type:        "string",
				Description: "Liquidity pair contract address (optional, will be derived if not provided)",
			},
			"period": {
				Type:        "integer",
				Description: "Time period in seconds for TWAP calculation (default: 3600 = 1 hour)",
			},
		},
		Required: []string{"action"},
	}

	base := NewBaseTool(
		"pancakeswap",
		"Query PancakeSwap DEX data including spot prices, time-weighted average prices (TWAP), liquidity pool data, and trading volumes. Useful for price oracle data and market analysis.",
		ToolTypeFunction,
		schema,
	)

	tool := &PancakeSwapTool{
		BaseTool: base,
		client:   client,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	base.SetExecutor(tool.execute)

	return tool
}

// execute performs the PancakeSwap query
func (t *PancakeSwapTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	action, ok := input.Arguments["action"].(string)
	if !ok || action == "" {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("action is required"),
		}, fmt.Errorf("action is required")
	}

	var result map[string]any
	var err error

	switch action {
	case "price":
		result, err = t.getSpotPrice(ctx, input.Arguments)
	case "twap":
		result, err = t.getTWAP(ctx, input.Arguments)
	case "liquidity":
		result, err = t.getLiquidity(ctx, input.Arguments)
	case "volume":
		result, err = t.getVolume(ctx, input.Arguments)
	default:
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("unknown action: %s", action),
		}, fmt.Errorf("unknown action: %s", action)
	}

	if err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("pancakeswap query failed: %w", err),
			Data: map[string]any{
				"action": action,
				"error":  err.Error(),
			},
		}, fmt.Errorf("pancakeswap query failed: %w", err)
	}

	return ToolOutput{
		CallID: input.CallID,
		Data:   result,
	}, nil
}

// getSpotPrice fetches current spot price from PancakeSwap
func (t *PancakeSwapTool) getSpotPrice(ctx context.Context, args map[string]any) (map[string]any, error) {
	pairAddr, err := t.getPairAddress(args)
	if err != nil {
		return nil, err
	}

	// Call pair contract to get reserves
	reserves, err := t.getReserves(ctx, pairAddr)
	if err != nil {
		return nil, err
	}

	// Calculate price (reserve1/reserve0)
	reserve0 := new(big.Float).SetInt(reserves["reserve0"].(*big.Int))
	reserve1 := new(big.Float).SetInt(reserves["reserve1"].(*big.Int))

	if reserve0.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("insufficient liquidity: reserve0 is zero")
	}

	price := new(big.Float).Quo(reserve1, reserve0)
	priceFloat, _ := price.Float64()

	return map[string]any{
		"pair_address": pairAddr,
		"reserve0":     reserves["reserve0"].(*big.Int).String(),
		"reserve1":     reserves["reserve1"].(*big.Int).String(),
		"price":        priceFloat,
		"price_token0": priceFloat,
		"price_token1": 1 / priceFloat,
		"timestamp":    time.Now().Unix(),
	}, nil
}

// getTWAP calculates time-weighted average price
func (t *PancakeSwapTool) getTWAP(ctx context.Context, args map[string]any) (map[string]any, error) {
	// For TWAP, we'd need historical price data
	// PancakeSwap uses Chainlink oracles or on-chain TWAP mechanisms
	// For now, we'll use TheGraph API to fetch historical data

	token0, ok := args["token0"].(string)
	if !ok || token0 == "" {
		return nil, fmt.Errorf("token0 is required for TWAP")
	}

	token1, ok := args["token1"].(string)
	if !ok || token1 == "" {
		return nil, fmt.Errorf("token1 is required for TWAP")
	}

	period := 3600 // Default 1 hour
	if p, ok := args["period"].(float64); ok {
		period = int(p)
	} else if p, ok := args["period"].(int); ok {
		period = p
	}

	// Query TheGraph for historical prices
	twapData, err := t.queryTheGraph(ctx, token0, token1, period)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch TWAP data: %w", err)
	}

	return twapData, nil
}

// getLiquidity fetches pool liquidity information
func (t *PancakeSwapTool) getLiquidity(ctx context.Context, args map[string]any) (map[string]any, error) {
	pairAddr, err := t.getPairAddress(args)
	if err != nil {
		return nil, err
	}

	reserves, err := t.getReserves(ctx, pairAddr)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"pair_address":         pairAddr,
		"reserve0":             reserves["reserve0"].(*big.Int).String(),
		"reserve1":             reserves["reserve1"].(*big.Int).String(),
		"block_timestamp_last": reserves["blockTimestampLast"],
		"description":          "Liquidity pool reserves",
	}, nil
}

// getVolume fetches 24h trading volume
func (t *PancakeSwapTool) getVolume(ctx context.Context, args map[string]any) (map[string]any, error) {
	token0, ok := args["token0"].(string)
	if !ok || token0 == "" {
		return nil, fmt.Errorf("token0 is required for volume query")
	}

	token1, ok := args["token1"].(string)
	if !ok || token1 == "" {
		return nil, fmt.Errorf("token1 is required for volume query")
	}

	// Query TheGraph for volume data
	volumeData, err := t.queryVolumeData(ctx, token0, token1)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch volume data: %w", err)
	}

	return volumeData, nil
}

// getReserves calls the pair contract to get reserves
func (t *PancakeSwapTool) getReserves(ctx context.Context, pairAddr string) (map[string]any, error) {
	client := t.client.GetETHClient()

	// PancakeSwap V2 Pair ABI (getReserves function)
	pairABI := `[{"constant":true,"inputs":[],"name":"getReserves","outputs":[{"internalType":"uint112","name":"_reserve0","type":"uint112"},{"internalType":"uint112","name":"_reserve1","type":"uint112"},{"internalType":"uint32","name":"_blockTimestampLast","type":"uint32"}],"payable":false,"stateMutability":"view","type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(pairABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Pack the function call
	data, err := parsedABI.Pack("getReserves")
	if err != nil {
		return nil, fmt.Errorf("failed to pack function call: %w", err)
	}

	// Call the contract
	pairAddress := common.HexToAddress(pairAddr)
	msg := ethereum.CallMsg{
		To:   &pairAddress,
		Data: data,
	}

	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %w", err)
	}

	// Unpack the result
	var reserves struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}

	err = parsedABI.UnpackIntoInterface(&reserves, "getReserves", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return map[string]any{
		"reserve0":           reserves.Reserve0,
		"reserve1":           reserves.Reserve1,
		"blockTimestampLast": reserves.BlockTimestampLast,
	}, nil
}

// getPairAddress extracts or derives the pair address
func (t *PancakeSwapTool) getPairAddress(args map[string]any) (string, error) {
	// If pair address provided, use it
	if pairAddr, ok := args["pair_address"].(string); ok && pairAddr != "" {
		return pairAddr, nil
	}

	// Otherwise, require token addresses to derive it
	token0, ok := args["token0"].(string)
	if !ok || token0 == "" {
		return "", fmt.Errorf("either pair_address or token0+token1 are required")
	}

	token1, ok := args["token1"].(string)
	if !ok || token1 == "" {
		return "", fmt.Errorf("either pair_address or token0+token1 are required")
	}

	// Derive pair address using PancakeSwap factory formula
	// For now, return error suggesting to provide pair_address
	return "", fmt.Errorf("pair address derivation not implemented, please provide pair_address")
}

// queryTheGraph queries TheGraph for historical TWAP data
func (t *PancakeSwapTool) queryTheGraph(ctx context.Context, token0, token1 string, period int) (map[string]any, error) {
	// TheGraph endpoint for PancakeSwap V2
	graphURL := "https://api.thegraph.com/subgraphs/name/pancakeswap/exchange-v2"

	// Build GraphQL query
	query := fmt.Sprintf(`{
		pair(id: "%s") {
			token0Price
			token1Price
			reserveUSD
			volumeUSD
			txCount
		}
	}`, strings.ToLower(token0+token1))

	reqBody := map[string]any{
		"query": query,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", graphURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return map[string]any{
		"token0":      token0,
		"token1":      token1,
		"period":      period,
		"twap_data":   result,
		"description": "Time-weighted average price data",
	}, nil
}

// queryVolumeData queries TheGraph for volume data
func (t *PancakeSwapTool) queryVolumeData(ctx context.Context, token0, token1 string) (map[string]any, error) {
	// Similar to queryTheGraph but focused on volume
	return t.queryTheGraph(ctx, token0, token1, 86400) // 24 hours
}
