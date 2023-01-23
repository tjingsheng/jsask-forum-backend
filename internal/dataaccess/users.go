package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func ListCurrUser(db *gorm.DB, username string, password string) (models.User, error) {
	var user models.User
	result := db.Where(models.User{Username: username, Password: password}).First(&user)
	return user, result.Error
}

func ListSalt(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	result := db.Where(models.User{Username: username}).First(&user)
	return user, result.Error
}

func CreateUser(db *gorm.DB, newUser models.User) error {
	result := db.Create(&newUser)
	return result.Error
}
