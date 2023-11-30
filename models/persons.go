package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"história" gorm:"tableName:persons"`
}

var Persons []Person
