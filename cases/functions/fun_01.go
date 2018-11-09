package functions

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, a float64, b float64) float64 {
	return fn(a, b)
}

func Fun01() {
	fmt.Printf("-> 3*4=%v\n", compute(func(a float64, b float64) float64 { return a * b }, 3, 4))
	fmt.Printf("-> 3^4=%v\n", compute(math.Pow, 3, 4))
	fmt.Printf("-> max(3,4)=%v\n", compute(math.Max, 3, 4))
}
