package currpostview

import (
	"strconv"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/currpostviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

type ListView struct {
	Post     postsview.ListView `json:"post"`
	Comments []models.Post      `json:"comments"`
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

func isNotCurrUser(userId string) func(models.PostPreference) bool {
	return func(postPreference models.PostPreference) bool {
		targetUserId, err := strconv.Atoi(userId)
		if err != nil {
			return false
		}
		return targetUserId != postPreference.UserID
	}
}

func ListFrom(post currpostviewmodel.ListView, userId string) ListView {
	postPreferencesCurrUser := utils.FilterStruct(post.PostPreferences, isCurrUser(userId))
	postPreferencesNoCurrUser := utils.FilterStruct(post.PostPreferences, isNotCurrUser(userId))
	postTags := make([]string, len(post.Tags))
	for i := range post.Tags {
		postTags[i] = post.Tags[i].TagName
	}
	postUsername := post.User.Username
	postCommentCount := len(post.Comments)
	postLikes := len(utils.FilterStruct(postPreferencesNoCurrUser, isLiked)) - len(utils.FilterStruct(postPreferencesNoCurrUser, isDisliked))
	postIsLikeSelected := len(utils.FilterStruct(postPreferencesCurrUser, isLiked)) != 0
	postIsDislikeSelected := len(utils.FilterStruct(postPreferencesCurrUser, isDisliked)) != 0

	currPost := postsview.ListView{
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

	return ListView{
		Post:     currPost,
		Comments: post.Comments,
	}
}
