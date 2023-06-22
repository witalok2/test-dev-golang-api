package rabbitmq

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRabbitMQ struct {
	mock.Mock
}

func (ref *MockRabbitMQ) PublishMessage(ctx context.Context, message []byte) error {
	args := ref.Called(ctx, message)
	return args.Error(0)
}
