package numericalMethods

import (
	"errors"
	"math"
)

// Bisection1D is for solving the 1D root finding bisection method
func Bisection1D(a float64, b float64, TOL float64, maxIteration int, f func(x float64) float64) (float64, error) {
	fOfA := f(a)
	currentX := a + (b-a)/float64(2)
	fOfCurrentX := f(currentX)
	root := float64(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if fOfCurrentX == 0 || (b-a)/float64(2) < TOL {
			root = currentX
			solutionFound = true
			break
		}

		if fOfA*fOfCurrentX > 0 {
			a = currentX
			fOfA = fOfCurrentX
		} else {
			b = currentX
		}

		currentX = a + (b-a)/float64(2)
		fOfCurrentX = f(currentX)
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}

// FixedPointIteration1D is for solving the 1D root finding fixed point iteration method
func FixedPointIteration1D(initialApprox float64, TOL float64, maxIteration int, f func(x float64) float64) (float64, error) {
	previousApprox := initialApprox
	currentApprox := f(previousApprox)
	root := float64(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if math.Abs(currentApprox-previousApprox) < TOL {
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
func Newton1D(initialApprox float64, TOL float64, maxIteration int, f func(x float64) float64, df func(x float64) float64) (float64, error) {
	previousApprox := initialApprox
	currentApprox := previousApprox - f(previousApprox)/df(previousApprox)
	root := float64(0)
	solutionFound := false

	for i := 0; i < maxIteration; i++ {
		if math.Abs(currentApprox-previousApprox) < TOL {
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

// Secant1D is for solving the 1D root finding secant method
func Secant1D(initialApprox1 float64, intitialApprox2 float64, TOL float64, maxIteration int, f func(x float64) float64) (float64, error) {
	previousApprox1 := initialApprox1
	previousAprrox2 := intitialApprox2
	currentApprox := previousAprrox2 - f(previousAprrox2)*(previousAprrox2-previousApprox1)/(f(previousAprrox2)-f(previousApprox1))
	root := float64(0)
	solutionFound := false

	for i := 1; i < maxIteration; i++ {
		if math.Abs(currentApprox-previousAprrox2) < TOL {
			root = currentApprox
			solutionFound = true
			break
		}

		previousApprox1 = previousAprrox2
		previousAprrox2 = currentApprox
		currentApprox = previousAprrox2 - f(previousAprrox2)*(previousAprrox2-previousApprox1)/(f(previousAprrox2)-f(previousApprox1))
	}

	if solutionFound {
		return root, nil
	}

	return root, errors.New("Unable to find root of given function")
}
