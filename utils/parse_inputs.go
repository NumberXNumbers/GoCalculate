package utils

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

var (
	isComplex = regexp.MustCompile(`^\d+\.*\d*[\+\-]\d+\.*\d*i$`)
	plusMinus = regexp.MustCompile(`[\+\-]`)
)

// StringToValueParser will parses inputs from the command line for real functions
// TODO: this needs to be implemented. Below is place holder code.
func StringToValueParser(s string) (output gcv.Value, err error) {
	if f, e := strconv.ParseFloat(s, 64); e == nil {
		output = gcv.MakeValue(f)
		return
	} else if isComplex.MatchString(s) {
		s = s[:len(s)-1]
		sSlice := plusMinus.Split(s, 2)
		var real float64
		var image float64
		if i, e := strconv.ParseInt(sSlice[0], 10, 64); e == nil {
			real = float64(i)
		} else if f, e := strconv.ParseFloat(sSlice[0], 64); e == nil {
			real = f
		}
		if i, e := strconv.ParseInt(sSlice[1], 10, 64); e == nil {
			image = float64(i)
		} else if f, e := strconv.ParseFloat(sSlice[1], 64); e == nil {
			image = f
		}
		output = gcv.MakeValue(complex(real, image))
		return
	}
	err = errors.New("String is not type int, float or complex")
	return
}
