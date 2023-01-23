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
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/saltviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/curruserview"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/saltview"
)

func GetCurrUser(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	password := chi.URLParam(req, "password")
	currUser, err := dataaccess.ListCurrUser(database.DB, utils.UsernameFormatter(username), password)

	currUserViewModel := curruserviewmodel.ListFrom(currUser)
	currUserView := curruserview.ListFrom(currUserViewModel)

	response, _ := utils.HandlerFormatter(err, currUserView, "GetCurrUser", constants.SuccessfulGetMessage)
	json.NewEncoder(w).Encode(response)
}

func GetSalt(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	currSalt, err := dataaccess.ListSalt(database.DB, utils.UsernameFormatter(username))

	saltViewModel := saltviewmodel.ListFrom(currSalt)
	saltView := saltview.ListFrom(saltViewModel)

	response, _ := utils.HandlerFormatter(err, saltView, "GetCurrUser", constants.SuccessfulGetMessage)
	json.NewEncoder(w).Encode(response)
}

func PostCurrUser(w http.ResponseWriter, req *http.Request) {
	type PostCurrUserRequest struct {
		Username string `json:"username"`
		Salt     string `json:"salt"`
		Password string `json:"password"`
	}

	var request PostCurrUserRequest
	json.NewDecoder(req.Body).Decode(&request)

	newUser := models.User{
		Username:     utils.UsernameFormatter(request.Username),
		UserDatetime: time.Now().Format(constants.Go2PostgresqlDatetime),
		Salt:         request.Salt,
		Password:     request.Password,
	}

	err := dataaccess.CreateUser(database.DB, newUser)

	response, _ := utils.HandlerFormatter(err, newUser, "PostCurrUser", constants.SuccessfulPostMessage)
	json.NewEncoder(w).Encode(response)
}
