package database

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		godotenv.Load(".env")
		DB_URL := os.Getenv("DB_URL")

		var err error
		DB, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{
			PrepareStmt: false,
		})

		if err != nil {
			log.Fatalln(err)
		}
	})
}
