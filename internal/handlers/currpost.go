package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/currpostviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/currpostview"
)

func GetCurrPost(w http.ResponseWriter, r *http.Request, userId string, postId string) (*api.Response, error) {
	currPost, err := dataaccess.ListCurrPost(database.DB, postId)
	currpostviewmodel := currpostviewmodel.ListFrom(currPost)
	currpostview := currpostview.ListFrom(currpostviewmodel, userId)
	return utils.HandlerFormat(err, currpostview, "GetAllPosts")
}
