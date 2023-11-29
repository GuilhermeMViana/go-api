package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/GuilhermeMViana/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persons", controllers.AllPersons).Methods("Get")
	r.HandleFunc("/api/persons/{Id}", controllers.IdFilter).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
