package models

type Post struct {
	Post_ID       int    `json:"postId"`
	User_ID       int    `json:"userId"`
	Post_Datetime string `json:"postDatetime"`
	Post_Title    string `json:"postTitle"`
	Post_Content  string `json:"postContent"`
	Parent_Post   int    `json:""`
}
