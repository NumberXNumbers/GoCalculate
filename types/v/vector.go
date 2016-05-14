package v

import (
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
)

// Space is the space in which a vector lives. Either Column or Row space.
type Space int

const (
	// RowSpace is a space Row Vector
	RowSpace Space = iota
	// ColSpace is a space Column Vector
	ColSpace
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
	Space() Space

	// Returns the core value type of the vector
	Type() gcv.Type

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
	space    Space
	coreType gcv.Type
	elements gcv.Values
}

// implementation of Len method
func (v *vector) Len() int { return v.elements.Len() }

// implementation of Len method
func (v *vector) Get(index int) gcv.Value { return v.elements.Get(index) }

// implementation of Set method
func (v *vector) Set(index int, val gcv.Value) {
	if v.coreType < val.Type() {
		v.coreType = val.Type()
	}
	v.elements.Set(index, val)
}

// implementation of Space method
func (v *vector) Space() Space { return v.space }

// implementation of Type method
func (v *vector) Type() gcv.Type { return v.elements.Type() }

// implementation of Copy method
func (v *vector) Copy() Vector { return MakeVectorAlt(v.Space(), v.Elements()) }

// implementation of Elements method
func (v *vector) Elements() gcv.Values { return v.elements }

// implementation of Append method
func (v *vector) Append(val gcv.Value) { v.Elements().Append(val) }

// implementation of Trans method
func (v *vector) Trans() {
	if v.Type() == gcv.Complex {
		for i := 0; i < v.Len(); i++ {
			value := v.Get(i)
			value.SetValue(gcvops.Conj(value))
			v.Set(i, value)
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
	dotProduct := gcv.NewValue()
	conjVector := MakeConjVector(v)
	for i := 0; i < v.Len(); i++ {
		dotProduct = gcvops.Add(dotProduct, gcvops.Mult(v.Get(i), conjVector.Get(i)))
	}
	return gcvops.Sqrt(dotProduct)
}

// implementation of IndexOf method
func (v *vector) IndexOf(val gcv.Value) int { return v.Elements().IndexOf(val) }

// NewVector returns zero vector of size length
func NewVector(space Space, length int) Vector {
	vector := new(vector)
	vector.space = space
	vector.coreType = gcv.Real
	elements := make([]gcv.Value, length)
	for index := range elements {
		val := gcv.MakeValue(0)
		elements[index] = val
	}
	vector.elements = gcv.MakeValues(elements...)
	return vector
}

// MakeVectorAlt returns Vector, requires a framework Values type
func MakeVectorAlt(space Space, elements gcv.Values) Vector {
	vector := new(vector)
	vector.space = space
	vector.coreType = elements.Type()
	vector.elements = elements.Copy()
	return vector
}

// MakeVectorPure returns Vector, requires a pure interfaces
func MakeVectorPure(space Space, elements ...interface{}) Vector {
	values := gcv.MakeValuesPure(elements...)
	vector := MakeVectorAlt(space, values)
	return vector
}

// MakeVector returns Vector, requires a framework Value type
func MakeVector(space Space, elements ...gcv.Value) Vector {
	values := gcv.MakeValuesAlt(elements)
	vector := MakeVectorAlt(space, values)
	return vector
}

// MakeConjVector returns a new conjugate vector of vector
func MakeConjVector(v Vector) Vector {
	conjVector := v.Copy()
	conjVector.Trans()
	return conjVector
}
