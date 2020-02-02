package routes

import (
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	models "../models"
	util "../utils"
)

//GET

func loginIndexHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/login/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Login", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func registerIndexHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/login/register.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Register", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

//POST

func loginUserHandler(w http.ResponseWriter, req *http.Request) {

}

func registerUserHandler(w http.ResponseWriter, req *http.Request) {
	var user util.User_json
    json.NewDecoder(req.Body).Decode(&user)
	models.InsertUser(user)
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(true)
}