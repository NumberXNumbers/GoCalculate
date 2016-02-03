package coreTypes

import (
	"reflect"
	"testing"
)

func TestGetAndSetMethodsVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeComplexVector(3, ColVector)
	testVectorC := MakeVector(4, RowVector)
	testVectorD := MakeComplexVector(4, RowVector)

	if testVectorA.Get(0) != float64(0) {
		t.Errorf("Expected %f, recieved %f", float64(0), testVectorA.Get(0))
	}

	if testVectorB.Get(0) != complex128(0) {
		t.Errorf("Expected %f, recieved %f", complex128(0), testVectorB.Get(0))
	}

	testVectorC.Set(0, 2.0)
	testVectorD.Set(0, 2+2i)

	if testVectorC.Get(0) != 2.0 {
		t.Errorf("Expected %f, recieved %f", 2.0, testVectorC.Get(0))
	}

	if testVectorD.Get(0) != 2+2i {
		t.Errorf("Expected %f, recieved %f", 2+2i, testVectorD.Get(0))
	}
}

func TestTypeMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeComplexVector(3, ColVector)
	testVectorC := MakeVector(4, RowVector)
	testVectorD := MakeComplexVector(4, RowVector)

	if testVectorA.Type() != ColVector {
		t.Errorf("Expected %s, recieved %s", ColVector, testVectorA.Type())
	}

	if testVectorB.Type() != ColVector {
		t.Errorf("Expected %s, recieved %s", ColVector, testVectorB.Type())
	}

	if testVectorC.Type() != RowVector {
		t.Errorf("Expected %s, recieved %s", RowVector, testVectorC.Type())
	}

	if testVectorD.Type() != RowVector {
		t.Errorf("Expected %s, recieved %s", RowVector, testVectorD.Type())
	}
}

func TestCopyMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeComplexVector(4, RowVector)

	if !reflect.DeepEqual(testVectorA.Copy(), testVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorA.Copy(), testVectorA))
	}

	if !reflect.DeepEqual(testVectorB.Copy(), testVectorB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorB.Copy(), testVectorB))
	}
}

func TestDimMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeComplexVector(4, RowVector)

	if testVectorA.Dim() != 3 {
		t.Errorf("Expected %d, recieved %d", 3, testVectorA.Dim())
	}

	if testVectorB.Dim() != 4 {
		t.Errorf("Expected %d, recieved %d", 4, testVectorB.Dim())
	}
}

func TestTransMethodVector(t *testing.T) {
	testElementsA := []float64{1, 2, 3}
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)
	testTransVectorA := MakeVectorWithElements(testElementsA, RowVector)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}

	testElementsB := []complex128{1 + 1i, 2, 3 - 3i}
	testTransElementsB := []complex128{1 - 1i, 2, 3 + 3i}
	testVectorB := MakeComplexVectorWithElements(testElementsB, ColVector)
	testTransVectorB := MakeComplexVectorWithElements(testTransElementsB, RowVector)

	testVectorB.Trans()

	if !reflect.DeepEqual(testVectorB, testTransVectorB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorB, testTransVectorB))
	}

	testElementsC := []float64{1, 2, 3}
	testVectorC := MakeVectorWithElements(testElementsC, RowVector)
	testTransVectorC := MakeVectorWithElements(testElementsC, ColVector)

	testVectorC.Trans()

	if !reflect.DeepEqual(testVectorC, testTransVectorC) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorC, testTransVectorC))
	}

	testElementsD := []complex128{1 + 1i, 2, 3 - 3i}
	testTransElementsD := []complex128{1 - 1i, 2, 3 + 3i}
	testVectorD := MakeComplexVectorWithElements(testElementsD, RowVector)
	testTransVectorD := MakeComplexVectorWithElements(testTransElementsD, ColVector)

	testVectorD.Trans()

	if !reflect.DeepEqual(testVectorD, testTransVectorD) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorD, testTransVectorD))
	}
}

func TestGetElementsMethodVector(t *testing.T) {
	testElementsA := []float64{3, 4}
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	if !reflect.DeepEqual(testVectorA.GetElements(), testElementsA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorA.GetElements(), testElementsA))
	}

	testElementsB := []complex128{2 + 2i, 1}
	testVectorB := MakeComplexVectorWithElements(testElementsB, ColVector)

	if !reflect.DeepEqual(testVectorB.GetElements(), testElementsB) {
		t.Errorf("Expected %v, recieved %v", true, !reflect.DeepEqual(testVectorB.GetElements(), testElementsB))
	}
}

func TestNormMethodVector(t *testing.T) {
	testElementsA := []float64{3, 4}
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	if testVectorA.Norm() != float64(5) {
		t.Errorf("Expected %f, recieved %f", float64(5), testVectorA.Norm())
	}

	testElementsB := []complex128{2 + 2i, 1}
	testVectorB := MakeComplexVectorWithElements(testElementsB, ColVector)

	if testVectorB.Norm() != complex128(3) {
		t.Errorf("Expected %f, recieved %f", complex128(3), testVectorB.Norm())
	}
}

func TestMakeNewConjVector(t *testing.T) {
	testVectorA := MakeComplexVector(4, RowVector)
	testTransVectorA := MakeNewConjVector(testVectorA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}
}
