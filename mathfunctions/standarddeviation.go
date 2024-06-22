package mathfunctions

import "math"

// StandardDeviation calculates the standard deviation given the variance.
// It returns the square root of the variance.
func StandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
