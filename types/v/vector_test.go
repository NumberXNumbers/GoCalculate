package v

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
)

func TestGetAndSetMethodsVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)
	testVectorB := NewVector(RowSpace, 4)

	if !reflect.DeepEqual(testVectorA.Get(0), gcv.NewValue()) {
		t.Errorf("Expected %v, received %v", gcv.NewValue(), testVectorA.Get(0))
	}

	testVectorB.Set(0, gcv.MakeValue(2+2i))

	if !reflect.DeepEqual(testVectorB.Get(0), gcv.MakeValue(2+2i)) {
		t.Errorf("Expected %v, received %v", gcv.MakeValue(2+2i), testVectorB.Get(0))
	}
}

func TestSetValuesVectors(t *testing.T) {
	testVector := NewVector(ColSpace, 3)
	testVectors := MakeVectors(ColSpace, testVector, testVector, testVector)
	testValue := gcv.MakeValue(3.0 + 1i)

	testVectors.SetValue(0, 1, testValue)

	if !reflect.DeepEqual(testVectors.Get(0).Get(1), testValue) {
		t.Fail()
	}
}

func TestSpaceMethodVector(t *testing.T) {
	testVectorA := NewVector(ColSpace, 3)
	testVectorB := NewVector(RowSpace, 4)

	if testVectorA.Space() != ColSpace {
		t.Errorf("Expected %v, received %v", ColSpace, testVectorA.Space())
	}

	if testVectorB.Space() != RowSpace {
		t.Errorf("Expected %v, received %v", RowSpace, testVectorB.Space())
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

func TestTransMethods(t *testing.T) {
	testElementsA := gcv.MakeValues(1, 2, 3)
	testVectorA := MakeVectorAlt(ColSpace, testElementsA)
	testTransVectorA := MakeVectorAlt(RowSpace, testElementsA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", testVectorA, testTransVectorA)
	}

	testVectorB := MakeVector(ColSpace, 1+1i, 2, 3-3i)
	testTransVectorB := MakeVector(RowSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2), gcv.MakeValue(3+3i))

	testVectorB.ConjTrans()

	if !reflect.DeepEqual(testVectorB.Get(0), testTransVectorB.Get(0)) &&
		!reflect.DeepEqual(testVectorB.Get(1), testTransVectorB.Get(1)) &&
		!reflect.DeepEqual(testVectorB.Get(2), testTransVectorB.Get(2)) {
		t.Errorf("Expected %v, received %v", testTransVectorB, testVectorB)
	}

	testVectorC := MakeVectorAlt(RowSpace, testElementsA)
	testTransVectorC := MakeVectorAlt(ColSpace, testElementsA)

	testVectorC.Trans()

	if !reflect.DeepEqual(testVectorC, testTransVectorC) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorC, testTransVectorC))
	}
}

func TestNormMethodVector(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4))

	if !reflect.DeepEqual(testVectorA.Norm(), gcv.MakeValue(5.0)) {
		t.Errorf("Expected %v, received %v", gcv.MakeValue(5), testVectorA.Norm())
	}

	testVectorB := MakeVector(ColSpace, gcv.MakeValue(2+2i), gcv.MakeValue(1))

	if !reflect.DeepEqual(testVectorB.Norm(), gcv.MakeValue(3+0i)) {
		t.Errorf("Expected %v, received %v", gcv.MakeValue(3), testVectorB.Norm())
	}
}

func TestMakeNewTransVector(t *testing.T) {
	testVectorA := NewVector(RowSpace, 4)
	testTransVectorA := MakeTransVector(testVectorA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}
}

func TestMakeNewConjTransVector(t *testing.T) {
	testVectorA := NewVector(RowSpace, 4)
	testTransVectorA := MakeConjTransVector(testVectorA)

	testVectorA.Trans()

	if !reflect.DeepEqual(testVectorA, testTransVectorA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testVectorA, testTransVectorA))
	}
}

func TestNewVectorsGet(t *testing.T) {
	testVectors := []Vector{MakeVector(RowSpace, 3, 4), MakeVector(ColSpace, 6, 7)}
	testVectorsA := MakeVectorsAlt(RowSpace, testVectors)

	if !reflect.DeepEqual(testVectors[0], testVectorsA.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectors[0], testVectorsA.Get(0))
	}

	if reflect.DeepEqual(testVectors[1], testVectorsA.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectors[1], testVectorsA.Get(1))
	}

	testVectorC := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4))

	testVectorD := MakeVector(RowSpace, gcv.MakeValue(6+3i), gcv.MakeValue(7-4i))

	testVectorsB := MakeVectors(RowSpace, testVectorC, testVectorD)

	if reflect.DeepEqual(testVectorC, testVectorsB.Get(0)) {
		t.Errorf("Expected %v, received %v", testVectorC, testVectorsB.Get(0))
	}

	if !reflect.DeepEqual(testVectorD, testVectorsB.Get(1)) {
		t.Errorf("Expected %v, received %v", testVectorD, testVectorsB.Get(1))
	}
}

func TestVectorIndexOf(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0))

	testIndexA := testVectorA.IndexOf(gcv.MakeValue(3))
	if testIndexA != 0 {
		t.Errorf("Expected %v, received %v", 0, testIndexA)
	}

	testIndexB := testVectorA.IndexOf(gcv.MakeValue(5 + 1i))
	if testIndexB != -1 {
		t.Errorf("Expected %v, received %v", 0, testIndexB)
	}
}

func TestVectorsIndexOf(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5+1i))

	testVectorB := MakeVector(ColSpace, gcv.MakeValue(6), gcv.MakeValue(7))

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	indexOfTestVectorA := testVectorsA.IndexOf(testVectorA)
	if indexOfTestVectorA != 0 {
		t.Errorf("Expected %v, received %v", 0, indexOfTestVectorA)
	}

	indexOfTestVectorB := testVectorsA.IndexOf(testVectorB)
	if indexOfTestVectorB != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorB)
	}

	testVectorC := MakeVector(ColSpace, gcv.MakeValue(4), gcv.MakeValue(5.0), gcv.MakeValue(8+1i))

	indexOfTestVectorC := testVectorsA.IndexOf(testVectorC)
	if indexOfTestVectorC != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testVectorD := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(5.0), gcv.MakeValue(8+1i))

	indexOfTestVectorD := testVectorsA.IndexOf(testVectorD)
	if indexOfTestVectorD != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorC)
	}

	testVectorE := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(8+1i))

	indexOfTestVectorE := testVectorsA.IndexOf(testVectorE)
	if indexOfTestVectorE != -1 {
		t.Errorf("Expected %v, received %v", -1, indexOfTestVectorE)
	}
}

func TestSetAndSpaceVectors(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5))

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testVectorB := MakeVector(RowSpace, gcv.MakeValue(6), gcv.MakeValue(7+1i))

	testVectorsA.Set(0, testVectorB)

	if !reflect.DeepEqual(testVectorB, testVectorsA.Get(0)) {
		t.Fail()
	}

	if testVectorsA.Space() != ColSpace {
		t.Fail()
	}
}

func TestCopyVectors(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5+1i))

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testCopyVector := testVectorsA.Copy()

	if !reflect.DeepEqual(testCopyVector, testVectorsA) {
		t.Fail()
	}
}

func TestAppendVectors(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5+1i))

	testVectorsA := MakeVectors(ColSpace, testVectorA)

	testVectorB := MakeVector(ColSpace, gcv.MakeValue(6), gcv.MakeValue(7))

	testVectorB.Append(gcv.MakeValue(8))

	if !reflect.DeepEqual(testVectorB.Get(2), gcv.MakeValue(8)) {
		t.Fail()
	}

	testVectorsA.Append(testVectorB)

	if !reflect.DeepEqual(testVectorsA.Get(1), testVectorB) {
		t.Fail()
	}

	testVectorsB := MakeVectors(ColSpace)

	testVectorsB.Append(testVectorB)

	if !reflect.DeepEqual(testVectorsB.Get(0), testVectorB) {
		t.Fail()
	}
}

func TestSubsetAndLenVectors(t *testing.T) {
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5+1i))

	testVectorB := MakeVector(ColSpace, gcv.MakeValue(6), gcv.MakeValue(7))

	testVectorC := MakeVector(ColSpace, gcv.MakeValue(42))

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
	testVectorA := MakeVector(ColSpace, gcv.MakeValue(3), gcv.MakeValue(4.0), gcv.MakeValue(5+1i))

	testVectorB := MakeVector(ColSpace, gcv.MakeValue(6), gcv.MakeValue(7))

	testVectorsA := MakeVectors(ColSpace, testVectorA, testVectorB)

	if !reflect.DeepEqual(testVectorsA.Vectors()[0], testVectorA) || !reflect.DeepEqual(testVectorsA.Vectors()[1], testVectorB) {
		t.Fail()
	}

	testVectorsB := NewVectors(RowSpace, 3, 4)

	if testVectorsB.Space() != RowSpace {
		t.Error("Expected Type RowSpace")
	}

	if !reflect.DeepEqual(testVectorsB.Get(2).Get(3), gcv.MakeValue(0)) {
		t.Fail()
	}

}
