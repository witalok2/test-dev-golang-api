package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/witalok2/test-dev-golang-api/adapter/secondary/database"
	amqp "github.com/witalok2/test-dev-golang-api/adapter/secondary/rabbitMQ"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

type ServiceInterface interface {
	ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error)
	GetClient(ctx context.Context, clientID uuid.UUID) (clients entity.Client, err error)
	CreateClient(ctx context.Context, client *entity.Client) error
	UpdateClient(ctx context.Context, client *entity.Client) error
	DeleteClient(ctx context.Context, clientID uuid.UUID) error
}

type service struct {
	db       database.Repository
	rabbitMQ amqp.QueueClient
}

func NewService(ctx context.Context, db database.Repository, rabbitMQ amqp.QueueClient) *service {
	return &service{
		db:       db,
		rabbitMQ: rabbitMQ,
	}
}
