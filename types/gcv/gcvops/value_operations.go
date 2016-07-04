package gcvops

import (
	"errors"
	"math"
	"math/cmplx"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

// Add will add two gcv Values together
func Add(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(valueA.Complex() + valueB.Complex())
	}
	return gcv.MakeValue(valueA.Real() + valueB.Real())
}

// Sub will subtract two gcv Values together
func Sub(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(valueA.Complex() - valueB.Complex())
	}
	return gcv.MakeValue(valueA.Real() - valueB.Real())
}

// Mult will multiply two gcv Values together
func Mult(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(valueA.Complex() * valueB.Complex())
	}
	return gcv.MakeValue(valueA.Real() * valueB.Real())
}

// Div will divide two gcv Values together
func Div(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(valueA.Complex() / valueB.Complex())
	}
	return gcv.MakeValue(valueA.Real() / valueB.Real())
}

// Sqrt returns the square root of a gcv Value
func Sqrt(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Sqrt(value.Complex()))
	}
	return gcv.MakeValue(math.Sqrt(value.Real()))
}

// Abs returns the absolute value of a gcv Value
func Abs(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Abs(value.Complex()))
	}
	return gcv.MakeValue(math.Abs(value.Real()))
}

// Conj returns the conjugate of a gcv Value
func Conj(value gcv.Value) gcv.Value {
	return gcv.MakeValue(cmplx.Conj(value.Complex()))
}

// Cot returns the cot of a gcv Value, meant for Value of type Complex
func Cot(value gcv.Value) gcv.Value {
	return gcv.MakeValue(cmplx.Cot(value.Complex()))
}

// Sin returns the sine of a function
func Sin(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Sin(value.Complex()))
	}
	return gcv.MakeValue(math.Sin(value.Real()))
}

// Cos returns the cosine of a function
func Cos(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Cos(value.Complex()))
	}
	return gcv.MakeValue(math.Cos(value.Real()))
}

// Tan returns the tangent of a function
func Tan(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Tan(value.Complex()))
	}
	return gcv.MakeValue(math.Tan(value.Real()))
}

// Asin returns the arcsine of a function
func Asin(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Asin(value.Complex()))
	}
	return gcv.MakeValue(math.Asin(value.Real()))
}

// Acos returns the arccosine of a gcv Value
func Acos(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Acos(value.Complex()))
	}
	return gcv.MakeValue(math.Acos(value.Real()))
}

// Atan returns the arctangent of a gcv Value
func Atan(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Atan(value.Complex()))
	}
	return gcv.MakeValue(math.Atan(value.Real()))
}

// Sinh returns the hyperbolicSine of a gcv Value
func Sinh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Sinh(value.Complex()))
	}
	return gcv.MakeValue(math.Sinh(value.Real()))
}

// Cosh returns the hyperbolicCosine of a gcv Value
func Cosh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Cosh(value.Complex()))
	}
	return gcv.MakeValue(math.Cosh(value.Real()))
}

// Tanh returns the hyperbolicTangent of a gcv Value
func Tanh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Tanh(value.Complex()))
	}
	return gcv.MakeValue(math.Tanh(value.Real()))
}

// Asinh returns the inverseHyperbolicSine of a gcv Value
func Asinh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Asinh(value.Complex()))
	}
	return gcv.MakeValue(math.Asinh(value.Real()))
}

// Acosh returns the inverseHyperbolicCosine of a gcv Value
func Acosh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Acosh(value.Complex()))
	}
	return gcv.MakeValue(math.Acosh(value.Real()))
}

// Atanh returns the inverseHyperbolicTangent of a gcv Value
func Atanh(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Atanh(value.Complex()))
	}
	return gcv.MakeValue(math.Atanh(value.Real()))
}

// Log returns the natural log of gcv Value
func Log(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Log(value.Complex()))
	}
	return gcv.MakeValue(math.Log(value.Real()))
}

// Log10 returns the log base 10 of gcv Value
func Log10(value gcv.Value) gcv.Value {
	if value.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Log10(value.Complex()))
	}
	return gcv.MakeValue(math.Log10(value.Real()))
}

// LogBase returns the log of gcv Value valueA in base of gcv Value valueB
func LogBase(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Log(valueA.Complex()) / cmplx.Log(valueB.Complex()))
	}
	return gcv.MakeValue(math.Log(valueA.Real()) / math.Log(valueB.Real()))
}

// Pow returns the power of gcv Value valueA raised to the power of gcv Value valueB
func Pow(valueA gcv.Value, valueB gcv.Value) gcv.Value {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return gcv.MakeValue(cmplx.Pow(valueA.Complex(), valueB.Complex()))
	}
	return gcv.MakeValue(math.Pow(valueA.Real(), valueB.Real()))
}

// Mod returns the modulo of a real Value valueA by a real Value valueB.
// if either Value is of type Complex an error is returned
func Mod(valueA gcv.Value, valueB gcv.Value) (gcv.Value, error) {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return nil, errors.New("Modulo is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Mod(valueA.Real(), valueB.Real())), nil
}

// Floor returns the floor (rounded down) of a gcv Value.
// if either Value is of type Complex an error is returned
func Floor(value gcv.Value) (gcv.Value, error) {
	if value.Type() == gcv.Complex {
		return nil, errors.New("Floor is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Floor(value.Real())), nil
}

// Ceil returns the ceil (rounded up) of a gcv Value
// if either Value is of type Complex an error is returned
func Ceil(value gcv.Value) (gcv.Value, error) {
	if value.Type() == gcv.Complex {
		return nil, errors.New("Ceil is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Ceil(value.Real())), nil
}

// Max returns the max of two gcv Value
// if either Value is of type Complex an error is returned
func Max(valueA gcv.Value, valueB gcv.Value) (gcv.Value, error) {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return nil, errors.New("Max is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Max(valueA.Real(), valueB.Real())), nil
}

// Min returns the minimum of two gcv Value.
// if either Value is of type Complex an error is returned
func Min(valueA gcv.Value, valueB gcv.Value) (gcv.Value, error) {
	if valueA.Type() == gcv.Complex || valueB.Type() == gcv.Complex {
		return nil, errors.New("Min is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Min(valueA.Real(), valueB.Real())), nil
}

// Erf returns the error function of a gcv Value.
// if either Value is of type Complex an error is returned
func Erf(value gcv.Value) (gcv.Value, error) {
	if value.Type() == gcv.Complex {
		return nil, errors.New("Erf is not supported for Complex numbers")
	}
	return gcv.MakeValue(math.Erf(value.Real())), nil
}
