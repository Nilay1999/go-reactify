package controllers

import (
	"net/http"

	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (a AuthController) Signin(ctx *gin.Context) {
	var payload types.AuthType
	if validationError := ctx.BindJSON(&payload); validationError != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(validationError)})
		return
	}

	response, serviceError := userService.Authenticate(payload)
	if serviceError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}
	if response.Token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": map[string]string{"message": response.Message}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": map[string]string{"message": response.Message, "token": response.Token}})
}

func (a AuthController) Signup(ctx *gin.Context) {
	var payloadType types.UserType
	if err := ctx.BindJSON(&payloadType); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	user, err := userService.Create(payloadType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
