package tools

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DateTimeTool performs date and time calculations
// All operations are implemented in pure Go
type DateTimeTool struct {
	*BaseTool
}

// NewDateTimeTool creates a new datetime tool
func NewDateTimeTool() *DateTimeTool {
	schema := &ToolSchema{
		Type: "object",
		Properties: map[string]Property{
			"operation": {
				Type:        "string",
				Description: "The operation to perform: parse, compare, time_until, time_since, is_before, is_after, current_timestamp, format",
				Enum:        []string{"parse", "compare", "time_until", "time_since", "is_before", "is_after", "current_timestamp", "format"},
			},
			"timestamp": {
				Type:        "integer",
				Description: "Unix timestamp in seconds (for operations that need a single timestamp)",
			},
			"timestamp1": {
				Type:        "integer",
				Description: "First Unix timestamp in seconds (for comparison operations)",
			},
			"timestamp2": {
				Type:        "integer",
				Description: "Second Unix timestamp in seconds (for comparison operations)",
			},
			"date_string": {
				Type:        "string",
				Description: "Date string to parse (formats: RFC3339, 2006-01-02, 2006-01-02T15:04:05)",
			},
		},
		Required: []string{"operation"},
	}

	base := NewBaseTool(
		"datetime",
		"Perform date and time calculations including parsing dates, comparing timestamps, calculating time differences, and checking if events have occurred. Works with Unix timestamps (seconds).",
		ToolTypeFunction,
		schema,
	)

	tool := &DateTimeTool{
		BaseTool: base,
	}

	base.SetExecutor(tool.execute)

	return tool
}

// execute performs the datetime calculation
func (t *DateTimeTool) execute(ctx context.Context, input ToolInput) (ToolOutput, error) {
	// Extract operation
	operation, ok := input.Arguments["operation"].(string)
	if !ok || operation == "" {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("operation is required"),
		}, fmt.Errorf("operation is required")
	}

	// Perform operation
	result, err := t.performOperation(operation, input.Arguments)
	if err != nil {
		return ToolOutput{
			CallID: input.CallID,
			Error:  fmt.Errorf("datetime operation failed: %w", err),
			Data: map[string]any{
				"operation": operation,
				"error":     err.Error(),
			},
		}, fmt.Errorf("datetime operation failed: %w", err)
	}

	return ToolOutput{
		CallID: input.CallID,
		Data:   result,
	}, nil
}

// performOperation executes the specific datetime operation
func (t *DateTimeTool) performOperation(operation string, args map[string]any) (map[string]any, error) {
	now := time.Now()

	switch operation {
	case "current_timestamp":
		return map[string]any{
			"timestamp":   now.Unix(),
			"rfc3339":     now.Format(time.RFC3339),
			"description": "Current Unix timestamp in seconds",
		}, nil

	case "parse":
		dateStr, ok := args["date_string"].(string)
		if !ok || dateStr == "" {
			return nil, fmt.Errorf("date_string is required for parse operation")
		}

		parsedTime, layout, err := t.parseDate(dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %w", err)
		}

		return map[string]any{
			"timestamp": parsedTime.Unix(),
			"rfc3339":   parsedTime.Format(time.RFC3339),
			"layout":    layout,
			"year":      parsedTime.Year(),
			"month":     int(parsedTime.Month()),
			"day":       parsedTime.Day(),
			"hour":      parsedTime.Hour(),
			"minute":    parsedTime.Minute(),
		}, nil

	case "format":
		ts, err := t.getTimestamp(args, "timestamp")
		if err != nil {
			return nil, err
		}
		t := time.Unix(ts, 0).UTC()

		return map[string]any{
			"timestamp":   ts,
			"rfc3339":     t.Format(time.RFC3339),
			"date":        t.Format("2006-01-02"),
			"time":        t.Format("15:04:05"),
			"datetime":    t.Format("2006-01-02 15:04:05"),
			"year":        t.Year(),
			"month":       int(t.Month()),
			"day":         t.Day(),
			"hour":        t.Hour(),
			"minute":      t.Minute(),
			"second":      t.Second(),
			"day_of_week": t.Weekday().String(),
		}, nil

	case "compare":
		ts1, err := t.getTimestamp(args, "timestamp1")
		if err != nil {
			return nil, err
		}
		ts2, err := t.getTimestamp(args, "timestamp2")
		if err != nil {
			return nil, err
		}

		diff := ts1 - ts2
		return map[string]any{
			"timestamp1":         ts1,
			"timestamp2":         ts2,
			"difference_seconds": diff,
			"difference_minutes": diff / 60,
			"difference_hours":   diff / 3600,
			"difference_days":    diff / 86400,
			"timestamp1_before":  ts1 < ts2,
			"timestamp1_after":   ts1 > ts2,
			"equal":              ts1 == ts2,
		}, nil

	case "time_until":
		ts, err := t.getTimestamp(args, "timestamp")
		if err != nil {
			return nil, err
		}

		diff := ts - now.Unix()
		return map[string]any{
			"target_timestamp":  ts,
			"current_timestamp": now.Unix(),
			"seconds_until":     diff,
			"minutes_until":     diff / 60,
			"hours_until":       diff / 3600,
			"days_until":        diff / 86400,
			"has_passed":        diff < 0,
			"is_future":         diff > 0,
		}, nil

	case "time_since":
		ts, err := t.getTimestamp(args, "timestamp")
		if err != nil {
			return nil, err
		}

		diff := now.Unix() - ts
		return map[string]any{
			"target_timestamp":  ts,
			"current_timestamp": now.Unix(),
			"seconds_since":     diff,
			"minutes_since":     diff / 60,
			"hours_since":       diff / 3600,
			"days_since":        diff / 86400,
			"is_past":           diff > 0,
			"is_future":         diff < 0,
		}, nil

	case "is_before":
		ts, err := t.getTimestamp(args, "timestamp")
		if err != nil {
			return nil, err
		}

		isBefore := now.Unix() < ts
		return map[string]any{
			"current_timestamp": now.Unix(),
			"target_timestamp":  ts,
			"is_before":         isBefore,
			"is_after":          !isBefore,
		}, nil

	case "is_after":
		ts, err := t.getTimestamp(args, "timestamp")
		if err != nil {
			return nil, err
		}

		isAfter := now.Unix() > ts
		return map[string]any{
			"current_timestamp": now.Unix(),
			"target_timestamp":  ts,
			"is_after":          isAfter,
			"is_before":         !isAfter,
		}, nil

	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}
}

// getTimestamp extracts a timestamp from the arguments
func (t *DateTimeTool) getTimestamp(args map[string]any, key string) (int64, error) {
	val, ok := args[key]
	if !ok {
		return 0, fmt.Errorf("%s is required", key)
	}

	switch v := val.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid timestamp format: %w", err)
		}
		return ts, nil
	default:
		return 0, fmt.Errorf("invalid timestamp type: %T", v)
	}
}

// parseDate attempts to parse a date string in multiple common formats
func (t *DateTimeTool) parseDate(dateStr string) (time.Time, string, error) {
	// Remove common timezone abbreviations and extra whitespace
	dateStr = strings.TrimSpace(dateStr)

	// Try common date formats
	formats := []string{
		time.RFC3339,          // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,      // "2006-01-02T15:04:05.999999999Z07:00"
		"2006-01-02T15:04:05", // ISO 8601 without timezone
		"2006-01-02 15:04:05", // Common datetime format
		"2006-01-02",          // Date only
		"2006/01/02",          // Date with slashes
		"01/02/2006",          // US date format
		"02/01/2006",          // European date format
		"January 2, 2006",     // Long month name
		"Jan 2, 2006",         // Short month name
		"2 January 2006",      // Day first
		"2 Jan 2006",          // Day first, short month
		time.RFC1123,          // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,         // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC822,           // "02 Jan 06 15:04 MST"
		time.RFC822Z,          // "02 Jan 06 15:04 -0700"
		time.ANSIC,            // "Mon Jan _2 15:04:05 2006"
		time.UnixDate,         // "Mon Jan _2 15:04:05 MST 2006"
	}

	var lastErr error
	for _, format := range formats {
		parsed, err := time.Parse(format, dateStr)
		if err == nil {
			return parsed, format, nil
		}
		lastErr = err
	}

	// Try parsing as Unix timestamp
	if ts, err := strconv.ParseInt(dateStr, 10, 64); err == nil {
		return time.Unix(ts, 0), "unix_timestamp", nil
	}

	return time.Time{}, "", fmt.Errorf("unable to parse date: %w", lastErr)
}
