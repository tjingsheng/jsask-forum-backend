package models

type PostsTags struct {
	ID     int `json:"id" gorm:"primaryKey"`
	PostID int `json:"postId"`
	TagID  int `json:"tagId"`
}
