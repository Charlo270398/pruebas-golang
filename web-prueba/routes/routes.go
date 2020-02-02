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

	//STATIC RESOURCES
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(s)
	http.Handle("/", router)

	//ADMIN-GLOBAL
	router.HandleFunc("/user/gadmin", menuAdminHandler).Methods("GET")
	router.HandleFunc("/user/gadmin/users", getUserListAdminHandler).Methods("GET")
	router.HandleFunc("/user/gadmin/users/addform", addUserFormGadminHandler).Methods("GET")

	//LOGIN
	router.HandleFunc("/login", loginIndexHandler).Methods("GET")
	router.HandleFunc("/login", loginUserHandler).Methods("POST")
	router.HandleFunc("/register", registerIndexHandler).Methods("GET")
	router.HandleFunc("/register", registerUserHandler).Methods("POST")

	//HOME
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")

	//USER(GLOBAL)
	router.HandleFunc("/user/menu", menuUserHandler).Methods("GET")
	router.HandleFunc("/user/delete", deleteUserHandler).Methods("DELETE")

	//USER-PACIENTE
	router.HandleFunc("/user/patient", UserPatientHandler).Methods("GET")
	router.HandleFunc("/user/patient/edit", editUserPatientHandler).Methods("GET")
	router.HandleFunc("/user/patient/historial", historialUserHandler).Methods("GET")
	
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
