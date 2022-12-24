package models

type User struct {
	User_ID       int    `json:"userId"`
	Username      string `json:"username"`
	User_Datetime string `json:"userDatetime"`
}
