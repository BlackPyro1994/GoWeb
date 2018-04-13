
package main

import (
	"net/http"
	"log"
	"html/template"
	"strings"
)

func main() {

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})*/

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	// http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":80", nil)
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))
var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("image.tmpl"))

type Index struct {
	Title string
	Body  string
	Links []Link
}
type Link struct {
	URL, Title string
}
type Image struct {
	Title string
		URL   string
}

var images = map[string]*Image{
	"go":     {"The Go Gopher", "https://golang.org/doc/gopher/frontpage.png"},
	"google": {"The Google Logo", "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Image gallery",
		Body:  "Welcome to the image gallery.",
	}
	for name, img := range images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := images[strings.TrimPrefix(r.URL.Path, "/image/")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
