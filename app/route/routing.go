package route

import (
	"html/template"
	"net/http"
	"strings"
	"../model"
	"../controller"
)

//--------------------------------------------------------------------

type menu struct {
	Title     string
	Item1     string
	Item2     string
	Item3     string
	Basket    bool
	Name      string
	Type      string
	Profil    bool
	EmptySide bool
	Profile   bool
}

type Items struct {
	KundenItems   []model.Kunde
	VerleiheItems []model.Verleihe
}

type Clients struct {
	Items []client
}

type client struct {
	BildUrl      string
	Benutzername string
	KundenID     int
	Typ          string
	Bezeichnungen  []Bez
	Status       string
}

type Bez struct {
	Bezeichnung string
}

//--------------------------------------------------------------------

var funcMap = template.FuncMap{
	"split": func(s string, index int) string {
		// fmt.Println(s, index)
		arr := strings.Split(s, ",")

		if s == "" {
			return ""
		} else {
			return arr[index]
		}

	},
}

var artikelList = make(model.Artikels)

//------------------------------------------------------------------------

func index(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   false}

	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/index.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "index", p)

}

func login(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: true,
		Profile:   false}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/login.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "login", p)

}

func register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		controller.RegisterKunden(w, r)
		// index(w,r)
	} else {

		// REGISTER
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Login,login",
			Item3:     "",
			Basket:    false,
			Name:      "",
			Type:      "",
			EmptySide: true,
			Profile:   false}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/register.html", "template/layout.html", "template/header.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "register", p)
	}
}

func equipment(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "equipment", p)

}

func cart(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "cart", p)

}

func myequipment(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/myequipment.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "myequipment", p)

}

func profile(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/profile.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "profile", p)

}

func admin(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "Peter",
		Type:      "Verleiher",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "admin", p)

}

func adminEquipment(w http.ResponseWriter, r *http.Request) {

	// ADMIN
	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEquipment.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminEquipment", p)

}

func adminAddEquipment(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		controller.CreateArtikel(w, r)
		// equipment(w,r)
	} else {

	}
	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminAddEquipment.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminAddEquipment", p)

}

func adminClients(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		adminEditClients()
	} else {


		p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	KundenArr := controller.GetAllKunden()

	var ClientsArr= []client{}

	// for index := range ClientsArr {
	for _,element := range KundenArr {
		// ClientsArr = append(ClientsArr,client{controller.getKundenById(controller.getVerleihById(index).kundeID)).bildUrl,"asdasd","asdasd","asdasd","asdasd","asdasdad",},)

		artikelFromUser := controller.GetAllArtikelFromKunde(element.KundeID)

		var EquipmentString = []Bez{}

		for _,element := range artikelFromUser {

			EquipmentString = append(EquipmentString,Bez{element})
			}

			ClientsArr = append(ClientsArr, client{element.BildUrl, element.Benutzername, element.KundeID, element.Typ, EquipmentString, element.Status})
		}

		data := Clients{
			Items: ClientsArr,
		}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/clients.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", nil)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "clients", data)

	}

func adminEditClients(w http.ResponseWriter, r *http.Request) {

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditClients.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminEditClients", p)

}

//--------------------------------------------------------------------

func Handler() {

	http.HandleFunc("/", index)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/equipment", adminEquipment)
	http.HandleFunc("/admin/add", adminAddEquipment)
	http.HandleFunc("/admin/clients", adminClients)
	http.HandleFunc("/admin/edit-client", adminEditClients)
	http.HandleFunc("/login", login)
	http.HandleFunc("/equipment", equipment)
	http.HandleFunc("/myequipment", myequipment)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/register", register)
	http.HandleFunc("/cart", cart)

	return
}
