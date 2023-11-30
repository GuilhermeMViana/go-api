package main

import (
	"fmt"
	"github.com/GuilhermeMViana/go-rest-api/database"
	"github.com/GuilhermeMViana/go-rest-api/models"
	"github.com/GuilhermeMViana/go-rest-api/routes"
)

func main() {
	models.Persons = []models.Person{
		{Id: 1, Nome: "TBD", Historia: "TBD"},
		{Id: 2, Nome: "TBD1", Historia: "TBD1"},
	}

	database.ConnectDatabase()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
