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

func seedData[T any](db *gorm.DB, filename string) {
	data, err := csvToStruct[T](filename)
	if err != nil {
		log.Fatalf("Error converting CSV to struct: %v", err)
	}

	for _, record := range data {
		if err := db.Create(&record).Error; err != nil {
			log.Fatalf("Failed to insert record: %v", err)
		}
	}

	fmt.Printf("Successfully seeded %T from %s\n", *new(T), filename)
}

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

func seed(db *gorm.DB) {
	seedData[models.User](db, "seed/users.csv")
	seedData[models.Post](db, "seed/posts.csv")
	seedData[models.PostPreference](db, "seed/post_preferences.csv")
	seedData[models.PostsTags](db, "seed/posts_tags.csv")
	seedData[models.Tag](db, "seed/tags.csv")
}

func main() {
	db := database.InitDB()
	seed(db)
}
