package functions

import "math/cmplx"

func variableComplex(index int) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return x[index]
	}
	return
}

func negativeVariableComplex(index int) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return -x[index]
	}
	return
}

func constantComplex(constant complex128) (f func(x ...complex128) complex128) {
	f = func(x ...complex128) complex128 {
		return constant
	}
	return
}

func parensComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return (f(x...))
	}
	return
}

func argumentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return complex(cmplx.Phase(f(x...)), 0)
	}
	return
}

func cotComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cot(f(x...))
	}
	return
}

func conjugateComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Conj(f(x...))
	}
	return
}

func addComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) + g(x...)
	}
	return
}

func subtractComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) - g(x...)
	}
	return
}

func divideComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) / g(x...)
	}
	return
}

func multipleComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return f(x...) * g(x...)
	}
	return
}

func squareRootComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sqrt(f(x...))
	}
	return
}

func absoluteValueComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return complex(cmplx.Abs(f(x...)), 0)
	}
	return
}

func sineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sin(f(x...))
	}
	return
}

func cosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cos(f(x...))
	}
	return
}

func tangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Tan(f(x...))
	}
	return
}

func arcsineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Asin(f(x...))
	}
	return
}

func arccosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Acos(f(x...))
	}
	return
}

func arctangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Atan(f(x...))
	}
	return
}

func hyperbolicSineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Sinh(f(x...))
	}
	return
}

func hyperbolicCosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Cosh(f(x...))
	}
	return
}

func hyperbolicTangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Tanh(f(x...))
	}
	return
}

func inverseHyperbolicSineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Asinh(f(x...))
	}
	return
}

func inverseHyperbolicCosineComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Acosh(f(x...))
	}
	return
}

func inverseHyperbolicTangentComplex(f func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Atanh(f(x...))
	}
	return
}

// g(x) will typically be a constant function
func logBaseGxComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Log(f(x...)) / cmplx.Log(g(x...))
	}
	return
}

func powerComplex(f func(x ...complex128) complex128, g func(x ...complex128) complex128) (h func(x ...complex128) complex128) {
	h = func(x ...complex128) complex128 {
		return cmplx.Pow(f(x...), g(x...))
	}
	return
}
