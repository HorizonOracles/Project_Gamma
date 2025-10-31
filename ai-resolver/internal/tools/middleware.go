package tools

import (
	"context"
	"fmt"
	"log"
	"time"
)

// LoggingMiddleware logs tool execution
func LoggingMiddleware(logger *log.Logger) ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			if logger == nil {
				return next(ctx, input)
			}

			logger.Printf("[Tool] Starting execution (CallID: %s)", input.CallID)
			start := time.Now()

			output, err := next(ctx, input)

			duration := time.Since(start)
			if err != nil {
				logger.Printf("[Tool] Execution failed (CallID: %s, Duration: %v, Error: %v)",
					input.CallID, duration, err)
			} else {
				logger.Printf("[Tool] Execution completed (CallID: %s, Duration: %v)",
					input.CallID, duration)
			}

			return output, err
		}
	}
}

// TimingMiddleware adds execution timing information
func TimingMiddleware() ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			start := time.Now()
			output, err := next(ctx, input)
			output.ExecutionTime = time.Since(start)
			output.AddLog(fmt.Sprintf("Execution time: %v", output.ExecutionTime))
			return output, err
		}
	}
}

// TimeoutMiddleware enforces a timeout on tool execution
func TimeoutMiddleware(timeout time.Duration) ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			// Create context with timeout
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			// Channel to receive result
			resultChan := make(chan struct {
				output ToolOutput
				err    error
			}, 1)

			// Execute in goroutine
			go func() {
				output, err := next(ctx, input)
				resultChan <- struct {
					output ToolOutput
					err    error
				}{output, err}
			}()

			// Wait for result or timeout
			select {
			case result := <-resultChan:
				return result.output, result.err
			case <-ctx.Done():
				return ToolOutput{
					Error: ErrExecutionTimeout,
				}, ErrExecutionTimeout
			}
		}
	}
}

// RecoveryMiddleware recovers from panics during tool execution
func RecoveryMiddleware() ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (output ToolOutput, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("tool execution panic: %v", r)
					output = ToolOutput{
						Error: err,
					}
					output.AddLog(fmt.Sprintf("Recovered from panic: %v", r))
				}
			}()

			return next(ctx, input)
		}
	}
}

// RetryMiddleware retries failed tool executions
func RetryMiddleware(maxRetries int, delay time.Duration) ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			var output ToolOutput
			var err error

			for attempt := 0; attempt <= maxRetries; attempt++ {
				output, err = next(ctx, input)

				if err == nil {
					if attempt > 0 {
						output.AddLog(fmt.Sprintf("Succeeded after %d retries", attempt))
					}
					return output, nil
				}

				// Don't retry on context cancellation or validation errors
				if ctx.Err() != nil {
					return output, err
				}
				if _, ok := err.(*ValidationError); ok {
					return output, err
				}

				// Don't sleep after last attempt
				if attempt < maxRetries {
					output.AddLog(fmt.Sprintf("Retry attempt %d/%d after error: %v",
						attempt+1, maxRetries, err))
					time.Sleep(delay)
				}
			}

			return output, fmt.Errorf("failed after %d retries: %w", maxRetries, err)
		}
	}
}

// ValidationMiddleware adds additional validation logic
func ValidationMiddleware(validator ToolValidator) ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			// Run custom validation
			if err := validator(input); err != nil {
				return ToolOutput{
					Error: err,
				}, err
			}

			return next(ctx, input)
		}
	}
}

// MetricsMiddleware tracks basic metrics (simplified version without prometheus)
type MetricsCollector struct {
	ExecutionCount    map[string]int64
	ExecutionTime     map[string]time.Duration
	ErrorCount        map[string]int64
	LastExecutionTime map[string]time.Time
}

// NewMetricsCollector creates a new metrics collector
func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		ExecutionCount:    make(map[string]int64),
		ExecutionTime:     make(map[string]time.Duration),
		ErrorCount:        make(map[string]int64),
		LastExecutionTime: make(map[string]time.Time),
	}
}

// MetricsMiddleware creates middleware that collects execution metrics
func MetricsMiddleware(collector *MetricsCollector, toolName string) ToolMiddleware {
	return func(next ToolExecutor) ToolExecutor {
		return func(ctx context.Context, input ToolInput) (ToolOutput, error) {
			start := time.Now()
			output, err := next(ctx, input)
			duration := time.Since(start)

			// Update metrics
			collector.ExecutionCount[toolName]++
			collector.ExecutionTime[toolName] += duration
			collector.LastExecutionTime[toolName] = time.Now()

			if err != nil {
				collector.ErrorCount[toolName]++
			}

			return output, err
		}
	}
}

// GetAverageExecutionTime returns the average execution time for a tool
func (m *MetricsCollector) GetAverageExecutionTime(toolName string) time.Duration {
	count := m.ExecutionCount[toolName]
	if count == 0 {
		return 0
	}
	return m.ExecutionTime[toolName] / time.Duration(count)
}

// GetErrorRate returns the error rate for a tool
func (m *MetricsCollector) GetErrorRate(toolName string) float64 {
	count := m.ExecutionCount[toolName]
	if count == 0 {
		return 0
	}
	return float64(m.ErrorCount[toolName]) / float64(count)
}
