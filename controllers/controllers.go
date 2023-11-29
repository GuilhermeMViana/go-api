package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/GuilhermeMViana/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Persons)
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
