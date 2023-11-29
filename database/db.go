package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados!")
	}
}
