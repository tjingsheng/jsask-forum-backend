package postsview

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
)

type ListView struct {
	ID           int    `json:"postID"`
	UserID       int    `json:"userID"`
	PostDatetime string `json:"postDatetime"`
	PostTitle    string `json:"postTitle"`
	PostContent  string `json:"postContent"`
	ParentPost   int    `json:"parentPost"`

	Tags []string `json:"tags"`

	Username string `json:"username"`

	CommentCount int `json:"commentcount"`

	Likes             int  `json:"likes"`
	IsLikeSelected    bool `json:"isLikeSelected"`
	IsDislikeSelected bool `json:"isDislikeSelected"`
}

func ListFrom(post postsviewmodel.ListView) ListView {
	postTags := make([]string, len(post.Tags))
	for i := range post.Tags {
		postTags[i] = post.Tags[i].TagName
	}
	postUsername := post.User.Username
	postCommentCount := len(post.Comments)
	postLikes := 20
	postIsLikeSelected := true
	postIsDislikeSelected := true

	return ListView{
		ID:           post.ID,
		UserID:       post.UserID,
		PostDatetime: post.PostDatetime,
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		ParentPost:   post.ParentPost,

		Tags: postTags,

		Username: postUsername,

		CommentCount: postCommentCount,

		Likes:             postLikes,
		IsLikeSelected:    postIsLikeSelected,
		IsDislikeSelected: postIsDislikeSelected,
	}
}
