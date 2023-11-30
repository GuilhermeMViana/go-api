package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermeMViana/go-rest-api/database"
	"github.com/GuilhermeMViana/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	var p []models.Person
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func IdFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Person
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
