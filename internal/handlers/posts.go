package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	// "github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	var post []models.Post
	// var tags models.Tag
	result := database.DB.Model(&post).Preload("tags").Find(&post, "parent_post = ?", 0)

	return format(result.Error, post, ListTags)

	// posts, err := dataaccess.ListPosts(database.DB)

	// allPostsView := make([]postsview.ListView, len(posts))

	// for i := range posts {
	// 	allPostsView[i] = postsview.ListFrom(posts[i])
	// }

	// return format(err, allPostsView, ListTags)
}
