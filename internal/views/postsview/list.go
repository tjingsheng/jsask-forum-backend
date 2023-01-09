package postsview

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

	Tags []models.PostsTag `json:"tags" gorm:"foreignKey:PostID"`

	// Username string `json:"username"`

	// CommentCount int `json:"commentCount"`

	// Likes             int  `json:"likes"`
	// IsLikeSelected    bool `json:"isLikeSelected"`
	// IsDislikeSelected bool `json:"isDislikeSelected"`
}

func ListFrom(post ListView) ListView {
	return ListView{
		ID:           post.ID,
		UserID:       post.UserID,
		PostDatetime: post.PostDatetime,
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		ParentPost:   post.ParentPost,

		// Tags: post.Tags,

		// Username: post.Username,

		// Comment_Count: post.Comment_Count,

		// Likes:               post.Likes,
		// Is_Like_Selected:    post.Is_Like_Selected,
		// Is_Dislike_Selected: post.Is_Dislike_Selected,
	}
}
