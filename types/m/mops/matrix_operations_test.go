package mops

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestMatrixScalarMulti(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.NewValue(1), gcv.NewValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.NewValue(1), gcv.NewValue(2))
	testMatrix := m.MakeMatrix(testVectorAa, testVectorAb)

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
	testVectorAa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

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

	testVectorBa := v.MakeVector(v.RowSpace, gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

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

	testVectorCa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := MultiplicationSimple(testMatrixC, testMatrixC)

	if errC == nil {
		t.Fail()
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.NewValue(1), gcv.NewValue(0))
	testVectorDb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(1))
	testMatrixD := m.MakeMatrix(testVectorDa, testVectorDb)

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
	testVectorAa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

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

	testVectorBa := v.MakeVector(v.RowSpace, gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

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

	testVectorCa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := Addition(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}

func TestMatrixSubtraction(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

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

	testVectorBa := v.MakeVector(v.RowSpace, gcv.NewValue(2+0i), gcv.NewValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.NewValue(0+0i), gcv.NewValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

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

	testVectorCa := v.MakeVector(v.RowSpace, gcv.NewValue(2), gcv.NewValue(0), gcv.NewValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.NewValue(0), gcv.NewValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := Subtraction(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}
