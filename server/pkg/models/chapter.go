package models

import "github.com/jinzhu/gorm"

type Chapter struct {
	gorm.Model
	TextID        uint
	ChapterTitle  string
	ChapterNumber int
}
