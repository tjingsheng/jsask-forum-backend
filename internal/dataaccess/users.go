package users

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "Susan",
		},
		{
			ID:   2,
			Name: "Bobby",
		},
		{
			ID:   3,
			Name: "Harry",
		},
	}
	return users, nil
}
