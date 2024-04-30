package models

import (
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Gender   string `gorm:"not null" json:"gender"`
}

func (u User) Create(payload types.UserType) (*User, error) {
	user := User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Gender:   payload.Gender,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
