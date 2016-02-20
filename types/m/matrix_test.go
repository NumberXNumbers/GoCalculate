package m

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	testMatrixA, _, _ := NewMatrix(3, 3, reflect.Float64)
	_, testMatrixB, _ := NewMatrix(3, 3, reflect.Complex128)
	_, _, testMatrixC := NewMatrix(3, 3, reflect.Int)

	if testMatrixA.Type() != reflect.Float64 {
		t.Errorf("Expected %s, received %s", reflect.Float64, testMatrixA.Type())
	}

	if testMatrixB.Type() != reflect.Complex128 {
		t.Errorf("Expected %s, received %s", reflect.Complex128, testMatrixB.Type())
	}

	if testMatrixC == nil {
		t.Errorf("Expected error, received %s", testMatrixC)
	}
}

func TestGetNumColsAndGetNumRowsMethodsMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 4)
	testMatrixB := MakeComplexMatrix(4, 3)

	if testMatrixA.GetNumRows() != 3 {
		t.Errorf("Expected %d, received %d", 3, testMatrixA.GetNumRows())
	}

	if testMatrixA.GetNumCols() != 4 {
		t.Errorf("Expected %d, received %d", 4, testMatrixA.GetNumCols())
	}

	if testMatrixB.GetNumRows() != 4 {
		t.Errorf("Expected %d, received %d", 4, testMatrixB.GetNumRows())
	}

	if testMatrixB.GetNumCols() != 3 {
		t.Errorf("Expected %d, received %d", 3, testMatrixB.GetNumCols())
	}
}

func TestIsIdenityMethodMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeMatrix(4, 3)
	testMatrixC := MakeIdentityMatrix(3)

	if testMatrixA.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixA.IsIdentity())
	}

	if testMatrixB.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixB.IsIdentity())
	}

	if !testMatrixC.IsIdentity() {
		t.Errorf("Expected %v, received %v", true, testMatrixC.IsIdentity())
	}

	testMatrixD := MakeComplexMatrix(3, 3)
	testMatrixE := MakeComplexMatrix(3, 3)
	testMatrixF := MakeIdentityComplexMatrix(3)

	if testMatrixD.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixD.IsIdentity())
	}

	if testMatrixE.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixE.IsIdentity())
	}

	if !testMatrixF.IsIdentity() {
		t.Errorf("Expected %v, received %v", false, testMatrixF.IsIdentity())
	}
}

func TestCopyMethodMatrix(t *testing.T) {
	testElementsA := [][]float64{{1, 2}, {1, 2}}
	testMatrixA := MakeMatrixWithElements(testElementsA)

	if !reflect.DeepEqual(testMatrixA, testMatrixA.Copy()) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixA, testMatrixA.Copy()))
	}

	testElementsB := [][]complex128{{1 + 1i, 2}, {1, 2 - 1i}}
	testMatrixB := MakeComplexMatrixWithElements(testElementsB)

	if !reflect.DeepEqual(testMatrixB, testMatrixB.Copy()) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixB, testMatrixB.Copy()))
	}

	if reflect.DeepEqual(testMatrixA, testMatrixB.Copy()) {
		t.Errorf("Expected %v, received %v", false, reflect.DeepEqual(testMatrixA, testMatrixB.Copy()))
	}

	if reflect.DeepEqual(testMatrixB, testMatrixA.Copy()) {
		t.Errorf("Expected %v, received %v", false, reflect.DeepEqual(testMatrixB, testMatrixA.Copy()))
	}
}

func TestTransMethodMatrix(t *testing.T) {
	testElementsA := [][]float64{{1, 2}, {1, 2}}
	testElementsTransA := [][]float64{{1, 1}, {2, 2}}
	testMatrixA := MakeMatrixWithElements(testElementsA)
	testMatrixTransA := MakeMatrixWithElements(testElementsTransA)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixTransA, testMatrixA) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixTransA, testMatrixA))
	}

	testElementsB := [][]float64{{1, 2}, {1, 2}, {1, 2}}
	testElementsTransB := [][]float64{{1, 1, 1}, {2, 2, 2}}
	testMatrixB := MakeMatrixWithElements(testElementsB)
	testMatrixTransB := MakeMatrixWithElements(testElementsTransB)

	testMatrixB.Trans()
	if !reflect.DeepEqual(testMatrixTransB, testMatrixB) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixTransB, testMatrixB))
	}

	testElementsC := [][]complex128{{1 + 1i, 2}, {1, 2 + 1i}}
	testElementsTransC := [][]complex128{{1 - 1i, 1}, {2, 2 - 1i}}
	testMatrixC := MakeComplexMatrixWithElements(testElementsC)
	testMatrixTransC := MakeComplexMatrixWithElements(testElementsTransC)

	testMatrixC.Trans()
	if !reflect.DeepEqual(testMatrixTransC, testMatrixC) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixTransC, testMatrixC))
	}

	testElementsD := [][]complex128{{1 + 1i, 2}, {1, 2 + 1i}, {1 - 1i, 2}}
	testElementsTransD := [][]complex128{{1 - 1i, 1, 1 + 1i}, {2, 2 - 1i, 2}}
	testMatrixD := MakeComplexMatrixWithElements(testElementsD)
	testMatrixTransD := MakeComplexMatrixWithElements(testElementsTransD)

	testMatrixD.Trans()
	if !reflect.DeepEqual(testMatrixTransD, testMatrixD) {
		t.Errorf("Expected %v, received %v", false, reflect.DeepEqual(testMatrixTransD, testMatrixD))
	}
}

func TestTrMethodMatrix(t *testing.T) {
	testElementsA := [][]float64{{1, 0}, {0, 1}}
	testMatrixA := MakeMatrixWithElements(testElementsA)

	if trA, _ := testMatrixA.Tr(); trA != float64(2) {
		t.Errorf("Expected %f, received %f", float64(2), trA)
	}

	testMatrixB := MakeIdentityMatrix(2)

	if trB, _ := testMatrixB.Tr(); trB != float64(2) {
		t.Errorf("Expected %f, received %f", float64(2), trB)
	}

	testElementsC := [][]complex128{{1 + 1i, 0}, {0, 1 - 1i}}
	testMatrixC := MakeComplexMatrixWithElements(testElementsC)

	if trC, _ := testMatrixC.Tr(); trC != complex128(2) {
		t.Errorf("Expected %f, received %f", complex128(2), trC)
	}

	testMatrixD := MakeIdentityComplexMatrix(2)

	if trD, _ := testMatrixD.Tr(); trD != complex128(2) {
		t.Errorf("Expected %f, received %f", complex128(2), trD)
	}

	testMatrixE := MakeMatrix(3, 4)

	if _, err := testMatrixE.Tr(); err == nil {
		t.Errorf("Expected err, received %v", err)
	}

	testMatrixF := MakeComplexMatrix(3, 4)

	if _, err := testMatrixF.Tr(); err == nil {
		t.Errorf("Expected err, received %v", err)
	}
}

func TestSetAndGetMethodsMatrix(t *testing.T) {
	testElementsA := [][]float64{{0, 0}, {0, 0}}
	testMatrixA := MakeMatrixWithElements(testElementsA)

	if testMatrixA.Get(0, 0) != float64(0) {
		t.Errorf("Expected %f, received %f", float64(0), testMatrixA.Get(0, 0))
	}

	testMatrixA.Set(0, 0, 1)

	if testMatrixA.Get(0, 0) != float64(1) {
		t.Errorf("Expected %f, received %f", float64(1), testMatrixA.Get(0, 0))
	}

	testElementsB := [][]complex128{{1 + 1i, 0}, {0, 1 - 1i}}
	testMatrixB := MakeComplexMatrixWithElements(testElementsB)

	if testMatrixB.Get(0, 0) != complex128(1+1i) {
		t.Errorf("Expected %f, received %f", complex128(1+1i), testMatrixB.Get(0, 0))
	}

	testMatrixB.Set(0, 0, 1-1i)

	if testMatrixB.Get(0, 0) != complex128(1-1i) {
		t.Errorf("Expected %f, received %f", complex128(1-1i), testMatrixB.Get(0, 0))
	}
}

func TestTypeMethodMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeComplexMatrix(3, 3)

	if testMatrixA.Type() != reflect.Float64 {
		t.Errorf("Expected %s, received %s", reflect.Float64, testMatrixA.Type())
	}

	if testMatrixB.Type() != reflect.Complex128 {
		t.Errorf("Expected %s, received %s", reflect.Complex128, testMatrixB.Type())
	}
}

func TestIsSquareMethodMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeComplexMatrix(3, 3)
	testMatrixC := MakeMatrix(3, 4)
	testMatrixD := MakeComplexMatrix(4, 3)

	if !testMatrixA.IsSquare() {
		t.Errorf("Expected %v, received %v", true, testMatrixA.IsSquare())
	}

	if !testMatrixB.IsSquare() {
		t.Errorf("Expected %v, received %v", true, testMatrixB.IsSquare())
	}

	if testMatrixC.IsSquare() {
		t.Errorf("Expected %v, received %v", false, testMatrixC.IsSquare())
	}

	if testMatrixD.IsSquare() {
		t.Errorf("Expected %v, received %v", false, testMatrixD.IsSquare())
	}
}

func TestDimMethodMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeComplexMatrix(3, 3)
	testMatrixC := MakeMatrix(3, 4)
	testMatrixD := MakeComplexMatrix(4, 3)

	if rowsA, colsA := testMatrixA.Dim(); rowsA != 3 || colsA != 3 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 3, rowsA, colsA)
	}

	if rowsB, colsB := testMatrixB.Dim(); rowsB != 3 || colsB != 3 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 3, rowsB, colsB)
	}

	if rowsC, colsC := testMatrixC.Dim(); rowsC != 3 || colsC != 4 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 3, 4, rowsC, colsC)
	}

	if rowsD, colsD := testMatrixD.Dim(); rowsD != 4 || colsD != 3 {
		t.Errorf("Expected (%d, %d), received (%d, %d)", 4, 3, rowsD, colsD)
	}
}

func TestNumElementsMethodMatrix(t *testing.T) {
	testMatrixA := MakeMatrix(3, 3)
	testMatrixB := MakeComplexMatrix(3, 3)
	testMatrixC := MakeMatrix(3, 4)
	testMatrixD := MakeComplexMatrix(4, 3)

	if testMatrixA.NumElements() != 9 {
		t.Errorf("Expected %d, received %d", 9, testMatrixA.NumElements())
	}

	if testMatrixB.NumElements() != 9 {
		t.Errorf("Expected %d, received %d", 9, testMatrixB.NumElements())
	}

	if testMatrixC.NumElements() != 12 {
		t.Errorf("Expected %d, received %d", 12, testMatrixA.NumElements())
	}

	if testMatrixD.NumElements() != 12 {
		t.Errorf("Expected %d, received %d", 12, testMatrixB.NumElements())
	}
}

func TestGetElementsMethodMatrix(t *testing.T) {
	testElementsA := [][]float64{{1, 0}, {0, 1}}
	testMatrixA := MakeMatrixWithElements(testElementsA)
	testElementsB := [][]complex128{{1 + 1i, 0}, {0, 1 - 1i}}
	testMatrixB := MakeComplexMatrixWithElements(testElementsB)

	if !reflect.DeepEqual(testMatrixA.GetElements(), testElementsA) {
		t.Errorf("Expected %v, received %v", testElementsA, testMatrixA.GetElements())
	}

	if !reflect.DeepEqual(testMatrixB.GetElements(), testElementsB) {
		t.Errorf("Expected %v, received %v", testElementsB, testMatrixB.GetElements())
	}
}

func TestMakeNewConjMatrix(t *testing.T) {
	testElementsA := [][]complex128{{1 + 1i, 0}, {0, 1 - 1i}}
	testMatrixA := MakeComplexMatrixWithElements(testElementsA)
	testMatrixB := MakeNewConjMatrix(testMatrixA)

	testMatrixA.Trans()
	if !reflect.DeepEqual(testMatrixA, testMatrixB) {
		t.Errorf("Expected %v, received %v", true, reflect.DeepEqual(testMatrixA, testMatrixB))
	}
}
