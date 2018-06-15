package controller

import (
	"../../config"
	"net/http"
	"../model"
)


func RegisterKunden(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("user")
	email := r.FormValue("mail")
	password := r.FormValue("psw")

	statement := "insert into Kunde (Benutzername, Passwort, Email) values (?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(userName,email,password)

	return
}

// GetAll Kunden
func GetAllKunden() (kunden [] model.Kunde , err error) {
	rows, err := config.Db.Query("select * from Kunde")

	if err != nil {
		return
	}

	for rows.Next() {
		kunde := model.Kunde{}
		err = rows.Scan(&kunde.KundeID,&kunde.Benutzername,&kunde.BildUrl,&kunde.Typ,&kunde.Status,&kunde.Email, &kunde.Passwort)

		if err != nil {
			return
		}

		// fmt.Println(kunden)
		kunden = append(kunden, kunde)
	}
	rows.Close()
	return
}

/*func GetKundenById() (kunde model.Kunde , err error) {
	rows, err := config.Db.Query("select * from Kunde")

	if err != nil {
		return
	}

	for rows.Next() {
		kunde := model.Kunde{}
		err = rows.Scan(&kunde.KundeID,&kunde.Benutzername,&kunde.BildUrl,&kunde.Typ,&kunde.Status,&kunde.Email, &kunde.Passwort)

		if err != nil {
			return
		}

		fmt.Println(kunden)
		kunden = append(kunden, kunde)
	}
	rows.Close()
	return
}
*/