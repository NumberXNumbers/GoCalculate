package utils

import (
	"errors"
	"strings"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// StringToMatrixParser takes a string and returns a matrix if string is
// of matrix format, else error.
// Example of matrix format: [1 2 3: 2+2i 2 0: 3.0 2.3 0+3i]
func StringToMatrixParser(s string) (matrix m.Matrix, err error) {
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

	var vector v.Vector
	var vectors []v.Vector
	var value gcv.Value

	length := len(mSlice[0])
	// vSlice stands for vector string slice
	for _, vSlice := range mSlice {
		if len(vSlice) != length {
			err = errors.New("Inconsistent lengths")
			return
		}
		vector = v.NewVector(v.RowSpace, length)
		for index, val := range vSlice {
			value, err = StringToValueParser(val)
			if err != nil {
				return
			}
			vector.Set(index, value.Copy())
		}
		vectors = append(vectors, vector.Copy())
	}

	matrix = m.MakeMatrix(vectors...)
	// fmt.Println(matrix)
	return
}
