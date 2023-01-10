package handlers

import (
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/tagsviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/tagsview"
)

func GetAllTags(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	tags, err := dataaccess.ListTags(database.DB)
	allTagsViewModel := make([]tagsviewmodel.ListView, len(tags))
	allTagsView := make([]tagsview.ListView, len(tags))
	for i := range tags {
		allTagsViewModel[i] = tagsviewmodel.ListFrom(tags[i])
		allTagsView[i] = tagsview.ListFrom(allTagsViewModel[i])
	}
	return utils.HandlerFormat(err, allTagsView, "GetAllTags")
}
