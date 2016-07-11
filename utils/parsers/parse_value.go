package parsers

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

var (
	isComplex = regexp.MustCompile(`^\-?\d+\.*\d*[\+\-]\d+\.*\d*i$`)
	plus      = "+"
	minus     = "-"
)

// Value will parses inputs from the command line into gcv Value types
func Value(s string) (output gcv.Value, err error) {
	if f, e := strconv.ParseFloat(s, 64); e == nil {
		output = gcv.MakeValue(f)
		return
	} else if isComplex.MatchString(s) {
		s = s[:len(s)-1]
		var sSlice []string
		count := strings.Count(s, minus)
		if count == 2 {
			s = s[1:len(s)]
			sSlice = strings.Split(s, minus)
			sSlice[0] = "-" + sSlice[0]
			sSlice[1] = "-" + sSlice[1]
		} else if count == 1 && strings.Index(s, minus) == 0 {
			sSlice = strings.Split(s, plus)
		} else if count == 1 {
			sSlice = strings.Split(s, minus)
			sSlice[1] = "-" + sSlice[1]
		} else {
			sSlice = strings.Split(s, plus)
		}

		var real float64
		var image float64
		if f, e := strconv.ParseFloat(sSlice[0], 64); e == nil {
			real = f
		}
		if f, e := strconv.ParseFloat(sSlice[1], 64); e == nil {
			image = f
		}
		output = gcv.MakeValue(complex(real, image))
		return
	}
	err = errors.New("String is not type real or complex")
	return
}
