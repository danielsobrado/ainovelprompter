package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username       string `gorm:"unique"`
    Email          string `gorm:"unique"`
    HashedPassword string
    Role           string
    IsActive       bool
}