package calculators

import (
	"errors"
	"math"
	"strconv"
)

func calculate(firstValue, secondValue float64, s string) (result float64, err error) {
	switch s {
	case "+":
		result = firstValue + secondValue
	case "x", "X", "*":
		result = firstValue * secondValue
	case "/":
		result = firstValue / secondValue
	case "-":
		result = firstValue - secondValue
	case "exp":
		result = math.Pow(firstValue, secondValue)
	case "%":
		result = math.Mod(firstValue, secondValue)
	default:
		err = errors.New("IllegalArgumentException")
	}
	return
}

func stackBuilder(stack []float64, s string) ([]float64, error) {
	size := len(stack)

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		stack = append(stack, f)
	} else {
		if size > 1 {
			value, err := calculate(stack[size-2], stack[size-1], s)
			if err != nil {
				return stack, err
			}
			stack = stack[:size-2]
			stack = append(stack, value)
		} else {
			return stack, errors.New("IndexOutOfBoundsException")
		}
	}

	return stack, nil
}

// ReversePolishCalculator is a simple reverse polish calculator
func ReversePolishCalculator(args []string) (value float64, err error) {
	var stack []float64

	if len(args) <= 2 {
		err = errors.New("Not Enough Arguments")
		return
	}

	for _, argument := range args {
		stack, err = stackBuilder(stack, argument)
		if err != nil {
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
