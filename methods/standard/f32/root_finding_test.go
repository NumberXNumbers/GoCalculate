package methods

import (
	"math"
	"testing"
)

func TestBisection1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Pow(float64(x), 3)) + 5*float32(math.Pow(float64(x), 2)) + x - 5
	}

	rootA, errA := Bisection1D(0, 2, float32(math.Pow(10, -4)), 100, testFunction)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(float64(10), -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := Bisection1D(0, 2, float32(math.Pow(10, -4)), 5, testFunction)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestFixedPointIteration1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Sqrt(float64((5 - x - float32(math.Pow(float64(x), 3))) / 5.0)))
	}

	rootA, errA := FixedPointIteration1D(0.7, float32(math.Pow(10, -4)), 100, testFunction)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := FixedPointIteration1D(0.7, float32(math.Pow(10, -4)), 5, testFunction)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestNewton1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Pow(float64(x), 3)) + 5*float32(math.Pow(float64(x), 2)) + x - 5
	}

	testFunctionD := func(x float32) float32 {
		return 3*float32(math.Pow(float64(x), 2)) + 10*x + 1
	}

	rootA, errA := Newton1D(0.7, float32(math.Pow(10, -4)), 5, testFunction, testFunctionD)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := Newton1D(0.7, float32(math.Pow(10, -4)), 2, testFunction, testFunctionD)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestModifiedNewton1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Pow(float64(x), 3)) + 5*float32(math.Pow(float64(x), 2)) + x - 5
	}

	testFunctionD := func(x float32) float32 {
		return 3*float32(math.Pow(float64(x), 2)) + 10*x + 1
	}

	testFunctionDD := func(x float32) float32 {
		return 6*x + 10
	}

	rootA, errA := ModifiedNewton1D(0.7, float32(math.Pow(10, -4)), 5, testFunction, testFunctionD, testFunctionDD)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := ModifiedNewton1D(0.7, float32(math.Pow(10, -4)), 2, testFunction, testFunctionD, testFunctionDD)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestSecant1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Pow(float64(x), 3)) + 5*float32(math.Pow(float64(x), 2)) + x - 5
	}

	rootA, errA := Secant1D(0.7, 0.75, float32(math.Pow(10, -4)), 5, testFunction)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := Secant1D(0.7, 0.75, float32(math.Pow(10, -4)), 2, testFunction)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestFalsePosition1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Pow(float64(x), 3)) + 5*float32(math.Pow(float64(x), 2)) + x - 5
	}

	rootA, errA := FalsePosition1D(0.7, 0.75, float32(math.Pow(10, -4)), 5, testFunction)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := FalsePosition1D(0.7, 0.75, float32(math.Pow(10, -4)), 2, testFunction)

	if errB == nil {
		t.Error("Expected error")
	}
}

func TestSteffensen1D(t *testing.T) {
	testFunction := func(x float32) float32 {
		return float32(math.Sqrt((5 - float64(x) - math.Pow(float64(x), 3)) / 5))
	}

	rootA, errA := Steffensen1D(0.7, float32(math.Pow(10, -4)), 5, testFunction)

	if errA != nil {
		t.Errorf("Unexpected error, %v", errA)
	}

	if math.Abs(float64(rootA-0.8434)) >= math.Pow(10, -4) {
		t.Errorf("Expected %v, received %v", math.Pow(10, -4), math.Abs(float64(rootA-0.8434)))
	}

	_, errB := Steffensen1D(0.7, float32(math.Pow(10, -4)), 2, testFunction)

	if errB == nil {
		t.Error("Expected error")
	}
}
