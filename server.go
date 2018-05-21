
package main

import (
	"net/http"
	"html/template"
)

type menu struct {
	Title string
	Item1 string
	Item2 string
	Item3 string
	Basket bool
	Name string
	Type string
	Profil bool
}

func index(w http.ResponseWriter, r *http.Request) {
	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Login",
		Item3: "",
		Basket: false,
		Name: "",
		Type: "",
		Profil: false}

	tmpl := template.Must(template.ParseFiles("template/index.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "index", p)


}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}