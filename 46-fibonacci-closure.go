package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current, previous := -1, 0
	return func() int {
		if current < 1 {
			current++
			return current
		}
		previous, current = current, current+previous
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 21; i++ {
		fmt.Println(f())
	}
}
