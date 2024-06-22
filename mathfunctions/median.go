package mathfunctions

import "sort"

// Median calculates the median of a slice of float64 values.
func Median(data []float64) float64 {
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	n := len(sortedData)
	if n%2 == 0 {
		return (sortedData[n/2-1] + sortedData[n/2]) / 2
	}
	return sortedData[n/2]
}
