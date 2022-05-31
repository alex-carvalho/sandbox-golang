package tests

import (
	"fmt"
	"strconv"
)

func Average(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}

	average := total / float64(len(numbers))
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return result
}

func Sum(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}

	return total
}
