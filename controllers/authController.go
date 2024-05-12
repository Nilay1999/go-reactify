package controllers

import (
	"net/http"

	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (a AuthController) Login(ctx *gin.Context) {
	var payload types.AuthType
	if validationError := ctx.BindJSON(&payload); validationError != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(validationError)})
		return
	}

	response, serviceError := userModel.Authenticate(payload)
	if serviceError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": map[string]string{"message": response.Message, "token": response.Token}})
}
