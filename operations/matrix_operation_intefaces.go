package operations

import "github.com/traviscox1990/GoCalculate/coreTypes"

// Matrix is an interface over the coreTypes Matrix Type meant for operations
type Matrix interface {
	coreTypes.Matrix
}

// MatrixComplex is an interface over the coreTypes MatrixComplex Type meant for operations
type MatrixComplex interface {
	coreTypes.MatrixComplex
}

func makeMatrix(rows int, cols int) (matrix Matrix) {
	matrix = coreTypes.MakeMatrix(rows, cols)
	return
}

func makeComplexMatrix(rows int, cols int) (matrix MatrixComplex) {
	matrix = coreTypes.MakeComplexMatrix(rows, cols)
	return
}
