package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	postview "github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	posts, err := dataaccess.ListPosts(database.DB)

	allPostsView := make([]postview.ListView, len(posts))

	for i := range posts {
		allPostsView[i] = postview.ListFrom(posts[i])
	}

	return format(err, allPostsView, ListTags)
}
