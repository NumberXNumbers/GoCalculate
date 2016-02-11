package vops

import (
	"math"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestVectorScalarMulti(t *testing.T) {
	testElementsA := []float64{1, 2, 3}
	testVectorA := v.MakeVectorWithElements(testElementsA, ct.RowVector)

	testScalarA := 2.0

	solutionElementsA := []float64{2, 4, 6}
	solutionVectorA := v.MakeVectorWithElements(solutionElementsA, ct.RowVector)

	resultVectorA := VectorScalarMulti(testScalarA, testVectorA)

	if !reflect.DeepEqual(resultVectorA, solutionVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorA, solutionVectorA))
	}

	testElementsB := []complex128{1, 2, 3}
	testVectorB := v.MakeComplexVectorWithElements(testElementsB, ct.RowVector)

	testScalarB := 2.0 + 1i

	solutionElementsB := []complex128{2 + 1i, 4 + 2i, 6 + 3i}
	solutionVectorB := v.MakeComplexVectorWithElements(solutionElementsB, ct.RowVector)

	resultVectorB := VectorComplexScalarMulti(testScalarB, testVectorB)

	if !reflect.DeepEqual(resultVectorB, solutionVectorB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorB, solutionVectorB))
	}
}

func TestVectorAddition(t *testing.T) {
	testElementsAa := []float64{1, 2, 3}
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ct.RowVector)

	testElementsAb := []float64{1, 2, 3}
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ct.RowVector)

	solutionElementsA := []float64{2, 4, 6}
	solutionVectorA := v.MakeVectorWithElements(solutionElementsA, ct.RowVector)

	resultVectorA, errA := VectorAddition(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA, solutionVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorA, solutionVectorA))
	}

	testElementsBa := []complex128{1, 2, 3}
	testVectorBa := v.MakeComplexVectorWithElements(testElementsBa, ct.RowVector)

	testElementsBb := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
	testVectorBb := v.MakeComplexVectorWithElements(testElementsBb, ct.RowVector)

	solutionElementsB := []complex128{2 + 1i, 4 + 2i, 6 + 3i}
	solutionVectorB := v.MakeComplexVectorWithElements(solutionElementsB, ct.RowVector)

	resultVectorB, errB := VectorComplexAddition(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB, solutionVectorB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorB, solutionVectorB))
	}

	testElementsCa := []float64{1, 2, 3}
	testVectorCa := v.MakeVectorWithElements(testElementsCa, ct.RowVector)

	testElementsCb := []float64{2, 4, 6}
	testVectorCb := v.MakeVectorWithElements(testElementsCb, ct.ColVector)

	_, errC := VectorAddition(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := []complex128{1, 2, 3}
	testVectorDa := v.MakeComplexVectorWithElements(testElementsDa, ct.ColVector)

	testElementsDb := []complex128{2 + 1i, 4 + 2i, 6 + 3i}
	testVectorDb := v.MakeComplexVectorWithElements(testElementsDb, ct.RowVector)

	_, errD := VectorComplexAddition(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := []float64{1, 2}
	testVectorEa := v.MakeVectorWithElements(testElementsEa, ct.RowVector)

	testElementsEb := []float64{2, 4, 6}
	testVectorEb := v.MakeVectorWithElements(testElementsEb, ct.RowVector)

	_, errE := VectorAddition(testVectorEa, testVectorEb)

	if errE == nil {
		t.Error("Expected error")
	}

	testElementsFa := []complex128{1, 2, 3}
	testVectorFa := v.MakeComplexVectorWithElements(testElementsFa, ct.ColVector)

	testElementsFb := []complex128{2 + 1i, 6 + 3i}
	testVectorFb := v.MakeComplexVectorWithElements(testElementsFb, ct.ColVector)

	_, errF := VectorComplexAddition(testVectorFa, testVectorFb)

	if errF == nil {
		t.Error("Expected error")
	}
}

func TestVectorSubtraction(t *testing.T) {
	testElementsAa := []float64{1, 2, 3}
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ct.RowVector)

	testElementsAb := []float64{1, 2, 3}
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ct.RowVector)

	solutionElementsA := []float64{0, 0, 0}
	solutionVectorA := v.MakeVectorWithElements(solutionElementsA, ct.RowVector)

	resultVectorA, errA := VectorSubtraction(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorA, solutionVectorA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorA, solutionVectorA))
	}

	testElementsBa := []complex128{1, 2, 3}
	testVectorBa := v.MakeComplexVectorWithElements(testElementsBa, ct.RowVector)

	testElementsBb := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
	testVectorBb := v.MakeComplexVectorWithElements(testElementsBb, ct.RowVector)

	solutionElementsB := []complex128{-1i, -2i, -3i}
	solutionVectorB := v.MakeComplexVectorWithElements(solutionElementsB, ct.RowVector)

	resultVectorB, errB := VectorComplexSubtraction(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(resultVectorB, solutionVectorB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(resultVectorB, solutionVectorB))
	}

	testElementsCa := []float64{1, 2, 3}
	testVectorCa := v.MakeVectorWithElements(testElementsCa, ct.RowVector)

	testElementsCb := []float64{2, 4, 6}
	testVectorCb := v.MakeVectorWithElements(testElementsCb, ct.ColVector)

	_, errC := VectorSubtraction(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := []complex128{1, 2, 3}
	testVectorDa := v.MakeComplexVectorWithElements(testElementsDa, ct.ColVector)

	testElementsDb := []complex128{2 + 1i, 4 + 2i, 6 + 3i}
	testVectorDb := v.MakeComplexVectorWithElements(testElementsDb, ct.RowVector)

	_, errD := VectorComplexSubtraction(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := []float64{1, 2}
	testVectorEa := v.MakeVectorWithElements(testElementsEa, ct.RowVector)

	testElementsEb := []float64{2, 4, 6}
	testVectorEb := v.MakeVectorWithElements(testElementsEb, ct.RowVector)

	_, errE := VectorSubtraction(testVectorEa, testVectorEb)

	if errE == nil {
		t.Error("Expected error")
	}

	testElementsFa := []complex128{1, 2, 3}
	testVectorFa := v.MakeComplexVectorWithElements(testElementsFa, ct.ColVector)

	testElementsFb := []complex128{2 + 1i, 6 + 3i}
	testVectorFb := v.MakeComplexVectorWithElements(testElementsFb, ct.ColVector)

	_, errF := VectorComplexSubtraction(testVectorFa, testVectorFb)

	if errF == nil {
		t.Error("Expected error")
	}
}

func TestInnerProduct(t *testing.T) {
	testElementsAa := []float64{1, 2, 3}
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ct.RowVector)

	testElementsAb := []float64{1, 2, 3}
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ct.ColVector)

	solutionA := float64(14)

	resultA, errA := InnerProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if resultA != solutionA {
		t.Errorf("Expected %v, recieved %v", solutionA, resultA)
	}

	testElementsBa := []complex128{1, 2, 3}
	testVectorBa := v.MakeComplexVectorWithElements(testElementsBa, ct.RowVector)

	testElementsBb := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
	testVectorBb := v.MakeComplexVectorWithElements(testElementsBb, ct.ColVector)

	solutionB := complex128(14 + 14i)

	resultB, errB := InnerProductComplex(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if resultB != solutionB {
		t.Errorf("Expected %v, recieved %v", solutionB, resultB)
	}

	testElementsCa := []float64{1, 2, 3}
	testVectorCa := v.MakeVectorWithElements(testElementsCa, ct.ColVector)

	testElementsCb := []float64{2, 4, 6}
	testVectorCb := v.MakeVectorWithElements(testElementsCb, ct.RowVector)

	_, errC := InnerProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := []complex128{1, 2, 3}
	testVectorDa := v.MakeComplexVectorWithElements(testElementsDa, ct.ColVector)

	testElementsDb := []complex128{2 + 1i, 4 + 2i, 6 + 3i}
	testVectorDb := v.MakeComplexVectorWithElements(testElementsDb, ct.RowVector)

	_, errD := InnerProductComplex(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := []float64{1, 2}
	testVectorEa := v.MakeVectorWithElements(testElementsEa, ct.RowVector)

	testElementsEb := []float64{2, 4, 6}
	testVectorEb := v.MakeVectorWithElements(testElementsEb, ct.RowVector)

	_, errE := InnerProduct(testVectorEa, testVectorEb)

	if errE == nil {
		t.Error("Expected error")
	}

	testElementsFa := []complex128{1, 2, 3}
	testVectorFa := v.MakeComplexVectorWithElements(testElementsFa, ct.ColVector)

	testElementsFb := []complex128{2 + 1i, 6 + 3i}
	testVectorFb := v.MakeComplexVectorWithElements(testElementsFb, ct.ColVector)

	_, errF := InnerProductComplex(testVectorFa, testVectorFb)

	if errF == nil {
		t.Error("Expected error")
	}
}

func TestAngleTheta(t *testing.T) {
	testElementsAa := []float64{1, 0}
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ct.RowVector)

	testElementsAb := []float64{0, 1}
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ct.ColVector)

	solutionA := float64(math.Pi / float64(2))

	resultA, errA := AngleTheta(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if resultA != solutionA {
		t.Errorf("Expected %v, recieved %v", solutionA, resultA)
	}

	testElementsBa := []complex128{1, 0}
	testVectorBa := v.MakeComplexVectorWithElements(testElementsBa, ct.RowVector)

	testElementsBb := []complex128{0, 1}
	testVectorBb := v.MakeComplexVectorWithElements(testElementsBb, ct.ColVector)

	solutionB := complex128(math.Pi / complex128(2))

	resultB, errB := AngleThetaComplex(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if resultB != solutionB {
		t.Errorf("Expected %v, recieved %v", solutionB, resultB)
	}

	testElementsCa := []float64{0, 0}
	testVectorCa := v.MakeVectorWithElements(testElementsCa, ct.ColVector)

	testElementsCb := []float64{0, 0}
	testVectorCb := v.MakeVectorWithElements(testElementsCb, ct.RowVector)

	_, errC := AngleTheta(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := []complex128{0, 0}
	testVectorDa := v.MakeComplexVectorWithElements(testElementsDa, ct.ColVector)

	testElementsDb := []complex128{0, 0}
	testVectorDb := v.MakeComplexVectorWithElements(testElementsDb, ct.RowVector)

	_, errD := AngleThetaComplex(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}

	testElementsEa := []float64{1, 2}
	testVectorEa := v.MakeVectorWithElements(testElementsEa, ct.RowVector)

	testElementsEb := []float64{2, 4, 6}
	testVectorEb := v.MakeVectorWithElements(testElementsEb, ct.RowVector)

	_, errE := AngleTheta(testVectorEa, testVectorEb)

	if errE == nil {
		t.Error("Expected error")
	}

	testElementsFa := []complex128{1, 2, 3}
	testVectorFa := v.MakeComplexVectorWithElements(testElementsFa, ct.ColVector)

	testElementsFb := []complex128{2 + 1i, 6 + 3i}
	testVectorFb := v.MakeComplexVectorWithElements(testElementsFb, ct.ColVector)

	_, errF := AngleThetaComplex(testVectorFa, testVectorFb)

	if errF == nil {
		t.Error("Expected error")
	}
}

func TestOuterProduct(t *testing.T) {
	testElementsAa := []float64{1, 0}
	testVectorAa := v.MakeVectorWithElements(testElementsAa, ct.ColVector)

	testElementsAb := []float64{0, 1}
	testVectorAb := v.MakeVectorWithElements(testElementsAb, ct.RowVector)

	solutionElementsA := [][]float64{{0, 1}, {0, 0}}
	solutionMatrixA := m.MakeMatrixWithElements(solutionElementsA)

	resultMatrixA, errA := OuterProduct(testVectorAa, testVectorAb)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixA, resultMatrixA) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixA, resultMatrixA))
	}

	testElementsBa := []complex128{1, 0}
	testVectorBa := v.MakeComplexVectorWithElements(testElementsBa, ct.ColVector)

	testElementsBb := []complex128{0, 1}
	testVectorBb := v.MakeComplexVectorWithElements(testElementsBb, ct.RowVector)

	solutionElementsB := [][]complex128{{0, 1}, {0, 0}}
	solutionMatrixB := m.MakeComplexMatrixWithElements(solutionElementsB)

	resultMatrixB, errB := OuterProductComplex(testVectorBa, testVectorBb)

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(solutionMatrixB, resultMatrixB) {
		t.Errorf("Expected %v, recieved %v", true, reflect.DeepEqual(solutionMatrixB, resultMatrixB))
	}

	testElementsCa := []float64{1, 2}
	testVectorCa := v.MakeVectorWithElements(testElementsCa, ct.RowVector)

	testElementsCb := []float64{2, 4, 6}
	testVectorCb := v.MakeVectorWithElements(testElementsCb, ct.RowVector)

	_, errC := OuterProduct(testVectorCa, testVectorCb)

	if errC == nil {
		t.Error("Expected error")
	}

	testElementsDa := []complex128{1, 2, 3}
	testVectorDa := v.MakeComplexVectorWithElements(testElementsDa, ct.ColVector)

	testElementsDb := []complex128{2 + 1i, 6 + 3i}
	testVectorDb := v.MakeComplexVectorWithElements(testElementsDb, ct.ColVector)

	_, errD := OuterProductComplex(testVectorDa, testVectorDb)

	if errD == nil {
		t.Error("Expected error")
	}
}
