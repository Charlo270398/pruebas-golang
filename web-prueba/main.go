package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body string
}

func main() {
	router := mux.NewRouter()
	//HOME 
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/login/", loginHandler).Methods("GET")
	
	//STATIC RESOURCES
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
    router.PathPrefix("/public/").Handler(s)
    http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

// indexHandler uses a template to create an index.html.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/home/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Menu", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(
		template.New("").ParseFiles("public/templates/user/index.html", "public/templates/layouts/base.html"),
	)
	if err := tmp.ExecuteTemplate(w, "base", &Page{Title: "Login", Body: "body"}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
