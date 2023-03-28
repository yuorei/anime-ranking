package table

import "gorm.io/gorm"

type AnimeRanking struct {
	gorm.Model
	UserID        int    `gorm:"not null"`
	Title         string `gorm:"not null"`
	Rank          int    `gorm:"not null"`
	Description   string
	AnimeImageURL string `gorm:"not null"`
}
