package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"gorm.io/gorm"

	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func csvToStruct[T any](filename string) ([]T, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %v", err)
	}

	var result []T
	structType := reflect.TypeOf(*new(T))

	for _, record := range records {
		obj := reflect.New(structType).Elem()

		for i := 0; i < structType.NumField(); i++ {
			field := obj.Field(i)

			if field.Kind() == reflect.Int {
				val, _ := strconv.Atoi(record[i])
				field.SetInt(int64(val))
			} else if field.Kind() == reflect.String {
				field.SetString(record[i])
			}
		}

		result = append(result, obj.Interface().(T))
	}

	return result, nil
}

func seedAndSyncSequence[T any](tx *gorm.DB, filename, tableName string) error {
	data, err := csvToStruct[T](filename)
	if err != nil {
		return fmt.Errorf("error converting CSV to struct: %v", err)
	}

	for _, record := range data {
		if err := tx.Create(&record).Error; err != nil {
			return fmt.Errorf("failed to insert record: %v", err)
		}
	}

	fmt.Printf("Successfully seeded %T from %s\n", *new(T), filename)

	// We set the sequence number to 10000 to avoid conflicts with existing data as they are seeded with IDs specified in the CSV files.
	query := fmt.Sprintf(
		`SELECT setval(
			COALESCE(pg_get_serial_sequence('%s', 'id'), ''), 
			10000, 
			true
		);`,
		tableName,
	)

	if err := tx.Exec(query).Error; err != nil {
		return fmt.Errorf("failed to reset sequence for %s: %v", tableName, err)
	}

	fmt.Printf("Successfully reset sequence for %s\n", tableName)
	return nil
}

func seed(db *gorm.DB) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}

	tx.Exec("DEALLOCATE ALL;")

	err := seedAndSyncSequence[models.User](tx, "seed/users.csv", "users")
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error seeding users: %v", err)
	}

	err = seedAndSyncSequence[models.Tag](tx, "seed/tags.csv", "tags")
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error seeding tags: %v", err)
	}

	err = seedAndSyncSequence[models.Post](tx, "seed/posts.csv", "posts")
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error seeding posts: %v", err)
	}

	err = seedAndSyncSequence[models.PostPreference](tx, "seed/post_preferences.csv", "post_preferences")
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error seeding post preferences: %v", err)
	}

	err = seedAndSyncSequence[models.PostsTags](tx, "seed/posts_tags.csv", "posts_tags")
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error seeding posts_tags: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Seeding completed successfully!")
}

func main() {
	database.InitDB()
	seed(database.DB)

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()
}
