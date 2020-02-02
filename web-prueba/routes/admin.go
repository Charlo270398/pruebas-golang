package routes

import (
	"html/template"
	"log"
	"net/http"
	models "../models"
)


//GET
func menuAdminHandler(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/admin/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Menú administrador", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getUserListAdminHandler(w http.ResponseWriter, req *http.Request) {
	models.ShowAllUsers()
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/list.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Listado de usuarios", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}