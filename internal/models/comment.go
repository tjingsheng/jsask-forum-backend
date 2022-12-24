package models

type Comment struct {
	Post     Post   `json:"post"`
	Comments []Post `json:"comments"`
}
