package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/tagsview"
)

func GetAllUniqueTags(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	tags, err := dataaccess.ListTags(database.DB)

	allTagsView := make([]tagsview.ListView, len(tags))
	for i := range tags {
		allTagsView[i] = tagsview.ListFrom(tags[i])
	}

	keys := make(map[tagsview.ListView]bool)
	uniqueTags := []tagsview.ListView{}
	for _, i := range allTagsView {
		if _, obj := keys[i]; !obj {
			keys[i] = true
			uniqueTags = append(uniqueTags, i)
		}
	}

	return format(err, uniqueTags, ListTags)
}
