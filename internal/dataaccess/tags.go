package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"gorm.io/gorm"
)

func ListTags(db *gorm.DB) ([]models.Tag, error) {
	var tag []models.Tag
	result := db.Find(&tag)
	return tag, result.Error
}

func CreateTags(db *gorm.DB, rawTags []string, postId int) {
	db.Where("post_id = ?", postId).Delete(&models.PostsTags{})
	for _, rawTag := range rawTags {
		var tag = models.Tag{}
		db.FirstOrCreate(&tag, models.Tag{TagName: utils.CapitalizeWords(rawTag)})
		var postTag = models.PostsTags{PostID: postId, TagID: tag.ID}
		db.Create(&postTag)
	}
}
