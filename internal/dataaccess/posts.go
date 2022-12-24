package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func ListPosts(db *database.Database) ([]models.Post, error) {
	posts := []models.Post{
		{
			Post_ID:       5,
			User_ID:       1,
			Post_Datetime: "12 December 16:40",
			Post_Title:    "This is my fifth forum post",
			Post_Content:  "This is life 5 by one",
			Parent_Post:   -1,

			Tags: []string{"tag 1", "tag 3"},

			Username: "Bobby Lee One",

			Comment_Count: 23,

			Likes:               29,
			Is_Like_Selected:    true,
			Is_Dislike_Selected: false,
		},
		{
			Post_ID:       4,
			User_ID:       3,
			Post_Datetime: "11 December 15:40",
			Post_Title:    "This is my fourth forum post",
			Post_Content:  "This is life 4 by three",
			Parent_Post:   -1,

			Tags: []string{"tag 2", "tag 4"},

			Username: "Bobby Lee Three",

			Comment_Count: 63,

			Likes:               30,
			Is_Like_Selected:    false,
			Is_Dislike_Selected: false,
		},
		{
			Post_ID:       3,
			User_ID:       2,
			Post_Datetime: "10 December 13:40",
			Post_Title:    "This is my Third forum post",
			Post_Content:  "This is life 3 by two",
			Parent_Post:   -1,

			Tags: []string{"tag 1", "tag 2", "tag 3", "tag 4", "tag 5"},

			Username: "Bobby Lee Two",

			Comment_Count: 14,

			Likes:               15,
			Is_Like_Selected:    false,
			Is_Dislike_Selected: true,
		},
		{
			Post_ID:       2,
			User_ID:       1,
			Post_Datetime: "4 December 12:40",
			Post_Title:    "This is my second forum post",
			Post_Content:  "This is life 2 by one",
			Parent_Post:   -1,

			Tags: []string{"tag 2", "tag 1", "tag 3"},

			Username: "Bobby Lee One",

			Comment_Count: 7,

			Likes:               44,
			Is_Like_Selected:    false,
			Is_Dislike_Selected: false,
		},
		{
			Post_ID:       1,
			User_ID:       5,
			Post_Datetime: "1 December 12:40",
			Post_Title:    "This is my first forum post",
			Post_Content:  "This is life 1 by five",
			Parent_Post:   -1,

			Tags: []string{"tag 4", "tag 2"},

			Username: "Bobby Lee Five",

			Comment_Count: 16,

			Likes:               30,
			Is_Like_Selected:    false,
			Is_Dislike_Selected: false,
		},
	}
	return posts, nil
}
