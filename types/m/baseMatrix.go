package m

import (
	"errors"
	"reflect"
)

type baseMatrix interface {
	// Returns dimensions of Matrix.
	Dim() (rows, cols int)

	// Returns number of elements of Matrix
	NumElements() int

	// Returns true is a Matrix is square, Else false
	IsSquare() bool

	// Returns if a matrix is an Identity Matrix
	IsIdentity() bool

	// Returns Number of rows
	GetNumRows() int

	// Returns Number of columns
	GetNumCols() int

	// Type of Matrix
	Type() reflect.Kind

	// Transpose of a Matrix
	Trans()
}

type matrixBase struct {
	numRows int
	numCols int
}

// NewMatrix will return either a Matrix, MatrixComplex or Error depending on matrixType given
func NewMatrix(row int, col int, matrixType reflect.Kind) (Matrix, MatrixComplex, error) {
	switch {
	default:
		return nil, nil, errors.New("Unsupported Matrix Type")
	case matrixType == reflect.Float64:
		return MakeMatrix(row, col), nil, nil
	case matrixType == reflect.Complex128:
		return nil, MakeComplexMatrix(row, col), nil
	}

}
