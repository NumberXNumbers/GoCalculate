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
		t.Errorf("Expected %v, received %v", gcv.NewValue(0), testVectorA.Get(0))
	}

	testVectorB.Set(0, gcv.NewValue(2+2i))

	if !reflect.DeepEqual(testVectorB.Get(0), gcv.NewValue(2+2i)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(2+2i), testVectorB.Get(0))
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

	testElementsB := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2), gcv.NewValue(3-3i))
	testTransElementsB := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2), gcv.NewValue(3+3i))
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

	testElementsB := gcv.NewValues(gcv.NewValue(2+2i), gcv.NewValue(1))
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

func TestNewVectorsGet(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA, testVectorB)

	if !reflect.DeepEqual(testVectorA, testVectorsA.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectorA, testVectorsA.Get(0))
	}

	if !reflect.DeepEqual(testVectorB, testVectorsA.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectorB, testVectorsA.Get(1))
	}

	testElementsC := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorC := MakeVectorWithElements(testElementsC, RowVector)

	testElementsD := gcv.NewValues(gcv.NewValue(6+3i), gcv.NewValue(7-4i))
	testVectorD := MakeVectorWithElements(testElementsD, ColVector)

	testVectorsB := NewVectors(RowSpace, testVectorC, testVectorD)

	if !reflect.DeepEqual(testVectorC, testVectorsB.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectorA, testVectorsA)
	}

	if !reflect.DeepEqual(testVectorD, testVectorsB.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectorD, testVectorsB.Get(1))
	}
}

func TestVectorIndexOf(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testIndexA := testVectorA.IndexOf(gcv.NewValue(3))
	if testIndexA != 0 {
		t.Errorf("Expected %v, received %v", 0, testIndexA)
	}

	testIndexB := testVectorA.IndexOf(gcv.NewValue(5 + 1i))
	if testIndexB != -1 {
		t.Errorf("Expected %v, received %v", 0, testIndexB)
	}
}

func TestVectorsIndexOf(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA)

	indexOfTestVectorA := testVectorsA.IndexOf(testVectorA)
	if indexOfTestVectorA != 0 {
		t.Errorf("Expected %v, received %v", 0, indexOfTestVectorA)
	}

	indexOfTestVectorB := testVectorsA.IndexOf(testVectorB)
	if indexOfTestVectorB != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorB)
	}

	testElementsC := gcv.NewValues(gcv.NewValue(4), gcv.NewValue(5.0), gcv.NewValue(8+1i))
	testVectorC := MakeVectorWithElements(testElementsC, ColVector)

	indexOfTestVectorC := testVectorsA.IndexOf(testVectorC)
	if indexOfTestVectorC != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testElementsD := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(5.0), gcv.NewValue(8+1i))
	testVectorD := MakeVectorWithElements(testElementsD, ColVector)

	indexOfTestVectorD := testVectorsA.IndexOf(testVectorD)
	if indexOfTestVectorD != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testElementsE := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(8+1i))
	testVectorE := MakeVectorWithElements(testElementsE, ColVector)

	indexOfTestVectorE := testVectorsA.IndexOf(testVectorE)
	if indexOfTestVectorE != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorE)
	}
}

func TestSetAndSpaceVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, RowVector)

	testVectorsA.Set(0, testVectorB)

	if !reflect.DeepEqual(testVectorB, testVectorsA.Get(0)) {
		t.Fail()
	}

	if testVectorsA.Space() != ColSpace {
		t.Fail()
	}
}

func TestCopyVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA)

	testCopyVector := testVectorsA.Copy()

	if !reflect.DeepEqual(testCopyVector, testVectorsA) {
		t.Fail()
	}
}

func TestAppendVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	testVectorsA.Append(testVectorB)

	if !reflect.DeepEqual(testVectorsA.Get(1), testVectorB) {
		t.Fail()
	}
}

func TestSubsetAndLenVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	testElementsC := gcv.NewValues(gcv.NewValue(42))
	testVectorC := MakeVectorWithElements(testElementsC, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA, testVectorB, testVectorC)

	lenA := testVectorsA.Len()

	testVectorsB := testVectorsA.Subset(1, 2)

	lenB := testVectorsB.Len()

	if !reflect.DeepEqual(testVectorsB.Get(0), testVectorB) || !reflect.DeepEqual(testVectorsB.Get(1), testVectorC) {
		t.Fail()
	}

	if lenB != lenA-1 {
		t.Fail()
	}
}

func TestVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVectorWithElements(testElementsA, ColVector)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVectorWithElements(testElementsB, ColVector)

	testVectorsA := NewVectors(ColSpace, testVectorA, testVectorB)

	if !reflect.DeepEqual(testVectorsA.Vectors()[0], testVectorA) || !reflect.DeepEqual(testVectorsA.Vectors()[1], testVectorB) {
		t.Fail()
	}
}
