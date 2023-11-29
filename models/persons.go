package models

type Person struct {
	Id      int    `json:"Id"`
	Name    string `json:"Nome"`
	History string `json:"História"`
}

var Persons []Person
