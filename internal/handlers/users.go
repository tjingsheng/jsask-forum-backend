package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/userview"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request, username string) (*api.Response, error) {
	currUser, err := dataaccess.ListCurentUser(database.DB, username)
	currUserView := userview.ListFrom(currUser)
	return format(err, currUserView, ListCurrentUser)
}
