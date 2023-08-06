package models

import "github.com/jinzhu/gorm"

type People struct {
	gorm.Model

	ID int64 `json:"id"`
	Name string `json:"name"`
	LastName string `json:"last name"`
	Address string `json:"adress"`
	Telephone string `json:"telephone"`
}




