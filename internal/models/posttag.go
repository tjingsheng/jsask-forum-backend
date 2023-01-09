package models

type PostsTag struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	PostID string `json:"postId"`
	TagID  string `json:"tagId"`
}
