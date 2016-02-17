package calculators

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func calculate(firstValue, secondValue float64, s string) float64 {
	switch s {
	case "+":
		return firstValue + secondValue
	case "x", "X":
		return firstValue * secondValue
	case "/":
		return firstValue / secondValue
	case "-":
		return firstValue - secondValue
	case "exp":
		return math.Pow(firstValue, secondValue)
	case "%":
		return math.Mod(firstValue, secondValue)
	default:
		log.Fatal("IllegalArgumentException")
	}
	return 0.0
}

func stackBuilder(stack []float64, s string) []float64 {
	size := len(stack)

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		stack = append(stack, f)
	} else {
		if size > 1 {
			value := calculate(stack[size-1], stack[size-2], s)
			stack = stack[:size-2]
			stack = append(stack, value)
		} else {
			log.Fatal("IndexOutOfBoundsException")
		}
	}

	return stack
}

// ReversePolishCalculator is a simple reverse polish calculator
func ReversePolishCalculator(args []string) {
	if len(args) == 0 {
		log.Fatal("Zero Arguments passed in.")
	}

	fmt.Print("Arguments passed in: ")
	fmt.Println(args)

	var stack []float64
	for _, argument := range args {
		stack = stackBuilder(stack, argument)
	}

	fmt.Print("Final Value(s): ")
	for _, value := range stack {
		fmt.Println(value)
	}
}
