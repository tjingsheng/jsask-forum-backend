package users

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   0,
			Name: "Susan",
		},
		{
			ID:   1,
			Name: "Bobby",
		},
	}
	return users, nil
}
