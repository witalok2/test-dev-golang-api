package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

func TestListClient(t *testing.T) {
	repo := new(MockDatabase)

	clients := []entity.Client{
		{ID: uuid.New(), Name: "John Doe"},
		{ID: uuid.New(), Name: "Jane Smith"},
	}
	pagination := &entity.Pagination{
		Page:       1,
		Limit:      10,
		TotalPages: 1,
		TotalItems: len(clients),
	}
	repo.On("ListClient", mock.Anything, 1, 10).Return(&clients, pagination, nil)

	ctx := context.Background()

	result, resultPagination, err := repo.ListClient(ctx, 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, &clients, result)
	assert.Equal(t, pagination, resultPagination)

	// Verificar se o método mock foi chamado conforme esperado
	repo.AssertExpectations(t)
}

func TestGetClient(t *testing.T) {
	repo := new(MockDatabase)

	clientID := uuid.New()
	client := entity.Client{ID: clientID, Name: "John Doe"}
	repo.On("GetClient", mock.Anything, clientID).Return(client, nil)

	ctx := context.Background()

	result, err := repo.GetClient(ctx, clientID)

	assert.NoError(t, err)
	assert.Equal(t, client, result)

	repo.AssertExpectations(t)
}

func TestDeleteClient(t *testing.T) {
	repo := new(MockDatabase)

	clientID := uuid.New()
	repo.On("DeleteClient", mock.Anything, clientID).Return(nil)

	ctx := context.Background()

	err := repo.DeleteClient(ctx, clientID)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
