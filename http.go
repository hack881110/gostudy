package main

import (
	"fmt"
	"log"
	"net/http"
)

//String  string
type String string

//Struct hello
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}

func (h Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h.Who, h.Punct, h.Greeting)
}
func main() {

	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:9998", nil))
}
