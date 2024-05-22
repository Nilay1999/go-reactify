package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID          uint     `gorm:"not null" json:"postId"`
	Post            Post     `gorm:"foreignKey:PostID"`
	UserID          uint     `gorm:"not null" json:"userId"`
	User            User     `gorm:"foreignKey:UserID"`
	ParentCommentID uint     `gorm:"not null" json:"parentCommentId"`
	ParentComment   *Comment `gorm:"foreignKey:ParentCommentID"`
}
