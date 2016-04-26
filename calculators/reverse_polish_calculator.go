package calculators

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/utils"
)

// ReversePolishCalculator is a simple reverse polish calculator
func ReversePolishCalculator(args []string) (value gcv.Value, err error) {
	stack := gcv.NewValues()
	var operand1 gcv.Value
	var operand2 gcv.Value
	var result gcv.Value

	if len(args) <= 2 {
		err = errors.New("Not Enough Arguments")
		return
	}

	for _, argument := range args {
		if v, e := utils.StringToValueParser(argument); e == nil {
			stack.Append(v)
			continue
		}

		if stack.Len() > 1 {
			operand1, stack = pop(stack)
			operand2, stack = pop(stack)

			result, err = calculate(operand2, operand1, argument)

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
