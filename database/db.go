package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))

	if err != nil {
		log.Fatal("Problema ao conectar com banco de dados:", err)
	} else {
		fmt.Println("Sucesso ao conectar com o banco de dados.")
	}
}
