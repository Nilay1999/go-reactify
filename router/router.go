package router

import (
	"net/http"

	"github.com/Nilay1999/gin-gonic-server/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server running ...",
		})
	})

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/", user.GetPaginatedUser)
		}
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.GetUserById)
		}
		{
			user := new(controllers.UserController)
			userGroup.POST("/", user.CreateUser)
		}
		{
			user := new(controllers.UserController)
			userGroup.DELETE("/:id", user.DeleteUser)
		}
	}
	return router
}
