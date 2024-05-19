package router

import (
	"net/http"

	"github.com/Nilay1999/gin-gonic-server/controllers"
	"github.com/Nilay1999/gin-gonic-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server running ...",
		})
	})

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user").Use(middleware.AuthenticateRequest)
		{
			user := new(controllers.UserController)
			userGroup.GET("", user.GetPaginatedUser)
		}
		{
			user := new(controllers.UserController)
			userGroup.GET(":id", user.GetUserById)
		}
		{
			user := new(controllers.UserController)
			userGroup.DELETE(":id", user.DeleteUser)
		}

		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.POST("signin", auth.Signin)
		}
		{
			auth := new(controllers.AuthController)
			authGroup.POST("signup", auth.Signup)
		}
	}
	return router
}
