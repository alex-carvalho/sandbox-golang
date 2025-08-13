package tests

import (
	"fmt"
	"strconv"
)

func Average(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	var total float64
	for _, number := range numbers {
		total += number
	}

	average := total / float64(len(numbers))
	// Use math.Round for better performance
	return float64(int(average*100+0.5)) / 100
}

func Sum(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}

	return total
}
