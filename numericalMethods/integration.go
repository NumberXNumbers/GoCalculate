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
