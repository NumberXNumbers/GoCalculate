package functions

import (
	"errors"
	"fmt"
	"math"
)

// LegendrePolynomial returns the nth Legendre polynomial, P_n(x), as a reusable function
// will panic if integer n is less than 0
func LegendrePolynomial(n int) func(x float64) float64 {
	if n < 0 {
		panic(fmt.Sprint("Integer must be greater than 0"))
	}

	return func(x float64) float64 {
		var sum float64
		for k := 0; k <= n; k++ {
			sum += math.Pow(BinomialCoefficient(n, k), 2.0) * math.Pow(x-1.0, float64(n-k)) * math.Pow(x+1, float64(k))
		}
		return math.Exp2(-float64(n)) * sum
	}
}

// LegendrePolynomial2 returns the nth Legendre polynomial, P_n(x), as a reusable function
// will return an error if integer n is less than 0
func LegendrePolynomial2(n int) (Pn func(x float64) float64, err error) {
	if n < 0 {
		err = errors.New("Integer must be greater than 0")
		return
	}

	Pn = func(x float64) float64 {
		var sum float64
		for k := 0; k <= n; k++ {
			sum += math.Pow(BinomialCoefficient(n, k), 2.0) * math.Pow(x-1.0, float64(n-k)) * math.Pow(x+1, float64(k))
		}
		return math.Exp2(-float64(n)) * sum
	}

	return
}

// BinomialCoefficient returns n choose k
func BinomialCoefficient(n int, k int) float64 {
	prod := 1.0
	for i := 1; i <= k; i++ {
		prod *= float64(n-k+i) / float64(i)
	}

	return prod
}
