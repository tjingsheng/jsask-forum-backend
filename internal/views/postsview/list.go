package postsview

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	ID            int    `json:"postID"`
	User_ID       int    `json:"userID"`
	Post_Datetime string `json:"postDatetime"`
	Post_Title    string `json:"postTitle"`
	Post_Content  string `json:"postContent"`
	Parent_Post   int    `json:"parentPost"`

	Tags []models.Tag `json:"tags"`

	Username string `json:"username"`

	Comment_Count int `json:"commentCount"`

	Likes               int  `json:"likes"`
	Is_Like_Selected    bool `json:"isLikeSelected"`
	Is_Dislike_Selected bool `json:"isDislikeSelected"`
}

func ListFrom(post models.PostZ) ListView {
	return ListView{
		ID:            post.ID,
		User_ID:       post.User_ID,
		Post_Datetime: post.Post_Datetime,
		Post_Title:    post.Post_Title,
		Post_Content:  post.Post_Content,
		Parent_Post:   post.Parent_Post,

		Tags: post.Tags,

		Username: post.Username,

		Comment_Count: post.Comment_Count,

		Likes:               post.Likes,
		Is_Like_Selected:    post.Is_Like_Selected,
		Is_Dislike_Selected: post.Is_Dislike_Selected,
	}
}
