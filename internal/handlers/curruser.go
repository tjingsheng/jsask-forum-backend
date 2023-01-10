package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/curruserviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/curruserview"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request, username string) (*api.Response, error) {
	currUser, err := dataaccess.ListCurentUser(database.DB, username)
	currUserViewModel := curruserviewmodel.ListFrom(currUser)
	currUserView := curruserview.ListFrom(currUserViewModel)
	return utils.HandlerFormat(err, currUserView, "GetCurrentUser")
}
