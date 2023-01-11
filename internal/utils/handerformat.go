package utils

import (
	"encoding/json"
	"fmt"

	"github.com/tjingsheng/jsask-forum-backend/internal/api"
	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
)

func HandlerFormatGet(err error, x interface{}, handlerName string) (*api.Response, error) {
	if err != nil {
		return nil, fmt.Errorf(constants.ErrRetrieveData, handlerName)
	}

	data, err := json.Marshal(x)

	if err != nil {
		return nil, fmt.Errorf(constants.ErrEncodeView, handlerName)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(constants.SuccessfulListMessage, handlerName)},
	}, nil
}

func HandlerFormatPost(err error, x interface{}, handlerName string) (*api.Response, error) {
	if err != nil {
		return nil, fmt.Errorf(constants.ErrRetrieveData, handlerName)
	}

	data, err := json.Marshal(x)

	if err != nil {
		return nil, fmt.Errorf(constants.ErrEncodeView, handlerName)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(constants.SuccessfulPostedMessage, handlerName)},
	}, nil
}
