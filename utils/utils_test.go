package utils

import (
	"reflect"
	"testing"

	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func TestStringToValueParser(t *testing.T) {
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

	valueA, errA := StringToValueParser(testStrA)
	valueB, errB := StringToValueParser(testStrB)
	valueC, errC := StringToValueParser(testStrC)
	valueD, errD := StringToValueParser(testStrD)
	valueE, errE := StringToValueParser(testStrE)
	valueF, errF := StringToValueParser(testStrF)
	_, errG := StringToValueParser(testStrG)
	valueH, errH := StringToValueParser(testStrH)
	valueI, errI := StringToValueParser(testStrI)
	valueJ, errJ := StringToValueParser(testStrJ)

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

func TestStringToMatrixParser(t *testing.T) {
	testStrA := "[1 2 3: 4 5 6: 7 8 9]"
	testStrB := "[1.0 -5+4i 5.0: 4 3 2]"
	testStrC := "[2 3: 4 5: 6 6: 7 7]"
	testStrD := "[1000 123.0 2345.5 34.3: 3 4 5 6]"
	testStrE := "[1 2 3: 4 5 6: 7 8]"
	testStrFa := "1 2 3: 4 5 6: 7 8]"
	testStrFb := "[1 2 3: 4 5 6: 7 8"
	testStrGa := "[[1 2 3: [4 5 6: [7 8]"
	testStrGb := "[1 2 3] 4 5 6] 7 8]]"
	testStrH := "[1 2 3 4 5 6 7 8 9]"
	testStrI := "[1 2 3: 1 * 3]"

	valueA, errA := StringToMatrixParser(testStrA)
	valueB, errB := StringToMatrixParser(testStrB)
	valueC, errC := StringToMatrixParser(testStrC)
	valueD, errD := StringToMatrixParser(testStrD)
	_, errE := StringToMatrixParser(testStrE)
	_, errFa := StringToMatrixParser(testStrFa)
	_, errFb := StringToValueParser(testStrFb)
	_, errGa := StringToMatrixParser(testStrGa)
	_, errGb := StringToValueParser(testStrGb)
	_, errH := StringToMatrixParser(testStrH)
	_, errI := StringToMatrixParser(testStrI)

	solMatrixA := m.MakeMatrix(v.MakeVectorPure(v.RowSpace, 1, 2, 3), v.MakeVectorPure(v.RowSpace, 4, 5, 6), v.MakeVectorPure(v.RowSpace, 7, 8, 9))
	solMatrixB := m.MakeMatrix(v.MakeVectorPure(v.RowSpace, 1.0, -5+4i, 5.0), v.MakeVectorPure(v.RowSpace, 4, 3, 2))
	solMatrixC := m.MakeMatrix(v.MakeVectorPure(v.RowSpace, 2, 3), v.MakeVectorPure(v.RowSpace, 4, 5), v.MakeVectorPure(v.RowSpace, 6, 6), v.MakeVectorPure(v.RowSpace, 7, 7))
	solMatrixD := m.MakeMatrix(v.MakeVectorPure(v.RowSpace, 1000, 123.0, 2345.5, 34.3), v.MakeVectorPure(v.RowSpace, 3, 4, 5, 6))

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
}

func TestStringToVectorParser(t *testing.T) {
	testStrA := "[1 2 3 4 5 6 7 8 9]"
	testStrB := "[1.0 -5+4i 5.0 4 3 2]"
	testStrC := "[2 3 4 5 6 6 7 7]"
	testStrD := "[1000 123.0 2345.5 34.3 3 4 5 6]"
	testStrE := "[1 2 3: 4 5 6: 7 8]"
	testStrFa := "1 2 3 4 5 6 7 8]"
	testStrFb := "[1 2 3 4 5 6 7 8"
	testStrH := "[1 2 3 1 * 3]"

	valueA, errA := StringToVectorParser(testStrA)
	valueB, errB := StringToVectorParser(testStrB)
	valueC, errC := StringToVectorParser(testStrC)
	valueD, errD := StringToVectorParser(testStrD)
	_, errE := StringToVectorParser(testStrE)
	_, errFa := StringToVectorParser(testStrFa)
	_, errFb := StringToValueParser(testStrFb)
	_, errH := StringToVectorParser(testStrH)

	solVectorA := v.MakeVectorPure(v.RowSpace, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	solVectorB := v.MakeVectorPure(v.RowSpace, 1.0, -5+4i, 5.0, 4, 3, 2)
	solVectorC := v.MakeVectorPure(v.RowSpace, 2, 3, 4, 5, 6, 6, 7, 7)
	solVectorD := v.MakeVectorPure(v.RowSpace, 1000, 123.0, 2345.5, 34.3, 3, 4, 5, 6)

	if errA != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueA, solVectorA) {
		t.Errorf("Expected %v, received %v", solVectorA, valueA)
	}

	if errB != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueB, solVectorB) {
		t.Errorf("Expected %v, received %v", solVectorB, valueB)
	}

	if errC != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueC, solVectorC) {
		t.Errorf("Expected %v, received %v", solVectorC, valueC)
	}

	if errD != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(valueD, solVectorD) {
		t.Errorf("Expected %v, received %v", solVectorD, valueD)
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

	if errH == nil {
		t.Error("Expected Error")
	}
}
