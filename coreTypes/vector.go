package coreTypes

import (
	"math"

	"github.com/traviscox1990/GoCalculate/operations"
)

// Vector is the main vector interface for real vectors
type Vector interface {
	baseVector

	// Get returns element at index in vector
	Get(index int) float64

	// Set will set a value at index in a vector
	Set(index int, value float64)

	// Return norm of vectorBase
	Norm() float64

	// Returns a new Copy of vector
	Copy() Vector
}

type vector struct {
	vectorBase
	elements []float64
}

// implementation of Get method
func (v vector) Get(index int) float64 { return v.elements[index] }

// implementation of Set method
func (v vector) Set(index int, value float64) { v.elements[index] = value }

// implementation of Dim method
func (v vector) Dim() int { return len(v.elements) }

// implementation of Copy method
func (v vectorComplex) Copy() Vector { return MakeVector(v.Dim(), v.Type()) }

// implementation of Trans method
func (v vectorBase) Trans() {
	if v.Type() == column {
		v.vectorType = RowVector
	} else {
		v.vectorType = ColVector
	}
}

// implementation of Norm method
func (v vector) Norm() float64 {
	var norm float64
	var dotProduct float64
	dotProduct, _ = operations.InnerProduct(v, v)
	norm = math.Sqrt(dotProduct)
	return norm
}

// MakeVector returns zero vector of size length
func MakeVector(length int, vectorType string) Vector {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]float64, length)
	return v
}

// MakeVectorWithElements returns zero vector of size length
func MakeVectorWithElements(length int, vectorType string, elements []float64) VectorComplex {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]float64, length)
	copy(vector.elements, elements)
	return v
}
