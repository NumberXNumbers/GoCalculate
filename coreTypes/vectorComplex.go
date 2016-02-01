package coreTypes

import (
	"math"
	"math/cmplx"

	"github.com/traviscox1990/GoCalculate/operations"
)

// VectorComplex is the main vector interface for real vectors
type VectorComplex interface {
	baseVector

	// Get returns element at index in vector
	Get(index int) complex128

	// Set will set a value at index in a vector
	Set(index int, value complex128)

	// Returns a new Copy of vector
	Copy() VectorComplex
}

type vectorComplex struct {
	vectorBase
	elements []complex128
}

// implementation of Get method
func (v vectorComplex) Get(index int) complex128 { return v.elements[index] }

// implementation of Set method
func (v vectorComplex) Set(index int, value complex128) { v.elements[index] = value }

// implementation of Dim method
func (v vectorComplex) Dim() int { return len(v.elements) }

// implementation of Copy method
func (v vectorComplex) Copy() VectorComplex {
	return MakeComplexVectorWithElements(v.Dim(), v.Type(), v.GetElements())
}

// implementation of Trans method
func (v vectorBase) Trans() {
	if v.Type() == column {
		v.vectorType = RowVector
	} else {
		v.vectorType = ColVector
	}

	for i := 0; i < v.Dim(); i++ {
		v[i] = cmplx.Conj(v[i])
	}
}

// implementation of Norm method
func (v vectorComplex) Norm() complex128 {
	var norm complex128
	var dotProduct complex128
	conjVector := MakeNewConjVector(v)
	dotProduct, _ = operations.InnerProductComplex(v, conjVector)
	norm = math.Sqrt(dotProduct)
	return norm
}

// MakeComplexVector returns zero vector of size length
func MakeComplexVector(length int, vectorType string) VectorComplex {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]complex128, length)
	return v
}

// MakeComplexVectorWithElements returns zero vector of size length
func MakeComplexVectorWithElements(length int, vectorType string, elements []complex128) VectorComplex {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]complex128, length)
	copy(vector.elements, elements)
	return v
}

// MakeNewConjVector returns a new conj vector of vector
func MakeNewConjVector(v vectorComplex) VectorComplex {
	conjVector := v.Copy()
	conjVector.Trans()
	return conjVector
}
