package functions

import "math"

func variable(index int) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return x[index]
	}
	return
}

func negativeVariable(index int) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return -x[index]
	}
	return
}

func constant(constant float64) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return constant
	}
	return
}

func parens(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return (f(x...))
	}
	return
}

func add(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) + g(x...)
	}
	return
}

func subtract(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) - g(x...)
	}
	return
}

func divide(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) / g(x...)
	}
	return
}

func multiple(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) * g(x...)
	}
	return
}

func squareRoot(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sqrt(f(x...))
	}
	return
}

func absoluteValue(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Abs(f(x...))
	}
	return
}

func sine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sin(f(x...))
	}
	return
}

func cosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Cos(f(x...))
	}
	return
}

func tangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Tan(f(x...))
	}
	return
}

func arcsine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Asin(f(x...))
	}
	return
}

func arccosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Acos(f(x...))
	}
	return
}

func arctangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atan(f(x...))
	}
	return
}

func arctangent2(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atan2(f(x...), g(x...))
	}
	return
}

func hyperbolicSine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sinh(f(x...))
	}
	return
}

func hyperbolicCosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Cosh(f(x...))
	}
	return
}

func hyperbolicTangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Tanh(f(x...))
	}
	return
}

func inverseHyperbolicSine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Asinh(f(x...))
	}
	return
}

func inverseHyperbolicCosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Acosh(f(x...))
	}
	return
}

func inverseHyperbolicTangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atanh(f(x...))
	}
	return
}

// g(x) will typically be a constant function
func logBaseGx(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Log(f(x...)) / math.Log(g(x...))
	}
	return
}

func power(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Pow(f(x...), g(x...))
	}
	return
}

func modulo(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Mod(f(x...), g(x...))
	}
	return
}

func floor(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Floor(f(x...))
	}
	return
}

func ceil(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Ceil(f(x...))
	}
	return
}

func max(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Max(f(x...), g(x...))
	}
	return
}

func min(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Min(f(x...), g(x...))
	}
	return
}

func errorFunction(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Erf(f(x...))
	}
	return
}
