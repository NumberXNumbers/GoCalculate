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
	testVector := v.MakeVector(v.ColSpace, testElements)

	testScalar := gcv.NewValue(2.0 + 1i)

	resultVector := ScalarMultiplication(testScalar, testVector)

	if !reflect.DeepEqual(resultVector.Get(0), gcv.NewValue(2+1i)) ||
		!reflect.DeepEqual(resultVector.Get(1), gcv.NewValue(4+2i)) ||
		resultVector.Space() != v.ColSpace {
		t.Fail()
	}
}

func TestVectorAddition(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorA := v.MakeVector(v.ColSpace, testElementsA)

	resultVectorA, errA := Addition(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.NewValue(2.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.NewValue(4.0)) ||
		resultVectorA.Space() != v.ColSpace {
		t.Fail()
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorB := v.MakeVector(v.RowSpace, testElementsB)

	resultVectorB, errB := Addition(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.NewValue(2+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.NewValue(4+2i)) ||
		resultVectorB.Space() != v.RowSpace {
		t.Fail()
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorCa := v.MakeVector(v.RowSpace, testElementsCa)

	testElementsCb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorCb := v.MakeVector(v.ColSpace, testElementsCb)

	_, errC := Addition(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorDa := v.MakeVector(v.RowSpace, testElementsDa)

	testElementsDb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorDb := v.MakeVector(v.RowSpace, testElementsDb)

	_, errD := Addition(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestVectorSubtraction(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorA := v.MakeVector(v.ColSpace, testElementsA)

	resultVectorA, errA := Subtraction(testVectorA, testVectorA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.NewValue(0.0)) ||
		resultVectorA.Space() != v.ColSpace {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.NewValue(0.0), gcv.NewValue(0.0),
			v.ColSpace, resultVectorA.Get(0), resultVectorA.Get(1), resultVectorA.Type())
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testVectorB := v.MakeVector(v.RowSpace, testElementsB)

	resultVectorB, errB := Subtraction(testVectorB, testVectorB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB.Get(0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultVectorB.Get(1), gcv.NewValue(0+0i)) ||
		resultVectorB.Space() != v.RowSpace {
		t.Errorf("Expected %v, %v and %v, received %v, %v and %v", gcv.NewValue(0+0i), gcv.NewValue(0+0i),
			v.RowSpace, resultVectorB.Get(0), resultVectorB.Get(1), resultVectorB.Type())
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorCa := v.MakeVector(v.RowSpace, testElementsCa)

	testElementsCb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorCb := v.MakeVector(v.ColSpace, testElementsCb)

	_, errC := Subtraction(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorDa := v.MakeVector(v.RowSpace, testElementsDa)

	testElementsDb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorDb := v.MakeVector(v.RowSpace, testElementsDb)

	_, errD := Subtraction(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestInnerProduct(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsA)
	testVectorAb := v.MakeVector(v.ColSpace, testElementsA)

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
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.ColSpace, testElementsBb)

	solutionB := gcv.NewValue(7 + 0i)

	resultB, errB := InnerProduct(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVector(v.ColSpace, testElementsA)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsA)

	_, errC := InnerProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testElementsDb := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2-1i), gcv.NewValue(3-1i))
	testVectorDa := v.MakeVector(v.RowSpace, testElementsDa)
	testVectorDb := v.MakeVector(v.ColSpace, testElementsDb)

	_, errD := InnerProduct(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}

func TestAngleTheta(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.ColSpace, testElementsAb)

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
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.ColSpace, testElementsBb)

	solutionB := gcv.NewValue(complex128(math.Pi / complex128(2)))

	resultB, errB := AngleTheta(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultB, solutionB) {
		t.Errorf("Expected %v, received %v", solutionB, resultB)
	}

	testVectorCa := v.MakeVector(v.ColSpace, testElementsAa)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsAb)

	_, errC := AngleTheta(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2+1i))
	testElementsDb := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2-1i), gcv.NewValue(3-1i))
	testVectorDa := v.MakeVector(v.RowSpace, testElementsDa)
	testVectorDb := v.MakeVector(v.ColSpace, testElementsDb)

	_, errD := AngleTheta(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testVectorE := v.NewVector(v.ColSpace, 2)

	_, errE := AngleTheta(testVectorE, testVectorE)

	if errE == nil {
		t.Error("Expected error")
	}
}

func TestOuterProduct(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	testVectorAa := v.MakeVector(v.ColSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)

	solutionElementsAa := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	solutionElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(0))
	solutionVectorAa := v.MakeVector(v.RowSpace, solutionElementsAa)
	solutionVectorAb := v.MakeVector(v.RowSpace, solutionElementsAb)
	solutionVectorsA := v.MakeVectors(v.RowSpace, solutionVectorAa, solutionVectorAb)
	solutionMatrixA := m.MakeMatrix(solutionVectorsA)

	resultMatrixA, errA := OuterProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	// fmt.Println(solutionMatrixA.Get(0, 0), solutionMatrixA.Get(0, 1), solutionMatrixA.Get(1, 0), solutionMatrixA.Get(1, 1))

	if solutionMatrixA.Get(0, 0).Complex128() != resultMatrixA.Get(0, 0).Complex128() ||
		solutionMatrixA.Get(1, 1).Complex128() != resultMatrixA.Get(1, 1).Complex128() ||
		solutionMatrixA.Get(0, 1).Complex128() != resultMatrixA.Get(0, 1).Complex128() ||
		solutionMatrixA.Get(1, 0).Complex128() != resultMatrixA.Get(1, 0).Complex128() {
		t.Fail()
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0+0i), gcv.NewValue(1+0i))
	testVectorBa := v.MakeVector(v.ColSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)

	resultMatrixB, errB := OuterProduct(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if solutionMatrixA.Get(0, 0).Complex128() != resultMatrixB.Get(0, 0).Complex128() ||
		solutionMatrixA.Get(1, 1).Complex128() != resultMatrixB.Get(1, 1).Complex128() ||
		solutionMatrixA.Get(0, 1).Complex128() != resultMatrixB.Get(0, 1).Complex128() ||
		solutionMatrixA.Get(1, 0).Complex128() != resultMatrixB.Get(1, 0).Complex128() {
		t.Fail()
	}

	testVectorCa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsAb)

	_, errC := OuterProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}
}
