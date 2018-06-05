package main

import (
	"html/template"
	"net/http"
	"strings"
	"fmt"
	"./app/model"
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

type Ware struct {
	Items []Item
}

type Item struct {
	A string
	B string
	C string
	D string
	E string
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

	/* liste := list{
	[5]string{"x.png,Kamera1,Inv.Nr.1234,25.054.123123",
		"x.png,Kamera1,Inv.Nr.1234,25.054.123123",
		"x.png,Kamera2,Inv.Nr.1234,25.054.123123",
		"x.png,Kamera3,Inv.Nr.1234,25.054.123123",
		"x.png,Kamera4,Inv.Nr.1234,25.054.123123"}}*/

	lol := Ware{
		Items: []Item{
			{A: "x.png", B: "Kamera1", C: "Inv.Nr.12", D: "34", E: "25.054.12"},
			{A: "x.png", B: "Kamera1", C: "Inv.Nr.12", D: "34", E: "25.054.12"},
		},
	}

	// fmt.Println(lol.items[0])

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/clients.html", "template/header.html", "template/layout.html"))

	tmpl.ExecuteTemplate(w, "main", nil)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "clients", lol)

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

var artikelList = make(model.Artikels)

func main() {

	fmt.Println("ALLE ARTIKEL")
	fmt.Println(model.ReadAllArtikel())
	fmt.Println()
	fmt.Println("UPDATE ARTIKEL MIT ID 5")
	model.UpdateArtikel(5, "ABC", "KATEGORIE", "LAGERORT", 99, "HINWEIS", "BILDURL")
	fmt.Println()
	fmt.Println("ALLE ARTIKEL 2.0")
	fmt.Println(model.ReadAllArtikel())
	fmt.Println()
	fmt.Println("DELETE ARTIKEL MIT ID 5")
	model.DeleteArtikel(5)
	fmt.Println()
	fmt.Println("ADD NEUER ARTIKEL")
	artikel := model.Artikel{0,"Bez","Kat", "LAGERORT", 99, "hinweis","url"}
	artikel.CreateArtikel()
	fmt.Println(model.ReadAllArtikel())

	fmt.Println()
	fmt.Println("#####################################")
	fmt.Println()

	fmt.Println("ALLE KUNDEN")
	fmt.Println(model.ReadAllKunden())
	fmt.Println()
	fmt.Println("UPDATE KUNDE MIT ID 2")
	model.UpdateKunde(2, "ABC", "PASSWORT", "test@test.de")
	fmt.Println()
	fmt.Println("ALLE KUNDEN 2.0")
	fmt.Println(model.ReadAllKunden())
	fmt.Println()
	fmt.Println("DELETE KUNDE MIT ID 2")
	model.DeleteKunde(2)
	fmt.Println()
	fmt.Println("ALLE KUNDEN 3.0")
	fmt.Println(model.ReadAllKunden())

	fmt.Println()
	fmt.Println("#####################################")
	fmt.Println()

	fmt.Println("ALLE VERLEIHER")
	fmt.Println(model.ReadAllVerleiher())
	fmt.Println()
	fmt.Println("UPDATE VERLEIHER MIT ID 3")
	model.UpdateVerleiher(3, "BNAME", "test@test.de")
	fmt.Println()
	fmt.Println("ALLE VERLEIHER 2.0")
	fmt.Println(model.ReadAllVerleiher())
	fmt.Println()
	fmt.Println("DELETE VERLEIHER MIT ID 3")
	model.DeleteVerleiher(3)
	fmt.Println()
	fmt.Println("ALLE VERLEIHER 3.0")
	fmt.Println(model.ReadAllVerleiher())

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

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}
