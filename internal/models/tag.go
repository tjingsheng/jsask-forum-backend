package models

type Tag struct {
	Tag_ID   int    `json:"tagId"`
	Post_ID  int    `json:"postId"`
	Tag_Name string `json:"tagName"`
}
