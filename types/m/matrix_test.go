package m

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestGetNumColsAndGetNumRowsMethodsMatrix(t *testing.T) {
	testMatrixA := NewMatrix(3, 4)

	if testMatrixA.GetNumRows() != 3 {
		t.Errorf("Expected %d, received %d", 3, testMatrixA.GetNumRows())
	}

	if testMatrixA.GetNumCols() != 4 {
		t.Errorf("Expected %d, received %d", 4, testMatrixA.GetNumCols())
	}
}

func TestIsIdenityMethodMatrix(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(4, 3)
	testMatrixC := NewIdentityMatrix(3)

	if testMatrixA.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixA.IsIdentity())
	}

	if testMatrixB.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixB.IsIdentity())
	}

	if !testMatrixC.IsIdentity() {
		t.Errorf("Expected %v, received %v", true, testMatrixC.IsIdentity())
	}
}

func TestCopyMethodMatrix(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(1), gcv.NewValue(2))
	testElementsAb := gcv.NewValues(gcv.NewValue(0), gcv.NewValue(0))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrix(testVectorsA)

	if !reflect.DeepEqual(testMatrixA, testMatrixA.Copy()) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixA, testMatrixA.Copy()))
	}
}

func TestTransMethodMatrix(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testElementsAb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrix(testVectorsA)

	testElementsAc := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(1))
	testElementsAd := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(2))
	testVectorAc := v.MakeVector(v.RowSpace, testElementsAc)
	testVectorAd := v.MakeVector(v.RowSpace, testElementsAd)
	testVectorsTransA := v.MakeVectors(v.RowSpace, testVectorAc, testVectorAd)
	testMatrixTransA := MakeMatrix(testVectorsTransA)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixTransA, testMatrixA) {
		t.Errorf("Expected %v, received %v", testMatrixTransA, testMatrixA)
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1+1i), gcv.NewValue(2))
	testElementsBb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2+1i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := MakeMatrix(testVectorsB)

	testElementsBc := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(1))
	testElementsBd := gcv.NewValues(gcv.NewValue(2), gcv.NewValue(2-1i))
	testVectorBc := v.MakeVector(v.RowSpace, testElementsBc)
	testVectorBd := v.MakeVector(v.RowSpace, testElementsBd)
	testVectorsTransB := v.MakeVectors(v.RowSpace, testVectorBc, testVectorBd)
	testMatrixTransB := MakeMatrix(testVectorsTransB)

	testMatrixB.Trans()
	if !reflect.DeepEqual(testMatrixTransB.Get(0, 0), testMatrixB.Get(0, 0)) ||
		testMatrixTransB.Get(0, 1).Complex128() != testMatrixB.Get(0, 1).Complex128() ||
		testMatrixTransB.Get(1, 0).Complex128() != testMatrixB.Get(1, 0).Complex128() ||
		!reflect.DeepEqual(testMatrixTransB.Get(1, 1), testMatrixB.Get(1, 1)) ||
		testMatrixTransB.Type() != gcv.Complex {
		t.Errorf("Expected %v, received %v", testMatrixTransB, testMatrixB)
	}
}

func TestTrMethodMatrix(t *testing.T) {
	testMatrixA := NewIdentityMatrix(2)

	if trA, _ := testMatrixA.Tr(); trA.Float64() != 2 {
		t.Errorf("Expected %d, received %f", 2, trA.Float64())
	}

	testElementsBa := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2))
	testElementsBb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(1+1i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := MakeMatrix(testVectorsB)

	if trB, _ := testMatrixB.Tr(); trB.Complex128() != complex128(2) {
		t.Errorf("Expected %v, received %v", complex128(2), trB.Complex128())
	}

	testMatrixC := NewMatrix(3, 4)

	if _, errC := testMatrixC.Tr(); errC == nil {
		t.Errorf("Expected err, received %v", errC)
	}
}

func TestSetAndGetMethodsMatrix(t *testing.T) {
	testMatrixA := NewMatrix(2, 2)

	testMatrixA.Set(0, 0, gcv.NewValue(1))

	if testMatrixA.Get(0, 0).Float64() != 1 {
		t.Errorf("Expected %d, received %f", 1, testMatrixA.Get(0, 0).Float64())
	}

	testMatrixA.Set(0, 1, gcv.NewValue(1+1i))

	if testMatrixA.Get(0, 1).Complex128() != 1+1i {
		t.Errorf("Expected %d, received %f", 1, testMatrixA.Get(0, 0).Complex128())
	}
}

func TestTypeMethodMatrix(t *testing.T) {
	testElementsBa := gcv.NewValues(gcv.NewValue(1-1i), gcv.NewValue(2))
	testElementsBb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(1+1i))
	testVectorBa := v.MakeVector(v.RowSpace, testElementsBa)
	testVectorBb := v.MakeVector(v.RowSpace, testElementsBb)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := MakeMatrix(testVectorsB)

	if testMatrixB.Type() != gcv.Complex {
		t.Errorf("Expected %s, received %s", gcv.Complex, testMatrixB.Type())
	}
}

func TestIsSquareMethodMatrix(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if !testMatrixA.IsSquare() {
		t.Errorf("Expected %v, received %v", true, testMatrixA.IsSquare())
	}

	if testMatrixB.IsSquare() {
		t.Errorf("Expected %v, received %v", false, testMatrixB.IsSquare())
	}
}

func TestDimMethodMatrix(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if rowsA, colsA := testMatrixA.Dim(); rowsA != 3 || colsA != 3 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 3, rowsA, colsA)
	}

	if rowsB, colsB := testMatrixB.Dim(); rowsB != 3 || colsB != 4 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 4, rowsB, colsB)
	}
}

func TestNumElementsMethodMatrix(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if testMatrixA.TotalElements() != 9 {
		t.Errorf("Expected %d, received %d", 9, testMatrixA.TotalElements())
	}

	if testMatrixB.TotalElements() != 12 {
		t.Errorf("Expected %d, received %d", 12, testMatrixB.TotalElements())
	}
}

func TestGetElementsMethodMatrix(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testElementsAb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrix(testVectorsA)

	if !reflect.DeepEqual(testMatrixA.Elements(), testVectorsA) {
		t.Errorf("Expected %v, received %v", testVectorsA, testMatrixA.Elements())
	}
}

func TestMakeConjMatrix(t *testing.T) {
	testElementsAa := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testElementsAb := gcv.NewValues(gcv.NewValue(1), gcv.NewValue(2))
	testVectorAa := v.MakeVector(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVector(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrix(testVectorsA)
	solutionMatrix := MakeConjMatrix(testMatrixA)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixA, solutionMatrix) {
		t.Errorf("Expected %v, received %v", solutionMatrix, testMatrixA)
	}
}
