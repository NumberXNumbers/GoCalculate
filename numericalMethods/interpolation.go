package numericalMethods

import "errors"

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
