package models

type User struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Username      string `json:"username"`
	User_Datetime string `json:"userDatetime"`
}
