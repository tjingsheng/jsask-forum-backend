package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	// "github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	posts, err := dataaccess.ListPosts(database.DB)
	return format(err, posts, ListPosts)
}
