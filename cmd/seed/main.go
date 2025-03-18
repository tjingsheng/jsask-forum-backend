package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"gorm.io/gorm"

	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
)

func SeedUsers(db *gorm.DB, filename string) {
  users, err := CSVToStruct[models.User](filename)
  if err != nil {
    log.Fatalf("Error converting CSV to struct: %v", err)
  }

  for _, user := range users {
    if err := db.Create(&user).Error; err != nil {
      log.Fatalf("Failed to insert user: %v", err)
    }
  }

  fmt.Println("Successfully seeded users from CSV")
}

func CSVToStruct[T any](filename string) ([]T, error) {
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


func Seed(db *gorm.DB) {
	SeedUsers(db, "db/users.csv")
}

func main() {
	db := database.InitDB()
	Seed(db)
}
