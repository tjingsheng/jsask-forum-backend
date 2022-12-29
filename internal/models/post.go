package models

type Post struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	User_ID       int    `json:"user_id"`
	Post_Datetime string `json:"post_datetime"`
	Post_Title    string `json:"post_title"`
	Post_Content  string `json:"post_content"`
	Parent_Post   int    `json:"parent_post"`

	// Tags []string `json:"tags"`

	// Username string `json:"username"`

	// Comment_Count int `json:"commentCount"`

	// Likes               int  `json:"likes"`
	// Is_Like_Selected    bool `json:"isLikeSelected"`
	// Is_Dislike_Selected bool `json:"isDislikeSelected"`
}
