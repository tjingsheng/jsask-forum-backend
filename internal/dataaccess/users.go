package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
)

func ListUsers(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			User_ID:       1,
			Username:      "Bobby Lee One",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
		{
			User_ID:       2,
			Username:      "Bobby Lee Two",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
		{
			User_ID:       3,
			Username:      "Bobby Lee Three",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
		{
			User_ID:       4,
			Username:      "Harry Potter One",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
		{
			User_ID:       5,
			Username:      "Bobby Lee Five",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
		{
			User_ID:       6,
			Username:      "Harry Potter Two",
			User_Datetime: "2014-03-12T13:37:27+00:00",
		},
	}
	return users, nil
}
