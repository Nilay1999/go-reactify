package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nilay1999/gin-gonic-server/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userService = new(services.User)

func (u UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := userService.GetById(id)

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

	users, err := userService.Get(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (u UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := userService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
