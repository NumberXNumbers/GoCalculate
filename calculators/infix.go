package calculators

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/utils"
)

// InfixCalculator will calculate an infix calculation
func InfixCalculator(args []string) (value gcv.Value, err error) {
	stack := gcv.MakeValues()
	var operatorStack []string
	var operand1 gcv.Value
	var operand2 gcv.Value
	var result gcv.Value

	var leftParens uint
	var rightParens uint

	var subArgs []string
	var arg string

	index := 0
	for index < len(args) {
		arg = args[index]
		if v, e := utils.StringToValueParser(arg); e == nil {
			stack.Append(v)
		} else if arg == "(" {
			leftParens++
			for i := index + 1; i < len(args); i++ {
				if args[i] == ")" {
					rightParens++
				} else if args[i] == "(" {
					leftParens++
				}
				if rightParens != leftParens {
					subArgs = append(subArgs, args[i])
					continue
				}
				break
			}
			index = index + len(subArgs) + 1
			result, err = InfixCalculator(subArgs)
			if err != nil {
				return
			}
			stack.Append(result)
		} else {
			operatorStack = append(operatorStack, arg)
		}

		if stack.Len() > 1 && len(operatorStack) > 0 {
			operand1, stack = pop(stack)
			operand2, stack = pop(stack)

			operator := operatorStack[0]
			operatorStack = operatorStack[1:]
			result, err = calculate(operand2, operand1, operator)

			if err != nil {
				return
			}

			stack.Append(result)
		}

		if stack.Len() == 1 && len(operatorStack) > 1 {
			err = errors.New("IndexOutOfBoundsException")
			return
		}
		index++
	}

	value = stack.Get(0)
	return
}
