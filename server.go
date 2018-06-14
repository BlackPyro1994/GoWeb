package main

import (
	"net/http"
	"fmt"
	"./app/model"
	"./app/route"
)

func main() {

	fmt.Println("ALLE ARTIKEL")
	fmt.Println(model.ReadAllArtikel())
	fmt.Println()
	fmt.Println("UPDATE ARTIKEL MIT ID 5")
	model.UpdateArtikel(5, "ABC", "KATEGORIE", "LAGERORT", 99, "HINWEIS", "BILDURL")
	fmt.Println()
	fmt.Println("ALLE ARTIKEL 2.0")
	fmt.Println(model.ReadAllArtikel())
	fmt.Println()
	fmt.Println("DELETE ARTIKEL MIT ID 5")
	model.DeleteArtikel(5)
	fmt.Println()
	fmt.Println("ADD NEUER ARTIKEL")
	artikel := model.Artikel{0,"Bez","Kat", "LAGERORT", 99, "hinweis","url"}
	artikel.CreateArtikel()
	fmt.Println(model.ReadAllArtikel())

	fmt.Println()
	fmt.Println("#####################################")
	fmt.Println()



	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}
