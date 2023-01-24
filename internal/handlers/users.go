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
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/saltviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/userviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/saltview"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/userview"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	password := chi.URLParam(req, "password")
	user, err := dataaccess.ListCurrUser(database.DB, utils.UsernameFormatter(username), password)

	userViewModel := userviewmodel.ListFrom(user)
	userView := userview.ListFrom(userViewModel)

	response, _ := utils.HandlerFormatter(err, userView, "GetUser", constants.SuccessfulGetMessage)
	json.NewEncoder(w).Encode(response)
}

func GetUserSalt(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	currSalt, err := dataaccess.ListSalt(database.DB, utils.UsernameFormatter(username))

	saltViewModel := saltviewmodel.ListFrom(currSalt)
	saltView := saltview.ListFrom(saltViewModel)

	response, _ := utils.HandlerFormatter(err, saltView, "GetCurrUser", constants.SuccessfulGetMessage)
	json.NewEncoder(w).Encode(response)
}

func PostUser(w http.ResponseWriter, req *http.Request) {
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
