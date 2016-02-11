package numericalMethods

import "math"

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

// Simpson38Rule is Simpson's 3/8ths rule for solving numerical integration
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

// BooleRule is Boole's rule for solving numerical integration
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

// RungeKutta2 or midpoint method returns a solution found using the 2nd order runge-kutta
func RungeKutta2(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float64
	var kappa2 float64

	for i := 1; i < N+1; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2.0, omega+kappa/2.0)

		omega += kappa2
		theta += stepSize

		solutionSet[i][0] = theta
		solutionSet[i][1] = omega
	}

	return solutionSet
}

// ModifiedEulerMethod returns a [][]float64
func ModifiedEulerMethod(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float64
	var kappa2 float64

	for i := 1; i < N+1; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta, omega+kappa)

		omega += (kappa + kappa2) / 2.0
		theta += stepSize

		solutionSet[i][0] = theta
		solutionSet[i][1] = omega
	}

	return solutionSet
}

// HeunMethod returns a solution to the 3rd order runge-kutta method
func HeunMethod(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float64
	var kappa2 float64
	var kappa3 float64

	for i := 1; i < N+1; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/3.0, omega+kappa/3.0)
		kappa3 = stepSize * f(theta+2.0*stepSize/3.0, omega+2.0*kappa2/3.0)

		omega += (kappa + 3.0*kappa3) / 4.0
		theta += stepSize

		solutionSet[i][0] = theta
		solutionSet[i][1] = omega
	}

	return solutionSet
}

// RungeKutta4 returns a solution found using the 4th order runge-kutta method
func RungeKutta4(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64

	for i := 1; i < N+1; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2.0, omega+kappa/2.0)
		kappa3 = stepSize * f(theta+stepSize/2.0, omega+kappa2/2.0)
		kappa4 = stepSize * f(theta+stepSize, omega+kappa3)

		omega += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
		theta += stepSize

		solutionSet[i][0] = theta
		solutionSet[i][1] = omega
	}

	return solutionSet
}

// RungeKuttaFehlbery returns a solution to the runge-kutta-fehlbery method
func RungeKuttaFehlbery(a float64, b float64, initialCondition float64,
	TOL float64, maxStep float64, minStep float64,
	f func(x, y float64) float64) [][]float64 {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false

	var solutionSet [][]float64

	solutionSet = append(solutionSet, []float64{theta, omega, stepSize})

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64
	var kappa5 float64
	var kappa6 float64

	var remainder float64
	var delta float64

	for !done {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/4.0, omega+kappa/4.0)
		kappa3 = stepSize * f(theta+3.0*stepSize/8.0, omega+3.0*kappa/32.0+9.0*kappa2/32.0)
		kappa4 = stepSize * f(theta+12.0*stepSize/13.0, omega+1932.0*kappa/2197.0-
			7200.0*kappa2/2197.0+7296.0*kappa3/2197.0)
		kappa5 = stepSize * f(theta+stepSize, omega+439.0*kappa-8.0*kappa2+
			3680.0*kappa3/513.0-845.0*kappa4/4104.0)
		kappa6 = stepSize * f(theta+stepSize/2.0, omega-8.0*kappa/27.0+2.0*kappa2-
			3544.0*kappa3/2565.0+1859.0*kappa4/4104.0-11.0*kappa5/40.0)

		remainder = math.Abs(kappa/360.0-128.0*kappa3/4275.0-2197.0*kappa4/75240.0+kappa5/50.0+2.0*kappa6/55.0) / stepSize

		if remainder < TOL {
			theta += stepSize
			omega += 25.0*kappa/216.0 + 1408.0*kappa3/2565.0 + 2197.0*kappa4/4104.0 - kappa5/5.0

			solutionSet = append(solutionSet, []float64{theta, omega, stepSize})
		}

		delta = 0.84 * math.Pow(TOL/remainder, 1.0/4.0)

		if delta <= 0.1 {
			stepSize = 0.1 * stepSize
		} else if delta >= 4 {
			stepSize = 4.0 * stepSize
		} else {
			stepSize = delta * stepSize
		}

		if stepSize > maxStep {
			stepSize = maxStep
		}

		if theta > b {
			done = true
		} else if theta+stepSize > b {
			stepSize = b - theta
		} else if stepSize < minStep {
			done = true
		}
	}

	return solutionSet
}
