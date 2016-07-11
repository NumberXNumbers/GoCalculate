package calculators

import (
	"github.com/NumberXNumbers/GoCalculate/types/gcf"
	"github.com/NumberXNumbers/GoCalculate/utils/parsers"
)

var (
	// C style order of opertations. Mod is of same level as multiplication and division
	orderOfOperations = map[string]uint{
		"exp": 3,
		"*":   2,
		"x":   2,
		"/":   2,
		"%":   2,
		"+":   1,
		"-":   1,
	}
)

// InfixCalculator will calculate an infix calculation
func InfixCalculator(args []string) gcf.Const {
	var argsForCalculation []interface{}

	for _, arg := range args {
		if v, e := parsers.Value(arg); e == nil {
			argsForCalculation = append(argsForCalculation, v)
		} else {
			argsForCalculation = append(argsForCalculation, arg)
		}
	}

	calculation := gcf.Calculate(argsForCalculation...)

	return calculation
}
