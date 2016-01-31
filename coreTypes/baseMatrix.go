package coreTypes

type baseMatrix interface {
	// Returns dimensions of Matrix.
	Size() (rows, cols int)

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
}
