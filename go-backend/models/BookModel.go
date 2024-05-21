package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID            uint
	Title         string
	Author        string
	Genre         string
	Summary       string
	PublishedYear uint
}

func (Book) TableName() string {
	return "books"
}
