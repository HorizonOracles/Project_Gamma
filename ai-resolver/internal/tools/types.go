package tools

import (
	"time"
)

// ToolType represents the type of tool
type ToolType string

const (
	// ToolTypeFunction represents a function tool with JSON schema parameters
	ToolTypeFunction ToolType = "function"

	// ToolTypeCustom represents a custom tool that accepts arbitrary string input
	ToolTypeCustom ToolType = "custom"

	// ToolTypeWebSearchPreview represents OpenAI's web search tool
	ToolTypeWebSearchPreview ToolType = "web_search_preview"
)

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
	CallID        string
	ExecutionTime time.Duration
	Logs          []string
}

// AddLog adds a log entry to the output
func (o *ToolOutput) AddLog(log string) {
	o.Logs = append(o.Logs, log)
}

// HasError returns true if the output contains an error
func (o *ToolOutput) HasError() bool {
	return o.Error != nil
}
