package tools

import (
	"context"
	"fmt"
	"math"
	"strconv"
)

// CalculatorTool performs mathematical and statistical calculations
// All operations are implemented in pure Go for safety and performance
type CalculatorTool struct {
	*BaseTool
}

// NewCalculatorTool creates a new calculator tool
func NewCalculatorTool() *CalculatorTool {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"operation": {
				Type:        "string",
				Description: "The calculation to perform: add, subtract, multiply, divide, power, sqrt, percentage, probability_multiply, probability_complement, mean, median",
				Enum:        []string{"add", "subtract", "multiply", "divide", "power", "sqrt", "percentage", "probability_multiply", "probability_complement", "mean", "median"},
			},
			"values": {
				Type:        "array",
				Description: "The numbers to calculate with (1-2 values for basic operations, any number for mean/median)",
				Items: &Property{
					Type: "number",
				},
			},
		},
		Required: []string{"operation", "values"},
	}

	base := NewBaseTool(
		"calculate",
		"Perform mathematical and statistical calculations including basic arithmetic, probability calculations, and statistical operations. Returns a numeric result.",
		ToolTypeFunction,
		schema,
	)

	tool := &CalculatorTool{
		BaseTool: base,
	}

	base.SetExecutor(tool.execute)

	return tool
}

// execute performs the calculation in pure Go
func (t *CalculatorTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	// Extract operation
	operation, ok := input.Arguments["operation"].(string)
	if !ok || operation == "" {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("operation is required"),
		}, fmt.Errorf("operation is required")
	}

	// Extract values
	valuesInterface, ok := input.Arguments["values"]
	if !ok {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("values are required"),
		}, fmt.Errorf("values are required")
	}

	// Convert values to float64 slice
	var values []float64
	switch v := valuesInterface.(type) {
	case []interface{}:
		for i, val := range v {
			floatVal, err := toFloat64(val)
			if err != nil {
				return ToolOutput{
					CallID: input.CallID,
					Error:  fmt.Errorf("invalid value at index %d: %w", i, err),
				}, fmt.Errorf("invalid value at index %d: %w", i, err)
			}
			values = append(values, floatVal)
		}
	case []float64:
		values = v
	default:
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("values must be an array of numbers"),
		}, fmt.Errorf("values must be an array of numbers")
	}

	// Perform calculation
	result, err := t.calculate(operation, values)
	if err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("calculation failed: %w", err),
			Data: map[string]any{
				"operation": operation,
				"values":    values,
				"error":     err.Error(),
			},
		}, fmt.Errorf("calculation failed: %w", err)
	}

	return ToolOutput{
		CallID: input.CallID,
		Data: map[string]any{
			"operation": operation,
			"values":    values,
			"result":    result,
		},
	}, nil
}

// calculate performs the actual calculation based on operation
func (t *CalculatorTool) calculate(operation string, values []float64) (float64, error) {
	switch operation {
	case "add":
		if len(values) < 2 {
			return 0, fmt.Errorf("add requires at least 2 values")
		}
		result := values[0]
		for i := 1; i < len(values); i++ {
			result += values[i]
		}
		return result, nil

	case "subtract":
		if len(values) != 2 {
			return 0, fmt.Errorf("subtract requires exactly 2 values")
		}
		return values[0] - values[1], nil

	case "multiply":
		if len(values) < 2 {
			return 0, fmt.Errorf("multiply requires at least 2 values")
		}
		result := values[0]
		for i := 1; i < len(values); i++ {
			result *= values[i]
		}
		return result, nil

	case "divide":
		if len(values) != 2 {
			return 0, fmt.Errorf("divide requires exactly 2 values")
		}
		if values[1] == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return values[0] / values[1], nil

	case "power":
		if len(values) != 2 {
			return 0, fmt.Errorf("power requires exactly 2 values (base, exponent)")
		}
		return math.Pow(values[0], values[1]), nil

	case "sqrt":
		if len(values) != 1 {
			return 0, fmt.Errorf("sqrt requires exactly 1 value")
		}
		if values[0] < 0 {
			return 0, fmt.Errorf("cannot take square root of negative number")
		}
		return math.Sqrt(values[0]), nil

	case "percentage":
		if len(values) != 2 {
			return 0, fmt.Errorf("percentage requires exactly 2 values (part, whole)")
		}
		if values[1] == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return (values[0] / values[1]) * 100, nil

	case "probability_multiply":
		// Multiply probabilities (for independent events)
		if len(values) < 2 {
			return 0, fmt.Errorf("probability_multiply requires at least 2 values")
		}
		result := values[0]
		for i := 1; i < len(values); i++ {
			if values[i] < 0 || values[i] > 1 {
				return 0, fmt.Errorf("probability values must be between 0 and 1")
			}
			result *= values[i]
		}
		return result, nil

	case "probability_complement":
		// Calculate complement probability (1 - p)
		if len(values) != 1 {
			return 0, fmt.Errorf("probability_complement requires exactly 1 value")
		}
		if values[0] < 0 || values[0] > 1 {
			return 0, fmt.Errorf("probability must be between 0 and 1")
		}
		return 1 - values[0], nil

	case "mean":
		if len(values) == 0 {
			return 0, fmt.Errorf("mean requires at least 1 value")
		}
		sum := 0.0
		for _, v := range values {
			sum += v
		}
		return sum / float64(len(values)), nil

	case "median":
		if len(values) == 0 {
			return 0, fmt.Errorf("median requires at least 1 value")
		}
		// Sort values (simple bubble sort for small arrays)
		sorted := make([]float64, len(values))
		copy(sorted, values)
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i] > sorted[j] {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}
		// Return median
		n := len(sorted)
		if n%2 == 0 {
			return (sorted[n/2-1] + sorted[n/2]) / 2, nil
		}
		return sorted[n/2], nil

	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

// toFloat64 converts various numeric types to float64
func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse string as number: %w", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported numeric type: %T", v)
	}
}
