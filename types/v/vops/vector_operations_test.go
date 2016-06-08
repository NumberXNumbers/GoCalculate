package vops

import (
	"math"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestSMult(t *testing.T) {
	testVector := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2))

	testScalar := gcv.MakeValue(2.0 + 1i)

	resultVector := SMult(testScalar, testVector)

	if !reflect.DeepEqual(resultVector.Get(0), gcv.MakeValue(2+1i)) ||
		!reflect.DeepEqual(resultVector.Get(1), gcv.MakeValue(4+2i)) ||
		resultVector.Space() != v.ColSpace {
		t.Fail()
	}
}

func TestSDiv(t *testing.T) {
	testVector := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2))

	testScalar := gcv.MakeValue(2.0)

	resultVector := SDiv(testScalar, testVector)

	if !reflect.DeepEqual(resultVector.Get(0), gcv.MakeValue(0.5)) ||
		!reflect.DeepEqual(resultVector.Get(1), gcv.MakeValue(1)) ||
		resultVector.Space() != v.ColSpace {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	testVectorA := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2))

	resultVectorA, errA := Add(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.MakeValue(2.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.MakeValue(4.0)) ||
		resultVectorA.Space() != v.ColSpace {
		t.Fail()
	}

	testVectorB := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))

	resultVectorB, errB := Add(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.MakeValue(2+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.MakeValue(4+2i)) ||
		resultVectorB.Space() != v.RowSpace {
		t.Fail()
	}

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))

	testVectorCb := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))

	_, errC := Add(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2), gcv.MakeValue(3))
	testVectorDb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))

	_, errD := Add(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestSub(t *testing.T) {
	testVectorA := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2))

	resultVectorA, errA := Sub(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.MakeValue(0.0)) ||
		resultVectorA.Space() != v.ColSpace {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.MakeValue(0.0), gcv.MakeValue(0.0),
			v.ColSpace, resultVectorA.Get(0), resultVectorA.Get(1), resultVectorA.Type())
	}

	testVectorB := v.MakeVector(v.RowSpace, gcv.MakeValue(1+1i), gcv.MakeValue(2+1i))

	resultVectorB, errB := Sub(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.MakeValue(0+0i)) ||
		resultVectorB.Space() != v.RowSpace {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.MakeValue(0+0i), gcv.MakeValue(0+0i),
			v.RowSpace, resultVectorB.Get(0), resultVectorB.Get(1), resultVectorB.Type())
	}

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorCb := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))

	_, errC := Sub(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2), gcv.MakeValue(3))
	testVectorDb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))

	_, errD := Sub(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestInnerProduct(t *testing.T) {
	testElementsA := gcv.MakeValues(gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAa := v.MakeVectorAlt(v.RowSpace, testElementsA)
	testVectorAb := v.MakeVectorAlt(v.ColSpace, testElementsA)

	solutionA := gcv.MakeValue(5.0)

	resultA, errA := InnerProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultA, solutionA) {
		t.Errorf("Expected %v, received %v", solutionA, resultA)
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(1+1i), gcv.MakeValue(2+1i))
	testVectorBb := v.MakeVector(v.ColSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2-1i))

	solutionB := gcv.MakeValue(7 + 0i)

	resultB, errB := InnerProduct(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVectorAlt(v.ColSpace, testElementsA)
	testVectorCb := v.MakeVectorAlt(v.RowSpace, testElementsA)

	_, errC := InnerProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.MakeValue(1+1i), gcv.MakeValue(2+1i))
	testVectorDb := v.MakeVector(v.ColSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2-1i), gcv.MakeValue(3-1i))

	_, errD := InnerProduct(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestAcos(t *testing.T) {
	testElementsAa := gcv.MakeValues(gcv.MakeValue(1), gcv.MakeValue(0))
	testElementsAb := gcv.MakeValues(gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorAa := v.MakeVectorAlt(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVectorAlt(v.ColSpace, testElementsAb)

	solutionA := gcv.MakeValue(float64(math.Pi / float64(2)))

	resultA, errA := Acos(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultA, solutionA) {
		t.Errorf("Expected %v, received %v", solutionA, resultA)
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(1+0i), gcv.MakeValue(0+0i))
	testVectorBb := v.MakeVector(v.ColSpace, gcv.MakeValue(0-0i), gcv.MakeValue(1-0i))

	solutionB := gcv.MakeValue(complex128(math.Pi / complex128(2)))

	resultB, errB := Acos(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVectorAlt(v.ColSpace, testElementsAa)
	testVectorCb := v.MakeVectorAlt(v.RowSpace, testElementsAb)

	_, errC := Acos(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.MakeValue(1+1i), gcv.MakeValue(2+1i))
	testVectorDb := v.MakeVector(v.ColSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2-1i), gcv.MakeValue(3-1i))

	_, errD := Acos(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testVectorE := v.NewVector(v.ColSpace, 2)

	_, errE := Acos(testVectorE, testVectorE)

	if errE == nil {
		t.Error("Expected error")
	}
}

func TestOuterProduct(t *testing.T) {
	testElementsAa := gcv.MakeValues(gcv.MakeValue(1), gcv.MakeValue(0))
	testElementsAb := gcv.MakeValues(gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorAa := v.MakeVectorAlt(v.ColSpace, testElementsAa)
	testVectorAb := v.MakeVectorAlt(v.RowSpace, testElementsAb)

	solutionVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(1))
	solutionVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(0))
	solutionMatrixA := m.MakeMatrix(solutionVectorAa, solutionVectorAb)

	resultMatrixA, errA := OuterProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	// fmt.Println(solutionMatrixA.Get(0, 0), solutionMatrixA.Get(0, 1), solutionMatrixA.Get(1, 0), solutionMatrixA.Get(1, 1))

	if solutionMatrixA.Get(0, 0).Complex() != resultMatrixA.Get(0, 0).Complex() ||
		solutionMatrixA.Get(1, 1).Complex() != resultMatrixA.Get(1, 1).Complex() ||
		solutionMatrixA.Get(0, 1).Complex() != resultMatrixA.Get(0, 1).Complex() ||
		solutionMatrixA.Get(1, 0).Complex() != resultMatrixA.Get(1, 0).Complex() {
		t.Fail()
	}

	testVectorBa := v.MakeVector(v.ColSpace, gcv.MakeValue(1+0i), gcv.MakeValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(0+0i), gcv.MakeValue(1+0i))

	resultMatrixB, errB := OuterProduct(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if solutionMatrixA.Get(0, 0).Complex() != resultMatrixB.Get(0, 0).Complex() ||
		solutionMatrixA.Get(1, 1).Complex() != resultMatrixB.Get(1, 1).Complex() ||
		solutionMatrixA.Get(0, 1).Complex() != resultMatrixB.Get(0, 1).Complex() ||
		solutionMatrixA.Get(1, 0).Complex() != resultMatrixB.Get(1, 0).Complex() {
		t.Fail()
	}

	testVectorCa := v.MakeVectorAlt(v.RowSpace, testElementsAa)
	testVectorCb := v.MakeVectorAlt(v.RowSpace, testElementsAb)

	_, errC := OuterProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}
}
