package tools

import (
	"encoding/json"
	"fmt"
)

// ToolSchema defines the JSON schema for tool parameters
type ToolSchema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
}

// Property represents a schema property
type Property struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Enum        []string  `json:"enum,omitempty"`
	Items       *Property `json:"items,omitempty"`
	Default     any       `json:"default,omitempty"`
}

// Validate checks if the schema is valid
func (s *ToolSchema) Validate() error {
	if s == nil {
		return nil // nil schema is valid (for custom tools)
	}

	if s.Type != "object" {
		return NewValidationError("type", "schema type must be 'object'", s.Type)
	}

	// Validate required fields exist in properties
	for _, req := range s.Required {
		if _, exists := s.Properties[req]; !exists {
			return NewValidationError("required", fmt.Sprintf("required field '%s' not found in properties", req), req)
		}
	}

	// Validate each property
	for name, prop := range s.Properties {
		if err := prop.Validate(); err != nil {
			return NewValidationError(name, "invalid property", err)
		}
	}

	return nil
}

// Validate checks if a property definition is valid
func (p *Property) Validate() error {
	validTypes := map[string]bool{
		"string":  true,
		"number":  true,
		"integer": true,
		"boolean": true,
		"array":   true,
		"object":  true,
	}

	if !validTypes[p.Type] {
		return NewValidationError("type", fmt.Sprintf("invalid property type '%s'", p.Type), p.Type)
	}

	// If type is array, items must be defined
	if p.Type == "array" && p.Items == nil {
		return NewValidationError("items", "array type must have items defined", nil)
	}

	// Validate nested items
	if p.Items != nil {
		if err := p.Items.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ValidateInput validates input data against the schema
func (s *ToolSchema) ValidateInput(input map[string]any) error {
	if s == nil {
		return nil // No schema means no validation
	}

	// Check required fields
	for _, req := range s.Required {
		if _, exists := input[req]; !exists {
			return NewValidationError(req, "required field missing", nil)
		}
	}

	// Validate each input field
	for key, value := range input {
		prop, exists := s.Properties[key]
		if !exists {
			// Unknown field - we can either ignore or reject
			// For now, we'll be lenient and ignore unknown fields
			continue
		}

		if err := prop.ValidateValue(key, value); err != nil {
			return err
		}
	}

	return nil
}

// ValidateValue validates a value against this property definition
func (p *Property) ValidateValue(name string, value any) error {
	if value == nil {
		return nil // nil is valid for optional fields
	}

	switch p.Type {
	case "string":
		if _, ok := value.(string); !ok {
			return NewValidationError(name, fmt.Sprintf("expected string, got %T", value), value)
		}
		// Validate enum if present
		if len(p.Enum) > 0 {
			str := value.(string)
			valid := false
			for _, enumVal := range p.Enum {
				if str == enumVal {
					valid = true
					break
				}
			}
			if !valid {
				return NewValidationError(name, fmt.Sprintf("value must be one of %v", p.Enum), value)
			}
		}

	case "number":
		switch value.(type) {
		case float64, float32, int, int32, int64, json.Number:
			// Valid number types
		default:
			return NewValidationError(name, fmt.Sprintf("expected number, got %T", value), value)
		}

	case "integer":
		switch value.(type) {
		case int, int32, int64, uint, uint32, uint64:
			// Valid integer types
		case float64:
			// Check if it's a whole number
			f := value.(float64)
			if f != float64(int64(f)) {
				return NewValidationError(name, "expected integer, got float", value)
			}
		default:
			return NewValidationError(name, fmt.Sprintf("expected integer, got %T", value), value)
		}

	case "boolean":
		if _, ok := value.(bool); !ok {
			return NewValidationError(name, fmt.Sprintf("expected boolean, got %T", value), value)
		}

	case "array":
		arr, ok := value.([]any)
		if !ok {
			return NewValidationError(name, fmt.Sprintf("expected array, got %T", value), value)
		}
		// Validate each item if items schema is defined
		if p.Items != nil {
			for i, item := range arr {
				if err := p.Items.ValidateValue(fmt.Sprintf("%s[%d]", name, i), item); err != nil {
					return err
				}
			}
		}

	case "object":
		if _, ok := value.(map[string]any); !ok {
			return NewValidationError(name, fmt.Sprintf("expected object, got %T", value), value)
		}
	}

	return nil
}

// ToOpenAIFormat converts the schema to OpenAI function tool format
func (s *ToolSchema) ToOpenAIFormat() map[string]any {
	if s == nil {
		return nil
	}

	result := map[string]any{
		"type":       s.Type,
		"properties": make(map[string]any),
	}

	for name, prop := range s.Properties {
		result["properties"].(map[string]any)[name] = prop.ToOpenAIFormat()
	}

	if len(s.Required) > 0 {
		result["required"] = s.Required
	}

	return result
}

// ToOpenAIFormat converts the property to OpenAI format
func (p *Property) ToOpenAIFormat() map[string]any {
	result := map[string]any{
		"type":        p.Type,
		"description": p.Description,
	}

	if len(p.Enum) > 0 {
		result["enum"] = p.Enum
	}

	if p.Items != nil {
		result["items"] = p.Items.ToOpenAIFormat()
	}

	if p.Default != nil {
		result["default"] = p.Default
	}

	return result
}
