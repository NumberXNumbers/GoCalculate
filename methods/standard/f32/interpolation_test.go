package methods

import (
	"math"
	"reflect"
	"testing"
)

func TestNewtonDividedDifference(t *testing.T) {
	testTableA, errA := NewtonDividedDifference([]float32{1, 1.3, 1.6, 1.9, 2.2}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionTable := [][]float32{{0.9153827}, {0.4873198, -1.4268765}, {0.8960778, 1.3625264, 4.649005}, {0.2769871, -2.0636358, -5.7102704, -11.510306}, {0.7866039, 1.6987225, 6.2705965, 13.312074, 20.685316}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionTable, testTableA) {
		t.Errorf("Expected %v, received %v", solutionTable, testTableA)
	}

	_, errB := NewtonDividedDifference([]float32{1, 1.3, 1.6, 1.9}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestNewtonForwardDividedDifference(t *testing.T) {
	testSetA, errA := NewtonForwardDividedDifference([]float32{1, 1.3, 1.6, 1.9, 2.2}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := []float32{0.9153827, -1.4268765, 4.649005, -11.510306, 20.685316}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := NewtonForwardDividedDifference([]float32{1, 1.3, 1.6, 1.9}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestNewtonBackwardsDividedDifference(t *testing.T) {
	testSetA, errA := NewtonBackwardsDividedDifference([]float32{1, 1.3, 1.6, 1.9, 2.2}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := []float32{0.7866039, 1.6987225, 6.2705965, 13.312074, 20.685316}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := NewtonBackwardsDividedDifference([]float32{1, 1.3, 1.6, 1.9}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestStrilingCenterDividedDifference(t *testing.T) {
	testSetA, errA := StirlingCenterDividedDifference([]float32{1, 1.3, 1.6, 1.9, 2.2}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := [][]float32{{0.8960778}, {1.3625264, -2.0636358}, {-5.7102704}, {-11.510306, 13.312074}, {20.685316}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	testSetB, errB := StirlingCenterDividedDifference([]float32{1, 1.3, 1.6, 1.9, 2.2, 2.5}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039, 0.0837184})

	if errB != nil {
		t.Errorf("Error %v", errB)
	}

	if !reflect.DeepEqual(solutionSet, testSetB) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetB)
	}

	_, errC := StirlingCenterDividedDifference([]float32{1, 1.3, 1.6, 1.9}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errC == nil {
		t.Error("Expected Error")
	}
}

// []float32{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623}
func TestNevilleIterated(t *testing.T) {
	testTableA, errA := NevilleIterated(1.5, []float32{1, 1.3, 1.6, 1.9, 2.2}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := [][]float32{{0.9153827}, {0.4873198, 0.20194444}, {0.8960778, 0.7598251, 0.66684496}, {0.2769871, 1.1024414, 0.87403053, 0.78194803}, {0.7866039, -0.40250182, 1.3532653, 0.98052716, 0.8646893}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := NevilleIterated(1.5, []float32{1, 1.3, 1.6, 1.9}, []float32{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestHermite(t *testing.T) {
	testSetA, errA := Hermite([]float32{1.3, 1.6, 1.9}, []float32{0.4873198, 0.8960778, 0.2769871}, []float32{-0.0293884, 1.3455501, -0.741541})

	solutionSet := []float32{0.4873198, -0.0293884, 4.639715, -15.654339, -5.3187847, 207.24077}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := Hermite([]float32{1.3, 1.6}, []float32{0.4873198, 0.8960778, 0.2769871}, []float32{-0.0293884, 1.3455501, -0.741541})

	if errB == nil {
		t.Error("Expected Error")
	}

	_, errC := Hermite([]float32{1.3, 1.6, 1.9}, []float32{0.4873198, 0.8960778, 0.2769871}, []float32{-0.0293884, 1.3455501})

	if errC == nil {
		t.Error("Expected Error")
	}
}

func TestNaturalCubicSpline(t *testing.T) {
	testTableA, errA := NaturalCubicSpline([]float32{0, 1, 2, 3, 4}, []float32{1, float32(math.Exp(1)), float32(math.Exp(2)), float32(math.Exp(3)), float32(math.Exp(4))})

	solutionSet := [][]float32{{1, 2.7182817, 7.389056, 20.085537}, {1.1111264, 2.9325924, 6.3256726, 23.866482}, {0, 1.8214658, 1.5716147, 15.969195}, {0.60715526, -0.0832837, 4.799194, -5.3230653}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := NaturalCubicSpline([]float32{0, 1, 2, 3}, []float32{1, float32(math.Exp(1)), float32(math.Exp(2)), float32(math.Exp(3)), float32(math.Exp(4))})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestClampedCubicSpline(t *testing.T) {
	testTableA, errA := ClampedCubicSpline([]float32{0, 1, 2, 3, 4}, []float32{1, float32(math.Exp(1)), float32(math.Exp(2)), float32(math.Exp(3)), float32(math.Exp(4))}, 1, float32(math.Exp(4)))

	solutionSet := [][]float32{{1, 2.7182817, 7.389056, 20.085537}, {1, 2.6987429, 7.372197, 19.91423}, {0.4561025, 1.2426403, 3.4308145, 9.111221}, {0.26217926, 0.7293914, 1.8934689, 5.4871583}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := ClampedCubicSpline([]float32{0, 1, 2, 3}, []float32{1, float32(math.Exp(1)), float32(math.Exp(2)), float32(math.Exp(3)), float32(math.Exp(4))}, 1, float32(math.Exp(4)))

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestBezierCurve(t *testing.T) {
	testSetA, errA := BezierCurve([][2]float32{{0, 0}, {1, 0}}, [][2]float32{{2, 1}}, [][2]float32{{0, 1}})

	solutionSet := [][][4]float32{{{0, 6, -12, 7}}, {{0, 3, -3, 0}}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := BezierCurve([][2]float32{{0, 0}, {1, 0}}, [][2]float32{{2, 1}, {4, 1}}, [][2]float32{{0, 1}, {0, 4}})

	if errB == nil {
		t.Error("Expected Error")
	}

	_, errC := BezierCurve([][2]float32{{0, 0}, {1, 0}}, [][2]float32{{2, 1}}, [][2]float32{{0, 1}, {0, 4}})

	if errC == nil {
		t.Error("Expected Error")
	}
}
