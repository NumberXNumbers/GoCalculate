package methods

import (
	"math"
	"testing"
)

func TestEuler1D(t *testing.T) {
	f := func(x, omega float32) float32 {
		return omega - x*x + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initValue := float32(0.5)
	result := Euler1D(a, b, N, initValue, f)
	if (result - 4.865784) > 0.000001 {
		t.Fail()
	}
}

func TestTrapezoidRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := TrapezoidRule(a, b, f)
	if (result - 0.2776801) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpsonRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := SimpsonRule(a, b, f)
	if (result - 0.2929326) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpson38Rule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := Simpson38Rule(a, b, f)
	if (result - 0.2929107) > 0.0000001 {
		t.Fail()
	}
}

func TestBooleRule(t *testing.T) {
	f := func(x float32) float32 {
		return float32(math.Sin(float64(x)))
	}
	a := float32(0.0)
	b := float32(math.Pi / 4)
	result := BooleRule(a, b, f)
	if (result - 0.29289318) > 0.0000001 {
		t.Fail()
	}
}

func TestRungeKutta(t *testing.T) {
	f := func(x, y float32) float32 {
		return y - float32(math.Pow(float64(x), 2)) + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initialCondition := float32(0.5)
	solutionMatrixA := RungeKutta2(a, b, N, initialCondition, f)
	if result := float64(solutionMatrixA[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := RungeKutta4(a, b, N, initialCondition, f)
	if result := float64(solutionMatrixB[10][1]); math.Abs(result-5.3054720) > 1e-2 {
		t.Fail()
	}
	TOL := float32(1e-5)
	maxStep := float32(0.25)
	minStep := float32(0.01)
	solutionMatrixC := RungeKuttaFehlbery(a, b, initialCondition, TOL, maxStep, minStep, f)
	if result := float64(solutionMatrixC[9][1]); math.Abs(result-5.3054720) > 1e-4 {
		t.Fail()
	}
}

func TestModifiedEuler(t *testing.T) {
	f := func(x, y float32) float32 {
		return y - float32(math.Pow(float64(x), 2)) + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initialCondition := float32(0.5)
	solutionMatrix := ModifiedEuler(a, b, N, initialCondition, f)
	if result := float64(solutionMatrix[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestHeun(t *testing.T) {
	f := func(x, y float32) float32 {
		return y - float32(math.Pow(float64(x), 2)) + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initialCondition := float32(0.5)
	solutionMatrix := Heun(a, b, N, initialCondition, f)
	if result := float64(solutionMatrix[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestAdamsBashforth(t *testing.T) {
	f := func(x, y float32) float32 {
		return y - float32(math.Pow(float64(x), 2)) + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initialCondition1 := float32(0.5)
	initialCondition2 := float32(0.8292986)
	initialCondition3 := float32(1.2140877)
	initialCondition4 := float32(1.6489406)
	initialCondition5 := float32(2.1272295)
	solutionMatrixA := AdamsBashforth2(a, b, N, initialCondition1, initialCondition2, f)
	if result := float64(solutionMatrixA[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := AdamsBashforth3(a, b, N, initialCondition1, initialCondition2, initialCondition3, f)
	if result := float64(solutionMatrixB[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixC := AdamsBashforth4(a, b, N, initialCondition1, initialCondition2, initialCondition3, initialCondition4, f)

	if result := float64(solutionMatrixC[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixD := AdamsBashforth5(a, b, N, initialCondition1,
		initialCondition2, initialCondition3, initialCondition4, initialCondition5, f)
	if result := float64(solutionMatrixD[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestAdamsBashforthMoulton(t *testing.T) {
	f := func(x, y float32) float32 {
		return y - float32(math.Pow(float64(x), 2)) + 1
	}
	a := float32(0.0)
	b := float32(2.0)
	N := 10
	initialCondition := float32(0.5)
	solutionMatrixA := AdamsBashforthMoulton3(a, b, N, initialCondition, f)
	if result := float64(solutionMatrixA[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := AdamsBashforthMoulton4(a, b, N, initialCondition, f)
	if result := float64(solutionMatrixB[10][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	maxStep := float32(0.2)
	minStep := float32(0.01)
	TOL := float32(1e-5)
	solutionMatrixC, err := AdamsBashforthMoulton(a, b, initialCondition, TOL, maxStep, minStep, f)
	if err != nil {
		t.Fail()
	}
	if result := float64(solutionMatrixC[21][1]); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}
