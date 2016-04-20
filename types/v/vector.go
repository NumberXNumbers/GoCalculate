package v

import (
	"math"
	"math/cmplx"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

const (
	// ColSpace is a space Column Vector
	ColSpace = "column"
	// RowSpace is a space Row Vector
	RowSpace = "row"
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

	// Returns Space of vector. Either Column or Row vector
	Space() string

	// Returns the core value type of the vector
	Type() string

	// Returns the len of the values
	Len() int

	// Transpose of vector. i.e changes a row into a column vector and vice versa
	Trans()

	// Append a value to vector
	Append(val gcv.Value)

	// Returns the elements in the vector
	Elements() gcv.Values
}

type vector struct {
	space    string
	coreType string
	elements gcv.Values
}

// implementation of Len method
func (v *vector) Len() int { return v.elements.Len() }

// implementation of Len method
func (v *vector) Get(index int) gcv.Value { return v.elements.Get(index) }

// implementation of Set method
func (v *vector) Set(index int, val gcv.Value) {
	if len(v.coreType) < len(val.GetValueType()) {
		v.coreType = val.GetValueType()
	}
	v.elements.Set(index, val)
}

// implementation of Space method
func (v *vector) Space() string { return v.space }

// implementation of Type method
func (v *vector) Type() string { return v.elements.Type() }

// implementation of Copy method
func (v *vector) Copy() Vector { return MakeVector(v.Space(), v.elements) }

// implementation of Elements method
func (v *vector) Elements() gcv.Values { return v.elements }

// implementation of Append method
func (v *vector) Append(val gcv.Value) { v.Elements().Append(val) }

// implementation of Trans method
func (v *vector) Trans() {
	if v.Type() == gcv.Complex {
		for i := 0; i < v.Len(); i++ {
			v.Set(i, gcv.NewValue(cmplx.Conj(v.Get(i).Complex128())))
		}
	}

	if v.Space() == ColSpace {
		v.space = RowSpace
	} else {
		v.space = ColSpace
	}
}

// implementation of Norm method
func (v *vector) Norm() gcv.Value {
	if v.Type() == gcv.Complex {
		var dotProduct complex128
		conjVector := MakeConjVector(v)
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
func (v *vector) IndexOf(val gcv.Value) int { return v.Elements().IndexOf(val) }

// NewVector returns zero vector of size length
func NewVector(space string, length int) Vector {
	vector := new(vector)
	if space != RowSpace && space != ColSpace {
		space = RowSpace
	}
	vector.space = space
	vector.coreType = gcv.Int
	elements := make([]gcv.Value, length)
	for index := range elements {
		val := gcv.NewValue(0)
		elements[index] = val
	}
	vector.elements = gcv.NewValues(elements...)
	return vector
}

// MakeVector returns zero vector of size length
func MakeVector(space string, elements gcv.Values) Vector {
	vector := new(vector)
	if space != RowSpace && space != ColSpace {
		space = RowSpace
	}
	vector.space = space
	vector.coreType = elements.Type()
	vector.elements = elements.Copy()
	return vector
}

// MakeConjVector returns a new conjugate vector of vector
func MakeConjVector(v Vector) Vector {
	conjVector := v.Copy()
	conjVector.Trans()
	return conjVector
}
