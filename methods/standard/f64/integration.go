package methods

import (
	"errors"
	"math"
)

// Euler1D is for solving the numerical integration 1D euler method
func Euler1D(a float64, b float64, N int, initValue float64, f func(float64, float64) float64) float64 {
	h := (b - a) / float64(N)
	x := a
	omega := initValue

	for i := 0; i < N; i++ {
		omega += h * f(x, omega)
		x += h
	}

	return omega
}

// TrapezoidRule is for solving the numerical integration using the trapezoid rule
func TrapezoidRule(a float64, b float64, f func(float64) float64) float64 {
	var omega float64
	h := (b - a)
	x := a

	omega = f(x+h) + f(x)

	return h / 2 * omega
}

// SimpsonRule for solving numerical integration
func SimpsonRule(a float64, b float64, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / 2
	x := a

	omega = f(x) + 4*f(x+h) + f(x+2*h)

	return h / 3 * omega
}

// Simpson38Rule is Simpson's 3/8ths rule for solving numerical integration
func Simpson38Rule(a float64, b float64, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / 3
	x := a

	omega += f(x) + 3*f(x+h) + 3*f(x+2*h) + f(x+3*h)

	return 3 * h / 8 * omega
}

// BooleRule is Boole's rule for solving numerical integration
func BooleRule(a float64, b float64, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / 4
	x := a

	omega = 7*f(x) + 32*f(x+h) + 12*f(x+2*h) + 32*f(x+3*h) + 7*f(x+4*h)

	return 2 * h / 45 * omega
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

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2.0, omega+kappa/2.0)

		omega += kappa2
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// ModifiedEuler returns a [][]float64
func ModifiedEuler(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
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

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		theta += stepSize
		kappa2 = stepSize * f(theta, omega+kappa)

		omega += (kappa + kappa2) / 2.0

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// Heun returns a solution to the 3rd order runge-kutta method
func Heun(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
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

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/3.0, omega+kappa/3.0)
		kappa3 = stepSize * f(theta+2.0*stepSize/3.0, omega+2.0*kappa2/3.0)

		omega += (kappa + 3.0*kappa3) / 4.0
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
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

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2.0, omega+kappa/2.0)
		kappa3 = stepSize * f(theta+stepSize/2.0, omega+kappa2/2.0)
		kappa4 = stepSize * f(theta+stepSize, omega+kappa3)

		omega += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
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

	solutionSet = append(solutionSet, []float64{theta, omega})

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
		kappa5 = stepSize * f(theta+stepSize, omega+439.0*kappa/216.0-8.0*kappa2+
			3680.0*kappa3/513.0-845.0*kappa4/4104.0)
		kappa6 = stepSize * f(theta+stepSize/2.0, omega-8.0*kappa/27.0+2.0*kappa2-
			3544.0*kappa3/2565.0+1859.0*kappa4/4104.0-11.0*kappa5/40.0)

		remainder = math.Abs(kappa/360.0-128.0*kappa3/4275.0-2197.0*kappa4/75240.0+kappa5/50.0+2.0*kappa6/55.0) / stepSize

		if remainder <= TOL {
			theta += stepSize
			omega += 25.0*kappa/216.0 + 1408.0*kappa3/2565.0 + 2197.0*kappa4/4104.0 - kappa5/5.0

			solutionSet = append(solutionSet, []float64{theta, omega})
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

		if theta >= b {
			done = true
		} else if theta+stepSize > b {
			stepSize = b - theta
		} else if stepSize < minStep {
			done = true
		}
	}

	return solutionSet
}

// AdamsBashforth2 returns a solution found using the 2nd order Adams-Bashforth method
// Algorithm from Numerical Analysis - By Burden and Faires
func AdamsBashforth2(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	for i := 0; i < 2; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2

	omega := omega2

	var kappa float64
	var kappa2 float64

	for i := 1; i < N; i++ {
		kappa = 3.0 * f(solutionSet[i][0], solutionSet[i][1])
		kappa2 = f(solutionSet[i-1][0], solutionSet[i-1][1])

		omega += stepSize * (kappa - kappa2) / 2
		theta = stepSize + solutionSet[i][0]

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforth3 returns a solution found using the 3rd order Adams-Bashforth method
func AdamsBashforth3(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	for i := 0; i < 3; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2
	solutionSet[2][1] = omega3

	omega := omega3

	var kappa float64
	var kappa2 float64
	var kappa3 float64

	for i := 2; i < N; i++ {
		kappa = 23 * f(solutionSet[i][0], solutionSet[i][1])
		kappa2 = 16 * f(solutionSet[i-1][0], solutionSet[i-1][1])
		kappa3 = 5 * f(solutionSet[i-2][0], solutionSet[i-2][1])

		omega += stepSize * (kappa - kappa2 + kappa3) / 12
		theta = stepSize + solutionSet[i][0]

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforth4 returns a solution found using the 4th order Adams-Bashforth method
func AdamsBashforth4(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, initialCondition4 float64,
	f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	for i := 0; i < 4; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2
	solutionSet[2][1] = omega3
	solutionSet[3][1] = omega4

	omega := omega4

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64

	for i := 3; i < N; i++ {
		kappa = 55 * f(solutionSet[i][0], solutionSet[i][1])
		kappa2 = 59 * f(solutionSet[i-1][0], solutionSet[i-1][1])
		kappa3 = 37 * f(solutionSet[i-2][0], solutionSet[i-2][1])
		kappa4 = 9 * f(solutionSet[i-3][0], solutionSet[i-3][1])

		omega += stepSize * (kappa - kappa2 + kappa3 - kappa4) / 24
		theta = stepSize + solutionSet[i][0]

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforth5 returns a solution found using the 5th order Adams-Bashforth method
func AdamsBashforth5(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, initialCondition4 float64,
	initialCondition5 float64, f func(x, y float64) float64) [][]float64 {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4
	omega5 := initialCondition5

	solutionSet := make([][]float64, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float64, 2)
	}

	for i := 0; i < 5; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2
	solutionSet[2][1] = omega3
	solutionSet[3][1] = omega4
	solutionSet[4][1] = omega5

	omega := omega5

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64
	var kappa5 float64

	for i := 4; i < N; i++ {
		kappa = 1901 * f(solutionSet[i][0], solutionSet[i][1])
		kappa2 = 2774 * f(solutionSet[i-1][0], solutionSet[i-1][1])
		kappa3 = 2616 * f(solutionSet[i-2][0], solutionSet[i-2][1])
		kappa4 = 1274 * f(solutionSet[i-3][0], solutionSet[i-3][1])
		kappa5 = 251 * f(solutionSet[i-4][0], solutionSet[i-4][1])

		omega += stepSize * (kappa - kappa2 + kappa3 - kappa4 + kappa5) / 720
		theta = stepSize + solutionSet[i][0]

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforthMoulton3 returns solutions for the third order Adams-Bashforth-Moulton predictor-corrector method
func AdamsBashforthMoulton3(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
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

	for i := 0; i < 2; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/3.0, omega+kappa/3.0)
		kappa3 = stepSize * f(theta+2.0*stepSize/3.0, omega+2.0*kappa2/3.0)

		omega += (kappa + 3.0*kappa3) / 4.0
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	for i := 2; i < N; i++ {
		theta += stepSize
		omega = solutionSet[i][1] + stepSize*(23.0*f(solutionSet[i][0], solutionSet[i][1])-
			16.0*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			5.0*f(solutionSet[i-2][0], solutionSet[i-2][1]))/12.0
		omega = solutionSet[i][1] + stepSize*(5.0*f(theta, omega)+
			8.0*f(solutionSet[i][0], solutionSet[i][1])-
			f(solutionSet[i-1][0], solutionSet[i-1][1]))/12.0

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforthMoulton4 returns solutions for the fourth order Adams-Bashforth-Moulton predictor-corrector method
func AdamsBashforthMoulton4(a float64, b float64, N int, initialCondition float64, f func(x, y float64) float64) [][]float64 {
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

	for i := 0; i < 3; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2.0, omega+kappa/2.0)
		kappa3 = stepSize * f(theta+stepSize/2.0, omega+kappa2/2.0)
		kappa4 = stepSize * f(theta+stepSize, omega+kappa3)

		omega += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	for i := 3; i < N; i++ {
		theta += stepSize
		omega = solutionSet[i][1] + stepSize*(55.0*f(solutionSet[i][0], solutionSet[i][1])-
			59.0*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			37.0*f(solutionSet[i-2][0], solutionSet[i-2][1])-
			9.0*f(solutionSet[i-3][0], solutionSet[i-3][1]))/24.0
		omega = solutionSet[i][1] + stepSize*(9.0*f(theta, omega)+
			19.0*f(solutionSet[i][0], solutionSet[i][1])-
			5.0*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			f(solutionSet[i-2][0], solutionSet[i-2][1]))/24.0

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforthMoulton returns a solution from the variable step Adams-Bashforth-Moulton method
func AdamsBashforthMoulton(a float64, b float64, initialCondition float64,
	TOL float64, maxStep float64, minStep float64, f func(x, y float64) float64) ([][]float64, error) {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false
	rk4Done := false
	lastValueCalc := false

	var thetas []float64
	var omegas []float64

	thetas = append(thetas, theta)
	omegas = append(omegas, omega)

	RK4 := func(h float64, tSet, oSet []float64, f func(x, y float64) float64) ([]float64, []float64) {
		var kappa float64
		var kappa2 float64
		var kappa3 float64
		var kappa4 float64
		var t float64
		var o float64

		for i := 0; i < 3; i++ {
			t = tSet[len(tSet)-1]
			o = oSet[len(oSet)-1]

			kappa = h * f(t, o)
			kappa2 = h * f(t+h/2.0, o+kappa/2.0)
			kappa3 = h * f(t+h/2.0, o+kappa2/2.0)
			kappa4 = h * f(t+h, o+kappa3)

			o += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
			t = tSet[len(tSet)-1] + h

			tSet, oSet = append(tSet, t), append(oSet, o)
		}

		return tSet, oSet
	}

	thetas, omegas = RK4(stepSize, thetas, omegas, f)
	rk4Done = true

	theta = thetas[len(thetas)-1] + stepSize
	var predictor float64
	var corrector float64
	var sigma float64
	var zeta float64

	for !done {
		predictor = omegas[len(thetas)-1] + stepSize*(55.0*f(thetas[len(thetas)-1], omegas[len(omegas)-1])-
			59.0*f(thetas[len(thetas)-2], omegas[len(omegas)-2])+
			37.0*f(thetas[len(thetas)-3], omegas[len(omegas)-3])-
			9.0*f(thetas[len(thetas)-4], omegas[len(omegas)-4]))/24.0
		corrector = omegas[len(thetas)-1] + stepSize*(9.0*f(theta, predictor)+
			19.0*f(thetas[len(thetas)-1], omegas[len(omegas)-1])-
			5.0*f(thetas[len(thetas)-2], omegas[len(omegas)-2])+
			f(thetas[len(thetas)-3], omegas[len(omegas)-3]))/24.0

		sigma = 19.0 * math.Abs(corrector-predictor) / (270.0 * stepSize)
		if sigma <= TOL {
			omega = corrector

			thetas, omegas = append(thetas, theta), append(omegas, omega)

			if lastValueCalc {
				done = true
			} else {
				if sigma <= 0.1*TOL || thetas[len(thetas)-1]+stepSize > b {
					zeta = math.Pow(TOL/(2.0*sigma), 1.0/4.0)
					if zeta > 4 {
						stepSize = 4.0 * stepSize
					} else {
						stepSize = zeta * stepSize
					}

					if stepSize > maxStep {
						stepSize = maxStep
					}

					if thetas[len(thetas)-1]+4.0*stepSize > b {
						stepSize = (b - thetas[len(thetas)-1]) / 4.0
						lastValueCalc = true
					}

					thetas, omegas = RK4(stepSize, thetas, omegas, f)
					rk4Done = true
				}
			}
		} else {
			zeta = math.Pow(TOL/(2.0*sigma), 1.0/4.0)

			if zeta < 0.1 {
				stepSize = 0.1 * stepSize
			} else {
				stepSize = zeta * stepSize
			}

			if stepSize < minStep {
				done = true
			} else {
				if rk4Done {
					thetas = thetas[:len(thetas)-3]
					omegas = omegas[:len(omegas)-3]
				}

				thetas, omegas = RK4(stepSize, thetas, omegas, f)
				rk4Done = true
			}
		}

		theta = thetas[len(thetas)-1] + stepSize
	}

	var solutionSet [][]float64

	if !lastValueCalc {
		return solutionSet, errors.New("Minimum step size exceeded")
	}

	for i := 0; i < len(thetas); i++ {
		solutionSet = append(solutionSet, []float64{thetas[i], omegas[i]})
	}

	return solutionSet, nil
}
