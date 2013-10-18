package main

import (
	"fmt"
	"math"
)

const delta = 1e-15

type sqrtFunc func(float64) (float64, int)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z, 10
}

func SqrtDelta(x float64) (float64, int) {
	iter := 0
	oldZ, z := 0.0, 1.0
	for {
		iter++
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-oldZ) <= delta {
			break
		}
		oldZ = z
	}
	return z, iter
}

func newton(funcName string, sqrt sqrtFunc) {
	fmt.Printf("| X   | Iterations | %-21s | %-21s | %-21s |\n", funcName, "math.Sqrt", "Diff")
	for i := 1; i <= 20; i++ {
		x := float64(i)
		sqrtValue, iterations := sqrt(x)
		mathSqrt := math.Sqrt(x)
		diff := math.Abs(mathSqrt - sqrtValue)
		fmt.Printf("| %-3v | %-10v | %-21v | %-21v | %-21v |\n", i, iterations, sqrtValue, mathSqrt, diff)
	}
}

func main() {
	newton("Sqrt", Sqrt)
	fmt.Println("------------------")
	newton("SqrtDelta", SqrtDelta)
}
