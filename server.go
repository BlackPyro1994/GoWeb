
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

func index(w http.ResponseWriter, r *http.Request) {
	p := Test {Title: "THIS IS A TEST", News: "some news" }

	tmpl := template.Must(template.ParseFiles("template/index.html", "template/header.html", "template/content_index.html"))

	tmpl.Execute(w, p)
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WORLD")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}