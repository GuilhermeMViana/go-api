package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Id       int    `json:"Id"`
	Nome     string `json:"Nome"`
	Historia string `json:"História"`
}

var Persons []Person
