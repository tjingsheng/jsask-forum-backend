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
		},
		{
			Post_ID:       4,
			User_ID:       3,
			Post_Datetime: "11 December 15:40",
			Post_Title:    "This is my fourth forum post",
			Post_Content:  "This is life 4 by three",
			Parent_Post:   -1,
		},
		{
			Post_ID:       2,
			User_ID:       1,
			Post_Datetime: "4 December 12:40",
			Post_Title:    "This is my second forum post",
			Post_Content:  "This is life 2 by one",
			Parent_Post:   -1,
		},
		{
			Post_ID:       1,
			User_ID:       5,
			Post_Datetime: "1 December 12:40",
			Post_Title:    "This is my first forum post",
			Post_Content:  "This is life 1 by five",
			Parent_Post:   -1,
		},
	}
	return posts, nil
}
