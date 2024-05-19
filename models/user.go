package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Gender   string `gorm:"size:255;not null;" json:"gender"`
	Age      uint   `gorm:"size:255;not null;" json:"age"`
	Posts    []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
