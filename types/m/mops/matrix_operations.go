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
			if matrix.Type() == gcv.Complex || scalar.GetValueType() == gcv.Complex {
				newMatrix.Set(i, j, gcv.NewValue(scalar.Complex128()*matrix.Get(i, j).Complex128()))
				continue
			}
			newMatrix.Set(i, j, gcv.NewValue(scalar.Float64()*matrix.Get(i, j).Float64()))
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
					sumComplex += matrixA.Get(i, k).Complex128() * matrixB.Get(k, j).Complex128()
				}
				sum = gcv.NewValue(sumComplex)
			} else {
				var sumFloat float64
				for k := 0; k < matrixA.GetNumCols(); k++ {
					sumFloat += matrixA.Get(i, k).Float64() * matrixB.Get(k, j).Float64()
				}
				sum = gcv.NewValue(sumFloat)
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
				matrixAB.Set(i, j, gcv.NewValue(matrixA.Get(i, j).Complex128()+matrixB.Get(i, j).Complex128()))
				continue
			}
			matrixAB.Set(i, j, gcv.NewValue(matrixA.Get(i, j).Float64()+matrixB.Get(i, j).Float64()))
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
				matrixAB.Set(i, j, gcv.NewValue(matrixA.Get(i, j).Complex128()-matrixB.Get(i, j).Complex128()))
				continue
			}
			matrixAB.Set(i, j, gcv.NewValue(matrixA.Get(i, j).Float64()-matrixB.Get(i, j).Float64()))
		}
	}
	return matrixAB, nil
}
