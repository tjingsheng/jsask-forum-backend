package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"gorm.io/gorm"
)

func UpdateTables(db *gorm.DB) {
	CleanPostTags(db)
	CleanTags(db)
	CleanPostPreferences(db)
}

// Delete all tags that are not tagged to any post
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

// Delete all invalid post-tag entries i.e. if post/tag has been deleted
func CleanPostTags(db *gorm.DB) {
	var postIds []int
	var jPostIds []int
	db.Model(&models.Post{}).Pluck("id", &postIds)
	db.Model(&models.PostsTags{}).Pluck("post_id", &jPostIds)
	postIdDiff := utils.SetDifference(jPostIds, postIds)
	if len(postIdDiff) > 0 {
		db.Where("post_id IN (?)", postIdDiff).Delete(&models.PostsTags{})
	}

	var tagIds []int
	var jTagIds []int
	db.Model(&models.Tag{}).Pluck("id", &tagIds)
	db.Model(&models.PostsTags{}).Pluck("tag_id", &jTagIds)
	tagIdDiff := utils.SetDifference(jTagIds, tagIds)
	if len(tagIdDiff) > 0 {
		db.Where("tag_id IN (?)", tagIdDiff).Delete(&models.PostsTags{})
	}
}

// Delete all invalid postpreferences i.e. if post/user has been deleted
func CleanPostPreferences(db *gorm.DB) {
	var postIds []int
	var jPostIds []int
	db.Model(&models.Post{}).Pluck("id", &postIds)
	db.Model(&models.PostPreference{}).Pluck("post_id", &jPostIds)
	postIdDiff := utils.SetDifference(jPostIds, postIds)
	if len(postIdDiff) > 0 {
		db.Where("post_id IN (?)", postIdDiff).Delete(&models.PostPreference{})
	}

	var userIds []int
	var jUserIds []int
	db.Model(&models.User{}).Pluck("id", &userIds)
	db.Model(&models.PostPreference{}).Pluck("user_id", &jUserIds)
	userIdDiff := utils.SetDifference(jUserIds, userIds)
	if len(userIdDiff) > 0 {
		db.Where("user_id IN (?)", userIdDiff).Delete(&models.PostPreference{})
	}
}
