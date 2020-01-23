package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"./vendor/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  string
}

func createDB() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "golang:golang@(127.0.0.1:3306)/golang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	query := `
    CREATE TABLE IF NOT EXISTS usuarios (
		id INT AUTO_INCREMENT,
		dni VARCHAR(9) UNIQUE,
		nombre VARCHAR(25) NOT NULL,
		apellidos VARCHAR(50) NOT NULL,
        password VARCHAR(30) NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
	);`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database OK")
}
func printLog(text string) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf(text)
}

func main() {
	createDB()
	router := mux.NewRouter()
	//HOME
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("GET")

	//STATIC RESOURCES
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(s)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("Defaulting to port", port)
		printLog("Initialized at port " + port)
	}

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
