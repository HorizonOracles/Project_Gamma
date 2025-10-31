package tools

import (
	"testing"
)

// TestToolSchemaValidation tests schema validation
func TestToolSchemaValidation(t *testing.T) {
	t.Run("valid schema", func(t *testing.T) {
		schema := &ToolSchema{
			Type: "object",
			Properties: map[string]Property{
				"name": {
					Type:        "string",
					Description: "Name parameter",
				},
				"age": {
					Type:        "integer",
					Description: "Age parameter",
				},
			},
			Required: []string{"name"},
		}

		err := schema.Validate()
		if err != nil {
			t.Errorf("expected valid schema, got error: %v", err)
		}
	})

	t.Run("invalid schema type", func(t *testing.T) {
		schema := &ToolSchema{
			Type:       "string", // Should be "object"
			Properties: map[string]Property{},
		}

		err := schema.Validate()
		if err == nil {
			t.Error("expected error for invalid schema type")
		}
	})

	t.Run("required field not in properties", func(t *testing.T) {
		schema := &ToolSchema{
			Type: "object",
			Properties: map[string]Property{
				"name": {
					Type:        "string",
					Description: "Name parameter",
				},
			},
			Required: []string{"age"}, // age not in properties
		}

		err := schema.Validate()
		if err == nil {
			t.Error("expected error for required field not in properties")
		}
	})

	t.Run("nil schema", func(t *testing.T) {
		var schema *ToolSchema
		err := schema.Validate()
		if err != nil {
			t.Errorf("expected nil schema to be valid, got: %v", err)
		}
	})
}

// TestPropertyValidation tests property validation
func TestPropertyValidation(t *testing.T) {
	tests := []struct {
		name      string
		prop      Property
		wantError bool
	}{
		{
			name: "valid string property",
			prop: Property{
				Type:        "string",
				Description: "A string",
			},
			wantError: false,
		},
		{
			name: "valid number property",
			prop: Property{
				Type:        "number",
				Description: "A number",
			},
			wantError: false,
		},
		{
			name: "valid array property",
			prop: Property{
				Type:        "array",
				Description: "An array",
				Items: &Property{
					Type: "string",
				},
			},
			wantError: false,
		},
		{
			name: "invalid property type",
			prop: Property{
				Type:        "invalid",
				Description: "Invalid type",
			},
			wantError: true,
		},
		{
			name: "array without items",
			prop: Property{
				Type:        "array",
				Description: "Array without items",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prop.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestSchemaValidateInput tests input validation against schema
func TestSchemaValidateInput(t *testing.T) {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"name": {
				Type:        "string",
				Description: "Name",
			},
			"age": {
				Type:        "integer",
				Description: "Age",
			},
			"active": {
				Type:        "boolean",
				Description: "Active status",
			},
			"tags": {
				Type:        "array",
				Description: "Tags",
				Items: &Property{
					Type: "string",
				},
			},
		},
		Required: []string{"name", "age"},
	}

	t.Run("valid input", func(t *testing.T) {
		input := map[string]any{
			"name":   "John",
			"age":    30,
			"active": true,
			"tags":   []any{"tag1", "tag2"},
		}

		err := schema.ValidateInput(input)
		if err != nil {
			t.Errorf("expected valid input, got error: %v", err)
		}
	})

	t.Run("missing required field", func(t *testing.T) {
		input := map[string]any{
			"name": "John",
			// Missing "age"
		}

		err := schema.ValidateInput(input)
		if err == nil {
			t.Error("expected error for missing required field")
		}
	})

	t.Run("wrong type", func(t *testing.T) {
		input := map[string]any{
			"name": "John",
			"age":  "thirty", // Should be integer
		}

		err := schema.ValidateInput(input)
		if err == nil {
			t.Error("expected error for wrong type")
		}
	})

	t.Run("unknown field allowed", func(t *testing.T) {
		input := map[string]any{
			"name":    "John",
			"age":     30,
			"unknown": "value", // Unknown field
		}

		err := schema.ValidateInput(input)
		if err != nil {
			t.Errorf("expected unknown fields to be allowed, got error: %v", err)
		}
	})
}

// TestPropertyValidateValue tests value validation
func TestPropertyValidateValue(t *testing.T) {
	tests := []struct {
		name      string
		prop      Property
		value     any
		wantError bool
	}{
		{
			name:      "valid string",
			prop:      Property{Type: "string"},
			value:     "hello",
			wantError: false,
		},
		{
			name:      "invalid string",
			prop:      Property{Type: "string"},
			value:     123,
			wantError: true,
		},
		{
			name:      "valid number",
			prop:      Property{Type: "number"},
			value:     42.5,
			wantError: false,
		},
		{
			name:      "valid integer",
			prop:      Property{Type: "integer"},
			value:     42,
			wantError: false,
		},
		{
			name:      "invalid integer (float)",
			prop:      Property{Type: "integer"},
			value:     42.5,
			wantError: true,
		},
		{
			name:      "valid boolean",
			prop:      Property{Type: "boolean"},
			value:     true,
			wantError: false,
		},
		{
			name:      "invalid boolean",
			prop:      Property{Type: "boolean"},
			value:     "true",
			wantError: true,
		},
		{
			name: "valid array",
			prop: Property{
				Type:  "array",
				Items: &Property{Type: "string"},
			},
			value:     []any{"a", "b", "c"},
			wantError: false,
		},
		{
			name: "invalid array",
			prop: Property{
				Type:  "array",
				Items: &Property{Type: "string"},
			},
			value:     "not an array",
			wantError: true,
		},
		{
			name:      "valid enum",
			prop:      Property{Type: "string", Enum: []string{"red", "green", "blue"}},
			value:     "red",
			wantError: false,
		},
		{
			name:      "invalid enum",
			prop:      Property{Type: "string", Enum: []string{"red", "green", "blue"}},
			value:     "yellow",
			wantError: true,
		},
		{
			name:      "nil value",
			prop:      Property{Type: "string"},
			value:     nil,
			wantError: false, // nil is valid for optional fields
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prop.ValidateValue("test", tt.value)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateValue() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestSchemaToOpenAIFormat tests schema conversion to OpenAI format
func TestSchemaToOpenAIFormat(t *testing.T) {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"query": {
				Type:        "string",
				Description: "Search query",
			},
			"limit": {
				Type:        "integer",
				Description: "Result limit",
				Default:     10,
			},
		},
		Required: []string{"query"},
	}

	format := schema.ToOpenAIFormat()

	if format["type"] != "object" {
		t.Errorf("expected type 'object', got %v", format["type"])
	}

	props, ok := format["properties"].(map[string]any)
	if !ok {
		t.Fatal("expected properties to be map[string]any")
	}

	if len(props) != 2 {
		t.Errorf("expected 2 properties, got %d", len(props))
	}

	required, ok := format["required"].([]string)
	if !ok {
		t.Fatal("expected required to be []string")
	}

	if len(required) != 1 || required[0] != "query" {
		t.Errorf("expected required=['query'], got %v", required)
	}
}

// TestPropertyToOpenAIFormat tests property conversion to OpenAI format
func TestPropertyToOpenAIFormat(t *testing.T) {
	prop := Property{
		Type:        "string",
		Description: "A test property",
		Enum:        []string{"a", "b", "c"},
		Default:     "a",
	}

	format := prop.ToOpenAIFormat()

	if format["type"] != "string" {
		t.Errorf("expected type 'string', got %v", format["type"])
	}

	if format["description"] != "A test property" {
		t.Errorf("expected description, got %v", format["description"])
	}

	enum, ok := format["enum"].([]string)
	if !ok || len(enum) != 3 {
		t.Errorf("expected enum with 3 values, got %v", format["enum"])
	}

	if format["default"] != "a" {
		t.Errorf("expected default 'a', got %v", format["default"])
	}
}

// TestNilSchemaToOpenAIFormat tests nil schema conversion
func TestNilSchemaToOpenAIFormat(t *testing.T) {
	var schema *ToolSchema
	format := schema.ToOpenAIFormat()

	if format != nil {
		t.Errorf("expected nil format for nil schema, got %v", format)
	}
}
