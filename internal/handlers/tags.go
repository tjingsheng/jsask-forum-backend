package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/tagsview"
)

func GetAllTags(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	tags, err := dataaccess.ListTags(database.DB)
	allTagsView := make([]tagsview.ListView, len(tags))
	for i := range tags {
		allTagsView[i] = tagsview.ListFrom(tags[i])
	}
	return format(err, allTagsView, ListTags)
}
