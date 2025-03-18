package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	godotenv.Load(".env")
	DB_URL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
