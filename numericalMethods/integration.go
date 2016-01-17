package numericalMethods

// Euler1D is for solving the numerical integration 1D euler method
func Euler1D(a float64, b float64, N int, initValue float64, f func(float64, float64) float64) float64 {
	h := (b - a) / float64(N)
	t := a
	omega := initValue

	for i := 0; i < N; i++ {
		omega += h * f(t, omega)
		t += h
	}

	return omega
}

// TrapezoidRule is for solving the numerical integration using the trapezoid rule
func TrapezoidRule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += f(x+h) + f(x)
		x += h
	}

	return h / 2 * omega
}

// SimpsonRule for solving numerical integration
func SimpsonRule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += f(x) + 4*f(x+h) + f(x+2*h)
		x += h
	}
	return h / 6 * omega
}

// Simpson's 3/8ths rule for solving numerical integration
func Simpson38Rule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += f(x) + 3*f(x+h) + 3*f(x+2*h) + f(x+3*h)
		x += h
	}
	return h / 8 * omega
}

// Boole's rule for solving numerical integration
func BooleRule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += 7*f(x) + 32*f(x+h) + 12*f(x+2*h) + 32*f(x+3*h) + 7*f(x+4*h)
		x += h
	}
	return h / 90 * omega
}
