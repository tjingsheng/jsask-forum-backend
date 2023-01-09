package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
)

func format(err error, x interface{}, handlerName string) (*api.Response, error) {
	if err != nil {
		return nil, fmt.Errorf(ErrRetrieveData, handlerName)
	}

	data, err := json.Marshal(x)

	if err != nil {
		return nil, fmt.Errorf(ErrEncodeView, handlerName)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulListMessage, handlerName)},
	}, nil
}
