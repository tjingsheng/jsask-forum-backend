package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/currpostsview"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetCurrPost(w http.ResponseWriter, r *http.Request, userId string, postId string) (*api.Response, error) {
	currPosts, err := dataaccess.ListCurrPost(database.DB, postId)
	allCurrPostsViewModel := make([]postsviewmodel.ListView, len(currPosts))
	allCurrPostsView := make([]postsview.ListView, len(currPosts))
	for i := range currPosts {
		allCurrPostsViewModel[i] = postsviewmodel.ListFrom(currPosts[i])
		allCurrPostsView[i] = postsview.ListFrom(allCurrPostsViewModel[i], userId)
	}

	currPostView := currpostsview.ListFrom(allCurrPostsView)

	return utils.HandlerFormat(err, currPostView, "GetAllPosts")
}
