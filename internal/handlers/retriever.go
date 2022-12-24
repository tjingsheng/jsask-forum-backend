package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"errors"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	users "github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
)

const (
	SuccessfulListTargetMessage = "Successfully listed %s"
	ErrRetrieveDatabase         = "Failed to retrieve database. Therefore fail to retrieve %s in %s.HandleList"
	ErrRetrieveTarget           = "Failed to retrieve %s in %s.HandleList"
	ErrEncodeView               = "Failed to retrieve %s in %s.HandleList"
)

func HandleList(w http.ResponseWriter, r *http.Request, target string) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrRetrieveDatabase, target, target))
	}

	users, err := users.List(db)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrRetrieveTarget, target, target))
	}

	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrEncodeView, target, target))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulListTargetMessage, target)},
	}, nil
}
