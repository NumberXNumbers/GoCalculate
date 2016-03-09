package fops

import (
	"math"
	"math/cmplx"
	"testing"
)

func TestVariable(t *testing.T) {
	testVariableA := Variable(0)
	testVariableB := Variable(1)
	testVariableC := VariableComplex(0)
	testVariableD := VariableComplex(1)

	testFunctionA := Add(testVariableA, testVariableB)
	testFunctionB := AddComplex(testVariableC, testVariableD)

	if testVariableA(1, 0, 0) != 1.0 {
		t.Errorf("Expected %d, received %v", 1, testFunctionA(1, 0, 0))
	}

	if testVariableB(0, 1, 0) != 1.0 {
		t.Errorf("Expected %v, received %v", 1, testFunctionB(0, 1, 0))
	}

	if testVariableC(1, 0, 0) != complex128(1) {
		t.Errorf("Expected %v, received %v", complex128(1), testVariableC(1, 0, 0))
	}

	if testVariableD(0, 1, 0) != complex128(1) {
		t.Errorf("Expected %v, received %v", complex128(1), testVariableD(0, 1, 0))
	}

	if testFunctionA(1, 2) != 3.0 {
		t.Errorf("Expected %v, received %v", 3, testFunctionA(1, 2))
	}

	if testFunctionB(1, 2) != complex128(3) {
		t.Errorf("Expected %v, received %v", complex128(3), testFunctionA(1, 2))
	}
}

func TestNegativeVariable(t *testing.T) {
	testVariableA := NegativeVariable(0)
	testVariableB := NegativeVariable(1)
	testVariableC := NegativeVariableComplex(0)
	testVariableD := NegativeVariableComplex(1)

	testFunctionA := Add(testVariableA, testVariableB)
	testFunctionB := AddComplex(testVariableC, testVariableD)

	if testVariableA(1, 0, 0) != -1.0 {
		t.Errorf("Expected %d, received %v", -1, testFunctionA(1, 0, 0))
	}

	if testVariableB(0, 1, 0) != -1.0 {
		t.Errorf("Expected %v, received %v", -1, testFunctionB(0, 1, 0))
	}

	if testVariableC(1, 0, 0) != complex128(-1) {
		t.Errorf("Expected %v, received %v", complex128(-1), testVariableC(1, 0, 0))
	}

	if testVariableD(0, 1, 0) != complex128(-1) {
		t.Errorf("Expected %v, received %v", complex128(-1), testVariableD(0, 1, 0))
	}

	if testFunctionA(1, 2) != -3.0 {
		t.Errorf("Expected %v, received %v", -3, testFunctionA(1, 2))
	}

	if testFunctionB(1, 2) != complex128(-3) {
		t.Errorf("Expected %v, received %v", complex128(-3), testFunctionA(1, 2))
	}
}

func TestConstant(t *testing.T) {
	testVariableA := Constant(0)
	testVariableB := Constant(1)
	testVariableC := ConstantComplex(0)
	testVariableD := ConstantComplex(1)

	testFunctionA := Add(testVariableA, testVariableB)
	testFunctionB := AddComplex(testVariableC, testVariableD)

	if testVariableA(1, 1, 1) != 0.0 {
		t.Errorf("Expected %d, received %v", 0, testFunctionA(1, 1, 1))
	}

	if testVariableB(0, 0, 0) != 1.0 {
		t.Errorf("Expected %v, received %v", 1, testFunctionB(0, 0, 0))
	}

	if testVariableC(1, 1, 1) != complex128(0) {
		t.Errorf("Expected %v, received %v", complex128(0), testVariableC(1, 1, 1))
	}

	if testVariableD(0, 0, 0) != complex128(1) {
		t.Errorf("Expected %v, received %v", complex128(0), testVariableD(0, 0, 0))
	}

	if testFunctionA(1, 2) != 1.0 {
		t.Errorf("Expected %v, received %v", 1, testFunctionA(1, 2))
	}

	if testFunctionB(1, 2) != complex128(1) {
		t.Errorf("Expected %v, received %v", complex128(1), testFunctionA(1, 2))
	}
}

func TestAddSubtractMultiplyDivideParens(t *testing.T) {
	f := func(x ...float64) float64 {
		return x[0]
	}

	g := func(x ...float64) float64 {
		return x[1]
	}

	i := func(x ...complex128) complex128 {
		return x[0]
	}

	j := func(x ...complex128) complex128 {
		return x[1]
	}

	testFunctionA := Add(f, g)
	testFunctionB := AddComplex(i, j)

	testFunctionC := Subtract(f, g)
	testFunctionD := SubtractComplex(i, j)

	testFunctionE := Multiple(f, g)
	testFunctionF := MultipleComplex(i, j)

	testFunctionG := Divide(f, g)
	testFunctionH := DivideComplex(i, j)

	testFunctionI := Parens(f)
	testFunctionJ := ParensComplex(i)

	if testFunctionA(1, 2) != 3.0 {
		t.Errorf("Expected %v, received %v", 3, testFunctionA(1, 2))
	}

	if testFunctionB(1, 2) != complex128(3) {
		t.Errorf("Expected %v, received %v", complex128(3), testFunctionA(1, 2))
	}

	if testFunctionC(2, 1) != 1.0 {
		t.Errorf("Expected %v, received %v", 1, testFunctionC(2, 1))
	}

	if testFunctionD(2, 1) != complex128(1) {
		t.Errorf("Expected %v, received %v", complex128(1), testFunctionD(2, 1))
	}

	if testFunctionE(1, 2) != 2.0 {
		t.Errorf("Expected %v, received %v", 2, testFunctionE(1, 2))
	}

	if testFunctionF(1, 2) != complex128(2) {
		t.Errorf("Expected %v, received %v", complex128(2), testFunctionF(1, 2))
	}

	if testFunctionG(2, 1) != 2.0 {
		t.Errorf("Expected %v, received %v", 2, testFunctionG(2, 1))
	}

	if testFunctionH(2, 1) != complex128(2) {
		t.Errorf("Expected %v, received %v", complex128(1), testFunctionH(2, 1))
	}

	if testFunctionI(2) != 2.0 {
		t.Errorf("Expected %v, received %v", 2, testFunctionG(2, 1))
	}

	if testFunctionJ(2) != complex128(2) {
		t.Errorf("Expected %v, received %v", complex128(1), testFunctionH(2, 1))
	}
}

func TestTriFunctions(t *testing.T) {
	f := func(x ...float64) float64 {
		return x[0]
	}

	g := func(x ...complex128) complex128 {
		return x[0]
	}

	testFunctionA := Sine(f)
	testFunctionB := SineComplex(g)

	testFunctionC := Cosine(f)
	testFunctionD := CosineComplex(g)

	testFunctionE := Tangent(f)
	testFunctionF := TangentComplex(g)

	testFunctionG := Arcsine(f)
	testFunctionH := ArcsineComplex(g)

	testFunctionI := Arccosine(f)
	testFunctionJ := ArccosineComplex(g)

	testFunctionK := Arctangent(f)
	testFunctionL := ArctangentComplex(g)

	if math.Abs(testFunctionA(math.Pi)-0.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionA(math.Pi)-0.0), 0.0001)
	}

	if cmplx.Abs(testFunctionB(math.Pi)-complex128(0)) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionB(math.Pi)-complex128(0)), 0.0001)
	}

	if testFunctionC(math.Pi) != -1.0 {
		t.Errorf("Expected %v, received %v", -1, testFunctionC(math.Pi))
	}

	if testFunctionD(math.Pi) != complex128(-1) {
		t.Errorf("Expected %v, received %v", complex128(-1), testFunctionD(math.Pi))
	}

	if math.Abs(testFunctionE(math.Pi)-0.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionE(math.Pi)-0.0), 0.0001)
	}

	if cmplx.Abs(testFunctionF(math.Pi)-complex128(0)) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionF(math.Pi)-complex128(0)), 0.0001)
	}

	if testFunctionG(0) != 0.0 {
		t.Errorf("Expected %v, received %v", 0, testFunctionG(0))
	}

	if testFunctionH(0) != complex128(0) {
		t.Errorf("Expected %v, received %v", complex128(0), testFunctionH(0))
	}

	if testFunctionI(0) != math.Pi/2.0 {
		t.Errorf("Expected %v, received %v", math.Pi/2.0, testFunctionG(0))
	}

	if testFunctionJ(0) != complex128(math.Pi/2.0) {
		t.Errorf("Expected %v, received %v", complex128(math.Pi/2.0), testFunctionH(0))
	}

	if testFunctionK(0) != 0.0 {
		t.Errorf("Expected %v, received %v", 0, testFunctionK(0))
	}

	if testFunctionL(0) != complex128(0) {
		t.Errorf("Expected %v, received %v", complex128(0), testFunctionL(0))
	}

	// Special functions
	h := func(x ...float64) float64 {
		return x[1]
	}

	testFunctionM := Arctangent2(f, h)
	testFunctionN := CotComplex(g)

	if math.Abs(testFunctionM(math.Pi, 2.0)-1.00388) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionM(math.Pi, 2.0)-0.0), 0.0001)
	}

	if cmplx.Abs(testFunctionN(math.Pi/2.0)-complex128(0)) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionF(math.Pi/2.0)-complex128(0)), 0.0001)
	}
}

func TestHyperTriFunctions(t *testing.T) {
	f := func(x ...float64) float64 {
		return x[0]
	}

	g := func(x ...complex128) complex128 {
		return x[0]
	}

	testFunctionA := HyperbolicSine(f)
	testFunctionB := HyperbolicSineComplex(g)

	testFunctionC := HyperbolicCosine(f)
	testFunctionD := HyperbolicCosineComplex(g)

	testFunctionE := HyperbolicTangent(f)
	testFunctionF := HyperbolicTangentComplex(g)

	testFunctionG := InverseHyperbolicSine(f)
	testFunctionH := InverseHyperbolicSineComplex(g)

	testFunctionI := InverseHyperbolicCosine(f)
	testFunctionJ := InverseHyperbolicCosineComplex(g)

	testFunctionK := InverseHyperbolicTangent(f)
	testFunctionL := InverseHyperbolicTangentComplex(g)

	if math.Abs(testFunctionA(1)-1.1752011936) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionA(1)-1.1752011936), 0.0001)
	}

	if cmplx.Abs(testFunctionB(1)-1.1752011936) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionB(1)-1.1752011936), 0.0001)
	}

	if math.Abs(testFunctionC(1)-1.5430806348) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionC(1)-1.5430806348), 0.0001)
	}

	if cmplx.Abs(testFunctionD(1)-1.5430806348) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionD(1)-1.5430806348), 0.0001)
	}

	if math.Abs(testFunctionE(1)-0.7615941559) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionE(1)-0.7615941559), 0.0001)
	}

	if cmplx.Abs(testFunctionF(1)-0.7615941559) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionF(1)-0.7615941559), 0.0001)
	}

	if math.Abs(testFunctionG(1)-0.881373587) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionG(1)-0.881373587), 0.0001)
	}

	if cmplx.Abs(testFunctionH(1)-0.881373587) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionH(1)-0.881373587), 0.0001)
	}

	if math.Abs(testFunctionI(1)-0.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionI(1)-0.0), 0.0001)
	}

	if cmplx.Abs(testFunctionJ(1)-0.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionJ(1)-0.0), 0.0001)
	}

	if math.Abs(testFunctionK(1.0/2.0)-0.549306144) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionK(1.0/2.0)-0.549306144), 0.0001)
	}

	if cmplx.Abs(testFunctionL(1.0/2.0)-0.549306144) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionL(1.0/2.0)-0.549306144), 0.0001)
	}
}

func TestSquareRootAbsoluteValueLogPower(t *testing.T) {
	f := func(x ...float64) float64 {
		return x[0]
	}

	g := func(x ...float64) float64 {
		return x[1]
	}

	i := func(x ...complex128) complex128 {
		return x[0]
	}

	j := func(x ...complex128) complex128 {
		return x[1]
	}

	testFunctionA := SquareRoot(f)
	testFunctionB := SquareRootComplex(i)

	testFunctionC := AbsoluteValue(f)
	testFunctionD := AbsoluteValueComplex(i)

	testFunctionE := LogBaseGx(f, g)
	testFunctionF := LogBaseGxComplex(i, j)

	testFunctionG := Power(f, g)
	testFunctionH := PowerComplex(i, j)

	if math.Abs(testFunctionA(4)-2.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionA(4)-2.0), 0.0001)
	}

	if cmplx.Abs(testFunctionB(4)-2.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionB(4)-complex128(2)), 0.0001)
	}

	if math.Abs(testFunctionC(-1)-1.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionC(-1)-1.0), 0.0001)
	}

	if cmplx.Abs(testFunctionD(-1)-1.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionD(-1)-1.0), 0.0001)
	}

	if math.Abs(testFunctionE(10, 10)-1.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionE(10, 10)-1.0), 0.0001)
	}

	if cmplx.Abs(testFunctionF(10, 10)-1.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionF(10, 10)-1.0), 0.0001)
	}

	if math.Abs(testFunctionG(2, 2)-4.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionG(2, 2)-4.0), 0.0001)
	}

	if cmplx.Abs(testFunctionH(2, 2)-4.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionH(2, 2)-4.0), 0.0001)
	}
}

func TestModuloMaxMinFloorCielErrorFunction(t *testing.T) {
	f := func(x ...float64) float64 {
		return x[0]
	}

	g := func(x ...float64) float64 {
		return x[1]
	}

	testFunctionA := Modulo(f, g)
	testFunctionB := Max(f, g)
	testFunctionC := Min(f, g)
	testFunctionD := Ceil(f)
	testFunctionE := Floor(f)
	testFunctionF := ErrorFunction(f)

	if math.Abs(testFunctionA(4, 3)-1.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionA(4, 3)-1.0), 0.0001)
	}

	if math.Abs(testFunctionB(11, 10)-11.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionB(11, 10)-11.0), 0.0001)
	}

	if math.Abs(testFunctionC(11, 10)-10.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionC(11, 10)-10.0), 0.0001)
	}

	if math.Abs(testFunctionD(10.6)-11.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionD(10.6)-11.0), 0.0001)
	}

	if math.Abs(testFunctionE(10.4)-10.0) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionE(10.4)-10.0), 0.0001)
	}

	if math.Abs(testFunctionF(1)-0.84270079) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", math.Abs(testFunctionF(1)-0.84270079), 0.0001)
	}
}

func TestArgumentConjugate(t *testing.T) {
	f := func(x ...complex128) complex128 {
		return x[0]
	}

	testFunctionA := ArgumentComplex(f)
	testFunctionB := ConjugateComplex(f)

	if cmplx.Abs(testFunctionA(1+2i)-1.107148717) >= 0.0001 {
		t.Errorf("Expected %v to be less than %v", cmplx.Abs(testFunctionA(1+2i)-1.107148717), 0.0001)
	}

	if testFunctionB(1+2i) != 1-2i {
		t.Errorf("Expected %v, received %v", 1-2i, testFunctionB(1+2i))
	}
}
