package calculators

import "testing"

func TestReversePolishCalculator(t *testing.T) {
	testArgsA := []string{"4", "3", "+"}
	testArgsB := []string{"4", "3", "2", "-", "/"}
	testArgsC := []string{"11", "3", "2", "1", "*", "exp", "%"}
	testArgsD := []string{"3", "2", "&"}
	testArgsE := []string{"1"}
	testArgsF := []string{"1", "2", "3", "+"}
	testArgsG := []string{"1", "+", "+"}

	testValueA, errA := ReversePolishCalculator(testArgsA)
	testValueB, errB := ReversePolishCalculator(testArgsB)
	testValueC, errC := ReversePolishCalculator(testArgsC)
	_, errD := ReversePolishCalculator(testArgsD)
	_, errE := ReversePolishCalculator(testArgsE)
	_, errF := ReversePolishCalculator(testArgsF)
	_, errG := ReversePolishCalculator(testArgsG)

	if errA != nil {
		t.Errorf("Expect there to be no error, received %s", errA)
	}

	if testValueA != 7.0 {
		t.Errorf("Expect %d, received %v", 7, testValueA)
	}

	if errB != nil {
		t.Errorf("Expect there to be no error, received %s", errB)
	}

	if testValueB != 4.0 {
		t.Errorf("Expect %d, received %v", 4, testValueB)
	}

	if errC != nil {
		t.Errorf("Expect there to be no error, received %s", errC)
	}

	if testValueC != 2.0 {
		t.Errorf("Expect %d, received %v", 2, testValueC)
	}

	if errD == nil {
		t.Error("Expect there to be no error")
	}

	if errE == nil {
		t.Error("Expect there to be no error")
	}

	if errF == nil {
		t.Error("Expect there to be no error")
	}

	if errG == nil {
		t.Error("Expect there to be no error")
	}
}
