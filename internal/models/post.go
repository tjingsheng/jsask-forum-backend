package models

type Post struct {
	Post_ID       int    `json:"postID"`
	User_ID       int    `json:"userID"`
	Post_Datetime string `json:"postDatetime"`
	Post_Title    string `json:"postTitle"`
	Post_Content  string `json:"postContent"`
	Parent_Post   int    `json:""`

	Tags []string `json:"tags"`

	Username string `json:"username"`

	Comment_Count int `json:"commentCount"`

	Likes               int  `json:"likes"`
	Is_Like_Selected    bool `json:"isLikeSelected"`
	Is_Dislike_Selected bool `json:"isDislikeSelected"`
}
