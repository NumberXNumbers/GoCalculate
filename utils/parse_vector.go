package utils

import (
	"errors"
	"strings"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// StringToVectorParser takes a string and returns a vector if string is
// of vector format, else error.
// Example of matrix format: [1 2 3 2+2i 2 0 3.0 2.3 0+3i]
func StringToVectorParser(s string) (vector v.Vector, err error) {
	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		err = errors.New("String is not of type Vector")
		return
	}

	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")

	// vSlice stands for vector string slice
	vSlice := strings.Split(s, " ")
	var value gcv.Value

	vector = v.NewVector(v.RowSpace, len(vSlice))
	for index, val := range vSlice {
		value, err = StringToValueParser(val)
		if err != nil {
			return
		}
		vector.Set(index, value.Copy())
	}
	return
}
