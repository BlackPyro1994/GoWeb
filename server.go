
package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
func hello(w http.ResponseWriter, r *http.Request) {

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl := template.Must(template.ParseFiles("test.tmpl"))
	tmpl.Execute(w, data)
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WORLD")
}

func main() {
	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}