// Remember to `go get code.google.com/p/go-tour/pic` first

package main

import "code.google.com/p/go-tour/pic"

func PicMult(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res[y][x] = uint8(x * y)
		}
	}
	return res
}

func PicHalfSum(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res[y][x] = uint8((x + y) / 2)
		}
	}
	return res
}

func PicPow(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res[y][x] = uint8(x ^ y)
		}
	}
	return res
}

func main() {
	pic.Show(PicMult)
	pic.Show(PicHalfSum)
	pic.Show(PicPow)
}
