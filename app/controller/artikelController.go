package controller

import (
	"../model"
	"../../config"
	"net/http"
)

// GetAll Artikel
func ReadAllArtikel() (artikels [] model.Artikel, err error) {
	rows, err := config.Db.Query("select * from Artikel")

	if err != nil {
		return
	}

	for rows.Next() {
		artikel := model.Artikel{}
		err = rows.Scan(&artikel.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)

		if err != nil {
			return
		}
		artikels = append(artikels, artikel)
	}
	rows.Close()
	return
}

func CreateArtikel(w http.ResponseWriter, r *http.Request) {

	bezeichnung := r.FormValue("bz")
	kategorie := r.FormValue("kat")
	inventarNummer:= r.FormValue("invNum")
	lagerort := r.FormValue("lgo")
	inhalt:= r.FormValue("inhalt")
	hinweis:= r.FormValue("hinweis")
	anzahl := r.FormValue("anz")

	statement := "insert into Artikel (Bezeichnung, Kategorie,InventarNummer,Lagerort, Anzahl,Hinweis,BildURL) values (?,?,?,?,?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(bezeichnung,kategorie,inventarNummer,lagerort,inhalt,hinweis,anzahl)

	return
}

// Read Artikel by ArtikelID
func ReadArtikel(id int) (artikel model.Artikel, err error) {
	artikel = model.Artikel{}
	err = config.Db.QueryRow("select *  from Artikel where ArtikelID = $1", id).Scan(&artikel.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)
	return
}

// Update the Artikel-Bezeichnung by id
func UpdateArtikel(id int, bez string, kat string, lago string, anz int, hin string, url string) (err error) {
	_, err = config.Db.Exec("update Artikel set Bezeichnung = $1 where ArtikelID = $2", bez, id)
	_, err = config.Db.Exec("update Artikel set Kategorie = $1 where ArtikelID = $2", kat, id)
	_, err = config.Db.Exec("update Artikel set Lagerort = $1 where ArtikelID = $2", lago, id)
	_, err = config.Db.Exec("update Artikel set Anzahl = $1 where ArtikelID = $2", anz, id)
	_, err = config.Db.Exec("update Artikel set Hinweis = $1 where ArtikelID = $2", hin, id)
	_, err = config.Db.Exec("update Artikel set BildURL = $1 where ArtikelID = $2", url, id)
	return
}

// Delete Artikel by id
func DeleteArtikel(id int) (err error) {
	_, err = config.Db.Exec("delete from Artikel where ArtikelID = $1", id)
	return
}
