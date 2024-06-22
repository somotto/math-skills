package mathfunctions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ReadData reads numeric data from a file specified by the filePath parameter
// and returns a slice of float64 and an error.

func ReadData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := strings.TrimSpace(scanner.Text())
		if value == "" {
			continue
		}
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		if number > math.MaxInt {
			fmt.Println("There is an overflow")
			os.Exit(0)
		}
		data = append(data, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no valid numeric values found in the data.txt")
	}

	return data, nil
}
