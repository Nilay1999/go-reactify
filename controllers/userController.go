package controllers

import (
	"net/http"
	"strconv"

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
	if err := ctx.BindJSON(&payloadType); err != nil {
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

func (u UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := userModel.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (u UserController) GetPaginatedUser(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset := (page - 1) * limit

	users, err := userModel.Get(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (u UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := userModel.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
