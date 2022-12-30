package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func ListTags(db *gorm.DB) ([]models.Tag, error) {
	var tag []models.Tag
	result := db.Find(&tag)
	return tag, result.Error
}
