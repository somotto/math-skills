package mathfunctions

import "math"

// Variance calculates the variance of a slice of float64 values.
func Variance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}
	return sum / float64(len(data))
}
