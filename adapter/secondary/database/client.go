package database

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

func (r *repository) ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error) {
	offset := (page - 1) * limit

	rows, err := r.db.QueryContext(ctx, SqlListClient, limit, offset)
	if err != nil {
		logger.WithError(err).Error("error listing clients with pagination")
		return nil, nil, err
	}
	defer rows.Close()

	clients := []entity.Client{}
	for rows.Next() {
		client := entity.Client{}
		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.LastName,
			&client.Contact,
			&client.Address,
			&client.Birthday,
			&client.CPF,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			logger.WithError(err).Error("error on scan to client struct")
			return nil, nil, err
		}

		clients = append(clients, client)
	}

	totalItems, err := r.getTotalClientCount(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &clients, &entity.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: int(math.Ceil(float64(totalItems) / float64(limit))),
		TotalItems: totalItems,
	}, nil
}

func (r *repository) GetClient(ctx context.Context, clientID uuid.UUID) (client entity.Client, err error) {
	err = r.db.QueryRow(SqlGetClientById, clientID).Scan(
		&client.ID,
		&client.Name,
		&client.LastName,
		&client.Contact,
		&client.Address,
		&client.Birthday,
		&client.CPF,
		&client.CreatedAt,
		&client.UpdatedAt,
		&client.DeletedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			logger.WithError(err).Error(sql.ErrNoRows)
			return entity.Client{}, errors.New("client not found")
		}

		logger.WithError(err).Error("error getting client")
		return entity.Client{}, err
	}

	return client, nil
}

func (r *repository) DeleteClient(ctx context.Context, clientId uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, SqlDeleteClient, clientId)
	if err != nil {
		logger.WithError(err).Error("error delete client")
		return err
	}

	return nil
}

func (r *repository) getTotalClientCount(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, SqlCountClient).Scan(&count)
	if err != nil {
		logger.WithError(err).Error("error count client")
		return 0, err
	}
	return count, nil
}
