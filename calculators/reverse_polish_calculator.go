package calculators

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/utils/parsers"
)

// ReversePolishCalculator is a simple reverse polish calculator
func ReversePolishCalculator(args []string) (value gcv.Value, err error) {
	stack := gcv.MakeValues()
	var operand1 gcv.Value
	var operand2 gcv.Value
	var result gcv.Value

	if len(args) <= 2 {
		err = errors.New("Not Enough Arguments")
		return
	}

	for _, argument := range args {
		if v, e := parsers.Value(argument); e == nil {
			stack.Append(v)
			continue
		}

		if stack.Len() > 1 {
			operand2, stack = pop(stack)
			operand1, stack = pop(stack)

			result, err = calculateV(operand1, operand2, argument)

			if err != nil {
				return
			}

			stack.Append(result)
		} else {
			err = errors.New("IndexOutOfBoundsException")
			return
		}
	}

	if stack.Len() != 1 {
		err = errors.New("Multiple Final Values")
		return
	}

	value = stack.Get(0)
	return
}
