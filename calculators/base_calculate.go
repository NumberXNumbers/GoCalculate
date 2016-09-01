package calculators

import (
	"errors"

	gcv "github.com/NumberXNumbers/types/gc/values"
	gcvops "github.com/NumberXNumbers/types/gc/values/ops"
)

func pop(stack gcv.Values) (gcv.Value, gcv.Values) {
	return stack.Get(stack.Len() - 1), stack.Subset(0, stack.Len()-2)
}

func dequeue(stack gcv.Values) (gcv.Value, gcv.Values) {
	return stack.Get(0), stack.Subset(1, stack.Len()-1)
}

const (
	add    = "+"
	sub    = "-"
	times1 = "x"
	times2 = "X"
	times3 = "*"
	div    = "/"
	pow    = "exp"
	mod    = "%"
)

// calculateV will calculate two values together
func calculateV(firstValue, secondValue gcv.Value, s string) (result gcv.Value, err error) {
	switch s {
	case add:
		result = gcvops.Add(firstValue, secondValue)
	case times1, times2, times3:
		result = gcvops.Mult(firstValue, secondValue)
	case div:
		result = gcvops.Div(firstValue, secondValue)
	case sub:
		result = gcvops.Sub(firstValue, secondValue)
	case pow:
		result = gcvops.Pow(firstValue, secondValue)
	case mod:
		result, err = gcvops.Mod(firstValue, secondValue)
	default:
		err = errors.New("IllegalArgumentException")
	}
	return
}
