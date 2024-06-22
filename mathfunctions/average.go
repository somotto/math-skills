package mathfunctions

// Average calculates the average of a slice of float64 values
func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data)) // Divide the sum by the number of elements
}
