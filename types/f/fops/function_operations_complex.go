package fops

import "math/cmplx"

// VariableComplex return a complex variable
func VariableComplex(index int) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return x[index]
	}
	return
}

// NegativeVariableComplex returns a negative complex variable
func NegativeVariableComplex(index int) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return -x[index]
	}
	return
}

// ConstantComplex returns a complex constand
func ConstantComplex(constant complex128) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return constant
	}
	return
}

// ParensComplex returns the Identity of f(x...)
func ParensComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return (f(x...))
	}
	return
}

// ArgumentComplex returns the argument of a complex function
func ArgumentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return complex(cmplx.Phase(f(x...)), 0)
	}
	return
}

// CotComplex returns the cot of a complex function
func CotComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cot(f(x...))
	}
	return
}

// ConjugateComplex returns the conjugate of a complex function
func ConjugateComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Conj(f(x...))
	}
	return
}

// AddComplex returns the addition of two complex functions
func AddComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) + g(x...)
	}
	return
}

// SubtractComplex returns the straction of two complex functions
func SubtractComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) - g(x...)
	}
	return
}

// DivideComplex returns the division of two complex functions
func DivideComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) / g(x...)
	}
	return
}

// MultipleComplex returns the multiplication of two complex functions
func MultipleComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) * g(x...)
	}
	return
}

// SquareRootComplex returns the square root of a complex function
func SquareRootComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sqrt(f(x...))
	}
	return
}

// AbsoluteValueComplex returns the absolute value of a complex function
func AbsoluteValueComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return complex(cmplx.Abs(f(x...)), 0)
	}
	return
}

// SineComplex returns the sine of a complex function
func SineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sin(f(x...))
	}
	return
}

// CosineComplex returns the cosine of a complex function
func CosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cos(f(x...))
	}
	return
}

// TangentComplex returns the tangent of a complex function
func TangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Tan(f(x...))
	}
	return
}

// ArcsineComplex returns the arcsine of a complex function
func ArcsineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Asin(f(x...))
	}
	return
}

// ArccosineComplex returns the arccosine of a complex function
func ArccosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Acos(f(x...))
	}
	return
}

// ArctangentComplex returns the arctangent of a complex function
func ArctangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Atan(f(x...))
	}
	return
}

// HyperbolicSineComplex returns the hyperbolicSine of a complex function
func HyperbolicSineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sinh(f(x...))
	}
	return
}

// HyperbolicCosineComplex returns the hyperbolicCosine of a complex function
func HyperbolicCosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cosh(f(x...))
	}
	return
}

// HyperbolicTangentComplex returns the hyperbolicTangent of a complex function
func HyperbolicTangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Tanh(f(x...))
	}
	return
}

// InverseHyperbolicSineComplex return the inverseHyperbolicSine of a complex function
func InverseHyperbolicSineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Asinh(f(x...))
	}
	return
}

// InverseHyperbolicCosineComplex returns the inverseHyperbolicCosine of a complex function
func InverseHyperbolicCosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Acosh(f(x...))
	}
	return
}

// InverseHyperbolicTangentComplex returns the inverseHyperbolicTangent of a complex function
func InverseHyperbolicTangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Atanh(f(x...))
	}
	return
}

// LogBaseGxComplex returns the log of two complex functions
// g(x) will typically be a constant function
func LogBaseGxComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Log(f(x...)) / cmplx.Log(g(x...))
	}
	return
}

// PowerComplex return the power of a complex function raised to another complex function
func PowerComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Pow(f(x...), g(x...))
	}
	return
}
