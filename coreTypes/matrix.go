package coreTypes

import (
	"errors"
	"reflect"
)

// Matrix is the base matirx interface.
type Matrix interface {
	// Returns dimensions of Matrix.
	Size() (rows, cols int)

	// Returns number of elements of Matrix
	NumElements() int

	// Returns true is a Matrix is square, Else false
	IsSquare() bool

	// Type of Matrix
	Type() reflect.Type
}

// MatrixReal is a Matrix with only real Elements
type MatrixReal interface {
	Matrix

	// Trace of Matrix. Returns error if matrix is not square
	Tr() (float64, error)

	// TODO Determinate of Matrix. Returns error is there is no determinate
	// Det() (float64, error)

	// TODO Inverse of Matrix. Returns error if there is no inverse
	// Inv() (float64, error)
}

// MatrixComplex is a Matrix with complex values.
type MatrixComplex interface {
	Matrix

	// Trace of Matrix. Returns error if matrix is not square
	Tr() (complex128, error)

	// TODO Determinate of Matrix. Returns error is there is no determinate
	// Det() (complex128, error)

	// TODO Inverse of Matrix. Returns error if there is no inverse
	// Inv() (complex128, error)
}

// MatrixInt is a MatrixReal with int elements
type MatrixInt interface {
	MatrixReal

	// Get Method for int MatrixInt
	Get(row int, col int) (int, error)
}

// MatrixFloat32 is a MatrixReal with float32 elements
type MatrixFloat32 interface {
	MatrixReal

	// Get Method for int MatrixFloat32
	Get(row int, col int) (float32, error)
}

// MatrixFloat64 is a MatrixReal with float64 elements
type MatrixFloat64 interface {
	MatrixReal

	// Get Method for int MatrixFloat64
	Get(row int, col int) (float64, error)
}

// MatrixComplex64 is a MatrixComplex with complex64 elements
type MatrixComplex64 interface {
	MatrixComplex

	// Get Method for int MatrixComplex64
	Get(row int, col int) (complex64, error)
}

// MatrixComplex128 is a MatrixComplex with complex128 elements
type MatrixComplex128 interface {
	MatrixComplex

	// Get Method for int MatrixComplex128
	Get(row int, col int) (complex128, error)
}

type matrix struct {
	numRows int
	numCols int
}

type matrixInt struct {
	matrix
	elements [][]int
}

type matrixFloat32 struct {
	matrix
	elements [][]float32
}

type matrixFloat64 struct {
	matrix
	elements [][]float64
}

type matrixComplex64 struct {
	matrix
	elements [][]complex64
}

type matrixComplex128 struct {
	matrix
	elements [][]complex128
}

// implementation of Size methods
func (m matrix) Size() (rows, cols int) { return m.numRows, m.numCols }

// implementation of NumElements methods
func (m matrix) NumElements() int { return m.numCols * m.numRows }

// implementation of Type methods
func (m matrixInt) Type() reflect.Type        { return reflect.TypeOf(m.elements) }
func (m matrixFloat32) Type() reflect.Type    { return reflect.TypeOf(m.elements) }
func (m matrixFloat64) Type() reflect.Type    { return reflect.TypeOf(m.elements) }
func (m matrixComplex64) Type() reflect.Type  { return reflect.TypeOf(m.elements) }
func (m matrixComplex128) Type() reflect.Type { return reflect.TypeOf(m.elements) }

// implementation of IsSquare methods
func (m matrix) IsSquare() bool {
	if m.numCols == m.numRows {
		return true
	}

	return false
}

// implementation of Get methods
func (m matrixInt) Get(row int, col int) (int, error) {
	var element int

	if row >= m.numRows || row < 0 {
		return element, errors.New("Row value is out of bounds")
	}

	if col >= m.numCols || col < 0 {
		return element, errors.New("Column value is out of bounds")
	}

	return m.elements[row][col], nil
}

func (m matrixFloat32) Get(row int, col int) (float32, error) {
	var element float32

	if row >= m.numRows || row < 0 {
		return element, errors.New("Row value is out of bounds")
	}

	if col >= m.numCols || col < 0 {
		return element, errors.New("Column value is out of bounds")
	}

	return m.elements[row][col], nil
}

func (m matrixFloat64) Get(row int, col int) (float64, error) {
	var element float64

	if row >= m.numRows || row < 0 {
		return element, errors.New("Row value is out of bounds")
	}

	if col >= m.numCols || col < 0 {
		return element, errors.New("Column value is out of bounds")
	}

	return m.elements[row][col], nil
}

func (m matrixComplex64) Get(row int, col int) (complex64, error) {
	var element complex64

	if row >= m.numRows || row < 0 {
		return element, errors.New("Row value is out of bounds")
	}

	if col >= m.numCols || col < 0 {
		return element, errors.New("Column value is out of bounds")
	}

	return m.elements[row][col], nil
}

func (m matrixComplex128) Get(row int, col int) (complex128, error) {
	var element complex128

	if row >= m.numRows || row < 0 {
		return element, errors.New("Row value is out of bounds")
	}

	if col >= m.numCols || col < 0 {
		return element, errors.New("Column value is out of bounds")
	}

	return m.elements[row][col], nil
}

// implementation of Tr methods
func (m matrixInt) Tr() (float64, error) {
	var trace float64
	var currentValue int

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.elements); i++ {
		currentValue, _ = m.Get(i, i)
		trace += float64(currentValue)
	}

	return trace, nil
}

func (m matrixFloat32) Tr() (float64, error) {
	var trace float64
	var currentValue float32

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.elements); i++ {
		currentValue, _ = m.Get(i, i)
		trace += float64(currentValue)
	}

	return trace, nil
}

func (m matrixFloat64) Tr() (float64, error) {
	var trace float64
	var currentValue float64

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.elements); i++ {
		currentValue, _ = m.Get(i, i)
		trace += currentValue
	}

	return trace, nil
}

func (m matrixComplex64) Tr() (complex128, error) {
	var trace complex128
	var currentValue complex64

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.elements); i++ {
		currentValue, _ = m.Get(i, i)
		trace += complex128(currentValue)
	}

	return trace, nil
}

func (m matrixComplex128) Tr() (complex128, error) {
	var trace complex128
	var currentValue complex128

	if !m.IsSquare() {
		return trace, errors.New("Matrix is not square")
	}

	for i := 0; i < len(m.elements); i++ {
		currentValue, _ = m.Get(i, i)
		trace += currentValue
	}

	return trace, nil
}
