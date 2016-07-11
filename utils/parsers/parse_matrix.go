package parsers

import (
	"errors"
	"strings"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
)

// Matrix takes a string and returns a matrix if string is
// of matrix format, else error.
// Example of matrix format: [1 2 3: 2+2i 2 0: 3.0 2.3 0+3i]
func Matrix(s string) (matrix m.Matrix, err error) {
	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		err = errors.New("String is not of type Matrix")
		return
	}

	if strings.Count(s, "[") != 1 || strings.Count(s, "]") != 1 {
		err = errors.New("String is not of type Matrix")
		return
	}

	if strings.Count(s, ":") == 0 {
		err = errors.New("String is of type Vector")
		return
	}
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
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

	// fmt.Println(matrix)
	return
}
