package m

import (
	"errors"
	"reflect"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
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

	// Swaps two matrix rows
	Swap(rowA, rowB int)

	// Returns the Determinate of a matrix or error if matrix is not square.
	Det() (gcv.Value, error)

	// TODO Inverse of Matrix. Returns error if there is no inverse
	// Inv() (Matrix, error)

	// Get element at location (row, col)
	Get(row int, col int) gcv.Value

	// Set element at location (row, col)
	Set(row int, col int, value interface{})

	// Aug will take either a Vector or a Matrix and will create a new augmented matrix
	Aug(b interface{}) Matrix

	// Sub will return a sub matrix. top are the number or rows to cut from the top.
	// bottom are the number of rows to cut from the bottom.
	// left are columns to cut from the left.
	// right are columns to cut from the right.
	Sub(top, bottom, left, right int) Matrix
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
func (m *matrix) Set(row int, col int, value interface{}) {
	val := gcv.NewValue()
	val.SetValue(value)
	if m.Type() < val.Type() {
		m.coreType = val.Type()
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
	trace := gcv.NewValue()

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < m.elements.Len(); i++ {
		trace = gcvops.Add(trace, m.Get(i, i))
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
			transMatrixElements.SetValue(i, j, gcvops.Conj(m.Get(j, i)))
		}
	}

	m.numRows = transMatrixNumRows
	m.numCols = transMatrixNumCols
	m.elements = transMatrixElements
}

// implementation of Swap
func (m *matrix) Swap(rowA, rowB int) {
	vectorA := m.Elements().Get(rowA).Copy()
	vectorB := m.Elements().Get(rowB).Copy()
	m.elements.Set(rowA, vectorB)
	m.elements.Set(rowB, vectorA)
}

// implementation of Det method
func (m *matrix) Det() (gcv.Value, error) {
	if !m.IsSquare() {
		return nil, errors.New("Matrix is not square")
	}

	pivots := 0
	var det gcv.Value
	rows, cols := m.Dim()
	for i := 0; i < cols; i++ {
		for j := i + 1; j < rows; j++ {
			valueA := m.Get(i, i).Copy()
			valueB := m.Get(j, i).Copy()
			if valueA.Real() < valueB.Real() || valueA.Complex() == 0 {
				m.Swap(i, j)
				pivots++
				vector := sub(m.Elements().Get(j), sMult(valueA, sDiv(valueB, m.Elements().Get(i))))
				m.elements.Set(j, vector)
			} else {
				vector := sub(m.Elements().Get(j), sMult(valueB, sDiv(valueA, m.Elements().Get(i))))
				m.elements.Set(j, vector)
			}
		}
		if i > 1 {
			det = gcvops.Mult(m.Get(i, i), det)
		} else if i == 1 {
			det = gcvops.Mult(m.Get(1, 1), m.Get(0, 0))
		}
	}
	if pivots%2 != 0 {
		det = gcvops.Mult(gcv.MakeValue(-1), det)
	}
	return det, nil
}

// implementation of Aug method
func (m *matrix) Aug(b interface{}) Matrix {
	var augmentedMatrix Matrix
	switch b.(type) {
	case Matrix:
		rowsA, colsA := m.Dim()
		matrixB := b.(Matrix)
		rowsB, colsB := matrixB.Dim()
		if rowsA != rowsB {
			panic("Number of rows in b not equal to rows in matrix to be augmented")
		}
		augmentedMatrix = NewMatrix(rowsA, colsA+colsB)
		for i := 0; i < rowsA; i++ {
			for j := 0; j < colsA+colsB; j++ {
				if j < colsA {
					augmentedMatrix.Set(i, j, m.Get(i, j))
				} else {
					augmentedMatrix.Set(i, j, matrixB.Get(i, j-colsA))
				}
			}
		}
	case v.Vector:
		vector := b.(v.Vector)
		if vector.Space() != v.ColSpace {
			panic("Vector not in ColSpace")
		}
		rowsA, colsA := m.Dim()
		if vector.Len() != rowsA {
			panic("Vector Length not equal to Number of rows of matrix to be augmented")
		}
		augmentedMatrix = NewMatrix(rowsA, colsA+1)
		for i := 0; i < rowsA; i++ {
			for j := 0; j < colsA+1; j++ {
				if j < colsA {
					augmentedMatrix.Set(i, j, m.Get(i, j))
				} else {
					augmentedMatrix.Set(i, j, vector.Get(j-colsA))
				}
			}
		}
	default:
		panic("Type of b is not supported. Must be either Vector or Matrix")
	}
	return augmentedMatrix
}

func (m *matrix) Sub(top, bottom, left, right int) Matrix {
	rows, cols := m.Dim()
	tPlusB := (top + bottom)
	lPlusR := (left + right)
	if rows < tPlusB || cols < lPlusR {
		panic("Requested dimensions are greater than dimensions of primary matrix")
	}

	subMatrix := NewMatrix(rows-tPlusB, cols-lPlusR)

	for i := 0; i < rows-tPlusB; i++ {
		for j := 0; j < cols-lPlusR; j++ {
			subMatrix.Set(i, j, m.Get(i+top, j+left))
		}
	}

	return subMatrix
}

// NewMatrix returns a new matrix of type Matrix
func NewMatrix(rows int, cols int) Matrix {
	matrix := new(matrix)
	matrix.numRows = rows
	matrix.numCols = cols
	matrix.coreType = gcv.Real
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
	vectors := v.MakeVectorsAlt(v.RowSpace, elements)
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
