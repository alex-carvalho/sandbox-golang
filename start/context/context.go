package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation demonstrates the core pattern: racing between work completion and context cancellation
// This simulates any long-running operation (DB query, HTTP request, file processing)
func slowOperation(ctx context.Context) error {
	// select allows goroutine to respond to whichever channel receives first
	select {
	case <-time.After(3 * time.Second): // Work takes 3 seconds to complete
		fmt.Println("Operation completed")
		return nil
	case <-ctx.Done(): // Context was cancelled or timed out
		return ctx.Err() // Return specific cancellation reason
	}
}

func withTimeout() {
	// Creates context that automatically cancels after 2 seconds
	// Useful for preventing operations from hanging indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to release resources, even if timeout occurs first
	
	// Since operation needs 3s but timeout is 2s, this will always timeout
	if err := slowOperation(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func withCancel() {
	// Manual cancellation - useful when you need to stop operations based on external events
	// (user clicks cancel, system shutdown, dependency failure)
	ctx, cancel := context.WithCancel(context.Background())
	
	// Simulate external event triggering cancellation
	go func() {
		time.Sleep(1 * time.Second)
		cancel() // Manually trigger cancellation after 1 second
	}()
	
	// Operation will be cancelled before it can complete (1s < 3s)
	if err := slowOperation(ctx); err != nil {
		fmt.Printf("Cancelled: %v\n", err)
	}
}

func withValue() {
	// Context values carry request-scoped data through call chains
	// Common use: user ID, request ID, authentication tokens, tracing info
	ctx := context.WithValue(context.Background(), "userID", 123)
	
	// Values flow down the call stack - child functions can access parent context values
	if userID := ctx.Value("userID"); userID != nil {
		fmt.Printf("User ID: %v\n", userID)
	}
}

func main() {
	fmt.Println("=== Timeout Example ===")
	withTimeout()
	
	fmt.Println("\n=== Cancel Example ===")
	withCancel()
	
	fmt.Println("\n=== Value Example ===")
	withValue()
}