package mops

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestMatrixScalarMulti(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testElementsAb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectors := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrix := m.MakeMatrix(testVectors)

	testScalarA := gcv.NewValue(2.0)
	testScalarB := gcv.NewValue(2 + 1i)

	resultMatrixA := ScalarMultiplication(testScalarA, testMatrix)
	resultMatrixB := ScalarMultiplication(testScalarB, testMatrix)

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.NewValue(2.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.NewValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.NewValue(2.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.NewValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.NewValue(2+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.NewValue(4+2i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.NewValue(2+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.NewValue(4+2i)) {
		t.Error("Failure: Test 2")
	}
}

func TestMatrixMultiSimple(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := m.MakeMatrix(testVectorsA)

	resultMatrixA, errA := MultiplicationSimple(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.NewValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.NewValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := m.MakeMatrix(testVectorsB)

	resultMatrixB, errB := MultiplicationSimple(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.NewValue(4+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.NewValue(4+0i)) {
		t.Error("Failure: Test2")
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testElementsCb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorCa := v.MakeVector(v.RowSpace, testElementsCa)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsCb)
	testVectorsC := v.MakeVectors(v.RowSpace, testVectorCa, testVectorCb)
	testMatrixC := m.MakeMatrix(testVectorsC)

	_, errC := MultiplicationSimple(testMatrixC, testMatrixC)

	if errC == nil {
		t.Fail()
	}

	testElementsDa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(0))
	testElementsDb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1))
	testVectorDa := v.MakeVector(v.RowSpace, testElementsDa)
	testVectorDb := v.MakeVector(v.RowSpace, testElementsDb)
	testVectorsD := v.MakeVectors(v.RowSpace, testVectorDa, testVectorDb)
	testMatrixD := m.MakeMatrix(testVectorsD)

	resultMatrixD, errD := MultiplicationSimple(testMatrixD, testMatrixD)

	if errD != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixD.Get(0, 0), gcv.NewValue(1)) ||
		!reflect.DeepEqual(resultMatrixD.Get(0, 1), gcv.NewValue(0)) ||
		!reflect.DeepEqual(resultMatrixD.Get(1, 0), gcv.NewValue(0)) ||
		!reflect.DeepEqual(resultMatrixD.Get(1, 1), gcv.NewValue(1)) {
		t.Error("Failure: Test 1")
	}
}

func TestMatrixAddition(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := m.MakeMatrix(testVectorsA)

	resultMatrixA, errA := Addition(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.NewValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.NewValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := m.MakeMatrix(testVectorsB)

	resultMatrixB, errB := Addition(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.NewValue(4+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.NewValue(4+0i)) {
		t.Error("Failure: Test2")
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testElementsCb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorCa := v.MakeVector(v.RowSpace, testElementsCa)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsCb)
	testVectorsC := v.MakeVectors(v.RowSpace, testVectorCa, testVectorCb)
	testMatrixC := m.MakeMatrix(testVectorsC)

	_, errC := Addition(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}

func TestMatrixSubtraction(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := m.MakeMatrix(testVectorsA)

	resultMatrixA, errA := Subtraction(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.NewValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.NewValue(0.0)) {
		t.Error("Failure: Test 1")
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testElementsBb := gcv.NewValues(gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := m.MakeMatrix(testVectorsB)

	resultMatrixB, errB := Subtraction(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.NewValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.NewValue(0+0i)) {
		t.Error("Failure: Test2")
	}

	testElementsCa := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testElementsCb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(2))
	testVectorCa := v.MakeVector(v.RowSpace, testElementsCa)
	testVectorCb := v.MakeVector(v.RowSpace, testElementsCb)
	testVectorsC := v.MakeVectors(v.RowSpace, testVectorCa, testVectorCb)
	testMatrixC := m.MakeMatrix(testVectorsC)

	_, errC := Subtraction(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}
