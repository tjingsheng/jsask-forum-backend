package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func ListTags(db *database.Database) ([]models.Tag, error) {
	posts := []models.Tag{
		{
			Tag_ID:   1,
			Post_ID:  1,
			Tag_Name: "tag 4",
		},
		{
			Tag_ID:   2,
			Post_ID:  1,
			Tag_Name: "tag 2",
		},
		{
			Tag_ID:   3,
			Post_ID:  2,
			Tag_Name: "tag 2",
		},
		{
			Tag_ID:   4,
			Post_ID:  2,
			Tag_Name: "tag 1",
		},
		{
			Tag_ID:   5,
			Post_ID:  2,
			Tag_Name: "tag 3",
		},
		{
			Tag_ID:   6,
			Post_ID:  3,
			Tag_Name: "tag 1",
		},
		{
			Tag_ID:   7,
			Post_ID:  3,
			Tag_Name: "tag 2",
		},
		{
			Tag_ID:   8,
			Post_ID:  3,
			Tag_Name: "tag 3",
		},
		{
			Tag_ID:   9,
			Post_ID:  3,
			Tag_Name: "tag 4",
		},
		{
			Tag_ID:   10,
			Post_ID:  3,
			Tag_Name: "tag 5",
		},
		{
			Tag_ID:   11,
			Post_ID:  4,
			Tag_Name: "tag 2",
		},
		{
			Tag_ID:   12,
			Post_ID:  4,
			Tag_Name: "tag 4",
		},
		{
			Tag_ID:   13,
			Post_ID:  5,
			Tag_Name: "tag 1",
		},
		{
			Tag_ID:   14,
			Post_ID:  5,
			Tag_Name: "tag 3",
		},
	}
	return posts, nil
}
