package models

type PostPreference struct {
	ID                  int    `json:"id" gorm:"primaryKey"`
	User_ID             int    `json:"userId"`
	Post_ID             string `json:"postId"`
	Preference_Datetime string `json:"preferenceDatetime"`
	Preference          string `json:"preference"`
}
