package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i:=1; i<100; i++ {
		z = z - (math.Pow(z, 2) - x)/2/z
		fmt.Println("Cycle:", z)
	}
	return z
}

func main() {
	fmt.Println("Res:", Sqrt(2))
}
