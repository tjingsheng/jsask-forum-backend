package models

type Post_Tag struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Post_ID string `json:"postId"`
	Tag_ID  string `json:"tagId"`
}
