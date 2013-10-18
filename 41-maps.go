// Remember to `go get code.google.com/p/go-tour/wc` first

package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	split := strings.Fields(s)
	for _, word := range split {
		(ret[word])++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
	fmt.Println()
}
