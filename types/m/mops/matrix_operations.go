package mops

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
)

// ScalarMultiplication is an operation for multiplying a Matrix by a scalar float64 value
func ScalarMultiplication(scalar gcv.Value, matrix m.Matrix) m.Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			if matrix.Type() == gcv.Complex || scalar.Type() == gcv.Complex {
				newMatrix.Set(i, j, gcv.MakeValue(scalar.Complex()*matrix.Get(i, j).Complex()))
				continue
			}
			newMatrix.Set(i, j, gcv.MakeValue(scalar.Real()*matrix.Get(i, j).Real()))
		}
	}

	return newMatrix
}

// MultiplicationSimple is an operation that will multiple two matrices of any size together
func MultiplicationSimple(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumRows() {
		return nil, errors.New("Length of columns of matrix A not equal to length of rows of matrix B")
	}

	if matrixA.IsIdentity() || matrixB.IsIdentity() {
		_, degree := matrixA.Dim()
		matrixAB := m.NewIdentityMatrix(degree)
		return matrixAB, nil
	}

	matrixAB := m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())
	var sum gcv.Value
	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixB.GetNumCols(); j++ {
			if matrixA.Type() == gcv.Complex || matrixB.Type() == gcv.Complex {
				var sumComplex complex128
				for k := 0; k < matrixA.GetNumCols(); k++ {
					sumComplex += matrixA.Get(i, k).Complex() * matrixB.Get(k, j).Complex()
				}
				sum = gcv.MakeValue(sumComplex)
			} else {
				var sumFloat float64
				for k := 0; k < matrixA.GetNumCols(); k++ {
					sumFloat += matrixA.Get(i, k).Real() * matrixB.Get(k, j).Real()
				}
				sum = gcv.MakeValue(sumFloat)
			}
			matrixAB.Set(i, j, sum)
		}
	}
	return matrixAB, nil
}

// Addition is an operation that will add two matrices together
func Addition(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			if matrixA.Type() == gcv.Complex || matrixB.Type() == gcv.Complex {
				matrixAB.Set(i, j, gcv.MakeValue(matrixA.Get(i, j).Complex()+matrixB.Get(i, j).Complex()))
				continue
			}
			matrixAB.Set(i, j, gcv.MakeValue(matrixA.Get(i, j).Real()+matrixB.Get(i, j).Real()))
		}
	}
	return matrixAB, nil
}

// Subtraction is an operation that will subtract two matrices from one another
func Subtraction(matrixA m.Matrix, matrixB m.Matrix) (m.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := m.NewMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			if matrixA.Type() == gcv.Complex || matrixB.Type() == gcv.Complex {
				matrixAB.Set(i, j, gcv.MakeValue(matrixA.Get(i, j).Complex()-matrixB.Get(i, j).Complex()))
				continue
			}
			matrixAB.Set(i, j, gcv.MakeValue(matrixA.Get(i, j).Real()-matrixB.Get(i, j).Real()))
		}
	}
	return matrixAB, nil
}
