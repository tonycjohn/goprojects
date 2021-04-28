package main

import (
	"fmt"
)

func divideNumbers(numerator, denominator float64) (float64, error) {
	if denominator == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return numerator / denominator, nil
}
