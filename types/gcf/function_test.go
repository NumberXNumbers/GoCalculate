package gcf

import (
	"fmt"
	"math"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestFn0(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, x)
	value := function.MustEval(4)
	if value.Value().Real() != 4 {
		t.Fail()
	}
}

func TestFn1(t *testing.T) {
	x := NewVar(Value)
	y := NewVar(Value)
	regVars := []Var{x, y}
	constant := MakeConst(0)
	function := MakeFunc(regVars, MakeConst(4), "+", 5, "-", y, "*", MakeConst(3), "/", MakeConst(7), "+", x, "+", MakeConst(constant))
	value := function.MustEval(4, 3)
	if value.Value().Real() != 11.714285714285715 {
		t.Fail()
	}
}

func TestFn2(t *testing.T) {
	x := NewVar(Value)
	y := NewVar(Matrix)
	regVars := []Var{x, y}
	vector := v.MakeVector(v.RowSpace, 2, 4, 6)
	constVect := MakeConst(vector)
	function := MakeFunc(regVars, constVect, "*", y, "*", x, "/", MakeConst(4))
	matrix := m.NewIdentityMatrix(3)
	value := function.MustEval(2, matrix)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 2 ||
		value.Vector().Get(2).Real() != 3 {
		t.Fail()
	}
}

func TestFn3(t *testing.T) {
	x := NewVar(Vector)
	regVars := []Var{x}
	vector := v.MakeVector(v.RowSpace, 2, 4, 6)
	function := MakeFunc(regVars, x, "+", "(", MakeConst(5), "*", x, ")")
	value := function.MustEval(vector)
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
	vectorA := v.MakeVector(v.RowSpace, 2, 4, 6)
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	function := MakeFunc(regVars, a, "*", "(", x, "+", "(", y, "*", MakeConst(2), "-", x, ")", "/", MakeConst(2), ")", "-", b, "/", MakeConst(2))
	value := function.MustEval(matrixA, matrixB, vectorA, vectorB)
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
	vectorA := v.MakeVector(v.RowSpace, 1, 0)
	vectorB := v.MakeVector(v.ColSpace, 0, 1)
	function := MakeFunc(regVars, a, "*", b, "*", x, "+", y, "+", b, "*", a)

	value := function.MustEval(matrixA, matrixB, vectorA, vectorB)
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
	vectorA := v.MakeVector(v.ColSpace, 1, 0)
	function := MakeFunc(regVars, x, "*", y, "*", a)
	value := function.MustEval(matrixA, matrixB, vectorA)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 0 {
		t.Fail()
	}
}

func TestFn7(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Sin", "(", x, ")")
	value := function.MustEval(math.Pi)
	if value.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestFn8(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Sin", x)
	value := function.MustEval(math.Pi)
	if value.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestFn9(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, x, "*", "Sin", "(", x, ")", "+", x)
	// fmt.Println(function.args)
	// fmt.Println(function.inputTypes)
	value := function.MustEval(math.Pi)
	if math.Abs(value.Value().Real()-3.1415926535897936) > 10e-15 {
		t.Fail()
	}
}

func TestFn10(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Sqrt", "(", x, "^", x, ")")
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-2) > 10e-15 {
		t.Fail()
	}
}

func TestFn11(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Cos", "(", x, ")", "*", "Sin", "(", x, ")")
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.5000000) > 10e-15 {
		t.Fail()
	}
}

func TestFn12(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Cos", "(", "Sin", "(", x, ")", ")")
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestFn13(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, "Cos", "Sin", x)
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestFn14(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, x, pow, "Sin", math.Pi/2, pow, x)
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-2) > 10e-6 {
		t.Fail()
	}
}

func TestFn15(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}
	function := MakeFunc(regVars, x, pow, 3, pow, x)
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-512) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateA(t *testing.T) {
	matrixA := m.NewIdentityMatrix(3)
	matrixB := m.NewIdentityMatrix(3)
	vectorA := v.MakeVector(v.RowSpace, 2, 4, 6)
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	calculation := MustCalculate(vectorA, "*", "(", matrixA, "+", "(", matrixB, "*", MakeConst(2), "-", matrixA, ")", "/", 2, ")", "-", vectorB, "/", gcv.MakeValue(2))
	if calculation.Vector().Get(0).Real() != 2 ||
		calculation.Vector().Get(1).Real() != 4 ||
		calculation.Vector().Get(2).Real() != 6 {
		t.Fail()
	}
}

func TestMustCalculateB(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	calculation := MustCalculate("(", MakeConst(2), ")", "*", vectorB)
	if calculation.Vector().Get(0).Real() != 4 ||
		calculation.Vector().Get(1).Real() != 8 ||
		calculation.Vector().Get(2).Real() != 12 {
		t.Fail()
	}
}

func TestMustCalculateC(t *testing.T) {
	calculation := MustCalculate("Sin", "(", math.Pi, ")")
	if calculation.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateD(t *testing.T) {
	calculation := MustCalculate(2, pow, 2, pow, 3)
	if calculation.Value().Real() == 8 {
		t.Fail()
	}
}

func TestMustCalculateE(t *testing.T) {
	calculation := MustCalculate("Cos", "(", "Sin", "(", math.Pi/4, ")", ")")
	if math.Abs(calculation.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateF(t *testing.T) {
	calculation := MustCalculate(math.Pi, "*", "Sin", "(", 2, "*", math.Pi, "-", math.Pi, ")", "+", math.Pi)
	// fmt.Println(function.args)
	// fmt.Println(function.inputTypes)
	if math.Abs(calculation.Value().Real()-3.1415926535897936) > 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateG(t *testing.T) {
	calculation := MustCalculate(2, "*", 2, "/", 2, "*", 2)
	if math.Abs(calculation.Value().Real()-4) > 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateH(t *testing.T) {
	calculation := MustCalculate("Cos", "Sin", math.Pi/4)
	if math.Abs(calculation.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateI(t *testing.T) {
	calculation := MustCalculate(2, pow, "Sin", math.Pi/2, pow, 2)
	if math.Abs(calculation.Value().Real()-2) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateJ(t *testing.T) {
	calculation := MustCalculate(2, "+", "Sin", math.Pi/2, "+", 2)
	if math.Abs(calculation.Value().Real()-5) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateK(t *testing.T) {
	calculation := MustCalculate("Sqrt", "(", "Sin", math.Pi/2, pow, 2, "+", "Cos", math.Pi/2, pow, 2, ")")
	if math.Abs(calculation.Value().Real()-1) > 10e-6 {
		t.Fail()
	}
}

func TestFunctionPanicOperatorNotSupported(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	regVars := []Var{}
	function := MakeFunc(regVars, "(", MakeConst(2), ")", "=", vectorB)

	value := function.MustEval()

	if value.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestFunctionPanicOperatorParensMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	regVars := []Var{}
	function := MakeFunc(regVars, 2, "+", MakeConst(2), ")", "+", vectorB)

	value := function.MustEval()

	if value.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicOperatorsOperandMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate("(", MakeConst(2), ")", ")", "*", vectorB)

	if calculation != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicOperatorParensMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate(2, "+", MakeConst(2), ")", "+", vectorB)

	if calculation.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicUnsupportedType(t *testing.T) {
	vectorB := NewVar(Vector)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate("(", MakeConst(2), ")", "*", vectorB)

	if calculation != nil {
		t.Error("Expected Panic")
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

	solutionV := MustAdd(v1, v2)

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

	solutionM := MustAdd(m1, m2)

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

	solution := MustAdd(v1, m2)

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

	solutionV := MustSub(v1, v2)

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

	solutionM := MustSub(m1, m2)

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

	solution := MustSub(v1, m2)

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

	solution := MustDiv(v1, m2)

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

	solution := MustMult(v1, v2)

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

	solution := MustMult(v1, v2)

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

	solution := MustMult(v1, m2)

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

	solution := MustMult(m2, v1)

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

	solution := MustMult(m1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadPow(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustPow(v1, v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadSqrt(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSqrt(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadSin(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSin(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadCos(t *testing.T) {
	v1 := MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustCos(v1)

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

	solution := m.MustEval(gcv.NewValue())

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
	x := uint(2)
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

	function.MustEval()

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadgetOpFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}

	function := MakeFunc(regVars, x)

	_, err := function.getOp(0)

	if err == nil {
		t.Error("Expected Error")
	}
}

func TestPanicBadgetVarFunc(t *testing.T) {
	x := NewVar(Value)
	regVars := []Var{x}

	function := MakeFunc(regVars, x, "+", x)

	_, err := function.getVar(2)

	if err == nil {
		t.Error("Expected Error")
	}
}
