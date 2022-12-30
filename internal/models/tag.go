package models

type Tag struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Post_ID  string `json:"postId"`
	Tag_Name string `json:"tagName"`
}
