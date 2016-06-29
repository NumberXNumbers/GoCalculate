package m

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestGetNumColsAndGetNumRows(t *testing.T) {
	testMatrixA := NewMatrix(3, 4)

	if testMatrixA.GetNumRows() != 3 {
		t.Errorf("Expected %d, received %d", 3, testMatrixA.GetNumRows())
	}

	if testMatrixA.GetNumCols() != 4 {
		t.Errorf("Expected %d, received %d", 4, testMatrixA.GetNumCols())
	}
}

func TestIsIdenity(t *testing.T) {
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

func TestCopy(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(0), gcv.MakeValue(0))
	testMatrixA := MakeMatrix(testVectorAa, testVectorAb)

	if !reflect.DeepEqual(testMatrixA, testMatrixA.Copy()) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixA, testMatrixA.Copy()))
	}
}

func TestTrans(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testMatrixA := MakeMatrix(testVectorAa, testVectorAb)

	testVectorAc := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1))
	testVectorAd := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(2))
	testMatrixTransA := MakeMatrix(testVectorAc, testVectorAd)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixTransA, testMatrixA) {
		t.Errorf("Expected %v, received %v", testMatrixTransA, testMatrixA)
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(1+1i), gcv.MakeValue(2))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2+1i))
	testMatrixB := MakeMatrix(testVectorBa, testVectorBb)

	testVectorBc := v.MakeVector(v.RowSpace, gcv.MakeValue(1-1i), gcv.MakeValue(1))
	testVectorBd := v.MakeVector(v.RowSpace, gcv.MakeValue(2), gcv.MakeValue(2-1i))
	testMatrixTransB := MakeMatrix(testVectorBc, testVectorBd)

	testMatrixB.Trans()
	if !reflect.DeepEqual(testMatrixTransB.Get(0, 0), testMatrixB.Get(0, 0)) ||
		testMatrixTransB.Get(0, 1).Complex() != testMatrixB.Get(0, 1).Complex() ||
		testMatrixTransB.Get(1, 0).Complex() != testMatrixB.Get(1, 0).Complex() ||
		!reflect.DeepEqual(testMatrixTransB.Get(1, 1), testMatrixB.Get(1, 1)) ||
		testMatrixTransB.Type() != gcv.Complex {
		t.Errorf("Expected %v, received %v", testMatrixTransB, testMatrixB)
	}
}

func TestTr(t *testing.T) {
	testMatrixA := NewIdentityMatrix(2)

	if trA, _ := testMatrixA.Tr(); trA.Real() != 2 {
		t.Errorf("Expected %d, received %f", 2, trA.Real())
	}

	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1+1i))
	testMatrixB := MakeMatrix(testVectorBa, testVectorBb)

	if trB, _ := testMatrixB.Tr(); trB.Complex() != complex128(2) {
		t.Errorf("Expected %v, received %v", complex128(2), trB.Complex())
	}

	testMatrixC := NewMatrix(3, 4)

	if _, errC := testMatrixC.Tr(); errC == nil {
		t.Errorf("Expected err, received %v", errC)
	}
}

func TestSetAndGet(t *testing.T) {
	testMatrixA := NewMatrix(2, 2)

	testMatrixA.Set(0, 0, gcv.MakeValue(1))

	if testMatrixA.Get(0, 0).Real() != 1 {
		t.Errorf("Expected %d, received %f", 1, testMatrixA.Get(0, 0).Real())
	}

	testMatrixA.Set(0, 1, gcv.MakeValue(1+1i))

	if testMatrixA.Get(0, 1).Complex() != 1+1i {
		t.Errorf("Expected %d, received %f", 1, testMatrixA.Get(0, 0).Complex())
	}
}

func TestType(t *testing.T) {
	testVectorBa := v.MakeVector(v.RowSpace, gcv.MakeValue(1-1i), gcv.MakeValue(2))
	testVectorBb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1+1i))
	testMatrixB := MakeMatrix(testVectorBa, testVectorBb)

	if testMatrixB.Type() != gcv.Complex {
		t.Errorf("Expected %v, received %v", gcv.Complex, testMatrixB.Type())
	}
}

func TestIsSquare(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if !testMatrixA.IsSquare() {
		t.Errorf("Expected %v, received %v", true, testMatrixA.IsSquare())
	}

	if testMatrixB.IsSquare() {
		t.Errorf("Expected %v, received %v", false, testMatrixB.IsSquare())
	}
}

func TestDim(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if rowsA, colsA := testMatrixA.Dim(); rowsA != 3 || colsA != 3 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 3, rowsA, colsA)
	}

	if rowsB, colsB := testMatrixB.Dim(); rowsB != 3 || colsB != 4 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 4, rowsB, colsB)
	}
}

func TestNumElements(t *testing.T) {
	testMatrixA := NewMatrix(3, 3)
	testMatrixB := NewMatrix(3, 4)

	if testMatrixA.TotalElements() != 9 {
		t.Errorf("Expected %d, received %d", 9, testMatrixA.TotalElements())
	}

	if testMatrixB.TotalElements() != 12 {
		t.Errorf("Expected %d, received %d", 12, testMatrixB.TotalElements())
	}
}

func TestGetElements(t *testing.T) {
	testElementsAa := gcv.MakeValues(gcv.MakeValue(1), gcv.MakeValue(2))
	testElementsAb := gcv.MakeValues(gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAa := v.MakeVectorAlt(v.RowSpace, testElementsAa)
	testVectorAb := v.MakeVectorAlt(v.RowSpace, testElementsAb)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrixAlt(testVectorsA)

	if !reflect.DeepEqual(testMatrixA.Elements(), testVectorsA) {
		t.Errorf("Expected %v, received %v", testVectorsA, testMatrixA.Elements())
	}
}

func TestMakeConj(t *testing.T) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrixAlt(testVectorsA)
	solutionMatrix := MakeConjMatrix(testMatrixA)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixA, solutionMatrix) {
		t.Errorf("Expected %v, received %v", solutionMatrix, testMatrixA)
	}
}

func TestSwap(t *testing.T) {
	testVectorAa := v.MakeVectorPure(v.RowSpace, 1, 2)
	testVectorAb := v.MakeVectorPure(v.RowSpace, 3, 2)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrixAlt(testVectorsA)
	solutionMatrix := MakeMatrix(testVectorAb, testVectorAa)

	testMatrixA.Swap(0, 1)
	if !reflect.DeepEqual(testMatrixA, solutionMatrix) {
		t.Errorf("Expected %v, received %v", solutionMatrix, testMatrixA)
	}
}

func TestDet(t *testing.T) {
	testVectorAa := v.MakeVectorPure(v.RowSpace, 1, 2)
	testVectorAb := v.MakeVectorPure(v.RowSpace, 3, 2)
	testVectorsA := v.MakeVectors(v.RowSpace, testVectorAa, testVectorAb)
	testMatrixA := MakeMatrixAlt(testVectorsA)
	detA, errA := testMatrixA.Det()

	if errA != nil {
		t.Fail()
	}

	if detA.Real() != -4 {
		t.Errorf("Expected %v, received %v", -4, detA)
	}

	testVectorBa := v.MakeVectorPure(v.RowSpace, 4, 2)
	testVectorBb := v.MakeVectorPure(v.RowSpace, 3, 2)
	testVectorsB := v.MakeVectors(v.RowSpace, testVectorBa, testVectorBb)
	testMatrixB := MakeMatrixAlt(testVectorsB)
	detB, errB := testMatrixB.Det()

	if errB != nil {
		t.Fail()
	}

	if detB.Real() != 2 {
		t.Errorf("Expected %v, received %v", 2, detB)
	}

	testVectorCa := v.MakeVectorPure(v.RowSpace, 1, 2, 1)
	testVectorCb := v.MakeVectorPure(v.RowSpace, 3, 2, 3)
	testVectorCc := v.MakeVectorPure(v.RowSpace, 2, 3, 4)
	testVectorsC := v.MakeVectors(v.RowSpace, testVectorCa, testVectorCb, testVectorCc)
	testMatrixC := MakeMatrixAlt(testVectorsC)
	detC, errC := testMatrixC.Det()

	if errC != nil {
		t.Fail()
	}

	if detC.Real() != -8 {
		t.Errorf("Expected %v, received %v", -8, detC)
	}

	testVectorsD := v.MakeVectors(v.RowSpace, testVectorCa, testVectorCb)
	testMatrixD := MakeMatrixAlt(testVectorsD)
	_, errD := testMatrixD.Det()

	if errD == nil {
		t.Error("Expected error")
	}
}
