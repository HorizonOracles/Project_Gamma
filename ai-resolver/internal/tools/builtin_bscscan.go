package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// BSCScanTool fetches blockchain data from BSCScan API
type BSCScanTool struct {
	*BaseTool
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewBSCScanTool creates a new BSCScan API tool
func NewBSCScanTool(apiKey string) *BSCScanTool {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"action": {
				Type:        "string",
				Description: "The BSCScan API action: balance (get BNB balance), tokenbalance (get token balance), txlist (get transaction list), tokentx (get token transfers), price (get BNB price)",
				Enum:        []string{"balance", "tokenbalance", "txlist", "tokentx", "price"},
			},
			"address": {
				Type:        "string",
				Description: "Wallet or contract address (required for balance, tokenbalance, txlist, tokentx)",
			},
			"contract_address": {
				Type:        "string",
				Description: "Token contract address (required for tokenbalance and tokentx)",
			},
			"startblock": {
				Type:        "integer",
				Description: "Starting block number for transaction queries (optional, default: 0)",
			},
			"endblock": {
				Type:        "integer",
				Description: "Ending block number for transaction queries (optional, default: latest)",
			},
		},
		Required: []string{"action"},
	}

	base := NewBaseTool(
		"bscscan",
		"Query BSC blockchain data via BSCScan API including wallet balances, token balances, transaction history, and BNB price. Returns real-time on-chain data.",
		ToolTypeFunction,
		schema,
	)

	tool := &BSCScanTool{
		BaseTool: base,
		apiKey:   apiKey,
		baseURL:  "https://api.bscscan.com/api",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	base.SetExecutor(tool.execute)

	return tool
}

// execute performs the BSCScan API call
func (t *BSCScanTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	action, ok := input.Arguments["action"].(string)
	if !ok || action == "" {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("action is required"),
		}, fmt.Errorf("action is required")
	}

	// Build API parameters based on action
	params := url.Values{}
	params.Set("apikey", t.apiKey)

	var result map[string]any
	var err error

	switch action {
	case "balance":
		result, err = t.getBalance(ctx, input.Arguments, params)
	case "tokenbalance":
		result, err = t.getTokenBalance(ctx, input.Arguments, params)
	case "txlist":
		result, err = t.getTransactionList(ctx, input.Arguments, params)
	case "tokentx":
		result, err = t.getTokenTransfers(ctx, input.Arguments, params)
	case "price":
		result, err = t.getBNBPrice(ctx, params)
	default:
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("unknown action: %s", action),
		}, fmt.Errorf("unknown action: %s", action)
	}

	if err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("bscscan API call failed: %w", err),
			Data: map[string]any{
				"action": action,
				"error":  err.Error(),
			},
		}, fmt.Errorf("bscscan API call failed: %w", err)
	}

	return ToolOutput{
		CallID: input.CallID,
		Data:   result,
	}, nil
}

// getBalance fetches BNB balance for an address
func (t *BSCScanTool) getBalance(ctx context.Context, args map[string]any, params url.Values) (map[string]any, error) {
	address, ok := args["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address is required for balance action")
	}

	params.Set("module", "account")
	params.Set("action", "balance")
	params.Set("address", address)
	params.Set("tag", "latest")

	resp, err := t.makeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"address": address,
		"balance": resp["result"],
		"unit":    "wei",
	}, nil
}

// getTokenBalance fetches token balance for an address
func (t *BSCScanTool) getTokenBalance(ctx context.Context, args map[string]any, params url.Values) (map[string]any, error) {
	address, ok := args["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address is required for tokenbalance action")
	}

	contractAddress, ok := args["contract_address"].(string)
	if !ok || contractAddress == "" {
		return nil, fmt.Errorf("contract_address is required for tokenbalance action")
	}

	params.Set("module", "account")
	params.Set("action", "tokenbalance")
	params.Set("address", address)
	params.Set("contractaddress", contractAddress)
	params.Set("tag", "latest")

	resp, err := t.makeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"address":          address,
		"contract_address": contractAddress,
		"balance":          resp["result"],
		"unit":             "wei",
	}, nil
}

// getTransactionList fetches transaction history
func (t *BSCScanTool) getTransactionList(ctx context.Context, args map[string]any, params url.Values) (map[string]any, error) {
	address, ok := args["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address is required for txlist action")
	}

	params.Set("module", "account")
	params.Set("action", "txlist")
	params.Set("address", address)
	params.Set("startblock", t.getOptionalString(args, "startblock", "0"))
	params.Set("endblock", t.getOptionalString(args, "endblock", "99999999"))
	params.Set("page", "1")
	params.Set("offset", "10") // Limit to 10 recent transactions
	params.Set("sort", "desc")

	resp, err := t.makeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"address":      address,
		"transactions": resp["result"],
		"count":        len(resp["result"].([]any)),
	}, nil
}

// getTokenTransfers fetches token transfer history
func (t *BSCScanTool) getTokenTransfers(ctx context.Context, args map[string]any, params url.Values) (map[string]any, error) {
	address, ok := args["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address is required for tokentx action")
	}

	params.Set("module", "account")
	params.Set("action", "tokentx")
	params.Set("address", address)
	params.Set("startblock", t.getOptionalString(args, "startblock", "0"))
	params.Set("endblock", t.getOptionalString(args, "endblock", "99999999"))
	params.Set("page", "1")
	params.Set("offset", "10")
	params.Set("sort", "desc")

	if contractAddr, ok := args["contract_address"].(string); ok && contractAddr != "" {
		params.Set("contractaddress", contractAddr)
	}

	resp, err := t.makeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"address":   address,
		"transfers": resp["result"],
		"count":     len(resp["result"].([]any)),
	}, nil
}

// getBNBPrice fetches current BNB price
func (t *BSCScanTool) getBNBPrice(ctx context.Context, params url.Values) (map[string]any, error) {
	params.Set("module", "stats")
	params.Set("action", "bnbprice")

	resp, err := t.makeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	result := resp["result"].(map[string]any)
	return map[string]any{
		"bnb_btc":      result["ethbtc"],
		"bnb_usd":      result["ethusd"],
		"last_updated": result["ethusd_timestamp"],
		"description":  "Current BNB price in BTC and USD",
	}, nil
}

// makeRequest performs the HTTP request to BSCScan API
func (t *BSCScanTool) makeRequest(ctx context.Context, params url.Values) (map[string]any, error) {
	reqURL := fmt.Sprintf("%s?%s", t.baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check API status
	if status, ok := result["status"].(string); ok && status != "1" {
		errMsg := "unknown error"
		if msg, ok := result["message"].(string); ok {
			errMsg = msg
		}
		return nil, fmt.Errorf("API error: %s", errMsg)
	}

	return result, nil
}

// getOptionalString gets a string parameter with a default value
func (t *BSCScanTool) getOptionalString(args map[string]any, key, defaultVal string) string {
	if val, ok := args[key]; ok {
		switch v := val.(type) {
		case string:
			return v
		case int, int64, float64:
			return fmt.Sprintf("%v", v)
		}
	}
	return defaultVal
}
