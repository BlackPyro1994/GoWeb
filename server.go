
package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type Test struct {
	Title string
	News string
}

func hello(w http.ResponseWriter, r *http.Request) {
	p := Test {Title: "THIS IS A TEST", News: "some news" }

	tmpl := template.Must(template.ParseFiles("template/test.html", "template/layout.html", "template/header.html"))

	tmpl.Execute(w, p)
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WORLD")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	http.ListenAndServe(":8080", nil)
}