package parsers

import (
	"errors"
	"strings"

	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
)

// Matrix takes a string and returns a matrix if string is
// of matrix format, else error.
// Example of matrix format: [1 2 3: 2+2i 2 0: 3.0 2.3 0+3i]
// [1 2 3: 2+2i 2 0: 3.0 2.3 0+3i]' will return the transpose of the matrix
// [1 2 3: 2+2i 2 0: 3.0 2.3 0+3i]* will return the conjugate transpose of the matrix
func Matrix(s string) (matrix m.Matrix, err error) {
	if !strings.HasPrefix(s, leftBracket) {
		err = errors.New("String is not of type Matrix")
		return
	}

	if strings.Count(s, leftBracket) != 1 || strings.Count(s, rightBracket) != 1 {
		err = errors.New("String is not of type Matrix")
		return
	}

	if strings.Count(s, ":") == 0 {
		err = errors.New("String is of type Vector")
		return
	}

	var transposeMatrix bool
	var conjTransposeMatrix bool

	if strings.HasSuffix(s, apostrophe) {
		if strings.Index(s, rightBracket) == len(s)-offsetApostrophe {
			transposeMatrix = true
			s = strings.TrimRight(s, apostrophe)
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	if strings.HasSuffix(s, star) {
		if strings.Index(s, rightBracket) == len(s)-offsetStar {
			conjTransposeMatrix = true
			s = strings.TrimRight(s, star)
		} else {
			err = errors.New("String is not of type Vector")
			return
		}
	}

	s = strings.TrimLeft(s, leftBracket)
	s = strings.TrimRight(s, rightBracket)
	s = strings.Replace(s, ": ", ":", -1)
	s = strings.Replace(s, " ", ",", -1)
	sSlice := strings.Split(s, ":")

	// mSlice stands for matrix string slice
	var mSlice [][]string

	for _, subSlice := range sSlice {
		mSlice = append(mSlice, strings.Split(subSlice, ","))
	}

	length := len(mSlice[0])
	matrix = m.NewMatrix(len(mSlice), length)
	var value gcv.Value

	// vSlice stands for vector string slice
	for indexI, vSlice := range mSlice {
		if len(vSlice) != length {
			err = errors.New("Inconsistent lengths")
			return
		}
		for indexJ, val := range vSlice {
			value, err = Value(val)
			if err != nil {
				return
			}
			matrix.Set(indexI, indexJ, value.Copy())
		}
	}

	if transposeMatrix {
		matrix.Trans()
	}

	if conjTransposeMatrix {
		matrix.ConjTrans()
	}

	return
}
