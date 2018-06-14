package model

import (
	_ "github.com/mattn/go-sqlite3"
	"../../config"
	"database/sql"
)

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

var Db *sql.DB = config.Db

type Artikels map[int]*Artikel

// GetAll Artikel
func ReadAllArtikel() (artikels []Artikel, err error) {
	rows, err := Db.Query("select * from Artikel")

	if err != nil {
		return
	}

	for rows.Next() {
		artikel := Artikel{}
		err = rows.Scan(&artikel.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)

		if err != nil {
			return
		}
		artikels = append(artikels, artikel)
	}
	rows.Close()
	return
}

// Create Artikel
func  (artikel *Artikel) CreateArtikel()  (err error) {
	statement := "insert into Artikel (Bezeichnung, Kategorie, Lagerort, Anzahl, Hinweis, BildURL) values (?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(&artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)
	//_, err = Db.Exec("insert into Artikel (Bezeichnung, Kategorie, Lagerort, Anzahl, Hinweis, BildURL) values (bez, kat, lago, anz, hin, url)")
	return
}

// Read Artikel by ArtikelID
func ReadArtikel(id int) (artikel Artikel, err error) {
	artikel = Artikel{}
	err = Db.QueryRow("select *  from Artikel where ArtikelID = $1", id).Scan(&artikel.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)
	return
}

// Update the Artikel-Bezeichnung by id
func  UpdateArtikel(id int, bez string, kat string, lago string, anz int, hin string, url string) (err error) {
	_, err = Db.Exec("update Artikel set Bezeichnung = $1 where ArtikelID = $2",bez, id)
	_, err = Db.Exec("update Artikel set Kategorie = $1 where ArtikelID = $2",kat, id)
	_, err = Db.Exec("update Artikel set Lagerort = $1 where ArtikelID = $2",lago, id)
	_, err = Db.Exec("update Artikel set Anzahl = $1 where ArtikelID = $2",anz, id)
	_, err = Db.Exec("update Artikel set Hinweis = $1 where ArtikelID = $2",hin, id)
	_, err = Db.Exec("update Artikel set BildURL = $1 where ArtikelID = $2",url, id)
	return
}

// Delete Artikel by id
func DeleteArtikel(id int) (err error) {
	_, err = Db.Exec("delete from Artikel where ArtikelID = $1", id)
	return
}
