package models

import "github.com/jinzhu/gorm"

type Text struct {
	gorm.Model
	UserID   uint
	TextType string
	Content  string
}
