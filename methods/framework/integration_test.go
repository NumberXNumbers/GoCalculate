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
