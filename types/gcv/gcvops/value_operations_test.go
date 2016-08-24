package gcvops

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

var (
	testValueA = gcv.MakeValue(1)
	testValueB = gcv.MakeValue(2)
	testValueC = gcv.MakeValue(1 + 1i)
	testValueD = gcv.MakeValue(1 - 1i)
	result     gcv.Value
	solution   gcv.Value
)

func TestAdd(t *testing.T) {
	result = Add(testValueA, testValueB)
	solution = gcv.MakeValue(3)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Add(testValueC, testValueD)
	solution = gcv.MakeValue(2)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestSub(t *testing.T) {
	result = Sub(testValueA, testValueB)
	solution = gcv.MakeValue(-1)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Sub(testValueC, testValueD)
	solution = gcv.MakeValue(2i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestDiv(t *testing.T) {
	result = Div(testValueA, testValueB)
	solution = gcv.MakeValue(0.5)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Div(testValueC, testValueD)
	solution = gcv.MakeValue(1i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestMult(t *testing.T) {
	result = Mult(testValueA, testValueB)
	solution = gcv.MakeValue(2)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Mult(testValueC, testValueD)
	solution = gcv.MakeValue(2)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestSqrt(t *testing.T) {
	result = Sqrt(testValueB)
	solution = gcv.MakeValue(1.4142135623730951)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Sqrt(testValueC)
	solution = gcv.MakeValue(1.09868411346781 + 0.45508986056222733i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAbs(t *testing.T) {
	result = Abs(testValueB)
	solution = gcv.MakeValue(2)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Abs(testValueC)
	solution = gcv.MakeValue(1.4142135623730951)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestConj(t *testing.T) {
	result = Conj(testValueB)
	solution = gcv.MakeValue(2)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Conj(testValueC)
	solution = testValueD
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestCot(t *testing.T) {
	result = Cot(testValueB)
	solution = gcv.MakeValue(-0.45765755436028577)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Cot(testValueC)
	solution = gcv.MakeValue(0.21762156185440268 - 0.868014142895925i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestSin(t *testing.T) {
	result = Sin(testValueA)
	solution = gcv.MakeValue(0.8414709848078965)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Sin(testValueC)
	solution = gcv.MakeValue(1.2984575814159773 + 0.6349639147847361i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestCos(t *testing.T) {
	result = Cos(testValueA)
	solution = gcv.MakeValue(0.5403023058681398)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Cos(testValueC)
	solution = gcv.MakeValue(0.8337300251311491 - 0.9888977057628651i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestTan(t *testing.T) {
	result = Tan(testValueA)
	solution = gcv.MakeValue(1.557407724654902)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Tan(testValueC)
	solution = gcv.MakeValue(0.2717525853195117 + 1.0839233273386948i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAsin(t *testing.T) {
	result = Asin(testValueA)
	solution = gcv.MakeValue(1.5707963267948966)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Asin(testValueC)
	solution = gcv.MakeValue(0.6662394324925153 + 1.0612750619050355i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAcos(t *testing.T) {
	result = Acos(testValueA)
	solution = gcv.MakeValue(0)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Acos(testValueC)
	solution = gcv.MakeValue(0.9045568943023813 - 1.0612750619050355i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAtan(t *testing.T) {
	result = Atan(testValueA)
	solution = gcv.MakeValue(0.7853981633974483)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Atan(testValueC)
	solution = gcv.MakeValue(1.0172219678978514 + 0.40235947810852507i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAsinh(t *testing.T) {
	result = Asinh(testValueA)
	solution = gcv.MakeValue(0.8813735870195432)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Asinh(testValueC)
	solution = gcv.MakeValue(1.0612750619050357 + 0.6662394324925153i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAcosh(t *testing.T) {
	result = Acosh(testValueA)
	solution = gcv.MakeValue(0)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Acosh(testValueC)
	solution = gcv.MakeValue(1.0612750619050355 + 0.9045568943023813i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestAtanh(t *testing.T) {
	result = Atanh(testValueA)
	solution = gcv.MakeValue(math.Inf(1))
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Atanh(testValueC)
	solution = gcv.MakeValue(0.40235947810852507 + 1.0172219678978514i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestSinh(t *testing.T) {
	result = Sinh(testValueA)
	solution = gcv.MakeValue(1.1752011936438014)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Sinh(testValueC)
	solution = gcv.MakeValue(0.6349639147847361 + 1.2984575814159773i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestCosh(t *testing.T) {
	result = Cosh(testValueA)
	solution = gcv.MakeValue(1.5430806348152437)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Cosh(testValueC)
	solution = gcv.MakeValue(0.8337300251311491 + 0.9888977057628651i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestTanh(t *testing.T) {
	result = Tanh(testValueA)
	solution = gcv.MakeValue(0.7615941559557649)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Tanh(testValueC)
	solution = gcv.MakeValue(1.0839233273386948 + 0.2717525853195117i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestLog(t *testing.T) {
	result = Log(testValueA)
	solution = gcv.MakeValue(0)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Log(testValueC)
	solution = gcv.MakeValue(0.3465735902799727 + 0.7853981633974483i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestLog10(t *testing.T) {
	result = Log10(testValueA)
	solution = gcv.MakeValue(0)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Log10(testValueC)
	solution = gcv.MakeValue(0.1505149978319906 + 0.3410940884604603i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestLogBase(t *testing.T) {
	result = LogBase(testValueA, testValueB)
	solution = gcv.MakeValue(0)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = LogBase(testValueC, testValueD)
	solution = gcv.MakeValue(-0.6740320278365401 + 0.738702122272951i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestPow(t *testing.T) {
	result = Pow(testValueA, testValueB)
	solution = gcv.MakeValue(1)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}

	result = Pow(testValueC, testValueD)
	solution = gcv.MakeValue(2.807879297260629 + 1.3178651729011805i)
	if !reflect.DeepEqual(result, solution) {
		t.Errorf("Expected %v, received %v", solution, result)
	}
}

func TestMod(t *testing.T) {
	resultA, errA := Mod(testValueB, testValueA)
	solution = gcv.MakeValue(0)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Mod(testValueC, testValueD)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustMod(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustMod(testValueC, testValueD)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestMax(t *testing.T) {
	resultA, errA := Max(testValueB, testValueA)
	solution = gcv.MakeValue(2)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Max(testValueC, testValueD)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustMax(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustMax(testValueC, testValueD)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestMin(t *testing.T) {
	resultA, errA := Min(testValueB, testValueA)
	solution = gcv.MakeValue(1)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Min(testValueC, testValueD)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustMin(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustMin(testValueC, testValueD)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestCeil(t *testing.T) {
	resultA, errA := Ceil(testValueB)
	solution = gcv.MakeValue(2)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Ceil(testValueC)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustCeil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustCeil(testValueC)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestFloor(t *testing.T) {
	resultA, errA := Floor(testValueB)
	solution = gcv.MakeValue(2)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Floor(testValueC)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustFloor(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustFloor(testValueC)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestErf(t *testing.T) {
	resultA, errA := Erf(testValueB)
	solution = gcv.MakeValue(0.9953222650189527)
	if errA != nil {
		t.Errorf("Unexpected Error: %v", errA)
	}
	if !reflect.DeepEqual(resultA, solution) {
		t.Errorf("Expected %v, received %v", solution, resultA)
	}

	_, errB := Erf(testValueC)
	if errB == nil {
		t.Error("Expected An Error")
	}
}

func TestMustErf(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustErf(testValueC)

	if result != nil {
		t.Error("Expected Panic")
	}
}
