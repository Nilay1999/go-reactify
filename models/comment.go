package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostId          uint     `gorm:"not null"`
	Upvotes         int      `gorm:"default:0"`
	Downvotes       int      `gorm:"default:0"`
	Post            uint     `gorm:"not null"`
	UserID          uint     `gorm:"not null"`
	ParentCommentID *uint    `gorm:"not null" json:"parent_comment_id"`
	ParentComment   *Comment `gorm:"foreignKey:ParentCommentID"`
}
