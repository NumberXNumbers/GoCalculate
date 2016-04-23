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

	intSolution := 1
	floatSolution := 1.0
	complexSolutionA := 1 + 0i
	complexSolutionB := 1 - 0i

	if errA != nil {
		t.Fail()
	}

	if valueA.Int() != intSolution {
		t.Errorf("Expecte %d, received %d", intSolution, valueA.Int())
	}

	if errB != nil {
		t.Fail()
	}

	if valueB.Float64() != floatSolution {
		t.Errorf("Expecte %f, received %f", floatSolution, valueB.Float64())
	}

	if errC != nil {
		t.Fail()
	}

	if valueC.Complex128() != complexSolutionA {
		t.Errorf("Expecte %f, received %f", complexSolutionA, valueB.Complex128())
	}

	if errD != nil {
		t.Fail()
	}

	if valueD.Complex128() != complexSolutionA {
		t.Errorf("Expecte %f, received %f", complexSolutionA, valueD.Complex128())
	}

	if errE != nil {
		t.Fail()
	}

	if valueE.Complex128() != complexSolutionB {
		t.Errorf("Expecte %f, received %f", complexSolutionB, valueE.Complex128())
	}

	if errF != nil {
		t.Fail()
	}

	if valueF.Complex128() != complexSolutionB {
		t.Errorf("Expecte %f, received %f", complexSolutionB, valueF.Complex128())
	}

	if errG == nil {
		t.Error("Expected Error")
	}

}
