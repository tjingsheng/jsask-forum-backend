package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request, userId string) (*api.Response, error) {
	posts, err := dataaccess.ListPosts(database.DB)
	allPostsViewModel := make([]postsviewmodel.ListView, len(posts))
	allPostsView := make([]postsview.ListView, len(posts))
	for i := range posts {
		allPostsViewModel[i] = postsviewmodel.ListFrom(posts[i])
		allPostsView[i] = postsview.ListFrom(allPostsViewModel[i], userId)
	}
	return utils.HandlerFormat(err, allPostsView, "GetAllPosts")
}
