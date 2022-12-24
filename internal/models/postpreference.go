package models

type PostPreference struct {
	Post_Preference_ID  int    `json:"postPreferenceId"`
	User_ID             int    `json:"userId"`
	Post_ID             int    `json:"postId"`
	Preference_Datetime string `json:"preferenceDatetime"`
	Preference          string `json:"preference"`
}
