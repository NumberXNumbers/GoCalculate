package v

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

func TestGetAndSetMethodsVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)
	testVectorB := NewVector("blank", 4)

	if !reflect.DeepEqual(testVectorA.Get(0), gcv.NewValue(0)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(0), testVectorA.Get(0))
	}

	testVectorB.Set(0, gcv.NewValue(2+2i))

	if !reflect.DeepEqual(testVectorB.Get(0), gcv.NewValue(2+2i)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(2+2i), testVectorB.Get(0))
	}
}

func TestSetValuesVectors(t *testing.T) {
	testVector := NewVector(ColSpace, 3)
	testVectors := MakeVectors(ColSpace, testVector, testVector, testVector)
	testValue := gcv.NewValue(3.0 + 1i)

	testVectors.SetValue(0, 1, testValue)

	if !reflect.DeepEqual(testVectors.Get(0).Get(1), testValue) {
		t.Fail()
	}
}

func TestSpaceMethodVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)
	testVectorB := NewVector(RowSpace, 4)

	if testVectorA.Space() != ColSpace {
		t.Errorf("Expected %s, received %s", ColSpace, testVectorA.Space())
	}

	if testVectorB.Space() != RowSpace {
		t.Errorf("Expected %s, received %s", RowSpace, testVectorB.Space())
	}
}

func TestCopyMethodVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)

	if !reflect.DeepEqual(testVectorA.Copy(), testVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA.Copy(), testVectorA))
	}
}

func TestLenMethodVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)
	testVectorB := NewVector(RowSpace, 4)

	if testVectorA.Len() != 3 {
		t.Errorf("Expected %d, received %d", 3, testVectorA.Len())
	}

	if testVectorB.Len() != 4 {
		t.Errorf("Expected %d, received %d", 4, testVectorB.Len())
	}
}

func TestTransMethodVector(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2), gcv.NewValue(3))
	testVectorA := MakeVector(ColSpace, testElementsA)
	testTransVectorA := MakeVector("blank", testElementsA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", testVectorA, testTransVectorA)
	}

	testElementsB := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2), gcv.NewValue(3-3i))
	testTransElementsB := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2), gcv.NewValue(3+3i))
	testVectorB := MakeVector(ColSpace, testElementsB)
	testTransVectorB := MakeVector(RowSpace, testTransElementsB)

	testVectorB.Trans()

	if !reflect.DeepEqual(testVectorB.Get(0), testTransVectorB.Get(0)) &&
		!reflect.DeepEqual(testVectorB.Get(1), testTransVectorB.Get(1)) &&
		!reflect.DeepEqual(testVectorB.Get(2), testTransVectorB.Get(2)) {
		t.Errorf("Expected %v, received %v", testTransVectorB, testVectorB)
	}

	testVectorC := MakeVector(RowSpace, testElementsA)
	testTransVectorC := MakeVector(ColSpace, testElementsA)

	testVectorC.Trans()

	if !reflect.DeepEqual(testVectorC, testTransVectorC) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorC, testTransVectorC))
	}
}

func TestNormMethodVector(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorA := MakeVector(ColSpace, testElementsA)

	if !reflect.DeepEqual(testVectorA.Norm(), gcv.NewValue(5.0)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(5), testVectorA.Norm())
	}

	testElementsB := gcv.NewValues(gcv.NewValue(2+2i), gcv.NewValue(1))
	testVectorB := MakeVector(ColSpace, testElementsB)

	if !reflect.DeepEqual(testVectorB.Norm(), gcv.NewValue(3+0i)) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(3), testVectorB.Norm())
	}
}

func TestMakeNewConjVector(t *testing.T) {
	testVectorA := NewVector(RowSpace, 4)
	testTransVectorA := MakeConjVector(testVectorA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}
}

func TestNewVectorsGet(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorA := MakeVector(ColSpace, testElementsA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVector(ColSpace, testElementsB)

	testVectorsA := MakeVectors("blank", testVectorA, testVectorB)

	if !reflect.DeepEqual(testVectorA, testVectorsA.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectorA, testVectorsA.Get(0))
	}

	if !reflect.DeepEqual(testVectorB, testVectorsA.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectorB, testVectorsA.Get(1))
	}

	testElementsC := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4))
	testVectorC := MakeVector(RowSpace, testElementsC)

	testElementsD := gcv.NewValues(gcv.NewValue(6+3i), gcv.NewValue(7-4i))
	testVectorD := MakeVector(ColSpace, testElementsD)

	testVectorsB := MakeVectors(RowSpace, testVectorC, testVectorD)

	if !reflect.DeepEqual(testVectorC, testVectorsB.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectorA, testVectorsA)
	}

	if !reflect.DeepEqual(testVectorD, testVectorsB.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectorD, testVectorsB.Get(1))
	}
}

func TestVectorIndexOf(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0))
	testVectorA := MakeVector(ColSpace, testElementsA)

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
	testVectorA := MakeVector(ColSpace, testElementsA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVector(ColSpace, testElementsB)

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	indexOfTestVectorA := testVectorsA.IndexOf(testVectorA)
	if indexOfTestVectorA != 0 {
		t.Errorf("Expected %v, received %v", 0, indexOfTestVectorA)
	}

	indexOfTestVectorB := testVectorsA.IndexOf(testVectorB)
	if indexOfTestVectorB != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorB)
	}

	testElementsC := gcv.NewValues(gcv.NewValue(4), gcv.NewValue(5.0), gcv.NewValue(8+1i))
	testVectorC := MakeVector(ColSpace, testElementsC)

	indexOfTestVectorC := testVectorsA.IndexOf(testVectorC)
	if indexOfTestVectorC != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testElementsD := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(5.0), gcv.NewValue(8+1i))
	testVectorD := MakeVector(ColSpace, testElementsD)

	indexOfTestVectorD := testVectorsA.IndexOf(testVectorD)
	if indexOfTestVectorD != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testElementsE := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(8+1i))
	testVectorE := MakeVector(ColSpace, testElementsE)

	indexOfTestVectorE := testVectorsA.IndexOf(testVectorE)
	if indexOfTestVectorE != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorE)
	}
}

func TestSetAndSpaceVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5))
	testVectorA := MakeVector(ColSpace, testElementsA)

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7+1i))
	testVectorB := MakeVector(RowSpace, testElementsB)

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
	testVectorA := MakeVector(ColSpace, testElementsA)

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testCopyVector := testVectorsA.Copy()

	if !reflect.DeepEqual(testCopyVector, testVectorsA) {
		t.Fail()
	}
}

func TestAppendVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVector(ColSpace, testElementsA)

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVector(ColSpace, testElementsB)

	testVectorB.Append(gcv.NewValue(8))

	if !reflect.DeepEqual(testVectorB.Get(2), gcv.NewValue(8)) {
		t.Fail()
	}

	testVectorsA.Append(testVectorB)

	if !reflect.DeepEqual(testVectorsA.Get(1), testVectorB) {
		t.Fail()
	}
}

func TestSubsetAndLenVectors(t *testing.T) {
	testElementsA := gcv.NewValues(gcv.NewValue(3), gcv.NewValue(4.0), gcv.NewValue(5+1i))
	testVectorA := MakeVector(ColSpace, testElementsA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVector(ColSpace, testElementsB)

	testElementsC := gcv.NewValues(gcv.NewValue(42))
	testVectorC := MakeVector(ColSpace, testElementsC)

	testVectorsA := MakeVectors(ColSpace, testVectorA, testVectorB, testVectorC)

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
	testVectorA := MakeVector(ColSpace, testElementsA)

	testElementsB := gcv.NewValues(gcv.NewValue(6), gcv.NewValue(7))
	testVectorB := MakeVector(ColSpace, testElementsB)

	testVectorsA := MakeVectors(ColSpace, testVectorA, testVectorB)

	if !reflect.DeepEqual(testVectorsA.Vectors()[0], testVectorA) || !reflect.DeepEqual(testVectorsA.Vectors()[1], testVectorB) {
		t.Fail()
	}

	testVectorsB := NewVectors("blank", 3, 4)

	if testVectorsB.Space() != RowSpace {
		t.Error("Expected Type RowSpace")
	}

	if !reflect.DeepEqual(testVectorsB.Get(2).Get(3), gcv.NewValue(0)) {
		t.Fail()
	}

}
