package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("file not found")
var ErrPermission = errors.New("permission denied")

func readFile(filename string) error {
	if filename == "missing.txt" {
		return ErrNotFound
	}
	if filename == "protected.txt" {
		return ErrPermission
	}
	return nil
}

func processFile(filename string) error {
	if err := readFile(filename); err != nil {
		return fmt.Errorf("processing %s failed: %w", filename, err)
	}
	return nil
}

func handleFile(filename string) error {
	if err := processFile(filename); err != nil {
		return fmt.Errorf("handle operation failed: %w", err)
	}
	return nil
}

func main() {
	files := []string{"missing.txt", "protected.txt", "valid.txt"}
	
	for _, file := range files {
		if err := handleFile(file); err != nil {
			fmt.Printf("Error with %s: %v\n", file, err)
			
			// Check for specific errors
			if errors.Is(err, ErrNotFound) {
				fmt.Println("  -> File does not exist")
			}
			if errors.Is(err, ErrPermission) {
				fmt.Println("  -> Access denied")
			}
			
			// Unwrap to get original error
			if unwrapped := errors.Unwrap(err); unwrapped != nil {
				fmt.Printf("  -> Original: %v\n", unwrapped)
			}
			
			fmt.Println()
		} else {
			fmt.Printf("%s processed successfully\n", file)
		}
	}
	
	// errors.Join (Go 1.20+)
	err1 := errors.New("first error")
	err2 := errors.New("second error")
	combined := errors.Join(err1, err2)
	fmt.Printf("Combined errors: %v\n", combined)
}