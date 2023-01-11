package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/curruserviewmodel"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/curruserview"
)

func GetCurrUser(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	currUser, err := dataaccess.ListCurrUser(database.DB, username)
	currUserViewModel := curruserviewmodel.ListFrom(currUser)
	currUserView := curruserview.ListFrom(currUserViewModel)
	response, _ := utils.HandlerFormatGet(err, currUserView, "GetCurrUser")
	json.NewEncoder(w).Encode(response)
}

func PostNewUser(w http.ResponseWriter, req *http.Request) {
	username := chi.URLParam(req, "username")
	currUser, err := dataaccess.ListCurrUser(database.DB, username)
	currUserViewModel := curruserviewmodel.ListFrom(currUser)
	currUserView := curruserview.ListFrom(currUserViewModel)
	response, _ := utils.HandlerFormatGet(err, currUserView, "GetCurrUser")
	json.NewEncoder(w).Encode(response)
}
