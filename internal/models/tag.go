package models

type Tag struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Tag_Name string `json:"tagName"`
}
