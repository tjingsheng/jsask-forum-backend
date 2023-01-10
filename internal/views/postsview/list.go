package postsview

import (
	"strconv"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
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

func isLiked(postPreference models.PostPreference) bool {
	return postPreference.Preference == constants.IsLiked
}

func isDisliked(postPreference models.PostPreference) bool {
	return postPreference.Preference == constants.IsDisliked
}

func isCurrUser(userId string) func(models.PostPreference) bool {
	return func(postPreference models.PostPreference) bool {
		targetUserId, err := strconv.Atoi(userId)
		if err != nil {
			return false
		}
		return targetUserId == postPreference.UserID
	}
}

func ListFrom(post postsviewmodel.ListView, userId string) ListView {
	postTags := make([]string, len(post.Tags))
	for i := range post.Tags {
		postTags[i] = post.Tags[i].TagName
	}
	postUsername := post.User.Username
	postCommentCount := len(post.Comments)
	postLikes := len(utils.FilterStruct(post.PostPreferences, isLiked)) - len(utils.FilterStruct(post.PostPreferences, isDisliked))
	postIsLikeSelected := len(utils.FilterStruct(utils.FilterStruct(post.PostPreferences, isCurrUser(userId)), isLiked)) != 0
	postIsDislikeSelected := len(utils.FilterStruct(utils.FilterStruct(post.PostPreferences, isCurrUser(userId)), isDisliked)) != 0

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
