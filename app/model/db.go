package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

// Kunde data structure
type Kunde struct {
	KundeID int
	Benutzername string
	Passwort string
	Email string
}

// Kunde data structure
type Verleiher struct {
	VerleiherID int
	Benutzername string
	Email string
}

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media")
	if err != nil {
		panic(err)
	}
}

// GetAll Artikel
func GetAllArtikel() (artikels []Artikel, err error) {
	rows, err := Db.Query("select ArtikelID, Bezeichnung, Kategorie, Lagerort, Anzahl, Hinweis, BildUrl from Artikel")

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

// Get Artikel by ArtikelID
func Get(id int) (artikel Artikel, err error) {
	artikel = Artikel{}
	err = Db.QueryRow("select ArtikelID, Bezeichnung, Kategorie, Lagerort, Anzahl, Hinweis, BildUrl from Artikel where ArtikelID = $1", id).Scan(&artikel.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)
	return
}

// Add Artikel
func  Add(bez string, kat string, lago string, anz int, hin string, url string)  (err error) {
	//defer stmt.Close()
	//_, err = stmt.Exec(&artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)
	_, err = Db.Exec("insert into Artikel (Bezeichnung, Kategorie, Lagerort, Anzahl, Hinweis, BildURL) values (bez, kat, lago, anz, hin, url)")
	return
}

// ToggleStatus toggles the completion status of the Artikel by id
func  ToggleStatus(bez string, id int) (err error) {
	_, err = Db.Exec("update Artikel set Bezeichnung = $1 where ArtikelID = $2",bez,id)
	return
}

// Delete Artikel by id from the list of artikels
func Delete(id int) (err error) {
	_, err = Db.Exec("delete from Artikel where ArtikelID = $1", id)
	return
}
