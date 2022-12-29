package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Init() *gorm.DB {
	test := goDotEnvVariable("TEST_TEST")
	fmt.Printf("godotenv : %s = %s \n", "TEST_TEST", test)

	dbURL := "postgres://pg:pass@localhost:5432/crud" //FIX
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.PostPreference{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.User{})

	return db
}

type Database struct {
}

func GetDB() (*Database, error) {
	return &Database{}, nil
}
