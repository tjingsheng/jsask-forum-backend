package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
)

func PutPostPreference(w http.ResponseWriter, req *http.Request) {

	type PostPreferenceRequest struct {
		UserID     int `json:"currUserId"`
		PostID     int `json:"currPostId"`
		Preference int `json:"currPreference"`
	}

	var request PostPreferenceRequest
	json.NewDecoder(req.Body).Decode(&request)

	newPostPreference := models.PostPreference{
		UserID:             request.UserID,
		PostID:             request.PostID,
		PreferenceDatetime: time.Now().Format(constants.Go2PostgresqlDatetime),
		Preference:         request.Preference,
	}

	err := dataaccess.UpdatePostPreference(database.DB, newPostPreference)
	response, _ := utils.HandlerFormatPost(err, newPostPreference, "PutPostPreference")
	json.NewEncoder(w).Encode(response)
}
