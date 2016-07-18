package methods

import (
	"math"
	"testing"
)

func TestEuler1D(t *testing.T) {
	f := func(x, omega float64) float64 {
		return omega - x*x + 1
	}
	a := 0.0
	b := 2.0
	N := 10
	initValue := 0.5
	result := Euler1D(a, b, N, initValue, f)
	if (result - 4.865784) > 0.000001 {
		t.Fail()
	}
}

func TestTrapezoidRule(t *testing.T) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	a := 0.0
	b := math.Pi / 4
	result := TrapezoidRule(a, b, f)
	if (result - 0.2776801) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpsonRule(t *testing.T) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	a := 0.0
	b := math.Pi / 4
	result := SimpsonRule(a, b, f)
	if (result - 0.2929326) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpson38Rule(t *testing.T) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	a := 0.0
	b := math.Pi / 4
	result := Simpson38Rule(a, b, f)
	if (result - 0.2929107) > 0.0000001 {
		t.Fail()
	}
}

func TestBooleRule(t *testing.T) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	a := 0.0
	b := math.Pi / 4
	result := BooleRule(a, b, f)
	if (result - 0.29289318) > 0.0000001 {
		t.Fail()
	}
}

func TestRungeKutta(t *testing.T) {
	f := func(x, y float64) float64 {
		return y - math.Pow(float64(x), 2) + 1
	}
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrixA := RungeKutta2(a, b, N, initialCondition, f)
	if result := solutionMatrixA[10][1]; math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := RungeKutta4(a, b, N, initialCondition, f)
	if result := solutionMatrixB[10][1]; math.Abs(result-5.3054720) > 1e-2 {
		t.Fail()
	}
	TOL := 1e-5
	maxStep := 0.25
	minStep := 0.01
	solutionMatrixC := RungeKuttaFehlbery(a, b, initialCondition, TOL, maxStep, minStep, f)
	if result := solutionMatrixC[9][1]; math.Abs(result-5.3054720) > 1e-4 {
		t.Fail()
	}
}

func TestModifiedEuler(t *testing.T) {
	f := func(x, y float64) float64 {
		return y - math.Pow(x, 2) + 1
	}
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrix := ModifiedEuler(a, b, N, initialCondition, f)
	if result := solutionMatrix[10][1]; math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestHeun(t *testing.T) {
	f := func(x, y float64) float64 {
		return y - math.Pow(x, 2) + 1
	}
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrix := Heun(a, b, N, initialCondition, f)
	if result := solutionMatrix[10][1]; math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}
