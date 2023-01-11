package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/currpostviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
	"gorm.io/gorm"
)

func ListPosts(db *gorm.DB) ([]postsviewmodel.ListView, error) {
	var posts []postsviewmodel.ListView
	result := db.Model(&models.Post{}).Preload("Tags").Preload("User").Preload("Comments").Preload("PostPreferences").Find(&posts, "parent_post = ?", 0)
	return posts, result.Error
}

func ListCurrPost(db *gorm.DB, postId string) (currpostviewmodel.ListView, error) {
	var posts currpostviewmodel.ListView
	result := db.Model(&models.Post{}).Preload("Tags").Preload("User").Preload("Comments").Preload("PostPreferences").Where("id = ?", postId).Find(&posts, "parent_post = ?", 0)
	return posts, result.Error
}
