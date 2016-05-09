package calculators

import (
	"errors"
	"math"
	"math/cmplx"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

func pop(stack gcv.Values) (gcv.Value, gcv.Values) {
	return stack.Get(stack.Len() - 1), stack.Subset(0, stack.Len()-2)
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

func calculate(firstValue, secondValue gcv.Value, s string) (result gcv.Value, err error) {
	if firstValue.GetValueType() == gcv.Complex || secondValue.GetValueType() == gcv.Complex {
		result, err = calculateComplex(firstValue.Complex128(), secondValue.Complex128(), s)
	} else if firstValue.GetValueType() == gcv.Float || secondValue.GetValueType() == gcv.Float {
		result, err = calculateFloat(firstValue.Float64(), secondValue.Float64(), s)
	} else {
		result, err = calculateInt(firstValue.Int(), secondValue.Int(), s)
	}
	return
}

func calculateComplex(firstValue, secondValue complex128, s string) (result gcv.Value, err error) {
	switch s {
	case add:
		result = gcv.MakeValue(firstValue + secondValue)
	case times1, times2, times3:
		result = gcv.MakeValue(firstValue * secondValue)
	case div:
		result = gcv.MakeValue(firstValue / secondValue)
	case sub:
		result = gcv.MakeValue(firstValue - secondValue)
	case pow:
		result = gcv.MakeValue(cmplx.Pow(firstValue, secondValue))
	default:
		err = errors.New("IllegalArgumentException")
	}
	return
}

func calculateFloat(firstValue, secondValue float64, s string) (result gcv.Value, err error) {
	switch s {
	case add:
		result = gcv.MakeValue(firstValue + secondValue)
	case times1, times2, times3:
		result = gcv.MakeValue(firstValue * secondValue)
	case div:
		result = gcv.MakeValue(firstValue / secondValue)
	case sub:
		result = gcv.MakeValue(firstValue - secondValue)
	case pow:
		result = gcv.MakeValue(math.Pow(firstValue, secondValue))
	case mod:
		result = gcv.MakeValue(math.Mod(firstValue, secondValue))
	default:
		err = errors.New("IllegalArgumentException")
	}
	return
}

func calculateInt(firstValue, secondValue int, s string) (result gcv.Value, err error) {
	switch s {
	case add:
		result = gcv.MakeValue(firstValue + secondValue)
	case times1, times2, times3:
		result = gcv.MakeValue(firstValue * secondValue)
	case div:
		result = gcv.MakeValue(firstValue / secondValue)
	case sub:
		result = gcv.MakeValue(firstValue - secondValue)
	case pow:
		result = gcv.MakeValue(int(math.Pow(float64(firstValue), float64(secondValue))))
	case mod:
		result = gcv.MakeValue(int(math.Mod(float64(firstValue), float64(secondValue))))
	default:
		err = errors.New("IllegalArgumentException")
	}
	return
}
