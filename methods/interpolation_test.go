package methods

import (
	"math"
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

	_, errB := NewtonDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

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

	_, errB := NewtonForwardDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

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

	_, errB := NewtonBackwardsDividedDifference([]float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

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

// []float64{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623}
func TestNevilleIterated(t *testing.T) {
	testTableA, errA := NevilleIterated(1.5, []float64{1, 1.3, 1.6, 1.9, 2.2}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	solutionSet := [][]float64{{0.9153827}, {0.4873198, 0.2019445333333335}, {0.8960778, 0.7598251333333332, 0.6668450333333332}, {0.2769871, 1.102441366666667, 0.8740305444444444, 0.7819480950617284}, {0.7866039, -0.40250196666666593, 1.3532652555555562, 0.9805271469135803, 0.8646893666666667}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := NevilleIterated(1.5, []float64{1, 1.3, 1.6, 1.9}, []float64{0.9153827, 0.4873198, 0.8960778, 0.2769871, 0.7866039})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestHermite(t *testing.T) {
	testSetA, errA := Hermite([]float64{1.3, 1.6, 1.9}, []float64{0.4873198, 0.8960778, 0.2769871}, []float64{-0.0293884, 1.3455501, -0.741541})

	solutionSet := []float64{0.4873198, -0.0293884, 4.639716888888888, -15.65435148148147, -5.318758641975368, 207.24067901234616}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := Hermite([]float64{1.3, 1.6}, []float64{0.4873198, 0.8960778, 0.2769871}, []float64{-0.0293884, 1.3455501, -0.741541})

	if errB == nil {
		t.Error("Expected Error")
	}

	_, errC := Hermite([]float64{1.3, 1.6, 1.9}, []float64{0.4873198, 0.8960778, 0.2769871}, []float64{-0.0293884, 1.3455501})

	if errC == nil {
		t.Error("Expected Error")
	}
}

func TestNaturalCubicSpline(t *testing.T) {
	testTableA, errA := NaturalCubicSpline([]float64{0, 1, 2, 3, 4}, []float64{1, math.Exp(1), math.Exp(2), math.Exp(3), math.Exp(4)})

	solutionSet := [][]float64{{1, 2.718281828459045, 7.38905609893065, 20.085536923187668}, {1.1111266016600037, 2.9325922820571275, 6.325672566903437, 23.86648273451499}, {0, 1.8214656803971239, 1.571614604449186, 15.969195563162367}, {0.6071552267990413, -0.08328369198264592, 4.7991936529043935, -5.323065187720789}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := NaturalCubicSpline([]float64{0, 1, 2, 3}, []float64{1, math.Exp(1), math.Exp(2), math.Exp(3), math.Exp(4)})

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestClampedCubicSpline(t *testing.T) {
	testTableA, errA := ClampedCubicSpline([]float64{0, 1, 2, 3, 4}, []float64{1, math.Exp(1), math.Exp(2), math.Exp(3), math.Exp(4)}, 1, math.Exp(4))

	solutionSet := [][]float64{{1, 2.718281828459045, 7.38905609893065, 20.085536923187668}, {0.9999999999999999, 2.698742769368434, 7.372197219318214, 19.914233637544577}, {0.4561027160087011, 1.2426400533597335, 3.430814396590047, 9.111222021636316}, {0.26217911245034414, 0.7293914477434379, 1.8934692083487563, 5.487157450775675}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testTableA) {
		t.Errorf("Expected %v, received %v", solutionSet, testTableA)
	}

	_, errB := ClampedCubicSpline([]float64{0, 1, 2, 3}, []float64{1, math.Exp(1), math.Exp(2), math.Exp(3), math.Exp(4)}, 1, math.Exp(4))

	if errB == nil {
		t.Error("Expected Error")
	}
}

func TestBezierCurve(t *testing.T) {
	testSetA, errA := BezierCurve([][2]float64{{0, 0}, {1, 0}}, [][2]float64{{2, 1}}, [][2]float64{{0, 1}})

	solutionSet := [][][4]float64{{{0, 6, -12, 7}}, {{0, 3, -3, 0}}}

	if errA != nil {
		t.Errorf("Error %v", errA)
	}

	if !reflect.DeepEqual(solutionSet, testSetA) {
		t.Errorf("Expected %v, received %v", solutionSet, testSetA)
	}

	_, errB := BezierCurve([][2]float64{{0, 0}, {1, 0}}, [][2]float64{{2, 1}, {4, 1}}, [][2]float64{{0, 1}, {0, 4}})

	if errB == nil {
		t.Error("Expected Error")
	}

	_, errC := BezierCurve([][2]float64{{0, 0}, {1, 0}}, [][2]float64{{2, 1}}, [][2]float64{{0, 1}, {0, 4}})

	if errC == nil {
		t.Error("Expected Error")
	}
}
