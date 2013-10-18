package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str)
}

func (h *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v%v %v", h.Greeting, h.Punct, h.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct-go", &Struct{"Hello", ":", "Gophers!"})
	http.Handle("/struct-ruby", &Struct{"Hello", ":", "Rubysts!"})
	http.ListenAndServe("localhost:4000", nil)
}
