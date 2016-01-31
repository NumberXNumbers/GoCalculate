package coreTypes

import (
	"errors"
	"reflect"
)

// MatrixComplex is the main matirx interface for complex matrices.
type MatrixComplex interface {
	baseMatrix

	// Transpose of a Matrix
	Trans() MatrixComplex

	// Returns all Elements
	GetElements() [][]complex128

	// Trace of Matrix. Returns error if matrix is not square
	Tr() (complex128, error)

	// TODO Determinate of Matrix. Returns error is there is no determinate
	// Det() (complex128, error)

	// TODO Inverse of Matrix. Returns error if there is no inverse
	// Inv() (MatrixComplex, error)

	// Get element at location (row, col)
	Get(row int, col int) complex128

	// Set element at location (row, col)
	Set(row int, col int, value complex128)
}

type matrixComplex struct {
	numRows    int
	numCols    int
	matrixType reflect.Kind
	elements   [][]complex128
}

// implementation of Size method
func (m matrixComplex) Size() (rows, cols int) { return m.numRows, m.numCols }

// implementation of NumElements method
func (m matrixComplex) NumElements() int { return m.numCols * m.numRows }

// implementation of Type method
func (m matrixComplex) Type() reflect.Kind { return m.matrixType }

// implementation of GetRows method
func (m matrixComplex) GetNumRows() int { return m.numRows }

// implementation of GetColumns method
func (m matrixComplex) GetNumCols() int { return m.numCols }

// implementation of IsSquare method
func (m matrixComplex) IsSquare() bool { return m.GetNumCols() == m.GetNumRows() }

// implementation of GetElements method
func (m matrixComplex) GetElements() [][]complex128 { return m.elements }

// implementation of Get method
func (m matrixComplex) Get(row int, col int) complex128 { return m.elements[row][col] }

// implementation of Set method
func (m matrixComplex) Set(row int, col int, value complex128) { m.elements[row][col] = value }

// implementation of IsIdentity method
func (m matrixComplex) IsIdentity() bool {
	return reflect.DeepEqual(m, MakeIdentityComplexMatrix(m.GetNumRows()))
}

// implementation of Tr method
func (m matrixComplex) Tr() (complex128, error) {
	var trace complex128

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.GetElements()); i++ {
		trace += m.Get(i, i)
	}

	return trace, nil
}

// implementation of Trans method
func (m matrixComplex) Trans() MatrixComplex {
	transMatrixNumCols := m.numRows
	transMatrixNumRows := m.numCols
	transMatrixType := m.Type()
	transMatrixElements := make([][]complex128, transMatrixNumRows)

	for i := 0; i < len(transMatrixElements); i++ {
		transMatrixElements[i] = make([]complex128, transMatrixNumCols)
	}

	for i := 0; i < transMatrixNumRows; i++ {
		for j := 0; j < transMatrixNumCols; j++ {
			transMatrixElements[i][j] = m.Get(j, i)
		}
	}

	transposeMatrix := MakeComplexMatrixWithElements(transMatrixNumRows, transMatrixNumCols, transMatrixType, transMatrixElements)

	return transposeMatrix
}

// MakeComplexMatrix returns a new matrix of type MatrixComplex128
func MakeComplexMatrix(rows int, cols int, matrixType reflect.Kind) MatrixComplex {
	matrix := new(matrixComplex)
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.matrixType = matrixType
	matrix.matrixType = matrixType

	matrixElements := make([][]complex128, rows)

	for i := 0; i < rows; i++ {
		matrixElements[i] = make([]complex128, cols)
	}

	matrix.elements = matrixElements

	return matrix
}

// MakeComplexMatrixWithElements returns a new matrix of type MatrixComplex128 with predefined elements
func MakeComplexMatrixWithElements(rows int, cols int, matrixType reflect.Kind, elements [][]complex128) MatrixComplex {
	matrix := new(matrixComplex)
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.matrixType = matrixType
	matrix.elements = elements
	return matrix
}

// MakeIdentityComplexMatrix returns a new Identity Matrix with complex128 diagonal values of size (degree, degree)
func MakeIdentityComplexMatrix(degree int) MatrixComplex {
	matrix := MakeComplexMatrix(degree, degree, reflect.Float64)

	for i := 0; i < degree; i++ {
		matrix.Set(i, i, complex128(1))
	}

	return matrix
}
