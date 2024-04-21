package models

type TraitType struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	TraitType   string `gorm:"column:trait_type" json:"traitType"`
	Description string `gorm:"column:description" json:"description"`
	TriggerText string `gorm:"column:trigger_text" json:"triggerText"`
}
