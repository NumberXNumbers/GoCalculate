package methods

import (
	"errors"
	"math"

	"github.com/NumberXNumbers/GoCalculate/types/gcf"
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
)

// Bisection1D is for solving the 1D root finding bisection method
func Bisection1D(intervalBegin float64, intervalEnd float64, TOL float64, maxIteration int, f *gcf.Function) (gcv.Value, error) {
	fOfA := f.MustEval(intervalBegin).Value()
	currentX := intervalBegin + (intervalEnd-intervalBegin)/float64(2)
	fOfCurrentX := f.MustEval(currentX).Value()
	var root gcv.Value
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if fOfCurrentX.Real() == 0 || (intervalEnd-intervalBegin)/float64(2) < TOL {
			root = gcv.MakeValue(currentX)
			solutionFound = true
			break
		}

		if gcvops.Mult(fOfA, fOfCurrentX).Real() > 0 {
			intervalBegin = currentX
			fOfA = fOfCurrentX
		} else {
			intervalEnd = currentX
		}

		currentX = intervalBegin + (intervalEnd-intervalBegin)/float64(2)
		fOfCurrentX = f.MustEval(currentX).Value()
	}

	if solutionFound {
		return root, nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// FixedPointIteration1D is for solving the 1D root finding fixed point iteration method
func FixedPointIteration1D(initialApprox float64, TOL float64, maxIteration int, f *gcf.Function) (gcv.Value, error) {
	previousApprox := gcv.MakeValue(initialApprox)
	currentApprox := f.MustEval(previousApprox).Value()
	var root gcv.Value
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if gcvops.Abs(gcvops.Sub(currentApprox, previousApprox)).Real() < TOL {
			root = gcv.MakeValue(currentApprox)
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		currentApprox = f.MustEval(previousApprox).Value()
	}

	if solutionFound {
		return root, nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// Newton1D is for solving the 1D  root finding newton's method
func Newton1D(initialApprox float64, TOL float64, maxIteration int, f *gcf.Function, df *gcf.Function) (gcv.Value, error) {
	previousApprox := gcv.MakeValue(initialApprox)
	currentApprox := gcvops.Sub(previousApprox, gcf.MustDiv(f.MustEval(previousApprox), df.MustEval(previousApprox)).Value())
	var root gcv.Value
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if gcvops.Abs(gcvops.Sub(currentApprox, previousApprox)).Real() < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		currentApprox = gcvops.Sub(previousApprox, gcf.MustDiv(f.MustEval(previousApprox), df.MustEval(previousApprox)).Value())
	}

	if solutionFound {
		return root, nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// ModifiedNewton1D is a modification for solving the 1D  root finding newton's method
func ModifiedNewton1D(initialApprox float64, TOL float64, maxIteration int, f *gcf.Function,
	df *gcf.Function, ddf *gcf.Function) (gcv.Value, error) {
	previousApprox := gcv.MakeValue(initialApprox)
	two := gcv.MakeValue(2)
	fOfPreviousApproxConst := f.MustEval(previousApprox)
	dfOfPreviousApproxConst := df.MustEval(previousApprox)
	ratioA := gcf.MustMult(fOfPreviousApproxConst, dfOfPreviousApproxConst).Value()
	ratioB := gcf.MustMult(fOfPreviousApproxConst, ddf.MustEval(previousApprox)).Value()
	currentApprox := gcvops.Sub(previousApprox, gcvops.Div(ratioA, gcvops.Sub(gcvops.Pow(dfOfPreviousApproxConst.Value(), two), ratioB)))
	var root gcv.Value
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if gcvops.Abs(gcvops.Sub(currentApprox, previousApprox)).Real() < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		fOfPreviousApproxConst := f.MustEval(previousApprox)
		dfOfPreviousApproxConst := df.MustEval(previousApprox)
		ratioA := gcf.MustMult(fOfPreviousApproxConst, dfOfPreviousApproxConst).Value()
		ratioB := gcf.MustMult(fOfPreviousApproxConst, ddf.MustEval(previousApprox)).Value()
		currentApprox = gcvops.Sub(previousApprox, gcvops.Div(ratioA, gcvops.Sub(gcvops.Pow(dfOfPreviousApproxConst.Value(), two), ratioB)))
	}

	if solutionFound {
		return root, nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// Secant1D is for solving the 1D root finding secant method
func Secant1D(initialApprox1 float64, intitialApprox2 float64, TOL float64, maxIteration int, f *gcf.Function) (gcv.Value, error) {
	previousApprox1 := gcv.MakeValue(initialApprox1)
	previousApprox2 := gcv.MakeValue(intitialApprox2)
	fOfPreviousApprox1 := f.MustEval(previousApprox1).Value()
	fOfPreviousApprox2 := f.MustEval(previousApprox2).Value()
	ratioA := gcvops.Div(gcvops.Sub(previousApprox2, previousApprox1), gcvops.Sub(fOfPreviousApprox2, fOfPreviousApprox1))
	currentApprox := gcvops.Sub(previousApprox2, gcvops.Mult(fOfPreviousApprox2, ratioA))
	var root gcv.Value
	solutionFound := false

	for i := 1; i < maxIteration; i++ {
		if gcvops.Abs(gcvops.Sub(currentApprox, previousApprox2)).Real() < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox1 = previousApprox2
		previousApprox2 = currentApprox
		fOfPreviousApprox1 := f.MustEval(previousApprox1).Value()
		fOfPreviousApprox2 := f.MustEval(previousApprox2).Value()
		ratioA := gcvops.Div(gcvops.Sub(previousApprox2, previousApprox1), gcvops.Sub(fOfPreviousApprox2, fOfPreviousApprox1))
		currentApprox = gcvops.Sub(previousApprox2, gcvops.Mult(fOfPreviousApprox2, ratioA))
	}

	if solutionFound {
		return root, nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// FalsePosition1D is for solving the 1D root finding false position method
func FalsePosition1D(initialApprox1 float64, initialApprox2 float64, TOL float64, maxIteration int, f *gcf.Function) (gcv.Value, error) {
	previousApprox1 := initialApprox1
	previousApprox2 := initialApprox2
	fOfApprox1 := f.MustEval(previousApprox1).Value().Real()
	fOfApprox2 := f.MustEval(previousApprox2).Value().Real()
	currentApprox := previousApprox2 - fOfApprox2*(previousApprox2-previousApprox1)/(fOfApprox2-fOfApprox1)
	fOfCurrentApprox := f.MustEval(currentApprox).Value().Real()
	root := float64(0)
	solutionFound := false

	for i := 1; i < maxIteration; i++ {
		if math.Abs(currentApprox-previousApprox2) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		if fOfCurrentApprox*fOfApprox2 < 0 {
			previousApprox1 = previousApprox2
			fOfApprox1 = fOfApprox2
		}

		previousApprox2 = currentApprox
		fOfApprox2 = fOfCurrentApprox

		currentApprox = previousApprox2 - fOfApprox2*(previousApprox2-previousApprox1)/(fOfApprox2-fOfApprox1)
		fOfCurrentApprox = f.MustEval(currentApprox).Value().Real()
	}

	if solutionFound {
		return gcv.MakeValue(root), nil
	}

	return nil, errors.New("Unable to find root of given function")
}

// Steffensen1D is for solving the 1D root finding Steffensen's mehtod
func Steffensen1D(initialApprox float64, TOL float64, maxIteration int, f *gcf.Function) (gcv.Value, error) {
	previousApprox1 := initialApprox
	previousApprox2 := f.MustEval(previousApprox1).Value().Real()
	previousApprox3 := f.MustEval(previousApprox2).Value().Real()
	currentApprox := previousApprox1 - math.Pow((previousApprox2-previousApprox1), 2)/(previousApprox3-2*previousApprox2+previousApprox1)
	root := float64(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if math.Abs(currentApprox-previousApprox1) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox1 = currentApprox
		previousApprox2 = f.MustEval(previousApprox1).Value().Real()
		previousApprox3 = f.MustEval(previousApprox2).Value().Real()

		currentApprox = previousApprox1 - math.Pow((previousApprox2-previousApprox1), 2)/(previousApprox3-2*previousApprox2+previousApprox1)
	}

	if solutionFound {
		return gcv.MakeValue(root), nil
	}

	return nil, errors.New("Unable to find root of given function")
}
