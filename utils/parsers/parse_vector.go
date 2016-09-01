package parsers

import (
	"errors"
	"strings"

	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

var (
	leftBracket       = "["
	rightBracket      = "]"
	apostrophe        = "'"
	graveAccent       = "`"
	star              = "*"
	offsetGraveAccent = 1
	offsetApostrophe  = 2
	offsetStar        = 2
)

// Vector takes a string and returns a vector if string is
// of vector format, else error.
// Example of matrix format: [1 2 3 2+2i 2 0 3.0 2.3 0+3i] for row vector
// or [1 2 3 2+2i 2 0 3.0 2.3 0+3i]' for the transpose column vector
// or [1 2 3 2+2i 2 0 3.0 2.3 0+3i]* for the conjugate transpose column vector
func Vector(s string) (vector v.Vector, err error) {
	if !strings.HasPrefix(s, leftBracket) {
		err = errors.New("String is not of type Vector")
		return
	}

	if strings.Count(s, leftBracket) != 1 ||
		strings.Count(s, rightBracket) != 1 {
		err = errors.New("String is not of type Vector")
		return
	}

	var transposeVector bool
	var conjTransposeVector bool

	if strings.HasSuffix(s, apostrophe) {
		if strings.Index(s, rightBracket) == len(s)-offsetApostrophe {
			transposeVector = true
			s = strings.TrimRight(s, apostrophe)
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	if strings.HasSuffix(s, star) {
		if strings.Index(s, rightBracket) == len(s)-offsetStar {
			conjTransposeVector = true
			s = strings.TrimRight(s, star)
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	s = strings.TrimLeft(s, leftBracket)
	s = strings.TrimRight(s, rightBracket)

	// vSlice stands for vector string slice
	vSlice := strings.Split(s, " ")
	var value gcv.Value

	vector = v.NewVector(v.RowSpace, len(vSlice))

	for index, val := range vSlice {
		value, err = Value(val)
		if err != nil {
			return
		}
		vector.Set(index, value.Copy())
	}

	if transposeVector {
		vector.Trans()
	}

	if conjTransposeVector {
		vector.ConjTrans()
	}

	return
}
