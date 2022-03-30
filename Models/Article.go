package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (e *Article) TableName() string {
	return "articles"
}
 

