package model

import (
	_ "github.com/mattn/go-sqlite3"
)

// Kunde data structure
type Kunde struct {
	KundeID int
	Benutzername string
	Passwort string
	Email string
}
// GetAll Kunde
func ReadAllKunden() (kunden []Kunde, err error) {
	rows, err := Db.Query("select * from Kunde")

	if err != nil {
		return
	}

	for rows.Next() {
		kunde := Kunde{}
		err = rows.Scan(&kunde.KundeID, &kunde.Benutzername, &kunde.Passwort, &kunde.Email)

		if err != nil {
			return
		}
		kunden = append(kunden, kunde)
	}
	rows.Close()
	return
}

// Create Kunde
func  CreateKunde(bname string, psw string, mail string)  (err error) {
	//defer stmt.Close()
	//_, err = stmt.Exec(&kunde.Bezeichnung, &kunde.Kategorie, &kunde.Lagerort, &kunde.Anzahl, &kunde.Hinweis, &kunde.BildURL)
	_, err = Db.Exec("insert into Kunde (Benutzername, Passwort, Email) values (bname, psw, mail)")
	return
}

// Read Kunde by KundeID
func ReadKunde(id int) (kunde Kunde, err error) {
	kunde = Kunde{}
	err = Db.QueryRow("select *  from Kunde where KundeID = $1", id).Scan(&kunde.KundeID, &kunde.Benutzername, &kunde.Passwort, &kunde.Email)
	return
}

// Update the Kunde by id
func  UpdateKunde(id int, bname string, psw string, mail string) (err error) {
	_, err = Db.Exec("update Kunde set Benutzername = $1 where KundeID = $2",bname, id)
	_, err = Db.Exec("update Kunde set Passwort = $1 where KundeID = $2",psw, id)
	_, err = Db.Exec("update Kunde set Email = $1 where KundeID = $2",mail, id)
	return
}

// Delete Kunde by id
func DeleteKunde(id int) (err error) {
	_, err = Db.Exec("delete from Kunde where KundeID = $1", id)
	return
}