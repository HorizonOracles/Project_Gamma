package tools

import (
	"context"
	"fmt"
	"sync"
)

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

	// Count returns the number of registered tools
	Count() int

	// Has checks if a tool is registered
	Has(name string) bool
}

// DefaultRegistry is a thread-safe implementation of Registry
type DefaultRegistry struct {
	mu    sync.RWMutex
	tools map[string]Tool
}

// NewRegistry creates a new DefaultRegistry
func NewRegistry() Registry {
	return &DefaultRegistry{
		tools: make(map[string]Tool),
	}
}

// Register adds a new tool to the registry
func (r *DefaultRegistry) Register(tool Tool) error {
	if tool == nil {
		return fmt.Errorf("cannot register nil tool")
	}

	name := tool.Name()
	if name == "" {
		return ErrInvalidToolName
	}

	// Validate tool schema if present
	if schema := tool.Schema(); schema != nil {
		if err := schema.Validate(); err != nil {
			return NewToolError(name, "register", fmt.Errorf("%w: %v", ErrInvalidSchema, err))
		}
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tools[name]; exists {
		return NewToolError(name, "register", ErrToolAlreadyRegistered)
	}

	r.tools[name] = tool
	return nil
}

// Unregister removes a tool from the registry
func (r *DefaultRegistry) Unregister(name string) error {
	if name == "" {
		return ErrInvalidToolName
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tools[name]; !exists {
		return NewToolError(name, "unregister", ErrToolNotFound)
	}

	delete(r.tools, name)
	return nil
}

// Get retrieves a tool by name
func (r *DefaultRegistry) Get(name string) (Tool, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tool, exists := r.tools[name]
	return tool, exists
}

// List returns all registered tools
func (r *DefaultRegistry) List() []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// ListByType returns tools of a specific type
func (r *DefaultRegistry) ListByType(toolType ToolType) []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tools := make([]Tool, 0)
	for _, tool := range r.tools {
		if tool.Type() == toolType {
			tools = append(tools, tool)
		}
	}
	return tools
}

// ExecuteTool runs a tool by name with the provided input
func (r *DefaultRegistry) ExecuteTool(ctx context.Context, name string, input ToolInput) (ToolOutput, error) {
	tool, exists := r.Get(name)
	if !exists {
		return ToolOutput{
			Error: ErrToolNotFound,
		}, NewToolError(name, "execute", ErrToolNotFound)
	}

	return tool.Execute(ctx, input)
}

// ToOpenAIFormat converts registered tools to OpenAI API format
func (r *DefaultRegistry) ToOpenAIFormat() []map[string]any {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]map[string]any, 0, len(r.tools))
	for _, tool := range r.tools {
		// Use tool's ToOpenAIFormat method if available
		if bt, ok := tool.(*BaseTool); ok {
			result = append(result, bt.ToOpenAIFormat())
		} else {
			// Fallback for custom Tool implementations
			format := map[string]any{
				"type": string(tool.Type()),
				"name": tool.Name(),
			}
			if tool.Type() == ToolTypeFunction {
				format["function"] = map[string]any{
					"name":        tool.Name(),
					"description": tool.Description(),
				}
				if schema := tool.Schema(); schema != nil {
					format["function"].(map[string]any)["parameters"] = schema.ToOpenAIFormat()
				}
			} else if tool.Type() == ToolTypeCustom {
				format["description"] = tool.Description()
			}
			result = append(result, format)
		}
	}
	return result
}

// Count returns the number of registered tools
func (r *DefaultRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.tools)
}

// Has checks if a tool is registered
func (r *DefaultRegistry) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.tools[name]
	return exists
}

// Clear removes all tools from the registry (useful for testing)
func (r *DefaultRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tools = make(map[string]Tool)
}
