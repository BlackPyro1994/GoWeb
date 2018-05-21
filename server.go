
package main

import (
	"net/http"
	"html/template"
)

type menu struct {
	Title string
	Item1 string
	Item2 string
	Item3 string
	Basket bool
	Name string
	Type string
	Profil bool
	EmptySide bool
	Profile bool
}

func index(w http.ResponseWriter, r *http.Request) {
	// INDEX
	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Login",
		Item3: "",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: false }

	tmpl := template.Must(template.ParseFiles("template/index.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "index", p)


}
func register(w http.ResponseWriter, r *http.Request) {

	// REGISTER
	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Login",
		Item3: "",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: true,
		Profile: false }

	tmpl := template.Must(template.ParseFiles("template/register.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "register", p)


}
func equipment(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Meine Geräte",
		Item3: "Logout",
		Basket: true,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/equipment.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "equipment", p)


}
func cart(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Meine Geräte",
		Item3: "Logout",
		Basket: true,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/cart.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "cart", p)


}
func myequipment(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Meine Geräte",
		Item3: "Logout",
		Basket: true,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/myequipment.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "myequipment", p)


}
func profile(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Meine Geräte",
		Item3: "Logout",
		Basket: true,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/profile.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "profile", p)


}
func admin(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Kunden",
		Item3: "Logout",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/admin.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "admin", p)


}
func adminEquipment(w http.ResponseWriter, r *http.Request) {

	// ADMIN
	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Kunden",
		Item3: "Logout",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }
	tmpl := template.Must(template.ParseFiles("template/adminEquipment.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminEquipment", p)


}
func addEquipment(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Kunden",
		Item3: "Logout",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/addEquipment.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "addEquipment", p)


}
func adminClients(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Kunden",
		Item3: "Logout",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/adminClients.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "admin", p)


}
func adminEditClients(w http.ResponseWriter, r *http.Request) {

	p := menu {
		Title: "borgdir.media",
		Item1: "Equipment",
		Item2: "Kunden",
		Item3: "Logout",
		Basket: false,
		Name: "",
		Type: "",
		EmptySide: false,
		Profile: true }

	tmpl := template.Must(template.ParseFiles("template/adminEditClients.html", "template/header.html"))

	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminEditClients", p)


}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/admin", index)
	http.HandleFunc("/equipment", index)
	http.HandleFunc("/kunden", index)
	http.HandleFunc("/login", index)
	http.HandleFunc("/meinegeräte", index)
	http.HandleFunc("/profilbearbeiten", index)
	http.HandleFunc("/registrieren", index)
	http.HandleFunc("/warenkorb", index)

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}