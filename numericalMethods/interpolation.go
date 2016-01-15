package numericalMethods

import "errors"

// NewtonForwardDividedDifference is for calculating the coefficients for newton's forward divided-difference interpolating polynomial
func NewtonForwardDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
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

// NewtonDividedDifference is for calculating the coefficients for newton's divided-difference interpolating polynomial
func NewtonDividedDifference(xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	tableValues := make([][]float64, size)

	for i := 0; i < size; i++ {
		tableValues[i] = make([]float64, i+1)
		tableValues[i][0] = functionValues[i]
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= i; j++ {
			tableValues[i][j] = (tableValues[i][j-1] - tableValues[i-1][j-1]) / (xValues[i] - xValues[i-j])
		}
	}

	return tableValues, nil
}

// NevilleIterated is for determining the table values of neville iterated interpolation
func NevilleIterated(valueToApprox float64, xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	tableValues := make([][]float64, size)

	for i := 0; i < size; i++ {
		tableValues[i] = make([]float64, i+1)
		tableValues[i][0] = functionValues[i]
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= i; j++ {
			tableValues[i][j] = ((valueToApprox-xValues[i-j])*tableValues[i][j-1] - (valueToApprox-xValues[i])*tableValues[i-1][j-1]) / (xValues[i] - xValues[i-j])
		}
	}

	return tableValues, nil
}
