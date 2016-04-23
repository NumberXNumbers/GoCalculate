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
	testArgsH := []string{"3.4", "4.4", "2.3", "3.3", "1.0", "+", "-", "*", "/"}
	testArgsI := []string{"3.4+3i", "4.4-3i", "2.3+1i", "4+2i", "7.2+2i", "+", "-", "*", "/"}
	testArgsJ := []string{"3.0", "2.0", "1.0", "exp", "%"}
	testArgsK := []string{"3.0", "3.0", "&"}
	testArgsL := []string{"3+0i", "2-0i", "exp"}
	testArgsM := []string{"3+0i", "2-0i", "%"}

	testValueA, errA := ReversePolishCalculator(testArgsA)
	testValueB, errB := ReversePolishCalculator(testArgsB)
	testValueC, errC := ReversePolishCalculator(testArgsC)
	_, errD := ReversePolishCalculator(testArgsD)
	_, errE := ReversePolishCalculator(testArgsE)
	_, errF := ReversePolishCalculator(testArgsF)
	_, errG := ReversePolishCalculator(testArgsG)
	testValueH, errH := ReversePolishCalculator(testArgsH)
	testValueI, errI := ReversePolishCalculator(testArgsI)
	testValueJ, errJ := ReversePolishCalculator(testArgsJ)
	_, errK := ReversePolishCalculator(testArgsK)
	testValueL, errL := ReversePolishCalculator(testArgsL)
	_, errM := ReversePolishCalculator(testArgsM)

	if errA != nil {
		t.Errorf("Expect there to be no error, received %s", errA)
	}

	if testValueA.Float64() != 7.0 {
		t.Errorf("Expect %d, received %v", 7, testValueA)
	}

	if errB != nil {
		t.Errorf("Expect there to be no error, received %s", errB)
	}

	if testValueB.Float64() != 4.0 {
		t.Errorf("Expect %d, received %v", 4, testValueB)
	}

	if errC != nil {
		t.Errorf("Expect there to be no error, received %s", errC)
	}

	if testValueC.Float64() != 2.0 {
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

	if errH != nil {
		t.Errorf("Expect there to be no error, received %s", errH)
	}

	if testValueH.Float64() != -0.3863636363636363 {
		t.Errorf("Expect %d, received %v", 7, testValueH)
	}

	if errI != nil {
		t.Errorf("Expect there to be no error, received %s", errI)
	}

	if testValueI.Complex128() != -0.08883947766013564+0.01806018430502029i {
		t.Errorf("Expect %d, received %v", 7, testValueI)
	}

	if errJ != nil {
		t.Errorf("Expect there to be no error, received %s", errJ)
	}

	if testValueJ.Float64() != 1 {
		t.Errorf("Expect %d, received %v", 7, testValueJ)
	}

	if errK == nil {
		t.Error("Expect there to be no error")
	}

	if errL != nil {
		t.Errorf("Expect there to be no error, received %s", errL)
	}

	if testValueL.Complex128() != 9+0i {
		t.Errorf("Expect %d, received %v", 7, testValueL)
	}

	if errM == nil {
		t.Error("Expect there to be no error")
	}
}

func TestPolishCalculator(t *testing.T) {
	testArgsA := []string{"+", "4", "3"}
	testArgsB := []string{"/", "4", "-", "3", "2"}
	testArgsC := []string{"%", "11", "exp", "3", "*", "2", "1"}
	testArgsD := []string{"&", "2", "3"}
	testArgsE := []string{"1"}
	testArgsF := []string{"+", "3", "2", "1"}
	testArgsG := []string{"+", "+", "1"}

	testValueA, errA := PolishCalculator(testArgsA)
	testValueB, errB := PolishCalculator(testArgsB)
	testValueC, errC := PolishCalculator(testArgsC)
	_, errD := PolishCalculator(testArgsD)
	_, errE := PolishCalculator(testArgsE)
	_, errF := PolishCalculator(testArgsF)
	_, errG := PolishCalculator(testArgsG)

	if errA != nil {
		t.Errorf("Expect there to be no error, received %s", errA)
	}

	if testValueA.Float64() != 7.0 {
		t.Errorf("Expect %d, received %v", 7, testValueA)
	}

	if errB != nil {
		t.Errorf("Expect there to be no error, received %s", errB)
	}

	if testValueB.Float64() != 4.0 {
		t.Errorf("Expect %d, received %v", 4, testValueB)
	}

	if errC != nil {
		t.Errorf("Expect there to be no error, received %s", errC)
	}

	if testValueC.Float64() != 2.0 {
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
