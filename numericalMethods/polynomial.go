package numericalMethods

import "math"

// LegendrePolynomial returns P_n(x) where P_n is the nth degree legendre polynomial
func LegendrePolynomial(n int, x float64) float64 {
	var sum float64
	for k := 0; k <= n; k++ {
		sum += math.Pow(BinomialCoefficient(n, k), 2.0) * math.Pow(x-1.0, float64(n-k)) * math.Pow(x+1, float64(k))
	}
	return math.Exp2(-float64(n)) * sum
}

// LegendrePolynomialNGenerator returns a function that takes x and returns P_n(x)
func LegendrePolynomialNGenerator(n int) func(x float64) float64 {
	return func(x float64) float64 {
		var sum float64
		for k := 0; k <= n; k++ {
			sum += math.Pow(BinomialCoefficient(n, k), 2.0) * math.Pow(x-1.0, float64(n-k)) * math.Pow(x+1, float64(k))
		}
		return math.Exp2(-float64(n)) * sum
	}
}

// BinomialCoefficient returns n choose k
func BinomialCoefficient(n int, k int) float64 {
	prod := 1.0
	for i := 1; i <= k; i++ {
		prod *= float64(n-k+i) / float64(i)
	}

	return prod
}
