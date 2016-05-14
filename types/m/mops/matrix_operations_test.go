package mops

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestMatrixScalarMulti(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testMatrix := m.MakeMatrix(testVectorAa, testVectorAb)

	testScalarA := gcv.MakeValue(2.0)
	testScalarB := gcv.MakeValue(2 + 1i)

	resultMatrixA := SMult(testScalarA, testMatrix)
	resultMatrixB := SMult(testScalarB, testMatrix)

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(2.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(2.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.MakeValue(2+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.MakeValue(4+2i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.MakeValue(2+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.MakeValue(4+2i)) {
		t.Error("Failure: Test 2")
	}
}

func TestMatrixMultSimple(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	resultMatrixA, errA := MultSimple(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(2+0i), gcv.MakeValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(0+0i), gcv.MakeValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

	resultMatrixB, errB := MultSimple(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.MakeValue(4+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.MakeValue(4+0i)) {
		t.Error("Failure: Test2")
	}

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := MultSimple(testMatrixC, testMatrixC)

	if errC == nil {
		t.Fail()
	}

	testVectorDa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(0))
	testVectorDb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(1))
	testMatrixD := m.MakeMatrix(testVectorDa, testVectorDb)

	resultMatrixD, errD := MultSimple(testMatrixD, testMatrixD)

	if errD != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixD.Get(0, 0), gcv.MakeValue(1)) ||
		!reflect.DeepEqual(resultMatrixD.Get(0, 1), gcv.MakeValue(0)) ||
		!reflect.DeepEqual(resultMatrixD.Get(1, 0), gcv.MakeValue(0)) ||
		!reflect.DeepEqual(resultMatrixD.Get(1, 1), gcv.MakeValue(1)) {
		t.Error("Failure: Test 1")
	}

	resultMatrixE, errE := MultSimple(testMatrixA, testMatrixD)

	if errE != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixE.Get(0, 0), gcv.MakeValue(2)) ||
		!reflect.DeepEqual(resultMatrixE.Get(0, 1), gcv.MakeValue(0)) ||
		!reflect.DeepEqual(resultMatrixE.Get(1, 0), gcv.MakeValue(0)) ||
		!reflect.DeepEqual(resultMatrixE.Get(1, 1), gcv.MakeValue(2)) {
		t.Error("Failure: Test 1")
	}
}

func TestMatrixAddition(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	resultMatrixA, errA := Add(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(4.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(4.0)) {
		t.Error("Failure: Test 1")
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(2+0i), gcv.MakeValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(0+0i), gcv.MakeValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

	resultMatrixB, errB := Add(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.MakeValue(4+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.MakeValue(4+0i)) {
		t.Error("Failure: Test2")
	}

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := Add(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}

func TestMatrixSubtraction(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	resultMatrixA, errA := Sub(testMatrixA, testMatrixA)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(0.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(0.0)) {
		t.Error("Failure: Test 1")
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(2+0i), gcv.MakeValue(0+0i))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(0+0i), gcv.MakeValue(2+0i))
	testMatrixB := m.MakeMatrix(testVectorBa, testVectorBb)

	resultMatrixB, errB := Sub(testMatrixB, testMatrixB)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.MakeValue(0+0i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.MakeValue(0+0i)) {
		t.Error("Failure: Test2")
	}

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	_, errC := Sub(testMatrixA, testMatrixC)

	if errC == nil {
		t.Fail()
	}
}
