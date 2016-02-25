package calculators

import (
	"errors"
	"strconv"
)

// PolishCalculator is a simple reverse polish calculator
func PolishCalculator(args []string) (value float64, err error) {
	for left, right := 0, len(args)-1; left < right; left, right = left+1, right-1 {
		args[left], args[right] = args[right], args[left]
	}

	var stack []float64
	var operand1 float64
	var operand2 float64
	var result float64

	if len(args) <= 2 {
		err = errors.New("Not Enough Arguments")
		return
	}

	for _, argument := range args {
		if f, err := strconv.ParseFloat(argument, 64); err == nil {
			stack = append(stack, f)
			continue
		}

		if len(stack) > 1 {
			operand1, stack = popFloat64(stack)
			operand2, stack = popFloat64(stack)

			result, err = calculate(operand1, operand2, argument)

			if err != nil {
				return
			}

			stack = append(stack, result)
		} else {
			err = errors.New("IndexOutOfBoundsException")
			return
		}
	}

	if len(stack) != 1 {
		err = errors.New("Multiple Final Values")
		return
	}

	value = stack[0]
	return
}
