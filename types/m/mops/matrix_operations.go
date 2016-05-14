package mops

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/m"
)

// SMult is an operation for multiplying a Matrix by a scalar float64 value
func SMult(scalar gcv.Value, matrix m.Matrix) m.Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, gcvops.Mult(scalar, matrix.Get(i, j)))
		}
	}

	return newMatrix
}

// MultSimple is an operation that will multiple two matrices of any size (m X k) and (k X n) together
func MultSimple(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumRows() {
		return nil, errors.New("Length of columns of matrix A not equal to length of rows of matrix B")
	}

	var matrixAB m.Matrix
	if matrixA.IsIdentity() {
		matrixAB = matrixB.Copy()
		return matrixAB, nil
	}

	if matrixB.IsIdentity() {
		matrixAB = matrixA.Copy()
		return matrixAB, nil
	}

	matrixAB = m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())
	var sum gcv.Value
	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixB.GetNumCols(); j++ {
			sum = gcv.NewValue()
			for k := 0; k < matrixA.GetNumCols(); k++ {
				sum = gcvops.Add(sum, gcvops.Mult(matrixA.Get(i, k), matrixB.Get(k, j)))
			}
			matrixAB.Set(i, j, sum)
		}
	}
	return matrixAB, nil
}

// Add is an operation that will add two matrices together
func Add(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, gcvops.Add(matrixA.Get(i, j), matrixB.Get(i, j)))
		}
	}
	return matrixAB, nil
}

// Sub is an operation that will subtract two matrices from one another
func Sub(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, gcvops.Sub(matrixA.Get(i, j), matrixB.Get(i, j)))
		}
	}
	return matrixAB, nil
}
