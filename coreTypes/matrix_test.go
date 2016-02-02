package coreTypes

import (
	"reflect"
	"testing"
)

func TestIsIdenityMethod(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeMatrix(4, 3)
	testMatrixC := MakeIdentityMatrix(3)

	if testMatrixA.IsIdentity() {
		t.Fail()
	}

	if testMatrixB.IsIdentity() {
		t.Fail()
	}

	if !testMatrixC.IsIdentity() {
		t.Fail()
	}

	testMatrixD := MakeComplexMatrix(3, 3)
	testMatrixE := MakeComplexMatrix(3, 3)
	testMatrixF := MakeIdentityComplexMatrix(3)

	if testMatrixD.IsIdentity() {
		t.Fail()
	}

	if testMatrixE.IsIdentity() {
		t.Fail()
	}

	if !testMatrixF.IsIdentity() {
		t.Fail()
	}
}

func TestNewMatrix(t *testing.T) {
	testMatrixA, _, _ := NewMatrix(3, 3, reflect.Float64)
	_, testMatrixB, _ := NewMatrix(3, 3, reflect.Complex128)
	_, _, testMatrixC := NewMatrix(3, 3, reflect.Int)

	if testMatrixA.Type() != reflect.Float64 {
		t.Fail()
	}

	if testMatrixB.Type() != reflect.Complex128 {
		t.Fail()
	}

	if testMatrixC == nil {
		t.Fail()
	}
}

func TestTransMethod(t *testing.T) {
	testElementsA := [][]float64{{1, 2}, {1, 2}}
	testElementsTransA := [][]float64{{1, 1}, {2, 2}}
	testMatrixA := MakeMatrixWithElements(testElementsA)
	testMatrixTransA := MakeMatrixWithElements(testElementsTransA)

	if !reflect.DeepEqual(testMatrixTransA, testMatrixA.Trans()) {
		t.Fail()
	}

	testElementsB := [][]float64{{1, 2}, {1, 2}, {1, 2}}
	testElementsTransB := [][]float64{{1, 1, 1}, {2, 2, 2}}
	testMatrixB := MakeMatrixWithElements(testElementsB)
	testMatrixTransB := MakeMatrixWithElements(testElementsTransB)

	if !reflect.DeepEqual(testMatrixTransB, testMatrixB.Trans()) {
		t.Fail()
	}

	testElementsC := [][]complex128{{1 + 1i, 2}, {1, 2 + 1i}}
	testElementsTransC := [][]complex128{{1 - 1i, 1}, {2, 2 - 1i}}
	testMatrixC := MakeComplexMatrixWithElements(testElementsC)
	testMatrixTransC := MakeComplexMatrixWithElements(testElementsTransC)

	if !reflect.DeepEqual(testMatrixTransC, testMatrixC.Trans()) {
		t.Fail()
	}

	testElementsD := [][]complex128{{1 + 1i, 2}, {1, 2 + 1i}, {1 - 1i, 2}}
	testElementsTransD := [][]complex128{{1 - 1i, 1, 1 + 1i}, {2, 2 - 1i, 2}}
	testMatrixD := MakeComplexMatrixWithElements(testElementsD)
	testMatrixTransD := MakeComplexMatrixWithElements(testElementsTransD)

	if !reflect.DeepEqual(testMatrixTransD, testMatrixD.Trans()) {
		t.Fail()
	}
}
