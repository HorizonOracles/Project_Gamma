package tools

import (
	"context"
)

// WebSearchTool wraps OpenAI's web_search_preview functionality
// This tool type doesn't require parameters - OpenAI handles it automatically
type WebSearchTool struct {
	*BaseTool
}

// NewWebSearchTool creates a new web search tool
func NewWebSearchTool() *WebSearchTool {
	base := NewBaseTool(
		"web_search",
		"Search the web for current information using OpenAI's integrated web search. The LLM will automatically use this when it needs up-to-date information from the internet.",
		ToolTypeWebSearchPreview,
		nil, // No schema needed - OpenAI handles this internally
	)

	tool := &WebSearchTool{
		BaseTool: base,
	}

	// Set executor - for web_search_preview, execution is handled by OpenAI
	// This executor is a no-op that returns metadata
	base.SetExecutor(tool.execute)

	return tool
}

// execute is a no-op executor since web search is handled by OpenAI
func (t *WebSearchTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	// Web search is handled entirely by OpenAI's API
	// This method shouldn't be called in normal flow
	return ToolOutput{
		CallID: input.CallID,
		Data: map[string]any{
			"message": "Web search is handled by OpenAI API",
			"status":  "delegated_to_openai",
		},
	}, nil
}
