package parsers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestValue(t *testing.T) {
	testStrA := "1"
	testStrB := "1.0"
	testStrC := "1+0i"
	testStrD := "1+0.0i"
	testStrE := "1.0-0i"
	testStrF := "1.0-0.0i"
	testStrG := "bad"
	testStrH := "-5-3i"
	testStrI := "-5+3i"
	testStrJ := "5-3i"

	valueA, errA := Value(testStrA)
	valueB, errB := Value(testStrB)
	valueC, errC := Value(testStrC)
	valueD, errD := Value(testStrD)
	valueE, errE := Value(testStrE)
	valueF, errF := Value(testStrF)
	_, errG := Value(testStrG)
	valueH, errH := Value(testStrH)
	valueI, errI := Value(testStrI)
	valueJ, errJ := Value(testStrJ)

	realSolution := 1.0
	complexSolutionA := 1 + 0i
	complexSolutionB := 1 - 0i
	complexSolutionC := -5 - 3i
	complexSolutionD := -5 + 3i
	complexSolutionE := 5 - 3i

	if errA != nil {
		t.Fail()
	}

	if valueA.Real() != realSolution {
		t.Errorf("Expecte %f, received %f", realSolution, valueA.Real())
	}

	if errB != nil {
		t.Fail()
	}

	if valueB.Real() != realSolution {
		t.Errorf("Expecte %f, received %f", realSolution, valueB.Real())
	}

	if errC != nil {
		t.Fail()
	}

	if valueC.Complex() != complexSolutionA {
		t.Errorf("Expecte %f, received %f", complexSolutionA, valueB.Complex())
	}

	if errD != nil {
		t.Fail()
	}

	if valueD.Complex() != complexSolutionA {
		t.Errorf("Expecte %f, received %f", complexSolutionA, valueD.Complex())
	}

	if errE != nil {
		t.Fail()
	}

	if valueE.Complex() != complexSolutionB {
		t.Errorf("Expecte %f, received %f", complexSolutionB, valueE.Complex())
	}

	if errF != nil {
		t.Fail()
	}

	if valueF.Complex() != complexSolutionB {
		t.Errorf("Expecte %f, received %f", complexSolutionB, valueF.Complex())
	}

	if errG == nil {
		t.Error("Expected Error")
	}

	if errH != nil {
		t.Fail()
	}

	if valueH.Complex() != complexSolutionC {
		t.Errorf("Expecte %f, received %f", complexSolutionC, valueH.Complex())
	}

	if errI != nil {
		t.Fail()
	}

	if valueI.Complex() != complexSolutionD {
		t.Errorf("Expecte %f, received %f", complexSolutionD, valueI.Complex())
	}

	if errJ != nil {
		t.Fail()
	}

	if valueJ.Complex() != complexSolutionE {
		t.Errorf("Expecte %f, received %f", complexSolutionE, valueJ.Complex())
	}
}

func TestMatrix(t *testing.T) {
	testStrA := "[1 2 3: 4 5 6: 7 8 9]"
	testStrB := "[1.0 -5+4i 5.0: 4 3 2]*"
	testStrC := "[2 3: 4 5: 6 6: 7 7]'"
	testStrD := "[1000 123.0 2345.5 34.3: 3 4 5 6]"
	testStrE := "[1 2 3: 4 5 6: 7 8]"
	testStrFa := "1 2 3: 4 5 6: 7 8]"
	testStrFb := "[1 2 3: 4 5 6: 7 8"
	testStrGa := "[[1 2 3: [4 5 6: [7 8]"
	testStrGb := "[1 2 3] 4 5 6] 7 8]]"
	testStrH := "[1 2 3 4 5 6 7 8 9]"
	testStrI := "[1 2 3: 1 * 3]"
	testStrJ := "[1 2 3: 1 2 3]9*"
	testStrK := "[1 2 3: 1 2 3]9'"

	valueA, errA := Matrix(testStrA)
	valueB, errB := Matrix(testStrB)
	valueC, errC := Matrix(testStrC)
	valueD, errD := Matrix(testStrD)
	_, errE := Matrix(testStrE)
	_, errFa := Matrix(testStrFa)
	_, errFb := Value(testStrFb)
	_, errGa := Matrix(testStrGa)
	_, errGb := Value(testStrGb)
	_, errH := Matrix(testStrH)
	_, errI := Matrix(testStrI)
	_, errJ := Matrix(testStrJ)
	_, errK := Matrix(testStrK)

	solMatrixA := m.MakeMatrix(v.MakeVector(v.RowSpace, 1, 2, 3), v.MakeVector(v.RowSpace, 4, 5, 6), v.MakeVector(v.RowSpace, 7, 8, 9))
	solMatrixB := m.MakeMatrix(v.MakeVector(v.RowSpace, 1.0, 4), v.MakeVector(v.RowSpace, -5-4i, 3), v.MakeVector(v.RowSpace, 5, 2))
	solMatrixC := m.MakeMatrix(v.MakeVector(v.RowSpace, 2, 4, 6, 7), v.MakeVector(v.RowSpace, 3, 5, 6, 7))
	solMatrixD := m.MakeMatrix(v.MakeVector(v.RowSpace, 1000, 123.0, 2345.5, 34.3), v.MakeVector(v.RowSpace, 3, 4, 5, 6))

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueA, solMatrixA) {
		t.Errorf("Expected %v, received %v", solMatrixA, valueA)
	}

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueB, solMatrixB) {
		t.Errorf("Expected %v, received %v", solMatrixB, valueB)
	}

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueC, solMatrixC) {
		t.Errorf("Expected %v, received %v", solMatrixC, valueC)
	}

	if errD != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueD, solMatrixD) {
		t.Errorf("Expected %v, received %v", solMatrixD, valueD)
	}

	if errE == nil {
		t.Error("Expected Error")
	}

	if errFa == nil {
		t.Error("Expected Error")
	}

	if errFb == nil {
		t.Error("Expected Error")
	}

	if errGa == nil {
		t.Error("Expected Error")
	}

	if errGb == nil {
		t.Error("Expected Error")
	}

	if errH == nil {
		t.Error("Expected Error")
	}

	if errI == nil {
		t.Error("Expected Error")
	}

	fmt.Println(errJ)
	if errJ == nil {

		t.Error("Expected Error")
	}

	if errK == nil {
		t.Error("Expected Error")
	}
}

func TestVector(t *testing.T) {
	testStrA := "[1 2 3 4 5 6 7 8 9]"
	testStrB := "[1.0 -5+4i 5.0 4 3 2]*"
	testStrC := "[2 3 4 5 6 6 7 7]'"
	testStrE := "[1 2 3: 4 5 6: 7 8]"
	testStrFa := "1 2 3 4 5 6 7 8]"
	testStrFb := "[1 2 3 4 5 6 7 8"
	testStrG := "1.0 [ -5+4i 5.0 4 3 2]"
	testStrH := "[1 2 3 1 * 3]"
	testStrI := "[2 3 4 5 6 6 7 ] 7'"
	testStrJ := "[2 3 4 5 6 6 7 ] 7*"

	vectorA, errA := Vector(testStrA)
	vectorB, errB := Vector(testStrB)
	vectorC, errC := Vector(testStrC)
	_, errD := Vector(testStrE)
	_, errEa := Vector(testStrFa)
	_, errEb := Vector(testStrFb)
	_, errF := Vector(testStrG)
	_, errG := Vector(testStrH)
	_, errH := Vector(testStrI)
	_, errI := Vector(testStrJ)

	solVectorA := v.MakeVector(v.RowSpace, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	solVectorB := v.MakeVector(v.ColSpace, 1.0, -5-4i, 5.0, 4, 3, 2)
	solVectorC := v.MakeVector(v.ColSpace, 2, 3, 4, 5, 6, 6, 7, 7)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(vectorA, solVectorA) {
		t.Errorf("Expected %v, received %v", solVectorA, vectorA)
	}

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(vectorB, solVectorB) {
		t.Errorf("Expected %v, received %v", solVectorB, vectorB)
	}

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(vectorC, solVectorC) {
		t.Errorf("Expected %v, received %v", solVectorC, vectorC)
	}

	if errD == nil {
		t.Error("Expected Error")
	}

	if errEa == nil {
		t.Error("Expected Error")
	}

	if errEb == nil {
		t.Error("Expected Error")
	}

	if errF == nil {
		t.Error("Expected Error")
	}

	if errG == nil {
		t.Error("Expected Error")
	}

	if errH == nil {
		t.Error("Expected Error")
	}

	if errI == nil {
		t.Error("Expected Error")
	}
}
