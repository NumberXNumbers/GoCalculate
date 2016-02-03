package coreTypes

import "math"

// Vector is the main vector interface for real vectors
type Vector interface {
	baseVector

	// Returns Elements of Vector
	GetElements() []float64

	// Get returns element at index in vector
	Get(index int) float64

	// Set will set a value at index in a vector
	Set(index int, value float64)

	// Return norm of vector
	Norm() float64

	// Returns a new Copy of vector
	Copy() Vector
}

type vector struct {
	vectorBase
	elements []float64
}

// implementation of Get method
func (v *vector) Get(index int) float64 { return v.elements[index] }

// implementation of Set method
func (v *vector) Set(index int, value float64) { v.elements[index] = value }

// implementation of Dim method
func (v *vector) Dim() int { return len(v.elements) }

// implementation of GetElements method
func (v *vector) GetElements() []float64 { return v.elements }

// implementation of Copy method
func (v *vector) Copy() Vector { return MakeVectorWithElements(v.GetElements(), v.Type()) }

// implementation of Trans method
func (v *vector) Trans() {
	if v.Type() == ColVector {
		v.vectorType = RowVector
	} else {
		v.vectorType = ColVector
	}
}

// implementation of Norm method
func (v *vector) Norm() float64 {
	var norm float64
	var dotProduct float64
	for i := 0; i < v.Dim(); i++ {
		dotProduct += v.Get(i) * v.Get(i)
	}
	norm = math.Sqrt(dotProduct)
	return norm
}

// MakeVector returns zero vector of size length
func MakeVector(length int, vectorType string) Vector {
	vector := new(vector)
	vector.vectorType = vectorType
	vector.elements = make([]float64, length)
	return vector
}

// MakeVectorWithElements returns zero vector of size length
func MakeVectorWithElements(elements []float64, vectorType string) Vector {
	vector := new(vector)
	vector.vectorType = vectorType
	vector.elements = make([]float64, len(elements))
	copy(vector.elements, elements)
	return vector
}
