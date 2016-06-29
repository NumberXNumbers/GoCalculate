package calculators

import (
	"fmt"
	"testing"
)

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
	testArgsL := []string{"3+1i", "2-1i", "exp"}
	testArgsM := []string{"3+1i", "2-1i", "%"}

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

	if testValueA.Real() != 7.0 {
		t.Errorf("Expect %d, received %v", 7, testValueA.Real())
	}

	if errB != nil {
		t.Errorf("Expect there to be no error, received %s", errB)
	}

	if testValueB.Real() != 4.0 {
		t.Errorf("Expect %d, received %v", 4, testValueB.Real())
	}

	if errC != nil {
		t.Errorf("Expect there to be no error, received %s", errC)
	}

	if testValueC.Real() != 2.0 {
		t.Errorf("Expect %d, received %v", 2, testValueC.Real())
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

	if testValueH.Real() != -0.3863636363636363 {
		t.Errorf("Expect %v, received %v", -0.3863636363636363, testValueH.Real())
	}

	if errI != nil {
		t.Errorf("Expect there to be no error, received %s", errI)
	}

	if testValueI.Complex() != -0.049265368625230635-0.07610221088954763i {
		t.Errorf("Expect %v, received %v", -0.049265368625230635-0.07610221088954763i, testValueI.Complex())
	}

	if errJ != nil {
		t.Errorf("Expect there to be no error, received %s", errJ)
	}

	if testValueJ.Real() != 1 {
		t.Errorf("Expect %d, received %v", 1, testValueJ.Real())
	}

	if errK == nil {
		t.Error("Expect there to be no error")
	}

	if errL != nil {
		t.Errorf("Expect there to be no error, received %s", errL)
	}

	if testValueL.Complex() != 12.054709343346778-6.707996187723516i {
		t.Errorf("Expect %v, received %v", 12.054709343346778-6.707996187723516i, testValueL.Complex())
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

	if testValueA.Real() != 7.0 {
		t.Errorf("Expect %d, received %v", 7, testValueA)
	}

	if errB != nil {
		t.Errorf("Expect there to be no error, received %s", errB)
	}

	if testValueB.Real() != 4.0 {
		t.Errorf("Expect %d, received %v", 4, testValueB)
	}

	if errC != nil {
		t.Errorf("Expect there to be no error, received %s", errC)
	}

	if testValueC.Real() != 2.0 {
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

func TestInfixCalculator(t *testing.T) {
	testArgsA := []string{"4", "+", "3"}
	testArgsB := []string{"5", "+", "1", "*", "(", "2", "-", "1", ")"}
	testArgsD := []string{"5", "*", "1", "+", "(", "2", "-", "(", "1", ")", ")"}

	testValueA := InfixCalculator(testArgsA)
	testValueB := InfixCalculator(testArgsB)
	testValueD := InfixCalculator(testArgsD)

	if testValueA.Value().Real() != 7 {
		t.Errorf("Expect %d, received %v", 7, testValueA.Value().Real())
	}

	if testValueB.Value().Real() != 6 {
		t.Errorf("Expect %d, received %v", 6, testValueB.Value().Real())
	}

	if testValueD.Value().Real() != 6 {
		t.Errorf("Expect %d, received %v", 6, testValueD.Value().Real())
	}
}

func TestPanicInfixC(t *testing.T) {
	testArgsC := []string{"5", "+", "-"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("InfixCalculator Error: %v\n", r)
		}
	}()

	errC := InfixCalculator(testArgsC)

	if errC != nil {
		t.Error("Expected panic")
	}
}

func TestPanicInfixE(t *testing.T) {
	testArgsE := []string{"5", "*", "(", "2", "&", "1", ")"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("InfixCalculator Error: %v\n", r)
		}
	}()

	errE := InfixCalculator(testArgsE)

	if errE != nil {
		t.Error("Expected panic")
	}
}

func TestPanicInfixF(t *testing.T) {
	testArgsF := []string{"5", "+", "2", "%", "10+2i"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("InfixCalculator Error: %v\n", r)
		}
	}()

	errF := InfixCalculator(testArgsF)

	if errF != nil {
		t.Error("Expected panic")
	}
}

func TestPanicInfixG(t *testing.T) {
	testArgsG := []string{"5", "%", "10+2i", "+", "2"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("InfixCalculator Error: %v\n", r)
		}
	}()

	errG := InfixCalculator(testArgsG)

	if errG != nil {
		t.Error("Expected panic")
	}
}
