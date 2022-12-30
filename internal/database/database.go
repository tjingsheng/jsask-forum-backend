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
	DB := goDotEnvVariable("DB")
	DB_USER := goDotEnvVariable("DB_USER")
	DB_PASS := goDotEnvVariable("DB_PASS")
	dbURL := "postgres://" + DB_USER + ":" + DB_PASS + "@" + DB + "/JsaskDatabase"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
