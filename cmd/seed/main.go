package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	faker "github.com/go-faker/faker/v4"
	"gorm.io/gorm"

	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
)

// TODO: Use brcypt instead of GenerateSalt and HashPassword
func GenerateSalt() string {
	salt := make([]byte, 16) 
	_, err := rand.Read(salt)
	if err != nil {
		panic(err) 
	}
	return hex.EncodeToString(salt)
}

func HashPassword(salt, password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(salt + password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SeedUsers(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		password := "1234"
		salt := GenerateSalt()
		hashedPassword := HashPassword(salt, password)

		user := models.User{
			Username:     faker.FirstName(),
			UserDatetime: time.Now().Format("2006-01-02 15:04:05"),
			Salt:         salt,
			Password:     hashedPassword, 
		}

		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("failed to seed user: %v", err)
		}
	}
	fmt.Println("Successfully seeded users")
}

func Seed(db *gorm.DB) {
	SeedUsers(db, 10)
}

func main() {
	db := database.InitDB()
	Seed(db)
}
