package numericalMethods

import "errors"

// NewtonDividedDifference is for calculating the coefficients for newton's divided-difference interpolating polynomial
func NewtonDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	currentFunctionValues := make([]float64, size)

	for i := 1; i < size; i++ {
		// fmt.Println(solutionArray)
		for j := i; j < size; j++ {
			currentFunctionValues[j] = (functionValues[j] - functionValues[j-1]) / (xValues[j] - xValues[j-i])
		}

		for j := i; j < size; j++ {
			functionValues[j] = currentFunctionValues[j]
		}
	}

	return functionValues, nil
}
