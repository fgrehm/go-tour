package main

import "code.google.com/p/go-tour/pic"

type funcPic func(int, int) int

func picMult(x, y int) int {
	return x * y
}

func picHalfSum(x, y int) int {
	return (x + y) / 2
}

func picPow(x, y int) int {
	return x ^ y
}

func picSum(x, y int) int {
	return x + y
}

func Pic(funcForPic funcPic) (func(int, int) [][]uint8) {
	return func(dx, dy int) [][]uint8 {
		res := make([][]uint8, dy)
		for y := 0; y < dy; y++ {
			res[y] = make([]uint8, dx)
			for x := 0; x < dx; x++ {
				res[y][x] = uint8(funcForPic(x, y))
			}
		}
		return res
	}
}

func main() {
	// Pick one at a time, use the go tour to see the images
	pic.Show(Pic(picMult))
	//pic.Show(Pic(picSum))
	//pic.Show(Pic(picHalfSum))
	//pic.Show(Pic(picPow))
}
