package v

import "github.com/NumberXNumbers/GoCalculate/types/gcv"

// Vectors returns the set of vectors in a particular vector space
type Vectors interface {
	// Returns the index of vect. If vect is not in Values it returns -1.
	IndexOf(vect Vector) int

	// Set the Vector vect at index. If vect is not in the space of other vectors
	// Trans() method will be called on it.
	Set(index int, vect Vector)

	// Set value j of vector i in vectors to val
	SetValue(i int, j int, val gcv.Value)

	// Append Vector vect to Vectors
	Append(vect Vector)

	// Returns the Vector at index
	Get(index int) Vector

	// Returns the Raw Vector slice. Used mainly for use with range
	Vectors() []Vector

	// Returns a subset of Vectors from start to finish
	Subset(start, finish int) Vectors

	// Returns a copy of Vectors.
	Copy() Vectors

	// Returns the length of Vectors slice. this is the same as the number of vectors in Vectors
	Len() int

	// Returns the length of the individual vectors in Vectors
	InnerLen() int

	// Returns the space that Vectors is in
	Space() string

	// Returns the highest ranking vector type
	Type() string
}

type vectors struct {
	vects       []Vector
	length      int
	innerLength int
	space       string
	coreType    string
}

func (v *vectors) setVectors(vects []Vector, space string) {
	v.length = len(vects)
	if space != RowSpace && space != ColSpace {
		space = RowSpace
	}
	v.space = space
	v.coreType = gcv.Int
	v.vects = make([]Vector, v.Len())
	for index, vect := range vects {
		copyVect := vect.Copy()
		if len(v.Type()) < len(copyVect.Type()) {
			v.coreType = vect.Type()
		}
		if v.InnerLen() < copyVect.Len() {
			v.innerLength = vect.Len()
		}
		if copyVect.Space() != space {
			copyVect.Trans()
		}
		v.vects[index] = copyVect
	}
}

func (v *vectors) Len() int { return v.length }

func (v *vectors) InnerLen() int { return v.innerLength }

func (v *vectors) Space() string { return v.space }

func (v *vectors) Type() string { return v.coreType }

func (v *vectors) Get(index int) Vector { return v.vects[index] }

func (v *vectors) Vectors() []Vector { return v.vects }

func (v *vectors) Set(index int, vect Vector) {
	if len(v.Type()) < len(vect.Type()) {
		v.coreType = vect.Type()
	}
	if vect.Space() != v.Space() {
		vect.Trans()
	}
	v.vects[index] = vect
}

func (v *vectors) SetValue(i int, j int, value gcv.Value) {
	if len(v.Type()) < len(value.GetValueType()) {
		v.coreType = value.GetValueType()
	}
	vector := v.vects[i]
	values := vector.Elements()
	values.Values()[j] = value
}

func (v *vectors) Append(vect Vector) {
	v.setVectors(append(v.Vectors(), vect), v.Space())
}

func (v *vectors) Copy() Vectors {
	vects := new(vectors)
	vElements := make([]Vector, len(v.vects))
	for index, vect := range v.Vectors() {
		vElements[index] = vect.Copy()
	}
	vects.length = v.Len()
	vects.innerLength = v.InnerLen()
	vects.coreType = v.Type()
	vects.space = v.Space()
	vects.vects = vElements
	return vects
}

func (v *vectors) Subset(start, finish int) Vectors {
	vects := new(vectors)
	subVects := make([]Vector, len(v.vects[start:finish+1]))
	copy(subVects, v.vects[start:finish+1])
	vects.setVectors(subVects, v.Space())
	return vects
}

func (v *vectors) IndexOf(vect Vector) int {
	found := false
	values := vect.Elements().Values()
	for index, vector := range v.Vectors() {
		if vect.Len() != vector.Len() {
			continue
		}
		for valIndex, value := range values {
			tempValue := vector.Get(valIndex)
			if value.GetValueType() == gcv.Complex && value.Complex128() != tempValue.Complex128() {
				found = false
				break
			}
			if value.GetValueType() == gcv.Float && value.Float64() != tempValue.Float64() {
				found = false
				break
			}
			if value.GetValueType() == gcv.Int && value.Int() != tempValue.Int() {
				found = false
				break
			}
			found = true
		}
		if found {
			return index
		}
	}
	return -1
}

// NewVectors will return the Zero Vectors. or numVectors of zero vectors of length lenVectors
func NewVectors(space string, numVectors, lenVectors int) Vectors {
	vectors := new(vectors)
	vects := make([]Vector, numVectors)
	vectors.length = numVectors
	vectors.innerLength = lenVectors
	vectors.vects = vects
	if space != RowSpace && space != ColSpace {
		space = RowSpace
	}
	vectors.space = space
	vectors.coreType = gcv.Int
	for i := 0; i < vectors.Len(); i++ {
		vect := NewVector(space, lenVectors)
		vectors.Set(i, vect)
	}
	return vectors
}

// MakeVectors will return a Vectors type. All vectors will be in vector space, space.
// If inputed vector is not in that vector space, Trans() will be called on it
func MakeVectors(space string, vects ...Vector) Vectors {
	vectors := new(vectors)
	vectors.setVectors(vects, space)
	return vectors
}
