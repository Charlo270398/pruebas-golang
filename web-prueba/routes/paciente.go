package routes

import (
	"html/template"
	"log"
	"net/http"
	util "../utils"
)

//GET
func editUserPatientHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/paciente/edit.html", "public/templates/layouts/menuPaciente.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Mis datos", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		util.PrintErrorLog(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}