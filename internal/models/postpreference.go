package models

type PostPreference struct {
	ID                 int    `json:"id" gorm:"primaryKey"`
	UserID             int    `json:"userId"`
	PostID             string `json:"postId"`
	PreferenceDatetime string `json:"preferenceDatetime"`
	Preference         string `json:"preference"`
}
