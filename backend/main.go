package main

import (
	"os"

	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/router"
)

func init() {
	if os.Getenv("STAGE") != "production" {
		initializers.LoadEnvVariables()
	}
	initializers.Init()
}

func main() {
	r := router.InitRouter()
	r.Run(":" + os.Getenv("PORT"))
}
