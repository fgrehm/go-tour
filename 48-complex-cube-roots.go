package main

import (
	"fmt"
	"math/cmplx"
)

const delta = 1e-15

func Cbrt(x complex128) complex128 {
	oldZ, z := complex128(0.0), complex128(1.0)
	for {
		z = z - (z*z*z-x)/(3*z*z)
		if cmplx.Abs(z-oldZ) <= delta {
			break
		}
		oldZ = z
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 1.0/3))
}
