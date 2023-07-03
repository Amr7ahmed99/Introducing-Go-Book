package math

import "math"

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	var listLength = len(xs)
	if listLength <= 0 {
		return 0.0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(listLength)
}

// Finds the max of a series of numbers
func FindMax(args ...int) int {
	maxValue := math.MinInt
	for _, value := range args {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// Finds the min of a series of numbers
func FindMin(args ...int) int {
	minValue := math.MaxInt
	for _, value := range args {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}
