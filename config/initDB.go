package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Db handle
var Db *sql.DB
var err error

func initSQLiteDB() {

	Db, err = sql.Open("sqlite3", "./data/borgdir.media")
	if err != nil {
		panic(err)
	}
}

func initPostgresDB() {
	connStr := "user=postgres dbname=postgres password=postgres host=localhost port=5431 sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)

	}
}
