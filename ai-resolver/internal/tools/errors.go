package tools

import (
	"errors"
	"fmt"
)

var (
	// ErrToolNotFound is returned when a tool cannot be found in the registry
	ErrToolNotFound = errors.New("tool not found")

	// ErrToolAlreadyRegistered is returned when attempting to register a duplicate tool
	ErrToolAlreadyRegistered = errors.New("tool already registered")

	// ErrInvalidToolName is returned when a tool name is empty or invalid
	ErrInvalidToolName = errors.New("invalid tool name")

	// ErrInvalidInput is returned when tool input validation fails
	ErrInvalidInput = errors.New("invalid tool input")

	// ErrExecutionFailed is returned when tool execution fails
	ErrExecutionFailed = errors.New("tool execution failed")

	// ErrExecutionTimeout is returned when tool execution times out
	ErrExecutionTimeout = errors.New("tool execution timeout")

	// ErrUnauthorized is returned when a user lacks permission to execute a tool
	ErrUnauthorized = errors.New("unauthorized tool access")

	// ErrInvalidSchema is returned when a tool schema is invalid
	ErrInvalidSchema = errors.New("invalid tool schema")

	// ErrMissingRequired is returned when required parameters are missing
	ErrMissingRequired = errors.New("missing required parameters")
)

// ToolError wraps an error with additional context about the tool
type ToolError struct {
	ToolName string
	Op       string // Operation that failed (e.g., "execute", "validate")
	Err      error
}

func (e *ToolError) Error() string {
	if e.Op != "" {
		return fmt.Sprintf("tool %s: %s: %v", e.ToolName, e.Op, e.Err)
	}
	return fmt.Sprintf("tool %s: %v", e.ToolName, e.Err)
}

func (e *ToolError) Unwrap() error {
	return e.Err
}

// NewToolError creates a new ToolError
func NewToolError(toolName, op string, err error) *ToolError {
	return &ToolError{
		ToolName: toolName,
		Op:       op,
		Err:      err,
	}
}

// ValidationError represents a validation failure with details
type ValidationError struct {
	Field   string
	Message string
	Value   any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field %s: %s (value: %v)", e.Field, e.Message, e.Value)
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string, value any) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}
