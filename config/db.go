package config

import (
	"log"
	"os"

	"github.com/herizal95/golang-jwt-mux/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("error Loading .env File")

	}
	conn := os.Getenv("Database_Connection")

	DB, err = gorm.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	// migrate models to database postgres

	DB.AutoMigrate(&models.User{})

	if err != nil {
		return
	}
}
