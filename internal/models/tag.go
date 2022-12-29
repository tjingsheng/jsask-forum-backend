package models

type Tag struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Post_ID  string `json:"post_id"`
	Tag_Name string `json:"tag_name"`
}
