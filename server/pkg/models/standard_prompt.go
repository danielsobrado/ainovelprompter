package models

import "time"

type StandardPrompt struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	StandardName string    `gorm:"not null;index" json:"standard_name"`
	Title        string    `gorm:"not null" json:"title"`
	Prompt       string    `gorm:"not null" json:"prompt"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Version      int       `gorm:"default:1" json:"version"`
}
