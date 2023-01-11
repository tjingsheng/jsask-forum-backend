package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/postsviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/currpostsview"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetAllPosts(w http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")
	posts, err := dataaccess.ListPosts(database.DB)
	allPostsViewModel := make([]postsviewmodel.ListView, len(posts))
	allPostsView := make([]postsview.ListView, len(posts))
	for i := range posts {
		allPostsViewModel[i] = postsviewmodel.ListFrom(posts[i])
		allPostsView[i] = postsview.ListFrom(allPostsViewModel[i], userId)
	}
	response, _ := utils.HandlerFormatGet(err, allPostsView, "GetAllPosts")
	json.NewEncoder(w).Encode(response)
}

func GetCurrPost(w http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")
	postId := chi.URLParam(req, "postId")
	currPosts, err := dataaccess.ListCurrPost(database.DB, postId)
	allCurrPostsViewModel := make([]postsviewmodel.ListView, len(currPosts))
	allCurrPostsView := make([]postsview.ListView, len(currPosts))
	for i := range currPosts {
		allCurrPostsViewModel[i] = postsviewmodel.ListFrom(currPosts[i])
		allCurrPostsView[i] = postsview.ListFrom(allCurrPostsViewModel[i], userId)
	}
	currPostView := currpostsview.ListFrom(allCurrPostsView)
	response, _ := utils.HandlerFormatGet(err, currPostView, "GetAllPosts")
	json.NewEncoder(w).Encode(response)
}
