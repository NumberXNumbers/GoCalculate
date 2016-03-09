package methods

import (
	"reflect"
	"testing"
)

func TestNewtonDividedDifference(t *testing.T) {
	testTableA, errA := NewtonDividedDifference([]float64{1, 1.3, 1.6, 1.9, 2.2}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionTable := [][]float64{{0.9153827}, {0.4873198, -1.426876333333333}, {0.8960778, 1.3625266666666664, 4.649004999999998}, {0.2769871, -2.063635666666668, -5.710270555555559, -11.51030617283951}, {0.7866039, 1.6987226666666655, 6.2705972222222215, 13.312075308641978, 20.68531790123457}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionTable, testTableA) {
		t.Errorf("Expected %v, received %v", solutionTable, testTableA)
	}

	_, errB := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestNewtonForwardDividedDifference(t *testing.T) {
	testSetA, errA := NewtonForwardDividedDifference([]float64{1, 1.3, 1.6, 1.9, 2.2}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := []float64{0.9153827, -1.426876333333333, 4.649004999999998, -11.51030617283951, 20.68531790123457}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestNewtonBackwardsDividedDifference(t *testing.T) {
	testSetA, errA := NewtonBackwardsDividedDifference([]float64{1, 1.3, 1.6, 1.9, 2.2}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := []float64{0.7866039, 1.6987226666666655, 6.2705972222222215, 13.312075308641978, 20.68531790123457}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestStrilingCenterDividedDifference(t *testing.T) {
	testSetA, errA := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9, 2.2}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := [][]float64{{0.8960778}, {1.3625266666666664, -2.063635666666668}, {-5.710270555555559}, {-11.51030617283951, 13.312075308641978}, {20.68531790123457}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	testSetB, errB := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9, 2.2, 2.5}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039, 0.0837184})

	if errB != nil {
		t.Errorf("Error %v", errB)
	}

	if !reflect.DeepEqual(solutionSet, testSetB) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetB)
	}

	_, errC := StirlingCenterDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errC == nil {
		t.Error("Expected Error")
	}
}
