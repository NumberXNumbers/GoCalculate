package numericalMethods

import (
	"errors"
	"fmt"
)

// Bisection1D is for solving the 1D root finding bisection method
func Bisection1D(a float64, b float64, TOL float64, maxIteration int, f func(x float64) float64) (float64, error) {
	fOfA := f(a)
	currentX := a + (b-a)/float64(2)
	fOfCurrentX := f(currentX)
	root := float64(0)
	solutionFound := false
	fmt.Println(TOL)

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
