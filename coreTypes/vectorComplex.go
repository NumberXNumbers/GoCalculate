package coreTypes

import "math/cmplx"

// VectorComplex is the main vector interface for real vectors
type VectorComplex interface {
	baseVector

	// Returns Elements of VectorComplex
	GetElements() []complex128

	// Get returns element at index in vector
	Get(index int) complex128

	// Set will set a value at index in a vector
	Set(index int, value complex128)

	// Return norm of vectorBase
	Norm() complex128

	// Returns a new Copy of vector
	Copy() VectorComplex
}

type vectorComplex struct {
	vectorBase
	elements []complex128
}

// implementation of Get method
func (v *vectorComplex) Get(index int) complex128 { return v.elements[index] }

// implementation of Set method
func (v *vectorComplex) Set(index int, value complex128) { v.elements[index] = value }

// implementation of Dim method
func (v *vectorComplex) Dim() int { return len(v.elements) }

// implementation of GetElements()
func (v *vectorComplex) GetElements() []complex128 { return v.elements }

// implementation of Copy method
func (v *vectorComplex) Copy() VectorComplex {
	return MakeComplexVectorWithElements(v.GetElements(), v.Type())
}

// implementation of Trans method
func (v *vectorComplex) Trans() {
	if v.Type() == ColVector {
		v.vectorType = RowVector
	} else {
		v.vectorType = ColVector
	}

	for i := 0; i < v.Dim(); i++ {
		v.elements[i] = cmplx.Conj(v.elements[i])
	}
}

// implementation of Norm method
func (v *vectorComplex) Norm() complex128 {
	var norm complex128
	var dotProduct complex128
	conjVector := MakeNewConjVector(v)
	for i := 0; i < v.Dim(); i++ {
		dotProduct += v.Get(i) * conjVector.Get(i)
	}
	norm = cmplx.Sqrt(dotProduct)
	return norm
}

// MakeComplexVector returns zero vector of size length
func MakeComplexVector(length int, vectorType string) VectorComplex {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]complex128, length)
	return vector
}

// MakeComplexVectorWithElements returns zero vector of size length
func MakeComplexVectorWithElements(elements []complex128, vectorType string) VectorComplex {
	vector := new(vectorComplex)
	vector.vectorType = vectorType
	vector.elements = make([]complex128, len(elements))
	copy(vector.elements, elements)
	return vector
}

// MakeNewConjVector returns a new conj vector of vector
func MakeNewConjVector(v VectorComplex) VectorComplex {
	conjVector := v.Copy()
	conjVector.Trans()
	return conjVector
}
