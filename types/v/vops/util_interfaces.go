package vops

import (
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

const (
	// RowVector for operations
	RowVector = v.RowVector
	// ColVector for operations
	ColVector = v.ColVector
	// Complex for operations
	Complex = gcv.Complex
)

// Value is an interface over type Value
type Value interface {
	gcv.Value
}

func newValue(val interface{}) (value Value) {
	value = gcv.NewValue(val)
	return
}

// Vector is an interface over coreTypes Vector for operations
type Vector interface {
	v.Vector
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
