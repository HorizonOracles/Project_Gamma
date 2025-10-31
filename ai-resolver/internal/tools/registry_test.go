package tools

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestToolTypes tests the tool type constants
func TestToolTypes(t *testing.T) {
	tests := []struct {
		name     string
		toolType ToolType
		expected string
	}{
		{"function type", ToolTypeFunction, "function"},
		{"custom type", ToolTypeCustom, "custom"},
		{"web search type", ToolTypeWebSearchPreview, "web_search_preview"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.toolType) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, string(tt.toolType))
			}
		})
	}
}

// TestToolInput tests the ToolInput structure
func TestToolInput(t *testing.T) {
	input := ToolInput{
		Arguments: map[string]any{
			"key1": "value1",
			"key2": 42,
		},
		RawInput:  "test input",
		CallID:    "call_123",
		Timestamp: time.Now().Unix(),
	}

	if input.Arguments["key1"] != "value1" {
		t.Errorf("expected key1=value1, got %v", input.Arguments["key1"])
	}
	if input.RawInput != "test input" {
		t.Errorf("expected 'test input', got %s", input.RawInput)
	}
}

// TestToolOutput tests the ToolOutput structure
func TestToolOutput(t *testing.T) {
	output := ToolOutput{
		Data:          map[string]string{"result": "success"},
		CallID:        "call_123",
		ExecutionTime: 100 * time.Millisecond,
	}

	// Test AddLog
	output.AddLog("log1")
	output.AddLog("log2")

	if len(output.Logs) != 2 {
		t.Errorf("expected 2 logs, got %d", len(output.Logs))
	}

	// Test HasError
	if output.HasError() {
		t.Error("expected no error")
	}

	output.Error = errors.New("test error")
	if !output.HasError() {
		t.Error("expected error")
	}
}

// TestNewRegistry tests registry creation
func TestNewRegistry(t *testing.T) {
	registry := NewRegistry()
	if registry == nil {
		t.Fatal("expected non-nil registry")
	}

	if registry.Count() != 0 {
		t.Errorf("expected empty registry, got %d tools", registry.Count())
	}
}

// TestRegisterTool tests tool registration
func TestRegisterTool(t *testing.T) {
	registry := NewRegistry()

	// Create a simple tool
	tool := NewBaseTool("test_tool", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	// Register tool
	err := registry.Register(tool)
	if err != nil {
		t.Fatalf("failed to register tool: %v", err)
	}

	if registry.Count() != 1 {
		t.Errorf("expected 1 tool, got %d", registry.Count())
	}

	// Try to register the same tool again
	err = registry.Register(tool)
	if err == nil {
		t.Error("expected error when registering duplicate tool")
	}
	if !errors.Is(err, ErrToolAlreadyRegistered) {
		t.Errorf("expected ErrToolAlreadyRegistered, got %v", err)
	}
}

// TestRegisterNilTool tests registering a nil tool
func TestRegisterNilTool(t *testing.T) {
	registry := NewRegistry()
	err := registry.Register(nil)
	if err == nil {
		t.Error("expected error when registering nil tool")
	}
}

// TestGetTool tests retrieving tools from registry
func TestGetTool(t *testing.T) {
	registry := NewRegistry()
	tool := NewBaseTool("test_tool", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	registry.Register(tool)

	// Get existing tool
	retrieved, exists := registry.Get("test_tool")
	if !exists {
		t.Error("expected tool to exist")
	}
	if retrieved.Name() != "test_tool" {
		t.Errorf("expected test_tool, got %s", retrieved.Name())
	}

	// Get non-existing tool
	_, exists = registry.Get("nonexistent")
	if exists {
		t.Error("expected tool to not exist")
	}
}

// TestUnregisterTool tests removing tools from registry
func TestUnregisterTool(t *testing.T) {
	registry := NewRegistry()
	tool := NewBaseTool("test_tool", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	registry.Register(tool)

	// Unregister existing tool
	err := registry.Unregister("test_tool")
	if err != nil {
		t.Errorf("failed to unregister tool: %v", err)
	}

	if registry.Count() != 0 {
		t.Errorf("expected 0 tools, got %d", registry.Count())
	}

	// Unregister non-existing tool
	err = registry.Unregister("nonexistent")
	if err == nil {
		t.Error("expected error when unregistering non-existing tool")
	}
}

// TestListTools tests listing all tools
func TestListTools(t *testing.T) {
	registry := NewRegistry()

	tool1 := NewBaseTool("tool1", "Tool 1", ToolTypeFunction, nil)
	tool1.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	tool2 := NewBaseTool("tool2", "Tool 2", ToolTypeCustom, nil)
	tool2.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	registry.Register(tool1)
	registry.Register(tool2)

	tools := registry.List()
	if len(tools) != 2 {
		t.Errorf("expected 2 tools, got %d", len(tools))
	}
}

// TestListToolsByType tests filtering tools by type
func TestListToolsByType(t *testing.T) {
	registry := NewRegistry()

	tool1 := NewBaseTool("tool1", "Tool 1", ToolTypeFunction, nil)
	tool1.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	tool2 := NewBaseTool("tool2", "Tool 2", ToolTypeCustom, nil)
	tool2.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	tool3 := NewBaseTool("tool3", "Tool 3", ToolTypeFunction, nil)
	tool3.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	registry.Register(tool1)
	registry.Register(tool2)
	registry.Register(tool3)

	functionTools := registry.ListByType(ToolTypeFunction)
	if len(functionTools) != 2 {
		t.Errorf("expected 2 function tools, got %d", len(functionTools))
	}

	customTools := registry.ListByType(ToolTypeCustom)
	if len(customTools) != 1 {
		t.Errorf("expected 1 custom tool, got %d", len(customTools))
	}
}

// TestExecuteTool tests tool execution through registry
func TestExecuteTool(t *testing.T) {
	registry := NewRegistry()
	ctx := context.Background()

	tool := NewBaseTool("test_tool", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: map[string]string{"result": "success"}}, nil
	})

	registry.Register(tool)

	// Execute existing tool
	output, err := registry.ExecuteTool(ctx, "test_tool", ToolInput{
		Arguments: map[string]any{},
	})

	if err != nil {
		t.Errorf("failed to execute tool: %v", err)
	}

	data, ok := output.Data.(map[string]string)
	if !ok {
		t.Fatal("expected map[string]string data")
	}
	if data["result"] != "success" {
		t.Errorf("expected success, got %s", data["result"])
	}

	// Execute non-existing tool
	_, err = registry.ExecuteTool(ctx, "nonexistent", ToolInput{})
	if err == nil {
		t.Error("expected error when executing non-existing tool")
	}
}

// TestBaseTool tests the BaseTool implementation
func TestBaseTool(t *testing.T) {
	tool := NewBaseTool("test", "Test tool", ToolTypeFunction, nil)

	if tool.Name() != "test" {
		t.Errorf("expected name 'test', got %s", tool.Name())
	}

	if tool.Description() != "Test tool" {
		t.Errorf("expected description 'Test tool', got %s", tool.Description())
	}

	if tool.Type() != ToolTypeFunction {
		t.Errorf("expected type function, got %s", tool.Type())
	}

	if tool.Schema() != nil {
		t.Error("expected nil schema")
	}
}

// TestBaseToolExecution tests tool execution
func TestBaseToolExecution(t *testing.T) {
	ctx := context.Background()
	called := false

	tool := NewBaseTool("test", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		called = true
		return ToolOutput{Data: "executed"}, nil
	})

	output, err := tool.Execute(ctx, ToolInput{
		Arguments: map[string]any{},
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !called {
		t.Error("executor was not called")
	}

	if output.Data != "executed" {
		t.Errorf("expected 'executed', got %v", output.Data)
	}

	if output.CallID == "" {
		t.Error("expected non-empty call ID")
	}

	if output.ExecutionTime == 0 {
		t.Error("expected non-zero execution time")
	}
}

// TestBaseToolNoExecutor tests executing a tool without an executor
func TestBaseToolNoExecutor(t *testing.T) {
	ctx := context.Background()
	tool := NewBaseTool("test", "Test tool", ToolTypeFunction, nil)

	_, err := tool.Execute(ctx, ToolInput{
		Arguments: map[string]any{},
	})

	if err == nil {
		t.Error("expected error when executing tool without executor")
	}
}

// TestBaseToolMiddleware tests middleware functionality
func TestBaseToolMiddleware(t *testing.T) {
	ctx := context.Background()
	middlewareCalled := false

	tool := NewBaseTool("test", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "executed"}, nil
	})

	// Add middleware
	tool.WithMiddleware(func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			middlewareCalled = true
			return next(ctx, input)
		}
	})

	_, err := tool.Execute(ctx, ToolInput{
		Arguments: map[string]any{},
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !middlewareCalled {
		t.Error("middleware was not called")
	}
}

// TestToOpenAIFormat tests OpenAI format conversion
func TestToOpenAIFormat(t *testing.T) {
	registry := NewRegistry()

	// Function tool with schema
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"query": {
				Type:        "string",
				Description: "Search query",
			},
		},
		Required: []string{"query"},
	}

	functionTool := NewBaseTool("search", "Search the web", ToolTypeFunction, schema)
	functionTool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	// Custom tool
	customTool := NewBaseTool("exec", "Execute code", ToolTypeCustom, nil)
	customTool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	registry.Register(functionTool)
	registry.Register(customTool)

	format := registry.ToOpenAIFormat()

	if len(format) != 2 {
		t.Errorf("expected 2 tools in format, got %d", len(format))
	}

	// Verify function tool format
	var functionFormat map[string]any
	for _, f := range format {
		if f["type"] == "function" {
			functionFormat = f
			break
		}
	}

	if functionFormat == nil {
		t.Fatal("function tool not found in format")
	}

	function, ok := functionFormat["function"].(map[string]any)
	if !ok {
		t.Fatal("expected function object")
	}

	if function["name"] != "search" {
		t.Errorf("expected name 'search', got %v", function["name"])
	}
}

// TestHasTool tests the Has method
func TestHasTool(t *testing.T) {
	registry := NewRegistry()
	tool := NewBaseTool("test", "Test tool", ToolTypeFunction, nil)
	tool.SetExecutor(func(ctx context.Context, input ToolInput) (ToolOutput, error) {
		return ToolOutput{Data: "success"}, nil
	})

	if registry.Has("test") {
		t.Error("expected false for non-existing tool")
	}

	registry.Register(tool)

	if !registry.Has("test") {
		t.Error("expected true for existing tool")
	}
}
