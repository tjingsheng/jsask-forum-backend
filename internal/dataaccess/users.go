package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func ListCurentUser(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}
