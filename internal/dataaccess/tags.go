package dataaccess

import (
	"fmt"

	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"gorm.io/gorm"
)

func ListTags(db *gorm.DB) ([]models.Tag, error) {
	CleanPostTags(db)
	CleanTags(db)
	var tag []models.Tag
	result := db.Find(&tag)
	return tag, result.Error
}

func CreateTags(db *gorm.DB, rawTags []string, postId int) {
	CleanPostTags(db)
	CleanTags(db)
	db.Where("post_id = ?", postId).Delete(&models.PostsTags{})
	for _, rawTag := range rawTags {
		var tag = models.Tag{}
		db.FirstOrCreate(&tag, models.Tag{TagName: rawTag})
		var postTag = models.PostsTags{PostID: postId, TagID: tag.ID}
		db.Create(&postTag)
	}
}

func CleanTags(db *gorm.DB) {
	var tagIds []int
	var jTagIds []int
	db.Model(&models.Tag{}).Pluck("id", &tagIds)
	db.Model(&models.PostsTags{}).Pluck("tag_id", &jTagIds)
	tagIdDiff := utils.SetDifference(tagIds, jTagIds)
	if len(tagIdDiff) > 0 {
		db.Where("id IN (?)", tagIdDiff).Delete(&models.Tag{})
	}
}

func CleanPostTags(db *gorm.DB) {
	var postIds []int
	var jPostIds []int
	db.Model(&models.Post{}).Pluck("id", &postIds)
	db.Model(&models.PostsTags{}).Pluck("post_id", &jPostIds)
	postIdDiff := utils.SetDifference(jPostIds, postIds)
	if len(postIdDiff) > 0 {
		fmt.Println("hi")
		db.Where("post_id IN (?)", postIdDiff).Delete(&models.PostsTags{})
		fmt.Println(db.Where("post_id IN (?)", postIdDiff).Delete(&models.PostsTags{}))
	}

	var tagIds []int
	var jTagIds []int
	db.Model(&models.Tag{}).Pluck("id", &tagIds)
	db.Model(&models.PostsTags{}).Pluck("tag_id", &jTagIds)
	tagIdDiff := utils.SetDifference(jTagIds, tagIds)
	if len(tagIdDiff) > 0 {
		db.Where("tag_id IN ?", tagIdDiff).Delete(&models.PostsTags{})
	}
}
