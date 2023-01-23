package models

type User struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Username     string `json:"username"`
	UserDatetime string `json:"userDatetime"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
}
