package gcf

import (
	"fmt"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestFn0(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, x)
	value := function.Eval(4)
	if value.Value().Real() != 4 {
		t.Fail()
	}
}

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

func TestPanicAddVector(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionV := Add(v1, v2)

	if solutionV != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicAddMatrix(t *testing.T) {
	m1 := MakeConst(m.NewMatrix(2, 2))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionM := Add(m1, m2)

	if solutionM != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicAddMismatch(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Add(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicSubVector(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionV := Sub(v1, v2)

	if solutionV != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicSubMatrix(t *testing.T) {
	m1 := MakeConst(m.NewMatrix(2, 2))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionM := Sub(m1, m2)

	if solutionM != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicSubMismatch(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Sub(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicDivMismatch(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Div(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultDoubleRowVector(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Mult(v1, v2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultDoubleColVector(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.ColSpace, 3))
	v2 := MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Mult(v1, v2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultVM(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.ColSpace, 3))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Mult(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultMV(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Mult(m2, v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultMatrix(t *testing.T) {
	m1 := MakeConst(m.NewMatrix(2, 2))
	m2 := MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := Mult(m1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadValue(t *testing.T) {
	v := MakeConst(m.NewMatrix(2, 2))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := v.Value()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadVector(t *testing.T) {
	v := MakeConst(m.NewMatrix(2, 2))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := v.Vector()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadMatrix(t *testing.T) {
	m := MakeConst(gcv.NewValue())

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := m.Matrix()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadVariable(t *testing.T) {
	m := NewVar(Matrix)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := m.Eval(gcv.NewValue())

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicDuplicateRegVarsForFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x, x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, MakeConst(4), "+", MakeConst(5))

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotRegVarsForFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	y := NewVar(Value)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, y)

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotSupportedTypeForFunc(t *testing.T) {
	x := gcv.NewValue()
	regVars := []Var{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, x)

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotEnoughArgsFunc(t *testing.T) {
	x := NewVar(Matrix)
	regVars := []Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, x)

	function.Eval()

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicOperatorOperandMismatchFuncA(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, "+", x, x)

	function.Eval(1)

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicOperatorOperandMismatchFuncB(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, x, x, "+")

	function.Eval(1)

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadgetOpFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, x)

	function.getOp(0)

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadgetVarFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFunc(regVars, x, "+", x)

	function.getVar(1)

	if function != nil {
		t.Error("Expected Panic")
	}
}
