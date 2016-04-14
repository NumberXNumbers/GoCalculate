package v

import "github.com/NumberXNumbers/GoCalculate/types/gcv"

const (
	// RowSpace are vectors in row space
	RowSpace = "row"
	// ColSpace are vectors in column space
	ColSpace = "column"
)

// Vectors returns the set of vectors in a particular vector space
type Vectors interface {
	// Returns the index of vect. If vect is not in Values it returns -1.
	IndexOf(vect Vector) int

	// Set the Vector vect at index. If vect is not in the space of other vectors
	// Trans() method will be called on it.
	Set(index int, vect Vector)

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

	// Returns the length of Vectors
	Len() int

	// Returns the space that Vectors is in
	Space() string
}

type vectors struct {
	vects  []Vector
	length int
	space  string
}

func (v *vectors) setVectors(vects []Vector, space string) {
	v.length = len(vects)
	v.space = space
	for index, vect := range vects {
		if vect.Type() != space {
			vect.Trans()
			vects[index] = vect
		}
	}
	v.vects = vects
}

func (v *vectors) Len() int { return v.length }

func (v *vectors) Space() string { return v.space }

func (v *vectors) Get(index int) Vector { return v.vects[index] }

func (v *vectors) Vectors() []Vector { return v.vects }

func (v *vectors) Set(index int, vect Vector) {
	if vect.Type() != v.Space() {
		vect.Trans()
	}
	v.vects[index] = vect
}

func (v *vectors) Append(vect Vector) {
	v.setVectors(append(v.Vectors(), vect), v.Space())
}

func (v *vectors) Copy() Vectors {
	vects := new(vectors)
	vElements := make([]Vector, len(v.vects))
	copy(vElements, v.vects)
	vects.setVectors(vElements, v.Space())
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

// NewVectors will return a Vectors type. All vectors will be in vector space, space.
// If inputed vector is not in that vector space, Trans() will be called on it
func NewVectors(space string, vects ...Vector) Vectors {
	newVectors := new(vectors)
	newVectors.setVectors(vects, space)
	return newVectors
}
