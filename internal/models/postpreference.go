package models

type PostPreference struct {
	ID                  int    `json:"id" gorm:"primaryKey"`
	User_ID             int    `json:"user_id"`
	Post_ID             string `json:"post_id"`
	Preference_Datetime string `json:"preference_datetime"`
	Preference          string `json:"preference"`
}
