package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		char := p[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			p[i] += 13
		} else if (char > 'M' && char <= 'Z') || (char > 'm' && char <= 'z') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
