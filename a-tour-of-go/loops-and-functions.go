// Exercise: Loops and Functions
// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	delta := 1e-6
	previous := float64(0)
	z := float64(x / 2)

	for i := 0; ; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(previous-z) < delta {
			break
		}
		previous = z
	}
	return z
}

func main() {
	values := []float64{10001347, 1000.1587}

	for _, v := range values {
		fmt.Printf("v: %f. Sqrt: %f, math.Sqrt: %f, diff: %F\n", v, Sqrt(v), math.Sqrt(v), math.Abs(Sqrt(v)-math.Sqrt(v)))
	}
}
