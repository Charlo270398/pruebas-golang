package routes

import (
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	models "../models"
	util "../utils"
)

type JSON_Return struct {
	Result    string
	Error     string
} 

//GET

func loginIndexHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/login/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Login", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		util.PrintErrorLog(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func registerIndexHandler(w http.ResponseWriter, req *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/login/register.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Register", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		util.PrintErrorLog(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

//POST

type JSON_Credentials struct {
	Password string `json:"password"`
	Email string `json:"email"`
} 

func loginUserHandler(w http.ResponseWriter, req *http.Request) {
	var creds JSON_Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//COMPROBAMOS USER Y PASS
	jsonReturn := JSON_Return{"", ""}
	correctLogin := models.LoginUser(creds.Email,creds.Password)
	if correctLogin == true {
		jsonReturn = JSON_Return{"Sesión Iniciada", ""}
	}else{
		jsonReturn = JSON_Return{"", "Usuario y contraseña incorrectos"}
	}
	js, err := json.Marshal(jsonReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func registerUserHandler(w http.ResponseWriter, req *http.Request) {
	var user util.User_json
	json.NewDecoder(req.Body).Decode(&user)
	inserted, err := models.InsertUser(user)
	jsonReturn := JSON_Return{"", ""}
	if inserted == true {
		jsonReturn = JSON_Return{"Hola", ""}
	}else{
		jsonReturn = JSON_Return{"", "El usuario no se ha podido registrar"}
	}

	js, err := json.Marshal(jsonReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}