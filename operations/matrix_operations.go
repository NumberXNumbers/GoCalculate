package operations

import (
	"errors"

	ct "github.com/traviscox1990/GoCalculate/coreTypes"
)

// MatrixScalarMulti is an operation for multiplying a Matrix by a scalar float64 value
func MatrixScalarMulti(scalar float64, matrix ct.Matrix) ct.Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, scalar*matrix.Get(i, j))
		}
	}

	return newMatrix
}

// MatrixComplexScalarMulti is an operation for multiplying a Matrix by a scalar float64 value
func MatrixComplexScalarMulti(scalar complex128, matrix ct.MatrixComplex) ct.MatrixComplex {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, scalar*matrix.Get(i, j))
		}
	}

	return newMatrix
}

// MatrixMultiSimple is an operation that will multiple two matrices of any size together
func MatrixMultiSimple(matrixA ct.Matrix, matrixB ct.Matrix) (ct.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumRows() {
		return nil, errors.New("Length of columns of matrix A not equal to length of rows of matrix B")
	}

	if matrixA.IsIdentity() {
		return matrixB, nil
	}

	if matrixB.IsIdentity() {
		return matrixA, nil
	}

	var sum float64

	matrixAB := ct.MakeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixB.GetNumCols(); j++ {
			sum = 0.0
			for k := 0; k < matrixA.GetNumCols(); k++ {
				sum += matrixA.Get(i, k) * matrixB.Get(k, j)
			}
			matrixAB.Set(i, j, sum)
		}
	}

	return matrixAB, nil
}

// MatrixComplexMultiSimple is an operation that will multiple two complex matrices of any size together
func MatrixComplexMultiSimple(matrixA ct.MatrixComplex, matrixB ct.MatrixComplex) (ct.MatrixComplex, error) {
	if matrixA.GetNumCols() != matrixB.GetNumRows() {
		return nil, errors.New("Length of columns of matrix A not equal to length of rows of matrix B")
	}

	if matrixA.IsIdentity() {
		return matrixB, nil
	}

	if matrixB.IsIdentity() {
		return matrixA, nil
	}

	var sum complex128

	matrixAB := ct.MakeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumRows())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixB.GetNumCols(); j++ {
			sum = 0.0
			for k := 0; k < matrixA.GetNumCols(); k++ {
				sum += matrixA.Get(i, k) * matrixB.Get(k, j)
			}
			matrixAB.Set(i, j, sum)
		}
	}

	return matrixAB, nil
}

// MatrixAddition is an operation that will add two matrices together
func MatrixAddition(matrixA ct.Matrix, matrixB ct.Matrix) (ct.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() && matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := ct.MakeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)+matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixComplexAddition is an operation that will add two complex matrices together
func MatrixComplexAddition(matrixA ct.MatrixComplex, matrixB ct.MatrixComplex) (ct.MatrixComplex, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() && matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := ct.MakeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)+matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixSubtraction is an operation that will subtract two matrices from one another
func MatrixSubtraction(matrixA ct.Matrix, matrixB ct.Matrix) (ct.Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() && matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := ct.MakeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)-matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixComplexSubtraction is an operation that will subtract two complex matrices from one another
func MatrixComplexSubtraction(matrixA ct.MatrixComplex, matrixB ct.MatrixComplex) (ct.MatrixComplex, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() && matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := ct.MakeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)-matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}
