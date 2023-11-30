package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GuilhermeMViana/go-rest-api/database"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/GuilhermeMViana/go-rest-api/models"
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
	Id := vars["Id"]

	for _, person := range models.Persons {
		if strconv.Itoa(person.Id) == Id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
