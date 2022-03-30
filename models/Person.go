package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (e *Person) TableName() string {
	return "person"
}
