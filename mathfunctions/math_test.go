package mathfunctions

import (
	"math"
	"os"
	"testing"
)

func TestReadData(t *testing.T) {
	// Create a temporary file with test data
	testData := []byte("10\n20\n30\n40\n50\n")
	tmpfile, err := os.CreateTemp("", "test_data_*.txt")
	if err != nil {
		t.Fatalf("error creating temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	if _, err := tmpfile.Write(testData); err != nil {
		tmpfile.Close()
		t.Fatalf("error writing to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("error closing temporary file: %v", err)
	}

	// Test reading data from the temporary file
	data, err := ReadData(tmpfile.Name())
	if err != nil {
		t.Fatalf("error reading data: %v", err)
	}

	expectedData := []float64{10, 20, 30, 40, 50}
	for i, v := range data {
		if v != expectedData[i] {
			t.Errorf("expected: %f, got: %f", expectedData[i], v)
		}
	}
}

func TestAverage(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	expected := 30.0 // (10 + 20 + 30 + 40 + 50) / 5
	avg := Average(data)
	if avg != expected {
		t.Errorf("expected: %f, got: %f", expected, avg)
	}
}

func TestMedian(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	expected := 30.0 // median of {10, 20, 30, 40, 50} is 30
	median := Median(data)
	if median != expected {
		t.Errorf("expected: %f, got: %f", expected, median)
	}
}

func TestVariance(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	mean := Average(data)
	expected := 200.0 // variance of {10, 20, 30, 40, 50} is 200
	variance := Variance(data, mean)
	if variance != expected {
		t.Errorf("expected: %f, got: %f", expected, variance)
	}
}

func TestStandardDeviation(t *testing.T) {
	variance := 200.0
	expected := math.Sqrt(200) // standard deviation of variance 200 is sqrt(200)
	stdDev := StandardDeviation(variance)
	if stdDev != expected {
		t.Errorf("expected: %f, got: %f", expected, stdDev)
	}
}
