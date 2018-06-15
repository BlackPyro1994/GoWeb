package main

import (
	"net/http"
	"./app/route"
	"./config"
)

func main() {

	config.InitSQLiteDB()

	route.Handler()

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}
