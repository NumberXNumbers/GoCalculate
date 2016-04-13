package v

import (
	"math"
	"math/cmplx"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

const (
	// ColVector is a vectorType Column Vector
	ColVector = "column"
	// RowVector is a vectorType Row Vector
	RowVector = "row"
)

// Vector is the main vector interface for real vectors
type Vector interface {
	// Returns the index of vect. If vect is not in Values it returns -1.
	IndexOf(val gcv.Value) int

	// Returns the Value at index
	Get(index int) gcv.Value

	// Set the Value val at index
	Set(index int, val gcv.Value)

	// Return norm of vector
	Norm() gcv.Value

	// Returns a new Copy of vector
	Copy() Vector

	// Returns Type of vector. Either Column or Row vector
	Type() string

	// Returns the len of the values
	Len() int

	// Transpose of vector. i.e changes a row into a column vector and vice versa
	Trans()

	// Returns the elements in the vector
	Elements() gcv.Values
}

type vector struct {
	vectorType string
	elements   gcv.Values
}

// implementation of Len method
func (v *vector) Len() int { return v.elements.Len() }

// implementation of Len method
func (v *vector) Get(index int) gcv.Value { return v.elements.Get(index) }

func (v *vector) Set(index int, val gcv.Value) { v.elements.Set(index, val) }

// implementation of Type method
func (v *vector) Type() string { return v.vectorType }

// implementation of Copy method
func (v *vector) Copy() Vector { return MakeVectorWithElements(v.elements, v.Type()) }

// implementation of Elements method
func (v *vector) Elements() gcv.Values { return v.elements }

// implementation of Trans method
func (v *vector) Trans() {
	if v.elements.CoreType() == gcv.Complex {
		for i := 0; i < v.Len(); i++ {
			v.Set(i, gcv.NewValue(cmplx.Conj(v.Get(i).Complex128())))
		}
	}

	if v.Type() == ColVector {
		v.vectorType = RowVector
	} else {
		v.vectorType = ColVector
	}
}

// implementation of Norm method
func (v *vector) Norm() gcv.Value {
	if v.elements.CoreType() == gcv.Complex {
		var dotProduct complex128
		conjVector := MakeNewConjVector(v)
		for i := 0; i < v.Len(); i++ {
			dotProduct += v.Get(i).Complex128() * conjVector.Get(i).Complex128()
		}
		return gcv.NewValue(cmplx.Sqrt(dotProduct))
	}

	var dotProduct float64
	for i := 0; i < v.Len(); i++ {
		dotProduct += v.Get(i).Float64() * v.Get(i).Float64()
	}
	return gcv.NewValue(math.Sqrt(dotProduct))
}

// implementation of IndexOf method
func (v *vector) IndexOf(val gvc.Value) int { return v.Elements().IndexOf(val) }

// MakeVector returns zero vector of size length
func MakeVector(length int, vectorType string) Vector {
	vector := new(vector)
	vector.vectorType = vectorType
	emptyElements := make([]gcv.Value, length)
	for index := range emptyElements {
		val := gcv.NewValue(0)
		emptyElements[index] = val
	}
	vector.elements = gcv.NewValues(emptyElements...)
	return vector
}

// MakeVectorWithElements returns zero vector of size length
func MakeVectorWithElements(elements gcv.Values, vectorType string) Vector {
	vector := new(vector)
	vector.vectorType = vectorType
	vector.elements = elements.Copy()
	return vector
}

// MakeNewConjVector returns a new conjugate vector of vector
func MakeNewConjVector(v Vector) Vector {
	conjVector := v.Copy()
	conjVector.Trans()
	return conjVector
}
