package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Artikels map[int]* Artikel

// Artikel data structure
type Artikel struct {
	ArtikelID int
	Bezeichnung string
	Kategorie string
	Lagerort string
	Anzahl int
	Hinweis string
	BildURL string
}



