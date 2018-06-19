package main

import (
	"net/http"
	"./app/route"
	"./config"

)

func main() {

	// config.InitSQLiteDB()
	config.InitPostgresDB()

	// controller.GetAllArtikelFromKunde(1)

	// temp:=""

	// println(rows)
	// println(temp)
	// println(err)

	route.Handler()

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}
