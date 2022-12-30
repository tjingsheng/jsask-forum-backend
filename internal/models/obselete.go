package models

// type Comment struct {
// 	Post     Post   `json:"post"`
// 	Comments []Post `json:"comments"`
// }

type PostZ struct {
	ID            int    `json:"id"`
	User_ID       int    `json:"userID"`
	Post_Datetime string `json:"postDatetime"`
	Post_Title    string `json:"postTitle"`
	Post_Content  string `json:"postContent"`
	Parent_Post   int    `json:""`

	Tags []Tag `json:"tags" gorm:"many2many:post_tag"`

	Username string `json:"username"`

	Comment_Count int `json:"commentCount"`

	Likes               int  `json:"likes"`
	Is_Like_Selected    bool `json:"isLikeSelected"`
	Is_Dislike_Selected bool `json:"isDislikeSelected"`
}

// type User struct {
// 	User_ID       int    `json:"userID"`
// 	Username      string `json:"username"`
// 	User_Datetime string `json:"userDatetime"`
// }
