// package numericalMethods

package main

import "errors"
import "log"
import "fmt"

// NewtonDividedDifference is for calculating the coefficients for newton's divided-difference interpolating polynomial
func NewtonDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	previousValue := float64(0)
	tempValue := float64(0)

	for i := 1; i < size; i++ {
		previousValue = functionValues[i-1]
		for j := i; j < size; j++ {
			tempValue = functionValues[j]
			functionValues[j] = (functionValues[j] - previousValue) / (xValues[j] - xValues[j-i])
			previousValue = tempValue
		}
	}

	return functionValues, nil
}

func main() {
	xValues := []float64{1, 1.3, 1.6, 1.9, 2.2}
	functionValues := []float64{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623}
	results, err := NewtonDividedDifference(xValues, functionValues)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(results)
}
