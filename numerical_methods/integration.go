package numericalMethods

import (
    "fmt"
)

func Euler1D(a float64, b float64, N int, initValue float64, f func(float64, float64) float64) float64 {
    h := (b - a) / float64(N)
    t := a
    omega := initValue

    for i:= 0; i < N; i++ {
    	omega += h*f(t, omega)
	t += h
    }

    return omega
}

/* An example use case - output : 4.437500
func main() {
    result := Euler1D(0, 2, 4, 0.5, func(t float64, y float64) float64 { return y - math.Pow(t,2) + 1 })

    fmt.Printf("%f\n", result)
}
*/
