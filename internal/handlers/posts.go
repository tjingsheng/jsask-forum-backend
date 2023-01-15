package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
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

func PostNewPost(w http.ResponseWriter, req *http.Request) {

	type PostNewPostRequest struct {
		UserID      int    `json:"userId"`
		PostTitle   string `json:"postTitle"`
		PostContent string `json:"postContent"`
		ParentPost  int    `json:"parentPost"`

		Tags []string `json:"tags"`
	}

	var request PostNewPostRequest
	json.NewDecoder(req.Body).Decode(&request)
	fmt.Println(request)

	newPost := models.Post{
		UserID:       request.UserID,
		PostDatetime: time.Now().Format(constants.Go2PostgresqlDatetime),
		PostTitle:    request.PostTitle,
		PostContent:  request.PostContent,
		ParentPost:   request.ParentPost,
	}

	err := dataaccess.CreatePost(database.DB, newPost)
	response, _ := utils.HandlerFormatPost(err, newPost, "PostNewPost")
	json.NewEncoder(w).Encode(response)
}
