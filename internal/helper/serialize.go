package helper

import (
	"encoding/json"
	"errors"

	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

func PreperamentQueue(input interface{}, param string) ([]byte, error) {
	queueRequest := entity.QueueRequest{
		Param: param,
		Data:  input,
	}

	message, err := json.Marshal(queueRequest)
	if err != nil {
		return nil, errors.New("failed to serialize Queue")
	}

	return message, nil
}
