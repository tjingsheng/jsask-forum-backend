package retriever

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	users "github.com/tjingsheng/jsask-forum-backend/internal/dataaccess"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
)

const (
	SuccessfulListTargetMessage = "successfully listed %s"
	ErrRetrieveDatabase         = "failed to retrieve database > failed to retrieve %s in %s.HandleList"
	ErrRetrieveTarget           = "failed to retrieve %s in %s.HandleList"
	ErrEncodeView               = "failed to retrieve %s in %s.HandleList"
)

func HandleList(w http.ResponseWriter, r *http.Request, target string) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, fmt.Errorf(ErrRetrieveDatabase, target, target)
	}

	users, err := users.List(db)
	if err != nil {
		return nil, fmt.Errorf(ErrRetrieveTarget, target, target)
	}

	data, err := json.Marshal(users)
	if err != nil {
		return nil, fmt.Errorf(ErrEncodeView, target, target)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulListTargetMessage, target)},
	}, nil
}
