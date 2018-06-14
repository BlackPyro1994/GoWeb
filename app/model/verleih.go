package model

import (
	_ "github.com/mattn/go-sqlite3"
)

// Verleiher data structure
type Verleiher struct {
	VerleiherID int
	Benutzername string
	Email string
}

// GetAll Verleiher
func ReadAllVerleiher() (verleihers []Verleiher, err error) {
	rows, err := Db.Query("select * from Verleiher")

	if err != nil {
		return
	}

	for rows.Next() {
		verleiher := Verleiher{}
		err = rows.Scan(&verleiher.VerleiherID, &verleiher.Benutzername, &verleiher.Email)

		if err != nil {
			return
		}
		verleihers = append(verleihers, verleiher)
	}
	rows.Close()
	return
}

// Create Verleiher
func  CreateVerleiher(bname string, mail string)  (err error) {
	//defer stmt.Close()
	_, err = Db.Exec("insert into Verleiher (Benutzername, Email) values (bname, mail)")
	return
}

// Read Verleiher by VerleiherID
func ReadVerleiher(id int) (verleiher Verleiher, err error) {
	verleiher = Verleiher{}
	err = Db.QueryRow("select *  from Verleiher where VerleiherID = $1", id).Scan(&verleiher.VerleiherID, &verleiher.Benutzername, &verleiher.Email)
	return
}

// Update the Verleiher by id
func  UpdateVerleiher(id int, bname string, mail string) (err error) {
	_, err = Db.Exec("update Verleiher set Benutzername = $1 where VerleiherID = $2",bname, id)
	_, err = Db.Exec("update Verleiher set Email = $1 where VerleiherID = $2",mail, id)
	return
}

// Delete Verleiher by id
func DeleteVerleiher(id int) (err error) {
	_, err = Db.Exec("delete from Verleiher where VerleiherID = $1", id)
	return
}