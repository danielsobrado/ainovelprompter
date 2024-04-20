package models

type TraitType struct {
	TraitTypeID             uint   `gorm:"primary_key"`
	TraitType               string `gorm:"not null"`
	Description             string `gorm:"not null"`
	AnalyzeTraitDescription string `gorm:"not null"`
}
