package main

import (
	"fmt"
	"log/slog"
	"maps"
	"slices"
)

// Go 1.18+ Generics
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Go 1.18+ Generic constraint
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Numeric](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	// Go 1.21+ slices package
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Original:", numbers)
	
	// Sort using slices package
	sorted := slices.Clone(numbers)
	slices.Sort(sorted)
	fmt.Println("Sorted:", sorted)
	
	// Find max using slices package
	fmt.Println("Max:", slices.Max(numbers))
	
	// Check if contains
	fmt.Println("Contains 5:", slices.Contains(numbers, 5))
	
	// Go 1.21+ maps package
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	
	// Clone map
	cloned := maps.Clone(m1)
	fmt.Println("Cloned map:", cloned)
	
	// Get all keys and values
	keys := make([]string, 0, len(m1))
	for k := range maps.Keys(m1) {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	fmt.Println("Keys:", keys)
	
	values := make([]int, 0, len(m1))
	for v := range maps.Values(m1) {
		values = append(values, v)
	}
	slices.Sort(values)
	fmt.Println("Values:", values)
	
	// Generics example
	strings := []string{"hello", "world", "go"}
	lengths := Map(strings, func(s string) int { return len(s) })
	fmt.Println("String lengths:", lengths)
	
	// Generic sum
	intSum := Sum([]int{1, 2, 3, 4, 5})
	floatSum := Sum([]float64{1.1, 2.2, 3.3})
	fmt.Printf("Int sum: %d, Float sum: %.1f\n", intSum, floatSum)
	
	// Go 1.21+ structured logging with slog
	logger := slog.Default()
	logger.Info("Application started",
		"version", "1.0.0",
		"numbers_processed", len(numbers),
	)
	
	// Go 1.22+ range over integers
	fmt.Print("Range over 5: ")
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}