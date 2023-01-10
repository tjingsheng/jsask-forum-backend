package postsviewmodel

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

type ListView struct {
	ID           int    `json:"postID"`
	UserID       int    `json:"userID"`
	PostDatetime string `json:"postDatetime"`
	PostTitle    string `json:"postTitle"`
	PostContent  string `json:"postContent"`
	ParentPost   int    `json:"parentPost"`

	Tags []models.Tag `json:"tags" gorm:"many2many:posts_tags;joinForeignKey:id;"`

	User models.User `json:"user" gorm:"foreignKey:id"`

	Comments []models.Post `json:"comments" gorm:"foreignKey:parent_post"`

	PostPreferences []models.PostPreference `json:"postPreferences"  gorm:"foreignKey:post_id"`
}

func ListFrom(post ListView) ListView {
	return ListView{
		ID:           post.ID,
		UserID:       post.UserID,
		PostDatetime: post.PostDatetime,
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		ParentPost:   post.ParentPost,

		Tags: post.Tags,

		User: post.User,

		Comments: post.Comments,

		PostPreferences: post.PostPreferences,
	}
}
