package fops

import "math"

// Variable returns a real variable
func Variable(index int) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return x[index]
	}
	return
}

// NegativeVariable returns a negative real variable
func NegativeVariable(index int) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return -x[index]
	}
	return
}

// Constant returns a constant
func Constant(constant float64) (f func(x ...float64) float64) {
	f = func(x ...float64) float64 {
		return constant
	}
	return
}

// Parens returns the identity of f(x...)
func Parens(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return (f(x...))
	}
	return
}

// Add returns the addition of two functions
func Add(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) + g(x...)
	}
	return
}

// Subtract returns the subtraction of two functions
func Subtract(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) - g(x...)
	}
	return
}

// Divide returns the division of two functions
func Divide(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) / g(x...)
	}
	return
}

// Multiple returns the multiplication of two functions
func Multiple(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return f(x...) * g(x...)
	}
	return
}

// SquareRoot returns the square root of a function
func SquareRoot(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sqrt(f(x...))
	}
	return
}

// AbsoluteValue returns the absolute value of a funciton
func AbsoluteValue(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Abs(f(x...))
	}
	return
}

// Sine returns the sine of a function
func Sine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sin(f(x...))
	}
	return
}

// Cosine returns the cosine of a function
func Cosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Cos(f(x...))
	}
	return
}

// Tangent returns the tangent of a function
func Tangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Tan(f(x...))
	}
	return
}

// Arcsine returns the arcsine of a function
func Arcsine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Asin(f(x...))
	}
	return
}

// Arccosine returns the arccosine of a function
func Arccosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Acos(f(x...))
	}
	return
}

// Arctangent returns the arctangent of a function
func Arctangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atan(f(x...))
	}
	return
}

// Arctangent2 returns the arctangent using the signs of the two functions to tell which
// quadrant the resulting function is in. h(x...) = Atan(f(xx...)/g(x...))
func Arctangent2(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atan2(f(x...), g(x...))
	}
	return
}

// HyperbolicSine returns the hyperbolicSine of a function
func HyperbolicSine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Sinh(f(x...))
	}
	return
}

// HyperbolicCosine returns the hyperbolicCosine of a function
func HyperbolicCosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Cosh(f(x...))
	}
	return
}

// HyperbolicTangent returns the hyperbolicTangent of a function
func HyperbolicTangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Tanh(f(x...))
	}
	return
}

// InverseHyperbolicSine returns the inverseHyperbolicSine of a function
func InverseHyperbolicSine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Asinh(f(x...))
	}
	return
}

// InverseHyperbolicCosine returns the inverseHyperbolicCosine of a function
func InverseHyperbolicCosine(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Acosh(f(x...))
	}
	return
}

// InverseHyperbolicTangent returns the inverseHyperbolicTangent of a function
func InverseHyperbolicTangent(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Atanh(f(x...))
	}
	return
}

// LogBaseGx returns the log of two functions
// g(x) will typically be a constant function
func LogBaseGx(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Log(f(x...)) / math.Log(g(x...))
	}
	return
}

// Power returns the power of a function raised two another function
func Power(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Pow(f(x...), g(x...))
	}
	return
}

// Modulo returns the modulo of a function
func Modulo(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Mod(f(x...), g(x...))
	}
	return
}

// Floor returns the floor (rounded down) of a function
func Floor(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Floor(f(x...))
	}
	return
}

// Ceil returns the ceil (rounded up) of a function
func Ceil(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Ceil(f(x...))
	}
	return
}

// Max returns the max of two functions
func Max(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Max(f(x...), g(x...))
	}
	return
}

// Min returns the minimum
func Min(f func(x ...float64) float64, g func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Min(f(x...), g(x...))
	}
	return
}

// ErrorFunction returns the error function of a function.
func ErrorFunction(f func(x ...float64) float64) (h func(x ...float64) float64) {
	h = func(x ...float64) float64 {
		return math.Erf(f(x...))
	}
	return
}
