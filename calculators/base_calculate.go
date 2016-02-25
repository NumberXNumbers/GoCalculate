package calculators

import (
	"errors"
	"math"
)

func popFloat64(stack []float64) (float64, []float64) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

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
