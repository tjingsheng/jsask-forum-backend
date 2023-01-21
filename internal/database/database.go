package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	DB := os.Getenv("DB")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	// DB := goDotEnvVariable("DB")
	// DB_USER := goDotEnvVariable("DB_USER")
	// DB_PASS := goDotEnvVariable("DB_PASS")
	dbURL := "postgres://" + DB_USER + ":" + DB_PASS + "@" + DB + "/JsaskDatabase"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// func goDotEnvVariable(key string) string {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }
