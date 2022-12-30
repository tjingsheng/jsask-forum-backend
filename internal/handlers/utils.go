package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
)

func format(err error, x interface{}, handlerName string) (*api.Response, error) {
	if err != nil {
		return nil, fmt.Errorf(ErrRetrieveData, ListTags)
	}

	data, err := json.Marshal(x)

	if err != nil {
		return nil, fmt.Errorf(ErrEncodeView, ListTags)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListMessage},
	}, nil
}
