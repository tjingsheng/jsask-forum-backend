package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Post{}, &models.PostPreference{}, &models.PostsTags{}, &models.Tag{}, &models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Database migrated successfully!")
	}
}

func main() {
	db := database.InitDB() 
	Migrate(db)
}
