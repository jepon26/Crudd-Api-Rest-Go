package models

import "github.com/jinzhu/gorm"

type people struct {
	gorm.Model

	ID int64 `json:"string"`
}
