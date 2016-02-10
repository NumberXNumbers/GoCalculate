package matrixOps

import "errors"

// MatrixScalarMulti is an operation for multiplying a Matrix by a scalar float64 value
func MatrixScalarMulti(scalar float64, matrix Matrix) Matrix {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, scalar*matrix.Get(i, j))
		}
	}

	return newMatrix
}

// MatrixComplexScalarMulti is an operation for multiplying a Matrix by a scalar float64 value
func MatrixComplexScalarMulti(scalar complex128, matrix MatrixComplex) MatrixComplex {
	newMatrix := matrix.Copy()
	for i := 0; i < matrix.GetNumRows(); i++ {
		for j := 0; j < matrix.GetNumCols(); j++ {
			newMatrix.Set(i, j, scalar*matrix.Get(i, j))
		}
	}

	return newMatrix
}

// MatrixMultiSimple is an operation that will multiple two matrices of any size together
func MatrixMultiSimple(matrixA Matrix, matrixB Matrix) (Matrix, error) {
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

	matrixAB := makeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

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
func MatrixComplexMultiSimple(matrixA MatrixComplex, matrixB MatrixComplex) (MatrixComplex, error) {
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

	matrixAB := makeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumRows())

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
func MatrixAddition(matrixA Matrix, matrixB Matrix) (Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := makeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)+matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixComplexAddition is an operation that will add two complex matrices together
func MatrixComplexAddition(matrixA MatrixComplex, matrixB MatrixComplex) (MatrixComplex, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := makeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)+matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixSubtraction is an operation that will subtract two matrices from one another
func MatrixSubtraction(matrixA Matrix, matrixB Matrix) (Matrix, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := makeMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)-matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}

// MatrixComplexSubtraction is an operation that will subtract two complex matrices from one another
func MatrixComplexSubtraction(matrixA MatrixComplex, matrixB MatrixComplex) (MatrixComplex, error) {
	if matrixA.GetNumCols() != matrixB.GetNumCols() || matrixA.GetNumRows() != matrixB.GetNumRows() {
		return nil, errors.New("Matrices do not have equivalent dimensions")
	}

	matrixAB := makeComplexMatrix(matrixA.GetNumRows(), matrixB.GetNumCols())

	for i := 0; i < matrixA.GetNumRows(); i++ {
		for j := 0; j < matrixA.GetNumCols(); j++ {
			matrixAB.Set(i, j, matrixA.Get(i, j)-matrixB.Get(i, j))
		}
	}

	return matrixAB, nil
}
