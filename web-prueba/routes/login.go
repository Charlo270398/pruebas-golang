package routes

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	models "../models"
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

type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}

func registerUserHandler(w http.ResponseWriter, req *http.Request) {
	var user User
    json.NewDecoder(req.Body).Decode(&user)
    fmt.Println("Hola")
	models.InsertUser()
}