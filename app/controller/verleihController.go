package controller

/*
func GetAllVerleihe() (verleihe [] model.Verleih) {
	rows, err := config.Db.Query("select * from Verleih")

	if err != nil {
		return
	}

	for rows.Next() {
		verleih := model.Verleih{}
		err = rows.Scan(&verleih.ArtikelID, &artikel.Bezeichnung, &artikel.Kategorie, &artikel.Lagerort, &artikel.Anzahl, &artikel.Hinweis, &artikel.BildURL)

		if err != nil {
			return
		}
		artikels = append(artikels, artikel)
	}
	rows.Close()
	return
}

*/