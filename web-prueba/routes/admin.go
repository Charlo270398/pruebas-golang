package routes

import (
	"html/template"
	"log"
	"net/http"
	models "../models"
)

func menuAdminHandler(w http.ResponseWriter, r *http.Request) {
	models.ShowAllUsers()
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/admin/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Menú administrador", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func usersAdminHandler(w http.ResponseWriter, r *http.Request) {
	models.ShowAllUsers()
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/admin/usersList.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Listado de usuarios", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}