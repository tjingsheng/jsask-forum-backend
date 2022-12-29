package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/tjingsheng/jsask-forum-backend/internal/api"
// 	"github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
// 	"github.com/tjingsheng/jsask-forum-backend/internal/database"
// )

// func (h handler) HandleListPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
// 	db, err := database.GetDB()

// 	if err != nil {
// 		return nil, fmt.Errorf(ErrRetrieveDatabase, ListPosts)
// 	}

// 	users, err := dataaccess.ListPosts(db)

// 	if err != nil {
// 		return nil, fmt.Errorf(ErrRetrieveUsers, ListPosts)
// 	}

// 	data, err := json.Marshal(users)
// 	if err != nil {
// 		return nil, fmt.Errorf(ErrEncodeView, ListPosts)
// 	}

// 	return &api.Response{
// 		Payload: api.Payload{
// 			Data: data,
// 		},
// 		Messages: []string{SuccessfulListMessage},
// 	}, nil
// }
