package gcf

import (
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestFn1(t *testing.T) {
	x := NewVar(Value)
	y := NewVar(Value)
	regVars := []Var{x, y}
	constant := MakeConst(0)
	function := MakeFunc(regVars, MakeConst(4), "+", MakeConst(5), "-", y, "*", MakeConst(3), "/", MakeConst(7), "+", x, "+", MakeConst(constant))
	value := function.Eval(4, 3)
	if value.Value().Real() != 11.714285714285715 {
		t.Fail()
	}
}

func TestFn2(t *testing.T) {
	x := NewVar(Value)
	y := NewVar(Matrix)
	regVars := []Var{x, y}
	vector := v.MakeVectorPure(v.RowSpace, 2, 4, 6)
	constVect := MakeConst(vector)
	function := MakeFunc(regVars, constVect, "*", y, "*", x, "/", MakeConst(4))
	matrix := m.NewIdentityMatrix(3)
	value := function.Eval(2, matrix)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 2 ||
		value.Vector().Get(2).Real() != 3 {
		t.Fail()
	}
}

func TestFn3(t *testing.T) {
	x := NewVar(Vector)
	regVars := []Var{x}
	vector := v.MakeVectorPure(v.RowSpace, 2, 4, 6)
	function := MakeFunc(regVars, x, "+", "(", MakeConst(5), "*", x, ")")
	value := function.Eval(vector)
	if value.Vector().Get(0).Real() != 12 ||
		value.Vector().Get(1).Real() != 24 ||
		value.Vector().Get(2).Real() != 36 {
		t.Fail()
	}
}

func TestFn4(t *testing.T) {
	x := NewVar(Matrix)
	y := NewVar(Matrix)
	a := NewVar(Vector)
	b := NewVar(Vector)
	regVars := []Var{x, y, a, b}
	matrixA := m.NewIdentityMatrix(3)
	matrixB := m.NewIdentityMatrix(3)
	vectorA := v.MakeVectorPure(v.RowSpace, 2, 4, 6)
	vectorB := v.MakeVectorPure(v.RowSpace, 2, 4, 6)
	function := MakeFunc(regVars, a, "*", "(", x, "+", "(", y, "*", MakeConst(2), "-", x, ")", "/", MakeConst(2), ")", "-", b, "/", MakeConst(2))
	value := function.Eval(matrixA, matrixB, vectorA, vectorB)
	if value.Vector().Get(0).Real() != 2 ||
		value.Vector().Get(1).Real() != 4 ||
		value.Vector().Get(2).Real() != 6 {
		t.Fail()
	}
}

func TestFn5(t *testing.T) {
	x := NewVar(Matrix)
	y := NewVar(Matrix)
	a := NewVar(Vector)
	b := NewVar(Vector)
	regVars := []Var{x, y, a, b}
	matrixA := m.NewIdentityMatrix(2)
	matrixB := m.NewIdentityMatrix(2)
	vectorA := v.MakeVectorPure(v.RowSpace, 1, 0)
	vectorB := v.MakeVectorPure(v.ColSpace, 0, 1)
	function := MakeFunc(regVars, a, "*", b, "*", x, "+", y, "+", b, "*", a)
	value := function.Eval(matrixA, matrixB, vectorA, vectorB)
	if value.Matrix().Get(0, 0).Real() != 1 ||
		value.Matrix().Get(0, 1).Real() != 0 ||
		value.Matrix().Get(1, 0).Real() != 1 ||
		value.Matrix().Get(1, 1).Real() != 1 {
		t.Fail()
	}
}

func TestFn6(t *testing.T) {
	x := NewVar(Matrix)
	y := NewVar(Matrix)
	a := NewVar(Vector)
	regVars := []Var{x, y, a}
	matrixA := m.NewIdentityMatrix(2)
	matrixB := m.NewIdentityMatrix(2)
	vectorA := v.MakeVectorPure(v.ColSpace, 1, 0)
	function := MakeFunc(regVars, x, "*", y, "*", a)
	value := function.Eval(matrixA, matrixB, vectorA)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 0 {
		t.Fail()
	}
}
