package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
	"github.com/witalok2/test-dev-golang-api/internal/helper"
)

func (s service) ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error) {
	return s.db.ListClient(ctx, page, limit)
}

func (s service) GetClient(ctx context.Context, clientID uuid.UUID) (client entity.Client, err error) {
	client, err = s.db.GetClient(ctx, clientID)
	if err != nil {
		return entity.Client{}, err
	}

	return client, nil
}

func (s service) CreateClient(ctx context.Context, client *entity.Client) error {
	message, err := helper.PreperamentQueue(client, entity.CREATE_CLIENT)
	if err != nil {
		return err
	}

	err = s.rabbitMQ.PublishMessage(ctx, message)
	if err != nil {
		return errors.New("erro ao publicar evento")
	}

	return nil
}

func (s service) UpdateClient(ctx context.Context, client *entity.Client) error {
	message, err := helper.PreperamentQueue(client, entity.UPDATE_CLIENT)
	if err != nil {
		return err
	}

	err = s.rabbitMQ.PublishMessage(ctx, message)
	if err != nil {
		return errors.New("erro ao publicar evento")
	}

	return nil
}

func (s service) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	return s.db.DeleteClient(ctx, clientID)
}
