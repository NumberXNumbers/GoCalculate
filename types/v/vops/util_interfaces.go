package vops

import (
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

const (
	// RowVector for operations
	RowVector = v.RowVector
	// ColVector for operations
	ColVector = v.ColVector
)

// Vector is an interface over coreTypes Vector for operations
type Vector interface {
	v.Vector
}

// VectorComplex is an interface over coreTypes VectorComplex for operations
type VectorComplex interface {
	v.VectorComplex
}

func makeComplexVector(length int, vectorType string) (vector VectorComplex) {
	vector = v.MakeComplexVector(length, vectorType)
	return
}

func makeVector(length int, vectorType string) (vector Vector) {
	vector = v.MakeVector(length, vectorType)
	return
}

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
