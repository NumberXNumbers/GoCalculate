package parsers

import (
	"errors"
	"strings"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

var (
	leftBracket       = "["
	rightBracket      = "]"
	apostrophe        = "'"
	graveAccent       = "`"
	offsetGraveAccent = 1
	offsetApostrophe  = 2
)

// Vector takes a string and returns a vector if string is
// of vector format, else error.
// Example of matrix format: [1 2 3 2+2i 2 0 3.0 2.3 0+3i] for row vector
// or [1 2 3 2+2i 2 0 3.0 2.3 0+3i]' for the transpose column vector
// or `[1 2 3 2+2i 2 0 3.0 2.3 0+3i] for the column vector
// or `[1 2 3 2+2i 2 0 3.0 2.3 0+3i]' for the transpose row vector
func Vector(s string) (vector v.Vector, err error) {
	if strings.Count(s, leftBracket) != 1 ||
		strings.Count(s, rightBracket) != 1 {
		err = errors.New("String is not of type Vector")
		return
	}

	var isColumnVector bool
	var transposeVector bool

	if strings.HasPrefix(s, graveAccent) {
		if strings.Index(s, leftBracket) == offsetGraveAccent {
			isColumnVector = true
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	if strings.HasSuffix(s, apostrophe) {
		if strings.Index(s, rightBracket) == len(s)-offsetApostrophe {
			transposeVector = true
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	if (!transposeVector && !strings.HasSuffix(s, rightBracket)) ||
		(!isColumnVector && !strings.HasPrefix(s, leftBracket)) {
		err = errors.New("String is not of type Vector")
		return
	}

	s = strings.TrimLeft(s, leftBracket)
	s = strings.TrimRight(s, rightBracket)

	// vSlice stands for vector string slice
	vSlice := strings.Split(s, " ")
	var value gcv.Value

	if isColumnVector {
		vector = v.NewVector(v.ColSpace, len(vSlice))
	} else {
		vector = v.NewVector(v.RowSpace, len(vSlice))
	}

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

	return
}
