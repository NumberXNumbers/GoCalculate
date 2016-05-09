package m

import (
	"errors"
	"math/cmplx"
	"reflect"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// Matrix is the main matirx interface for real matrices.
type Matrix interface {
	// Returns dimensions of Matrix.
	Dim() (rows, cols int)

	// Returns number of elements of Matrix
	TotalElements() int

	// Returns true is a Matrix is square, Else false
	IsSquare() bool

	// Returns if a matrix is an Identity Matrix
	IsIdentity() bool

	// Returns Number of rows
	GetNumRows() int

	// Returns Number of columns
	GetNumCols() int

	// Type of Matrix
	Type() gcv.Type

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
	numRows  int
	numCols  int
	coreType gcv.Type
	elements v.Vectors
}

// implementation of Dim method
func (m *matrix) Dim() (rows, cols int) { return m.numRows, m.numCols }

// implementation of TotalElements method
func (m *matrix) TotalElements() int { return m.numCols * m.numRows }

// implementation of Type method
func (m *matrix) Type() gcv.Type { return m.coreType }

// implementation of GetRows method
func (m *matrix) GetNumRows() int { return m.numRows }

// implementation of GetColumns method
func (m *matrix) GetNumCols() int { return m.numCols }

// implementation of IsSquare method
func (m *matrix) IsSquare() bool { return m.GetNumCols() == m.GetNumRows() }

// implementation of Elements method
func (m *matrix) Elements() v.Vectors { return m.elements }

// implementation of Get method
func (m *matrix) Get(row int, col int) gcv.Value { return m.elements.Get(row).Get(col) }

// implementation of Set method
func (m *matrix) Set(row int, col int, val gcv.Value) {
	if m.Type() < val.GetValueType() {
		m.coreType = val.GetValueType()
	}
	m.elements.SetValue(row, col, val)
}

// implementation of IsIdentity method
func (m *matrix) IsIdentity() bool {
	return reflect.DeepEqual(m, NewIdentityMatrix(m.GetNumRows()))
}

//implementation of Copy method
func (m matrix) Copy() Matrix {
	return MakeMatrixAlt(m.Elements())
}

// implementation of Tr method
func (m *matrix) Tr() (gcv.Value, error) {
	var trace gcv.Value

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	if m.Type() == gcv.Complex {
		var complexTrace complex128
		for i := 0; i < m.elements.Len(); i++ {
			complexTrace += m.Get(i, i).Complex128()
		}
		trace = gcv.MakeValue(complexTrace)
	} else {
		var floatTrace float64
		for i := 0; i < m.elements.Len(); i++ {
			floatTrace += m.Get(i, i).Float64()
		}
		trace = gcv.MakeValue(floatTrace)
	}

	return trace, nil
}

// implementation of Trans method
func (m *matrix) Trans() {
	transMatrixNumCols := m.numRows
	transMatrixNumRows := m.numCols
	transMatrixElements := v.NewVectors(v.RowSpace, transMatrixNumRows, transMatrixNumCols)

	for i := 0; i < transMatrixNumRows; i++ {
		for j := 0; j < transMatrixNumCols; j++ {
			if m.Type() == gcv.Complex {
				transMatrixElements.SetValue(i, j, gcv.MakeValue(cmplx.Conj(m.Get(j, i).Complex128())))
				continue
			}
			transMatrixElements.SetValue(i, j, m.Get(j, i))
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
	matrix.coreType = gcv.Int
	matrix.elements = v.NewVectors(v.RowSpace, rows, cols)
	return matrix
}

// MakeMatrixAlt returns a new matrix of type Matrix with predefined elements. takes type Vectors as arg
func MakeMatrixAlt(elements v.Vectors) Matrix {
	matrix := new(matrix)
	rows := elements.Len()
	cols := elements.InnerLen()
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.coreType = elements.Type()
	matrixElements := v.NewVectors(v.RowSpace, matrix.numRows, matrix.numCols)

	for i := 0; i < matrix.numRows; i++ {
		vector := elements.Get(i).Copy()
		lenVector := vector.Len()
		for j := 0; j < matrix.numCols; j++ {
			if lenVector <= j {
				break
			}
			matrixElements.SetValue(i, j, vector.Get(j).Copy())
		}
	}

	matrix.elements = matrixElements
	return matrix
}

// MakeMatrix returns a new matrix. takes individual vectors as args
func MakeMatrix(elements ...v.Vector) Matrix {
	vectors := v.MakeVectors(v.RowSpace, elements...)
	matrix := MakeMatrixAlt(vectors)
	return matrix
}

// NewIdentityMatrix returns a new Identity Matrix of size (degree, degree)
func NewIdentityMatrix(degree int) Matrix {
	matrix := NewMatrix(degree, degree)

	for i := 0; i < degree; i++ {
		matrix.Set(i, i, gcv.MakeValue(1))
	}

	return matrix
}

// MakeConjMatrix returns a new conj matrix of matrix
func MakeConjMatrix(m Matrix) Matrix {
	conjMatrix := m.Copy()
	conjMatrix.Trans()
	return conjMatrix
}
