package models

type TraitKey struct {
	ID          uint   `gorm:"primary_key"`
	TraitKey    string `gorm:"unique;not null"`
	Description string
	TriggerText string
}
