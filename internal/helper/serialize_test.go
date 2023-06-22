package helper

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

func TestPreperamentQueue(t *testing.T) {
	input := "test"
	param := "param"

	expectedQueueRequest := entity.QueueRequest{
		Param: param,
		Data:  input,
	}

	expectedMessage, err := json.Marshal(expectedQueueRequest)
	if err != nil {
		t.Fatalf("Failed to marshal expected message: %v", err)
	}

	message, err := PreperamentQueue(input, param)
	if err != nil {
		t.Fatalf("Failed to prepare queue: %v", err)
	}

	if !reflect.DeepEqual(message, expectedMessage) {
		t.Errorf("Expected message: %s\nGot message: %s", expectedMessage, message)
	}
}
