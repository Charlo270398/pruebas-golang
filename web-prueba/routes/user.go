package routes

import (
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	models "../models"
	util "../utils"
)

//GET

func menuUserHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Usuario", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		util.PrintErrorLog(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func historialUserHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/paciente/historial/index.html", "public/templates/layouts/menuUsuario.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Historial", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		util.PrintErrorLog(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

//DELETE

func deleteUserHandler(w http.ResponseWriter, req *http.Request) {
	var user util.User_id_json
	json.NewDecoder(req.Body).Decode(&user)
	result := models.DeleteUser(user.Id)
	if result == true{
		util.PrintLog("Usuario con ID '" + strconv.Itoa(user.Id) + "' borrado.")
	}else{
		util.PrintLog("Error borrando usuario con ID '" + strconv.Itoa(user.Id) + "'")
	}
}