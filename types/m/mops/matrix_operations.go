package mops

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// SMult is an operation for multiplying a Matrix by a scalar Value
func SMult(scalar gcv.Value, matrix m.Matrix) m.Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, gcvops.Mult(scalar, matrix.Get(i, j)))
		}
	}

	return newMatrix
}

// SDiv will divide a Matrix by a scalar Value
func SDiv(scalar gcv.Value, matrix m.Matrix) m.Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, gcvops.Div(matrix.Get(i, j), scalar))
		}
	}

	return newMatrix
}

// VMMult will multiply a vector V and a matrix M together by V*M
func VMMult(vector v.Vector, matrix m.Matrix) (v.Vector, error) {
	if vector.Space() != v.RowSpace {
		return nil, errors.New("Vector is not in Row Space")
	}
	rows, _ := matrix.Dim()
	if vector.Len() != rows {
		return nil, errors.New("Vector Length not equal to the number of rows in Matrix")
	}
	newVector := vector.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		sum := gcv.NewValue()
		for j := 0; j < matrix.GetNumCols(); j++ {
			sum = gcvops.Add(sum, gcvops.Mult(vector.Get(j), matrix.Get(j, i)))
		}
		newVector.Set(i, sum)
	}
	return newVector, nil
}

// MVMult will multiply a vector V and a matrix M together by M*V
func MVMult(vector v.Vector, matrix m.Matrix) (v.Vector, error) {
	if vector.Space() != v.ColSpace {
		return nil, errors.New("Vector is not in Column Space")
	}
	_, cols := matrix.Dim()
	if vector.Len() != cols {
		return nil, errors.New("Vector Length not equal to the number of columns in Matrix")
	}
	newVector := vector.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		sum := gcv.NewValue()
		for j := 0; j < matrix.GetNumCols(); j++ {
			sum = gcvops.Add(sum, gcvops.Mult(matrix.Get(i, j), vector.Get(j)))
		}
		newVector.Set(i, sum)
	}
	return newVector, nil
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

func squareAndMultiplyHelper(matrixA, matrixB m.Matrix, n int) (m.Matrix, error) {
	if n < 0 {
		newMatrix, err := matrixB.Copy().Inv()
		if err != nil {
			return nil, err
		}
		return squareAndMultiplyHelper(matrixA, newMatrix, -n)
	} else if n == 0 {
		return matrixA, nil
	} else if n == 1 {
		return MultSimple(matrixA, matrixB)
	} else if n%2 == 0 {
		newMatrix, err := MultSimple(matrixB, matrixB)
		if err != nil {
			return nil, err
		}
		return squareAndMultiplyHelper(matrixA, newMatrix, n/2)
	}

	newMatrixA, errA := MultSimple(matrixB, matrixA)
	if errA != nil {
		return nil, errA
	}
	newMatrixB, errB := MultSimple(matrixB, matrixB)
	if errB != nil {
		return nil, errB
	}
	return squareAndMultiplyHelper(newMatrixA, newMatrixB, (n-1)/2)

}

// squareAndMultiply will solve the power of a matrix by squaring and multiplying
func squareAndMultply(matrix m.Matrix, n int) (m.Matrix, error) {
	degree, _ := matrix.Dim()
	identityMatrix := m.NewIdentityMatrix(degree)
	return squareAndMultiplyHelper(identityMatrix, matrix, n)
}

// Pow is an operation that will raise a square Matrix to int n
func Pow(matrix m.Matrix, n int) (m.Matrix, error) {
	if !matrix.IsSquare() {
		return nil, errors.New("Matrix is not square")
	}

	if matrix.IsIdentity() {
		return matrix, nil
	}

	return squareAndMultply(matrix, n)
}

//Exp will give the matrix exponential of matrix
func Exp(matrix m.Matrix) m.Matrix {
	return matrix
}
