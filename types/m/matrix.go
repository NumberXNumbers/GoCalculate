package m

import (
	"errors"
	"reflect"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// Matrix is the main matirx interface for real matrices.
type Matrix interface {
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
	Type() string

	// Transpose of a Matrix
	Trans()

	// Returns a copy of Matrix
	Copy() Matrix

	// Returns all Elements
	Elements() v.Vectors

	// Trace of Matrix. Returns error if matrix is not square
	Tr() (gcv.Value, error)

	// TODO Determinate of Matrix. Returns error is there is no determinate
	// Det() (float64, error)

	// TODO Inverse of Matrix. Returns error if there is no inverse
	// Inv() (MatrixReal, error)

	// Get element at location (row, col)
	Get(row int, col int) gcv.Value

	// Set element at location (row, col)
	Set(row int, col int, value gcv.Value)
}

type matrix struct {
	numRows    int
	numCols    int
	matrixType string
	elements   v.Vectors
}

// implementation of Dim method
func (m *matrix) Dim() (rows, cols int) { return m.numRows, m.numCols }

// implementation of NumElements method
func (m *matrix) NumElements() int { return m.numCols * m.numRows }

// implementation of Type method
func (m *matrix) Type() string { return m.matrixType }

// implementation of GetRows method
func (m *matrix) GetNumRows() int { return m.numRows }

// implementation of GetColumns method
func (m *matrix) GetNumCols() int { return m.numCols }

// implementation of IsSquare method
func (m *matrix) IsSquare() bool { return m.GetNumCols() == m.GetNumRows() }

// implementation of GetElements method
func (m *matrix) GetElements() v.Vectors { return m.elements }

// implementation of Get method
func (m *matrix) Get(row int, col int) gcv.Value { return m.elements.Get(row).Get(col) }

// implementation of Set method
func (m *matrix) Set(row int, col int, val gcv.Value) { m.elements.SetValue(row, col, val) }

// implementation of IsIdentity method
func (m *matrix) IsIdentity() bool {
	return reflect.DeepEqual(m, MakeIdentityMatrix(m.GetNumRows()))
}

//implementation of Copy method
func (m matrix) Copy() Matrix {
	return MakeMatrixWithElements(m.Elements())
}

// implementation of Tr method
func (m *matrix) Tr() (gcv.Value, error) {
	var trace gcv.Value

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < m.elements.Len(); i++ {
		trace += m.Get(i, i)
	}

	return trace, nil
}

// implementation of Trans method
func (m *matrix) Trans() {
	transMatrixNumCols := m.numRows
	transMatrixNumRows := m.numCols
	transMatrixElements := make([][]float64, transMatrixNumRows)

	for i := 0; i < len(transMatrixElements); i++ {
		transMatrixElements[i] = make([]float64, transMatrixNumCols)
	}

	for i := 0; i < transMatrixNumRows; i++ {
		for j := 0; j < transMatrixNumCols; j++ {
			transMatrixElements[i][j] = m.Get(j, i)
		}
	}

	m.numRows = transMatrixNumRows
	m.numCols = transMatrixNumCols
	m.elements = transMatrixElements
}

// NewMatrix returns a new matrix of type Matrix
func NewMatrix(rows int, cols int) Matrix {
	matrix := new(matrix)
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.matrixType = gcv.Int
	matrix.elements = v.NewVectors(v.RowSpace)
	return matrix
}

// MakeMatrixWithElements returns a new matrix of type Matrix with predefined elements
func MakeMatrixWithElements(elements v.Vectors) Matrix {
	matrix := new(matrix)
	rows := elements.Len()
	cols := elements.Get(0).Len()
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.matrixType = reflect.Float64

	matrixElements := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		matrixElements[i] = make([]float64, cols)
		copy(matrixElements[i], elements[i])
	}

	matrix.elements = matrixElements
	return matrix
}

// MakeIdentityMatrix returns a new Identity Matrix of size (degree, degree)
func MakeIdentityMatrix(degree int) Matrix {
	matrix := MakeMatrix(degree, degree)

	for i := 0; i < degree; i++ {
		matrix.Set(i, i, float64(1))
	}

	return matrix
}
