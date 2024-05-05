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
	Gender   string `gorm:"size:255;not null;" json:"gender"`
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

func (u User) GetById(id string) (*User, error) {
	var user User

	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) Get(offset int, limit int) ([]User, error) {
	var users []User
	result := initializers.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u User) Delete(id string) (string, error) {
	var user User
	var message string = "User deleted !"
	var errorMessage string = "User deletion failed !"

	result := initializers.DB.Delete(&user, id)
	if result.Error != nil {
		return errorMessage, result.Error
	}
	return message, nil
}
