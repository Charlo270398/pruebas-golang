package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"log"
)

type Page struct {
	Title string
	Body  string
}

func LoadRouter() {
	router := mux.NewRouter()

	//ADMIN
	router.HandleFunc("/administrador/usuarios", usersAdminHandler).Methods("GET")

	//LOGIN
	router.HandleFunc("/login", loginIndexHandler).Methods("GET")
	router.HandleFunc("/login", loginUserHandler).Methods("POST")
	router.HandleFunc("/register", registerIndexHandler).Methods("GET")
	router.HandleFunc("/register", registerUserHandler).Methods("POST")

	//HOME
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")

	
	
	

	//STATIC RESOURCES
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(s)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("Defaulting to port", port)
		//printLog("Initialized at port " + port)
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
