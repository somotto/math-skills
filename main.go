package main

import (
	"fmt"
	"os"

	"math-skills/mathfunctions"
)

// Check if the correct number of command-line arguments is provided.
// Read data from the specified file
func main() {
	if len(os.Args) != 2 || os.Args[1] != "data.txt" {
		fmt.Println("Usage: go run your_program.go data.txt")
		os.Exit(0)
	}

	filePath := os.Args[1]
	data, err := mathfunctions.ReadData(filePath)
	if err != nil {
		fmt.Println("Error reading data:", err)
		os.Exit(0)
	}

	avg := mathfunctions.Average(data)
	median := mathfunctions.Median(data)
	mean := mathfunctions.Average(data)
	variance := mathfunctions.Variance(data, mean)
	stdDev := mathfunctions.StandardDeviation(variance)

	fmt.Println("Average:", int(avg+0.5)) // rounding to the nearest integer
	fmt.Println("Median:", int(median+0.5))
	fmt.Println("Variance:", int(variance+0.5))
	fmt.Println("Standard Deviation:", int(stdDev+0.5))
}
