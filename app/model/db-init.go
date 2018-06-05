package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media")
	if err != nil {
		panic(err)
	}
}
