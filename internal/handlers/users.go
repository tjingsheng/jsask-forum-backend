package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/curruserviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/curruserview"
	"gorm.io/gorm"
)

func GetCurrUser(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	currUser, err := dataaccess.ListCurrUser(database.DB, username)

	if err == gorm.ErrRecordNotFound {
		newUser := models.User{
			Username:     username,
			UserDatetime: time.Now().Format(constants.Go2PostgresqlDatetime),
		}
		dataaccess.CreateUser(database.DB, newUser)
		currUser, err = dataaccess.ListCurrUser(database.DB, username)
	}

	currUserViewModel := curruserviewmodel.ListFrom(currUser)
	currUserView := curruserview.ListFrom(currUserViewModel)
	response, _ := utils.HandlerFormatGet(err, currUserView, "GetCurrUser")
	json.NewEncoder(w).Encode(response)
}

// func PostNewUser(w http.ResponseWriter, req *http.Request) {
// 	username := chi.URLParam(req, "username")
// 	newUser := models.User{
// 		Username:     username,
// 		UserDatetime: time.Now().Format(constants.Go2PostgresqlDatetime),
// 	}
// 	err := dataaccess.CreateUser(database.DB, newUser)
// 	response, _ := utils.HandlerFormatPost(err, newUser, "PostNewUser")
// 	json.NewEncoder(w).Encode(response)
// }
