package models

type Tag struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	TagName string `json:"tagName"`
}
