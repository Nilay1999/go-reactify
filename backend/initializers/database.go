package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repository *gorm.DB

func Init() {
	DatabaseURL := os.Getenv("DATABASE_URL")

	repository, err := gorm.Open(postgres.Open(DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	Repository = repository
}
