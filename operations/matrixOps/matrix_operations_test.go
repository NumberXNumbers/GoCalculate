package matrixOps

import (
	"reflect"
	"testing"

	ct "github.com/traviscox1990/GoCalculate/coreTypes"
)

func TestMatrixScalarMulti(t *testing.T) {
	testElementsA := [][]float64{{1, 0}, {0, 1}}
	testMatrixA := ct.MakeMatrixWithElements(testElementsA)

	testScalarA := 2.0

	resultMatrixA := MatrixScalarMulti(testScalarA, testMatrixA)

	solutionElementsA := [][]float64{{2, 0}, {0, 2}}
	solutionMatrixA := ct.MakeMatrixWithElements(solutionElementsA)

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsB := [][]complex128{{1, 0}, {0, 1}}
	testMatrixB := ct.MakeComplexMatrixWithElements(testElementsB)

	testScalarB := 2.0 + 1i

	resultMatrixB := MatrixComplexScalarMulti(testScalarB, testMatrixB)

	solutionElementsB := [][]complex128{{2 + 1i, 0}, {0, 2 + 1i}}
	solutionMatrixB := ct.MakeComplexMatrixWithElements(solutionElementsB)

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixB, resultMatrixB))
	}
}

func TestMatrixMultiSimple(t *testing.T) {
	testElementsAa := [][]float64{{1, 0}, {0, 1}}
	testMatrixAa := ct.MakeMatrixWithElements(testElementsAa)

	testElementsAb := [][]float64{{0, 1}, {1, 0}}
	testMatrixAb := ct.MakeMatrixWithElements(testElementsAb)

	solutionElementsA := [][]float64{{0, 1}, {1, 0}}
	solutionMatrixA := ct.MakeMatrixWithElements(solutionElementsA)

	resultMatrixA, errA := MatrixMultiSimple(testMatrixAa, testMatrixAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsBa := [][]float64{{2, 0}, {0, 2}}
	testMatrixBa := ct.MakeMatrixWithElements(testElementsBa)

	testElementsBb := [][]float64{{0, 1}, {1, 0}}
	testMatrixBb := ct.MakeMatrixWithElements(testElementsBb)

	solutionElementsB := [][]float64{{0, 2}, {2, 0}}
	solutionMatrixB := ct.MakeMatrixWithElements(solutionElementsB)

	resultMatrixB, errB := MatrixMultiSimple(testMatrixBa, testMatrixBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsCa := [][]float64{{0, 1}, {1, 0}}
	testMatrixCa := ct.MakeMatrixWithElements(testElementsCa)

	testElementsCb := [][]float64{{1, 0}, {0, 1}}
	testMatrixCb := ct.MakeMatrixWithElements(testElementsCb)

	solutionElementsC := [][]float64{{0, 1}, {1, 0}}
	solutionMatrixC := ct.MakeMatrixWithElements(solutionElementsC)

	resultMatrixC, errC := MatrixMultiSimple(testMatrixCa, testMatrixCb)

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixC, resultMatrixC) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixC, resultMatrixC))
	}

	testElementsDa := [][]float64{{1, 0}, {0, 1}, {0, 1}}
	testMatrixDa := ct.MakeMatrixWithElements(testElementsDa)

	testElementsDb := [][]float64{{0, 1}, {1, 0}, {0, 1}}
	testMatrixDb := ct.MakeMatrixWithElements(testElementsDb)

	_, errD := MatrixMultiSimple(testMatrixDa, testMatrixDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := [][]complex128{{1, 0}, {0, 1}}
	testMatrixEa := ct.MakeComplexMatrixWithElements(testElementsEa)

	testElementsEb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixEb := ct.MakeComplexMatrixWithElements(testElementsEb)

	solutionElementsE := [][]complex128{{0, 1}, {1, 0}}
	solutionMatrixE := ct.MakeComplexMatrixWithElements(solutionElementsE)

	resultMatrixE, errE := MatrixComplexMultiSimple(testMatrixEa, testMatrixEb)

	if errE != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixE, resultMatrixE) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixE, resultMatrixE))
	}

	testElementsFa := [][]complex128{{2, 0}, {0, 2}}
	testMatrixFa := ct.MakeComplexMatrixWithElements(testElementsFa)

	testElementsFb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixFb := ct.MakeComplexMatrixWithElements(testElementsFb)

	solutionElementsF := [][]complex128{{0, 2}, {2, 0}}
	solutionMatrixF := ct.MakeComplexMatrixWithElements(solutionElementsF)

	resultMatrixF, errF := MatrixComplexMultiSimple(testMatrixFa, testMatrixFb)

	if errF != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixF, resultMatrixF) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsGa := [][]complex128{{0, 1}, {1, 0}}
	testMatrixGa := ct.MakeComplexMatrixWithElements(testElementsGa)

	testElementsGb := [][]complex128{{1, 0}, {0, 1}}
	testMatrixGb := ct.MakeComplexMatrixWithElements(testElementsGb)

	solutionElementsG := [][]complex128{{0, 1}, {1, 0}}
	solutionMatrixG := ct.MakeComplexMatrixWithElements(solutionElementsG)

	resultMatrixG, errG := MatrixComplexMultiSimple(testMatrixGa, testMatrixGb)

	if errG != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixG, resultMatrixG) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixG, resultMatrixG))
	}

	testElementsHa := [][]complex128{{1, 0}, {0, 1}, {0, 1}}
	testMatrixHa := ct.MakeComplexMatrixWithElements(testElementsHa)

	testElementsHb := [][]complex128{{0, 1}, {1, 0}, {0, 1}}
	testMatrixHb := ct.MakeComplexMatrixWithElements(testElementsHb)

	_, errH := MatrixComplexMultiSimple(testMatrixHa, testMatrixHb)

	if errH == nil {
		t.Error("Expected error")
	}
}

func TestMatrixAddition(t *testing.T) {
	testElementsAa := [][]float64{{1, 0}, {0, 1}}
	testMatrixAa := ct.MakeMatrixWithElements(testElementsAa)

	testElementsAb := [][]float64{{0, 1}, {1, 0}}
	testMatrixAb := ct.MakeMatrixWithElements(testElementsAb)

	solutionElementsA := [][]float64{{1, 1}, {1, 1}}
	solutionMatrixA := ct.MakeMatrixWithElements(solutionElementsA)

	resultMatrixA, errA := MatrixAddition(testMatrixAa, testMatrixAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsBa := [][]float64{{2, 0}, {0, 2}, {2, 0}}
	testMatrixBa := ct.MakeMatrixWithElements(testElementsBa)

	testElementsBb := [][]float64{{0, 1}, {1, 0}, {1, 0}}
	testMatrixBb := ct.MakeMatrixWithElements(testElementsBb)

	solutionElementsB := [][]float64{{2, 1}, {1, 2}, {3, 0}}
	solutionMatrixB := ct.MakeMatrixWithElements(solutionElementsB)

	resultMatrixB, errB := MatrixAddition(testMatrixBa, testMatrixBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsCa := [][]float64{{0, 1, 0}, {1, 0, 1}}
	testMatrixCa := ct.MakeMatrixWithElements(testElementsCa)

	testElementsCb := [][]float64{{1, 0, 1}, {0, 1, 0}}
	testMatrixCb := ct.MakeMatrixWithElements(testElementsCb)

	solutionElementsC := [][]float64{{1, 1, 1}, {1, 1, 1}}
	solutionMatrixC := ct.MakeMatrixWithElements(solutionElementsC)

	resultMatrixC, errC := MatrixAddition(testMatrixCa, testMatrixCb)

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixC, resultMatrixC) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixC, resultMatrixC))
	}

	testElementsDa := [][]float64{{1, 0}, {0, 1}, {1, 0}}
	testMatrixDa := ct.MakeMatrixWithElements(testElementsDa)

	testElementsDb := [][]float64{{0, 1, 1}, {1, 0, 1}}
	testMatrixDb := ct.MakeMatrixWithElements(testElementsDb)

	_, errD := MatrixAddition(testMatrixDa, testMatrixDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := [][]complex128{{1, 0, 1}, {0, 1, 1}}
	testMatrixEa := ct.MakeComplexMatrixWithElements(testElementsEa)

	testElementsEb := [][]complex128{{0, 1 + 1i, 2 + 2i}, {1, 0, 0}}
	testMatrixEb := ct.MakeComplexMatrixWithElements(testElementsEb)

	solutionElementsE := [][]complex128{{1, 1 + 1i, 3 + 2i}, {1, 1, 1}}
	solutionMatrixE := ct.MakeComplexMatrixWithElements(solutionElementsE)

	resultMatrixE, errE := MatrixComplexAddition(testMatrixEa, testMatrixEb)

	if errE != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixE, resultMatrixE) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixE, resultMatrixE))
	}

	testElementsFa := [][]complex128{{2, 0}, {0, 2}}
	testMatrixFa := ct.MakeComplexMatrixWithElements(testElementsFa)

	testElementsFb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixFb := ct.MakeComplexMatrixWithElements(testElementsFb)

	solutionElementsF := [][]complex128{{2, 1}, {1, 2}}
	solutionMatrixF := ct.MakeComplexMatrixWithElements(solutionElementsF)

	resultMatrixF, errF := MatrixComplexAddition(testMatrixFa, testMatrixFb)

	if errF != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixF, resultMatrixF) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsGa := [][]complex128{{0, 1}, {1, 0}, {1 + 5i, 5 - 4i}}
	testMatrixGa := ct.MakeComplexMatrixWithElements(testElementsGa)

	testElementsGb := [][]complex128{{1, 0}, {0, 1}, {0, 0}}
	testMatrixGb := ct.MakeComplexMatrixWithElements(testElementsGb)

	solutionElementsG := [][]complex128{{1, 1}, {1, 1}, {1 + 5i, 5 - 4i}}
	solutionMatrixG := ct.MakeComplexMatrixWithElements(solutionElementsG)

	resultMatrixG, errG := MatrixComplexAddition(testMatrixGa, testMatrixGb)

	if errG != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixG, resultMatrixG) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixG, resultMatrixG))
	}

	testElementsHa := [][]complex128{{1, 0}, {0, 1}, {0, 1}}
	testMatrixHa := ct.MakeComplexMatrixWithElements(testElementsHa)

	testElementsHb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixHb := ct.MakeComplexMatrixWithElements(testElementsHb)

	_, errH := MatrixComplexAddition(testMatrixHa, testMatrixHb)

	if errH == nil {
		t.Error("Expected error")
	}
}

func TestMatrixSubtraction(t *testing.T) {
	testElementsAa := [][]float64{{1, 0}, {0, 1}}
	testMatrixAa := ct.MakeMatrixWithElements(testElementsAa)

	testElementsAb := [][]float64{{0, 1}, {1, 0}}
	testMatrixAb := ct.MakeMatrixWithElements(testElementsAb)

	solutionElementsA := [][]float64{{1, -1}, {-1, 1}}
	solutionMatrixA := ct.MakeMatrixWithElements(solutionElementsA)

	resultMatrixA, errA := MatrixSubtraction(testMatrixAa, testMatrixAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsBa := [][]float64{{2, 0}, {0, 2}, {2, 0}}
	testMatrixBa := ct.MakeMatrixWithElements(testElementsBa)

	testElementsBb := [][]float64{{0, 1}, {1, 0}, {1, 0}}
	testMatrixBb := ct.MakeMatrixWithElements(testElementsBb)

	solutionElementsB := [][]float64{{2, -1}, {-1, 2}, {1, 0}}
	solutionMatrixB := ct.MakeMatrixWithElements(solutionElementsB)

	resultMatrixB, errB := MatrixSubtraction(testMatrixBa, testMatrixBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsCa := [][]float64{{0, 1, 0}, {1, 0, 1}}
	testMatrixCa := ct.MakeMatrixWithElements(testElementsCa)

	testElementsCb := [][]float64{{1, 0, 1}, {0, 1, 0}}
	testMatrixCb := ct.MakeMatrixWithElements(testElementsCb)

	solutionElementsC := [][]float64{{-1, 1, -1}, {1, -1, 1}}
	solutionMatrixC := ct.MakeMatrixWithElements(solutionElementsC)

	resultMatrixC, errC := MatrixSubtraction(testMatrixCa, testMatrixCb)

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixC, resultMatrixC) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixC, resultMatrixC))
	}

	testElementsDa := [][]float64{{1, 0}, {0, 1}, {1, 0}}
	testMatrixDa := ct.MakeMatrixWithElements(testElementsDa)

	testElementsDb := [][]float64{{0, 1, 1}, {1, 0, 1}}
	testMatrixDb := ct.MakeMatrixWithElements(testElementsDb)

	_, errD := MatrixSubtraction(testMatrixDa, testMatrixDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := [][]complex128{{1, 0, 1}, {0, 1, 1}}
	testMatrixEa := ct.MakeComplexMatrixWithElements(testElementsEa)

	testElementsEb := [][]complex128{{0, 1 + 1i, 2 + 2i}, {1, 0, 0}}
	testMatrixEb := ct.MakeComplexMatrixWithElements(testElementsEb)

	solutionElementsE := [][]complex128{{1, -1 - 1i, -1 - 2i}, {-1, 1, 1}}
	solutionMatrixE := ct.MakeComplexMatrixWithElements(solutionElementsE)

	resultMatrixE, errE := MatrixComplexSubtraction(testMatrixEa, testMatrixEb)

	if errE != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixE, resultMatrixE) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixE, resultMatrixE))
	}

	testElementsFa := [][]complex128{{2, 0}, {0, 2}}
	testMatrixFa := ct.MakeComplexMatrixWithElements(testElementsFa)

	testElementsFb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixFb := ct.MakeComplexMatrixWithElements(testElementsFb)

	solutionElementsF := [][]complex128{{2, -1}, {-1, 2}}
	solutionMatrixF := ct.MakeComplexMatrixWithElements(solutionElementsF)

	resultMatrixF, errF := MatrixComplexSubtraction(testMatrixFa, testMatrixFb)

	if errF != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixF, resultMatrixF) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsGa := [][]complex128{{0, 1}, {1, 0}, {1 + 5i, 5 - 4i}}
	testMatrixGa := ct.MakeComplexMatrixWithElements(testElementsGa)

	testElementsGb := [][]complex128{{1, 0}, {0, 1}, {0, 0}}
	testMatrixGb := ct.MakeComplexMatrixWithElements(testElementsGb)

	solutionElementsG := [][]complex128{{-1, 1}, {1, -1}, {1 + 5i, 5 - 4i}}
	solutionMatrixG := ct.MakeComplexMatrixWithElements(solutionElementsG)

	resultMatrixG, errG := MatrixComplexSubtraction(testMatrixGa, testMatrixGb)

	if errG != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixG, resultMatrixG) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixG, resultMatrixG))
	}

	testElementsHa := [][]complex128{{1, 0}, {0, 1}, {0, 1}}
	testMatrixHa := ct.MakeComplexMatrixWithElements(testElementsHa)

	testElementsHb := [][]complex128{{0, 1}, {1, 0}}
	testMatrixHb := ct.MakeComplexMatrixWithElements(testElementsHb)

	_, errH := MatrixComplexSubtraction(testMatrixHa, testMatrixHb)

	if errH == nil {
		t.Error("Expected error")
	}
}
