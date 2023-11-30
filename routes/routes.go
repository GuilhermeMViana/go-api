package routes

import (
	"log"
	"net/http"

	"github.com/GuilhermeMViana/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persons", controllers.AllPersons).Methods("Get")
	r.HandleFunc("/api/persons/{id}", controllers.IdFilter).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
