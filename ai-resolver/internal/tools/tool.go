package tools

import (
	"context"
	"fmt"
	"time"
)

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

	// ToOpenAIFormat converts the tool to OpenAI API format
	ToOpenAIFormat() map[string]any
}

// ToolExecutor is the function signature for tool execution logic
type ToolExecutor func(ctx context.Context, input ToolInput) (ToolOutput, error)

// ToolValidator validates tool input
type ToolValidator func(input ToolInput) error

// ToolMiddleware wraps tool execution with additional functionality
type ToolMiddleware func(next ToolExecutor) ToolExecutor

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

// NewBaseTool creates a new BaseTool
func NewBaseTool(name, description string, toolType ToolType, schema *ToolSchema) *BaseTool {
	return &BaseTool{
		name:        name,
		description: description,
		toolType:    toolType,
		schema:      schema,
		middleware:  make([]ToolMiddleware, 0),
	}
}

// Name returns the tool name
func (t *BaseTool) Name() string {
	return t.name
}

// Description returns the tool description
func (t *BaseTool) Description() string {
	return t.description
}

// Type returns the tool type
func (t *BaseTool) Type() ToolType {
	return t.toolType
}

// Schema returns the tool schema
func (t *BaseTool) Schema() *ToolSchema {
	return t.schema
}

// SetExecutor sets the execution function for the tool
func (t *BaseTool) SetExecutor(executor ToolExecutor) *BaseTool {
	t.executor = executor
	return t
}

// SetValidator sets the validation function for the tool
func (t *BaseTool) SetValidator(validator ToolValidator) *BaseTool {
	t.validator = validator
	return t
}

// WithMiddleware adds middleware to the tool
func (t *BaseTool) WithMiddleware(middleware ToolMiddleware) *BaseTool {
	t.middleware = append(t.middleware, middleware)
	return t
}

// Execute runs the tool with the provided input
func (t *BaseTool) Execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	// Set call ID and timestamp if not set
	if input.CallID == "" {
		input.CallID = fmt.Sprintf("call_%d", time.Now().UnixNano())
	}
	if input.Timestamp == 0 {
		input.Timestamp = time.Now().Unix()
	}

	// Validate input first
	if err := t.Validate(input); err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  err,
		}, NewToolError(t.name, "validate", err)
	}

	// Check if executor is set
	if t.executor == nil {
		err := fmt.Errorf("no executor set for tool %s", t.name)
		return ToolOutput{
			CallID: input.CallID,
			Error:  err,
		}, NewToolError(t.name, "execute", err)
	}

	// Build middleware chain
	executor := t.executor
	for i := len(t.middleware) - 1; i >= 0; i-- {
		executor = t.middleware[i](executor)
	}

	// Execute with timing
	start := time.Now()
	output, err := executor(ctx, input)
	output.ExecutionTime = time.Since(start)
	output.CallID = input.CallID

	if err != nil {
		output.Error = err
		return output, NewToolError(t.name, "execute", err)
	}

	return output, nil
}

// Validate checks if the input is valid for this tool
func (t *BaseTool) Validate(input ToolInput) error {
	// Use custom validator if set
	if t.validator != nil {
		if err := t.validator(input); err != nil {
			return err
		}
	}

	// For function tools, validate against schema
	if t.toolType == ToolTypeFunction && t.schema != nil {
		if input.Arguments == nil {
			return NewValidationError("arguments", "arguments required for function tool", nil)
		}
		if err := t.schema.ValidateInput(input.Arguments); err != nil {
			return err
		}
	}

	// For custom tools, ensure raw input is provided
	if t.toolType == ToolTypeCustom {
		if input.RawInput == "" {
			return NewValidationError("rawInput", "raw input required for custom tool", nil)
		}
	}

	return nil
}

// ToOpenAIFormat converts the tool to OpenAI API format
func (t *BaseTool) ToOpenAIFormat() map[string]any {
	switch t.toolType {
	case ToolTypeFunction:
		// Responses API format: flat structure with type, name, description, parameters
		result := map[string]any{
			"type":        "function",
			"name":        t.name,
			"description": t.description,
		}

		// Only enable strict mode if all properties are required
		// (OpenAI strict mode requires all properties to be in the required array)
		if t.schema != nil {
			result["parameters"] = t.schema.ToOpenAIFormat()

			// Check if all properties are required
			allRequired := len(t.schema.Properties) == len(t.schema.Required)
			if allRequired {
				result["strict"] = true
			}
		}

		return result

	case ToolTypeCustom:
		return map[string]any{
			"type":        "custom",
			"name":        t.name,
			"description": t.description,
		}

	case ToolTypeWebSearchPreview:
		return map[string]any{
			"type": "web_search_preview",
		}

	default:
		return map[string]any{
			"type":        string(t.toolType),
			"name":        t.name,
			"description": t.description,
		}
	}
}
