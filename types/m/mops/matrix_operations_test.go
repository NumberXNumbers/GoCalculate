package mops

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestSMult(t *testing.T) {
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

func TestVMMult(t *testing.T) {
	testVectorA := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1))
	testMatrix := m.MakeMatrix(testVectorA, testVectorA)

	resultVectorA, errA := VMMult(testVectorA, testMatrix)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA.Get(0), gcv.MakeValue(2.0)) ||
		!reflect.DeepEqual(resultVectorA.Get(1), gcv.MakeValue(2.0)) {
		t.Error("Failure: Test 1")
	}

	testVectorA.Trans()
	_, errB := VMMult(testVectorA, testMatrix)

	if errB == nil {
		t.Fail()
	}

	testVectorA.Trans()
	testVectorA.Append(gcv.NewValue())
	_, errC := VMMult(testVectorA, testMatrix)

	if errC == nil {
		t.Fail()
	}
}

func TestMustVMMult(t *testing.T) {
	testVectorA := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(1))
	testMatrix := m.MakeMatrix(testVectorA, testVectorA)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustVMMult(testVectorA, testMatrix)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestMVMult(t *testing.T) {
	testVectorA := v.MakeVector(v.ColSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testMatrix := m.MakeMatrix(testVectorA, testVectorA)

	resultMatrixA, errA := MVMult(testVectorA, testMatrix)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0), gcv.MakeValue(5.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1), gcv.MakeValue(5.0)) {
		t.Error("Failure: Test 1")
	}

	testVectorA.Trans()
	_, errB := MVMult(testVectorA, testMatrix)

	if errB == nil {
		t.Fail()
	}

	testVectorA.Trans()
	testVectorA.Append(gcv.NewValue())
	_, errC := MVMult(testVectorA, testMatrix)

	if errC == nil {
		t.Fail()
	}
}

func TestMustMVMult(t *testing.T) {
	testVectorA := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1))
	testMatrix := m.MakeMatrix(testVectorA, testVectorA)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustMVMult(testVectorA, testMatrix)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestSDiv(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(2+2i), gcv.MakeValue(2-4i))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAa)
	testMatrixB := m.MakeMatrix(testVectorAb, testVectorAb)

	testScalar := gcv.MakeValue(2.0)

	resultMatrixA := SDiv(testScalar, testMatrixA)
	resultMatrixB := SDiv(testScalar, testMatrixB)

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(0.5)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(1.0)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(0.5)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(1.0)) {
		t.Error("Failure: Test 1")
	}

	if !reflect.DeepEqual(resultMatrixB.Get(0, 0), gcv.MakeValue(1+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(0, 1), gcv.MakeValue(1-2i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 0), gcv.MakeValue(1+1i)) ||
		!reflect.DeepEqual(resultMatrixB.Get(1, 1), gcv.MakeValue(1-2i)) {
		t.Error("Failure: Test 2")
	}
}

func TestMultSimple(t *testing.T) {
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

func TestMustMultSimple(t *testing.T) {
	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustMultSimple(testMatrixC, testMatrixC)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestAdd(t *testing.T) {
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

func TestMustAdd(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustAdd(testMatrixA, testMatrixC)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestSub(t *testing.T) {
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

func TestMustSub(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustSub(testMatrixA, testMatrixC)

	if result != nil {
		t.Error("Expected Panic")
	}
}

func TestPow(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(1))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testMatrixA := m.MakeMatrix(testVectorAa, testVectorAb)

	resultMatrixA, errA := Pow(testMatrixA, 20)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultMatrixA.Get(0, 0), gcv.MakeValue(1743392201)) ||
		!reflect.DeepEqual(resultMatrixA.Get(0, 1), gcv.MakeValue(1743392200)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 0), gcv.MakeValue(1743392200)) ||
		!reflect.DeepEqual(resultMatrixA.Get(1, 1), gcv.MakeValue(1743392201)) {
		t.Error("Failure: Test2")
	}
}

func TestMustPow(t *testing.T) {
	testVectorCa := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(0), gcv.MakeValue(1))
	testVectorCb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(2))
	testMatrixC := m.MakeMatrix(testVectorCa, testVectorCb)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	result := MustPow(testMatrixC, 3)

	if result != nil {
		t.Error("Expected Panic")
	}
}
