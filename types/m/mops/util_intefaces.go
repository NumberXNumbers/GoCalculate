package mops

import "github.com/NumberXNumbers/GoCalculate/types/m"

// Matrix is an interface over the coreTypes Matrix Type meant for operations
type Matrix interface {
	m.Matrix
}

// MatrixComplex is an interface over the coreTypes MatrixComplex Type meant for operations
type MatrixComplex interface {
	m.MatrixComplex
}

func makeMatrix(rows int, cols int) (matrix Matrix) {
	matrix = m.MakeMatrix(rows, cols)
	return
}

func makeComplexMatrix(rows int, cols int) (matrix MatrixComplex) {
	matrix = m.MakeComplexMatrix(rows, cols)
	return
}
