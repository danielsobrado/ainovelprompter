package models

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	ChapterID uint
	UserID    *uint
	Content   string
	Rating    *int
}
