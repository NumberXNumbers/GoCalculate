package utils

import "testing"

func TestStringToValueParser(t *testing.T) {
	testStrA := "1"
	testStrB := "1.0"
	testStrC := "1+0i"
	testStrD := "1+0.0i"
	testStrE := "1.0-0i"
	testStrF := "1.0-0.0i"
	testStrG := "bad"

	valueA, errA := StringToValueParser(testStrA)
	valueB, errB := StringToValueParser(testStrB)
	valueC, errC := StringToValueParser(testStrC)
	valueD, errD := StringToValueParser(testStrD)
	valueE, errE := StringToValueParser(testStrE)
	valueF, errF := StringToValueParser(testStrF)
	_, errG := StringToValueParser(testStrG)

	realSolution := 1.0
	complexSolutionA := 1 + 0i
	complexSolutionB := 1 - 0i

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

}
