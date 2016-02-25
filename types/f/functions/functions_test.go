package functions

import (
	"fmt"
	"testing"
)

func TestBinomialCoefficient(t *testing.T) {
	testResultA := BinomialCoefficient(5, 3)
	testResultB := BinomialCoefficient(5, 1)
	testResultC := BinomialCoefficient(5, 5)
	testResultD := BinomialCoefficient(2, 3)

	if testResultA != 10 {
		t.Errorf("Expected %v, received %v", 10, testResultA)
	}

	if testResultB != 5 {
		t.Errorf("Expected %v, received %v", 5, testResultB)
	}

	if testResultC != 1 {
		t.Errorf("Expected %v, received %v", 1, testResultC)
	}

	if testResultD != 0 {
		t.Errorf("Expected %v, received %v", 0, testResultD)
	}
}

func TestLegendrePolynomial2(t *testing.T) {
	testFunctionA, errA := LegendrePolynomial2(5)
	testFunctionB, errB := LegendrePolynomial2(3)
	testFunctionC, errC := LegendrePolynomial2(0)
	_, errD := LegendrePolynomial2(-1)

	if errA != nil {
		t.Error("No error expected")
	}

	if testFunctionA(2) != 185.75 {
		t.Errorf("Expected %v, received %v", 185.75, testFunctionA(2))
	}

	if errB != nil {
		t.Error("No error expected")
	}

	if testFunctionB(2) != 17 {
		t.Errorf("Expected %v, received %v", 17, testFunctionB(2))
	}

	if errC != nil {
		t.Error("No error expected")
	}

	if testFunctionC(2) != 1 {
		t.Errorf("Expected %v, received %v", 1, testFunctionC(2))
	}

	if errD == nil {
		t.Error("Error expected")
	}
}

func TestLegendrePolynomial(t *testing.T) {
	testFunctionA := LegendrePolynomial(5)
	testFunctionB := LegendrePolynomial(3)
	testFunctionC := LegendrePolynomial(0)

	if testFunctionA(2) != 185.75 {
		t.Errorf("Expected %v, received %v", 185.75, testFunctionA(2))
	}

	if testFunctionB(2) != 17 {
		t.Errorf("Expected %v, received %v", 17, testFunctionB(2))
	}

	if testFunctionC(2) != 1 {
		t.Errorf("Expected %v, received %v", 1, testFunctionC(2))
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	testFunctionD := LegendrePolynomial(-1)

	if testFunctionD != nil {
		t.Error("Expected Panic")
	}
}
