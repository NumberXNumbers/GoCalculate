package vops

import (
	"math"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestVectorScalarMulti(t *testing.T) {
	testElements := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVector := v.MakeVectorWithElements(testElements, ColVector)

	testScalar := gcv.NewValue(2.0 + 1i)

	resultVector := VectorScalarMulti(testScalar, testVector)

	if !reflect.DeepEqual(resultVector.Get(0), gcv.NewValue(2+1i)) ||
		!reflect.DeepEqual(resultVector.Get(1), gcv.NewValue(4+2i)) ||
		resultVector.Type() != ColVector {
		t.Fail()
	}
}

func TestVectorAddition(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorA := v.MakeVectorWithElements(testElementsA, ColVector)

	resultVectorA, errA := VectorAddition(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.NewValue(2.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.NewValue(4.0)) ||
		resultVectorA.Type() != ColVector {
		t.Fail()
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorB := v.MakeVectorWithElements(testElementsB, RowVector)

	resultVectorB, errB := VectorAddition(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.NewValue(2+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.NewValue(4+2i)) ||
		resultVectorB.Type() != RowVector {
		t.Fail()
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorCa := v.MakeVectorWithElements(testElementsCa, v.RowVector)

	testElementsCb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorCb := v.MakeVectorWithElements(testElementsCb, v.ColVector)

	_, errC := VectorAddition(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorDa := v.MakeVectorWithElements(testElementsDa, v.RowVector)

	testElementsDb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorDb := v.MakeVectorWithElements(testElementsDb, v.RowVector)

	_, errD := VectorAddition(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestVectorSubtraction(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorA := v.MakeVectorWithElements(testElementsA, ColVector)

	resultVectorA, errA := VectorSubtraction(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.NewValue(0.0)) ||
		resultVectorA.Type() != ColVector {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.NewValue(0.0), gcv.NewValue(0.0),
			ColVector, resultVectorA.Get(0), resultVectorA.Get(1), resultVectorA.Type())
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testVectorB := v.MakeVectorWithElements(testElementsB, RowVector)

	resultVectorB, errB := VectorSubtraction(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.NewValue(0+0i)) ||
		resultVectorB.Type() != RowVector {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.NewValue(0+0i), gcv.NewValue(0+0i),
			RowVector, resultVectorB.Get(0), resultVectorB.Get(1), resultVectorB.Type())
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorCa := v.MakeVectorWithElements(testElementsCa, v.RowVector)

	testElementsCb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorCb := v.MakeVectorWithElements(testElementsCb, v.ColVector)

	_, errC := VectorSubtraction(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorDa := v.MakeVectorWithElements(testElementsDa, v.RowVector)

	testElementsDb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorDb := v.MakeVectorWithElements(testElementsDb, v.RowVector)

	_, errD := VectorSubtraction(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestInnerProduct(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVectorWithElements(testElementsA, RowVector)
	testVectorAb := v.MakeVectorWithElements(testElementsA, ColVector)

	solutionA := gcv.NewValue(5.0)

	resultA, errA := InnerProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultA, solutionA) {
		t.Errorf("Expected %v, received %v", solutionA, resultA)
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testElementsBb := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2-1i))
	testVectorBa := v.MakeVectorWithElements(testElementsBa, RowVector)
	testVectorBb := v.MakeVectorWithElements(testElementsBb, ColVector)

	solutionB := gcv.NewValue(7 + 0i)

	resultB, errB := InnerProduct(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVectorWithElements(testElementsA, ColVector)
	testVectorCb := v.MakeVectorWithElements(testElementsA, RowVector)

	_, errC := InnerProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testElementsDb := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2-1i), gcv.NewValue(3-1i))
	testVectorDa := v.MakeVectorWithElements(testElementsDa, RowVector)
	testVectorDb := v.MakeVectorWithElements(testElementsDb, ColVector)

	_, errD := InnerProduct(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestAngleTheta(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	testVectorAa := v.MakeVectorWithElements(testElementsAa, RowVector)
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ColVector)

	solutionA := gcv.NewValue(float64(math.Pi / float64(2)))

	resultA, errA := AngleTheta(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultA, solutionA) {
		t.Errorf("Expected %v, received %v", solutionA, resultA)
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0-0i), gcv.NewValue(1-0i))
	testVectorBa := v.MakeVectorWithElements(testElementsBa, RowVector)
	testVectorBb := v.MakeVectorWithElements(testElementsBb, ColVector)

	solutionB := gcv.NewValue(complex128(math.Pi / complex128(2)))

	resultB, errB := AngleTheta(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVectorWithElements(testElementsAa, ColVector)
	testVectorCb := v.MakeVectorWithElements(testElementsAb, RowVector)

	_, errC := AngleTheta(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testElementsDb := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2-1i), gcv.NewValue(3-1i))
	testVectorDa := v.MakeVectorWithElements(testElementsDa, RowVector)
	testVectorDb := v.MakeVectorWithElements(testElementsDb, ColVector)

	_, errD := AngleTheta(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestOuterProduct(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ColVector)
	testVectorAb := v.MakeVectorWithElements(testElementsAb, RowVector)

	solutionElementsA := [][]float64{{0, 1}, {0, 0}}
	solutionMatrixA := m.MakeMatrixWithElements(solutionElementsA)

	resultMatrixA, errA := OuterProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0-0i), gcv.NewValue(1-0i))
	testVectorBa := v.MakeVectorWithElements(testElementsBa, ColVector)
	testVectorBb := v.MakeVectorWithElements(testElementsBb, RowVector)

	solutionElementsB := [][]complex128{{0, 1}, {0, 0}}
	solutionMatrixB := m.MakeComplexMatrixWithElements(solutionElementsB)

	resultMatrixB, errB := OuterProductComplex(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(solutionMatrixB, resultMatrixB))
	}

	testVectorCa := v.MakeVectorWithElements(testElementsAa, RowVector)
	testVectorCb := v.MakeVectorWithElements(testElementsAb, RowVector)

	_, errC := OuterProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testVectorDa := v.MakeVectorWithElements(testElementsBa, ColVector)
	testVectorDb := v.MakeVectorWithElements(testElementsBb, ColVector)

	_, errD := OuterProductComplex(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}
