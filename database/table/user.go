package table

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"unique:not null"`
	Password       string `gorm:"not null"`
	Description    string
	ProfieImageURL string `gorm:"not null"`
}
