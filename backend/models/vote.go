package models

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	UserID   uint   `gorm:"not null" json:"userId"`
	User     User   `gorm:"foreignKey:UserID"`
	PostID   uint   `gorm:"not null" json:"postId"`
	Post     Post   `gorm:"foreignKey:PostID"`
	VoteType string `gorm:"size:255;not null;" json:"voteType"`
}
