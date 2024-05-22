package main

import (
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Init()
}

func main() {
	initializers.Repository.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Vote{})
}
