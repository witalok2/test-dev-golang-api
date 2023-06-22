package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

type MockService struct {
	mock.Mock
}

func (ref *MockService) ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error) {
	args := ref.Called(ctx, page, limit)
	return args.Get(0).(*[]entity.Client), args.Get(1).(*entity.Pagination), args.Error(2)
}

func (ref *MockService) GetClient(ctx context.Context, clientID uuid.UUID) (entity.Client, error) {
	args := ref.Called(ctx, clientID)
	return args.Get(0).(entity.Client), args.Error(1)
}

func (ref *MockService) CreateClient(ctx context.Context, client *entity.Client) error {
	args := ref.Called(ctx, client)
	return args.Error(0)
}

func (ref *MockService) UpdateClient(ctx context.Context, client *entity.Client) error {
	args := ref.Called(ctx, client)
	return args.Error(0)
}

func (ref *MockService) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	args := ref.Called(ctx, clientID)
	return args.Error(0)
}
