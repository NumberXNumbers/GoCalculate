package methods

import (
	"math"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcf"
)

func TestEuler1D(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	omega := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, omega}
	f := gcf.MakeFunc(regVars, omega, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initValue := 0.5
	result := Euler1D(a, b, N, initValue, f)
	if (result.Real() - 4.865784) > 0.000001 {
		t.Fail()
	}
}

func TestTrapezoidRule(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x}
	f := gcf.MakeFunc(regVars, "Sin", "(", x, ")")
	a := 0.0
	b := math.Pi / 4
	result := TrapezoidRule(a, b, f)
	if (result.Real() - 0.2776801) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpsonRule(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x}
	f := gcf.MakeFunc(regVars, "Sin", "(", x, ")")
	a := 0.0
	b := math.Pi / 4
	result := SimpsonRule(a, b, f)
	if (result.Real() - 0.2929326) > 0.0000001 {
		t.Fail()
	}
}

func TestSimpson38Rule(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x}
	f := gcf.MakeFunc(regVars, "Sin", "(", x, ")")
	a := 0.0
	b := math.Pi / 4
	result := Simpson38Rule(a, b, f)
	if (result.Real() - 0.2929107) > 0.0000001 {
		t.Fail()
	}
}

func TestBooleRule(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x}
	f := gcf.MakeFunc(regVars, "Sin", "(", x, ")")
	a := 0.0
	b := math.Pi / 4
	result := BooleRule(a, b, f)
	if (result.Real() - 0.29289318) > 0.0000001 {
		t.Fail()
	}
}

func TestRungeKutta(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	y := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, y}
	f := gcf.MakeFunc(regVars, y, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrixA := RungeKutta2(a, b, N, initialCondition, f)
	if result := solutionMatrixA.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := RungeKutta4(a, b, N, initialCondition, f)
	if result := solutionMatrixB.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-2 {
		t.Fail()
	}
	TOL := 1e-5
	maxStep := 0.25
	minStep := 0.01
	solutionMatrixC := RungeKuttaFehlbery(a, b, initialCondition, TOL, maxStep, minStep, f)
	if result := solutionMatrixC.Get(9, 1).Real(); math.Abs(result-5.3054720) > 1e-4 {
		t.Fail()
	}
}

func TestModifiedEuler(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	y := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, y}
	f := gcf.MakeFunc(regVars, y, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrix := ModifiedEuler(a, b, N, initialCondition, f)
	if result := solutionMatrix.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestHeun(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	y := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, y}
	f := gcf.MakeFunc(regVars, y, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrix := Heun(a, b, N, initialCondition, f)
	if result := solutionMatrix.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestAdamsBashforth(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	y := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, y}
	f := gcf.MakeFunc(regVars, y, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initialCondition1 := 0.5
	initialCondition2 := 0.8292986
	initialCondition3 := 1.2140877
	initialCondition4 := 1.6489406
	initialCondition5 := 2.1272295
	solutionMatrixA := AdamsBashforth2(a, b, N, initialCondition1, initialCondition2, f)
	if result := solutionMatrixA.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := AdamsBashforth3(a, b, N, initialCondition1, initialCondition2, initialCondition3, f)
	if result := solutionMatrixB.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixC := AdamsBashforth4(a, b, N, initialCondition1, initialCondition2, initialCondition3, initialCondition4, f)

	if result := solutionMatrixC.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixD := AdamsBashforth5(a, b, N, initialCondition1,
		initialCondition2, initialCondition3, initialCondition4, initialCondition5, f)
	if result := solutionMatrixD.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}

func TestAdamsBashforthMoulton(t *testing.T) {
	x := gcf.NewVar(gcf.Value)
	y := gcf.NewVar(gcf.Value)
	regVars := []gcf.Var{x, y}
	f := gcf.MakeFunc(regVars, y, "-", x, "^", 2, "+", 1)
	a := 0.0
	b := 2.0
	N := 10
	initialCondition := 0.5
	solutionMatrixA := AdamsBashforthMoulton3(a, b, N, initialCondition, f)
	if result := solutionMatrixA.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	solutionMatrixB := AdamsBashforthMoulton4(a, b, N, initialCondition, f)
	if result := solutionMatrixB.Get(10, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
	maxStep := 0.2
	minStep := 0.01
	TOL := 1e-5
	solutionMatrixC, err := AdamsBashforthMoulton(a, b, initialCondition, TOL, maxStep, minStep, f)
	if err != nil {
		t.Fail()
	}
	if result := solutionMatrixC.Get(21, 1).Real(); math.Abs(result-5.3054720) > 1e-1 {
		t.Fail()
	}
}
