package methods

import (
	"errors"
	"math"
)

// NewtonForwardDividedDifference is for calculating the coefficients for newton's forward divided-difference interpolating polynomial
func NewtonForwardDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
	tableValues, err := NewtonDividedDifference(xValues, functionValues)

	if err != nil {
		return nil, err
	}

	solutionSet := make([]float64, len(tableValues))

	for i := 0; i < len(tableValues); i++ {
		solutionSet[i] = tableValues[i][i]
	}

	return solutionSet, nil
}

// NewtonBackwardsDividedDifference is for calculating the coefficients for newton's backwards divided-difference interpolating polynomial
func NewtonBackwardsDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
	tableValues, err := NewtonDividedDifference(xValues, functionValues)

	if err != nil {
		return nil, err
	}

	solutionSet := make([]float64, len(tableValues))

	for i := 0; i < len(tableValues); i++ {
		solutionSet[i] = tableValues[len(tableValues)-1][i]
	}

	return solutionSet, nil
}

// StirlingCenterDividedDifference is for calculating the coefficients for stirling's center divided-difference interpolating polynomial
// if xValues or functionValues has an even number of elements, the last elements will be removed.
func StirlingCenterDividedDifference(xValues []float64, functionValues []float64) ([][]float64, error) {
	if len(xValues)%2 == 0 || len(functionValues)%2 == 0 {
		xValues = xValues[:len(xValues)-1]
		functionValues = functionValues[:len(functionValues)-1]
	}

	tableValues, err := NewtonDividedDifference(xValues, functionValues)

	if err != nil {
		return nil, err
	}

	solutionSet := make([][]float64, len(tableValues))

	for i := 0; i < len(tableValues); i++ {
		if i%2 == 0 {
			solutionSet[i] = make([]float64, 1)
		} else {
			solutionSet[i] = make([]float64, 2)
		}

	}

	middle := int(math.Floor(float64(len(tableValues)) / 2))

	for i := 0; i < len(tableValues); i++ {
		if i%2 == 1 {
			solutionSet[i][0] = tableValues[middle][i]
			middle++
			solutionSet[i][1] = tableValues[middle][i]
		} else {
			solutionSet[i][0] = tableValues[middle][i]
		}
	}

	return solutionSet, nil
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

// Hermite is for determining the coefficients of the hermite interpolation polynomial
func Hermite(xValues []float64, functionValues []float64, dfunctionValues []float64) ([]float64, error) {
	size := len(xValues)

	if size != len(functionValues) || size != len(dfunctionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	valueDoubleSet := make([]float64, 2*size)

	tableValues := make([][]float64, 2*size)

	for i := 0; i < size; i++ {
		tableValues[2*i] = make([]float64, 2*i+1)
		tableValues[2*i+1] = make([]float64, 2*(i+1))
	}

	valueDoubleSet[0] = xValues[0]
	valueDoubleSet[1] = xValues[0]

	tableValues[0][0] = functionValues[0]
	tableValues[1][0] = functionValues[0]
	tableValues[1][1] = dfunctionValues[0]

	for i := 1; i < size; i++ {
		valueDoubleSet[2*i] = xValues[i]
		valueDoubleSet[2*i+1] = xValues[i]

		tableValues[2*i][0] = functionValues[i]
		tableValues[2*i+1][0] = functionValues[i]
		tableValues[2*i+1][1] = dfunctionValues[i]

		tableValues[2*i][1] = (tableValues[2*i][0] - tableValues[2*i-1][0]) / (valueDoubleSet[2*i] - valueDoubleSet[2*i-1])
	}

	for i := 2; i < 2*size; i++ {
		for j := 2; j <= i; j++ {
			tableValues[i][j] = (tableValues[i][j-1] - tableValues[i-1][j-1]) / (valueDoubleSet[i] - valueDoubleSet[i-j])
		}
	}

	solutionSet := make([]float64, 2*size)

	for i := 0; i < 2*size; i++ {
		solutionSet[i] = tableValues[i][i]
	}

	return solutionSet, nil
}

// NaturalCubicSpline is used for finding the coefficients solution set of the natural cubic spline
func NaturalCubicSpline(xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	stepLengthSet := make([]float64, size-1)

	for i := 0; i < size-1; i++ {
		stepLengthSet[i] = xValues[i+1] - xValues[i]
	}

	alpha := make([]float64, size-1)

	for i := 1; i < size-1; i++ {
		alpha[i] = 3.0*(functionValues[i+1]-functionValues[i])/stepLengthSet[i] - 3.0*(functionValues[i]-functionValues[i-1])/stepLengthSet[i-1]
	}

	solvingSetA := make([]float64, size)
	solvingSetB := make([]float64, size-1)
	solvingSetC := make([]float64, size)

	solvingSetA[0] = 1
	solvingSetB[0] = 0
	solvingSetC[0] = 0

	for i := 1; i < size-1; i++ {
		solvingSetA[i] = 2.0*(xValues[i+1]-xValues[i-1]) - stepLengthSet[i-1]*solvingSetB[i-1]
		solvingSetB[i] = stepLengthSet[i] / solvingSetA[i]
		solvingSetC[i] = (alpha[i] - stepLengthSet[i-1]*solvingSetC[i-1]) / solvingSetA[i]
	}

	solvingSetA[size-1] = 1
	solvingSetC[size-1] = 0

	var solutionSetA []float64
	solutionSetA = functionValues[:size-1]
	solutionSetB := make([]float64, size-1)
	solutionSetC := make([]float64, size)
	solutionSetD := make([]float64, size-1)

	solutionSetC[size-1] = 0

	sizeSolutionSet := len(solutionSetB)

	for i := sizeSolutionSet - 1; i >= 0; i-- {
		solutionSetC[i] = solvingSetC[i] - solvingSetB[i]*solutionSetC[i+1]
		solutionSetB[i] = (functionValues[i+1]-functionValues[i])/stepLengthSet[i] - stepLengthSet[i]*(solutionSetC[i+1]+2.0*solutionSetC[i])/3.0
		solutionSetD[i] = (solutionSetC[i+1] - solutionSetC[i]) / (3.0 * stepLengthSet[i])
	}

	solutionSetC = solutionSetC[:size-1]

	solutionTable := [][]float64{}

	solutionTable = append(solutionTable, solutionSetA)
	solutionTable = append(solutionTable, solutionSetB)
	solutionTable = append(solutionTable, solutionSetC)
	solutionTable = append(solutionTable, solutionSetD)

	return solutionTable, nil
}

// ClampedCubicSpline is for finding the coefficients solution set of the clamped cubic spline
func ClampedCubicSpline(xValues []float64, functionValues []float64, df0 float64, dfN float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	stepLengthSet := make([]float64, size-1)

	for i := 0; i < size-1; i++ {
		stepLengthSet[i] = xValues[i+1] - xValues[i]
	}

	alpha := make([]float64, size)

	alpha[0] = 3.0*(functionValues[1]-functionValues[0])/stepLengthSet[0] - 3.0*df0
	alpha[size-1] = 3.0*dfN - 3.0*(functionValues[size-1]-functionValues[size-2])/stepLengthSet[size-2]

	for i := 1; i < size-1; i++ {
		alpha[i] = 3.0*(functionValues[i+1]-functionValues[i])/stepLengthSet[i] - 3.0*(functionValues[i]-functionValues[i-1])/stepLengthSet[i-1]
	}

	solvingSetA := make([]float64, size)
	solvingSetB := make([]float64, size-1)
	solvingSetC := make([]float64, size)

	solvingSetA[0] = 2.0 * stepLengthSet[0]
	solvingSetB[0] = 0.5
	solvingSetC[0] = alpha[0] / solvingSetA[0]

	for i := 1; i < size-1; i++ {
		solvingSetA[i] = 2.0*(xValues[i+1]-xValues[i-1]) - stepLengthSet[i-1]*solvingSetB[i-1]
		solvingSetB[i] = stepLengthSet[i] / solvingSetA[i]
		solvingSetC[i] = (alpha[i] - stepLengthSet[i-1]*solvingSetC[i-1]) / solvingSetA[i]
	}

	solvingSetA[size-1] = stepLengthSet[size-2] * (2.0 - solvingSetB[size-2])
	solvingSetC[size-1] = (alpha[size-1] - stepLengthSet[size-2]*solvingSetC[size-2]) / solvingSetA[size-1]

	var solutionSetA []float64
	solutionSetA = functionValues[:size-1]
	solutionSetB := make([]float64, size-1)
	solutionSetC := make([]float64, size)
	solutionSetD := make([]float64, size-1)

	solutionSetC[size-1] = solvingSetC[size-1]

	sizeSolutionSet := len(solutionSetB)

	for i := sizeSolutionSet - 1; i >= 0; i-- {
		solutionSetC[i] = solvingSetC[i] - solvingSetB[i]*solutionSetC[i+1]
		solutionSetB[i] = (functionValues[i+1]-functionValues[i])/stepLengthSet[i] - stepLengthSet[i]*(solutionSetC[i+1]+2.0*solutionSetC[i])/3.0
		solutionSetD[i] = (solutionSetC[i+1] - solutionSetC[i]) / (3.0 * stepLengthSet[i])
	}

	solutionSetC = solutionSetC[:size-1]

	solutionTable := [][]float64{}

	solutionTable = append(solutionTable, solutionSetA)
	solutionTable = append(solutionTable, solutionSetB)
	solutionTable = append(solutionTable, solutionSetC)
	solutionTable = append(solutionTable, solutionSetD)

	return solutionTable, nil
}

// BezierCurve is for constructing the cubic bezier curves in parametric form
func BezierCurve(endpoints [][2]float64, leftGuidepoints [][2]float64, rightGuidepoints [][2]float64) ([][][4]float64, error) {
	size := len(endpoints)

	if len(leftGuidepoints) != len(rightGuidepoints) {
		return nil, errors.New("Left and Right Guide Points sets do not have the same size")
	}

	if size-1 != len(leftGuidepoints) && size-1 != len(rightGuidepoints) {
		return nil, errors.New("Endpoints lengths are not correct")
	}

	solutionSetA := make([][4]float64, size-1)
	solutionSetB := make([][4]float64, size-1)

	for i := 0; i < size-1; i++ {
		solutionSetA[i][0] = endpoints[i][0]
		solutionSetB[i][0] = endpoints[i][1]
		solutionSetA[i][1] = 3.0 * (leftGuidepoints[i][0] - endpoints[i][0])
		solutionSetB[i][1] = 3.0 * (leftGuidepoints[i][1] - endpoints[i][1])
		solutionSetA[i][2] = 3.0 * (endpoints[i][0] + rightGuidepoints[i][0] - 2.0*leftGuidepoints[i][0])
		solutionSetB[i][2] = 3.0 * (endpoints[i][1] + rightGuidepoints[i][1] - 2.0*leftGuidepoints[i][1])
		solutionSetA[i][3] = endpoints[i+1][0] - endpoints[i][0] + 3.0*(leftGuidepoints[i][0]-rightGuidepoints[i][0])
		solutionSetB[i][3] = endpoints[i+1][1] - endpoints[i][1] + 3.0*(leftGuidepoints[i][1]-rightGuidepoints[i][1])
	}

	solutionTable := make([][][4]float64, 2)

	solutionTable[0] = solutionSetA
	solutionTable[1] = solutionSetB

	return solutionTable, nil
}
