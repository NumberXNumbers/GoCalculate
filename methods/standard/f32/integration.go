package methods

import (
	"errors"
	"math"
)

// Euler1D is for solving the numerical integration 1D euler method
func Euler1D(a float32, b float32, N int, initValue float32, f func(float32, float32) float32) float32 {
	h := (b - a) / float32(N)
	t := a
	omega := initValue

	for i := 0; i < N; i++ {
		omega += h * f(t, omega)
		t += h
	}

	return omega
}

// TrapezoidRule is for solving the numerical integration using the trapezoid rule
func TrapezoidRule(a float32, b float32, f func(float32) float32) float32 {
	var omega float32
	h := (b - a)
	x := a

	omega = f(x+h) + f(x)

	return h / 2 * omega
}

// SimpsonRule for solving numerical integration
func SimpsonRule(a float32, b float32, f func(float32) float32) float32 {
	var omega float32
	h := (b - a) / 2
	x := a

	omega = f(x) + 4*f(x+h) + f(x+2*h)

	return h / 3 * omega
}

// Simpson38Rule is Simpson's 3/8ths rule for solving numerical integration
func Simpson38Rule(a float32, b float32, f func(float32) float32) float32 {
	var omega float32
	h := (b - a) / 3
	x := a

	omega += f(x) + 3*f(x+h) + 3*f(x+2*h) + f(x+3*h)

	return 3 * h / 8 * omega
}

// BooleRule is Boole's rule for solving numerical integration
func BooleRule(a float32, b float32, f func(float32) float32) float32 {
	var omega float32
	h := (b - a) / 4
	x := a

	omega = 7*f(x) + 32*f(x+h) + 12*f(x+2*h) + 32*f(x+3*h) + 7*f(x+4*h)

	return 2 * h / 45 * omega
}

// RungeKutta2 or midpoint method returns a solution found using the 2nd order runge-kutta
func RungeKutta2(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2, omega+kappa/2)

		omega += kappa2
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// ModifiedEuler returns a [][]float32
func ModifiedEuler(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		theta += stepSize
		kappa2 = stepSize * f(theta, omega+kappa)

		omega += (kappa + kappa2) / 2

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// Heun returns a solution to the 3rd order runge-kutta method
func Heun(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32
	var kappa3 float32

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/3, omega+kappa/3)
		kappa3 = stepSize * f(theta+2*stepSize/3, omega+2*kappa2/3)

		omega += (kappa + 3*kappa3) / 4
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// RungeKutta4 returns a solution found using the 4th order runge-kutta method
func RungeKutta4(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32
	var kappa3 float32
	var kappa4 float32

	for i := 0; i < N; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2, omega+kappa/2)
		kappa3 = stepSize * f(theta+stepSize/2, omega+kappa2/2)
		kappa4 = stepSize * f(theta+stepSize, omega+kappa3)

		omega += (kappa + 2*kappa2 + 2*kappa3 + kappa4) / 6
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// RungeKuttaFehlbery returns a solution to the runge-kutta-fehlbery method
// Algorithm from Numerical Analysis - By Burden and Faires
func RungeKuttaFehlbery(a float32, b float32, initialCondition float32,
	TOL float32, maxStep float32, minStep float32,
	f func(x, y float32) float32) [][]float32 {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false

	var solutionSet [][]float32

	solutionSet = append(solutionSet, []float32{theta, omega})

	var kappa float32
	var kappa2 float32
	var kappa3 float32
	var kappa4 float32
	var kappa5 float32
	var kappa6 float32

	var remainder float32
	var delta float32

	for !done {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/4, omega+kappa/4)
		kappa3 = stepSize * f(theta+3*stepSize/8, omega+3*kappa/32+9*kappa2/32)
		kappa4 = stepSize * f(theta+12*stepSize/13, omega+1932*kappa/2197-
			7200*kappa2/2197+7296*kappa3/2197)
		kappa5 = stepSize * f(theta+stepSize, omega+439*kappa/216-8*kappa2+
			3680*kappa3/513-845*kappa4/4104)
		kappa6 = stepSize * f(theta+stepSize/2, omega-8*kappa/27+2*kappa2-
			3544*kappa3/2565+1859*kappa4/4104-11*kappa5/40)

		remainder = float32(math.Abs(float64(kappa/360-128*kappa3/4275-2197*kappa4/75240+kappa5/50+2*kappa6/55))) / stepSize

		if remainder <= TOL {
			theta += stepSize
			omega += 25*kappa/216 + 1408*kappa3/2565 + 2197*kappa4/4104 - kappa5/5

			solutionSet = append(solutionSet, []float32{theta, omega})
		}

		delta = 0.84 * float32(math.Pow(float64(TOL/remainder), 1.0/4.0))

		if delta <= 0.1 {
			stepSize = 0.1 * stepSize
		} else if delta >= 4 {
			stepSize = 4 * stepSize
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
func AdamsBashforth2(a float32, b float32, N int, initialCondition1 float32,
	initialCondition2 float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	for i := 0; i < 2; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2

	omega := omega2

	var kappa float32
	var kappa2 float32

	for i := 1; i < N; i++ {
		kappa = 3 * f(solutionSet[i][0], solutionSet[i][1])
		kappa2 = f(solutionSet[i-1][0], solutionSet[i-1][1])

		omega += stepSize * (kappa - kappa2) / 2
		theta = stepSize + solutionSet[i][0]

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforth3 returns a solution found using the 3rd order Adams-Bashforth method
func AdamsBashforth3(a float32, b float32, N int, initialCondition1 float32,
	initialCondition2 float32, initialCondition3 float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	for i := 0; i < 3; i++ {
		solutionSet[i][0] = theta
		theta += stepSize
	}

	solutionSet[0][1] = omega1
	solutionSet[1][1] = omega2
	solutionSet[2][1] = omega3

	omega := omega3

	var kappa float32
	var kappa2 float32
	var kappa3 float32

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
func AdamsBashforth4(a float32, b float32, N int, initialCondition1 float32,
	initialCondition2 float32, initialCondition3 float32, initialCondition4 float32,
	f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
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

	var kappa float32
	var kappa2 float32
	var kappa3 float32
	var kappa4 float32

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
func AdamsBashforth5(a float32, b float32, N int, initialCondition1 float32,
	initialCondition2 float32, initialCondition3 float32, initialCondition4 float32,
	initialCondition5 float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4
	omega5 := initialCondition5

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
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

	var kappa float32
	var kappa2 float32
	var kappa3 float32
	var kappa4 float32
	var kappa5 float32

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
func AdamsBashforthMoulton3(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32
	var kappa3 float32

	for i := 0; i < 2; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/3, omega+kappa/3)
		kappa3 = stepSize * f(theta+2*stepSize/3, omega+2*kappa2/3)

		omega += (kappa + 3*kappa3) / 4
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	for i := 2; i < N; i++ {
		theta += stepSize
		omega = solutionSet[i][1] + stepSize*(23*f(solutionSet[i][0], solutionSet[i][1])-
			16*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			5*f(solutionSet[i-2][0], solutionSet[i-2][1]))/12
		omega = solutionSet[i][1] + stepSize*(5*f(theta, omega)+
			8*f(solutionSet[i][0], solutionSet[i][1])-
			f(solutionSet[i-1][0], solutionSet[i-1][1]))/12

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforthMoulton4 returns solutions for the fourth order Adams-Bashforth-Moulton predictor-corrector method
func AdamsBashforthMoulton4(a float32, b float32, N int, initialCondition float32, f func(x, y float32) float32) [][]float32 {
	stepSize := (b - a) / float32(N)
	theta := a
	omega := initialCondition

	solutionSet := make([][]float32, N+1)

	for i := 0; i < N+1; i++ {
		solutionSet[i] = make([]float32, 2)
	}

	solutionSet[0][0] = theta
	solutionSet[0][1] = omega

	var kappa float32
	var kappa2 float32
	var kappa3 float32
	var kappa4 float32

	for i := 0; i < 3; i++ {
		kappa = stepSize * f(theta, omega)
		kappa2 = stepSize * f(theta+stepSize/2, omega+kappa/2)
		kappa3 = stepSize * f(theta+stepSize/2, omega+kappa2/2)
		kappa4 = stepSize * f(theta+stepSize, omega+kappa3)

		omega += (kappa + 2*kappa2 + 2*kappa3 + kappa4) / 6
		theta += stepSize

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	for i := 3; i < N; i++ {
		theta += stepSize
		omega = solutionSet[i][1] + stepSize*(55*f(solutionSet[i][0], solutionSet[i][1])-
			59*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			37*f(solutionSet[i-2][0], solutionSet[i-2][1])-
			9*f(solutionSet[i-3][0], solutionSet[i-3][1]))/24
		omega = solutionSet[i][1] + stepSize*(9*f(theta, omega)+
			19*f(solutionSet[i][0], solutionSet[i][1])-
			5*f(solutionSet[i-1][0], solutionSet[i-1][1])+
			f(solutionSet[i-2][0], solutionSet[i-2][1]))/24

		solutionSet[i+1][0] = theta
		solutionSet[i+1][1] = omega
	}

	return solutionSet
}

// AdamsBashforthMoulton returns a solution from the variable step Adams-Bashforth-Moulton method
func AdamsBashforthMoulton(a float32, b float32, initialCondition float32,
	TOL float32, maxStep float32, minStep float32, f func(x, y float32) float32) ([][]float32, error) {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false
	rk4Done := false
	lastValueCalc := false

	var thetas []float32
	var omegas []float32

	thetas = append(thetas, theta)
	omegas = append(omegas, omega)

	RK4 := func(h float32, tSet, oSet []float32, f func(x, y float32) float32) ([]float32, []float32) {
		var kappa float32
		var kappa2 float32
		var kappa3 float32
		var kappa4 float32
		var t float32
		var o float32

		for i := 0; i < 3; i++ {
			t = tSet[len(tSet)-1]
			o = oSet[len(oSet)-1]

			kappa = h * f(t, o)
			kappa2 = h * f(t+h/2, o+kappa/2)
			kappa3 = h * f(t+h/2, o+kappa2/2)
			kappa4 = h * f(t+h, o+kappa3)

			o += (kappa + 2*kappa2 + 2*kappa3 + kappa4) / 6.0
			t = tSet[len(tSet)-1] + h

			tSet, oSet = append(tSet, t), append(oSet, o)
		}

		return tSet, oSet
	}

	thetas, omegas = RK4(stepSize, thetas, omegas, f)

	rk4Done = true

	theta = thetas[len(thetas)-1] + stepSize
	var predictor float32
	var corrector float32
	var sigma float32
	var zeta float32

	for !done {
		predictor = omegas[len(thetas)-1] + stepSize*(55*f(thetas[len(thetas)-1], omegas[len(omegas)-1])-
			59*f(thetas[len(thetas)-2], omegas[len(omegas)-2])+
			37*f(thetas[len(thetas)-3], omegas[len(omegas)-3])-
			9*f(thetas[len(thetas)-4], omegas[len(omegas)-4]))/24
		corrector = omegas[len(thetas)-1] + stepSize*(9*f(theta, predictor)+
			19*f(thetas[len(thetas)-1], omegas[len(omegas)-1])-
			5*f(thetas[len(thetas)-2], omegas[len(omegas)-2])+
			f(thetas[len(thetas)-3], omegas[len(omegas)-3]))/24

		sigma = 19 * float32(math.Abs(float64(corrector-predictor))) / (270 * stepSize)
		if sigma <= TOL {
			omega = corrector

			thetas, omegas = append(thetas, theta), append(omegas, omega)

			if lastValueCalc {
				done = true
			} else {
				if sigma <= 0.1*TOL || thetas[len(thetas)-1]+stepSize > b {
					zeta = float32(math.Pow(float64(TOL/(2*sigma)), 1/4.0))
					if zeta > 4 {
						stepSize = 4 * stepSize
					} else {
						stepSize = zeta * stepSize
					}

					if stepSize > maxStep {
						stepSize = maxStep
					}

					if thetas[len(thetas)-1]+4*stepSize > b {
						stepSize = (b - thetas[len(thetas)-1]) / 4
						lastValueCalc = true
					}

					thetas, omegas = RK4(stepSize, thetas, omegas, f)
					rk4Done = true
				}
			}
		} else {
			zeta = float32(math.Pow(float64(TOL/(2*sigma)), 1/4.0))

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

	var solutionSet [][]float32

	if !lastValueCalc {
		return solutionSet, errors.New("Minimum step size exceeded")
	}

	for i := 0; i < len(thetas); i++ {
		solutionSet = append(solutionSet, []float32{thetas[i], omegas[i]})
	}

	return solutionSet, nil
}
