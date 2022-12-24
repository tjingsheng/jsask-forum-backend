package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func ListTags(db *database.Database) ([]models.Tag, error) {
	tags := []models.Tag{
		{
			Tag_Name: "tag 1",
		},
		{
			Tag_Name: "tag 2",
		},
		{
			Tag_Name: "tag 3",
		},
		{
			Tag_Name: "tag 4",
		},
		{
			Tag_Name: "tag 5",
		},
	}
	return tags, nil
}
