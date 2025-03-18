package models

type Post struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	UserID       int    `json:"userId"`
	PostDatetime string `json:"postDatetime"`
	PostTitle    string `json:"postTitle"`
	PostContent  string `json:"postContent"`
	ParentPost   int    `json:"parentPost"`
}
