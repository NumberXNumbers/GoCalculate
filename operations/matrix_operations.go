package operations

import "errors"

// MatrixMultiplicationSimple is an operation that will multiple two matrices of any size together
func MatrixMultiplicationSimple(matrixA [][]float64, matrixB [][]float64) ([][]float64, error) {
	numberRowsMatrixA := len(matrixA)
	numberColsMatrixA := len(matrixA[0])
	numberRowsMatrixB := len(matrixB)
	numberColsMatrixB := len(matrixB[0])

	if numberColsMatrixA != numberRowsMatrixB {
		return nil, errors.New("Length of columns of matrix A not equal to length of rows of matrix B")
	}

	matrixAB := make([][]float64, numberRowsMatrixA)

	for i := 0; i < len(matrixAB); i++ {
		matrixAB[i] = make([]float64, numberColsMatrixB)
	}

	var sum float64

	for i := 0; i < numberRowsMatrixA; i++ {
		for j := 0; j < numberColsMatrixB; j++ {
			sum = float64(0)
			for k := 0; k < numberColsMatrixA; k++ {
				sum += matrixA[i][k] * matrixB[k][j]
			}
			matrixAB[i][j] = sum
		}
	}

	return matrixAB, nil
}
