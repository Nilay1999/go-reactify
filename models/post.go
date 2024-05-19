package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string `gorm:"size:255;not null;" json:"title"`
	Body      string `gorm:"size:255;not null;" json:"body"`
	Upvotes   int    `gorm:"size:255;not null;default:0" json:"upvotes"`
	Downvotes int    `gorm:"size:255;not null;default:0" json:"downvotes"`
	UserID    uint   `gorm:"not null"`
}
