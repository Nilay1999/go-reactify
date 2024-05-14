package services

import (
	"os"
	"time"

	"github.com/Nilay1999/gin-gonic-server/helpers"
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenResponse struct {
	Message string
	Token   string
}

type User struct {
	models.User
}

func (u User) Create(payload types.UserType) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		User: models.User{
			Username: payload.Username,
			Password: string(hashedPassword),
			Email:    payload.Email,
			Gender:   payload.Gender,
		},
	}

	result := initializers.Repository.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) GetById(id string) (*User, error) {
	var user User

	result := initializers.Repository.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u User) Get(offset int, limit int) ([]User, error) {
	var users []User
	result := initializers.Repository.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u User) Delete(id string) (string, error) {
	var user User
	var message string = "User deleted !"
	var errorMessage string = "User deletion failed !"

	result := initializers.Repository.Delete(&user, id)
	if result.Error != nil {
		return errorMessage, result.Error
	}
	return message, nil
}

func (u User) Authenticate(payload types.AuthType) (TokenResponse, error) {
	var user User

	var message string = "Authentication successful !"
	var errorMessage string = "User not found with given credentials !"

	isIdentifierEmail := helpers.ValidateEmail(payload.Identifier)
	if isIdentifierEmail {
		result := initializers.Repository.Where("email = ?", payload.Identifier).First(&user)
		if result.Error != nil {
			return TokenResponse{Message: errorMessage, Token: ""}, result.Error
		}
		error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

		if error == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user.ID,
				"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})
			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
			return TokenResponse{Message: message, Token: tokenString}, err
		} else {
			errorMessage = "Incorrect password!"
			return TokenResponse{Message: errorMessage, Token: ""}, nil
		}

	} else {
		result := initializers.Repository.Where("username = ?", payload.Identifier).First(&user)
		if result.Error != nil {
			return TokenResponse{Message: errorMessage, Token: ""}, result.Error
		}

		error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
		if error == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user.ID,
				"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})
			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
			return TokenResponse{Message: message, Token: tokenString}, err
		} else {
			errorMessage = "Incorrect password!"
			return TokenResponse{Message: errorMessage, Token: ""}, nil
		}
	}
}
