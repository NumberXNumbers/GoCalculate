package methods

import (
	"errors"
	"math"

	"github.com/NumberXNumbers/GoCalculate/types/gcf"
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// Euler1D is for solving the numerical integration 1D euler method
func Euler1D(a float64, b float64, N int, initValue float64, f *gcf.Function) gcv.Value {
	h := (b - a) / float64(N)
	x := a
	omega := initValue

	for i := 0; i < N; i++ {
		omega += h * f.Eval(x, omega).Value().Real()
		x += h
	}

	return gcv.MakeValue(omega)
}

// TrapezoidRule is for solving the numerical integration using the trapezoid rule
func TrapezoidRule(a float64, b float64, f *gcf.Function) gcv.Value {
	var omega float64
	h := (b - a)
	x := a

	omega = f.Eval(x+h).Value().Real() + f.Eval(x).Value().Real()

	return gcv.MakeValue(h / 2 * omega)
}

// SimpsonRule for solving numerical integration
func SimpsonRule(a float64, b float64, f *gcf.Function) gcv.Value {
	var omega float64
	h := (b - a) / 2
	x := a

	omega = f.Eval(x).Value().Real() + 4*f.Eval(x+h).Value().Real() + f.Eval(x+2*h).Value().Real()

	return gcv.MakeValue(h / 3 * omega)
}

// Simpson38Rule is Simpson's 3/8ths rule for solving numerical integration
func Simpson38Rule(a float64, b float64, f *gcf.Function) gcv.Value {
	var omega float64
	h := (b - a) / 3
	x := a

	omega = f.Eval(x).Value().Real() + 3*f.Eval(x+h).Value().Real() + 3*f.Eval(x+2*h).Value().Real() + f.Eval(x+3*h).Value().Real()

	return gcv.MakeValue(3 * h / 8 * omega)
}

// BooleRule is Boole's rule for solving numerical integration
func BooleRule(a float64, b float64, f *gcf.Function) gcv.Value {
	var omega float64
	h := (b - a) / 4
	x := a

	omega = 7*f.Eval(x).Value().Real() +
		32*f.Eval(x+h).Value().Real() +
		12*f.Eval(x+2*h).Value().Real() +
		32*f.Eval(x+3*h).Value().Real() +
		7*f.Eval(x+4*h).Value().Real()

	return gcv.MakeValue(2 * h / 45 * omega)
}

// RungeKutta2 or midpoint method returns a solution found using the 2nd order runge-kutta
func RungeKutta2(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64

	for i := 0; i < N; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/2.0, omega+kappa/2.0).Value().Real()

		omega += kappa2
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// ModifiedEuler returns a solution to the ModifiedEuler method
func ModifiedEuler(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64

	for i := 0; i < N; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		theta += stepSize
		kappa2 = stepSize * f.Eval(theta, omega+kappa).Value().Real()

		omega += (kappa + kappa2) / 2.0

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// Heun returns a solution to the 3rd order runge-kutta method (Heun method)
func Heun(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64
	var kappa3 float64

	for i := 0; i < N; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/3.0, omega+kappa/3.0).Value().Real()
		kappa3 = stepSize * f.Eval(theta+2.0*stepSize/3.0, omega+2.0*kappa2/3.0).Value().Real()

		omega += (kappa + 3.0*kappa3) / 4.0
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// RungeKutta4 returns a solution found using the 4th order runge-kutta method
func RungeKutta4(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64

	for i := 0; i < N; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/2.0, omega+kappa/2.0).Value().Real()
		kappa3 = stepSize * f.Eval(theta+stepSize/2.0, omega+kappa2/2.0).Value().Real()
		kappa4 = stepSize * f.Eval(theta+stepSize, omega+kappa3).Value().Real()

		omega += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// RungeKuttaFehlbery returns a solution to the runge-kutta-fehlbery method
// Algorithm from Numerical Analysis - By Burden and Faires
func RungeKuttaFehlbery(a float64, b float64, initialCondition float64,
	TOL float64, maxStep float64, minStep float64, f *gcf.Function) m.Matrix {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false

	solutionSet := v.MakeVectors(v.RowSpace, v.MakeVectorPure(v.RowSpace, theta, omega))

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64
	var kappa5 float64
	var kappa6 float64

	var remainder float64
	var delta float64

	for !done {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/4.0, omega+kappa/4.0).Value().Real()
		kappa3 = stepSize * f.Eval(theta+3.0*stepSize/8.0, omega+3.0*kappa/32.0+9.0*kappa2/32.0).Value().Real()
		kappa4 = stepSize * f.Eval(theta+12.0*stepSize/13.0, omega+1932.0*kappa/2197.0-
			7200.0*kappa2/2197.0+7296.0*kappa3/2197.0).Value().Real()
		kappa5 = stepSize * f.Eval(theta+stepSize, omega+439.0*kappa/216.0-8.0*kappa2+
			3680.0*kappa3/513.0-845.0*kappa4/4104.0).Value().Real()
		kappa6 = stepSize * f.Eval(theta+stepSize/2.0, omega-8.0*kappa/27.0+2.0*kappa2-
			3544.0*kappa3/2565.0+1859.0*kappa4/4104.0-11.0*kappa5/40.0).Value().Real()

		remainder = math.Abs(kappa/360.0-128.0*kappa3/4275.0-2197.0*kappa4/75240.0+kappa5/50.0+2.0*kappa6/55.0) / stepSize
		if remainder <= TOL {
			theta += stepSize
			omega += 25.0*kappa/216.0 + 1408.0*kappa3/2565.0 + 2197.0*kappa4/4104.0 - kappa5/5.0

			solutionSet.Append(v.MakeVectorPure(v.RowSpace, theta, omega))
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

	return m.MakeMatrixAlt(solutionSet)
}

// AdamsBashforth2 returns a solution found using the 2nd order Adams-Bashforth method
func AdamsBashforth2(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2

	solutionSet := m.NewMatrix(N+1, 2)

	for i := 0; i < 2; i++ {
		solutionSet.Set(i, 0, theta)
		theta += stepSize
	}

	solutionSet.Set(0, 1, omega1)
	solutionSet.Set(1, 1, omega2)

	omega := omega2
	scalingConstantFactor := stepSize / 2.0

	var kappa float64
	var kappa2 float64

	for i := 2; i < N; i++ {
		kappa = scalingConstantFactor * 3.0 * f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()
		kappa2 = scalingConstantFactor * f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()

		omega += (kappa - kappa2)
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforth3 returns a solution found using the 3rd order Adams-Bashforth method
func AdamsBashforth3(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3

	solutionSet := m.NewMatrix(N+1, 2)

	for i := 0; i < 3; i++ {
		solutionSet.Set(i, 0, theta)
		theta += stepSize
	}

	solutionSet.Set(0, 1, omega1)
	solutionSet.Set(1, 1, omega2)
	solutionSet.Set(2, 1, omega3)

	omega := omega3
	scalingConstantFactor := stepSize / 12.0

	var kappa float64
	var kappa2 float64
	var kappa3 float64

	for i := 2; i < N; i++ {
		kappa = scalingConstantFactor * 23.0 * f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()
		kappa2 = scalingConstantFactor * 16.0 * f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()
		kappa3 = scalingConstantFactor * 5.0 * f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real()

		omega += (kappa - kappa2 + kappa3)
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforth4 returns a solution found using the 4th order Adams-Bashforth method
func AdamsBashforth4(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, initialCondition4 float64,
	f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4

	solutionSet := m.NewMatrix(N+1, 2)

	for i := 0; i < 4; i++ {
		solutionSet.Set(i, 0, theta)
		theta += stepSize
	}

	solutionSet.Set(0, 1, omega1)
	solutionSet.Set(1, 1, omega2)
	solutionSet.Set(2, 1, omega3)
	solutionSet.Set(3, 1, omega4)

	omega := omega4
	scalingConstantFactor := stepSize / 24.0

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64

	for i := 3; i < N; i++ {
		kappa = scalingConstantFactor * 55.0 * f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()
		kappa2 = scalingConstantFactor * 59.0 * f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()
		kappa3 = scalingConstantFactor * 37.0 * f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real()
		kappa4 = scalingConstantFactor * 9.0 * f.Eval(solutionSet.Get(i-3, 0), solutionSet.Get(i-3, 1)).Value().Real()

		omega += (kappa - kappa2 + kappa3 - kappa4)
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforth5 returns a solution found using the 5th order Adams-Bashforth method
func AdamsBashforth5(a float64, b float64, N int, initialCondition1 float64,
	initialCondition2 float64, initialCondition3 float64, initialCondition4 float64,
	initialCondition5 float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega1 := initialCondition1
	omega2 := initialCondition2
	omega3 := initialCondition3
	omega4 := initialCondition4
	omega5 := initialCondition5

	solutionSet := m.NewMatrix(N+1, 2)

	for i := 0; i < 5; i++ {
		solutionSet.Set(i, 0, theta)
		theta += stepSize
	}

	solutionSet.Set(0, 1, omega1)
	solutionSet.Set(1, 1, omega2)
	solutionSet.Set(2, 1, omega3)
	solutionSet.Set(3, 1, omega4)
	solutionSet.Set(4, 1, omega5)

	omega := omega5
	scalingConstantFactor := stepSize / 720.0

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64
	var kappa5 float64

	for i := 3; i < N; i++ {
		kappa = scalingConstantFactor * 1901.0 * f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()
		kappa2 = scalingConstantFactor * 2774.0 * f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()
		kappa3 = scalingConstantFactor * 2616.0 * f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real()
		kappa4 = scalingConstantFactor * 1274.0 * f.Eval(solutionSet.Get(i-3, 0), solutionSet.Get(i-3, 1)).Value().Real()
		kappa5 = scalingConstantFactor * 251.0 * f.Eval(solutionSet.Get(i-3, 0), solutionSet.Get(i-3, 1)).Value().Real()

		omega += (kappa - kappa2 + kappa3 - kappa4 + kappa5)
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforthMoulton3 returns solutions for the third order Adams-Bashforth-Moulton predictor-corrector method
func AdamsBashforthMoulton3(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64
	var kappa3 float64

	for i := 0; i < 2; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/3.0, omega+kappa/3.0).Value().Real()
		kappa3 = stepSize * f.Eval(theta+2.0*stepSize/3.0, omega+2.0*kappa2/3.0).Value().Real()

		omega += (kappa + 3.0*kappa3) / 4.0
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	for i := 2; i < N; i++ {
		theta += stepSize
		omega = solutionSet.Get(i, 1).Real() + stepSize*(23.0*f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()-
			16.0*f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()+
			5.0*f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real())/12.0
		omega = solutionSet.Get(i, 1).Real() + stepSize*(5.0*f.Eval(theta, omega).Value().Real()+
			8.0*f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()-
			f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real())/12.0

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforthMoulton4 returns solutions for the fourth order Adams-Bashforth-Moulton predictor-corrector method
func AdamsBashforthMoulton4(a float64, b float64, N int, initialCondition float64, f *gcf.Function) m.Matrix {
	stepSize := (b - a) / float64(N)
	theta := a
	omega := initialCondition

	solutionSet := m.NewMatrix(N+1, 2)

	solutionSet.Set(0, 0, theta)
	solutionSet.Set(0, 1, omega)

	var kappa float64
	var kappa2 float64
	var kappa3 float64
	var kappa4 float64

	for i := 0; i < 3; i++ {
		kappa = stepSize * f.Eval(theta, omega).Value().Real()
		kappa2 = stepSize * f.Eval(theta+stepSize/2.0, omega+kappa/2.0).Value().Real()
		kappa3 = stepSize * f.Eval(theta+stepSize/2.0, omega+kappa2/2.0).Value().Real()
		kappa4 = stepSize * f.Eval(theta+stepSize, omega+kappa3).Value().Real()

		omega += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
		theta += stepSize

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	for i := 3; i < N; i++ {
		theta += stepSize
		omega = solutionSet.Get(i, 1).Real() + stepSize*(55.0*f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()-
			59.0*f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()+
			37.0*f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real()-
			9.0*f.Eval(solutionSet.Get(i-3, 0), solutionSet.Get(i-3, 1)).Value().Real())/24.0
		omega = solutionSet.Get(i, 1).Real() + stepSize*(9.0*f.Eval(theta, omega).Value().Real()+
			19.0*f.Eval(solutionSet.Get(i, 0), solutionSet.Get(i, 1)).Value().Real()-
			5.0*f.Eval(solutionSet.Get(i-1, 0), solutionSet.Get(i-1, 1)).Value().Real()+
			f.Eval(solutionSet.Get(i-2, 0), solutionSet.Get(i-2, 1)).Value().Real())/24.0

		solutionSet.Set(i+1, 0, theta)
		solutionSet.Set(i+1, 1, omega)
	}

	return solutionSet
}

// AdamsBashforthMoulton returns a solution from the variable step Adams-Bashforth-Moulton method
func AdamsBashforthMoulton(a float64, b float64, initialCondition float64,
	TOL float64, maxStep float64, minStep float64, f *gcf.Function) (m.Matrix, error) {
	stepSize := maxStep
	theta := a
	omega := initialCondition
	done := false
	rk4Done := false
	lastValueCalc := false

	set := v.MakeVectors(v.RowSpace, v.MakeVectorPure(v.RowSpace, theta, omega))

	RK4 := func(h float64, set v.Vectors, f *gcf.Function) v.Vectors {
		var kappa float64
		var kappa2 float64
		var kappa3 float64
		var kappa4 float64
		var t float64
		var o float64

		for i := 0; i < 3; i++ {
			t = set.Get(set.Len() - 1).Get(0).Real()
			o = set.Get(set.Len() - 1).Get(1).Real()

			kappa = h * f.Eval(t, o).Value().Real()
			kappa2 = h * f.Eval(t+h/2.0, o+kappa/2.0).Value().Real()
			kappa3 = h * f.Eval(t+h/2.0, o+kappa2/2.0).Value().Real()
			kappa4 = h * f.Eval(t+h, o+kappa3).Value().Real()

			o += (kappa + 2.0*kappa2 + 2.0*kappa3 + kappa4) / 6.0
			t += h

			set.Append(v.MakeVectorPure(v.RowSpace, t, o))
		}

		return set
	}

	set = RK4(stepSize, set, f)
	rk4Done = true

	theta = set.Get(set.Len()-1).Get(0).Real() + stepSize
	var predictor float64
	var corrector float64
	var sigma float64
	var zeta float64

	for !done {
		theta1, omega1 := set.Get(set.Len()-1).Get(0).Real(), set.Get(set.Len()-1).Get(1).Real()
		theta2, omega2 := set.Get(set.Len()-2).Get(0).Real(), set.Get(set.Len()-2).Get(1).Real()
		theta3, omega3 := set.Get(set.Len()-3).Get(0).Real(), set.Get(set.Len()-3).Get(1).Real()
		theta4, omega4 := set.Get(set.Len()-4).Get(0).Real(), set.Get(set.Len()-4).Get(1).Real()
		predictor = omega1 + stepSize*(55.0*f.Eval(theta1, omega1).Value().Real()-
			59.0*f.Eval(theta2, omega2).Value().Real()+
			37.0*f.Eval(theta3, omega3).Value().Real()-
			9.0*f.Eval(theta4, omega4).Value().Real())/24.0
		corrector = omega1 + stepSize*(9.0*f.Eval(theta, predictor).Value().Real()+
			19.0*f.Eval(theta1, omega1).Value().Real()-
			5.0*f.Eval(theta2, omega2).Value().Real()+
			f.Eval(theta3, omega3).Value().Real())/24.0

		sigma = 19.0 * math.Abs(corrector-predictor) / (270.0 * stepSize)
		if sigma <= TOL {
			omega = corrector

			set.Append(v.MakeVectorPure(v.RowSpace, theta, omega))

			if lastValueCalc {
				done = true
			} else {
				if sigma <= 0.1*TOL || theta+stepSize > b {
					zeta = math.Pow(TOL/(2.0*sigma), 1.0/4.0)
					if zeta > 4 {
						stepSize = 4.0 * stepSize
					} else {
						stepSize = zeta * stepSize
					}

					if stepSize > maxStep {
						stepSize = maxStep
					}

					if theta+4.0*stepSize > b {
						stepSize = (b - theta) / 4.0
						lastValueCalc = true
					}

					set = RK4(stepSize, set, f)
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
					set = set.Subset(0, set.Len()-4)
				}

				set = RK4(stepSize, set, f)
				rk4Done = true
			}
		}

		theta = set.Get(set.Len()-1).Get(0).Real() + stepSize
	}

	if !lastValueCalc {
		return nil, errors.New("Minimum step size exceeded")
	}

	return m.MakeMatrixAlt(set), nil
}
