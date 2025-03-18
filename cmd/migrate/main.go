package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func enableRLS(db *gorm.DB, tableName string) {
	query := fmt.Sprintf("ALTER TABLE %s ENABLE ROW LEVEL SECURITY;", tableName)
	if err := db.Exec(query).Error; err != nil {
		log.Fatalf("Failed to enable RLS for table %s: %v", tableName, err)
	}
}

func Migrate(db *gorm.DB) {
	db.Exec("DEALLOCATE ALL;")

	err := db.AutoMigrate(
		&models.Post{},
		&models.PostPreference{},
		&models.PostsTags{},
		&models.Tag{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Database migrated successfully!")
	}

	tables := []string{"posts", "post_preferences", "posts_tags", "tags", "users"}
	for _, table := range tables {
		enableRLS(db, table)
	}
	log.Printf("RLS enabled for table %v", tables)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB connection: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	} else {
		log.Println("Database connection closed after migration.")
	}
}

func main() {
	db := database.InitDB()
	Migrate(db)
}
