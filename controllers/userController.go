package controllers

import (
	"net/http"

	"github.com/Nilay1999/gin-gonic-server/models"
	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

var (
	g          = galidator.New()
	customizer = g.Validator(types.UserType{})
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) CreateUser(ctx *gin.Context) {
	var payloadType types.UserType
	if err := ctx.BindJSON(&types.UserType{}); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}
	user, err := userModel.Create(payloadType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
