package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
	"gorm.io/gorm"
)

func ListPosts(db *gorm.DB) ([]postsviewmodel.ListView, error) {
	var posts []postsviewmodel.ListView
	result := db.Model(&models.Post{}).
		Preload("Tags").
		Preload("User").
		Preload("Comments").
		Preload("PostPreferences").
		Order("post_datetime DESC").
		Find(&posts, "parent_post = ?", 0)
	return posts, result.Error
}

func ListCurrPost(db *gorm.DB, postId string) ([]postsviewmodel.ListView, error) {
	var posts []postsviewmodel.ListView
	result := db.Model(&models.Post{}).
		Preload("Tags").
		Preload("User").
		Preload("Comments").
		Preload("PostPreferences").
		Where("parent_post = ?", postId).
		Or("id = ?", postId).Find(&posts)
	return posts, result.Error
}

func CreatePost(db *gorm.DB, newPost models.Post) error {
	result := db.Create(&newPost)
	return result.Error
}
