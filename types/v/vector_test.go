package v

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

func TestGetAndSetMethodsVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeVector(4, RowVector)

	if !reflect.DeepEqual(testVectorA.Get(0), gcv.NewValue(0)) {
		t.Errorf("Expected %f, received %f", gcv.NewValue(0), testVectorA.Get(0))
	}

	testVectorB.Set(0, gcv.NewValue(2+2i))

	if !reflect.DeepEqual(testVectorB.Get(0), gcv.NewValue(2+2i)) {
		t.Errorf("Expected %f, received %f", gcv.NewValue(2+2i), testVectorB.Get(0))
	}
}

func TestTypeMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeVector(4, RowVector)

	if testVectorA.Type() != ColVector {
		t.Errorf("Expected %s, received %s", ColVector, testVectorA.Type())
	}

	if testVectorB.Type() != RowVector {
		t.Errorf("Expected %s, received %s", RowVector, testVectorB.Type())
	}
}

func TestCopyMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)

	if !reflect.DeepEqual(testVectorA.Copy(), testVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA.Copy(), testVectorA))
	}
}

func TestLenMethodVector(t *testing.T) {
	testVectorA := MakeVector(3, ColVector)
	testVectorB := MakeVector(4, RowVector)

	if testVectorA.Len() != 3 {
		t.Errorf("Expected %d, received %d", 3, testVectorA.Len())
	}

	if testVectorB.Len() != 4 {
		t.Errorf("Expected %d, received %d", 4, testVectorB.Len())
	}
}

func TestTransMethodVector(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)
	testTransVectorA := MakeVectorWithElements(testElementsA, RowVector)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1 + 1i), gcv.NewValue(2), gcv.NewValue(3 - 3i))
	testTransElementsB := gcv.NewValues(gcv.NewValue(1 - 1i), gcv.NewValue(2), gcv.NewValue(3 + 3i))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)
	testTransVectorB := MakeVectorWithElements(testTransElementsB, RowVector)

	testVectorB.Trans()

	if !reflect.DeepEqual(testVectorB.Get(0), testTransVectorB.Get(0)) &&
	   !reflect.DeepEqual(testVectorB.Get(1), testTransVectorB.Get(1)) &&
	   !reflect.DeepEqual(testVectorB.Get(2), testTransVectorB.Get(2)) {
		t.Errorf("Expected %v, received %v", testTransVectorB, testVectorB)
	}

	testVectorC := MakeVectorWithElements(testElementsA, RowVector)
	testTransVectorC := MakeVectorWithElements(testElementsA, ColVector)

	testVectorC.Trans()

	if !reflect.DeepEqual(testVectorC, testTransVectorC) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorC, testTransVectorC))
	}
}

func TestNormMethodVector(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	if !reflect.DeepEqual(testVectorA.Norm(), gcv.NewValue(5.0)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(5), testVectorA.Norm())
	}

	testElementsB := gcv.NewValues(gcv.NewValue(2 + 2i), gcv.NewValue(1))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	if !reflect.DeepEqual(testVectorB.Norm(), gcv.NewValue(3+0i)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(3), testVectorB.Norm())
	}
}

func TestMakeNewConjVector(t *testing.T) {
	testVectorA := MakeVector(4, RowVector)
	testTransVectorA := MakeNewConjVector(testVectorA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}
}
