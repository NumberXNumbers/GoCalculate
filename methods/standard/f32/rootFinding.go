package methods

import (
	"errors"
	"math"
)

// Bisection1D is for solving the 1D root finding bisection method
func Bisection1D(intervalBegin float32, intervalEnd float32, TOL float32, maxIteration int, f func(x float32) float32) (float32, error) {
	fOfA := f(intervalBegin)
	currentX := intervalBegin + (intervalEnd-intervalBegin)/float32(2)
	fOfCurrentX := f(currentX)
	root := float32(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if fOfCurrentX == 0 || (intervalEnd-intervalBegin)/float32(2) < TOL {
			root = currentX
			solutionFound = true
			break
		}

		if fOfA*fOfCurrentX > 0 {
			intervalBegin = currentX
			fOfA = fOfCurrentX
		} else {
			intervalEnd = currentX
		}

		currentX = intervalBegin + (intervalEnd-intervalBegin)/float32(2)
		fOfCurrentX = f(currentX)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// FixedPointIteration1D is for solving the 1D root finding fixed point iteration method
func FixedPointIteration1D(initialApprox float32, TOL float32, maxIteration int, f func(x float32) float32) (float32, error) {
	previousApprox := initialApprox
	currentApprox := f(previousApprox)
	root := float32(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox))) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		currentApprox = f(previousApprox)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// Newton1D is for solving the 1D  root finding newton's method
func Newton1D(initialApprox float32, TOL float32, maxIteration int, f func(x float32) float32, df func(x float32) float32) (float32, error) {
	previousApprox := initialApprox
	currentApprox := previousApprox - f(previousApprox)/df(previousApprox)
	root := float32(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox))) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		currentApprox = previousApprox - f(previousApprox)/df(previousApprox)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// ModifiedNewton1D is a modification for solving the 1D  root finding newton's method
func ModifiedNewton1D(initialApprox float32, TOL float32, maxIteration int, f func(x float32) float32,
	df func(x float32) float32, ddf func(x float32) float32) (float32, error) {
	previousApprox := initialApprox
	currentApprox := previousApprox - f(previousApprox)*df(previousApprox)/(float32(math.Pow(float64(df(previousApprox)), 2))-f(previousApprox)*ddf(previousApprox))
	root := float32(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox))) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox = currentApprox
		currentApprox = previousApprox - f(previousApprox)*df(previousApprox)/(float32(math.Pow(float64(df(previousApprox)), 2))-f(previousApprox)*ddf(previousApprox))
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// Secant1D is for solving the 1D root finding secant method
func Secant1D(initialApprox1 float32, intitialApprox2 float32, TOL float32, maxIteration int, f func(x float32) float32) (float32, error) {
	previousApprox1 := initialApprox1
	previousApprox2 := intitialApprox2
	currentApprox := previousApprox2 - f(previousApprox2)*(previousApprox2-previousApprox1)/(f(previousApprox2)-f(previousApprox1))
	root := float32(0)
	solutionFound := false

	for i := 1; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox2))) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox1 = previousApprox2
		previousApprox2 = currentApprox
		currentApprox = previousApprox2 - f(previousApprox2)*(previousApprox2-previousApprox1)/(f(previousApprox2)-f(previousApprox1))
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// FalsePosition1D is for solving the 1D root finding false position method
func FalsePosition1D(initialApprox1 float32, initialApprox2 float32, TOL float32, maxIteration int, f func(x float32) float32) (float32, error) {
	previousApprox1 := initialApprox1
	previousApprox2 := initialApprox2
	fOfApprox1 := f(previousApprox1)
	fOfApprox2 := f(previousApprox2)
	currentApprox := previousApprox2 - fOfApprox2*(previousApprox2-previousApprox1)/(fOfApprox2-fOfApprox1)
	fOfCurrentApprox := f(currentApprox)
	root := float32(0)
	solutionFound := false

	for i := 1; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox2))) < TOL {
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
		fOfCurrentApprox = f(currentApprox)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// Steffensen1D is for solving the 1D root finding Steffensen's mehtod
func Steffensen1D(initialApprox float32, TOL float32, maxIteration int, f func(x float32) float32) (float32, error) {
	previousApprox1 := initialApprox
	previousApprox2 := f(previousApprox1)
	previousApprox3 := f(previousApprox2)
	currentApprox := previousApprox1 - float32(math.Pow(float64(previousApprox2-previousApprox1), 2))/(previousApprox3-2*previousApprox2+previousApprox1)
	root := float32(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if float32(math.Abs(float64(currentApprox-previousApprox1))) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox1 = currentApprox
		previousApprox2 = f(previousApprox1)
		previousApprox3 = f(previousApprox2)

		currentApprox = previousApprox1 - float32(math.Pow(float64(previousApprox2-previousApprox1), 2))/(previousApprox3-2*previousApprox2+previousApprox1)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}
