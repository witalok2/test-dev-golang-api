package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).(*[]entity.Client), args.Get(1).(*entity.Pagination), args.Error(2)
}

func (m *MockDatabase) GetClient(ctx context.Context, clientID uuid.UUID) (entity.Client, error) {
	args := m.Called(ctx, clientID)
	return args.Get(0).(entity.Client), args.Error(1)
}

func (m *MockDatabase) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	args := m.Called(ctx, clientID)
	return args.Error(0)
}

func (m *MockDatabase) Close() {}
